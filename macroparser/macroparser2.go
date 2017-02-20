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

func (p *MacroParser) parseAny() ast.Node {
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
