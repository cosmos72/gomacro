/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2020 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * value.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	r "reflect"
)

// re-export reflect.Kind from this package
type Kind = r.Kind

// re-export reflect.SelectDir from this package
type SelectDir = r.SelectDir

// xreflect.Value is a replacement for reflect.Value
// that hides the presence of xreflect.Forward
// due to emulated recursive types
type Value struct {
	rv r.Value
}

const (
	Invalid       Kind = r.Invalid
	Bool          Kind = r.Bool
	Int           Kind = r.Int
	Int8          Kind = r.Int8
	Int16         Kind = r.Int16
	Int32         Kind = r.Int32
	Int64         Kind = r.Int64
	Uint          Kind = r.Uint
	Uint8         Kind = r.Uint8
	Uint16        Kind = r.Uint16
	Uint32        Kind = r.Uint32
	Uint64        Kind = r.Uint64
	Uintptr       Kind = r.Uintptr
	Float32       Kind = r.Float32
	Float64       Kind = r.Float64
	Complex64     Kind = r.Complex64
	Complex128    Kind = r.Complex128
	Array         Kind = r.Array
	Chan          Kind = r.Chan
	Func          Kind = r.Func
	Interface     Kind = r.Interface
	Map           Kind = r.Map
	Ptr           Kind = r.Ptr
	Slice         Kind = r.Slice
	String        Kind = r.String
	Struct        Kind = r.Struct
	UnsafePointer Kind = r.UnsafePointer
)

const (
	SelectSend    SelectDir = r.SelectSend
	SelectRecv    SelectDir = r.SelectRecv
	SelectDefault SelectDir = r.SelectDefault
)

func FromReflectValues(rv []r.Value) []Value {
	v := make([]Value, len(rv))
	for i := range rv {
		v[i] = Value{rv[i]}
	}
	return v
}

func ToReflectValues(v []Value) []r.Value {
	rv := make([]r.Value, len(v))
	for i := range v {
		rv[i] = v[i].fwd()
	}
	return rv
}

func isfwd(rv r.Value) bool {
	return rv.Kind() == r.Interface && rv.Type() == rTypeOfForward
}

func (v Value) isfwd() bool {
	return isfwd(v.rv)
}

func (v Value) fwd() r.Value {
	rv := v.rv
	if isfwd(rv) {
		rv = rv.Elem()
	}
	return rv
}

func (v Value) ReflectValue() r.Value {
	return v.fwd()
}

// -----------------------------------------------------------------------------

func (v Value) Addr() Value {
	// values wrapped in Forward are not addressable
	// (because values wrapped in interfaces are not addressable)
	// thus just return the pointer to Forward
	return Value{v.rv.Addr()}
}

func (v Value) Bool() bool {
	return v.fwd().Bool()
}

func (v Value) Bytes() []byte {
	return v.fwd().Bytes()
}

func (v Value) Call(in []Value) []Value {
	rv := ToReflectValues(in)
	rv = v.fwd().Call(rv)
	return FromReflectValues(rv)
}

func (v Value) CallSlice(in []Value) []Value {
	rv := ToReflectValues(in)
	rv = v.fwd().CallSlice(rv)
	return FromReflectValues(rv)
}

func (v Value) CanAddr() bool {
	return v.fwd().CanAddr()
}

func (v Value) CanInterface() bool {
	return v.fwd().CanInterface()
}

func (v Value) CanSet() bool {
	return v.rv.CanSet()
}

func (v Value) Cap() int {
	return v.fwd().Cap()
}

func (v Value) Close() {
	v.fwd().Close()
}

func (v Value) Complex() complex128 {
	return v.fwd().Complex()
}

func (v Value) Convert(rt r.Type) Value {
	return Value{v.fwd().Convert(rt)}
}

func (v Value) Elem() Value {
	return Value{v.fwd().Elem()}
}

func (v Value) Field(i int) Value {
	return Value{v.fwd().Field(i)}
}

func (v Value) FieldByIndex(index []int) Value {
	return Value{v.fwd().FieldByIndex(index)}
}

func (v Value) FieldByName(name string) Value {
	return Value{v.fwd().FieldByName(name)}
}

func (v Value) FieldByNameFunc(match func(string) bool) Value {
	return Value{v.fwd().FieldByNameFunc(match)}
}

func (v Value) Float() float64 {
	return v.fwd().Float()
}

func (v Value) Index(i int) Value {
	return Value{v.fwd().Index(i)}
}

