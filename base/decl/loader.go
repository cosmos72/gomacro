/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * loader.go
 *
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package decl

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base"
)

func (l *Loader) Ast(form ast2.Ast) {
	switch form := form.(type) {
	case nil:
	case ast2.AstWithNode:
		l.Node(form.Node())
	case ast2.AstWithSlice:
		n := form.Size()
		for i := 0; i < n; i++ {
			l.Ast(form.Get(i))
		}
	default:
		base.Errorf("Loader.Ast(): unsupported ast2.Ast node, expecting ast2.AstWithNode or ast2.AstWithSlice, found %v // %T", form, form)
	}
}

func (l *Loader) Node(node ast.Node) {
	switch node := node.(type) {
	case nil:
	case ast.Decl:
		l.Decl(node)
	case ast.Expr:
		// nothing to do
	case ast.Stmt:
		// nothing to do
	case *ast.File:
		l.File(node)
	default:
		base.Errorf("Loader.Ast(): unsupported node type, expecting ast.Decl, ast.Expr, ast.Stmt or *ast.File, found %v // %T", node, node)
	}
}

func (l *Loader) Nodes(nodes []ast.Node) {
	l.Ast(ast2.NodeSlice{nodes})
}

func (l *Loader) Decl(node ast.Node) {
	switch node := node.(type) {
	case nil:
	case *ast.GenDecl:
		l.GenDecl(node)
	case *ast.FuncDecl:
		l.Func(node)
	default:
		base.Errorf("Loader.Decl(): unsupported declaration, expecting *ast.GenDecl or *ast.FuncDecl, found: %v // %T", node, node)
	}
}

func (l *Loader) File(node *ast.File) {
	if node != nil {
		for _, decl := range node.Decls {
			l.Decl(decl)
		}
	}
}

func (l *Loader) GenDecl(node *ast.GenDecl) {
	switch node.Tok {
	case token.IMPORT:
	case token.CONST:
		var defaultDeps []string
		iota := 0
		for _, decl := range node.Specs {
			deps := l.Consts(decl, iota, defaultDeps)
			if valueSpec, ok := decl.(*ast.ValueSpec); ok && valueSpec.Values != nil {
				// if expressions are omitted, they default to the last ones found (with their type, if any)
				defaultDeps = deps
			}
			iota++
		}
	case token.TYPE:
		for _, decl := range node.Specs {
			l.Type(decl)
		}
	case token.VAR:
		for _, decl := range node.Specs {
			l.Vars(decl)
		}
	case token.PACKAGE:
	default:
		base.Errorf("Loader.GenDecl(): unsupported declaration kind, expecting token.IMPORT, token.PACKAGE, token.CONST, token.TYPE or token.VAR, found %v: %v // %T",
			node.Tok, node, node)
	}
}

// constants
func (l *Loader) Consts(node ast.Spec, iota int, deps []string) []string {
	if node, ok := node.(*ast.ValueSpec); ok {
		// if expressions are omitted, they default to the last ones found (with their type, if any)
		if node.Type != nil || node.Values != nil {
			deps = append(l.Expr(node.Type), l.Exprs(node.Values)...)
		}
		for _, ident := range node.Names {
			// node appears multiple times... to simplify, all occurrences have the same dependencies
			l.Const(node, ident.Name, iota, deps)
		}
	} else {
		base.Errorf("Loader.Consts(): unsupported constant declaration: expecting *ast.ValueSpec, found: %v // %T", node, node)
	}
	return deps
}

// constant
func (l *Loader) Const(node ast.Spec, name string, iota int, deps []string) {
	l.Decls[name] = &Decl{Kind: Const, Name: name, Node: node, Deps: sort_unique(deps), Iota: iota}
}

// variables
func (l *Loader) Vars(node ast.Spec) {
	if node, ok := node.(*ast.ValueSpec); ok {
		deps := append(l.Expr(node.Type), l.Exprs(node.Values)...)
		for _, ident := range node.Names {
			// node appears multiple times... to simplify, all occurrences have the same dependencies
			l.Var(node, ident.Name, deps)
		}
	} else {
		base.Errorf("Loader.Vars(): unsupported variable declaration: expecting *ast.ValueSpec, found: %v // %T", node, node)
	}
}

// variable
func (l *Loader) Var(node ast.Spec, name string, deps []string) {
	l.Decls[name] = &Decl{Kind: Var, Name: name, Node: node, Deps: sort_unique(deps)}
}

// function or method
func (l *Loader) Func(node *ast.FuncDecl) {
	name := node.Name.Name
	deps := l.Expr(node.Type)
	kind := Func
	if node.Recv != nil && len(node.Recv.List) != 0 {
		types := l.Expr(node.Recv)
		// method names are not global!
		// without this, a method Foo.String would overwrite a func String in l.Decls[]
		//
		// also makes it impossible to depend on a method, but nothing can depend on a method,
		// Except the constant returned by unsafe.Sizeof(type.method),
		// but we do not support unsafe.Sizeof() yet and all methods have the same size anyway
		if len(types) == 1 {
			name = fmt.Sprintf("%s.%s", types[0], name)
		} else {
			name = fmt.Sprintf("%d.%s", l.Gensym, name)
			l.Gensym++
		}

		deps = append(deps, types...)
		kind = Method
	}
	l.Decls[name] = &Decl{Kind: kind, Name: name, Node: node, Deps: sort_unique(deps)}
}

// type
func (l *Loader) Type(node ast.Spec) {
	if node, ok := node.(*ast.TypeSpec); ok {
		name := node.Name.Name
		deps := l.Expr(node.Type)
		deps = sort_unique(deps)

		// support self-referencing types, as for example: type List struct { First int; Rest *List }
		deps = remove_item_inplace(name, deps)

		l.Decls[name] = &Decl{Kind: Type, Name: name, Node: node, Deps: deps}
	} else {
		base.Errorf("Loader.Type(): unexpected declaration type, expecting *ast.TypeSpec, found: %v // %T", node, node)
	}
}

func (l *Loader) Expr(node ast.Node) []string {
	if node == nil {
		return nil
	}
	return l.AstExpr(ast2.ToAst(node))
}

func (l *Loader) Exprs(list []ast.Expr) []string {
	if len(list) == 0 {
		return nil
	}
	return l.AstExpr(ast2.ExprSlice{list})
}

func (l *Loader) AstExpr(in ast2.Ast) []string {
	if in == nil {
		return nil
	}
	var deps []string
	switch node := in.Interface().(type) {
	case ast.Stmt, ast.Decl, ast.Spec:
		// in expression context, do not enter statements or declarations
		return nil
	case *ast.SelectorExpr:
		// one of: package.symbol, type.method, type.field.
		// only the part *before* the dot may be a local declaration.
		return l.Expr(node.X)
	case *ast.FieldList:
		// collect (struct field, param...) types but ignore their names
		for _, field := range node.List {
			deps = append(deps, l.Expr(field.Type)...)
		}
		return deps
	}
	if form, ok := in.(ast2.Ident); ok {
		deps = append(deps, form.X.Name)
	}

	for i, n := 0, in.Size(); i < n; i++ {
		form := in.Get(i)
		if form != nil {
			deps = append(deps, l.AstExpr(form)...)
		}
	}
	return deps
}
