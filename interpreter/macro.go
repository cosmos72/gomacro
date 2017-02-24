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

package interpreter

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
		if len(node.List) == 1 {
			ret = node.List[0]
			if ret2, ok := ret.(*ast.ExprStmt); ok {
				// unwrap expressions... they fit in more places and make the life easier to MacroExpand
				ret = ret2
			}
		} else {
			// be lenient, and accept quote{a;b;c} as nicer alias for quote{{a;b;c}}
			ret = node
		}
	case mp.UNQUOTE, mp.UNQUOTE_SPLICE:
		return env.Errorf("unimplemented %s: %v", mp.TokenString(op), node)
	}
	return r.ValueOf(&ret).Elem(), nil
}

func isMacroCall(node *ast.BinaryExpr) bool {
	return node.Op == mp.MACRO
}

// MacroExpandCodewalk traverses the whole AST tree using pre-order traversal,
// and replaces each node with the result of MacroExpand(node).
// It implements the macroexpansion phase
// Warning: it modifies the AST tree in place!
func (env *Env) MacroExpandCodewalk(node ast.Node) ast.Node {
	return env.Walk(env.macroExpandReturnNilOnQuote, node)
}

// MacroExpand repeatedly invokes MacroExpand1
// as long as the node represents a macro call.
// it returns the resulting node.
func (env *Env) macroExpandReturnNilOnQuote(node ast.Node) ast.Node {
	if node == nil {
		return nil
	}
	for {
		ret := env.macroExpand1(node, true)
		if ret == nil || ret == node {
			return ret
		}
		node = ret
	}
}

// MacroExpand repeatedly invokes MacroExpand1
// as long as the node represents a macro call.
// it returns the resulting node.
func (env *Env) MacroExpand(node ast.Node) (ast.Node, bool) {
	if node == nil {
		return nil, false
	}
	save := node
	for {
		ret := env.macroExpand1(node, false)
		if ret == nil || ret == node {
			return ret, ret != save
		}
		node = ret
	}
}

// if node represents a macro call, MacroExpand1 executes it
// and returns the resulting node.
// Otherwise returns the node argument unchanged
func (env *Env) MacroExpand1(node ast.Node) (ast.Node, bool) {
	ret := env.macroExpand1(node, false)
	// if ret == node {
	//    env.Debugf("MacroExpand1() not a macro: %v <%v>", node, r.TypeOf(node))
	//}
	return ret, ret != node
}

func (env *Env) macroExpand1(node ast.Node, returnNilOnQuote bool) ast.Node {
	var expr *ast.BinaryExpr
Again:
	switch x := node.(type) {
	case *ast.BinaryExpr:
		if x.Op != mp.MACRO {
			// not a macro call, return unchanged
			return node
		}
		expr = x
		// env.Debugf("macroExpand1() found macro call: %v", expr)

	case *ast.UnaryExpr:
		// not a macro. maybe it's a QUOTE ?
		if returnNilOnQuote && x.Op == mp.QUOTE {
			// env.Debugf("macroExpand1() found QUOTE, returning nil")
			return nil
		}
		return node

	case *ast.ExprStmt:
		// expressions are sometimes wrapped in statements... unwrap them
		node = x.X
		goto Again

	default:
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

func (env *Env) badMacro(node *ast.BinaryExpr) ast.Expr {
	env.Errorf("macroexpansion of non-macro: %v", node)
	return nil
}
