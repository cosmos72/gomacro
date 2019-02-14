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
 * alias.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"github.com/cosmos72/gomacro/jit/asm"
)

type (
	ArchId    = asm.ArchId
	Arch      = asm.Arch
	Arg       = asm.Arg
	Asm       = asm.Asm
	Const     = asm.Const
	Kind      = asm.Kind
	Mem       = asm.Mem
	RegId     = asm.RegId
	Reg       = asm.Reg
	Save      = asm.Save
	Size      = asm.Size
	SoftRegId = asm.SoftRegId
)

const (
	ASM_SUPPORTED  = asm.ARCH_SUPPORTED
	MMAP_SUPPORTED = asm.MMAP_SUPPORTED
	SUPPORTED      = ASM_SUPPORTED && MMAP_SUPPORTED
	NAME           = asm.NAME

	// ArchId
	NOARCH = asm.NOARCH
	AMD64  = asm.AMD64
	ARM64  = asm.ARM64

	// Kind
	Invalid = asm.Invalid
	Bool    = asm.Bool
	Int     = asm.Int
	Int8    = asm.Int8
	Int16   = asm.Int16
	Int32   = asm.Int32
	Int64   = asm.Int64
	Uint    = asm.Uint
	Uint8   = asm.Uint8
	Uint16  = asm.Uint16
	Uint32  = asm.Uint32
	Uint64  = asm.Uint64
	Uintptr = asm.Uintptr
	Float32 = asm.Float32
	Float64 = asm.Float64
	Ptr     = asm.Ptr
	KLo     = asm.KLo
	KHi     = asm.KHi

	// RegId
	NoRegId = asm.NoRegId
)

// map[ArchId]Arch is a handle, changes effect asm.Archs
var Archs = asm.Archs

func MakeConst(val int64, kind Kind) Const {
	return asm.MakeConst(val, kind)
}

func MakeMem(off int32, id RegId, kind Kind) Mem {
	return asm.MakeMem(off, id, kind)
}

func MakeReg(id RegId, kind Kind) Reg {
	return asm.MakeReg(id, kind)
}

func SizeOf(a Arg) Size {
	return asm.SizeOf(a)
}

// local variable
func (c *Comp) MakeVar(idx uint16, kind Kind) Mem {
	// TODO support fast.Env local variables with upn > 0
	return asm.MakeMem(int32(idx)*8, c.RVAR, kind)
}

// function parameter or return value
func (c *Comp) MakeParam(off int32, kind Kind) Mem {
	return asm.MakeMem(off, c.RSP, kind)
}
