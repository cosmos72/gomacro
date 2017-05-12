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
 * global.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package type2

import (
	"go/types"
	"reflect"
)

type Package types.Package

// InterfaceHeader is the internal header of interpreted interfaces
type InterfaceHeader struct {
	// val and typ must be private! otherwise interpreted code may mess with them and break type safety
	val reflect.Value
	typ Type
}

type Method struct {
	Name  string
	Pkg   *Package
	Type  Type          // method type
	Func  reflect.Value // func with receiver as first argument
	Index int           // index for Type.Method
}

type StructField struct {
	// Name is the field name.
	Name string
	// Pkg is the package that qualifies a lower case (unexported)
	// field name. It may be nil for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	Pkg       *Package
	Type      Type              // field type
	Tag       reflect.StructTag // field tag string
	Offset    uintptr           // offset within struct, in bytes
	Index     []int             // index sequence for Type.FieldByIndex
	Anonymous bool              // is an embedded field
}

type xtype struct {
	kind  reflect.Kind
	gtype types.Type
	rtype reflect.Type
	// underlying *timpl // needed? computing it for elements
	//   of arrays, maps, slices, structs is painful.
	//   for function/method arguments is even worse.
}

type Type interface {
	// AssignableTo reports whether a value of the type is assignable to type u.
	AssignableTo(u Type) bool
	// ConvertibleTo reports whether a value of the type is convertible to type u.
	ConvertibleTo(u Type) bool
	// Comparable reports whether values of this type are comparable.
	Comparable() bool
	// GoType returns the go/types.Type corresponding to the given type.
	GoType() types.Type
	// Implements reports whether the type implements the interface type u.
	// It panics if u's Kind is not Interface
	Implements(u Type) bool
	// Name returns the type's name within its package.
	// It returns an empty string for unnamed types.
	Name() string
	// Named returns whether the type is named.
	// It returns false for unnamed types.
	Named() bool
	// Pkg returns a named type's package, that is, the package where it was defined.
	// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
	// Pkg will return nil.
	Pkg() *Package
	// PkgName returns a named type's package name, that is,
	// the default name that the package provides when imported.
	// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
	// the package name will be the empty string.
	PkgName() string
	// PkgPath returns a named type's package path, that is, the import path
	// that uniquely identifies the package, such as "encoding/base64".
	// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
	// the package path will be the empty string.
	PkgPath() string
	// ReflectType returns a best-effort reflect.Type that approximates the type.
	// It may be inexact for the following reasons:
	// 1) missing reflect.NamedOf(): no way to programmatically create named types, or to access the underlying type of a named type
	// 2) missing reflect.InterfaceOf(): interface types created at runtime will be approximated by structs
	// 3) missing reflect.MethodOf(): method types created at runtime will be approximated by functions
	//    whose first parameter is the receiver
	// 4) reflect.StructOf() does not support embedded or unexported fields
	// 5) go/reflect lacks the ability to create self-referencing types:
	//    references to the type itself will be replaced by interface{}.
	//
	// Examples:
	//    after invoking at runtime type2.NewStruct() and type2.NewNamed()
	//    to create the following type:
	//        type List struct { Elem int; Rest *List }
	//    ReflectType will return a reflect.Type equivalent to:
	//        struct { Elem int; Rest interface{} }
	//    i.e. the type name will be missing due to limitation 1 above,
	//    and the field 'Rest' will have type interface{} instead of *List due to limitation 5.
	ReflectType() reflect.Type
	// Size returns the number of bytes needed to store
	// a value of the given type; it is analogous to unsafe.Sizeof.
	Size() uintptr
	// String returns a string representation of a type.
	String() string

	// AddMethod adds method with given name and signature to type, unless it is already in the method list.
	// It panics if the type is unnamed, or if the signature is not a function type.
	AddMethod(name string, signature Type)
	// Bits returns the size of the type in bits.
	// It panics if the type's Kind is not one of the
	// sized or unsized Int, Uint, Float, or Complex kinds.
	Bits() int
	// ChanDir returns a channel type's direction.
	// It panics if the type's Kind is not Chan.
	ChanDir() reflect.ChanDir
	// Complete marks an interface type as complete and computes wrapper methods for embedded fields.
	// It must be called by users of InterfaceOf after the interface's embedded types are fully defined
	// and before using the interface type in any way other than to form other types.
	// Complete returns the receiver.
	Complete() Type
	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	Elem() Type
	// ExplicitMethod return the i-th explicitly declared method of named type or interface t.
	// Wrapper methods for embedded fields or embedded interfaces are not returned - use Method() to get them.
	// It panics if the type is unnamed, or if the type's Kind is not Interface
	ExplicitMethod(i int) Method
	// Field returns a struct type's i-th field.
	// It panics if the type's Kind is not Struct.
	// It panics if i is not in the range [0, NumField()).
	Field(i int) StructField
	// IsMethod reports whether a function type's contains a receiver, i.e. is a method.
	// It panics if the type's Kind is not Func.
	IsMethod() bool
	// IsVariadic reports whether a function type's final input parameter is a "..." parameter.
	// If so, t.In(t.NumIn() - 1) returns the parameter's implicit actual type []T.
	// IsVariadic panics if the type's Kind is not Func.
	IsVariadic() bool
	// Key returns a map type's key type.
	// It panics if the type's Kind is not Map.
	Key() Type
	// Kind returns the specific kind of the type.
	Kind() reflect.Kind
	// Len returns an array type's length.
	// It panics if the type's Kind is not Array.
	Len() int
	// In returns the type of a function type's i'th input parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumIn()).
	In(i int) Type
	// NumExplicitMethods returns the number of explicitly declared methods of named type or interface t.
	// Wrapper methods for embedded fields or embedded interfaces are not counted - use NumMethods() to include them.
	NumExplicitMethods() int
	// NumField returns a struct type's field count.
	// It panics if the type's Kind is not Struct.
	NumField() int
	// NumIn returns a function type's input parameter count.
	// It panics if the type's Kind is not Func.
	NumIn() int
	// NumOut returns a function type's output parameter count.
	// It panics if the type's Kind is not Func.
	NumOut() int
	// Out returns the type of a function type's i'th output parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumOut()).
	Out(i int) Type
	// Recv returns the type of a method type's receiver parameter.
	// It panics if the type's Kind is not Func.
	// It returns nil if t has no receiver.
	Recv() Type
	// SetUnderlying sets the underlying type of a named type and marks it as complete.
	// It panics if the type is unnamed, or if the underlying type is named,
	// or if SetUnderlying() was already invoked on the named type.
	SetUnderlying(underlying Type)
	// underlying returns the underlying types.Type of a type.
	// TODO implement Underlying() Type ?
	// Synthetizing the underlying reflect.Type is not possible for interface types,
	// or for struct types with embedded or unexported fields.
	underlying() types.Type
}
