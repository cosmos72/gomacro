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
	collatz_n   = 837799 // sequence climbs to 1487492288, which also fits 32-bit ints
	sum_n       = 1000
	fib_n       = 12
	bigswitch_n = 100
)

var verbose = false

/*
	--------- 2017-05-21: results on Intel Core i7 4770 ---------------

	BenchmarkFibonacciCompiler-8            	 3000000	       501 ns/op
	BenchmarkFibonacciFast-8                	  100000	     15774 ns/op
	BenchmarkFibonacciFast2-8               	  100000	     15141 ns/op
	BenchmarkFibonacciClassic-8             	    2000	    915990 ns/op
	BenchmarkFibonacciClassic2-8            	    2000	    912180 ns/op
	BenchmarkFibonacciClosureValues-8       	    5000	    259074 ns/op
	BenchmarkFibonacciClosureInterfaces-8   	   10000	    193098 ns/op
	BenchmarkFibonacciClosureMaps-8         	    5000	    358345 ns/op
	BenchmarkShellSortCompiler-8            	20000000	        74.0 ns/op
	BenchmarkShellSortFast-8                	  200000	      7790 ns/op
	BenchmarkShellSortClassic-8             	    5000	    276673 ns/op
	BenchmarkSwitchCompiler-8               	 1000000	      2363 ns/op
	BenchmarkSwitchFast-8                   	   50000	     37773 ns/op
	BenchmarkSwitchClassic-8                	     500	   3454461 ns/op
	BenchmarkArithCompiler1-8               	200000000	         8.41 ns/op
	BenchmarkArithCompiler2-8               	200000000	         8.41 ns/op
	BenchmarkArithFast-8                    	50000000	        30.8 ns/op
	BenchmarkArithFast2-8                   	30000000	        50.6 ns/op
	BenchmarkArithFastConst-8               	100000000	        15.2 ns/op
	BenchmarkArithFastCompileLoop-8         	  100000	     21442 ns/op
	BenchmarkArithClassic-8                 	 1000000	      1686 ns/op
	BenchmarkArithClassic2-8                	  500000	      2916 ns/op
	BenchmarkCollatzCompiler-8              	 5000000	       265 ns/op
	BenchmarkCollatzFast-8                  	  200000	     11812 ns/op
	BenchmarkCollatzClassic-8               	    2000	    654139 ns/op
	BenchmarkCollatzBytecodeInterfaces-8    	   50000	     30203 ns/op
	BenchmarkCollatzClosureValues-8         	  100000	     16570 ns/op
	BenchmarkSumCompiler-8                  	 5000000	       294 ns/op
	BenchmarkSumFast-8                      	  100000	     20789 ns/op
	BenchmarkSumFast2-8                     	  100000	     20720 ns/op
	BenchmarkSumClassic-8                   	    1000	   1223624 ns/op
	BenchmarkSumBytecodeValues-8            	   20000	     76201 ns/op
	BenchmarkSumBytecodeInterfaces-8        	   30000	     53031 ns/op
	BenchmarkSumClosureValues-8             	   30000	     41124 ns/op
	BenchmarkSumClosureInterfaces-8         	   10000	    147109 ns/op
	BenchmarkSumClosureMaps-8               	   20000	     93320 ns/op
*/

// ---------------------- recursion: fibonacci ----------------------

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
	if verbose {
		println(total)
	}
}

func BenchmarkFibonacciFast(b *testing.B) {
	ce := fast.New()
	ce.Eval(fibonacci_source_string)

	// compile the call to fibonacci(fib_n)
	expr := ce.Compile(fmt.Sprintf("fibonacci(%d)", fib_n))
	fun := expr.Fun.(func(*fast.Env) int)
	env := ce.PrepareEnv()
	fun(env)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(env)
	}
}

func BenchmarkFibonacciFast2(b *testing.B) {
	ce := fast.New()
	ce.Eval(fibonacci_source_string)

	// alternative: extract the function fibonacci, and call it ourselves
	//
	// ValueOf is the method to retrieve constants, functions and variables from the classic and fast interpreters
	// (if you set the same interpreter variable repeatedly, use the address returned by AddressOfVar)
	fun := ce.ValueOf("fibonacci").Interface().(func(int) int)
	fun(fib_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(fib_n)
	}
}

