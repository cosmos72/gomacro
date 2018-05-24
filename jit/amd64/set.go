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
 * var_set_value.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

func (asm *Asm) Mov(dst hwReg, r hwReg) *Asm {
	if dst == r {
		return asm
	}
	slo, shi := r.lohi()
	dlo, dhi := dst.lohi()
	return asm.Bytes(0x48|dhi|shi*4, 0x89, 0xc0+dlo+slo*8) //  movq   %reg_src,%reg_dst
}

func (asm *Asm) Store(dst *Var, r hwReg) *Asm {
	lo, hi := r.lohi()
	return asm.Bytes(0x48|hi*4, 0x89, 0x87|lo*8).Idx(dst) //   movq   %reg,dst(%rdi)
}

func (asm *Asm) Load(dst hwReg, a Arg) *Asm {
	switch a := a.(type) {
	case hwReg:
		return asm.Mov(dst, a)
	case *Const:
		return asm.LoadConst(dst, a.val)
	case *Var:
		lo, hi := dst.lohi()
		return asm.Bytes(0x48|hi*4, 0x8b, 0x87|lo*8).Idx(a) //   movq   src(%rdi),%reg
	default:
		errorf("invalid arg type: %#v // %T", a, a)
		return nil
	}
}

func (asm *Asm) LoadConst(r hwReg, val int64) *Asm {
	lo, hi := r.lohi()
	if val == int64(uint32(val)) {
		if hi != 0 {
			asm.Bytes(0x41)
		}
		return asm.Bytes(0xb8 + lo).Uint32(uint32(val)) //            movl   $val,%regl // zero extend
	} else if val == int64(int32(val)) {
		return asm.Bytes(0x48|hi, 0xc7, 0xc0|lo).Int32(int32(val)) // movq   $val,%reg  // sign extend
	} else {
		return asm.Bytes(0x48|hi, 0xb8+lo).Int64(val) //              movabs $val,%reg
	}
}

func (asm *Asm) Zero(dst *Var) *Asm {
	return asm.Set(dst, Int64(0))
}

func (asm *Asm) Set(dst *Var, a Arg) *Asm {
	switch a := a.(type) {
	case *Const:
		if val := a.val; val == int64(int32(val)) {
			return asm.Bytes(0x48, 0xc7, 0x87).Idx(dst).Int32(int32(val)) //  movq   $val,z(%rdi)
		}
	case *Var:
		if dst.desc == a.desc {
			return asm
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	return asm.Store(dst, tmp).hwFree(tmp, alloc)
}
