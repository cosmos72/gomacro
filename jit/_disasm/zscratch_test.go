/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * zscratch_test.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"testing"

	"github.com/cosmos72/gomacro/jit"
	"github.com/cosmos72/gomacro/jit/asm"
)

func parsehexdigit(ch byte) uint8 {
	switch {
	case ch >= '0' && ch <= '9':
		return ch - '0'
	case ch >= 'A' && ch <= 'F':
		return ch - 'A' + 10
	case ch >= 'a' && ch <= 'f':
		return ch - 'a' + 10
	default:
		return 0
	}
}

func HexToBinary(hex string) []byte {
	buf := make([]byte, len(hex)/2)
	for i := 0; i < len(hex)/2; i++ {
		buf[i] = parsehexdigit(hex[i*2])<<4 | parsehexdigit(hex[i*2+1])
	}
	return buf
}

func TestAmd64Shift2(t *testing.T) {
	code := asm.MachineCode{
		asm.AMD64,
		HexToBinary("488b742408488b761866c16e1805668b4e1866d36e18"),
	}
	PrintDisasm(t, code)

	_5 := asm.MakeConst(5, jit.Uint16)

	var c jit.Comp
	for _, archid := range [...]asm.ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archid)
		m := c.MakeVar(0, 0, jit.Uint16)
		c.Compile(jit.Source{
			jit.SHR_ASSIGN, m, jit.NewExpr2(jit.SHR, m, _5),
		})
		t.Log(c.Code())

		PrintDisasm(t, c.Assemble())
	}
}

func TestAmd64Collatz(t *testing.T) {
	PrintDisasm(t, asm.MachineCode{
		asm.AMD64,
		HexToBinary("488b742408488b7618486b1e0348ffc34889df48d1ef48893e"),
	})
}
