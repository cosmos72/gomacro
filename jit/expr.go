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
	return c.expr(e, nil)
}

func (c *Comp) expr(e Expr, dst Expr) (Expr, SoftReg) {
	var dstsoft SoftReg
	switch e := e.(type) {
	case *Expr1:
		return c.expr1(e, dst)
	case *Expr2:
		return c.expr2(e, dst)
	case Const, Reg, Mem, SoftReg:
		dst = e
	default:
		errorf("unknown expression type %T: %v", e, e)
	}
	return dst, dstsoft
}

// compile unary expression
func (c *Comp) expr1(e *Expr1, dst Expr) (Expr, SoftReg) {
	dsoft, _ := dst.(SoftReg)
	var dto Expr
	if dsoft.Valid() {
		// forward the request to write into dsoft
		dto = dst
	}
	src, ssoft := c.expr(e.X, dto)
	if !dsoft.Valid() {
		if ssoft.Valid() {
			dsoft = SoftReg{ssoft.id, e.K}
		} else {
			dsoft = c.AllocSoftReg(e.K)
		}
	}
	c.code.Op1(e.Op, src, dsoft)
	if ssoft.id != dsoft.id {
		c.FreeSoftReg(ssoft)
	}
	if dst == nil {
		// no destination requested
		dst = dsoft
	} else if dsoft != dst {
		// copy dsoft to the requested destination
		// and free it
		c.code.Inst2(ASSIGN, dsoft, dst)
		c.FreeSoftReg(dsoft)
		dsoft = SoftReg{}
	}
	return dst, dsoft
}

// compile binary expression
func (c *Comp) expr2(e *Expr2, dst Expr) (Expr, SoftReg) {
	dsoft, _ := dst.(SoftReg)
	var dto Expr
	if dsoft.Valid() {
		// forward the request to write into dsoft
		dto = dst
	}
	src1, soft1 := c.expr(e.X, dto)
	src2, soft2 := c.Expr(e.Y)
	if !dsoft.Valid() {
		if soft1.Valid() {
			dsoft = SoftReg{soft1.id, e.K}
		} else if soft2.Valid() && e.Op.IsCommutative() {
			dsoft = SoftReg{soft2.id, e.K}
		} else {
			dsoft = c.AllocSoftReg(e.K)
		}
	}
	dst = dsoft
	c.code.Op2(e.Op, src1, src2, dsoft)
	if soft1.id != dsoft.id {
		c.FreeSoftReg(soft1)
	}
	if soft2.id != dsoft.id {
		c.FreeSoftReg(soft2)
	}
	if dst == nil {
		// no destination requested
		dst = dsoft
	} else if dsoft != dst {
		// copy dsoft to the requested destination
		// and free it
		c.code.Inst2(ASSIGN, dsoft, dst)
		c.FreeSoftReg(dsoft)
		dsoft = SoftReg{}
	}
	return dst, dsoft
}
