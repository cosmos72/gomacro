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
		return Errorf("unimplemented declaration: %#v", node)
	}
}

func (env *Env) evalDeclGen(node *ast.GenDecl) (r.Value, []r.Value) {
	tok := node.Tok
	var ret r.Value
	var rets []r.Value
	for _, decl := range node.Specs {
		switch tok {
		case token.IMPORT:
			ret, rets = env.evalImport(decl)
		case token.CONST:
			ret, rets = env.evalDeclConst(decl)
		case token.TYPE:
			ret, rets = env.evalDeclType(decl)
		case token.VAR:
			ret, rets = env.evalDeclVar(decl)
		default:
			return Errorf("unimplemented declaration: %#v", decl)
		}
	}
	return ret, rets
}

func (env *Env) evalImport(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	default:
		return Errorf("unimplemented import %#v", node)
	}
}

func (env *Env) evalDeclConst(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	default:
		return Errorf("unimplemented constant declaration %#v", node)
	}
}

func (env *Env) evalDeclType(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	default:
		return Errorf("unimplemented type declaration %#v", node)
	}
}

func (env *Env) evalDeclVar(node ast.Spec) (r.Value, []r.Value) {
	var ret r.Value
	var rets []r.Value
	switch node := node.(type) {
	case *ast.ValueSpec:
		idents := node.Names
		values := node.Values
		t := env.evalType(node.Type)
		var zero r.Value
		if t != nil {
			// t can be nil when inferring types
			zero = r.Zero(t)
		}
		for i, ident := range idents {
			if values != nil && len(values) >= i {
				value, _ := env.Eval(values[i])
				ret, rets = env.defineVar(ident.Name, t, value)
			} else if t == nil {
				return Errorf("invalid variable declaration, type OR initializer required: %#v", node)
			} else {
				ret, rets = env.defineVar(ident.Name, t, zero)
			}
		}
	default:
		return Errorf("unimplemented variable declaration: %#v", node)
	}
	return ret, rets
}

func (env *Env) defineVarConvert(name string, t r.Type, value r.Value) (r.Value, []r.Value) {
	value_as_t, ok := toType(value, t)
	if !ok {
		return Errorf("failed to cast %#v to %v in variable declaration: var %s %v = %#v", value, t, name, t, value)
	}
	return env.defineVar(name, t, value_as_t)
}

func (env *Env) defineVar(name string, t r.Type, value r.Value) (r.Value, []r.Value) {
	if name == "_" {
		// never define bindings for "_"
		return value, nil
	}
	if t == nil {
		t = value.Type()
		// fmt.Printf("debug: defineVar() type inference: var %s <%v> = %#v\n", name, t, value.Interface())
	} else {
		// fmt.Printf("debug: defineVar() var %s %v = %#v\n", name, t, value.Interface())
	}
	if _, exists := env.Binds[name]; exists {
		Warnf("redefined identifier: %v\n", name)
	}
	addr := r.New(t)

	env.assign(addr.Elem(), token.ASSIGN, value)
	env.Binds[name] = addr.Elem()
	// fmt.Printf("debug: defineVar() added %#v to %#v\n", name, env.Binds)
	return value, nil
}
