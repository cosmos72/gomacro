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
	"fmt"
	"go/ast"
	_ "go/token"
	r "reflect"

	mp "github.com/cosmos72/gomacro/parser"
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
func (env *Env) evalQuasiquote(node *ast.BlockStmt) ast.Node {
	// we invoke simplifyNodeForQuote() at the end, not at the beginning.
	// reason: to support quasiquote{unquote_splice ...}
	toUnwrap := node != simplifyNodeForQuote(node, true)

	in := ToAst(node)
	out := env.evalQuasiquoteAst(in, 1)
	ret := ToNode(out)
	return simplifyNodeForQuote(ret, toUnwrap)
}

// evalQuasiquoteAst evaluates the body of a quasiquote{} represented as Ast
// use unified API to traverse ast.Node... every other solution is a nightmare
func (env *Env) evalQuasiquoteAst(inout Ast, depth int) Ast {
	if inout == nil {
		return nil
	}
	withSlice, canSplice := inout.(AstWithSlice)
	form := inout
	env.debugQuasiQuote("evaluating", depth, canSplice, form.Interface())
	form = unwrapTrivialAst(form) // drill through DeclStmt, ExprStmt, ParenExpr
	if form == nil || form.Size() == 0 {
		return inout
	}

	if !canSplice {
		if form, ok := form.(UnaryExpr); ok {
			switch form.Op() {
			case mt.QUASIQUOTE:
				// equivalent to ToAst(form.p.X.(*ast.FuncLit).Body)
				toexpand := form.Get(0).Get(1)
				expansion := env.evalQuasiquoteAst(toexpand, depth+1)
				form.Get(0).Set(1, expansion)
				return form
			case mt.UNQUOTE:
				if depth <= 1 {
					y := env.evalUnquote(form)
					return AnyToAst(y, "unquote")
				} else {
					// equivalent to ToAst(form.p.X.(*ast.FuncLit).Body)
					toexpand := form.Get(0).Get(1)
					expansion := env.evalQuasiquoteAst(toexpand, depth-1)
					form.Get(0).Set(1, expansion)
					return form
				}
			case mt.UNQUOTE_SPLICE:
				y := form.Interface()
				env.Errorf("quasiquote: cannot splice in single-statement context: %v <%v>", y, r.TypeOf(y))
				return nil
			}
		}

		ni := form.Size()
		for i := 0; i < ni; i++ {
			// general case: recurse on child
			child := form.Get(i)
			child = env.evalQuasiquoteAst(child, depth)
			form.Set(i, child)
		}
		// we modified form destructively... return form, not inout!
		return form
	}

	ni := form.Size()
	ret := make([]Ast, 0, ni)
	for i := 0; i < ni; i++ {
		// drill through DeclStmt, ExprStmt, ParenExpr
		child := unwrapTrivialAst(form.Get(i))
		switch child := child.(type) {
		case UnaryExpr:
			switch child.Op() {
			case mt.QUASIQUOTE:
				// equivalent to ToAst(form.p.X.(*ast.FuncLit).Body)
				toexpand := child.Get(0).Get(1)
				expansion := env.evalQuasiquoteAst(toexpand, depth+1)
				child.Get(0).Set(1, expansion)
				ret = append(ret, child)
				goto PrintDebug
			case mt.UNQUOTE, mt.UNQUOTE_SPLICE:
				// complication: in Common Lisp, the right-most unquote pairs with the left-most comma!
				// we implement the same mechanics, so we must drill down to the last unquote/unquote_splice
				// and, for unquote_splice, create a copy of the unquote/unquote_splice stack for each result.
				// Example:
				//   x:=quote{7; 8}
				//   quasiquote{quasiquote{1; unquote{2}; unquote{unquote_splice{x}}}}
				// must return
				//   quasiquote{1; unquote{2}; unquote{7}; unquote{8}}
				lastUnquote, unquoteDepth := env.descendNestedUnquotes(child)

				env.debugQuasiQuote(fmt.Sprintf("found %s (unquoteDepth = %d)", mt.String(lastUnquote.Op()), unquoteDepth),
					depth, canSplice, child)

				op := lastUnquote.Op()
				if unquoteDepth > depth {
					env.Errorf("%s not inside quasiquote: %v <%v>", mt.String(op), lastUnquote, r.TypeOf(lastUnquote))
					return nil
				} else if unquoteDepth < depth {
					toexpand := child.Get(0).Get(1)
					expansion := env.evalQuasiquoteAst(toexpand, depth-1)
					child.Get(0).Set(1, expansion)
					ret = append(ret, child)
				} else {
					toInsert := AnyToAst(env.evalUnquote(lastUnquote), mt.String(op))
					if op == mt.UNQUOTE {
						stack := duplicateNestedUnquotes(child, unquoteDepth-1, toInsert)
						ret = append(ret, stack)
					} else {
						toSplice := ToAstWithSlice(toInsert, "unquote_splice")
						nj := toSplice.Size()
						for j := 0; j < nj; j++ {
							stack := duplicateNestedUnquotes(child, unquoteDepth-1, toSplice.Get(j))
							ret = append(ret, stack)
						}
					}
				}
				goto PrintDebug
			}
		}
		// general case: recurse on child
		child = env.evalQuasiquoteAst(child, depth)
		ret = append(ret, child)
	PrintDebug:
		env.debugQuasiQuote("accumulated", depth, canSplice, ret)
	}
	withSlice.Slice(0, 0)
	for _, node := range ret {
		withSlice.Append(node)
	}
	return withSlice
}

