/*
 * gomacro - A Go intepreter with Lisp-like macros
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
 * compile.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
	r "reflect"
	"strings"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
)

func New() *CompEnv {
	top := NewCompEnv(nil, "builtin")
	top.growEnv(128)
	env := NewCompEnv(top, "main")
	env.growEnv(1024)
	return env
}

func (c *Comp) CompileAst(in Ast) func(*Env) (r.Value, []r.Value) {
	for {
		switch form := in.(type) {
		case nil:
			return nil
		case AstWithNode:
			return c.Compile(form.Node())
		case AstWithSlice:
			switch n := form.Size(); n {
			case 0:
				return nil
			case 1:
				in = form.Get(0)
				continue
			default:
				var list []func(*Env) (r.Value, []r.Value)
				for i := 0; i < n; i++ {
					fun := c.CompileAst(form.Get(i))
					if fun != nil {
						list = append(list, fun)
					}
				}
				return func(env *Env) (r.Value, []r.Value) {
					n_1 := len(list) - 1
					for i := 0; i < n_1; i++ {
						list[i](env)
					}
					return list[n_1](env)
				}
			}
		}
		c.Errorf("CompileAst: unsupported value, expecting <AstWithNode> or <AstWithSlice>, found %v <%v>", in, r.TypeOf(in))
		return nil
	}
}

func (c *Comp) Compile(in ast.Node) func(*Env) (r.Value, []r.Value) {
	switch node := in.(type) {
	case ast.Decl:
		return ToXV(c.Decl(node), c.CompileOptions)
	case ast.Expr:
		return c.Expr(node).AsXV(c.CompileOptions)
	case *ast.ExprStmt:
		// special case of statement
		return c.Expr(node.X).AsXV(c.CompileOptions)
	case ast.Stmt:
		// TODO return c.Statement(node)
		c.Errorf("Compile: unimplemented <ast.Stmt>, found %v <%v>", in, r.TypeOf(in))
	case *ast.File:
		c.Errorf("Compile: unimplemented <*ast.File>, found %v <%v>", in, r.TypeOf(in))
		// TODO return c.File(node)
	default:
	}
	c.Errorf("Compile: unsupported expression, expecting <ast.Decl>, <ast.Expr>, <ast.Stmt> or <*ast.File>, found %v <%v>", in, r.TypeOf(in))
	return nil
}

func NewCompEnv(outer *CompEnv, path string) *CompEnv {
	name := path[1+strings.LastIndexByte(path, '/'):]

	var namedTypes map[r.Type]NamedType
	var outerComp *Comp
	var outerEnv *Env
	var base *InterpreterBase
	if outer != nil {
		outerComp = outer.Comp
		outerEnv = outer.Env
		if outerComp != nil {
			namedTypes = outer.NamedTypes
		}
		base = outer.InterpreterBase
	}
	if base == nil {
		b := MakeInterpreterBase()
		base = &b
	}
	if namedTypes == nil {
		namedTypes = make(map[r.Type]NamedType)
	}
	c := &CompEnv{
		Comp: &Comp{
			NamedTypes:      namedTypes,
			Outer:           outerComp,
			Name:            name,
			Path:            path,
			InterpreterBase: base,
		},
		Env: &Env{Outer: outerEnv},
	}
	if outer == nil {
		c.addBuiltins()
	}
	return c

}

func NewComp(outer *Comp) *Comp {
	if outer == nil {
		return &Comp{}
	}
	return &Comp{
		NamedTypes:     outer.NamedTypes,
		Outer:          outer,
		CompileOptions: outer.CompileOptions,
	}
}

func (c *Comp) Top() *Comp {
	for ; c != nil; c = c.Outer {
		if c.Outer == nil {
			break
		}
	}
	return c
}

func (c *Comp) File() *Comp {
	for ; c != nil; c = c.Outer {
		outer := c.Outer
		if outer == nil || outer.Outer == nil {
			break
		}
	}
	return c
}

func NewEnv(outer *Env, nbinds int, nintbinds int) *Env {
	return &Env{
		Binds:    make([]r.Value, nbinds),
		IntBinds: make([]uint64, nintbinds),
		Outer:    outer,
	}
}
