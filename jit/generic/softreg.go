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
 *  Created on Feb 09, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// soft register.
// may be mapped by assembler to an actual machine register
// or to a memory location
type SoftReg uint32

const (
	// SLo = 0       // lowest SoftReg available to user code
	SHi = 0x7FFFFFFF // highest SoftReg available to user code
)

func (s SoftReg) Validate() {
	if s > SHi {
		errorf("invalid register: %v", s)
	}
}

func (s SoftReg) String() string {
	return fmt.Sprintf("S%d", uint32(s))
}

// use Asm to convert softreg to Reg or Mem
func (s SoftReg) Arg(asm *Asm) Arg {
	return asm.SoftReg(s)
}

func (s SoftReg) RegId(asm *Asm) RegId {
	return asm.SoftReg(s).RegId()
}

func (s SoftReg) Kind(asm *Asm) Kind {
	return asm.SoftReg(s).Kind()
}

// ===================================

type SoftRegs map[SoftReg]Arg // SoftReg -> Arg
