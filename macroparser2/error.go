// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package MacroParser implements a MacroParser for Go source files. Input may be
// provided in a variety of forms (see the various Parse* functions); the
// output is an abstract syntax tree (AST) representing the Go source. The
// MacroParser is invoked through one of the Parse* functions.
//
// The MacroParser accepts a larger language than is syntactically permitted by
// the Go spec, for simplicity, and for improved robustness in the presence
// of syntax errors. For instance, in method declarations, the receiver is
// treated like an ordinary parameter list and thus may contain multiple
// entries where the spec permits exactly one. Consequently, the corresponding
// field in the AST (ast.FuncDecl.Recv) field is not restricted to one entry.
//
package macroparser2

import (
	"go/token"
)

// A bailout panic is raised to indicate early termination.
type Bailout struct{}

func (p *MacroParser) badExpr(pos token.Pos, end token.Pos) Slice {
	return Slice{BAD_s}.set(p, Attrs{POS: pos, END: end})
}

func (p *MacroParser) error(pos token.Pos, msg string) {
	epos := p.file.Position(pos)

	// If AllErrors is not set, discard errors reported on the same line
	// as the last recorded error and stop parsing if there are more than
	// 10 errors.
	if p.mode&AllErrors == 0 {
		n := len(p.errors)
		if n > 0 && p.errors[n-1].Pos.Line == epos.Line {
			return // discard - likely a spurious error
		}
		if n > 10 {
			panic(Bailout{})
		}
	}

	p.errors.Add(epos, msg)
}

func (p *MacroParser) errorExpected(pos token.Pos, msg string) {
	msg = "expected " + msg
	if pos == p.pos {
		// the error happened at the current position;
		// make the error message more specific
		if p.tok == token.SEMICOLON && p.lit == "\n" {
			msg += ", found newline"
		} else {
			msg += ", found '" + TokenString(p.tok) + "'"
			if p.tok.IsLiteral() {
				msg += " " + p.lit
			}
		}
	}
	p.error(pos, msg)
}

func (p *MacroParser) expect(tok token.Token) token.Pos {
	pos := p.pos
	if p.tok != tok {
		p.errorExpected(pos, "'"+TokenString(tok)+"'")
	}
	p.next() // make progress
	return pos
}

// expectClosing is like expect but provides a better error message
// for the common case of a missing comma before a newline.
//
func (p *MacroParser) expectClosing(tok token.Token, context string) token.Pos {
	if p.tok != tok && p.tok == token.SEMICOLON && p.lit == "\n" {
		p.error(p.pos, "missing ',' before newline in "+context)
		p.next()
	}
	return p.expect(tok)
}

func (p *MacroParser) expectSemi() {
	// semicolon is optional before a closing ')' or '}'
	if p.tok != token.RPAREN && p.tok != token.RBRACE {
		switch p.tok {
		case token.COMMA:
			// permit a ',' instead of a ';' but complain
			p.errorExpected(p.pos, "';'")
			fallthrough
		case token.SEMICOLON:
			p.next()
		default:
			p.errorExpected(p.pos, "';'")
			syncStmt(p)
		}
	}
}
