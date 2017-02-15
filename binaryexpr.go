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
 * binaryexpr.go
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

func unsupportedBinaryExpr(xf interface{}, op token.Token, yf interface{}) error {
	return errors.New(fmt.Sprintf("unsupported binary operation %s between %T and %T: %#v %s %#v\n", op, xf, yf, xf, op, yf))
}

func (ir *Interpreter) evalBinaryExpr(expr *ast.BinaryExpr) (r.Value, error) {
	xv, err := ir.Eval(expr.X)
	if err != nil {
		return Nil, err
	}
	yv, err := ir.Eval(expr.Y)
	if err != nil {
		return Nil, err
	}
	op := expr.Op
	switch xv.Kind() {
	case r.Bool:
		return ir.evalBinaryExprBool(xv.Bool(), op, yv)
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return ir.evalBinaryExprInt(xv.Int(), op, yv)
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		return ir.evalBinaryExprUint(xv.Uint(), op, yv)
	case r.Float32, r.Float64:
		return ir.evalBinaryExprFloat(xv.Float(), op, yv)
	case r.Complex64, r.Complex128:
		return ir.evalBinaryExprComplex(xv.Complex(), op, yv)
	default:
		// TODO string
		return Nil, unsupportedBinaryExpr(xv.Interface(), op, yv.Interface())
	}
}

func (ir *Interpreter) evalBinaryExprBool(x bool, op token.Token, yv r.Value) (r.Value, error) {
	var ret bool
	switch yv.Kind() {
	case r.Bool:
		y := yv.Bool()
		switch op {
		case token.EQL:
			ret = x == y
		case token.NEQ:
			ret = x != y
		default:
			return Nil, unsupportedBinaryExpr(x, op, y)
		}
		return r.ValueOf(ret), nil
	default:
		return Nil, unsupportedBinaryExpr(x, op, yv.Interface())
	}
}

func (ir *Interpreter) evalBinaryExprUint(x uint64, op token.Token, yv r.Value) (r.Value, error) {
	switch yv.Kind() {

	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		var ret interface{}
		y := yv.Uint()

		switch op {
		case token.ADD:
			ret = x + y
		case token.SUB:
			ret = x - y
		case token.MUL:
			ret = x * y
		case token.QUO:
			ret = x / y
		case token.REM:
			ret = x % y
		case token.AND:
			ret = x & y
		case token.OR:
			ret = x | y
		case token.XOR:
			ret = x ^ y
		case token.SHL:
			ret = x << y
		case token.SHR:
			ret = x >> y
		case token.AND_NOT:
			ret = x &^ y
		case token.EQL:
			ret = x == y
		case token.LSS:
			ret = x < y
		case token.GTR:
			ret = x < y
		case token.NEQ:
			ret = x != y
		case token.LEQ:
			ret = x <= y
		case token.GEQ:
			ret = x >= y
		default:
			return Nil, unsupportedBinaryExpr(x, op, yv.Interface())
		}
		return r.ValueOf(ret), nil

	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return ir.evalBinaryExprIntInt(int64(x), op, yv.Int())

	case r.Float32, r.Float64:
		return ir.evalBinaryExprFloat(float64(x), op, yv)

	case r.Complex64, r.Complex128:
		return ir.evalBinaryExprComplex(complex(float64(x), 0.0), op, yv)

	default:
		return Nil, unsupportedBinaryExpr(x, op, yv.Interface())
	}
}

func (ir *Interpreter) evalBinaryExprInt(x int64, op token.Token, yv r.Value) (r.Value, error) {
	switch yv.Kind() {
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		return ir.evalBinaryExprIntInt(x, op, int64(yv.Uint()))
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return ir.evalBinaryExprIntInt(x, op, yv.Int())
	case r.Float32, r.Float64:
		return ir.evalBinaryExprFloat(float64(x), op, yv)
	case r.Complex64, r.Complex128:
		return ir.evalBinaryExprComplex(complex(float64(x), 0.0), op, yv)
	default:
		return Nil, unsupportedBinaryExpr(x, op, yv.Interface())
	}
}

func (ir *Interpreter) evalBinaryExprIntInt(x int64, op token.Token, y int64) (r.Value, error) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x + y
	case token.SUB:
		ret = x - y
	case token.MUL:
		ret = x * y
	case token.QUO:
		ret = x / y
	case token.REM:
		ret = x % y
	case token.AND:
		ret = x & y
	case token.OR:
		ret = x | y
	case token.XOR:
		ret = x ^ y
	case token.SHL:
		// in Go, x << y and x >> y require y to be unsigned
		ret = x << uint64(y)
	case token.SHR:
		ret = x >> uint64(y)
	case token.AND_NOT:
		ret = x &^ y
	case token.EQL:
		ret = x == y
	case token.LSS:
		ret = x < y
	case token.GTR:
		ret = x < y
	case token.NEQ:
		ret = x != y
	case token.LEQ:
		ret = x <= y
	case token.GEQ:
		ret = x >= y
	default:
		return Nil, unsupportedBinaryExpr(x, op, y)
	}
	return r.ValueOf(ret), nil
}

func (ir *Interpreter) evalBinaryExprFloat(x float64, op token.Token, yv r.Value) (r.Value, error) {
	var ret interface{}
	if y, ok := toFloat(yv); ok {
		switch op {
		case token.ADD:
			ret = x + y
		case token.SUB:
			ret = x - y
		case token.MUL:
			ret = x * y
		case token.QUO:
			ret = x / y
		case token.EQL:
			ret = x == y
		case token.LSS:
			ret = x < y
		case token.GTR:
			ret = x < y
		case token.NEQ:
			ret = x != y
		case token.LEQ:
			ret = x <= y
		case token.GEQ:
			ret = x >= y
		default:
			return Nil, unsupportedBinaryExpr(x, op, yv.Interface())
		}
		return r.ValueOf(ret), nil
	}
	if k := yv.Kind(); k == r.Complex64 || k == r.Complex128 {
		return ir.evalBinaryExprComplex(complex(x, 0.0), op, yv)
	}
	return Nil, unsupportedBinaryExpr(x, op, yv.Interface())
}

func (ir *Interpreter) evalBinaryExprComplex(x complex128, op token.Token, yv r.Value) (r.Value, error) {
	var ret interface{}
	if y, ok := toComplex(yv); ok {
		switch op {
		case token.ADD:
			ret = x + y
		case token.SUB:
			ret = x - y
		case token.MUL:
			ret = x * y
		case token.QUO:
			ret = x / y
		case token.EQL:
			ret = x == y
		case token.NEQ:
			ret = x != y
		default:
			return Nil, unsupportedBinaryExpr(x, op, yv.Interface())
		}
		return r.ValueOf(ret), nil
	}
	return Nil, unsupportedBinaryExpr(x, op, yv.Interface())
}
