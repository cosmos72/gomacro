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
	"go/token"
	r "reflect"

	mt "github.com/cosmos72/gomacro/token"
)

func (env *Env) unsupportedBinaryExpr(xv r.Value, op token.Token, yv r.Value) r.Value {
	opstr := mt.String(op)
	ret, _ := env.Errorf("unsupported binary operation %s between <%v> and <%v>: %v %s %v", opstr, TypeOf(xv), TypeOf(yv), xv, opstr, yv)
	return ret
}

func (env *Env) evalBinaryExpr(xv r.Value, op token.Token, yv r.Value) r.Value {
	switch xv.Kind() {
	case r.Bool:
		switch yv.Kind() {
		case r.Bool:
			return env.evalBinaryExprBoolBool(xv, op, yv)
		}
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		x := xv.Int()
		switch yv.Kind() {
		case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
			return env.evalBinaryExprIntInt(xv, op, yv)
		case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
			return env.evalBinaryExprIntInt(xv, op, r.ValueOf(int64(yv.Uint())))
		case r.Float32, r.Float64:
			xv = r.ValueOf(float64(x)).Convert(yv.Type())
			return env.evalBinaryExprFloat(xv, op, yv)
		case r.Complex64, r.Complex128:
			xv = r.ValueOf(complex(float64(x), 0.0)).Convert(yv.Type())
			return env.evalBinaryExprComplex(xv, op, yv)
		}
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		x := xv.Uint()
		switch yv.Kind() {
		case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
			if yv.Int() < 0 {
				return env.evalBinaryExprIntInt(r.ValueOf(int64(x)), op, yv)
			} else {
				return env.evalBinaryExprUintUint(xv, op, r.ValueOf(uint64(yv.Int())))
			}
		case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
			return env.evalBinaryExprUintUint(xv, op, yv)
		case r.Float32, r.Float64:
			xv = r.ValueOf(float64(x)).Convert(yv.Type())
			return env.evalBinaryExprFloat(xv, op, yv)
		case r.Complex64, r.Complex128:
			xv = r.ValueOf(complex(float64(x), 0.0)).Convert(yv.Type())
			return env.evalBinaryExprComplex(xv, op, yv)
		}
	case r.Float32, r.Float64:
		return env.evalBinaryExprFloat(xv, op, yv)
	case r.Complex64, r.Complex128:
		return env.evalBinaryExprComplex(xv, op, yv)
	case r.String:
		return env.evalBinaryExprString(xv, op, yv)
	}
	return env.unsupportedBinaryExpr(xv, op, yv)
}

func (env *Env) evalBinaryExprBoolBool(xv r.Value, op token.Token, yv r.Value) r.Value {
	x := xv.Bool()
	y := yv.Bool()
	var b bool
	switch op {
	case token.LAND: // for a short-circuit implementation, see evalExpr
		b = x && y
	case token.LOR: // for a short-circuit implementation, see evalExpr
		b = x || y
	case token.EQL:
		b = x == y
	case token.NEQ:
		b = x != y
	default:
		env.unsupportedBinaryExpr(xv, op, yv)
	}
	return r.ValueOf(b)
}

func (env *Env) evalBinaryExprIntInt(xv r.Value, op token.Token, yv r.Value) r.Value {
	x := xv.Int()
	y := yv.Int()
	var ret int64
	switch op {
	case token.ADD, token.ADD_ASSIGN:
		ret = x + y
	case token.SUB, token.SUB_ASSIGN:
		ret = x - y
	case token.MUL, token.MUL_ASSIGN:
		ret = x * y
	case token.QUO, token.QUO_ASSIGN:
		ret = x / y
	case token.REM, token.REM_ASSIGN:
		ret = x % y
	case token.AND, token.AND_ASSIGN:
		ret = x & y
	case token.OR, token.OR_ASSIGN:
		ret = x | y
	case token.XOR, token.XOR_ASSIGN:
		ret = x ^ y
	case token.SHL, token.SHL_ASSIGN:
		// in Go, x << y and x >> y require y to be unsigned
		ret = x << uint64(y)
	case token.SHR, token.SHR_ASSIGN:
		ret = x >> uint64(y)
	case token.AND_NOT, token.AND_NOT_ASSIGN:
		ret = x &^ y
	default:
		goto PART2
	}
	return env.valueToType(r.ValueOf(ret), xv.Type())

PART2:
	var b bool
	switch op {
	case token.EQL:
		b = x == y
	case token.LSS:
		b = x < y
	case token.GTR:
		b = x > y
	case token.NEQ:
		b = x != y
	case token.LEQ:
		b = x <= y
	case token.GEQ:
		b = x >= y
	default:
		return env.unsupportedBinaryExpr(r.ValueOf(x), op, r.ValueOf(y))
	}
	return r.ValueOf(b)
}

