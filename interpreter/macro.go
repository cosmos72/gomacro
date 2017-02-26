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

func simplifyNodeForQuote(in ast.Node, unwrapTrivialBlocks bool) ast.Node {
	// unwrap expressions... they fit in more places and make the life easier
	// to MacroExpand and evalQuasiquote
	// also, only for quote{},
	// unwrap single-element blocks { foo } to foo
	// unless their only element is itself a block
	for {
		switch node := in.(type) {
		case *ast.BlockStmt:
			if unwrapTrivialBlocks {
				switch len(node.List) {
				case 0:
					return &ast.EmptyStmt{Semicolon: node.End(), Implicit: false}
				case 1:
					in = node.List[0]
					unwrapTrivialBlocks = false
					continue
				}
			}
			return node
		case *ast.ExprStmt:
			return node.X
		case *ast.ParenExpr:
			return node.X
		case *ast.DeclStmt:
			return node.Decl
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
	return simplifyNodeForQuote(node, true)
}

// evalQuasiquote evaluates the body of a quasiquote{} represented as ast.Node
func (env *Env) evalQuasiquote(node ast.Node) ast.Node {
	// we invoke simplifyNodeForQuote() at the end, not at the beginning.
	// reason: to support quasiquote{unquote_splice ...}
	toUnwrap := node != simplifyNodeForQuote(node, true)

	// use unified API to traverse ast.Node... every other solution is a nightmare
	in := ToAst(node)
	out := env.evalQuasiquoteAst(in)
	node = ToNode(out)
	return simplifyNodeForQuote(node, toUnwrap)
}

// evalQuasiquoteAst evaluates the body of a quasiquote{} represented as Ast
func (env *Env) evalQuasiquoteAst(inout Ast) Ast {
	withSlice, canSplice := inout.(AstWithSlice)
	form := inout
	if env.Options&OptDebugQuasiquote != 0 {
		x := form.Interface()
		env.Debugf("quasiquote: evaluating %v <%v> (canSplice = %v)", x, r.TypeOf(x), canSplice)
	}
	if form == nil {
		return nil
	}
	ni := form.Size()
	if !canSplice {
		switch form := form.(type) {
		case UnaryExpr:
			switch form.Op() {
			case mt.UNQUOTE:
				y := env.evalUnquote(form)
				return AnyToAst(y, "unquote")
			case mt.UNQUOTE_SPLICE:
				x := form.Interface()
				env.Errorf("quasiquote: cannot splice inout single-statement context: %v <%v>", x, r.TypeOf(x))
			}
		}
		for i := 0; i < ni; i++ {
			// general case: recurse on child
			child := form.Get(i)
			child = env.evalQuasiquoteAst(child)
			form.Set(i, child)
		}
		// we modified inout destructively... return inout, not form!
		return inout
	}
	for {
		if ni == 0 {
			return inout
		} else if ni == 1 && isTrivialAst(form) {
			form = form.Get(0)
			ni = form.Size()
			if env.Options&OptDebugQuasiquote != 0 {
				x := form.Interface()
				env.Debugf("quasiquote: simplified to %v <%v> (canSplice = %v)", x, r.TypeOf(x), canSplice)
			}
		} else {
			break
		}
	}
	ret := make([]Ast, 0, ni)
	for i := 0; i < ni; i++ {
		child := form.Get(i)
		for isTrivialAst(child) {
			// drill through DeclStmt, ExprStmt, ParenExpr
			child = child.Get(0)
		}
		switch child := child.(type) {
		case UnaryExpr:
			switch child.Op() {
			case mt.UNQUOTE:
				toInsert := AnyToAst(env.evalUnquote(child), "unquote")
				ret = append(ret, toInsert)
				if env.Options&OptDebugQuasiquote != 0 {
					y := child.Interface()
					env.Debugf("quasiquote: accumulated expansion after unquote(%v <%v>) is: %v <%v> (canSplice = %v)",
						y, r.TypeOf(y), ret, r.TypeOf(ret), canSplice)
				}
				continue
			case mt.UNQUOTE_SPLICE:
				y := env.evalUnquote(child)
				toSplice := AnyToAstWithSlice(y, "unquote_splice")
				nj := toSplice.Size()
				for j := 0; j < nj; j++ {
					ret = append(ret, toSplice.Get(j))
				}
				if env.Options&OptDebugQuasiquote != 0 {
					y := child.Interface()
					env.Debugf("quasiquote: accumulated expansion after unquote_splice(%v <%v>) is: %v <%v> (canSplice = %v)",
						y, r.TypeOf(y), ret, r.TypeOf(ret), canSplice)
				}
				continue
			}
		}
		// general case: recurse on child
		child = env.evalQuasiquoteAst(child)
		ret = append(ret, child)
		if env.Options&OptDebugQuasiquote != 0 {
			y := child.Interface()
			env.Debugf("quasiquote: accumulated expansion after quasiquote(%v <%v>) is: %v <%v> (canSplice = %v)",
				y, r.TypeOf(y), ret, r.TypeOf(ret), canSplice)
		}
	}
	withSlice.Slice(0, 0)
	for _, node := range ret {
		withSlice.Append(node)
	}
	return withSlice
}

// ParenExpr, ExprStmt, DeclStmt are trivial wrappers for their contents...
func isTrivialAst(form Ast) bool {
	switch form.(type) {
	case ParenExpr, ExprStmt, DeclStmt:
		// TODO is InterfaceType trivial?
		return true
	default:
		return false
	}
}

// evalUnquote performs expansion inside a QUASIQUOTE
func (env *Env) evalUnquote(inout UnaryExpr) interface{} {
	block := inout.p.X.(*ast.FuncLit).Body

	ret, extraValues := env.evalBlock(block)
	if len(extraValues) > 1 {
		env.Warnf("unquote returned %d values, only the first one will be used: %v", len(extraValues), block)
	}
	if ret == None || ret == Nil {
		return nil
	}
	return ret.Interface()
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
// Warning: it modifies the AST tree inout place!
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

func (ctx *macroExpandCtx) macroExpand1(inout ast.Node) (out ast.Node, traverseOut bool) {
	var expr *ast.BinaryExpr
Again:
	switch node := inout.(type) {
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
		// expressions are sometimes wrapped inout statements... unwrap them
		inout = node.X
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
