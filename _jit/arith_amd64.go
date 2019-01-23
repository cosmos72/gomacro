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
 * arith_amd64.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

// dst = a + b
func (asm *Asm) Add(dst Arg, a Arg, b Arg) *Asm {
	if dst.Eq(a) {
		return asm.add2(dst, b)
	} else if dst.Eq(b) {
		return asm.add2(dst, a)
	}
	r, alloc := asm.hwAlloc(dst)
	asm.load(r, a)
	asm.hwOp2(hwADD, r, asm.hw(b))
	return asm.hwFree(dst, r, alloc)
}

// dst += src
func (asm *Asm) add2(dst Arg, src Arg) *Asm {
	r, alloc := asm.hwAlloc(src)
	asm.hwOp2(hwADD, asm.hw(dst), r)
	return asm.hwFree(src, r, alloc)
}

// dst = a - b
func (asm *Asm) Sub(dst Arg, a Arg, b Arg) *Asm {
	if dst.Eq(a) {
		return asm.sub2(dst, b)
	}
	r, alloc := asm.hwAlloc(dst)
	asm.load(r, a)
	asm.hwOp2(hwSUB, r, asm.hw(b))
	return asm.hwFree(dst, r, alloc)
}

// dst -= src
func (asm *Asm) sub2(dst Arg, src Arg) *Asm {
	r, alloc := asm.hwAlloc(src)
	asm.hwOp2(hwSUB, asm.hw(dst), r)
	return asm.hwFree(src, r, alloc)
}

// dst = a * b
func (asm *Asm) SMul(dst Arg, a Arg, b Arg) *Asm {
	if dst.Eq(a) {
		return asm.smul2(dst, b)
	} else if dst.Eq(b) {
		return asm.smul2(dst, a)
	}
	// TODO
	return asm
}

// dst *= src
func (asm *Asm) smul2(dst Arg, src Arg) *Asm {
	// TODO
	return asm
}

// ---------------- DIV --------------------

// dst = a / b // signed division
func (asm *Asm) SDiv(dst Arg, a Arg, b Arg) *Asm {
	return asm.divrem(dst, a, b, div|signed)
}

// dst = a / b // unsigned division
func (asm *Asm) UDiv(dst Arg, a Arg, b Arg) *Asm {
	return asm.divrem(dst, a, b, div|unsigned)
}

// ---------------- REM --------------------

// dst = a % b // signed remainder
func (asm *Asm) SRem(dst Arg, a Arg, b Arg) *Asm {
	return asm.divrem(dst, a, b, rem|signed)
}

// dst = a % b // unsigned remainder
func (asm *Asm) URem(dst Arg, a Arg, b Arg) *Asm {
	return asm.divrem(dst, a, b, rem|unsigned)
}

// FIXME: golang remainder rules are NOT the same as C !
func (asm *Asm) divrem(dst Arg, a Arg, b Arg, k divkind) *Asm {
	tosave := newHwRegs(rDX)
	rz := asm.reg(dst)
	if rz != rAX {
		tosave.Set(rAX)
	}
	tosave = asm.pushRegs(tosave)
	var rb Reg
	ra := a.reg(asm)
	if tosave.Contains(ra) {
		rb = asm.alloc()
		asm.Load(rb, a)
		ra = rb
	}
	asm.hwMov(rAX, rz) // nop if z == AX

	switch a := a.(type) {
	case *Var:
		if k&unsigned != 0 {
			asm.Bytes(0x31, 0xd2)              //  xor    %edx,%edx
			asm.Bytes(0x48, 0xf7, 0xb7).Idx(a) //  divq   a(%rdi)
		} else {
			asm.Bytes(0x48, 0x99)              //  cqto
			asm.Bytes(0x48, 0xf7, 0xbf).Idx(a) //  idivq  a(%rdi)
		}
	default:
		tmp, alloc := asm.hwAlloc(a)
		if k&unsigned != 0 {
			asm.Bytes(0x31, 0xd2)                         //  xor    %edx,%edx
			asm.Bytes(0x48+tmp.hi(), 0xf7, 0xf0+tmp.lo()) //  div    %reg_tmp
		} else {
			asm.Bytes(0x48, 0x99)                         //  cqto
			asm.Bytes(0x48+tmp.hi(), 0xf7, 0xf8+tmp.lo()) //  idiv   %reg_tmp
		}
		asm.hwFree(tmp, alloc)
	}
	if b != NoReg {
		asm.Free(b)
	}
	if k&rem != 0 {
		asm.mov(rz, rDX) // nop if z == DX
	} else {
		asm.mov(rz, rAX) // nop if z == AX
	}
	asm.popRegs(tosave)
	return asm
}

// dst = - src
func (asm *Asm) Neg(dst Arg, src Arg) *Asm {
	return asm
}
