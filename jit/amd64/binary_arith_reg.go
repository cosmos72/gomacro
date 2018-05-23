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
 * binary_arith_reg.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

import (
	"reflect"

	"github.com/cosmos72/gomacro/base"
)

// %rax += a
func (asm *Asm) Add_ax(a *Var) *Asm {
	if a.Const {
		val := a.Val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48, 0x05).Int32(int32(val)) // addq $val,%rax
		}
	}
	asm.load_rdx(a)
	asm.Bytes(0x48, 0x01, 0xd0) //  addq   %rdx,%rax
	return asm
}

// %rax -= a
func (asm *Asm) Sub_ax(a *Var) *Asm {
	if a.Const {
		val := a.Val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48, 0x2d).Int32(int32(val)) // subq   $val,%rax
		}
	}
	asm.load_rdx(a)
	asm.Bytes(0x48, 0x29, 0xd0) //   subq   %rdx,%rax
	return asm
}

// %rax *= a
func (asm *Asm) Mul_ax(a *Var) *Asm {
	if a.Const {
		val := a.Val
		if val == 0 {
			return asm.load_rax_const(0)
		} else if val == 1 {
			return asm
		} else if val == int64(int32(val)) {
			asm.Bytes(0x48, 0x69, 0xc0).Int32(int32(val)) //  imul   $val,%rax,%rax
			return asm
		}
	}
	asm.load_rdx(a)
	asm.Bytes(0x48, 0x0f, 0xaf, 0xc2) // imul   %rdx,%rax
	return asm
}

// ---------------- QUO --------------------

// FIXME: golang remainder rules are NOT the same as C !
// %rax /= a
func (asm *Asm) Quo_ax(a *Var) *Asm {
	if a.Const {
		// clobbers %rdx, %rsi
		asm.load_rsi_const(a.Val)
		switch base.KindToCategory(a.Kind) {
		case reflect.Uint:
			asm.Bytes(0x48, 0xf7, 0xfe) //         idiv   %rsi
		default:
			asm.Bytes(0x48, 0xf7, 0xf6) //         div    %rsi
		}
	} else {
		// clobbers %rdx
		switch base.KindToCategory(a.Kind) {
		case reflect.Uint:
			asm.Bytes(0x31, 0xd2)              //  xor    %edx,%edx
			asm.Bytes(0x48, 0xf7, 0xb7).Idx(a) //  divq   a(%rdi)
		default:
			asm.Bytes(0x48, 0x99)              //  cqto
			asm.Bytes(0x48, 0xf7, 0xbf).Idx(a) //  idivq  a(%rdi)
		}
	}
	return asm
}

// ---------------- REM --------------------

// FIXME: golang remainder rules are NOT the same as C !
// %rax %= a
func (asm *Asm) Rem_ax(a *Var) *Asm {
	if a.Const {
		// clobbers %rdx, %rsi
		asm.load_rsi_const(a.Val)
		switch base.KindToCategory(a.Kind) {
		case reflect.Uint:
			asm.Bytes(0x48, 0xf7, 0xfe) //         idiv   %rsi
		default:
			asm.Bytes(0x48, 0xf7, 0xf6) //         div    %rsi
		}
	} else {
		// clobbers %rdx
		switch base.KindToCategory(a.Kind) {
		case reflect.Uint:
			asm.Bytes(0x31, 0xd2)              //  xor    %edx,%edx
			asm.Bytes(0x48, 0xf7, 0xb7).Idx(a) //  divq   a(%rdi)
		default:
			asm.Bytes(0x48, 0x99)              //  cqto
			asm.Bytes(0x48, 0xf7, 0xbf).Idx(a) //  idivq  a(%rdi)
		}
	}
	asm.Bytes(0x48, 0x89, 0xd0) //  movq   %rdx,%rax
	return asm
}
