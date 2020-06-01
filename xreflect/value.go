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
 * value.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	r "reflect"
)

// xreflect.Value is a replacement for reflect.Value
// that hides the presence of xreflect.Forward
// due to emulated recursive types
type Value struct {
	rv r.Value
}

func fromReflectValues(rv []r.Value) []Value {
	v := make([]Value, len(rv))
	for i := range rv {
		v[i] = Value{rv[i]}
	}
	return v
}

func toReflectValues(v []Value) []r.Value {
	rv := make([]r.Value, len(v))
	for i := range v {
		rv[i] = v[i].fwd()
	}
	return rv
}

func (v Value) fwd() r.Value {
	rv := v.rv
	if rv.Kind() == r.Interface && rv.Type() == rTypeOfForward {
		rv = rv.Elem()
	}
	return rv
}

func (v Value) ReflectValue() r.Value {
	return v.fwd()
}

// -----------------------------------------------------------------------------

func (v Value) Addr() Value {
	return Value{v.fwd().Addr()}
}

func (v Value) Bool() bool {
	return v.fwd().Bool()
}

func (v Value) Bytes() []byte {
	return v.fwd().Bytes()
}

func (v Value) Call(in []Value) []Value {
	rv := toReflectValues(in)
	rv = v.fwd().Call(rv)
	return fromReflectValues(rv)
}

func (v Value) CallSlice(in []Value) []Value {
	rv := toReflectValues(in)
	rv = v.fwd().CallSlice(rv)
	return fromReflectValues(rv)
}

func (v Value) CanAddr() bool {
	return v.fwd().CanAddr()
}

func (v Value) CanInterface() bool {
	return v.fwd().CanInterface()
}

func (v Value) CanSet() bool {
	return v.fwd().CanSet()
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
	return fromReflectValues(v.fwd().MapKeys())
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

func (v Value) Set(x Value) {
	v.fwd().Set(x.fwd())
}

func (v Value) SetBool(x bool) {
	v.fwd().SetBool(x)
}

func (v Value) SetBytes(x []byte) {
	v.fwd().SetBytes(x)
}

func (v Value) SetCap(n int) {
	v.fwd().SetCap(n)
}

func (v Value) SetComplex(x complex128) {
	v.fwd().SetComplex(x)
}

func (v Value) SetFloat(x float64) {
	v.fwd().SetFloat(x)
}

func (v Value) SetInt(x int64) {
	v.fwd().SetInt(x)
}

func (v Value) SetLen(n int) {
	v.fwd().SetLen(n)
}

func (v Value) SetMapIndex(key, elem Value) {
	v.fwd().SetMapIndex(key.fwd(), elem.fwd())
}

// func (v Value) SetPointer(x unsafe.Pointer) { ... }

func (v Value) SetString(x string) {
	v.fwd().SetString(x)
}

func (v Value) SetUint(x uint64) {
	v.fwd().SetUint(x)
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
