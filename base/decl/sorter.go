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
 * loader.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package decl

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base"
)

func (s *Sorter) LoadAst(form ast2.Ast) {
	s.Loader.Ast(form)
}

func (s *Sorter) LoadNode(node ast.Node) {
	s.Loader.Node(node)
}

func (s *Sorter) LoadNodes(nodes []ast.Node) {
	s.Loader.Nodes(nodes)
}

// remove dependencies that cannot be resolved, i.e. not present among the declarations
func (s *Sorter) removeAllUnresolvableDeps() {
	for _, decl := range s.Loader.Decls {
		s.removeUnresolvableDeps(decl)
	}
}

// remove dependencies that cannot be resolved, i.e. not present among the declarations
func (s *Sorter) removeUnresolvableDeps(decl *Decl) {
	decls := s.Loader.Decls

	decl.Deps = filter_if_inplace(decl.Deps, func(name string) bool {
		_, ok := decls[name]
		return ok
	})
}

func (s *Sorter) Sort() []*Decl {
	decls := s.Loader.Decls
	ret := make([]*Decl, 0, len(decls))

	for _, decl := range decls {
		// preserve all deps for later analysis / debugging
		decl.AllDeps = dup(decl.Deps)
	}

	for len(decls) != 0 {
		var list []*Decl
		s.removeAllUnresolvableDeps()

		for _, decl := range decls {
			if len(decl.Deps) == 0 {
				// no dependencies still to resolve
				list = append(list, decl)
				// cannot delete(decls, decl) while iterating
			}
		}
		if len(list) == 0 {
			s.circularDependencyError()
		}
		sort_decls(list)
		ret = append(ret, list...)
		for _, decl := range list {
			delete(decls, decl.Name)
		}
	}
	return ret
}

func (s *Sorter) circularDependencyError() {
	var buf strings.Builder
	buf.WriteString("---- circular dependency! ----\n")
	for _, decl := range s.Loader.Decls {
		// methods cannot cause a circular dependency
		if decl.Kind != Method {
			fmt.Fprintf(&buf, "%s\t%s\t%v\n", decl.Name, decl.Kind, decl)
		}
	}
	base.Errorf("%s", buf.String())
}
