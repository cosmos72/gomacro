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

package compiler

import (
	"go/ast"
	"go/token"
	r "reflect"
)

// Decl compiles a constant, variable, function or type declaration - or an import
func (c *Comp) Decl(node ast.Decl) X {
	switch node := node.(type) {
	case *ast.GenDecl:
		return ExprStmtsToX(c.DeclGen(node))
	case *ast.FuncDecl:
		return c.DeclFunc(node, node.Type, node.Body)
	default:
		c.Errorf("Compile: unsupported declaration, expecting <*ast.GenDecl> or <*ast.FuncDecl>, found: %v <%v>", node, TypeOf(node))
		return nil
	}
}

// Decl compiles a constant, variable or type declaration - or an import
func (c *Comp) DeclGen(node *ast.GenDecl) []X {
	var funs []X
	switch node.Tok {
	case token.IMPORT:
		for _, decl := range node.Specs {
			c.Import(decl)
		}
	case token.CONST:
		for _, decl := range node.Specs {
			c.DeclConsts(decl)
		}
	case token.TYPE:
		for _, decl := range node.Specs {
			c.DeclType(decl)
		}
	case token.VAR:
		for _, decl := range node.Specs {
			funs = append(funs, c.DeclVars(decl)...)
		}
	default:
		c.Errorf("Compile: unsupported declaration kind, expecting token.IMPORT, token.CONST, token.TYPE or token.VAR, found %v: %v %<v>",
			node.Tok, node, TypeOf(node))
	}
	return funs
}

// DeclConsts compiles a set of constant declarations
func (c *Comp) DeclConsts(node ast.Spec) {
	switch node := node.(type) {
	case *ast.ValueSpec:
		names, t, inits := c.prepareDeclConstsOrVars(node.Names, node.Type, node.Values)
		c.DeclConsts0(names, t, inits)
	default:
		c.Errorf("Compile: unsupported constant declaration: expecting <*ast.ValueSpec>, found: %v <%v>", node, TypeOf(node))
	}
}

// DeclConsts compiles a set of constant declarations
func (c *Comp) DeclVars(node ast.Spec) []X {
	switch node := node.(type) {
	case *ast.ValueSpec:
		names, t, inits := c.prepareDeclConstsOrVars(node.Names, node.Type, node.Values)
		return c.DeclVars0(names, t, inits)
	default:
		c.Errorf("Compile: unsupported variable declaration: expecting <*ast.ValueSpec>, found: %v <%v>", node, TypeOf(node))
		return nil
	}
}

func (c *Comp) prepareDeclConstsOrVars(idents []*ast.Ident, typ ast.Expr, exprs []ast.Expr) (names []string, t r.Type, inits []*Expr) {
	n := len(idents)
	names = make([]string, n)
	for i, ident := range idents {
		names[i] = ident.Name
	}
	if typ != nil {
		t = c.CompileType(typ)
	}
	if exprs != nil {
		inits = c.ExprsMultipleValues(exprs, n)
	}
	return names, t, inits
}

func (c *Comp) DeclConsts0(names []string, t r.Type, inits []*Expr) {
	n := len(names)
	if inits == nil {
		c.Errorf("constants without initialization: %v", names)
	} else if len(inits) != n {
		c.Errorf("cannot declare %d constants with %d initializers: %v", names)
	}
	for i, name := range names {
		init := inits[i]
		if !init.Const() {
			c.Errorf("const initializer for %q is not a constant", name)
		}
		c.DeclConst0(name, t, init.Value)
	}
}

func (c *Comp) DeclVars0(names []string, t r.Type, inits []*Expr) []X {
	n := len(names)
	ni := len(inits)
	var funs []X
	if ni == 0 {
		funs = make([]X, n)
		for i := 0; i < n; i++ {
			funs[i] = c.DeclVar0(names[i], t, nil)
		}
	} else if ni == n {
		funs = make([]X, n)
		for i := 0; i < n; i++ {
			funs[i] = c.DeclVar0(names[i], t, inits[i])
		}
	} else if ni == 1 && n > 1 {
		funs = append(funs, c.DeclMultiVar0(names, t, inits[0]))
	} else {
		c.Errorf("cannot declare %d variables from %d expressions: %v", n, ni, names)
	}
	return funs
}

