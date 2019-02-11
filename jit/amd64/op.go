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

// ============================================================================
// two-arg instruction
type Op2 uint8

const (
	ADD = Op2(token.ADD)
	SUB = Op2(token.SUB)
	ADC = Op2(token.ADD + token.VAR) // add with carry
	SBB = Op2(token.SUB + token.VAR) // subtract with borrow
	MUL = Op2(token.MUL)
	DIV = Op2(token.QUO) // divide
	QUO = DIV            // alias for DIV
	REM = Op2(token.REM) // remainder

	AND     = Op2(token.AND)
	OR      = Op2(token.OR)
	XOR     = Op2(token.XOR)
	SHL     = Op2(token.SHL)
	SHR     = Op2(token.SHR)
	AND_NOT = Op2(token.AND_NOT)

	LAND = Op2(token.LAND) // &&
	LOR  = Op2(token.LOR)  // ||

	MOV  = Op2(token.ASSIGN) // =
	CAST = Op2(token.TYPE)   // somewhat arbitrary choice
	LEA  = Op2(token.ARROW)  // amd64 only. somewhat arbitrary choice
	// CMP  = ??
	// XCHG = ??

	// two-arg versions of NOT, NEG above
	NEG2 = Op2(NEG)
	NOT2 = Op2(NOT)
)

var op2Name = map[Op2]string{
	ADD: "ADD",
	SUB: "SUB",
	ADC: "ADC",
	SBB: "SBB",
	MUL: "MUL",
	DIV: "DIV",
	REM: "REM",

	AND:     "AND",
	OR:      "OR",
	XOR:     "XOR",
	SHL:     "SHL",
	SHR:     "SHR",
	AND_NOT: "AND_NOT",

	LAND: "LAND",
	LOR:  "LOR",

	MOV:  "MOV",
	CAST: "CAST",
	LEA:  "LEA",
	// CMP:  "CMP",
	// XCHG: "XCHG",
	NEG2: "NEG2",
	NOT2: "NOT2",
}

func (op Op2) String() string {
	s, ok := op2Name[op]
	if !ok {
		s = fmt.Sprintf("Op2(%d)", int(op))
	}
	return s
}
