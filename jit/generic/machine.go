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
 * machine.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package arch

const (
	SUPPORTED = false
	Name      = "generic"

	NoRegId RegId = iota
	RLo           = NoRegId
	RHi           = NoRegId
)

func (r RegId) Valid() bool {
	return false
}

var alwaysLiveRegIds RegIds // empty

type Op0 struct{}
type Op1 struct{}
type Op2 struct{}
type Op3 struct{}
type Op4 struct{}

func (asm *Asm) Op0(op Op0) *Asm {
	return asm
}

func (asm *Asm) Op1(op Op1, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Op2(op Op2, src Arg, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Op3(op Op3, a Arg, b Arg, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Op4(op Op4, a Arg, b Arg, c Arg, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Mov(src Arg, dst Arg) *Asm {
	return asm
}

func (asm *Asm) Prologue() *Asm {
	return asm
}

func (asm *Asm) Epilogue() *Asm {
	return asm
}

func (asm *Asm) archPush(id RegId) {
}

func (asm *Asm) archPop(id RegId) {
}

func (s *Save) ArchInit(start uint16, end uint16) {
}
