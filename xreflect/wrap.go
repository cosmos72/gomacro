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
 * wrap.go
 *
 *  Created on Jun 01, 2020
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	r "reflect"

	"github.com/cosmos72/gomacro/go/types"
)

type SelectCase = r.SelectCase

func Append(s Value, x ...Value) Value {
	rx := ToReflectValues(x)
	return Value{r.Append(s.fwd(), rx...)}
}

func MakeChan(t Type, n int) Value {
	return Value{r.MakeChan(t.ReflectType(), n)}
}

func MakeFunc(t Type, fn func([]Value) []Value) Value {
	rtype := t.ReflectType()
	if rtype == rtypeOfForward {
		rtype = funcOfForward(t.NumIn(), t.NumOut(), t.IsVariadic())
	}

	rfn := func(args []r.Value) []r.Value {
		v := FromReflectValues(args)
		v = fn(v)
		return ToReflectValues(v)
	}
	return Value{r.MakeFunc(rtype, rfn)}
}

func MakeMap(t Type) Value {
	return Value{r.MakeMap(t.ReflectType())}
}

func MakeMapWithSize(t Type, n int) Value {
	return Value{r.MakeMapWithSize(t.ReflectType(), n)}
}

func MakeSlice(t Type, len, cap int) Value {
	rv := r.MakeSlice(t.ReflectType(), len, cap)
	return Value{fillForward(rv, t)}
}

func NewR(rt r.Type) Value {
	// TODO: recursively initialize any xreflect.Forward ?
	return Value{r.New(rt)}
}

func New(t Type) Value {
	rv := r.New(t.ReflectType())
	fillForward(rv.Elem(), t)
	return Value{rv}
}

func Select(cases []SelectCase) (chosen int, recv Value, ok bool) {
	chosen, v, ok := r.Select(cases)
	return chosen, Value{v}, ok
}

func ValueOf(x interface{}) Value {
	return Value{r.ValueOf(x)}
}

// ZeroT returns a Value representing the zero value for the specified xreflect.Type.
func Zero(t Type) Value {
	return ZeroR(t.ReflectType())
}

// Zero returns a Value representing the zero value for the specified reflect.Type.
func ZeroR(typ r.Type) Value {
	// TODO: recursively initialize any xreflect.Forward ?
	return Value{r.Zero(typ)}
}

// -----------------------------------------------------------------------------

func MakeValue(rv r.Value) Value {
	return Value{rv}
}

// --------------------- MakeFunc() helpers ------------------------------------

func funcOfForward(nin int, nout int, variadic bool) r.Type {
	touts := sliceOfForward(nout)
	if variadic {
		touts = append([]r.Type{}, touts...) // make a copy
		touts[nout-1] = r.SliceOf(rtypeOfForward)
	}
	return r.FuncOf(sliceOfForward(nin), touts, variadic)
}

var (
	cacheSliceForward []r.Type
	rtypeOfForward    = r.TypeOf((*Forward)(nil)).Elem()
)

func sliceOfForward(n int) []r.Type {
	for len(cacheSliceForward) < n {
		cacheSliceForward = append(cacheSliceForward, rtypeOfForward)
	}
	return cacheSliceForward[:n]
}

// --------------------- New() and Make*() helpers -----------------------------

// recursively initialize any xreflect.Forward
func fillForward(rv r.Value, t Type) r.Value {
	xt := unwrap(t)
	// debugf("fillForward: value %+v, type %v, %v", rv.Interface(), t, xt.option)
	if xt == nil || xt.option == OptDefault {
		return rv
	}
	switch rv.Kind() {
	case Array:
		fillForwardSlice(rv, t)
	case Interface:
		if isfwd(rv) {
			rv = fillForwardInterface(rv, t)
		}
	case Ptr:
		rv = fillForwardPtr(rv, t)
	case Slice:
		fillForwardSlice(rv, t)
	case Struct:
		fillForwardStruct(rv, t)
	}
	return rv
}

func fillForwardInterface(rv r.Value, t Type) r.Value {
	xt := unwrap(t)
	rt := xt.rtype
	// debugf("fillForwardInterface: (step 1) value %+v, type %v, kind %v, rtype %v, %v", rv.Interface(), t, xt.kind, rt, xt.option)
	// time.Sleep(time.Second / 10)
	switch xt.kind {
	case Chan:
		relem := xt.elem().resolveFwd().ReflectType()
		rt = r.ChanOf(xt.ChanDir(), relem)
	case Func:
		rt = resolveFwdFunc(xt)
	case Map:
		rkey := xt.Key().resolveFwd().ReflectType()
		relem := xt.elem().resolveFwd().ReflectType()
		rt = r.MapOf(rkey, relem)
	case Ptr:
		relem := xt.elem().resolveFwd().ReflectType()
		rt = r.PtrTo(relem)
	case Slice:
		relem := xt.elem().resolveFwd().ReflectType()
		rt = r.SliceOf(relem)
	}
	// debugf("fillForwardInterface: (step 2) value %+v, type %v, kind %v, rtype %v, %v", rv.Interface(), t, xt.kind, rt, xt.option)
	fill := r.Zero(rt)
	if rv.CanSet() {
		rv.Set(fill)
	} else {
		rv = fill
	}
	return rv
}

// convert reflect type rTypeOfForward to concrete function type
func resolveFwdFunc(xt *xtype) r.Type {
	v := xt.universe
	gsig := xt.gunderlying().(*types.Signature)
	var in []r.Type
	if grecv := gsig.Recv(); grecv != nil {
		in = append(in, v.resolve(grecv.Type()).ReflectType())
	}
	in = append(in, resolveFwdTuple(v, gsig.Params())...)
	out := resolveFwdTuple(v, gsig.Results())
	return r.FuncOf(in, out, gsig.Variadic())
}

func resolveFwdTuple(v *Universe, tuple *types.Tuple) []r.Type {
	if tuple == nil {
		return nil
	}
	rtypes := make([]r.Type, tuple.Len())
	for i, n := 0, tuple.Len(); i < n; i++ {
		rtypes[i] = v.resolve(tuple.At(i).Type()).ReflectType()
	}
	return rtypes
}

func fillForwardSlice(rv r.Value, t Type) {
	telem := t.Elem()
	for i, n := 0, rv.Len(); i < n; i++ {
		rvi := rv.Index(i)
		rvi.Set(fillForward(rvi, telem))
	}
}

func fillForwardPtr(rv r.Value, t Type) r.Value {
	// debugf("fillForwardPtr: %+v type %v", rv.Interface(), t)
	fill := r.Zero(t.ReflectType())
	if rv.CanSet() {
		rv.Set(fill)
	} else {
		rv = fill
	}
	return rv
}

func fillForwardStruct(rv r.Value, t Type) {
	// debugf("fillForwardStruct: %+v type %v", rv.Interface(), t)
	rt := t.ReflectType()
	for i, n := 0, rv.NumField(); i < n; i++ {
		field := t.Field(i)
		tfield := field.Type
		rfield := rv.Field(i)
		// debugf("fillForwardStruct field %d: %q %+v type %v", i, field.Name, rv.Interface(), t)

		if rt.Field(i).Type == rTypeOfForward {
			fillForwardInterface(rfield, tfield)
		} else if unwrap(tfield).option != OptDefault {
			fillForward(rfield, tfield)
		}
	}
}
