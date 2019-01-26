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
 * asm_amd64_mov.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

func (asm *Asm) Mov(dst Arg, src Arg) *Asm {
	return asm.Op2(MOV, dst, src)
}

// %reg_dst = const
func (asm *Asm) movRegConst(dst Reg, c Const) *Asm {
	if c.val == 0 {
		return asm.xorRegSelf(dst)
	}
	dlo, dhi := dst.lohi()
	// 32-bit signed immediate constants, use mov
	if c.val == int64(int32(c.val)) {
		return asm.Bytes(0x48|dhi, 0xC7, 0xC0|dlo).Int32(int32(c.val))
	}
	// 64-bit constant, must use movabs
	return asm.Bytes(0x48|dhi, 0xB8|dlo).Int64(c.val)
}

// %reg_dst ^= %reg_dst // compact way to zero a register
func (asm *Asm) xorRegSelf(dst Reg) *Asm {
	dlo, dhi := dst.lohi()
	if dhi == 0 {
		return asm.Bytes(0x31, 0xC0|dlo|dlo<<3)
	} else {
		return asm.Bytes(0x48|dhi<<1|dhi<<2, 0x31, 0xC0|dlo|dlo<<3)
	}
}

// movsx, movzx or mov
func (asm *Asm) Cast(dst Arg, src Arg) *Asm {
	if src == dst {
		return asm
	} else if SizeOf(src) == SizeOf(dst) {
		return asm.Op2(MOV, dst, src)
	}
	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Reg:
			asm.castRegReg(dst, src)
		case Mem:
			asm.castRegMem(dst, src)
		case Const:
			src = src.Cast(dst.kind)
			asm.movRegConst(dst, src)
		default:
			errorf("Cast: unsupported source type, expecting Reg, Mem or Const: %v %v %v", CAST, dst, src)
		}
	case Mem:
		switch src := src.(type) {
		case Reg:
			asm.castMemReg(dst, src)
		case Mem:
			asm.castMemMem(dst, src)
		case Const:
			src = src.Cast(dst.Kind())
			asm.op2MemConst(MOV, dst, src)
		default:
			errorf("Cast: unsupported source type, expecting Reg, Mem or Const: %v %v %v", CAST, dst, src)
		}
	case Const:
		errorf("Cast: destination cannot be a constant: %v %v %v", CAST, dst, src)
	default:
		errorf("Cast: unsupported destination type, expecting Reg or Mem: %v %v %v", CAST, dst, src)
	}
	return asm
}

func (asm *Asm) castRegReg(dst Reg, src Reg) *Asm {
	var op uint8 = 0xB6 // movzx
	if dst.kind.Signed() {
		op = 0xBE // movsx
	}
	dlo, dhi := dst.lohi()
	slo, shi := src.lohi()
	switch SizeOf(src) {
	case 1:
		// movzbq, movsbq
	case 2:
		op++ // movzwq, movswq
	case 4:
		if dst.kind.Signed() {
			// movsd i.e. movslq
			return asm.Bytes(0x48|dhi<<2|shi, 0x63, 0xC0|dlo<<3|slo)
		}
		// amd64 does not have zero-extend 32bit -> 64bit
		// because operations that write into 32bit registers
		// already zero the upper 32 bits.
		// So just compile as a regular MOV
		fallthrough
	case 8:
		return asm.op2RegReg(MOV, MakeReg(dst.id, src.kind), src)
	default:
		errorf("invalid register size %v, expecting 1, 2, 4 or 8: %v %v %v",
			SizeOf(src), CAST, dst, src)
	}
	// for simplicity, assume Sizeof(dst) == 8
	return asm.Bytes(0x48|dhi<<2|shi, 0x0F, op, 0xC0|dlo<<3|slo)
}

func (asm *Asm) castRegMem(dst Reg, m Mem) *Asm {
	src := m.reg

	var op uint8 = 0xB6 // movzx
	if src.kind.Signed() {
		op = 0xBE // movsx
	}
	dlo, dhi := dst.lohi()
	slo, shi := src.lohi()
	offlen := m.offlen(src.id)
	debugf("castRegMem() dst = %v, src = %v", dst, src)
	switch SizeOf(src) {
	case 1:
		// movzbq, movsbq
	case 2:
		op++ // movzwq, movswq
	case 4:
		if src.kind.Signed() {
			// sign-extend 32bit -> 64bit
			// movsd i.e. movslq
			switch offlen {
			case 0:
				asm.Bytes(0x48|dhi<<2|shi, 0x63, dlo<<3|slo).
					quirk24(src)
			case 1:
				asm.Bytes(0x48|dhi<<2|shi, 0x63, 0x40|dlo<<3|slo).
					quirk24(src).Int8(int8(m.off))
			case 4:
				asm.Bytes(0x48|dhi<<2|shi, 0x63, 0x80|dlo<<3|slo).
					quirk24(src).Int32(m.off)
			}
			return asm
		}
		// amd64 does not have zero-extend 32bit -> 64bit
		// because operations that write into 32bit registers
		// already zero the upper 32 bits.
		// So just compile as a regular MOV
		debugf("zero-extend 32bit -> 64bit: dst = %v, src = %v", dst, m)
		fallthrough
	case 8:
		return asm.op2RegMem(MOV, MakeReg(dst.id, src.kind), m)
	default:
		errorf("invalid register size %v, expecting 1, 2, 4 or 8: %v %v %v",
			SizeOf(src), CAST, dst, src)
	}
	// for simplicity, assume Sizeof(dst) == 8
	switch offlen {
	case 0:
		asm.Bytes(0x48|dhi<<2|shi, 0x0F, op, dlo<<3|slo).
			quirk24(src)
	case 1:
		asm.Bytes(0x48|dhi<<2|shi, 0x0F, op, 0x40|dlo<<3|slo).
			quirk24(src).Int8(int8(m.off))
	case 4:
		asm.Bytes(0x48|dhi<<2|shi, 0x0F, op, 0x80|dlo<<3|slo).
			quirk24(src).Int32(m.off)
	}
	return asm
}

func (asm *Asm) castMemReg(m Mem, src Reg) *Asm {
	dst := m.reg
	// assume that user code cannot use the same register
	// twice with different kinds
	r := MakeReg(src.id, dst.kind)
	asm.castRegReg(r, src)
	return asm.op2MemReg(MOV, m, r)
}

func (asm *Asm) castMemMem(dst Mem, src Mem) *Asm {
	r := asm.RegAlloc(dst.Kind())
	asm.castRegMem(r, src)
	asm.op2MemReg(MOV, dst, r)
	return asm.RegFree(r)
}
