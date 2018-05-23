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
 * binary_bitwise.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

// %rax &= a
func (asm *Asm) And_ax(a *Var) *Asm {
	if a.Const {
		val := uint64(a.Val)
		if val == 0 {
			return asm.load_rax_const(0)
		} else if val == ^uint64(0) {
			return asm
		} else if val == uint64(uint32(val)) {
			asm.Bytes(0x25).Uint32(uint32(val)) //  andl   $val,%eax
			return asm
		}
	}
	asm.load_rdx(a)
	asm.Bytes(0x48, 0x21, 0xd0) //  andq   %rdx,%rax
	return asm
}

// %rax |= a
func (asm *Asm) Or_ax(a *Var) *Asm {
	if a.Const {
		val := a.Val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			asm.Bytes(0x48, 0x0d).Int32(int32(val)) //  orq    $val,%rax
			return asm
		}
	}
	asm.load_rdx(a)
	asm.Bytes(0x48, 0x09, 0xd0) //  orq    %rdx,%rax
	return asm
}

// %rax ^= a
func (asm *Asm) Xor_ax(a *Var) *Asm {
	if a.Const {
		val := uint64(a.Val)
		if val == 0 {
			return asm
		} else if val == uint64(uint32(val)) {
			asm.Bytes(0x48, 0x35).Uint32(uint32(val)) //  xorq   $val,%rax
			return asm
		}
	}
	asm.load_rdx(a)
	asm.Bytes(0x48, 0x31, 0xd0) //  xorq   %rdx,%rax
	return asm
}

// %rax &^= a
func (asm *Asm) Andnot_ax(a *Var) *Asm {
	if a.Const {
		val := ^uint64(a.Val) // negate val!
		if val == 0 {
			return asm.load_rax_const(0)
		} else if val == ^uint64(0) {
			return asm
		} else if val == uint64(uint32(val)) {
			asm.Bytes(0x25).Uint32(uint32(val)) //  andl   $val,%eax
			return asm
		}
		asm.load_rdx_const(int64(val))
	} else {
		asm.load_rdx(a)
		asm.Bytes(0x48, 0xf7, 0xd2) //  not    %rdx
	}
	asm.Bytes(0x48, 0x21, 0xd0) //  andq   %rdx,%rax
	return asm
}
