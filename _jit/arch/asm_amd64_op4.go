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
 * asm_amd64_op4.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

func (asm *Asm) Op4(op Op4, a Arg, b Arg, c Arg, d Arg) *Asm {

	dst := a.(Reg)
	assert(SizeOf(dst) == 8)
	m := b.(Mem)
	assert(SizeOf(m) == 8)

	var reg Reg
	var scale int64
	if c != nil {
		assert(SizeOf(c) == 8)
		reg = c.(Reg)
	}
	if d != nil {
		assert(SizeOf(d) == 8)
		scale = d.(Const).val
	}
	if reg.id == NoRegId || scale == 0 {
		return asm.op2RegMem(LEA, dst, m)
	}
	if reg.id == RSP {
		errorf("LEA: register RSP cannot be scaled: %v %v %v %v %v",
			op, dst, m, reg, scale)
	}
	var scalebit uint8
	switch scale {
	case 1:
		scalebit = 0
	case 2:
		scalebit = 0x40
	case 4:
		scalebit = 0x80
	case 8:
		scalebit = 0xC0
	default:
		errorf("LEA: unsupported scale %v, expecting 1,2,4 or 8: %v %v %v %v %v",
			op, dst, m, reg, scale)
	}
	dlo, dhi := dst.lohi()
	mlo, mhi := m.reg.lohi()
	rlo, rhi := reg.lohi()
	offlen, offbit := m.offlen(m.reg.id)

	asm.Bytes(0x48|dhi<<2|rhi<<1|mhi, uint8(op), offbit|0x04|dlo<<3, scalebit|rlo<<3|mlo)
	switch offlen {
	case 1:
		asm.Int8(int8(m.off))
	case 4:
		asm.Int32(m.off)
	}
	return asm
}
