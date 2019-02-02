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
 * arg.go
 *
 *  Created on Jan 29, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	arch "github.com/cosmos72/gomacro/jit/native"
)

type Arg interface {
	Reg(asm *Asm) arch.Reg // arch.Reg{} if not a register
	Kind(asm *Asm) arch.Kind
	Const() bool
}
