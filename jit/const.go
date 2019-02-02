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
 * const.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"github.com/cosmos72/gomacro/jit/arch"
)

type Const struct {
	val  int64
	kind arch.Kind
}

func ConstInt64(val int64) *Const {
	return &Const{val: val, kind: arch.Int64}
}

// implement Arg interface
func (c Const) Reg(asm *Asm) arch.Reg {
	return arch.Reg{}
}

func (c Const) Kind() arch.Kind {
	return c.kind
}

func (c Const) Const() bool {
	return true
}
