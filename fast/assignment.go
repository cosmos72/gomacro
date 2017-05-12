/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * declaration.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
)

// Assign compiles an *ast.AssignStmt into an assignment to one or more place
func (c *Comp) Assign(node *ast.AssignStmt) {
	c.Pos = node.Pos()

	lhs, rhs := node.Lhs, node.Rhs
	if node.Tok == token.DEFINE {
		c.DeclVarsShort(lhs, rhs)
		return
	}
	ln, rn := len(lhs), len(rhs)
	if ln > 1 && rn == 1 {
		c.Errorf("unimplemented: assignment of multiple places with a single multi-valued expression: %v", node)
	} else if ln != rn {
		c.Errorf("invalid assignment, cannot assign %d values to %d places: %v", node)
	}

	// the naive loop
	//   for i := range lhs { c.Assign1(lhs[i], node.Tok, rhs[i]) }
	// is bugged. It breaks, among others, the common Go idiom to swap two values: a,b = b,a
	//
	// More in general, Go guarantees that all assignments happen *as if*
	// the rhs values were copied to temporary locations before the assignments.
	// That's exactly what we must do.
	//
	// To be completely accurate, https://golang.org/ref/spec#Assignments states:
	//
	// "The assignment proceeds in two phases. First, the operands of index expressions
	// and pointer indirections (including implicit pointer indirections in selectors)
	// on the left and the expressions on the right are all evaluated in the usual order.
	// Second, the assignments are carried out in left-to-right order."
	//
	// so we must also cache the results of index expressions, pointer indirections, selectors on left side

	places := make([]*Place, ln)
	exprs := make([]*Expr, rn)
	needcache := false
	for i, li := range lhs {
		places[i] = c.Place(li)
		needcache = needcache || !places[i].IsVar() // ach, needed. see the example i := 0; i, x[i] = 1, 2  // set i = 1, x[0] = 2
	}
	for i, ri := range rhs {
		exprs[i] = c.Expr1(ri)
		needcache = needcache || !exprs[i].Const()
	}
	if ln <= 1 {
		// optimization
		needcache = false
	}
	if needcache {
		var inits []func(*Env)
		for i, li := range lhs {
			for {
				switch node := li.(type) {
				case *ast.ParenExpr:
					li = node.X
					continue
				case *ast.IndexExpr, *ast.StarExpr, *ast.SelectorExpr:
					fun, _, index := places[i].Cache(CacheCopy)
					if fun != nil {
						inits = append(inits, fun)
					}
					if index != nil {
						inits = append(inits, index)
					}
				}
				break
			}
		}
		for _, expr := range exprs {
			init := expr.Cache(CacheCopy)
			// init is nil for constant expressions: nothing to cache
			if init != nil {
				inits = append(inits, init)
			}
		}
		if len(inits) != 0 {
			c.Code.Append(func(env *Env) (Stmt, *Env) {
				for _, init := range inits {
					init(env)
				}
				env.IP++
				return env.Code[env.IP], env
			})
		}
	}
	for i := range lhs {
		c.assign1(lhs[i], node.Tok, rhs[i], places[i], exprs[i])
	}
}

// assign1 compiles a single assignment to a place
func (c *Comp) assign1(lhs ast.Expr, op token.Token, rhs ast.Expr, place *Place, init *Expr) {
	panicking := true
	defer func() {
		if !panicking {
			return
		}
		rec := recover()
		node := &ast.AssignStmt{Lhs: []ast.Expr{lhs}, Tok: op, Rhs: []ast.Expr{rhs}} // for nice error messages
		c.Errorf("error compiling assignment: %v\n    %v", node, rec)
	}()
	if place.IsVar() {
		c.SetVar(&place.Var, op, init)
	} else {
		c.SetPlace(place, op, init)
	}
	panicking = false
}

// LookupVar compiles the left-hand-side of an assignment, in case it's an identifier (i.e. a variable name)
func (c *Comp) LookupVar(name string) *Var {
	if name == "_" {
		return &Var{}
	}
	sym := c.Resolve(name)
	return sym.AsVar(PlaceSettable)
}

