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
	case Op0:
		asm.Op0(op)
		n = 0
	case Op1:
		asm.Op1(op, args[1].(Arg))
		n = 1
	case Op2:
		asm.Op2(op, args[1].(Arg), args[2].(Arg))
		n = 2
	case Op3:
		asm.Op3(op, args[1].(Arg), args[2].(Arg), args[3].(Arg))
		n = 3
	case Op4:
		asm.Op4(op, args[1].(Arg), args[2].(Arg), args[3].(Arg), args[4].(Arg))
		n = 4
	default:
		errorf("syntax error: expecting Op0,Op1,Op2,Op3 or Op4 [args], found %v", op)
	}
	return n
}
