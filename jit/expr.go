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
	return &Expr1{x, op, x.Kind()}
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

// implement Expr interface
func (e *Expr2) Kind() Kind {
	return e.K
}

func (e *Expr2) Const() bool {
	return false
}

func IsLeaf(e Expr) bool {
	switch e.(type) {
	case *Expr1, *Expr2:
		return false
	default:
		return true
	}
}
