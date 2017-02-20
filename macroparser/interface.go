// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the exported entry points for invoking the parser.

package macroparser

import (
	"go/ast"
	"go/token"
)

// A Mode value is a set of flags (or 0).
// They control the amount of source code parsed and other optional
// parser functionality.
//
type Mode uint

const (
	PackageClauseOnly Mode             = 1 << iota // stop parsing after package clause
	ImportsOnly                                    // stop parsing after import declarations
	ParseComments                                  // parse comments and add them to AST
	Trace                                          // print a trace of parsed productions
	DeclarationErrors                              // report declaration errors
	SpuriousErrors                                 // same as AllErrors, for backward-compatibility
	AllErrors         = SpuriousErrors             // report all errors (not just the first 10 on different lines)
)

const (
	MACRO token.Token = token.VAR + 101 + iota
	QUOTE
	QUASIQUOTE
	UNQUOTE
	UNQUOTE_SPLICE
)

var tokens map[token.Token]string

var keywords map[string]token.Token

func init() {
	tokens = make(map[token.Token]string)
	keywords = make(map[string]token.Token)

	tokens[MACRO] = "macro"
	tokens[QUOTE] = "quote"
	tokens[QUASIQUOTE] = "quasiquote"
	tokens[UNQUOTE] = "unquote"
	tokens[UNQUOTE_SPLICE] = "unquote_splice"

	for k, v := range tokens {
		keywords[v] = k
	}
}

func TokenString(tok token.Token) string {
	if str, ok := tokens[tok]; ok {
		return str
	}
	return tok.String()
}

func Parse(fileset *token.FileSet, filename string, src []byte, mode Mode) (node ast.Node, err error) {
	var p MacroParser
	defer func() {
		if e := recover(); e != nil {
			// resume same panic if it's not a bailout
			if _, ok := e.(Bailout); !ok {
				panic(e)
			}
		}
		p.errors.Sort()
		err = p.errors.Err()
	}()

	// parse expr
	p.init(fileset, filename, src, mode)
	// Set up pkg-level scopes to avoid nil-pointer errors.
	// This is not needed for a correct expression x as the
	// parser will be ok with a nil topScope, but be cautious
	// in case of an erroneous x.
	p.openScope()
	p.pkgScope = p.topScope
	node = p.parseAny()
	p.closeScope()
	assert(p.topScope == nil, "unbalanced scopes")

	if p.errors.Len() > 0 {
		p.errors.Sort()
		err = p.errors.Err()
	}
	return node, err
}
