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
 * softreg.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	arch "github.com/cosmos72/gomacro/jit/native"
)

// subset of Arg interface
type SoftReg struct {
	id   arch.SoftReg
	kind Kind
}

func (s SoftReg) Kind() Kind {
	return s.kind
}

func (s SoftReg) Const() bool {
	return false
}

func (s SoftReg) Valid() bool {
	return s.kind != Invalid
}
