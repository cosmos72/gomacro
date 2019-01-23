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
 * asm_amd64_mov.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

func (asm *Asm) Asm(args ...interface{}) *Asm {
	n := len(args)
	for i := 0; i < n; i++ {
		i += asm.Op(args[i:]...)
	}
	return asm
}

func (asm *Asm) Op(args ...interface{}) int {
	var n int
	switch op := args[0].(type) {
	case TernaryOp:
		asm.Op3(op, args[1].(Arg), args[2].(Arg), args[3].(Arg))
		n = 3
	case Op:
		asm.Op2(op, args[1].(Arg), args[2].(Arg))
		n = 2
	case UnaryOp:
		asm.Op1(op, args[1].(Arg))
		n = 1
	default:
		errorf("syntax error: expecting TernaryOp, Op or UnaryOp [args], found %v", op)
	}
	return n
}
