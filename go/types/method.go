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

func (b *Basic) NumMethods() int     { return len(b.methods) }
func (a *Array) NumMethods() int     { a.initMethods(); return len(a.methods) }
func (s *Slice) NumMethods() int     { s.initMethods(); return len(s.methods) }
func (s *Struct) NumMethods() int    { return 0 }
func (p *Pointer) NumMethods() int   { return 0 }
func (t *Tuple) NumMethods() int     { return 0 }
func (s *Signature) NumMethods() int { return 0 }
func (m *Map) NumMethods() int       { m.initMethods(); return len(m.methods) }
func (c *Chan) NumMethods() int      { c.initMethods(); return len(c.methods) }

func (b *Basic) Method(i int) *Func     { return b.methods[i] }
func (a *Array) Method(i int) *Func     { a.initMethods(); return a.methods[i] }
func (s *Slice) Method(i int) *Func     { s.initMethods(); return s.methods[i] }
func (s *Struct) Method(i int) *Func    { return ([]*Func)(nil)[i] }
func (p *Pointer) Method(i int) *Func   { return ([]*Func)(nil)[i] }
func (t *Tuple) Method(i int) *Func     { return ([]*Func)(nil)[i] }
func (s *Signature) Method(i int) *Func { return ([]*Func)(nil)[i] }
func (m *Map) Method(i int) *Func       { m.initMethods(); return m.methods[i] }
func (c *Chan) Method(i int) *Func      { c.initMethods(); return c.methods[i] }