func (env *Env) evalBinaryExprUintUint(xv r.Value, op token.Token, yv r.Value) r.Value {
	x := xv.Uint()
	y := yv.Uint()
	var ret uint64

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
	default:
		goto PART2
	}
	return env.valueToType(r.ValueOf(ret), xv.Type())

PART2:
	var b bool
	switch op {
	case token.EQL:
		b = x == y
	case token.LSS:
		b = x < y
	case token.GTR:
		b = x > y
	case token.NEQ:
		b = x != y
	case token.LEQ:
		b = x <= y
	case token.GEQ:
		b = x >= y
	default:
		return env.unsupportedBinaryExpr(xv, op, yv)
	}
	return r.ValueOf(b)
}

func (env *Env) evalBinaryExprFloat(xv r.Value, op token.Token, yv r.Value) r.Value {
	x := xv.Float()
	y, ok := env.toFloat(yv)
	if ok {
		var ret float64
		switch op {
		case token.ADD, token.ADD_ASSIGN:
			ret = x + y
		case token.SUB, token.SUB_ASSIGN:
			ret = x - y
		case token.MUL, token.MUL_ASSIGN:
			ret = x * y
		case token.QUO, token.QUO_ASSIGN:
			ret = x / y
		default:
			goto PART2
		}
		if xv.Kind() == r.Float32 {
			return r.ValueOf(float32(ret))
		}
		return r.ValueOf(ret)
	PART2:
		var b bool
		switch op {
		case token.EQL:
			b = x == y
		case token.LSS:
			b = x < y
		case token.GTR:
			b = x > y
		case token.NEQ:
			b = x != y
		case token.LEQ:
			b = x <= y
		case token.GEQ:
			b = x >= y
		default:
			return env.unsupportedBinaryExpr(xv, op, yv)
		}
		return r.ValueOf(b)
	}
	if yv.Kind() == r.Complex64 || yv.Kind() == r.Complex128 {
		xv = r.ValueOf(complex(x, 0.0)).Convert(yv.Type())
		return env.evalBinaryExprComplex(xv, op, yv)
	}
	return env.unsupportedBinaryExpr(xv, op, yv)
}

func (env *Env) evalBinaryExprComplex(xv r.Value, op token.Token, yv r.Value) r.Value {
	x := xv.Complex()
	y, ok := env.toComplex(yv)
	if ok {
		var ret complex128
		switch op {
		case token.ADD, token.ADD_ASSIGN:
			ret = x + y
		case token.SUB, token.SUB_ASSIGN:
			ret = x - y
		case token.MUL, token.MUL_ASSIGN:
			ret = x * y
		case token.QUO, token.QUO_ASSIGN:
			ret = x / y
		default:
			goto PART2
		}
		if xv.Kind() == r.Complex64 {
			return r.ValueOf(complex64(ret))
		}
		return r.ValueOf(ret)
	PART2:
		var b bool
		switch op {
		case token.EQL:
			b = x == y
		case token.NEQ:
			b = x != y
		default:
			return env.unsupportedBinaryExpr(xv, op, yv)
		}
		return r.ValueOf(b)
	}
	return env.unsupportedBinaryExpr(xv, op, yv)
}

func (env *Env) evalBinaryExprString(xv r.Value, op token.Token, yv r.Value) r.Value {
	if xv.Kind() == r.String && yv.Kind() == r.String && op == token.ADD {
		return r.ValueOf(xv.String() + yv.String())
	}
	return env.unsupportedBinaryExpr(xv, op, yv)
}
