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
 * reg.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

func (asm *Asm) load_rax(a Bind) *Asm {
	return asm.Bytes(0x48, 0x8b, 0x87).Idx(a) //    movq   a(%rdi),%rax
}

func (asm *Asm) load_rax_const(val int64) *Asm {
	if val == int64(int32(val)) {
		return asm.Bytes(0xba).Int32(int32(val)) // movq $val,%rax
	}
	return asm.Bytes(0x48, 0xb8).Int64(val) //      movabs $val,%rax
}

func (asm *Asm) store_rax(z Bind) *Asm {
	return asm.Bytes(0x48, 0x89, 0x87).Idx(z) //    movq   %rax,z(%rdi)
}

func (asm *Asm) store_rdx(z Bind) *Asm {
	return asm.Bytes(0x48, 0x89, 0x97).Idx(z) //    movq   %rdx,z(%rdi)
}
