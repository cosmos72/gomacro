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
 * binaryexpr.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package compiler

import (
	"go/ast"
	"go/token"

	mt "github.com/cosmos72/gomacro/token"
)

func (c *Comp) BinaryExpr(node *ast.BinaryExpr) I {
	x := c.Expr(node.X)
	y := c.Expr(node.Y)
	bothLiteral := isLiteral(x) && isLiteral(y)
	var z I

	switch op := node.Op; op {
	case token.ADD, token.ADD_ASSIGN:
		z = c.Add(op, x, y)
	case token.SUB, token.SUB_ASSIGN:
		z = c.Sub(op, x, y)
	case token.MUL, token.MUL_ASSIGN:
		z = c.Mul(op, x, y)
	case token.QUO, token.QUO_ASSIGN:
		z = c.Quo(op, x, y)
	case token.REM, token.REM_ASSIGN:
		z = c.Rem(op, x, y)
	case token.AND, token.AND_ASSIGN:
		z = c.And(op, x, y)
	case token.OR, token.OR_ASSIGN:
		z = c.Or(op, x, y)
	case token.XOR, token.XOR_ASSIGN:
		z = c.Xor(op, x, y)
	case token.SHL, token.SHL_ASSIGN:
		z = c.Shl(op, x, y)
	case token.SHR, token.SHR_ASSIGN:
		z = c.Shr(op, x, y)
	case token.AND_NOT, token.AND_NOT_ASSIGN:
		z = c.Andnot(op, x, y)
	case token.LAND:
		z = c.Land(op, x, y)
	case token.LOR:
		z = c.Lor(op, x, y)
	case token.EQL:
		z = c.Eql(op, x, y)
	case token.LSS:
		z = c.Lss(op, x, y)
	case token.GTR:
		z = c.Gtr(op, x, y)
	case token.NEQ:
		z = c.Neq(op, x, y)
	case token.LEQ:
		z = c.Leq(op, x, y)
	case token.GEQ:
		z = c.Geq(op, x, y)
	default:
		return c.unsupportedBinaryExpr(node.Op, x, y)
	}
	if bothLiteral {
		// constant propagation
		return c.evalConst(z)
	} else {
		return z
	}
}

// convert x and y to the same type. needed to emulate Go untyped constants
func toSameFuncType(op token.Token, x, y I) (I, I) {
	xlit, ylit := isLiteral(x), isLiteral(y)
	if ylit {
		if xlit {
			x, y = literalsToSameType(op, x, y)
			t := TypeOf(x)
			return wrapLiteral(x, t), wrapLiteral(y, t)
		} else {
			xt := RetOf(x)
			return x, wrapLiteral(y, xt)
		}
	} else {
		if xlit {
			yt := RetOf(y)
			return wrapLiteral(x, yt), y
		} else {
			return x, y
		}
	}
}

func (c *Comp) Shl(op token.Token, x I, y I) I {
	return c.unsupportedBinaryExpr(op, x, y)
}

func (c *Comp) Shr(op token.Token, x I, y I) I {
	return c.unsupportedBinaryExpr(op, x, y)
}

func toPred(expr I) (val bool, pred func(*Env) bool, err bool) {
	switch expr := expr.(type) {
	case bool:
		return expr, nil, false
	case func(*Env) bool:
		return false, expr, false
	default:
		return false, nil, true
	}
}

func (c *Comp) Land(op token.Token, x I, y I) I {
	xval, xfun, xerr := toPred(x)
	yval, yfun, yerr := toPred(y)
	if xerr || yerr {
		return c.unsupportedBinaryExpr(op, x, y)
	}
	// optimize short-circuit logic
	if xfun == nil {
		if xval {
			return y
		}
		return false
	}
	if yfun == nil {
		if yval {
			return x
		}
		return func(env *Env) bool {
			return xfun(env) && false
		}
	}
	return func(env *Env) bool {
		return xfun(env) && yfun(env)
	}
}

func (c *Comp) Lor(op token.Token, x I, y I) I {
	xval, xfun, xerr := toPred(x)
	yval, yfun, yerr := toPred(y)
	if xerr || yerr {
		return c.unsupportedBinaryExpr(op, x, y)
	}
	// optimize short-circuit logic
	if xfun == nil {
		if xval {
			return true
		}
		return y
	}
	if yfun == nil {
		if yval {
			return func(env *Env) bool {
				return xfun(env) || true
			}
		}
		return x
	}
	return func(env *Env) bool {
		return xfun(env) || yfun(env)
	}
}

func (c *Comp) unsupportedBinaryExpr(op token.Token, x I, y I) I {
	str := mt.String(op)
	c.Errorf("unsupported binary operation <%v> %s <%v> in expression: %v %s %v",
		RetOf(x), str, RetOf(y), x, str, y)
	return nil
}
