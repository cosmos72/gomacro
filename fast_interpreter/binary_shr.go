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
 * binary_shr.go
 *
 *  Created on Apr 08, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
)

func (c *Comp) Shr(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
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
			fun = shrXYInt(x, y)
		case func(*Env) int8:
			fun = shrXYInt8(x, y)
		case func(*Env) int16:
			fun = shrXYInt16(x, y)
		case func(*Env) int32:
			fun = shrXYInt32(x, y)
		case func(*Env) int64:
			fun = shrXYInt64(x, y)
		case func(*Env) uint:
			fun = shrXYUint(x, y)
		case func(*Env) uint8:
			fun = shrXYUint8(x, y)
		case func(*Env) uint16:
			fun = shrXYUint16(x, y)
		case func(*Env) uint32:
			fun = shrXYUint32(x, y)
		case func(*Env) uint64:
			fun = shrXYUint64(x, y)
		case func(*Env) uintptr:
			fun = shrXYUintptr(x, y)
		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 0) {
			return xe
		}
		switch x := x.(type) {
		case func(*Env) int:
			fun = shrXInt(x, y)
		case func(*Env) int8:
			fun = shrXInt8(x, y)
		case func(*Env) int16:
			fun = shrXInt16(x, y)
		case func(*Env) int32:
			fun = shrXInt32(x, y)
		case func(*Env) int64:
			fun = shrXInt64(x, y)
		case func(*Env) uint:
			fun = shrXUint(x, y)
		case func(*Env) uint8:
			fun = shrXUint8(x, y)
		case func(*Env) uint16:
			fun = shrXUint16(x, y)
		case func(*Env) uint32:
			fun = shrXUint32(x, y)
		case func(*Env) uint64:
			fun = shrXUint64(x, y)
		case func(*Env) uintptr:
			fun = shrXUintptr(x, y)
		}
	} else {
		x := xe.Value
		y := ye.Fun
		// cannot optimize 0 << y to 0 because y may have side effects
		switch x := x.(type) {
		case int:
			fun = shrYInt(x, y)
		case int8:
			fun = shrYInt8(x, y)
		case int16:
			fun = shrYInt16(x, y)
		case int32:
			fun = shrYInt32(x, y)
		case int64:
			fun = shrYInt64(x, y)
		case uint:
			fun = shrYUint(x, y)
		case uint8:
			fun = shrYUint8(x, y)
		case uint16:
			fun = shrYUint16(x, y)
		case uint32:
			fun = shrYUint32(x, y)
		case uint64:
			fun = shrYUint64(x, y)
		case uintptr:
			fun = shrYUintptr(x, y)
		}
	}
	if fun == nil {
		return c.invalidBinaryExpr(node, xe, ye)
	}
	return ExprFun(xe.Type, fun)
}
