// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *
 * func2ret0.go
 *
 *  Created on Apr 16, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	"unsafe"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) func2ret0(t xr.Type, m *funcMaker) func(*Env) r.Value {
	karg0 := t.In(0).Kind()
	karg1 := t.In(1).Kind()

	if !IsOptimizedKind(karg0) || !IsOptimizedKind(karg1) {
		return nil
	}

	indexes := &[2]int{
		m.Param[0].Desc.Index(),
		m.Param[1].Desc.Index(),
	}
	var debugC *Comp
	if c.Globals.Options&OptDebugger != 0 {
		debugC = c
	}

	var ret func(*Env) r.Value
	switch karg0 {
	case r.Bool:
		ret = func2ret0Bool(m, indexes, karg1, debugC)

	case r.Int:
		ret = func2ret0Int(m, indexes, karg1, debugC)

	case r.Int8:
		ret = func2ret0Int8(m, indexes, karg1, debugC)

	case r.Int16:
		ret = func2ret0Int16(m, indexes, karg1, debugC)

	case r.Int32:
		ret = func2ret0Int32(m, indexes, karg1, debugC)

	case r.Int64:
		ret = func2ret0Int64(m, indexes, karg1, debugC)

	case r.Uint:
		ret = func2ret0Uint(m, indexes, karg1, debugC)

	case r.Uint8:
		ret = func2ret0Uint8(m, indexes, karg1, debugC)

	case r.Uint16:
		ret = func2ret0Uint16(m, indexes, karg1, debugC)

	case r.Uint32:
		ret = func2ret0Uint32(m, indexes, karg1, debugC)

	case r.Uint64:
		ret = func2ret0Uint64(m, indexes, karg1, debugC)

	case r.Uintptr:
		ret = func2ret0Uintptr(m, indexes, karg1, debugC)

	case r.Float32:
		ret = func2ret0Float32(m, indexes, karg1, debugC)

	case r.Float64:
		ret = func2ret0Float64(m, indexes, karg1, debugC)

	case r.Complex64:
		ret = func2ret0Complex64(m, indexes, karg1, debugC)

	case r.Complex128:
		ret = func2ret0Complex128(m, indexes, karg1, debugC)

	case r.String:
		ret = func2ret0String(m, indexes, karg1, debugC)

	}
	return ret
}
func func2ret0Bool(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(bool,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 bool,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Int(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Int8(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int8,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int8,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Int16(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int16,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int16,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Int32(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int32,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int32,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Int64(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(int64,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 int64,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uint(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uint8(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint8,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint8,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uint16(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint16,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint16,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uint32(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint32,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint32,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uint64(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uint64,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uint64,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[indexes[0]] = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Uintptr(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(uintptr,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 uintptr,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Float32(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float32,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float32,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Float64(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(float64,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 float64,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Complex64(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex64,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex64,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[0]])) = arg0

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0Complex128(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(complex128,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 complex128,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
func func2ret0String(m *funcMaker, indexes *[2]int, karg1 r.Kind, debugC *Comp) func(*Env) r.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	var ret func(*Env) r.Value
	switch karg1 {
	case r.Bool:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						bool) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*bool)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						int) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int8:

		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						int8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						int16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						int32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Int64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						int64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*int64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						uint) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint8:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						uint8) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint8)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint16:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						uint16) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint16)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						uint32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uint32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uint64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						uint64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					env.Ints[indexes[1]] = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Uintptr:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						uintptr) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*uintptr)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float32:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						float32) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*float32)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Float64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						float64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*float64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex64:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						complex64) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}

					*(*complex64)(unsafe.Pointer(&env.Ints[indexes[1]])) = arg1

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.Complex128:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						complex128) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}
					{
						place := r.New(TypeOfComplex128).Elem()
						place.SetComplex(arg1,
						)
						env.Vals[indexes[1]] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case r.String:
		{
			if funcbody == nil {
				ret = func(env *Env) r.Value {
					return r.ValueOf(func(string,

						string) {
					})
				}

				break
			}
			ret = func(env *Env) r.Value {

				env.MarkUsedByClosure()
				return r.ValueOf(func(arg0 string,

					arg1 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg0,
						)
						env.Vals[indexes[0]] = place
					}
					{
						place := r.New(TypeOfString).Elem()
						place.SetString(arg1,
						)
						env.Vals[indexes[1]] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	}
	return ret
}
