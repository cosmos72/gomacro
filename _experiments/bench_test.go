// +build gomacro_jit

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * bench_test.go
 *
 *  Created on: Jun 06 2018
 *      Author: Massimiliano Ghilardi
 */
package experiments

import (
	"testing"
	"unsafe"

	"github.com/cosmos72/gomacro/_experiments/bytecode_interfaces"
	"github.com/cosmos72/gomacro/_experiments/bytecode_values"
	"github.com/cosmos72/gomacro/_experiments/closure_interfaces"
	"github.com/cosmos72/gomacro/_experiments/closure_ints"
	"github.com/cosmos72/gomacro/_experiments/closure_maps"
	"github.com/cosmos72/gomacro/_experiments/closure_values"
	"github.com/cosmos72/gomacro/_experiments/jit"
)

func arithJitEmulate(uenv *uint64) {
	env := (*[3]int64)(unsafe.Pointer(uenv))
	a := env[0]
	a *= 2
	a += 3
	a |= 4
	a &^= 5
	a ^= 6
	b := env[0]
	b &= 2
	b |= 1
	a /= b
	env[1] = a
}

func BenchmarkArithJitEmul(b *testing.B) {
	benchArithJit(b, arithJitEmulate)
}

func BenchmarkArithJit(b *testing.B) {
	if !jit.SUPPORTED {
		b.SkipNow()
		return
	}
	benchArithJit(b, jit.DeclArith(5))
}

func benchArithJit(b *testing.B, f func(*uint64)) {
	total := 0
	var env [5]uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env[0] = uint64(b.N)
		f(&env[0])
		total += int(env[1])
	}
	if verbose {
		println(total)
	}
}

// --------------------------------------------------------------

func BenchmarkSumJit(b *testing.B) {
	if !jit.SUPPORTED {
		b.SkipNow()
		return
	}
	sum := jit.DeclSum()
	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(sum_arg)
	}
}

func TestFibonacciClosureInts(t *testing.T) {
	env := closure_ints.NewEnv(nil)
	f := closure_ints.DeclFibonacci(env)

	expected := fibonacci(fib_arg)
	actual := f(fib_arg)
	if actual != expected {
		t.Errorf("expecting %v, found %v\n", expected, actual)
	}
}

func BenchmarkFibonacciClosureInts(b *testing.B) {
	env := closure_ints.NewEnv(nil)
	fib := closure_ints.DeclFibonacci(env)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(fib_arg)
	}
}

func BenchmarkFibonacciClosureValues(b *testing.B) {
	env := closure_values.NewEnv(nil)
	fib := closure_values.DeclFibonacci(env, 0)
	n := r.ValueOf(fib_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func BenchmarkFibonacciClosureInterfaces(b *testing.B) {
	env := closure_interfaces.NewEnv(nil)
	fib := closure_interfaces.DeclFibonacci(env, 0)
	var n interface{} = fib_arg

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func BenchmarkFibonacciClosureMaps(b *testing.B) {
	env := closure_maps.NewEnv(nil)
	fib := closure_maps.DeclFibonacci(env, "fib")
	n := r.ValueOf(fib_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func TestFibonacciClosureInts(t *testing.T) {
	env := closure_ints.NewEnv(nil)
	f := closure_ints.DeclFibonacci(env)

	expected := fibonacci(fib_arg)
	actual := f(fib_arg)
	if actual != expected {
		t.Errorf("expecting %v, found %v\n", expected, actual)
	}
}

func BenchmarkFibonacciClosureInts(b *testing.B) {
	env := closure_ints.NewEnv(nil)
	fib := closure_ints.DeclFibonacci(env)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(fib_arg)
	}
}

func BenchmarkFibonacciClosureValues(b *testing.B) {
	env := closure_values.NewEnv(nil)
	fib := closure_values.DeclFibonacci(env, 0)
	n := r.ValueOf(fib_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func BenchmarkFibonacciClosureInterfaces(b *testing.B) {
	env := closure_interfaces.NewEnv(nil)
	fib := closure_interfaces.DeclFibonacci(env, 0)
	var n interface{} = fib_arg

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}

func BenchmarkFibonacciClosureMaps(b *testing.B) {
	env := closure_maps.NewEnv(nil)
	fib := closure_maps.DeclFibonacci(env, "fib")
	n := r.ValueOf(fib_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += fib(n)
	}
}
func BenchmarkSumBytecodeValues(b *testing.B) {
	sum := bytecode_values.BytecodeSum(sum_arg)
	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(sum.Exec(0)[0].Int())
	}
}

func BenchmarkSumBytecodeInterfaces(b *testing.B) {
	p := bytecode_interfaces.BytecodeSum(sum_arg)
	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += p.Exec(0)[0].(int)
	}
}

func off_TestSumClosureInts(t *testing.T) {
	env := closure_ints.NewEnv(nil)
	f := closure_ints.DeclSum(env)

	expected := sum(sum_arg)
	actual := f(sum_arg)
	if actual != expected {
		t.Errorf("expecting %v, found %v\n", expected, actual)
	}
}

func BenchmarkSumClosureInts(b *testing.B) {
	env := closure_ints.NewEnv(nil)
	sum := closure_ints.DeclSum(env)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(sum_arg)
	}
}

func BenchmarkSumClosureValues(b *testing.B) {
	env := closure_values.NewEnv(nil)
	sum := closure_values.DeclSum(env, 0)
	n := r.ValueOf(sum_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}

func BenchmarkSumClosureInterfaces(b *testing.B) {
	env := closure_interfaces.NewEnv(nil)
	sum := closure_interfaces.DeclSum(env, 0)
	var n interface{} = sum_arg

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}

func BenchmarkSumClosureMaps(b *testing.B) {
	env := closure_maps.NewEnv(nil)
	sum := closure_maps.DeclSum(env, "sum")
	n := r.ValueOf(sum_arg)

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(n)
	}
}
