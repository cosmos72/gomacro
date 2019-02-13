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
 * op2.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arm64

// ============================================================================
// two-arg instruction

func (op Op2) val() uint32 {
	var val uint32
	switch op {
	case NEG2:
		val = 0x4B0003E0
	case NOT2:
		val = 0x2A2003E0
	default:
		errorf("unknown Op2 instruction: %v", op)
	}
	return val
}

// ============================================================================
func (arch Arm64) Op2(asm *Asm, op Op2, src Arg, dst Arg) {
	arch.op2(asm, op, src, dst)
}

func (arch Arm64) op2(asm *Asm, op Op2, src Arg, dst Arg) Arm64 {
	switch op {
	case CAST:
		if SizeOf(src) != SizeOf(dst) {
			return arch.cast(asm, src, dst)
		}
		fallthrough
	case MOV:
		return arch.mov(asm, src, dst)
	case NEG2, NOT2:
		break
	default:
		// dst OP= src
		//    translates to
		// dst = dst OP src
		//    note the argument order
		return arch.op3(asm, Op3(op), dst, src, dst)
	}

	op.val() // validate op

	assert(src.Kind() == dst.Kind())
	if dst.Const() {
		errorf("destination cannot be a constant: %v %v, %v", op, src, dst)
	}

	switch src := src.(type) {
	case Reg:
		switch dst := dst.(type) {
		case Reg:
			arch.op2RegReg(asm, op, src, dst)
		case Mem:
			r := asm.RegAlloc(dst.Kind())
			arch.op2RegReg(asm, op, src, r).store(asm, r, dst)
			asm.RegFree(r)
		default:
			errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v", dst, op, src, dst)
		}
	case Mem:
		switch dst := dst.(type) {
		case Reg:
			arch.load(asm, src, dst).op2RegReg(asm, op, dst, dst)
		case Mem:
			r := asm.RegAlloc(dst.Kind())
			arch.load(asm, src, r).op2RegReg(asm, op, r, r).store(asm, r, dst)
			asm.RegFree(r)
		default:
			errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v", dst, op, src, dst)
		}
	case Const:
		if op == NEG2 {
			src.val = -src.val
		} else {
			src.val = ^src.val
		}
		return arch.mov(asm, src, dst)
	default:
		errorf("unknown argument type %T, expecting Const, Reg or Mem: %v %v, %v", src, op, src, dst)
	}
	return arch
}

func (arch Arm64) op2RegReg(asm *Asm, op Op2, src Reg, dst Reg) Arm64 {
	asm.Uint32(dst.kind.arm64_kbit() | op.val() | src.arm64_val()<<16 | dst.arm64_val())
	return arch
}
