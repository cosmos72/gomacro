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
 * place.go
 *
 *  Created on Mar 03, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

// used by fast/jit.go
type Place struct {
	Base   Expr
	Index  Expr
	tofree []SoftReg
}
