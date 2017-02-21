// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the exported entry points for invoking the parser.

package macroparser2

import (
	"go/token"
)

// A Mode value is a set of flags (or 0).
// They control the amount of source code parsed and other optional
// parser functionality.
//
type Mode uint

type Ident struct {
	*string // not string because we need the pointer as key to index its metadata
}

type Slice []interface{}

type Attr int
type Attrs map[Attr]interface{}

type Node interface {
	Attr(p *MacroParser, key Attr) interface{}
	Pos(p *MacroParser) token.Pos
	End(p *MacroParser) token.Pos
	Bad() bool
}

func slice(x ...interface{}) Slice {
	return x
}

const (
	// parse mode
	PackageClauseOnly Mode             = 1 << iota // stop parsing after package clause
	ImportsOnly                                    // stop parsing after import declarations
	ParseComments                                  // parse comments and add them to AST
	Trace                                          // print a trace of parsed productions
	DeclarationErrors                              // report declaration errors
	SpuriousErrors                                 // same as AllErrors, for backward-compatibility
	AllErrors         = SpuriousErrors             // report all errors (not just the first 10 on different lines)
)

const (
	// keyword tokens
	MACRO token.Token = token.VAR + 101 + iota
	QUOTE
	QUASIQUOTE
	UNQUOTE
	UNQUOTE_SPLICE
)

const (
	// Ident and Slice attributes
	DOC Attr = iota
	POS      // start position
	END      // end   position
	SCOPE
	IMPORTS
	UNRESOLVED
	COMMENTS
	OBJ // for identifiers
	TAG // for struct fields
)

const (
	// reserved function, macro and operator names
	BAD_s    = "???"
	BLOCK_s  = "{}"
	PERIOD_s = "."
	DOT_S    = "."
	INDEX_s  = "[]"
	PAREN_s  = "()"

	ADD_s  = "+"
	SUB_s  = "-"
	MUL_s  = "*"
	STAR_s = "*"
	QUO_s  = "/"
	REM_s  = "%"

	AND_s     = "&"
	OR_s      = "|"
	XOR_s     = "^"
	SHL_s     = "<<"
	SHR_s     = ">>"
	AND_XOR_s = "&^"

	ADD_ASSIGN_s = "+="
	SUB_ASSIGN_s = "-="
	MUL_ASSIGN_s = "*="
	QUO_ASSIGN_s = "/="
	REM_ASSIGN_s = "%="

	AND_ASSIGN_s     = "&="
	OR_ASSIGN_s      = "|="
	XOR_ASSIGN_s     = "^="
	SHL_ASSIGN_s     = "<<="
	SHR_ASSIGN_s     = ">>="
	AND_NOT_ASSIGN_s = "&^="
	INDEX_ASSIGN_s   = "[]="

	LAND_s  // &&
	LOR_s   // ||
	ARROW_s = "<-"
	INC_s   // ++
	DEC_s   // --

	EQL_s    = "=="
	LSS_s    = "<"
	GTR_s    = ">"
	ASSIGN_s = "="
	NOT_s    = "!"

	NEQ_s      = "!="
	LEQ_s      = "<="
	GEQ_s      = ">="
	DEFINE_s   = ":="
	ELLIPSIS_s = "..."

	PACKAGE_s   = "package"
	IMPORT_s    = "import"
	CONST_s     = "const"
	FUNC_s      = "func"
	LAMBDA_s    = "λ" // "func()"
	INTERFACE_s = "interface"
	STRUCT_s    = "struct"
	TYPE_s      = "type"
	VAR_s       = "var"
	CHAN_s      = "chan"
	MAP_s       = "map"

	FIELDS_s  = "field[]"
	FIELDSS_s = "field[][]"
	PARAMS_s  = "param[]"
	PARAMSS_s = "param[][]"

	LABEL_s              = ":"
	SEMICOLON_s          = ";"
	SEMICOLON_IMPLICIT_s = "\n"

	BREAK_s       = "break"
	CONTINUE_s    = "continue"
	DEFER_s       = "defer"
	FALLTHROUGH_s = "fallthrough"
	FOR_s         = "for"
	GO_s          = "go"
	GOTO_s        = "goto"
	IF_s          = "if"
	RETURN_s      = "return"
	SWITCH_s      = "switch"
	SELECT_s      = "select"

	MACRO_s          = "macro"
	QUASIQUOTE_s     = "quasiquote"
	QUOTE_s          = "quote"
	UNQUOTE_s        = "unquote"
	UNQUOTE_SPLICE_s = "unquote_splice"
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

func TokenIsMacroKeyword(tok token.Token) bool {
	_, ok := tokens[tok]
	return ok
}

/*
func Parse(fileset *token.FileSet, filename string, src []byte, mode Mode) (node Any, err error) {
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
*/
