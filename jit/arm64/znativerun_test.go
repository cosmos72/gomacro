// +build arm64

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
 * znativerun_test.go
 *
 *  Created on Feb 07, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"testing"
)

func Init(asm *Asm) *Asm {
	return asm.Init().RegIncUse(X29).Asm(MOV, MakeMem(8, XSP, Uint64), MakeReg(X29, Uint64))
}

func TestNop(t *testing.T) {
	var f func()
	var asm Asm
	asm.Init().Func(&f)
	f()
}

func TestZero(t *testing.T) {
	var f func() uint64
	var asm Asm

	asm.Init()
	asm.Asm( //
		ZERO, MakeMem(8, XSP, Uint64),
	).Func(&f)

	actual := f()
	var expected uint64
	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestConst(t *testing.T) {
	var f func() uint64
	var asm Asm
	var expected uint64 = 7

	asm.Init()
	asm.Asm( //
		MOV, ConstUint64(expected), MakeMem(8, XSP, Uint64),
	).Func(&f)

	actual := f()
	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestLoadStore(t *testing.T) {
	var f func() uint64
	var asm Asm
	var expected uint64 = 0x12345678abcdef0

	r := asm.Init().RegAlloc(Uint64)
	asm.Asm( //
		MOV, ConstUint64(expected), r,
		MOV, r, MakeMem(8, XSP, Uint64),
	).Func(&f)

	actual := f()
	if actual != expected {
		t.Errorf("expected 0x%x, actual 0x%x", expected, actual)
	}
}
