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
	"io"
	"reflect"
	"testing"
)

func fail(t *testing.T, actual interface{}, expected interface{}) {
	t.Errorf("expecting %v <%T>, found %v <%T>\n", expected, expected, actual, actual)
}

func fail2(t *testing.T, actual interface{}, expected interface{}) {
	t.Errorf("expecting %#v <%T>,\n\tfound %#v <%T>\n", expected, expected, actual, actual)
}

func is(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		fail(t, actual, expected)
	}
}

func isdeepequal(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		fail2(t, actual, expected)
	}
}

func istype(t *testing.T, actual interface{}, expected interface{}) {
	is(t, reflect.TypeOf(actual), reflect.TypeOf(expected))
}

func TestBasic(t *testing.T) {
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
		typ := BasicTypes[kind]
		is(t, typ.Kind(), rtype.Kind())
		is(t, typ.Name(), rtype.Name())
		is(t, typ.ReflectType(), rtype)
		istype(t, typ.GoType(), (*types.Basic)(nil))

		basic := typ.GoType().(*types.Basic)
		k := gbasickindToKind(basic.Kind())
		is(t, k, rtype.Kind())
	}
}

func TestArray(t *testing.T) {
	typ := ArrayOf(7, TypeOfUint8)
	rtype := reflect.TypeOf([7]uint8{})
	is(t, typ.Kind(), reflect.Array)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Array)(nil))
	is(t, typ.String(), "[7]uint8")
}

func TestFunction(t *testing.T) {
	typ := FuncOf([]Type{TypeOfBool, TypeOfInt16}, []Type{TypeOfString}, false)
	rtype := reflect.TypeOf(func(bool, int16) string { return "" })
	is(t, typ.Kind(), reflect.Func)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Signature)(nil))
	is(t, typ.String(), "func(bool, int16) string")
}

func TestInterface(t *testing.T) {
	methodtyp := FuncOf(nil, []Type{TypeOfInt}, false)
	typ := InterfaceOf([]string{"Cap", "Len"}, []Type{methodtyp, methodtyp}, nil).Complete()

	is(t, typ.Kind(), reflect.Interface)
	is(t, typ.Name(), "")
	is(t, typ.NumExplicitMethods(), 2)
	actual := typ.ExplicitMethod(0)
	is(t, actual.Name, "Cap")
	is(t, true, types.Identical(methodtyp.GoType(), actual.Type.GoType()))
	actual = typ.ExplicitMethod(1)
	is(t, actual.Name, "Len")
	is(t, true, types.Identical(methodtyp.GoType(), actual.Type.GoType()))
	istype(t, typ.GoType(), (*types.Interface)(nil))

	rtype := reflect.StructOf([]reflect.StructField{
		reflect.StructField{Name: StrGensymInterface, Type: TypeOfInterface.ReflectType()},
		reflect.StructField{Name: "Cap", Type: methodtyp.ReflectType()},
		reflect.StructField{Name: "Len", Type: methodtyp.ReflectType()},
	})
	is(t, typ.ReflectType(), rtype)
	is(t, typ.String(), "interface{Cap() int; Len() int}")
}

func TestMap(t *testing.T) {
	typ := MapOf(TypeOfInterface, TypeOfBool)
	rtype := reflect.TypeOf(map[interface{}]bool{})
	is(t, typ.Kind(), reflect.Map)
	is(t, typ.Name(), "")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Map)(nil))
}

func TestNamed(t *testing.T) {
	typ := NamedOf("MyMap", NewPackage("main", ""))
	underlying := MapOf(TypeOfInterface, TypeOfBool)
	typ.SetUnderlying(underlying)
	rtype := reflect.TypeOf(map[interface{}]bool{})
	is(t, typ.Kind(), reflect.Map)
	is(t, typ.Name(), "MyMap")
	is(t, typ.ReflectType(), rtype)
	istype(t, typ.GoType(), (*types.Named)(nil))
}

func TestSelfReference(t *testing.T) {
	typ := NamedOf("List", NewPackage("main", ""))
	underlying := StructOf([]StructField{
		StructField{Name: "First", Type: TypeOfInt},
		StructField{Name: "Rest", Type: typ},
	})
	typ.SetUnderlying(underlying)
	rtype := reflect.TypeOf(struct {
		First int
		Rest  interface{}
	}{})
	is(t, typ.Kind(), reflect.Struct)
	is(t, typ.Name(), "List")
	is(t, typ.ReflectType(), rtype)
	is(t, true, types.Identical(typ.Field(1).Type.GoType(), typ.GoType()))
	istype(t, typ.GoType(), (*types.Named)(nil))

	is(t, typ.String(), "main.List")
	is(t, typ.underlying().String(), "struct{First int; Rest main.List}")
}

func TestStruct(t *testing.T) {
	typ := StructOf([]StructField{
		StructField{Name: "First", Type: TypeOfInt},
		StructField{Name: "Rest", Type: TypeOfInterface},
	})
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
	is(t, typ.String(), "struct{First int; Rest interface{}}")
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
		G []float64
		M map[string]*complex64
	}
	in := reflect.TypeOf(Bag{})
	expected := reflect.TypeOf(struct {
		C <-chan bool
		I int32
		U uintptr
		F [3]float32
		G []float64
		M map[string]*complex64
	}{})
	typ := FromReflectType(in, RebuildReflectType)
	actual := typ.ReflectType()
	is(t, typ.Kind(), reflect.Struct)
	is(t, typ.Name(), "Bag")
	is(t, actual, expected)
	is(t, actual.ConvertibleTo(in), true)
	is(t, in.ConvertibleTo(actual), true)
	is(t, actual.AssignableTo(in), true)
	is(t, in.AssignableTo(actual), true)
}

func TestFromReflect3(t *testing.T) {
	rtype := reflect.TypeOf((*io.Reader)(nil)).Elem()
	typ := FromReflectType(rtype, RebuildReflectType)

	rmethod := reflect.FuncOf(
		[]reflect.Type{reflect.SliceOf(TypeOfUint8.rtype)},
		[]reflect.Type{TypeOfInt.rtype, TypeOfError.rtype},
		false,
	)
	expected := reflect.StructOf([]reflect.StructField{
		reflect.StructField{Name: StrGensymInterface, Type: TypeOfInterface.rtype},
		reflect.StructField{Name: "Read", Type: rmethod},
	})
	actual := typ.ReflectType()
	is(t, typ.Kind(), reflect.Interface)
	is(t, actual, expected)
	is(t, typ.underlying().String(), "interface{Read([]uint8) (int, error)}")
}
