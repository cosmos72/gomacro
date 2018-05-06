/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * api.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

import (
	"fmt"
	"go/ast"
	"go/token"
)

// Support for out-of-order code

type Kind int

const (
	Const Kind = iota
	Expr
	Func
	Import
	Macro
	Method
	Package
	Stmt
	Type
	TypeFwd
	Var
	VarMulti
)

var kinds = map[Kind]string{
	Const:    "Const",
	Expr:     "Expr",
	Func:     "Func",
	Import:   "Import",
	Macro:    "Macro",
	Method:   "Method",
	Package:  "Package",
	Stmt:     "Stmt",
	Type:     "Type",
	TypeFwd:  "TypeFwd", // forward type declaration
	Var:      "Var",
	VarMulti: "VarMulti", // variables initialized with multi-value expression
}

func (k Kind) String() string {
	name, ok := kinds[k]
	if ok {
		return name
	}
	return fmt.Sprintf("Kind%d", int(k))
}

// for multiple const or var declarations in a single *ast.ValueSpec
type Extra struct {
	Ident *ast.Ident
	Type  ast.Expr
	Value ast.Expr
	Iota  int // for constants, value of iota to use
}

// convert *Extra to ast.Spec
func (extra *Extra) Spec() *ast.ValueSpec {
	spec := &ast.ValueSpec{
		Names: []*ast.Ident{extra.Ident},
		Type:  extra.Type,
	}
	if extra.Value != nil {
		spec.Values = []ast.Expr{extra.Value}
	}
	return spec
}

type Decl struct {
	Kind  Kind
	Name  string
	Node  ast.Node // nil for multiple const or var declarations in a single *ast.ValueSpec - in such case, see Extra
	Deps  []string // names of types, constants and variables used in Node's declaration
	Pos   token.Pos
	Extra *Extra
}

func NewDecl(kind Kind, name string, node ast.Node, pos token.Pos, deps []string) *Decl {
	return &Decl{Kind: kind, Name: name, Node: node, Deps: sort_unique_inplace(deps), Pos: pos}
}

type DeclMap map[string]*Decl

type DeclList []*Decl

type Scope struct {
	Decls  DeclMap
	Outer  *Scope
	Gensym int
}

func NewScope(outer *Scope) *Scope {
	return &Scope{
		Decls: make(map[string]*Decl),
		Outer: outer,
	}
}

type Sorter struct {
	scope Scope
	queue []ast.Node
}

func NewSorter() *Sorter {
	return &Sorter{
		scope: Scope{
			Decls: make(map[string]*Decl),
		},
	}
}

// Sorter resolves top-level constant, type, function and var
// declaration order by analyzing their dependencies.
//
// also resolves top-level var initialization order
// analyzing their dependencies.
