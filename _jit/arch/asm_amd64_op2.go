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

import (
	"fmt"
)

// ============================================================================
// binary operation
type Op2 uint8

const (
	ADD Op2 = 0
	OR  Op2 = 0x08
	ADC Op2 = 0x10 // add with carry
	SBB Op2 = 0x18 // subtract with borrow
	AND Op2 = 0x20
	SUB Op2 = 0x28
	XOR Op2 = 0x30
	// CMP Op = 0x38 // compare, set flags
	// XCHG Op = 0x86 // exchange. xchg %reg, %reg has different encoding
	MOV  Op2 = 0x88
	LEA  Op2 = 0x8D
	CAST Op2 = 0xB6 // sign extend, zero extend or narrow
	SHL  Op2 = 0xE0 // shift left. has different encoding
	SHR  Op2 = 0xE8 // shift right. has different encoding
	MUL  Op2 = 0xF6
	DIV  Op2 = 0xFE // TODO divide
	REM  Op2 = 0xFF // TODO remainder
)

var op2Name = map[Op2]string{
	ADD: "ADD",
	OR:  "OR",
	ADC: "ADC",
	SBB: "SBB",
	AND: "AND",
	SUB: "SUB",
	XOR: "XOR",
	// CMP: "CMP",
	// XCHG: "XCHG",
	MOV:  "MOV",
	CAST: "CAST",
	MUL:  "MUL",
	DIV:  "DIV",
	REM:  "REM",
}

func (op Op2) String() string {
	s, ok := op2Name[op]
	if !ok {
		s = fmt.Sprintf("Op2(%d)", int(op))
	}
	return s
}

// ============================================================================
func (asm *Asm) Op2(op Op2, src Arg, dst Arg) *Asm {
	if op == CAST {
		if SizeOf(src) != SizeOf(dst) {
			return asm.Cast(src, dst)
		}
		op = MOV
	}
	if op == MOV {
		assert(SizeOf(src) == SizeOf(dst))
	} else if op != SHL && op != SHR {
		assert(src.Kind() == dst.Kind())
	}
	if asm.optimize(op, src, dst) {
		return asm
	}
	if op == DIV || op == REM {
		errorf("unimplemented operation %v: %v %v %v", op, op, src, dst)
	}
	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Reg:
			asm.op2RegReg(op, src, dst)
		case Mem:
			asm.op2MemReg(op, src, dst)
		case Const:
			asm.op2ConstReg(op, src, dst)
		default:
			errorf("unsupported source type %T, expecting Reg, Mem or Const: %v %v %v", op, src, dst)
		}
	case Mem:
		switch src := src.(type) {
		case Reg:
			asm.op2RegMem(op, src, dst)
		case Mem:
			asm.op2MemMem(op, src, dst)
		case Const:
			asm.op2ConstMem(op, src, dst)
		default:
			errorf("unsupported source type %T, expecting Reg, Mem or Const: %v %v %v", src, op, src, dst)
		}
	case Const:
		errorf("destination cannot be a constant: %v %v %v", op, src, dst)
	default:
		errorf("unsupported destination type %T, expecting Reg or Mem: %v %v %v", dst, op, src, dst)
	}
	return asm
}

// %reg_dst OP= const
func (asm *Asm) op2ConstReg(op Op2, c Const, dst Reg) *Asm {
	switch op {
	case MOV:
		return asm.movConstReg(c, dst)
	case SHL, SHR:
		return asm.shiftConstReg(op, c, dst)
	case MUL:
		return asm.mul2ConstReg(c, dst)
	}
	assert(op != LEA)

	switch dst.kind.Size() {
	case 1:
		asm.op2ConstReg8(op, c, dst)
	case 2:
		asm.op2ConstReg16(op, c, dst)
	case 4:
		asm.op2ConstReg32(op, c, dst)
	case 8:
		asm.op2ConstReg64(op, c, dst)
	}
	return asm
}

func (asm *Asm) op2ConstReg8(op Op2, c Const, dst Reg) *Asm {
	op_ := uint8(op)
	dlo, dhi := dst.lohi()
	val := c.val
	if val == int64(int8(val)) {
		if dst.id == RAX {
			asm.Bytes(0x04 | op_)
		} else if dst.id < RSP {
			asm.Bytes(0x80, 0xC0|op_|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0xC0|op_|dlo)
		}
		asm.Int8(int8(val))
	} else {
		errorf("sign-extended constant overflows 8-bit destination: %v %v %v", op, c, dst)
	}
	return asm
}

