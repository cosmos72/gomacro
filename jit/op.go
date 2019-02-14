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
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"github.com/cosmos72/gomacro/jit/asm"
)

type Op1 uint8 // unary expression operator
type Op2 uint8 // binary expression operator

type Inst1 uint8 // unary statement operator
type Inst2 uint8 // binary statement operator

const (
	ADD = Op2(asm.ADD3)
	SUB = Op2(asm.SUB3)
	MUL = Op2(asm.MUL3)
	QUO = Op2(asm.DIV3)
	REM = Op2(asm.REM3)
	AND = Op2(asm.AND3)
	OR  = Op2(asm.OR3)
	XOR = Op2(asm.XOR3)
	SHL = Op2(asm.SHL3)
	SHR = Op2(asm.SHR3)
	/*
		AND_NOT = Op2(asm.AND_NOT) // &^
		LAND    = Op2(asm.LAND)    // &&
		LOR     = Op2(asm.LOR)     // ||
		EQL     = Op2(asm.EQL)
		LSS     = Op2(asm.LSS)
		GTR     = Op2(asm.GTR)
		NEQ     = Op2(asm.NEQ)
		LEQ     = Op2(asm.LEQ)
		GEQ     = Op2(asm.GEQ)
	*/

	NEG = Op1(asm.NEG2)
	NOT = Op1(asm.NOT2)

	INC  = Inst1(asm.INC)  // ++
	DEC  = Inst1(asm.DEC)  // --
	ZERO = Inst1(asm.ZERO) // = 0

	ASSIGN     = Inst2(asm.MOV)
	ADD_ASSIGN = Inst2(asm.ADD)
	SUB_ASSIGN = Inst2(asm.SUB)
	MUL_ASSIGN = Inst2(asm.MUL)
	QUO_ASSIGN = Inst2(asm.DIV)
	REM_ASSIGN = Inst2(asm.REM)
	AND_ASSIGN = Inst2(asm.AND)
	OR_ASSIGN  = Inst2(asm.OR)
	XOR_ASSIGN = Inst2(asm.XOR)
	SHL_ASSIGN = Inst2(asm.SHL)
	SHR_ASSIGN = Inst2(asm.SHR)
	/*
		AND_NOT_ASSIGN = Inst2(asm.AND_NOT)
	*/
)

func (op Op1) String() string {
	return asm.Op2(op).String()
}

func (op Op2) String() string {
	return asm.Op3(op).String()
}

func (op Op2) IsCommutative() bool {
	return asm.Op3(op).IsCommutative()
}

func (inst Inst1) String() string {
	return asm.Op1(inst).String()
}

func (inst Inst2) String() string {
	return asm.Op2(inst).String()
}
