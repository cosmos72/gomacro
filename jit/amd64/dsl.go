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

package amd64

type Op uint8

const (
	ADD Op = iota
	SUB
	MUL
	QUO
	REM

	AND
	OR
	XOR
	ANDNOT

	LOAD
	STORE
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
	switch op {
	case ADD, SUB, MUL, QUO, REM, AND, OR, XOR, ANDNOT:
		if len(args) < 2 {
			errorf("syntax error: expecting OP arg1 arg2, found %v", append([]interface{}{op}, args...)...)
		}
		asm.Op2(op, args[0].(hwReg), args[1].(Arg))
		return 2
	case LOAD:
		asm.Load(args[0].(hwReg), args[1].(Arg))
		return 2
	case STORE:
		asm.Store(args[0].(*Var), args[1].(hwReg))
		return 2
	default:
		errorf("unknown operator: %v", op)
	}
	return 0
}

func (asm *Asm) Op2(op Op, z hwReg, a Arg) *Asm {
	switch op {
	case ADD:
		asm.Add(z, a)
	case SUB:
		asm.Sub(z, a)
	case MUL:
		asm.Mul(z, a)
	case QUO:
		asm.Quo(z, a)
	case REM:
		asm.Rem(z, a)
	case AND:
		asm.And(z, a)
	case OR:
		asm.Or(z, a)
	case XOR:
		asm.Xor(z, a)
	case ANDNOT:
		asm.Andnot(z, a)
	default:
		errorf("unknown binary operator: %v", op)
	}
	return asm
}