func (v Value) Int() int64 {
	return v.fwd().Int()
}

func (v Value) Interface() interface{} {
	return v.fwd().Interface()
}

func (v Value) InterfaceData() [2]uintptr {
	return v.fwd().InterfaceData()
}

func (v Value) IsNil() bool {
	return v.fwd().IsNil()
}

func (v Value) IsValid() bool {
	return v.fwd().IsValid()
}

func (v Value) IsZero() bool {
	return v.fwd().IsZero()
}

func (v Value) Kind() r.Kind {
	return v.fwd().Kind()
}

func (v Value) Len() int {
	return v.fwd().Len()
}

func (v Value) MapIndex(key Value) Value {
	return Value{v.fwd().MapIndex(key.fwd())}
}

func (v Value) MapKeys() []Value {
	return FromReflectValues(v.fwd().MapKeys())
}

// requires Go 1.12
// func (v Value) MapRange() *MapIter { ... }

func (v Value) Method(i int) Value {
	return Value{v.fwd().Method(i)}
}

func (v Value) MethodByName(name string) Value {
	return Value{v.fwd().MethodByName(name)}
}

func (v Value) NumField() int {
	return v.fwd().NumField()
}

func (v Value) NumMethod() int {
	return v.fwd().NumMethod()
}

func (v Value) OverflowComplex(x complex128) bool {
	return v.fwd().OverflowComplex(x)
}

func (v Value) OverflowFloat(x float64) bool {
	return v.fwd().OverflowFloat(x)
}

func (v Value) OverflowInt(x int64) bool {
	return v.fwd().OverflowInt(x)
}

func (v Value) OverflowUint(x uint64) bool {
	return v.fwd().OverflowUint(x)
}

func (v Value) Pointer() uintptr {
	return v.fwd().Pointer()
}

func (v Value) Recv() (x Value, ok bool) {
	rv, ok := v.fwd().Recv()
	return Value{rv}, ok
}

func (v Value) Send(x Value) {
	v.fwd().Send(x.fwd())
}

func (v Value) set(x interface{}) {
	v.rv.Set(r.ValueOf(x))
}

func (v Value) Set(x Value) {
	v.rv.Set(x.fwd())
}

func (v Value) SetBool(x bool) {
	if v.isfwd() {
		v.set(x)
	} else {
		v.rv.SetBool(x)
	}
}

func (v Value) SetBytes(x []byte) {
	if v.isfwd() {
		v.set(x)
	} else {
		v.rv.SetBytes(x)
	}
}

func (v Value) SetCap(n int) {
	v.fwd().SetCap(n)
}

func (v Value) SetComplex(x complex128) {
	if v.isfwd() {
		v.set(x)
	} else {
		v.rv.SetComplex(x)
	}
}

func (v Value) SetFloat(x float64) {
	if v.isfwd() {
		v.set(x)
	} else {
		v.rv.SetFloat(x)
	}
}

func (v Value) SetInt(x int64) {
	if v.isfwd() {
		v.set(x)
	} else {
		v.rv.SetInt(x)
	}
}

func (v Value) SetLen(n int) {
	v.fwd().SetLen(n)
}

func (v Value) SetMapIndex(key, elem Value) {
	v.fwd().SetMapIndex(key.fwd(), elem.fwd())
}

// func (v Value) SetPointer(x unsafe.Pointer) { ... }

func (v Value) SetString(x string) {
	if v.isfwd() {
		v.set(x)
	} else {
		v.rv.SetString(x)
	}
}

func (v Value) SetUint(x uint64) {
	if v.isfwd() {
		v.set(x)
	} else {
		v.rv.SetUint(x)
	}
}

func (v Value) Slice(i, j int) Value {
	return Value{v.fwd().Slice(i, j)}
}

func (v Value) Slice3(i, j, k int) Value {
	return Value{v.fwd().Slice3(i, j, k)}
}

func (v Value) String() string {
	return v.fwd().String()
}

func (v Value) TryRecv() (x Value, ok bool) {
	rv, ok := v.fwd().TryRecv()
	return Value{rv}, ok
}

func (v Value) TrySend(x Value) bool {
	return v.fwd().TrySend(x.fwd())
}

func (v Value) Type() r.Type {
	return v.fwd().Type()
}

func (v Value) Uint() uint64 {
	return v.fwd().Uint()
}

// func (v Value) UnsafeAddr() uintptr { ... }
