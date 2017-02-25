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

package interpreter

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"

	mp "github.com/cosmos72/gomacro/parser"
)

type FileSet struct {
	Fileset *token.FileSet
}

type Interpreter struct {
	FileSet
	Packagename string
	Filename    string
	Options     Options
	ParserMode  parser.Mode
	ParserScope *ast.Scope
	Stdout      io.Writer
	Stderr      io.Writer
}

func NewInterpreter() *Interpreter {
	ir := Interpreter{}
	ir.Packagename = "main"
	ir.Filename = "main.go"
	ir.Fileset = token.NewFileSet()
	// using both os.Stdout and os.Stderr can interleave impredictably
	// normal output and diagnostic messages - ugly in interactive use
	ir.Stdout = os.Stdout
	ir.Stderr = os.Stdout
	return &ir
}

func (ir *Interpreter) ParseN(src interface{}) []ast.Node {
	bytes := ir.ReadFromSource(src)
	nodes, err := ir.parseOrError(bytes)
	if err != nil {
		Error(err)
		return nil
	}
	return nodes
}

func (ir *Interpreter) parseOrError(src []byte) ([]ast.Node, error) {
	var parser mp.Parser
	ir.ParserScope = parser.Init(ir.Fileset, ir.Filename, src, mp.Mode(ir.ParserMode), ir.ParserScope)

	return parser.Parse()
}

//
//
// no longer used:
//
//

func (ir *Interpreter) parseOrError_OrigVersion(src []byte) ([]ast.Node, error) {
	node, err := ir.parseOrError1_OrigVersion(src)
	return []ast.Node{node}, err
}

func (ir *Interpreter) parseOrError1_OrigVersion(src []byte) (ast.Node, error) {
	pos := findFirstToken(src)
	src = src[pos:]
	expr, err := parser.ParseExprFrom(ir.Fileset, ir.Filename, src, 0)
	if err == nil {
		if ir.ParserMode == 0 {
			return expr, nil
		}
		return parser.ParseExprFrom(ir.Fileset, ir.Filename, src, ir.ParserMode)
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
	return parser.ParseFile(ir.Fileset, ir.Filename, src, ir.ParserMode)
}
