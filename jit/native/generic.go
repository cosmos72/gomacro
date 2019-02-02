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
