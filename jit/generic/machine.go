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
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

const ASM_SUPPORTED = false
const Name = "generic"

const (
	NoRegId = RegId(0)
	RLo     = NoRegId
	RHi     = NoRegId
	RSP     = NoRegId
	RVAR    = NoRegId
)

func (r RegId) Valid() bool {
	return false
}

var alwaysLiveRegIds RegIds // empty

type Op0 uint8
type Op1 uint8
type Op2 uint8
type Op3 uint8
type Op4 uint8

const (
	// Op0
	RET Op0 = 0x00
	NOP Op0 = 0x01

	// Op1
	ZERO Op1 = 0x10
	INC  Op1 = 0x11
	DEC  Op1 = 0x12
	NEG  Op1 = 0x13
	NOT  Op1 = 0x14

	// Op2
	ADD  Op2 = 0x20
	AND  Op2 = 0x21
	ADC  Op2 = 0x22
	OR   Op2 = 0x23
	XOR  Op2 = 0x24
	SUB  Op2 = 0x25
	SBB  Op2 = 0x26
	SHL  Op2 = 0x27
	SHR  Op2 = 0x28
	MUL  Op2 = 0x29
	DIV  Op2 = 0x2A
	REM  Op2 = 0x2B
	MOV  Op2 = 0x2C
	CAST Op2 = 0x2D
	NEG2     = Op2(NEG)
	NOT2     = Op2(NOT)
	/*
		CMP  = arch.CMP
		XCHG = arch.XCHG
	*/

	// Op3
	ADD3 = Op3(ADD)
	OR3  = Op3(OR)
	ADC3 = Op3(ADC)
	SBB3 = Op3(SBB)
	AND3 = Op3(AND)
	SUB3 = Op3(SUB)
	XOR3 = Op3(XOR)
	SHL3 = Op3(SHL)
	SHR3 = Op3(SHR)
	MUL3 = Op3(MUL)
	DIV3 = Op3(DIV)
	REM3 = Op3(REM)

	// Op4
)

func (op Op0) String() string {
	return fmt.Sprintf("Op0(%d)", uint8(op))
}

func (op Op1) String() string {
	return fmt.Sprintf("Op1(%d)", uint8(op))
}

func (op Op2) String() string {
	return fmt.Sprintf("Op2(%d)", uint8(op))
}

func (op Op3) String() string {
	return fmt.Sprintf("Op3(%d)", uint8(op))
}

func (op Op4) String() string {
	return fmt.Sprintf("Op4(%d)", uint8(op))
}

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

func (asm *Asm) Zero(dst Arg) *Asm {
	return asm
}

func (asm *Asm) Mov(src Arg, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Load(src Mem, dst Reg) *Asm {
	return asm
}

func (asm *Asm) Store(src Reg, dst Mem) *Asm {
	return asm
}

func (asm *Asm) Prologue() *Asm {
	return asm
}

func (asm *Asm) Epilogue() *Asm {
	return asm
}

func (s *Save) ArchInit(start SaveSlot, end SaveSlot) {
}
