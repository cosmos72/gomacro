// +build amd64

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
 * z_amd64_test.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"fmt"
	"testing"

	. "github.com/cosmos72/gomacro/_jit/arch"
)

func TestDisasm(t *testing.T) {
	engine, _ := New()
	var asm Asm

	v1, v2, v3 := MakeVar0(0), MakeVar0(1), MakeVar0(2)

	for id := RLo; id <= RHi; id++ {
		asm.Init()
		if asm.RegIsUsed(id) {
			continue
		}
		r := MakeReg(id, Int64)
		asm.Asm(MOV, v1, r, //
			NEG, r, //
			NOT, r, //
			ADD, v2, r, //
			NOT, r, //
			NEG, r, //
			MOV, r, v3, //
		)

		insns, err := engine.Disasm(asm.Code(), 0x10000, 0)

		if err == nil {
			fmt.Printf("Disasm:\n")
			for _, insn := range insns {
				Show(insn)
			}
		}
	}
}

func TestDisasmSum(t *testing.T) {
	var asm Asm

	Total, I := MakeVar0(1), MakeVar0(2)
	asm.Init().Asm( //
		MOV, ConstInt64(0xFF), I,
		ADD, ConstInt64(2), I,
		ADD, I, Total)

	PrintDisasm(asm.Code())
}

func TestDisasmCast(t *testing.T) {
	N := [...]Mem{
		MakeVar0K(0, Uint64),
		MakeVar0K(1, Uint8), MakeVar0K(2, Uint16), MakeVar0K(3, Uint32),
		MakeVar0K(4, Int8), MakeVar0K(5, Int16), MakeVar0K(6, Int32),
	}
	V := [...]Mem{
		MakeVar0K(8, Uint64),
		MakeVar0K(9, Uint64), MakeVar0K(10, Uint64), MakeVar0K(11, Uint64),
		MakeVar0K(12, Uint64), MakeVar0K(13, Uint64), MakeVar0K(14, Uint64),
	}
	var asm Asm
	asm.Init()
	asm.Asm(
		NOP,
		CAST, N[1], V[1],
		CAST, N[2], V[2],
		CAST, N[3], V[3],
		CAST, N[4], V[4],
		CAST, N[5], V[5],
		CAST, N[6], V[6],
		NOP,
		CAST, V[1], N[1],
		CAST, V[2], N[2],
		CAST, V[3], N[3],
		CAST, V[4], N[4],
		CAST, V[5], N[5],
		CAST, V[6], N[6],
		RET,
	)

	PrintDisasm(asm.Code())
}

func TestDisasmLea(t *testing.T) {
	const (
		m int64 = 9
	)
	N := MakeVar0(0)
	M := MakeVar0(1)

	var asm Asm
	r := asm.Init().RegAlloc(N.Kind())
	asm.Asm(
		MUL, ConstInt64(m), N,
		LEA, N, r,
		LEA, M, r)
	asm.RegFree(r)

	PrintDisasm(asm.Code())
}

func TestDisasmShift(t *testing.T) {
	N := MakeVar0(0)
	M := MakeVar0(1)

	var asm Asm
	asm.Init()
	asm.RegIncUse(RCX)
	r := MakeReg(RCX, Uint8)
	asm.Asm(
		SHL, ConstInt64(0), M, // nop
		SHL, ConstInt64(1), M,
		SHL, r, N,
		SHR, ConstInt64(3), M,
		SHR, r, N,
	)
	asm.RegDecUse(RCX)

	PrintDisasm(asm.Code())
}
