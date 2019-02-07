// +build arm64

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
 * znativerun_test.go
 *
 *  Created on Feb 07, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"testing"
)

func TestArm64Sample(t *testing.T) {
	var asm Asm

	xzr := MakeReg(XZR, Uint64)
	asm.Init()
	asm.Asm( //
		MOV, xzr, MakeMem(8, XSP, Uint64),
		RET
	)

}
