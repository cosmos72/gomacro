/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
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
	"fmt"
)

type Stmt interface {
	stmt()
}

// unary statement X Inst,
// for example X++
type Stmt1 struct {
	Dst  Expr
	Inst Inst1
}

// binary statement X Inst Y,
// for example X += Y
type Stmt2 struct {
	Dst  Expr
	Src  Expr
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
func (t *Stmt1) stmt() {
}

func (t *Stmt2) stmt() {
}

func (t *StmtN) stmt() {
}

func (t *Stmt1) String() string {
	switch t.Inst {
	case NOP:
		return fmt.Sprintf("_ = %v;", t.Dst)
	default:
		return fmt.Sprintf("%v%v;", t.Dst, t.Inst)
	}
}

func (t *Stmt2) String() string {
	return fmt.Sprintf("%v %v %v;", t.Dst, t.Inst, t.Src)
}

// compile statement
func (c *Comp) Stmt(t Stmt) {
	switch t := t.(type) {
	case *Stmt1:
		c.Stmt1(t.Inst, t.Dst)
	case *Stmt2:
		c.Stmt2(t.Inst, t.Dst, t.Src)
	case *StmtN:
		c.StmtN(t.Dst, t.Src)
	default:
		errorf("unknown Stmt type %T: %v", t, t)
	}
}

// compile unary statement
func (c *Comp) Stmt1(inst Inst1, tdst Expr) {
	dst, soft := c.Expr(tdst)
	if inst != NOP {
		checkAssignable(dst)
	}
	c.code.Inst1(inst, dst)
	c.FreeSoftReg(soft)
}

// compile binary statement
func (c *Comp) Stmt2(inst Inst2, tdst Expr, tsrc Expr) {
	// evaluate left-hand side first
	dst, dsoft := c.Expr(tdst)
	checkAssignable(dst)
	src, ssoft := c.expr(tsrc, dst)
	c.code.Inst2(inst, src, dst)
	c.FreeSoftReg(dsoft)
	if ssoft.id != dsoft.id {
		c.FreeSoftReg(ssoft)
	}
}

// compile n-ary statement
func (c *Comp) StmtN(tdst []Expr, tsrc []Expr) {
	n := len(tdst)
	if n != len(tsrc) {
		errorf("assignment mismatch: %d variables but %d values: %v = %v", n, len(tsrc), tdst, tsrc)
	}
	dst := make([]Expr, n)
	src := make([]Expr, n)
	// evaluate left-hand side first
	for i, x := range tdst {
		e, _ := c.Expr(x)
		checkAssignable(e)
		dst[i] = e
	}
	for i, x := range tsrc {
		e, soft := c.Expr(x)
		if _, ok := e.(Mem); ok && !soft.Valid() {
			// source is a local variable. we must evaluate it,
			// in case it also appears in left-hand side
			soft = c.AllocSoftReg(e.Kind())
			c.code.Inst2(ASSIGN, e, soft)
			e = soft
		}
		src[i] = e
	}
	for i := range src {
		c.code.Inst2(ASSIGN, src[i], dst[i])
	}
	for i := n - 1; i >= 0; i-- {
		if soft, ok := src[i].(SoftReg); ok && soft.Valid() {
			c.FreeSoftReg(soft)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if soft, ok := dst[i].(SoftReg); ok && soft.Valid() {
			c.FreeSoftReg(soft)
		}
	}
}
