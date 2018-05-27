// +build arm64

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
 * arith_amd64.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

// %reg_z += a
func (asm *Asm) Add(z Reg, a Arg) *Asm {
	if a.Const() {
		if val := uint64(a.(*Const).val); val < 4096 {
			return asm.Uint32(0x91<<24|uint32(val)<<10|asm.lo(z)*0x21) // add  %reg_z,%reg_z,$val
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Uint32(0x8b<<24|tmp.lo()<<16|asm.lo(z)*0x21) //  add  %reg_z,%reg_z,%reg_tmp
	asm.hwFree(tmp, alloc)
	return asm
}

// %reg_z -= a
func (asm *Asm) Sub(z Reg, a Arg) *Asm {
	return asm
}

// %reg_z *= a
func (asm *Asm) Mul(z Reg, a Arg) *Asm {
	return asm
}

// %reg_z /= a
func (asm *Asm) Quo(z Reg, a Arg) *Asm {
	return asm
}

// %reg %= a
func (asm *Asm) Rem(z Reg, a Arg) *Asm {
	return asm
}

// %reg_z = - %reg_z
func (asm *Asm) Neg(z Reg) *Asm {
	return asm
}
