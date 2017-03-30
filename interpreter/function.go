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

package interpreter

import (
	"fmt"
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

func (env *Env) evalDeclFunction(decl *ast.FuncDecl, funcType *ast.FuncType, body *ast.BlockStmt) (r.Value, []r.Value) {
	isMacro := false
	var recv *ast.Field

	if decl != nil && decl.Recv != nil {
		recvList := decl.Recv.List
		if recvList != nil && len(recvList) == 0 {
			isMacro = true
		} else {
			recv = recvList[0]
		}
	}
	tFunc, tFuncOrMethod, argNames, resultNames := env.evalTypeFunctionOrMethod(recv, funcType)
	tret := tFuncOrMethod

	var funcName string
	if decl == nil {
		funcName = makeFuncNameForEnv(decl, isMacro)
	} else {
		funcName = decl.Name.Name
	}

	closure := func(args []r.Value) (results []r.Value) {
		return env.evalFuncCall(funcName, body, tFuncOrMethod, argNames, args, resultNames)
	}
	var ret r.Value
	if isMacro {
		// env.Debugf("defined macro %v, type %v, args (%v), returns (%v)", decl.Name.Name, t, strings.Join(argNames, ", "), strings.Join(resultNames, ", "))
		ret = r.ValueOf(Macro{Closure: closure, ArgNum: len(argNames)})
		tret = ret.Type()
	} else {
		ret = r.MakeFunc(tFuncOrMethod, closure)

		if decl != nil && recv != nil {
			recvType := tFuncOrMethod.In(0)
			// register tFunc, i.e. without the receiver, to allow comparison with Interface methods
			env.registerMethod(recvType, funcName, tFunc, ret)
		}
	}
	if decl != nil && recv == nil {
		// register named functions and macros (NOT methods) in the current environment
		ret = env.defineFunc(funcName, tret, ret)
	}
	return ret, nil
}

func makeFuncNameForEnv(decl *ast.FuncDecl, isMacro bool) string {
	var prefix, space, suffix string = "func", "", ""
	if isMacro {
		prefix = "macro"
	}
	if decl != nil {
		space = " "
		suffix = decl.Name.Name
	}
	return fmt.Sprintf("%s%s%s()", prefix, space, suffix)
}
