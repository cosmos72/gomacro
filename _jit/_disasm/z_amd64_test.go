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
		if asm.RegIds.Contains(id) {
			continue
		}
		r := MakeReg(id, Int64)
		asm.Asm(MOV, r, v1, //
			NEG, r, //
			NOT, r, //
			ADD, r, v2, //
			NOT, r, //
			NEG, r, //
			MOV, v3, r, //
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
	engine, _ := New()
	var asm Asm

	Total, I := MakeVar0(1), MakeVar0(2)
	asm.Init().Asm( //
		MOV, I, ConstInt64(1),
		ADD, I, ConstInt64(1),
		ADD, Total, I)

	insns, err := engine.Disasm(asm.Code(), 0x10000, 0)

	if err == nil {
		fmt.Printf("Disasm:\n")
		for _, insn := range insns {
			Show(insn)
		}
	}
}

func TestDisasmCast(t *testing.T) {
	const n uint64 = 0xEFCDAB8967452301
	N := [...]Mem{
		MakeVar0K(0, Uint64),
		MakeVar0K(0, Uint8), MakeVar0K(0, Uint16), MakeVar0K(0, Uint32),
		MakeVar0K(0, Int8), MakeVar0K(0, Int16), MakeVar0K(0, Int32),
	}
	V := [...]Mem{
		MakeVar0K(0, Uint64),
		MakeVar0K(1, Uint64), MakeVar0K(2, Uint64), MakeVar0K(3, Uint64),
		MakeVar0K(4, Uint64), MakeVar0K(5, Uint64), MakeVar0K(6, Uint64),
	}
	var asm Asm
	asm.Init()
	r := asm.RegAlloc(Uint64)
	asm.Asm(
		CAST, V[1], N[1], // MOV, V[1], r,
		CAST, V[2], N[2], // MOV, V[2], r,
		CAST, V[3], N[3], // MOV, V[3], r,
		CAST, V[4], N[4], // MOV, V[4], r,
		CAST, V[5], N[5], // MOV, V[5], r,
		CAST, V[6], N[6], // MOV, V[6], r,
	).RegFree(r)

	engine, _ := New()
	insns, err := engine.Disasm(asm.Code(), 0x10000, 0)

	if err == nil {
		fmt.Printf("Disasm:\n")
		for _, insn := range insns {
			Show(insn)
		}
	}
}
