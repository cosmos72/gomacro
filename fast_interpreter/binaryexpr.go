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

package fast_interpreter

import (
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	mt "github.com/cosmos72/gomacro/token"
)

func (c *Comp) BinaryExpr(node *ast.BinaryExpr) *Expr {
	x := c.Expr(node.X)
	y := c.Expr(node.Y)
	if x.NumOut() == 0 {
		c.Errorf("operand returns no values, cannot use in binary expression: %v", node.X)
	} else if y.NumOut() == 0 {
		c.Errorf("operand returns no values, cannot use in binary expression: %v", node.Y)
	}

	bothConst := x.Const() && y.Const()
	var z *Expr

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
		/*
			case token.SHL, token.SHL_ASSIGN:
				z = c.Shl(op, x, y)
			case token.SHR, token.SHR_ASSIGN:
				z = c.Shr(op, x, y)
		*/
	case token.AND_NOT, token.AND_NOT_ASSIGN:
		z = c.Andnot(op, x, y)
		/*
			case token.LAND:
				z = c.Land(op, x, y)
			case token.LOR:
				z = c.Lor(op, x, y)
		*/
	case token.EQL:
		z = c.Eql(op, x, y)
		/*
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
		*/
	default:
		c.invalidBinaryExpr(node.Op, x, y)
		return z
	}
	if bothConst {
		// constant propagation
		z.EvalConst()
	}
	return z
}

func (c *Comp) Shl(op token.Token, x *Expr, y *Expr) I {
	return c.invalidBinaryExpr(op, x, y)
}

func (c *Comp) Shr(op token.Token, x *Expr, y *Expr) I {
	return c.invalidBinaryExpr(op, x, y)
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

func (c *Comp) Land(op token.Token, x *Expr, y *Expr) I {
	xval, xfun, xerr := toPred(x)
	yval, yfun, yerr := toPred(y)
	if xerr || yerr {
		return c.invalidBinaryExpr(op, x, y)
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

func (c *Comp) Lor(op token.Token, x *Expr, y *Expr) *Expr {
	xval, xfun, xerr := toPred(x)
	yval, yfun, yerr := toPred(y)
	if xerr || yerr {
		return c.invalidBinaryExpr(op, x, y)
	}
	// optimize short-circuit logic
	if xfun == nil {
		if xval {
			return ExprValue(true)
		}
		return y
	}
	if yfun == nil {
		if yval {
			return ExprBool(func(env *Env) bool {
				return xfun(env) || true
			})
		}
		return x
	}
	return ExprBool(func(env *Env) bool {
		return xfun(env) || yfun(env)
	})
}

func (c *Comp) invalidBinaryExpr(op token.Token, x *Expr, y *Expr) *Expr {
	opstr := mt.String(op)
	xstr := "<expr>"
	ystr := xstr
	if x.Const() {
		if x.IsNil {
			xstr = "nil"
		} else {
			xstr = fmt.Sprintf("%v", x.Value)
		}
	}
	if y.Const() {
		if y.IsNil {
			ystr = "nil"
		} else {
			ystr = fmt.Sprintf("%v", y.Value)
		}
	}
	c.Errorf("invalid binary operation <%v> %s <%v> in expression: %v %s %v",
		x.Type, opstr, y.Type, xstr, opstr, ystr)
	return nil
}

// convert x and y to the same single-valued expression type. needed to emulate Go untyped constants
func toSameFuncType(op token.Token, x, y *Expr) {
	x.CheckX1()
	y.CheckX1()
	xc, yc := x.Const(), y.Const()
	if yc {
		if xc {
			xi, yi := constsToSameType(op, x.Value, y.Value)
			x.SetWithFun(xi)
			y.SetWithFun(yi)
		} else {
			y.ConstTo(x.Type)
		}
	} else {
		if xc {
			x.ConstTo(y.Type)
		} else {
			if x.Type != y.Type {
				errorf("unsupported binary operation <%v> %s <%v>", x.Type, mt.String(op), y.Type)
			}
		}
	}
}

func constsToSameType(op token.Token, x, y I) (I, I) {
	if x == nil {
		if y == nil {
			return x, y
		} else {
			str := mt.String(op)
			errorf("unsupported binary operation <%T> %s <%T>: %v %s %v",
				x, str, y, x, str, y)
		}
	}
	xt, yt := r.TypeOf(x), r.TypeOf(y)
	xk, yk := xt.Kind(), yt.Kind()
	if xk == yk {
		return x, y
	}
	xc, yc := kindToCategory(xk), kindToCategory(yk)
	if xc == yc {
		// same class, only the number of bits differs
		if xk < yk {
			x = r.ValueOf(x).Convert(yt).Interface()
		} else {
			y = r.ValueOf(y).Convert(xt).Interface()
		}
		return x, y
	}
	if xc == r.Int && yc == r.Uint {
		// mixing signed and unsigned integers
		xi := r.ValueOf(x).Int()
		yi := r.ValueOf(y).Uint()
		return intsToSameType(xi, yi, xt, yt)
	} else if xc == r.Uint && yc == r.Int {
		// mixing signed and unsigned integers
		xi := r.ValueOf(x).Uint()
		yi := r.ValueOf(y).Int()
		y, x = intsToSameType(yi, xi, yt, xt)
		return x, y
	}
	// at least one is a float or complex... or a non-integer
	if xc == r.Complex128 || yc == r.Complex128 {
		if xc != r.Complex128 {
			x = complex(r.ValueOf(x).Convert(TypeOfFloat64).Float(), 0.0)
		}
		if yc != r.Complex128 {
			y = complex(r.ValueOf(y).Convert(TypeOfFloat64).Float(), 0.0)
		}
	} else if xc == r.Float64 || yc == r.Float64 {
		x = r.ValueOf(x).Convert(TypeOfFloat64).Float()
		y = r.ValueOf(y).Convert(TypeOfFloat64).Float()
	} else {
		errorf("cannot convert  to the same type: %v <%T> and %v <%T>", x, x, y, y)
	}
	return x, y
}

func intsToSameType(x int64, y uint64, xt r.Type, yt r.Type) (I, I) {
	// try int for both
	if x == int64(int(x)) && y == uint64(int(y)) {
		return int(x), int(y)
	}
	// try uint for both
	if x >= 0 && x == int64(uint(x)) && y == uint64(uint(y)) {
		return uint(x), uint(y)
	}
	// try int64 for both
	if y == uint64(int64(y)) {
		return x, int64(y)
	}
	// try uint64 for both
	if x >= 0 && x == int64(uint64(x)) {
		return uint64(x), y
	}
	return float64(x), float64(y)
}
