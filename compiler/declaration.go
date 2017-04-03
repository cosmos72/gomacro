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
		return c.DeclGen(node)
	case *ast.FuncDecl:
		return c.DeclFunc(node, node.Type, node.Body)
	default:
		c.Errorf("Compile: unsupported declaration, expecting <*ast.GenDecl> or <*ast.FuncDecl>, found: %v <%v>", node, TypeOf(node))
		return nil
	}
}

// Decl compiles a constant, variable or type declaration - or an import
func (c *Comp) DeclGen(node *ast.GenDecl) X {
	var ret X
	switch node.Tok {
	case token.IMPORT:
		for _, decl := range node.Specs {
			ret = c.Import(decl)
		}
	case token.CONST:
		for _, decl := range node.Specs {
			ret = c.DeclConsts(decl)
		}
	case token.TYPE:
		for _, decl := range node.Specs {
			return c.DeclType(decl)
		}
	case token.VAR:
		for _, decl := range node.Specs {
			ret = c.DeclVars(decl)
		}
	default:
		c.Errorf("Compile: unsupported declaration kind, expecting token.IMPORT, token.CONST, token.TYPE or token.VAR, found %v: %v %<v>",
			node.Tok, node, TypeOf(node))
		return nil
	}
	return ret
}

// DeclConsts compiles a set of constant declarations
func (c *Comp) DeclConsts(node ast.Spec) X {
	switch node := node.(type) {
	case *ast.ValueSpec:
		names, t, inits := c.prepareDeclConstsOrVars(node.Names, node.Type, node.Values)
		return c.DeclConsts0(names, t, inits)
	default:
		return c.Errorf("Compile: unsupported constant declaration: expecting <*ast.ValueSpec>, found: %v <%v>", node, TypeOf(node))
	}
}

// DeclConsts compiles a set of constant declarations
func (c *Comp) DeclVars(node ast.Spec) X {
	switch node := node.(type) {
	case *ast.ValueSpec:
		names, t, inits := c.prepareDeclConstsOrVars(node.Names, node.Type, node.Values)
		return c.DeclVars0(names, t, inits)
	default:
		return c.Errorf("Compile: unsupported variable declaration: expecting <*ast.ValueSpec>, found: %v <%v>", node, TypeOf(node))
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

func (c *Comp) DeclConsts0(names []string, t r.Type, inits []*Expr) X {
	n := len(names)
	if inits == nil {
		c.Errorf("constants without initialization: %v", names)
	} else if len(inits) != n {
		c.Errorf("cannot declare %d constants with %d initializers: %v", names)
	}
	value0 := None
	values := make([]r.Value, n)
	for i, name := range names {
		init := inits[i]
		if !init.Const() {
			c.Errorf("const initializer for %q is not a constant", name)
		}
		values[i] = r.ValueOf(c.DeclConst0(name, t, init.Value))
	}
	if n != 0 {
		value0 = values[0]
	}
	return func(env *Env) (r.Value, []r.Value) {
		return value0, values
	}
}

func (c *Comp) DeclVars0(names []string, t r.Type, inits []*Expr) X {
	n := len(names)
	ni := len(inits)
	funs := make([]X, n)
	if ni == 0 {
		inits = make([]*Expr, n)
		for i := 0; i < n; i++ {
			funs[i] = c.DeclVar0(names[i], t, nil)
		}
	} else if ni == n {
		for i := 0; i < n; i++ {
			funs[i] = c.DeclVar0(names[i], t, inits[i])
		}
	} else if ni == 1 && n > 1 {
		return c.DeclMultiVar0(names, t, inits[0])
	} else {
		c.Errorf("cannot declare %d variables from %d expressions: %v", n, ni, names)
	}
	return MultipleX(funs)
}

// DeclConst0 compiles a constant declaration
func (c *Comp) DeclConst0(name string, t r.Type, value I) I {
	if !isLiteral(value) {
		c.Errorf("const initializer for %q is not a constant", name)
		return nil
	}
	bind := BindValue(value)
	if t != nil {
		value = bind.ConstTo(t)
	}
	// never define bindings for "_"
	if name == "_" {
		return value
	}
	if _, ok := c.Binds[name]; ok {
		c.Warnf("redefined identifier: %v", name)
	} else if c.Binds == nil {
		c.Binds = make(map[string]Bind)
	}
	c.Binds[name] = bind
	return value
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
	idx := c.AddBind(name, t)
	if init == nil || init.Const() {
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
		if idx >= 0 {
			return func(env *Env) (r.Value, []r.Value) {
				place := r.New(t).Elem()
				place.Set(value)
				env.Binds[idx] = place
				return value, nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return value, nil
			}
		}
	}
	x := ToX(init.Fun)
	// never define bindings for "_"
	if idx >= 0 {
		return func(env *Env) (r.Value, []r.Value) {
			value, _ := x(env)
			place := r.New(t).Elem()
			place.Set(value)
			env.Binds[idx] = place
			return value, nil
		}
	} else {
		return func(env *Env) (r.Value, []r.Value) {
			value, _ := x(env)
			return value, nil
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
	indexes := make([]int, n)
	ts := make([]r.Type, n)
	for i, name := range names {
		ts[i] = init.Out(i)
		if t != nil && t != ts[i] {
			c.Errorf("cannot use <%v> as <%v> in variable declaration: %v", ts[i], t, names)
		}
		indexes[i] = c.AddBind(name, ts[i])
	}
	init.WithFun()
	fun := ToX(init.Fun)

	return func(env *Env) (r.Value, []r.Value) {
		ret, rets := fun(env)
		values := PackValues(ret, rets)
		for i, idx := range indexes {
			place := r.New(ts[i]).Elem()
			place.Set(values[i])
			env.Binds[idx] = place
		}
		return ret, rets
	}
}

// AddBind reserves space for a subsequent constant, variable or function declaration
// returns < 0 if name == "_"
func (c *Comp) AddBind(name string, t r.Type) int {
	if name == "_" {
		return -1
	}
	if bind, ok := c.Binds[name]; ok {
		c.Warnf("redefined identifier: %v", name)
		return bind.Index
	} else if c.Binds == nil {
		c.Binds = make(map[string]Bind)
	}
	idx := c.BindNum
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
