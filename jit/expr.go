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
 * expr.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"fmt"
)

// subset of Arg interface
type Expr interface {
	Kind() Kind
	Const() bool
}

// unary expression OP X
type Expr1 struct {
	X  Expr
	Op Op1
	K  Kind
}

// binary expression X OP Y
type Expr2 struct {
	X  Expr
	Y  Expr
	Op Op2
	K  Kind
}

func NewExpr1(op Op1, x Expr) *Expr1 {
	kind := x.Kind()
	if op.IsCast() {
		// cast Ops have the same values
		// as the corresponding Kind
		kind = Kind(op)
	}
	return &Expr1{x, op, kind}
}

func NewExpr2(op Op2, x Expr, y Expr) *Expr2 {
	return &Expr2{x, y, op, x.Kind()}
}

// implement Expr interface
func (e *Expr1) Kind() Kind {
	return e.K
}

func (e *Expr1) Const() bool {
	return false
}

func (e *Expr1) String() string {
	if e.Op.IsCast() {
		return fmt.Sprintf("%v(%v)", e.Op, e.X)
	}
	return fmt.Sprintf("(%v %v)", e.Op, e.X)
}

// implement Expr interface
func (e *Expr2) Kind() Kind {
	return e.K
}

func (e *Expr2) Const() bool {
	return false
}

func (e *Expr2) String() string {
	return fmt.Sprintf("(%v %v %v)", e.X, e.Op, e.Y)
}

func IsLeaf(e Expr) bool {
	switch e.(type) {
	case *Expr1, *Expr2:
		return false
	default:
		return true
	}
}

// compile expression
func (c *Comp) Expr(e Expr) (Expr, SoftReg) {
	var dst Expr
	var dstsoft SoftReg
	switch e := e.(type) {
	case *Expr1:
		return c.expr1(e)
	case *Expr2:
		return c.expr2(e)
	case Const, Reg, Mem:
		dst = e
	}
	return dst, dstsoft
}

// compile unary expression
func (c *Comp) expr1(e *Expr1) (Expr, SoftReg) {
	var dst Expr
	var dstsoft SoftReg
	src, soft1 := c.Expr(e.X)
	if soft1.Valid() {
		dstsoft = SoftReg{soft1.id, e.K}
	} else {
		dstsoft = c.AllocSoftReg(e.K)
	}
	dst = dstsoft
	c.code.Op1(e.Op, src, dstsoft)
	if soft1.Valid() && soft1.id != dstsoft.id {
		c.FreeSoftReg(soft1)
	}
	return dst, dstsoft
}

// compile binary expression
func (c *Comp) expr2(e *Expr2) (Expr, SoftReg) {
	var dst Expr
	var dstsoft SoftReg
	src1, soft1 := c.Expr(e.X)
	src2, soft2 := c.Expr(e.Y)
	if soft1.Valid() {
		dstsoft = SoftReg{soft1.id, e.K}
	} else if soft2.Valid() && e.Op.IsCommutative() {
		dstsoft = SoftReg{soft2.id, e.K}
	} else {
		dstsoft = c.AllocSoftReg(e.K)
	}
	dst = dstsoft
	c.code.Op2(e.Op, src1, src2, dstsoft)
	if soft1.Valid() && soft1.id != dstsoft.id {
		c.FreeSoftReg(soft1)
	}
	if soft2.Valid() && soft2.id != dstsoft.id {
		c.FreeSoftReg(soft2)
	}
	return dst, dstsoft
}
