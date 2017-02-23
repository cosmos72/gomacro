// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parser implements a parser for Go source files. Input may be
// provided in a variety of forms (see the various Parse* functions); the
// output is an abstract syntax tree (AST) representing the Go source. The
// parser is invoked through one of the Parse* functions.
//
// The parser accepts a larger language than is syntactically permitted by
// the Go spec, for simplicity, and for improved robustness in the presence
// of syntax errors. For instance, in method declarations, the receiver is
// treated like an ordinary parameter list and thus may contain multiple
// entries where the spec permits exactly one. Consequently, the corresponding
// field in the AST (ast.FuncDecl.Recv) field is not restricted to one entry.
//
package macroparser

import (
	"go/ast"
	"go/token"
)

func (p *Parser) parseAny() ast.Node {
	var node ast.Node

	if p.tok == token.COMMENT {
		// advance to the next non-comment token
		p.next()
	}
	switch p.tok {
	case token.PACKAGE:
		node = p.parseFile()
	case token.IMPORT:
		node = p.parseGenDecl(token.IMPORT, p.parseImportSpec)
	case token.CONST, token.FUNC, token.TYPE, token.VAR, MACRO:
		node = p.parseDecl(syncDecl)
	default:
		node = p.parseStmt()
		if expr, ok := node.(*ast.ExprStmt); ok {
			// unwrap expressions
			node = expr.X
		}
	}
	return node
}

// patch: quote and friends
func (p *Parser) parseQuote() ast.Expr {
	if p.trace {
		defer un(trace(p, "Quote"))
	}

	tok := p.tok
	pos := p.pos
	name := p.lit
	p.next()
	switch p.tok {
	case token.EOF, token.RPAREN, token.RBRACK, token.RBRACE,
		token.COMMA, token.PERIOD, token.SEMICOLON, token.COLON:

		// no applicable expression after QUOTE: just return the QUOTE identifier itself
		return &ast.Ident{NamePos: pos, Name: name}
	}

	body := p.parseBlockStmt()

	// due to go/ast strictly typed model, there is only one mechanism
	// to insert a statement inside an expression: use a closure.
	// so we return a unary expression: QUOTE (func() { /*block*/ })
	typ := &ast.FuncType{Func: token.NoPos, Params: &ast.FieldList{}}
	fun := &ast.FuncLit{Type: typ, Body: body}
	return &ast.UnaryExpr{OpPos: pos, Op: tok, X: fun}
}

func isMacroDecl(decl *ast.FuncDecl) bool {
	return decl != nil && decl.Recv != nil && decl.Recv.List != nil && len(decl.Recv.List) == 0
}

func funcDeclNumParams(decl *ast.FuncDecl) int {
	ret := 0
	if params := decl.Type.Params; params != nil {
		ret = params.NumFields()
	}
	return ret
}

func (p *Parser) tryParseMacroStmt() ast.Stmt {
	if p.trace {
		defer un(trace(p, "tryMacroStmt"))
	}
	if expr := p.tryParseMacroExpr(); expr != nil {
		return &ast.ExprStmt{X: expr}
	}
	return nil
}

// if current token is an identifier currently defined as a macro,
// retrieve the number of arguments it expects and parse it accordingly
func (p *Parser) tryParseMacroExpr() ast.Expr {
	if p.trace {
		defer un(trace(p, "tryMacroExpr"))
	}
	pos := p.pos
	if p.tok != token.IDENT {
		p.errorExpected(pos, "'"+token.IDENT.String()+"'")
		return nil
	}
	name := p.lit
	ident := &ast.Ident{NamePos: pos, Name: name}

	p.tryResolve(ident, false)
	if ident.Obj == nil || ident.Obj.Decl == nil {
		return nil
	}
	switch decl := ident.Obj.Decl.(type) {
	case *ast.FuncDecl:
		if isMacroDecl(decl) {
			n := funcDeclNumParams(decl)
			return p.parseMacro(ident, n)
		}
	}
	return nil
}

func (p *Parser) parseMacro(ident *ast.Ident, numParams int) ast.Expr {
	if p.trace {
		defer un(trace(p, "Macro"))
	}
	p.expect(token.IDENT)

	// we could execute the macro here - but this is a parser, not an interpreter.
	// Also, before executing it, we would also need to parse its arguments
	// WHILE DISABLING macro execution... and after the macro executed,
	// walk through its output to find and execute all macro calls found.
	// Considering all of these points, the task is definitely beyond MacroParser objective,
	// so we just insert the macro call into the generated AST

	lbrace := p.pos
	list := make([]ast.Stmt, numParams)
	for i := 0; i < numParams; i++ {
		list[i] = p.parseStmt() // rely on parseStmt() error and EOF detection
	}
	rbrace := p.pos

	body := &ast.BlockStmt{Lbrace: lbrace, List: list, Rbrace: rbrace}

	// due to go/ast strictly typed model, there is only one mechanism
	// to insert a statement inside an expression: use a closure.
	// so we return a binary expression: ident MACRO (func() { /*block*/ })
	typ := &ast.FuncType{Func: token.NoPos, Params: &ast.FieldList{}}
	fun := &ast.FuncLit{Type: typ, Body: body}
	return &ast.BinaryExpr{X: ident, OpPos: ident.Pos(), Op: MACRO, Y: fun}
}
