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
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"fmt"
	"reflect"

	"github.com/cosmos72/gomacro/jit/asm"
	"github.com/cosmos72/gomacro/jit/common"
)

type (
	ArchId      = common.ArchId
	Arch        = common.Arch
	Arg         = common.Arg
	Asm         = common.Asm
	AsmCode     = common.AsmCode
	Const       = common.Const
	Kind        = common.Kind
	MachineCode = common.MachineCode
	Mem         = common.Mem
	Reg         = common.Reg
	RegId       = common.RegId
	RegIdConfig = common.RegIdConfig
	Save        = common.Save
	Size        = common.Size
	SoftRegId   = common.SoftRegId
)

const (
	ASM_SUPPORTED  = asm.ARCH_SUPPORTED
	MMAP_SUPPORTED = asm.MMAP_SUPPORTED
	SUPPORTED      = asm.SUPPORTED
	NAME           = asm.NAME

	// ArchId
	NOARCH = common.NOARCH
	AMD64  = common.AMD64
	ARM64  = common.ARM64

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
)

// map[ArchId]Arch is a handle, changes effect common.Archs
var Archs = common.Archs

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

func SizeOf(a Arg) Size {
	return common.SizeOf(a)
}

var errMakeVarUpn = fmt.Errorf("unimplemented: jit.MakeVar with upn != 0")
var errMakeVarIdx = fmt.Errorf("jit.MakeVar: index too large, the byte offset overflows int32")
var errMakeVarKind = fmt.Errorf("jit.MakeVar: invalid kind")

// local variable
func MakeVar(idx int, upn int, kind Kind, config RegIdConfig) (Mem, error) {
	// TODO support fast.Env local variables with upn > 0
	var mem Mem
	if upn != 0 {
		return mem, errMakeVarUpn
	}
	if kind.Size() == 0 {
		return mem, errMakeVarKind
	}
	off := int32(idx) * 8
	if int(off/8) != idx {
		return mem, errMakeVarIdx
	}
	return common.MakeMem(int32(idx)*8, config.RVAR, kind), nil
}

// function parameter or return value
func MakeParam(off int32, kind Kind, config RegIdConfig) Mem {
	return common.MakeMem(off, config.RSP, kind)
}
