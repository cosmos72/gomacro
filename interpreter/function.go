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

func (env *Env) evalDeclNamedFunction(node *ast.FuncDecl) (r.Value, []r.Value) {
	name := node.Name.Name
	if name == temporaryFunctionName {
		// do *NOT* use env.evalBlock(), because it would create all bindings
		// in its block scope -> they are lost after env.evalBlock() returns
		return env.evalStatements(node.Body.List)
	}

	fun, t := env.evalDeclFunction(node, node.Type, node.Body)
	ret := env.defineFunc(name, t, fun)
	return ret, nil
}

func (env *Env) evalDeclFunction(decl *ast.FuncDecl, funcType *ast.FuncType, body *ast.BlockStmt) (r.Value, r.Type) {
	var ret r.Value
	isMacro := false

	if decl != nil && decl.Recv != nil {
		recvList := decl.Recv.List
		if recvList != nil && len(recvList) == 0 {
			isMacro = true
		} else {
			// TODO implement receiver
			env.errorf("unimplemented: method declarations (i.e. functions with receiver): %v", decl)
			return ret, nil
		}
	}
	t, argNames, resultNames := env.evalTypeFunction(funcType)
	tret := t

	funcName := makeFuncNameForEnv(decl, isMacro)

	closure := func(args []r.Value) (results []r.Value) {
		return env.evalFuncCall(funcName, body, t, argNames, args, resultNames)
	}
	if isMacro {
		// env.Debugf("defined macro %v, type %v, args (%v), returns (%v)", decl.Name.Name, t, strings.Join(argNames, ", "), strings.Join(resultNames, ", "))
		ret = r.ValueOf(Macro{Closure: closure, ArgNum: len(argNames)})
		tret = typeOf(ret) // do NOT change t, is needed by the closure above
	} else {
		ret = r.MakeFunc(t, closure)
	}
	return ret, tret
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

// eval an interpreted function
func (env *Env) evalFuncCall(envName string, body *ast.BlockStmt, t r.Type, argNames []string, args []r.Value, resultNames []string) (results []r.Value) {
	if t.Kind() != r.Func {
		return env.packErrorf("call of non-function type %v", t)
	}
	env = NewEnv(env, envName)
	debugCall := env.Options&OptDebugCallStack != 0
	// register this function call in the call stack prepared by evalCall()
	{
		stack := env.CallStack.Frame
		depth := len(stack) - 1
		env.funcData = &funcData{CallDepth: depth}
		if depth >= 0 {
			stack[depth].FuncEnv = env
		}
		if debugCall {
			env.debugf("func starting: %s, args = %v, call stack is:", envName, args)
			env.showStack()
		}
	}

	panicking := true // use a flag to distinguish non-panic from panic(nil)
	defer func() {
		if debugCall {
			env.debugf("func exiting:  %s, panicking = %v, stack length = %d",
				envName, panicking, len(env.CallStack.Frame))
		}
		f := env.funcData
		if panicking {
			pan := recover()
			switch p := pan.(type) {
			case eReturn:
				// return is implemented with a panic(Return{})
				results = env.convertFuncCallResults(t, p.results, true)
			default: // some interpreted or compiled code invoked panic()
				if env.Options&OptDebugPanicRecover != 0 {
					env.debugf("captured panic for defers: env = %v, panic = %#v", env.Name, p)
				}
				f.panick = &p
			}
		}
		if len(f.defers) != 0 {
			env.runDefers()
		}
		if debugCall {
			env.debugf("func exited:   %s, panic     = %v, stack length = %d",
				envName, f.panick, len(env.CallStack.Frame))
		}
		if pp := f.panick; pp != nil {
			panic(*pp)
		}
	}()

	for i, resultName := range resultNames {
		env.defineVar(resultName, t.Out(i), r.Zero(t.Out(i)))
	}
	for i, argName := range argNames {
		env.defineVar(argName, t.In(i), args[i])
	}
	// use evalStatements(), not evalBlock(): in Go, the function arguments and body are in the same scope
	rets := packValues(env.evalStatements(body.List))
	results = env.convertFuncCallResults(t, rets, false)
	panicking = false
	return results
}

func (env *Env) convertFuncCallResults(t r.Type, rets []r.Value, warn bool) []r.Value {
	retsN := len(rets)
	expectedN := t.NumOut()
	if retsN < expectedN {
		if warn {
			env.warnf("not enough return values: expected %d, found %d: %v", expectedN, retsN, rets)
		}
		tmp := make([]r.Value, expectedN)
		copy(tmp, rets)
		rets = tmp
	} else if retsN > expectedN {
		if warn {
			env.warnf("too many return values: expected %d, found %d: %v", expectedN, retsN, rets)
		}
		rets = rets[:expectedN]
	}
	for i := range rets {
		rets[i] = rets[i].Convert(t.Out(i))
	}
	return rets
}

func (env *Env) runDefers() {
	// execute defers last-to-first
	env.funcData.runningDefers = true
	if env.Options&OptDebugCallStack != 0 {
		env.debugf("func defers:   %s, panic = %v, stack length = %d",
			env.Name, env.funcData.panick, len(env.CallStack.Frame))
	}
	defers := env.funcData.defers
	for i := len(defers) - 1; i >= 0; i-- {
		env.runDefer(defers[i])
	}
}

func (env *Env) runDefer(deferred func()) {
	// invoking panic() inside a deferred function exits it with a panic,
	// but the previously-installed deferred functions are still executed
	// and can recover() such panic

	panicking := true // use a flag to distinguish non-panic from panic(nil)
	env.CallStack.Frame = append(env.CallStack.Frame, CallFrame{CallerEnv: env})
	defer func() {
		env.CallStack.Frame = env.CallStack.Frame[0 : len(env.CallStack.Frame)-1]
		if panicking {
			pan := recover()
			env.funcData.panick = &pan
		}
	}()
	deferred()
	panicking = false
}