// Place compiles the left-hand-side of an assignment
func (c *Comp) Place(node ast.Expr) *Place {
	return c.placeOrAddress(node, false)
}

// PlaceOrAddress compiles the left-hand-side of an assignment or the location of an address-of
func (c *Comp) placeOrAddress(in ast.Expr, opt PlaceOption) *Place {
	for {
		if in != nil {
			c.Pos = in.Pos()
		}
		switch node := in.(type) {
		case *ast.ParenExpr:
			in = node.X
			continue
		case *ast.Ident:
			return c.IdentPlace(node.Name, opt)
		case *ast.IndexExpr:
			return c.IndexPlace(node, opt)
		case *ast.StarExpr:
			e := c.Expr1(node.X)
			if e.Const() {
				c.Errorf("%s a constant: %v <%v>", opt, node, e.Type)
				return nil
			} else if e.Sym != nil && opt == PlaceAddress {
				// optimize &*variable -> variable: it's already the address we want,
				// remember to dereference its type
				//
				// we cannot do this optimization when opt == PlaceSettable,
				// because in such case the code to compile is *variable - not an identifier
				va := *e.Sym.AsVar(opt)
				va.Type = va.Type.Elem()
				return &Place{Var: va}
			} else {
				// e.Fun is already the address we want,
				// remember to dereference its type
				t := e.Type.Elem()
				addr := e.AsX1()
				fun := func(env *Env) r.Value {
					return addr(env).Elem()
				}
				return &Place{Var: Var{Type: t}, Fun: fun, Addr: addr}
			}
		case *ast.SelectorExpr:
			return c.SelectorPlace(node, opt)
		default:
			c.Errorf("%s: %v", opt, in)
			return nil
		}
	}
}

// placeForSideEffects compiles the left-hand-side of a do-nothing assignment,
// as for example *addressOfInt() += 0, in order to apply its side effects
func (c *Comp) placeForSideEffects(place *Place) {
	if place.IsVar() {
		return
	}
	var ret Stmt
	fun := place.Fun
	if mapkey := place.MapKey; mapkey != nil {
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			mapkey(env)
			// no need to call obj.MapIndex(key): it has no side effects and cannot panic.
			// obj := fun(env)
			// key := mapkey(env)
			// obj.MapIndex(key)
			env.IP++
			return env.Code[env.IP], env
		}
	} else {
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	}
	c.Code.Append(ret)
}

type CacheOption bool

const (
	CacheNoCopy CacheOption = false
	CacheCopy   CacheOption = true
)

// Expr.Cache is a wrapper around CacheFun that operates on expressions. it modifies 'e' !
func (e *Expr) Cache(option CacheOption) (setfun func(*Env)) {
	if e.Const() {
		return nil
	}
	setfun, e.Fun = CacheFun(e.Fun, e.Type, option)
	return
}

// Place.Cache is a wrapper around CacheFun that operates on places. it modifies 'p' !
func (p *Place) Cache(option CacheOption) (setfun func(*Env), addrfun func(*Env), mapkeyfun func(*Env)) {
	if p.IsVar() {
		return nil, nil, nil
	}
	t := p.Type
	if p.Addr != nil {
		addrfun, p.Addr = cacheFunX1(p.Addr, r.PtrTo(t), option)
	}
	if p.MapKey != nil {
		tmap := p.MapType
		setfun, p.Fun = cacheFunX1(p.Fun, tmap, option)
		mapkeyfun, p.MapKey = cacheFunX1(p.MapKey, tmap.Key(), option)
	} else {
		setfun, p.Fun = cacheFunX1(p.Fun, t, CacheNoCopy) // cannot copy, it would become unsettable
	}
	return
}

