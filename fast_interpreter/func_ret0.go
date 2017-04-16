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
 * func_ret0.go
 *
 *  Created on Apr 16, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	r "reflect"
	"unsafe"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) func_ret0(t r.Type, m *funcMaker) func(*Env) r.Value {

	nbinds := m.nbinds
	nintbinds := m.nintbinds
	parambinds := m.parambinds
	funcbody := m.funcbody

	switch t.NumIn() {
	case 0:
		if funcbody == nil {
			return func(env *Env) r.Value { return r.ValueOf(INone) }
		} else {
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func() {
					env := NewEnv4Func(env, nbinds, nintbinds)

					funcbody(env)

					env.FreeEnv()
				})
			}
		}
	case 1:
		targ0 := t.In(0)
		karg0 := targ0.Kind()
		switch karg0 {
		case r.Bool:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(bool) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(bool) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 bool) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*bool)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Int:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(int) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*int)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Int8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int8) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(int8) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int8) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*int8)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Int16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int16) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(int16) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int16) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*int16)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Int32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int32) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(int32) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int32) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*int32)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Int64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(int64) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(int64) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 int64) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*int64)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Uint:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(uint) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*uint)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Uint8:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint8) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(uint8) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint8) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*uint8)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Uint16:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint16) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(uint16) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint16) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*uint16)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Uint32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint32) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(uint32) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint32) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*uint32)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Uint64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uint64) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(uint64) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uint64) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							env.IntBinds[param0index] = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Uintptr:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(uintptr) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(uintptr) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 uintptr) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*uintptr)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Float32:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float32) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(float32) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float32) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*float32)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Float64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(float64) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(float64) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 float64) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*float64)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Complex64:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex64) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(complex64) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex64) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							*(*complex64)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.Complex128:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(complex128) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(complex128) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 complex128) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfComplex128).Elem()
								place.SetComplex(arg0,
								)
								env.Binds[param0index] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		case r.String:
			{
				if funcbody == nil {
					return func(env *Env) r.Value {
						return r.ValueOf(func(string) {
						})
					}
				}

				param0index := parambinds[0].Desc.Index()
				if param0index == NoIndex {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(string) {
							env := NewEnv4Func(env, nbinds, nintbinds)

							funcbody(env)

							env.FreeEnv()
						})
					}
				} else {
					return func(env *Env) r.Value {

						env.MarkUsedByClosure()
						return r.ValueOf(func(arg0 string) {
							env := NewEnv4Func(env, nbinds, nintbinds)
							{
								place := r.New(TypeOfString).Elem()
								place.SetString(arg0,
								)
								env.Binds[param0index] = place
							}

							funcbody(env)

							env.FreeEnv()
						})
					}
				}

			}
		default:

			if funcbody == nil {
				return func(env *Env) r.Value {
					return r.MakeFunc(t, func([]r.Value) []r.Value { return ZeroValues },
					)
				}
			} else {
				param0index := parambinds[0].Desc.Index()
				return func(env *Env) r.Value {

					env.MarkUsedByClosure()
					return r.MakeFunc(t, func(args []r.Value) []r.Value {
						env := NewEnv4Func(env, nbinds, nintbinds)

						if param0index != NoIndex {
							place := r.New(targ0).Elem()
							if arg0 := args[0]; arg0 != Nil && arg0 != None {
								place.Set(arg0.Convert(targ0))
							}

							env.Binds[param0index] = place
						}

						funcbody(env)
						return ZeroValues
					})
				}
			}
		}
	case 2:
		targ0 := t.In(0)
		targ1 := t.In(1)
		karg0 := targ0.Kind()
		karg1 := targ1.Kind()

		targs := [2]r.Type{targ0, targ1}
		paramindexes := [2]int{
			parambinds[0].Desc.Index(),
			parambinds[1].Desc.Index(),
		}

		if IsOptimizedKind(karg0) && IsOptimizedKind(karg1) {
			argdecls := [2]func(*Env, r.Value){
				c.DeclBindRuntimeValue(m.paramnames[0], parambinds[0]),
				c.DeclBindRuntimeValue(m.paramnames[1], parambinds[1]),
			}
			switch karg0 {
			case r.Bool:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool, bool) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 bool,

									arg1 bool) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 bool,

									arg1 bool) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*bool)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 bool,

									arg1 bool) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*bool)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 bool,

									arg1 bool) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*bool)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*bool)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,
									int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,

									uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,

									uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,

									uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,

									uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,

									uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,

									uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,

									float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,

									float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,

									complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,

									complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(bool,

									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 bool,

								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Int:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int, int) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int,

									arg1 int) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int,

									arg1 int) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int,

									arg1 int) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int,

									arg1 int) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*int)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int,
									uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int,

									uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int,

									uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int,

									uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int,

									uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int,

									uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int,

									float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int,

									float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int,

									complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int,

									complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int,

									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int,

								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Int8:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8, int8) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int8,

									arg1 int8) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int8,

									arg1 int8) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int8)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int8,

									arg1 int8) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int8)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int8,

									arg1 int8) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int8)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*int8)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8,
									uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8,

									uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8,

									uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8,

									uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8,

									uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8,

									float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8,

									float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8,

									complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8,

									complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int8,

									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int8,

								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Int16:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16, int16) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int16,

									arg1 int16) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int16,

									arg1 int16) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int16)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int16,

									arg1 int16) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int16)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int16,

									arg1 int16) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int16)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*int16)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16,
									uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16,

									uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16,

									uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16,

									uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16,

									float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16,

									float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16,

									complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16,

									complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int16,

									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int16,

								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Int32:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32, int32) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int32,

									arg1 int32) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int32,

									arg1 int32) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int32)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int32,

									arg1 int32) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int32)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int32,

									arg1 int32) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int32)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*int32)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32,
									uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32,

									uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32,

									uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32,

									float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32,

									float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32,

									complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32,

									complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int32,

									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int32,

								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Int64:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64, int64) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int64,

									arg1 int64) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int64,

									arg1 int64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int64)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int64,

									arg1 int64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int64)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 int64,

									arg1 int64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*int64)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*int64)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64, uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64,
									uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64,

									uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64,

									float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64,

									float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64,

									complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64,

									complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(int64,

									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 int64,

								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Uint:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint, uint) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint,

									arg1 uint) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint,

									arg1 uint) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint,

									arg1 uint) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint,

									arg1 uint) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*uint)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint, uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint, uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint,
									uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint,

									float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint,

									float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint,

									complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint,

									complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint,

									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint,

								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Uint8:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, uint8) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint8,

									arg1 uint8) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint8,

									arg1 uint8) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint8)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint8,

									arg1 uint8) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint8)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint8,

									arg1 uint8) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint8)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*uint8)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8, uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8,
									float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8,

									float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8,

									complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8,

									complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint8,

									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint8,

								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Uint16:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, uint16) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint16,

									arg1 uint16) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint16,

									arg1 uint16) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint16)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint16,

									arg1 uint16) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint16)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint16,

									arg1 uint16) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint16)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*uint16)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16, float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16,
									float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16,

									complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16,

									complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint16,

									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint16,

								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Uint32:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, uint32) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint32,

									arg1 uint32) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint32,

									arg1 uint32) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint32)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint32,

									arg1 uint32) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint32)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint32,

									arg1 uint32) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uint32)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*uint32)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32, float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32,
									complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32,

									complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint32,

									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint32,

								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Uint64:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, uint64) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint64,
									arg1 uint64) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint64,

									arg1 uint64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									env.IntBinds[paramindexes[0]] = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint64,

									arg1 uint64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									env.IntBinds[paramindexes[1]] = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uint64,

									arg1 uint64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									env.IntBinds[paramindexes[0]] = arg0

									env.IntBinds[paramindexes[1]] = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64, complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64,
									complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uint64,

									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uint64,

								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Uintptr:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, uintptr) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uintptr, arg1 uintptr) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uintptr,

									arg1 uintptr) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uintptr)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uintptr,

									arg1 uintptr) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uintptr)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 uintptr,

									arg1 uintptr) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*uintptr)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*uintptr)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr, complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(uintptr,
									string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 uintptr,
								arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Float32:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, float32) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 float32, arg1 float32) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 float32,

									arg1 float32) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*float32)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 float32,

									arg1 float32) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*float32)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 float32,

									arg1 float32) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*float32)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*float32)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float32, string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float32, arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Float64:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, float64) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 float64, arg1 float64) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 float64,

									arg1 float64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*float64)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 float64,

									arg1 float64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*float64)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 float64,

									arg1 float64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*float64)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*float64)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(float64, string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 float64, arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}
			case r.Complex64:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, complex64) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 complex64, arg1 complex64) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 complex64,

									arg1 complex64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*complex64)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 complex64,

									arg1 complex64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*complex64)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 complex64,

									arg1 complex64) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									*(*complex64)(unsafe.Pointer(&env.IntBinds[paramindexes[0]])) = arg0

									*(*complex64)(unsafe.Pointer(&env.IntBinds[paramindexes[1]])) = arg0

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex64, string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex64, arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}

			case r.Complex128:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, complex128) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 complex128, arg1 complex128) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 complex128,

									arg1 complex128) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									{
										place := r.New(TypeOfComplex128).Elem()
										place.SetComplex(arg0,
										)
										env.Binds[paramindexes[0]] = place
									}

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 complex128,

									arg1 complex128) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									{
										place := r.New(TypeOfComplex128).Elem()
										place.SetComplex(arg0,
										)
										env.Binds[paramindexes[1]] = place
									}

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 complex128,

									arg1 complex128) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									{
										place := r.New(TypeOfComplex128).Elem()
										place.SetComplex(arg0,
										)
										env.Binds[paramindexes[0]] = place
									}
									{
										place := r.New(TypeOfComplex128).Elem()
										place.SetComplex(arg0,
										)
										env.Binds[paramindexes[1]] = place
									}

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(complex128, string) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 complex128, arg1 string) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				}

			case r.String:
				switch karg0 {
				case r.Bool:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, bool) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 bool) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, int) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 int) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int8:

					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, int8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 int8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, int16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 int16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, int32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 int32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Int64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, int64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 int64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, uint) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 uint) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint8:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, uint8) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 uint8) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint16:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, uint16) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 uint16) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, uint32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 uint32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uint64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, uint64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 uint64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Uintptr:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, uintptr) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 uintptr) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float32:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, float32) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 float32) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Float64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, float64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 float64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex64:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, complex64) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 complex64) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.Complex128:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, complex128) {
								})
							}
						}
						return func(env *Env) r.Value {

							if paramindexes[0] != NoIndex || paramindexes[1] != NoIndex {
								env.MarkUsedByClosure()
							}
							return r.ValueOf(func(arg0 string, arg1 complex128) {
								env := NewEnv4Func(env, nbinds, nintbinds)

								argdecls[0](env, r.ValueOf(arg0))
								argdecls[1](env, r.ValueOf(arg1))

								funcbody(env)

								env.FreeEnv()
							})
						}
					}
				case r.String:
					{
						if funcbody == nil {
							return func(env *Env) r.Value {
								return r.ValueOf(func(string, string) {
								})
							}
						}

						if paramindexes[0] == NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 string, arg1 string) {
									env := NewEnv4Func(env, nbinds, nintbinds)

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] != NoIndex && paramindexes[1] == NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 string,

									arg1 string) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									{
										place := r.New(TypeOfString).Elem()
										place.SetString(arg0,
										)
										env.Binds[paramindexes[0]] = place
									}

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else if paramindexes[0] == NoIndex && paramindexes[1] != NoIndex {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 string,

									arg1 string) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									{
										place := r.New(TypeOfString).Elem()
										place.SetString(arg0,
										)
										env.Binds[paramindexes[1]] = place
									}

									funcbody(env)

									env.FreeEnv()
								})
							}
						} else {
							return func(env *Env) r.Value {

								env.MarkUsedByClosure()
								return r.ValueOf(func(arg0 string,

									arg1 string) {
									env := NewEnv4Func(env, nbinds, nintbinds)
									{
										place := r.New(TypeOfString).Elem()
										place.SetString(arg0,
										)
										env.Binds[paramindexes[0]] = place
									}
									{
										place := r.New(TypeOfString).Elem()
										place.SetString(arg0,
										)
										env.Binds[paramindexes[1]] = place
									}

									funcbody(env)

									env.FreeEnv()
								})
							}
						}

					}
				}

			}
		}

		if funcbody == nil {
			return func(env *Env) r.Value {
				return r.MakeFunc(t, func([]r.Value) []r.Value { return ZeroValues },
				)
			}
		} else {
			return func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.MakeFunc(t, func(args []r.Value) []r.Value {
					env := NewEnv4Func(env, nbinds, nintbinds)

					for i := range targs {
						if idx := paramindexes[i]; idx != NoIndex {
							place := r.New(targs[i]).Elem()
							if arg := args[i]; arg != Nil && arg != None {
								place.Set(arg.Convert(targs[i]))
							}

							env.Binds[idx] = place
						}
					}

					funcbody(env)
					return ZeroValues
				})
			}
		}
	}
	return c.funcGeneric(t, m)
}
