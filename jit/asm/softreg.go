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

package asm

import (
	"fmt"
)

// soft register.
// may be mapped by assembler to an actual machine register
// or to a memory location
type SoftRegId uint32

const (
	// SLo = 0       // lowest SoftRegId available to user code
	SHi = 0x7FFFFFFF // highest SoftRegId available to user code
)

func (s SoftRegId) Validate() {
	if s > SHi {
		errorf("invalid register: %v", s)
	}
}

func (s SoftRegId) String() string {
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

// ===================================

type SoftRegIds map[SoftRegId]Arg // SoftRegId -> Arg