// CacheFun creates and returns two functions:
// the function 'setfun' will evaluate the given function 'fun' and store its result in a cache,
// the function 'getfun' will return the cached value.
// If 'option' == CacheCopy, setfun will also make a copy of the value before storing it in the cache.
// Used to create temporary storage for multiple assignments, as for example a,b = b,a
func CacheFun(fun I, t r.Type, option CacheOption) (setfun func(*Env), getfun I) {
	switch t.Kind() {
	case r.Bool:
		fun := fun.(func(*Env) bool)
		var cache bool
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) bool {
			return cache
		}
	case r.Int:
		fun := fun.(func(*Env) int)
		var cache int
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) int {
			return cache
		}
	case r.Int8:
		fun := fun.(func(*Env) int8)
		var cache int8
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) int8 {
			return cache
		}
	case r.Int16:
		fun := fun.(func(*Env) int16)
		var cache int16
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) int16 {
			return cache
		}
	case r.Int32:
		fun := fun.(func(*Env) int32)
		var cache int32
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) int32 {
			return cache
		}
	case r.Int64:
		fun := fun.(func(*Env) int64)
		var cache int64
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) int64 {
			return cache
		}
	case r.Uint:
		fun := fun.(func(*Env) uint)
		var cache uint
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) uint {
			return cache
		}
	case r.Uint8:
		fun := fun.(func(*Env) uint8)
		var cache uint8
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) uint8 {
			return cache
		}
	case r.Uint16:
		fun := fun.(func(*Env) uint16)
		var cache uint16
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) uint16 {
			return cache
		}
	case r.Uint32:
		fun := fun.(func(*Env) uint32)
		var cache uint32
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) uint32 {
			return cache
		}
	case r.Uint64:
		fun := fun.(func(*Env) uint64)
		var cache uint64
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) uint64 {
			return cache
		}
	case r.Uintptr:
		fun := fun.(func(*Env) uintptr)
		var cache uintptr
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) uintptr {
			return cache
		}
	case r.Float32:
		fun := fun.(func(*Env) float32)
		var cache float32
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) float32 {
			return cache
		}
	case r.Float64:
		fun := fun.(func(*Env) float64)
		var cache float64
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) float64 {
			return cache
		}
	case r.Complex64:
		fun := fun.(func(*Env) complex64)
		var cache complex64
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) complex64 {
			return cache
		}
	case r.Complex128:
		fun := fun.(func(*Env) complex128)
		var cache complex128
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) complex128 {
			return cache
		}
	case r.String:
		fun := fun.(func(*Env) string)
		var cache string
		setfun = func(env *Env) {
			cache = fun(env)
		}
		getfun = func(env *Env) string {
			return cache
		}
	default:
		if fun, ok := fun.(func(env *Env) (r.Value, []r.Value)); ok {
			setfun, getfun = cacheFunXV(fun, t, option)
		} else {
			funx1 := funAsX1(fun, t)
			setfun, getfun = cacheFunX1(funx1, t, option)
		}
	}
	return setfun, getfun
}

// special case of CacheFun
func cacheFunXV(fun func(env *Env) (r.Value, []r.Value), t r.Type, option CacheOption) (setfun func(*Env), getfun func(env *Env) r.Value) {
	var cache r.Value
	if option == CacheCopy {
		setfun = func(env *Env) {
			cache, _ = fun(env)
			if cache != base.Nil && cache.CanSet() {
				// make a copy. how? Convert() the r.Value to its expected type
				cache = cache.Convert(t)
			}
		}
	} else {
		setfun = func(env *Env) {
			cache, _ = fun(env)
		}
	}
	getfun = func(env *Env) r.Value {
		return cache
	}
	return
}

// special case of CacheFun
func cacheFunX1(fun func(env *Env) r.Value, t r.Type, option CacheOption) (setfun func(*Env), getfun func(env *Env) r.Value) {
	var cache r.Value
	if option == CacheCopy {
		setfun = func(env *Env) {
			cache = fun(env)
			if cache != base.Nil && cache.CanSet() {
				// make a copy. how? Convert() the r.Value to its expected type
				cache = cache.Convert(t)
			}
		}
	} else {
		setfun = func(env *Env) {
			cache = fun(env)
		}
	}
	getfun = func(env *Env) r.Value {
		return cache
	}
	return
}
