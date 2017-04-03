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
 * all_test.go
 *
 *  Created on: Mar 06 2017
 *      Author: Massimiliano Ghilardi
 */
package main

import (
	"go/ast"
	"go/token"
	r "reflect"
	"testing"

	. "github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/compiler"
	"github.com/cosmos72/gomacro/constants"
	ir "github.com/cosmos72/gomacro/interpreter"
)

type TestFor int

const (
	I TestFor = 1 << iota
	C TestFor = 0 // temporarily DISABLE compiler tests
	A TestFor = I | C
)

type TestCase struct {
	testfor TestFor
	name    string
	program string
	result0 interface{}
	results []interface{}
}

func TestCompiler(t *testing.T) {
	env := ir.New()
	comp := compiler.New()
	for _, test := range tests {
		if test.testfor&C != 0 {
			c := test
			t.Run(c.name, func(t *testing.T) { c.compile(t, comp, env) })
		}
	}
}

func TestInterpreter(t *testing.T) {
	env := ir.New()
	// env.Options |= OptDebugCallStack | OptDebugPanicRecover
	for _, test := range tests {
		if test.testfor&I != 0 {
			c := test
			t.Run(c.name, func(t *testing.T) { c.interpret(t, env) })
		}
	}
}

func (c *TestCase) compile(t *testing.T, comp *compiler.CompEnv, env *ir.Env) {
	// parse + macroexpansion phase
	form := env.ParseAst(c.program)

	// compile phase
	f := comp.CompileAst(form)
	rets := compiler.PackValues(comp.Run(f))

	c.compareResults(t, rets)
}

func (c *TestCase) interpret(t *testing.T, env *ir.Env) {
	// parse + macroexpansion phase
	form := env.ParseAst(c.program)
	// eval phase
	rets := ir.PackValues(env.EvalAst(form))

	c.compareResults(t, rets)
}

const sum_s = "func sum(n int) int { total := 0; for i := 1; i <= n; i++ { total += i }; return total }"
const fib_s = "func fibonacci(n uint) uint { if n <= 2 { return 1 }; return fibonacci(n-1) + fibonacci(n-2) }"

var ti = r.StructOf(
	[]r.StructField{
		r.StructField{Name: "\u0080", Type: r.TypeOf((*interface{})(nil)).Elem()},
		r.StructField{Name: "String", Type: r.TypeOf((*func() string)(nil)).Elem()},
	},
)
var si = r.Zero(ti).Interface()

