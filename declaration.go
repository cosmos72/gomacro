/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * declaration.go
 *
 *  Created on: Feb 13, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	"go/token"
	r "reflect"
)

func (env *Env) evalDecl(node ast.Decl) (r.Value, []r.Value) {
	switch node := node.(type) {
	case *ast.GenDecl:
		return env.evalDeclGen(node)
	case *ast.FuncDecl:
		return env.evalDeclFunc(node)
	default:
		return env.Errorf("unimplemented declaration: %#v", node)
	}
}

func (env *Env) evalDeclGen(node *ast.GenDecl) (r.Value, []r.Value) {
	tok := node.Tok
	var ret r.Value
	var rets []r.Value
	for _, decl := range node.Specs {
		switch tok {
		case token.IMPORT:
			ret, rets = env.evalImports(decl)
		case token.CONST:
			ret, rets = env.evalDeclConsts(decl)
		case token.TYPE:
			ret, rets = env.evalDeclTypes(decl)
		case token.VAR:
			ret, rets = env.evalDeclVars(decl)
		default:
			return env.Errorf("unimplemented declaration: %v", decl)
		}
	}
	return ret, rets
}

func (env *Env) evalImports(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	default:
		return env.Errorf("unimplemented: import: %v", node)
	}
}

func (env *Env) evalDeclConsts(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	default:
		return env.Errorf("unimplemented constant declaration %#v", node)
	}
}

func (env *Env) evalDeclTypes(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	default:
		return env.Errorf("unimplemented type declaration %#v", node)
	}
}

func (env *Env) evalDeclVars(node ast.Spec) (r.Value, []r.Value) {
	var ret r.Value
	var rets []r.Value
	switch node := node.(type) {
	case *ast.ValueSpec:
		return env.evalDeclVarsExpr(node.Names, node.Type, node.Values)
	default:
		return env.Errorf("unimplemented variable declaration: %v", node)
	}
	return ret, rets
}

func (env *Env) evalDeclVarsExpr(idents []*ast.Ident, typ ast.Expr, exprs []ast.Expr) (r.Value, []r.Value) {
	n := len(idents)
	names := make([]string, n)
	for i, ident := range idents {
		names[i] = ident.Name
	}
	t := env.evalType(typ)

	var values []r.Value
	if exprs != nil {
		values = env.evalExprsMultipleValues(exprs, n)
	}
	return env.defineVars(names, t, values)
}

func (env *Env) defineVars(names []string, t r.Type, values []r.Value) (r.Value, []r.Value) {
	n := len(names)
	if values == nil {
		values = make([]r.Value, n)
		zero := r.Zero(t)
		for i := 0; i < n; i++ {
			values[i] = env.defineVar(names[i], t, zero)
		}
	} else {
		for i := 0; i < n; i++ {
			values[i] = env.defineVar(names[i], t, values[i])
		}
	}
	return unpackValues(values)
}

func (env *Env) defineVar(name string, t r.Type, value r.Value) r.Value {
	if name == "_" {
		// never define bindings for "_"
		if t != nil {
			value = env.toType(value, t)
		}
		return value
	}
	if t == nil {
		t = value.Type()
		// Debugf("defineVar() type inference: var %s <%v> = %#v", name, t, value.Interface())
	} else {
		// Debugf("defineVar() var %s %v = %#v", name, t, value.Interface())
	}
	if _, exists := env.binds[name]; exists {
		env.Warnf("redefined identifier: %v", name)
	}
	addr := r.New(t)

	value = env.assignPlace(addr.Elem(), token.ASSIGN, value)
	env.binds[name] = addr.Elem()
	// Debugf("defineVar() added %#v to %#v", name, env.Binds)
	return value
}
