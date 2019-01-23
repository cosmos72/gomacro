// +build amd64

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
 * machine_amd64.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"errors"
	"reflect"
)

const SUPPORTED = true

type TernaryOp struct{} // not used on amd64

type Op uint8

const (
	ADD Op = 0
	OR  Op = 0x08
	// ADC Op = 0x10 // add with carry
	// SBB Op = 0x18 // subtract with borrow
	AND Op = 0x20
	SUB Op = 0x28
	XOR Op = 0x30
	// CMP Op = 0x38 // compare, set flags
	// XCHG Op = 0x86 // exchange. xchg %reg, %reg has different encoding
	MOV Op = 0x88
)

type UnaryOp uint8

const (
	NOT UnaryOp = 0x10
	NEG UnaryOp = 0x18
)

const (
	NoReg Reg = iota
	RAX
	RCX
	RDX
	RBX
	RSP
	RBP
	RSI
	RDI
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
	RLo Reg = RAX
	RHi Reg = R15
)

var alwaysLiveRegs = Regs{RSP: 1, RBP: 1, RDI: 1 /* &Env.IntBinds[0] */}

var regName = [...]string{
	NoReg: "unknown register",
	RAX:   "%rax",
	RCX:   "%rcx",
	RDX:   "%rdx",
	RBX:   "%rbx",
	RSP:   "%rsp",
	RBP:   "%rbp",
	RSI:   "%rsi",
	RDI:   "%rdi",
	R8:    "%r8",
	R9:    "%r9",
	R10:   "%r10",
	R11:   "%r11",
	R12:   "%r12",
	R13:   "%r13",
	R14:   "%r14",
	R15:   "%r15",
}

func (r Reg) Valid() bool {
	return r >= RLo && r <= RHi
}

func (r Reg) Validate() {
	if !r.Valid() {
		giveupf("invalid register: %d", r)
	}
}

func (r Reg) String() string {
	if !r.Valid() {
		r = NoReg
	}
	return regName[r]
}

// implement Arg interface
func (r Reg) Reg() Reg {
	return r
}

func (r Reg) Kind() reflect.Kind {
	// update after implementing floating-point registers
	return reflect.Int64
}

func (r Reg) Const() bool {
	return false
}

func (r Reg) bits() uint8 {
	r.Validate()
	return uint8(r) - 1
}

func (r Reg) lo() uint8 {
	return r.bits() & 0x7
}

func (r Reg) hi() uint8 {
	return (r.bits() & 0x8) >> 3
}

func (r Reg) lohi() (uint8, uint8) {
	bits := r.bits()
	return bits & 0x7, (bits & 0x8) >> 3
}

func (asm *Asm) Prologue() *Asm {
	return asm.Bytes(0x48, 0x8b, 0x7c, 0x24, 0x08) // movq 0x8(%rsp), %rdi
}

func (asm *Asm) Epilogue() *Asm {
	return asm.Bytes(0xc3) // ret
}

var assertError = errors.New("jit/amd64 internal error, assertion failed")

func assert(flag bool) {
	if !flag {
		panic(assertError)
	}
}
