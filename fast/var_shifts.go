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
 * var_shifts.go
 *
 *  Created on May 17, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
	r "reflect"
	"unsafe"

	"github.com/cosmos72/gomacro/base/reflect"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) varShlConst(va *Var, ival I) Stmt {
	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind

	t2 := r.TypeOf(ival)
	if t2 == nil {
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHL, t, t2)
	} else if cat := reflect.Category(t2.Kind()); cat != r.Int && cat != r.Uint {
		c.Errorf(`invalid operator %s= between <%v> and <%v> // %v`, token.SHL, t, t2, t2.Kind())
	}

	if isLiteralNumber(ival, 0) {
		return nil
	}
	val, ok := constAsUint64(ival)
	if !ok {
		c.Errorf(`invalid shift amount: %v %v %v`, va, token.SHL, ival)
	}

	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
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
					*(*int)(unsafe.Pointer(&o.Ints[index])) <<= val

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
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
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
					*(*int8)(unsafe.Pointer(&o.Ints[index])) <<= val

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
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
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
					*(*int16)(unsafe.Pointer(&o.Ints[index])) <<= val

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
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
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
					*(*int32)(unsafe.Pointer(&o.Ints[index])) <<= val

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
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							val,
						)
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
					*(*int64)(unsafe.Pointer(&o.Ints[index])) <<= val

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
						lhs.SetInt(lhs.Int() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
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
					*(*uint)(unsafe.Pointer(&o.Ints[index])) <<= val

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
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
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
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) <<= val

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
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
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
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) <<= val

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
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
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
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) <<= val

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
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
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
						Ints[index] <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
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
						Ints[index] <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
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

					o.Ints[index] <<= val

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
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							val,
						)
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
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) <<= val

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
						lhs.SetUint(lhs.Uint() <<
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHL, t, t2)

	}
	return ret
}
func (c *Comp) varShlExpr(va *Var, e *Expr) Stmt {
	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind

	t2 := funTypeOut(e.Fun)
	if t2 == nil {
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHL, t, t2)
	} else if cat := reflect.Category(t2.Kind()); cat != r.Int && cat != r.Uint {
		c.Errorf(`invalid operator %s= between <%v> and <%v> // %v`, token.SHL, t, t2, t2.Kind())
	}

	fun := e.AsUint64()
	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
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
					*(*int)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

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
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
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
					*(*int8)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

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
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
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
					*(*int16)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

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
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
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
					*(*int32)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

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
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
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
					*(*int64)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

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
						lhs.SetInt(lhs.Int() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
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
					*(*uint)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

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
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
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
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

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
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
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
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

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
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
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
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

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
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
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
						Ints[index] <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
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
						Ints[index] <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
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

					o.Ints[index] <<= fun(env)

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
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) <<= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
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
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) <<= fun(env)

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
						lhs.SetUint(lhs.Uint() <<
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHL, t, t2)

	}
	return ret
}
func (c *Comp) varShrConst(va *Var, ival I) Stmt {
	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind

	t2 := r.TypeOf(ival)
	if t2 == nil {
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHL, t, t2)
	} else if cat := reflect.Category(t2.Kind()); cat != r.Int && cat != r.Uint {
		c.Errorf(`invalid operator %s= between <%v> and <%v> // %v`, token.SHL, t, t2, t2.Kind())
	}

	if isLiteralNumber(ival, 0) {
		return nil
	}
	val, ok := constAsUint64(ival)
	if !ok {
		c.Errorf(`invalid shift amount: %v %v %v`, va, token.SHL, ival)
	}

	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
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
					*(*int)(unsafe.Pointer(&o.Ints[index])) >>= val

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
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
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
					*(*int8)(unsafe.Pointer(&o.Ints[index])) >>= val

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
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
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
					*(*int16)(unsafe.Pointer(&o.Ints[index])) >>= val

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
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
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
					*(*int32)(unsafe.Pointer(&o.Ints[index])) >>= val

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
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							val,
						)
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
					*(*int64)(unsafe.Pointer(&o.Ints[index])) >>= val

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
						lhs.SetInt(lhs.Int() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
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
					*(*uint)(unsafe.Pointer(&o.Ints[index])) >>= val

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
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
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
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) >>= val

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
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
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
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) >>= val

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
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
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
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) >>= val

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
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
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
						Ints[index] >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
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
						Ints[index] >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
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

					o.Ints[index] >>= val

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
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= val

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							val,
						)
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
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) >>= val

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
						lhs.SetUint(lhs.Uint() >>
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHR, t, t2)

	}
	return ret
}
func (c *Comp) varShrExpr(va *Var, e *Expr) Stmt {
	t := va.Type
	upn := va.Upn
	index := va.Desc.Index()
	intbinds := va.Desc.Class() == IntBind

	t2 := funTypeOut(e.Fun)
	if t2 == nil {
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHL, t, t2)
	} else if cat := reflect.Category(t2.Kind()); cat != r.Int && cat != r.Uint {
		c.Errorf(`invalid operator %s= between <%v> and <%v> // %v`, token.SHL, t, t2, t2.Kind())
	}

	fun := e.AsUint64()
	var ret Stmt
	switch t.Kind() {
	case xr.Int:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
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
					*(*int)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

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
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
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
					*(*int8)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

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
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
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
					*(*int16)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

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
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
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
					*(*int32)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

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
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Int64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
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
					*(*int64)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

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
						lhs.SetInt(lhs.Int() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
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
					*(*uint)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

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
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint8:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
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
					*(*uint8)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

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
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint16:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
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
					*(*uint16)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

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
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint32:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
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
					*(*uint32)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

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
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uint64:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.
						Ints[index] >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
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
						Ints[index] >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
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
						Ints[index] >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					env.FileEnv.
						Ints[index] >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
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

					o.Ints[index] >>= fun(env)

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
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	case xr.Uintptr:
		switch upn {
		case 0:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case 2:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case c.Depth - 1:

			if intbinds {
				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.FileEnv.Ints[index])) >>= fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			} else {
				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.FileEnv.
							Vals[index]
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
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
					*(*uintptr)(unsafe.Pointer(&o.Ints[index])) >>= fun(env)

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
						lhs.SetUint(lhs.Uint() >>
							fun(env),
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SHR, t, t2)

	}
	return ret
}
