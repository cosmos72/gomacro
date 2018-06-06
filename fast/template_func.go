/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * template_function.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import "go/ast"

type TemplateFunc struct {
	Decl           *ast.FuncDecl
	TypeParams     []string // template type names
	SpecializedFor []ast.Expr
}

// TemplateFuncDecl stores a template function or method declaration
// for later instantiation
func (c *Comp) TemplateFuncDecl(funcdecl *ast.FuncDecl) {
	n := 0
	if funcdecl.Recv != nil {
		n = len(funcdecl.Recv.List)
	}
	if n < 2 {
		c.Errorf("invalid template function or method declaration: expecting at least 2 receivers, found %d: %v", n, funcdecl)
	}
	funcname := funcdecl.Name.Name

	funcbind := c.NewBind(funcname, TemplateFuncBind, c.TypeOfPtrTemplateFunc())

	typeParams := make([]string, n-2)
	_ = funcbind
	_ = typeParams

	c.Errorf("template functions and methods not yet implemented: %v", funcdecl)
}
