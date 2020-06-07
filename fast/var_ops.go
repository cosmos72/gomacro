// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * var_setops.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/token"
	r "reflect"
	"unsafe"

	"github.com/cosmos72/gomacro/base/reflect"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) varAddConst(va *Var, val I) Stmt {
	if isLiteralNumber(val, 0) || val == "" {
		return nil
	}

	t := va.Type

	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		val := int(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		val := int8(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		val := int16(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		val := int32(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		val := xr.ValueOf(val).Int()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() + int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		val := uint(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		val := uint8(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		val := uint16(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		val := uint32(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		val := xr.ValueOf(val).Uint()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		val := uintptr(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Float32:
		val :=

			float32(xr.ValueOf(val).Float())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Float64:
		val := xr.ValueOf(val).Float()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Complex64:
		val :=

			complex64(xr.ValueOf(val).Complex())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Complex128:
		val := xr.ValueOf(val).Complex()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.FileEnv.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex128)(unsafe.Pointer(&o.Ints[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.String:
		val := xr.ValueOf(val).String()
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Vals[index]
					lhs.SetString(lhs.String() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Vals[index]
					lhs.SetString(lhs.String() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Vals[index]
					lhs.SetString(lhs.String() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.FileEnv.
						Vals[index]
					lhs.SetString(lhs.String() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Vals[index]
					lhs.SetString(lhs.String() +
						val,
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.ADD, t)

	}
	return ret
}
func (c *Comp) varAddExpr(va *Var, fun I) Stmt {

	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() + int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() + uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() + float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.FileEnv.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex128)(unsafe.Pointer(&o.Ints[index])) += fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() + complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) string:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Vals[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Vals[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Vals[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.FileEnv.
						Vals[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Vals[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.ADD, t)

	}
	return ret
}
func (c *Comp) varSubConst(va *Var, val I) Stmt {
	if isLiteralNumber(val, 0) {
		return nil
	}

	t := va.Type

	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		val := int(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		val := int8(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		val := int16(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		val := int32(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		val := xr.ValueOf(val).Int()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() - int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		val := uint(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		val := uint8(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		val := uint16(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		val := uint32(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		val := xr.ValueOf(val).Uint()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		val := uintptr(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Float32:
		val :=

			float32(xr.ValueOf(val).Float())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Float64:
		val := xr.ValueOf(val).Float()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Complex64:
		val :=

			complex64(xr.ValueOf(val).Complex())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Complex128:
		val := xr.ValueOf(val).Complex()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex128)(unsafe.Pointer(&o.Ints[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.SUB, t)

	}
	return ret
}
func (c *Comp) varSubExpr(va *Var, fun I) Stmt {

	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() - int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() - uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() - float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.FileEnv.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex128)(unsafe.Pointer(&o.Ints[index])) -= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() - complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.SUB, t)

	}
	return ret
}
func (c *Comp) varMulConst(va *Var, val I) Stmt {
	if isLiteralNumber(val, 0) {
		return c.varSetZero(va)
	} else if isLiteralNumber(val, 1) {
		return nil
	}

	t := va.Type

	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		val := int(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		val := int8(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		val := int16(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		val := int32(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		val := xr.ValueOf(val).Int()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() * int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		val := uint(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		val := uint8(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		val := uint16(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		val := uint32(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		val := xr.ValueOf(val).Uint()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		val := uintptr(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Float32:
		val :=

			float32(xr.ValueOf(val).Float())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Float64:
		val := xr.ValueOf(val).Float()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Complex64:
		val :=

			complex64(xr.ValueOf(val).Complex())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Complex128:
		val := xr.ValueOf(val).Complex()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex128)(unsafe.Pointer(&o.Ints[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.MUL, t)

	}
	return ret
}
func (c *Comp) varMulExpr(va *Var, fun I) Stmt {

	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() * int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() * uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() * float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.FileEnv.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex128)(unsafe.Pointer(&o.Ints[index])) *= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() * complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.MUL, t)

	}
	return ret
}
func (c *Comp) varQuoPow2(va *Var, val I) Stmt {
	t := va.Type
	if isLiteralNumber(val, 0) {
		c.Errorf("division by %v <%v>", val, t)
		return nil
	} else if isLiteralNumber(val, 1) {
		return nil
	}

	ypositive := true
	yv := xr.ValueOf(val)
	var y uint64
	switch reflect.Category(yv.Kind()) {
	case xr.Int:
		sy := yv.Int()
		if sy < 0 {
			ypositive = false
			y = uint64(-sy)
		} else {
			y = uint64(sy)
		}

	case xr.Uint:
		y = yv.Uint()
	default:
		return nil
	}
	if !isPowerOfTwo(y) {
		return nil
	}

	shift := integerLen(y) - 1
	upn := va.Upn
	index := va.Desc.Index()
	var ret Stmt

	switch t.Kind() {
	case xr.Int:
		switch upn {
		case 0:

			y_1 :=

				int(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int)(unsafe.Pointer(&env.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int)(unsafe.Pointer(&env.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			y_1 :=

				int(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int)(unsafe.Pointer(&env.
						Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int)(unsafe.Pointer(&env.
						Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			y_1 :=

				int(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:
			y_1 :=

				int(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int)(unsafe.Pointer(&env.FileEnv.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int)(unsafe.Pointer(&env.FileEnv.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			y_1 :=

				int(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					addr := (*int)(unsafe.Pointer(&o.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					addr := (*int)(unsafe.Pointer(&o.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		}

	case xr.Int8:
		switch upn {
		case 0:

			y_1 :=

				int8(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int8)(unsafe.Pointer(&env.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int8)(unsafe.Pointer(&env.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			y_1 :=

				int8(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int8)(unsafe.Pointer(&env.
						Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int8)(unsafe.Pointer(&env.
						Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			y_1 :=

				int8(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:
			y_1 :=

				int8(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int8)(unsafe.Pointer(&env.FileEnv.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int8)(unsafe.Pointer(&env.FileEnv.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			y_1 :=

				int8(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					addr := (*int8)(unsafe.Pointer(&o.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					addr := (*int8)(unsafe.Pointer(&o.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		}

	case xr.Int16:
		switch upn {
		case 0:

			y_1 :=

				int16(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int16)(unsafe.Pointer(&env.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int16)(unsafe.Pointer(&env.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			y_1 :=

				int16(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int16)(unsafe.Pointer(&env.
						Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int16)(unsafe.Pointer(&env.
						Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			y_1 :=

				int16(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:
			y_1 :=

				int16(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int16)(unsafe.Pointer(&env.FileEnv.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int16)(unsafe.Pointer(&env.FileEnv.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			y_1 :=

				int16(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					addr := (*int16)(unsafe.Pointer(&o.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					addr := (*int16)(unsafe.Pointer(&o.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		}

	case xr.Int32:
		switch upn {
		case 0:

			y_1 :=

				int32(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int32)(unsafe.Pointer(&env.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int32)(unsafe.Pointer(&env.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			y_1 :=

				int32(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int32)(unsafe.Pointer(&env.
						Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int32)(unsafe.Pointer(&env.
						Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			y_1 :=

				int32(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:
			y_1 :=

				int32(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int32)(unsafe.Pointer(&env.FileEnv.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int32)(unsafe.Pointer(&env.FileEnv.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			y_1 :=

				int32(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					addr := (*int32)(unsafe.Pointer(&o.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					addr := (*int32)(unsafe.Pointer(&o.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		}

	case xr.Int64:
		switch upn {
		case 0:

			y_1 :=

				int64(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int64)(unsafe.Pointer(&env.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int64)(unsafe.Pointer(&env.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			y_1 :=

				int64(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int64)(unsafe.Pointer(&env.
						Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int64)(unsafe.Pointer(&env.
						Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			y_1 :=

				int64(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:
			y_1 :=

				int64(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int64)(unsafe.Pointer(&env.FileEnv.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {

					addr := (*int64)(unsafe.Pointer(&env.FileEnv.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			y_1 :=

				int64(y - 1)
			if ypositive {
				ret = func(env *Env) (Stmt, *Env) {

					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					addr := (*int64)(unsafe.Pointer(&o.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = n >> shift
					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					addr := (*int64)(unsafe.Pointer(&o.Ints[index]))

					n := *addr
					if n < 0 {
						n += y_1
					}

					*addr = -(n >> shift)
					env.IP++
					return env.Code[env.IP], env
				}
			}
		}

	case xr.Uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint)(unsafe.Pointer(&o.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case xr.Uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint8)(unsafe.Pointer(&o.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case xr.Uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint16)(unsafe.Pointer(&o.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case xr.Uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uint32)(unsafe.Pointer(&o.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case xr.Uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.Ints[index] >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Ints[index] >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.Ints[index] >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {

				env.FileEnv.Ints[index] >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.Ints[index] >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case xr.Uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {

				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}
				*(*uintptr)(unsafe.Pointer(&o.Ints[index])) >>= shift
				env.IP++
				return env.Code[env.IP], env
			}
		}

	}
	return ret
}
func (c *Comp) varQuoConst(va *Var, val I) Stmt {
	if isLiteralNumber(val, 0) {
		c.Errorf("division by %v <%T>", val, val)
		return nil
	} else if isLiteralNumber(val, 1) {
		return nil
	} else if isLiteralNumber(val, -1) {
		return c.varMulConst(va, val)
	}

	if stmt := c.varQuoPow2(va, val); stmt != nil {
		return stmt
	}

	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		val := int(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		val := int8(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		val := int16(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		val := int32(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		val := xr.ValueOf(val).Int()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() / int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		val := uint(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		val := uint8(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		val := uint16(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		val := uint32(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		val := xr.ValueOf(val).Uint()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		val := uintptr(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Float32:
		val :=

			float32(xr.ValueOf(val).Float())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Float64:
		val := xr.ValueOf(val).Float()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Complex64:
		val :=

			complex64(xr.ValueOf(val).Complex())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Complex128:
		val := xr.ValueOf(val).Complex()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex128)(unsafe.Pointer(&o.Ints[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.QUO, t)

	}
	return ret
}
func (c *Comp) varQuoExpr(va *Var, fun I) Stmt {

	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() / int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() / uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float32)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*float64)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetFloat(lhs.Float() / float64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex64)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*complex128)(unsafe.Pointer(&env.FileEnv.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*complex128)(unsafe.Pointer(&o.Ints[index])) /= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetComplex(lhs.Complex() / complex128(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.QUO, t)

	}
	return ret
}
func (c *Comp) varRemConst(va *Var, val I) Stmt {
	t := va.Type
	if reflect.IsCategory(t.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, 0) {
			c.Errorf("division by %v <%v>", val, t)
			return nil
		} else if isLiteralNumber(val, 1) {
			return c.varSetZero(va)
		}
	}

	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		val := int(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		val := int8(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		val := int16(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		val := int32(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		val := xr.ValueOf(val).Int()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() % int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		val := uint(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		val := uint8(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		val := uint16(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		val := uint32(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		val := xr.ValueOf(val).Uint()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		val := uintptr(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.REM, t)

	}
	return ret
}
func (c *Comp) varRemExpr(va *Var, fun I) Stmt {

	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() % int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) %= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() % uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.REM, t)

	}
	return ret
}
func (c *Comp) varAndConst(va *Var, val I) Stmt {
	t := va.Type
	if reflect.IsCategory(t.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, -1) {
			return nil
		} else if isLiteralNumber(val, 0) {
			return c.varSetZero(va)
		}
	}

	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		val := int(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		val := int8(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		val := int16(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		val := int32(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		val := xr.ValueOf(val).Int()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() & int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		val := uint(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		val := uint8(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		val := uint16(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		val := uint32(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		val := xr.ValueOf(val).Uint()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		val := uintptr(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.AND, t)

	}
	return ret
}
func (c *Comp) varAndExpr(va *Var, fun I) Stmt {

	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() & int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) &= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() & uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.AND, t)

	}
	return ret
}
func (c *Comp) varOrConst(va *Var, val I) Stmt {
	t := va.Type
	if reflect.IsCategory(t.Kind(), r.Int, r.Uint) && isLiteralNumber(val, 0) {
		return nil
	}

	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		val := int(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		val := int8(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		val := int16(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		val := int32(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		val := xr.ValueOf(val).Int()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() | int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		val := uint(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		val := uint8(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		val := uint16(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		val := uint32(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		val := xr.ValueOf(val).Uint()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		val := uintptr(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.OR, t)

	}
	return ret
}
func (c *Comp) varOrExpr(va *Var, fun I) Stmt {

	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() | int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) |= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() | uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.OR, t)

	}
	return ret
}
func (c *Comp) varXorConst(va *Var, val I) Stmt {
	t := va.Type
	if reflect.IsCategory(t.Kind(), r.Int, r.Uint) && isLiteralNumber(val, 0) {
		return nil
	}

	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		val := int(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		val := int8(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		val := int16(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		val := int32(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		val := xr.ValueOf(val).Int()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		val := uint(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		val := uint8(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		val := uint16(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		val := uint32(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		val := xr.ValueOf(val).Uint()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		val := uintptr(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

	}
	return ret
}
func (c *Comp) varXorExpr(va *Var, fun I) Stmt {

	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() ^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) ^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() ^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

	}
	return ret
}
func (c *Comp) varAndnotConst(va *Var, val I) Stmt {
	t := va.Type
	if reflect.IsCategory(t.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, -1) {
			return c.varSetZero(va)
		} else if isLiteralNumber(val, 0) {
			return nil
		}
	}

	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		val := int(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		val := int8(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		val := int16(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		val := int32(xr.ValueOf(val).Int())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		val := xr.ValueOf(val).Int()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		val := uint(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		val := uint8(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		val := uint16(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		val := uint32(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		val := xr.ValueOf(val).Uint()
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		val := uintptr(xr.ValueOf(val).Uint())
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(val))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.AND_NOT, t)

	}
	return ret
}
func (c *Comp) varAndnotExpr(va *Var, fun I) Stmt {

	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int)(unsafe.Pointer(&o.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int8)(unsafe.Pointer(&o.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int16)(unsafe.Pointer(&o.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int32)(unsafe.Pointer(&o.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*int64)(unsafe.Pointer(&o.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					{
						lhs :=

							o.Vals[index]
						lhs.SetInt(lhs.Int() &^ int64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint)(unsafe.Pointer(&o.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Ints[index] &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Ints[index] &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Ints[index] &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) &^= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Vals[index]
						lhs.SetUint(lhs.Uint() &^ uint64(fun(env)))
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= on <%v>`, token.AND_NOT, t)

	}
	return ret
}
func (c *Comp) setVar(va *Var, op token.Token, init *Expr) Stmt {
	t := va.Type
	var shift bool
	var err interface{} = ""
	switch op {
	case token.SHL, token.SHL_ASSIGN, token.SHR, token.SHR_ASSIGN:
		shift = true
		if init.Untyped() {
			init.ConstTo(c.TypeOfUint64())
			err = nil
		} else if init.Type == nil {
			err = fmt.Sprintf("\n\treason: type is %v, expecting integer", init.Type)
		} else if c := reflect.Category(init.Type.Kind()); c != r.Int && c != r.Uint {
			err = fmt.Sprintf("\n\treason: type %v is %v, expecting integer", init.Type, init.Type.Kind())
		} else {
			err = nil
		}

	default:
		if init.Const() {
			init.ConstTo(t)
			err = nil
		} else if init.Type == nil {
			if op != token.ASSIGN {
				err = fmt.Sprintf("\n\treason: invalid operation %s nil", op)
			} else if !reflect.IsNillableKind(t.Kind()) {
				err = fmt.Sprintf("\n\treason: cannot assign nil to %v", t)
			}
		} else if !init.Type.AssignableTo(t) {
			err = interfaceMissingMethod(init.Type, t)
		} else {
			err = nil
		}

	}
	if err != nil {
		c.Errorf("incompatible types in assignment: %v %s %v%v", t, op, init.Type, err)
		return nil
	}
	class := va.Desc.Class()
	if class != VarBind && class != IntBind {
		c.Errorf("invalid operator %s on %v", op, class)
		return nil
	}
	index := va.Desc.Index()
	if index == NoIndex {
		if op != token.ASSIGN {
			c.Errorf("invalid operator %s on _", op)
		}

		if init.Const() {
			return nil
		}
		return init.AsStmt(c)
	}
	if init.Const() {
		rt := t.ReflectType()
		val := init.Value
		v := xr.ValueOf(val)
		if !v.IsValid() || v == None {
			v = xr.ZeroR(rt)
			val = v.Interface()
		} else if v.Type() != rt && !shift {
			v = convert(v, rt)
			val = v.Interface()
		}
		switch op {
		case token.ASSIGN:
			return c.varSetConst(va, val)
		case token.ADD, token.ADD_ASSIGN:
			return c.varAddConst(va, val)
		case token.SUB, token.SUB_ASSIGN:
			return c.varSubConst(va, val)
		case token.MUL, token.MUL_ASSIGN:
			return c.varMulConst(va, val)
		case token.QUO, token.QUO_ASSIGN:
			return c.varQuoConst(va, val)
		case token.REM, token.REM_ASSIGN:
			return c.varRemConst(va, val)
		case token.AND, token.AND_ASSIGN:
			return c.varAndConst(va, val)
		case token.OR, token.OR_ASSIGN:
			return c.varOrConst(va, val)
		case token.XOR, token.XOR_ASSIGN:
			return c.varXorConst(va, val)
		case token.SHL, token.SHL_ASSIGN:
			return c.varShlConst(va, val)
		case token.SHR, token.SHR_ASSIGN:
			return c.varShrConst(va, val)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			return c.varAndnotConst(va, val)
		}
	} else {
		fun := init.Fun
		switch op {
		case token.ASSIGN:
			return c.varSetExpr(va, init)
		case token.ADD, token.ADD_ASSIGN:
			return c.varAddExpr(va, fun)
		case token.SUB, token.SUB_ASSIGN:
			return c.varSubExpr(va, fun)
		case token.MUL, token.MUL_ASSIGN:
			return c.varMulExpr(va, fun)
		case token.QUO, token.QUO_ASSIGN:
			return c.varQuoExpr(va, fun)
		case token.REM, token.REM_ASSIGN:
			return c.varRemExpr(va, fun)
		case token.AND, token.AND_ASSIGN:
			return c.varAndExpr(va, fun)
		case token.OR, token.OR_ASSIGN:
			return c.varOrExpr(va, fun)
		case token.XOR, token.XOR_ASSIGN:
			return c.varXorExpr(va, fun)
		case token.SHL, token.SHL_ASSIGN:
			return c.varShlExpr(va, init)
		case token.SHR, token.SHR_ASSIGN:
			return c.varShrExpr(va, init)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			return c.varAndnotExpr(va, fun)
		}
	}
	c.Errorf("invalid operator %s", op)
	return nil
}
