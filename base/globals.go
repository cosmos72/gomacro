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
 * globals.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"os"
	r "reflect"
	"strings"

	mp "github.com/cosmos72/gomacro/parser"

	. "github.com/cosmos72/gomacro/ast2"
)

type Globals struct {
	Output
	Options      Options
	Packagename  string
	Filename     string
	GensymN      uint
	Importer     Importer
	Imports      []*ast.GenDecl
	Declarations []ast.Decl
	Statements   []ast.Stmt
	ParserMode   mp.Mode
	SpecialChar  rune
}

func (g *Globals) Init() {
	g.Output = Output{
		Stringer: Stringer{
			Fileset:    token.NewFileSet(),
			NamedTypes: make(map[r.Type]string),
		},
		// using both os.Stdout and os.Stderr can interleave impredictably
		// normal output and diagnostic messages - ugly in interactive use
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	g.Options = OptTrapPanic // set by default
	g.Packagename = "main"
	g.Filename = "repl.go"
	g.GensymN = 0
	g.Importer = DefaultImporter()
	g.Imports = nil
	g.Declarations = nil
	g.Statements = nil
	g.ParserMode = 0
	g.SpecialChar = '~'
}

func IsReflectGensym(name string) bool {
	return strings.HasPrefix(name, ReflectGensymPrefix)
}

func IsGensym(name string) bool {
	return strings.HasPrefix(name, GensymPrefix)
}

func (g *Globals) Gensym() string {
	n := g.GensymN
	g.GensymN++
	return fmt.Sprintf("%s%d", GensymPrefix, n)
}

func (g *Globals) ParseBytes(src []byte) []ast.Node {
	var parser mp.Parser

	parser.Fileset = g.Fileset
	parser.Mode = mp.Mode(g.ParserMode)
	parser.SpecialChar = g.SpecialChar

	parser.Init(g.Filename, src)

	nodes, err := parser.Parse()
	if err != nil {
		Error(err)
		return nil
	}
	return nodes
}

// CollectAst accumulates declarations in ir.Decls and statements in ir.Stmts
// allows generating a *.go file on user request
func (g *Globals) CollectAst(form Ast) {
	if g.Options&(OptCollectDeclarations|OptCollectStatements) == 0 {
		return
	}

	switch form := form.(type) {
	case AstWithNode:
		g.CollectNode(form.Node())
	case AstWithSlice:
		n := form.Size()
		for i := 0; i < n; i++ {
			g.CollectAst(form.Get(i))
		}
	}
}

func (g *Globals) CollectNode(node ast.Node) {
	collectDecl := g.Options&OptCollectDeclarations != 0
	collectStmt := g.Options&OptCollectStatements != 0

	switch node := node.(type) {
	case *ast.GenDecl:
		if collectDecl {
			if node.Tok == token.IMPORT {
				g.Imports = append(g.Imports, node)
			} else {
				g.Declarations = append(g.Declarations, node)
			}
		}
	case *ast.FuncDecl:
		if collectDecl {
			if node.Recv == nil || len(node.Recv.List) != 0 {
				// function or method declaration.
				// skip macro declarations, Go compilers would choke on them
				g.Declarations = append(g.Declarations, node)
			}
		}
	case ast.Decl:
		if collectDecl {
			g.Declarations = append(g.Declarations, node)
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
				g.Declarations = append(g.Declarations, decl)
			}
		} else {
			if collectStmt {
				g.Statements = append(g.Statements, node)
			}
		}
	case ast.Stmt:
		if collectStmt {
			g.Statements = append(g.Statements, node)
		}
	case ast.Expr:
		if unary, ok := node.(*ast.UnaryExpr); ok && collectDecl {
			if unary.Op == token.PACKAGE && unary.X != nil {
				if ident, ok := unary.X.(*ast.Ident); ok {
					g.Packagename = ident.Name
					break
				}
			}
		}
		if collectStmt {
			stmt := &ast.ExprStmt{X: node}
			g.Statements = append(g.Statements, stmt)
		}
	}
}

func (g *Globals) WriteDeclsToFile(filename string, prologue ...string) {
	f, err := os.Create(filename)
	if err != nil {
		g.Errorf("failed to create file %q: %v", filename, err)
	}
	defer f.Close()
	for _, str := range prologue {
		f.WriteString(str)
	}
	g.WriteDeclsToStream(f)
}

func (g *Globals) WriteDeclsToStream(out io.Writer) {
	fmt.Fprintf(out, "package %s\n\n", g.Packagename)

	for _, imp := range g.Imports {
		fmt.Fprintln(out, g.toPrintable(imp))
	}
	if len(g.Imports) != 0 {
		fmt.Fprintln(out)
	}
	for _, decl := range g.Declarations {
		fmt.Fprintln(out, g.toPrintable(decl))
	}
	if len(g.Statements) != 0 {
		fmt.Fprint(out, "\nfunc init() {\n")
		config.Indent = 1
		defer func() {
			config.Indent = 0
		}()
		for _, stmt := range g.Statements {
			fmt.Fprintln(out, g.toPrintable(stmt))
		}
		fmt.Fprint(out, "}\n")
	}
}
