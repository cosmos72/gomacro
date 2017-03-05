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

type Place struct {
	obj    r.Value // the map to modify, or a settable r.Value
	mapkey r.Value // the map key to set, or Nil
}

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

func (env *Env) evalPlaces(node []ast.Expr) []Place {
	n := len(node)
	places := make([]Place, n)
	for i := 0; i < n; i++ {
		places[i] = env.evalPlace(node[i])
	}
	return places
}

func (env *Env) evalPlace(node ast.Expr) Place {
	obj := Nil
	// ignore parenthesis: (expr) = value is the same as expr = value
	for {
		if paren, ok := node.(*ast.ParenExpr); ok {
			node = paren.X
		} else {
			break
		}
	}
	switch node := node.(type) {
	case *ast.IndexExpr:
		obj = env.evalExpr1(node.X)
		index := env.evalExpr1(node.Index)

		switch obj.Kind() {
		case r.Map:
			return Place{obj, index}
		case r.Array, r.Slice, r.String:
			i, ok := env.toInt(index)
			if !ok {
				env.Errorf("invalid index, expecting an int: %v <%v>", index, TypeOf(index))
				return Place{}
			}
			obj = obj.Index(int(i))
		default:
			env.Errorf("unsupported index operation: %v [ %v ]. not an array, map, slice or string: %v <%v>",
				node.X, index, obj, TypeOf(obj))
			return Place{}
		}
	default:
		obj = env.evalExpr1(node)
	}
	if !obj.CanSet() {
		env.Errorf("cannot assign to read-only location: %v", node)
		return Place{}
	}
	return Place{obj, Nil}
}

func (env *Env) assignPlaces(places []Place, op token.Token, values []r.Value) (r.Value, []r.Value) {
	n := len(places)
	for i := 0; i < n; i++ {
		values[i] = env.assignPlace(places[i], op, values[i])
	}
	return unpackValues(values)
}

func (env *Env) assignPlace(place Place, op token.Token, value r.Value) r.Value {
	obj := place.obj
	key := place.mapkey
	if key == Nil {
		t := TypeOf(obj)
		value = env.valueToType(value, t)
		if op != token.ASSIGN {
			value = env.evalBinaryExpr(obj, op, value)
		}
		obj.Set(value)
		return value
	}
	// map[key] OP value
	env.Debugf("setting map[key]: %v <%v> [%v <%v>] %s %v <%v>",
		obj, TypeOf(obj), key, TypeOf(key), op, value, TypeOf(value))
	currValue, _, t := MapIndex(obj, key)
	value = env.valueToType(value, t)
	if op != token.ASSIGN {
		value = env.evalBinaryExpr(currValue, op, value)
	}
	obj.SetMapIndex(key, value)
	return value
}
