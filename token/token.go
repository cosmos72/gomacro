// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
//
package token

import (
	base "go/token"
)

const (
	MACRO base.Token = (base.VAR+127)&^127 + iota
	QUOTE
	QUASIQUOTE
	UNQUOTE
	UNQUOTE_SPLICE
	SPLICE
)

var tokens map[base.Token]string

var keywords map[string]base.Token

func init() {
	tokens = map[base.Token]string{
		MACRO:          "macro",
		SPLICE:         "splice",
		QUOTE:          "quote",
		QUASIQUOTE:     "quasiquote",
		UNQUOTE:        "unquote",
		UNQUOTE_SPLICE: "unquote_splice",
	}

	keywords = make(map[string]base.Token)
	for k, v := range tokens {
		keywords[v] = k
	}
}

// Find maps an identifier to its keyword token or IDENT (if not a keyword).
//
func Lookup(lit string) base.Token {
	if tok, ok := keywords[lit]; ok {
		return tok
	}
	return base.Lookup(lit)
}

func String(tok base.Token) string {
	if str, ok := tokens[tok]; ok {
		return str
	}
	return tok.String()
}

// Predicates

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
//
func IsLiteral(tok base.Token) bool {
	return tok.IsLiteral()
}

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
//
func IsOperator(tok base.Token) bool {
	return tok.IsOperator()
}

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
//
func IsKeyword(tok base.Token) bool {
	return tok.IsKeyword()
}

// IsMacroKeyword returns true for tokens corresponding to macro-related keywords;
// it returns false otherwise.
//
func IsMacroKeyword(tok base.Token) bool {
	_, ok := tokens[tok]
	return ok
}
