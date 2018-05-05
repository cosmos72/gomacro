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
 * sorter.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

import (
	"go/ast"

	"github.com/cosmos72/gomacro/ast2"
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

// remove all dependencies that cannot be resolved, i.e. not present among s.Loader.Decls
func (s *Sorter) RemoveUnresolvableDeps() {
	s.Loader.Decls.RemoveUnresolvableDeps()
}

func (s *Sorter) Sort() DeclList {
	s.RemoveUnresolvableDeps()

	m := s.Loader.Decls.Dup()
	g := graph{
		Nodes: m,
		Edges: m.depMap(),
	}
	return g.Sort()
}
