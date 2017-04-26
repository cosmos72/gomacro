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
	"github.com/cosmos72/gomacro/classic"
	bi "github.com/cosmos72/gomacro/experiments/bytecode_interfaces"
	bv "github.com/cosmos72/gomacro/experiments/bytecode_values"
	ci "github.com/cosmos72/gomacro/experiments/closure_interfaces"
	cm "github.com/cosmos72/gomacro/experiments/closure_maps"
	cv "github.com/cosmos72/gomacro/experiments/closure_values"
	"github.com/cosmos72/gomacro/fast"
)

const (
	collatz_n = 837799 // sequence climbs to 1487492288, which also fits 32-bit ints
	sum_n     = 1000
	fib_n     = 12
)

var verbose = false

/*
	--------- results on Intel Core 2 Duo P8400 ---------------

	BenchmarkFibonacciCompiler-2             	  500000	      2447 ns/op
	BenchmarkFibonacciFastInterpreter-2      	   30000	     41455 ns/op
	BenchmarkFibonacciClassicInterpreter-2   	    1000	   2192651 ns/op
	BenchmarkFibonacciClosureValues-2        	    2000	    719944 ns/op
	BenchmarkFibonacciClosureInterfaces-2    	    3000	    580711 ns/op
	BenchmarkFibonacciClosureMaps-2          	    2000	   1121442 ns/op
	BenchmarkArithCompiler1-2                	100000000	        17.7 ns/op
	BenchmarkArithCompiler2-2                	100000000	        17.7 ns/op
	BenchmarkArithFastInterpreter-2          	10000000	       170 ns/op
	BenchmarkArithFastInterpreterBis-2       	10000000	       127 ns/op
	BenchmarkArithClassicInterpreter-2       	  500000	      3086 ns/op
	BenchmarkArithClassicInterpreterBis-2    	  300000	      5632 ns/op
	BenchmarkCollatzCompiler-2               	 1000000	      1915 ns/op
	BenchmarkCollatzFastInterpreter-2        	   50000	     35823 ns/op
	BenchmarkCollatzClassicInterpreter-2     	    1000	   1277431 ns/op
	BenchmarkCollatzBytecodeInterfaces-2     	   20000	     76116 ns/op
	BenchmarkCollatzClosureValues-2          	   50000	     37573 ns/op
	BenchmarkSumCompiler-2                   	 1000000	      1304 ns/op
	BenchmarkSumFastInterpreter-2            	   30000	     55481 ns/op
	BenchmarkSumClassicInterpreter-2         	    1000	   2338576 ns/op
	BenchmarkSumBytecodeValues-2             	   10000	    188303 ns/op
	BenchmarkSumBytecodeInterfaces-2         	   10000	    136112 ns/op
	BenchmarkSumClosureValues-2              	   20000	     98944 ns/op
	BenchmarkSumClosureInterfaces-2          	    3000	    407104 ns/op
	BenchmarkSumClosureMaps-2                	    5000	    256734 ns/op
*/

// recursion: fibonacci. fib(n) => if (n <= 2) { return 1 } else { return fib(n-1) + fib(n-2) }

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

func BenchmarkFibonacciFastInterpreter(b *testing.B) {
	ce := fast.New()
	c := ce.Comp
	ce.Eval(fib_s)

	// compile the call to fibonacci(fib_n)
	fun := c.CompileAst(c.ParseAst(fmt.Sprintf("fibonacci(%d)", fib_n)))
	env := ce.PrepareEnv()
	fun(env)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		retv, _ := fun(env)
		total += int(retv.Int())
	}
}

func BenchmarkFibonacciFastInterpreterBis(b *testing.B) {
	ce := fast.New()
	ce.Eval(fib_s)

	// alternative: extract the function fibonacci, and call it ourselves
	//
	// ValueOf is the method to retrieve constants, functions and variables from the classic and fast interpreters
	// (if you set the same variable repeatedly, use the address returned by AddressOfVar)
	fun := ce.ValueOf("fibonacci").Interface().(func(int) int)
	fun(fib_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(fib_n)
	}
}

func BenchmarkFibonacciClassicInterpreter(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst(fib_s))

	// compile the call to fibonacci(fib_n)
	form := env.ParseAst(fmt.Sprintf("fibonacci(%d)", fib_n))

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(env.EvalAst1(form).Int())
	}
}

