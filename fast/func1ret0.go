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
 * func1ret0.go
 *
 *  Created on Apr 16, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"unsafe"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) func1ret0(t xr.Type, m *funcMaker) func(*Env) xr.Value {

	nbind := m.nbind
	nintbind := m.nintbind
	funcbody := m.funcbody
	param0index := m.Param[0].Desc.Index()

	var debugC *Comp
	if c.Globals.Options&base.OptDebugger != 0 {
		debugC = c
	}

	targ0 := t.In(0)
	karg0 := targ0.Kind()
	switch karg0 {
	case xr.Bool:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					bool) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 bool) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*bool)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case xr.Int:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					int) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 int) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case xr.Int8:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					int8) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 int8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int8)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case xr.Int16:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					int16) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 int16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int16)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case xr.Int32:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					int32) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 int32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int32)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case xr.Int64:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					int64) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 int64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*int64)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case xr.Uint:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					uint) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 uint) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case xr.Uint8:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					uint8) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 uint8) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint8)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case xr.Uint16:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					uint16) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 uint16) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint16)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case xr.Uint32:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					uint32) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 uint32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uint32)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}
		}
	case xr.Uint64:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					uint64) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 uint64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					env.Ints[param0index] = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}

		}
	case xr.Uintptr:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					uintptr) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 uintptr) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*uintptr)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}

		}
	case xr.Float32:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					float32) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 float32) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float32)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}

		}
	case xr.Float64:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					float64) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 float64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*float64)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}

		}
	case xr.Complex64:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					complex64) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 complex64) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					*(*complex64)(unsafe.Pointer(&env.Ints[param0index])) = arg0

					funcbody(env)

					env.freeEnv4Func()
				})
			}

		}
	case xr.Complex128:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					complex128) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 complex128) {
					env := newEnv4Func(env, nbind, nintbind, debugC)
					{
						place := xr.NewR(base.TypeOfComplex128).Elem()
						place.SetComplex(arg0)
						env.Vals[param0index] = place
					}

					funcbody(env)

					env.freeEnv4Func()
				})
			}

		}
	case xr.String:
		{
			if funcbody == nil {
				funv := xr.ValueOf(func(

					string) {
				})
				return func(env *Env) xr.Value { return funv }

			}
			return func(env *Env) xr.Value {

				env.MarkUsedByClosure()
				return xr.ValueOf(func(arg0 string) {
					env := newEnv4Func(env, nbind, nintbind, debugC)

					{
						place := xr.NewR(base.TypeOfString).Elem()
						place.SetString(arg0)
						env.Vals[param0index] = place
					}
					funcbody(env)

					env.freeEnv4Func()
				})
			}

		}
	default:
		{
			if funcbody == nil {
				funv := xr.MakeFunc(t, func([]xr.Value) []xr.Value { return nil })
				return func(env *Env) xr.Value { return funv }

			} else {
				return func(env *Env) xr.Value {

					env.MarkUsedByClosure()
					return xr.MakeFunc(t, func(args []xr.Value) []xr.Value {
						env := newEnv4Func(env, nbind, nintbind, debugC)

						if param0index != NoIndex {
							place := xr.New(targ0).Elem()
							if arg0 := args[0]; arg0.IsValid() && arg0 != None {
								place.Set(convert(arg0, targ0.ReflectType()))
							}

							env.Vals[param0index] = place
						}

						funcbody(env)
						return ZeroValues
					})
				}
			}
		}
	}
}
