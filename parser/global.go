// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the exported entry points for invoking the parser.

package parser

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
	TraceMacro                                     // print debug statements while parsing macro invokations
)

func (p *Parser) Parse() (node []ast.Node, err error) {
	if p.file == nil || p.pkgScope == nil {
		panic("Parser.Parse(): parser is not initialized, call Parser.Init() first")
	}

	defer func() {
		if e := recover(); e != nil {
			// resume same panic if it's not a bailout
			if _, ok := e.(Bailout); !ok {
				panic(e)
			}
		}
		p.errors.Sort()
		err = p.errors.Err()
		p.file = nil
		p.pkgScope = nil
	}()

	topScope := p.topScope

	list := make([]ast.Node, 0)
	for p.tok != token.EOF {
		list = append(list, p.parseAny())
	}

	assert(topScope == p.topScope, "unbalanced scopes")

	if p.errors.Len() > 0 {
		p.errors.Sort()
		err = p.errors.Err()
	}
	return list, err
}
