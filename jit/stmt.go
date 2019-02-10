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

// subset of Arg interface
type Stmt interface {
	stmt()
}

type Stmt1 struct {
	x    Expr
	inst Inst1
}

type Stmt2 struct {
	x    Expr
	y    Expr
	inst Inst2
}

func (stmt *Stmt1) stmt() {
}

func (stmt *Stmt2) stmt() {
}
