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

package fast_interpreter

import (
	r "reflect"
	"unsafe"
)

func (c *Comp) Resolve(name string) (upn int, bind Bind) {
	ok := false
	upn, bind, ok = c.TryResolve(name)
	if !ok {
		c.Errorf("undefined identifier: %v", name)
	}
	return upn, bind
}

func (c *Comp) TryResolve(name string) (upn int, bind Bind, ok bool) {
	for ; c != nil; c = c.Outer {
		if b, ok := c.Binds[name]; ok {
			return upn, b, true
		}
		upn += c.UpCost // zero if *Comp has no local variables/functions so it will NOT have a corresponding *Env at runtime
	}
	return upn, bind, false
}

// Ident compiles a read operation on a constant, variable or function
func (c *Comp) Ident(name string) *Expr {
	upn, bind := c.Resolve(name)
	desc := bind.Desc
	switch desc.Class() {
	case ConstBind:
		return ExprLit(bind.Lit)
	case IntBind:
		return c.identIntBind(upn, bind)
	default:
		return c.identBind(upn, bind)
	}
}

func (c *Comp) identBind(upn int, bind Bind) *Expr {
	idx := bind.Desc.Index()
	var fun I
	switch bind.Type.Kind() {
	case r.Complex128:
		switch upn {
		case 0:
			fun = func(env *Env) complex128 {
				return env.Binds[idx].Complex()
			}
		case 1:
			fun = func(env *Env) complex128 {
				return env.Outer.Binds[idx].Complex()
			}
		case 2:
			fun = func(env *Env) complex128 {
				return env.Outer.Outer.Binds[idx].Complex()
			}
		default:
			fun = func(env *Env) complex128 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].Complex()
			}
		}
	case r.String:
		switch upn {
		case 0:
			fun = func(env *Env) string {
				return env.Binds[idx].String()
			}
		case 1:
			fun = func(env *Env) string {
				return env.Outer.Binds[idx].String()
			}
		case 2:
			fun = func(env *Env) string {
				return env.Outer.Outer.Binds[idx].String()
			}
		default:
			fun = func(env *Env) string {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx].String()
			}
		}
	default:
		switch upn {
		case 0:
			fun = func(env *Env) r.Value {
				return env.Binds[idx]
			}
		case 1:
			fun = func(env *Env) r.Value {
				return env.Outer.Binds[idx]
			}
		case 2:
			fun = func(env *Env) r.Value {
				return env.Outer.Outer.Binds[idx]
			}
		default:
			fun = func(env *Env) r.Value {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.Outer.Binds[idx]
			}
		}
	}
	return &Expr{Lit: Lit{Type: bind.Type}, Fun: fun}
}

