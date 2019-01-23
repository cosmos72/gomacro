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

	v1, v2, v3 := MakeVar(0), MakeVar(1), MakeVar(2)

	for r := RLo; r <= RHi; r++ {
		asm.Init()
		if asm.Regs.Contains(r) {
			continue
		}
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
				fmt.Printf("0x%x:\t%s\t\t%s\n", insn.Address, insn.Mnemonic, insn.OpStr)
			}
		}
	}
}
