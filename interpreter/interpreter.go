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
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"os"

	. "github.com/cosmos72/gomacro/ast2"
	mp "github.com/cosmos72/gomacro/parser"
)

type InterpreterCommon struct {
	output
	Options     Options
	Importer    Importer
	Packagename string
	Filename    string
	Decls       []ast.Decl
	Stmts       []ast.Stmt
	ParserMode  mp.Mode
	SpecialChar rune
}

func NewInterpreterCommon() *InterpreterCommon {
	return &InterpreterCommon{
		output: output{
			fileSet: fileSet{token.NewFileSet()},
			// using both os.Stdout and os.Stderr can interleave impredictably
			// normal output and diagnostic messages - ugly in interactive use
			Stdout: os.Stdout,
			Stderr: os.Stdout,
		},
		Options:     OptTrapPanic, // set by default
		Importer:    DefaultImporter(),
		Packagename: "main",
		Filename:    "main.go",
		SpecialChar: '~',
	}
}

func (ir *InterpreterCommon) ParseAst(src interface{}) Ast {
	// parse phase
	bytes := ReadBytes(src)
	nodes := ir.ParseBytes(bytes)

	if env.Options&OptShowAfterParse != 0 {
		env.debugf("after parse: %v", ast.Interface())
	}

	var form Ast
	switch len(nodes) {
	case 0:
		return nil
	case 1:
		form = ToAst(nodes[0])
	default:
		form = NodeSlice{X: nodes}
	}

	// macroexpansion phase.
	form, _ = env.MacroExpandAstCodewalk(form)

	if env.Options&OptShowAfterMacroExpansion != 0 {
		env.debugf("after macroexpansion: %v", form.Interface())
	}
	if env.Options&(OptCollectDeclarations|OptCollectStatements) != 0 {
		env.collectAst(form)
	}
	return form
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

// collectAst accumulates declarations in ir.Decls and statements in ir.Stmts
// allows generating a *.go file on user request
func (ir *InterpreterCommon) collectAst(form Ast) {
	if ir.Options&(OptCollectDeclarations|OptCollectStatements) == 0 {
		return
	}

	switch form := form.(type) {
	case AstWithNode:
		ir.collectNode(form.Node())
	case AstWithSlice:
		n := form.Size()
		for i := 0; i < n; i++ {
			ir.collectAst(form.Get(i))
		}
	}
}

func (ir *InterpreterCommon) collectNode(node ast.Node) {
	collectDecl := ir.Options&OptCollectDeclarations != 0
	collectStmt := ir.Options&OptCollectStatements != 0

	switch node := node.(type) {
	case *ast.FuncDecl:
		if collectDecl {
			if node.Recv == nil || len(node.Recv.List) != 0 {
				// function or method declaration.
				// skip macro declarations, Go compilers would choke on them
				ir.Decls = append(ir.Decls, node)
			}
		}
	case ast.Decl:
		if collectDecl {
			ir.Decls = append(ir.Decls, node)
		}
	case *ast.AssignStmt:
		if node.Tok == token.DEFINE {
			if collectDecl {
				idents := make([]*ast.Ident, len(node.Lhs))
				for i, lhs := range node.Lhs {
					idents[i] = lhs.(*ast.Ident)
				}
				decl := &ast.GenDecl{
					TokPos: node.Pos(),
					Tok:    token.VAR,
					Specs: []ast.Spec{
						&ast.ValueSpec{
							Names:  idents,
							Type:   nil,
							Values: node.Rhs,
						},
					},
				}
				ir.Decls = append(ir.Decls, decl)
			}
		} else {
			if collectStmt {
				ir.Stmts = append(ir.Stmts, node)
			}
		}
	case ast.Stmt:
		if collectStmt {
			ir.Stmts = append(ir.Stmts, node)
		}
	case ast.Expr:
		if collectStmt {
			stmt := &ast.ExprStmt{X: node}
			ir.Stmts = append(ir.Stmts, stmt)
		}
	}
}

func (ir *InterpreterCommon) writeDecls(out io.Writer, filename string) {
	if out == nil {
		f, err := os.Create(filename)
		if err != nil {
			ir.errorf("failed to create file %q: %v", filename, err)
		}
		out = f
	}
	for _, decl := range ir.Decls {
		fmt.Fprintln(out, ir.toPrintable(decl))
	}
	if len(ir.Stmts) == 0 {
		return
	}
	fmt.Fprint(out, "\nfunc init() {\n")
	config.Indent = 1
	defer func() {
		config.Indent = 0
	}()
	for _, stmt := range ir.Stmts {
		fmt.Fprintln(out, ir.toPrintable(stmt))
	}
	fmt.Fprint(out, "}\n")
}
