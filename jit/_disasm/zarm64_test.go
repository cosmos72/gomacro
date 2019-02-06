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

	. "github.com/cosmos72/gomacro/jit/arm64"
)

func TestArm64Sample(t *testing.T) {
	var asm Asm

	for id := RLo; id+2 <= RHi; id++ {
		asm.Init()
		if asm.RegIsUsed(id) || asm.RegIsUsed(id+1) || asm.RegIsUsed(id+2) {
			continue
		}
		r := MakeReg(id+0, Uint64)
		s := MakeReg(id+1, Uint64)
		t := MakeReg(id+2, Uint64)
		m := MakeMem(8, id, Uint64)
		c := ConstUint64(0xFFFF000000000000)
		asm.RegIncUse(id).RegIncUse(id + 1).RegIncUse(id + 2)
		asm.Asm(MOV, c, r, //
			// MOV, c, m, //
			MOV, m, r, //
			NOP,           //
			ADD3, r, s, t, //
			SUB3, r, s, t, //
			AND3, r, s, t, //
			OR3, r, s, t, //
			XOR3, r, s, t, //
			SHL3, r, s, t, //
			SHR3, r, s, t, //
			NOP,           //
			ADD3, c, r, t, // test commutativity optimization
			SUB3, r, c, t, //
			AND3, c, r, t, //
			OR3, c, r, t, //
			XOR3, r, c, t, //
		).Epilogue()
		asm.RegDecUse(id).RegDecUse(id + 1).RegDecUse(id + 2)

		PrintDisasm(ARM64, asm.Code())
	}
}
