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
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
	mc "github.com/cosmos72/gomacro/compiler"
	mp "github.com/cosmos72/gomacro/parser"
)

type InterpreterCommon struct {
	output
	Options      Options
	AllMethods   map[r.Type]Methods // methods implemented by interpreted code
	Importer     Importer
	Packagename  string
	Filename     string
	Imports      []*ast.GenDecl
	Declarations []ast.Decl
	Statements   []ast.Stmt
	ParserMode   mp.Mode
	SpecialChar  rune
	CompEnv      *mc.CompEnv // temporary...
}

func NewInterpreterCommon() *InterpreterCommon {
	return &InterpreterCommon{
		output: output{
			stringer: stringer{
				Fileset:    token.NewFileSet(),
				NamedTypes: make(map[r.Type]string),
			},
			// using both os.Stdout and os.Stderr can interleave impredictably
			// normal output and diagnostic messages - ugly in interactive use
			Stdout: os.Stdout,
			Stderr: os.Stdout,
		},
		Options:     OptTrapPanic, // set by default
		AllMethods:  make(map[r.Type]Methods),
		Importer:    DefaultImporter(),
		Packagename: "main",
		Filename:    "main.go",
		SpecialChar: '~',
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
	case *ast.GenDecl:
		if collectDecl {
			if node.Tok == token.IMPORT {
				ir.Imports = append(ir.Imports, node)
			} else {
				ir.Declarations = append(ir.Declarations, node)
			}
		}
	case *ast.FuncDecl:
		if collectDecl {
			if node.Recv == nil || len(node.Recv.List) != 0 {
				// function or method declaration.
				// skip macro declarations, Go compilers would choke on them
				ir.Declarations = append(ir.Declarations, node)
			}
		}
	case ast.Decl:
		if collectDecl {
			ir.Declarations = append(ir.Declarations, node)
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
				ir.Declarations = append(ir.Declarations, decl)
			}
		} else {
			if collectStmt {
				ir.Statements = append(ir.Statements, node)
			}
		}
	case ast.Stmt:
		if collectStmt {
			ir.Statements = append(ir.Statements, node)
		}
	case ast.Expr:
		if unary, ok := node.(*ast.UnaryExpr); ok && collectDecl {
			if unary.Op == token.PACKAGE && unary.X != nil {
				if ident, ok := unary.X.(*ast.Ident); ok {
					ir.Packagename = ident.Name
					break
				}
			}
		}
		if collectStmt {
			stmt := &ast.ExprStmt{X: node}
			ir.Statements = append(ir.Statements, stmt)
		}
	}
}

func (ir *InterpreterCommon) writeDeclsToFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		ir.Errorf("failed to create file %q: %v", filename, err)
	}
	defer f.Close()
	ir.writeDeclsToStream(f)
}

func (ir *InterpreterCommon) writeDeclsToStream(out io.Writer) {
	fmt.Fprintf(out, "package %s\n\n", ir.Packagename)

	for _, imp := range ir.Imports {
		fmt.Fprintln(out, ir.toPrintable(imp))
	}
	if len(ir.Imports) != 0 {
		fmt.Println()
	}
	for _, decl := range ir.Declarations {
		fmt.Fprintln(out, ir.toPrintable(decl))
	}
	if len(ir.Statements) != 0 {
		fmt.Fprint(out, "\nfunc init() {\n")
		config.Indent = 1
		defer func() {
			config.Indent = 0
		}()
		for _, stmt := range ir.Statements {
			fmt.Fprintln(out, ir.toPrintable(stmt))
		}
		fmt.Fprint(out, "}\n")
	}
}
