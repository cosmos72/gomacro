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
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

type Size uint8 // 1, 2, 4 or 8

type Arg interface {
	RegId() RegId // NoReg if not a register
	Kind() Kind
	Const() bool
}

type Code []uint8

type Save struct {
	start, idx, end uint16 // memory area where spill registers can be saved
}

type Asm struct {
	code      Code
	RegIds      RegIds
	NextRegId RegId // first available register among usable ones
	Save      Save
}

func SizeOf(a Arg) Size {
	size := a.Kind().Size()
	if size == 0 {
		errorf("unsupported Kind %v", a.Kind())
	}
	return size
}
