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

func (asm *Asm) Zero(z *Var) *Asm {
	return asm.Set(z, Int64(0))
}

func (asm *Asm) Set(z, a *Var) *Asm {
	if a.Const {
		if val := a.Val; val == int64(int32(val)) {
			return asm.Bytes(0x48, 0xc7, 0x87).Idx(z).Int32(int32(val)) //  movq   $val,z(%rdi)
		}
	} else if a.Desc == z.Desc {
		return asm
	}
	return asm.load_rax(a).store_rax(z)
}
