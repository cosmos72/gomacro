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
func (asm *Asm) Add(z hwReg, a Arg) *Asm {
	lo, hi := z.lohi()
	if a.Const() {
		val := a.(*Const).val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xc0+lo).Int32(int32(val)) // add  $val,%reg_z // sign extend
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x01, 0xc0+lo+tmp.lo()*8) //      add  %reg_tmp,%reg_z
	asm.hwFree(tmp, alloc)
	return asm
}

// %reg_z -= a
func (asm *Asm) Sub(z hwReg, a Arg) *Asm {
	lo, hi := z.lohi()
	if a.Const() {
		val := a.(*Const).val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xe8+lo).Int32(int32(val)) // sub  $val,%reg_z // sign extend
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x29, 0xc0+lo+tmp.lo()*8) //      sub  %reg_tmp,%reg_z
	asm.hwFree(tmp, alloc)
	return asm
}

// %reg_z *= a
func (asm *Asm) Mul(z hwReg, a Arg) *Asm {
	lo, hi := z.lohi()
	if a.Const() {
		val := a.(*Const).val
		if val == 0 {
			return asm.LoadConst(z, 0)
		} else if val == 1 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi*5, 0x69, 0xc0+lo*9).Int32(int32(val)) // imul  $val,%reg_z,%reg_z // sign extend
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Bytes(0x48+hi*4+tmp.hi(), 0x0f, 0xaf+lo*8+tmp.lo()) //      imul  %reg_tmp,%reg_z
	asm.hwFree(tmp, alloc)
	return asm
}

// ---------------- QUO --------------------

// %reg_z /= a
func (asm *Asm) Quo(z hwReg, a Arg) *Asm {
	return asm.quorem(z, a, false)
}

// ---------------- REM --------------------

// %reg %= a
func (asm *Asm) Rem(z hwReg, a Arg) *Asm {
	return asm.quorem(z, a, true)
}

// FIXME: golang remainder rules are NOT the same as C !
func (asm *Asm) quorem(z hwReg, a Arg, rem bool) *Asm {
	tosave := newHwRegs(rDX)
	if z != rAX {
		tosave.Set(rAX)
	}
	tosave = asm.PushRegs(tosave)
	var b hwReg
	if a.Reg().Valid() && tosave.Contains(a.Reg()) {
		b = asm.liveRegs.Alloc()
		asm.Mov(b, a.Reg())
		a = b
	}
	asm.Mov(rAX, z) // nop if z == AX

	switch a := a.(type) {
	case *Var:
		switch base.KindToCategory(a.Kind()) {
		case reflect.Uint:
			asm.Bytes(0x31, 0xd2)              //  xor    %edx,%edx
			asm.Bytes(0x48, 0xf7, 0xb7).Idx(a) //  divq   a(%rdi)
		default:
			asm.Bytes(0x48, 0x99)              //  cqto
			asm.Bytes(0x48, 0xf7, 0xbf).Idx(a) //  idivq  a(%rdi)
		}
	default:
		tmp, alloc := asm.hwAlloc(a)
		switch base.KindToCategory(a.Kind()) {
		case reflect.Uint:
			asm.Bytes(0x31, 0xd2)                         //  xor    %edx,%edx
			asm.Bytes(0x48+tmp.hi(), 0xf7, 0xf0+tmp.lo()) //  div    %reg_tmp
		default:
			asm.Bytes(0x48, 0x99)                         //  cqto
			asm.Bytes(0x48+tmp.hi(), 0xf7, 0xf8+tmp.lo()) //  idiv   %reg_tmp
		}
		asm.hwFree(tmp, alloc)
	}
	if b.Valid() {
		asm.liveRegs.Free(b)
	}
	if rem {
		asm.Mov(z, rDX) // nop if z == DX
	} else {
		asm.Mov(z, rAX) // nop if z == AX
	}
	asm.PopRegs(tosave)
	return asm
}
