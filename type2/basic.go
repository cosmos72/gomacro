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

func makebasictypes() map[string]Type {
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
	rmap := map[string]reflect.Type{
		"bool":       reflect.TypeOf(bool(false)),
		"int":        reflect.TypeOf(int(0)),
		"int8":       reflect.TypeOf(int8(0)),
		"int16":      reflect.TypeOf(int16(0)),
		"int32":      reflect.TypeOf(int32(0)),
		"int64":      reflect.TypeOf(int64(0)),
		"uint":       reflect.TypeOf(uint(0)),
		"uint8":      reflect.TypeOf(uint8(0)),
		"uint16":     reflect.TypeOf(uint16(0)),
		"uint32":     reflect.TypeOf(uint32(0)),
		"uint64":     reflect.TypeOf(uint64(0)),
		"uintptr":    reflect.TypeOf(uintptr(0)),
		"float32":    reflect.TypeOf(float32(0)),
		"float64":    reflect.TypeOf(float64(0)),
		"complex64":  reflect.TypeOf(complex64(0)),
		"complex128": reflect.TypeOf(complex128(0)),
		"string":     reflect.TypeOf(string("")),
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

	m := make(map[string]Type)
	for x := range typemap {
		if call, _ := x.(*ast.CallExpr); call != nil {
			t := typemap[call].Type
			str := types.ExprString(call)
			if t == nil {
				errorf("no type recorded for %s", str)
			}
			switch t := t.(type) {
			case *types.Basic:
				name := t.Name()
				m[name] = maketype(t, rmap[name])
				continue
			case *types.Named:
				name := t.Obj().Name()
				if name == "error" {
					TypeOfError = maketype(t, reflect.TypeOf((*error)(nil)).Elem())
					fmt.Printf("%s\t %#v\n", name, t)
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
	TypeOfBool       = basictypes["bool"]
	TypeOfInt        = basictypes["int"]
	TypeOfInt8       = basictypes["int8"]
	TypeOfInt16      = basictypes["int16"]
	TypeOfInt32      = basictypes["int32"]
	TypeOfInt64      = basictypes["int64"]
	TypeOfUint       = basictypes["uint"]
	TypeOfUint8      = basictypes["uint8"]
	TypeOfUint16     = basictypes["uint16"]
	TypeOfUint32     = basictypes["uint32"]
	TypeOfUint64     = basictypes["uint64"]
	TypeOfUintptr    = basictypes["uintptr"]
	TypeOfFloat32    = basictypes["float32"]
	TypeOfFloat64    = basictypes["float64"]
	TypeOfComplex64  = basictypes["complex64"]
	TypeOfComplex128 = basictypes["complex128"]
	TypeOfString     = basictypes["string"]
)

func BasicType(name string) Type {
	return basictypes[name]
}

// Bits returns the size of the type in bits.
// It panics if the type's Kind is not one of the
// sized or unsized Int, Uint, Float, or Complex kinds.
func (t *timpl) Bits() int {
	return t.rtype.Bits()
}
