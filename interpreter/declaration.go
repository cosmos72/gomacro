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
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

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
		return env.evalDeclNamedFunction(node)
	default:
		return env.Errorf("unimplemented declaration: %v", node)
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
			ret, rets = env.evalDeclConsts(decl)
		case token.TYPE:
			return env.evalDeclType(decl)
		case token.VAR:
			ret, rets = env.evalDeclVars(decl)
		default:
			return env.Errorf("unimplemented declaration: %v", decl)
		}
	}
	return ret, rets
}

func (env *Env) evalDeclConsts(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	case *ast.ValueSpec:
		return env.evalDeclConstsOrVars(node.Names, node.Type, node.Values, true)
	default:
		return env.Errorf("unexpected constant declaration: expecting *ast.ValueSpec, found: %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) evalDeclType(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	case *ast.TypeSpec:
		name := node.Name.Name
		t := env.evalType(node.Type)
		if _, ok := env.Types[name]; ok {
			env.Warnf("redefined type: %v", name)
		} else if env.Types == nil {
			env.Types = make(map[string]r.Type)
		}
		env.Types[name] = t
		return r.ValueOf(&t).Elem(), nil // always return a reflect.Type

	default:
		return env.Errorf("unexpected type declaration: expecting *ast.TypeSpec, found: %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) evalDeclVars(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	case *ast.ValueSpec:
		return env.evalDeclConstsOrVars(node.Names, node.Type, node.Values, false)
	default:
		return env.Errorf("unexpected variable declaration: expecting *ast.ValueSpec, found: %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) evalDeclConstsOrVars(idents []*ast.Ident, typ ast.Expr, exprs []ast.Expr, constant bool) (r.Value, []r.Value) {
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
	return env.defineConstsVarsOrFuncs(names, t, values, constant)
}

func (env *Env) defineConstsVarsOrFuncs(names []string, t r.Type, values []r.Value, constant bool) (r.Value, []r.Value) {
	n := len(names)
	if values == nil {
		if t == nil {
			return env.Errorf("no values and no type: cannot define %v", names)
		}
		values = make([]r.Value, n)
		zero := r.Zero(t)
		for i := 0; i < n; i++ {
			values[i] = env.defineConstVarOrFunc(names[i], t, zero, constant)
		}
	} else {
		for i := 0; i < n; i++ {
			values[i] = env.defineConstVarOrFunc(names[i], t, values[i], constant)
		}
	}
	return unpackValues(values)
}

func (env *Env) defineConst(name string, t r.Type, value r.Value) r.Value {
	return env.defineConstVarOrFunc(name, t, value, true)
}

func (env *Env) defineVar(name string, t r.Type, value r.Value) r.Value {
	return env.defineConstVarOrFunc(name, t, value, false)
}

func (env *Env) defineFunc(name string, t r.Type, value r.Value) r.Value {
	return env.defineConstVarOrFunc(name, t, value, true)
}

func (env *Env) defineConstVarOrFunc(name string, t r.Type, value r.Value, constant bool) r.Value {
	if name == "_" {
		// never define bindings for "_"
		if t != nil {
			value = env.valueToType(value, t)
		}
		return value
	}
	if t == nil {
		t = TypeOf(value)
	}
	if _, exists := env.Binds[name]; exists {
		env.Warnf("redefined identifier: %v", name)
	}
	if constant {
		value = value.Convert(t)
		env.Binds[name] = value
	} else {
		addr := r.New(t)
		value = env.assignPlace(Place{addr.Elem(), Nil}, token.ASSIGN, value)
		env.Binds[name] = addr.Elem()
	}
	// Debugf("defineConstVarOrFunc() added %#v to %#v", name, env.Binds)
	return value
}
