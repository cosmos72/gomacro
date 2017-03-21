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

	. "github.com/cosmos72/gomacro/ast2"
	mp "github.com/cosmos72/gomacro/parser"
)

type InterpreterCommon struct {
	output
	Packagename string
	Filename    string
	Options     Options
	Importer    Importer
	ParserMode  mp.Mode
	SpecialChar rune
}

func NewInterpreterCommon() *InterpreterCommon {
	ir := InterpreterCommon{}
	ir.Packagename = "main"
	ir.Filename = "main.go"
	ir.Fileset = token.NewFileSet()
	ir.Options = OptTrapPanic // set by default
	ir.SpecialChar = '~'
	// using both os.Stdout and os.Stderr can interleave impredictably
	// normal output and diagnostic messages - ugly in interactive use
	ir.Stdout = os.Stdout
	ir.Stderr = os.Stdout
	ir.Importer = DefaultImporter()
	return &ir
}

func (ir *InterpreterCommon) ParseAst(src interface{}) Ast {
	bytes := ReadBytes(src)
	nodes := ir.ParseBytes(bytes)
	switch len(nodes) {
	case 0:
		return nil
	case 1:
		return ToAst(nodes[0])
	default:
		return NodeSlice{X: nodes}
	}
}

func (ir *InterpreterCommon) ParseBytes(src []byte) []ast.Node {
	var parser mp.Parser

	parser.Fileset = ir.Fileset
	parser.Mode = mp.Mode(ir.ParserMode)
	parser.SpecialChar = ir.SpecialChar

	parser.Init(ir.Filename, src)

	nodes, err := parser.Parse()
	if err != nil {
		error_(err)
		return nil
	}
	return nodes
}

//
//
// no longer used:
//
//

func (ir *InterpreterCommon) ParseBytes_OrigVersion(src []byte) []ast.Node {
	node := ir.ParseBytes1_OrigVersion(src)
	return []ast.Node{node}
}

func (ir *InterpreterCommon) ParseBytes1_OrigVersion(src []byte) ast.Node {
	pos := findFirstToken(src)
	src = src[pos:]
	expr, err := parser.ParseExprFrom(ir.Fileset, ir.Filename, src, 0)
	if err == nil {
		if ir.ParserMode != 0 {
			// run again with user-specified ParserMode
			expr, err = parser.ParseExprFrom(ir.Fileset, ir.Filename, src, parser.Mode(ir.ParserMode))
			if err != nil {
				error_(err)
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
	node, err := parser.ParseFile(ir.Fileset, ir.Filename, src, parser.Mode(ir.ParserMode))
	if err != nil {
		error_(err)
		return nil
	}
	return node
}
