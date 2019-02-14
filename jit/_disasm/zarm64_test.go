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
 * zarm_64_test.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"testing"

	. "github.com/cosmos72/gomacro/jit/old/arm64"
)

func TestArm64Sample(T *testing.T) {
	var asm Asm

	for id := RLo; id+2 <= RHi; id++ {
		asm.Init()
		if asm.RegIsUsed(id) || asm.RegIsUsed(id+1) || asm.RegIsUsed(id+2) {
			continue
		}
		r := MakeReg(id+0, Int64)
		s := MakeReg(id+1, Int64)
		t := MakeReg(id+2, Int64)
		m := MakeMem(8, id, Int64)
		c := ConstInt64(0xFFF)
		one := ConstUint8(1)
		ur := MakeReg(id+0, Uint64)
		us := MakeReg(id+1, Uint64)
		ut := MakeReg(id+2, Uint64)
		br := MakeReg(id+0, Uint8)
		bt := MakeReg(id+2, Uint8)
		asm.RegIncUse(id).RegIncUse(id + 1).RegIncUse(id + 2)
		asm.Asm(MOV, c, r, //
			MOV, c, m, //
			MOV, m, r, //
			NOP,           //
			ADD3, r, s, t, //
			SUB3, r, s, t, //
			AND3, r, s, t, //
			OR3, r, s, t, //
			XOR3, r, s, t, //
			SHL3, r, us, t, //
			SHR3, ur, us, ut, //
			SHR3, r, us, t, //
			NOP, //
			// test commutativity optimization
			ADD3, c, r, t, //
			SUB3, r, c, t, //
			AND3, c, r, t, //
			OR3, c, r, t, //
			XOR3, r, c, t, //
			SHL3, r, one, t, //
			SHR3, ur, one, ut, //
			SHR3, r, one, t, //
			NOP, //
			NOP, //
			// test 8-bit registers
			ADD3, one, br, bt, //
			SUB3, br, one, bt, //
			AND3, one, br, bt, //
			OR3, one, br, bt, //
			XOR3, br, one, bt, //
			SHL3, br, one, bt, //
			SHR3, br, one, bt, //
		).Epilogue()
		asm.RegDecUse(id).RegDecUse(id + 1).RegDecUse(id + 2)

		if id == RLo || id == RHi {
			PrintDisasm(T, ARM64, asm.Code())
		}
	}
}

func TestArm64Zero(t *testing.T) {
	r := MakeReg(RLo, Uint64)
	xzr := MakeReg(XZR, Uint64)
	m := MakeMem(8, XSP, Uint64)

	var asm Asm
	asm.Init().Asm(
		ZERO, r,
		MOV, xzr, r,
		ZERO, m,
		RET)

	PrintDisasm(t, ARM64, asm.Code())
}

func TestArm64Cast(t *testing.T) {
	var asm Asm
	asm.Init()
	for _, skind := range [...]Kind{
		Int8, Int16, Int32, Int64,
		Uint8, Uint16, Uint32, Uint64,
	} {

		src := MakeReg(RLo, skind)
		for _, dkind := range [...]Kind{Uint8, Uint16, Uint32, Uint64} {
			dst := MakeReg(RLo, dkind)
			asm.Asm(CAST, src, dst)
		}
	}
	PrintDisasm(t, ARM64, asm.Code())
}

func TestArm64Mem(t *testing.T) {
	var asm Asm
	asm.Init()

	id := RLo
	asm.RegIncUse(id)

	for _, skind := range [...]Kind{
		Int8, Int16, Int32, Int64,
		Uint8, Uint16, Uint32, Uint64,
	} {
		s := MakeMem(0, id, skind)
		c := MakeConst(0xFF, skind)
		for _, dkind := range [...]Kind{Uint8, Uint16, Uint32, Uint64} {

			d := MakeMem(8, id, dkind)
			if skind == dkind {
				asm.Asm(ADD3, s, c, d)
			} else {
				asm.Asm(CAST, s, d)
			}
		}
		asm.Asm(NOP)
	}
	asm.Epilogue()
	PrintDisasm(t, ARM64, asm.Code())
}

func TestArm64Unary(t *testing.T) {

	var asm Asm
	asm.Init()
	r := MakeReg(X27, Uint64)
	s := MakeReg(X28, Uint64)
	v := MakeMem(0, X29, Uint64)

	asm.Asm( //
		MOV, v, r,
		NEG2, r, s,
		NOT2, s, r,
		MOV, r, v,
	)
	asm.Epilogue()
	PrintDisasm(t, ARM64, asm.Code())
}

func TestArm64SoftReg(t *testing.T) {
	var asm Asm
	asm.Init()

	var a, b, c SoftRegId = 0, 1, 2
	asm.Asm(
		ALLOC, a, Uint64,
		ALLOC, b, Uint64,
		ALLOC, c, Uint64,
		MOV, ConstUint64(1), a,
		MOV, ConstUint64(2), b,
		ADD3, a, b, c,
		FREE, a, Uint64,
		FREE, b, Uint64,
		FREE, c, Uint64,
	).Epilogue()
	PrintDisasm(t, ARM64, asm.Code())
}
