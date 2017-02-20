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
 * parser.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"

	mp "github.com/cosmos72/gomacro/macroparser"
)

type Interpreter struct {
	Packagename string
	Filename    string
	Fileset     *token.FileSet
	Parsermode  parser.Mode
	Stdout      io.Writer
	Stderr      io.Writer
}

func NewInterpreter() *Interpreter {
	p := Interpreter{}
	p.Packagename = "main"
	p.Filename = "main.go"
	p.Fileset = token.NewFileSet()
	// using both os.Stdout and os.Stderr can interleave impredictably
	// normal output and diagnostic messages - ugly in interactive use
	p.Stdout = os.Stdout
	p.Stderr = os.Stdout
	return &p
}

func (p *Interpreter) Parse(src interface{}) ast.Node {
	bytes := p.ReadFromSource(src)
	node, err := p.parseOrError(bytes)
	if err != nil {
		Errore(err)
		return nil
	}
	return node
}

func (p *Interpreter) parseOrError_OrigVersion(src []byte) (ast.Node, error) {
	pos := findFirstToken(src)
	src = src[pos:]
	expr, err := parser.ParseExprFrom(p.Fileset, p.Filename, src, 0)
	if err == nil {
		if p.Parsermode == 0 {
			return expr, nil
		}
		return parser.ParseExprFrom(p.Fileset, p.Filename, src, p.Parsermode)
	}
	firstIdent := string(extractFirstIdentifier(src))
	switch firstIdent {
	case "package":
		// nothing to do
	case "const", "func", "import", "type", "var":
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "package %s; ", p.Packagename)
		buf.Write(src)
		src = buf.Bytes()
	default:
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "package %s; func %s() { ", p.Packagename, temporaryFunctionName)
		buf.Write(src)
		buf.WriteString(" }")
		src = buf.Bytes()
	}
	return parser.ParseFile(p.Fileset, p.Filename, src, p.Parsermode)
}

func (ir *Interpreter) parseOrError(src []byte) (node ast.Node, err error) {

	return mp.Parse(ir.Fileset, ir.Filename, src, mp.Mode(ir.Parsermode))
}
