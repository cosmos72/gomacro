// +build amd64

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
 * amd64.go
 *
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

package native

import (
	arch "github.com/cosmos72/gomacro/jit/amd64"
)

type (
	Asm   = arch.Asm
	Arg   = arch.Arg
	Code  = arch.Code
	Const = arch.Const
	Kind  = arch.Kind
	Mem   = arch.Mem
	RegId = arch.RegId
	Reg   = arch.Reg
	Save  = arch.Save
	Size  = arch.Size

	Op0     = arch.Op0
	Op1     = arch.Op1
	Op2     = arch.Op2
	Op2Misc = arch.Op2Misc
	Op3     = arch.Op3
	Op4     = arch.Op4
)

const (
	SUPPORTED = arch.SUPPORTED
	Name      = arch.Name

	// Kind
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
	RAX     = arch.RAX
	RCX     = arch.RCX
	RDX     = arch.RDX
	RBX     = arch.RBX
	RSP     = arch.RSP
	RBP     = arch.RBP
	RSI     = arch.RSI
	RDI     = arch.RDI
	R8      = arch.R8
	R9      = arch.R9
	R10     = arch.R10
	R11     = arch.R11
	R12     = arch.R12
	R13     = arch.R13
	R14     = arch.R14
	R15     = arch.R15
	RLo     = arch.RLo
	RHi     = arch.RHi

	// Op0
	RET = arch.RET
	NOP = arch.NOP

	// Op1
	NOT = arch.NOT
	NEG = arch.NEG
	INC = arch.INC
	DEC = arch.DEC

	// Op2
	ADD = arch.ADD
	OR  = arch.OR
	ADC = arch.ADC
	SBB = arch.SBB
	AND = arch.AND
	SUB = arch.SUB
	XOR = arch.XOR
	// CMP  = arch.CMP
	// XCHG  = arch.XCHG
	MOV  = arch.MOV
	LEA  = arch.LEA
	CAST = arch.CAST
	SHL  = arch.SHL
	SHR  = arch.SHR
	MUL  = arch.MUL
	DIV  = arch.DIV
	REM  = arch.REM

	// Op3
	ADD3  = arch.ADD3
	OR3   = arch.OR3
	ADC3  = arch.ADC3
	SBB3  = arch.SBB3
	AND3  = arch.AND3
	SUB3  = arch.SUB3
	XOR3  = arch.XOR3
	CAST3 = arch.CAST3
	SHL3  = arch.SHL3
	SHR3  = arch.SHR3
	MUL3  = arch.MUL3
	DIV3  = arch.DIV3
	REM3  = arch.REM3

	// Op4
	LEA4 = arch.LEA4
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

func Sizeof(a Arg) Size {
	return arch.SizeOf(a)
}
