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
 *  Created on: Feb 15, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"errors"
	"fmt"
	"go/ast"
	r "reflect"
)

func (ir *Interpreter) evalCall(node *ast.CallExpr) (r.Value, error) {
	fun, err := ir.evalExpr(node.Fun)
	if err != nil {
		return Nil, err
	}
	if fun.Kind() != r.Func {
		return Nil, errors.New(fmt.Sprintf("call of non-function %#v", node))
	}
	args, err := ir.evalExprs(node.Args)
	if err != nil {
		return Nil, err
	}
	rets := fun.Call(args)
	if len(rets) == 0 {
		return Nil, nil
	}
	// TODO return multiple values
	return rets[0], nil
	// return Nil, errors.New(fmt.Sprintf("unsupported function call: %#v", node))
}
