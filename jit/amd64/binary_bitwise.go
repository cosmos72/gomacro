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

// z &= a
func (asm *Asm) And(z, a *Var) *Asm {
	if a.Const {
		val := uint64(a.Val)
		if val == 0 {
			return asm.Zero(z)
		} else if val == ^uint64(0) {
			return asm
		} else if val == uint64(uint32(val)) {
			asm.Bytes(0x48, 0x81, 0xa7).Idx(z).Uint32(uint32(val)) //  andq   $val,z(%rdi)
			return asm
		}
	}
	asm.load_rax(a)
	asm.Bytes(0x48, 0x21, 0x87).Idx(z) //  andq   %rax,z(%rdi)
	return asm
}

// z |= a
func (asm *Asm) Or(z, a *Var) *Asm {
	if a.Const {
		val := a.Val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			asm.Bytes(0x48, 0x81, 0x8f).Idx(z).Int32(int32(val)) //  orq    $val,z(%rdi)
			return asm
		}
	}
	asm.load_rax(a)
	asm.Bytes(0x48, 0x09, 0x87).Idx(z) //  orq    %rax,z(%rdi)
	return asm
}

// z ^= a
func (asm *Asm) Xor(z, a *Var) *Asm {
	if a.Const {
		val := uint64(a.Val)
		if val == 0 {
			return asm
		} else if val == uint64(uint32(val)) {
			asm.Bytes(0x48, 0x81, 0x87).Idx(z).Uint32(uint32(val)) //  xorq   $val,z(%rdi)
			return asm
		}
	}
	asm.load_rax(a)
	asm.Bytes(0x48, 0x31, 0x87).Idx(z) //  xorq   %rax,z(%rdi)
	return asm
}

// z &^= a
func (asm *Asm) Andnot(z, a *Var) *Asm {
	if a.Const {
		val := ^a.Val // negate val!
		if val == 0 {
			return asm.Zero(z)
		} else if val == -1 {
			return asm
		} else if val == int64(int32(val)) {
			asm.Bytes(0x48, 0x81, 0xa7).Idx(z).Int32(int32(val)) //  andq   $val,z(%rdi)
			return asm
		}
		asm.load_rax_const(val)
	} else {
		asm.load_rax(a)
		asm.Bytes(0x48, 0xf7, 0xd0) //  not    %rax
	}
	asm.Bytes(0x48, 0x21, 0x87).Idx(z) //  andq   %rax,z(%rdi)
	return asm
}

