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
	RLo     = arch.RLo
	RHi     = arch.RHi

	// Op0
	RET Op0 = 0x00
	NOP Op0 = 0x01

	// Op1
	ZERO Op1 = 0x10
	INC  Op1 = 0x11
	DEC  Op1 = 0x12
	NEG  Op1 = 0x13
	NOT  Op1 = 0x14

	// Op2
	ADD  Op2 = 0x20
	AND  Op2 = 0x21
	ADC  Op2 = 0x22
	OR   Op2 = 0x23
	XOR  Op2 = 0x24
	SUB  Op2 = 0x25
	SBB  Op2 = 0x26
	SHL  Op2 = 0x27
	SHR  Op2 = 0x28
	MUL  Op2 = 0x29
	DIV  Op2 = 0x2A
	REM  Op2 = 0x2B
	MOV  Op2 = 0x2C
	CAST Op2 = 0x2D
	NEG2     = Op2(NEG)
	NOT2     = Op2(NOT)
	/*
		CMP  = arch.CMP
		XCHG = arch.XCHG
	*/

	// Op2Misc
	ALLOC Op2Misc = 0x01
	FREE  Op2Misc = 0x02
	PUSH  Op2Misc = 0x03
	POP   Op2Misc = 0x04

	// Op3
	ADD3 = Op3(ADD)
	OR3  = Op3(OR)
	ADC3 = Op3(ADC)
	SBB3 = Op3(SBB)
	AND3 = Op3(AND)
	SUB3 = Op3(SUB)
	XOR3 = Op3(XOR)
	SHL3 = Op3(SHL)
	SHR3 = Op3(SHR)
	MUL3 = Op3(MUL)
	DIV3 = Op3(DIV)
	REM3 = Op3(REM)

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

func Sizeof(a Arg) Size {
	return arch.SizeOf(a)
}