// DeclConst0 compiles a constant declaration
func (c *Comp) DeclConst0(name string, t r.Type, value I) {
	if !isLiteral(value) {
		c.Errorf("const initializer for %q is not a constant", name)
		return
	}
	bind := BindValue(value)
	if t != nil {
		value = bind.ConstTo(t)
	}
	// never define bindings for "_"
	if name == "_" {
		return
	}
	if _, ok := c.Binds[name]; ok {
		c.Warnf("redefined identifier: %v", name)
	} else if c.Binds == nil {
		c.Binds = make(map[string]Bind)
	}
	c.Binds[name] = bind
}

// DeclVar0 compiles a variable declaration
func (c *Comp) DeclVar0(name string, t r.Type, init *Expr) X {
	if t == nil {
		if init == nil {
			c.Errorf("no value and no type, cannot declare variable: %v", name)
		}
		t = init.Type
		if t == nil {
			c.Errorf("cannot declare variable as untyped nil: %v", name)
		}
		n := init.NumOut()
		if n == 0 {
			c.Errorf("initializer returns no values, cannot declare variable: %v", name)
		} else if n > 1 {
			c.Errorf("initializer returns %d values, using only the first one to declare variable: %v", name)
		}
	}
	if init == nil || init.Const() {
		idx := c.AddBind(name, t)
		index := idx.Index()
		var value r.Value
		if init == nil || init.Value == nil {
			// no initializer... use the zero-value of t
			value = r.Zero(t)
		} else {
			if init.Type != t {
				init.ConstTo(t)
			}
			value = r.ValueOf(init.Value)
		}
		// never define bindings for "_"
		if idx == NoBind {
			return nil
		} else if idx.IntBind() {
			var n uint64
			switch kindToClass(t.Kind()) {
			case r.Bool:
				if value.Bool() {
					n = 1
				}
			case r.Int:
				n = uint64(value.Int())
			case r.Uint:
				n = value.Uint()
			}
			return func(env *Env) {
				env.IntBinds[index] = n
			}
		} else {
			return func(env *Env) {
				place := r.New(t).Elem()
				place.Set(value)
				env.Binds[index] = place
			}
		}
	}
	if t != init.Type {
		c.Errorf("cannot initialize variable %v <%v> with <%v>", name, t, init.Type)
	}
	idx := c.AddBind(name, t)
	index := idx.Index()
	// never define bindings for "_"
	if idx == NoBind {
		return ToX(init.Fun)
	}
	if idx.IntBind() {
		switch x := init.Fun.(type) {
		case func(*Env) bool:
			return func(env *Env) {
				var n uint64
				if x(env) {
					n = 1
				}
				env.IntBinds[index] = n
			}
		case func(*Env) int:
			return func(env *Env) {
				env.IntBinds[index] = uint64(x(env))
			}
		case func(*Env) int8:
			return func(env *Env) {
				env.IntBinds[index] = uint64(x(env))
			}
		case func(*Env) int16:
			return func(env *Env) {
				env.IntBinds[index] = uint64(x(env))
			}
		case func(*Env) int32:
			return func(env *Env) {
				env.IntBinds[index] = uint64(x(env))
			}
		case func(*Env) int64:
			return func(env *Env) {
				env.IntBinds[index] = uint64(x(env))
			}
		case func(*Env) uint:
			return func(env *Env) {
				env.IntBinds[index] = uint64(x(env))
			}
		case func(*Env) uint8:
			return func(env *Env) {
				env.IntBinds[index] = uint64(x(env))
			}
		case func(*Env) uint16:
			return func(env *Env) {
				env.IntBinds[index] = uint64(x(env))
			}
		case func(*Env) uint32:
			return func(env *Env) {
				env.IntBinds[index] = uint64(x(env))
			}
		case func(*Env) uint64:
			return func(env *Env) {
				env.IntBinds[index] = x(env)
			}
		case func(*Env) uintptr:
			return func(env *Env) {
				env.IntBinds[index] = uint64(x(env))
			}
		default:
			c.Errorf("unsupported expression, cannot use as integer initializer: %v <%T>",
				x, x)
			return nil
		}
	} else {
		x := ToXV(init.Fun)
		return func(env *Env) {
			value, _ := x(env)
			place := r.New(t).Elem()
			place.Set(value)
			env.Binds[index] = place
		}
	}
}

