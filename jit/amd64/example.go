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
	_, Total, I := Bind{Idx: n}, Bind{Idx: total}, Bind{Idx: i}

	var asm Asm
	init := asm.Init().SetInt64(I, 1).Func()
	pred := func(env *[3]uint64) bool {
		return int(env[i]) <= int(env[n])
	}
	next := asm.Init().IncInt64c(I, 1).Func()
	loop := asm.Init().IncInt64(Total, I).Func()

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
		return ((n*2+3)&4 | 5 ^ 6) / (n | 1)
	}
*/
func DeclArith() func(arg int) int {
	// TODO
	return nil
}
