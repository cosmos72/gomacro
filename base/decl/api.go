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

package decl

import (
	"fmt"
	"go/ast"
)

// Support for out-of-order code

type Kind int

const (
	// sort by typical order of appearance
	Type Kind = iota
	Const
	Var
	Func
	Method
)

var kinds = map[Kind]string{
	Type:   "Type",
	Const:  "Const",
	Var:    "Var",
	Func:   "Func",
	Method: "Method",
}

func (k Kind) String() string {
	name, ok := kinds[k]
	if ok {
		return name
	}
	return fmt.Sprintf("Kind%d", int(k))
}

type Decl struct {
	Kind    Kind
	Name    string
	Node    ast.Node
	AllDeps []string // names of types, constants and variables used in Node's declaration
	Deps    []string
	Iota    int // for constants, value of iota to use
}

type Loader struct {
	Decls  map[string]*Decl
	Gensym int
}

func NewLoader() *Loader {
	return &Loader{
		Decls: make(map[string]*Decl),
	}
}

type Sorter struct {
	Loader Loader
}

func NewSorter() *Sorter {
	return &Sorter{
		Loader: Loader{
			Decls: make(map[string]*Decl),
		},
	}
}

// Sorter resolves top-level constant, type, function and var
// declaration order by analyzing their dependencies.
//
// also resolves top-level var initialization order
// analyzing their dependencies.
