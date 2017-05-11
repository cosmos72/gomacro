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
	} else if ln == 1 {
		l := lhs[0]
		r := rhs[0]
		c.assign1(l, node.Tok, r, c.Place(l), c.Expr1(r))
		return
	}

	// the naive loop
	//   for i := range lhs { c.Assign1(lhs[i], node.Tok, rhs[i]) }
	// is bugged. It breaks, among others, the common Go idiom to swap two values: a,b = b,a
	//
	// More in general, Go guarantees that all assignments happen *as if*
	// the rhs values were copied to temporary locations before the assignments.
	// That's exactly what we must do.
	places := make([]*Place, ln)
	for i, li := range lhs {
		places[i] = c.Place(li)
	}
	inits := make([]func(*Env), 0, rn)
	caches := make([]*Expr, rn)
	for i, ri := range rhs {
		init, cache := c.Cache(c.Expr1(ri), CacheCopy)
		// init is nil for constant expressions: nothing to cache
		if init != nil {
			inits = append(inits, init)
		}
		caches[i] = cache
	}
	// TODO indexing, selectors, dereferencing and other side effects of places
	// should happen BEFORE evaluating inits!
	c.Code.Append(func(env *Env) (Stmt, *Env) {
		for _, init := range inits {
			init(env)
		}
		env.IP++
		return env.Code[env.IP], env
	})
	for i := range lhs {
		c.assign1(lhs[i], node.Tok, rhs[i], places[i], caches[i])
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
	CacheRef  CacheOption = false
	CacheCopy CacheOption = true
)

// Cache creates and returns a function and an expression:
// the function 'setfun' will evaluate the given expression 'expr' and store its result in a cache,
// the expression 'getter' will return the value stored in the cache.
// If 'option' == CacheCopy, setfun will also make a copy of the value before storing it in the cache.
// Used to create temporary storage for multiple assignments, as for example a,b = b,a
func (c *Comp) Cache(expr *Expr, option CacheOption) (setfun func(*Env), getter *Expr) {
	if expr.Const() {
		return nil, expr
	}
	efun := expr.Fun
	var getfun I
	switch expr.Type.Kind() {
	case r.Bool:
		efun := efun.(func(*Env) bool)
		var cache bool
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) bool {
			return cache
		}
	case r.Int:
		efun := efun.(func(*Env) int)
		var cache int
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) int {
			return cache
		}
	case r.Int8:
		efun := efun.(func(*Env) int8)
		var cache int8
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) int8 {
			return cache
		}
	case r.Int16:
		efun := efun.(func(*Env) int16)
		var cache int16
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) int16 {
			return cache
		}
	case r.Int32:
		efun := efun.(func(*Env) int32)
		var cache int32
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) int32 {
			return cache
		}
	case r.Int64:
		efun := efun.(func(*Env) int64)
		var cache int64
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) int64 {
			return cache
		}
	case r.Uint:
		efun := efun.(func(*Env) uint)
		var cache uint
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) uint {
			return cache
		}
	case r.Uint8:
		efun := efun.(func(*Env) uint8)
		var cache uint8
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) uint8 {
			return cache
		}
	case r.Uint16:
		efun := efun.(func(*Env) uint16)
		var cache uint16
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) uint16 {
			return cache
		}
	case r.Uint32:
		efun := efun.(func(*Env) uint32)
		var cache uint32
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) uint32 {
			return cache
		}
	case r.Uint64:
		efun := efun.(func(*Env) uint64)
		var cache uint64
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) uint64 {
			return cache
		}
	case r.Uintptr:
		efun := efun.(func(*Env) uintptr)
		var cache uintptr
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) uintptr {
			return cache
		}
	case r.Float32:
		efun := efun.(func(*Env) float32)
		var cache float32
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) float32 {
			return cache
		}
	case r.Float64:
		efun := efun.(func(*Env) float64)
		var cache float64
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) float64 {
			return cache
		}
	case r.Complex64:
		efun := efun.(func(*Env) complex64)
		var cache complex64
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) complex64 {
			return cache
		}
	case r.Complex128:
		efun := efun.(func(*Env) complex128)
		var cache complex128
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) complex128 {
			return cache
		}
	case r.String:
		efun := efun.(func(*Env) string)
		var cache string
		setfun = func(env *Env) {
			cache = efun(env)
		}
		getfun = func(env *Env) string {
			return cache
		}
	default:
		t := expr.Type
		var cache r.Value
		if efun, ok := efun.(func(env *Env) (r.Value, []r.Value)); ok {
			if option == CacheCopy {
				setfun = func(env *Env) {
					v, _ := efun(env)
					if v != base.Nil && v != base.None {
						// to make a copy, Convert() an r.Value to its own type
						cache = v.Convert(t)
					}
				}
			} else {
				setfun = func(env *Env) {
					cache, _ = efun(env)
				}
			}
		} else {
			efunx1 := expr.AsX1()
			if option == CacheCopy {
				t := expr.Type
				setfun = func(env *Env) {
					v := efunx1(env)
					if v != base.Nil && v != base.None {
						// to make a copy, Convert() an r.Value to its own type
						cache = v.Convert(t)
					}
				}
			} else {
				setfun = func(env *Env) {
					cache = efunx1(env)
				}
			}
		}
		getfun = func(env *Env) r.Value {
			return cache
		}
	}
	return setfun, exprFun(expr.Type, getfun)
}

/*
......
......
......
......
......
......
......
......
......
*/
