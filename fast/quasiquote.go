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
	mp "github.com/cosmos72/gomacro/parser"
	mt "github.com/cosmos72/gomacro/token"
)

func (c *Comp) QuasiquoteUnary(unary *ast.UnaryExpr) ast.Node {
	block := unary.X.(*ast.FuncLit).Body

	// we invoke SimplifyNodeForQuote() at the end, not at the beginning.
	// reason: to support quasiquote{unquote_splice ...}
	toUnwrap := block != SimplifyNodeForQuote(block, true)

	in := ToAst(block)
	out, _ := c.Quasiquote(in)
	node := ToNode(out)
	return SimplifyNodeForQuote(node, toUnwrap)
}

// Quasiquote expands ~quasiquote, if Ast starts with it
func (c *Comp) Quasiquote(in Ast) (Ast, bool) {
	switch form := in.(type) {
	case UnaryExpr:
		if form.Op() == mt.QUASIQUOTE {
			body := form.X.X.(*ast.FuncLit).Body
			return c.quasiquote(ToAst(body)), true
		}
	}
	return in, false
}

func concatAst(a, b AstWithSlice) {
	if b != nil {
		n := b.Size()
		for i := 0; i < n; i++ {
			a.Append(b.Get(i))
		}
	}
}

func quote(node ast.Node) *ast.UnaryExpr {
	unary, _ := mp.MakeQuote(nil, mt.QUOTE, node.Pos(), node)
	return unary
}

func (c *Comp) quasiquoteSlice(in Ast) AstWithSlice {
	switch form := in.(type) {
	case UnaryExpr:
		switch form.Op() {
		case mt.UNQUOTE:
			node := SimplifyNodeForQuote(form.X.X.(*ast.FuncLit).Body, true)
			return NodeSlice{X: []ast.Node{node}}
		case mt.UNQUOTE_SPLICE:
			body := form.X.X.(*ast.FuncLit).Body
			return BlockStmt{X: body}
		}
	}
	form := c.quasiquote(in)
	return AstSlice{X: []Ast{form}}
}

// quasiquote expands the contents of a ~quasiquote
func (c *Comp) quasiquote(in Ast) Ast {
	switch form := in.(type) {
	case AstWithSlice:
		ni := form.Size()

		out := form.New().(AstWithSlice)
		for i := 0; i < ni; i++ {
			concatAst(out, c.quasiquoteSlice(form.Get(i)))
		}
		return out
	case UnaryExpr:
		switch form.Op() {
		case mt.UNQUOTE:
			node := SimplifyNodeForQuote(form.X.X.(*ast.FuncLit).Body, true)
			return ToAst(node)
		case mt.UNQUOTE_SPLICE:
			c.Pos = form.X.Pos()
			c.Errorf("quasiquote: cannot %s in single-node context: %v", mt.String(form.Op()), form.X)
		}
	case AstWithNode:
		return ToAst(quote(form.Node()))
	default:
		x := in.Interface()
		c.Errorf("quasiquote: unsupported node type: %v <%v>", x, r.TypeOf(x))
	}
	return nil
}
