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
 * op1.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

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
		s = fmt.Sprintf("Op1(%d)", int(op))
	}
	return s
}

func (op Op1) lohi() (uint8, uint8) {
	return uint8(op & 0x18), uint8(op >> 2)
}

// ============================================================================
func (asm *Asm) Op1(op Op1, a Arg) *Asm {
	switch a := a.(type) {
	case Reg:
		asm.op1Reg(op, a)
	case Mem:
		asm.op1Mem(op, a)
	case Const:
		errorf("destination cannot be a constant: %v %v", op, a)
	default:
		errorf("unsupported destination type %T, expecting Reg or Mem: %v %v", a, op, a)
	}
	return asm
}

// OP %reg_dst
func (asm *Asm) op1Reg(op Op1, r Reg) *Asm {
	rlo, rhi := r.lohi()
	oplo, ophi := op.lohi()

	switch SizeOf(r) {
	case 1:
		if r.id >= RSP {
			asm.Byte(0x40 | rhi)
		}
		asm.Bytes(0xF6|ophi, 0xC0|oplo|rlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if rhi != 0 {
			asm.Byte(0x41)
		}
		asm.Bytes(0xF7|ophi, 0xC0|oplo|rlo)
	case 8:
		asm.Bytes(0x48|rhi, 0xF7|ophi, 0xC0|oplo|rlo)
	default:
		errorf("unsupported register size %v, expecting 1,2,4 or 8 bytes: %v %v", SizeOf(r), op, r)
	}
	return asm
}

// OP off_m(%reg_m)
func (asm *Asm) op1Mem(op Op1, m Mem) *Asm {

	r := m.reg
	rlo, dhi := r.lohi()
	oplo, ophi := op.lohi()

	offlen, offbit := m.offlen(r.id)

	switch SizeOf(m) {
	case 1:
		if dhi != 0 {
			asm.Byte(0x41)
		}
		asm.Bytes(0xF6|ophi, offbit|oplo|rlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Byte(0x41)
		}
		asm.Bytes(0xF7|ophi, offbit|oplo|rlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xF7|ophi, offbit|oplo|rlo)
	default:
		errorf("unsupported memory size %v, expecting 1,2,4 or 8 bytes: %v %v", SizeOf(m), op, m)
	}
	asm.quirk24(r)
	switch offlen {
	case 1:
		asm.Int8(int8(m.off))
	case 4:
		asm.Int32(m.off)
	}
	return asm
}
