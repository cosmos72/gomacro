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
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"
	"testing"

	. "github.com/cosmos72/gomacro/ast2"
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
	TestCase{"int8_overflow", "int8(64)+64", int8(-128), nil},
	TestCase{"string", "\"foobar\"", "foobar", nil},
	TestCase{"var", "var v uint32 = 99", uint32(99), nil},
	TestCase{"pointer", "var p = 1.25; if *&p != p { p = -1 }; p", 1.25, nil},
	TestCase{"struct", "type Pair struct { A, B int }; var pair Pair; pair.A, pair.B = 1, 2; pair", struct{ A, B int }{1, 2}, nil},
	TestCase{"literal_struct", "Pair{A: 73, B: 94}", struct{ A, B int }{A: 73, B: 94}, nil},
	TestCase{"literal_array", "[3]int{1,2:3}", [3]int{1, 0, 3}, nil},
	TestCase{"literal_map", "map[int]string{1: \"foo\", 2: \"bar\"}", map[int]string{1: "foo", 2: "bar"}, nil},
	TestCase{"literal_slice", "[]rune{'a','b','c'}", []rune{'a', 'b', 'c'}, nil},
	TestCase{"make_chan", "c := make(chan interface{}, 2)", make(chan interface{}, 2), nil},
	TestCase{"make_map", "m := make(map[rune]bool)", make(map[rune]bool), nil},
	TestCase{"make_slice", "y := make([]uint8, 7); y[0] = 100; y[3] = 103; y", []uint8{100, 0, 0, 103, 0, 0, 0}, nil},
	TestCase{"expr_slice", "y = y[:4]", []uint8{100, 0, 0, 103}, nil},
	TestCase{"expr_slice3", "y = y[:3:4]", []uint8{100, 0, 0}, nil},
	TestCase{"send_recv", "c <- 'x'; <-c", nil, []interface{}{'x', true}},
	TestCase{"function", "func ident(x uint) uint { return x }; ident(42)", uint(42), nil},
	TestCase{"sum", sum_s + "; sum(100)", 5050, nil},
	TestCase{"fibonacci", fib_s + "; fibonacci(13)", uint(233), nil},
	TestCase{"multiple_values_1", "func twins(x float32) (float32,float32) { return x, x+1 }; twins(17.0)", nil, []interface{}{float32(17.0), float32(18.0)}},
	TestCase{"multiple_values_2", "func twins2(x float32) (float32,float32) { return twins(x) }; twins2(19.0)", nil, []interface{}{float32(19.0), float32(20.0)}},
	TestCase{"defer_1", "v = 0; func testdefer(x uint32) { if x != 0 { defer func() { v = x }() } }; testdefer(29); v", uint32(29), nil},
	TestCase{"defer_2", "v = 12; testdefer(0); v", uint32(12), nil},
	TestCase{"recover", `var vpanic interface{}
		func test_recover(rec bool, panick interface{}) {
			defer func() {
				if rec {
					vpanic = recover()
				}
			}()
			panic(panick)
		}
		test_recover(true, -3)
		vpanic
		`, -3, nil},
	TestCase{"nested_recover_1", `var vpanic2, vpanic3 interface{}
		func test_nested_recover(repanic bool, panick interface{}) {
			defer func() {
				vpanic = recover()
			}()
			defer func() {
				func() {
					vpanic3 = recover()
				}()
				vpanic2 = recover()
				if repanic {
					panic(vpanic2)
				}
			}()
			panic(panick)
		}
		test_nested_recover(false, -4)
		Values(vpanic, vpanic2, vpanic3)
		`, nil, []interface{}{nil, -4, nil}},
	TestCase{"nested_recover_2", `vpanic, vpanic2, vpanic3 = nil, nil, nil
		test_nested_recover(true, -5)
		Values(vpanic, vpanic2, vpanic3)
		`, nil, []interface{}{-5, -5, nil}},
	TestCase{"import", "import \"fmt\"", "fmt", nil},
	TestCase{"quote_1", "quote{7}", &ast.BasicLit{Kind: token.INT, Value: "7"}, nil},
	TestCase{"quote_2", "quote{x}", &ast.Ident{Name: "x"}, nil},
	TestCase{"quote_3", "ab:=quote{a;b}", &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "a"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "b"}},
	}}, nil},
	TestCase{"quote_4", "quote{\"foo\"+\"bar\"}", &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.STRING, Value: "\"foo\""},
		Y:  &ast.BasicLit{Kind: token.STRING, Value: "\"bar\""},
	}, nil},
	TestCase{"quasiquote", "~`{1+~,{2+3}}", &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.INT, Value: "1"},
		Y:  &ast.BasicLit{Kind: token.INT, Value: "5"},
	}, nil},
	TestCase{"unquote_splice", "~`{~,@ab ; c}", &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "a"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "b"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "c"}},
	}}, nil},
	TestCase{"macro", "macro second_arg(a,b,c interface{}) interface{} { return b }; 0", 0, nil},
	TestCase{"macro_call", "v = 98; second_arg;1;v;3", uint32(98), nil},
	TestCase{"macro_nested", "second_arg;1;{second_arg;2;3;4};5", 3, nil},
	TestCase{"values", "Values(3,4,5)", nil, []interface{}{3, 4, 5}},
	TestCase{"eval", "Eval(Values(3,4,5))", 3, nil},
	TestCase{"eval_quote", "Eval(quote{Values(3,4,5)})", nil, []interface{}{3, 4, 5}},
}

