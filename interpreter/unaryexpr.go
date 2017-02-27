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

func (env *Env) evalUnaryExpr(node *ast.UnaryExpr) (r.Value, []r.Value) {
	op := node.Op
	switch op {
	case mt.MACRO:
		// the various QUOTE special forms and the result of macroexpansion
		// are statements wrapped in a closure
		block := node.X.(*ast.FuncLit).Body
		return env.evalBlock(block)

	case mt.QUOTE:
		block := node.X.(*ast.FuncLit).Body
		ret := env.evalQuote(block)
		return r.ValueOf(ret), nil

	case mt.QUASIQUOTE:
		block := node.X.(*ast.FuncLit).Body
		ret := env.evalQuasiquote(block)
		return r.ValueOf(ret), nil

	case mt.UNQUOTE, mt.UNQUOTE_SPLICE:
		return env.Errorf("%s not inside quasiquote: %v <%v>", mt.String(op), node, r.TypeOf(node))
	}
	xv, _ := env.Eval(node.X)
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
	var ret uint64
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		iret := -int64(x)
		iret2 := int(iret)
		if int64(iret2) == iret {
			return r.ValueOf(iret2), nil
		} else {
			return r.ValueOf(iret), nil
		}
	case token.XOR:
		ret = ^x
	default:
		return env.unsupportedUnaryExpr(x, op)
	}
	ret2 := uint(ret)
	if uint64(ret2) == ret {
		return r.ValueOf(ret2), nil
	} else {
		return r.ValueOf(ret), nil
	}
}

func (env *Env) evalUnaryExprInt(x int64, op token.Token) (r.Value, []r.Value) {
	var ret int64
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
	ret2 := int(ret)
	if int64(ret2) == ret {
		return r.ValueOf(ret2), nil
	} else {
		return r.ValueOf(ret), nil
	}
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
