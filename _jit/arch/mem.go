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
	reg Reg // also defines kind, width and signedness
}

func (m Mem) String() string {
	return fmt.Sprintf("%v[%v+%v]", m.reg.kind, m.reg.id, m.off)
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

func MakeVar0(index uint16) Mem {
	return Mem{off: int32(index) * 8, reg: Reg{id: RDI, kind: Int64}}
}

func MakeVar0K(index uint16, k Kind) Mem {
	return Mem{off: int32(index) * 8, reg: Reg{id: RDI, kind: k}}
}

func MakeMem(off int32, id RegId, k Kind) Mem {
	return Mem{off: off, reg: Reg{id: id, kind: k}}
}
