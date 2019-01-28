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
 * asm_amd64_op2.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

func (asm *Asm) Op2(op Op2, dst Arg, src Arg) *Asm {
	if op == CAST {
		if dst.Kind() != src.Kind() {
			return asm.Cast(dst, src)
		}
		op = MOV
	}
	if op != SHL && op != SHR {
		assert(dst.Kind() == src.Kind())
	}
	if asm.optimize(op, dst, src) {
		return asm
	}
	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Reg:
			asm.op2RegReg(op, dst, src)
		case Mem:
			asm.op2RegMem(op, dst, src)
		case Const:
			asm.op2RegConst(op, dst, src)
		default:
			errorf("unsupported source type %T, expecting Reg, Mem or Const: %v %v %v", op, dst, src)
		}
	case Mem:
		switch src := src.(type) {
		case Reg:
			asm.op2MemReg(op, dst, src)
		case Mem:
			asm.op2MemMem(op, dst, src)
		case Const:
			asm.op2MemConst(op, dst, src)
		default:
			errorf("unsupported source type %T, expecting Reg, Mem or Const: %v %v %v", src, op, dst, src)
		}
	case Const:
		errorf("destination cannot be a constant: %v %v %v", op, dst, src)
	default:
		errorf("unsupported destination type %T, expecting Reg or Mem: %v %v %v", dst, op, dst, src)
	}
	return asm
}

// %reg_dst OP= const
func (asm *Asm) op2RegConst(op Op2, dst Reg, src Const) *Asm {
	switch op {
	case MOV:
		return asm.movRegConst(dst, src)
	case SHL, SHR:
		return asm.shiftRegConst(op, dst, src)
	case MUL:
		return asm.mul2RegConst(dst, src)
	}
	assert(op != LEA)

	dlo, dhi := dst.lohi()
	op_ := uint8(op)
	c := src.val
	if c == int64(int8(c)) {
		asm.Bytes(0x48|dhi, 0x83, 0xC0|op_|dlo, uint8(int8(c)))
	} else if c == int64(int32(c)) {
		if dst.id == RAX {
			asm.Bytes(0x48|dhi, 0x05|op_).Int32(int32(c))
		} else {
			asm.Bytes(0x48|dhi, 0x81, 0xC0|op_|dlo).Int32(int32(c))
		}
	} else {
		// constant is 64 bit wide, must load it in a register
		r := asm.RegAlloc(src.kind)
		asm.movRegConst(r, src)
		asm.op2RegReg(op, dst, r)
		asm.RegFree(r)
	}
	return asm
}

// %reg_dst OP= %reg_src
func (asm *Asm) op2RegReg(op Op2, dst Reg, src Reg) *Asm {
	switch op {
	case MUL:
		return asm.mul2RegReg(dst, src)
	case SHL, SHR:
		return asm.shiftRegReg(op, dst, src)
	}
	assert(op != LEA)

	dlo, dhi := dst.lohi()
	slo, shi := src.lohi()

	switch SizeOf(dst) { // == SizeOf(src)
	case 1:
		if dst.id < RSP && src.id < RSP {
			asm.Bytes(uint8(op), 0xC0|dlo|slo<<3)
		} else {
			asm.Bytes(0x40|dhi|shi<<2, uint8(op), 0xC0|dlo|slo<<3)
		}
	case 2:
		if dhi|shi<<2 == 0 {
			asm.Bytes(0x66, 0x01|uint8(op), 0xC0|dlo|slo<<3)
		} else {
			asm.Bytes(0x66, 0x40|dhi|shi<<2, 0x01|uint8(op), 0xC0|dlo|slo<<3)
		}
	case 4:
		if dhi|shi<<2 == 0 {
			asm.Bytes(0x01|uint8(op), 0xC0|dlo|slo<<3)
		} else {
			asm.Bytes(0x40|dhi|shi<<2, 0x01|uint8(op), 0xC0|dlo|slo<<3)
		}
	case 8:
		asm.Bytes(0x48|dhi|shi<<2, 0x01|uint8(op), 0xC0|dlo|slo<<3)
	}
	return asm
}

// off_m(%reg_m) OP= %reg_src
func (asm *Asm) op2MemReg(op Op2, m Mem, src Reg) *Asm {
	switch op {
	case MUL:
		return asm.mul2MemReg(m, src)
	case SHL, SHR:
		return asm.shiftMemReg(op, m, src)
	}
	assert(op != LEA)

	dst := m.reg
	dlo, dhi := dst.lohi()
	slo, shi := src.lohi()

	assert(SizeOf(m) == SizeOf(dst))
	siz := SizeOf(dst)
	offlen, offbit := m.offlen(dst.id)

	switch siz {
	case 1:
		if src.id < RSP && dhi == 0 {
			asm.Bytes(uint8(op), offbit|dlo|slo<<3)
		} else {
			asm.Bytes(0x40|dhi|shi<<2, uint8(op), offbit|dlo|slo<<3)
		}
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi|shi<<2 == 0 {
			asm.Bytes(0x01|uint8(op), offbit|dlo|slo<<3)
		} else {
			asm.Bytes(0x40|dhi|shi<<2, 0x01|uint8(op), offbit|dlo|slo<<3)
		}
	case 8:
		asm.Bytes(0x48|dhi|shi<<2, 0x01|uint8(op), offbit|dlo|slo<<3)
	}
	asm.quirk24(dst)
	switch offlen {
	case 1:
		asm.Int8(int8(m.off))
	case 4:
		asm.Int32(m.off)
	}
	return asm
}

