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
 * amd64_mul.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

func (asm *Asm) mul2ConstReg(c Const, dst Reg) *Asm {
	n := c.val
	switch n {
	case 0:
		return asm.Op2(XOR, dst, dst)
	case 1:
		return asm
	case 3, 5, 9:
		return asm.lea4(MakeMem(0, dst.id, dst.kind), dst, n-1, dst)
	}
	if n&(n-1) == 0 {
		// TODO shift
	}
	// TODO
	errorf("unimplemented: MUL %v %v", c, dst)
	return asm
}

func (asm *Asm) mul2RegReg(src Reg, dst Reg) *Asm {
	// TODO
	errorf("unimplemented: MUL %v %v", src, dst)
	return asm
}

func (asm *Asm) mul2MemReg(src_m Mem, dst Reg) *Asm {
	// TODO
	errorf("unimplemented: MUL %v %v", src_m, dst)
	return asm
}

func (asm *Asm) mul2ConstMem(c Const, m Mem) *Asm {
	switch c.val {
	case -1:
		return asm.op1Mem(NEG, m)
	case 0:
		return asm.movConstMem(c, m)
	case 1:
		return asm
	default:
		r, allocated := asm.AllocLoad(m)
		return asm.mul2ConstReg(c, r).StoreFree(r, allocated, m)
	}
}

func (asm *Asm) mul2RegMem(src Reg, dst_m Mem) *Asm {
	// TODO
	errorf("unimplemented: MUL %v %v", src, dst_m)
	return asm
}

func (asm *Asm) mul2MemMem(src_m Mem, dst_m Mem) *Asm {
	return asm
}
