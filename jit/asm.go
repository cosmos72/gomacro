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
 * asm.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	arch "github.com/cosmos72/gomacro/jit/native"
)

const (
	VERBOSE = false
)

type regUse struct {
	reg   arch.Reg
	count uint32
}

type Asm struct {
	arch    arch.Asm
	regs    map[Reg]*regUse // regUse structs are shared among aliases
	nextReg Reg             // first available register among usable ones
}

func (asm *Asm) Reg(z Reg) arch.Reg {
	var r arch.Reg
	use := asm.regs[z]
	if use != nil {
		r = use.reg
	}
	return r
}

// allocate a jit-reserved register
func (asm *Asm) NewReg(kind arch.Kind) Reg {
	z := asm.nextReg
	asm.nextReg++
	asm.Alloc(z, kind)
	return z
}

func (asm *Asm) Alloc(z Reg, kind arch.Kind) *Asm {
	use := asm.regs[z]
	if use != nil {
		errorf("register %v is already allocated as %v, cannot allocate it again",
			z, use.reg)
	}
	asm.regs[z] = &regUse{
		reg:   asm.arch.RegAlloc(kind),
		count: 1,
	}
	return asm
}

// combined Alloc + Load
func (asm *Asm) AllocLoad(dst Reg, src Arg) *Asm {
	use := asm.regs[dst]
	if use != nil {
		errorf("register %v is already allocated as %v, cannot allocate it again",
			dst, use.reg)
	}
	if z, ok := src.(Reg); ok {
		use := asm.regs[z]
		if use == nil {
			errorf("register %v is not allocated, cannot use it", z)
		}
		if dst != z {
			// allocate dst as alias for z
			asm.regs[dst] = use
		}
		return asm
	}
	return asm.Alloc(dst, src.Kind(asm)).Load(dst, src)
}

func (asm *Asm) Free(z Reg) *Asm {
	use := asm.regs[z]
	if use == nil {
		return asm
	}
	if use.count > 0 {
		use.count--
		if use.count == 0 {
			asm.arch.RegFree(use.reg)
		}
	}
	delete(asm.regs, z)
	return asm
}

// combined Store + Free
func (asm *Asm) StoreFree(dst Var, src Reg) *Asm {
	return asm.Store(dst, src).Free(src)
}
