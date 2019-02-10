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
 * op3.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// ============================================================================
// ternary operation
type Op3 uint8

const (
	ADD3 Op3 = Op3(ADD)
	OR3  Op3 = Op3(OR)
	ADC3 Op3 = Op3(ADC) // add with carry
	SBB3 Op3 = Op3(SBB) // subtract with borrow
	AND3 Op3 = Op3(AND)
	SUB3 Op3 = Op3(SUB)
	XOR3 Op3 = Op3(XOR)
	SHL3 Op3 = Op3(SHL) // shift left
	SHR3 Op3 = Op3(SHR) // shift right
	MUL3 Op3 = Op3(MUL)
	DIV3 Op3 = Op3(DIV)
	REM3 Op3 = Op3(REM)
)

var op3Name = map[Op3]string{
	ADD3: "ADD3",
	OR3:  "OR3",
	ADC3: "ADC3",
	SBB3: "SBB3",
	AND3: "AND3",
	SUB3: "SUB3",
	XOR3: "XOR3",
	SHL3: "SHL3",
	SHR3: "SHR3",
	MUL3: "MUL3",
	DIV3: "DIV3",
	REM3: "REM3",
}

func (op Op3) String() string {
	s := op3Name[op]
	if len(s) == 0 {
		s = fmt.Sprintf("Op3(%d)", int(op))
	}
	return s
}

// ============================================================================
func (asm *Asm) Op3(op Op3, a Arg, b Arg, dst Arg) *Asm {
	// validate kinds
	assert(a.Kind() == dst.Kind())
	switch op {
	case SHL3, SHR3:
		assert(!b.Kind().Signed())
	default:
		assert(b.Kind() == dst.Kind())
	}
	// validate dst
	switch dst.(type) {
	case Reg, Mem:
		break
	case Const:
		errorf("destination cannot be a constant: %v %v, %v, %v", op, a, b, dst)
	default:
		errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v, %v", dst, op, a, b, dst)
	}
	if asm.optimize3(op, a, b, dst) {
		return asm
	}
	if op == MUL3 {
		return asm.mul3(a, b, dst)
	}
	op2 := Op2(op)
	if a == dst {
		asm.Op2(op2, b, dst)
	} else if op.IsCommutative() && b == dst {
		asm.Op2(op2, a, dst)
	} else if r, ok := dst.(Reg); ok && r.id != b.RegId() {
		asm.Mov(a, dst).Op2(op2, b, dst)
	} else {
		r := asm.RegAlloc(dst.Kind())
		asm.Mov(a, r).Op2(op2, b, r).Mov(r, dst)
		asm.RegFree(r)
	}
	return asm
}

func (asm *Asm) mul3(a Arg, b Arg, dst Arg) *Asm {
	if a.Const() && !b.Const() {
		a, b = b, a
	}
	if a == dst {
		return asm.Op2(MUL, b, dst)
	} else if b == dst {
		return asm.Op2(MUL, a, dst)
	}
	rdst, rokdst := dst.(Reg)
	if !a.Const() && b.Const() {
		bval := b.(Const).val
		if bval == int64(int8(bval)) {
			// use amd64 3-argument multiplication
			if !rokdst {
				rdst = asm.RegAlloc(dst.Kind())
			}
			switch a := a.(type) {
			case Reg:
				asm.mul3RegConst8Reg(a, int8(bval), rdst)
			case Mem:
				asm.mul3MemConst8Reg(a, int8(bval), rdst)
			default:
				errorf("unknown argument type %T, expecting Const, Reg or Mem: %v %v, %v, %v", a, MUL3, a, b, dst)
			}
			if !rokdst {
				asm.Store(rdst, dst.(Mem)).RegFree(rdst)
			}
			return asm
		}
	}
	if rokdst && rdst.id != b.RegId() {
		return asm.Mov(a, dst).Op2(MUL, b, dst)
	}
	r := asm.RegAlloc(dst.Kind())
	asm.Mov(a, r).Op2(MUL, b, r).Mov(r, dst)
	return asm.RegFree(r)
}
