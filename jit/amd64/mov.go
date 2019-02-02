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
 * mov.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

// ============================================================================
func (asm *Asm) Mov(src Arg, dst Arg) *Asm {
	return asm.Op2(MOV, src, dst)
}

// zero a register: use XOR
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

// %reg_dst = const
func (asm *Asm) movConstReg(c Const, dst Reg) *Asm {
	if c.val == 0 {
		return asm.zeroReg(dst)
	}
	switch dst.kind.Size() {
	case 1:
		asm.movConstReg8(c, dst)
	case 2:
		asm.movConstReg16(c, dst)
	case 4:
		asm.movConstReg32(c, dst)
	case 8:
		asm.movConstReg64(c, dst)
	}
	return asm
}

func (asm *Asm) movConstReg8(c Const, dst Reg) *Asm {
	dlo, dhi := dst.lohi()
	val := c.val
	if val == int64(int8(val)) {
		if dst.id < RSP {
			asm.Bytes(0xB0 | dlo)
		} else {
			asm.Bytes(0x40|dhi, 0xB0|dlo)
		}
		asm.Int8(int8(val))
	} else {
		errorf("sign-extended constant overflows 8-bit destination: %v %v %v", MOV, c, dst)
	}
	return asm
}

func (asm *Asm) movConstReg16(c Const, dst Reg) *Asm {
	dlo, dhi := dst.lohi()
	val := c.val
	if val == int64(int16(val)) {
		if dhi == 0 {
			asm.Bytes(0x66, 0xB8|dlo)
		} else {
			asm.Bytes(0x66, 0x40|dhi, 0xB8|dlo)
		}
		asm.Int16(int16(val))
	} else {
		errorf("sign-extended constant overflows 16-bit destination: %v %v %v", MOV, c, dst)
	}
	return asm
}

func (asm *Asm) movConstReg32(c Const, dst Reg) *Asm {
	dlo, dhi := dst.lohi()
	val := c.val
	if val == int64(int32(val)) {
		if dhi == 0 {
			asm.Byte(0xB8 | dlo)
		} else {
			asm.Bytes(40|dhi, 0xB8|dlo)
		}
		asm.Int32(int32(val))
	} else {
		errorf("sign-extended constant overflows 16-bit destination: %v %v %v", MOV, c, dst)
	}
	return asm
}

func (asm *Asm) movConstReg64(c Const, dst Reg) *Asm {
	dlo, dhi := dst.lohi()
	val := c.val
	// 32-bit signed immediate constants, use mov
	if val == int64(int32(val)) {
		return asm.Bytes(0x48|dhi, 0xC7, 0xC0|dlo).Int32(int32(val))
	}
	// 64-bit constant, must use movabs
	return asm.Bytes(0x48|dhi, 0xB8|dlo).Int64(val)
}

// off_dst(%reg_dst) = const
func (asm *Asm) movConstMem(c Const, m Mem) *Asm {
	dst := m.reg
	dlo, dhi := dst.lohi()
	offlen, offbit := m.offlen(dst.id)
	val := c.val
	switch dst.kind.Size() {
	case 1:
		if dhi == 0 {
			asm.Bytes(0xC6, offbit|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0xC6, offbit|dlo)
		}
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi == 0 {
			asm.Bytes(0xC7, offbit|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0xC7, offbit|dlo)
		}
	case 8:
		if val == int64(int32(val)) {
			asm.Bytes(0x48|dhi, 0xC7, offbit|dlo)
		} else {
			r := asm.RegAlloc(dst.Kind())
			return asm.movConstReg64(c, r).Op2(MOV, r, dst).RegFree(r)
		}
	}
	asm.quirk24(dst)
	switch offlen {
	case 1:
		asm.Int8(int8(m.off))
	case 4:
		asm.Int32(m.off)
	}

	switch dst.kind.Size() {
	case 1:
		asm.Int8(int8(val))
	case 2:
		asm.Int16(int16(val))
	case 4, 8:
		asm.Int32(int32(val))
	}
	return asm
}

