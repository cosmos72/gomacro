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
 * mul.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

// amd64 has very constrained 8-bit multiply
// (it can only read/write %al), so use at least 16-bit
func (asm *Asm) mul2WidenReg(r Reg) Reg {
	switch r.kind {
	case Bool, Uint8:
		w := MakeReg(r.id, Uint16)
		asm.castRegReg(r, w)
		return w
	case Int8:
		w := MakeReg(r.id, Int16)
		asm.castRegReg(r, w)
		return w
	}
	return r
}

func (asm *Asm) mul2ConstReg(c Const, dst Reg) *Asm {
	n := c.val
	switch n {
	case -1:
		return asm.op1Reg(NEG, dst)
	case 0:
		return asm.zeroReg(dst)
	case 1:
		return asm
	case 4, 8:
		if SizeOf(dst) == 8 {
			return asm.lea4(MakeMem(0, NoRegId, dst.kind), dst, n, dst)
		}
	case 2, 3, 5, 9:
		if SizeOf(dst) == 8 {
			return asm.lea4(MakeMem(0, dst.id, dst.kind), dst, n-1, dst)
		}
	}
	if n&(n-1) == 0 {
		// TODO shift
	}
	if n == int64(int8(n)) {
		return asm.mul3RegConst8Reg(dst, int8(n), dst)
	}
	// constant is too large, must allocate a register
	dst = asm.mul2WidenReg(dst)
	r := asm.RegAlloc(dst.kind)
	asm.movConstReg64(MakeConst(c.val, dst.kind), r)
	return asm.mul2RegReg(r, dst).RegFree(r)
}

func (asm *Asm) mul2RegReg(src Reg, dst Reg) *Asm {
	src = asm.mul2WidenReg(src)
	dst = asm.mul2WidenReg(dst)
	slo, shi := src.lohi()
	dlo, dhi := dst.lohi()
	switch dst.kind.Size() {
	case 1:
		errorf("internal error, Asm.mul2WidenReg did not widen 8-bit registers: %v %v, %v", MUL, src, dst)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi<<2|shi == 0 {
			asm.Bytes(0x0F, 0xAF, 0xC0|dlo<<3|slo)
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x0F, 0xAF, 0xC0|dlo<<3|slo)
		}
	case 8:
		asm.Bytes(0x48|dhi<<2|shi, 0x0F, 0xAF, 0xC0|dlo<<3|slo)
	}
	return asm
}

func (asm *Asm) mul2MemReg(src_m Mem, dst Reg) *Asm {
	src := src_m.reg
	slo, shi := src.lohi()
	dlo, dhi := dst.lohi()
	offlen, offbit := src_m.offlen(src.id)
	switch dst.kind.Size() {
	case 1:
		// amd64 has very constrained 8-bit multiply
		// (it can only read/write %al), so copy 8-bit memory
		// to a register and use widening multiplication
		r := asm.RegAlloc(src_m.Kind())
		return asm.Load(src_m, r).mul2RegReg(r, dst).RegFree(r)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi<<2|shi == 0 {
			asm.Bytes(0x0F, 0xAF, offbit|dlo<<3|slo)
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x0F, 0xAF, offbit|dlo<<3|slo)
		}
	case 8:
		asm.Bytes(0x48|dhi<<2|shi, 0x0F, 0xAF, offbit|dlo<<3|slo)
	}
	asm.quirk24(src)
	switch offlen {
	case 1:
		asm.Int8(int8(src_m.off))
	case 4:
		asm.Int32(src_m.off)
	}
	return asm
}

func (asm *Asm) mul2ConstMem(c Const, m Mem) *Asm {
	switch c.val {
	case -1:
		return asm.op1Mem(NEG, m)
	case 0:
		return asm.zeroMem(m)
	case 1:
		return asm
	default:
		r := asm.RegAlloc(m.Kind())
		return asm.Load(m, r).mul2ConstReg(c, r).Store(r, m).RegFree(r)
	}
}

func (asm *Asm) mul2RegMem(src Reg, dst_m Mem) *Asm {
	// must allocate a register
	r := asm.RegAlloc(dst_m.Kind())
	return asm.Load(dst_m, r).mul2RegReg(src, r).Store(r, dst_m).RegFree(r)
}

func (asm *Asm) mul2MemMem(src_m Mem, dst_m Mem) *Asm {
	// must allocate a register
	r := asm.RegAlloc(dst_m.Kind())
	return asm.Load(dst_m, r).mul2MemReg(src_m, r).Store(r, dst_m).RegFree(r)
}

// =============== 3-argument MUL3 ==================

func (asm *Asm) mul3RegConst8Reg(src Reg, cval int8, dst Reg) *Asm {
	src = asm.mul2WidenReg(src)
	dst = asm.mul2WidenReg(dst)
	slo, shi := src.lohi()
	dlo, dhi := dst.lohi()
	switch dst.kind.Size() {
	case 1:
		errorf("internal error, Asm.mul2WidenReg did not widen 8-bit registers: %v %v, %v, %v", MUL, src, cval, dst)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi<<2|shi == 0 {
			asm.Bytes(0x6B, 0xC0|dlo<<3|slo, uint8(cval))
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x6B, 0xC0|dlo<<3|slo, uint8(cval))
		}
	case 8:
		asm.Bytes(0x48|dhi<<2|shi, 0x6B, 0xC0|dlo<<3|slo, uint8(cval))
	}
	return asm
}

func (asm *Asm) mul3MemConst8Reg(src_m Mem, cval int8, dst Reg) *Asm {
	src := src_m.reg
	slo, shi := src.lohi()
	dlo, dhi := dst.lohi()
	offlen, offbit := src_m.offlen(src.id)
	switch dst.kind.Size() {
	case 1:
		errorf("internal error, missing call to Asm.mul2WidenReg to widen 8-bit registers: %v %v, %v, %v", MUL, src_m, cval, dst)
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi<<2|shi == 0 {
			asm.Bytes(0x6B, offbit|dlo<<3|slo)
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x6B, offbit|dlo<<3|slo)
		}
	case 8:
		asm.Bytes(0x48|dhi<<2|shi, 0x6B, offbit|dlo<<3|slo)
	}
	asm.quirk24(src)
	switch offlen {
	case 1:
		asm.Int8(int8(src_m.off))
	case 4:
		asm.Int32(src_m.off)
	}
	return asm.Int8(cval)
}
