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
 * bitwise.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

// %rax &= a
func (asm *Asm) And(z hwReg, a Arg) *Asm {
	lo, hi := z.lohi()
	if a.Const() {
		val := a.(*Const).val
		if val == 0 {
			return asm.LoadConst(z, 0)
		} else if val == -1 {
			return asm
		} else if val == int64(uint32(val)) {
			if hi != 0 {
				asm.Bytes(0x41)
			}
			return asm.Bytes(0x81, 0xe0+lo).Uint32(uint32(val)) //        andl  $val,%reg_z // zero extend
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xe0+lo).Int32(int32(val)) // andq  $val,%reg_z // sign extend
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x21, 0xc0+lo+tmp.lo()*8) //      and  %reg_tmp,%reg_z
	asm.hwFree(tmp, alloc)
	return asm
}

// %rax |= a
func (asm *Asm) Or(z hwReg, a Arg) *Asm {
	lo, hi := z.lohi()
	if a.Const() {
		val := a.(*Const).val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xc8+lo).Int32(int32(val)) // orq   $val,%reg_z // sign extend
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x09, 0xc0+lo+tmp.lo()*8) //      or   %reg_tmp,%reg_z
	asm.hwFree(tmp, alloc)
	return asm
}

// %rax ^= a
func (asm *Asm) Xor(z hwReg, a Arg) *Asm {
	lo, hi := z.lohi()
	if a.Const() {
		val := a.(*Const).val
		if val == 0 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xf0+lo).Int32(int32(val)) // xorq  $val,%reg_z // sign extend
		}
	}
	tmp, alloc := asm.hwAlloc(a)
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x31, 0xc0+lo+tmp.lo()*8) //      xor  %reg_tmp,%reg_z
	asm.hwFree(tmp, alloc)
	return asm
}

// %rax &^= a
func (asm *Asm) Andnot(z hwReg, a Arg) *Asm {
	lo, hi := z.lohi()
	var tmp hwReg
	var alloc bool
	if a.Const() {
		val := ^a.(*Const).val // negate val!
		if val == 0 {
			return asm.LoadConst(z, 0)
		} else if val == -1 {
			return asm
		} else if val == int64(int32(val)) {
			return asm.Bytes(0x48+hi, 0x81, 0xe0+lo).Int32(int32(val)) // andq  $val,%reg_z // sign extend
		}
		tmp = asm.hwAllocConst(val)
		alloc = true
	} else {
		tmp, alloc = asm.hwAlloc(a)
		asm.Bytes(0x48, 0xf7, 0xd2) //  not    %reg_tmp
	}
	asm.Bytes(0x48+hi+tmp.hi()*4, 0x21, 0xc0+lo+tmp.lo()*8) //      and  %reg_tmp,%reg_z
	asm.hwFree(tmp, alloc)
	return asm
}
