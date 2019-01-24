/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * kind.go
 *
 *  Created on Jan 24, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"reflect"
)

type Kind uint8 // narrow version of reflect.Kind

const (
	KBool    = Kind(reflect.Bool)
	KInt     = Kind(reflect.Int)
	KInt8    = Kind(reflect.Int8)
	KInt16   = Kind(reflect.Int16)
	KInt32   = Kind(reflect.Int32)
	KInt64   = Kind(reflect.Int64)
	KUint    = Kind(reflect.Uint)
	KUint8   = Kind(reflect.Uint8)
	KUint16  = Kind(reflect.Uint16)
	KUint32  = Kind(reflect.Uint32)
	KUint64  = Kind(reflect.Uint64)
	KUintptr = Kind(reflect.Uintptr)
	KFloat32 = Kind(reflect.Float32)
	KFloat64 = Kind(reflect.Float64)
	KPtr     = Kind(reflect.Ptr)
	KLo      = KBool
	KHi      = KPtr
)

var ksize = [...]Size{
	KBool:    1,
	KInt8:    1,
	KInt16:   2,
	KInt32:   4,
	KInt64:   8,
	KUint8:   1,
	KUint16:  2,
	KUint32:  4,
	KUint64:  8,
	KFloat32: 4,
	KFloat64: 8,
	KInt:     Size(reflect.TypeOf(int(0)).Size()),
	KUint:    Size(reflect.TypeOf(uint(0)).Size()),
	KUintptr: Size(reflect.TypeOf(uintptr(0)).Size()),
	KPtr:     Size(reflect.TypeOf((*int)(nil)).Size()),
}

func (k Kind) Size() Size {
	if k >= KLo && k <= KHi {
		return ksize[k]
	}
	return 0
}

func (k Kind) String() string {
	return reflect.Kind(k).String()
}
