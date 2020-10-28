// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file sets up the pre-declared methods of a type.
// Needed by Go generics implementation "contracts are interfaces"

package types

import (
	"go/token"

	"github.com/cosmos72/gomacro/go/etoken"
)

// declare CTI methods on basic types, Array, Chan, Map, Slice
// and named types wrapping them

func (b *Basic) NumMethods() int     { return len(b.methods()) }
func (a *Array) NumMethods() int     { return len(a.methods()) }
func (c *Chan) NumMethods() int      { return len(c.methods()) }
func (m *Map) NumMethods() int       { return len(m.methods()) }
func (p *Pointer) NumMethods() int   { return 0 }
func (s *Signature) NumMethods() int { return 0 }
func (s *Slice) NumMethods() int     { return len(s.methods()) }
func (s *Struct) NumMethods() int    { return 0 }
func (t *Tuple) NumMethods() int     { return 0 }

func (b *Basic) Method(i int) *Func     { return b.methods()[i] }
func (a *Array) Method(i int) *Func     { return a.methods()[i] }
func (c *Chan) Method(i int) *Func      { return c.methods()[i] }
func (m *Map) Method(i int) *Func       { return m.methods()[i] }
func (p *Pointer) Method(i int) *Func   { return ([]*Func)(nil)[i] }
func (s *Signature) Method(i int) *Func { return ([]*Func)(nil)[i] }
func (s *Slice) Method(i int) *Func     { return s.methods()[i] }
func (s *Struct) Method(i int) *Func    { return ([]*Func)(nil)[i] }
func (t *Tuple) Method(i int) *Func     { return ([]*Func)(nil)[i] }

func (b *Basic) methods() []*Func {
	if !etoken.GENERICS.V2_CTI() {
		b.lazymethods = nil
	} else if len(b.lazymethods) == 0 {
		b.lazymethods = makeBasicMethods(b, b)
	}
	return b.lazymethods
}
func (a *Array) methods() []*Func {
	if !etoken.GENERICS.V2_CTI() {
		a.lazymethods = nil
	} else if len(a.lazymethods) == 0 {
		a.lazymethods = makeArrayMethods(a, a)
	}
	return a.lazymethods
}
func (c *Chan) methods() []*Func {
	if !etoken.GENERICS.V2_CTI() {
		c.lazymethods = nil
	} else if len(c.lazymethods) == 0 {
		c.lazymethods = makeChanMethods(c, c)
	}
	return c.lazymethods
}
func (m *Map) methods() []*Func {
	if !etoken.GENERICS.V2_CTI() {
		m.lazymethods = nil
	} else if len(m.lazymethods) == 0 {
		m.lazymethods = makeMapMethods(m, m)
	}
	return m.lazymethods
}
func (s *Slice) methods() []*Func {
	if !etoken.GENERICS.V2_CTI() {
		s.lazymethods = nil
	} else if len(s.lazymethods) == 0 {
		s.lazymethods = makeSliceMethods(s, s)
	}
	return s.lazymethods
}

func (t *Named) initMethods() {
	if etoken.GENERICS.V2_CTI() && len(t.methods) == 0 {
		var methods []*Func
		switch u := t.underlying.(type) {
		case *Basic:
			methods = makeBasicMethods(t, u)
		case *Array:
			methods = makeArrayMethods(t, u)
		case *Chan:
			methods = makeChanMethods(t, u)
		case *Map:
			methods = makeMapMethods(t, u)
		case *Slice:
			methods = makeSliceMethods(t, u)
		}
		t.methods = methods
	}
}

func newVar(t Type) *Var {
	return NewVar(token.NoPos, nil, "", t)
}

func newFunc(name string, sig *Signature) *Func {
	return NewFunc(token.NoPos, nil, name, sig)
}

