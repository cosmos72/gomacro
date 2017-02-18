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
 *  Created on: Feb 13, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	"go/token"
	r "reflect"
)

func (env *Env) evalAssignments(node *ast.AssignStmt) (r.Value, []r.Value) {
	lhs := node.Lhs
	rhs := node.Rhs
	op := node.Tok
	var ret r.Value
	var rets []r.Value
	for i, n := 0, len(lhs); i < n; i++ {
		ret, rets = env.evalAssignment(lhs[i], op, rhs[i])
	}
	return ret, rets
}

func (env *Env) evalAssignment(lhs ast.Expr, op token.Token, rhs ast.Expr) (r.Value, []r.Value) {
	// fmt.Printf("debug: evalAssignment() %v %s %v\n", lhs, op, rhs)

	value, _ := env.Eval(rhs)

	if op == token.DEFINE {
		// fmt.Printf("debug: evalAssignment() DEFINE %v %#v\n", lhs, value)
		ident, ok := lhs.(*ast.Ident)
		if ok {
			return env.defineVar(ident.Name, nil, value)
		}
	}
	place, _ := env.Eval(lhs)
	if !place.CanSet() {
		return Errorf("cannot assign to read-only location: %#v %s %#v", lhs, op, rhs)
	}
	return env.assign(place, token.ASSIGN, value)
}

func (env *Env) assign(place r.Value, op token.Token, value r.Value) (r.Value, []r.Value) {
	t := place.Type()
	if !value.Type().AssignableTo(t) {
		if !value.Type().ConvertibleTo(t) {
			return Errorf("failed to convert %#v to %v", value, t)
		}
		if primitiveTypeOverflows(value, place) {
			Warnf("value %#v overflows %v, truncating to %#v\n", value, t, value.Convert(t))
		}
		value = value.Convert(t)
	}
	// TODO op can be = += -= *= ...
	place.Set(value)
	return value, nil
}
