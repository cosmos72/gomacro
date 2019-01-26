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
)

const SUPPORTED = true

// ============================================================================
// unary operation
type Op1 uint8

const (
	NOT Op1 = 0x10
	NEG Op1 = 0x18
	INC Op1 = 0x20
	DEC Op1 = 0x28
)

var op1Name = map[Op1]string{
	NOT: "NOT",
	NEG: "NEG",
	INC: "INC",
	DEC: "DEC",
}

func (op Op1) String() string {
	s, ok := op1Name[op]
	if !ok {
		s = "unknown unary operation"
	}
	return s
}

func (op Op1) lohi() (uint8, uint8) {
	return uint8(op & 0x18), uint8(op >> 2)
}

// ============================================================================
// binary operation
type Op2 uint8

const (
	ADD Op2 = 0
	OR  Op2 = 0x08
	ADC Op2 = 0x10 // add with carry
	SBB Op2 = 0x18 // subtract with borrow
	AND Op2 = 0x20
	SUB Op2 = 0x28
	XOR Op2 = 0x30
	// CMP Op = 0x38 // compare, set flags
	// XCHG Op = 0x86 // exchange. xchg %reg, %reg has different encoding
	MOV  Op2 = 0x88
	CAST Op2 = 0xB6 // sign extend, zero extend or narrow
)

var op2Name = map[Op2]string{
	ADD: "ADD",
	OR:  "OR",
	ADC: "ADC",
	SBB: "SBB",
	AND: "AND",
	SUB: "SUB",
	XOR: "XOR",
	// CMP: "CMP",
	// XCHG: "XCHG",
	MOV:  "MOV",
	CAST: "CAST",
}

func (op Op2) String() string {
	s, ok := op2Name[op]
	if !ok {
		s = "unknown binary operation"
	}
	return s
}

// ============================================================================
// register
const (
	NoRegId RegId = iota
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
	RLo RegId = RAX
	RHi RegId = R15
)

var alwaysLiveRegIds = RegIds{RSP: 1, RBP: 1, RDI: 1 /* &Env.IntBinds[0] */}

var regName1 = [...]string{
	NoRegId: "unknown 1-byte register",
	RAX:     "%al",
	RCX:     "%cl",
	RDX:     "%dl",
	RBX:     "%bl",
	RSP:     "%spl",
	RBP:     "%bpl",
	RSI:     "%sil",
	RDI:     "%dil",
	R8:      "%r8b",
	R9:      "%r9b",
	R10:     "%r10b",
	R11:     "%r11b",
	R12:     "%r12b",
	R13:     "%r13b",
	R14:     "%r14b",
	R15:     "%r15b",
}
var regName2 = [...]string{
	NoRegId: "unknown 2-byte register",
	RAX:     "%ax",
	RCX:     "%cx",
	RDX:     "%dx",
	RBX:     "%bx",
	RSP:     "%sp",
	RBP:     "%bp",
	RSI:     "%si",
	RDI:     "%di",
	R8:      "%r8w",
	R9:      "%r9w",
	R10:     "%r10w",
	R11:     "%r11w",
	R12:     "%r12w",
	R13:     "%r13w",
	R14:     "%r14w",
	R15:     "%r15w",
}
var regName4 = [...]string{
	NoRegId: "unknown 4-byte register",
	RAX:     "%eax",
	RCX:     "%ecx",
	RDX:     "%edx",
	RBX:     "%ebx",
	RSP:     "%esp",
	RBP:     "%ebp",
	RSI:     "%esi",
	RDI:     "%edi",
	R8:      "%r8d",
	R9:      "%r9d",
	R10:     "%r10d",
	R11:     "%r11d",
	R12:     "%r12d",
	R13:     "%r13d",
	R14:     "%r14d",
	R15:     "%r15d",
}
var regName8 = [...]string{
	NoRegId: "unknown register",
	RAX:     "%rax",
	RCX:     "%rcx",
	RDX:     "%rdx",
	RBX:     "%rbx",
	RSP:     "%rsp",
	RBP:     "%rbp",
	RSI:     "%rsi",
	RDI:     "%rdi",
	R8:      "%r8",
	R9:      "%r9",
	R10:     "%r10",
	R11:     "%r11",
	R12:     "%r12",
	R13:     "%r13",
	R14:     "%r14",
	R15:     "%r15",
}

func (id RegId) Valid() bool {
	return id >= RLo && id <= RHi
}

func (id RegId) Validate() {
	if !id.Valid() {
		errorf("invalid register: %v", id)
	}
}

func (id RegId) String() string {
	if !id.Valid() {
		id = NoRegId
	}
	return regName8[id]
}

func (r Reg) Valid() bool {
	return r.id.Valid()
}

func (r Reg) Validate() {
	r.id.Validate()
}

func (r Reg) String() string {
	id := NoRegId
	if r.Valid() {
		id = r.id
	}
	switch r.kind.Size() {
	case 1:
		return regName1[id]
	case 2:
		return regName2[id]
	case 4:
		return regName4[id]
	default:
		return regName8[id]
	}
}

func (r Reg) bits() uint8 {
	r.Validate()
	return uint8(r.id) - 1
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

// return number of assembler bytes needed to encode m.off
func (m Mem) offlen(id RegId) (uint8, uint8) {
	switch {
	// (%rbp) and (%r13) registers must use 1-byte offset even if m.off == 0
	case m.off == 0 && id != RBP && id != R13:
		return 0, 0
	case m.off == int32(int8(m.off)):
		return 1, 0x40
	default:
		return 4, 0x80
	}
}

func (asm *Asm) quirk24(r Reg) *Asm {
	if r.id == RSP || r.id == R12 {
		asm.Bytes(0x24) // amd64 quirk
	}
	return asm
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
