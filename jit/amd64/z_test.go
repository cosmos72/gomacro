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
 * z_test.go
 *
 *  Created on Feb 07, 2019
 *      Author Massimiliano Ghilardi
 */

package amd64

import (
	"testing"
)

func Var(index uint16) Mem {
	return MakeMem(int32(index)*8, RSI, Int64)
}

func VarK(index uint16, k Kind) Mem {
	return MakeMem(int32(index)*8, RSI, k)
}

func TestAmd64SoftRegId(t *testing.T) {
	var asm Asm
	asm.InitArch(Amd64{})

	var a, b, c SoftRegId = 0, 1, 2
	code := []AsmCode{
		ALLOC, a, Uint64,
		ALLOC, b, Uint64,
		ALLOC, c, Uint64,
		MOV, ConstUint64(1), a,
		MOV, ConstUint64(2), b,
		ADD3, a, b, c,
		FREE, a, Uint64,
		FREE, b, Uint64,
		FREE, c, Uint64,
		RET,
	}
	asm.Asm(code...)
	// t.Log(code...)

	actual := asm.Code()
	expected := MachineCode{
		AMD64,
		[]uint8{
			0x48, 0xc7, 0xc3, 0x01, 0x00, 0x00, 0x00, //  movq	$1, %rbx
			0x48, 0xc7, 0xc6, 0x02, 0x00, 0x00, 0x00, //  movq	$2, %rsi
			0x48, 0x89, 0xdf, //                          movq	%rbx, %rdi
			0x48, 0x01, 0xf7, //                          addq	%rsi, %rdi
			0xc3, //                                      retq
		},
	}

	if !actual.Equal(expected) {
		t.Errorf("bad assembled code:\n\texpected %x\n\tactual   %x",
			expected, actual)
	}
}
