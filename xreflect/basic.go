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
 * basic.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"

	"go/types"
)

func errorf(format string, arg ...interface{}) {
	panic(errors.New(fmt.Sprintf(format, arg...)))
}

func makebasictypes() []Type {
	rmap := []reflect.Type{
		reflect.Bool:          reflect.TypeOf(bool(false)),
		reflect.Int:           reflect.TypeOf(int(0)),
		reflect.Int8:          reflect.TypeOf(int8(0)),
		reflect.Int16:         reflect.TypeOf(int16(0)),
		reflect.Int32:         reflect.TypeOf(int32(0)),
		reflect.Int64:         reflect.TypeOf(int64(0)),
		reflect.Uint:          reflect.TypeOf(uint(0)),
		reflect.Uint8:         reflect.TypeOf(uint8(0)),
		reflect.Uint16:        reflect.TypeOf(uint16(0)),
		reflect.Uint32:        reflect.TypeOf(uint32(0)),
		reflect.Uint64:        reflect.TypeOf(uint64(0)),
		reflect.Uintptr:       reflect.TypeOf(uintptr(0)),
		reflect.Float32:       reflect.TypeOf(float32(0)),
		reflect.Float64:       reflect.TypeOf(float64(0)),
		reflect.Complex64:     reflect.TypeOf(complex64(0)),
		reflect.Complex128:    reflect.TypeOf(complex128(0)),
		reflect.String:        reflect.TypeOf(string("")),
		reflect.UnsafePointer: reflect.TypeOf(unsafe.Pointer(nil)),
	}
	m := make([]Type, len(rmap))
	for gkind := types.Bool; gkind <= types.UnsafePointer; gkind++ {
		kind := gbasickindToKind(gkind)
		gtype := types.Typ[gkind]
		rtype := rmap[kind]
		if gtype == nil || rtype == nil {
			continue
		}
		m[kind] = maketype(gtype, rtype)
	}

	return m
}

const MaxUInt = ^uint(0)
const MaxInt = int(MaxUInt >> 1)
const MaxDepth = MaxInt

var BasicTypes = makebasictypes()

var (
	TypeOfBool       = BasicTypes[reflect.Bool]
	TypeOfInt        = BasicTypes[reflect.Int]
	TypeOfInt8       = BasicTypes[reflect.Int8]
	TypeOfInt16      = BasicTypes[reflect.Int16]
	TypeOfInt32      = BasicTypes[reflect.Int32]
	TypeOfInt64      = BasicTypes[reflect.Int64]
	TypeOfUint       = BasicTypes[reflect.Uint]
	TypeOfUint8      = BasicTypes[reflect.Uint8]
	TypeOfUint16     = BasicTypes[reflect.Uint16]
	TypeOfUint32     = BasicTypes[reflect.Uint32]
	TypeOfUint64     = BasicTypes[reflect.Uint64]
	TypeOfUintptr    = BasicTypes[reflect.Uintptr]
	TypeOfFloat32    = BasicTypes[reflect.Float32]
	TypeOfFloat64    = BasicTypes[reflect.Float64]
	TypeOfComplex64  = BasicTypes[reflect.Complex64]
	TypeOfComplex128 = BasicTypes[reflect.Complex128]
	TypeOfString     = BasicTypes[reflect.String]

	TypeOfByte          = TypeOfUint8
	TypeOfRune          = TypeOfInt32
	TypeOfUnsafePointer = BasicTypes[reflect.UnsafePointer]

	TypeOfError = maketype(
		types.Universe.Lookup("error").Type(),
		reflect.TypeOf((*error)(nil)).Elem(),
	)
	TypeOfInterface = maketype(
		types.NewInterface(nil, nil).Complete(),
		reflect.TypeOf((*interface{})(nil)).Elem(),
	)
	reflectTypeOfInterfaceHeader = reflect.TypeOf(InterfaceHeader{})
)

// Bits returns the size of the type in bits.
// It panics if the type's Kind is not one of the
// sized or unsized Int, Uint, Float, or Complex kinds.
func (t *xtype) Bits() int {
	return t.rtype.Bits()
}

// Align returns the alignment in bytes of a value of
// this type when allocated in memory.
func (t *xtype) Align() int {
	return t.rtype.Align()
}

// FieldAlign returns the alignment in bytes of a value of
// this type when used as a field in a struct.
func (t *xtype) FieldAlign() int {
	return t.rtype.FieldAlign()
}