// ============================================================================
// movsx, movzx or mov
func (asm *Asm) Cast(src Arg, dst Arg) *Asm {
	if src == dst {
		return asm
	} else if SizeOf(src) == SizeOf(dst) {
		return asm.Op2(MOV, src, dst)
	}
	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Reg:
			asm.castRegReg(src, dst)
		case Mem:
			asm.castMemReg(src, dst)
		case Const:
			src = src.Cast(dst.kind)
			asm.movConstReg(src, dst)
		default:
			errorf("unsupported source type %T, expecting Reg, Mem or Const: %v %v %v", src, CAST, src, dst)
		}
	case Mem:
		switch src := src.(type) {
		case Reg:
			asm.castRegMem(src, dst)
		case Mem:
			asm.castMemMem(src, dst)
		case Const:
			src = src.Cast(dst.Kind())
			asm.op2ConstMem(MOV, src, dst)
		default:
			errorf("unsupported source type %T, expecting Reg, Mem or Const: %v %v %v", src, CAST, src, dst)
		}
	case Const:
		errorf("destination cannot be a constant: %v %v %v", CAST, src, dst)
	default:
		errorf("unsupported destination type %T, expecting Reg or Mem: %v %v %v", dst, CAST, src, dst)
	}
	return asm
}

func (asm *Asm) castRegReg(src Reg, dst Reg) *Asm {
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
		return asm.op2RegReg(MOV, src, MakeReg(dst.id, src.kind))
	default:
		errorf("unsupported source register size %v, expecting 1, 2, 4 or 8: %v %v %v",
			SizeOf(src), CAST, src, dst)
	}
	// for simplicity, assume Sizeof(dst) == 8
	return asm.Bytes(0x48|dhi<<2|shi, 0x0F, op, 0xC0|dlo<<3|slo)
}

func (asm *Asm) castMemReg(src_m Mem, dst Reg) *Asm {
	src := src_m.reg

	var op uint8 = 0xB6 // movzx
	if src.kind.Signed() {
		op = 0xBE // movsx
	}
	dlo, dhi := dst.lohi()
	slo, shi := src.lohi()
	offlen, offbit := src_m.offlen(src.id)
	// debugf("castRegMem() src = %v, dst = %v", src, dst)
	switch SizeOf(src) {
	case 1:
		// movzbq, movsbq
	case 2:
		op++ // movzwq, movswq
	case 4:
		if src.kind.Signed() {
			// sign-extend 32bit -> 64bit
			// movsd i.e. movslq
			asm.Bytes(0x48|dhi<<2|shi, 0x63, offbit|dlo<<3|slo)
			asm.quirk24(src)
			switch offlen {
			case 1:
				asm.Int8(int8(src_m.off))
			case 4:
				asm.Int32(src_m.off)
			}
			return asm
		}
		// amd64 does not have zero-extend 32bit -> 64bit
		// because operations that write into 32bit registers
		// already zero the upper 32 bits.
		// So just compile as a regular MOV
		// debugf("zero-extend 32bit -> 64bit: dst = %v, src = %v", dst, m)
		fallthrough
	case 8:
		return asm.op2MemReg(MOV, src_m, MakeReg(dst.id, src.kind))
	default:
		errorf("invalid source register size %v, expecting 1, 2, 4 or 8: %v %v %v",
			SizeOf(src), CAST, src, dst)
	}
	// for simplicity, assume Sizeof(dst) == 8
	asm.Bytes(0x48|dhi<<2|shi, 0x0F, op, offbit|dlo<<3|slo)
	asm.quirk24(src)
	switch offlen {
	case 1:
		asm.Int8(int8(src_m.off))
	case 4:
		asm.Int32(src_m.off)
	}
	return asm
}

func (asm *Asm) castRegMem(src Reg, dst_m Mem) *Asm {
	dst := dst_m.reg
	// assume that user code cannot use the same register
	// multiple times with different kinds
	r := MakeReg(src.id, dst.kind)
	asm.castRegReg(src, r)
	return asm.op2RegMem(MOV, r, dst_m)
}

func (asm *Asm) castMemMem(src Mem, dst Mem) *Asm {
	if SizeOf(src) > SizeOf(dst) && !src.Kind().IsFloat() {
		// just read the lowest bytes from src
		asm.op2MemMem(MOV,
			MakeMem(src.off, src.reg.id, dst.reg.kind),
			dst)
	} else {
		r := asm.RegAlloc(dst.Kind())
		asm.castMemReg(src, r)
		asm.op2RegMem(MOV, r, dst)
		asm.RegFree(r)
	}
	return asm
}
