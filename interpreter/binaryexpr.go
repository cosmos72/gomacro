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
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
	r "reflect"

	mt "github.com/cosmos72/gomacro/token"
)

func (env *Env) unsupportedBinaryExpr(xv r.Value, op token.Token, yv r.Value) (r.Value, []r.Value) {
	opstr := mt.String(op)
	return env.Errorf("unsupported binary operation %s between %v <%v> and %v <%v>: %v %s %v", opstr, xv.Kind(), xv.Type(), yv.Kind(), yv.Type(), xv, opstr, yv)
}

func (env *Env) evalBinaryExpr(expr *ast.BinaryExpr) (r.Value, []r.Value) {
	xv := env.Eval1(expr.X)
	yv := env.Eval1(expr.Y)
	op := expr.Op
	switch xv.Kind() {
	case r.Bool:
		return env.evalBinaryExprBool(xv.Bool(), op, yv)
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return env.evalBinaryExprInt(xv.Int(), op, yv)
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		return env.evalBinaryExprUint(xv.Uint(), op, yv)
	case r.Float32, r.Float64:
		return env.evalBinaryExprFloat(xv.Float(), op, yv)
	case r.Complex64, r.Complex128:
		return env.evalBinaryExprComplex(xv.Complex(), op, yv)
	case r.String:
		return env.evalBinaryExprString(xv.String(), op, yv)
	default:
		return env.unsupportedBinaryExpr(xv, op, yv)
	}
}

func (env *Env) evalBinaryExprBool(x bool, op token.Token, yv r.Value) (r.Value, []r.Value) {
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
			env.unsupportedBinaryExpr(r.ValueOf(x), op, yv)
		}
		return r.ValueOf(ret), nil
	default:
		return env.unsupportedBinaryExpr(r.ValueOf(x), op, yv)
	}
}

func (env *Env) evalBinaryExprUint(x uint64, op token.Token, yv r.Value) (r.Value, []r.Value) {
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
			ret = x > y
		case token.NEQ:
			ret = x != y
		case token.LEQ:
			ret = x <= y
		case token.GEQ:
			ret = x >= y
		default:
			return env.unsupportedBinaryExpr(r.ValueOf(x), op, yv)
		}
		return r.ValueOf(ret), nil

	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return env.evalBinaryExprIntInt(int64(x), op, yv.Int())

	case r.Float32, r.Float64:
		return env.evalBinaryExprFloat(float64(x), op, yv)

	case r.Complex64, r.Complex128:
		return env.evalBinaryExprComplex(complex(float64(x), 0.0), op, yv)

	default:
		return env.unsupportedBinaryExpr(r.ValueOf(x), op, yv)
	}
}

func (env *Env) evalBinaryExprInt(x int64, op token.Token, yv r.Value) (r.Value, []r.Value) {
	switch yv.Kind() {
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		return env.evalBinaryExprIntInt(x, op, int64(yv.Uint()))
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return env.evalBinaryExprIntInt(x, op, yv.Int())
	case r.Float32, r.Float64:
		return env.evalBinaryExprFloat(float64(x), op, yv)
	case r.Complex64, r.Complex128:
		return env.evalBinaryExprComplex(complex(float64(x), 0.0), op, yv)
	default:
		return env.unsupportedBinaryExpr(r.ValueOf(x), op, yv)
	}
}

func (env *Env) evalBinaryExprIntInt(x int64, op token.Token, y int64) (r.Value, []r.Value) {
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
		ret = x > y
	case token.NEQ:
		ret = x != y
	case token.LEQ:
		ret = x <= y
	case token.GEQ:
		ret = x >= y
	default:
		return env.unsupportedBinaryExpr(r.ValueOf(x), op, r.ValueOf(y))
	}
	return r.ValueOf(ret), nil
}

func (env *Env) evalBinaryExprFloat(x float64, op token.Token, yv r.Value) (r.Value, []r.Value) {
	var ret interface{}
	if y, ok := env.toFloat(yv); ok {
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
			ret = x > y
		case token.NEQ:
			ret = x != y
		case token.LEQ:
			ret = x <= y
		case token.GEQ:
			ret = x >= y
		default:
			return env.unsupportedBinaryExpr(r.ValueOf(x), op, yv)
		}
		return r.ValueOf(ret), nil
	}
	if k := yv.Kind(); k == r.Complex64 || k == r.Complex128 {
		return env.evalBinaryExprComplex(complex(x, 0.0), op, yv)
	}
	return env.unsupportedBinaryExpr(r.ValueOf(x), op, yv)
}

func (env *Env) evalBinaryExprComplex(x complex128, op token.Token, yv r.Value) (r.Value, []r.Value) {
	var ret interface{}
	if y, ok := env.toComplex(yv); ok {
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
			return env.unsupportedBinaryExpr(r.ValueOf(x), op, yv)
		}
		return r.ValueOf(ret), nil
	}
	return env.unsupportedBinaryExpr(r.ValueOf(x), op, yv)
}

func (env *Env) evalBinaryExprString(x string, op token.Token, yv r.Value) (r.Value, []r.Value) {
	if yv.Kind() == r.String && op == token.ADD {
		return r.ValueOf(x + yv.String()), nil
	}
	return env.unsupportedBinaryExpr(r.ValueOf(x), op, yv)
}