// unwrapTrivialAst extract the content from ParenExpr, ExprStmt, DeclStmt:
// such nodes are trivial wrappers for their contents
func unwrapTrivialAst(form Ast) Ast {
	for {
		switch form.(type) {
		case ParenExpr, ExprStmt, DeclStmt:
			form = form.Get(0)
		default:
			return form
		}
	}
}

// unwrapTrivialNode is equivalent to unwrapTrivialAst,
// except it works on ast.Node instead of our Ast
func unwrapTrivialNode(in ast.Node) ast.Node {
	for {
		switch node := in.(type) {
		case *ast.ParenExpr:
			in = node.X
		case *ast.ExprStmt:
			in = node.X
		case *ast.DeclStmt:
			in = node.Decl
		default:
			return in
		}
	}
}

func (env *Env) debugQuasiQuote(msg string, depth int, canSplice bool, x interface{}) {
	if env.Options&OptDebugQuasiquote != 0 {
		env.Debugf("quasiquote: %s (depth = %d, canSplice = %v)\n%v <%v>", msg, depth, canSplice, x, r.TypeOf(x))
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

func (env *Env) descendNestedUnquotes(unquote UnaryExpr) (lastUnquote UnaryExpr, depth int) {
	depth = 1
	for {
		form := unquote.Get(0).Get(1)
		form = unwrapTrivialAst(form)

		if form != nil && form.Size() == 1 {
			if block, ok := form.(BlockStmt); ok {
				form = unwrapTrivialAst(block.Get(0))
				if form != nil && form.Size() == 1 {
					if expr, ok := form.(UnaryExpr); ok {
						if op := expr.Op(); op == mt.UNQUOTE || op == mt.UNQUOTE_SPLICE {
							unquote = expr
							depth++
							continue
						}
					}
				}
			}
		}
		return unquote, depth
	}
}

func duplicateNestedUnquotes(src UnaryExpr, depth int, content Ast) Ast {
	if depth == 0 {
		return content
	}
	head, tail := MakeQuote(src)
	var form Ast = src

	for ; depth > 1; depth-- {
		form = form.Get(0).Get(1)
		form = unwrapTrivialAst(form)

		src = form.(UnaryExpr)
		expr, newTail := MakeQuote(src)
		tail.Append(expr)
		tail = newTail
	}
	tail.Append(content)
	return head
}

// MakeQuote invokes parser.MakeQuote() and wraps the resulting ast.Node,
// which represents quote{<form>}, into an Ast struct
func MakeQuote(form UnaryExpr) (UnaryExpr, BlockStmt) {
	expr, block := (*mp.Parser)(nil).MakeQuote(form.p.Op, form.p.OpPos, nil)
	return UnaryExpr{expr}, BlockStmt{block}
}

type macroExpandCtx struct {
	env *Env
}

// MacroExpandCodewalk traverses the whole AST tree using pre-order traversal,
// and replaces each node with the result of MacroExpand(node).
// It implements the macroexpansion phase
// Warning: it destructively modifies the ast.Node !
func (env *Env) MacroExpandCodewalk(in ast.Node) (out ast.Node, anythingExpanded bool) {
	if in == nil {
		return nil, false
	}
	var form Ast = ToAst(in)
	form, anythingExpanded = env.macroExpandAstCodewalk(form, 0)
	out = ToNode(form)
	// if !anythingExpanded {
	//    env.Debugf("MacroExpand1() nothing to expand: %v <%v>", out, r.TypeOf(out))
	//}
	return out, anythingExpanded
}

func (env *Env) macroExpandAstCodewalk(form Ast, quasiquoteDepth int) (out Ast, anythingExpanded bool) {
	if form == nil || form.Size() == 0 {
		return form, false
	}
	if quasiquoteDepth <= 0 {
		form, anythingExpanded = env.macroExpandAst(form)
	}
	if form != nil {
		form = unwrapTrivialAst(form)
	}
	if form == nil {
		return form, anythingExpanded
	}
	saved := form
	for expr, ok := form.(UnaryExpr); ok; {
		switch expr.p.Op {
		case mt.QUOTE:
			// QUOTE prevents macroexpansion only if found outside any QUASIQUOTE
			if quasiquoteDepth == 0 {
				return saved, anythingExpanded
			}
		case mt.QUASIQUOTE:
			// extract the body of QUASIQUOTE
			quasiquoteDepth++
		case mt.UNQUOTE, mt.UNQUOTE_SPLICE:
			// extract the body of UNQUOTE or UNQUOTE_SPLICE
			quasiquoteDepth--
		default:
			goto Recurse
		}
		temp := unwrapTrivialAst(form.Get(0).Get(1))
		if env.Options&OptDebugMacroExpand != 0 {
			env.Debugf("MacroExpandCodewalk: unwrapped %v to %v", form, temp)
		}
		form = temp
	}
Recurse:
	if form == nil {
		return saved, anythingExpanded
	}
	if env.Options&OptDebugMacroExpand != 0 {
		env.Debugf("MacroExpandCodewalk: recursing on %v", form)
	}
	n := form.Size()
	var expanded bool
	for i := 0; i < n; i++ {
		child := unwrapTrivialAst(form.Get(i))
		if child == nil || child.Size() == 0 {
			continue
		}
		child, expanded = env.macroExpandAstCodewalk(child, quasiquoteDepth)
		if expanded {
			anythingExpanded = true
			form.Set(i, child)
		}
	}
	return saved, anythingExpanded
}

// MacroExpand repeatedly invokes MacroExpand1
// as long as the node represents a macro call.
// it returns the resulting node.
func (env *Env) MacroExpand(in ast.Node) (out ast.Node, everExpanded bool) {
	if in == nil {
		return nil, false
	}
	var form Ast = ToAst(in)
	form, everExpanded = env.macroExpandAst(form)
	out = ToNode(form)
	// if !everExpanded {
	//    env.Debugf("MacroExpand1() not a macro: %v <%v>", out, r.TypeOf(out))
	//}
	return out, everExpanded
}

func (env *Env) macroExpandAst(form Ast) (out Ast, everExpanded bool) {
	var expanded bool
	for {
		form, expanded = env.macroExpandAstOnce(form)
		if !expanded {
			return form, everExpanded
		}
		everExpanded = true
	}
}

// if node represents a macro call, MacroExpand1 executes it
// and returns the resulting node.
// Otherwise returns the node argument unchanged
func (env *Env) MacroExpand1(in ast.Node) (out ast.Node, expanded bool) {
	if in == nil {
		return nil, false
	}
	var form Ast = ToAst(in)
	form, expanded = env.macroExpandAstOnce(form)
	out = ToNode(form)
	// if !expanded {
	//    env.Debugf("MacroExpand1: not a macro: %v <%v>", out, r.TypeOf(out))
	//}
	return out, expanded
}

//
func (env *Env) extractMacroCall(stmt ast.Stmt) Macro {
	first := unwrapTrivialNode(stmt)
	switch first := first.(type) {
	case *ast.Ident:
		name := first.Name
		// we cannot use env.evalIdentifier() because it panics on failure
		for e := env; e != nil; e = e.Outer {
			if bind, exists := e.Binds[name]; exists {
				if bind.Kind() == r.Struct {
					switch value := bind.Interface().(type) {
					case Macro:
						return value
					}
				}
				break
			}
		}
	}
	return Macro{}
}

func (env *Env) macroExpandAstOnce(in Ast) (out Ast, expanded bool) {
	if in == nil {
		return nil, false
	}
	saved := in
	// unwrap trivial nodes: DeclStmt, ParenExpr, ExprStmt
	form := unwrapTrivialAst(in)
	block, ok := form.(BlockStmt)
	if !ok {
		return saved, false
	}
	if env.Options&OptDebugMacroExpand != 0 {
		env.Debugf("MacroExpand1: found block: %v", block)
	}
	list := block.p.List
	n := len(list)
	rets := make([]ast.Stmt, 0, n)

	// since macro calls are sequences of statements,
	// we must scan the whole statement list,
	// consume it as needed by the macros we find,
	// and build a new list accumulating the results of macroexpansion
	for i := 0; i < n; i++ {
		stmt := list[i]
		macro := env.extractMacroCall(stmt)
		if macro.Closure == nil {
			rets = append(rets, stmt)
			continue
		}
		argn := macro.ArgNum
		if argn > n-i-1 {
			env.Errorf("not enough arguments for macroexpansion of %v: expecting %d, found %d", list[i:], macro.ArgNum, n-i-1)
			return saved, false
		}
		if env.Options&OptDebugMacroExpand != 0 {
			env.Debugf("MacroExpand1: found macro call: %v", list[i:i+1+argn])
		}
		// wrap each ast.Stmt into a reflect.Value
		args := make([]r.Value, argn)
		for j := 0; j < argn; j++ {
			args[j] = r.ValueOf(list[i+j+1])
		}
		// invoke the macro
		results := macro.Closure(args)
		if env.Options&OptDebugMacroExpand != 0 {
			env.Debugf("MacroExpand1: macro expanded to: %v", results)
		}
		// convert and accumulate the results
		for _, result := range results {
			if result != None {
				ret := AnyToAst(result.Interface(), "macroexpansion")
				if ret == nil {
					// do not insert nil nodes... they would wreak havok, convert them to the identifier nil
					ret = Ident{&ast.Ident{Name: "nil"}}
				}
				rets = append(rets, ToStmt(ret))
			}
		}
		i += argn
		expanded = true
	}
	if !expanded {
		return saved, false
	}
	switch len(rets) {
	case 0:
		form = EmptyStmt{&ast.EmptyStmt{}}
	case 1:
		form = ToAst(unwrapTrivialNode(rets[0]))
	default:
		block.p.List = rets
		form = block
	}
	return form, true
}
