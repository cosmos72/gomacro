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
 * example.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

/*
  jit-compiled version of:

	func sum(n int) int {
		total := 0
		for i := 1; i <= n; i++ {
			total += i
		}
		return total
	}
*/
func DeclSum() func(arg int) int {
	const n, total, i = 0, 1, 2
	_, Total, I := NewVar(n), NewVar(total), NewVar(i)

	var asm Asm
	init := asm.Init().Set(I, Int64(1)).Func()
	pred := func(env *[3]uint64) bool {
		return int(env[i]) <= int(env[n])
	}
	var r Reg
	next := asm.Init().ToReg2(I, &r).Add(r, Int64(1)).Store(I, r).Func()
	loop := asm.Init().ToReg2(Total, &r).Add(r, I).Store(Total, r).Func()

	return func(arg int) int {
		env := [3]uint64{n: uint64(arg)}

		for init(&env[0]); pred(&env); next(&env[0]) {
			loop(&env[0])
		}
		return int(env[total])
	}
}

/*
  jit-compiled version of:

	func arith(n int) int {
		return ((((n*2+3)|4) &^ 5) ^ 6) / ((n & 2) | 1)
	}
*/
func DeclArith(envlen int) func(env *uint64) {
	const n, a, b = 0, 1, 2
	N, A, B := NewVar(n), NewVar(a), NewVar(b)

	var asm Asm
	var r, s Reg
	asm.Init2(2, uint16(envlen)).ToReg2(N, &r).Mul(r, Int64(2)).Add(r, Int64(3)). /*Or(r, Int64(4)).Andnot(r, Int64(5)).Xor(r, Int64(6)).*/ Store(A, r)
	asm.ToReg2(N, &s). /*And(s, Int64(2)).Or(s, Int64(1)).*/ Store(B, s).FreeReg(s)
	asm.Quo(r, B).FreeReg(r)
	_, _ = A, B
	return asm.Func()
}
