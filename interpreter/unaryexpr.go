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
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
	r "reflect"

	mp "github.com/cosmos72/gomacro/macroparser"
)

func (env *Env) unsupportedUnaryExpr(xf interface{}, op token.Token) (r.Value, []r.Value) {
	opstr := mp.TokenString(op)
	return env.Errorf("unsupported unary expression %s on %T: %s %#v", opstr, xf, opstr, xf)
}

func (env *Env) evalUnaryExpr(expr *ast.UnaryExpr) (r.Value, []r.Value) {
	switch op := expr.Op; op {
	case mp.MACRO:
		// result of macroexpansion: a statement wrapped in a closure
		return env.evalBlock(expr.X.(*ast.FuncLit).Body)

	case mp.QUOTE, mp.QUASIQUOTE, mp.UNQUOTE, mp.UNQUOTE_SPLICE:
		// result of quote and friends: a statement wrapped in a closure
		return env.evalQuote(op, expr.X.(*ast.FuncLit).Body)

	default:
		break
	}
	xv, _ := env.Eval(expr.X)
	op := expr.Op
	switch xv.Kind() {
	case r.Bool:
		return env.evalUnaryExprBool(xv.Bool(), op)
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return env.evalUnaryExprInt(xv.Int(), op)
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		return env.evalUnaryExprUint(xv.Uint(), op)
	case r.Float32, r.Float64:
		return env.evalUnaryExprFloat(xv.Float(), op)
	case r.Complex64, r.Complex128:
		return env.evalUnaryExprComplex(xv.Complex(), op)
	default:
		return env.unsupportedUnaryExpr(xv.Interface(), op)
	}
}

func (env *Env) evalUnaryExprBool(x bool, op token.Token) (r.Value, []r.Value) {
	var ret interface{}
	switch op {
	case token.NOT:
		ret = !x
	default:
		return env.unsupportedUnaryExpr(x, op)
	}
	return r.ValueOf(ret), nil
}

func (env *Env) evalUnaryExprUint(x uint64, op token.Token) (r.Value, []r.Value) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -int64(x)
	case token.XOR:
		ret = ^x
	default:
		return env.unsupportedUnaryExpr(x, op)
	}
	return r.ValueOf(ret), nil
}

func (env *Env) evalUnaryExprInt(x int64, op token.Token) (r.Value, []r.Value) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -x
	case token.XOR:
		ret = ^x
	default:
		return env.unsupportedUnaryExpr(x, op)
	}
	return r.ValueOf(ret), nil
}

func (env *Env) evalUnaryExprFloat(x float64, op token.Token) (r.Value, []r.Value) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -x
	default:
		return env.unsupportedUnaryExpr(x, op)
	}
	return r.ValueOf(ret), nil
}

func (env *Env) evalUnaryExprComplex(x complex128, op token.Token) (r.Value, []r.Value) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -x
	default:
		return env.unsupportedUnaryExpr(x, op)
	}
	return r.ValueOf(ret), nil
}
