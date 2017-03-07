/*
 * gomacro - A Go intepreter with Lisp-like macros
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
 * interpreter_test.go
 *
 *  Created on: Mar 06 2017
 *      Author: Massimiliano Ghilardi
 */
package interpreter

import (
	r "reflect"
	"testing"
)

type TestCase struct {
	name    string
	program string
	result0 interface{}
	results []interface{}
}

const sum_s = "func sum(n int) int { total := 0; for i := 1; i <= n; i++ { total += i }; return total }"
const fib_s = "func fibonacci(n uint) uint { if n <= 2 { return 1 }; return fibonacci(n-1) + fibonacci(n-2) }"

var testcases = []TestCase{
	TestCase{"1+1", "1+1", 2, nil},
	TestCase{"int8+1", "int8(1)+1", int8(2), nil},
	TestCase{"int8(128)", "int8(64)+64", int8(-128), nil},
	TestCase{"struct", "var pair struct { A, B int }; pair.A, pair.B = 1, 2; pair", struct{ A, B int }{1, 2}, nil},
	TestCase{"pointer", "var x = 1.25; if *&x != x { x = -1 }; x", 1.25, nil},
	TestCase{"function", "func ident(x uint) uint { return x }; ident(42)", uint(42), nil},
	TestCase{"sum", sum_s + "; sum(100)", 5050, nil},
	TestCase{"fibonacci", fib_s + "; fibonacci(13)", uint(233), nil},
	TestCase{"multiple-values", "func twins(x float32) (float32,float32) { return x, x+1 }; twins(17.0)", nil, []interface{}{float32(17.0), float32(18.0)}},
}

func TestInterpreter(t *testing.T) {
	env := New()
	for _, testcase := range testcases {
		c := testcase
		t.Run(c.name, func(t *testing.T) { c.run(t, env) })
	}
}

func (c *TestCase) run(t *testing.T, env *Env) {
	nodes := env.Parse(c.program)
	rets := packValues(env.EvalList(nodes))
	c.compareResults(t, rets)
}

func (c *TestCase) compareResults(t *testing.T, actual []r.Value) {
	expected := c.results
	if len(expected) == 0 {
		expected = []interface{}{c.result0}
	}
	if len(actual) != len(expected) {
		t.Fail()
		return
	}
	for i := range actual {
		c.compareResult(t, actual[i], expected[i])
	}
}

func (c *TestCase) compareResult(t *testing.T, actual r.Value, expected interface{}) {
	if actual == Nil || actual == None {
		if expected != nil {
			t.Fail()
		}
		return
	}
	if !r.DeepEqual(actual.Interface(), expected) {
		t.Fail()
	}
}

// fibonacci(30):
// compiled: 832040, elapsed time: 0.00285 s
// interpreted: 832040 <uint> // eval time 3.371358 s
// the interpreter is 1180 times slower than compiled code...

func BenchmarkFibonacciInterpreter(b *testing.B) {
	env := New()
	node := env.Parse1(fib_s)
	env.Eval1(node)
	node = env.Parse1("fibonacci(30)")

	b.ResetTimer()
	var total uint
	for i := 0; i < b.N; i++ {
		total += uint(env.Eval1(node).Uint())
	}
}

func BenchmarkFibonacciCompiler(b *testing.B) {
	var total uint
	for i := 0; i < b.N; i++ {
		total += fibonacci(30)
	}
}

func BenchmarkSumInterpreter(b *testing.B) {
	env := New()
	node := env.Parse1(sum_s)
	env.Eval1(node)
	node = env.Parse1("sum(10000)")

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(env.Eval1(node).Int())
	}
}

func BenchmarkSumCompiler(b *testing.B) {
	var total int
	for i := 0; i < b.N; i++ {
		total += sum(10000)
	}
}

func sum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

func fibonacci(n uint) uint {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
