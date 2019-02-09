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
 * shift.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

// %reg_dst SHIFT= const
func (asm *Asm) shiftConstReg(op Op2, c Const, dst Reg) *Asm {
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
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xC0|nbit, uint8(op)|dlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Byte(0x40 | dhi)
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
func (asm *Asm) shiftConstMem(op Op2, c Const, m Mem) *Asm {
	n := c.val
	assert(n > 0) // shift by 0 is optimized away
	size := SizeOf(m)
	if n >= 8*int64(size) {
		if m.Kind().Signed() {
			n = 8*int64(size) - 1
		} else {
			return asm.zeroMem(m)
		}
	}
	dst := m.reg
	dlo, dhi := dst.lohi()
	offlen, offbit := m.offlen(dst.id)
	op_ := uint8(op) &^ 0xC0

	var nbit uint8
	if n == 1 {
		nbit = 0x10
	}
	switch size {
	case 1:
		if dst.id >= RSP {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xC0|nbit, offbit|op_|dlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xC1|nbit, offbit|op_|dlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xC1|nbit, offbit|op_|dlo)
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
func (asm *Asm) shiftRegReg(op Op2, src Reg, dst Reg) *Asm {
	if dst.id == RCX {
		errorf("unimplemented shift RCX by Reg: %v %v %v", op, src, dst)
	}
	if src.id != RCX {
		asm.op2RegReg(MOV, src, MakeReg(RCX, src.kind))
	}
	siz := SizeOf(dst)
	dlo, dhi := dst.lohi()

	switch siz {
	case 1:
		if dst.id >= RSP {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xD2, uint8(op)|dlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xD3, uint8(op)|dlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xD3, uint8(op)|dlo)
	}
	return asm
}

// off_dst(%reg_dst) SHIFT= %reg_src
func (asm *Asm) shiftRegMem(op Op2, src Reg, dst_m Mem) *Asm {
	if dst_m.RegId() == RCX {
		errorf("unimplemented shift Mem[RCX] by Reg: %v %v %v", op, src, dst_m)
	}
	if src.id != RCX {
		asm.op2RegReg(MOV, src, MakeReg(RCX, src.kind))
	}
	siz := SizeOf(dst_m)
	dst := dst_m.reg
	dlo, dhi := dst.lohi()
	offlen, offbit := dst_m.offlen(dst.id)
	op_ := uint8(op) &^ 0xC0

	switch siz {
	case 1:
		if dst.id >= RSP {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xD2, offbit|op_|dlo)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Byte(0x40 | dhi)
		}
		asm.Bytes(0xD3, offbit|op_|dlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xD3, offbit|op_|dlo)
	}
	asm.quirk24(dst)
	switch offlen {
	case 1:
		asm.Int8(int8(dst_m.off))
	case 4:
		asm.Int32(dst_m.off)
	}
	return asm
}

// %reg_dst SHIFT= off_src(%reg_src)
func (asm *Asm) shiftMemReg(op Op2, src_m Mem, dst Reg) *Asm {
	if dst.id == RCX {
		errorf("unimplemented shift RCX by Mem: %v %v %v", op, src_m, dst)
	}
	r := MakeReg(RCX, src_m.Kind())
	asm.op2MemReg(MOV, src_m, r)
	asm.shiftRegReg(op, r, dst)
	return asm
}

// off_dst(%reg_dst) SHIFT= off_src(%reg_src)
func (asm *Asm) shiftMemMem(op Op2, src_m Mem, dst_m Mem) *Asm {
	if dst_m.RegId() == RCX {
		errorf("unimplemented shift Mem[RCX] by Mem: %v %v %v", op, src_m, dst_m)
	} else if src_m.RegId() == RCX {
		errorf("unimplemented shift Mem by Mem[RCX]: %v %v %v", op, src_m, dst_m)
	}
	r := MakeReg(RCX, src_m.Kind())
	asm.op2MemReg(MOV, src_m, r)
	asm.shiftRegMem(op, r, dst_m)
	return asm
}