func (c *Comp) identIntBind(upn int, bind Bind) *Expr {
	k := bind.Type.Kind()
	idx := bind.Desc.Index()
	switch upn {
	case 0:
		switch k {
		case r.Bool:
			return ExprBool(func(env *Env) bool {
				return *(*bool)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Int:
			return ExprInt(func(env *Env) int {
				return *(*int)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Int8:
			return ExprInt8(func(env *Env) int8 {
				return *(*int8)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Int16:
			return ExprInt16(func(env *Env) int16 {
				return *(*int16)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Int32:
			return ExprInt32(func(env *Env) int32 {
				return *(*int32)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Int64:
			return ExprInt64(func(env *Env) int64 {
				return *(*int64)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Uint:
			return ExprUint(func(env *Env) uint {
				return *(*uint)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Uint8:
			return ExprUint8(func(env *Env) uint8 {
				return *(*uint8)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Uint16:
			return ExprUint16(func(env *Env) uint16 {
				return *(*uint16)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Uint32:
			return ExprUint32(func(env *Env) uint32 {
				return *(*uint32)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Uint64:
			return ExprUint64(func(env *Env) uint64 {
				return env.IntBinds[idx]
			})
		case r.Uintptr:
			return ExprUintptr(func(env *Env) uintptr {
				return *(*uintptr)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Float32:
			return ExprFloat32(func(env *Env) float32 {
				return *(*float32)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Float64:
			return ExprFloat64(func(env *Env) float64 {
				return *(*float64)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		case r.Complex64:
			return ExprComplex64(func(env *Env) complex64 {
				return *(*complex64)(unsafe.Pointer(&env.IntBinds[idx]))
			})
		default:
			c.Errorf("unsupported variable type, cannot use for optimized read: <%v>", bind.Type)
			return nil
		}
	case 1:
		switch k {
		case r.Bool:
			return ExprBool(func(env *Env) bool {
				return *(*bool)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Int:
			return ExprInt(func(env *Env) int {
				return *(*int)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Int8:
			return ExprInt8(func(env *Env) int8 {
				return *(*int8)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Int16:
			return ExprInt16(func(env *Env) int16 {
				return *(*int16)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Int32:
			return ExprInt32(func(env *Env) int32 {
				return *(*int32)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Int64:
			return ExprInt64(func(env *Env) int64 {
				return *(*int64)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Uint:
			return ExprUint(func(env *Env) uint {
				return *(*uint)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Uint8:
			return ExprUint8(func(env *Env) uint8 {
				return *(*uint8)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Uint16:
			return ExprUint16(func(env *Env) uint16 {
				return *(*uint16)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Uint32:
			return ExprUint32(func(env *Env) uint32 {
				return *(*uint32)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Uint64:
			return ExprUint64(func(env *Env) uint64 {
				return env.Outer.IntBinds[idx]
			})
		case r.Uintptr:
			return ExprUintptr(func(env *Env) uintptr {
				return *(*uintptr)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Float32:
			return ExprFloat32(func(env *Env) float32 {
				return *(*float32)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Float64:
			return ExprFloat64(func(env *Env) float64 {
				return *(*float64)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		case r.Complex64:
			return ExprComplex64(func(env *Env) complex64 {
				return *(*complex64)(unsafe.Pointer(&env.Outer.IntBinds[idx]))
			})
		default:
			c.Errorf("unsupported variable type, cannot use for optimized read: <%v>", bind.Type)
			return nil
		}
	case 2:
		switch k {
		case r.Bool:
			return ExprBool(func(env *Env) bool {
				return *(*bool)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Int:
			return ExprInt(func(env *Env) int {
				return *(*int)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Int8:
			return ExprInt8(func(env *Env) int8 {
				return *(*int8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Int16:
			return ExprInt16(func(env *Env) int16 {
				return *(*int16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Int32:
			return ExprInt32(func(env *Env) int32 {
				return *(*int32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Int64:
			return ExprInt64(func(env *Env) int64 {
				return *(*int64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Uint:
			return ExprUint(func(env *Env) uint {
				return *(*uint)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Uint8:
			return ExprUint8(func(env *Env) uint8 {
				return *(*uint8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Uint16:
			return ExprUint16(func(env *Env) uint16 {
				return *(*uint16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Uint32:
			return ExprUint32(func(env *Env) uint32 {
				return *(*uint32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Uint64:
			return ExprUint64(func(env *Env) uint64 {
				return env.Outer.Outer.IntBinds[idx]
			})
		case r.Uintptr:
			return ExprUintptr(func(env *Env) uintptr {
				return *(*uintptr)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Float32:
			return ExprFloat32(func(env *Env) float32 {
				return *(*float32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Float64:
			return ExprFloat64(func(env *Env) float64 {
				return *(*float64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Complex64:
			return ExprComplex64(func(env *Env) complex64 {
				return *(*complex64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		default:
			c.Errorf("unsupported variable type, cannot use for optimized read: <%v>", bind.Type)
			return nil
		}
	default:
		switch k {
		case r.Bool:
			return ExprBool(func(env *Env) bool {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*bool)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Int:
			return ExprInt(func(env *Env) int {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*int)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Int8:
			return ExprInt8(func(env *Env) int8 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*int8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Int16:
			return ExprInt16(func(env *Env) int16 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*int16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Int32:
			return ExprInt32(func(env *Env) int32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*int32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Int64:
			return ExprInt64(func(env *Env) int64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*int64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Uint:
			return ExprUint(func(env *Env) uint {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*uint)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Uint8:
			return ExprUint8(func(env *Env) uint8 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*uint8)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Uint16:
			return ExprUint16(func(env *Env) uint16 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*uint16)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Uint32:
			return ExprUint32(func(env *Env) uint32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*uint32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Uint64:
			return ExprUint64(func(env *Env) uint64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return env.Outer.Outer.IntBinds[idx]
			})
		case r.Uintptr:
			return ExprUintptr(func(env *Env) uintptr {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*uintptr)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Float32:
			return ExprFloat32(func(env *Env) float32 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*float32)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Float64:
			return ExprFloat64(func(env *Env) float64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*float64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		case r.Complex64:
			return ExprComplex64(func(env *Env) complex64 {
				for i := 3; i < upn; i++ {
					env = env.Outer
				}
				return *(*complex64)(unsafe.Pointer(&env.Outer.Outer.IntBinds[idx]))
			})
		default:
			c.Errorf("unsupported variable type, cannot use for optimized read: <%v>", bind.Type)
			return nil
		}
	}
}
