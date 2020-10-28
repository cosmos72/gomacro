/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * macroexpand.go
 *
 *  Created on Jun 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base"
	etoken "github.com/cosmos72/gomacro/go/etoken"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// MacroExpandNodeCodewalk traverses the whole AST tree using pre-order traversal,
// and replaces each node with the result of MacroExpandNode(node).
// It implements the macroexpansion phase
func (c *Comp) MacroExpandNodeCodewalk(in ast.Node) (out ast.Node, anythingExpanded bool) {
	if in == nil {
		return nil, false
	}
	var form Ast = ToAst(in)
	form, anythingExpanded = c.MacroExpandCodewalk(form)
	out = ToNode(form)
	// if !anythingExpanded {
	//    c.Debugf("MacroExpand1() nothing to expand: %v <%v>", out, r.TypeOf(out))
	//}
	return out, anythingExpanded
}

// MacroExpandCodewalk traverses the whole AST tree using pre-order traversal,
// and replaces each node with the result of MacroExpand(node).
// It implements the macroexpansion phase
func (c *Comp) MacroExpandCodewalk(in Ast) (out Ast, anythingExpanded bool) {
	return c.macroExpandCodewalk(in, 0)
}

func (c *Comp) macroExpandCodewalk(in Ast, quasiquoteDepth int) (out Ast, anythingExpanded bool) {
	if in == nil || in.Size() == 0 {
		return in, false
	}
	debug := c.Options&base.OptDebugMacroExpand != 0
	if quasiquoteDepth <= 0 {
		if debug {
			c.Debugf("MacroExpandCodewalk: qq = %d, macroexpanding %v", quasiquoteDepth, in.Interface())
		}
		in, anythingExpanded = c.MacroExpand(in)
	}
	if in != nil {
		in = base.UnwrapTrivialAst(in)
	}
	if in == nil {
		return in, anythingExpanded
	}
	saved := in

	if expr, ok := in.(UnaryExpr); ok {
		op := expr.X.Op
		switch op {
		case etoken.MACRO:
			break
		case etoken.QUOTE:
			// QUOTE prevents macroexpansion only if found outside any QUASIQUOTE
			if quasiquoteDepth == 0 {
				return saved, anythingExpanded
			}
		case etoken.QUASIQUOTE:
			// extract the body of QUASIQUOTE
			quasiquoteDepth++
		case etoken.UNQUOTE, etoken.UNQUOTE_SPLICE:
			// extract the body of UNQUOTE or UNQUOTE_SPLICE
			quasiquoteDepth--
		default:
			goto Recurse
		}
		inChild := base.UnwrapTrivialAst(in.Get(0).Get(1))
		outChild, expanded := c.macroExpandCodewalk(inChild, quasiquoteDepth)
		if op == etoken.MACRO {
			return outChild, expanded
		}
		out := in
		if expanded {
			out = base.MakeQuote2(expr, outChild.(AstWithNode))
		}
		return out, expanded
	}
Recurse:
	if in == nil {
		return saved, anythingExpanded
	}
	if debug {
		c.Debugf("MacroExpandCodewalk: qq = %d, recursing on %v", quasiquoteDepth, in)
	}
	out = in.New()
	n := in.Size()
	if outSlice, appendable := out.(AstWithSlice); appendable {
		// New() returns zero-length slice... resize it
		for outSlice.Size() < n {
			outSlice = outSlice.Append(nil)
		}
		out = outSlice
	}
	for i := 0; i < n; i++ {
		child := base.UnwrapTrivialAst(in.Get(i))
		if child != nil {
			expanded := false
			if child.Size() != 0 {
				child, expanded = c.macroExpandCodewalk(child, quasiquoteDepth)
			}
			if expanded {
				anythingExpanded = true
			}
		}
		out.Set(i, child)
	}
	if debug {
		c.Debugf("MacroExpandCodewalk: qq = %d, expanded to %v", quasiquoteDepth, out)
	}
	return out, anythingExpanded
}

// MacroExpandNode repeatedly invokes MacroExpandNode1
// as long as the node represents a macro call.
// it returns the resulting node.
func (c *Comp) MacroExpandNode(in ast.Node) (out ast.Node, everExpanded bool) {
	if in == nil {
		return nil, false
	}
	inAst := ToAst(in)
	outAst, everExpanded := c.MacroExpand(inAst)
	out = ToNode(outAst)
	// if !everExpanded {
	//    c.Debugf("MacroExpand1() not a macro: %v <%v>", out, r.TypeOf(out))
	//}
	return out, everExpanded
}

