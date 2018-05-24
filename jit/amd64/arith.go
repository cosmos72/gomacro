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
 * arith.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

import (
	"reflect"

	"github.com/cosmos72/gomacro/base"
)

// %reg_z += a
func (asm *Asm) Add(z Reg, a *Var) *Asm {
	lo, hi := z.lohi()
	if a.Const {
		val := a.Val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xc0+lo).Int32(int32(val)) // add  $val,%reg_z // sign extend
		}
	}
	tmp := asm.ToReg(a)
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x01, 0xc0+lo+tmp.lo()*8) //      add  %reg_tmp,%reg_z
	asm.FreeReg(tmp)
	return asm
}

// %reg_z -= a
func (asm *Asm) Sub(z Reg, a *Var) *Asm {
	lo, hi := z.lohi()
	if a.Const {
		val := a.Val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xe8+lo).Int32(int32(val)) // sub  $val,%reg_z // sign extend
		}
	}
	tmp := asm.ToReg(a)
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x29, 0xc0+lo+tmp.lo()*8) //      sub  %reg_tmp,%reg_z
	asm.FreeReg(tmp)
	return asm
}

// %reg_z *= a
func (asm *Asm) Mul(z Reg, a *Var) *Asm {
	lo, hi := z.lohi()
	if a.Const {
		val := a.Val
		if val == 0 {
			return asm.LoadConst(z, 0)
		} else if val == 1 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi*5, 0x69, 0xc0+lo*9).Int32(int32(val)) // imul  $val,%reg_z,%reg_z // sign extend
		}
	}
	tmp := asm.ToReg(a)
	asm.Bytes(0x48+hi*4+tmp.hi(), 0x0f, 0xaf+lo*8+tmp.lo()) //      imul  %reg_tmp,%reg_z
	asm.FreeReg(tmp)
	return asm
}

// ---------------- QUO --------------------

// FIXME: golang remainder rules are NOT the same as C !
// %reg_z /= a
func (asm *Asm) Quo(z Reg, a *Var) *Asm {
	tosave := RegSet(DX)
	if z != AX {
		tosave.Set(AX)
	}
	tosave = asm.PushRegs(tosave)
	asm.Mov(AX, z) // nop if z == AX
	asm.quorem(a)
	asm.Mov(z, AX) // nop if z == AX
	asm.PopRegs(tosave)
	return asm
}

// ---------------- REM --------------------

// FIXME: golang remainder rules are NOT the same as C !
// %reg %= a
func (asm *Asm) Rem(z Reg, a *Var) *Asm {
	tosave := RegSet(AX)
	if z != DX {
		tosave.Set(DX)
	}
	tosave = asm.PushRegs(tosave)
	asm.Mov(AX, z) // nop if z == AX
	asm.quorem(a)
	asm.Mov(z, DX) // nop if z == DX
	asm.PopRegs(tosave)
	return asm
}

func (asm *Asm) quorem(a *Var) {
	if a.Const {
		tmp := asm.ToReg(a)
		switch base.KindToCategory(a.Kind) {
		case reflect.Uint:
			asm.Bytes(0x48+tmp.hi(), 0xf7, 0xf0+tmp.lo()) //         div    %reg_tmp
		default:
			asm.Bytes(0x48+tmp.hi(), 0xf7, 0xf8+tmp.lo()) //         idiv   %reg_tmp
		}
		asm.FreeReg(tmp)
	} else {
		switch base.KindToCategory(a.Kind) {
		case reflect.Uint:
			asm.Bytes(0x31, 0xd2)              //  xor    %edx,%edx
			asm.Bytes(0x48, 0xf7, 0xb7).Idx(a) //  divq   a(%rdi)
		default:
			asm.Bytes(0x48, 0x99)              //  cqto
			asm.Bytes(0x48, 0xf7, 0xbf).Idx(a) //  idivq  a(%rdi)
		}
	}
}