func TestInterpreter(t *testing.T) {

	env := New()
	// env.Options |= OptDebugCallStack | OptDebugPanicRecover
	for _, testcase := range testcases {
		c := testcase
		t.Run(c.name, func(t *testing.T) { c.run(t, env) })
	}
}

func (c *TestCase) run(t *testing.T, env *Env) {
	// parse phase
	form := env.ParseAst(c.program)
	// macroexpansion phase
	form, _ = env.MacroExpandAstCodewalk(form)
	// eval phase
	rets := packValues(env.EvalAst(form))

	c.compareResults(t, rets)
}

func (c *TestCase) compareResults(t *testing.T, actual []r.Value) {
	expected := c.results
	if len(expected) == 0 {
		expected = []interface{}{c.result0}
	}
	if len(actual) != len(expected) {
		c.fail(t, actual, expected)
		return
	}
	for i := range actual {
		c.compareResult(t, actual[i], expected[i])
	}
}

func (c *TestCase) compareResult(t *testing.T, actualv r.Value, expected interface{}) {
	if actualv == Nil || actualv == None {
		if expected != nil {
			c.fail(t, nil, expected)
		}
		return
	}
	actual := actualv.Interface()
	if !r.DeepEqual(actual, expected) {
		if r.TypeOf(actual) == r.TypeOf(expected) {
			if actualNode, ok := actual.(ast.Node); ok {
				if expectedNode, ok := expected.(ast.Node); ok {
					c.compareAst(t, ToAst(actualNode), ToAst(expectedNode))
					return
				}
			} else if actualv.Kind() == r.Chan {
				// for channels just check the type, length and capacity
				expectedv := r.ValueOf(expected)
				if actualv.Len() == expectedv.Len() && actualv.Cap() == expectedv.Cap() {
					return
				}
			}
		}
		c.fail(t, actual, expected)
	}
}

func (c *TestCase) compareAst(t *testing.T, actual Ast, expected Ast) {
	if r.TypeOf(actual) == r.TypeOf(expected) {
		switch actual := actual.(type) {
		case BadDecl, BadExpr, BadStmt:
			return
		case Ident:
			if actual.X.Name == expected.(Ident).X.Name {
				return
			}
		case BasicLit:
			actualp := actual.X
			expectedp := expected.(BasicLit).X
			if actualp.Kind == expectedp.Kind && actualp.Value == expectedp.Value {
				return
			}
		default:
			na := actual.Size()
			ne := expected.Size()
			if actual.Op() == expected.Op() && na == ne {
				for i := 0; i < na; i++ {
					c.compareAst(t, actual.Get(i), expected.Get(i))
				}
				return
			}
		}
	}
	c.fail(t, actual, expected)
}

func (c *TestCase) fail(t *testing.T, actual interface{}, expected interface{}) {
	fmt.Printf("--- FAIL: %s: expected %#v <%T>, found %#v <%T>\n", t.Name(), expected, expected, actual, actual)
	t.Fail()
}

// fibonacci(30):
// compiled: 832040, elapsed time: 0.00285 s
// interpreted: 832040 <uint> // eval time 3.371358 s
// the interpreter is 1180 times slower than compiled code...

func BenchmarkFibonacciInterpreter(b *testing.B) {
	env := New()
	env.EvalAst(env.ParseAst(fib_s))
	form := env.ParseAst("fibonacci(30)")

	b.ResetTimer()
	var total uint
	for i := 0; i < b.N; i++ {
		total += uint(env.EvalAst1(form).Uint())
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
	env.EvalAst(env.ParseAst(sum_s))
	form := env.ParseAst("sum(10000)")

	b.ResetTimer()
	var total int
	for i := 0; i < b.N; i++ {
		total += int(env.EvalAst1(form).Int())
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
