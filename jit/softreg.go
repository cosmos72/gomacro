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
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"fmt"
)

// SoftRegId wrapper, implements Expr
type SoftReg struct {
	id   SoftRegId
	kind Kind
}

func MakeSoftReg(id SoftRegId, kind Kind) SoftReg {
	return SoftReg{id, kind}
}

func (s SoftReg) Id() SoftRegId {
	return s.id
}

func (s SoftReg) Kind() Kind {
	return s.kind
}

func (s SoftReg) Const() bool {
	return false
}

func (s SoftReg) Valid() bool {
	return s.kind != Invalid
}

func (s SoftReg) Validate() {
	if !s.Valid() {
		errorf("invalid SoftReg: %v", s)
	}
}

func (s SoftReg) String() string {
	return fmt.Sprintf("T%d", uint32(s.id))

}
