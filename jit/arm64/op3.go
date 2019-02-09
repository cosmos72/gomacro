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
 * op3.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// ============================================================================
// ternary operation
type Op3 uint8

const (
	AND3 Op3 = 0x0A
	ADD3 Op3 = 0x0B
	ADC3 Op3 = 0x1A // add with carry
	OR3  Op3 = 0x2A
	XOR3 Op3 = 0x4A
	SUB3 Op3 = 0x4B
	SBB3 Op3 = 0x5A // subtract with borrow

	SHL3 Op3 = 0x30 // shift left
	SHR3 Op3 = 0x31 // shift right
	MUL3 Op3 = 0x32
	DIV3 Op3 = 0x33
	REM3 Op3 = 0x34
)

var op3Name = map[Op3]string{
	ADD3: "ADD3",
	OR3:  "OR3",
	ADC3: "ADC3",
	SBB3: "SBB3",
	AND3: "AND3",
	SUB3: "SUB3",
	XOR3: "XOR3",

	SHL3: "SHL3",
	SHR3: "SHR3",
	MUL3: "MUL3",
	DIV3: "DIV3",
	REM3: "REM3",
}

func (op Op3) String() string {
	s, ok := op3Name[op]
	if !ok {
		s = fmt.Sprintf("Op3(%d)", int(op))
	}
	return s
}

// return 32bit value used to encode operation on Reg,Reg,Reg
func (op Op3) val() uint32 {
	switch op {
	case SHL3:
		return 0x1AC02000
	case SHR3:
		// logical i.e. zero-extended right shift is 0x1AC02400
		// arithmetic i.e. sign-extended right shift is 0x1AC02800
		return 0x1AC02400
	case MUL3:
		// 0x1B007C00 because MUL3 a,b,c is an alias for MADD4 xzr,a,b,c
		return 0x1B007C00
	case DIV3:
		// unsigned division is 0x1AC00800
		// signed division is 0x1AC00C00
		return 0x1AC00800
	case REM3:
		errorf("internal error, operation %v needs to be implemented as {s|u}div followed by msub", op)
		return 0
	default:
		return uint32(op) << 24
	}
}

// return 32bit value used to encode operation on Reg,Const,Reg
func (op Op3) immval() uint32 {
	switch op {
	case AND3:
		return 0x12 << 24
	case ADD3:
		return 0x11 << 24
	case SHL3, SHR3:
		// immediate constant is encoded differently
		return 0x53 << 24
	case OR3:
		return 0x32 << 24
	case XOR3:
		return 0x52 << 24
	case SUB3:
		return 0x51 << 24
	default:
		errorf("cannot encode %v with immediate constant", op)
		return 0
	}
}

