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
 * var_set.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	"unsafe"

	"github.com/cosmos72/gomacro/base"
)

func (c *Comp) varSetZero(upn int, index int, t r.Type) {
	zero := r.Zero(t).Interface()
	c.varSetConst(upn, index, t, zero)
}
func (c *Comp) varSetConst(upn int, index int, t r.Type, val I) {
	v := r.ValueOf(val)
	if base.ValueType(v) == nil {
		v = r.Zero(t)
	} else {
		v = v.Convert(t)
	}

	var ret func(env *Env) (Stmt, *Env)
	switch upn {
	case 0:
		switch val := val.(type) {
		case bool:

			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex128:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Binds[index].SetComplex(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case string:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Binds[index].SetString(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case 1:
		switch val := val.(type) {
		case bool:

			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex128:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					Binds[index].SetComplex(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case string:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					Binds[index].SetString(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case 2:
		switch val := val.(type) {
		case bool:

			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex128:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					Binds[index].SetComplex(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case string:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					Binds[index].SetString(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		switch val := val.(type) {
		case bool:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*bool)(unsafe.Pointer(&o.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) = val

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
					IntBinds[index])) = val

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
					IntBinds[index])) = val

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
					IntBinds[index])) = val

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
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex128:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					Binds[index].SetComplex(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case string:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					Binds[index].SetString(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case c.Depth - 1:
		switch val := val.(type) {
		case bool:

			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex128:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					Binds[index].SetComplex(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case string:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					Binds[index].SetString(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case c.Depth:
		switch val := val.(type) {
		case bool:

			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.TopEnv.
					IntBinds[index] = val

				env.IP++
				return env.Code[env.IP], env
			}
		case uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = val

				env.IP++
				return env.Code[env.IP], env
			}
		case complex128:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.TopEnv.
					Binds[index].SetComplex(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case string:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.TopEnv.
					Binds[index].SetString(val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.TopEnv.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	}
	c.Code.Append(ret)
}
func (c *Comp) varSetExpr(upn int, index int, t r.Type, fun I) {
	var ret func(env *Env) (Stmt, *Env)
	switch upn {
	case 0:
		switch fun := fun.(type) {
		case func(env *Env) bool:

			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Binds[index].SetComplex(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Binds[index].SetString(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) r.Value:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Binds[index].Set(fun(env).Convert(t),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf("invalid expression type, cannot compile assignment: %v <%v> returns %v",
				fun, r.TypeOf(fun), funTypeOuts(fun))
			return
		}
	case 1:
		switch fun := fun.(type) {
		case func(env *Env) bool:

			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					Binds[index].SetComplex(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					Binds[index].SetString(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) r.Value:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					Binds[index].Set(fun(env).Convert(t),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf("invalid expression type, cannot compile assignment: %v <%v> returns %v",
				fun, r.TypeOf(fun), funTypeOuts(fun))
			return
		}
	case 2:
		switch fun := fun.(type) {
		case func(env *Env) bool:

			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					Binds[index].SetComplex(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					Binds[index].SetString(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) r.Value:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					Binds[index].Set(fun(env).Convert(t),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf("invalid expression type, cannot compile assignment: %v <%v> returns %v",
				fun, r.TypeOf(fun), funTypeOuts(fun))
			return
		}
	default:
		switch fun := fun.(type) {
		case func(env *Env) bool:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*bool)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					Binds[index].SetComplex(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					Binds[index].SetString(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) r.Value:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					Binds[index].Set(fun(env).Convert(t),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf("invalid expression type, cannot compile assignment: %v <%v> returns %v",
				fun, r.TypeOf(fun), funTypeOuts(fun))
			return
		}
	case c.Depth - 1:
		switch fun := fun.(type) {
		case func(env *Env) bool:

			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					Binds[index].SetComplex(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					Binds[index].SetString(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) r.Value:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					Binds[index].Set(fun(env).Convert(t),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf("invalid expression type, cannot compile assignment: %v <%v> returns %v",
				fun, r.TypeOf(fun), funTypeOuts(fun))
			return
		}
	case c.Depth:
		switch fun := fun.(type) {
		case func(env *Env) bool:

			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) int64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint8:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint16:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uint64:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.TopEnv.
					IntBinds[index] = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float32:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) float64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex64:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.TopEnv.
					IntBinds[index])) = fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) complex128:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.TopEnv.
					Binds[index].SetComplex(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) string:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.TopEnv.
					Binds[index].SetString(fun(env),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case func(env *Env) r.Value:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.TopEnv.
					Binds[index].Set(fun(env).Convert(t),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf("invalid expression type, cannot compile assignment: %v <%v> returns %v",
				fun, r.TypeOf(fun), funTypeOuts(fun))
			return
		}
	}
	c.Code.Append(ret)
}
