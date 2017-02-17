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
 * main.go
 *
 *  Created on: Feb 13, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	r "reflect"
	"strings"
)

const TemporaryFunctionName = "gorepl_temporary_function"

var Nil r.Value = r.ValueOf(nil)

type Binds map[string]r.Value

type Env struct {
	Binds
	Outer *Env
}

type Interpreter struct {
	*Env
	Packagename string
	Filename    string
	Fileset     *token.FileSet
	Parsermode  parser.Mode
	iotaOffset  int
	In          *bufio.Reader
	Out         io.Writer
	Eout        io.Writer
}

func New() *Interpreter {
	ir := Interpreter{}
	ir.PushEnv()
	ir.Packagename = "main"
	ir.Filename = "main.go"
	ir.Fileset = token.NewFileSet()
	ir.iotaOffset = 1
	addBuiltins(ir.Binds)
	return &ir
}

func (ir *Interpreter) PushEnv() {
	// fmt.Printf("debug: PushEnv() pushed new, empty bindings - outer bindings are %#v\n", ir.Env)
	ir.Env = &Env{make(Binds), ir.Env}
}

func (ir *Interpreter) PopEnv() {
	if ir.Env != nil {
		// fmt.Printf("debug: PushEnv() popped bindings %#v\n", ir.Env)
		ir.Env = ir.Env.Outer
	}
}

func (ir *Interpreter) Parse(src interface{}) (ast.Node, error) {
	expr, err := parser.ParseExprFrom(ir.Fileset, ir.Filename, src, 0)
	if err == nil {
		if ir.Parsermode == 0 {
			return expr, nil
		}
		return parser.ParseExprFrom(ir.Fileset, ir.Filename, src, ir.Parsermode)
	}
	str, ok := src.(string)
	if ok {
		str = strings.TrimLeft(str, " \f\t\r\n\v")
		firstWord := str
		space := strings.IndexAny(str, " \f\t\r\n\v")
		if space >= 0 {
			firstWord = str[:space]
		}
		switch firstWord {
		case "package":
			/* nothing to do */
		case "func", "var", "const", "type":
			str = fmt.Sprintf("package %s; %s", ir.Packagename, str)
		default:
			str = fmt.Sprintf("package %s; func %s() { %s }", ir.Packagename, TemporaryFunctionName, str)
		}
		src = str
	}
	return parser.ParseFile(ir.Fileset, ir.Filename, src, ir.Parsermode)
}

func (ir *Interpreter) PrintAst(out io.Writer, prefix string, node ast.Node) {
	fmt.Fprint(out, prefix)
	printer.Fprint(out, ir.Fileset, node)
	fmt.Fprintln(out)
}

func (ir *Interpreter) Print(out io.Writer, value r.Value) {
	if value == Nil {
		fmt.Fprint(out, "// no values\n")
		return
	}
	v := value.Interface()
	switch v.(type) {
	case uint, uint8, uint32, uint64, uintptr:
		fmt.Fprintf(out, "%d <%T>\n", v, v)
	default:
		fmt.Fprintf(out, "%#v <%T>\n", v, v)
	}
}

func (ir *Interpreter) Loop() {
	in := bufio.NewReader(os.Stdin)
	ir.Loop3(in, os.Stdout, os.Stderr)
}

func (ir *Interpreter) Loop1(in *bufio.Reader) {
	ir.Loop3(in, os.Stdout, os.Stderr)
}

func (ir *Interpreter) Loop2(in *bufio.Reader, out_and_err io.Writer) {
	ir.Loop3(in, out_and_err, out_and_err)
}

func (ir *Interpreter) Loop3(in *bufio.Reader, out io.Writer, eout io.Writer) {
	ir.In = in
	ir.Out = out
	ir.Eout = eout
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			fmt.Fprintln(eout, err)
			break
		}
		ast, err := ir.Parse(line)
		if err != nil {
			fmt.Fprintln(eout, err)
			continue
		}
		// ir.PrintAst(out, "parsed: ", ast)
		value, err := ir.Eval(ast)
		if err != nil {
			fmt.Fprintln(eout, err)
			continue
		}
		ir.Print(out, value)
	}
}

func main() {
	interpreter := New()
	// ir.Parsermode = parser.Trace
	interpreter.Loop()
}
