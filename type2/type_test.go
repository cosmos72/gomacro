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
 * type_test.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package type2

import (
	"go/types"
	"reflect"
	"testing"
)

func fail(t *testing.T, actual interface{}, expected interface{}) {
	t.Errorf("expecting %v <%T>, found %v <%T>\n", expected, expected, actual, actual)
}

func is(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		fail(t, actual, expected)
	}
}

func isdeepequal(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		fail(t, actual, expected)
	}
}

func istype(t *testing.T, actual interface{}, expected interface{}) {
	is(t, reflect.TypeOf(actual), reflect.TypeOf(expected))
}

func TestBasicTypes(t *testing.T) {
	rmap := []reflect.Type{
		reflect.Bool:       reflect.TypeOf(bool(false)),
		reflect.Int:        reflect.TypeOf(int(0)),
		reflect.Int8:       reflect.TypeOf(int8(0)),
		reflect.Int16:      reflect.TypeOf(int16(0)),
		reflect.Int32:      reflect.TypeOf(int32(0)),
		reflect.Int64:      reflect.TypeOf(int64(0)),
		reflect.Uint:       reflect.TypeOf(uint(0)),
		reflect.Uint8:      reflect.TypeOf(uint8(0)),
		reflect.Uint16:     reflect.TypeOf(uint16(0)),
		reflect.Uint32:     reflect.TypeOf(uint32(0)),
		reflect.Uint64:     reflect.TypeOf(uint64(0)),
		reflect.Uintptr:    reflect.TypeOf(uintptr(0)),
		reflect.Float32:    reflect.TypeOf(float32(0)),
		reflect.Float64:    reflect.TypeOf(float64(0)),
		reflect.Complex64:  reflect.TypeOf(complex64(0)),
		reflect.Complex128: reflect.TypeOf(complex128(0)),
		reflect.String:     reflect.TypeOf(string("")),
	}
	for i, rtype := range rmap {
		if rtype == nil {
			continue
		}
		kind := reflect.Kind(i)
		typ := BasicType(kind)
		is(t, typ.Kind(), rtype.Kind())
		is(t, typ.Name(), rtype.Name())
		is(t, typ.ReflectType(), rtype)
		istype(t, typ.GoType(), (*types.Basic)(nil))

		basic := typ.GoType().(*types.Basic)
		k := gbasickindToKind(basic.Kind())
		is(t, k, rtype.Kind())
	}
}

func TestCompositeTypes(t *testing.T) {
	typ := ArrayOf(7, TypeOfUint8)
	rtype := reflect.TypeOf([7]uint8{})
	is(t, typ.Kind(), reflect.Array)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Array)(nil))
}

func TestFunctionTypes(t *testing.T) {
	typ := FuncOf([]Type{TypeOfBool, TypeOfInt16}, []Type{TypeOfString}, false)
	rtype := reflect.TypeOf(func(bool, int16) string { return "" })
	is(t, typ.Kind(), reflect.Func)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Signature)(nil))
}

func TestStructTypes(t *testing.T) {
	fields := []StructField{
		StructField{
			Name: "First",
			Type: TypeOfInt,
		},
		StructField{
			Name: "Rest",
			Type: TypeOfInterface,
		},
	}
	typ := StructOf(fields)
	rtype := reflect.TypeOf(struct {
		First int
		Rest  interface{}
	}{})
	is(t, typ.Kind(), reflect.Struct)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Struct)(nil))
	is(t, typ.NumField(), rtype.NumField())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		rfield1 := field.toReflectField(false)
		rfield2 := rtype.Field(i)
		isdeepequal(t, rfield1, rfield2)
	}
}

func TestFromReflect1(t *testing.T) {
	rtype := reflect.TypeOf((*func(bool, int8, <-chan uint16, []float32, [2]float64, []complex64) map[interface{}]*string)(nil)).Elem()
	typ := FromReflectType(rtype, RebuildReflectType)
	is(t, typ.ReflectType(), rtype) // recreated 100% accurately?
}

func TestFromReflect2(t *testing.T) {
	type Bag struct {
		C <-chan bool
		I int32
		U uintptr
		F [3]float32
		G *float64
		M map[string]*[]complex64
	}
	rtype := reflect.TypeOf(Bag{})
	typ := FromReflectType(rtype, RebuildReflectType)
	rtype2 := typ.ReflectType()
	is(t, typ.Kind(), reflect.Struct)
	is(t, typ.Name(), "Bag")
	is(t, rtype.ConvertibleTo(rtype2), true)
	is(t, rtype2.ConvertibleTo(rtype), true)
	is(t, rtype.AssignableTo(rtype2), true)
	is(t, rtype2.AssignableTo(rtype), true)
}
