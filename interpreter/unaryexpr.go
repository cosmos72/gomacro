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
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	mt "github.com/cosmos72/gomacro/token"
)

func (env *Env) unsupportedUnaryExpr(xv r.Value, op token.Token) (r.Value, []r.Value) {
	opstr := mt.String(op)
	return env.Errorf("unsupported unary expression %s on <%v>: %s %v", opstr, typeOf(xv), opstr, xv)
}

func (env *Env) warnOverflowSignedMinus(x interface{}, ret interface{}) {
	str := fmt.Sprintf("%d", x)
	if len(str) > 0 && str[0] == '-' {
		str = str[1:]
	}
	env.Warnf("value %s overflows <%v>, result truncated to %d", str, r.TypeOf(x), ret)
}

func (env *Env) warnUnderflowUnsignedMinus(x interface{}, ret interface{}) {
	env.Warnf("value -%d underflows <%v>, result truncated to %d", x, r.TypeOf(x), ret)
}

func (env *Env) evalUnaryExpr(node *ast.UnaryExpr) (r.Value, []r.Value) {
	op := node.Op
	switch op {
	case token.AND:
		place := env.evalExpr1(node.X)
		if place == Nil || !place.CanAddr() {
			return env.Errorf("cannot take the address of: %v = %v <%v>", node.X, place, typeOf(place))
		}
		return place.Addr(), nil

	// the various QUOTE special forms, the result of macroexpansion,
	// and our extension "block statement inside expression" are:
	// a block statements, wrapped in a closure, wrapped in a unary expression "MACRO", i.e.:
	// MACRO func() { /*block*/ }
	case mt.MACRO:
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
	if xv == Nil || xv == None {
		return env.unsupportedUnaryExpr(xv, op)
	}
	var ret interface{}

	switch x := xv.Interface().(type) {
	case bool:
		if op == token.NOT {
			ret = !x
		}
	case int:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
			if x == -x {
				env.warnOverflowSignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case int8:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
			if x == -x {
				env.warnOverflowSignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case int16:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
			if x == -x {
				env.warnOverflowSignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case int32:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
			if x == -x {
				env.warnOverflowSignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case int64:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
			if x == -x {
				env.warnOverflowSignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case uint:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
			if x != 0 {
				env.warnUnderflowUnsignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case uint8:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
			if x != 0 {
				env.warnUnderflowUnsignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case uint16:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
			if x != 0 {
				env.warnUnderflowUnsignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case uint32:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
			if x != 0 {
				env.warnUnderflowUnsignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case uint64:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
			if x != 0 {
				env.warnUnderflowUnsignedMinus(x, ret)
			}
		case token.XOR:
			ret = ^x
		}
	case uintptr:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
			env.warnUnderflowUnsignedMinus(x, ret)
		case token.XOR:
			ret = ^x
		}
	case float32:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
		}
	case float64:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
		}
	case complex64:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
		}
	case complex128:
		switch op {
		case token.ADD:
			ret = x
		case token.SUB:
			ret = -x
		}
	default:
		if op == token.ARROW && xv.Kind() == r.Chan {
			ret, ok := xv.Recv()
			return ret, []r.Value{ret, r.ValueOf(ok)}
		}
	}
	if ret == nil {
		return env.unsupportedUnaryExpr(xv, op)
	}
	return r.ValueOf(ret), nil
}
