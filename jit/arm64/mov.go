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

	rsrc, rsok := src.(Reg)
	rdst, rdok := dst.(Reg)
	if !rsok {
		errorf("unimplemented source type %T, expecting Reg: %v %v %v", src, MOV, src, dst)
	}
	if _, ok := dst.(Const); ok {
		errorf("destination cannot be a constant: %v %v %v", dst, MOV, src, dst)
	}
	if !rdok {
		errorf("unimplemented destination type %T, expecting Reg: %v %v %v", dst, MOV, src, dst)
	}
	return asm.movRegReg(rsrc, rdst)
}

// zero a register: use XOR
func (asm *Asm) zeroReg(dst Reg) *Asm {
	return asm.movRegConst(MakeConst(0, dst.kind), dst)
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

func (asm *Asm) movRegConst(c Const, dst Reg) *Asm {
	var cbit, cval uint32
	if c.val >= 0 && c.val < 0x10000 {
		cval = uint32(c.val)
		cbit = 0x80 << 24
	} else if c.val < 0 && c.val >= -0x10000 {
		cval = uint32(^c.val)
	} else {
		errorf("unimplemented MOV const,reg with large constant: %v %v %v", MOV, c, dst)
	}
	return asm.Uint32(cbit | dst.kind.kbit() | 0x12800000 | cval<<5 | dst.val())
}

// ============================================================================
func (asm *Asm) Cast(src Arg, dst Arg) *Asm {
	if src == dst {
		return asm
	} else if SizeOf(src) == SizeOf(dst) {
		return asm.Op2(MOV, src, dst)
	}
	errorf("unimplemented: %v %v %v", CAST, src, dst)
	return asm
}
