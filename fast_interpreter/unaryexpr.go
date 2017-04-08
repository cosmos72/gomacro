/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * unaryexpr.go
 *
 *  Created on Apr 07, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
	"go/token"

	mt "github.com/cosmos72/gomacro/token"
)

func (c *Comp) UnaryExpr(node *ast.UnaryExpr) *Expr {
	x := c.Expr(node.X)
	if x.NumOut() == 0 {
		c.Errorf("operand returns no values, cannot use in unary expression: %v", node.X)
	}

	isConst := x.Const()
	var z *Expr

	switch node.Op {
	case token.ADD:
		z = c.UnaryPlus(node, x) // only checks x type, returns x itself
	/*
		case token.SUB:
			z = c.UnaryMinus(node, x)
		case token.NOT:
			z = c.UnaryNot(node, x)
		case token.XOR:
			z = c.UnaryXor()
		case token.AND:
			z = c.AddressOf(node, x)
		case token.ARROW:
			z = c.UnaryRecv(node, x)
		// case token.MUL: // not seen, the parser produces *ast.StarExpr instead
	*/
	default:
		return c.unimplementedUnaryExpr(node, x)
	}
	if isConst {
		// constant propagation
		z.EvalConst(CompileKeepUntyped)
	}
	return z
}

func (c *Comp) invalidUnaryExpr(node *ast.UnaryExpr, x *Expr) *Expr {
	return c.badUnaryExpr("invalid", node, x)
}

func (c *Comp) unimplementedUnaryExpr(node *ast.UnaryExpr, x *Expr) *Expr {
	return c.badUnaryExpr("unimplemented", node, x)
}

func (c *Comp) badUnaryExpr(reason string, node *ast.UnaryExpr, x *Expr) *Expr {
	opstr := mt.String(node.Op)
	c.Errorf("%s unary operation %s on <%v>: %s %v",
		reason, opstr, x.Type, opstr, node.X)
	return nil
}
