// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

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
 * binary_relops.go
 *
 *  Created on Apr 12, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
	r "reflect"
)

func (c *Comp) Lss(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun func(*Env) bool
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool {
					return x(env) < y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := y.(int)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := y.(int8)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := y.(int16)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := y.(int32)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := y.(int64)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(uint)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(uint8)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(uint16)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(uint32)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(uint64)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(uintptr)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(float32)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(float64)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}
		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(string)
				fun = func(env *Env) bool {
					return x(env) < y
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := x.(int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Int8:

			{
				x := x.(int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Int16:

			{
				x := x.(int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Int32:

			{
				x := x.(int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Int64:

			{
				x := x.(int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Uint:

			{
				x := x.(uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Uint8:

			{
				x := x.(uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Uint16:

			{
				x := x.(uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Uint32:

			{
				x := x.(uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Uint64:

			{
				x := x.(uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Uintptr:

			{
				x := x.(uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Float32:

			{
				x := x.(float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.Float64:

			{
				x := x.(float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		case r.String:

			{
				x := x.(string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool {
					return x < y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return ExprBool(fun)
}
func (c *Comp) Gtr(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun func(*Env) bool
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool {
					return x(env) > y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := y.(int)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := y.(int8)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := y.(int16)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := y.(int32)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := y.(int64)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(uint)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(uint8)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(uint16)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(uint32)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(uint64)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(uintptr)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(float32)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(float64)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}
		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(string)
				fun = func(env *Env) bool {
					return x(env) > y
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := x.(int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Int8:

			{
				x := x.(int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Int16:

			{
				x := x.(int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Int32:

			{
				x := x.(int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Int64:

			{
				x := x.(int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Uint:

			{
				x := x.(uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Uint8:

			{
				x := x.(uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Uint16:

			{
				x := x.(uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Uint32:

			{
				x := x.(uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Uint64:

			{
				x := x.(uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Uintptr:

			{
				x := x.(uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Float32:

			{
				x := x.(float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.Float64:

			{
				x := x.(float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		case r.String:

			{
				x := x.(string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool {
					return x > y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return ExprBool(fun)
}
func (c *Comp) Leq(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun func(*Env) bool
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool {
					return x(env) <= y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := y.(int)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := y.(int8)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := y.(int16)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := y.(int32)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := y.(int64)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(uint)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(uint8)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(uint16)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(uint32)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(uint64)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(uintptr)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(float32)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(float64)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}
		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(string)
				fun = func(env *Env) bool {
					return x(env) <= y
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := x.(int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Int8:

			{
				x := x.(int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Int16:

			{
				x := x.(int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Int32:

			{
				x := x.(int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Int64:

			{
				x := x.(int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Uint:

			{
				x := x.(uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Uint8:

			{
				x := x.(uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Uint16:

			{
				x := x.(uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Uint32:

			{
				x := x.(uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Uint64:

			{
				x := x.(uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Uintptr:

			{
				x := x.(uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Float32:

			{
				x := x.(float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.Float64:

			{
				x := x.(float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		case r.String:

			{
				x := x.(string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool {
					return x <= y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return ExprBool(fun)
}
func (c *Comp) Geq(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	c.toSameFuncType(node, xe, ye)
	k := xe.Type.Kind()

	var fun func(*Env) bool
	if xc == yc {
		x, y := xe.Fun, ye.Fun

		switch k {
		case r.Int:
			{
				x := x.(func(*Env) int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}
		case r.Int8:
			{
				x := x.(func(*Env) int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}
		case r.Int16:
			{
				x := x.(func(*Env) int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}
		case r.Int32:
			{
				x := x.(func(*Env) int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}
		case r.Int64:
			{
				x := x.(func(*Env) int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}

		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}

		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}

		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}

		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}

		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}

		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}

		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}

		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}

		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool {
					return x(env) >= y(env)
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else if yc {
		x := xe.Fun
		y := ye.Value

		switch k {
		case r.Int:

			{
				x := x.(func(*Env) int)
				y := y.(int)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Int8:

			{
				x := x.(func(*Env) int8)
				y := y.(int8)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Int16:

			{
				x := x.(func(*Env) int16)
				y := y.(int16)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Int32:

			{
				x := x.(func(*Env) int32)
				y := y.(int32)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Int64:

			{
				x := x.(func(*Env) int64)
				y := y.(int64)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Uint:
			{
				x := x.(func(*Env) uint)
				y := y.(uint)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Uint8:
			{
				x := x.(func(*Env) uint8)
				y := y.(uint8)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Uint16:
			{
				x := x.(func(*Env) uint16)
				y := y.(uint16)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Uint32:
			{
				x := x.(func(*Env) uint32)
				y := y.(uint32)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Uint64:
			{
				x := x.(func(*Env) uint64)
				y := y.(uint64)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Uintptr:
			{
				x := x.(func(*Env) uintptr)
				y := y.(uintptr)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Float32:
			{
				x := x.(func(*Env) float32)
				y := y.(float32)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.Float64:
			{
				x := x.(func(*Env) float64)
				y := y.(float64)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}
		case r.String:
			{
				x := x.(func(*Env) string)
				y := y.(string)
				fun = func(env *Env) bool {
					return x(env) >= y
				}

			}

		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	} else {
		x := xe.Value
		y := ye.Fun

		switch k {
		case r.Int:

			{
				x := x.(int)
				y := y.(func(*Env) int)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Int8:

			{
				x := x.(int8)
				y := y.(func(*Env) int8)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Int16:

			{
				x := x.(int16)
				y := y.(func(*Env) int16)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Int32:

			{
				x := x.(int32)
				y := y.(func(*Env) int32)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Int64:

			{
				x := x.(int64)
				y := y.(func(*Env) int64)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Uint:

			{
				x := x.(uint)
				y := y.(func(*Env) uint)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Uint8:

			{
				x := x.(uint8)
				y := y.(func(*Env) uint8)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Uint16:

			{
				x := x.(uint16)
				y := y.(func(*Env) uint16)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Uint32:

			{
				x := x.(uint32)
				y := y.(func(*Env) uint32)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Uint64:

			{
				x := x.(uint64)
				y := y.(func(*Env) uint64)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Uintptr:

			{
				x := x.(uintptr)
				y := y.(func(*Env) uintptr)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Float32:

			{
				x := x.(float32)
				y := y.(func(*Env) float32)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.Float64:

			{
				x := x.(float64)
				y := y.(func(*Env) float64)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		case r.String:

			{
				x := x.(string)
				y := y.(func(*Env) string)
				fun = func(env *Env) bool {
					return x >= y(env)
				}

			}
		default:
			return c.invalidBinaryExpr(node, xe, ye)

		}
	}
	return ExprBool(fun)
}
