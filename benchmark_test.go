/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * expr_test.go
 *
 *  Created on: Mar 06 2017
 *      Author: Massimiliano Ghilardi
 */
package main

import (
	"fmt"
	r "reflect"
	"testing"

	. "github.com/cosmos72/gomacro/base"
	bi "github.com/cosmos72/gomacro/experiments/bytecode_interfaces"
	bv "github.com/cosmos72/gomacro/experiments/bytecode_values"
	ci "github.com/cosmos72/gomacro/experiments/closure_interfaces"
	cm "github.com/cosmos72/gomacro/experiments/closure_maps"
	cv "github.com/cosmos72/gomacro/experiments/closure_values"
	fast "github.com/cosmos72/gomacro/fast_interpreter"
	classic "github.com/cosmos72/gomacro/interpreter"
)

const (
	collatz_n = 837799 // sequence climbs to 1487492288, which also fits 32-bit ints
	sum_n     = 1000
	fib_n     = 12
)

var verbose = false

/*
	BenchmarkCollatzCompiler-2              	 1000000	      1948 ns/op
	BenchmarkCollatzBytecodeInterfaces-2    	   20000	     79932 ns/op
	BenchmarkCollatzClosureValues-2         	   50000	     37442 ns/op
	BenchmarkSumCompiler-2                  	 1000000	      1302 ns/op
	BenchmarkSumBytecodeValues-2            	   10000	    186601 ns/op
	BenchmarkSumBytecodeInterfaces-2        	   10000	    139402 ns/op
	BenchmarkSumClosureValues-2             	   20000	    100682 ns/op
	BenchmarkSumClosureInterfaces-2         	    3000	    404691 ns/op
	BenchmarkSumClosureMaps-2               	    5000	    258179 ns/op
	BenchmarkSumInterpreter-2               	     500	   2865702 ns/op
	BenchmarkFibonacciCompiler-2            	 1000000	      2359 ns/op
	BenchmarkFibonacciClosureValues-2       	    2000	    725237 ns/op
	BenchmarkFibonacciClosureInterfaces-2   	    3000	    558534 ns/op
	BenchmarkFibonacciClosureMaps-2         	    2000	   1093970 ns/op
	BenchmarkFibonacciInterpreter-2         	     500	   2519917 ns/op
*/

func arith(n int) int {
	return ((n*2+3)&4 | 5 ^ 6) / (n | 1)
}

func BenchmarkArithCompiler1(b *testing.B) {
	total := 0
	for i := 0; i < b.N; i++ {
		n := b.N
		total += ((n*2+3)&4 | 5 ^ 6) / (n | 1)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithCompiler2(b *testing.B) {
	total := 0
	for i := 0; i < b.N; i++ {
		total += arith(b.N)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithFastInterpreter(b *testing.B) {
	c := fast.New()
	c.DefVar("n", TypeOfInt, 0)

	fun := c.CompileAst(c.ParseAst("((n*2+3)&4 | 5 ^ 6) / (n|1)"))

	setvar := c.SetVarValue0("n")
	value := r.New(TypeOfInt).Elem()
	var ret r.Value
	c.Run(fun)

	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		value.SetInt(int64(b.N))
		setvar(c.Env, value)
		ret, _ = c.Run(fun)
		total += int(ret.Int())
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithClassicInterpreter(b *testing.B) {
	ir := classic.New()
	ir.EvalAst(ir.ParseAst("n:=0"))

	form := ir.ParseAst("((n*2+3)&4 | 5 ^ 6) / (n|1)")

	value := ir.Binds["n"]
	var ret r.Value
	ir.EvalAst(form)

	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		value.SetInt(int64(b.N))
		ret, _ = ir.EvalAst(form)
		total += int(ret.Int())
	}
	if verbose {
		println(total)
	}
}

func BenchmarkForFastInterpreter(b *testing.B) {
	c := fast.New()
	c.Run(c.CompileAst(c.ParseAst("var i, n, total int")))

	// interpreted code performs iteration and arithmetic

	fun := c.CompileAst(c.ParseAst("total = 0; for i = 0; i < n; i=i+1 { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }; total"))
	// fun := c.CompileAst(c.ParseAst("total = 0; for i = 0; i < n; i=i+1 { }; total"))

	setvar := c.SetVarValue0("n")
	value := r.New(TypeOfInt).Elem()
	var ret r.Value
	c.Run(fun)

	b.ResetTimer()
	total := 0

	value.SetInt(int64(b.N))
	setvar(c.Env, value)
	ret, _ = c.Run(fun)
	total += int(ret.Int())

	if verbose {
		println(total)
	}
}

func BenchmarkForClassicInterpreter(b *testing.B) {
	ir := classic.New()
	ir.EvalAst(ir.ParseAst("var n, total int"))

	// interpreted code performs iteration and arithmetic
	form := ir.ParseAst("total = 0; for i:= 0; i < n; i++ { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }; total")

	value := ir.Binds["n"]
	ir.EvalAst(form)

	b.ResetTimer()
	value.SetInt(int64(b.N))
	ret, _ := ir.EvalAst(form)

	if verbose {
		println(ret.Int())
	}
}

// collatz conjecture

func collatz(n int) {
	for n > 1 {
		if n&1 != 0 {
			n = ((n * 3) + 1) / 2
		} else {
			n = n / 2
		}
	}
}

func BenchmarkCollatzCompiler(b *testing.B) {
	n := collatz_n
	for i := 0; i < b.N; i++ {
		collatz(n)
	}
}

func BenchmarkCollatzBytecodeInterfaces(b *testing.B) {
	coll := bi.BytecodeCollatz()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Vars[0] = collatz_n
		coll.Exec(0)
	}
}

func BenchmarkCollatzClosureValues(b *testing.B) {
	env := cv.NewEnv(nil)
	coll := cv.DeclCollatz(env, 0)
	n := r.ValueOf(collatz_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += coll(n)
	}
}

func BenchmarkCollatzFastInterpreter(b *testing.B) {
	c := fast.New()
	c.DefVar("n", TypeOfInt, 0)

	addr := c.Exec1(c.AddressOfVar("n")).Interface().(*int)

	fun := c.CompileAst(c.ParseAst("for n > 1 { if n&1 != 0 { n = ((n * 3) + 1) / 2 } else { n = n / 2 } }"))
	c.Run(fun)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		*addr = collatz_n
		fun(c.Env)
	}
}

func BenchmarkCollatzClassicInterpreter(b *testing.B) {
	ir := classic.New()
	ir.EvalAst(ir.ParseAst("func collatz(n int) { for n > 1 { if n&1 != 0 { n = ((n * 3) + 1) / 2 } else { n = n / 2 } } }"))

	form := ir.ParseAst(fmt.Sprintf("collatz(%d)", collatz_n))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ir.EvalAst(form)
	}
}

// looping: sum the integers from 1 to N

func sum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

func BenchmarkSumCompiler(b *testing.B) {
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(sum_n)
	}
}

