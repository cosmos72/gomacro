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
 * interpreter.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	mp "github.com/cosmos72/gomacro/parser"
)

type Interpreter struct {
	Output
	Packagename string
	Filename    string
	Options     Options
	Importer    Importer
	ParserMode  parser.Mode
	ParserScope *ast.Scope
	SpecialChar rune
}

func NewInterpreter() *Interpreter {
	ir := Interpreter{}
	ir.Packagename = "main"
	ir.Filename = "main.go"
	ir.Fileset = token.NewFileSet()
	ir.SpecialChar = '~'
	// using both os.Stdout and os.Stderr can interleave impredictably
	// normal output and diagnostic messages - ugly in interactive use
	ir.Stdout = os.Stdout
	ir.Stderr = os.Stdout
	ir.Importer = DefaultImporter()
	return &ir
}

func (ir *Interpreter) Parse(src interface{}) []ast.Node {
	bytes := ir.ReadFromSource(src)
	return ir.ParseBytes(bytes)
}

func (ir *Interpreter) Parse1(src interface{}) ast.Node {
	ret := ir.Parse(src)
	switch len(ret) {
	default:
		ir.Warnf("Interpreter.Parse() returned %d values, only the first one will be used", len(ret))
		fallthrough
	case 1:
		return ret[0]
	case 0:
		return nil
	}
}

func (ir *Interpreter) ParseBytes(src []byte) []ast.Node {
	var parser mp.Parser

	parser.Fileset = ir.Fileset
	parser.Mode = mp.Mode(ir.ParserMode)
	parser.SpecialChar = ir.SpecialChar

	ir.ParserScope = parser.Init(ir.Filename, src, ir.ParserScope)

	nodes, err := parser.Parse()
	if err != nil {
		Error(err)
		return nil
	}
	return nodes
}

//
//
// no longer used:
//
//

func (ir *Interpreter) ParseBytes_OrigVersion(src []byte) []ast.Node {
	node := ir.ParseBytes1_OrigVersion(src)
	return []ast.Node{node}
}

func (ir *Interpreter) ParseBytes1_OrigVersion(src []byte) ast.Node {
	pos := findFirstToken(src)
	src = src[pos:]
	expr, err := parser.ParseExprFrom(ir.Fileset, ir.Filename, src, 0)
	if err == nil {
		if ir.ParserMode != 0 {
			// run again with user-specified ParserMode
			expr, err = parser.ParseExprFrom(ir.Fileset, ir.Filename, src, ir.ParserMode)
			if err != nil {
				Error(err)
				return nil
			}
		}
		return expr
	}
	firstIdent := string(extractFirstIdentifier(src))
	switch firstIdent {
	case "package":
		// nothing to do
	case "const", "func", "import", "type", "var":
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "package %s; ", ir.Packagename)
		buf.Write(src)
		src = buf.Bytes()
	default:
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "package %s; func %s() { ", ir.Packagename, temporaryFunctionName)
		buf.Write(src)
		buf.WriteString(" }")
		src = buf.Bytes()
	}
	node, err := parser.ParseFile(ir.Fileset, ir.Filename, src, ir.ParserMode)
	if err != nil {
		Error(err)
		return nil
	}
	return node
}
