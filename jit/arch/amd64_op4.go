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
 * amd64_op4.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// ============================================================================
// quaternary operation
type Op4 uint8

const (
	LEA4 Op4 = 0x8D
)

func (op Op4) String() string {
	s := "LEA4"
	if op != LEA4 {
		s = fmt.Sprintf("Op4(%d)", int(op))
	}
	return s
}

// ============================================================================
func (asm *Asm) Op4(op Op4, a Arg, b Arg, c Arg, dst Arg) *Asm {
	assert(op == LEA4)

	src_m := a.(Mem)
	var reg Reg
	var scale int64
	if b != nil {
		reg = b.(Reg)
	}
	if c != nil {
		assert(SizeOf(c) == 8)
		scale = c.(Const).val
	}
	dreg := dst.(Reg)

	if reg.id == NoRegId || scale == 0 {
		return asm.op2MemReg(LEA, src_m, dreg)
	} else if src_m.reg.id == NoRegId && scale == 1 {
		return asm.op2MemReg(LEA, MakeMem(src_m.off, reg.id, src_m.reg.kind), dreg)
	}
	return asm.lea4(src_m, reg, scale, dreg)
}

func (asm *Asm) lea4(m Mem, reg Reg, scale int64, dst Reg) *Asm {
	op := LEA4
	assert(SizeOf(dst) == 8)
	assert(SizeOf(m) == 8)
	assert(SizeOf(reg) == 8)
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
			op, m, reg, scale, dst)
	}
	dlo, dhi := dst.lohi()
	var mlo, mhi uint8
	var offlen, offbit uint8
	if m.reg.Valid() {
		mlo, mhi = m.reg.lohi()
		offlen, offbit = m.offlen(m.reg.id)
	} else {
		// no mem register
		offlen = 4
		scalebit |= 0x05
	}
	if reg.id == RSP {
		errorf("LEA: register RSP cannot be scaled: %v %v %v %v %v",
			op, m, reg, scale, dst)
	}
	rlo, rhi := reg.lohi()

	asm.Bytes(0x48|dhi<<2|rhi<<1|mhi, uint8(op), offbit|0x04|dlo<<3, scalebit|rlo<<3|mlo)
	switch offlen {
	case 1:
		asm.Int8(int8(m.off))
	case 4:
		asm.Int32(m.off)
	}
	return asm
}
