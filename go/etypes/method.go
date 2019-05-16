// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file sets up the pre-declared methods of a type.

package etypes

import (
	"go/token"

	"github.com/cosmos72/gomacro/go/etoken"
)

func (b *Basic) NumMethods() int     { b.initMethods(); return len(b.methods) }
func (a *Array) NumMethods() int     { a.initMethods(); return len(a.methods) }
func (s *Slice) NumMethods() int     { s.initMethods(); return len(s.methods) }
func (s *Struct) NumMethods() int    { return 0 }
func (p *Pointer) NumMethods() int   { return 0 }
func (t *Tuple) NumMethods() int     { return 0 }
func (s *Signature) NumMethods() int { return 0 }
func (m *Map) NumMethods() int       { m.initMethods(); return len(m.methods) }
func (c *Chan) NumMethods() int      { c.initMethods(); return len(c.methods) }

func (b *Basic) Method(i int) *Func     { b.initMethods(); return b.methods[i] }
func (a *Array) Method(i int) *Func     { a.initMethods(); return a.methods[i] }
func (s *Slice) Method(i int) *Func     { s.initMethods(); return s.methods[i] }
func (s *Struct) Method(i int) *Func    { return ([]*Func)(nil)[i] }
func (p *Pointer) Method(i int) *Func   { return ([]*Func)(nil)[i] }
func (t *Tuple) Method(i int) *Func     { return ([]*Func)(nil)[i] }
func (s *Signature) Method(i int) *Func { return ([]*Func)(nil)[i] }
func (m *Map) Method(i int) *Func       { m.initMethods(); return m.methods[i] }
func (c *Chan) Method(i int) *Func      { c.initMethods(); return c.methods[i] }

