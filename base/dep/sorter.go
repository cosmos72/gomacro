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
	"go/token"

	"github.com/cosmos72/gomacro/ast2"
)

func (s *Sorter) LoadNode(node ast.Node) {
	s.LoadAst(ast2.ToAst(node))
}

func (s *Sorter) LoadNodes(nodes []ast.Node) {
	s.LoadAst(ast2.NodeSlice{nodes})
}

func (s *Sorter) LoadAst(form ast2.Ast) {
	s.queue = ast2.ToNodesAppend(s.queue, form)
}

// remove all dependencies that cannot be resolved, i.e. not present among s.Loader.Decls
func (s *Sorter) RemoveUnresolvableDeps() {
	s.scope.Decls.RemoveUnresolvableDeps()
}

// return one of:
// * a list of imports
// * a list of declarations
// * a list of expressions and statements
func (s *Sorter) Some() DeclList {

	decls := s.popPackages()
	if len(decls) == 0 {
		decls = s.popImports()
	}
	if len(decls) == 0 {
		decls = s.popDecls()
	}
	if len(decls) == 0 {
		decls = s.popStmts()
	}
	return decls
}

func (s *Sorter) popPackages() []*Decl {
	var specs []ast.Spec
	i, n := 0, len(s.queue)
loop:
	for ; i < n; i++ {
		node := s.queue[i]
		switch node := node.(type) {
		case nil:
			continue
		case *ast.GenDecl:
			if node != nil && node.Tok == token.PACKAGE {
				specs = append(specs, node.Specs...)
				continue
			}
		}
		// /*DELETEME*/ fmt.Printf("popPackages stopping at node: %v %T\n", node, node)
		break loop
	}
	if i > 0 {
		s.queue = s.queue[i:]
	}
	if len(specs) == 0 {
		return nil
	}
	s.scope.Decls = make(map[string]*Decl)

	for _, spec := range specs {
		s.scope.Package(spec)
	}
	list := s.scope.Decls.List().SortByPos()

	s.scope.Decls = nil
	return list
}

func (s *Sorter) popImports() []*Decl {
	var specs []ast.Spec
	i, n := 0, len(s.queue)
loop:
	for ; i < n; i++ {
		node := s.queue[i]
		switch node := node.(type) {
		case nil:
			continue
		case *ast.GenDecl:
			if node != nil && node.Tok == token.IMPORT {
				specs = append(specs, node.Specs...)
				continue
			}
		}
		// /*DELETEME*/ fmt.Printf("popImports stopping at node: %v %T\n", node, node)
		break loop
	}
	if i > 0 {
		s.queue = s.queue[i:]
	}
	if len(specs) == 0 {
		return nil
	}
	s.scope.Decls = make(map[string]*Decl)

	for _, spec := range specs {
		s.scope.Import(spec)
	}
	list := s.scope.Decls.List().SortByPos()

	s.scope.Decls = nil
	return list
}

func (s *Sorter) popDecls() []*Decl {
	var nodes []ast.Node
	i, n := 0, len(s.queue)
loop:
	for ; i < n; i++ {
		node := s.queue[i]
		switch node := node.(type) {
		case nil:
			continue
		case *ast.GenDecl:
			if node != nil && node.Tok != token.IMPORT && node.Tok != token.PACKAGE {
				nodes = append(nodes, node)
				continue
			}
		case ast.Decl:
			if node != nil {
				nodes = append(nodes, node)
				continue
			}
		}
		// /*DELETEME*/ fmt.Printf("popDecls stopping at node: %v %T\n", node, node)
		break loop
	}
	if i > 0 {
		s.queue = s.queue[i:]
	}
	if len(nodes) == 0 {
		return nil
	}
	s.scope.Decls = make(map[string]*Decl)

	s.scope.Nodes(nodes)
	s.scope.Decls.RemoveUnresolvableDeps()
	m := s.scope.Decls.Dup()

	s.scope.Decls = nil

	g := graph{
		Nodes: m,
		Edges: m.depMap(),
	}
	return g.Sort()
}

func (s *Sorter) popStmts() []*Decl {
	var nodes []ast.Node
	i, n := 0, len(s.queue)
loop:
	for ; i < n; i++ {
		node := s.queue[i]
		switch node := node.(type) {
		case nil:
			continue
		case ast.Expr:
			if node != nil {
				nodes = append(nodes, node)
				continue
			}
		case ast.Stmt:
			if node != nil {
				nodes = append(nodes, node)
				continue
			}
		}
		// /*DELETEME*/ fmt.Printf("popStmts stopping at node: %v %T\n", node, node)
		break loop
	}
	if i > 0 {
		s.queue = s.queue[i:]
	}
	if len(nodes) == 0 {
		return nil
	}
	s.scope.Decls = make(map[string]*Decl)

	s.scope.Nodes(nodes)
	list := s.scope.Decls.List().SortByPos()

	s.scope.Decls = nil
	return list
}
