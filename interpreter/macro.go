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
	_ "go/token"
	r "reflect"

	mt "github.com/cosmos72/gomacro/token"
)

func simplifyNodeForQuote(in ast.Node) ast.Node {
	for {
		switch node := in.(type) {
		case *ast.BlockStmt:
			list := node.List
			if len(list) == 1 {
				in = list[0]
				continue
			}
		case *ast.ExprStmt:
			return node.X
		}
		return in
	}
}

func isLeaf(node ast.Node) bool {
	switch node.(type) {
	case *ast.Ident, *ast.BasicLit:
		return true
	default:
		return false
	}
}

func (env *Env) evalQuote(node *ast.BlockStmt) ast.Node {
	// unwrap expressions... they fit in more places and make the life easier
	// to MacroExpand and evalQuasiquote
	// also be lenient, and accept quote{a;b;c} as nicer alias for quote{{a;b;c}}
	return simplifyNodeForQuote(node)
}

// evalQuasiquote evaluates the body of a quasiquote{} represented as ast.Node
func (env *Env) evalQuasiquote(node ast.Node) ast.Node {
	// use unified API to traverse ast.Node... every other solution is a nightmare
	in := ToAst(node)
	out := env.evalQuasiquoteAst(in)
	return ToNode(out)
}

// evalQuasiquoteAst evaluates the body of a quasiquote{} represented as Ast
func (env *Env) evalQuasiquoteAst(in Ast) Ast {
	switch ast := in.(type) {
	case nil:
		return nil
	case UnaryExpr:
		switch ast.Op() {
		case mt.UNQUOTE:
			return getQuoteContent(ast)
		case mt.UNQUOTE_SPLICE:
			env.Errorf("quasiquote(): cannot splice in single-statement context: %#v %<v>", ast, r.TypeOf(ast))
		}
	}

	ast := in
	ni := ast.Size()
	if ni == 0 {
		return ast
	}
	astWithResize, canResize := ast.(AstWithResize)
	rets := make([]Ast, 0, ni)
	for i := 0; i < ni; i++ {
		child := ast.Get(i)
		switch child := child.(type) {
		case UnaryExpr:
			switch child.Op() {
			case mt.UNQUOTE:
				toInsert := getQuoteContent(child)
				rets = append(rets, toInsert)
				continue
			case mt.UNQUOTE_SPLICE:
				if !canResize {
					env.Errorf("quasiquote(): cannot splice in fixed-length context: %#v %<v>", child, r.TypeOf(child))
					return nil
				}
				toSplice := getQuoteContent(child)
				nj := toSplice.Size()
				for j := 0; j < nj; j++ {
					rets = append(rets, toSplice.Get(j))
				}
				continue
			}
		}
		// general case: recurse on child
		child = env.evalQuasiquoteAst(child)
		rets = append(rets, child)
	}
	if len(rets) != ni {
		if !canResize {
			env.Errorf("quasiquote() internal error: attempt to splice in fixed-length context: %#v %<v>", ast, r.TypeOf(ast))
			return nil
		}
		astWithResize.Resize(len(rets))
	}
	for i, r := range rets {
		ast.Set(i, r)
	}
	return ast
}

// getQuoteContent returns the content of a QUOTE, QUASIQUOTE, UNQUOTE or UNQUOTE_SPLICE Ast
func getQuoteContent(ast UnaryExpr) Ast {
	return ast.Get(0).(FuncLit).Get(1)
}

func isMacroCall(node *ast.BinaryExpr) bool {
	return node.Op == mt.MACRO
}

type macroExpandCtx struct {
	env           *Env
	traverseQuote bool
}

// MacroExpandCodewalk traverses the whole AST tree using pre-order traversal,
// and replaces each node with the result of MacroExpand(node).
// It implements the macroexpansion phase
// Warning: it modifies the AST tree in place!
func (env *Env) MacroExpandCodewalk(node ast.Node) ast.Node {
	// TODO
	return node
}

// MacroExpand repeatedly invokes MacroExpand1
// as long as the node represents a macro call.
// it returns the resulting node.
func (env *Env) MacroExpand(node ast.Node) (ast.Node, bool) {
	if node == nil {
		return nil, false
	}
	save := node
	ctx := macroExpandCtx{env: env}
	for {
		ret, _ := ctx.macroExpand1(node)
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
	ctx := macroExpandCtx{env: env}
	ret, _ := ctx.macroExpand1(node)
	// if ret == node {
	//    env.Debugf("MacroExpand1() not a macro: %v <%v>", node, r.TypeOf(node))
	//}
	return ret, ret != node
}

func (ctx *macroExpandCtx) macroExpand1(in ast.Node) (out ast.Node, traverseOut bool) {
	var expr *ast.BinaryExpr
Again:
	switch node := in.(type) {
	case *ast.BinaryExpr:
		if node.Op != mt.MACRO {
			// not a macro call, return unchanged
			return node, true
		}
		expr = node
		// env.Debugf("macroExpand1() found macro call: %v", expr)

	case *ast.UnaryExpr:
		// not a macro. it could be QUOTE/QUASIQUOTE/UNQUOTE/UNQUOTE_SPLICE
		switch node.Op {
		case mt.QUOTE:
			return node, ctx.traverseQuote
		case mt.QUASIQUOTE:
			return ctx.env.evalQuasiquote(node.X.(*ast.FuncLit).Body), true
		case mt.UNQUOTE, mt.UNQUOTE_SPLICE:
			ctx.env.Errorf("%s not inside quasiquote", mt.String(node.Op))
		}
		return node, true

	case *ast.ExprStmt:
		// expressions are sometimes wrapped in statements... unwrap them
		in = node.X
		goto Again

	default:
		return node, true
	}

	// retrieve and validate the macro
	env := ctx.env
	macro := env.Eval1(expr.X)
	if macro == Nil || macro == None || macro.Kind() != r.Struct {
		return env.badMacro(expr), false
	}
	m, ok := macro.Interface().(Macro)
	if !ok || m.Closure == nil {
		return env.badMacro(expr), false
	}
	// validate the macroexpansion arguments
	fun, ok := expr.Y.(*ast.FuncLit)
	if !ok || len(fun.Type.Params.List) != 0 {
		return env.badMacro(expr), false
	}
	args := fun.Body.List
	n := len(args)
	if n > m.ArgNum {
		env.Errorf("too many arguments in macroexpansion of %v", expr)
		return expr, false
	} else if n > m.ArgNum {
		env.Errorf("not enough arguments in macroexpansion of %v", expr)
		return expr, false
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
		env.Warnf("macroexpansion returned %d values, only the first one will be used: %v", n, expr)
		fallthrough
	case 1:
		result := results[0]
		if result != Nil && result != None {
			ret, ok := result.Interface().(ast.Node)
			if ok {
				return ret, true
			}
			env.Errorf("macroexpansion returned a <%v>, not an <ast.Node>: %v", result.Type(), expr)
		}
		fallthrough
	case 0:
		return expr, false
	}
}

func (env *Env) badMacro(node *ast.BinaryExpr) ast.Expr {
	env.Errorf("macroexpansion of non-macro: %v", node)
	return nil
}
