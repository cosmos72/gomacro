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

// AssignVar0 compiles an assignment to a variable
func (c *Comp) AssignVar0(name string, bind Bind, init *Expr) {
	if init.Const() {
		c.assignVar0Const(name, bind, init)
	} else {
		c.assignVar0Expr(name, bind, init)
	}
}

// AssignVar0Const compiles an assignment to a variable with a value known at compile time
func (c *Comp) assignVar0Const(name string, bind Bind, init *Expr) {
	t := bind.Type
	if init.Type != t {
		init.ConstTo(t)
	}
	desc := bind.Desc
	var ret func(env *Env) (Stmt, *Env)
	switch desc.Class() {
	default:
		c.Errorf("cannot assign to %v", name)
		return
	case VarBind:
		index := desc.Index()
		if index == NoIndex {
			// assigning a value to _ has no effect at all
			return
		}
		v := r.ValueOf(init.Value)
		ret = func(env *Env) (Stmt, *Env) {
			env.Binds[index] = v
			env.IP++
			return env.Code[env.IP], env
		}
	case IntBind:
		index := desc.Index()
		if index == NoIndex {
			// assigning a value to _ has no effect at all
			return
		}
		switch value := init.Value.(type) {
		case bool:
			ret = func(env *Env) (Stmt, *Env) {
				*(*bool)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case int:
			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case int8:
			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case int16:
			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case int32:
			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case int64:
			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case uint:
			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case uint8:
			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case uint16:
			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case uint32:
			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case uint64:
			ret = func(env *Env) (Stmt, *Env) {
				env.IntBinds[index] = value
				env.IP++
				return env.Code[env.IP], env
			}
		case uintptr:
			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case float32:
			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case float64:
			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		case complex64:
			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.IntBinds[index])) = value
				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf("unsupported constant type, cannot use for optimized assignment: %v <%T>", value, value)
			return
		}
	}
	c.Code.Append(Stmt{ret})
}

// AssignVar0Expr compiles an assignment to a variable with an expression
func (c *Comp) assignVar0Expr(name string, bind Bind, init *Expr) {
	t := bind.Type
	if init.Type != t && !init.Type.AssignableTo(t) {
		c.Errorf("cannot assign <%v> to <%v>", init.Type, t)
		return
	}
	desc := bind.Desc
	var ret func(env *Env) (Stmt, *Env)
	switch desc.Class() {
	default:
		c.Errorf("cannot assign to %v", name)
		return
	case VarBind:
		index := desc.Index()
		if index == NoIndex {
			// assigning an expression to _
			// only keep the expression side effects
			c.Code.Append(init.AsStmt())
			return
		}
		fun := init.AsX1()
		ret = func(env *Env) (Stmt, *Env) {
			env.Binds[index].Set(fun(env).Convert(t))
			env.IP++
			return env.Code[env.IP], env
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
	c.Code.Append(Stmt{ret})
}

// AssignVar0Value compiles an assignment to a variable with a reflect.Value passed at runtime.
// Used to initialize variables with multi-valued expressions
func (c *Comp) AssignVar0Value(name string, bind Bind) func(*Env, r.Value) {
	desc := bind.Desc
	t := bind.Type
	switch desc.Class() {
	default:
		c.Errorf("cannot assign to %v", name)
		return nil
	case VarBind:
		index := desc.Index()
		if index == NoIndex {
			// assigning a value to _ has no effect at all
			return nil
		}
		return func(env *Env, v r.Value) {
			env.Binds[index].Set(v.Convert(t))
		}
	case IntBind:
		index := desc.Index()
		if index == NoIndex {
			// assigning a value to _ has no effect at all
			return nil
		}
		switch t.Kind() {
		case r.Bool:
			return func(env *Env, v r.Value) {
				*(*bool)(unsafe.Pointer(&env.IntBinds[index])) = v.Bool()
			}
		case r.Int:
			return func(env *Env, v r.Value) {
				*(*int)(unsafe.Pointer(&env.IntBinds[index])) = int(v.Int())
			}
		case r.Int8:
			return func(env *Env, v r.Value) {
				*(*int8)(unsafe.Pointer(&env.IntBinds[index])) = int8(v.Int())
			}
		case r.Int16:
			return func(env *Env, v r.Value) {
				*(*int16)(unsafe.Pointer(&env.IntBinds[index])) = int16(v.Int())
			}
		case r.Int32:
			return func(env *Env, v r.Value) {
				*(*int32)(unsafe.Pointer(&env.IntBinds[index])) = int32(v.Int())
			}
		case r.Int64:
			return func(env *Env, v r.Value) {
				*(*int64)(unsafe.Pointer(&env.IntBinds[index])) = v.Int()
			}
		case r.Uint:
			return func(env *Env, v r.Value) {
				*(*uint)(unsafe.Pointer(&env.IntBinds[index])) = uint(v.Uint())
			}
		case r.Uint8:
			return func(env *Env, v r.Value) {
				*(*uint8)(unsafe.Pointer(&env.IntBinds[index])) = uint8(v.Uint())
			}
		case r.Uint16:
			return func(env *Env, v r.Value) {
				*(*uint16)(unsafe.Pointer(&env.IntBinds[index])) = uint16(v.Uint())
			}
		case r.Uint32:
			return func(env *Env, v r.Value) {
				*(*uint32)(unsafe.Pointer(&env.IntBinds[index])) = uint32(v.Uint())
			}
		case r.Uint64:
			return func(env *Env, v r.Value) {
				env.IntBinds[index] = v.Uint()
			}
		case r.Uintptr:
			return func(env *Env, v r.Value) {
				*(*uintptr)(unsafe.Pointer(&env.IntBinds[index])) = uintptr(v.Uint())
			}
		case r.Float32:
			return func(env *Env, v r.Value) {
				*(*float32)(unsafe.Pointer(&env.IntBinds[index])) = float32(v.Float())
			}
		case r.Float64:
			return func(env *Env, v r.Value) {
				*(*float64)(unsafe.Pointer(&env.IntBinds[index])) = v.Float()
			}
		case r.Complex64:
			return func(env *Env, v r.Value) {
				*(*complex64)(unsafe.Pointer(&env.IntBinds[index])) = complex64(v.Complex())
			}
		default:
			c.Errorf("unsupported type, cannot use for optimized assignment: <%v>", t)
			return nil
		}
	}
}
