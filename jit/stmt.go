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
		c.stmt1(t)
	case *Stmt2:
		c.stmt2(t)
	case *StmtN:
		c.stmtN(t)
	default:
		errorf("unknown Stmt type %T: %v", t, t)
	}
}

// compile unary statement
func (c *Comp) stmt1(t *Stmt1) {
	dst, soft := c.Expr(t.Dst)
	if t.Inst != NOP {
		checkAssignable(dst)
	}
	c.code.Inst1(t.Inst, dst)
	c.FreeSoftReg(soft)
}

// compile binary statement
func (c *Comp) stmt2(t *Stmt2) {
	// evaluate left-hand side first
	dst, dsoft := c.Expr(t.Dst)
	checkAssignable(dst)
	src, ssoft := c.expr(t.Src, dst)
	c.code.Inst2(t.Inst, src, dst)
	c.FreeSoftReg(dsoft)
	if ssoft.id != dsoft.id {
		c.FreeSoftReg(ssoft)
	}
}

// compile n-ary statement
func (c *Comp) stmtN(t *StmtN) {
	n := len(t.Dst)
	if n != len(t.Src) {
		errorf("assignment mismatch: %d variables but %d values: %v", n, len(t.Src), t)
	}
	dst := make([]Expr, n)
	src := make([]Expr, n)
	// evaluate left-hand side first
	for i, x := range t.Dst {
		e, _ := c.Expr(x)
		checkAssignable(e)
		dst[i] = e
	}
	for i, x := range t.Src {
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
