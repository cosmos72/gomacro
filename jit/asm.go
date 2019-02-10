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
 * arch.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	arch "github.com/cosmos72/gomacro/jit/native"
)

type (
	Asm     = arch.Asm
	Arg     = arch.Arg
	AsmCode = arch.Code
	Const   = arch.Const
	Kind    = arch.Kind
	Mem     = arch.Mem
	RegId   = arch.RegId
	RegIds  = arch.RegIds
	Reg     = arch.Reg
	Save    = arch.Save
	Size    = arch.Size
)

const (
	ASM_SUPPORTED  = arch.ASM_SUPPORTED
	MMAP_SUPPORTED = arch.MMAP_SUPPORTED
	SUPPORTED      = ASM_SUPPORTED && MMAP_SUPPORTED
	Name           = arch.Name

	// Kind
	Invalid = arch.Invalid
	Bool    = arch.Bool
	Int     = arch.Int
	Int8    = arch.Int8
	Int16   = arch.Int16
	Int32   = arch.Int32
	Int64   = arch.Int64
	Uint    = arch.Uint
	Uint8   = arch.Uint8
	Uint16  = arch.Uint16
	Uint32  = arch.Uint32
	Uint64  = arch.Uint64
	Uintptr = arch.Uintptr
	Float32 = arch.Float32
	Float64 = arch.Float64
	Ptr     = arch.Ptr
	KLo     = arch.KLo
	KHi     = arch.KHi

	// RegId
	NoRegId = arch.NoRegId
	RLo     = arch.RLo
	RHi     = arch.RHi
)

func MakeConst(val int64, kind Kind) Const {
	return arch.MakeConst(val, kind)
}

func MakeMem(off int32, id RegId, kind Kind) Mem {
	return arch.MakeMem(off, id, kind)
}

func MakeReg(id RegId, kind Kind) Reg {
	return arch.MakeReg(id, kind)
}

func SizeOf(a Arg) Size {
	return arch.SizeOf(a)
}