// ============================================================================
func (asm *Asm) Op3(op Op3, a Arg, b Arg, dst Arg) *Asm {
	// validate kinds
	assert(a.Kind() == dst.Kind())
	switch op {
	case SHL3, SHR3:
		assert(!b.Kind().Signed())
	default:
		assert(b.Kind() == dst.Kind())
	}
	// validate dst
	switch dst.(type) {
	case Reg, Mem:
		break
	case Const:
		errorf("destination cannot be a constant: %v %v, %v, %v", op, a, b, dst)
	default:
		errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v, %v", dst, op, a, b, dst)
	}

	if asm.optimize3(op, a, b, dst) {
		return asm
	}
	var ra, rb, rdst Reg
	var ta, tdst bool // Reg is a temporary register?

	switch dst := dst.(type) {
	case Reg:
		rdst = dst
	case Mem:
		rdst = asm.RegAlloc(dst.Kind())
		defer asm.RegFree(rdst)
		tdst = true
	}
	if op.isCommutative() && a.Const() && !b.Const() {
		a, b = b, a
	}
	switch xa := a.(type) {
	case Reg:
		ra = xa
	case Mem:
		if tdst {
			// reuse temporary register rdst
			ra = rdst
		} else {
			ra = asm.RegAlloc(xa.Kind())
			defer asm.RegFree(ra)
		}
		ta = true
		asm.Load(xa, ra)
	case Const:
		ra = asm.RegAlloc(xa.kind)
		defer asm.RegFree(ra)
		asm.movConstReg(xa, ra)
	default:
		errorf("unknown argument type %T, expecting Const, Reg or Mem: %v %v, %v, %v", a, op, a, b, dst)
	}
	switch xb := b.(type) {
	case Reg:
		asm.op3RegRegReg(op, ra, xb, rdst)
	case Mem:
		if tdst && (!ta || ra != rdst) {
			// reuse temporary register rdst
			rb = rdst
		} else {
			rb = asm.RegAlloc(xb.Kind())
			defer asm.RegFree(rb)
		}
		asm.Load(xb, rb)
		asm.op3RegRegReg(op, ra, rb, rdst)
	case Const:
		asm.op3RegConstReg(op, ra, xb, rdst)
	default:
		errorf("unknown argument type %T, expecting Const, Reg or Mem: %v %v, %v, %v", b, op, a, b, dst)
	}
	if tdst {
		asm.Store(rdst, dst.(Mem))
	}
	return asm
}

func (asm *Asm) op3RegRegReg(op Op3, a Reg, b Reg, dst Reg) *Asm {
	var opbits uint32
	if dst.kind.Signed() {
		switch op {
		case SHR3:
			// arithmetic right shift
			opbits = 0xC00
		case DIV3:
			// signed division
			opbits = 0xC00
		}
	}
	asm.extendHighBits(op, a)
	asm.extendHighBits(op, b)
	// TODO: on arm64, division by zero returns zero instead of panic
	asm.Uint32(dst.kind.kbit() | (opbits ^ op.val()) | b.val()<<16 | a.val()<<5 | dst.val())
	return asm
}

func (asm *Asm) op3RegConstReg(op Op3, a Reg, cb Const, dst Reg) *Asm {
	if asm.tryOp3RegConstReg(op, a, uint64(cb.val), dst) {
		return asm
	}
	rb := asm.RegAlloc(cb.kind)
	return asm.movConstReg(cb, rb).op3RegRegReg(op, a, rb, dst).RegFree(rb)
}

// try to encode operation into a single instruction.
// return false if not possible because constant must be loaded in a register
func (asm *Asm) tryOp3RegConstReg(op Op3, a Reg, cval uint64, dst Reg) bool {
	imm3 := op.immediate()
	immcval, ok := imm3.Encode64(cval, dst.Kind())
	if !ok {
		return false
	}
	opval := op.immval()

	kbit := dst.kind.kbit()

	asm.extendHighBits(op, a)
	switch imm3 {
	case Imm3AddSub, Imm3Bitwise:
		// for op == OR3, also accept a == XZR
		asm.Uint32(kbit | opval | immcval | a.valOrX31(op == OR3)<<5 | dst.val())
	case Imm3Shift:
		asm.shiftRegConstReg(op, a, cval, dst)
	default:
		cb := ConstInt64(int64(cval))
		errorf("unknown constant encoding style %v for %v: %v %v, %v, %v", imm3, op, op, a, cb, dst)
	}
	return true
}

func (asm *Asm) shiftRegConstReg(op Op3, a Reg, cval uint64, dst Reg) {
	dsize := dst.kind.Size()
	if cval >= 8*uint64(dsize) {
		cb := ConstInt64(int64(cval))
		errorf("constant is out of range for shift: %v %v, %v, %v", op, a, cb, dst)
	}
	switch op {
	case SHL3:
		switch dsize {
		case 1, 2, 4:
			asm.Uint32(0x53000000 | uint32(32-cval)<<16 | uint32(31-cval)<<10 | a.val()<<5 | dst.val())
		case 8:
			asm.Uint32(0xD3400000 | uint32(64-cval)<<16 | uint32(63-cval)<<10 | a.val()<<5 | dst.val())
		}
	case SHR3:
		var unsignedbit uint32
		if !dst.kind.Signed() {
			unsignedbit = 0x40 << 24
		}
		switch dsize {
		case 1, 2, 4:
			asm.Uint32(unsignedbit | 0x13007C00 | uint32(cval)<<16 | a.val()<<5 | dst.val())
		case 8:
			asm.Uint32(unsignedbit | 0x9340FC00 | uint32(cval)<<16 | a.val()<<5 | dst.val())
		}
	}
}

