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

func (env *Env) evalDeclNamedFunction(node *ast.FuncDecl) (r.Value, []r.Value) {
	name := node.Name.Name
	if name == temporaryFunctionName {
		// do *NOT* use env.evalBlock(), because it would create all bindings
		// in its block scope -> they are lost after env.evalBlock() returns
		return env.evalStatements(node.Body.List)
	}

	fun, t := env.evalDeclFunction(node, node.Type, node.Body)
	ret := env.defineVar(name, t, fun)
	return ret, nil
}

func (env *Env) evalDeclFunction(nodeForReceiver *ast.FuncDecl, funcType *ast.FuncType, body *ast.BlockStmt) (r.Value, r.Type) {
	var ret r.Value
	isMacro := false
	if nodeForReceiver != nil && nodeForReceiver.Recv != nil {
		recvList := nodeForReceiver.Recv.List
		if recvList != nil && len(recvList) == 0 {
			isMacro = true
		} else {
			// TODO implement receiver
			env.Errorf("unimplemented: method declarations (i.e. functions with receiver): %v", nodeForReceiver)
			return ret, nil
		}
	}
	t, argNames, resultNames := env.evalTypeFunction(funcType)

	closure := func(args []r.Value) (results []r.Value) {
		return env.evalFuncCall(body, t, argNames, args, resultNames)
	}
	if isMacro {
		ret = r.ValueOf(Macro{Closure: closure, ArgN: len(argNames)})
		t = ret.Type()
	} else {
		ret = r.MakeFunc(t, closure)
	}
	return ret, t
}

// eval an interpreted function
func (env *Env) evalFuncCall(body *ast.BlockStmt, t r.Type, argNames []string, args []r.Value, resultNames []string) (results []r.Value) {
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