func (asm *Asm) op2ConstReg16(op Op2, c Const, dst Reg) *Asm {
	op_ := uint8(op)
	dlo, dhi := dst.lohi()
	val := c.val
	if val == int64(int8(val)) {
		if dhi == 0 {
			asm.Bytes(0x66, 0x83, 0xc0|op_|dlo)
		} else {
			asm.Bytes(0x66, 0x40|dhi, 0x83, 0xc0|op_|dlo)
		}
		asm.Int8(int8(val))
	} else if val == int64(int16(val)) {
		if dst.id == RAX {
			asm.Bytes(0x66, 0x05|op_)
		} else if dhi == 0 {
			asm.Bytes(0x66, 0x81, 0xc0|op_|dlo)
		} else {
			asm.Bytes(0x66, 0x40|dhi, 0x81, 0xc0|op_|dlo)
		}
		asm.Int16(int16(val))
	} else {
		errorf("sign-extended constant overflows 16-bit destination: %v %v %v", op, c, dst)
	}
	return asm
}

func (asm *Asm) op2ConstReg32(op Op2, c Const, dst Reg) *Asm {
	op_ := uint8(op)
	dlo, dhi := dst.lohi()
	val := c.val
	if val == int64(int8(val)) {
		if dhi == 0 {
			asm.Bytes(0x83, 0xc0|op_|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0x83, 0xc0|op_|dlo)
		}
		asm.Int8(int8(val))
	} else if val == int64(int32(val)) {
		if dst.id == RAX {
			asm.Bytes(0x05 | op_)
		} else if dhi == 0 {
			asm.Bytes(0x81, 0xc0|op_|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0x81, 0xc0|op_|dlo)
		}
		asm.Int32(int32(val))
	} else {
		errorf("sign-extended constant overflows 32-bit destination: %v %v %v", op, c, dst)
	}
	return asm
}

func (asm *Asm) op2ConstReg64(op Op2, c Const, dst Reg) *Asm {
	dlo, dhi := dst.lohi()
	op_ := uint8(op)
	val := c.val
	if val == int64(int8(val)) {
		asm.Bytes(0x48|dhi, 0x83, 0xC0|op_|dlo, uint8(int8(val)))
	} else if val == int64(int32(val)) {
		if dst.id == RAX {
			asm.Bytes(0x48|dhi, 0x05|op_)
		} else {
			asm.Bytes(0x48|dhi, 0x81, 0xC0|op_|dlo)
		}
		asm.Int32(int32(val))
	} else {
		// constant is 64 bit wide, must load it in a register
		r := asm.RegAlloc(c.kind)
		asm.movConstReg64(c, r)
		asm.op2RegReg(op, r, dst)
		asm.RegFree(r)
	}
	return asm
}

