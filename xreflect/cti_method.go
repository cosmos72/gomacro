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
 * cti_methods.go
 *
 *  Created on May 12, 2019
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	r "reflect"

	"github.com/cosmos72/gomacro/go/etoken"
)

// declare CTI methods on Array, Chan, Map, Slice

func (v *Universe) addTypeMethodsCTI(xt *xtype) {
	if !etoken.GENERICS_V2_CTI {
		return
	}
	k := xt.kind
	switch k {
	case r.Func, r.Interface, r.Ptr, r.Struct, r.UnsafePointer:
		return
	}
	rt := xt.rtype
	rbool := rbasictypes[r.Bool]
	rint := rbasictypes[r.Int]
	rkey := rint
	var relem r.Type

	if k == r.Map {
		rkey = rt.Key()
	}
	if k == r.Array || k == r.Chan || k == r.Map || k == r.Slice {
		relem = rt.Elem()
	} else if k == r.String {
		relem = rbasictypes[r.Uint8]
	}
	vt := []r.Type{rt}
	vtkey := []r.Type{rt, rkey}
	vint := []r.Type{rint}
	velem := []r.Type{relem}

	n := xt.NumExplicitMethod()
	m := xt.methodvalues
	if len(m) < n {
		xt.methodvalues = make([]r.Value, n)
		copy(xt.methodvalues, m)
		m = xt.methodvalues
	}
	for i := 0; i < n; i++ {
		switch xt.method(i).Name {

		// array, slice, string methods
		case "Append":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, r.SliceOf(relem)}, vt, true), ctiAppend)
		case "AppendString":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, rbasictypes[r.String]}, vt, false), ctiAppendString)
		case "Copy":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, rt}, nil, false), ctiCopy)
		case "CopyString":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, rbasictypes[r.String]}, nil, false), ctiCopy)
		case "Cap":
			m[i] = r.MakeFunc(r.FuncOf(vt, vint, false), ctiCap)
		case "Len":
			m[i] = r.MakeFunc(r.FuncOf(vt, vint, false), ctiLen)
		case "Slice":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, rkey, rkey}, velem, false), ctiSlice)
		case "Slice3":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, rkey, rkey, rkey}, velem, false), ctiSlice3)

			// indexing
		case "AddrIndex":
			sig := r.FuncOf(vtkey, []r.Type{r.PtrTo(relem)}, false)
			m[i] = r.MakeFunc(sig, ctiAddrIndex)
		case "DelIndex":
			m[i] = r.MakeFunc(r.FuncOf(vtkey, nil, false), ctiDelMapIndex)
		case "Index":
			sig := r.FuncOf(vtkey, []r.Type{relem}, false)
			if k == r.Map {
				zero := r.Zero(relem)
				m[i] = r.MakeFunc(sig,
					func(v []r.Value) []r.Value {
						ret := v[0].MapIndex(v[1])
						if !ret.IsValid() {
							ret = zero
						}
						return []r.Value{ret}
					})
			} else {
				m[i] = r.MakeFunc(sig, ctiIndex)
			}
		case "SetIndex":
			sig := r.FuncOf([]r.Type{rt, rkey, relem}, nil, false)
			if k == r.Map {
				m[i] = r.MakeFunc(sig, ctiSetMapIndex)
			} else {
				m[i] = r.MakeFunc(sig, ctiSetIndex)
			}
		case "TryIndex":
			sig := r.FuncOf(vtkey, []r.Type{relem, rbool}, false)

			zero := r.Zero(relem)
			vtrue := r.ValueOf(true)
			vfalse := r.ValueOf(false)
			m[i] = r.MakeFunc(sig,
				func(v []r.Value) []r.Value {
					elem := v[0].MapIndex(v[1])
					ok := vtrue
					if !elem.IsValid() {
						elem = zero
						ok = vfalse
					}
					return []r.Value{elem, ok}
				})

			// chan methods
		case "Recv":
			m[i] = r.MakeFunc(r.FuncOf(vt, []r.Type{relem, rbool}, false), ctiRecv)
		case "Send":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, relem}, nil, false), ctiSend)
		case "TryRecv":
			m[i] = r.MakeFunc(r.FuncOf(vt, []r.Type{relem, rbool}, false), ctiTryRecv)
		case "TrySend":
			m[i] = r.MakeFunc(r.FuncOf([]r.Type{rt, relem}, []r.Type{rbool}, false), ctiTrySend)
		case "Close":
			m[i] = r.MakeFunc(r.FuncOf(vt, nil, false), ctiClose)
		}
	}
}

// array, slice, string methods

func ctiAppend(v []r.Value) []r.Value {
	return []r.Value{
		r.AppendSlice(v[0], v[1]),
	}
}

var rTypeOfByteSlice = r.SliceOf(rbasictypes[r.Uint8])

func ctiAppendString(v []r.Value) []r.Value {
	vslice := v[0]
	t := vslice.Type()
	if t != rTypeOfByteSlice {
		vslice = vslice.Convert(rTypeOfByteSlice)
	}
	slice := vslice.Interface().([]byte)
	slice = append(slice, v[1].String()...)
	vslice = r.ValueOf(slice)
	if t != rTypeOfByteSlice {
		vslice = vslice.Convert(t)
	}
	return []r.Value{vslice}
}

func ctiCap(v []r.Value) []r.Value {
	return []r.Value{r.ValueOf(
		r.Indirect(v[0]).Cap(),
	)}
}

func ctiCopy(v []r.Value) []r.Value {
	r.Copy(r.Indirect(v[0]), v[1])
	return nil
}

func ctiLen(v []r.Value) []r.Value {
	return []r.Value{r.ValueOf(
		r.Indirect(v[0]).Len(),
	)}
}

func ctiSlice(v []r.Value) []r.Value {
	return []r.Value{
		r.Indirect(v[0]).Slice(int(v[1].Int()), int(v[2].Int())),
	}
}

func ctiSlice3(v []r.Value) []r.Value {
	return []r.Value{
		r.Indirect(v[0]).Slice3(int(v[1].Int()), int(v[2].Int()), int(v[3].Int())),
	}
}

// indexing

func ctiAddrIndex(v []r.Value) []r.Value {
	return []r.Value{
		r.Indirect(v[0]).Index(int(v[1].Int())).Addr(),
	}
}

func ctiDelMapIndex(v []r.Value) []r.Value {
	v[0].SetMapIndex(v[1], r.Value{})
	return nil
}

func ctiIndex(v []r.Value) []r.Value {
	return []r.Value{
		r.Indirect(v[0]).Index(int(v[1].Int())),
	}
}

func ctiSetMapIndex(v []r.Value) []r.Value {
	v[0].SetMapIndex(v[1], v[2])
	return nil
}

func ctiSetIndex(v []r.Value) []r.Value {
	r.Indirect(v[0]).Index(int(v[1].Int())).Set(v[2])
	return nil
}

// chan methods

func ctiRecv(v []r.Value) []r.Value {
	ret, ok := v[0].Recv()
	return []r.Value{
		ret, r.ValueOf(ok),
	}
}

func ctiSend(v []r.Value) []r.Value {
	v[0].Send(v[1])
	return nil
}

func ctiTryRecv(v []r.Value) []r.Value {
	ret, ok := v[0].TryRecv()
	return []r.Value{
		ret, r.ValueOf(ok),
	}
}

func ctiTrySend(v []r.Value) []r.Value {
	return []r.Value{r.ValueOf(
		v[0].TrySend(v[1]),
	)}
}

func ctiClose(v []r.Value) []r.Value {
	v[0].Close()
	return nil
}
