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
	NOP Op0 = 0xD5
	RET Op0 = 0xD6
)

var op0Name = map[Op0]string{
	RET: "RET",
	NOP: "NOP",
}

func (op Op0) String() string {
	s, ok := op0Name[op]
	if !ok {
		s = fmt.Sprintf("Op0(%d)", int(op))
	}
	return s
}

func (op Op0) val() uint32 {
	switch op {
	case NOP:
		return 0xD503201F
	case RET:
		return 0xD65F03C0
	default:
		return uint32(op) << 24
	}
}

// ============================================================================
func (asm *Asm) Op0(op Op0) *Asm {
	return asm.Uint32(op.val())
}
