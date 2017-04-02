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

package compiler

import (
	"go/ast"
	r "reflect"
	"strings"

	. "github.com/cosmos72/gomacro/ast2"
)

func New() *CompEnv {
	top := NewCompEnv(nil, "builtin")
	return NewCompEnv(top, "main")
}

func (c *Comp) CompileAst(in Ast) X {
	for {
		switch form := in.(type) {
		case AstWithNode:
			return ToX(c.Compile(form.Node()))
		case AstWithSlice:
			switch n := form.Size(); n {
			case 0:
				return XNop
			case 1:
				in = form.Get(0)
				continue
			default:
				list := make([]X, n)
				for i := 0; i < n; i++ {
					list[i] = c.CompileAst(form.Get(i))
				}
				return c.Block(list...)
			}
		}
		c.Errorf("CompileAst: unsupported value, expecting <AstWithNode> or <AstWithSlice>, found %v <%v>", in, TypeOf(in))
	}
}

func (c *Comp) Compile(in ast.Node) I {
	switch node := in.(type) {
	case ast.Decl:
		return c.Decl(node)
	case ast.Expr:
		return c.Expr(node)
	case ast.Stmt:
		// TODO return c.Statement(node)
	case *ast.File:
		// TODO return c.File(node)
	default:
	}
	c.Errorf("Compile: unsupported value, expecting <ast.Decl>, <ast.Expr>, <ast.Stmt> or <*ast.File>, found %v <%v>", in, r.TypeOf(in))
	return nil
}

func NewCompEnv(outer *CompEnv, path string) *CompEnv {
	name := path[1+strings.LastIndexByte(path, '/'):]

	var namedTypes map[r.Type]NamedType
	var outerComp *Comp
	var outerEnv *Env
	if outer != nil {
		outerComp = outer.Comp
		outerEnv = outer.Env
		if outerComp != nil {
			namedTypes = outer.NamedTypes
		}
	}
	if namedTypes == nil {
		namedTypes = make(map[r.Type]NamedType)
	}
	c := &CompEnv{
		Comp: &Comp{
			NamedTypes: namedTypes,
			Outer:      outerComp,
			Name:       name,
			Path:       path,
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
		NamedTypes: outer.NamedTypes,
		Outer:      outer,
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

func NewEnv(outer *Env, n int) *Env {
	return &Env{
		Binds: make([]r.Value, n),
		Outer: outer,
	}
}