func (b *Basic) initMethods() {
	if etoken.GENERICS_V2_CTI && len(b.methods) == 0 {
		b.methods = makeBasicMethods(b, b)
	}
}
func (a *Array) initMethods() {
	if etoken.GENERICS_V2_CTI && len(a.methods) == 0 {
		a.methods = makeArrayMethods(a, a)
	}
}
func (c *Chan) initMethods() {
	if etoken.GENERICS_V2_CTI && len(c.methods) == 0 {
		c.methods = makeChanMethods(c, c)
	}
}
func (m *Map) initMethods() {
	if etoken.GENERICS_V2_CTI && len(m.methods) == 0 {
		m.methods = makeMapMethods(m, m)
	}
}
func (s *Slice) initMethods() {
	if etoken.GENERICS_V2_CTI && len(s.methods) == 0 {
		s.methods = makeSliceMethods(s, s)
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
	if !etoken.GENERICS_V2_CTI || info&IsUntyped != 0 {
		return methods
	}
	v := newVar(t)
	vbool := newVar(Typ[Bool])
	tuple_v := NewTuple(v)
	tuple_bool := NewTuple(vbool)
	sig_vv := NewSignature(v, nil, tuple_v, false)
	sig_vvv := NewSignature(v, tuple_v, tuple_v, false)
	if info&IsNumeric != 0 {
		methods = append(methods,
			newFunc("Add", sig_vvv),
			newFunc("Sub", sig_vvv),
			newFunc("Mul", sig_vvv),
			newFunc("Quo", sig_vvv),
			newFunc("Neg", sig_vv),
		)
	} else if info&IsString != 0 {
		vint := newVar(Typ[Int])
		velem := newVar(Typ[Byte])
		tuple_int := NewTuple(vint)
		tuple_elem := NewTuple(velem)
		methods = append(methods,
			newFunc("Add", sig_vvv),
			newFunc("Get", NewSignature(v, tuple_int, tuple_elem, false)),
			newFunc("Len", NewSignature(v, nil, tuple_int, false)),
		)
	}
	if info&IsInteger != 0 {
		_8 := newVar(Typ[Uint8])
		tuple_8 := NewTuple(_8)
		sig_v8v := NewSignature(v, tuple_v, tuple_8, false)
		methods = append(methods,
			newFunc("Rem", sig_vvv),
			newFunc("And", sig_vvv),
			newFunc("Andnot", sig_vvv),
			newFunc("Or", sig_vvv),
			newFunc("Xor", sig_vvv),
			newFunc("Not", sig_vv), // unary ^
			newFunc("Shl", sig_v8v),
			newFunc("Shr", sig_v8v),
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
			newFunc("Not", sig_vv),
		)
	}
	sig_vvbool := NewSignature(v, tuple_v, tuple_bool, false)
	if info&IsOrdered != 0 {
		methods = append(methods,
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
	if !etoken.GENERICS_V2_CTI {
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
		newFunc("Get", NewSignature(vptr, tuple_int, tuple_elem, false)),
		newFunc("GetAddr", NewSignature(vptr, tuple_int, tuple_ptrelem, false)),
		newFunc("Len", NewSignature(vptr, nil, tuple_int, false)),
		newFunc("Set", NewSignature(vptr, tuple_int_elem, nil, false)),
		newFunc("Slice", NewSignature(vptr, tuple_int_int, tuple_slice, false)),
		newFunc("Slice3", NewSignature(vptr, tuple_int_int_int, tuple_slice, false)),
	}
}

func makeChanMethods(t Type, underlying *Chan) []*Func {
	var methods []*Func
	if !etoken.GENERICS_V2_CTI {
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
	if !etoken.GENERICS_V2_CTI {
		return methods
	}
	v := newVar(t)
	vint := newVar(Typ[Int])
	velem := newVar(underlying.elem)
	tuple_int := NewTuple(vint)
	tuple_elem := NewTuple(velem)
	tuple_int_elem := NewTuple(vint, velem)
	return []*Func{
		newFunc("Delete", NewSignature(v, tuple_int, nil, false)),
		newFunc("Get", NewSignature(v, tuple_int, tuple_elem, false)),
		newFunc("Len", NewSignature(v, nil, tuple_int, false)),
		newFunc("Set", NewSignature(v, tuple_int_elem, nil, false)),
	}
}

func makeSliceMethods(t Type, underlying *Slice) []*Func {
	var methods []*Func
	if !etoken.GENERICS_V2_CTI {
		return methods
	}
	elem := underlying.elem
	v := newVar(t)
	vint := newVar(Typ[Int])
	velem := newVar(elem)
	tuple_v := NewTuple(v)
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
			newFunc("Append", NewSignature(v, tuple_v, tuple_v, true)),
			newFunc("AppendString", NewSignature(v, tuple_string, tuple_v, false)),
			newFunc("Cap", NewSignature(v, nil, tuple_int, false)),
			newFunc("Copy", NewSignature(v, tuple_v, nil, false)),
			newFunc("CopyString", NewSignature(v, tuple_string, nil, false)),
			newFunc("Get", NewSignature(v, tuple_int, tuple_elem, false)),
			newFunc("GetAddr", NewSignature(v, tuple_int, tuple_ptrelem, false)),
			newFunc("Len", NewSignature(v, nil, tuple_int, false)),
			newFunc("Set", NewSignature(v, tuple_int_elem, nil, false)),
			newFunc("Slice", NewSignature(v, tuple_int_int, tuple_v, false)),
			newFunc("Slice3", NewSignature(v, tuple_int_int_int, tuple_v, false)),
		}
	}
	return []*Func{
		newFunc("Append", NewSignature(v, tuple_v, tuple_v, true)),
		newFunc("Cap", NewSignature(v, nil, tuple_int, false)),
		newFunc("Copy", NewSignature(v, tuple_v, nil, false)),
		newFunc("Get", NewSignature(v, tuple_int, tuple_elem, false)),
		newFunc("GetAddr", NewSignature(v, tuple_int, tuple_ptrelem, false)),
		newFunc("Len", NewSignature(v, nil, tuple_int, false)),
		newFunc("Set", NewSignature(v, tuple_int_elem, nil, false)),
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
		return t.methods
	case *Array:
		t.initMethods()
		return t.methods
	case *Slice:
		t.initMethods()
		return t.methods
	case *Map:
		t.initMethods()
		return t.methods
	case *Chan:
		t.initMethods()
		return t.methods
	default:
		return nil
	}
}
