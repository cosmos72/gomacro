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
 * callexpr.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	r "reflect"
)

func callExpr1Optimized(expr *Expr, args []*Expr, argfuns []func(*Env) r.Value) I {
	exprfun := expr.AsX1()
	t := expr.Type.Out(0)
	var call I

	if expr.Type.NumIn() == 1 && expr.Type.In(0) == t {
		switch t.Kind() {
		case r.Bool:

			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(bool)
					call = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(bool) bool)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) bool)
					call = func(env *Env) bool {
						fun := exprfun(env).Interface().(func(bool) bool)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Int:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(int)
					call = func(env *Env) int {
						fun := exprfun(env).Interface().(func(int) int)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) int)
					call = func(env *Env) int {
						fun := exprfun(env).Interface().(func(int) int)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Int8:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(int8)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(int8) int8)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) int8)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(int8) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Int16:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(int16)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(int16) int16)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) int16)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(int16) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Int32:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(int32)
					call = func(env *Env) int32 {
						fun := exprfun(env).Interface().(func(int32) int32)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) int32)
					call = func(env *Env) int32 {
						fun := exprfun(env).Interface().(func(int32) int32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Int64:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(int64)
					call = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(int64) int64)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) int64)
					call = func(env *Env) int64 {
						fun := exprfun(env).Interface().(func(int64) int64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Uint:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(uint)
					call = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(uint) uint)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) uint)
					call = func(env *Env) uint {
						fun := exprfun(env).Interface().(func(uint) uint)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Uint8:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(uint8)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(uint8) uint8)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) uint8)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(uint8) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Uint16:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(uint16)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(uint16) uint16)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) uint16)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(uint16) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Uint32:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(uint32)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(uint32) uint32)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) uint32)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(uint32) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Uint64:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(uint64)
					call = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(uint64) uint64)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) uint64)
					call = func(env *Env) uint64 {
						fun := exprfun(env).Interface().(func(uint64) uint64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Uintptr:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(uintptr)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(uintptr) uintptr)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) uintptr)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(uintptr) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Float32:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(float32)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(float32) float32)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) float32)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(float32) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Float64:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(float64)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(float64) float64)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) float64)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(float64) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Complex64:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(complex64)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(complex64) complex64)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) complex64)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(complex64) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.Complex128:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(complex128)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(complex128) complex128)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) complex128)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(complex128) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		case r.String:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(string)
					call = func(env *Env) string {
						fun := exprfun(env).Interface().(func(string) string)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) string)
					call = func(env *Env) string {
						fun := exprfun(env).Interface().(func(string) string)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		default:
			{
				arg0 := args[0]
				if arg0.Const() {
					argconst := args[0].Value.(r.Value)
					call = func(env *Env) r.Value {
						fun := exprfun(env).Interface().(func(r.Value) r.Value)
						return fun(argconst)
					}
				} else {
					argfun := args[0].Fun.(func(env *Env) r.Value)
					call = func(env *Env) r.Value {
						fun := exprfun(env).Interface().(func(r.Value) r.Value)
						arg := argfun(env)
						return fun(arg)
					}
				}
			}
		}
	} else {
		switch t.Kind() {
		case r.Bool:
			call = func(env *Env) bool {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Bool()
			}

		case r.Int:
			call = func(env *Env) int {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return int(ret0.Int())
			}

		case r.Int8:
			call = func(env *Env) int8 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return int8(ret0.Int())
			}

		case r.Int16:
			call = func(env *Env) int16 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return int16(ret0.Int())
			}

		case r.Int32:
			call = func(env *Env) int32 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return int32(ret0.Int())
			}

		case r.Int64:
			call = func(env *Env) int64 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Int()
			}

		case r.Uint:
			call = func(env *Env) uint {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return uint(ret0.Uint())
			}

		case r.Uint8:
			call = func(env *Env) uint8 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return uint8(ret0.Uint())
			}

		case r.Uint16:
			call = func(env *Env) uint16 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return uint16(ret0.Uint())
			}

		case r.Uint32:
			call = func(env *Env) uint32 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return uint32(ret0.Uint())
			}

		case r.Uint64:
			call = func(env *Env) uint64 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Uint()
			}

		case r.Uintptr:
			call = func(env *Env) uintptr {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return uintptr(ret0.Uint())
			}

		case r.Float32:
			call = func(env *Env) float32 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return float32(ret0.Float())
			}

		case r.Float64:
			call = func(env *Env) float64 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Float()
			}

		case r.Complex64:
			call = func(env *Env) complex64 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return complex64(ret0.Complex())
			}

		case r.Complex128:
			call = func(env *Env) complex128 {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Complex()
			}

		case r.String:
			call = func(env *Env) string {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return ret0.String()
			}

		default:
			call = func(env *Env) r.Value {
				funv := exprfun(env)
				argv := make([]r.Value, len(argfuns))
				for i, argfun := range argfuns {
					argv[i] = argfun(env)
				}

				ret0 := funv.Call(argv)[0]
				return ret0

			}

		}
	}
	return call
}
