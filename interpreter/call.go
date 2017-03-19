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
 * call.go
 *
 *  Created on: Feb 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
	r "reflect"
)

func (env *Env) evalCall(node *ast.CallExpr) (r.Value, []r.Value) {
	var fun r.Value
	var t r.Type
	if len(node.Args) == 1 {
		// may be a type conversion
		fun, t = env.evalExpr1OrType(node.Fun)
	} else {
		fun = env.evalExpr1(node.Fun)
	}

	if t != nil {
		val := env.evalExpr1(node.Args[0])
		return env.valueToType(val, t), nil
	} else {
		switch fun.Kind() {
		case r.Struct:
			switch fun := fun.Interface().(type) {
			case Builtin:
				if fun.ArgNum >= 0 && fun.ArgNum != len(node.Args) {
					return env.errorf("builtin %v expects %d arguments, found %d",
						node.Fun, fun.ArgNum, len(node.Args))
				}
				return fun.Exec(env, node.Args)
			case Function:
				if fun.ArgNum >= 0 && fun.ArgNum != len(node.Args) {
					return env.errorf("function %v expects %d arguments, found %d",
						node.Fun, fun.ArgNum, len(node.Args))
				}
				args := env.evalExprs(node.Args)
				return fun.Exec(env, args)
			}
		case r.Func:
			args := env.evalFuncArgs(fun, node)
			var rets []r.Value
			if node.Ellipsis == token.NoPos {
				rets = fun.Call(args)
			} else {
				rets = fun.CallSlice(args)
			}
			return unpackValues(rets)
		default:
			break
		}
	}
	return env.errorf("call of non-function: %v", node)
}

func (env *Env) evalFuncArgs(fun r.Value, node *ast.CallExpr) []r.Value {
	args := env.evalExprs(node.Args)
	funt := fun.Type()
	// TODO does Go have a special case fooAcceptsMultipleArgs( barReturnsMultipleValues() ) ???
	if !funt.IsVariadic() {
		if len(args) != funt.NumIn() {
			env.errorf("function %v expects %d arguments, found %d: %v", node.Fun, funt.NumIn(), len(args), args)
			return nil
		}
		for i, arg := range args {
			args[i] = env.valueToType(arg, funt.In(i))
		}
	}
	return args
}

func (env *Env) evalDefer(node *ast.CallExpr) (r.Value, []r.Value) {
	funcEnv := env.FuncEnv()
	if funcEnv == nil {
		return env.errorf("defer outside function: %v", node)
	}
	fun := env.evalExpr1(node.Fun)
	if fun.Kind() != r.Func {
		return env.errorf("defer of non-function: %v", node)
	}
	args := env.evalFuncArgs(fun, node)
	closure := func() {
		var rets []r.Value
		if node.Ellipsis == token.NoPos {
			rets = fun.Call(args)
		} else {
			rets = fun.CallSlice(args)
		}
		if len(rets) != 0 {
			env.warnf("call to deferred function %v returned %d values, expecting zero: %v", node, rets)
		}
	}
	funcEnv.funcData.defers = append(funcEnv.funcData.defers, closure)
	return None, nil
}
