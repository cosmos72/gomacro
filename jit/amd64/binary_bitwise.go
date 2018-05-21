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

func (asm *Asm) AndInt64(z, a, b Bind) *Asm {
	asm.load_rax(a)
	asm.Bytes(0x48, 0x23, 0x87).Idx(b) //  andq   b(%rdi),%rax
	asm.store_rax(z)
	return asm
}

func (asm *Asm) AndUint64(z, a, b Bind) *Asm {
	return asm.AndInt64(z, a, b)
}

func (asm *Asm) OrInt64(z, a, b Bind) *Asm {
	asm.load_rax(a)
	asm.Bytes(0x48, 0x0b, 0x87).Idx(b) //  orq    b(%rdi),%rax
	asm.store_rax(z)
	return asm
}

func (asm *Asm) OrUint64(z, a, b Bind) *Asm {
	return asm.OrInt64(z, a, b)
}

func (asm *Asm) XorInt64(z, a, b Bind) *Asm {
	asm.load_rax(a)
	asm.Bytes(0x48, 0x33, 0x87).Idx(b) //  xorq   b(%rdi),%rax
	asm.store_rax(z)
	return asm
}

func (asm *Asm) XorUint64(z, a, b Bind) *Asm {
	return asm.XorInt64(z, a, b)
}

func (asm *Asm) AndnotInt64(z, a, b Bind) *Asm {
	asm.load_rax(b)
	asm.Bytes(0x48, 0xf7, 0xd0)        //  not    %rax
	asm.Bytes(0x48, 0x23, 0x87).Idx(a) //  andq   a(%rdi),%rax
	asm.store_rax(z)
	return asm
}

func (asm *Asm) AndnotUint64(z, a, b Bind) *Asm {
	return asm.AndnotInt64(z, a, b)
}
