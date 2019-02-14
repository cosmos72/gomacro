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
	arch "github.com/cosmos72/gomacro/jit/old/redirect"
)

type Op1 uint8 // unary expression operator
type Op2 uint8 // binary expression operator

type Inst1 uint8 // unary statement operator
type Inst2 uint8 // binary statement operator

const (
	ADD = Op2(arch.ADD3)
	SUB = Op2(arch.SUB3)
	MUL = Op2(arch.MUL3)
	QUO = Op2(arch.DIV3)
	REM = Op2(arch.REM3)
	AND = Op2(arch.AND3)
	OR  = Op2(arch.OR3)
	XOR = Op2(arch.XOR3)
	SHL = Op2(arch.SHL3)
	SHR = Op2(arch.SHR3)
	/*
		AND_NOT = Op2(arch.AND_NOT) // &^
		LAND    = Op2(arch.LAND)    // &&
		LOR     = Op2(arch.LOR)     // ||
		EQL     = Op2(arch.EQL)
		LSS     = Op2(arch.LSS)
		GTR     = Op2(arch.GTR)
		NEQ     = Op2(arch.NEQ)
		LEQ     = Op2(arch.LEQ)
		GEQ     = Op2(arch.GEQ)
	*/

	NEG = Op1(arch.NEG2)
	NOT = Op1(arch.NOT2)

	INC  = Inst1(arch.INC)  // ++
	DEC  = Inst1(arch.DEC)  // --
	ZERO = Inst1(arch.ZERO) // = 0

	ASSIGN     = Inst2(arch.MOV)
	ADD_ASSIGN = Inst2(arch.ADD)
	SUB_ASSIGN = Inst2(arch.SUB)
	MUL_ASSIGN = Inst2(arch.MUL)
	QUO_ASSIGN = Inst2(arch.DIV)
	REM_ASSIGN = Inst2(arch.REM)
	AND_ASSIGN = Inst2(arch.AND)
	OR_ASSIGN  = Inst2(arch.OR)
	XOR_ASSIGN = Inst2(arch.XOR)
	SHL_ASSIGN = Inst2(arch.SHL)
	SHR_ASSIGN = Inst2(arch.SHR)
	/*
		AND_NOT_ASSIGN = Inst2(arch.AND_NOT)
	*/
)

func (op Op1) String() string {
	return arch.Op2(op).String()
}

func (op Op2) String() string {
	return arch.Op3(op).String()
}

func (op Op2) IsCommutative() bool {
	return arch.Op3(op).IsCommutative()
}

func (inst Inst1) String() string {
	return arch.Op1(inst).String()
}

func (inst Inst2) String() string {
	return arch.Op2(inst).String()
}
