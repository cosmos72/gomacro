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
 * softreg.go
 *
 *  Created on Feb 09, 2019
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"fmt"
)

// soft register.
// may be mapped by assembler to an actual machine register
// or to a memory location
type SoftRegId uint32

const (
	FirstTempRegId = ^SoftRegId(0) >> 1
)

func (s SoftRegId) Validate() {
}

func (s SoftRegId) String() string {
	if s >= FirstTempRegId {
		return fmt.Sprintf("T%d", uint32(s-FirstTempRegId))
	}
	return fmt.Sprintf("S%d", uint32(s))
}

// use Asm to convert softreg to Reg or Mem
func (s SoftRegId) Arg(asm *Asm) Arg {
	return asm.SoftRegId(s)
}

func (s SoftRegId) RegId(asm *Asm) RegId {
	return asm.SoftRegId(s).RegId()
}

func (s SoftRegId) Kind(asm *Asm) Kind {
	return asm.SoftRegId(s).Kind()
}

// implement AsmCode interface
func (s SoftRegId) asmcode() {
}

// ===================================

type SoftRegIds map[SoftRegId]Arg // SoftRegId -> Arg