// %reg_dst OP= %reg_src
func (asm *Asm) op2RegReg(op Op2, src Reg, dst Reg) *Asm {
	switch op {
	case MUL:
		return asm.mul2RegReg(src, dst)
	case SHL, SHR:
		return asm.shiftRegReg(op, src, dst)
	}
	assert(op != LEA)

	slo, shi := src.lohi()
	dlo, dhi := dst.lohi()

	switch SizeOf(src) { // == SizeOf(dst)
	case 1:
		if src.id < RSP && dst.id < RSP {
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
func (asm *Asm) op2RegMem(op Op2, src Reg, dst_m Mem) *Asm {
	switch op {
	case MUL:
		return asm.mul2RegMem(src, dst_m)
	case SHL, SHR:
		return asm.shiftRegMem(op, src, dst_m)
	}
	assert(op != LEA)

	dst := dst_m.reg
	dlo, dhi := dst.lohi()
	slo, shi := src.lohi()

	assert(SizeOf(dst_m) == SizeOf(dst))
	siz := SizeOf(dst)
	offlen, offbit := dst_m.offlen(dst.id)

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
		asm.Int8(int8(dst_m.off))
	case 4:
		asm.Int32(dst_m.off)
	}
	return asm
}

// %reg_dst OP= off_m(%reg_m)
func (asm *Asm) op2MemReg(op Op2, src_m Mem, dst Reg) *Asm {
	switch op {
	case MUL:
		return asm.mul2MemReg(src_m, dst)
	case SHL, SHR:
		return asm.shiftMemReg(op, src_m, dst)
	}
	src := src_m.reg
	dlo, dhi := dst.lohi()
	slo, shi := src.lohi()

	assert(SizeOf(src) == SizeOf(dst))
	siz := SizeOf(src)
	offlen, offbit := src_m.offlen(src.id)

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
		asm.Int8(int8(src_m.off))
	case 4:
		asm.Int32(src_m.off)
	}
	return asm
}

// off_dst(%reg_dst) OP= off_src(%reg_src)
func (asm *Asm) op2MemMem(op Op2, src_m Mem, dst_m Mem) *Asm {
	switch op {
	case MUL:
		return asm.mul2MemMem(src_m, dst_m)
	case SHL, SHR:
		return asm.shiftMemMem(op, src_m, dst_m)
	}
	assert(op != LEA)
	// not natively supported by amd64,
	// must load src in a register
	r := asm.RegAlloc(src_m.Kind())
	asm.op2MemReg(MOV, src_m, r)
	asm.op2RegMem(op, r, dst_m)
	return asm.RegFree(r)
}

// off_dst(%reg_dst) OP= const
func (asm *Asm) op2ConstMem(op Op2, c Const, m Mem) *Asm {
	switch op {
	case MOV:
		return asm.movConstMem(c, m)
	case SHL, SHR:
		return asm.shiftConstMem(op, c, m)
	case MUL:
		return asm.mul2ConstMem(c, m)
	}
	assert(op != LEA)
	op_ := uint8(op)
	dst := m.reg
	dlo, dhi := dst.lohi()
	offlen, offbit := m.offlen(dst.id)
	val := c.val
	switch dst.kind.Size() {
	case 1:
		if dhi == 0 {
			asm.Bytes(0x80, offbit|op_|dlo)
		} else {
			asm.Bytes(0x40|dhi, 0x80, offbit|op_|dlo)
		}
	case 2:
		asm.Byte(0x66)
		fallthrough
	case 4:
		if val == int64(int8(val)) {
			if dhi == 0 {
				asm.Bytes(0x83, offbit|op_|dlo)
			} else {
				asm.Bytes(0x40|dhi, 0x83, offbit|op_|dlo)
			}
		} else {
			if dhi == 0 {
				asm.Bytes(0x81, offbit|op_|dlo)
			} else {
				asm.Bytes(0x40|dhi, 0x81, offbit|op_|dlo)
			}
		}
	case 8:
		if val == int64(int8(val)) {
			asm.Bytes(0x48|dhi, 0x83, offbit|op_|dlo)
		} else if val == int64(int32(val)) {
			asm.Bytes(0x48|dhi, 0x81, offbit|op_|dlo)
		} else {
			// not natively supported by amd64,
			// must allocate a register
			r := asm.RegAlloc(c.kind)
			asm.movConstReg64(c, r).op2RegMem(op, r, m)
			return asm.RegFree(r)
		}
	}
	asm.quirk24(dst)
	switch offlen {
	case 1:
		asm.Int8(int8(m.off))
	case 4:
		asm.Int32(m.off)
	}

	if val == int64(int8(val)) {
		asm.Int8(int8(val))
	} else if dst.kind.Size() == 2 {
		asm.Int16(int16(val))
	} else {
		asm.Int32(int32(val))
	}
	return asm
}

func (asm *Asm) optimize(op Op2, src Arg, dst Arg) bool {
	if src == dst {
		switch op {
		case AND, OR, MOV, CAST:
			return true // operation is nop
		case SUB, XOR:
			asm.Op2(MOV, MakeConst(0, dst.Kind()), dst)
			return true
		}
	}
	c, ok := src.(Const)
	if !ok {
		return false
	}
	n := c.Cast(Int64).val
	src = MakeConst(n, dst.Kind())
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
			asm.Op2(MOV, src, dst)
			return true
		}
	case AND:
		switch n {
		case 0:
			asm.Op2(MOV, src, dst)
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
		asm.Op2(MOV, src, dst)
		return true
	case SHL, SHR:
		switch n {
		case 0:
			return true
		}
	case MUL:
		switch n {
		case 0:
			asm.Op2(MOV, src, dst)
			return true
		case 1:
			return true
		case -1:
			asm.Op1(NEG, dst)
			return true
		}
	}
	return false
}
