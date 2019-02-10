// +build !amd64,!arm64

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
 * asm_generic.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	arch "github.com/cosmos72/gomacro/jit/native"
)

const (
	// register pointing to local variables
	RVAR = arch.NoRegId

	// register pointing to stack
	RSP = arch.RSP
)

// local variable
func MakeVar(idx uint16, kind Kind) Mem {
	// TODO support upn > 0
	return Mem{}
}

// function parameter or return value
func MakeParam(off int32, kind Kind) Mem {
	return Mem{}
}
