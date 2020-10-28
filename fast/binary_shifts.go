// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * binary_shifts.go
 *
 *  Created on Apr 08, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"

	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) Shl(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	if ze := c.prepareShift(node, xe, ye); ze != nil {
		return ze
	}

	xc, yc := xe.Const(), ye.Const()
	xk := xe.Type.Kind()

	var fun I
	if xc == yc {
		x := xe.Fun
		y := ye.AsUint64()

		switch xk {
		case xr.Int:
			x := x.(func(*Env) int)

			fun = func(env *Env) int {
				return x(env) << y(env)
			}

		case xr.Int8:
			x := x.(func(*Env) int8)

			fun = func(env *Env) int8 {
				return x(env) << y(env)
			}

		case xr.Int16:
			x := x.(func(*Env) int16)

			fun = func(env *Env) int16 {
				return x(env) << y(env)
			}

		case xr.Int32:
			x := x.(func(*Env) int32)

			fun = func(env *Env) int32 {
				return x(env) << y(env)
			}

		case xr.Int64:
			x := x.(func(*Env) int64)

			fun = func(env *Env) int64 {
				return x(env) << y(env)
			}

		case xr.Uint:
			x := x.(func(*Env) uint)

			fun = func(env *Env) uint {
				return x(env) << y(env)
			}

		case xr.Uint8:
			x := x.(func(*Env) uint8)

			fun = func(env *Env) uint8 {
				return x(env) << y(env)
			}

		case xr.Uint16:
			x := x.(func(*Env) uint16)

			fun = func(env *Env) uint16 {
				return x(env) << y(env)
			}

		case xr.Uint32:
			x := x.(func(*Env) uint32)

			fun = func(env *Env) uint32 {
				return x(env) << y(env)
			}

		case xr.Uint64:
			x := x.(func(*Env) uint64)

			fun = func(env *Env) uint64 {
				return x(env) << y(env)
			}

		case xr.Uintptr:
			x := x.(func(*Env) uintptr)

			fun = func(env *Env) uintptr {
				return x(env) << y(env)
			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y, ok := constAsUint64(ye.Value)
		if !ok {
			c.invalidBinaryExpr(node, xe, ye)
		} else if y == 0 {
			return xe
		}

		switch xk {
		case xr.Int:
			x := x.(func(*Env) int)

			fun = func(env *Env) int {
				return x(env) << y
			}

		case xr.Int8:
			x := x.(func(*Env) int8)

			fun = func(env *Env) int8 {
				return x(env) << y
			}

		case xr.Int16:
			x := x.(func(*Env) int16)

			fun = func(env *Env) int16 {
				return x(env) << y
			}

		case xr.Int32:
			x := x.(func(*Env) int32)

			fun = func(env *Env) int32 {
				return x(env) << y
			}

		case xr.Int64:
			x := x.(func(*Env) int64)

			fun = func(env *Env) int64 {
				return x(env) << y
			}

		case xr.Uint:
			x := x.(func(*Env) uint)

			fun = func(env *Env) uint {
				return x(env) << y
			}

		case xr.Uint8:
			x := x.(func(*Env) uint8)

			fun = func(env *Env) uint8 {
				return x(env) << y
			}

		case xr.Uint16:
			x := x.(func(*Env) uint16)

			fun = func(env *Env) uint16 {
				return x(env) << y
			}

		case xr.Uint32:
			x := x.(func(*Env) uint32)

			fun = func(env *Env) uint32 {
				return x(env) << y
			}

		case xr.Uint64:
			x := x.(func(*Env) uint64)

			fun = func(env *Env) uint64 {
				return x(env) << y
			}

		case xr.Uintptr:
			x := x.(func(*Env) uintptr)

			fun = func(env *Env) uintptr {
				return x(env) << y
			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		xv := xr.ValueOf(xe.Value)
		y := ye.AsUint64()

		switch xk {
		case xr.Int:
			x :=

				int(xv.Int())
			fun = func(env *Env) int {
				return x << y(env)
			}
		case xr.Int8:
			x :=

				int8(xv.Int())
			fun = func(env *Env) int8 {
				return x << y(env)
			}
		case xr.Int16:
			x :=

				int16(xv.Int())
			fun = func(env *Env) int16 {
				return x << y(env)
			}
		case xr.Int32:
			x :=

				int32(xv.Int())
			fun = func(env *Env) int32 {
				return x << y(env)
			}
		case xr.Int64:
			x := xv.Int()
			fun = func(env *Env) int64 {
				return x << y(env)
			}

		case xr.Uint:
			x :=

				uint(xv.Uint())
			fun = func(env *Env) uint {
				return x << y(env)
			}

		case xr.Uint8:
			x :=

				uint8(xv.Uint())
			fun = func(env *Env) uint8 {
				return x << y(env)
			}

		case xr.Uint16:
			x :=

				uint16(xv.Uint())
			fun = func(env *Env) uint16 {
				return x << y(env)
			}

		case xr.Uint32:
			x :=

				uint32(xv.Uint())
			fun = func(env *Env) uint32 {
				return x << y(env)
			}

		case xr.Uint64:
			x := xv.Uint()
			fun = func(env *Env) uint64 {
				return x << y(env)
			}

		case xr.Uintptr:
			x :=

				uintptr(xv.Uint())
			fun = func(env *Env) uintptr {
				return x << y(env)
			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
func (c *Comp) Shr(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	if ze := c.prepareShift(node, xe, ye); ze != nil {
		return ze
	}

	xc, yc := xe.Const(), ye.Const()
	xk := xe.Type.Kind()

	var fun I
	if xc == yc {
		x := xe.Fun
		y := ye.AsUint64()

		switch xk {
		case xr.Int:
			x := x.(func(*Env) int)

			fun = func(env *Env) int {
				return x(env) >> y(env)
			}

		case xr.Int8:
			x := x.(func(*Env) int8)

			fun = func(env *Env) int8 {
				return x(env) >> y(env)
			}

		case xr.Int16:
			x := x.(func(*Env) int16)

			fun = func(env *Env) int16 {
				return x(env) >> y(env)
			}

		case xr.Int32:
			x := x.(func(*Env) int32)

			fun = func(env *Env) int32 {
				return x(env) >> y(env)
			}

		case xr.Int64:
			x := x.(func(*Env) int64)

			fun = func(env *Env) int64 {
				return x(env) >> y(env)
			}

		case xr.Uint:
			x := x.(func(*Env) uint)

			fun = func(env *Env) uint {
				return x(env) >> y(env)
			}

		case xr.Uint8:
			x := x.(func(*Env) uint8)

			fun = func(env *Env) uint8 {
				return x(env) >> y(env)
			}

		case xr.Uint16:
			x := x.(func(*Env) uint16)

			fun = func(env *Env) uint16 {
				return x(env) >> y(env)
			}

		case xr.Uint32:
			x := x.(func(*Env) uint32)

			fun = func(env *Env) uint32 {
				return x(env) >> y(env)
			}

		case xr.Uint64:
			x := x.(func(*Env) uint64)

			fun = func(env *Env) uint64 {
				return x(env) >> y(env)
			}

		case xr.Uintptr:
			x := x.(func(*Env) uintptr)

			fun = func(env *Env) uintptr {
				return x(env) >> y(env)
			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y, ok := constAsUint64(ye.Value)
		if !ok {
			c.invalidBinaryExpr(node, xe, ye)
		} else if y == 0 {
			return xe
		}

		switch xk {
		case xr.Int:
			x := x.(func(*Env) int)

			fun = func(env *Env) int {
				return x(env) >> y
			}

		case xr.Int8:
			x := x.(func(*Env) int8)

			fun = func(env *Env) int8 {
				return x(env) >> y
			}

		case xr.Int16:
			x := x.(func(*Env) int16)

			fun = func(env *Env) int16 {
				return x(env) >> y
			}

		case xr.Int32:
			x := x.(func(*Env) int32)

			fun = func(env *Env) int32 {
				return x(env) >> y
			}

		case xr.Int64:
			x := x.(func(*Env) int64)

			fun = func(env *Env) int64 {
				return x(env) >> y
			}

		case xr.Uint:
			x := x.(func(*Env) uint)

			fun = func(env *Env) uint {
				return x(env) >> y
			}

		case xr.Uint8:
			x := x.(func(*Env) uint8)

			fun = func(env *Env) uint8 {
				return x(env) >> y
			}

		case xr.Uint16:
			x := x.(func(*Env) uint16)

			fun = func(env *Env) uint16 {
				return x(env) >> y
			}

		case xr.Uint32:
			x := x.(func(*Env) uint32)

			fun = func(env *Env) uint32 {
				return x(env) >> y
			}

		case xr.Uint64:
			x := x.(func(*Env) uint64)

			fun = func(env *Env) uint64 {
				return x(env) >> y
			}

		case xr.Uintptr:
			x := x.(func(*Env) uintptr)

			fun = func(env *Env) uintptr {
				return x(env) >> y
			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		xv := xr.ValueOf(xe.Value)
		y := ye.AsUint64()

		switch xk {
		case xr.Int:
			x :=

				int(xv.Int())
			fun = func(env *Env) int {
				return x >> y(env)
			}
		case xr.Int8:
			x :=

				int8(xv.Int())
			fun = func(env *Env) int8 {
				return x >> y(env)
			}
		case xr.Int16:
			x :=

				int16(xv.Int())
			fun = func(env *Env) int16 {
				return x >> y(env)
			}
		case xr.Int32:
			x :=

				int32(xv.Int())
			fun = func(env *Env) int32 {
				return x >> y(env)
			}
		case xr.Int64:
			x := xv.Int()
			fun = func(env *Env) int64 {
				return x >> y(env)
			}

		case xr.Uint:
			x :=

				uint(xv.Uint())
			fun = func(env *Env) uint {
				return x >> y(env)
			}

		case xr.Uint8:
			x :=

				uint8(xv.Uint())
			fun = func(env *Env) uint8 {
				return x >> y(env)
			}

		case xr.Uint16:
			x :=

				uint16(xv.Uint())
			fun = func(env *Env) uint16 {
				return x >> y(env)
			}

		case xr.Uint32:
			x :=

				uint32(xv.Uint())
			fun = func(env *Env) uint32 {
				return x >> y(env)
			}

		case xr.Uint64:
			x := xv.Uint()
			fun = func(env *Env) uint64 {
				return x >> y(env)
			}

		case xr.Uintptr:
			x :=

				uintptr(xv.Uint())
			fun = func(env *Env) uintptr {
				return x >> y(env)
			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return exprFun(xe.Type, fun)
}
