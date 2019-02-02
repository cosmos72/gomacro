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
 * op3.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// ============================================================================
// ternary operation
type Op3 uint8

const (
	AND3 Op3 = 0x0A
	ADD3 Op3 = 0x0B
	ADC3 Op3 = 0x1A // add with carry
	MUL3 Op3 = 0x1B
	SHL3 Op3 = 0x1C // shift left
	SHR3 Op3 = 0x1D // shift right
	OR3  Op3 = 0x2A
	XOR3 Op3 = 0x4A
	SUB3 Op3 = 0x4B
	SBB3 Op3 = 0x5A // subtract with borrow

	// FIXME replace with correct values
	DIV3 Op3 = 0xFE
	REM3 Op3 = 0xFF
)

var op3Name = map[Op3]string{
	ADD3: "ADD3",
	OR3:  "OR3",
	ADC3: "ADC3",
	SBB3: "SBB3",
	AND3: "AND3",
	SUB3: "SUB3",
	XOR3: "XOR3",
	SHL3: "SHL3",
	SHR3: "SHR3",
	MUL3: "MUL3",
	DIV3: "DIV3",
	REM3: "REM3",
}

func (op Op3) String() string {
	s, ok := op3Name[op]
	if !ok {
		s = fmt.Sprintf("Op3(%d)", int(op))
	}
	return s
}

func (op Op3) val() uint32 {
	switch op {
	case SHL3:
		return 0x1AC02000
	case SHR3:
		// logical i.e. zero-extended right shift is 0x1AC02400
		// arithmetic i.e. sign-extended right shift is 0x1AC02800
		return 0x1AC02400
	case MUL3:
		// 0x1B007C00 because MUL3 a,b,c is an alias for MADD4 xzr,a,b,c
		return 0x1B007C00
	default:
		return uint32(op) << 24
	}
}

// ============================================================================
func (asm *Asm) Op3(op Op3, a Arg, b Arg, dst Arg) *Asm {
	assert(a.Kind() == dst.Kind())
	if op != SHL3 && op != SHR3 {
		assert(b.Kind() == dst.Kind())
	}
	if asm.optimize(op, a, b, dst) {
		return asm
	}
	ra, raok := a.(Reg)
	rb, rbok := b.(Reg)
	rdst, rdok := dst.(Reg)
	if !raok {
		errorf("unimplemented source type %T, expecting Reg: %v %v %v %v", a, op, a, b, dst)
	}
	if !rbok {
		errorf("unimplemented source type %T, expecting Reg: %v %v %v %v", b, op, a, b, dst)
	}
	if _, ok := dst.(Const); ok {
		errorf("destination cannot be a constant: %v %v %v %v", dst, op, a, b, dst)
	}
	if !rdok {
		errorf("unimplemented destination type %T, expecting Reg: %v %v %v %v", dst, op, a, b, dst)
	}
	return asm.op3RegRegReg(op, ra, rb, rdst)
}

func (asm *Asm) op3RegRegReg(op Op3, a Reg, b Reg, dst Reg) *Asm {
	var kbit, signedshr uint32
	if op == SHR3 && dst.kind.Signed() {
		signedshr = 0xC00
	}
	switch dst.kind.Size() {
	case 1, 2, 4:
		// TODO mask result for size 1, 2
		kbit = 0x80 << 24
		fallthrough
	case 8:
		asm.Uint32(kbit | (signedshr ^ op.val()) | b.val()<<16 | a.val()<<5 | dst.val())
	}
	return asm
}

func (op Op3) isCommutative() bool {
	switch op {
	case ADD3, OR3, ADC3, AND3, XOR3, MUL3:
		return true
	}
	return false
}

func (asm *Asm) optimize(op Op3, a Arg, b Arg, dst Arg) bool {
	// TODO
	return false
}
