// +build arm64

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
 * arm64.go
 *
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	arch "github.com/cosmos72/gomacro/jit/arm64"
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
	NoRegId  = arch.NoRegId
	X0       = arch.X0
	X1       = arch.X1
	X2       = arch.X2
	X3       = arch.X3
	X4       = arch.X4
	X5       = arch.X5
	X6       = arch.X6
	X7       = arch.X7
	X8       = arch.X8
	X9       = arch.X9
	X10      = arch.X10
	X11      = arch.X11
	X12      = arch.X12
	X13      = arch.X13
	X14      = arch.X14
	X15      = arch.X15
	X16      = arch.X16
	X17      = arch.X17
	X18      = arch.X18
	X19      = arch.X19
	X20      = arch.X20
	X21      = arch.X21
	X22      = arch.X22
	X23      = arch.X23
	X24      = arch.X24
	X25      = arch.X25
	X26      = arch.X26
	X27      = arch.X27
	X28      = arch.X28
	X29      = arch.X29
	X30      = arch.X30
	XZR, XSP = arch.XZR, arch.XSP
	RLo      = arch.RLo
	RHi      = arch.RHi
	RSP      = arch.RSP
	RVAR     = arch.RVAR

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
	ADD  = arch.ADD
	AND  = arch.AND
	ADC  = arch.ADC
	OR   = arch.OR
	XOR  = arch.XOR
	SUB  = arch.SUB
	SBB  = arch.SBB
	SHL  = arch.SHL
	SHR  = arch.SHR
	MUL  = arch.MUL
	DIV  = arch.DIV
	REM  = arch.REM
	MOV  = arch.MOV
	CAST = arch.CAST
	NEG2 = arch.NEG2
	NOT2 = arch.NOT2
	/*
		CMP  = arch.CMP
		XCHG = arch.XCHG
	*/

	// Op2Misc
	ALLOC = arch.ALLOC
	FREE  = arch.FREE
	PUSH  = arch.PUSH
	POP   = arch.POP

	// Op3
	ADD3 = arch.ADD3
	OR3  = arch.OR3
	ADC3 = arch.ADC3
	SBB3 = arch.SBB3
	AND3 = arch.AND3
	SUB3 = arch.SUB3
	XOR3 = arch.XOR3
	SHL3 = arch.SHL3
	SHR3 = arch.SHR3
	MUL3 = arch.MUL3
	DIV3 = arch.DIV3
	REM3 = arch.REM3

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
