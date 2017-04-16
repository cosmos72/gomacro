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
 * call_ret0.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	r "reflect"
)

func call_ret0(expr *Expr, args []*Expr, argfuns []func(*Env) r.Value) I {
	exprfun := expr.AsX1()
	var call I

	switch expr.Type.NumIn() {
	case 0:
		call = func(env *Env) {
			fun := exprfun(env).Interface().(func())
			fun()
		}
	case 1:
		t := expr.Type.In(0)
		switch t.Kind() {
		case r.Bool:

			if args[0].Const() {
				argconst := args[0].Value.(bool)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(bool))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) bool)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(bool))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Int:

			if args[0].Const() {
				argconst := args[0].Value.(int)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) int)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Int8:

			if args[0].Const() {
				argconst := args[0].Value.(int8)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int8))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) int8)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int8))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Int16:
			if args[0].Const() {
				argconst := args[0].Value.(int16)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int16))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) int16)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int16))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Int32:
			if args[0].Const() {
				argconst := args[0].Value.(int32)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int32))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) int32)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int32))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Int64:
			if args[0].Const() {
				argconst := args[0].Value.(int64)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int64))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) int64)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int64))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Uint:
			if args[0].Const() {
				argconst := args[0].Value.(uint)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) uint)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Uint8:
			if args[0].Const() {
				argconst := args[0].Value.(uint8)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint8))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) uint8)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint8))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Uint16:
			if args[0].Const() {
				argconst := args[0].Value.(uint16)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint16))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) uint16)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint16))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Uint32:
			if args[0].Const() {
				argconst := args[0].Value.(uint32)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint32))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) uint32)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint32))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Uint64:
			if args[0].Const() {
				argconst := args[0].Value.(uint64)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint64))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) uint64)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint64))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Uintptr:
			if args[0].Const() {
				argconst := args[0].Value.(uintptr)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uintptr))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) uintptr)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uintptr))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Float32:
			if args[0].Const() {
				argconst := args[0].Value.(float32)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(float32))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) float32)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(float32))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Float64:
			if args[0].Const() {
				argconst := args[0].Value.(float64)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(float64))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) float64)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(float64))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Complex64:
			if args[0].Const() {
				argconst := args[0].Value.(complex64)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(complex64))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) complex64)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(complex64))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.Complex128:
			if args[0].Const() {
				argconst := args[0].Value.(complex128)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(complex128))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) complex128)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(complex128))
					arg := argfun(env)

					fun(arg)
				}
			}
		case r.String:
			if args[0].Const() {
				argconst := args[0].Value.(string)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(string))

					fun(argconst)
				}
			} else {
				argfun := args[0].Fun.(func(env *Env) string)
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(string))
					arg := argfun(env)

					fun(arg)
				}
			}
		default:
			{
				argfun := argfuns[0]
				call = func(env *Env) {
					funv := exprfun(env)

					argv := []r.Value{
						argfun(env),
					}
					funv.Call(argv)
				}
			}

		}
	case 2:
		if expr.Type.In(0) == expr.Type.In(1) {
			t := expr.Type.In(0)
			switch t.Kind() {
			case r.Bool:
				if args[0].Const() {
					arg0const := args[0].Value.(bool)
					if args[1].Const() {
						arg1const := args[1].Value.(bool)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(bool, bool))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) bool)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(bool, bool))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) bool)
					if args[1].Const() {
						arg1const := args[1].Value.(bool)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(bool, bool))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) bool)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(bool, bool))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Int:
				if args[0].Const() {
					arg0const := args[0].Value.(int)
					if args[1].Const() {
						arg1const := args[1].Value.(int)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int, int))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) int)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int, int))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) int)
					if args[1].Const() {
						arg1const := args[1].Value.(int)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int, int))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) int)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int, int))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Int8:
				if args[0].Const() {
					arg0const := args[0].Value.(int8)
					if args[1].Const() {
						arg1const := args[1].Value.(int8)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int8, int8))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) int8)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int8, int8))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) int8)
					if args[1].Const() {
						arg1const := args[1].Value.(int8)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int8, int8))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) int8)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int8, int8))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Int16:
				if args[0].Const() {
					arg0const := args[0].Value.(int16)
					if args[1].Const() {
						arg1const := args[1].Value.(int16)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int16, int16))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) int16)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int16, int16))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) int16)
					if args[1].Const() {
						arg1const := args[1].Value.(int16)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int16, int16))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) int16)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int16, int16))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Int32:
				if args[0].Const() {
					arg0const := args[0].Value.(int32)
					if args[1].Const() {
						arg1const := args[1].Value.(int32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int32, int32))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) int32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int32, int32))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) int32)
					if args[1].Const() {
						arg1const := args[1].Value.(int32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int32, int32))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) int32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int32, int32))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Int64:
				if args[0].Const() {
					arg0const := args[0].Value.(int64)
					if args[1].Const() {
						arg1const := args[1].Value.(int64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int64, int64))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) int64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int64, int64))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) int64)
					if args[1].Const() {
						arg1const := args[1].Value.(int64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int64, int64))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) int64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(int64, int64))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Uint:
				if args[0].Const() {
					arg0const := args[0].Value.(uint)
					if args[1].Const() {
						arg1const := args[1].Value.(uint)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint, uint))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uint)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint, uint))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) uint)
					if args[1].Const() {
						arg1const := args[1].Value.(uint)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint, uint))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uint)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint, uint))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Uint8:
				if args[0].Const() {
					arg0const := args[0].Value.(uint8)
					if args[1].Const() {
						arg1const := args[1].Value.(uint8)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint8, uint8))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uint8)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint8, uint8))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) uint8)
					if args[1].Const() {
						arg1const := args[1].Value.(uint8)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint8, uint8))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uint8)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint8, uint8))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Uint16:
				if args[0].Const() {
					arg0const := args[0].Value.(uint16)
					if args[1].Const() {
						arg1const := args[1].Value.(uint16)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint16, uint16))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uint16)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint16, uint16))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) uint16)
					if args[1].Const() {
						arg1const := args[1].Value.(uint16)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint16, uint16))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uint16)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint16, uint16))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Uint32:
				if args[0].Const() {
					arg0const := args[0].Value.(uint32)
					if args[1].Const() {
						arg1const := args[1].Value.(uint32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint32, uint32))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uint32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint32, uint32))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) uint32)
					if args[1].Const() {
						arg1const := args[1].Value.(uint32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint32, uint32))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uint32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint32, uint32))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Uint64:
				if args[0].Const() {
					arg0const := args[0].Value.(uint64)
					if args[1].Const() {
						arg1const := args[1].Value.(uint64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint64, uint64))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uint64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint64, uint64))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) uint64)
					if args[1].Const() {
						arg1const := args[1].Value.(uint64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint64, uint64))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uint64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uint64, uint64))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Uintptr:
				if args[0].Const() {
					arg0const := args[0].Value.(uintptr)
					if args[1].Const() {
						arg1const := args[1].Value.(uintptr)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uintptr, uintptr))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uintptr)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uintptr, uintptr))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) uintptr)
					if args[1].Const() {
						arg1const := args[1].Value.(uintptr)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uintptr, uintptr))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) uintptr)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(uintptr, uintptr))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Float32:
				if args[0].Const() {
					arg0const := args[0].Value.(float32)
					if args[1].Const() {
						arg1const := args[1].Value.(float32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(float32, float32))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) float32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(float32, float32))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) float32)
					if args[1].Const() {
						arg1const := args[1].Value.(float32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(float32, float32))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) float32)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(float32, float32))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Float64:
				if args[0].Const() {
					arg0const := args[0].Value.(float64)
					if args[1].Const() {
						arg1const := args[1].Value.(float64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(float64, float64))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) float64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(float64, float64))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) float64)
					if args[1].Const() {
						arg1const := args[1].Value.(float64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(float64, float64))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) float64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(float64, float64))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Complex64:
				if args[0].Const() {
					arg0const := args[0].Value.(complex64)
					if args[1].Const() {
						arg1const := args[1].Value.(complex64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(complex64, complex64))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) complex64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(complex64, complex64))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) complex64)
					if args[1].Const() {
						arg1const := args[1].Value.(complex64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(complex64, complex64))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) complex64)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(complex64, complex64))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.Complex128:
				if args[0].Const() {
					arg0const := args[0].Value.(complex128)
					if args[1].Const() {
						arg1const := args[1].Value.(complex128)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(complex128, complex128))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) complex128)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(complex128, complex128))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) complex128)
					if args[1].Const() {
						arg1const := args[1].Value.(complex128)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(complex128, complex128))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) complex128)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(complex128, complex128))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}
			case r.String:
				if args[0].Const() {
					arg0const := args[0].Value.(string)
					if args[1].Const() {
						arg1const := args[1].Value.(string)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(string, string))
							fun(arg0const, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) string)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(string, string))
							arg1 := arg1fun(env)
							fun(arg0const, arg1)
						}
					}
				} else {
					arg0fun := args[0].Fun.(func(env *Env) string)
					if args[1].Const() {
						arg1const := args[1].Value.(string)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(string, string))
							arg0 := arg0fun(env)
							fun(arg0, arg1const)
						}
					} else {
						arg1fun := args[1].Fun.(func(env *Env) string)
						call = func(env *Env) {
							fun := exprfun(env).Interface().(func(string, string))
							arg0 := arg0fun(env)
							arg1 := arg1fun(env)
							fun(arg0, arg1)
						}
					}
				}

			}
		}
		if call == nil {
			call = func(env *Env) {
				funv := exprfun(env)

				argv := []r.Value{
					argfuns[0](env),
					argfuns[1](env),
				}
				funv.Call(argv)
			}
		}

	case 3:
		call = func(env *Env) {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
				argfuns[2](env),
			}
			funv.Call(argv)
		}
	default:
		call = func(env *Env) {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}

			funv.Call(argv)
		}
	}
	return call
}
