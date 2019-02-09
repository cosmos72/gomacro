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
 * api.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

type Size uint8 // 1, 2, 4 or 8

type Arg interface {
	RegId() RegId // register used by Arg, or NoReg if Arg is Const
	Kind() Kind
	Const() bool
}

type Code []uint8

type SaveSlot uint16

const (
	InvalidSlot = ^SaveSlot(0)
)

// memory area where spill registers can be saved
type Save struct {
	reg              Reg       // points to memory area
	start, next, end SaveSlot // memory area indexes
	bitmap           []bool    // bitmap of used/free indexes
}

type Asm struct {
	code        Code
	nextRegId   RegId   // first available register
	nextSoftReg SoftReg // first available soft register
	softRegs    SoftRegs
	save        Save
	regIds      RegIds
}

func New() *Asm {
	var asm Asm
	return asm.Init()
}

func SizeOf(a Arg) Size {
	size := a.Kind().Size()
	if size == 0 {
		errorf("unsupported register/memory kind: %v", a.Kind())
	}
	return size
}
