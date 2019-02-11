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

package arch

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
func (asm *Asm) Op2(op Op2, src Arg, dst Arg) *Asm {
	switch op {
	case CAST:
		if SizeOf(src) != SizeOf(dst) {
			return asm.Cast(src, dst)
		}
		fallthrough
	case MOV:
		return asm.Mov(src, dst)
	case NEG2, NOT2:
		return asm.op2(op, src, dst)
	default:
		// dst OP= src
		//    translates to
		// dst = dst OP src
		//    note the argument order
		return asm.Op3(Op3(op), dst, src, dst)
	}
}

func (asm *Asm) op2(op Op2, src Arg, dst Arg) *Asm {
	op.val() // validate op

	assert(src.Kind() == dst.Kind())
	if dst.Const() {
		errorf("destination cannot be a constant: %v %v, %v", op, src, dst)
	}

	switch src := src.(type) {
	case Reg:
		switch dst := dst.(type) {
		case Reg:
			asm.op2RegReg(op, src, dst)
		case Mem:
			r := asm.RegAlloc(dst.Kind())
			asm.op2RegReg(op, src, r).Store(r, dst).RegFree(r)
		default:
			errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v", dst, op, src, dst)
		}
	case Mem:
		switch dst := dst.(type) {
		case Reg:
			asm.Load(src, dst).op2RegReg(op, dst, dst)
		case Mem:
			r := asm.RegAlloc(dst.Kind())
			asm.Load(src, r).op2RegReg(op, r, r).Store(r, dst).RegFree(r)
		default:
			errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v", dst, op, src, dst)
		}
	case Const:
		if op == NEG2 {
			src.val = -src.val
		} else {
			src.val = ^src.val
		}
		return asm.Mov(src, dst)
	default:
		errorf("unknown argument type %T, expecting Const, Reg or Mem: %v %v, %v", src, op, src, dst)
	}

	return asm
}

func (asm *Asm) op2RegReg(op Op2, src Reg, dst Reg) *Asm {
	return asm.Uint32(dst.kind.kbit() | op.val() | src.val()<<16 | dst.val())
}
