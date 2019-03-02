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

	"github.com/cosmos72/gomacro/jit/common"
)

// SoftRegId wrapper, implements Expr
type SoftReg struct {
	id   SoftRegId
	kind Kind
}

const FirstTempRegId = common.FirstTempRegId

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

func (s SoftReg) isTemp() bool {
	return s.id >= FirstTempRegId
}

func (s SoftReg) Validate() {
	if !s.Valid() {
		errorf("invalid SoftReg: %v", s)
	}
}

func (s SoftReg) String() string {
	var suffix string
	switch s.kind.Size() {
	case 0:
		suffix = "(bad)"
	case 1:
		suffix = "b"
	case 2:
		suffix = "h"
	case 4:
		suffix = "w"
	case 8:
	}
	if s.id >= FirstTempRegId {
		return fmt.Sprintf("t%d%s", uint32(s.id-FirstTempRegId), suffix)
	}
	return fmt.Sprintf("s%d%s", uint32(s.id), suffix)
}

// =======================================================

func (c *Comp) NewSoftReg(kind Kind) SoftReg {
	id := c.nextSoftReg
	c.nextSoftReg++
	return c.code.SoftReg(common.ALLOC, id, kind)
}

func (c *Comp) newTempReg(kind Kind) SoftReg {
	id := c.nextTempReg
	c.nextTempReg++
	return c.code.SoftReg(common.ALLOC, id, kind)
}

func (c *Comp) FreeSoftReg(s SoftReg) {
	if s.Valid() && !s.isTemp() {
		if s.id+1 == c.nextSoftReg {
			c.nextSoftReg--
		}
		c.code.SoftReg(common.FREE, s.id, s.kind)
	}
}

func (c *Comp) freeTempReg(s SoftReg) {
	if s.Valid() && s.isTemp() {
		if s.id+1 == c.nextTempReg {
			c.nextTempReg--
		}
		c.code.SoftReg(common.FREE, s.id, s.kind)
	}
}

// alloc or free soft reg
func (c *Comp) SoftReg(inst Inst1Misc, s SoftReg) {
	if s.Valid() {
		c.code.SoftReg(inst.Asm(), s.id, s.kind)
	}
}
