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
 * macro.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	"go/token"
	r "reflect"

	mp "github.com/cosmos72/gomacro/macroparser"
)

func (env *Env) evalQuote(op token.Token, node *ast.BlockStmt) (r.Value, []r.Value) {
	var ret ast.Node
	switch op {
	case mp.QUOTE, mp.QUASIQUOTE:
		if len(node.List) != 1 {
			return env.Errorf("invalid %s: contains %d statements, expecting exactly one: %v", mp.TokenString(op), len(node.List), node)
		}
		ret = node.List[0]
	case mp.UNQUOTE, mp.UNQUOTE_SPLICE:
		return env.Errorf("unimplemented %s: %v", mp.TokenString(op), node)
	}
	return r.ValueOf(&ret).Elem(), nil
}

func isMacroCall(node *ast.BinaryExpr) bool {
	return node.Op == mp.MACRO
}

func (env *Env) MacroExpand(node ast.Node) ast.Node {
	for {
		ret := env.MacroExpand1(node)
		if ret == nil || ret == node {
			return ret
		}
		node = ret
	}
}

func (env *Env) MacroExpand1(node ast.Node) ast.Node {
	expr := env.extractMacroExpr(node)
	if expr == nil {
		return node
	}

	// retrieve and validate the macro
	macro := env.Eval1(expr.X)
	if macro == Nil || macro == None || macro.Kind() != r.Struct {
		return env.badMacro(expr)
	}
	m, ok := macro.Interface().(Macro)
	if !ok || m.Closure == nil {
		return env.badMacro(expr)
	}
	// validate the macroexpansion arguments
	fun, ok := expr.Y.(*ast.FuncLit)
	if !ok || len(fun.Type.Params.List) != 0 {
		return env.badMacro(expr)
	}
	args := fun.Body.List
	n := len(args)
	if n > m.ArgNum {
		env.Errorf("too many arguments in macroexpansion of %v", node)
		return nil
	} else if n > m.ArgNum {
		env.Errorf("not enough arguments in macroexpansion of %v", node)
		return nil
	}
	// wrap each ast.Stmt into a reflect.Value
	argsv := make([]r.Value, n)
	for i := 0; i < n; i++ {
		argsv[i] = r.ValueOf(args[i])
	}
	// invoke the macro
	results := m.Closure(argsv)
	// validate the results
	switch n = len(results); n {
	default:
		env.Warnf("macroexpansion returned %d values, only the first one will be used: %v", n, node)
		fallthrough
	case 1:
		result := results[0]
		if result != Nil && result != None {
			ret, ok := result.Interface().(ast.Node)
			if ok {
				return ret
			}
			env.Errorf("macroexpansion returned a <%v>, not an <ast.Node>: %v", result.Type(), node)
		}
		fallthrough
	case 0:
		return nil
	}
}

func (env *Env) extractMacroExpr(node ast.Node) *ast.BinaryExpr {
	for {
		switch x := node.(type) {
		case *ast.BinaryExpr:
			return x
		case *ast.ExprStmt:
			node = x.X
			continue
		default:
			break
		}
		return nil
	}
}

func (env *Env) badMacro(node *ast.BinaryExpr) ast.Expr {
	env.Errorf("macroexpansion of non-macro: %v", node)
	return nil
}

func (env *Env) nodeToExpr(node ast.Node) ast.Expr {
	switch node := node.(type) {
	case ast.Expr:
		return node
	case *ast.BlockStmt:
		return env.blockStmtToExpr(node)
	case ast.Stmt:
		list := []ast.Stmt{node}
		block := &ast.BlockStmt{List: list}
		return env.blockStmtToExpr(block)
	default:
		env.Errorf("macroexpansion returned a <%T>: unimplemented conversion to ast.Expr", node)
		return nil
	}
}

func (env *Env) blockStmtToExpr(node *ast.BlockStmt) ast.Expr {
	// due to go/ast strictly typed model, there is only one mechanism
	// to insert a statement inside an expression: use a closure.
	// so we return a unary expression: MACRO (func() { /*block*/ })
	typ := &ast.FuncType{Func: token.NoPos, Params: &ast.FieldList{}}
	fun := &ast.FuncLit{Type: typ, Body: node}
	return &ast.UnaryExpr{Op: mp.MACRO, X: fun}
}