// arm64 has no native operations to work on 8 bit and 16 bit registers.
// Actually, it only has ldr (load) and str (store), but no arithmetic
// or bitwise operations.
// So we emulate them similarly to what compilers do:
// use 32 bit registers and ignore high bits in operands and results.
// Exception: right-shift, division and remainder move data
// from high bits to low bits, so we must zero-extend or sign-extend
// the operands
func (asm *Asm) extendHighBits(op Op3, r Reg) {
	rkind := r.kind
	rsize := rkind.Size()
	if rsize > 2 {
		return
	}
	switch op {
	case SHR3, DIV3, REM3:
		if rkind.Signed() {
			asm.Cast(r, MakeReg(r.id, Int32))
		} else {
			asm.Cast(r, MakeReg(r.id, Uint32))
		}
	}
}

func (op Op3) isCommutative() bool {
	switch op {
	case ADD3, OR3, ADC3, AND3, XOR3, MUL3:
		return true
	}
	return false
}

func (asm *Asm) optimize3(op Op3, a Arg, b Arg, dst Arg) bool {
	if a == b {
		switch op {
		case AND3, OR3:
			if b == dst {
				// operation is NOP
				return true
			}
			asm.Mov(a, dst)
			return true
		case SUB3, XOR3:
			asm.Zero(dst)
			return true
		}
	}
	c, ok := b.(Const)
	if !ok {
		if op.isCommutative() {
			a, b = b, a
			c, ok = b.(Const)
		}
		if !ok {
			return false
		}
	}
	n := c.Cast(Int64).val
	c = MakeConst(n, dst.Kind())
	switch op {
	case ADD3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		}
	case OR3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		case -1:
			asm.Mov(c, dst)
			return true
		}
	case AND3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		case -1:
			asm.Mov(c, dst)
			return true
		}
	case SUB3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		}
	case XOR3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		case -1:
			asm.Op2(NOT2, a, dst)
			return true
		}
	case SHL3, SHR3:
		switch n {
		case 0:
			asm.Mov(a, dst)
			return true
		}
	case MUL3:
		switch n {
		case 0:
			asm.Zero(dst)
			return true
		case 1:
			asm.Mov(a, dst)
			return true
		case -1:
			asm.Op2(NEG2, a, dst)
			return true
		}
	case DIV3:
		switch n {
		case 1:
			asm.Mov(a, dst)
			return true
		case -1:
			asm.Op2(NEG2, a, dst)
			return true
		}
	}
	return false
}

// ============================================================================

// style of immediate constants
// embeddable in a single Op3 instruction
type Immediate3 uint8

const (
	Imm3None    Immediate3 = iota
	Imm3AddSub             // 12 bits wide, possibly shifted left by 12 bits
	Imm3Bitwise            // complicated
	Imm3Shift              // 0..63 for 64 bit registers; 0..31 for 32 bit registers
)

// return the style of immediate constants
// embeddable in a single Op3 instruction
func (op Op3) immediate() Immediate3 {
	switch op {
	case ADD3, SUB3:
		return Imm3AddSub
	case AND3, OR3, XOR3:
		return Imm3Bitwise
	case SHL3, SHR3:
		return Imm3Shift
	default:
		return Imm3None
	}
}

