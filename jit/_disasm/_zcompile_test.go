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
 * zcompile_test.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"testing"

	. "github.com/cosmos72/gomacro/jit"
	arch "github.com/cosmos72/gomacro/jit/old/redirect"
)

const (
	S0 arch.SoftRegId = iota
	S1
)

var disasmArch = redirectArchId()

// result depends on which assembler is actually imported
// by "github.com/cosmos72/gomacro/jit/redirect" above
func redirectArchId() ArchId {
	switch arch.Name {
	case "amd64":
		return AMD64
	case "arm64":
		return ARM64
	default:
		return 0
	}
}

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

func TestCompileExpr1(t *testing.T) {
	var c Comp
	c.Init()
	r := MakeReg(RLo, Uint64)
	c.Expr(
		NewExpr1(
			NEG, NewExpr1(NOT, r),
		),
	)
	actual := c.Code()
	expected := Code{
		arch.ALLOC, S0, Uint64,
		arch.NOT2, r, S0,
		arch.NEG2, S0, S0,
	}

	if !SameCode(actual, expected) {
		t.Errorf("miscompiled code:\n\texpected %v\n\tactual   %v",
			expected, actual)
	}
}

func TestCompileExpr2(t *testing.T) {
	var c Comp
	c.Init()

	var asm arch.Asm
	asm.Init()

	c7 := MakeConst(7, Uint64)
	c9 := MakeConst(9, Uint64)
	r1 := asm.RegAlloc(Uint64)
	r2 := asm.RegAlloc(Uint64)
	// compile
	c.Expr(
		NewExpr2(
			ADD, NewExpr2(MUL, c7, r1), NewExpr2(SUB, c9, r2),
		),
	)
	actual := c.Code()
	t.Log(actual...)

	expected := Code{
		arch.ALLOC, S0, Uint64,
		arch.MUL3, c7, r1, S0,
		arch.ALLOC, S1, Uint64,
		arch.SUB3, c9, r2, S1,
		arch.ADD3, S0, S1, S0,
		arch.FREE, S1, arch.Uint64,
	}

	if !SameCode(actual, expected) {
		t.Errorf("miscompiled code:\n\texpected %v\n\tactual   %v",
			expected, actual)
	}

	// assemble
	asm.Asm(c.Code()...)
	asm.Epilogue()
	PrintDisasm(t, disasmArch, asm.Code())
}
