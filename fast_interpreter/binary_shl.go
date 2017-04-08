/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * binary_shl.go
 *
 *  Created on Apr 08, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
)

func (c *Comp) Shl(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	if ze := c.prepareShift(node, xe, ye); ze != nil {
		return ze
	}
	xc, yc := xe.Const(), ye.Const()
	// if both x and y are constants, BinaryExpr will invoke EvalConst()
	// on our return value. no need to optimize that.
	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun
		switch x := x.(type) {
		case func(*Env) int:
			fun = shlXYInt(x, y)
		case func(*Env) int8:
			fun = shlXYInt8(x, y)
		case func(*Env) int16:
			fun = shlXYInt16(x, y)
		case func(*Env) int32:
			fun = shlXYInt32(x, y)
		case func(*Env) int64:
			fun = shlXYInt64(x, y)
		case func(*Env) uint:
			fun = shlXYUint(x, y)
		case func(*Env) uint8:
			fun = shlXYUint8(x, y)
		case func(*Env) uint16:
			fun = shlXYUint16(x, y)
		case func(*Env) uint32:
			fun = shlXYUint32(x, y)
		case func(*Env) uint64:
			fun = shlXYUint64(x, y)
		case func(*Env) uintptr:
			fun = shlXYUintptr(x, y)
		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 0) {
			return xe
		}
		switch x := x.(type) {
		case func(*Env) int:
			fun = shlXInt(x, y)
		case func(*Env) int8:
			fun = shlXInt8(x, y)
		case func(*Env) int16:
			fun = shlXInt16(x, y)
		case func(*Env) int32:
			fun = shlXInt32(x, y)
		case func(*Env) int64:
			fun = shlXInt64(x, y)
		case func(*Env) uint:
			fun = shlXUint(x, y)
		case func(*Env) uint8:
			fun = shlXUint8(x, y)
		case func(*Env) uint16:
			fun = shlXUint16(x, y)
		case func(*Env) uint32:
			fun = shlXUint32(x, y)
		case func(*Env) uint64:
			fun = shlXUint64(x, y)
		case func(*Env) uintptr:
			fun = shlXUintptr(x, y)
		}
	} else {
		x := xe.Value
		y := ye.Fun
		// cannot optimize 0 << y to 0 because y may have side effects
		switch x := x.(type) {
		case int:
			fun = shlYInt(x, y)
		case int8:
			fun = shlYInt8(x, y)
		case int16:
			fun = shlYInt16(x, y)
		case int32:
			fun = shlYInt32(x, y)
		case int64:
			fun = shlYInt64(x, y)
		case uint:
			fun = shlYUint(x, y)
		case uint8:
			fun = shlYUint8(x, y)
		case uint16:
			fun = shlYUint16(x, y)
		case uint32:
			fun = shlYUint32(x, y)
		case uint64:
			fun = shlYUint64(x, y)
		case uintptr:
			fun = shlYUintptr(x, y)
		}
	}
	if fun == nil {
		return c.invalidBinaryExpr(node, xe, ye)
	}
	return ExprFun(xe.Type, fun)
}
