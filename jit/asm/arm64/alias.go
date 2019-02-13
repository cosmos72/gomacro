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
 *  Created on Feb 13, 2019
 *      Author Massimiliano Ghilardi
 */

package arm64

import (
	"github.com/cosmos72/gomacro/jit/asm"
)

type (
	Arch       = asm.Arch
	ArchId     = asm.ArchId
	Arg        = asm.Arg
	Asm        = asm.Asm
	Code       = asm.Code
	Const      = asm.Const
	Kind       = asm.Kind
	Mem        = asm.Mem
	Reg        = asm.Reg
	RegId      = asm.RegId
	RegIdCfg   = asm.RegIdCfg
	RegIds     = asm.RegIds
	Save       = asm.Save
	SaveSlot   = asm.SaveSlot
	Size       = asm.Size
	SoftRegId  = asm.SoftRegId
	SoftRegIds = asm.SoftRegIds

	Op0     = asm.Op0
	Op1     = asm.Op1
	Op2     = asm.Op2
	Op2Misc = asm.Op2Misc
	Op3     = asm.Op3
	Op4     = asm.Op4
)

const (
	MMAP_SUPPORTED = asm.MMAP_SUPPORTED
	SUPPORTED      = MMAP_SUPPORTED
	NAME           = "arm64"

	// SaveSlot
	InvalidSlot = asm.InvalidSlot

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

	// Op0
	NOP = asm.NOP
	RET = asm.RET

	// Op1
	ZERO = asm.ZERO
	INC  = asm.INC
	DEC  = asm.DEC
	NOT  = asm.NOT
	NEG  = asm.NEG

	// Op2
	ADD     = asm.ADD
	SUB     = asm.SUB
	ADC     = asm.ADC
	SBB     = asm.SBB
	MUL     = asm.MUL
	DIV     = asm.DIV
	REM     = asm.REM
	AND     = asm.AND
	OR      = asm.OR
	XOR     = asm.XOR
	SHL     = asm.SHL
	SHR     = asm.SHR
	AND_NOT = asm.AND_NOT
	LAND    = asm.LAND
	LOR     = asm.LOR
	MOV     = asm.MOV
	CAST    = asm.CAST
	// CMP  = asm.CMP
	// XCHG = asm.XCHG
	NEG2 = asm.NEG2
	NOT2 = asm.NOT2

	// Op2Misc
	ALLOC = asm.ALLOC
	FREE  = asm.FREE
	PUSH  = asm.PUSH
	POP   = asm.POP

	// Op3
	ADD3     = asm.ADD3
	SUB3     = asm.SUB3
	ADC3     = asm.ADC3
	SBB3     = asm.SBB3
	MUL3     = asm.MUL3
	DIV3     = asm.DIV3
	REM3     = asm.REM3
	AND3     = asm.AND3
	OR3      = asm.OR3
	XOR3     = asm.XOR3
	SHL3     = asm.SHL3
	SHR3     = asm.SHR3
	AND_NOT3 = asm.AND_NOT3
	LAND3    = asm.LAND3
	LOR3     = asm.LOR3

	// Op4
)

func ConstInt64(val int64) Const {
	return asm.ConstInt64(val)
}

func ConstUint64(val uint64) Const {
	return asm.ConstUint64(val)
}

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

func New() *Asm {
	return asm.NewArch(Arm64{})
}
