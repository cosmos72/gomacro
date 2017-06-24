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
		fun := c.quasiquote1(in, 1, true).AsX1()

		return exprX1(c.TypeOfInterface(), func(env *Env) r.Value {
			x := fun(env).Interface()
			form := AnyToAst(x, "Quasiquote")
			if form, ok := form.(AstWithNode); ok {
				return r.ValueOf(SimplifyNodeForQuote(form.Node(), toUnwrap))
			}
			return r.ValueOf(form.Interface())
		})
	}
	return c.quasiquote1(ToAst(node), 1, true)
}

// Quasiquote expands and compiles ~quasiquote, if Ast starts with it
func (c *Comp) Quasiquote(in Ast) *Expr {
	switch form := in.(type) {
	case UnaryExpr:
		if form.Op() == mt.QUASIQUOTE {
			body := form.X.X.(*ast.FuncLit).Body
			return c.quasiquote1(ToAst(body), 1, true)
		}
	}
	return c.Compile(in)
}

func (c *Comp) quasiquote1(in Ast, depth int, can_splice bool) *Expr {
	expr, _ := c.quasiquote(in, depth, can_splice)
	return expr
}

// quasiquote expands and compiles the contents of a ~quasiquote
func (c *Comp) quasiquote(in Ast, depth int, can_splice bool) (*Expr, bool) {
	debug := c.Options&OptDebugQuasiquote != 0
	var label string
	if can_splice {
		label = " splice"
	}
	if debug {
		c.Debugf("Quasiquote[%d]%s expanding %s: %v <%v>", depth, label, mt.String(mt.QUASIQUOTE), in.Interface(), r.TypeOf(in.Interface()))
	}

	switch in := in.(type) {
	case AstWithSlice:
		n := in.Size()
		funs := make([]func(*Env) r.Value, 0, n)
		splices := make([]bool, 0, n)
		positions := make([]token.Position, 0, n)
		for i := 0; i < n; i++ {
			if form := in.Get(i); form != nil {
				form = UnwrapTrivialAst(form)
				expr, splice := c.quasiquote(form, depth, true)
				fun := expr.AsX1()
				if fun == nil {
					c.Warnf("Quasiquote[%d]%s: node expanded to nil: %v <%v>", depth, label, form.Interface(), r.TypeOf(form.Interface()))
					continue
				}
				funs = append(funs, fun)
				splices = append(splices, splice)
				var position token.Position
				if form, ok := form.(AstWithNode); ok {
					position = c.Fileset.Position(form.Node().Pos())
				}
				positions = append(positions, position)
			}
		}
		form := in.New().(AstWithSlice)

		typ := c.TypeOf(in.Interface()) // extract the concrete type implementing ast.Node
		rtype := typ.ReflectType()

		return exprX1(typ, func(env *Env) r.Value {
			out := form.New().(AstWithSlice)
			for i, fun := range funs {
				if splices[i] {
					xs := AnyToAstWithSlice(fun(env).Interface(), positions[i])
					n := xs.Size()
					for j := 0; j < n; j++ {
						if xj := xs.Get(j); xj != nil {
							out.Append(xj)
						}
					}
				} else {
					x := fun(env).Interface()
					out.Append(AnyToAst(x, positions[i]))
				}
			}
			return r.ValueOf(out.Interface()).Convert(rtype)
		}), false
	case UnaryExpr:
		unary := in.X
		switch op := unary.Op; op {
		case mt.QUOTE, mt.QUASIQUOTE, mt.UNQUOTE, mt.UNQUOTE_SPLICE:
			form := UnwrapTrivialAst(ToAst(unary.X.(*ast.FuncLit).Body))
			node := form.Interface()

			if op == mt.QUASIQUOTE {
				depth++
			} else if op == mt.UNQUOTE || op == mt.UNQUOTE_SPLICE {
				depth--
			}
			if depth <= 0 {
				if debug {
					c.Debugf("Quasiquote[%d]%s compiling %s: %v <%v>", depth, label, mt.String(op), node, r.TypeOf(node))
				}
				return c.Compile(form), op == mt.UNQUOTE_SPLICE
			}
			fun := c.quasiquote1(form, depth, true).AsX1()
			if fun == nil {
				c.Warnf("Quasiquote[%d]%s: node expanded to nil: %v <%v>", depth, label, node, r.TypeOf(node))
			}
			var position token.Position
			if node, ok := node.(ast.Node); ok {
				position = c.Fileset.Position(node.Pos())
			}
			return exprX1(c.Universe.FromReflectType(rtypeOfUnaryExpr), func(env *Env) r.Value {
				var node ast.Node
				if fun != nil {
					node = AnyToAstWithNode(fun(env).Interface(), position).Node()
				}
				ret, _ := mp.MakeQuote(nil, op, token.NoPos, node)
				return r.ValueOf(ret)
			}), false
		}
	}

	// Ast can still be a tree: just not a resizeable one, so support ~unquote but not ~unquote_splice
	in, ok := in.(AstWithNode)
	if !ok {
		x := in.Interface()
		c.Errorf("Quasiquote: unsupported node type, expecting AstWithNode or AstWithSlice: %v <%v>", x, r.TypeOf(x))
		return nil, false
	}
	if debug {
		c.Debugf("Quasiquote[%d] recursing: %v <%v>", depth, in.Interface(), r.TypeOf(in.Interface()))
	}
	form := in.New().(AstWithNode) // we must NOT retain input argument, so clone it
	n := in.Size()
	typ := c.TypeOf(in.Interface()) // extract the concrete type implementing ast.Node
	rtype := typ.ReflectType()

	if n == 0 {
		return exprX1(typ, func(env *Env) r.Value {
			return r.ValueOf(form.New().Interface()).Convert(rtype)
		}), false
	}
	funs := make([]func(*Env) r.Value, n)
	positions := make([]token.Position, n)
	for i := 0; i < n; i++ {
		if form := in.Get(i); form != nil {
			form = UnwrapTrivialAst(form)
			fun := c.quasiquote1(form, depth, false).AsX1()
			if fun == nil {
				c.Warnf("Quasiquote[%d]: node expanded to nil: %v", depth, form.Interface())
				continue
			}
			funs[i] = fun
			if form, ok := form.(AstWithNode); ok {
				positions[i] = c.Fileset.Position(form.Node().Pos())
			}
		}
	}

	return exprX1(typ, func(env *Env) r.Value {
		out := form.New().(AstWithNode)
		for i, fun := range funs {
			if fun != nil {
				x := fun(env).Interface()
				out.Set(i, AnyToAst(x, positions[i]))
			}
		}
		return r.ValueOf(out.Interface()).Convert(rtype)
	}), false
}