var tests = []TestCase{
	TestCase{A, "1+1", "1+1", 1 + 1, nil},
	TestCase{C, "1+'A'", "1+'A'", 'B', nil},
	TestCase{I, "1+'A'", "1+'A'", 66, nil}, // interpreter is not accurate in this case... returns <int> instead of <int32>
	TestCase{I, "int8+1", "int8(1)+1", int8(2), nil},
	TestCase{I, "int8_overflow", "int8(64)+64", int8(-128), nil},
	TestCase{I, "interface", "type Stringer interface { String() string }; var s Stringer", si, nil},
	TestCase{A, "string", "\"foobar\"", "foobar", nil},
	TestCase{A, "expr_and", "3 & 6", 3 & 6, nil},
	TestCase{A, "expr_or", "7 | 8", 7 | 8, nil},
	TestCase{A, "expr_xor", "0x1f ^ 0xf1", 0x1f ^ 0xf1, nil},
	TestCase{A, "expr_arith", "((1+2)*3^4|99)%112", ((1+2)*3 ^ 4 | 99) % 112, nil},
	TestCase{A, "complex_1", "7i", 7i, nil},
	TestCase{A, "complex_2", "0.5+1.75i", 0.5 + 1.75i, nil},
	TestCase{A, "complex_3", "1i * 2i", 1i * 2i, nil},
	TestCase{A, "const_1", "const _ = 11", 11, nil},
	TestCase{A, "const_2", "const c = 0xff&555+23/12.2", 0xff&555 + 23/12.2, nil},
	TestCase{A, "const_3", "const c2 = 0.1+0.2", float64(0.1) + float64(0.2), nil},    // the interpreter is not accurate in this case... missing exact arithmetic on constants
	TestCase{A, "const_4", "const c3 = c2/3", (float64(0.1) + float64(0.2)) / 3, nil}, // the interpreter is not accurate in this case... missing exact arithmetic on constants
	TestCase{A, "var_1", "var _ bool", false, nil},
	TestCase{A, "var_2", "var _ uint8 = 7", uint8(7), nil},
	TestCase{A, "var_3", "var _ uint16 = 65535", uint16(65535), nil},
	TestCase{A, "var_4", "var v uint32 = 99", uint32(99), nil},
	TestCase{A, "var_5", "var _ string", "", nil},
	TestCase{A, "var_6", "var _ float32", float32(0), nil},
	TestCase{A, "var_7", "var _ complex64", complex64(0), nil},
	TestCase{A, "var_8", "var err error", nil, nil},
	TestCase{A, "var_pointer", "var _ *string", (*string)(nil), nil},
	TestCase{A, "var_map", "var _ *map[error]bool", (*map[error]bool)(nil), nil},
	TestCase{A, "var_slice", "var _ []byte", ([]byte)(nil), nil},
	TestCase{A, "var_array", "var _ [2][]rune", [2][]rune{}, nil},
	TestCase{A, "var_interface{}", "var _ interface{} = 1", 1, nil},
	TestCase{A, "type_int8", "type _ int8", r.TypeOf(int8(0)), nil},
	TestCase{A, "type_complicated", "type _ func(int,int) func(error, func(bool)) string", r.TypeOf((func(int, int) func(error, func(bool)) string)(nil)), nil},
	TestCase{A, "type_struct", "type Pair struct { A, B int }", r.TypeOf(struct{ A, B int }{}), nil},
	TestCase{A, "struct_1", "var pair Pair", struct{ A, B int }{0, 0}, nil},
	TestCase{I, "struct_2", "pair.A, pair.B = 1, 2; pair", struct{ A, B int }{1, 2}, nil},
	TestCase{I, "pointer", "var p = 1.25; if *&p != p { p = -1 }; p", 1.25, nil},
	TestCase{I, "defer_1", "v = 0; func testdefer(x uint32) { if x != 0 { defer func() { v = x }() } }; testdefer(29); v", uint32(29), nil},
	TestCase{I, "defer_2", "v = 12; testdefer(0); v", uint32(12), nil},
	TestCase{I, "make_chan", "cx := make(chan interface{}, 2)", make(chan interface{}, 2), nil},
	TestCase{I, "make_map", "m := make(map[rune]bool)", make(map[rune]bool), nil},
	TestCase{I, "make_slice", "y := make([]uint8, 7); y[0] = 100; y[3] = 103; y", []uint8{100, 0, 0, 103, 0, 0, 0}, nil},
	TestCase{I, "expr_slice", "y = y[:4]", []uint8{100, 0, 0, 103}, nil},
	TestCase{I, "expr_slice3", "y = y[:3:4]", []uint8{100, 0, 0}, nil},
	TestCase{I, "for_range_chan", "i := 0; c := make(chan int, 2); c <- 1; c <- 2; close(c); for e := range c { i += e }; i", 3, nil},
	TestCase{I, "function", "func ident(x uint) uint { return x }; ident(42)", uint(42), nil},
	TestCase{I, "function_variadic", "func list_args(args ...interface{}) []interface{} { args }; list_args('x', 'y', 'z')", []interface{}{'x', 'y', 'z'}, nil},
	TestCase{I, "fibonacci", fib_s + "; fibonacci(13)", uint(233), nil},
	TestCase{I, "import", "import \"fmt\"", "fmt", nil},
	TestCase{I, "literal_struct", "Pair{A: 73, B: 94}", struct{ A, B int }{A: 73, B: 94}, nil},
	TestCase{I, "literal_array", "[3]int{1,2:3}", [3]int{1, 0, 3}, nil},
	TestCase{I, "literal_map", "map[int]string{1: \"foo\", 2: \"bar\"}", map[int]string{1: "foo", 2: "bar"}, nil},
	TestCase{I, "literal_slice", "[]rune{'a','b','c'}", []rune{'a', 'b', 'c'}, nil},
	TestCase{I, "method_on_ptr", "func (p *Pair) SetLhs(a int) { p.A = a }; pair.SetLhs(8); pair.A", 8, nil},
	TestCase{I, "method_on_value", "func (p Pair) SetLhs(a int) { p.A = a }; pair.SetLhs(11); pair.A", 8, nil}, // method on value gets a copy of the receiver - changes to not propagate
	TestCase{I, "multiple_values_1", "func twins(x float32) (float32,float32) { return x, x+1 }; twins(17.0)", nil, []interface{}{float32(17.0), float32(18.0)}},
	TestCase{I, "multiple_values_2", "func twins2(x float32) (float32,float32) { return twins(x) }; twins2(19.0)", nil, []interface{}{float32(19.0), float32(20.0)}},
	TestCase{A, "pred_bool_1", "false==false && true==true && true!=false", true, nil},
	TestCase{A, "pred_bool_2", "false!=false || true!=true || true==false", false, nil},
	TestCase{A, "pred_int", "1==1 && 1<=1 && 1>=1 && 1!=2 && 1<2 && 2>1 || 0==1", true, nil},
	TestCase{A, "pred_string_1", `"x"=="x" && "x"<="x" && "x">="x" && "x"!="y" && "x"<"y" && "y">"x"`, true, nil},
	TestCase{A, "pred_string_2", `"x"!="x" || "y"!="y" || "x">="y" || "y"<="x"`, false, nil},
	TestCase{I, "recover", `var vpanic interface{}
		func test_recover(rec bool, panick interface{}) {
			defer func() {
				if rec {
					vpanic = recover()
				}
			}()
			panic(panick)
		}
		test_recover(true, -3)
		vpanic`, -3, nil},
	TestCase{I, "recover_nested_1", `var vpanic2, vpanic3 interface{}
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
	TestCase{I, "recover_nested_2", `vpanic, vpanic2, vpanic3 = nil, nil, nil
		test_nested_recover(true, -5)
		Values(vpanic, vpanic2, vpanic3)
		`, nil, []interface{}{-5, -5, nil}},
	TestCase{I, "send_recv", "cx <- \"x\"; <-cx", nil, []interface{}{"x", true}},
	TestCase{I, "sum", sum_s + "; sum(100)", 5050, nil},

	TestCase{I, "select_1", "cx <- 1; { var x interface{}; select { case x=<-cx: x; default: } }", 1, nil},
	TestCase{I, "select_2", "cx <- m; select { case x:=<-cx: x; default: }", make(map[rune]bool), nil},
	TestCase{I, "select_3", "select { case cx<-1: 1; default: 0 }", 1, nil},
	TestCase{I, "select_4", "select { case cx<-2: 2; default: 0 }", 2, nil},
	TestCase{I, "select_5", "select { case cx<-3: 3; default: 0 }", 0, nil},
	TestCase{I, "select_6", "select { case cx<-4: 4; case x:=<-cx: x; default: 0 }", 1, nil},

	TestCase{I, "switch_1", "switch { case false: 0; default: 1 }", 1, nil},
	TestCase{I, "switch_2", "switch v:=20; v { case 20: '@' }", '@', nil},
	TestCase{I, "switch_fallthrough", "switch 0 { default: fallthrough; case 1: 10; fallthrough; case 2: 20 }", 20, nil},

	TestCase{I, "typeswitch_1", "var x interface{} = \"abc\"; switch y := x.(type) { default: 0; case string: 1 }", 1, nil},
	TestCase{I, "typeswitch_2", "switch x.(type) { default: 0; case interface{}: 2 }", 2, nil},
	TestCase{I, "typeswitch_3", "switch x.(type) { default: 0; case int: 3 }", 0, nil},
	TestCase{I, "typeswitch_4", "switch nil.(type) { default: 0; case nil: 4 }", 4, nil},

	TestCase{I, "quote_1", "quote{7}", &ast.BasicLit{Kind: token.INT, Value: "7"}, nil},
	TestCase{I, "quote_2", "quote{x}", &ast.Ident{Name: "x"}, nil},
	TestCase{I, "quote_3", "ab:=quote{a;b}", &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "a"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "b"}},
	}}, nil},
	TestCase{I, "quote_4", "quote{\"foo\"+\"bar\"}", &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.STRING, Value: "\"foo\""},
		Y:  &ast.BasicLit{Kind: token.STRING, Value: "\"bar\""},
	}, nil},
	TestCase{I, "quasiquote", "~`{1+~,{2+3}}", &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.INT, Value: "1"},
		Y:  &ast.BasicLit{Kind: token.INT, Value: "5"},
	}, nil},
	TestCase{I, "unquote_splice", "~`{~,@ab ; c}", &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "a"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "b"}},
		&ast.ExprStmt{X: &ast.Ident{Name: "c"}},
	}}, nil},
	TestCase{I, "macro", "macro second_arg(a,b,c interface{}) interface{} { return b }; 0", 0, nil},
	TestCase{I, "macro_call", "v = 98; second_arg;1;v;3", uint32(98), nil},
	TestCase{I, "macro_nested", "second_arg;1;{second_arg;2;3;4};5", 3, nil},
	TestCase{I, "values", "Values(3,4,5)", nil, []interface{}{3, 4, 5}},
	TestCase{I, "eval", "Eval(Values(3,4,5))", 3, nil},
	TestCase{I, "eval_quote", "Eval(quote{Values(3,4,5)})", nil, []interface{}{3, 4, 5}},
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
	if actualv == constants.Nil || actualv == constants.None {
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
	t.Errorf("expected %v <%T>, found %v <%T>\n", expected, expected, actual, actual)
}