func (b *Basic) initMethods() {
	if !etoken.GENERICS_V2_CTI || len(b.methods) != 0 || b.info&IsUntyped != 0 {
		return
	}
	info := b.info
	v := NewVar(token.NoPos, nil, "v", b)
	vbool := NewVar(token.NoPos, nil, "", Typ[Bool])
	tuple_v := NewTuple(v)
	tuple_bool := NewTuple(vbool)
	sig_vv := NewSignature(v, nil, tuple_v, false)
	sig_vvv := NewSignature(v, tuple_v, tuple_v, false)
	if info&IsNumeric != 0 {
		b.methods = append(b.methods,
			NewFunc(token.NoPos, nil, "Add", sig_vvv),
			NewFunc(token.NoPos, nil, "Sub", sig_vvv),
			NewFunc(token.NoPos, nil, "Mul", sig_vvv),
			NewFunc(token.NoPos, nil, "Div", sig_vvv),
			NewFunc(token.NoPos, nil, "Neg", sig_vv),
		)
	} else if info&IsString != 0 {
		b.methods = append(b.methods,
			NewFunc(token.NoPos, nil, "Add", sig_vvv),
		)
	}
	if info&IsInteger != 0 {
		_8 := NewVar(token.NoPos, nil, "", Typ[Uint8])
		tuple_8 := NewTuple(_8)
		sig_v8v := NewSignature(v, tuple_v, tuple_8, false)
		b.methods = append(b.methods,
			NewFunc(token.NoPos, nil, "Rem", sig_vvv),
			NewFunc(token.NoPos, nil, "And", sig_vvv),
			NewFunc(token.NoPos, nil, "Or", sig_vvv),
			NewFunc(token.NoPos, nil, "Xor", sig_vvv),
			NewFunc(token.NoPos, nil, "Shl", sig_v8v),
			NewFunc(token.NoPos, nil, "Shr", sig_v8v),
			NewFunc(token.NoPos, nil, "Andnot", sig_vvv),
			NewFunc(token.NoPos, nil, "Not", sig_vv), // unary ^
		)
	} else if info&IsComplex != 0 {
		var fl *Basic
		if b.kind == Complex64 {
			fl = Typ[Float32]
		} else {
			fl = Typ[Float64]
		}
		vfl := NewVar(token.NoPos, nil, "", fl)
		tuple_fl := NewTuple(vfl)
		sig_vfl := NewSignature(v, nil, tuple_fl, false)
		b.methods = append(b.methods,
			NewFunc(token.NoPos, nil, "Real", sig_vfl),
			NewFunc(token.NoPos, nil, "Imag", sig_vfl),
		)
	} else if info&IsBoolean != 0 {
		b.methods = append(b.methods,
			NewFunc(token.NoPos, nil, "Not", sig_vv),
		)
	}
	sig_vvbool := NewSignature(v, tuple_v, tuple_bool, false)
	if info&IsOrdered != 0 {
		b.methods = append(b.methods,
			NewFunc(token.NoPos, nil, "Equal", sig_vvbool),
			NewFunc(token.NoPos, nil, "Less", sig_vvbool),
		)
	} else {
		b.methods = append(b.methods,
			NewFunc(token.NoPos, nil, "Equal", sig_vvbool),
		)
	}
}
func (a *Array) initMethods() {
	if !etoken.GENERICS_V2_CTI || len(a.methods) != 0 {
		return
	}
	v := NewVar(token.NoPos, nil, "a", a)
	vptr := NewVar(token.NoPos, nil, "a", NewPointer(a))
	vint := NewVar(token.NoPos, nil, "", Typ[Int])
	velem := NewVar(token.NoPos, nil, "", a.elem)
	tuple_int := NewTuple(vint)
	tuple_elem := NewTuple(velem)
	tuple_ptrelem := NewTuple(NewVar(token.NoPos, nil, "", NewPointer(a.elem)))
	tuple_int_elem := NewTuple(vint, velem)
	a.methods = []*Func{
		NewFunc(token.NoPos, nil, "Cap", NewSignature(v, nil, tuple_int, false)),
		NewFunc(token.NoPos, nil, "Get", NewSignature(v, tuple_int, tuple_elem, false)),
		NewFunc(token.NoPos, nil, "GetAddr", NewSignature(vptr, tuple_int, tuple_ptrelem, false)),
		NewFunc(token.NoPos, nil, "Len", NewSignature(v, nil, tuple_int, false)),
		NewFunc(token.NoPos, nil, "Set", NewSignature(vptr, tuple_int_elem, nil, false)),
	}
}
func (s *Slice) initMethods() {
	if !etoken.GENERICS_V2_CTI || len(s.methods) != 0 {
		return
	}
	v := NewVar(token.NoPos, nil, "s", s)
	vint := NewVar(token.NoPos, nil, "", Typ[Int])
	velem := NewVar(token.NoPos, nil, "", s.elem)
	tuple_v := NewTuple(v)
	tuple_int := NewTuple(vint)
	tuple_elem := NewTuple(velem)
	tuple_ptrelem := NewTuple(NewVar(token.NoPos, nil, "", NewPointer(s.elem)))
	tuple_int_elem := NewTuple(vint, velem)
	s.methods = []*Func{
		NewFunc(token.NoPos, nil, "Append", NewSignature(v, tuple_v, tuple_v, true)),
		NewFunc(token.NoPos, nil, "Cap", NewSignature(v, nil, tuple_int, false)),
		NewFunc(token.NoPos, nil, "Copy", NewSignature(v, tuple_v, nil, false)),
		NewFunc(token.NoPos, nil, "Get", NewSignature(v, tuple_int, tuple_elem, false)),
		NewFunc(token.NoPos, nil, "GetAddr", NewSignature(v, tuple_int, tuple_ptrelem, false)),
		NewFunc(token.NoPos, nil, "Len", NewSignature(v, nil, tuple_int, false)),
		NewFunc(token.NoPos, nil, "Set", NewSignature(v, tuple_int_elem, nil, false)),
	}
}
func (m *Map) initMethods() {
	if !etoken.GENERICS_V2_CTI || len(m.methods) != 0 {
		return
	}
	v := NewVar(token.NoPos, nil, "m", m)
	vint := NewVar(token.NoPos, nil, "", Typ[Int])
	velem := NewVar(token.NoPos, nil, "", m.elem)
	tuple_int := NewTuple(vint)
	tuple_elem := NewTuple(velem)
	tuple_int_elem := NewTuple(vint, velem)
	m.methods = []*Func{
		NewFunc(token.NoPos, nil, "Delete", NewSignature(v, tuple_int, nil, false)),
		NewFunc(token.NoPos, nil, "Get", NewSignature(v, tuple_int, tuple_elem, false)),
		NewFunc(token.NoPos, nil, "Len", NewSignature(v, nil, tuple_int, false)),
		NewFunc(token.NoPos, nil, "Set", NewSignature(v, tuple_int_elem, nil, false)),
	}
}
func (c *Chan) initMethods() {
	if !etoken.GENERICS_V2_CTI || len(c.methods) != 0 {
		return
	}
	v := NewVar(token.NoPos, nil, "c", c)
	vbool := NewVar(token.NoPos, nil, "", Typ[Bool])
	vint := NewVar(token.NoPos, nil, "", Typ[Int])
	velem := NewVar(token.NoPos, nil, "", c.elem)
	tuple_int := NewTuple(vint)
	tuple_elem := NewTuple(velem)
	tuple_elem_bool := NewTuple(velem, vbool)
	c.methods = []*Func{
		NewFunc(token.NoPos, nil, "Cap", NewSignature(v, nil, tuple_int, false)),
		NewFunc(token.NoPos, nil, "Close", NewSignature(v, nil, nil, false)),
		NewFunc(token.NoPos, nil, "Len", NewSignature(v, nil, tuple_int, false)),
		NewFunc(token.NoPos, nil, "Recv", NewSignature(v, nil, tuple_elem_bool, false)),
		NewFunc(token.NoPos, nil, "Send", NewSignature(v, tuple_elem, nil, false)),
	}
}
