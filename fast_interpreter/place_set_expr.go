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
 * assign_set_expr.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"unsafe"
)

// placeSetExpr compiles 'place = expression'
func (c *Comp) placeSetExpr(place *Place, init *Expr) {
	if place.Fun != nil {
		c.Errorf("unimplemented assignment to place (only assignment to variables is currently implemented)")
	}
	t := place.Type
	if t != nil && init.Type != t {
		if t.Kind() != init.Type.Kind() || !init.Type.AssignableTo(t) {
			c.Errorf("cannot assign <%v> to <%v>", init.Type, t)
			return
		}
	}
	upn := place.Upn
	desc := place.Desc
	var ret func(env *Env) (Stmt, *Env)

	switch desc.Class() {
	default:
		c.Errorf("cannot assign to %v", desc.Class())
		return
	case VarBind:
		index := desc.Index()
		if index == NoIndex {
			// assigning an expression to _
			// only keep the expression side effects
			c.Code.Append(init.AsStmt())
			return
		}
		switch upn {
		case 0:
			switch fun := init.Fun.(type) {
			case func(env *Env) complex128:
				ret = func(env *Env) (Stmt, *Env) {
					env.Binds[index].SetComplex(fun(env))
					env.IP++
					return env.Code[env.IP], env
				}
			case func(env *Env) string:
				ret = func(env *Env) (Stmt, *Env) {
					env.Binds[index].SetString(fun(env))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				f := init.AsX1()
				ret = func(env *Env) (Stmt, *Env) {
					env.Binds[index].Set(f(env).Convert(t))
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:
			switch fun := init.Fun.(type) {
			case func(env *Env) complex128:
				ret = func(env *Env) (Stmt, *Env) {
					env.Outer.Binds[index].SetComplex(fun(env))
					env.IP++
					return env.Code[env.IP], env
				}
			case func(env *Env) string:
				ret = func(env *Env) (Stmt, *Env) {
					env.Outer.Binds[index].SetString(fun(env))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				f := init.AsX1()
				ret = func(env *Env) (Stmt, *Env) {
					env.Outer.Binds[index].Set(f(env).Convert(t))
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:
			switch fun := init.Fun.(type) {
			case func(env *Env) complex128:
				ret = func(env *Env) (Stmt, *Env) {
					env.Outer.Outer.Binds[index].SetComplex(fun(env))
					env.IP++
					return env.Code[env.IP], env
				}
			case func(env *Env) string:
				ret = func(env *Env) (Stmt, *Env) {
					env.Outer.Outer.Binds[index].SetString(fun(env))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				f := init.AsX1()
				ret = func(env *Env) (Stmt, *Env) {
					env.Outer.Outer.Binds[index].Set(f(env).Convert(t))
					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			switch fun := init.Fun.(type) {
			case func(env *Env) complex128:
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					o.Binds[index].SetComplex(fun(env))
					env.IP++
					return env.Code[env.IP], env
				}
			case func(env *Env) string:
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					o.Binds[index].SetString(fun(env))
					env.IP++
					return env.Code[env.IP], env
				}
			default:
				f := init.AsX1()
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					o.Binds[index].Set(f(env).Convert(t))
					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case IntBind:
		index := desc.Index()
		if index == NoIndex {
			// assigning an expression to _
			// only keep the expression side effects
			c.Code.Append(init.AsStmt())
			return
		}
		switch fun := init.Fun.(type) {
		case func(*Env) bool:
			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int:
			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int8:
			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int16:
			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int32:
			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) int64:
			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint:
			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint8:
			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint16:
			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint32:
			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uint64:
			ret = func(env *Env) (Stmt, *Env) {
				env.IntBinds[index] = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) uintptr:
			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float32:
			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) float64:
			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		case func(*Env) complex64:
			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf("unsupported expression type, cannot use for optimized assignment: %v <%T> returns %v",
				fun, fun, init.Outs())
			return
		}
	}
	c.Code.Append(ret)
}
