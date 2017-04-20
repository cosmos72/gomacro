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
 * call_ret1.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
)

func call2ret1(c *Call, maxdepth int) I {
	expr := c.Fun
	exprfun := expr.AsX1()
	karg0 := expr.Type.In(0).Kind()
	karg1 := expr.Type.In(1).Kind()
	kret := expr.Type.Out(0).Kind()
	args := c.Args
	argfuns := c.Argfuns
	var cachedfunv r.Value
	var call I
	switch kret {
	case r.Bool:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) bool

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) bool)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) bool)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) bool)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) bool {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) bool)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) bool {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Bool()
				}
			}

		}
	case r.Int:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) int

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) int {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return int(ret0.Int())
				}
			}

		}
	case r.Int8:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) int8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) int8 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return int8(ret0.Int())
				}
			}

		}
	case r.Int16:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) int16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) int16 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return int16(ret0.Int())
				}
			}

		}
	case r.Int32:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) int32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) int32 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return int32(ret0.Int())
				}
			}

		}
	case r.Int64:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) int64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) int64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) int64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) int64 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Int()
				}
			}

		}
	case r.Uint:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) uint

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uint {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return uint(ret0.Uint())
				}
			}

		}
	case r.Uint8:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) uint8

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint8)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint8)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint8)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint8 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint8)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uint8 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return uint8(ret0.Uint())
				}
			}

		}
	case r.Uint16:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) uint16

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint16)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint16)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint16)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint16 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint16)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uint16 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return uint16(ret0.Uint())
				}
			}

		}
	case r.Uint32:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) uint32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uint32 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return uint32(ret0.Uint())
				}
			}

		}
	case r.Uint64:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) uint64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uint64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uint64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uint64 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Uint()
				}
			}

		}
	case r.Uintptr:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) uintptr

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uintptr)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uintptr)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uintptr)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) uintptr {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) uintptr)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) uintptr {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return uintptr(ret0.Uint())
				}
			}

		}
	case r.Float32:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) float32

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) float32)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) float32)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) float32)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) float32 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) float32)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) float32 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return float32(ret0.Float())
				}
			}

		}
	case r.Float64:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) float64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) float64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) float64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) float64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) float64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) float64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) float64 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Float()
				}
			}

		}
	case r.Complex64:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) complex64

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) complex64)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) complex64)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) complex64)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) complex64 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) complex64)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) complex64 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return complex64(ret0.Complex())
				}
			}

		}
	case r.Complex128:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) complex128

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) complex128)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) complex128)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) complex128)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) complex128 {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) complex128)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) complex128 {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Complex()
				}
			}

		}
	case r.String:
		{
			if karg0 == karg1 {
				switch karg0 {
				case r.Bool:

					{
						var cachedfun func(bool, bool) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) bool)
							if arg1.Const() {
								arg1const := args[1].Value.(bool)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) bool)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(bool, bool) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int:

					{
						var cachedfun func(int, int) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int)
							if arg1.Const() {
								arg1const := args[1].Value.(int)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int, int) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int8:

					{
						var cachedfun func(int8, int8) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int8)
							if arg1.Const() {
								arg1const := args[1].Value.(int8)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int8)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int8, int8) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int16:

					{
						var cachedfun func(int16, int16) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int16)
							if arg1.Const() {
								arg1const := args[1].Value.(int16)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int16)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int16, int16) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int32:

					{
						var cachedfun func(int32, int32) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int32)
							if arg1.Const() {
								arg1const := args[1].Value.(int32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int32, int32) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Int64:

					{
						var cachedfun func(int64, int64) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) int64)
							if arg1.Const() {
								arg1const := args[1].Value.(int64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) int64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(int64, int64) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint:

					{
						var cachedfun func(uint, uint) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint)
							if arg1.Const() {
								arg1const := args[1].Value.(uint)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint, uint) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint8:

					{
						var cachedfun func(uint8, uint8) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint8)
							if arg1.Const() {
								arg1const := args[1].Value.(uint8)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint8)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint8, uint8) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint16:

					{
						var cachedfun func(uint16, uint16) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint16)
							if arg1.Const() {
								arg1const := args[1].Value.(uint16)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint16)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint16, uint16) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint32:

					{
						var cachedfun func(uint32, uint32) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint32)
							if arg1.Const() {
								arg1const := args[1].Value.(uint32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint32, uint32) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uint64:

					{
						var cachedfun func(uint64, uint64) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uint64)
							if arg1.Const() {
								arg1const := args[1].Value.(uint64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uint64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uint64, uint64) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Uintptr:

					{
						var cachedfun func(uintptr, uintptr) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) uintptr)
							if arg1.Const() {
								arg1const := args[1].Value.(uintptr)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) uintptr)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(uintptr, uintptr) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float32:

					{
						var cachedfun func(float32, float32) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float32)
							if arg1.Const() {
								arg1const := args[1].Value.(float32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float32)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float32, float32) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Float64:

					{
						var cachedfun func(float64, float64) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) float64)
							if arg1.Const() {
								arg1const := args[1].Value.(float64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) float64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(float64, float64) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex64:

					{
						var cachedfun func(complex64, complex64) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex64)
							if arg1.Const() {
								arg1const := args[1].Value.(complex64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex64)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex64, complex64) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.Complex128:

					{
						var cachedfun func(complex128, complex128) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) complex128)
							if arg1.Const() {
								arg1const := args[1].Value.(complex128)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) complex128)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(complex128, complex128) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				case r.String:

					{
						var cachedfun func(string, string) string

						arg0 := args[0]
						arg1 := args[1]
						if arg0.Const() {
							arg0const := args[0].Value.(string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) string)
									}
									return cachedfun(arg0const, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) string)
									}

									arg1 := arg1fun(env)
									return cachedfun(arg0const, arg1)
								}
							}
						} else {
							arg0fun := args[0].Fun.(func(env *Env) string)
							if arg1.Const() {
								arg1const := args[1].Value.(string)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) string)
									}

									arg0 := arg0fun(env)
									return cachedfun(arg0, arg1const)
								}
							} else {
								arg1fun := args[1].Fun.(func(env *Env) string)
								call = func(env *Env) string {
									funv := exprfun(env)
									if cachedfunv != funv {
										cachedfunv = funv
										cachedfun = funv.Interface().(func(string, string) string)
									}

									arg0 := arg0fun(env)
									arg1 := arg1fun(env)
									return cachedfun(arg0, arg1)
								}
							}
						}
					}
				}
			}

			if call == nil {
				call = func(env *Env) string {
					funv := exprfun(env)
					argv := []r.Value{
						argfuns[0](env),
						argfuns[1](env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.String()
				}
			}

		}
	default:
		call = func(env *Env) r.Value {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			return funv.Call(argv)[0]
		}

	}
	return call
}
