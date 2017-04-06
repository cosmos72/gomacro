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
 * address.go
 *
 *  Created on Apr 05, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	r "reflect"
	"unsafe"
)

type Place struct {
	addr    *r.Value // the reflect.Value to modify, or nil
	intaddr *uint64  // the int variable to modify, or nil
	mapobj  r.Value  // the map to modify, or Nil
	mapkey  r.Value  // the map key to set, or Nil
}

// IdentAddress compiles the expression "&name" where name is a variable
func (c *Comp) IdentAddress(name string) *Expr {
	upn, bind := c.resolve(name)
	switch bind.Desc.Class() {
	default:
		c.Errorf("cannot take the address of %v", name)
		return nil
	case VarBind:
		return c.identBindAddress(name, upn, bind)
	case IntBind:
		return c.identIntBindAddress(name, upn, bind)
	}
}

func (c *Comp) identBindAddress(name string, upn int, bind Bind) *Expr {
	idx := bind.Desc.Index()
	var fun I
	switch upn {
	case 0:
		fun = func(env *Env) r.Value {
			return env.Binds[idx].Addr()
		}
	case 1:
		fun = func(env *Env) r.Value {
			return env.Outer.Binds[idx].Addr()
		}
	case 2:
		fun = func(env *Env) r.Value {
			return env.Outer.Outer.Binds[idx].Addr()
		}
	default:
		fun = func(env *Env) r.Value {
			for i := 3; i < upn; i++ {
				env = env.Outer
			}
			return env.Outer.Outer.Outer.Binds[idx].Addr()
		}
	}
	return &Expr{Lit: Lit{Type: r.PtrTo(bind.Type)}, Fun: fun}
}

func (c *Comp) identIntBindAddress(name string, upn int, bind Bind) *Expr {
	k := bind.Type.Kind()
	idx := bind.Desc.Index()
	var fun func(env *Env) r.Value
	switch upn {
	case 0:
		switch k {
		case r.Bool:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*bool)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Int:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Int8:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int8)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Int16:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int16)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Int32:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int32)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Int64:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int64)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Uint:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Uint8:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint8)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Uint16:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint16)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Uint32:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint32)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Uint64:
			fun = func(env *Env) r.Value {
				return r.ValueOf(&env.IntBinds[idx])
			}
		case r.Uintptr:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uintptr)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Float32:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*float32)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Float64:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*float64)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		case r.Complex64:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*complex64)(unsafe.Pointer(&env.IntBinds[idx])))
			}
		default:
			c.Errorf("unsupported variable type, cannot use for optimized addressof: %v <%T>", name, bind.Type)
			return nil
		}
	case 1:
		switch k {
		case r.Bool:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*bool)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Int:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Int8:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int8)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Int16:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int16)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Int32:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int32)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Int64:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int64)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Uint:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Uint8:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint8)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Uint16:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint16)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Uint32:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint32)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Uint64:
			fun = func(env *Env) r.Value {
				return r.ValueOf(&env.Outer.IntBinds[idx])
			}
		case r.Uintptr:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uintptr)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Float32:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*float32)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Float64:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*float64)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		case r.Complex64:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*complex64)(unsafe.Pointer(&env.Outer.IntBinds[idx])))
			}
		default:
			c.Errorf("unsupported variable type, cannot use for optimized addressof: %v <%T>", name, bind.Type)
			return nil
		}
	case 2:
		switch k {
		case r.Bool:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*bool)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Int:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Int8:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Int16:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Int32:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Int64:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*int64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Uint:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Uint8:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Uint16:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Uint32:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uint32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Uint64:
			fun = func(env *Env) r.Value {
				return r.ValueOf(&env.Outer.Outer.IntBinds[idx])
			}
		case r.Uintptr:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*uintptr)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Float32:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*float32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Float64:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*float64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Complex64:
			fun = func(env *Env) r.Value {
				return r.ValueOf((*complex64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		default:
			c.Errorf("unsupported variable type, cannot use for optimized addressof: %v <%T>", name, bind.Type)
			return nil
		}
	default:
		switch k {
		case r.Bool:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*bool)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Int:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*int)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Int8:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*int8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Int16:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*int16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Int32:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*int32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Int64:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*int64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Uint:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*uint)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Uint8:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*uint8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Uint16:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*uint16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Uint32:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*uint32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Uint64:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf(&env.Outer.Outer.IntBinds[idx])
			}
		case r.Uintptr:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*uintptr)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Float32:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*float32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Float64:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*float64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		case r.Complex64:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return r.ValueOf((*complex64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx])))
			}
		default:
			c.Errorf("unsupported variable type, cannot use for optimized addressof: %v <%T>", name, bind.Type)
			return nil
		}
	}
	return ExprX1(bind.Type, fun)
}
