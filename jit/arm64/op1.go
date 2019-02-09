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
 * op1.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// ============================================================================
// unary operation
type Op1 uint8

const (
	ZERO = Op1(0x10)
	INC  = Op1(ADD3)
	DEC  = Op1(SUB3)
	NEG  = Op1(NEG2)
	NOT  = Op1(NOT2)
)

var op1Name = map[Op1]string{
	ZERO: "ZERO",
	INC:  "INC",
	DEC:  "DEC",
	NOT:  "NOT",
	NEG:  "NEG",
}

func (op Op1) String() string {
	s, ok := op1Name[op]
	if !ok {
		s = fmt.Sprintf("Op1(%d)", int(op))
	}
	return s
}

// ============================================================================
func (asm *Asm) Op1(op Op1, a Arg) *Asm {
	switch op {
	case ZERO:
		asm.Zero(a)
	case INC, DEC:
		asm.Op3(Op3(op), a, MakeConst(1, a.Kind()), a)
	case NEG, NOT:
		asm.Op2(Op2(op), a, a)
	default:
		errorf("unknown Op1 instruction: %v %v", op, a)
	}
	return asm
}

// zero a register or memory location
func (asm *Asm) Zero(dst Arg) *Asm {
	switch dst := dst.(type) {
	case Const:
		errorf("cannot zero a constant: %v %v", ZERO, dst)
	case Reg:
		asm.zeroReg(dst)
	case Mem:
		asm.zeroMem(dst)
	default:
		errorf("unknown destination type %T, expecting Reg or Mem: %v, %v", dst, ZERO, dst)
	}
	return asm
}

// zero a register
func (asm *Asm) zeroReg(dst Reg) *Asm {
	// alternative: return asm.movRegReg(MakeReg(XZR, dst.kind), dst)
	return asm.movConstReg(MakeConst(0, dst.kind), dst)
}

// zero a memory location
func (asm *Asm) zeroMem(dst Mem) *Asm {
	return asm.Store(MakeReg(XZR, dst.Kind()), dst)
}
