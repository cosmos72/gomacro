// +build !amd64,!arm64

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
 * generic.go
 *
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	arch "github.com/cosmos72/gomacro/jit/generic"
)

type (
	Asm        = arch.Asm
	Arg        = arch.Arg
	Code       = arch.Code
	Const      = arch.Const
	Kind       = arch.Kind
	Mem        = arch.Mem
	RegId      = arch.RegId
	RegIds     = arch.RegIds
	Reg        = arch.Reg
	Save       = arch.Save
	Size       = arch.Size
	SoftRegId  = arch.SoftRegId
	SoftRegIds = arch.SoftRegIds

	Op0     = arch.Op0
	Op1     = arch.Op1
	Op2     = arch.Op2
	Op2Misc = arch.Op2Misc
	Op3     = arch.Op3
	Op4     = arch.Op4
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
	RSP     = arch.RSP
	RVAR    = arch.RVAR

	// Op0
	NOP = arch.NOP
	RET = arch.RET

	// Op1
	ZERO = arch.ZERO
	INC  = arch.INC
	DEC  = arch.DEC
	NOT  = arch.NOT
	NEG  = arch.NEG

	// Op2
	ADD     = arch.ADD
	SUB     = arch.SUB
	ADC     = arch.ADC
	SBB     = arch.SBB
	MUL     = arch.MUL
	DIV     = arch.DIV
	REM     = arch.REM
	AND     = arch.AND
	OR      = arch.OR
	XOR     = arch.XOR
	SHL     = arch.SHL
	SHR     = arch.SHR
	AND_NOT = arch.AND_NOT
	LAND    = arch.LAND
	LOR     = arch.LOR
	MOV     = arch.MOV
	CAST    = arch.CAST
	// CMP  = arch.CMP
	// XCHG = arch.XCHG
	NEG2 = arch.NEG2
	NOT2 = arch.NOT2

	// Op2Misc
	ALLOC = arch.ALLOC
	FREE  = arch.FREE
	PUSH  = arch.PUSH
	POP   = arch.POP

	// Op3
	ADD3     = arch.ADD3
	SUB3     = arch.SUB3
	ADC3     = arch.ADC3
	SBB3     = arch.SBB3
	MUL3     = arch.MUL3
	DIV3     = arch.DIV3
	REM3     = arch.REM3
	AND3     = arch.AND3
	OR3      = arch.OR3
	XOR3     = arch.XOR3
	SHL3     = arch.SHL3
	SHR3     = arch.SHR3
	AND_NOT3 = arch.AND_NOT3
	LAND3    = arch.LAND3
	LOR3     = arch.LOR3

	// Op4
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

func NewAsm() *Asm {
	return arch.NewAsm()
}
