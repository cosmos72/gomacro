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
	fun := env.evalExpr1(node.Fun)
	if fun == Nil || fun == None || fun.Kind() != r.Func {
		return env.Errorf("call of non-function %v", node)
	}
	// TODO support the special case fooAcceptsMultipleArgs( barReturnsMultipleValues() )
	args := env.evalExprs(node.Args)
	funt := fun.Type()
	if !funt.IsVariadic() {
		for i, arg := range args {
			args[i] = env.valueToType(arg, funt.In(i))
		}
	}
	var rets []r.Value
	if node.Ellipsis == token.NoPos {
		rets = fun.Call(args)
	} else {
		rets = fun.CallSlice(args)
	}
	return unpackValues(rets)
}
