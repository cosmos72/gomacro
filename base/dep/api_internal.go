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
 * api_internal.go
 *
 *  Created on: May 05, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

type void struct{}

type set map[string]void

type depMap map[string]set

type fwdDeclList struct {
	List DeclList
	Set  set
}

type graph struct {
	Nodes DeclMap
	Edges depMap
}

type visitCtx struct {
	visiting   map[string]int
	visited    map[string]int
	beforeFunc func(node *Decl, ctx *visitCtx) // invoked once for each node, in visit pre-order
	afterFunc  func(node *Decl, ctx *visitCtx) // invoked once for each node, in visit post-order
	cycleFunc  func(node *Decl, ctx *visitCtx) // invoked when ctx.visiting[node.Name] exists already, i.e. for cycles
}
