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
 * machine.go
 *
 *  Created on May 26, 2018
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"errors"
	"fmt"
)

const SUPPORTED = false // not working yet
const Name = "arm64"

const (
	NoRegId RegId = iota
	X0
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	X29
	X30
	XZR, XSP = iota, iota // depending on context, zero register or stack pointer
	RLo      = X0
	RHi      = XZR
)

var (
	regName4 = makeRegNames("w")
	regName8 = makeRegNames("x")
)

var alwaysLiveRegIds = RegIds{
	X28: 1, // pointer to goroutine-local data
	X29: 1, // jit *uint64 pointer-to-variables
	X30: 1, // return address register
	XZR: 1, // zero register / stack pointer
}

func (id RegId) Valid() bool {
	return id >= RLo && id < RHi // XZR is valid only in few, hand-checked cases
}

func makeRegNames(prefix string) (ret [RHi + 1]string) {
	for id := RLo; id < RHi; id++ {
		ret[id] = fmt.Sprint(prefix, int(id)-1)
	}
	ret[RHi] = prefix + "zr"
	return ret
}

func (id RegId) String() string {
	var s string
	if id >= RLo && id <= RHi {
		s = regName8[id]
	}
	if s == "" {
		s = fmt.Sprintf("unknown_reg(%d)", uint8(id))
	}
	return s
}

func (r Reg) String() string {
	var s string
	id := r.id
	if id >= RLo && id <= RHi {
		switch r.kind.Size() {
		case 1, 2, 4:
			s = regName4[id]
		case 8:
			s = regName8[id]
		}
	}
	if s == "" {
		s = fmt.Sprintf("unknown_reg(%d,%v)", uint8(id), r.kind)
	}
	return s
}

func (r Reg) val() uint32 {
	r.Validate()
	return uint32(r.id) - 1
}

func (asm *Asm) Prologue() *Asm {
	return asm.Uint32(0xf94007fd) // ldr x29, [sp, #8]
}

func (asm *Asm) Epilogue() *Asm {
	return asm.Uint32(0xd65f03c0) // ret
}

func (asm *Asm) archPush(id RegId) {
	errorf("archPush not implemented")
}

func (asm *Asm) archPop(id RegId) {
	errorf("archPush not implemented")
}

func (s *Save) ArchInit(start uint16, end uint16) {
}

var assertError = errors.New("jit/arm64 internal error, assertion failed")

func assert(flag bool) {
	if !flag {
		panic(assertError)
	}
}
