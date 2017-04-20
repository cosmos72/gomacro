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
 * call0ret1.go
 *
 *  Created on Apr 20, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	. "github.com/cosmos72/gomacro/base"
)

func call0ret1(c *Call, maxdepth int) I {
	expr := c.Fun
	exprfun := expr.AsX1()
	funvar := c.Funvar
	var funupn, funindex int
	if funvar != nil {
		funupn = funvar.Upn
		funindex = funvar.Desc.Index()
	}
	kret := expr.Type.Out(0).Kind()
	var cachedfunv r.Value
	var call I
	switch kret {
	case r.Bool:

		{
			var cachedfun func() bool

			if funvar == nil {
				call = func(env *Env) bool {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() bool)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) bool {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() bool)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) bool {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() bool)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) bool {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() bool)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) bool {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() bool)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) bool {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() bool)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) bool {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() bool)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Int:

		{
			var cachedfun func() int

			if funvar == nil {
				call = func(env *Env) int {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() int)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) int {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) int {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) int {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) int {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) int {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) int {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Int8:

		{
			var cachedfun func() int8

			if funvar == nil {
				call = func(env *Env) int8 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() int8)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) int8 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int8)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) int8 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int8)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) int8 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int8)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) int8 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int8)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) int8 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int8)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) int8 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int8)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Int16:

		{
			var cachedfun func() int16

			if funvar == nil {
				call = func(env *Env) int16 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() int16)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) int16 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int16)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) int16 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int16)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) int16 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int16)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) int16 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int16)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) int16 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int16)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) int16 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int16)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Int32:

		{
			var cachedfun func() int32

			if funvar == nil {
				call = func(env *Env) int32 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() int32)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) int32 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int32)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) int32 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int32)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) int32 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int32)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) int32 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int32)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) int32 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int32)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) int32 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int32)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Int64:

		{
			var cachedfun func() int64

			if funvar == nil {
				call = func(env *Env) int64 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() int64)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) int64 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int64)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) int64 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int64)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) int64 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int64)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) int64 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int64)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) int64 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int64)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) int64 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() int64)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Uint:

		{
			var cachedfun func() uint

			if funvar == nil {
				call = func(env *Env) uint {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() uint)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) uint {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) uint {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) uint {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) uint {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) uint {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) uint {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Uint8:
		{
			var cachedfun func() uint8

			if funvar == nil {
				call = func(env *Env) uint8 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() uint8)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) uint8 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint8)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) uint8 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint8)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) uint8 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint8)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) uint8 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint8)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) uint8 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint8)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) uint8 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint8)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Uint16:
		{
			var cachedfun func() uint16

			if funvar == nil {
				call = func(env *Env) uint16 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() uint16)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) uint16 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint16)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) uint16 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint16)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) uint16 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint16)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) uint16 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint16)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) uint16 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint16)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) uint16 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint16)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Uint32:
		{
			var cachedfun func() uint32
			if funvar == nil {
				call = func(env *Env) uint32 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() uint32)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) uint32 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint32)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) uint32 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint32)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) uint32 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint32)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) uint32 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint32)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) uint32 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint32)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) uint32 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint32)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Uint64:
		{
			var cachedfun func() uint64
			if funvar == nil {
				call = func(env *Env) uint64 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() uint64)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) uint64 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint64)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) uint64 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint64)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) uint64 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint64)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) uint64 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint64)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) uint64 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint64)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) uint64 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uint64)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Uintptr:
		{
			var cachedfun func() uintptr
			if funvar == nil {
				call = func(env *Env) uintptr {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() uintptr)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) uintptr {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uintptr)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) uintptr {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uintptr)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) uintptr {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uintptr)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) uintptr {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uintptr)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) uintptr {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uintptr)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) uintptr {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() uintptr)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Float32:
		{
			var cachedfun func() float32
			if funvar == nil {
				call = func(env *Env) float32 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() float32)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) float32 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float32)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) float32 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float32)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) float32 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float32)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) float32 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float32)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) float32 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float32)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) float32 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float32)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Float64:
		{
			var cachedfun func() float64
			if funvar == nil {
				call = func(env *Env) float64 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() float64)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) float64 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float64)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) float64 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float64)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) float64 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float64)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) float64 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float64)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) float64 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float64)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) float64 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() float64)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Complex64:
		{
			var cachedfun func() complex64
			if funvar == nil {
				call = func(env *Env) complex64 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() complex64)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) complex64 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex64)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) complex64 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex64)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) complex64 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex64)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) complex64 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex64)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) complex64 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex64)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) complex64 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex64)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.Complex128:
		{
			var cachedfun func() complex128
			if funvar == nil {
				call = func(env *Env) complex128 {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() complex128)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) complex128 {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex128)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) complex128 {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex128)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) complex128 {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex128)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) complex128 {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex128)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) complex128 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex128)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) complex128 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() complex128)
						}
						return cachedfun()
					}
				}
			}

		}
	case r.String:
		{
			var cachedfun func() string
			if funvar == nil {
				call = func(env *Env) string {
					funv := exprfun(env)
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func() string)
					}
					return cachedfun()
				}
			} else {
				switch funupn {
				case 0:
					call = func(env *Env) string {
						funv := env.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() string)
						}
						return cachedfun()
					}
				case 1:
					call = func(env *Env) string {
						funv := env.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() string)
						}
						return cachedfun()
					}
				case 2:
					call = func(env *Env) string {
						funv := env.Outer.Outer.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() string)
						}
						return cachedfun()
					}
				default:
					call = func(env *Env) string {
						o := env
						for i := 0; i < funupn; i++ {
							o = o.Outer
						}

						funv := o.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() string)
						}
						return cachedfun()
					}
				case maxdepth - 1:
					call = func(env *Env) string {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() string)
						}
						return cachedfun()
					}
				case maxdepth:
					call = func(env *Env) string {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func() string)
						}
						return cachedfun()
					}
				}
			}

		}
	default:
		call = func(env *Env) r.Value {
			funv := exprfun(env)
			return funv.Call(ZeroValues)[0]
		}

	}
	return call
}
