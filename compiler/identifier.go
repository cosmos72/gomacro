/*
 * gomacro - A Go intepreter with Lisp-like macros
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
 * identifier.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package compiler

import (
	r "reflect"
)

func (c *Comp) resolve(name string) (upn int, bind Bind) {
	ok := false
	upn, bind, ok = c.tryResolve(name)
	if !ok {
		c.Errorf("undefined identifier: %v", name)
	}
	return upn, bind
}

func (c *Comp) tryResolve(name string) (upn int, bind Bind, ok bool) {
	for ; c != nil; c = c.Outer {
		if b, ok := c.Binds[name]; ok {
			return upn, b, true
		}
		upn++
	}
	return upn, bind, false
}

// Ident compiles a read operation on a constant, variable or function
func (c *Comp) Ident(name string) *Expr {
	upn, bind := c.resolve(name)
	if bind.Const() {
		return ExprLit(bind.Lit)
	}
	idx := bind.Index
	k := bind.Type.Kind()
	// be clever and extract primitive values from r.Value
	switch upn {
	case 0:
		switch k {
		case r.Bool:
			return ExprBool(func(env *Env) bool {
				return env.Binds[idx].Bool()
			})
		case r.Int:
			return ExprInt(func(env *Env) int {
				return int(env.Binds[idx].Int())
			})
		case r.Int8:
			return ExprInt8(func(env *Env) int8 {
				return int8(env.Binds[idx].Int())
			})
		case r.Int16:
			return ExprInt16(func(env *Env) int16 {
				return int16(env.Binds[idx].Int())
			})
		case r.Int32:
			return ExprInt32(func(env *Env) int32 {
				return int32(env.Binds[idx].Int())
			})
		case r.Int64:
			return ExprInt64(func(env *Env) int64 {
				return env.Binds[idx].Int()
			})
		case r.Uint:
			return ExprUint(func(env *Env) uint {
				return uint(env.Binds[idx].Uint())
			})
		case r.Uint8:
			return ExprUint8(func(env *Env) uint8 {
				return uint8(env.Binds[idx].Uint())
			})
		case r.Uint16:
			return ExprUint16(func(env *Env) uint16 {
				return uint16(env.Binds[idx].Uint())
			})
		case r.Uint32:
			return ExprUint32(func(env *Env) uint32 {
				return uint32(env.Binds[idx].Uint())
			})
		case r.Uint64:
			return ExprUint64(func(env *Env) uint64 {
				return env.Binds[idx].Uint()
			})
		case r.Uintptr:
			return ExprUintptr(func(env *Env) uintptr {
				return uintptr(env.Binds[idx].Uint())
			})
		case r.Float32:
			return ExprFloat32(func(env *Env) float32 {
				return float32(env.Binds[idx].Float())
			})
		case r.Float64:
			return ExprFloat64(func(env *Env) float64 {
				return env.Binds[idx].Float()
			})
		case r.Complex64:
			return ExprComplex64(func(env *Env) complex64 {
				return complex64(env.Binds[idx].Complex())
			})
		case r.Complex128:
			return ExprComplex128(func(env *Env) complex128 {
				return env.Binds[idx].Complex()
			})
		case r.String:
			return ExprString(func(env *Env) string {
				return env.Binds[idx].String()
			})
		default:
			return Expr1(bind.Type, func(env *Env) (r.Value, []r.Value) {
				return env.Binds[idx], nil
			})
		}
	case 1:
		switch k {
		case r.Bool:
			return ExprBool(func(env *Env) bool {
				return env.Outer.Binds[idx].Bool()
			})
		case r.Int:
			return ExprInt(func(env *Env) int {
				return int(env.Outer.Binds[idx].Int())
			})
		case r.Int8:
			return ExprInt8(func(env *Env) int8 {
				return int8(env.Outer.Binds[idx].Int())
			})
		case r.Int16:
			return ExprInt16(func(env *Env) int16 {
				return int16(env.Outer.Binds[idx].Int())
			})
		case r.Int32:
			return ExprInt32(func(env *Env) int32 {
				return int32(env.Outer.Binds[idx].Int())
			})
		case r.Int64:
			return ExprInt64(func(env *Env) int64 {
				return env.Outer.Binds[idx].Int()
			})
		case r.Uint:
			return ExprUint(func(env *Env) uint {
				return uint(env.Outer.Binds[idx].Uint())
			})
		case r.Uint8:
			return ExprUint8(func(env *Env) uint8 {
				return uint8(env.Outer.Binds[idx].Uint())
			})
		case r.Uint16:
			return ExprUint16(func(env *Env) uint16 {
				return uint16(env.Outer.Binds[idx].Uint())
			})
		case r.Uint32:
			return ExprUint32(func(env *Env) uint32 {
				return uint32(env.Outer.Binds[idx].Uint())
			})
		case r.Uint64:
			return ExprUint64(func(env *Env) uint64 {
				return env.Outer.Binds[idx].Uint()
			})
		case r.Uintptr:
			return ExprUintptr(func(env *Env) uintptr {
				return uintptr(env.Outer.Binds[idx].Uint())
			})
		case r.Float32:
			return ExprFloat32(func(env *Env) float32 {
				return float32(env.Outer.Binds[idx].Float())
			})
		case r.Float64:
			return ExprFloat64(func(env *Env) float64 {
				return env.Outer.Binds[idx].Float()
			})
		case r.Complex64:
			return ExprComplex64(func(env *Env) complex64 {
				return complex64(env.Outer.Binds[idx].Complex())
			})
		case r.Complex128:
			return ExprComplex128(func(env *Env) complex128 {
				return env.Outer.Binds[idx].Complex()
			})
		case r.String:
			return ExprString(func(env *Env) string {
				return env.Outer.Binds[idx].String()
			})
		default:
			return Expr1(bind.Type, func(env *Env) (r.Value, []r.Value) {
				return env.Outer.Binds[idx], nil
			})
		}
	case 2:
		switch k {
		case r.Bool:
			return ExprBool(func(env *Env) bool {
				return env.Outer.Outer.Binds[idx].Bool()
			})
		case r.Int:
			return ExprInt(func(env *Env) int {
				return int(env.Outer.Outer.Binds[idx].Int())
			})
		case r.Int8:
			return ExprInt8(func(env *Env) int8 {
				return int8(env.Outer.Outer.Binds[idx].Int())
			})
		case r.Int16:
			return ExprInt16(func(env *Env) int16 {
				return int16(env.Outer.Outer.Binds[idx].Int())
			})
		case r.Int32:
			return ExprInt32(func(env *Env) int32 {
				return int32(env.Outer.Outer.Binds[idx].Int())
			})
		case r.Int64:
			return ExprInt64(func(env *Env) int64 {
				return env.Outer.Outer.Binds[idx].Int()
			})
		case r.Uint:
			return ExprUint(func(env *Env) uint {
				return uint(env.Outer.Outer.Binds[idx].Uint())
			})
		case r.Uint8:
			return ExprUint8(func(env *Env) uint8 {
				return uint8(env.Outer.Outer.Binds[idx].Uint())
			})
		case r.Uint16:
			return ExprUint16(func(env *Env) uint16 {
				return uint16(env.Outer.Outer.Binds[idx].Uint())
			})
		case r.Uint32:
			return ExprUint32(func(env *Env) uint32 {
				return uint32(env.Outer.Outer.Binds[idx].Uint())
			})
		case r.Uint64:
			return ExprUint64(func(env *Env) uint64 {
				return env.Outer.Outer.Binds[idx].Uint()
			})
		case r.Uintptr:
			return ExprUintptr(func(env *Env) uintptr {
				return uintptr(env.Outer.Outer.Binds[idx].Uint())
			})
		case r.Float32:
			return ExprFloat32(func(env *Env) float32 {
				return float32(env.Outer.Outer.Binds[idx].Float())
			})
		case r.Float64:
			return ExprFloat64(func(env *Env) float64 {
				return env.Outer.Outer.Binds[idx].Float()
			})
		case r.Complex64:
			return ExprComplex64(func(env *Env) complex64 {
				return complex64(env.Outer.Outer.Binds[idx].Complex())
			})
		case r.Complex128:
			return ExprComplex128(func(env *Env) complex128 {
				return env.Outer.Outer.Binds[idx].Complex()
			})
		case r.String:
			return ExprString(func(env *Env) string {
				return env.Outer.Outer.Binds[idx].String()
			})
		default:
			return Expr1(bind.Type, func(env *Env) (r.Value, []r.Value) {
				return env.Outer.Outer.Binds[idx], nil
			})
		}
	default:
		switch k {
		case r.Bool:
			return ExprBool(func(env *Env) bool {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].Bool()
			})
		case r.Int:
			return ExprInt(func(env *Env) int {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return int(env.Outer.Outer.Outer.Binds[idx].Int())
			})
		case r.Int8:
			return ExprInt8(func(env *Env) int8 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return int8(env.Outer.Outer.Outer.Binds[idx].Int())
			})
		case r.Int16:
			return ExprInt16(func(env *Env) int16 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return int16(env.Outer.Outer.Outer.Binds[idx].Int())
			})
		case r.Int32:
			return ExprInt32(func(env *Env) int32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return int32(env.Outer.Outer.Outer.Binds[idx].Int())
			})
		case r.Int64:
			return ExprInt64(func(env *Env) int64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].Int()
			})
		case r.Uint:
			return ExprUint(func(env *Env) uint {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return uint(env.Outer.Outer.Outer.Binds[idx].Uint())
			})
		case r.Uint8:
			return ExprUint8(func(env *Env) uint8 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return uint8(env.Outer.Outer.Outer.Binds[idx].Uint())
			})
		case r.Uint16:
			return ExprUint16(func(env *Env) uint16 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return uint16(env.Outer.Outer.Outer.Binds[idx].Uint())
			})
		case r.Uint32:
			return ExprUint32(func(env *Env) uint32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return uint32(env.Outer.Outer.Outer.Binds[idx].Uint())
			})
		case r.Uint64:
			return ExprUint64(func(env *Env) uint64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].Uint()
			})
		case r.Uintptr:
			return ExprUintptr(func(env *Env) uintptr {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return uintptr(env.Outer.Outer.Outer.Binds[idx].Uint())
			})
		case r.Float32:
			return ExprFloat32(func(env *Env) float32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return float32(env.Outer.Outer.Outer.Binds[idx].Float())
			})
		case r.Float64:
			return ExprFloat64(func(env *Env) float64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].Float()
			})
		case r.Complex64:
			return ExprComplex64(func(env *Env) complex64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return complex64(env.Outer.Outer.Outer.Binds[idx].Complex())
			})
		case r.Complex128:
			return ExprComplex128(func(env *Env) complex128 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].Complex()
			})
		case r.String:
			return ExprString(func(env *Env) string {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].String()
			})
		default:
			return Expr1(bind.Type, func(env *Env) (r.Value, []r.Value) {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx], nil
			})
		}
	}
}
