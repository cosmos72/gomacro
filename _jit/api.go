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
 * api.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"reflect"

	"github.com/cosmos72/gomacro/_jit/arch"
)

// software-defined register. mapped to hardware register by Asm
type Reg uint32

type Const struct {
	kind reflect.Kind
	val  int64
}

type Var struct {
	kind reflect.Kind
	idx  uint16
	upn  uint16
}

type Arg interface {
	reg(asm *Asm) arch.Reg // noReg if not a register
	Kind() reflect.Kind
	Const() bool
}

type Asm struct {
	asm     arch.Asm
	Regs    map[Reg]arch.Reg
	NextReg Reg // first available register among usable ones
}
