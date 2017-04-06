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
 * declaration.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	r "reflect"
	"unsafe"
)

// AssignVar compiles an assignment to a variable
func (c *Comp) AssignVar0(name string, desc BindDescriptor, t r.Type, init *Expr) X {
	if init.Const() {
		return c.AssignVar0Const(name, desc, t, init)
	} else {
		return c.AssignVar0Expr(name, desc, t, init)
	}
}

func (c *Comp) AssignVar0Const(name string, desc BindDescriptor, t r.Type, init *Expr) X {
	if init.Type != t {
		init.ConstTo(t)
	}
	switch desc.Class() {
	case NoBind:
		// assigning a constant to _ has no effect at all
		return nil
	default: // case FuncBind:
		c.Errorf("cannot assign to %v", name)
		return nil
	case VarBind:
		index := desc.Index()
		v := r.ValueOf(init.Value)
		return func(env *Env) {
			env.Binds[index] = v
		}
	case IntBind:
		index := desc.Index()
		switch value := init.Value.(type) {
		case bool:
			return func(env *Env) {
				*(*bool)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case int:
			return func(env *Env) {
				*(*int)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case int8:
			return func(env *Env) {
				*(*int8)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case int16:
			return func(env *Env) {
				*(*int16)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case int32:
			return func(env *Env) {
				*(*int32)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case int64:
			return func(env *Env) {
				*(*int64)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case uint:
			return func(env *Env) {
				*(*uint)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case uint8:
			return func(env *Env) {
				*(*uint8)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case uint16:
			return func(env *Env) {
				*(*uint16)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case uint32:
			return func(env *Env) {
				*(*uint32)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case uint64:
			return func(env *Env) {
				env.IntBinds[index] = value
			}
		case uintptr:
			return func(env *Env) {
				*(*uintptr)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case float32:
			return func(env *Env) {
				*(*float32)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case float64:
			return func(env *Env) {
				*(*float64)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		case complex64:
			return func(env *Env) {
				*(*complex64)(unsafe.Pointer(&env.IntBinds[index])) = value
			}
		default:
			c.Errorf("unsupported constant type, cannot use for optimized assignment: %v <%T>", value, value)
			return nil
		}
	}
}

func (c *Comp) AssignVar0Expr(name string, desc BindDescriptor, t r.Type, init *Expr) X {
	if init.Type != t && !init.Type.AssignableTo(t) {
		c.Errorf("cannot assign <%v> to <%v>", init.Type, t)
	}

	switch desc.Class() {
	case NoBind:
		// assigning an expression to _
		// only keep the expression side effects
		return init.AsX()
	default: // case FuncBind:
		c.Errorf("cannot assign to %v", name)
		return nil
	case VarBind:
		index := desc.Index()
		fun := init.AsX1()
		if init.Type == t {
			return func(env *Env) {
				env.Binds[index] = fun(env)
			}
		} else {
			return func(env *Env) {
				env.Binds[index] = fun(env).Convert(t)
			}
		}
	case IntBind:
		index := desc.Index()
		switch fun := init.Fun.(type) {
		case func(*Env) bool:
			return func(env *Env) {
				*(*bool)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) int:
			return func(env *Env) {
				*(*int)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) int8:
			return func(env *Env) {
				*(*int8)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) int16:
			return func(env *Env) {
				*(*int16)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) int32:
			return func(env *Env) {
				*(*int32)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) int64:
			return func(env *Env) {
				*(*int64)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) uint:
			return func(env *Env) {
				*(*uint)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) uint8:
			return func(env *Env) {
				*(*uint8)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) uint16:
			return func(env *Env) {
				*(*uint16)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) uint32:
			return func(env *Env) {
				*(*uint32)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) uint64:
			return func(env *Env) {
				env.IntBinds[index] = fun(env)
			}
		case func(*Env) uintptr:
			return func(env *Env) {
				*(*uintptr)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) float32:
			return func(env *Env) {
				*(*float32)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) float64:
			return func(env *Env) {
				*(*float64)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		case func(*Env) complex64:
			return func(env *Env) {
				*(*complex64)(unsafe.Pointer(&env.IntBinds[index])) = fun(env)
			}
		default:
			c.Errorf("unsupported expression type, cannot use for optimized assignment: %v <%T> returns %v",
				fun, fun, init.Outs())
			return nil
		}
	}
}