func BenchmarkFibonacciClassic(b *testing.B) {
	env := classic.New()
	env.Eval(fibonacci_source_string)

	// compile the call to fibonacci(fib_n)
	form := env.Parse(fmt.Sprintf("fibonacci(%d)", fib_n))

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(env.EvalAst1(form).Int())
	}
}

func BenchmarkFibonacciClassic2(b *testing.B) {
	env := classic.New()
	env.Eval(fibonacci_source_string)

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

// ---------------------- arrays: shellsort ------------------------

var shellshort_gaps = []int{701, 301, 132, 57, 23, 10, 4, 1}

func shellsort(v []int) {
	var i, j, n, gap, temp int
	n = len(v)
	for _, gap = range shellshort_gaps {
		for i = gap; i < n; i++ {
			temp = v[i]
			for j = i; j >= gap && v[j-gap] > temp; j -= gap {
				v[j] = v[j-gap]
			}
			v[j] = temp
		}
	}
}

var sort_data = []int{97, 89, 3, 4, 7, 0, 36, 79, 1, 12, 2, 15, 70, 18, 35, 70, 15, 73}

func BenchmarkShellSortCompiler(b *testing.B) {
	benchmark_sort(b, shellsort)
}

func BenchmarkShellSortFast(b *testing.B) {
	ce := fast.New()
	ce.Eval(shellsort_source_string)

	// extract the function shellsort()
	sort := ce.ValueOf("shellsort").Interface().(func([]int))

	benchmark_sort(b, sort)
}

func BenchmarkShellSortClassic(b *testing.B) {
	env := classic.New()
	env.Eval(shellsort_source_string)

	// extract the function shellsort()
	sort := env.ValueOf("shellsort").Interface().(func([]int))

	benchmark_sort(b, sort)
}

func benchmark_sort(b *testing.B, sort func([]int)) {
	// call sort once for warm-up
	v := make([]int, len(sort_data))
	copy(v, sort_data)
	sort(v)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(v, sort_data)
		sort(v)
	}
	if verbose {
		fmt.Println(v)
	}
}

// ---------------------- switch ------------------------

func bigswitch(n int) int {
	for i := 0; i < 1000; i++ {
		switch n & 15 {
		case 0:
			n++
		case 1:
			n += 2
		case 2:
			n += 3
		case 3:
			n += 4
		case 4:
			n += 5
		case 5:
			n += 6
		case 6:
			n += 7
		case 7:
			n += 8
		case 8:
			n += 9
		case 9:
			n += 10
		case 10:
			n += 11
		case 11:
			n += 12
		case 12:
			n += 13
		case 13:
			n += 14
		case 14:
			n += 15
		case 15:
			n--
		}
	}
	return n
}

