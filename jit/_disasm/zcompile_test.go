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
 * zcompile_test.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"testing"

	. "github.com/cosmos72/gomacro/jit"
	"github.com/cosmos72/gomacro/jit/asm"
)

const (
	S0 SoftRegId = iota
	S1
)

func CompareCode(actual Code, expected Code) int {

	if n1, n2 := len(actual), len(expected); n1 != n2 {
		if n1 < n2 {
			return n1
		}
		return n2
	}
	for i := range actual {
		if actual[i] != expected[i] {
			return i
		}
	}
	return -1
}

func TestCompileExpr1(t *testing.T) {
	var c Comp
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)
		r := MakeReg(c.RLo, Uint64)
		e := NewExpr1(
			NEG, NewExpr1(NOT, r),
		)
		c.Expr(e)
		actual := c.Code()

		t.Log("expr: ", e)

		expected := Code{
			asm.ALLOC, S0, Uint64,
			asm.NOT2, r, S0,
			asm.NEG2, S0, S0,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log(actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}
}

func TestCompileExpr2(t *testing.T) {
	var c Comp
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)
		a := c.Asm()

		c7 := MakeConst(7, Uint64)
		c9 := MakeConst(9, Uint64)
		r1 := a.RegAlloc(Uint64)
		r2 := a.RegAlloc(Uint64)
		// compile
		e := NewExpr2(
			ADD, NewExpr2(MUL, c7, r1), NewExpr2(SUB, c9, r2),
		)
		c.Expr(e)
		actual := c.Code()

		t.Log("expr: ", e)

		expected := Code{
			asm.ALLOC, S0, Uint64,
			asm.MUL3, c7, r1, S0,
			asm.ALLOC, S1, Uint64,
			asm.SUB3, c9, r2, S1,
			asm.ADD3, S0, S1, S0,
			asm.FREE, S1, asm.Uint64,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log(actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}

}

func TestCompileExpr3(t *testing.T) {
	var c Comp
	var s0 SoftRegId
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)

		c_2 := MakeConst(-2, Int64)
		m := c.MakeVar(0, 0, Int64)
		// compile
		e := NewExpr2(
			AND_NOT,
			NewExpr2(DIV, m, c_2),
			m,
		)
		c.Expr(e)
		actual := c.Code()

		t.Log("expr: ", e)

		expected := Code{
			asm.ALLOC, s0, Int64,
			asm.DIV3, m, c_2, s0,
			asm.AND_NOT3, s0, m, s0,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log(actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}
}

func TestCompileStmt1(t *testing.T) {
	var c Comp
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)

		m1 := c.MakeVar(0, 0, Uint64)
		m2 := c.MakeVar(1, 0, Uint32)
		m3w := c.MakeVar(2, 0, Uint16)
		m3 := c.MakeVar(2, 0, Uint8)
		m4w := c.MakeVar(3, 0, Uint16)

		ts := []Stmt{
			NewStmt1(INC, m1),                           // m1++
			NewStmt1(DEC, m2),                           // m2--
			NewStmt1(ZERO, m3),                          // m3 = 0
			NewStmt2(ASSIGN, m3w, NewExpr1(UINT16, m3)), // m3w = uint16(m3)
			NewStmt1(NOP, m4w),                          // _ = m4w
			NewStmt2(ASSIGN, m4w, m3w),                  // m4w = m3w
		}
		c.Compile(ts...)
		actual := c.Code()

		t.Logf("stmt: %v", ts)

		expected := Code{
			asm.INC, m1,
			asm.DEC, m2,
			asm.ZERO, m3,
			asm.CAST, m3, m3w,
			// asm.NOP, m4w, // NOP is optimized away
			asm.MOV, m3w, m4w,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log(actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}
}

func TestCompileStmt2(t *testing.T) {
	var c Comp
	_7 := MakeConst(7, Int64)
	_5 := MakeConst(5, Int64)
	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)
		s0 := c.AllocSoftReg(Int64)
		s1 := c.AllocSoftReg(Int64)
		sid0, sid1 := s0.Id(), s1.Id()
		sid2 := sid1 + 1

		stmt := NewStmt2(ASSIGN, s0,
			NewExpr2(SUB,
				NewExpr2(MUL, s1, _7),
				NewExpr2(DIV, s1, _5),
			),
		)
		c.Stmt(stmt)
		actual := c.Code()

		t.Logf("stmt: %v", stmt)

		expected := Code{
			asm.ALLOC, sid0, Int64,
			asm.ALLOC, sid1, Int64,
			asm.MUL3, sid1, _7, sid0,
			asm.ALLOC, sid2, Int64,
			asm.DIV3, sid1, _5, sid2,
			asm.SUB3, sid0, sid2, sid0,
			asm.FREE, sid2, Int64,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log(actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}
}

func TestCompileGetidx(t *testing.T) {
	var c Comp

	for _, archId := range []ArchId{asm.AMD64, asm.ARM64} {
		c.InitArchId(archId)

		stack_in := c.MakeParam(8, Uint64)
		stack_out := c.MakeParam(16, Uint8)
		_42 := ConstUint8(42)
		idx := ConstUint64(3)

		c.Asm().RegIncUse(c.RegIdConfig.RVAR)
		// on amd64 and arm64, in a func(env *Env) ...
		// the parameter env is on the stack at [RSP+8]
		rvar := MakeReg(c.RegIdConfig.RVAR, Uint64)
		// env = stack[env_param]
		c.Stmt2(ASSIGN, rvar, stack_in)
		// rvar = env.Ints equivalent to rvar = &env.Ints[0]
		c.Stmt2(ASSIGN, rvar, NewExprIdx(rvar, idx, Uint64))
		// compile accumulated jit expression and copy result to stack.
		e := _42
		// on amd64 and arm64, in a func(env *Env) ...
		// the return value is on the stack at [RSP+16]
		c.Stmt2(ASSIGN, stack_out, e)

		actual := c.Code()

		expected := Code{
			asm.MOV, stack_in, rvar,
			asm.GETIDX, rvar, idx, rvar,
			asm.MOV, _42, stack_out,
		}

		if i := CompareCode(actual, expected); i >= 0 {
			t.Errorf("miscompiled code at index %d:\n\texpected %v\n\tactual   %v",
				i, expected, actual)
		} else {
			t.Log(actual)
		}

		c.Epilogue()
		PrintDisasm(t, c.Assemble())
	}
}