func BenchmarkFibonacciClassicInterpreterBis(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst(fib_s))

	// alternative: extract the function fibonacci, and call it ourselves
	fun := env.ValueOf("fibonacci").Interface().(func(int) int)
	fun(fib_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(fib_n)
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
	ce := fast.New()
	c := ce.Comp
	ce.DeclVar("n", TypeOfInt, 0)

	addr := ce.AddressOfVar("n").Interface().(*int)

	fun := c.CompileAst(c.ParseAst("((n*2+3)&4 | 5 ^ 6) / (n|1)"))
	env := ce.PrepareEnv()
	fun(env)
	var ret r.Value

	// interpreted code performs only arithmetic - iteration performed here
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		*addr = b.N
		ret, _ = fun(env)
		total += int(ret.Int())
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithFastInterpreterBis(b *testing.B) {
	ce := fast.New()
	c := ce.Comp
	ce.Eval("var i, n, total int")

	addr := ce.AddressOfVar("n").Interface().(*int)

	// interpreted code performs iteration and arithmetic
	fun := c.CompileAst(c.ParseAst("total = 0; for i = 0; i < n; i++ { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }; total"))

	env := ce.PrepareEnv()
	fun(env)
	var ret r.Value

	total := 0
	*addr = b.N
	b.ResetTimer()
	ret, _ = fun(env)
	total += int(ret.Int())

	if verbose {
		println(total)
	}
}

func BenchmarkArithClassicInterpreter(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst("n:=0"))

	form := env.ParseAst("((n*2+3)&4 | 5 ^ 6) / (n|1)")

	value := env.ValueOf("n")
	var ret r.Value
	env.EvalAst(form)

	// interpreted code performs only arithmetic - iteration performed here
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		value.SetInt(int64(b.N))
		ret, _ = env.EvalAst(form)
		total += int(ret.Int())
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithClassicInterpreterBis(b *testing.B) {
	ir := classic.New()
	ir.EvalAst(ir.ParseAst("var n, total int"))

	// interpreted code performs iteration and arithmetic
	form := ir.ParseAst("total = 0; for i:= 0; i < n; i++ { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }; total")

	value := ir.ValueOf("n")
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

func BenchmarkCollatzFastInterpreter(b *testing.B) {
	ce := fast.New()
	c := ce.Comp
	ce.DeclVar("n", TypeOfInt, 0)

	addr := ce.AddressOfVar("n").Interface().(*int)

	fun := c.CompileAst(c.ParseAst("for n > 1 { if n&1 != 0 { n = ((n * 3) + 1) / 2 } else { n = n / 2 } }"))
	env := ce.PrepareEnv()
	fun(env)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		*addr = collatz_n
		fun(env)
	}
}

func BenchmarkCollatzClassicInterpreter(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst("var n int"))
	n := env.ValueOf("n")

	form := env.ParseAst("for n > 1 { if n&1 != 0 { n = ((n * 3) + 1) / 2 } else { n = n / 2 } }")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.SetInt(collatz_n)
		env.EvalAst(form)
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
	if verbose {
		println(total)
	}
}

func BenchmarkSumFastInterpreter(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, total uint")
	ce.DeclConst("n", nil, uint(sum_n))
	ce.Apply()

	fun := ce.Compile("total = 0; for i = 1; i <= n; i++ { total += i }; total")
	env := ce.PrepareEnv()
	fun(env)

	var total uint
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret, _ := fun(env)
		total += uint(ret.Uint())
	}
	if verbose {
		println(total)
	}
}

func BenchmarkSumFastInterpreterBis(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, total uint")
	ce.DeclConst("n", nil, uint(sum_n))
	ce.Apply()

	fun := ce.Compile("for i = 1; i <= n; i++ { total += i }")
	env := ce.PrepareEnv()
	fun(env)
	total := ce.AddressOfVar("total").Interface().(*uint)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		*total = 0
		fun(env)
	}
	if verbose {
		println(*total)
	}
}

func BenchmarkSumClassicInterpreter(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.ParseAst("var i, n, total int"))
	env.ValueOf("n").SetInt(sum_n)
	form := env.ParseAst("total = 0; for i = 1; i <= n; i++ { total += i }; total")

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(env.EvalAst1(form).Int())
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
