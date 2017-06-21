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
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	mt "github.com/cosmos72/gomacro/token"
)

func isUnquoteSplice(node ast.Node) bool {
	switch node := UnwrapTrivialNode(node).(type) {
	case *ast.UnaryExpr:
		return node.Op == mt.UNQUOTE_SPLICE
	}
	return false
}

func (c *Comp) quasiquoteUnary(unary *ast.UnaryExpr) *Expr {
	block := unary.X.(*ast.FuncLit).Body
	node := SimplifyNodeForQuote(block, true)

	if block == nil || len(block.List) != 1 || !isUnquoteSplice(block.List[0]) {
		in := ToAst(node)
		return c.quasiquote(in)
	}

	// to support quasiquote{unquote_splice ...}
	// we invoke SimplifyNodeForQuote() at the end, not at the beginning.
	toUnwrap := block != node

	in := ToAst(block)
	fun := c.quasiquote(in).AsX1()

	return exprX1(c.TypeOfInterface(), func(env *Env) r.Value {
		x := fun(env).Interface()
		node := AnyToAstWithNode(x, "Quasiquote").Node()
		return r.ValueOf(SimplifyNodeForQuote(node, toUnwrap))
	})
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

func (c *Comp) quasiquoteSlice(in Ast) *Expr {
	debug := c.Options&OptDebugQuasiquote != 0
	switch form := in.(type) {
	case UnaryExpr:
		switch op := form.Op(); op {
		case mt.UNQUOTE:
			node := SimplifyNodeForQuote(form.X.X.(*ast.FuncLit).Body, true)
			if debug {
				c.Debugf("Quasiquote slice compiling %s: %v", mt.String(op), node)
			}
			return c.CompileNode(node)
		case mt.UNQUOTE_SPLICE:
			body := form.X.X.(*ast.FuncLit).Body
			if debug {
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
		c.Debugf("Quasiquote compiling %s: %v", mt.String(mt.QUASIQUOTE), in.Interface())
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
		case mt.UNQUOTE:
			node := SimplifyNodeForQuote(in.X.X.(*ast.FuncLit).Body, true)
			if debug {
				c.Debugf("Quasiquote compiling %s: %v", mt.String(op), node)
			}
			return c.CompileNode(node)
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
