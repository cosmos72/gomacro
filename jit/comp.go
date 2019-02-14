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
 * comp.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"github.com/cosmos72/gomacro/jit/asm"
)

type Comp struct {
	code        Code
	nextSoftReg SoftRegId
	arch        Arch
	asm.RegIdCfg
}

func New() *Comp {
	var c Comp
	return c.Init()
}

func NewArchId(archId ArchId) *Comp {
	var c Comp
	return c.InitArchId(archId)
}

func NewArch(arch Arch) *Comp {
	var c Comp
	return c.InitArch(arch)
}

func (c *Comp) Init() *Comp {
	return c.InitArchId(asm.ARCH_ID)
}

func (c *Comp) InitArchId(archId ArchId) *Comp {
	return c.InitArch(Archs[archId])
}

func (c *Comp) InitArch(arch Arch) *Comp {
	if arch == nil {
		errorf("unknown arch")
	}
	c.code = nil
	c.nextSoftReg = 0
	c.arch = arch
	c.RegIdCfg = arch.RegIdCfg()
	return c
}

func (c *Comp) Arch() Arch {
	return c.arch
}

func (c *Comp) ArchId() ArchId {
	if c.arch == nil {
		return asm.NOARCH
	}
	return c.arch.Id()
}

func (c *Comp) NewAsm() *Asm {
	return asm.NewArch(c.arch)
}

func (c *Comp) Code() Code {
	return c.code
}

func (c *Comp) AllocSoftReg(kind Kind) SoftReg {
	id := c.nextSoftReg
	c.nextSoftReg++
	return c.code.SoftReg(asm.ALLOC, id, kind)
}

func (c *Comp) FreeSoftReg(s SoftReg) {
	if s.id+1 == c.nextSoftReg {
		c.nextSoftReg--
	}
	c.code.SoftReg(asm.FREE, s.id, s.kind)
}

func (c *Comp) Stmt(t Stmt) {
	switch t := t.(type) {
	case *Stmt1:
		errorf("unimplemented: Stmt(Stmt1)")
	case *Stmt2:
		errorf("unimplemented: Stmt(Stmt2)")
	default:
		errorf("unknown Stmt type %T: %v", t, t)
	}
}

func (c *Comp) Expr(e Expr) (Expr, SoftReg) {
	var dst Expr
	var dstsoft SoftReg
	switch e := e.(type) {
	case *Expr1:
		src, soft1 := c.Expr(e.x)
		if soft1.Valid() {
			dstsoft = SoftReg{soft1.id, e.kind}
		} else {
			dstsoft = c.AllocSoftReg(e.kind)
		}
		dst = dstsoft
		c.code.Op1(e.op, src, dstsoft)
		if soft1.Valid() && soft1.id != dstsoft.id {
			c.FreeSoftReg(soft1)
		}
	case *Expr2:
		src1, soft1 := c.Expr(e.x)
		src2, soft2 := c.Expr(e.y)
		if soft1.Valid() {
			dstsoft = SoftReg{soft1.id, e.kind}
		} else if soft2.Valid() && e.op.IsCommutative() {
			dstsoft = SoftReg{soft2.id, e.kind}
		} else {
			dstsoft = c.AllocSoftReg(e.kind)
		}
		dst = dstsoft
		c.code.Op2(e.op, src1, src2, dstsoft)
		if soft1.Valid() && soft1.id != dstsoft.id {
			c.FreeSoftReg(soft1)
		}
		if soft2.Valid() && soft2.id != dstsoft.id {
			c.FreeSoftReg(soft2)
		}
	case Const, Reg, Mem:
		dst = e
	}
	return dst, dstsoft
}
