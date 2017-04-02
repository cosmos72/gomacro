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
 * quasiquote.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
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
func (env *Env) evalQuasiquoteAst(in Ast, depth int) (out Ast) {
	if in == nil {
		return nil
	}
	inSlice, canSplice := in.(AstWithSlice)
	env.debugQuasiQuote("evaluating", depth, canSplice, in.Interface())
	if !canSplice {
		in = unwrapTrivialAst(in) // drill through DeclStmt, ExprStmt, ParenExpr, one-element BlockStmt
	}
	if in == nil || in.Size() == 0 {
		return in
	}

	if !canSplice {
		if in, ok := in.(UnaryExpr); ok {
			switch in.Op() {
			case mt.QUASIQUOTE:
				// equivalent to ToAst(form.p.X.(*ast.FuncLit).Body)
				toexpand := in.Get(0).Get(1)
				env.debugQuasiQuote("recursing inside QUASIQUOTE", depth+1, canSplice, toexpand.Interface())
				expansion := env.evalQuasiquoteAst(toexpand, depth+1)
				return makeQuote2(in, expansion.(AstWithNode))
			case mt.UNQUOTE:
				if depth <= 1 {
					y := env.evalUnquote(in)
					return AnyToAst(y, "unquote")
				} else {
					// equivalent to ToAst(form.p.X.(*ast.FuncLit).Body)
					toexpand := in.Get(0).Get(1)
					env.debugQuasiQuote("recursing inside UNQUOTE", depth-1, canSplice, toexpand.Interface())
					expansion := env.evalQuasiquoteAst(toexpand, depth-1)
					return makeQuote2(in, expansion.(AstWithNode))
				}
			case mt.UNQUOTE_SPLICE:
				y := in.Interface()
				env.Errorf("quasiquote: cannot splice in single-statement context: %v <%v>", y, r.TypeOf(y))
				return nil
			}
		}

		out := in.New()
		ni := in.Size()
		for i := 0; i < ni; i++ {
			child := in.Get(i)
			if child == nil {
				env.debugQuasiQuote("child is nil", depth, canSplice, child)
			} else {
				env.debugQuasiQuote("general case: recurse on child", depth, canSplice, child.Interface())
				child = env.evalQuasiquoteAst(child, depth)
			}
			out.Set(i, child)
		}
		return out
	}

	outSlice := inSlice.New().(AstWithSlice)
	ni := inSlice.Size()
	for i := 0; i < ni; i++ {
		// drill through DeclStmt, ExprStmt, ParenExpr
		child := unwrapTrivialAstKeepBlocks(inSlice.Get(i))
		switch child := child.(type) {
		case UnaryExpr:
			switch child.Op() {
			case mt.QUASIQUOTE:
				// equivalent to ToAst(form.p.X.(*ast.FuncLit).Body)
				toexpand := child.Get(0).Get(1)
				env.debugQuasiQuote("recursing inside QUASIQUOTE", depth+1, canSplice, toexpand.Interface())
				expansion := env.evalQuasiquoteAst(toexpand, depth+1)
				child = makeQuote2(child, expansion.(AstWithNode))
				outSlice = outSlice.Append(child)
				goto Next
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

				op := lastUnquote.Op()

				env.debugQuasiQuote(fmt.Sprintf("inside %s, lastUnquote is %s (unquoteDepth = %d)",
					mt.String(child.Op()), mt.String(op), unquoteDepth), depth, canSplice, child)

				if unquoteDepth > depth {
					env.Errorf("%s not inside quasiquote: %v <%v>", mt.String(op), lastUnquote, r.TypeOf(lastUnquote))
					return nil
				} else if unquoteDepth < depth {
					toexpand := child.Get(0).Get(1)
					env.debugQuasiQuote(fmt.Sprintf("recursing inside %s, lastUnquote is %s", mt.String(child.Op()), mt.String(op)),
						depth-1, canSplice, toexpand.Interface())
					expansion := env.evalQuasiquoteAst(toexpand, depth-1)
					child = makeQuote2(child, expansion.(AstWithNode))
					outSlice = outSlice.Append(child)
				} else {
					env.debugQuasiQuote("calling unquote on", depth-unquoteDepth, canSplice, lastUnquote.Interface())
					toInsert := AnyToAst(env.evalUnquote(lastUnquote), mt.String(op))
					env.debugQuasiQuote("unquote returned", depth-unquoteDepth, canSplice, toInsert.Interface())
					if op == mt.UNQUOTE {
						stack := duplicateNestedUnquotes(child, unquoteDepth-1, toInsert)
						outSlice = outSlice.Append(stack)
					} else {
						toSplice := ToAstWithSlice(toInsert, "unquote_splice")
						nj := toSplice.Size()
						for j := 0; j < nj; j++ {
							stack := duplicateNestedUnquotes(child, unquoteDepth-1, toSplice.Get(j))
							outSlice = outSlice.Append(stack)
						}
					}
				}
				goto Next
			}
		}
		if child == nil {
			env.debugQuasiQuote("child is nil", depth, canSplice, child)
		} else {
			env.debugQuasiQuote("general case: recurse on child", depth, canSplice, child.Interface())
			child = env.evalQuasiquoteAst(child, depth)
		}
		outSlice = outSlice.Append(child)
	Next:
		env.debugQuasiQuote("accumulated", depth, canSplice, outSlice.Interface())
	}
	return outSlice
}

