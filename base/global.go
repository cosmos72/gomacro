/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * global.go
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
	"runtime"
	"strings"

	. "github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base/genimport"
	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/base/reflect"
	etoken "github.com/cosmos72/gomacro/go/etoken"
	mp "github.com/cosmos72/gomacro/go/parser"
	"github.com/cosmos72/gomacro/imports"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type Output = output.Output

type CmdOpt uint32

const (
	CmdOptQuit      = 1 << iota
	CmdOptForceEval // temporarily re-enable evaluation even if in macroexpand-only mode
)

type Inspector interface {
	Inspect(name string, val r.Value, typ r.Type, xtyp xr.Type, globals *Globals)
}

type Globals struct {
	Output
	Options      Options
	PackagePath  string
	Filepath     string
	Importer     *genimport.Importer
	Imports      []*ast.GenDecl
	Declarations []ast.Decl
	Statements   []ast.Stmt
	Prompt       string
	Readline     Readline
	GensymN      uint
	ParserMode   mp.Mode
	MacroChar    rune // prefix for macro-related keywords macro, quote, quasiquote, splice... The default is '~'
	ReplCmdChar  byte // prefix for special REPL commands env, help, inspect, quit, unload... The default is ':'
	Inspector    Inspector
	GoVersion    string
	GoOS         string
	GoArch       string
}

func NewGlobals() *Globals {
	var options Options = OptTrapPanic // set by default
	if GoModuleSupported {
		options |= OptModuleImport
	}

	g := &Globals{
		Output: Output{
			Stringer: output.Stringer{
				Fileset:    etoken.NewFileSet(),
				NamedTypes: make(map[r.Type]string),
			},
			// using both os.Stdout and os.Stderr can interleave impredictably
			// normal output and diagnostic messages - ugly in interactive use
			Stdout: os.Stdout,
			Stderr: os.Stdout,
		},
		Options:      options,
		PackagePath:  "main",
		Filepath:     "repl.go",
		Imports:      nil,
		Declarations: nil,
		Statements:   nil,
		Prompt:       "gomacro [In]: ",
		GensymN:      0,
		ParserMode:   0,
		MacroChar:    '~',
		ReplCmdChar:  ':', // Jupyter and gophernotes would probably set this to '%'
		GoVersion:    runtime.Version(),
		GoOS:         runtime.GOOS,
		GoArch:       runtime.GOARCH,
	}
	g.Importer = genimport.DefaultImporter(&g.Output)
	return g

}

func (g *Globals) Gensym() string {
	n := g.GensymN
	g.GensymN++
	return fmt.Sprintf("%s%d", StrGensym, n)
}

func (g *Globals) GensymAnonymous(name string) string {
	if len(name) == 0 {
		n := g.GensymN
		g.GensymN++
		name = fmt.Sprintf("%d", n)
	}
	return StrGensymAnonymous + name
}

func (g *Globals) GensymPrivate(name string) string {
	if len(name) == 0 {
		n := g.GensymN
		g.GensymN++
		name = fmt.Sprintf("%d", n)
	}
	return StrGensymPrivate + name
}

func IsGensym(name string) bool {
	return strings.HasPrefix(name, StrGensym)
}

func IsGensymInterface(name string) bool {
	return name == StrGensymInterface
}

func IsGensymAnonymous(name string) bool {
	return strings.HasPrefix(name, StrGensymAnonymous)
}

func IsGensymPrivate(name string) bool {
	return strings.HasPrefix(name, StrGensymPrivate)
}

// read phase
// return read string and position of first non-comment token.
// return "", -1 on EOF
func (g *Globals) ReadMultiline(opts ReadOptions, prompt string) (str string, firstToken int) {
	str, firstToken, err := ReadMultiline(g.Readline, opts, prompt)
	if err != nil && err != io.EOF {
		fmt.Fprintf(g.Stderr, "// read error: %s\n", err)
	}
	return str, firstToken
}

