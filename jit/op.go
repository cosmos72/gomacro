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
	"fmt"

	"github.com/cosmos72/gomacro/jit/asm"
)

type Op1 uint8 // unary expression operator
type Op2 uint8 // binary expression operator

const (
	// instead of a single CAST, we must implement
	// one Op1 per destination type:
	// INT8 ... INT64, UINT8 ... UINT64, etc.
	INT     = Op1(asm.Int)
	INT8    = Op1(asm.Int8)
	INT16   = Op1(asm.Int16)
	INT32   = Op1(asm.Int32)
	INT64   = Op1(asm.Int64)
	UINT    = Op1(asm.Uint)
	UINT8   = Op1(asm.Uint8)
	UINT16  = Op1(asm.Uint16)
	UINT32  = Op1(asm.Uint32)
	UINT64  = Op1(asm.Uint64)
	UINTPTR = Op1(asm.Uintptr)
	FLOAT32 = Op1(asm.Float32)
	FLOAT64 = Op1(asm.Float64)
	PTR     = Op1(asm.Ptr)
	NEG     = Op1(asm.NEG2)
	NOT     = Op1(asm.NOT2)

	ADD     = Op2(asm.ADD3)
	SUB     = Op2(asm.SUB3)
	MUL     = Op2(asm.MUL3)
	QUO     = Op2(asm.DIV3)
	REM     = Op2(asm.REM3)
	AND     = Op2(asm.AND3)
	OR      = Op2(asm.OR3)
	XOR     = Op2(asm.XOR3)
	SHL     = Op2(asm.SHL3)
	SHR     = Op2(asm.SHR3)
	AND_NOT = Op2(asm.AND_NOT3) // &^
	LAND    = Op2(asm.LAND3)    // &&
	LOR     = Op2(asm.LOR3)     // ||
	/*
		EQL     = Op2(asm.EQL3)
		LSS     = Op2(asm.LSS3)
		GTR     = Op2(asm.GTR3)
		NEQ     = Op2(asm.NEQ3)
		LEQ     = Op2(asm.LEQ3)
		GEQ     = Op2(asm.GEQ3)
	*/
)

var op1name = map[Op1]string{
	INT:     "int",
	INT8:    "int8",
	INT16:   "int16",
	INT32:   "int32",
	INT64:   "int64",
	UINT:    "uint",
	UINT8:   "uint8",
	UINT16:  "uint16",
	UINT32:  "uint32",
	UINT64:  "uint64",
	FLOAT32: "float32",
	FLOAT64: "float64",
	PTR:     "pointer",
	NEG:     "-",
	NOT:     "^",
}

var op2name = map[Op2]string{
	ADD:     "+",
	SUB:     "-",
	MUL:     "*",
	QUO:     "/",
	REM:     "%",
	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",
	LAND:    "&&",
	LOR:     "||",
	/*
		EQL    :"==",
		LSS    :"<",
		GTR    :"<",
		NEQ    :"!=",
		LEQ    :"<=",
		GEQ    :">=",
	*/
}

// =======================================================

func (op Op1) Valid() bool {
	_, ok := op1name[op]
	return ok
}

func (op Op1) Validate() {
	if !op.Valid() {
		errorf("unknown Op1: %v", op)
	}
}

func (op Op1) IsCast() bool {
	return op.Valid() && op >= INT && op <= PTR
}

// convert to asm.Op2
func (op Op1) Asm() asm.Op2 {
	op.Validate()
	if op.IsCast() {
		return asm.CAST
	}
	return asm.Op2(op)
}

func (op Op1) String() string {
	s, ok := op1name[op]
	if !ok {
		s = fmt.Sprintf("Op1(%d)", uint8(op))
	}
	return s
}

// =======================================================

func (op Op2) Valid() bool {
	_, ok := op2name[op]
	return ok
}

func (op Op2) Validate() {
	if !op.Valid() {
		errorf("unknown Op2: %v", op)
	}
}

// convert to asm.Op3
func (op Op2) Asm() asm.Op3 {
	op.Validate()
	return asm.Op3(op)
}

func (op Op2) String() string {
	s, ok := op2name[op]
	if !ok {
		s = fmt.Sprintf("Op2(%d)", uint8(op))
	}
	return s
}

func (op Op2) IsCommutative() bool {
	return op.Asm().IsCommutative()
}
