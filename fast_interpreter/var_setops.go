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
 * var_add.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/token"
	r "reflect"
	"unsafe"

	"github.com/cosmos72/gomacro/base"
)

func (c *Comp) varAddConst(upn int, index int, t r.Type, val I) {
	v := r.ValueOf(val)
	if base.ValueType(v) != t {
		v = v.Convert(t)
	}

	var ret Stmt
	switch upn {
	case 0:

		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case string:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetString(lhs.String() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.ADD, t)

		}
	case 1:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case string:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetString(lhs.String() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.ADD, t)

		}
	case 2:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case string:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetString(lhs.String() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.ADD, t)

		}
	default:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) += val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case string:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetString(lhs.String() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.ADD, t)

		}

	}
	c.Code.Append(ret)
}
func (c *Comp) varAddExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch upn {
	case 0:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.ADD, t, funTypeOuts(fun))

		}
	case 1:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.ADD, t, funTypeOuts(fun))

		}
	case 2:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.ADD, t, funTypeOuts(fun))

		}
	default:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.ADD, t, funTypeOuts(fun))

		}
	}
	c.Code.Append(ret)
}
func (c *Comp) varSubConst(upn int, index int, t r.Type, val I) {
	v := r.ValueOf(val)
	if base.ValueType(v) != t {
		v = v.Convert(t)
	}

	var ret Stmt
	switch upn {
	case 0:

		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.SUB, t)

		}
	case 1:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.SUB, t)

		}
	case 2:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.SUB, t)

		}
	default:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) -= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.SUB, t)

		}

	}
	c.Code.Append(ret)
}
func (c *Comp) varSubExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch upn {
	case 0:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SUB, t, funTypeOuts(fun))

		}
	case 1:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SUB, t, funTypeOuts(fun))

		}
	case 2:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SUB, t, funTypeOuts(fun))

		}
	default:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SUB, t, funTypeOuts(fun))

		}
	}
	c.Code.Append(ret)
}
func (c *Comp) varMulConst(upn int, index int, t r.Type, val I) {
	v := r.ValueOf(val)
	if base.ValueType(v) != t {
		v = v.Convert(t)
	}

	var ret Stmt
	switch upn {
	case 0:

		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.MUL, t)

		}
	case 1:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.MUL, t)

		}
	case 2:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.MUL, t)

		}
	default:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) *= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.MUL, t)

		}

	}
	c.Code.Append(ret)
}
func (c *Comp) varMulExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch upn {
	case 0:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.MUL, t, funTypeOuts(fun))

		}
	case 1:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.MUL, t, funTypeOuts(fun))

		}
	case 2:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.MUL, t, funTypeOuts(fun))

		}
	default:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.MUL, t, funTypeOuts(fun))

		}
	}
	c.Code.Append(ret)
}
func (c *Comp) varQuoConst(upn int, index int, t r.Type, val I) {
	v := r.ValueOf(val)
	if base.ValueType(v) != t {
		v = v.Convert(t)
	}

	var ret Stmt
	switch upn {
	case 0:

		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.QUO, t)

		}
	case 1:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.QUO, t)

		}
	case 2:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.QUO, t)

		}
	default:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) /= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.QUO, t)

		}

	}
	c.Code.Append(ret)
}
func (c *Comp) varQuoExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch upn {
	case 0:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.QUO, t, funTypeOuts(fun))

		}
	case 1:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.QUO, t, funTypeOuts(fun))

		}
	case 2:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.QUO, t, funTypeOuts(fun))

		}
	default:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.QUO, t, funTypeOuts(fun))

		}
	}
	c.Code.Append(ret)
}
func (c *Comp) varRemConst(upn int, index int, t r.Type, val I) {
	v := r.ValueOf(val)
	if base.ValueType(v) != t {
		v = v.Convert(t)
	}

	var ret Stmt
	switch upn {
	case 0:

		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.REM, t)

		}
	case 1:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.REM, t)

		}
	case 2:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.REM, t)

		}
	default:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] %= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) %= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.REM, t)

		}

	}
	c.Code.Append(ret)
}
func (c *Comp) varRemExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch upn {
	case 0:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.REM, t, funTypeOuts(fun))

		}
	case 1:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.REM, t, funTypeOuts(fun))

		}
	case 2:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.REM, t, funTypeOuts(fun))

		}
	default:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.REM, t, funTypeOuts(fun))

		}
	}
	c.Code.Append(ret)
}
func (c *Comp) varAndConst(upn int, index int, t r.Type, val I) {
	v := r.ValueOf(val)
	if base.ValueType(v) != t {
		v = v.Convert(t)
	}

	var ret Stmt
	switch upn {
	case 0:

		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.AND, t)

		}
	case 1:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.AND, t)

		}
	case 2:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.AND, t)

		}
	default:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] &= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) &= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.AND, t)

		}

	}
	c.Code.Append(ret)
}
func (c *Comp) varAndExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch upn {
	case 0:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND, t, funTypeOuts(fun))

		}
	case 1:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND, t, funTypeOuts(fun))

		}
	case 2:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND, t, funTypeOuts(fun))

		}
	default:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND, t, funTypeOuts(fun))

		}
	}
	c.Code.Append(ret)
}
func (c *Comp) varOrConst(upn int, index int, t r.Type, val I) {
	v := r.ValueOf(val)
	if base.ValueType(v) != t {
		v = v.Convert(t)
	}

	var ret Stmt
	switch upn {
	case 0:

		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.OR, t)

		}
	case 1:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.OR, t)

		}
	case 2:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.OR, t)

		}
	default:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] |= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) |= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.OR, t)

		}

	}
	c.Code.Append(ret)
}
func (c *Comp) varOrExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch upn {
	case 0:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.OR, t, funTypeOuts(fun))

		}
	case 1:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.OR, t, funTypeOuts(fun))

		}
	case 2:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.OR, t, funTypeOuts(fun))

		}
	default:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.OR, t, funTypeOuts(fun))

		}
	}
	c.Code.Append(ret)
}
func (c *Comp) varXorConst(upn int, index int, t r.Type, val I) {
	v := r.ValueOf(val)
	if base.ValueType(v) != t {
		v = v.Convert(t)
	}

	var ret Stmt
	switch upn {
	case 0:

		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

		}
	case 1:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

		}
	case 2:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

		}
	default:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

		}

	}
	c.Code.Append(ret)
}
func (c *Comp) varXorExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch upn {
	case 0:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, t, funTypeOuts(fun))

		}
	case 1:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, t, funTypeOuts(fun))

		}
	case 2:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, t, funTypeOuts(fun))

		}
	default:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, t, funTypeOuts(fun))

		}
	}
	c.Code.Append(ret)
}
func (c *Comp) varAndnotConst(upn int, index int, t r.Type, val I) {
	v := r.ValueOf(val)
	if base.ValueType(v) != t {
		v = v.Convert(t)
	}

	var ret Stmt
	switch upn {
	case 0:

		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

		}
	case 1:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

		}
	case 2:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

		}
	default:
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) ^= val

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

		}

	}
	c.Code.Append(ret)
}
func (c *Comp) varAndnotExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch upn {
	case 0:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, t, funTypeOuts(fun))

		}
	case 1:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, t, funTypeOuts(fun))

		}
	case 2:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, t, funTypeOuts(fun))

		}
	default:

		switch fun := fun.(type) {
		case func(*Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, t, funTypeOuts(fun))

		}
	}
	c.Code.Append(ret)
}
func (c *Comp) varSetOp(va *Var, op token.Token, init *Expr) {
	t := va.Type
	if init.Const() {
		init.ConstTo(t)
	} else if init.Type != t {
		if t.Kind() != init.Type.Kind() || !init.Type.AssignableTo(t) {
			c.Errorf("incompatible types in assignment: <%v> %s <%v>", t, op, init.Type)
			return
		}
	}

	class := va.Desc.Class()
	if class != VarBind && class != IntBind {
		c.Errorf("invalid operator %s on %v", op, class)
		return
	}
	upn := va.Upn
	index := va.Desc.Index()
	if index == NoIndex {
		c.Errorf("invalid operator %s on _", op)
		return
	}
	if init.Const() {
		switch op {
		case token.ADD, token.ADD_ASSIGN:
			c.varAddConst(upn, index, t, init.Value)
		case token.SUB, token.SUB_ASSIGN:
			c.varSubConst(upn, index, t, init.Value)
		case token.MUL, token.MUL_ASSIGN:
			c.varMulConst(upn, index, t, init.Value)
		case token.QUO, token.QUO_ASSIGN:
			c.varQuoConst(upn, index, t, init.Value)
		case token.REM, token.REM_ASSIGN:
			c.varRemConst(upn, index, t, init.Value)
		case token.AND, token.AND_ASSIGN:
			c.varAndConst(upn, index, t, init.Value)
		case token.OR, token.OR_ASSIGN:
			c.varOrConst(upn, index, t, init.Value)
		case token.XOR, token.XOR_ASSIGN:
			c.varAndConst(upn, index, t, init.Value)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			c.varAndnotConst(upn, index, t, init.Value)
		default:
			c.Errorf("operator %s is not implemented", op)
		}
	} else {
		switch op {
		case token.ADD, token.ADD_ASSIGN:
			c.varAddExpr(upn, index, t, init.Fun)
		case token.SUB, token.SUB_ASSIGN:
			c.varSubExpr(upn, index, t, init.Fun)
		case token.MUL, token.MUL_ASSIGN:
			c.varMulExpr(upn, index, t, init.Fun)
		case token.QUO, token.QUO_ASSIGN:
			c.varQuoExpr(upn, index, t, init.Fun)
		case token.REM, token.REM_ASSIGN:
			c.varRemExpr(upn, index, t, init.Fun)
		case token.AND, token.AND_ASSIGN:
			c.varAndExpr(upn, index, t, init.Fun)
		case token.OR, token.OR_ASSIGN:
			c.varOrExpr(upn, index, t, init.Fun)
		case token.XOR, token.XOR_ASSIGN:
			c.varAndExpr(upn, index, t, init.Fun)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			c.varAndnotExpr(upn, index, t, init.Fun)
		default:
			c.Errorf("operator %s is not implemented", op)
		}
	}

}
