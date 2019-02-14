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

package amd64

// ============================================================================
// no-arg instruction

var op0Val = [256]uint8{
	RET: 0xC3,
	NOP: 0x90,
}

func op0val(op Op0) uint8 {
	val := op0Val[op]
	if val == 0 {
		errorf("unknown Op0 instruction: %v", op)
	}
	return val
}

// ============================================================================
func (arch Amd64) Op0(asm *Asm, op Op0) *Asm {
	return asm.Byte(op0val(op))
}
