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

	mt "github.com/cosmos72/gomacro/token"
)

func (env *Env) unsupportedUnaryExpr(xf interface{}, op token.Token) (r.Value, []r.Value) {
	opstr := mt.String(op)
	return env.Errorf("unsupported unary expression %s on %T: %s %#v", opstr, xf, opstr, xf)
}

func (env *Env) evalUnaryExpr(expr *ast.UnaryExpr) (r.Value, []r.Value) {
	switch op := expr.Op; op {
	case mt.MACRO, mt.QUOTE, mt.QUASIQUOTE:
		// the various *QUOTE* special forms and the result of macroexpansion
		// are statements wrapped in a closure
		block := expr.X.(*ast.FuncLit).Body
		var ret r.Value
		var rets []r.Value
		switch op {
		case mt.MACRO:
			ret, rets = env.evalBlock(block)
		case mt.QUOTE:
			node := env.evalQuote(block)
			ret = r.ValueOf(&node).Elem()
		case mt.QUASIQUOTE:
			node := env.evalQuasiquote(block)
			ret = r.ValueOf(&node).Elem()
		}
		return ret, rets

	case mt.UNQUOTE, mt.UNQUOTE_SPLICE:
		return env.Errorf("%s outside quasiquote", mt.String(op))

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
