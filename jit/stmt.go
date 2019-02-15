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
 * stmt.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
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

func (inst Inst1) String() string {
	return asm.Op1(inst).String()
}

func (inst Inst2) String() string {
	return asm.Op2(inst).String()
}

type Stmt interface {
	stmt()
}

// unary statement X Inst,
// for example X++
type Stmt1 struct {
	Dst    Expr
	Inst Inst1
}

// binary statement X Inst Y,
// for example X += Y
type Stmt2 struct {
	Dst    Expr
	Src    Expr
	Inst Inst2
}

// N-ary assignment
// a,b,c... = p,q,r...
type StmtN struct {
	Dst []Expr
	Src []Expr
}

func NewStmt1(inst Inst1, dst Expr) *Stmt1 {
	return &Stmt1{dst, inst}
}

func NewStmt2(inst Inst2, dst Expr, src Expr) *Stmt2 {
	return &Stmt2{dst, src, inst}
}

func NewStmtN(dst []Expr, src []Expr) *StmtN {
	return &StmtN{dst, src}
}

// implement Stmt interface
func (stmt *Stmt1) stmt() {
}

func (stmt *Stmt2) stmt() {
}

func (stmt *StmtN) stmt() {
}
