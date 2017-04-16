/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * constants.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	r "reflect"
)

type none struct{}

const (
	ReflectGensymPrefix = "\u0080"     // prefix to generate names of extra struct fields needed by the interpreter in reflect.StructOf()
	GensymPrefix        = "\U000124AD" // prefix to generate names in macros - arbitrarily chosen U+124AD CUNEIFORM SIGN ERIN2 X - reasons:
	// * accepted by Go compiler identifier name in source code
	// * belongs to an ancient language no longer spoken, so hopefully low collision risk
	// * outside Unicode basic place, so hopefully lower collision risk
	// * relatively simple glyph picture
)

var (
	Nil = r.Value{}

	None = r.ValueOf(none{}) // used to indicate "no value"

	True  = r.ValueOf(true)
	False = r.ValueOf(false)

	One = r.ValueOf(1)

	TypeOfInt   = r.TypeOf(int(0))
	TypeOfInt8  = r.TypeOf(int8(0))
	TypeOfInt16 = r.TypeOf(int16(0))
	TypeOfInt32 = r.TypeOf(int32(0))
	TypeOfInt64 = r.TypeOf(int64(0))

	TypeOfUint    = r.TypeOf(uint(0))
	TypeOfUint8   = r.TypeOf(uint8(0))
	TypeOfUint16  = r.TypeOf(uint16(0))
	TypeOfUint32  = r.TypeOf(uint32(0))
	TypeOfUint64  = r.TypeOf(uint64(0))
	TypeOfUintptr = r.TypeOf(uintptr(0))

	TypeOfFloat32    = r.TypeOf(float32(0))
	TypeOfFloat64    = r.TypeOf(float64(0))
	TypeOfComplex64  = r.TypeOf(complex64(0))
	TypeOfComplex128 = r.TypeOf(complex128(0))

	TypeOfBool      = r.TypeOf(false)
	TypeOfByte      = r.TypeOf(byte(0))
	TypeOfRune      = r.TypeOf(rune(0))
	TypeOfString    = r.TypeOf("")
	TypeOfInterface = r.TypeOf((*interface{})(nil)).Elem()
	TypeOfError     = r.TypeOf((*error)(nil)).Elem()
	TypeOfDeferFunc = r.TypeOf(func() {})

	ZeroStrings = []string{}
	ZeroTypes   = []r.Type{}
	ZeroValues  = []r.Value{}
)
