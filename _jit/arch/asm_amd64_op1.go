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
 * asm_amd64_op1.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

func (asm *Asm) Op1(op Op1, dst Arg) *Asm {
	switch dst := dst.(type) {
	case Reg:
		asm.Op1Reg(op, dst)
	case Mem:
		asm.Op1Mem(op, dst)
	case Const:
		panicf("destination cannot be a constant: %v %v", op, dst)
	default:
		panicf("unsupported destination type, expecting Reg or Mem: %v %v", op, dst)
	}
	return asm
}

// unary operation OP, either NOT or NEG
// OP %reg_dst
func (asm *Asm) Op1Reg(op Op1, dst Reg) *Asm {
	dlo, dhi := dst.lohi()

	return asm.Bytes(0x48|dhi, 0xF7, 0xC0|uint8(op)|dlo)
}

// unary operation OP, either NOT or NEG
// OP off_m(%reg_m)
func (asm *Asm) Op1Mem(op Op1, m Mem) *Asm {

	dst := m.reg
	dlo, dhi := dst.lohi()
	op_ := uint8(op)

	var offlen, moff uint8
	switch {
	// (%rbp) and (%r13) sources must use 1-byte offset even if m.off == 0
	case m.off == 0 && dst.id != RBP && dst.id != R13:
		offlen = 0
		moff = 0
	case m.off == int32(int8(m.off)):
		offlen = 1
		moff = 0x40
	default:
		offlen = 4
		moff = 0x80
	}

	switch SizeOf(m) {
	case 1:
		if dhi != 0 {
			asm.Bytes(0x41)
		}
		asm.Bytes(0xF6, op_|moff|dlo)
	case 2:
		asm.Bytes(0x66)
		fallthrough
	case 4:
		if dhi != 0 {
			asm.Bytes(0x41)
		}
		asm.Bytes(0xF7, op_|moff|dlo)
	case 8:
		asm.Bytes(0x48|dhi, 0xF7, op_|moff|dlo)
	default:
		panicf("SizeOf(m) must be 1,2,4 or 8, found: %v", m)
	}
	if dst.id == RSP || dst.id == R12 {
		asm.Bytes(0x24) // amd64 quirk
	}
	switch offlen {
	case 1:
		asm.Int8(int8(m.off))
	case 4:
		asm.Int32(m.off)
	}
	return asm
}
