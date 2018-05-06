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
 * reader.go
 *
 *  Created on Mar 05, 2018
 *      Author Massimiliano Ghilardi
 */

package ast2

import (
	"go/ast"
)

type pair struct {
	form AstWithSlice
	idx  int
}

// zero value is a valid, empty Reader,
// use AppendAst() or AppendNode() to populate it.
// In alternative, create Reader with AstReader() or NodeReader
type Reader struct {
	queue []Ast
	stack []pair
	Node  ast.Node
}

func AstReader(form Ast) *Reader {
	r := Reader{}
	r.enqueue(form)
	return &r
}

func NodeReader(node ast.Node) *Reader {
	return AstReader(ToAst(node))
}

func (r *Reader) AppendAst(form Ast) *Reader {
	r.enqueue(form)
	return r
}

func (r *Reader) AppendNode(node ast.Node) *Reader {
	r.enqueue(ToAst(node))
	return r
}

func (r *Reader) Empty() bool {
	return len(r.queue) == 0 && len(r.stack) == 0 && r.Node == nil
}

func (r *Reader) Next() {
restart:
	if len(r.stack) == 0 {
		form := r.dequeue()
		if form == nil {
			r.Node = nil
			return
		}
		if r.enter(form) {
			return
		}
	}
	n := len(r.stack)
	for n != 0 {
		pair := &r.stack[n-1]
		pair.idx++
		if pair.idx >= pair.form.Size() {
			n--
			r.stack = r.stack[:n]
			continue
		}
		if r.enter(pair.form.Get(pair.idx)) {
			return
		}
	}
	goto restart
}

func (r *Reader) enqueue(form Ast) {
	if form != nil {
		r.queue = append(r.queue, form)
	}
}

func (r *Reader) dequeue() Ast {
	n := len(r.queue)
	if n == 0 {
		return nil
	}
	form := r.queue[n-1]
	r.queue = r.queue[:n-1]
	return form
}

func (r *Reader) enter(in Ast) bool {
	for {
		switch form := in.(type) {
		case AstWithNode:
			r.Node = form.Node()
			return true
		case AstWithSlice:
			if form.Size() == 0 {
				return false
			}
			r.stack = append(r.stack, pair{form, 0})
			in = form.Get(0)
			continue
		}
	}
}
