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
func (c *Comp) Decl(node ast.Decl) I {
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
func (c *Comp) DeclGen(node *ast.GenDecl) I {
	var ret I
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

func (c *Comp) prepareDeclConstsOrVars(idents []*ast.Ident, typ ast.Expr, exprs []ast.Expr) (names []string, t r.Type, inits []I) {
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
	if len(inits) == 0 && t == nil {
		c.Errorf("no values and no type: cannot declare %v", names)
	}
	return names, t, inits
}

func (c *Comp) DeclConsts0(names []string, t r.Type, inits []I) X {
	n := len(names)
	funs := make([]X, n)
	if inits == nil {
		for i := 0; i < n; i++ {
			funs[i] = c.DeclConst0(names[i], t, nil)
		}
	} else {
		for i := 0; i < n; i++ {
			funs[i] = c.DeclConst0(names[i], t, inits[i])
		}
	}
	return c.Block(funs...)
}

func (c *Comp) DeclVars0(names []string, t r.Type, inits []I) X {
	n := len(names)
	ni := len(inits)
	funs := make([]X, n)

	if ni == 0 {
		for i := 0; i < n; i++ {
			funs[i] = c.DeclVar0(names[i], t, nil)
		}
	} else if ni == n {
		for i := 0; i < n; i++ {
			funs[i] = c.DeclVar0(names[i], t, inits[i])
		}
	} else if ni == 1 && n > 1 {
		c.Errorf("unimplemented: multiple variable declarations from a single expression returning multiple values, found: %v", names)
	}
	return c.Block(funs...)
}

// DeclConst0 compiles a constant declaration
func (c *Comp) DeclConst0(name string, t r.Type, value I) X {
	if !isLiteral(value) {
		c.Errorf("const initializer for %q is not a constant", name)
		return nil
	}
	v := r.ValueOf(value)
	if t == nil {
		t = r.TypeOf(value)
	}
	idx := c.AddBind(name, t)
	if idx >= 0 {
		return func(env *Env) (r.Value, []r.Value) {
			env.Binds[idx] = v
			return v, nil
		}
	}
	// never define bindings for "_"
	return func(env *Env) (r.Value, []r.Value) {
		return v, nil
	}
}

// DeclVar0 compiles a variable declaration
func (c *Comp) DeclVar0(name string, t r.Type, expr I) X {
	if t == nil {
		t = RetOf(expr)
	}
	idx := c.AddBind(name, t)
	if isLiteral(expr) {
		var value r.Value
		if expr == nil {
			// no initializer... use the zero-value of t
			value = r.Zero(t)
		} else {
			expr = convertLiteral(expr, t)
			value = r.ValueOf(expr)
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
	x := ToX(expr)
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
	idx := len(c.Binds)
	c.Binds[name] = Bind{Index: idx, Type: t}
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
