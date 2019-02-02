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
 * op0.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// ============================================================================
// no-arg operation
type Op0 uint8

const (
/*
	RET Op0 = 0xC3
	NOP Op0 = 0x90
*/
)

var op0Name = map[Op0]string{
	/*
		RET: "RET",
		NOP: "NOP",
	*/
}

func (op Op0) String() string {
	s, ok := op0Name[op]
	if !ok {
		s = fmt.Sprintf("Op0(%d)", int(op))
	}
	return s
}

// ============================================================================
func (asm *Asm) Op0(op Op0) *Asm {
	return asm.Byte(uint8(op))
}
