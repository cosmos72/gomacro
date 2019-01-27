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
 * asm_amd64_mul.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

func (asm *Asm) mul2RegConst(dst Reg, c Const) *Asm {
	n := c.val
	switch n {
	case 0:
		return asm.Op2(XOR, dst, dst)
	case 1:
		return asm
	case 3, 5, 9:
		return asm.lea4(dst, MakeMem(0, dst.id, dst.kind), dst, n-1)
	}
	if n&(n-1) == 0 {
		// TODO shift
	}
	// TODO
	errorf("unimplemented: MUL %v %v", dst, c)
	return asm
}

func (asm *Asm) mul2RegReg(dst Reg, src Reg) *Asm {
	// TODO
	errorf("unimplemented: MUL %v %v", dst, src)
	return asm
}

func (asm *Asm) mul2RegMem(dst Reg, m Mem) *Asm {
	// TODO
	errorf("unimplemented: MUL %v %v", dst, m)
	return asm
}

func (asm *Asm) mul2MemConst(m Mem, c Const) *Asm {
	switch c.val {
	case -1:
		return asm.op1Mem(NEG, m)
	case 0:
		return asm.movMemConst(m, c)
	case 1:
		return asm
	default:
		r, allocated := asm.AllocLoad(m)
		return asm.mul2RegConst(r, c).StoreFree(m, r, allocated)
	}
}

func (asm *Asm) mul2MemReg(m Mem, src Reg) *Asm {
	// TODO
	return asm
}

func (asm *Asm) mul2MemMem(dst Mem, src Mem) *Asm {
	return asm
}