func BenchmarkSumBytecodeValues(b *testing.B) {
	sum := bv.BytecodeSum(sum_n)
	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(sum.Exec(0)[0].Int())
	}
}

func BenchmarkSumBytecodeInterfaces(b *testing.B) {
	p := bi.BytecodeSum(sum_n)
	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += p.Exec(0)[0].(int)
	}
}

func BenchmarkSumClosureValues(b *testing.B) {
	env := cv.NewEnv(nil)
	sum := cv.DeclSum(env, 0)
	n := r.ValueOf(sum_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}

func BenchmarkSumClosureInterfaces(b *testing.B) {
	env := ci.NewEnv(nil)
	sum := ci.DeclSum(env, 0)
	var n interface{} = sum_n

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}

func BenchmarkSumClosureMaps(b *testing.B) {
	env := cm.NewEnv(nil)
	sum := cm.DeclSum(env, "sum")
	n := r.ValueOf(sum_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}

func BenchmarkSumFastInterpreter(b *testing.B) {
	c := fast.New()
	c.Run(c.CompileAst(c.ParseAst("var i, n, total int")))

	addrexpr := c.LookupVar("n").Address()
	addr := c.Exec1(addrexpr).Interface().(*int)
	fun := c.CompileAst(c.ParseAst("total = 0; for i = 1; i <= n; i+=1 { total += i }; total"))
	c.Run(fun)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		*addr = sum_n
		fun(c.Env)
	}
}

func BenchmarkSumClassicInterpreter(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst(sum_s))
	form := env.ParseAst(fmt.Sprintf("sum(%d)", sum_n))

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(env.EvalAst1(form).Int())
	}
}

// recursion: fibonacci. fib(n) => if (n <= 2) { return 1 }; return fib(n-1) + fib(n-2)

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func BenchmarkFibonacciCompiler(b *testing.B) {
	var total int
	n := fib_n
	for i := 0; i < b.N; i++ {
		total += fibonacci(n)
	}
}

func BenchmarkFibonacciClosureValues(b *testing.B) {
	env := cv.NewEnv(nil)
	fib := cv.DeclFibonacci(env, 0)
	n := r.ValueOf(fib_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func BenchmarkFibonacciClosureInterfaces(b *testing.B) {
	env := ci.NewEnv(nil)
	fib := ci.DeclFibonacci(env, 0)
	var n interface{} = fib_n

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func BenchmarkFibonacciClosureMaps(b *testing.B) {
	env := cm.NewEnv(nil)
	fib := cm.DeclFibonacci(env, "fib")
	n := r.ValueOf(fib_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func BenchmarkFibonacciClassicInterpreter(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst(fib_s))
	form := env.ParseAst(fmt.Sprintf("fibonacci(%d)", fib_n))

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(env.EvalAst1(form).Uint())
	}
}
