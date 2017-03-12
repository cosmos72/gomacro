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
	"go/token"
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
				return MakeQuote2(in, expansion.(AstWithNode))
			case mt.UNQUOTE:
				if depth <= 1 {
					y := env.evalUnquote(in)
					return AnyToAst(y, "unquote")
				} else {
					// equivalent to ToAst(form.p.X.(*ast.FuncLit).Body)
					toexpand := in.Get(0).Get(1)
					env.debugQuasiQuote("recursing inside UNQUOTE", depth-1, canSplice, toexpand.Interface())
					expansion := env.evalQuasiquoteAst(toexpand, depth-1)
					return MakeQuote2(in, expansion.(AstWithNode))
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
				child = MakeQuote2(child, expansion.(AstWithNode))
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
					child = MakeQuote2(child, expansion.(AstWithNode))
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
func MakeQuote(form UnaryExpr) (UnaryExpr, BlockStmt) {
	expr, block := (*mp.Parser)(nil).MakeQuote(form.p.Op, form.p.OpPos, nil)
	return UnaryExpr{expr}, BlockStmt{block}
}

// MakeQuote3 invokes parser.MakeQuote() and wraps the resulting ast.Node,
// which represents quote{<form>}, into an Ast struct
func MakeQuote2(form UnaryExpr, toQuote AstWithNode) UnaryExpr {
	var node ast.Node
	if toQuote != nil {
		node = toQuote.Node()
	}
	expr, _ := (*mp.Parser)(nil).MakeQuote(form.p.Op, form.p.OpPos, node)
	return UnaryExpr{expr}
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
	form, anythingExpanded = env.MacroExpandAstCodewalk(form)
	out = ToNode(form)
	// if !anythingExpanded {
	//    env.Debugf("MacroExpand1() nothing to expand: %v <%v>", out, r.TypeOf(out))
	//}
	return out, anythingExpanded
}

func (env *Env) MacroExpandAstCodewalk(in Ast) (out Ast, anythingExpanded bool) {
	return env.macroExpandAstCodewalk(in, 0)
}

func (env *Env) macroExpandAstCodewalk(in Ast, quasiquoteDepth int) (out Ast, anythingExpanded bool) {
	if in == nil || in.Size() == 0 {
		return in, false
	}
	if quasiquoteDepth <= 0 {
		if env.Options&OptDebugMacroExpand != 0 {
			env.Debugf("MacroExpandCodewalk: qq = %d, macroexpanding %v", quasiquoteDepth, in.Interface())
		}
		in, anythingExpanded = env.macroExpandAst(in)
	}
	if in != nil {
		in = unwrapTrivialAst(in)
	}
	if in == nil {
		return in, anythingExpanded
	}
	saved := in

	if expr, ok := in.(UnaryExpr); ok {
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
		inChild := unwrapTrivialAst(in.Get(0).Get(1))
		outChild, expanded := env.macroExpandAstCodewalk(inChild, quasiquoteDepth)
		if !expanded {
			return in, false
		}
		return MakeQuote2(expr, outChild.(AstWithNode)), true
	}
Recurse:
	if in == nil {
		return saved, anythingExpanded
	}
	if env.Options&OptDebugMacroExpand != 0 {
		env.Debugf("MacroExpandCodewalk: qq = %d, recursing on %v", quasiquoteDepth, in)
	}
	out = in.New()
	n := in.Size()
	if outSlice, canAppend := out.(AstWithSlice); canAppend {
		// New() returns zero-length slice... resize it
		for i := 0; i < n; i++ {
			outSlice = outSlice.Append(nil)
		}
		out = outSlice
	}
	for i := 0; i < n; i++ {
		child := unwrapTrivialAst(in.Get(i))
		if child != nil {
			expanded := false
			if child.Size() != 0 {
				child, expanded = env.macroExpandAstCodewalk(child, quasiquoteDepth)
			}
			if expanded {
				anythingExpanded = true
			}
		}
		out.Set(i, child)
	}
	if env.Options&OptDebugMacroExpand != 0 {
		env.Debugf("MacroExpandCodewalk: qq = %d, expanded to %v", quasiquoteDepth, out)
	}
	return out, anythingExpanded
}

// MacroExpand repeatedly invokes MacroExpand1
// as long as the node represents a macro call.
// it returns the resulting node.
func (env *Env) MacroExpand(in ast.Node) (out ast.Node, everExpanded bool) {
	if in == nil {
		return nil, false
	}
	inAst := ToAst(in)
	outAst, everExpanded := env.macroExpandAst(inAst)
	out = ToNode(outAst)
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
func (env *Env) extractMacroCall(form Ast) Macro {
	form = unwrapTrivialAst(form)
	switch form := form.(type) {
	case Ident:
		bind, found := env.resolveIdentifier(form.p)
		if found && bind.Kind() == r.Struct {
			switch value := bind.Interface().(type) {
			case Macro:
				return value
			}
		}
	}
	return Macro{}
}

func (env *Env) macroExpandAstOnce(in Ast) (out Ast, expanded bool) {
	if in == nil {
		return nil, false
	}
	// unwrap trivial nodes: DeclStmt, ParenExpr, ExprStmt
	in = unwrapTrivialAst(in)
	ins, ok := in.(AstWithSlice)
	if !ok {
		return in, false
	}
	if env.Options&OptDebugMacroExpand != 0 {
		env.Debugf("MacroExpand1: found list: %v", ins.Interface())
	}
	outs := ins.New().(AstWithSlice)
	n := ins.Size()

	// since macro calls are sequences of statements,
	// we must scan the whole list,
	// consume it as needed by the macros we find,
	// and build a new list accumulating the results of macroexpansion
	for i := 0; i < n; i++ {
		elt := ins.Get(i)
		macro := env.extractMacroCall(elt)
		if macro.Closure == nil {
			outs = outs.Append(elt)
			continue
		}
		argn := macro.ArgNum
		leftn := n - i - 1
		var args []r.Value
		if argn > leftn {
			args := make([]r.Value, leftn+1) // include the macro itself
			for j := 0; j <= leftn; j++ {
				args[j] = r.ValueOf(ins.Get(i + j).Interface())
			}
			env.Errorf("not enough arguments for macroexpansion of %v: expecting %d, found %d", args, macro.ArgNum, leftn)
			return in, false
		}
		if env.Options&OptDebugMacroExpand != 0 {
			env.Debugf("MacroExpand1: found macro call %v at %d-th position of %v", elt.Interface(), i, ins.Interface())
		}
		// wrap each ast.Node into a reflect.Value
		args = make([]r.Value, argn)
		for j := 0; j < argn; j++ {
			args[j] = r.ValueOf(ToNode(ins.Get(i + j + 1)))
		}
		// invoke the macro
		results := macro.Closure(args)
		if env.Options&OptDebugMacroExpand != 0 {
			env.Debugf("MacroExpand1: macro expanded to: %v", results)
		}
		var out Ast
		switch len(results) {
		default:
			args = append([]r.Value{r.ValueOf(elt.Interface())}, args...)
			env.Warnf("macroexpansion returned %d values, using only the first one: %v %v returned %v",
				len(results), args, results)
			fallthrough
		case 1:
			any := results[0].Interface()
			if any != nil {
				out = AnyToAst(any, "macroexpansion")
				break
			}
			fallthrough
		case 0:
			// do not insert nil nodes... they would wreak havok, convert them to the identifier nil
			out = Ident{&ast.Ident{Name: "nil"}}
		}
		outs = outs.Append(out)
		i += argn
		expanded = true
	}
	if !expanded {
		return in, false
	}
	if outs.Size() == 0 {
		return EmptyStmt{&ast.EmptyStmt{}}, true
	}
	return unwrapTrivialAst(outs), true
}
