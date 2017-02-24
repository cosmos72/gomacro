/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * assignment.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
	r "reflect"
)

func (env *Env) evalAssignments(node *ast.AssignStmt) (r.Value, []r.Value) {
	left := node.Lhs
	right := node.Rhs
	op := node.Tok
	nleft := len(left)
	nright := len(right)

	if nright != 1 && nleft != nright {
		return env.Errorf("value count mismatch: cannot assign %d values to %d places: %v", nright, nleft, node)
	}

	// side effects happen left to right, with some unspecified cases,
	// so first Eval() all node.Lhs, then Eval() all node.Rhs
	// https://golang.org/ref/spec#Order_of_evaluation

	if op == token.DEFINE {
		names := make([]string, nleft)
		for i := 0; i < nleft; i++ {
			ident, ok := left[i].(*ast.Ident)
			if !ok {
				return env.Errorf("variable declaration: invalid identifier: %v", left[i])
			}
			names[i] = ident.Name
		}
		values := env.evalExprsMultipleValues(right, nleft)
		return env.defineVars(names, nil, values)

	} else {
		places := env.evalPlaces(left)
		values := env.evalExprsMultipleValues(right, nleft)
		return env.assignPlaces(places, op, values)
	}
}

func (env *Env) evalPlaces(node []ast.Expr) []r.Value {
	n := len(node)
	places := make([]r.Value, n)
	for i := 0; i < n; i++ {
		place := env.evalExpr1(&node[i])
		if !place.CanSet() {
			return env.PackErrorf("cannot assign to read-only location: %v", node[i])
		}
		places[i] = place
	}
	return places
}

func (env *Env) assignPlaces(places []r.Value, op token.Token, values []r.Value) (r.Value, []r.Value) {
	n := len(places)
	for i := 0; i < n; i++ {
		values[i] = env.assignPlace(places[i], op, values[i])
	}
	return unpackValues(values)
}

func (env *Env) assignPlace(place r.Value, op token.Token, value r.Value) r.Value {
	t := place.Type()
	value = env.valueToType(value, t)
	// TODO op can be = += -= *= ...
	place.Set(value)
	return value
}