func BenchmarkSwitchCompiler(b *testing.B) {
	var total int
	for i := 0; i < b.N; i++ {
		total += bigswitch(bigswitch_n)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkSwitchFast(b *testing.B) {
	ce := fast.New()
	ce.Eval(switch_source_string)

	fun := ce.ValueOf("bigswitch").Interface().(func(int) int)
	fun(bigswitch_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(bigswitch_n)
	}
}

func BenchmarkSwitchClassic(b *testing.B) {
	env := classic.New()
	env.Eval(switch_source_string)

	fun := env.ValueOf("bigswitch").Interface().(func(int) int)
	fun(bigswitch_n)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fun(bigswitch_n)
	}
}

// ---------------- simple arithmetic ------------------

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

func BenchmarkArithFast(b *testing.B) {
	ce := fast.New()
	ce.DeclVar("n", nil, int(0))

	addr := ce.AddressOfVar("n").Interface().(*int)

	expr := ce.Compile("((n*2+3)&4 | 5 ^ 6) / (n|1)")
	fun := expr.Fun.(func(*fast.Env) int)
	env := ce.PrepareEnv()
	fun(env)

	// interpreted code performs only arithmetic - iteration performed here
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		*addr = b.N
		total += fun(env)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkArithFast2(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, n, total int")

	n := ce.AddressOfVar("n").Interface().(*int)
	total := ce.AddressOfVar("total").Interface().(*int)

	// interpreted code performs iteration and arithmetic
	fun := ce.Compile("for i = 0; i < n; i++ { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }").AsX()
	env := ce.PrepareEnv()
	fun(env)

	b.ResetTimer()

	*n = b.N
	*total = 0
	fun(env)

	if verbose {
		println(*total)
	}
}

func BenchmarkArithFastConst(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, total int")
	// "cheat" a bit and declare n as a constant. checks if constant propagation works :)
	ce.DeclConst("n", nil, int(b.N))
	total := ce.AddressOfVar("total").Interface().(*int)

	// interpreted code performs iteration and arithmetic
	fun := ce.Compile("for i = 0; i < n; i++ { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }").AsX()
	env := ce.PrepareEnv()
	fun(env)

	b.ResetTimer()

	*total = 0
	fun(env)

	if verbose {
		println(*total)
	}
}

func BenchmarkArithFastCompileLoop(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, n, total int")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ce.Compile("total = 0; for i = 0; i < n; i++ { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }; total")
	}
}

func BenchmarkArithClassic(b *testing.B) {
	env := classic.New()
	env.Eval("n:=0")

	form := env.Parse("((n*2+3)&4 | 5 ^ 6) / (n|1)")

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

func BenchmarkArithClassic2(b *testing.B) {
	env := classic.New()
	env.Eval("var n, total int")

	// interpreted code performs iteration and arithmetic
	form := env.Parse("total = 0; for i:= 0; i < n; i++ { total += ((n*2+3)&4 | 5 ^ 6) / (n|1) }; total")

	value := env.ValueOf("n")
	env.EvalAst(form)

	b.ResetTimer()
	value.SetInt(int64(b.N))
	ret, _ := env.EvalAst(form)

	if verbose {
		println(ret.Int())
	}
}

// ---------------- collatz conjecture --------------------

func collatz(n uint) {
	for n > 1 {
		if n&1 != 0 {
			n = ((n * 3) + 1) >> 1
		} else {
			n >>= 1
		}
	}
}

func BenchmarkCollatzCompiler(b *testing.B) {
	var n uint = collatz_n
	for i := 0; i < b.N; i++ {
		collatz(n)
	}
}

func BenchmarkCollatzFast(b *testing.B) {
	ce := fast.New()
	ce.DeclVar("n", nil, uint(0))
	addr := ce.AddressOfVar("n").Interface().(*uint)

	fun := ce.Compile("for n > 1 { if n&1 != 0 { n = ((n * 3) + 1) >> 1 } else { n >>= 1 } }").AsX()
	env := ce.PrepareEnv()
	fun(env)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		*addr = collatz_n
		fun(env)
	}
}

func BenchmarkCollatzClassic(b *testing.B) {
	env := classic.New()
	env.EvalAst(env.Parse("var n uint"))
	addr := env.ValueOf("n").Addr().Interface().(*uint)

	form := env.Parse("for n > 1 { if n&1 != 0 { n = ((n * 3) + 1) >> 1 } else { n >>= 1 } }")
	env.EvalAst(form)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		*addr = collatz_n
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

// ------------- looping: sum the integers from 1 to N -------------------

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

func BenchmarkSumFast(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, total uint")
	ce.DeclConst("n", nil, uint(sum_n))

	expr := ce.Compile("total = 0; for i = 1; i <= n; i++ { total += i }; total")
	fun := expr.Fun.(func(*fast.Env) uint)
	env := ce.PrepareEnv()
	fun(env)

	var total uint
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += fun(env)
	}
	if verbose {
		println(total)
	}
}

func BenchmarkSumFast2(b *testing.B) {
	ce := fast.New()
	ce.Eval("var i, total uint")
	ce.DeclConst("n", nil, uint(sum_n))

	fun := ce.Compile("for i = 1; i <= n; i++ { total += i }").AsX()
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

func BenchmarkSumClassic(b *testing.B) {
	env := classic.New()
	env.Eval("var i, n, total int")
	env.ValueOf("n").SetInt(sum_n)
	form := env.Parse("total = 0; for i = 1; i <= n; i++ { total += i }; total")

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
