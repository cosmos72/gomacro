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
 * op1.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// ============================================================================
// unary operation
type Op1 uint8

const (
/*
	NOT Op1 = 0x10
	NEG Op1 = 0x18
	INC Op1 = 0x20
	DEC Op1 = 0x28
*/
)

var op1Name = map[Op1]string{
	/*
		NOT: "NOT",
		NEG: "NEG",
		INC: "INC",
		DEC: "DEC",
	*/
}

func (op Op1) String() string {
	s, ok := op1Name[op]
	if !ok {
		s = fmt.Sprintf("Op1(%d)", int(op))
	}
	return s
}

// ============================================================================
func (asm *Asm) Op1(op Op1, a Arg) *Asm {
	errorf("Op1 not implemented: %v, %v", op, a)
	return asm
}
