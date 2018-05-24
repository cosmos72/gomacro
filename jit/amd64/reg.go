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
 * reg.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

// ------------------- Reg -------------------------

type Reg uint8

type Regs uint16

const (
	NoReg Reg = iota
	AX
	CX
	DX
	BX
	SP
	BP
	SI
	DI
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
	FirstReg Reg = AX
	LastReg  Reg = R15
)

func (r Reg) Valid() bool {
	return r >= AX && r <= R15
}

func (r Reg) Validate() {
	if !r.Valid() {
		errorf("invalid register: %d", r)
	}
}

func (r Reg) bits() uint8 {
	r.Validate()
	return uint8(r) - 1
}

func (r Reg) mask() Regs {
	return Regs(1) << r.bits()
}

func (r Reg) lo() uint8 {
	return r.bits() & 0x7
}

func (r Reg) hi() uint8 {
	return (r.bits() & 0x8) >> 3
}

func (r Reg) lohi() (uint8, uint8) {
	bits := r.bits()
	return bits & 0x7, (bits & 0x8) >> 3
}

// ------------------- Regs -------------------------

func RegSet(rs ...Reg) Regs {
	var ret Regs
	for _, r := range rs {
		ret |= r.mask()
	}
	return ret
}

func (rs *Regs) Init() {
	*rs = SP.mask() | BP.mask() | DI.mask()
}

func (rs *Regs) Set(r Reg) {
	*rs |= r.mask()
}

func (rs *Regs) Unset(r Reg) {
	*rs &^= r.mask()
}

func (rs *Regs) Contains(r Reg) bool {
	return *rs&r.mask() != 0
}

func (rs *Regs) Alloc() Reg {
	for r := FirstReg; r <= LastReg; r++ {
		mask := r.mask()
		if *rs&mask == 0 {
			*rs |= mask
			return r
		}
	}
	errorf("no free registers")
	return NoReg
}

func (rs *Regs) Free(r Reg) {
	*rs &^= r.mask()
}

// ---------------- Asm.Mov, Asm.Store, Asm.Load -------------------------

func (asm *Asm) Mov(dst Reg, src Reg) *Asm {
	if dst == src {
		return asm
	}
	slo, shi := src.lohi()
	dlo, dhi := dst.lohi()
	return asm.Bytes(0x48|dhi|shi*4, 0x89, 0xc0+dlo+slo*8) //  movq   %reg_src,%reg_dst
}

func (asm *Asm) Store(dst *Var, r Reg) *Asm {
	lo, hi := r.lohi()
	return asm.Bytes(0x48|hi*4, 0x89, 0x87|lo*8).Idx(dst) //   movq   %reg,dst(%rdi)
}

func (asm *Asm) Load(r Reg, src *Var) *Asm {
	if src.Const {
		return asm.LoadConst(r, src.Val)
	}
	lo, hi := r.lohi()
	return asm.Bytes(0x48|hi*4, 0x8b, 0x87|lo*8).Idx(src) //   movq   src(%rdi),%reg
}

func (asm *Asm) LoadConst(r Reg, val int64) *Asm {
	lo, hi := r.lohi()
	if val == int64(uint32(val)) {
		if hi != 0 {
			asm.Bytes(0x41)
		}
		return asm.Bytes(0xb8 + lo).Uint32(uint32(val)) //            movl   $val,%regl // zero extend
	} else if val == int64(int32(val)) {
		return asm.Bytes(0x48|hi, 0xc7, 0xc0|lo).Int32(int32(val)) // movq   $val,%reg  // sign extend
	} else {
		return asm.Bytes(0x48|hi, 0xb8+lo).Int64(val) //              movabs $val,%reg
	}
}
