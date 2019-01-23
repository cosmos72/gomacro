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
 * mem.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
	"reflect"
)

func (m Mem) String() string {
	return fmt.Sprintf("[%v+%v]/*%v*/", m.reg, m.off, m.kind)
}

// implement Arg interface
func (m Mem) Reg() Reg {
	return NoReg
}

func (m Mem) Kind() reflect.Kind {
	return m.kind
}

func (m Mem) Const() bool {
	return false
}

func MakeVar(index uint16) Mem {
	return Mem{off: int32(index) * 8, kind: reflect.Int64, reg: RDI}
}

var sizeof = [...]Size{
	reflect.Bool:          1,
	reflect.Int8:          1,
	reflect.Int16:         2,
	reflect.Int32:         4,
	reflect.Int64:         8,
	reflect.Uint8:         1,
	reflect.Uint16:        2,
	reflect.Uint32:        4,
	reflect.Uint64:        8,
	reflect.Float32:       4,
	reflect.Float64:       8,
	reflect.Complex64:     8,
	reflect.Complex128:    16,
	reflect.Int:           Size(reflect.TypeOf(int(0)).Size()),
	reflect.Uint:          Size(reflect.TypeOf(uint(0)).Size()),
	reflect.Uintptr:       Size(reflect.TypeOf(uintptr(0)).Size()),
	reflect.Ptr:           Size(reflect.TypeOf((*int)(nil)).Size()),
	reflect.UnsafePointer: 0, // highest numbered reflect.Kind
}

func SizeOf(a Arg) Size {
	k := a.Kind()
	size := sizeof[k]
	if size == 0 {
		giveupf("unsupported reflect.Kind %v", k)
	}
	return size
}
