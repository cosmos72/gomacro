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
 * op2.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// ============================================================================
// binary operation
type Op2 uint8

const (
/*
	ADD Op2 = 0
	OR  Op2 = 0x08
	ADC Op2 = 0x10 // add with carry
	SBB Op2 = 0x18 // subtract with borrow
	AND Op2 = 0x20
	SUB Op2 = 0x28
	XOR Op2 = 0x30
	// CMP Op = 0x38 // compare, set flags
	// XCHG Op = 0x86 // exchange
	MOV  Op2 = 0x88
	LEA  Op2 = 0x8D
	CAST Op2 = 0xB6 // sign extend, zero extend or narrow
	SHL  Op2 = 0xE0 // shift left
	SHR  Op2 = 0xE8 // shift right
	MUL  Op2 = 0xF6
	DIV  Op2 = 0xFE // divide
	REM  Op2 = 0xFF // remainder
*/
)

var op2Name = map[Op2]string{
	/*
		ADD: "ADD",
		OR:  "OR",
		ADC: "ADC",
		SBB: "SBB",
		AND: "AND",
		SUB: "SUB",
		XOR: "XOR",
		// CMP: "CMP",
		// XCHG: "XCHG",
		MOV:  "MOV",
		CAST: "CAST",
		MUL:  "MUL",
		DIV:  "DIV",
		REM:  "REM",
	*/
}

func (op Op2) String() string {
	s, ok := op2Name[op]
	if !ok {
		s = fmt.Sprintf("Op2(%d)", int(op))
	}
	return s
}

// ============================================================================
func (asm *Asm) Mov(src Arg, dst Arg) *Asm {
	// return asm.Op2(MOV, src, dst)
	errorf("Mov not implemented: %v %v", src, dst)
	return asm
}

func (asm *Asm) Op2(op Op2, src Arg, dst Arg) *Asm {
	errorf("Op2 not implemented: %v %v %v", op, src, dst)
	return asm
}