// parse phase. no macroexpansion.
func (g *Globals) ParseBytes(src []byte) []ast.Node {
	var parser mp.Parser

	mode := g.ParserMode
	if g.Options&OptDebugParse != 0 {
		mode |= mp.Trace
	} else {
		mode &^= mp.Trace
	}
	if g.Options&OptDebugger != 0 {
		// to show source code in debugger
		mode |= mp.CopySources
	} else {
		mode &^= mp.CopySources
	}
	parser.Configure(mode, g.MacroChar)
	parser.Init(g.Fileset, g.Filepath, g.Line, src)

	nodes, err := parser.Parse()
	if err != nil {
		output.Error(err)
	}
	return nodes
}

// print phase
func (g *Globals) Print(values []r.Value, types []xr.Type) {
	opts := g.Options
	if opts&OptShowEval != 0 {
		if opts&OptShowEvalType != 0 {
			for i, vi := range values {
				var ti interface{}
				if types != nil && i < len(types) {
					ti = types[i]
				} else {
					ti = reflect.Type(vi)
				}
				g.Fprintf(g.Stdout, "       [Out]: %v\t// %v\n\n", vi, ti)
			}
		} else {
			for _, vi := range values {
				g.Fprintf(g.Stdout, "       [Out]: %v\n", vi)
			}
		}
	}
}

// remove package 'path' from the list of known packages.
// later attempts to import it again will trigger a recompile.
func (g *Globals) UnloadPackage(path string) {
	if n := len(path); n > 1 && path[0] == '"' && path[n-1] == '"' {
		path = path[1 : n-1] // remove quotes
	}
	slash := strings.IndexByte(path, '/')
	if _, found := imports.Packages[path]; !found {
		if slash < 0 {
			g.Debugf("nothing to unload: cannot find imported package %q. Remember to specify the full package path, not only its name", path)
		} else {
			g.Debugf("nothing to unload: cannot find imported package %q", path)
		}
	}
	delete(imports.Packages, path)
	dot := strings.IndexByte(path, '.')
	if slash < 0 || dot > slash {
		g.Warnf("unloaded standard library package %q. attempts to import it again will trigger a recompile", path)
		return
	}
	g.Debugf("unloaded package %q. attempts to import it again will trigger a recompile", path)
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
	default:
		g.Errorf("unable to collect AST type: %T", form)
	}
}

func (g *Globals) CollectNode(node ast.Node) {
	collectDecl := g.Options&OptCollectDeclarations != 0
	collectStmt := g.Options&OptCollectStatements != 0

	switch node := node.(type) {
	case *ast.GenDecl:
		if collectDecl {
			switch node.Tok {
			case token.IMPORT:
				g.Imports = append(g.Imports, node)
			case token.PACKAGE:
				/*
					modified parser converts 'package foo' to:

					ast.GenDecl{
						Tok: token.PACKAGE,
						Specs: []ast.Spec{
							&ast.ValueSpec{
								Names: []ast.Ident{
									&ast.Ident{
										Name:  "foo",
									},
								},
							},
						},
					}
				*/
				if len(node.Specs) == 1 {
					if decl, ok := node.Specs[0].(*ast.ValueSpec); ok {
						if len(decl.Names) == 1 {
							g.PackagePath = decl.Names[0].Name
						}
					}
				}
			case token.TYPE, token.VAR, token.CONST:
				g.Declarations = append(g.Declarations, node)
			default:
				g.Errorf("unable to collect AST declaration: %s", node.Tok)
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
	case ast.Spec:
		decl := &ast.GenDecl{
			TokPos: node.Pos(),
			Specs:  []ast.Spec{node},
		}
		switch node.(type) {
		case *ast.ImportSpec:
			decl.Tok = token.IMPORT
		case *ast.TypeSpec:
			decl.Tok = token.TYPE
		case *ast.ValueSpec:
			decl.Tok = token.VAR
		default:
			g.Errorf("unable to collect AST spec type: %T", node)
		}
		g.CollectNode(decl)
		return
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
					g.PackagePath = ident.Name
					break
				}
			}
		}
		if collectStmt {
			stmt := &ast.ExprStmt{X: node}
			g.Statements = append(g.Statements, stmt)
		}
	default:
		g.Errorf("unable to collect AST node type: %T", node)
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
	g.Output.WriteDeclsToStream(out, g.PackagePath, g.Imports, g.Declarations, g.Statements)
}
