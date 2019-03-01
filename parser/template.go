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

// set to false to disable parsing gomacro generics, version 1
const GENERICS_V1 = true

// parse prefix#[T1,T2...] as &ast.IndexExpr{ &ast.CompositeLit{Type: prefix, Elts: [T1, T2...]} }
func (p *parser) parseHash(prefix ast.Expr) ast.Expr {
	if p.trace {
		defer un(trace(p, "Hash"))
	}
	p.expect(mt.HASH)
	params := p.parseTemplateParams()
	return &ast.IndexExpr{
		X:      prefix,
		Lbrack: params.Lbrace,
		Index:  params,
		Rbrack: params.Rbrace,
	}
}

// parse template[T1,T2...] type ...
// and template[T1,T2...] func ...
func (p *parser) parseTemplateDecl(sync func(*parser)) ast.Decl {
	if p.trace {
		defer un(trace(p, "TemplateDecl"))
	}
	p.expect(mt.TEMPLATE)
	params := p.parseTemplateParams()

	var specialize *ast.CompositeLit
	if p.tok == token.FOR {
		p.next()
		specialize = p.parseTemplateParams()
		params.Elts = append(params.Elts, &ast.BadExpr{}, specialize)
	}
	switch tok := p.tok; tok {
	case token.TYPE:
		decl := p.parseGenDecl(tok, p.parseTypeSpec)
		return templateTypeDecl(params, decl)

	case token.FUNC, mt.FUNCTION:
		decl := p.parseFuncDecl(tok)
		return templateFuncDecl(params, decl)

	default:
		pos := p.pos
		if specialize == nil {
			p.errorExpected(pos, "'type', 'func' or 'for' after 'template[...]'")
		} else {
			p.errorExpected(pos, "'type' or 'func' after 'template[...] for[...]'")
		}
		sync(p)
		return &ast.BadDecl{From: pos, To: p.pos}
	}
}

// parse [T1,T2...] in a template declaration
func (p *parser) parseTemplateParams() *ast.CompositeLit {
	var list []ast.Expr

	lbrack := p.expect(token.LBRACK)
	if p.tok != token.RBRACK {
		list = append(list, p.parseRhsOrType())
		for p.tok == token.COMMA {
			p.next()
			list = append(list, p.parseRhsOrType())
		}
	}
	rbrack := p.expect(token.RBRACK)

	return &ast.CompositeLit{
		Lbrace: lbrack,
		Elts:   list,
		Rbrace: rbrack,
	}
}

func templateTypeDecl(params *ast.CompositeLit, decl *ast.GenDecl) *ast.GenDecl {
	for _, spec := range decl.Specs {
		if typespec, ok := spec.(*ast.TypeSpec); ok {
			// hack: store template params in *ast.CompositeLit.
			// it is never used inside *ast.TypeSpec and has exacly the required fields
			typespec.Type = &ast.CompositeLit{
				Type:   typespec.Type,
				Lbrace: params.Lbrace,
				Elts:   params.Elts,
				Rbrace: params.Rbrace,
			}
		}
	}
	return decl
}

func templateFuncDecl(params *ast.CompositeLit, decl *ast.FuncDecl) *ast.FuncDecl {
	// hack: store template types as second function receiver.
	// it's never used for functions and macros.
	recv := decl.Recv
	if recv == nil {
		recv = &ast.FieldList{Opening: params.Lbrace, Closing: params.Rbrace}
		decl.Recv = recv
	}
	list := []*ast.Field{
		nil,
		// add template types as second receiver
		&ast.Field{Type: params},
	}
	if len(recv.List) != 0 {
		list[0] = recv.List[0]
	}
	recv.List = list
	return decl
}
