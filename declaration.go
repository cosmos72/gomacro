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
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"
)

func (ir *Interpreter) evalDecl(node ast.Decl) (r.Value, error) {
	switch node := node.(type) {
	case *ast.GenDecl:
		return ir.evalDeclGen(node)
	case *ast.FuncDecl:
		return ir.evalDeclFunc(node)
	default:
		return Nil, errors.New(fmt.Sprintf("unsupported declaration: %#v", node))
	}
}

func (ir *Interpreter) evalDeclGen(node *ast.GenDecl) (r.Value, error) {
	tok := node.Tok
	var ret r.Value
	var err error
	for _, decl := range node.Specs {
		switch tok {
		case token.IMPORT:
			ret, err = ir.evalImport(decl)
		case token.CONST:
			ret, err = ir.evalDeclConst(decl)
		case token.TYPE:
			ret, err = ir.evalDeclType(decl)
		case token.VAR:
			ret, err = ir.evalDeclVar(decl)
		default:
			return Nil, errors.New(fmt.Sprintf("unsupported declaration: %#v", decl))
		}
		if err != nil {
			ir.PopEnv()
			return Nil, err
		}
	}
	return ret, nil
}

func (ir *Interpreter) evalImport(node ast.Spec) (r.Value, error) {
	switch node := node.(type) {
	default:
		return Nil, errors.New(fmt.Sprintf("unsupported import %#v", node))
	}
}

func (ir *Interpreter) evalDeclConst(node ast.Spec) (r.Value, error) {
	switch node := node.(type) {
	default:
		return Nil, errors.New(fmt.Sprintf("unsupported constant declaration %#v", node))
	}
}

func (ir *Interpreter) evalDeclType(node ast.Spec) (r.Value, error) {
	switch node := node.(type) {
	default:
		return Nil, errors.New(fmt.Sprintf("unsupported type declaration %#v", node))
	}
}

func (ir *Interpreter) evalDeclVar(node ast.Spec) (r.Value, error) {
	var ret r.Value
	switch node := node.(type) {
	case *ast.ValueSpec:
		idents := node.Names
		values := node.Values
		t, err := ir.evalType(node.Type)
		if err != nil {
			return Nil, err
		}
		var zero r.Value
		if t != nil {
			// t can be nil when inferring types
			zero = r.Zero(t)
		}
		for i, ident := range idents {
			if values != nil && len(values) >= i {
				value, err := ir.Eval(values[i])
				if err == nil {
					ret, err = ir.defineVar(ident.Name, t, value)
				}
				if err != nil {
					return Nil, err
				}
			} else if t == nil {
				return Nil, errors.New(fmt.Sprintf("invalid variable declaration, type OR initializer required: %#v", node))
			} else {
				ret, err = ir.defineVar(ident.Name, t, zero)
				if err != nil {
					return Nil, err
				}
			}
		}
	default:
		return Nil, errors.New(fmt.Sprintf("unsupported variable declaration: %#v", node))
	}
	return ret, nil
}

func (ir *Interpreter) defineVarConvert(name string, t r.Type, value r.Value) (r.Value, error) {
	value_as_t, ok := toType(value, t)
	if !ok {
		return Nil, errors.New(fmt.Sprintf("failed to cast %#v to %v in variable declaration: var %s %v = %#v", value, t, name, t, value))
	}
	return ir.defineVar(name, t, value_as_t)
}

func (ir *Interpreter) defineVar(name string, t r.Type, value r.Value) (r.Value, error) {

	if t == nil {
		t = value.Type()
		// fmt.Printf("debug: defineVar() type inference: var %s <%v> = %#v\n", name, t, value.Interface())
	} else {
		// fmt.Printf("debug: defineVar() var %s %v = %#v\n", name, t, value.Interface())
	}
	addr := r.New(t)

	_, err := ir.assign(addr.Elem(), token.ASSIGN, value)
	if err != nil {
		return Nil, err
	}
	ir.Binds[name] = addr.Elem()
	// fmt.Printf("debug: defineVar() added %#v to %#v\n", name, ir.Binds)
	return r.ValueOf(name), nil
}
