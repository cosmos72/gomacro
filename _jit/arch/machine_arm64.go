// +build arm64

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
 * machine_arm64.go
 *
 *  Created on May 26, 2018
 *      Author Massimiliano Ghilardi
 */

package arch

const SUPPORTED = false

type Op0 struct{}
type Op1 struct{}
type Op2 struct{}
type Op3 struct{}
type Op4 struct{}

func (asm *Asm) Op0(op Op0) *Asm {
	return asm
}

func (asm *Asm) Op1(op Op1, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Op2(op Op2, src Arg, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Op3(op Op3, a Arg, b Arg, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Op4(op Op4, a Arg, b Arg, c Arg, dst Arg) *Asm {
	return asm
}

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
	X31
	RLo = X0
	RHi = X31
)

var alwaysLiveRegs = RegIds{
	X28: 1, // pointer to goroutine-local data
	X29: 1, // jit *uint64 pointer-to-variables
	X30: 1, // link register?
	X31: 1, // instruction pointer? return address?
}

func (r Reg) Valid() bool {
	return r >= rLo && r <= rHi
}

func (r Reg) Validate() {
	if !r.Valid() {
		errorf("invalid register: %d", r)
	}
}

func (r Reg) lo() uint32 {
	r.Validate()
	return uint32(r) - 1
}

func (asm *Asm) Prologue() *Asm {
	return asm.Uint32(0xf94007fd) // ldr x29, [sp, #8]
}

func (asm *Asm) Epilogue() *Asm {
	return asm.Uint32(0xd65f03c0) // ret
}

var assertError = errors.New("jit/arm64 internal error, assertion failed")

func assert(flag bool) {
	if !flag {
		panic(assertError)
	}
}
