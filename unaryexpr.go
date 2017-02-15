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
 * unaryexpr.go
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

func unsupportedUnaryExpr(xf interface{}, op token.Token) error {
	return errors.New(fmt.Sprintf("unsupported unary expression %s on %T: %s %#v\n", op, xf, op, xf))
}

func (ir *Interpreter) evalUnaryExpr(expr *ast.UnaryExpr) (r.Value, error) {
	xv, err := ir.Eval(expr.X)
	if err != nil {
		return Nil, err
	}
	op := expr.Op
	switch xv.Kind() {
	case r.Bool:
		return ir.evalUnaryExprBool(xv.Bool(), op)
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return ir.evalUnaryExprInt(xv.Int(), op)
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		return ir.evalUnaryExprUint(xv.Uint(), op)
	case r.Float32, r.Float64:
		return ir.evalUnaryExprFloat(xv.Float(), op)
	case r.Complex64, r.Complex128:
		return ir.evalUnaryExprComplex(xv.Complex(), op)
	default:
		return Nil, unsupportedUnaryExpr(xv.Interface(), op)
	}
}

func (ir *Interpreter) evalUnaryExprBool(x bool, op token.Token) (r.Value, error) {
	var ret interface{}
	switch op {
	case token.NOT:
		ret = !x
	default:
		return Nil, unsupportedUnaryExpr(x, op)
	}
	return r.ValueOf(ret), nil
}

func (ir *Interpreter) evalUnaryExprUint(x uint64, op token.Token) (r.Value, error) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -int64(x)
	case token.XOR:
		ret = ^x
	default:
		return Nil, unsupportedUnaryExpr(x, op)
	}
	return r.ValueOf(ret), nil
}

func (ir *Interpreter) evalUnaryExprInt(x int64, op token.Token) (r.Value, error) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -x
	case token.XOR:
		ret = ^x
	default:
		return Nil, unsupportedUnaryExpr(x, op)
	}
	return r.ValueOf(ret), nil
}

func (ir *Interpreter) evalUnaryExprFloat(x float64, op token.Token) (r.Value, error) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -x
	default:
		return Nil, unsupportedUnaryExpr(x, op)
	}
	return r.ValueOf(ret), nil
}

func (ir *Interpreter) evalUnaryExprComplex(x complex128, op token.Token) (r.Value, error) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -x
	default:
		return Nil, unsupportedUnaryExpr(x, op)
	}
	return r.ValueOf(ret), nil
}
