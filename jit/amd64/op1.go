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

// ============================================================================
// one-arg operation

var op1val = map[Op1]uint8{
	ZERO: 0x08,
	NOT:  0x10,
	NEG:  0x18,
	INC:  0x20,
	DEC:  0x28,
}

func (op Op1) lohi() (uint8, uint8) {
	val, ok := op1val[op]
	if !ok {
		errorf("unknown Op0 instruction: %v", op)
	}
	return val & 0x18, val >> 2
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
		errorf("unknown destination type %T, expecting Reg or Mem: %v %v", a, op, a)
	}
	return asm
}

// OP %reg_dst
func (asm *Asm) op1Reg(op Op1, r Reg) *Asm {
	if op == ZERO {
		return asm.zeroReg(r)
	}
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
	if op == ZERO {
		return asm.zeroMem(m)
	}
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

// zero a register or memory location
func (asm *Asm) Zero(dst Arg) *Asm {
	switch dst := dst.(type) {
	case Reg:
		asm.zeroReg(dst)
	case Mem:
		asm.zeroMem(dst)
	case Const:
		errorf("destination cannot be a constant: %v %v", ZERO, dst)
	default:
		errorf("unknown destination type %T, expecting Reg or Mem: %v %v", dst, ZERO, dst)
	}
	return asm
}

func (asm *Asm) zeroReg(dst Reg) *Asm {
	dlo, dhi := dst.lohi()
	if dhi == 0 {
		return asm.Bytes(0x31, 0xC0|dlo|dlo<<3)
	} else {
		return asm.Bytes(0x48|dhi<<1|dhi<<2, 0x31, 0xC0|dlo|dlo<<3)
	}
}

// zero a memory location
func (asm *Asm) zeroMem(dst Mem) *Asm {
	return asm.movConstMem(Const{val: 0, kind: dst.Kind()}, dst)
}
