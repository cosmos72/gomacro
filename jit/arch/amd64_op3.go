// +build amd64

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
 * amd64_op3.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// ============================================================================
// ternary operation
type Op3 uint8

const (
	ADD3  Op3 = Op3(ADD)
	OR3   Op3 = Op3(OR)
	ADC3  Op3 = Op3(ADC) // add with carry
	SBB3  Op3 = Op3(SBB) // subtract with borrow
	AND3  Op3 = Op3(AND)
	SUB3  Op3 = Op3(SUB)
	XOR3  Op3 = Op3(XOR)
	CAST3 Op3 = Op3(CAST) // sign extend, zero extend or narrow
	SHL3  Op3 = Op3(SHL)  // shift left
	SHR3  Op3 = Op3(SHR)  // shift right
	MUL3  Op3 = Op3(MUL)
	DIV3  Op3 = Op3(DIV)
	REM3  Op3 = Op3(REM)
)

func (op Op3) String() string {
	s, ok := op2Name[Op2(op)]
	if ok {
		s += "3"
	} else {
		s = fmt.Sprintf("Op3(%d)", int(op))
	}
	return s
}

func (op Op3) isCommutative() bool {
	switch op {
	case ADD3, OR3, ADC3, AND3, XOR3, MUL3:
		return true
	default:
		return false
	}
}

// ============================================================================
func (asm *Asm) Op3(op Op3, a Arg, b Arg, dst Arg) *Asm {
	op2 := Op2(op)
	if a == dst {
		asm.Op2(op2, b, dst)
	} else if op.isCommutative() && b == dst {
		asm.Op2(op2, a, dst)
	} else if r, ok := dst.(Reg); ok && r.id != dst.UsedRegId() {
		asm.Mov(a, dst).Op2(op2, b, dst)
	} else {
		r := asm.RegAlloc(dst.Kind())
		asm.Mov(a, r).Op2(op2, b, r).Mov(r, dst)
		asm.RegFree(r)
	}
	return asm
}
