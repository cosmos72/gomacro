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

func (asm *Asm) store_rax(z *Var) *Asm {
	return asm.Bytes(0x48, 0x89, 0x87).Idx(z) //    movq   %rax,z(%rdi)
}

func (asm *Asm) store_rdx(z *Var) *Asm {
	return asm.Bytes(0x48, 0x89, 0x97).Idx(z) //    movq   %rdx,z(%rdi)
}

func (asm *Asm) load_rax(a *Var) *Asm {
	if a.Const {
		return asm.load_rax_const(a.Val)
	}
	return asm.Bytes(0x48, 0x8b, 0x87).Idx(a) //    movq   a(%rdi),%rax
}

func (asm *Asm) load_rax_const(val int64) *Asm {
	if val != int64(int32(val)) {
		return asm.Bytes(0x48, 0xb8).Int64(val) //              movabs $val,%rax
	} else if val < 0 {
		return asm.Bytes(0x48, 0xc7, 0xc0).Int32(int32(val)) // movq   $val,%rax
	}
	return asm.Bytes(0xb8).Int32(int32(val)) //                     movl   $val,%eax
}

func (asm *Asm) load_rdx(a *Var) *Asm {
	if a.Const {
		return asm.load_rdx_const(a.Val)
	}
	return asm.Bytes(0x48, 0x8b, 0x97).Idx(a) //    movq   a(%rdi),%rdx
}

func (asm *Asm) load_rdx_const(val int64) *Asm {
	if val != int64(int32(val)) {
		return asm.Bytes(0x48, 0xba).Int64(val) //              movabs $val,%rdx
	} else if val < 0 {
		return asm.Bytes(0x48, 0xc7, 0xc2).Int32(int32(val)) // movq   $val,%rdx
	}
	return asm.Bytes(0xba).Int32(int32(val)) //                     movl   $val,%edx
}

func (asm *Asm) load_rsi(a *Var) *Asm {
	if a.Const {
		return asm.load_rsi_const(a.Val)
	}
	return asm.Bytes(0x48, 0x8b, 0xb7).Idx(a) // mov    a(%rdi),%rsi            
}

func (asm *Asm) load_rsi_const(val int64) *Asm {
	if val != int64(int32(val)) {
		return asm.Bytes(0x48, 0xbe).Int64(val) //              movabs $val,%rsi
	} else if val < 0 {
		return asm.Bytes(0x48, 0xc7, 0xc6).Int32(int32(val)) // movq   $val,%rsi
	}
	return asm.Bytes(0xc7, 0x06).Int32(int32(val))       //         movl   $val,%rsi
}

