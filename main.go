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
	"go/token"
	"os"
	r "reflect"
	"strings"
)

const TemporaryFunctionName = "gorepl_temporary_function"

type Binds map[string]r.Value

type Parser struct {
	Packagename string
	Filename    string
	Fileset     *token.FileSet
	Parsermode  parser.Mode
}

type Env struct {
	*Parser
	binds      Binds
	Outer      *Env
	iotaOffset int
}

func NewParser() *Parser {
	p := Parser{}
	p.Packagename = "main"
	p.Filename = "main.go"
	p.Fileset = token.NewFileSet()
	return &p
}

func NewEnv(outer *Env) *Env {
	env := Env{}
	env.binds = make(map[string]r.Value)
	env.iotaOffset = 1
	if outer == nil {
		env.Parser = NewParser()
		env.addBuiltins()
		env.addInterpretedBuiltins()
	} else {
		env.Parser = outer.Parser
		env.Outer = outer
	}
	return &env
}

func (p *Parser) Parse(src interface{}) ast.Node {
	node, err := p.parseOrError(src)
	if err != nil {
		Errore(err)
		return nil
	}
	return node
}

func (p *Parser) parseOrError(src interface{}) (ast.Node, error) {
	expr, err := parser.ParseExprFrom(p.Fileset, p.Filename, src, 0)
	if err == nil {
		if p.Parsermode == 0 {
			return expr, nil
		}
		return parser.ParseExprFrom(p.Fileset, p.Filename, src, p.Parsermode)
	}
	str, ok := src.(string)
	if ok {
		str = strings.TrimSpace(str)
		str = skipComments(str)
		firstWord := str
		space := strings.IndexAny(str, " \f\t\r\n\v")
		if space >= 0 {
			firstWord = str[:space]
		}
		switch firstWord {
		case "package":
			/* nothing to do */
		case "func", "var", "const", "type":
			str = fmt.Sprintf("package %s; %s", p.Packagename, str)
		default:
			str = fmt.Sprintf("package %s; func %s() { %s }", p.Packagename, TemporaryFunctionName, str)
		}
		src = str
	}
	return parser.ParseFile(p.Fileset, p.Filename, src, p.Parsermode)
}

func (env *Env) Repl() {
	in := bufio.NewReader(os.Stdin)
	env.Repl1(in)
}

func (env *Env) Repl1(in *bufio.Reader) {
	for env.ReadEvalPrint(in) {
	}
}

func (env *Env) ReadEvalPrint(in *bufio.Reader) (ret bool) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Fprintln(Stderr, rec)
			ret = true
		}
	}()

	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Fprintln(Stderr, err)
		return false
	}
	line = strings.TrimSpace(line)
	if line == ":quit" {
		return false
	}
	ast := env.Parse(line)
	// env.FprintMultipleValues(Stdout, r.ValueOf(ast))
	value, values := env.Eval(ast)
	if len(values) != 0 {
		env.FprintMultipleValues(Stdout, values...)
	} else {
		env.FprintMultipleValues(Stdout, value)
	}
	return true
}

// func pair(a, b int) (int, int) { return a, b }
// var a, b, c = pair(1, 2), 3, 4

func main() {
	env := NewEnv(nil)
	// env.Parser.Parsermode = parser.Trace
	env.Repl()
}
