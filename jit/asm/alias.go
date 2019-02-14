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
	"github.com/cosmos72/gomacro/jit/common"
	// ensure Arch implementations are loaded and registered
	// _ "github.com/cosmos72/gomacro/jit/amd64"
	_ "github.com/cosmos72/gomacro/jit/arm64"
)

type (
	Arch       = common.Arch
	ArchId     = common.ArchId
	Arg        = common.Arg
	Asm        = common.Asm
	Code       = common.Code
	Const      = common.Const
	Kind       = common.Kind
	Mem        = common.Mem
	Reg        = common.Reg
	RegId      = common.RegId
	RegIdCfg   = common.RegIdCfg
	RegIds     = common.RegIds
	Save       = common.Save
	SaveSlot   = common.SaveSlot
	Size       = common.Size
	SoftRegId  = common.SoftRegId
	SoftRegIds = common.SoftRegIds

	Op0     = common.Op0
	Op1     = common.Op1
	Op2     = common.Op2
	Op2Misc = common.Op2Misc
	Op3     = common.Op3
	Op4     = common.Op4
)

const (
	MMAP_SUPPORTED = common.MMAP_SUPPORTED
	SUPPORTED      = MMAP_SUPPORTED
	NAME           = "arm64"

	// SaveSlot
	InvalidSlot = common.InvalidSlot

	// Kind
	Invalid = common.Invalid
	Bool    = common.Bool
	Int     = common.Int
	Int8    = common.Int8
	Int16   = common.Int16
	Int32   = common.Int32
	Int64   = common.Int64
	Uint    = common.Uint
	Uint8   = common.Uint8
	Uint16  = common.Uint16
	Uint32  = common.Uint32
	Uint64  = common.Uint64
	Uintptr = common.Uintptr
	Float32 = common.Float32
	Float64 = common.Float64
	Ptr     = common.Ptr
	KLo     = common.KLo
	KHi     = common.KHi

	// RegId
	NoRegId = common.NoRegId

	// Op0
	NOP = common.NOP
	RET = common.RET

	// Op1
	ZERO = common.ZERO
	INC  = common.INC
	DEC  = common.DEC
	NOT  = common.NOT
	NEG  = common.NEG

	// Op2
	ADD     = common.ADD
	SUB     = common.SUB
	ADC     = common.ADC
	SBB     = common.SBB
	MUL     = common.MUL
	DIV     = common.DIV
	REM     = common.REM
	AND     = common.AND
	OR      = common.OR
	XOR     = common.XOR
	SHL     = common.SHL
	SHR     = common.SHR
	AND_NOT = common.AND_NOT
	LAND    = common.LAND
	LOR     = common.LOR
	MOV     = common.MOV
	CAST    = common.CAST
	// CMP  = common.CMP
	// XCHG = common.XCHG
	NEG2 = common.NEG2
	NOT2 = common.NOT2

	// Op2Misc
	ALLOC = common.ALLOC
	FREE  = common.FREE
	PUSH  = common.PUSH
	POP   = common.POP

	// Op3
	ADD3     = common.ADD3
	SUB3     = common.SUB3
	ADC3     = common.ADC3
	SBB3     = common.SBB3
	MUL3     = common.MUL3
	DIV3     = common.DIV3
	REM3     = common.REM3
	AND3     = common.AND3
	OR3      = common.OR3
	XOR3     = common.XOR3
	SHL3     = common.SHL3
	SHR3     = common.SHR3
	AND_NOT3 = common.AND_NOT3
	LAND3    = common.LAND3
	LOR3     = common.LOR3

	// Op4
)

func ConstInt64(val int64) Const {
	return common.ConstInt64(val)
}

func ConstUint64(val uint64) Const {
	return common.ConstUint64(val)
}

func MakeConst(val int64, kind Kind) Const {
	return common.MakeConst(val, kind)
}

func MakeMem(off int32, id RegId, kind Kind) Mem {
	return common.MakeMem(off, id, kind)
}

func MakeReg(id RegId, kind Kind) Reg {
	return common.MakeReg(id, kind)
}

func SizeOf(a Arg) Size {
	return common.SizeOf(a)
}

func New(archId ArchId) *Asm {
	return common.New(archId)
}

func NewArch(arch Arch) *Asm {
	return common.NewArch(arch)
}
