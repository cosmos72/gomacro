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

func MakeCode(instr ...uint32) Code {
	code := make(Code, len(instr)*4)
	for i, inst := range instr {
		code[4*i+0] = byte(inst >> 0)
		code[4*i+1] = byte(inst >> 8)
		code[4*i+2] = byte(inst >> 16)
		code[4*i+3] = byte(inst >> 24)
	}
	return code
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

func TestArm64Sample(t *testing.T) {
	var asm Asm

	id := RLo
	asm.Init()
	x := MakeReg(id+0, Uint64)
	y := MakeReg(id+1, Uint64)
	z := MakeReg(id+2, Uint64)
	m := MakeMem(8, id, Uint64)
	c := ConstUint64(0xFFF)
	asm.RegIncUse(id).RegIncUse(id + 1).RegIncUse(id + 2)
	asm.Asm( //
		MOV, MakeMem(8, XSP, Uint64), MakeReg(X29, Uint64),
		MOV, c, x, //
		MOV, c, m, //
		MOV, m, x, //
		NOP,           //
		ADD3, x, y, z, //
		SUB3, x, y, z, //
		AND3, x, y, z, //
		OR3, x, y, z, //
		XOR3, x, y, z, //
		SHL3, x, y, z, //
		SHR3, x, y, z, //
		NOP,           //
		ADD3, c, x, z, // test commutativity optimization
		SUB3, x, c, z, //
		AND3, c, x, z, //
		OR3, c, x, z, //
		XOR3, x, c, z, //
	).Epilogue()
	asm.RegDecUse(id).RegDecUse(id + 1).RegDecUse(id + 2)

	actual := asm.Code()
	expected := MakeCode(
		0xf94007fd, //	ldr	x29, [sp, #8]
		0xd281ffe0, //	mov	x0, #0xfff
		0xd281ffe3, //	mov	x3, #0xfff
		0xf9000403, //	str	x3, [x0, #8]
		0xf9400400, //	ldr	x0, [x0, #8]
		0xd503201f, //	nop
		0x8b010002, //	add	x2, x0, x1
		0xcb010002, //	sub	x2, x0, x1
		0x8a010002, //	and	x2, x0, x1
		0xaa010002, //	orr	x2, x0, x1
		0xca010002, //	eor	x2, x0, x1
		0x9ac12002, //	lsl	x2, x0, x1
		0x9ac12402, //	lsr	x2, x0, x1
		0xd503201f, //	nop
		0x913ffc02, //	add	x2, x0, #0xfff
		0xd13ffc02, //	sub	x2, x0, #0xfff
		0x92402c02, //	and	x2, x0, #0xfff
		0xb2402c02, //	orr	x2, x0, #0xfff
		0xd2402c02, //	eor	x2, x0, #0xfff
		0xd65f03c0, //	ret
	)

	if !SameCode(actual, expected) {
		t.Errorf("miscompiled code:\n\texpected %#v\n\tactual   %#v",
			expected, actual)
	}
}

func TestArm64Cast(t *testing.T) {
	var asm Asm

	id := RLo
	asm.Init()

	for _, skind := range [...]Kind{
		Int8, Int16, Int32, Int64,
		Uint8, Uint16, Uint32, Uint64,
	} {
		src := MakeReg(id, skind)
		for _, dkind := range [...]Kind{Uint8, Uint16, Uint32, Uint64} {
			dst := MakeReg(id, dkind)
			asm.Asm(CAST, src, dst)
		}
	}

	actual := asm.Code()
	expected := MakeCode(
		0x13001c00, // sxtb	w0, w0
		0x13001c00, // sxtb	w0, w0
		0x93401c00, // sxtb	x0, w0
		0x13003c00, // sxth	w0, w0
		0x93403c00, // sxth	x0, w0
		0x93407c00, // sxtw	x0, w0
		0x12001c00, // and	w0, w0, #0xff
		0x12001c00, // and	w0, w0, #0xff
		0x92401c00, // and	x0, x0, #0xff
		0x12003c00, // and	w0, w0, #0xffff
		0x92403c00, // and	x0, x0, #0xffff
		0x2a0003e0, // mov	w0, w0
	)

	if !SameCode(actual, expected) {
		t.Errorf("miscompiled code:\n\texpected %#v\n\tactual   %#v",
			expected, actual)
	}
}
