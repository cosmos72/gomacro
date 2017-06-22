/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 *
 * quasiquote.go
 *
 *  Created on Jun 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	mp "github.com/cosmos72/gomacro/parser"
	mt "github.com/cosmos72/gomacro/token"
)

func isUnquoteOrUnquoteSplice(node ast.Node) bool {
	op := UnwrapTrivialAst(ToAst(node)).Op()
	return op == mt.UNQUOTE || op == mt.UNQUOTE_SPLICE
}

var (
	rtypeOfNode      = r.TypeOf((*ast.Node)(nil)).Elem()
	rtypeOfUnaryExpr = r.TypeOf((*ast.UnaryExpr)(nil))
)

func (c *Comp) quasiquoteUnary(unary *ast.UnaryExpr) *Expr {
	block := unary.X.(*ast.FuncLit).Body
	node := SimplifyNodeForQuote(block, true)

	if block != nil && len(block.List) == 1 && isUnquoteOrUnquoteSplice(block.List[0]) {
		// to support quasiquote{unquote ...} and quasiquote{unquote_splice ...}
		// we invoke SimplifyNodeForQuote() at the end, not at the beginning.
		toUnwrap := block != node

		in := ToAst(block)
		fun := c.quasiquote(in).AsX1()

		return exprX1(c.Universe.FromReflectType(rtypeOfNode), func(env *Env) r.Value {
			x := fun(env).Interface()
			node := AnyToAstWithNode(x, "Quasiquote").Node()
			return r.ValueOf(SimplifyNodeForQuote(node, toUnwrap))
		})
	}

	in := ToAst(node)
	return c.quasiquote(in)
}

// Quasiquote expands and compiles ~quasiquote, if Ast starts with it
func (c *Comp) Quasiquote(in Ast) *Expr {
	switch form := in.(type) {
	case UnaryExpr:
		if form.Op() == mt.QUASIQUOTE {
			body := form.X.X.(*ast.FuncLit).Body
			return c.quasiquote(ToAst(body))
		}
	}
	return c.Compile(in)
}

func (c *Comp) quasiquoteSimple(unary *ast.UnaryExpr, label string) *Expr {
	node := SimplifyNodeForQuote(unary.X.(*ast.FuncLit).Body, true)
	op := unary.Op
	if c.Options&OptDebugQuasiquote != 0 {
		c.Debugf("Quasiquote%s compiling %s: %v", label, mt.String(op), node)
	}
	if op == mt.UNQUOTE {
		return c.CompileNode(node)
	}
	fun := c.quasiquote(ToAst(node)).AsX1()
	return exprX1(c.Universe.FromReflectType(rtypeOfUnaryExpr), func(env *Env) r.Value {
		ret, _ := mp.MakeQuote(nil, op, token.NoPos, fun(env).Interface().(ast.Node))
		return r.ValueOf(ret)
	})
}

func (c *Comp) quasiquoteSlice(in Ast) *Expr {
	switch form := in.(type) {
	case UnaryExpr:
		switch op := form.Op(); op {
		case mt.QUOTE, mt.QUASIQUOTE, mt.UNQUOTE:
			return c.quasiquoteSimple(form.X, " slice")

		case mt.UNQUOTE_SPLICE:
			body := form.X.X.(*ast.FuncLit).Body
			if c.Options&OptDebugQuasiquote != 0 {
				c.Debugf("Quasiquote slice compiling %s: %v", mt.String(op), body)
			}
			return c.CompileNode(body)
		}
	}
	return c.quasiquote(in)
}

// quasiquote expands and compiles the contents of a ~quasiquote
func (c *Comp) quasiquote(in Ast) *Expr {
	debug := c.Options&OptDebugQuasiquote != 0
	if debug {
		c.Debugf("Quasiquote expanding %s: %v", mt.String(mt.QUASIQUOTE), in.Interface())
	}
	switch in := in.(type) {
	case AstWithSlice:
		n := in.Size()
		funs := make([]func(*Env) r.Value, n)
		for i := 0; i < n; i++ {
			funs[i] = c.quasiquoteSlice(in.Get(i)).AsX1()
		}
		form := in.New().(AstWithSlice)

		typ := c.TypeOf(in.Interface()) // extract the concrete type implementing ast.Node
		rtype := typ.ReflectType()

		return exprX1(typ, func(env *Env) r.Value {
			out := form.New().(AstWithSlice)
			for _, fun := range funs {
				x := fun(env).Interface()
				out.Append(AnyToAst(x, "Quasiquote"))
			}
			return r.ValueOf(out.Interface()).Convert(rtype)
		})
	case UnaryExpr:
		switch op := in.Op(); op {
		case mt.QUOTE, mt.QUASIQUOTE, mt.UNQUOTE:
			return c.quasiquoteSimple(in.X, "")

		case mt.UNQUOTE_SPLICE:
			c.Pos = in.X.Pos()
			c.Errorf("Quasiquote: cannot %s in single-node context: %v", mt.String(in.Op()), in.X)
			return nil
		}
	}

	// Ast can still be a tree: just not a resizeable one, so support ~unquote but not ~unquote_splice
	in, ok := in.(AstWithNode)
	if !ok {
		x := in.Interface()
		c.Errorf("Quasiquote: unsupported node type, expecting AstWithNode or AstWithSlice: %v <%v>", x, r.TypeOf(x))
		return nil
	}
	if debug {
		c.Debugf("Quasiquote compiling: %v", in.Interface())
	}
	form := in.New().(AstWithNode) // we must NOT retain input argument, so clone it
	n := in.Size()
	typ := c.TypeOf(in.Interface()) // extract the concrete type implementing ast.Node
	rtype := typ.ReflectType()

	if n == 0 {
		return exprX1(typ, func(env *Env) r.Value {
			return r.ValueOf(form.New().Interface()).Convert(rtype)
		})
	}
	funs := make([]func(*Env) r.Value, n)
	for i := 0; i < n; i++ {
		funs[i] = c.quasiquote(in.Get(i)).AsX1()
	}

	return exprX1(typ, func(env *Env) r.Value {
		out := form.New().(AstWithNode)
		for i, fun := range funs {
			x := fun(env).Interface()
			out.Set(i, AnyToAst(x, "Quasiquote"))
		}
		return r.ValueOf(out.Interface()).Convert(rtype)
	})
}
