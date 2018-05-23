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
 * binary_arith.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

import (
	"reflect"

	"github.com/cosmos72/gomacro/base"
)

// z += a
func (asm *Asm) Add(z, a *Var) *Asm {
	if a.Const {
		val := a.Val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48, 0x81, 0x87).Idx(z).Int32(int32(val)) // addq   $val,z(%rdi)
		}
	}
	asm.load_rax(a)
	asm.Bytes(0x48, 0x01, 0x87).Idx(z) //   addq    %rax,z(%rdi)
	return asm
}

// z -= a
func (asm *Asm) Sub(z, a *Var) *Asm {
	if a.Desc == z.Desc {
		return asm.Zero(z)
	}
        if a.Const {
		val := a.Val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48, 0x81, 0xaf).Idx(z).Int32(int32(val)) // subq   $val,z(%rdi)
		}
	}
	asm.load_rax(a)
	asm.Bytes(0x48, 0x29, 0x87).Idx(z) //   subq    %rax,z(%rdi)
	return asm
}

// z *= a
func (asm *Asm) Mul(z, a *Var) *Asm {
	if a.Const {
		val := a.Val
		if val == 0 {
			return asm.Zero(z)
		} else if val == 1 {
			return asm
		} else if val == int64(int32(val)) {
			asm.Bytes(0x48, 0x69, 0x87).Idx(z).Int32(int32(val)) //  imul   $val,z(%rdi),%rax
			asm.store_rax(z)
			return asm
		}
	}
	asm.load_rax(a)
	asm.Bytes(0x48, 0x0f, 0xaf, 0x87).Idx(z) //  imul   z(%rdi),%rax
	asm.store_rax(z)
	return asm
}

// ---------------- QUO --------------------

// FIXME: golang remainder rules are NOT the same as C !
// z /= a
func (asm *Asm) Quo(z, a *Var) *Asm {
	asm.load_rax(z)
	switch base.KindToCategory(a.Kind) {
	case reflect.Uint:
		asm.Bytes(0x31, 0xd2)              //  xor    %edx,%edx
		asm.Bytes(0x48, 0xf7, 0xb7).Idx(a) //  divq   a(%rdi)
	default:
		asm.Bytes(0x48, 0x99)              //  cqto
		asm.Bytes(0x48, 0xf7, 0xbf).Idx(a) //  idivq  a(%rdi)
	}
	asm.store_rax(z)
	return asm
}

// ---------------- REM --------------------

// FIXME: golang remainder rules are NOT the same as C !
// z %= a
func (asm *Asm) Rem(z, a *Var) *Asm {
	asm.load_rax(z)
	switch base.KindToCategory(a.Kind) {
	case reflect.Uint:
		asm.Bytes(0x31, 0xd2)              //  xor    %edx,%edx
		asm.Bytes(0x48, 0xf7, 0xb7).Idx(a) //  divq   a(%rdi)
	default:
		asm.Bytes(0x48, 0x99)              //  cqto
		asm.Bytes(0x48, 0xf7, 0xbf).Idx(a) //  idivq  a(%rdi)
	}
	asm.store_rdx(z)
	return asm
}
