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
package parser

import (
	"go/ast"
	"go/token"

	mt "github.com/cosmos72/gomacro/token"
)

// parse prefix#[T1,T2...] as &ast.IndexExpr{ &ast.CompositeLit{Type: Foo, Elts: [T1, T2...]} }
func (p *parser) parseHash(prefix ast.Expr) ast.Expr {
	p.next()
	lbrack := p.expect(token.LBRACK)
	list := p.parseTypeList()
	rbrack := p.expect(token.RBRACK)
	return &ast.IndexExpr{
		X:      prefix,
		Lbrack: lbrack,
		Index:  &ast.CompositeLit{Type: nil, Lbrace: lbrack, Elts: list, Rbrace: rbrack},
		Rbrack: rbrack,
	}
}

// parse template[T1,T2...] type ...
// and template[T1,T2...] func ...
func (p *parser) parseTemplateDecl(sync func(*parser)) ast.Decl {
	if p.trace {
		defer un(trace(p, "TemplateDecl"))
	}
	var lbrack, rbrack token.Pos
	var templateTypes []ast.Expr

	p.expect(mt.TEMPLATE)
	lbrack = p.expect(token.LBRACK)

	bad := func() ast.Decl {
		pos := p.expect(token.RBRACK)
		sync(p)
		return &ast.BadDecl{From: pos, To: p.pos}
	}
loop:
	for {
		tok := p.tok
		switch tok {
		case token.RBRACK:
			rbrack = p.pos
			p.next()
			break loop
		case token.ILLEGAL, token.EOF, token.RPAREN, token.RBRACE:
			return bad()
		}

		templateTypes = append(templateTypes, p.parseType())

		tok = p.tok
		if tok == token.RBRACK {
			continue
		} else if tok == token.COMMA {
			p.next()
		} else {
			return bad()
		}
	}
	switch tok := p.tok; tok {
	case token.TYPE:
		decl := p.parseGenDecl(tok, p.parseTypeSpec)
		return templateTypeDecl(lbrack, templateTypes, rbrack, decl)

	case token.FUNC, mt.FUNCTION:
		decl := p.parseFuncDecl(tok)
		return templateFuncDecl(lbrack, templateTypes, rbrack, decl)

	default:
		pos := p.pos
		p.errorExpected(pos, "type or func")
		sync(p)
		return &ast.BadDecl{From: pos, To: p.pos}
	}
}

func templateTypeDecl(lbrack token.Pos, templateTypes []ast.Expr, rbrack token.Pos, decl *ast.GenDecl) *ast.GenDecl {
	for _, spec := range decl.Specs {
		if typespec, ok := spec.(*ast.TypeSpec); ok {
			// hack: store template types in *ast.CompositeLit.
			// it is never used inside *ast.TypeSpec and has exacly the required fields
			typespec.Type = &ast.CompositeLit{
				Type:   typespec.Type,
				Lbrace: lbrack,
				Elts:   templateTypes,
				Rbrace: rbrack,
			}
		}
	}
	return decl
}

func templateFuncDecl(lbrack token.Pos, templateTypes []ast.Expr, rbrack token.Pos, decl *ast.FuncDecl) *ast.FuncDecl {
	// hack: store template types as second function receiver.
	// it's never used for functions and macros.
	recv := decl.Recv
	if recv == nil {
		recv = &ast.FieldList{Opening: lbrack, Closing: rbrack}
		decl.Recv = recv
	}
	list := []*ast.Field{
		nil,
		// add template types as second receiver
		&ast.Field{Type: &ast.CompositeLit{Elts: templateTypes}},
	}
	if len(recv.List) != 0 {
		list[0] = recv.List[0]
	}
	recv.List = list
	return decl
}
