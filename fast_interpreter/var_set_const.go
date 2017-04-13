// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated

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
 * var_set_const.gomacro
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	r "reflect"
	"unsafe"

	"github.com/cosmos72/gomacro/base"
)

func (c *Comp) varSetConst(va *Var, init *Expr) {
	t := va.Type
	if !init.Const() {
		c.Errorf("internal error: varSetConst() invoked with constant expression. must call varSetExpr() instead")
	}

	init.ConstTo(t)

	upn := va.Upn
	desc := va.Desc
	var ret func(env *Env) (Stmt, *Env)

	switch desc.Class() {
	default:
		c.Errorf("cannot assign to %v", desc.Class())
		return
	case VarBind, IntBind:
		index := desc.Index()
		if index == NoIndex {
			return
		}

		val := init.Value
		v := r.ValueOf(val)
		if base.ValueType(v) == nil {
			v = r.Zero(t)
		} else {
			v = v.Convert(t)
		}

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
						Binds[index].SetComplex(val)

					env.IP++
					return env.Code[env.IP], env
				}
			case string:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].SetString(val)

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].Set(v)

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
						Binds[index].SetComplex(val)

					env.IP++
					return env.Code[env.IP], env
				}
			case string:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].SetString(val)

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].Set(v)

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
						Binds[index].SetComplex(val)

					env.IP++
					return env.Code[env.IP], env
				}
			case string:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].SetString(val)

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].Set(v)

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
						Binds[index].SetComplex(val)

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
						Binds[index].SetString(val)

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
						Binds[index].Set(v)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	}
	c.Code.Append(ret)
}