// DeclMultiVar0 compiles multiple variable declarations from a single multi-valued expression
func (c *Comp) DeclMultiVar0(names []string, t r.Type, init *Expr) X {
	if t == nil {
		if init == nil {
			c.Errorf("no value and no type, cannot declare variables: %v", names)
		}
	}
	n := len(names)
	ni := init.NumOut()
	if ni < n {
		c.Errorf("cannot declare %d variables from expression returning %d values: %v", n, ni, names)
	} else if ni > n {
		c.Warnf("declaring %d variables from expression returning %d values: %v", n, ni, names)
	}
	indexes := make([]BindIndex, n)
	ts := make([]r.Type, n)
	for i, name := range names {
		ts[i] = init.Out(i)
		if t != nil && t != ts[i] {
			c.Errorf("cannot use <%v> as <%v> in variable declaration: %v", ts[i], t, names)
		}
		indexes[i] = c.AddBind(name, ts[i])
	}
	init.WithFun()
	fun := ToXV(init.Fun)

	return func(env *Env) {
		values := PackValues(fun(env))
		for i, idx := range indexes {
			v := values[i]
			index := idx.Index()
			if idx.IntBind() {
				var n uint64
				switch kindToClass(ts[i].Kind()) {
				case r.Bool:
					if v.Bool() {
						n = 1
					}
				case r.Int:
					n = uint64(v.Int())
				case r.Uint:
					n = v.Uint()
				}
				env.IntBinds[index] = n
			} else {
				place := r.New(ts[i]).Elem()
				place.Set(v)
				env.Binds[index] = place
			}
		}
	}
}

// AddBind reserves space for a subsequent constant, variable or function declaration
// returns NoBind if name == "_"
func (c *Comp) AddBind(name string, t r.Type) BindIndex {
	if name == "_" {
		return NoBind
	}
	if bind, ok := c.Binds[name]; ok {
		c.Warnf("redefined identifier: %v", name)
		return bind.Index
	} else if c.Binds == nil {
		c.Binds = make(map[string]Bind)
	}
	var idx BindIndex
	// allocate a slot either in IntBinds or in Binds
	if isClass(t.Kind(), r.Bool, r.Int, r.Uint) {
		idx = MakeIntBindIndex(c.IntBindNum)
		c.IntBindNum++
	} else {
		idx = BindIndex(c.BindNum)
		c.BindNum++
	}
	c.Binds[name] = Bind{Lit: Lit{Type: t}, Index: idx}
	c.BindNum++
	return idx
}

/*
func (c *Comp) DeclFunc(name string, paramTypes []r.Type, body X) X {
	idx := c.AddBind(name, typeOfInterface) // FIXME need accurate function type
	xf := c.MakeFunc(paramTypes, body)
	return func(env *Env) (r.Value, []r.Value) {
		f := xf(env)
		env.Binds[idx] = r.ValueOf(f)
		return r.ValueOf(f), nil
	}
}

func (c *Comp) MakeFunc(paramTypes []r.Type, body X) XFunc {
	return func(env *Env) Func {
		return func(args ...r.Value) (ret r.Value, rets []r.Value) {
			fenv := NewEnv(env, 10)
			panicking := true // use a flag to distinguish non-panic from panic(nil)
			defer func() {
				if panicking {
					pan := recover()
					switch p := pan.(type) {
					case SReturn:
						// return is implemented with a panic(SReturn{})
						ret = p.result0
						rets = p.results
					default:
						panic(pan)
					}
				}
			}()
			for i, paramType := range paramTypes {
				place := r.New(paramType).Elem()
				place.Set(args[i])
				fenv.Binds[i] = place
			}
			ret, rets = body(fenv)
			panicking = false
			return ret, rets
		}
	}
}

func MakeFuncInt(paramTypes []r.Type, body X) XFuncInt {
	return func(env *Env) FuncInt {
		return func(args ...r.Value) (ret int) {
			fenv := NewEnv(env, 10)
			panicking := true // use a flag to distinguish non-panic from panic(nil)
			defer func() {
				if panicking {
					pan := recover()
					switch p := pan.(type) {
					case SReturn:
						// return is implemented with a panic(cReturn{})
						ret = int(p.result0.Int())
					default:
						panic(pan)
					}
				}
			}()
			for i, paramType := range paramTypes {
				place := r.New(paramType).Elem()
				place.Set(args[i])
				fenv.Binds[i] = place
			}
			ret0, _ := body(fenv)
			panicking = false
			return int(ret0.Int())
		}
	}
}
*/
