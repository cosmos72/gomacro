/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * declaration.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

func varInt0(env *Env) int {
	return int(env.Binds[0].Int())
}

/*
  interpreted version of:

	func collatz(n int) {
		for n > 1 {
			if n&1 != 0 {
				n = ((n * 3) + 1) / 2
			} else {
				n = n / 2
			}
		}
	}
*/
/*
func DeclCollatz(env *Env, idx int) FuncInt {
	const (
		n = 0
	)
	return DeclFuncInt(
		idx, []r.Type{typeOfInt},
		Block(
			For(nil, LessInt(Int(1), VarInt(n)), nil,
				If(NeqInt(BitandInt(VarInt(n), Int(1)), Int(0)),
					VarSetInt(n,
						RshiftInt(
							AddInt(
								MulInt(VarInt(n), Int(3)),
								Int(1),
							),
							Int(1),
						),
					),
					VarSetInt(n,
						RshiftInt(
							VarInt(n),
							Int(1),
						),
					),
				),
			),
			ReturnInt(Int(0)),
		),
	)(env)
}
*/

/*
  interpreted version of:

	func sum(n int) int {
		total := 0
		for i := 1; i <= n; i++ {
			total += i
		}
		return total
	}
*/
/*
func DeclSum(env *Env, idx int) FuncInt {
	const (
		n     = 0
		total = 1
		i     = 2
	)
	return DeclFuncInt(
		idx, []r.Type{typeOfInt},
		Block(
			DeclVar(total, Const(0)),
			For(DeclVar(i, Const(1)), LeqInt(VarInt(i), VarInt(n)), VarIncInt(i),
				VarSetInt(total,
					AddInt(
						VarInt(total), VarInt(i),
					),
				),
			),
			ReturnInt(VarInt(total)),
		),
	)(env)
}
*/

/*
  interpreted version of:

	func fibonacci(n int) int {
		if (n <= 2) {
			return 1
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}
*/
/*
func DeclFibonacci(env *Env, idx int) FuncInt {
	const (
		n = 0
	)
	return DeclFuncInt(
		idx, []r.Type{typeOfInt},
		Block(
			If(LessInt(VarInt0, Int(2)),
				ReturnInt(Int(1)),
				ReturnInt(
					AddInt(
						CallInt(Var(1, idx), IntToX(SubInt(VarInt0, Int(1)))),
						CallInt(Var(1, idx), IntToX(SubInt(VarInt0, Int(2)))),
					),
				),
			),
		),
	)(env)
}
*/
