// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file sets up the pre-declared methods of a type.

package etypes

import (
	"go/token"
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
	if len(b.methods) != 0 {
		return
	}
	// TODO
}
func (a *Array) initMethods() {
	if len(a.methods) != 0 {
		return
	}
	v := NewVar(token.NoPos, nil, "a", a)
	vint := NewVar(token.NoPos, nil, "", Typ[Int])
	velem := NewVar(token.NoPos, nil, "", a.elem)
	tuple_int := NewTuple(vint)
	tuple_elem := NewTuple(velem)
	tuple_ptrelem := NewTuple(NewVar(token.NoPos, nil, "", NewPointer(a.elem)))
	tuple_int_elem := NewTuple(vint, velem)
	a.methods = []*Func{
		NewFunc(token.NoPos, nil, "Cap", NewSignature(v, nil, tuple_int, false)),
		NewFunc(token.NoPos, nil, "Get", NewSignature(v, tuple_int, tuple_elem, false)),
		NewFunc(token.NoPos, nil, "GetAddr", NewSignature(v, tuple_int, tuple_ptrelem, false)),
		NewFunc(token.NoPos, nil, "Len", NewSignature(v, nil, tuple_int, false)),
		NewFunc(token.NoPos, nil, "Set", NewSignature(v, tuple_int_elem, nil, false)),
	}
}
func (s *Slice) initMethods() {
	if len(s.methods) != 0 {
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
	if len(m.methods) != 0 {
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
	if len(c.methods) != 0 {
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
