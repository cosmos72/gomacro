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

package asm

import (
	"github.com/cosmos72/gomacro/jit/common"
	// ensure all known Arch implementations are loaded and registered
	_ "github.com/cosmos72/gomacro/jit/amd64"
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

	// ArchId
	NOARCH = common.NOARCH
	AMD64  = common.AMD64
	ARM64  = common.ARM64

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
	NOT1 = common.NOT1
	NEG1 = common.NEG1

	// Op2
	ADD2     = common.ADD2
	SUB2     = common.SUB2
	ADC2     = common.ADC2
	SBB2     = common.SBB2
	MUL2     = common.MUL2
	DIV2     = common.DIV2
	REM2     = common.REM2
	AND2     = common.AND2
	OR2      = common.OR2
	XOR2     = common.XOR2
	SHL2     = common.SHL2
	SHR2     = common.SHR2
	AND_NOT2 = common.AND_NOT2
	LAND2    = common.LAND2
	LOR2     = common.LOR2
	MOV      = common.MOV
	CAST     = common.CAST
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

// map[ArchId]Arch is a handle, changes effect common.Archs
var Archs = common.Archs

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

func New() *Asm {
	return common.New(ARCH_ID)
}

func NewArchId(archId ArchId) *Asm {
	return common.New(archId)
}

func NewArch(arch Arch) *Asm {
	return common.NewArch(arch)
}