func makeBasicMethods(t Type, underlying *Basic) []*Func {
	var methods []*Func
	info := underlying.info
	if !etoken.GENERICS.V2_CTI() || info&IsUntyped != 0 {
		return methods
	}
	v := newVar(t)
	vbool := newVar(Typ[Bool])
	vint := newVar(Typ[Int])
	tuple_v := NewTuple(v)
	tuple_vv := NewTuple(v, v)
	tuple_bool := NewTuple(vbool)
	tuple_int := NewTuple(vint)
	sig_unary := NewSignature(v, tuple_v, tuple_v, false)
	sig_binary := NewSignature(v, tuple_vv, tuple_v, false)
	if info&IsNumeric != 0 {
		methods = append(methods,
			newFunc("Add", sig_binary),
			newFunc("Sub", sig_binary),
			newFunc("Mul", sig_binary),
			newFunc("Quo", sig_binary),
			newFunc("Neg", sig_unary),
		)
	} else if info&IsString != 0 {
		velem := newVar(Typ[Byte])
		tuple_int_int := NewTuple(vint, vint)
		tuple_elem := NewTuple(velem)
		methods = append(methods,
			newFunc("Add", sig_binary),
			newFunc("Index", NewSignature(v, tuple_int, tuple_elem, false)),
			newFunc("Len", NewSignature(v, nil, tuple_int, false)),
			newFunc("Slice", NewSignature(v, tuple_int_int, tuple_v, false)),
		)
	}
	if info&IsInteger != 0 {
		_8 := newVar(Typ[Uint8])
		tuple_v8 := NewTuple(v, _8)
		sig_vv8v := NewSignature(v, tuple_v8, tuple_v, false)
		methods = append(methods,
			newFunc("Rem", sig_binary),
			newFunc("And", sig_binary),
			newFunc("AndNot", sig_binary),
			newFunc("Or", sig_binary),
			newFunc("Xor", sig_binary),
			newFunc("Not", sig_unary), // unary ^
			newFunc("Lsh", sig_vv8v),  // left shift <<
			newFunc("Rsh", sig_vv8v),  // right shift >>
		)
	} else if info&IsComplex != 0 {
		var fl *Basic
		if underlying.kind == Complex64 {
			fl = Typ[Float32]
		} else {
			fl = Typ[Float64]
		}
		vfl := newVar(fl)
		tuple_fl := NewTuple(vfl)
		sig_vfl := NewSignature(v, nil, tuple_fl, false)
		methods = append(methods,
			newFunc("Real", sig_vfl),
			newFunc("Imag", sig_vfl),
		)
	} else if info&IsBoolean != 0 {
		methods = append(methods,
			newFunc("Not", sig_unary),
		)
	}
	sig_vvbool := NewSignature(v, tuple_v, tuple_bool, false)
	if info&IsOrdered != 0 {
		sig_vvint := NewSignature(v, tuple_v, tuple_int, false)
		methods = append(methods,
			newFunc("Cmp", sig_vvint),
			newFunc("Equal", sig_vvbool),
			newFunc("Less", sig_vvbool),
		)
	} else {
		methods = append(methods,
			newFunc("Equal", sig_vvbool),
		)
	}
	shellsortFuncs(methods)
	return methods
}

func makeArrayMethods(t Type, underlying *Array) []*Func {
	var methods []*Func
	if !etoken.GENERICS.V2_CTI() {
		return methods
	}
	vptr := newVar(NewPointer(t))
	vint := newVar(Typ[Int])
	velem := newVar(underlying.elem)
	vslice := newVar(NewSlice(underlying.elem))
	tuple_int := NewTuple(vint)
	tuple_int_int := NewTuple(vint, vint)
	tuple_int_int_int := NewTuple(vint, vint, vint)
	tuple_elem := NewTuple(velem)
	tuple_ptrelem := NewTuple(newVar(NewPointer(underlying.elem)))
	tuple_int_elem := NewTuple(vint, velem)
	tuple_slice := NewTuple(vslice)
	// receiver is pointer-to-array to avoid hidden O(N) cost of array copy
	return []*Func{
		newFunc("Cap", NewSignature(vptr, nil, tuple_int, false)),
		newFunc("Copy", NewSignature(vptr, tuple_slice, nil, false)),
		// TODO CopyString
		newFunc("Index", NewSignature(vptr, tuple_int, tuple_elem, false)),
		newFunc("AddrIndex", NewSignature(vptr, tuple_int, tuple_ptrelem, false)),
		newFunc("Len", NewSignature(vptr, nil, tuple_int, false)),
		newFunc("SetIndex", NewSignature(vptr, tuple_int_elem, nil, false)),
		newFunc("Slice", NewSignature(vptr, tuple_int_int, tuple_slice, false)),
		newFunc("Slice3", NewSignature(vptr, tuple_int_int_int, tuple_slice, false)),
	}
}

func makeChanMethods(t Type, underlying *Chan) []*Func {
	var methods []*Func
	if !etoken.GENERICS.V2_CTI() {
		return methods
	}
	v := newVar(t)
	vbool := newVar(Typ[Bool])
	vint := newVar(Typ[Int])
	velem := newVar(underlying.elem)
	tuple_int := NewTuple(vint)
	tuple_bool := NewTuple(vbool)
	tuple_elem := NewTuple(velem)
	tuple_elem_bool := NewTuple(velem, vbool)
	methods = []*Func{
		newFunc("Cap", NewSignature(v, nil, tuple_int, false)),
		newFunc("Close", NewSignature(v, nil, nil, false)),
		newFunc("Len", NewSignature(v, nil, tuple_int, false)),
	}
	dir := underlying.dir
	if dir == SendRecv || dir == RecvOnly {
		methods = append(methods,
			newFunc("Recv", NewSignature(v, nil, tuple_elem_bool, false)),
			newFunc("TryRecv", NewSignature(v, nil, tuple_elem_bool, false)),
		)
	}
	if dir == SendRecv || dir == SendOnly {
		methods = append(methods,
			newFunc("Send", NewSignature(v, tuple_elem, nil, false)),
			newFunc("TrySend", NewSignature(v, tuple_elem, tuple_bool, false)),
		)
	}
	return methods
}

