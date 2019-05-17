// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package etypes

import (
	"go/token"
	"testing"

	"github.com/cosmos72/gomacro/go/etoken"
)

func mkVar(t Type) *Var {
	return NewVar(token.NoPos, nil, "", t)
}

func mkTuple(ts ...Type) *Tuple {
	vs := make([]*Var, len(ts))
	for i := range ts {
		vs[i] = mkVar(ts[i])
	}
	return NewTuple(vs...)
}

func mkFunc(name string, params *Tuple, results *Tuple) *Func {
	return NewFunc(token.NoPos, nil, name, NewSignature(nil, params, results, false))
}

func mkInterface(fs ...*Func) *Interface {
	return NewInterface(fs, nil).Complete()
}

/**
 * return
 * interface {
 *   Cap() int
 *   Len() int
 * }
 */
func mkInterfaceCapLen() *Interface {
	return mkInterface(
		mkFunc("Cap", nil, mkTuple(Typ[Int])),
		mkFunc("Len", nil, mkTuple(Typ[Int])),
	)
}

/**
 * return
 * interface {
 *   GetAddr(k Key) *Value
 * }
 */
func mkInterfaceGetAddr(key, value Type) *Interface {
	return mkInterface(
		mkFunc("GetAddr", mkTuple(key), mkTuple(NewPointer(value))),
	)
}

/**
 * return
 * interface {
 *   Get(k Key) Value
 *   Len() int
 * }
 */
func mkInterfaceGetLen(key, value Type) *Interface {
	return mkInterface(
		mkFunc("Get", mkTuple(key), mkTuple(value)),
		mkFunc("Len", nil, mkTuple(Typ[Int])),
	)
}

/**
 * return
 * interface {
 *   Send(e Elem)
 *   Recv() (Elem, bool)
 * }
 */
func mkInterfaceSendRecv(elem Type) *Interface {
	return mkInterface(
		mkFunc("Send", mkTuple(elem), nil),
		mkFunc("Recv", mkTuple(elem, Typ[Bool]), nil),
	)
}

/**
 * return
 * interface {
 *   Set(k Key, v Value)
 * }
 */
func mkInterfaceSet(key, value Type) *Interface {
	return mkInterface(
		mkFunc("Set", mkTuple(key, value), nil),
	)
}

type tcase struct {
	typ        Type
	interfaces []*Interface
}

func mkcase(typ Type, interfaces ...*Interface) tcase {
	return tcase{typ, interfaces}
}

func TestBasicMethodsForGenerics(t *testing.T) {
	if !etoken.GENERICS_V2_CTI {
		t.SkipNow()
		return
	}
	checkImplements := func(typ Type, v *Interface) {
		m, _ := MissingMethod(typ, v, true)
		if m != nil {
			t.Errorf("type %v does not implement %v: missing method %v", typ, v, m)
		}
	}
	checkNotImplements := func(typ Type, v *Interface) {
		m, _ := MissingMethod(typ, v, true)
		if m == nil {
			t.Errorf("type %v implements %v: this should not happen", typ, v)
		}
	}
	caplen := mkInterfaceCapLen()
	getaddr := mkInterfaceGetAddr(Typ[Int], Typ[Uint8])
	getlen := mkInterfaceGetLen(Typ[Int], Typ[Uint8])
	set := mkInterfaceSet(Typ[Int], Typ[Uint8])
	sendrecv := mkInterfaceSendRecv(Typ[Int])
	allifaces := []*Interface{
		caplen, getaddr, getlen, set,
	}
	contains := func(slice []*Interface, key *Interface) bool {
		for _, elem := range slice {
			if elem == key {
				return true
			}
		}
		return false
	}

	tcases := []tcase{
		mkcase(Typ[String], getlen),
		mkcase(NewPointer(NewArray(Typ[Uint8], 0)), caplen, getaddr, getlen, set),
		mkcase(NewSlice(Typ[Uint8]), caplen, getaddr, getlen, set),
		mkcase(NewMap(Typ[Int], Typ[Uint8]), getlen, set),
		mkcase(NewChan(SendRecv, Typ[Int]), caplen, sendrecv),
	}

	for _, c := range tcases {
		t := c.typ
		for _, iface := range allifaces {
			if contains(c.interfaces, iface) {
				checkImplements(t, iface)
			} else {
				checkNotImplements(t, iface)
			}
		}
	}
}
