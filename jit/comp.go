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
	if s.Valid() {
		if s.id+1 == c.nextSoftReg {
			c.nextSoftReg--
		}
		c.code.SoftReg(asm.FREE, s.id, s.kind)
	}
}

func isMem(e Expr) bool {
	_, ok := e.(Mem)
	return ok
}

func (c *Comp) Expr(e Expr) (Expr, SoftReg) {
	var dst Expr
	var dstsoft SoftReg
	switch e := e.(type) {
	case *Expr1:
		src, soft1 := c.Expr(e.X)
		if soft1.Valid() {
			dstsoft = SoftReg{soft1.id, e.K}
		} else {
			dstsoft = c.AllocSoftReg(e.K)
		}
		dst = dstsoft
		c.code.Op1(e.Op, src, dstsoft)
		if soft1.Valid() && soft1.id != dstsoft.id {
			c.FreeSoftReg(soft1)
		}
	case *Expr2:
		src1, soft1 := c.Expr(e.X)
		src2, soft2 := c.Expr(e.Y)
		if soft1.Valid() {
			dstsoft = SoftReg{soft1.id, e.K}
		} else if soft2.Valid() && e.Op.IsCommutative() {
			dstsoft = SoftReg{soft2.id, e.K}
		} else {
			dstsoft = c.AllocSoftReg(e.K)
		}
		dst = dstsoft
		c.code.Op2(e.Op, src1, src2, dstsoft)
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

func (c *Comp) Stmt(t Stmt) {
	switch t := t.(type) {
	case *Stmt1:
		dst, soft := c.Expr(t.Dst)
		if t.Inst != NOP && !isMem(dst) {
			errorf("cannot assign to %v", t.Dst)
		}
		c.code.Inst1(t.Inst, dst)
		c.FreeSoftReg(soft)
	case *Stmt2:
		// evaluate left-hand side first
		dst, dsoft := c.Expr(t.Dst)
		if !isMem(dst) {
			errorf("cannot assign to %v", t.Dst)
		}
		src, ssoft := c.Expr(t.Src)
		c.code.Inst2(t.Inst, src, dst)
		c.FreeSoftReg(dsoft)
		c.FreeSoftReg(ssoft)
	case *StmtN:
		n := len(t.Dst)
		if n != len(t.Src) {
			errorf("assignment mismatch: %d variables but %d values: %v", n, len(t.Src), t)
		}
		dst := make([]Expr, n)
		src := make([]Expr, n)
		// evaluate left-hand side first
		for i, x := range t.Dst {
			e, _ := c.Expr(x)
			if !isMem(e) {
				errorf("cannot assign to %v", x)
			}
			dst[i] = e
		}
		for i, x := range t.Src {
			e, soft := c.Expr(x)
			if _, ok := e.(Mem); ok && !soft.Valid() {
				// source is a local variable. we must evaluate it,
				// in case it also appears in left-hand side
				soft = c.AllocSoftReg(e.Kind())
				c.code.Inst2(ASSIGN, e, soft)
				e = soft
			}
			src[i] = e
		}
		for i := range src {
			c.code.Inst2(ASSIGN, src[i], dst[i])
		}
		for i := n - 1; i >= 0; i-- {
			if soft, ok := src[i].(SoftReg); ok && soft.Valid() {
				c.FreeSoftReg(soft)
			}
		}
		for i := n - 1; i >= 0; i-- {
			if soft, ok := dst[i].(SoftReg); ok && soft.Valid() {
				c.FreeSoftReg(soft)
			}
		}
	default:
		errorf("unknown Stmt type %T: %v", t, t)
	}
}

// compile list of Stmt
func (c *Comp) Compile(ts ...Stmt) {
	for _, t := range ts {
		c.Stmt(t)
	}
}
