// +build arm64

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
 * z_exec_test.go
 *
 *  Created on Feb 07, 2019
 *      Author Massimiliano Ghilardi
 */

package arm64

import (
	"testing"
)

func Init(asm *Asm) *Asm {
	asm.InitArch(Arm64{})
	asm.RegIncUse(X29)
	asm.Load(MakeMem(8, XSP, Uint64), MakeReg(X29, Uint64))
	return asm
}

func TestExecNop(t *testing.T) {
	var f func()
	var asm Asm
	asm.InitArch(Arm64{}).Func(&f)
	f()
}

func TestExecZero(t *testing.T) {
	var f func() uint64
	var asm Asm
	asm.InitArch(Arm64{})

	asm.Asm( //
		ZERO, MakeMem(8, XSP, Uint64),
	).Func(&f)

	actual := f()
	expected := uint64(0)
	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestExecConst(t *testing.T) {
	var f func() uint64
	var asm Asm
	var expected uint64 = 7

	asm.InitArch(Arm64{})
	asm.Asm( //
		MOV, ConstUint64(expected), MakeMem(8, XSP, Uint64),
	).Func(&f)

	actual := f()
	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestExecLoadStore(t *testing.T) {
	var f func() uint64
	var asm Asm
	var expected uint64 = 0x12345678abcdef0

	r := asm.InitArch(Arm64{}).RegAlloc(Uint64)
	asm.Asm( //
		MOV, ConstUint64(expected), r,
		MOV, r, MakeMem(8, XSP, Uint64),
	).Func(&f)

	actual := f()
	if actual != expected {
		t.Errorf("expected 0x%x, actual 0x%x", expected, actual)
	}
}

func TestExecUnary(t *testing.T) {
	var c uint64 = 0x64776657f7754abc
	binds := [...]uint64{c}

	var asm Asm
	r := Init(&asm).RegAlloc(Uint64)
	v := MakeMem(0, X29, Uint64)

	var f func(*uint64)
	asm.Asm( //
		MOV, v, r,
		NEG1, r,
		NOT1, r,
		MOV, r, v,
	).Func(&f)
	f(&binds[0])

	expected := ^-c
	actual := binds[0]

	if actual != expected {
		t.Errorf("expected 0x%x, actual 0x%x", expected, actual)
	}
}
