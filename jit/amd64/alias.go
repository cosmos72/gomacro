/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
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

package amd64

import (
	"reflect"

	"github.com/cosmos72/gomacro/jit/common"
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
	RegIdConfig   = common.RegIdConfig
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
	NAME           = "amd64"

	// ArchId
	NOARCH = common.NOARCH
	AMD64  = common.AMD64

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
	BAD = common.BAD
	NOP = common.NOP
	RET = common.RET

	// Op1
	ZERO = common.ZERO
	INC  = common.INC
	DEC  = common.DEC
	NOT  = common.NOT1
	NEG  = common.NEG1

	// Op2
	ADD     = common.ADD2
	SUB     = common.SUB2
	ADC     = common.ADC2
	SBB     = common.SBB2
	MUL     = common.MUL2
	DIV     = common.DIV2
	REM     = common.REM2
	AND     = common.AND2
	OR      = common.OR2
	XOR     = common.XOR2
	SHL     = common.SHL2
	SHR     = common.SHR2
	AND_NOT = common.AND_NOT2
	LAND    = common.LAND2
	LOR     = common.LOR2
	MOV     = common.MOV
	CAST    = common.CAST
	LEA2    = common.LEA2
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
	LEA4 = common.LEA4
)

func ConstInt8(val int8) Const {
	return common.ConstInt8(val)
}

func ConstInt16(val int16) Const {
	return common.ConstInt16(val)
}

func ConstInt32(val int32) Const {
	return common.ConstInt32(val)
}

func ConstInt64(val int64) Const {
	return common.ConstInt64(val)
}

func ConstUint8(val uint8) Const {
	return common.ConstUint8(val)
}

func ConstUint16(val uint16) Const {
	return common.ConstUint16(val)
}

func ConstUint32(val uint32) Const {
	return common.ConstUint32(val)
}

func ConstUint64(val uint64) Const {
	return common.ConstUint64(val)
}

func ConstInterface(ival interface{}, t reflect.Type) (Const, error) {
	return common.ConstInterface(ival, t)
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

func New() *Asm {
	return common.NewArch(Amd64{})
}

func SizeOf(a Arg) Size {
	return common.SizeOf(a)
}

func log2uint(n uint64) (uint8, bool) {
	return common.Log2Uint(n)
}
