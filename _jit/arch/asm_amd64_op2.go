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
 * asm_amd64_op2.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

func (asm *Asm) Op2(op Op2, dst Arg, src Arg) *Asm {
	if op == CAST {
		if dst.Kind() != src.Kind() {
			return asm.Cast(dst, src)
		}
		op = MOV
	}
	assert(dst.Kind() == src.Kind())
	if isNop2(op, dst, src) {
		return asm
	}
	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Reg:
			asm.op2RegReg(op, dst, src)
		case Mem:
			asm.op2RegMem(op, dst, src)
		case Const:
			asm.op2RegConst(op, dst, src)
		default:
			errorf("unsupported source type, expecting Reg, Mem or Const: %v %v %v", op, dst, src)
		}
	case Mem:
		switch src := src.(type) {
		case Reg:
			asm.op2MemReg(op, dst, src)
		case Mem:
			asm.op2MemMem(op, dst, src)
		case Const:
			asm.op2MemConst(op, dst, src)
		default:
			errorf("unsupported source type, expecting Reg, Mem or Const: %v %v %v", op, dst, src)
		}
	case Const:
		errorf("destination cannot be a constant: %v %v %v", op, dst, src)
	default:
		errorf("unsupported destination type, expecting Reg or Mem: %v %v %v", op, dst, src)
	}
	return asm
}

// %reg_dst OP= const
func (asm *Asm) op2RegConst(op Op2, dst Reg, src Const) *Asm {
	if op == MOV {
		return asm.movRegConst(dst, src)
	}
	dlo, dhi := dst.lohi()
	op_ := uint8(op)
	c := src.val
	if c == int64(int8(c)) {
		asm.Bytes(0x48|dhi, 0x83, 0xC0|op_|dlo, uint8(int8(c)))
	} else if c == int64(int32(c)) {
		if dst.id == RAX {
			asm.Bytes(0x48|dhi, 0x05|op_).Int32(int32(c))
		} else {
			asm.Bytes(0x48|dhi, 0x81, 0xC0|op_|dlo).Int32(int32(c))
		}
	} else {
		// constant is 64 bit wide, must load it in a register
		r := asm.RegAlloc(src.kind)
		asm.movRegConst(r, src)
		asm.op2RegReg(op, dst, r)
		asm.RegFree(r)
	}
	return asm
}

// %reg_dst OP= %reg_src
func (asm *Asm) op2RegReg(op Op2, dst Reg, src Reg) *Asm {
	if isNop2(op, dst, src) {
		return asm
	}
	dlo, dhi := dst.lohi()
	slo, shi := src.lohi()

	switch SizeOf(dst) { // == SizeOf(src)
	case 1:
		if dst.id < RSP && src.id < RSP {
			asm.Bytes(uint8(op), 0xC0|dlo|slo<<3)
		} else {
			asm.Bytes(0x40|dhi|shi<<2, uint8(op), 0xC0|dlo|slo<<3)
		}
	case 2:
		if dhi|shi<<2 == 0 {
			asm.Bytes(0x66, 0x01|uint8(op), 0xC0|dlo|slo<<3)
		} else {
			asm.Bytes(0x66, 0x40|dhi|shi<<2, 0x01|uint8(op), 0xC0|dlo|slo<<3)
		}
	case 4:
		if dhi|shi<<2 == 0 {
			asm.Bytes(0x01|uint8(op), 0xC0|dlo|slo<<3)
		} else {
			asm.Bytes(0x40|dhi|shi<<2, 0x01|uint8(op), 0xC0|dlo|slo<<3)
		}
	case 8:
		asm.Bytes(0x48|dhi|shi<<2, 0x01|uint8(op), 0xC0|dlo|slo<<3)
	}
	return asm
}

// off_m(%reg_m) OP= %reg_src
func (asm *Asm) op2MemReg(op Op2, m Mem, src Reg) *Asm {
	dst := m.reg
	dlo, dhi := dst.lohi()
	slo, shi := src.lohi()
	op_ := uint8(op)
	// assert(SizeOf(m) == 8) // TODO mem access by 1, 2 or 4 bytes

	switch m.offlen(dst.id) {
	case 0:
		asm.Bytes(0x48|dhi|shi<<2, 0x01|op_, dlo|slo<<3).
			quirk24(dst)
	case 1:
		asm.Bytes(0x48|dhi|shi<<2, 0x01|op_, 0x40|dlo|slo<<3).
			quirk24(dst).Int8(int8(m.off))
	case 4:
		asm.Bytes(0x48|dhi|shi<<2, 0x01|op_, 0x80|dlo|slo<<3).
			quirk24(dst).Int32(m.off)
	}
	return asm
}

// %reg_dst OP= off_m(%reg_m)
func (asm *Asm) op2RegMem(op Op2, dst Reg, m Mem) *Asm {
	src := m.reg
	dlo, dhi := dst.lohi()
	slo, shi := src.lohi()
	op_ := uint8(op)
	// assert(SizeOf(m) == 8) // TODO mem access by 1, 2 or 4 bytes

	offlen := m.offlen(dst.id)
	switch offlen {
	case 0:
		asm.Bytes(0x48|dhi<<2|shi, 0x03|op_, dlo<<3|slo).
			quirk24(src)
	case 1:
		asm.Bytes(0x48|dhi<<2|shi, 0x03|op_, 0x40|dlo<<3|slo).
			quirk24(src).Int8(int8(m.off))
	case 4:
		asm.Bytes(0x48|dhi<<2|shi, 0x03|op_, 0x80|dlo<<3|slo).
			quirk24(src).Int32(m.off)
	}
	return asm
}

// off_dst(%reg_dst) OP= off_src(%reg_src)
func (asm *Asm) op2MemMem(op Op2, dst Mem, src Mem) *Asm {
	if isNop2(op, dst, src) {
		return asm
	}
	// not natively supported by amd64,
	// must load src in a register
	r := asm.RegAlloc(src.Kind())
	asm.op2RegMem(MOV, r, src)
	asm.op2MemReg(op, dst, r)
	return asm.RegFree(r)
}

// off_dst(%reg_dst) OP= const
func (asm *Asm) op2MemConst(op Op2, dst Mem, src Const) *Asm {
	// not natively supported by amd64,
	// must load src in a register
	r := asm.RegAlloc(src.kind)
	asm.movRegConst(r, src)
	asm.op2MemReg(op, dst, r)
	return asm.RegFree(r)
}

// dst OP= src is a nop ?
func isNop2(op Op2, dst Arg, src Arg) bool {
	if dst == src {
		switch op {
		case AND, OR, MOV, CAST:
			return true
		default:
			return false
		}
	}
	if c, ok := src.(Const); ok {
		switch op {
		case AND:
			return c.val == -1
		case ADD, OR, SUB, XOR:
			return c.val == 0
		default:
			return false
		}
	}
	return false
}
