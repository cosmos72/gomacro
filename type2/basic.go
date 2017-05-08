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

package type2

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"

	"go/types"
)

func errorf(format string, arg ...interface{}) {
	panic(errors.New(fmt.Sprintf(format, arg...)))
}

var TypeOfError Type
var TypeOfInterface = maketype(types.NewInterface(nil, nil), reflect.TypeOf((*interface{})(nil)).Elem())

func makebasictypes() []Type {
	src := `package main
func _() {
	_ = bool(false)
	_ = int(0)
	_ = int8(0)
	_ = int16(0)
	_ = int32(0)
	_ = int64(0)
	_ = uint(0)
	_ = uint8(0)
	_ = uint16(0)
	_ = uint32(0)
	_ = uint64(0)
	_ = uintptr(0)
	_ = float32(0)
	_ = float64(0)
	_ = complex64(0)
	_ = complex128(0)
	_ = string("")
	_ = error(nil)
}`
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
		reflect.UnsafePointer: nil, // to set the length
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		errorf("%s: %s", src, err)
	}

	typemap := make(map[ast.Expr]types.TypeAndValue)
	conf := types.Config{}
	_, err = conf.Check(f.Name.Name, fset, []*ast.File{f}, &types.Info{Types: typemap})
	if err != nil {
		errorf("%s: %s", src, err)
	}

	m := make([]Type, len(rmap))
	for x := range typemap {
		if call, _ := x.(*ast.CallExpr); call != nil {
			t := typemap[call].Type
			str := types.ExprString(call)
			if t == nil {
				errorf("no type recorded for %s", str)
			}
			switch t := t.(type) {
			case *types.Basic:
				kind := gbasickindToKind(t.Kind())
				m[kind] = maketype(t, rmap[kind])
				continue
			case *types.Named:
				name := t.Obj().Name()
				if name == "error" {
					TypeOfError = maketype(t, reflect.TypeOf((*error)(nil)).Elem())
					continue
				}
			}
			errorf("type recorded for %s is not a known type: %#v", str, t)
		}
	}
	return m
}

var basictypes = makebasictypes()

var (
	TypeOfBool       = basictypes[reflect.Bool]
	TypeOfInt        = basictypes[reflect.Int]
	TypeOfInt8       = basictypes[reflect.Int8]
	TypeOfInt16      = basictypes[reflect.Int16]
	TypeOfInt32      = basictypes[reflect.Int32]
	TypeOfInt64      = basictypes[reflect.Int64]
	TypeOfUint       = basictypes[reflect.Uint]
	TypeOfUint8      = basictypes[reflect.Uint8]
	TypeOfUint16     = basictypes[reflect.Uint16]
	TypeOfUint32     = basictypes[reflect.Uint32]
	TypeOfUint64     = basictypes[reflect.Uint64]
	TypeOfUintptr    = basictypes[reflect.Uintptr]
	TypeOfFloat32    = basictypes[reflect.Float32]
	TypeOfFloat64    = basictypes[reflect.Float64]
	TypeOfComplex64  = basictypes[reflect.Complex64]
	TypeOfComplex128 = basictypes[reflect.Complex128]
	TypeOfString     = basictypes[reflect.String]
)

func BasicType(kind reflect.Kind) Type {
	return basictypes[kind]
}

// Bits returns the size of the type in bits.
// It panics if the type's Kind is not one of the
// sized or unsized Int, Uint, Float, or Complex kinds.
func (t Type) Bits() int {
	return t.rtype.Bits()
}
