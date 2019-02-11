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
 * op.go
 *
 *  Created on Feb 11, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
	"go/token"
)

// ============================================================================
// no-arg instruction
type Op0 uint8

const (
	NOP = Op0(token.SEMICOLON) // somewhat arbitrary choice
	RET = Op0(token.RETURN)
)

var op0Name = map[Op0]string{
	RET: "RET",
	NOP: "NOP",
}

func (op Op0) String() string {
	s, ok := op0Name[op]
	if !ok {
		s = fmt.Sprintf("Op0(%d)", uint8(op))
	}
	return s
}

// ============================================================================
// one-arg instruction
type Op1 uint8

const (
	ZERO = Op1(token.DEFAULT) // somewhat arbitrary choice
	INC  = Op1(token.INC)     // ++
	DEC  = Op1(token.DEC)     // --
	NEG  = Op1(1)             // - // avoid conflict between NEG2 and SUB
	NOT  = Op1(2)             // ^ // avoid conflict between NOT2 and XOR
)

var op1Name = map[Op1]string{
	ZERO: "ZERO",
	INC:  "INC",
	DEC:  "DEC",
	NOT:  "NOT",
	NEG:  "NEG",
}

func (op Op1) String() string {
	s, ok := op1Name[op]
	if !ok {
		s = fmt.Sprintf("Op1(%d)", uint8(op))
	}
	return s
}
