// +build !amd64

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
 * arith_dummy.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

// %reg_z += a
func (asm *Asm) Add(z Reg, a Arg) *Asm {
	return asm
}

// %reg_z -= a
func (asm *Asm) Sub(z Reg, a Arg) *Asm {
	return asm
}

// %reg_z *= a
func (asm *Asm) Mul(z Reg, a Arg) *Asm {
	return asm
}

// %reg_z /= a
func (asm *Asm) Quo(z Reg, a Arg) *Asm {
	return asm
}

// %reg %= a
func (asm *Asm) Rem(z Reg, a Arg) *Asm {
	return asm
}

// %reg_z = - %reg_z
func (asm *Asm) Neg(z Reg) *Asm {
	return asm
}
