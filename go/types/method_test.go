// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"go/token"
	"testing"

	"github.com/cosmos72/gomacro/go/etoken"
)

func mkvar(t Type) *Var {
	return NewVar(token.NoPos, nil, "", t)
}

func mktuple(ts ...Type) *Tuple {
	vs := make([]*Var, len(ts))
	for i := range ts {
		vs[i] = mkvar(ts[i])
	}
	return NewTuple(vs...)
}

func mkfunc(name string, params *Tuple, results *Tuple) *Func {
	return NewFunc(token.NoPos, nil, name, NewSignature(nil, params, results, false))
}

func mkinterface(fs ...*Func) *Interface {
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
	return mkinterface(
		mkfunc("Cap", nil, mktuple(Typ[Int])),
		mkfunc("Len", nil, mktuple(Typ[Int])),
	)
}

/**
 * return
 * interface {
 *   GetAddr(k Key) *Value
 * }
 */
func mkInterfaceGetAddr(key, value Type) *Interface {
	return mkinterface(
		mkfunc("GetAddr", mktuple(key), mktuple(NewPointer(value))),
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
	return mkinterface(
		mkfunc("Get", mktuple(key), mktuple(value)),
		mkfunc("Len", nil, mktuple(Typ[Int])),
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
	return mkinterface(
		mkfunc("Send", mktuple(elem), nil),
		mkfunc("Recv", mktuple(elem, Typ[Bool]), nil),
	)
}

/**
 * return
 * interface {
 *   Set(k Key, v Value)
 * }
 */
func mkInterfaceSet(key, value Type) *Interface {
	return mkinterface(
		mkfunc("Set", mktuple(key, value), nil),
	)
}

func mkNamed(name string, underlying Type) *Named {
	return NewNamed(NewTypeName(token.NoPos, nil, name, nil), underlying, nil)
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

	tarray := NewArray(Typ[Uint8], 0)
	tchan := NewChan(SendRecv, Typ[Int])
	tmap := NewMap(Typ[Int], Typ[Uint8])
	tslice := NewSlice(Typ[Uint8])
	tstring := Typ[String]

	tchannamed := mkNamed("ChanInt", tchan)
	tmapnamed := mkNamed("MapIntUint8", tmap)
	tslicenamed := mkNamed("SliceUint8", tslice)
	tstringnamed := mkNamed("String", tstring)

	tcases := []tcase{
		mkcase(NewPointer(tarray), caplen, getaddr, getlen, set),
		mkcase(tchan, caplen, sendrecv),
		mkcase(tchannamed, caplen, sendrecv),
		mkcase(tmap, getlen, set),
		mkcase(tmapnamed, getlen, set),
		mkcase(tslice, caplen, getaddr, getlen, set),
		mkcase(tslicenamed, caplen, getaddr, getlen, set),
		mkcase(tstring, getlen),
		mkcase(tstringnamed, getlen),
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
