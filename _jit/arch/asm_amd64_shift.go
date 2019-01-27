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
 * asm_amd64_shift.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

// %reg_dst SHIFT= const
func (asm *Asm) shiftRegConst(op Op2, dst Reg, c Const) *Asm {
	n := c.val
	assert(n > 0) // shift by 0 is optimized away
	siz := SizeOf(dst)
	if n >= 8*int64(siz) {
		return asm.zeroReg(dst)
	}
	var nbit uint8
	if n == 1 {
		nbit = 0x10
	}
	dlo, dhi := dst.lohi()
	switch siz {
	case 1:
		if dst.id >= RSP {
			asm.Uint8(0x40 | dhi)
		}
		asm.Bytes(0xC0|nbit, uint8(op)|dlo)
	case 2:
		asm.Bytes(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Uint8(0x40 | dhi)
		}
		asm.Bytes(0xC1|nbit, uint8(op)|dlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xC1|nbit, uint8(op)|dlo)
	}
	if n != 1 {
		asm.Uint8(uint8(n))
	}
	return asm
}

// off_dst(%reg_dst) SHIFT= const
func (asm *Asm) shiftMemConst(op Op2, m Mem, c Const) *Asm {
	n := c.val
	assert(n > 0) // shift by 0 is optimized away
	siz := SizeOf(m)
	if n >= 8*int64(siz) {
		return asm.zeroMem(m)
	}

	dst := m.reg
	dlo, dhi := dst.lohi()
	offlen, offbit := m.offlen(dst.id)
	op_ := uint8(op) &^ 0xC0
	var nbit uint8
	if n == 1 {
		nbit = 0x10
	}
	switch siz {
	case 1:
		if dhi != 0 {
			asm.Uint8(0x40 | dhi)
		}
		asm.Bytes(0xC0|nbit, offbit|op_|dlo)
	case 2:
		errorf("unimplemented: %v %v %v", op, m, c)
	case 4:
		errorf("unimplemented: %v %v %v", op, m, c)
	case 8:
		errorf("unimplemented: %v %v %v", op, m, c)
	}
	asm.quirk24(dst)
	switch offlen {
	case 1:
		asm.Int8(int8(m.off))
	case 4:
		asm.Int32(m.off)
	}
	if n != 1 {
		asm.Uint8(uint8(n))
	}
	return asm
}

// %reg_dst SHIFT= %reg_src
func (asm *Asm) shiftRegReg(op Op2, dst Reg, src Reg) *Asm {
	errorf("unimplemented: %v %v %v", op, dst, src)
	return asm
}

// %reg_dst SHIFT= off_src(%reg_src)
func (asm *Asm) shiftRegMem(op Op2, dst Reg, m Mem) *Asm {
	errorf("unimplemented: %v %v %v", op, dst, m)
	return asm
}

// off_dst(%reg_dst) SHIFT= %reg_src
func (asm *Asm) shiftMemReg(op Op2, m Mem, src Reg) *Asm {
	errorf("unimplemented: %v %v %v", op, m, src)
	return asm
}

// off_dst(%reg_dst) SHIFT= off_src(%reg_src)
func (asm *Asm) shiftMemMem(op Op2, dst Mem, src Mem) *Asm {
	errorf("unimplemented: %v %v %v", op, dst, src)
	return asm
}