// MacroExpand repeatedly invokes MacroExpand
// as long as the node represents a macro call.
// it returns the resulting node.
func (c *Comp) MacroExpand(form Ast) (out Ast, everExpanded bool) {
	var expanded bool
	for {
		form, expanded = c.MacroExpand1(form)
		if !expanded {
			return form, everExpanded
		}
		everExpanded = true
	}
}

// if node represents a macro call, MacroExpandNode1 executes it
// and returns the resulting node.
// Otherwise returns the node argument unchanged
func (c *Comp) MacroExpandNode1(in ast.Node) (out ast.Node, expanded bool) {
	if in == nil {
		return nil, false
	}
	var form Ast = ToAst(in)
	form, expanded = c.MacroExpand1(form)
	out = ToNode(form)
	// if !expanded {
	//    c.Debugf("MacroExpandNode1: not a macro: %v <%v>", out, r.TypeOf(out))
	//}
	return out, expanded
}

func (c *Comp) extractMacroCall(form Ast) Macro {
	form = base.UnwrapTrivialAst(form)
	switch form := form.(type) {
	case Ident:
		sym := c.TryResolve(form.X.Name)
		if sym != nil && sym.Bind.Desc.Class() == ConstBind && sym.Type != nil && sym.Type.Kind() == r.Struct {
			switch value := sym.Value.(type) {
			case Macro:
				if c.Options&base.OptDebugMacroExpand != 0 {
					c.Debugf("MacroExpand1: found macro: %v", form.X.Name)
				}
				return value
			}
		}
	}
	return Macro{}
}

// if node represents a macro call, MacroExpandNode1 executes it
// and returns the resulting node.
// Otherwise returns the node argument unchanged
func (c *Comp) MacroExpand1(in Ast) (out Ast, expanded bool) {
	if in == nil {
		return nil, false
	}
	// unwrap trivial nodes: DeclStmt, ParenExpr, ExprStmt
	in = base.UnwrapTrivialAstKeepBlocks(in)
	ins, ok := in.(AstWithSlice)
	if !ok {
		return in, false
	}
	debug := c.Options&base.OptDebugMacroExpand != 0
	if debug {
		c.Debugf("MacroExpand1: found list: %v", ins.Interface())
	}
	n := ins.Size()
	outs := ins.New().(AstWithSlice)

	// since macro calls are sequences of statements,
	// we must scan the whole list,
	// consume it as needed by the macros we find,
	// and build a new list accumulating the results of macroexpansion
	for i := 0; i < n; i++ {
		elt := ins.Get(i)
		macro := c.extractMacroCall(elt)
		if macro.closure == nil {
			outs = outs.Append(elt)
			continue
		}
		argn := macro.argNum
		leftn := n - i - 1
		var args []xr.Value
		if argn > leftn {
			args := make([]xr.Value, leftn+1) // include the macro itself
			for j := 0; j <= leftn; j++ {
				args[j] = xr.ValueOf(ins.Get(i + j).Interface())
			}
			c.Errorf("not enough arguments for macroexpansion of %v: expecting %d, found %d", args, macro.argNum, leftn)
			return in, false
		}
		if debug {
			c.Debugf("MacroExpand1: found macro call %v at %d-th position of %v", elt.Interface(), i, ins.Interface())
		}
		// wrap each ast.Node into a reflect.Value
		args = make([]xr.Value, argn)
		for j := 0; j < argn; j++ {
			args[j] = xr.ValueOf(ToNode(ins.Get(i + j + 1)))
		}
		// invoke the macro
		results := macro.closure(args)
		if debug {
			c.Debugf("MacroExpand1: macro expanded to: %v", results)
		}
		// a macro expansion can return multiple values.
		// each value can be:
		// * ast.Node or something that implements ast.Node
		// * slice of: ast.Node or something that implements ast.Node
		// * Ast or something that implements Ast
		for _, result := range results {
			if !result.IsValid() || !result.CanInterface() {
				c.Warnf("MacroExpand1: cannot extract interface{} from reflect.Value result: %v", result)
				continue
			}
			if result == None {
				continue
			}
			res := AnyToAst(result.Interface(), "macroexpansion")
			switch res := res.(type) {
			case AstWithSlice:
				n := res.Size()
				for i := 0; i < n; i++ {
					outs = outs.Append(res.Get(i))
				}
			case Ast:
				outs = outs.Append(res)
			case nil:
			default:
				c.Warnf("MacroExpand1: cannot convert result to Ast: %v", result)
				continue
			}
		}
		i += argn
		expanded = true
	}
	if !expanded {
		return in, false
	}
	if outs.Size() == 0 {
		return EmptyStmt{&ast.EmptyStmt{}}, true
	}
	return base.UnwrapTrivialAst(outs), true
}
