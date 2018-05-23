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
	next := asm.Init().Add(I, Int64(1)).Func()
	loop := asm.Init().Add(Total, I).Func()

	return func(arg int) int {
		var env = [3]uint64{n: uint64(arg)}

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
func DeclArithMem() func(env *uint64) {
	const n, a, b = 0, 1, 2
	N, A, B := NewVar(n), NewVar(a), NewVar(b)

	var asm Asm
	asm.Init().Set(A, N).Mul(A, Int64(2)).Add(A, Int64(3)).Or(A, Int64(4)).Andnot(A, Int64(5)).Xor(A, Int64(6))
	asm.Set(B, N).And(B, Int64(2)).Or(B, Int64(1))
	asm.Quo(A, B)
	return asm.Func()
}

func DeclArithReg() func(env *uint64) {
	const n, a, b = 0, 1, 2
	N, A, B := NewVar(n), NewVar(a), NewVar(b)

	var asm Asm
	asm.Init().load_rax(N).Mul_ax(Int64(2)).Add_ax(Int64(3)).Or_ax(Int64(4)).Andnot_ax(Int64(5)).Xor_ax(Int64(6)).store_rax(A)
	asm.load_rax(N).And_ax(Int64(2)).Or_ax(Int64(1)).store_rax(B)
	asm.Quo(A, B)
	return asm.Func()
}