func makeMapMethods(t Type, underlying *Map) []*Func {
	var methods []*Func
	if !etoken.GENERICS.V2_CTI() {
		return methods
	}
	v := newVar(t)
	vbool := newVar(Typ[Bool])
	vint := newVar(Typ[Int])
	vkey := newVar(underlying.key)
	velem := newVar(underlying.elem)
	tuple_int := NewTuple(vint)
	tuple_key := NewTuple(vkey)
	tuple_elem := NewTuple(velem)
	tuple_elem_bool := NewTuple(velem, vbool)
	tuple_key_elem := NewTuple(vkey, velem)
	return []*Func{
		newFunc("DelIndex", NewSignature(v, tuple_key, nil, false)),
		newFunc("Index", NewSignature(v, tuple_key, tuple_elem, false)),
		newFunc("Len", NewSignature(v, nil, tuple_int, false)),
		newFunc("SetIndex", NewSignature(v, tuple_key_elem, nil, false)),
		newFunc("TryIndex", NewSignature(v, tuple_key, tuple_elem_bool, false)),
	}
}

func makeSliceMethods(t Type, underlying *Slice) []*Func {
	var methods []*Func
	if !etoken.GENERICS.V2_CTI() {
		return methods
	}
	elem := underlying.elem
	v := newVar(t)
	vint := newVar(Typ[Int])
	velem := newVar(elem)
	tuple_v := NewTuple(v)
	tuple_slice := tuple_v
	if _, ok := t.(*Slice); !ok {
		// last argument of variadic method Append must be unnamed slice
		tuple_slice = NewTuple(newVar(NewSlice(elem)))
	}
	tuple_int := NewTuple(vint)
	tuple_int_int := NewTuple(vint, vint)
	tuple_int_int_int := NewTuple(vint, vint, vint)
	tuple_elem := NewTuple(velem)
	tuple_ptrelem := NewTuple(newVar(NewPointer(elem)))
	tuple_int_elem := NewTuple(vint, velem)
	if elem == Typ[Uint8] || elem == Universe.Lookup("byte").Type() {
		// special case: also has methods AppendString and CopyString
		tuple_string := NewTuple(newVar(Typ[String]))
		return []*Func{
			newFunc("Append", NewSignature(v, tuple_slice, tuple_v, true)),
			newFunc("AppendString", NewSignature(v, tuple_string, tuple_v, false)),
			newFunc("Cap", NewSignature(v, nil, tuple_int, false)),
			newFunc("Copy", NewSignature(v, tuple_v, nil, false)),
			newFunc("CopyString", NewSignature(v, tuple_string, nil, false)),
			newFunc("Index", NewSignature(v, tuple_int, tuple_elem, false)),
			newFunc("AddrIndex", NewSignature(v, tuple_int, tuple_ptrelem, false)),
			newFunc("Len", NewSignature(v, nil, tuple_int, false)),
			newFunc("SetIndex", NewSignature(v, tuple_int_elem, nil, false)),
			newFunc("Slice", NewSignature(v, tuple_int_int, tuple_v, false)),
			newFunc("Slice3", NewSignature(v, tuple_int_int_int, tuple_v, false)),
		}
	}
	return []*Func{
		newFunc("Append", NewSignature(v, tuple_slice, tuple_v, true)),
		newFunc("Cap", NewSignature(v, nil, tuple_int, false)),
		newFunc("Copy", NewSignature(v, tuple_v, nil, false)),
		newFunc("Index", NewSignature(v, tuple_int, tuple_elem, false)),
		newFunc("AddrIndex", NewSignature(v, tuple_int, tuple_ptrelem, false)),
		newFunc("Len", NewSignature(v, nil, tuple_int, false)),
		newFunc("SetIndex", NewSignature(v, tuple_int_elem, nil, false)),
		newFunc("Slice", NewSignature(v, tuple_int_int, tuple_v, false)),
		newFunc("Slice3", NewSignature(v, tuple_int_int_int, tuple_v, false)),
	}
}

// array indexing is faster that slice indexing,
// provided the array is *not* copied. so use a pointer to array
var shellshort_gaps = &[...]int{ /*701, 301, 132, 57,*/ 23, 10, 4, 1}

func shellsortFuncs(vf []*Func) {
	var i, j, n, gap int
	var f *Func
	n = len(vf)
	for _, gap = range shellshort_gaps {
		for i = gap; i < n; i++ {
			f = vf[i]
			for j = i; j >= gap && vf[j-gap].name > f.name; j -= gap {
				vf[j] = vf[j-gap]
			}
			vf[j] = f
		}
	}
}

func declaredMethods(t Type) []*Func {
	switch t := t.(type) {
	case *Named:
		return t.methods
	case *Basic:
		return t.methods()
	case *Array:
		return t.methods()
	case *Slice:
		return t.methods()
	case *Map:
		return t.methods()
	case *Chan:
		return t.methods()
	default:
		return nil
	}
}
