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
 * asm_amd64_shift.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

// %reg_dst SHIFT= const
func (asm *Asm) shiftRegConst(op Op2, dst Reg, c Const) *Asm {
	errorf("unimplemented: %v %v %v", op, dst, c)
	return asm
}

// off_dst(%reg_dst) SHIFT= const
func (asm *Asm) shiftMemConst(op Op2, m Mem, c Const) *Asm {
	errorf("unimplemented: %v %v %v", op, m, c)
	return asm
}

// %reg_dst SHIFT= %reg_src
func (asm *Asm) shiftRegReg(op Op2, dst Reg, src Reg) *Asm {
	errorf("unimplemented: %v %v %v", op, dst, src)
	return asm
}

// %reg_dst SHIFT= off_src(%reg_src)
func (asm *Asm) shiftRegMem(op Op2, dst Reg, m Mem) *Asm {
	errorf("unimplemented: %v %v %v", op, dst, m)
	return asm
}

// off_dst(%reg_dst) SHIFT= %reg_src
func (asm *Asm) shiftMemReg(op Op2, m Mem, src Reg) *Asm {
	errorf("unimplemented: %v %v %v", op, m, src)
	return asm
}

// off_dst(%reg_dst) SHIFT= off_src(%reg_src)
func (asm *Asm) shiftMemMem(op Op2, dst Mem, src Mem) *Asm {
	errorf("unimplemented: %v %v %v", op, dst, src)
	return asm
}
