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

package native

import (
	arch "github.com/cosmos72/gomacro/jit/arm64"
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
	X0      = arch.X0
	X1      = arch.X1
	X2      = arch.X2
	X3      = arch.X3
	X4      = arch.X4
	X5      = arch.X5
	X6      = arch.X6
	X7      = arch.X7
	X8      = arch.X8
	X9      = arch.X9
	X10     = arch.X10
	X11     = arch.X11
	X12     = arch.X12
	X13     = arch.X13
	X14     = arch.X14
	X15     = arch.X15
	X16     = arch.X16
	X17     = arch.X17
	X18     = arch.X18
	X19     = arch.X19
	X20     = arch.X20
	X21     = arch.X21
	X22     = arch.X22
	X23     = arch.X23
	X24     = arch.X24
	X25     = arch.X25
	X26     = arch.X26
	X27     = arch.X27
	X28     = arch.X28
	X29     = arch.X29
	X30     = arch.X30
	X31     = arch.X31
	RLo     = arch.RLo
	RHi     = arch.RHi

	// Op0

	// Op1

	// Op2

	// Op3

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