// unwrapTrivialAst extract the content from ParenExpr, ExprStmt, DeclStmt:
// such nodes are trivial wrappers for their contents
func unwrapTrivialAst(in Ast) Ast {
	return unwrapTrivialAst2(in, true)
}

func unwrapTrivialAstKeepBlocks(in Ast) Ast {
	return unwrapTrivialAst2(in, false)
}

func unwrapTrivialAst2(in Ast, unwrapTrivialBlockStmt bool) Ast {
	for {
		switch form := in.(type) {
		case BlockStmt:
			if !unwrapTrivialBlockStmt || form.Size() != 1 {
				return form
			}
			// a one-element block is trivial UNLESS it contains a declaration.
			// reason: the declaration alters its scope with new bindings.
			// unwrapping it would alters the OUTER scope.
			// i.e. { var x = foo() } and var x = foo() give different scopes
			// to the variable 'x' so they are not equivalent.
			//
			// same reasoning for { x := foo() } versus x := foo()
			child := form.Get(0)
			switch child := child.(type) {
			case DeclStmt:
				return in
			case AssignStmt:
				if child.Op() == token.DEFINE {
					return in
				}
			}
			// fmt.Printf("// debug: unwrapTrivialAst(block) unwrapping %#v <%T>\n\tto %#v <%T>\n", form.Interface(), form.Interface(), child.Interface(), child.Interface())
			in = child
		case ParenExpr, ExprStmt, DeclStmt:
			child := form.Get(0)
			// fmt.Printf("// debug: unwrapTrivialAst(1) unwrapped %#v <%T>\n\tto %#v <%T>\n", form.Interface(), form.Interface(), child.Interface(), child.Interface())
			in = child
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
	block := inout.X.X.(*ast.FuncLit).Body

	ret, extraValues := env.evalBlock(block)
	if len(extraValues) > 1 {
		env.warnf("unquote returned %d values, only the first one will be used: %v", len(extraValues), block)
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
	head, tail := makeQuote(src)
	var form Ast = src

	for ; depth > 1; depth-- {
		form = form.Get(0).Get(1)
		form = unwrapTrivialAst(form)

		src = form.(UnaryExpr)
		expr, newTail := makeQuote(src)
		// cheat: we know that BlockStmt.Append() always returns the receiver unmodified
		tail.Append(expr)
		tail = newTail
	}
	// cheat: we know that BlockStmt.Append() always returns the receiver unmodified
	tail.Append(content)
	return head
}

// MakeQuote invokes parser.MakeQuote() and wraps the resulting ast.Node,
// which represents quote{<form>}, into an Ast struct
func makeQuote(form UnaryExpr) (UnaryExpr, BlockStmt) {
	expr, block := mp.MakeQuote(nil, form.X.Op, form.X.OpPos, nil)
	return UnaryExpr{expr}, BlockStmt{block}
}

// MakeQuote2 invokes parser.MakeQuote() and wraps the resulting ast.Node,
// which represents quote{<form>}, into an Ast struct
func makeQuote2(form UnaryExpr, toQuote AstWithNode) UnaryExpr {
	var node ast.Node
	if toQuote != nil {
		node = toQuote.Node()
	}
	fmt.Printf("node   = %#v\n", node)
	fmt.Printf("form   = %#v\n", form)
	fmt.Printf("form.X = %#v\n", form.X)
	expr, _ := mp.MakeQuote(nil, form.X.Op, form.X.OpPos, node)
	return UnaryExpr{expr}
}
