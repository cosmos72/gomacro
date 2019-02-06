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
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

// ============================================================================
func (asm *Asm) Mov(src Arg, dst Arg) *Asm {
	assert(SizeOf(src) == SizeOf(dst))

	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Const:
			asm.movConstReg(src, dst)
		case Reg:
			asm.movRegReg(src, dst)
		case Mem:
			asm.movMemReg(src, dst)
		default:
			errorf("unknown source type %T, expecting Const, Reg or Mem: %v %v, %v", src, MOV, src, dst)
		}
	case Mem:
		switch src := src.(type) {
		case Const:
			asm.movConstMem(src, dst)
		case Reg:
			asm.movRegMem(src, dst)
		case Mem:
			asm.movMemMem(src, dst)
		default:
			errorf("unknown source type %T, expecting Const, Reg or Mem: %v %v, %v", src, MOV, src, dst)
		}
	case Const:
		errorf("destination cannot be a constant: %v %v, %v", MOV, src, dst)
	default:
		errorf("unimplemented destination type %T, expecting Reg: %v %v, %v", dst, MOV, src, dst)
	}
	return asm
}

// handy alias
func (asm *Asm) load(src Mem, dst Reg) *Asm {
	return asm.movMemReg(src, dst)
}

// handy alias
func (asm *Asm) store(src Reg, dst Mem) *Asm {
	return asm.movRegMem(src, dst)
}

// zero a register
func (asm *Asm) zeroReg(dst Reg) *Asm {
	return asm.movConstReg(MakeConst(0, dst.kind), dst)
}

// zero a memory location
func (asm *Asm) zeroMem(dst Mem) *Asm {
	errorf("unimplemented: zeroMem")
	return asm
}

func (asm *Asm) movRegReg(src Reg, dst Reg) *Asm {
	// arm64 implements "mov src,dst" as "orr xzr,src,dst"
	return asm.Uint32(dst.kind.kbit() | 0x2A0003E0 | src.val()<<16 | dst.val())
}

func (asm *Asm) movConstReg(c Const, dst Reg) *Asm {
	xzr := MakeReg(XZR, dst.kind)
	var immcval uint32
	var movk bool
	if c.val >= 0 && c.val < 0x10000 {
		immcval = 0x40<<19 | uint32(c.val)
	} else if c.val < 0 && c.val >= -0x10000 {
		immcval = uint32(^c.val)
	} else if asm.tryOp3RegConstReg(OR3, xzr, uint64(c.val), dst) {
		return asm
	} else if asm.tryOp3RegConstReg(OR3, xzr, uint64(uint32(c.val)), dst) {
		if dst.kind.Size() == 8 {
			asm.movk(uint16(c.val>>32), 32, dst)
			asm.movk(uint16(c.val>>48), 48, dst)
		}
		return asm
	} else {
		immcval = 0x40<<19 | uint32(c.val&0xFFFF)
		movk = true
	}
	asm.Uint32(dst.kind.kbit() | 0x12800000 | immcval<<5 | dst.val())
	if movk {
		asm.movk(uint16(c.val>>16), 16, dst)
		if dst.kind.Size() == 8 {
			asm.movk(uint16(c.val>>32), 32, dst)
			asm.movk(uint16(c.val>>48), 48, dst)
		}
	}
	return asm
}

// set some bits of dst, preserving others
func (asm *Asm) movk(val uint16, shift uint8, dst Reg) *Asm {
	if val != 0 {
		asm.Uint32(dst.kind.kbit() | 0xF2800000 | uint32(shift)<<17 | uint32(val)<<5 | dst.val())
	}
	return asm
}

func (asm *Asm) movMemReg(src Mem, dst Reg) *Asm {
	off := src.off
	var sizebit uint32
	switch src.Kind().Size() {
	case 1:
		if off >= 0 && off <= 4095 {
			return asm.Uint32(0x39400000 | uint32(src.off)<<10 | src.reg.valOrX31(true)<<5 | dst.val())
		}
	case 2:
		if off >= 0 && off <= 8190 && off%2 == 0 {
			return asm.Uint32(0x79400000 | uint32(src.off)<<9 | src.reg.valOrX31(true)<<5 | dst.val())
		}
		sizebit = 4
	case 4:
		if off >= 0 && off <= 16380 && off%4 == 0 {
			return asm.Uint32(0xB9400000 | uint32(src.off)<<8 | src.reg.valOrX31(true)<<5 | dst.val())
		}
		sizebit = 8
	case 8:
		if off >= 0 && off <= 32760 && off%8 == 0 {
			return asm.Uint32(0xF9400000 | uint32(src.off)<<7 | src.reg.valOrX31(true)<<5 | dst.val())
		}
		sizebit = 12
	}
	// load offset in a register. we could also try "ldur"...
	r := asm.RegAlloc(Uint64)
	asm.movConstReg(ConstInt64(int64(off)), r)

	asm.Uint32(sizebit<<28 | 0x38606800 | r.val()<<16 | src.reg.valOrX31(true)<<5 | dst.val())

	return asm.RegFree(r)
}

func (asm *Asm) movConstMem(c Const, dst Mem) *Asm {
	r := asm.RegAlloc(dst.Kind())
	return asm.movConstReg(c, r).movRegMem(r, dst).RegFree(r)
}

func (asm *Asm) movRegMem(src Reg, dst Mem) *Asm {
	errorf("unimplemented: %v %v, %v", MOV, src, dst)
	return asm
}

func (asm *Asm) movMemMem(src Mem, dst Mem) *Asm {
	r := asm.RegAlloc(src.Kind())
	return asm.movMemReg(src, r).movRegMem(r, dst).RegFree(r)
}

// ============================================================================
func (asm *Asm) Cast(src Arg, dst Arg) *Asm {
	if src == dst {
		return asm
	} else if SizeOf(src) == SizeOf(dst) {
		return asm.Op2(MOV, src, dst)
	}
	errorf("unimplemented: %v %v, %v", CAST, src, dst)
	return asm
}
