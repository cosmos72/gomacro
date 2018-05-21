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

func (asm *Asm) Zero(z Bind) *Asm {
	return asm.SetInt64(z, 0)
}

func (asm *Asm) SetInt64(z Bind, val int64) *Asm {
	if val == int64(int32(val)) {
		return asm.Bytes(0x48, 0xc7, 0x87).Idx(z).Int32(int32(val)) //  movq   $val,z(%rdi)
	}
	return asm.load_rax_const(val).store_rax(z)
}

func (asm *Asm) SetUint64(z Bind, val uint64) *Asm {
	return asm.SetInt64(z, int64(val))
}

func (asm *Asm) Set(z, a Bind) *Asm {
	if a != z {
		asm.load_rax(z).store_rax(z)
	}
	return asm
}