// %reg_dst OP= off_m(%reg_m)
func (asm *Asm) op2RegMem(op Op2, dst Reg, m Mem) *Asm {
	switch op {
	case MUL:
		return asm.mul2RegMem(dst, m)
	case SHL, SHR:
		return asm.shiftRegMem(op, dst, m)
	}
	src := m.reg
	dlo, dhi := dst.lohi()
	slo, shi := src.lohi()

	assert(SizeOf(src) == SizeOf(dst))
	siz := SizeOf(src)
	offlen, offbit := m.offlen(src.id)

	if op == LEA {
		assert(siz == 8)
	}

	switch siz {
	case 1:
		if dst.id < RSP && shi == 0 {
			asm.Bytes(0x02|uint8(op), offbit|dlo<<3|slo)
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x02|uint8(op), offbit|dlo<<3|slo)
		}
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if dhi|shi<<2 == 0 {
			asm.Bytes(0x03|uint8(op), offbit|dlo<<3|slo)
		} else {
			asm.Bytes(0x40|dhi<<2|shi, 0x03|uint8(op), offbit|dlo<<3|slo)
		}
	case 8:
		op_ := uint8(op)
		if op != LEA {
			op_ |= 0x03
		}
		asm.Bytes(0x48|dhi<<2|shi, op_, offbit|dlo<<3|slo)
	}
	asm.quirk24(src)
	switch offlen {
	case 1:
		asm.Int8(int8(m.off))
	case 4:
		asm.Int32(m.off)
	}
	return asm
}

// off_dst(%reg_dst) OP= off_src(%reg_src)
func (asm *Asm) op2MemMem(op Op2, dst Mem, src Mem) *Asm {
	switch op {
	case MUL:
		return asm.mul2MemMem(dst, src)
	case SHL, SHR:
		return asm.shiftMemMem(op, dst, src)
	}
	assert(op != LEA)
	// not natively supported by amd64,
	// must load src in a register
	r := asm.RegAlloc(src.Kind())
	asm.op2RegMem(MOV, r, src)
	asm.op2MemReg(op, dst, r)
	return asm.RegFree(r)
}

// off_dst(%reg_dst) OP= const
func (asm *Asm) op2MemConst(op Op2, dst Mem, src Const) *Asm {
	switch op {
	case MOV:
		return asm.movMemConst(dst, src)
	case SHL, SHR:
		return asm.shiftMemConst(op, dst, src)
	case MUL:
		return asm.mul2MemConst(dst, src)
	}
	assert(op != LEA)
	// not natively supported by amd64,
	// must load src in a register
	r := asm.RegAlloc(src.kind)
	asm.movRegConst(r, src)
	asm.op2MemReg(op, dst, r)
	return asm.RegFree(r)
}

func (asm *Asm) optimize(op Op2, dst Arg, src Arg) bool {
	if dst == src {
		switch op {
		case AND, OR, MOV, CAST:
			return true // operation is nop
		case SUB, XOR:
			asm.Op2(MOV, dst, Const{val: 0, kind: dst.Kind()})
			return true
		}
	}
	c, ok := src.(Const)
	if !ok {
		return false
	}
	n := c.Cast(Int64).val
	src = Const{val: n, kind: dst.Kind()}
	switch op {
	case ADD:
		switch n {
		case 0:
			return true
		case 1:
			asm.Op1(INC, dst)
			return true
		case -1:
			asm.Op1(DEC, dst)
			return true
		}
	case OR:
		switch n {
		case 0:
			return true
		case -1:
			asm.Op2(MOV, dst, src)
			return true
		}
	case AND:
		switch n {
		case 0:
			asm.Op2(MOV, dst, src)
			return true
		case -1:
			return true
		}
	case SUB:
		switch n {
		case 0:
			return true
		case 1:
			asm.Op1(DEC, dst)
			return true
		case -1:
			asm.Op1(INC, dst)
			return true
		}
	case XOR:
		switch n {
		case 0:
			return true
		case -1:
			asm.Op1(NOT, dst)
			return true
		}
	case CAST:
		asm.Op2(MOV, dst, src)
		return true
	case SHL, SHR:
		switch n {
		case 0:
			return true
		}
	case MUL:
		switch n {
		case 0:
			asm.Op2(MOV, dst, src)
			return true
		case 1:
			return true
		}
	}
	return false
}
