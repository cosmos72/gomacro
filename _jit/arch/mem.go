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
 * mem.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// hardware memory location.
type Mem struct {
	off int32
	reg Reg // also defines width and signedness
}

func (m Mem) String() string {
	return fmt.Sprintf("[%v+%v]/*%v*/", m.reg.id, m.off, m.reg.kind)
}

// implement Arg interface
func (m Mem) RegId() RegId {
	return NoRegId
}

func (m Mem) Kind() Kind {
	return m.reg.kind
}

func (m Mem) Const() bool {
	return false
}

func MakeVar(index uint16) Mem {
	return Mem{off: int32(index) * 8, reg: Reg{id: RDI, kind: Int64}}
}
