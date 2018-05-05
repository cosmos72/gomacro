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
 * util.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

import (
	"fmt"
	"io"
	"os"
	"sort"
)

// ===================== DeclMap =====================

func (m DeclMap) Dup() DeclMap {
	ret := make(DeclMap, len(m))
	for name, decl := range m {
		ret[name] = decl
	}
	return ret
}

func (m DeclMap) List() DeclList {
	list := make(DeclList, len(m))
	i := 0
	for _, e := range m {
		list[i] = e
		i++
	}
	return list
}

// remove all dependencies that cannot be resolved, i.e. not present among m
func (m DeclMap) RemoveUnresolvableDeps() {
	for _, decl := range m {
		decl.RemoveUnresolvableDeps(m)
	}
}

func (m DeclMap) Print() {
	m.List().SortByPos().Print()
}

func (m DeclMap) depMap() depMap {
	ret := make(depMap, len(m))
	for _, decl := range m {
		ret[decl.Name] = decl.depSet()
	}
	return ret
}

// ===================== DeclList ====================

func (list DeclList) Map() DeclMap {
	m := make(DeclMap, len(list))
	for _, e := range list {
		m[e.Name] = e
	}
	return m
}

func (list DeclList) SortByPos() DeclList {
	sort.Slice(list, func(i, j int) bool {
		a, b := list[i], list[j]
		return a.Pos < b.Pos
	})
	return list
}

func (list DeclList) Reverse() DeclList {
	n := len(list)
	for i := 0; i < n/2; i++ {
		temp := list[i]
		j := n - i - 1
		list[i] = list[j]
		list[j] = temp
	}
	return list
}

func (list DeclList) Print() {
	for _, decl := range list {
		decl.Print()
	}
}

// ======================= Decl ======================

func (decl *Decl) depSet() set {
	ret := make(set, len(decl.Deps))
	for _, dep := range decl.Deps {
		ret[dep] = void{}
	}
	return ret
}

// remove all dependencies that cannot be resolved, i.e. not present among m
func (decl *Decl) RemoveUnresolvableDeps(m DeclMap) {
	decl.Deps = filter_if_inplace(decl.Deps, func(name string) bool {
		return m[name] != nil
	})
}

func (decl *Decl) Fprint(out io.Writer) {
	fmt.Fprintf(out, "%s%s%s\t%v\n", decl.Name, spaces(decl.Name), decl.Kind, decl)
}

func (decl *Decl) Print() {
	decl.Fprint(os.Stdout)
}

const _spaces = "                                "

func spaces(name string) string {
	return _spaces[len(name)%32:]
}
