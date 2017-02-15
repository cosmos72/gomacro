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
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"
)

func (ir *Interpreter) evalAssignments(node *ast.AssignStmt) (r.Value, error) {
	lhs := node.Lhs
	rhs := node.Rhs
	op := node.Tok
	var ret r.Value
	var err error
	for i, n := 0, len(lhs); i < n; i++ {
		ret, err = ir.evalAssignment(lhs[i], op, rhs[i])
		if err != nil {
			return Nil, err
		}
	}
	return ret, nil
}

func (ir *Interpreter) evalAssignment(lhs ast.Expr, op token.Token, rhs ast.Expr) (r.Value, error) {
	// fmt.Printf("debug: evalAssignment() [%v] %s [%v]\n", lhs, op, rhs)
	place, err := ir.Eval(lhs)
	if err != nil {
		return Nil, err
	}
	if !place.CanSet() {
		return Nil, errors.New(fmt.Sprintf("cannot assign to read-only location: %#v %s %#v", lhs, op, rhs))
	}
	value, err := ir.Eval(rhs)
	if err != nil {
		return Nil, err
	}
	return ir.assign(place, token.ASSIGN, value)
}

func (ir *Interpreter) assign(place r.Value, op token.Token, value r.Value) (r.Value, error) {
	t := place.Type()
	if !value.Type().AssignableTo(t) {
		if !value.Type().ConvertibleTo(t) {
			return Nil, errors.New(fmt.Sprintf("failed to convert %#v to %v", value, t))
		}
		if primitiveTypeOverflows(value, place) {
			fmt.Fprintf(ir.Eout, "warning: value %#v overflows %v, truncating to %#v\n", value, t, value.Convert(t))
		}
		value = value.Convert(t)
	}
	// TODO op can be = += -= *= ...
	place.Set(value)
	return value, nil
}
