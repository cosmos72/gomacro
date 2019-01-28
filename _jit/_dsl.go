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
 * dsl.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

type Op uint8

const (
	MOV Op = iota

	ALLOC
	FREE

	ADD
	SUB
	SMUL // signed   multiplication
	UMUL // unsigned multiplication
	SDIV // signed   quotient
	UDIV // unsigned quotient
	SREM // signed   remainder
	UREM // unsigned remainder

	AND
	OR
	XOR
	ANDNOT

	NEG
	NOT
)

type divkind int

const (
	signed, unsigned divkind = 0, 1
	div, rem         divkind = 0, 2
)

func (asm *Asm) Asm(args ...interface{}) *Asm {
	n := len(args)
	for i := 0; i < n; i++ {
		op, ok := args[i].(Op)
		if !ok {
			errorf("syntax error: expecting OP [args], found %v", args[i])
		}
		i += asm.Op(op, args[i+1:]...)
	}
	return asm
}

func (asm *Asm) Op(op Op, args ...interface{}) int {
	var n int
	switch op {
	case ADD, SUB, SMUL, UMUL, SDIV, UDIV, SREM, UREM, AND, OR, XOR, ANDNOT:
		asm.Op2(op, args[0].(Arg), args[1].(Arg), args[2].(Arg))
		n = 2
	case MOV:
		asm.Mov(args[0].(Arg), args[1].(Arg))
		n = 2
	case ALLOC:
		asm.Alloc(args[0].(Reg))
		n = 1
	case FREE:
		asm.Free(args[0].(Reg))
		n = 1
	case NEG, NOT:
		asm.Op1(op, args[0].(Arg), args[1].(Arg))
		n = 1
	default:
		errorf("unknown operator: %v", op)
	}
	return n
}

func (asm *Asm) Op1(op Op, a Arg) *Asm {
	switch op {
	case NEG:
		asm.Neg(a)
	case NOT:
		asm.Not(a)
	default:
		errorf("unknown unary operator: %v", op)
	}
	return asm
}

func (asm *Asm) Op2(op Op, dst Arg, a Arg, b Arg) *Asm {
	switch op {
	case ADD:
		asm.Add(dst, a, b)
	case SUB:
		asm.Sub(dst, a, b)
	case SMUL:
		asm.SMul(dst, a, b)
	case UMUL:
		asm.UMul(dst, a, b)
	case SDIV:
		asm.SDiv(dst, a, b)
	case UDIV:
		asm.UDiv(dst, a, b)
	case SREM:
		asm.SRem(dst, a, b)
	case UREM:
		asm.URem(dst, a, b)
	case AND:
		asm.And(dst, a, b)
	case OR:
		asm.Or(dst, a, b)
	case XOR:
		asm.Xor(dst, a, b)
	case ANDNOT:
		asm.Andnot(dst, a, b)
	default:
		errorf("unknown binary operator: %v", op)
	}
	return asm
}
