/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
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

package amd64

// ============================================================================
// tree-arg instruction

func (arch Amd64) Op3(asm *Asm, op Op3, a Arg, b Arg, dst Arg) *Asm {
	arch.op3(asm, op, a, b, dst)
	return asm
}

func (arch Amd64) op3(asm *Asm, op Op3, a Arg, b Arg, dst Arg) Amd64 {
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
	if asm.Optimize3(op, a, b, dst) {
		return arch
	}
	if op == MUL3 {
		return arch.mul3(asm, a, b, dst)
	}
	op2 := Op2(op)
	if a == dst {
		arch.op2(asm, op2, b, dst)
	} else if op.IsCommutative() && b == dst {
		arch.op2(asm, op2, a, dst)
	} else if r, ok := dst.(Reg); ok && r.RegId() != b.RegId() {
		arch.mov(asm, a, dst).op2(asm, op2, b, dst)
	} else {
		r := asm.RegAlloc(dst.Kind())
		arch.mov(asm, a, r).op2(asm, op2, b, r).mov(asm, r, dst)
		asm.RegFree(r)
	}
	return arch
}

func (arch Amd64) mul3(asm *Asm, a Arg, b Arg, dst Arg) Amd64 {
	if a.Const() && !b.Const() {
		a, b = b, a
	}
	if a == dst {
		return arch.op2(asm, MUL, b, dst)
	} else if b == dst {
		return arch.op2(asm, MUL, a, dst)
	}
	rdst, rokdst := dst.(Reg)
	if !a.Const() && b.Const() {
		bval := b.(Const).Val()
		if bval == int64(int8(bval)) {
			// use amd64 3-argument multiplication
			if !rokdst {
				rdst = asm.RegAlloc(dst.Kind())
			}
			switch a := a.(type) {
			case Reg:
				arch.mul3RegConst8Reg(asm, a, int8(bval), rdst)
			case Mem:
				arch.mul3MemConst8Reg(asm, a, int8(bval), rdst)
			default:
				errorf("unknown argument type %T, expecting Const, Reg or Mem: %v %v, %v, %v", a, MUL3, a, b, dst)
			}
			if !rokdst {
				arch.store(asm, rdst, dst.(Mem))
				asm.RegFree(rdst)
			}
			return arch
		}
	}
	if rokdst && rdst.RegId() != b.RegId() {
		return arch.mov(asm, a, dst).op2(asm, MUL, b, dst)
	}
	r := asm.RegAlloc(dst.Kind())
	arch.mov(asm, a, r).op2(asm, MUL, b, r).mov(asm, r, dst)
	asm.RegFree(r)
	return arch
}
