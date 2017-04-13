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
 * var_set_expr.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	r "reflect"
	"unsafe"
)

func (c *Comp) varSetExpr(va *Var, init *Expr) {
	t := va.Type
	if t == nil {
		c.Errorf("invalid assignment: variable has type <%v>", t)
	}

	if init.Const() {
		c.Errorf("internal error: varSetExpr() invoked with constant expression. must call varSetConst() instead")
	} else if init.Type != t {
		if t.Kind() != init.Type.Kind() || !init.Type.AssignableTo(t) {
			c.Errorf("incompatible types in assignment: <%v> = <%v>", t, init.Type)
			return
		}
	}

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

			c.Code.Append(init.AsStmt())
			return
		}
		switch upn {
		case 0:
			switch fun := init.Fun.(type) {
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
						Binds[index].SetComplex(fun(env))

					env.IP++
					return env.Code[env.IP], env
				}
			case func(env *Env) string:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].SetString(fun(env))

					env.IP++
					return env.Code[env.IP], env
				}
			case func(env *Env) r.Value:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].Set(fun(env).Convert(t))

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf("unsupported expression type, cannot compile assignment: %v <%v> returns %v",
					fun, r.TypeOf(fun), init.Outs())
				return
			}

		case 1:
			switch fun := init.Fun.(type) {
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
						Binds[index].SetComplex(fun(env))

					env.IP++
					return env.Code[env.IP], env
				}
			case func(env *Env) string:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].SetString(fun(env))

					env.IP++
					return env.Code[env.IP], env
				}
			case func(env *Env) r.Value:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].Set(fun(env).Convert(t))

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf("unsupported expression type, cannot compile assignment: %v <%v> returns %v",
					fun, r.TypeOf(fun), init.Outs())
				return
			}

		case 2:
			switch fun := init.Fun.(type) {
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
						Binds[index].SetComplex(fun(env))

					env.IP++
					return env.Code[env.IP], env
				}
			case func(env *Env) string:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].SetString(fun(env))

					env.IP++
					return env.Code[env.IP], env
				}
			case func(env *Env) r.Value:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].Set(fun(env).Convert(t))

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf("unsupported expression type, cannot compile assignment: %v <%v> returns %v",
					fun, r.TypeOf(fun), init.Outs())
				return
			}

		default:
			switch fun := init.Fun.(type) {
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
						Binds[index].SetComplex(fun(env))

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
						Binds[index].SetString(fun(env))

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
						Binds[index].Set(fun(env).Convert(t))

					env.IP++
					return env.Code[env.IP], env
				}
			default:
				c.Errorf("unsupported expression type, cannot compile assignment: %v <%v> returns %v",
					fun, r.TypeOf(fun), init.Outs())
				return
			}

		}
	}
	c.Code.Append(ret)
}
