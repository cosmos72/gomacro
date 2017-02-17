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
 * function.go
 *
 *  Created on: Feb 13, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	r "reflect"
)

type Return struct {
	Values []r.Value
}

func (ir *Interpreter) evalDeclFunc(node *ast.FuncDecl) (r.Value, error) {
	name := node.Name.Name
	if name == TemporaryFunctionName {
		// do *NOT* use ir.evalBlock(), because it would create all bindings
		// in its block scope -> they are lost after ir.evalBlock() returns
		return ir.evalStatements(node.Body.List)
	}

	// TODO methods also use node.Recv
	t, err := ir.evalType(node.Type)
	if err != nil {
		return Nil, err
	}
	argNames := ir.evalTypeFieldsNames(node.Type.Params)

	closure := func(args []r.Value) (results []r.Value) {
		return ir.evalFunc(node.Body, t, argNames, args)
	}
	fun := r.MakeFunc(t, closure)

	return ir.defineVar(name, t, fun)
}

// eval an interpreted function
func (ir *Interpreter) evalFunc(body *ast.BlockStmt, t r.Type, argNames []string, args []r.Value) (results []r.Value) {
	ir.PushEnv()
	defer ir.PopEnv()
	defer func() {
		if rec := recover(); rec != nil {
			if ret, ok := rec.(Return); ok {
				results = ret.Values
			} else {
				panic(rec)
			}
		}
	}()

	for i, argName := range argNames {
		ir.defineVar(argName, t.In(i), args[i])
	}
	value, err := ir.evalBlock(body)
	if err != nil {
		panic(err)
	}
	return []r.Value{value}
}
