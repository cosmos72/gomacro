// +build !gomacro_xreflect_strict

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
 * easy.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"reflect"
)

var (
	BasicTypes = universe.BasicTypes

	TypeOfBool          = BasicTypes[reflect.Bool]
	TypeOfInt           = BasicTypes[reflect.Int]
	TypeOfInt8          = BasicTypes[reflect.Int8]
	TypeOfInt16         = BasicTypes[reflect.Int16]
	TypeOfInt32         = BasicTypes[reflect.Int32]
	TypeOfInt64         = BasicTypes[reflect.Int64]
	TypeOfUint          = BasicTypes[reflect.Uint]
	TypeOfUint8         = BasicTypes[reflect.Uint8]
	TypeOfUint16        = BasicTypes[reflect.Uint16]
	TypeOfUint32        = BasicTypes[reflect.Uint32]
	TypeOfUint64        = BasicTypes[reflect.Uint64]
	TypeOfUintptr       = BasicTypes[reflect.Uintptr]
	TypeOfFloat32       = BasicTypes[reflect.Float32]
	TypeOfFloat64       = BasicTypes[reflect.Float64]
	TypeOfComplex64     = BasicTypes[reflect.Complex64]
	TypeOfComplex128    = BasicTypes[reflect.Complex128]
	TypeOfString        = BasicTypes[reflect.String]
	TypeOfUnsafePointer = BasicTypes[reflect.UnsafePointer]
	TypeOfError         = universe.TypeOfError
	TypeOfInterface     = universe.TypeOfInterface
	// TypeOfByte       = TypeOfUint8
	// TypeOfRune       = TypeOfInt32
)

// TypeOf creates a Type corresponding to reflect.TypeOf() of given value.
// Note: conversions from Type to reflect.Type and back are not exact,
// because of the reasons listed in Type.ReflectType()
// Conversions from reflect.Type to Type and back are not exact for the same reasons.
func TypeOf(rvalue interface{}) Type {
	return universe.FromReflectType(reflect.TypeOf(rvalue))
}

// FromReflectType creates a Type corresponding to given reflect.Type
// Note: conversions from Type to reflect.Type and back are not exact,
// because of the reasons listed in Type.ReflectType()
// Conversions from reflect.Type to Type and back are not exact for the same reasons.
func FromReflectType(rtype reflect.Type) Type {
	return universe.FromReflectType(rtype)
}

// NamedOf returns a new named type for the given type name and package.
// Initially, the underlying type is set to interface{} - use SetUnderlying to change it.
// These two steps are separate to allow creating self-referencing types,
// as for example type List struct { Elem int; Rest *List }
func NamedOf(name, pkgpath string) Type {
	return universe.NamedOf(name, pkgpath)
}

func NewPackage(path, name string) *Package {
	return universe.NewPackage(path, name)
}

func MakeType(gtype types.Type, rtype reflect.Type) Type {
	return universe.MakeType(gtype, rtype)
}
