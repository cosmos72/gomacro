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
 * binary_eql.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package compiler

import (
	"go/token"
	r "reflect"
)

func (c *Comp) Lss(op token.Token, x I, y I) I {
	xlit, ylit := isLiteral(x), isLiteral(y)
	x, y = toSameFuncType(op, x, y)
	if !isClass(RetOf(x).Kind(), r.Int, r.Uint, r.Float64, r.Complex128, r.String) {
		return c.unsupportedBinaryExpr(op, x, y)
	}
	// if both x and y are literals, BinaryExpr will invoke evalConst()
	// on our return value. no need to optimize that.
	if xlit == ylit {
		switch x := x.(type) {
		case func(*Env) int:
			y := y.(func(*Env) int)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) int8:
			y := y.(func(*Env) int8)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) int16:
			y := y.(func(*Env) int16)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) int32:
			y := y.(func(*Env) int32)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) int64:
			y := y.(func(*Env) int64)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) uint:
			y := y.(func(*Env) uint)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) uint8:
			y := y.(func(*Env) uint8)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) uint16:
			y := y.(func(*Env) uint16)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) uint32:
			y := y.(func(*Env) uint32)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) uint64:
			y := y.(func(*Env) uint64)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) uintptr:
			y := y.(func(*Env) uintptr)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) float32:
			y := y.(func(*Env) float32)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) float64:
			y := y.(func(*Env) float64)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		case func(*Env) string:
			y := y.(func(*Env) string)
			return func(env *Env) bool {
				return x(env) < y(env)
			}
		default:
			return c.unsupportedBinaryExpr(op, x, y)
		}
	} else if ylit {
		y = c.evalConst(y)
		switch x := x.(type) {
		case func(*Env) int:
			y := y.(int)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) int8:
			y := y.(int8)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) int16:
			y := y.(int16)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) int32:
			y := y.(int32)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) int64:
			y := y.(int64)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) uint:
			y := y.(uint)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) uint8:
			y := y.(uint8)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) uint16:
			y := y.(uint16)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) uint32:
			y := y.(uint32)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) uint64:
			y := y.(uint64)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) uintptr:
			y := y.(uintptr)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) float32:
			y := y.(float32)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) float64:
			y := y.(float64)
			return func(env *Env) bool {
				return x(env) < y
			}
		case func(*Env) string:
			y := y.(string)
			return func(env *Env) bool {
				return x(env) < y
			}
		default:
			return c.unsupportedBinaryExpr(op, x, y)
		}
	} else {
		x = c.evalConst(x)
		switch y := y.(type) {
		case func(*Env) int8:
			x := x.(int8)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) int16:
			x := x.(int16)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) int32:
			x := x.(int32)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) int64:
			x := x.(int64)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) uint:
			x := x.(uint)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) uint8:
			x := x.(uint8)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) uint16:
			x := x.(uint16)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) uint32:
			x := x.(uint32)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) uint64:
			x := x.(uint64)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) uintptr:
			x := x.(uintptr)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) float32:
			x := x.(float32)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) float64:
			x := x.(float64)
			return func(env *Env) bool {
				return x < y(env)
			}
		case func(*Env) string:
			x := x.(string)
			return func(env *Env) bool {
				return x < y(env)
			}
		default:
			return c.unsupportedBinaryExpr(op, x, y)
		}
	}
}
