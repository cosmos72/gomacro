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
 * z_test.go
 *
 *  Created on Feb 07, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"testing"
)

func SameCode(actual Code, expected Code) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i := range actual {
		if actual[i] != expected[i] {
			return false
		}
	}
	return true
}

func TestAmd64SoftReg(t *testing.T) {
	var asm Asm
	asm.Init()

	var a, b, c SoftReg = 0, 1, 2
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

	actual := asm.Code()
	expected := Code{
		0x48, 0xc7, 0xc0, 0x01, 0x00, 0x00, 0x00, //  movq	$1, %rax
		0x48, 0xc7, 0xc2, 0x02, 0x00, 0x00, 0x00, //  movq	$2, %rdx
		0x48, 0x89, 0xc3, //                          movq	%rax, %rbx
		0x48, 0x01, 0xd3, //                          addq	%rdx, %rbx
		0xc3, //                                      retq
	}

	if !SameCode(actual, expected) {
		t.Errorf("miscompiled code:\n\texpected %x\n\tactual   %x",
			expected, actual)
	}
}
