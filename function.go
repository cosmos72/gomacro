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
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	r "reflect"
)

func packValues(val0 r.Value, vals []r.Value) []r.Value {
	if len(vals) == 0 && val0 != None {
		vals = []r.Value{val0}
	}
	return vals
}

func unpackValues(vals []r.Value) (r.Value, []r.Value) {
	val0 := None
	if len(vals) > 0 {
		val0 = vals[0]
	}
	return val0, vals
}

func (env *Env) evalDeclFunc(node *ast.FuncDecl) (r.Value, []r.Value) {
	name := node.Name.Name
	if name == temporaryFunctionName {
		// do *NOT* use env.evalBlock(), because it would create all bindings
		// in its block scope -> they are lost after env.evalBlock() returns
		return env.evalStatements(node.Body.List)
	}

	// TODO methods also use node.Recv
	t, argNames, resultNames := env.evalTypeFunction(node.Type)

	closure := func(args []r.Value) (results []r.Value) {
		return env.evalFunc(node.Body, t, argNames, args, resultNames)
	}
	fun := r.MakeFunc(t, closure)
	ret := env.defineVar(name, t, fun)
	return ret, nil
}

// eval an interpreted function
func (env *Env) evalFunc(body *ast.BlockStmt, t r.Type, argNames []string, args []r.Value, resultNames []string) (results []r.Value) {
	env = NewEnv(env)
	defer func() {
		if rec := recover(); rec != nil {
			if ret, ok := rec.(Return); ok {
				results = ret.Results
			} else {
				panic(rec)
			}
		}
	}()

	for i, resultName := range resultNames {
		env.defineVar(resultName, t.Out(i), r.Zero(t.Out(i)))
	}
	for i, argName := range argNames {
		env.defineVar(argName, t.In(i), args[i])
	}
	rets := packValues(env.evalBlock(body))
	for i := range rets {
		rets[i] = rets[i].Convert(t.Out(i))
	}
	results = rets
	return results
}