// return false if val cannot be encoded using imm style
func (imm Immediate3) Encode64(val uint64, kind Kind) (e uint32, ok bool) {
	kbits := kind.Size() * 8
	switch imm {
	case Imm3AddSub:
		// 12 bits wide, possibly shifted left by 12 bits
		if val == val&0xFFF {
			return uint32(val << 10), true
		} else if val == val&0xFFF000 {
			return 0x400000 | uint32(val>>2), true
		}
	case Imm3Bitwise:
		// complicated
		if kbits <= 32 {
			e, ok = imm3Bitwise32[val]
		} else {
			e, ok = imm3Bitwise64[val]
		}
		return e, ok
	case Imm3Shift:
		if val >= 0 && val < uint64(kbits) {
			// actual encoding is complicated
			return uint32(val), true
		}
	}
	return 0, false
}

var imm3Bitwise32 = makeImm3Bitwise32()
var imm3Bitwise64 = makeImm3Bitwise64()

// compute all immediate constants that can be encoded
// in and, orr, eor on 32-bit registers
func makeImm3Bitwise32() map[uint64]uint32 {
	result := make(map[uint64]uint32)
	var bitmask uint64
	var size, length, e, rotation uint32
	for size = 2; size <= 32; size *= 2 {
		for length = 1; length < size; length++ {
			bitmask = 0xffffffff >> (32 - length)
			for e = size; e < 32; e *= 2 {
				bitmask |= bitmask << e
			}
			for rotation = 0; rotation < size; rotation++ {
				result[bitmask] = (size&64|rotation)<<16 | (0x7800*size)&0xF000 | (length-1)<<10
				bitmask = (bitmask >> 1) | (bitmask << 31)
			}
		}
	}
	return result
}

// compute all immediate constants that can be encoded
// in and, orr, eor on 64-bit registers
func makeImm3Bitwise64() map[uint64]uint32 {
	result := make(map[uint64]uint32)
	var bitmask uint64
	var size, length, e, rotation uint32
	for size = 2; size <= 64; size *= 2 {
		for length = 1; length < size; length++ {
			bitmask = 0xffffffffffffffff >> (64 - length)
			for e = size; e < 64; e *= 2 {
				bitmask |= bitmask << e
			}
			for rotation = 0; rotation < size; rotation++ {
				// #0x5555555555555555 => size=2, length=1, rotation=0 => 0x00f000
				// #0xaaaaaaaaaaaaaaaa => size=2, length=1, rotation=1 => 0x01f000
				// #0x1111111111111111 => size=4, length=1, rotation=0 => 0x00e000
				// #0x8888888888888888 => size=4, length=1, rotation=1 => 0x01e000
				// #0x4444444444444444 => size=4, length=1, rotation=2 => 0x02e000
				// #0x2222222222222222 => size=4, length=1, rotation=3 => 0x03e000
				// #0x3333333333333333 => size=4, length=2, rotation=0 => 0x00e400
				// #0x7777777777777777 => size=4, length=3, rotation=0 => 0x00e800
				// #0x0101010101010101 => size=8, length=1, rotation=0 => 0x00c000
				// #0x0303030303030303 => size=8, length=2, rotation=0 => 0x00c400
				// #0x0707070707070707 => size=8, length=3, rotation=0 => 0x00c800
				// #0x0f0f0f0f0f0f0f0f => size=8, length=4, rotation=0 => 0x00cc00
				// #0x1f1f1f1f1f1f1f1f => size=8, length=5, rotation=0 => 0x00d000
				// #0x3f3f3f3f3f3f3f3f => size=8, length=6, rotation=0 => 0x00d400
				// #0x7f7f7f7f7f7f7f7f => size=8, length=7, rotation=0 => 0x00d800
				// ...
				result[bitmask] = (size&64|rotation)<<16 | (0x7800*size)&0xF000 | (length-1)<<10
				bitmask = (bitmask >> 1) | (bitmask << 63)
			}
		}
	}
	return result
}
