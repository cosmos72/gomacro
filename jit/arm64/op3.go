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
 * op3.go
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
/*
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
	DIV3
*/
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
	/*
		switch op {
		case ADD3, OR3, ADC3, AND3, XOR3, MUL3:
			return true
		}
	*/
	return false
}

// ============================================================================
func (asm *Asm) Op3(op Op3, a Arg, b Arg, dst Arg) *Asm {
	errorf("Op3 not implemented: %v %v %v %v", op, a, b, dst)
	return asm
}
