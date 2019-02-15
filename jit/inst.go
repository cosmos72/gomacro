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
 * inst.go
 *
 *  Created on Feb 15, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"fmt"

	"github.com/cosmos72/gomacro/jit/asm"
)

type Inst1 uint8 // unary statement operator
type Inst2 uint8 // binary statement operator

const (
	INC  = Inst1(asm.INC)  // ++
	DEC  = Inst1(asm.DEC)  // --
	ZERO = Inst1(asm.ZERO) // = 0
	NOP  = Inst1(asm.NOP)  // used to wrap an expression into a statement

	ASSIGN         = Inst2(asm.MOV)
	ADD_ASSIGN     = Inst2(asm.ADD2)
	SUB_ASSIGN     = Inst2(asm.SUB2)
	MUL_ASSIGN     = Inst2(asm.MUL2)
	QUO_ASSIGN     = Inst2(asm.DIV2)
	REM_ASSIGN     = Inst2(asm.REM2)
	AND_ASSIGN     = Inst2(asm.AND2)
	OR_ASSIGN      = Inst2(asm.OR2)
	XOR_ASSIGN     = Inst2(asm.XOR2)
	SHL_ASSIGN     = Inst2(asm.SHL2)
	SHR_ASSIGN     = Inst2(asm.SHR2)
	AND_NOT_ASSIGN = Inst2(asm.AND_NOT2)
	LAND_ASSIGN    = Inst2(asm.LAND2)
	LOR_ASSIGN     = Inst2(asm.LOR2)
)

var inst1name = map[Inst1]string{
	INC:  "++",
	DEC:  "--",
	ZERO: " = 0",
	NOP:  "",
}

var inst2name = map[Inst2]string{
	ASSIGN:         "=",
	ADD_ASSIGN:     "+=",
	SUB_ASSIGN:     "-=",
	MUL_ASSIGN:     "*=",
	QUO_ASSIGN:     "/=",
	REM_ASSIGN:     "%=",
	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",
	LAND_ASSIGN:    "&&=",
	LOR_ASSIGN:     "||=",
}

func (inst Inst1) Valid() bool {
	_, ok := inst1name[inst]
	return ok
}

func (inst Inst1) String() string {
	s, ok := inst1name[inst]
	if !ok {
		s = fmt.Sprintf("Inst1(%d)", uint8(inst))
	}
	return s
}

func (inst Inst2) Valid() bool {
	_, ok := inst2name[inst]
	return ok
}

func (inst Inst2) String() string {
	s, ok := inst2name[inst]
	if !ok {
		s = fmt.Sprintf("Inst2(%d)", uint8(inst))
	}
	return s
}
