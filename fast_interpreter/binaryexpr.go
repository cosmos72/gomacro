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
	"go/ast"
	"go/constant"
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

	if x.Untyped() && y.Untyped() {
		return c.BinaryExprUntyped(node, x.Value.(UntypedLit), y.Value.(UntypedLit))
	}
	bothConst := x.Const() && y.Const()
	var z *Expr

	switch node.Op {
	case token.ADD, token.ADD_ASSIGN:
		z = c.Add(node, x, y)
	case token.SUB, token.SUB_ASSIGN:
		z = c.Sub(node, x, y)
	case token.MUL, token.MUL_ASSIGN:
		z = c.Mul(node, x, y)
	case token.QUO, token.QUO_ASSIGN:
		z = c.Quo(node, x, y)
	case token.REM, token.REM_ASSIGN:
		z = c.Rem(node, x, y)
	case token.AND, token.AND_ASSIGN:
		z = c.And(node, x, y)
	case token.OR, token.OR_ASSIGN:
		z = c.Or(node, x, y)
	case token.XOR, token.XOR_ASSIGN:
		z = c.Xor(node, x, y)
		/*
			case token.SHL, token.SHL_ASSIGN:
				z = c.Shl(node, x, y)
			case token.SHR, token.SHR_ASSIGN:
				z = c.Shr(node, x, y)
		*/
	case token.AND_NOT, token.AND_NOT_ASSIGN:
		z = c.Andnot(node, x, y)
	case token.LAND:
		z = c.Land(node, x, y)
	case token.LOR:
		z = c.Lor(node, x, y)
	case token.EQL:
		z = c.Eql(node, x, y)
	case token.LSS:
		z = c.Lss(node, x, y)
	case token.GTR:
		z = c.Gtr(node, x, y)
	case token.NEQ:
		z = c.Neq(node, x, y)
	case token.LEQ:
		z = c.Leq(node, x, y)
	case token.GEQ:
		z = c.Geq(node, x, y)
	default:
		return c.unimplementedBinaryExpr(node, x, y)
	}
	if bothConst {
		// constant propagation
		z.EvalConst()
	}
	return z
}

func (c *Comp) BinaryExprUntyped(node *ast.BinaryExpr, x UntypedLit, y UntypedLit) *Expr {
	op := node.Op
	switch op {
	case token.LAND, token.LOR, token.EQL, token.LSS, token.GTR, token.NEQ, token.LEQ, token.GEQ:
		return ExprValue(constant.Compare(x.Obj, op, y.Obj))
	case token.SHL_ASSIGN:
		op = token.SHL
	case token.SHR_ASSIGN:
		op = token.SHL
	case token.SHL, token.SHR:
	default:
		zobj := constant.BinaryOp(x.Obj, op, y.Obj)
		var zkind r.Kind
		if zobj.Kind() == constant.Int {
			// reflect.Int32 (i.e. rune) has precedence over reflect.Int
			if x.Kind != r.Int {
				zkind = x.Kind
			} else if y.Kind != r.Int {
				zkind = y.Kind
			} else {
				zkind = r.Int
			}
		} else {
			zkind = constantKindToUntypedLitKind(zobj.Kind())
		}
		if zkind == r.Invalid {
			c.Errorf("invalid binary operation: %v %v %v", x.Obj, op, y.Obj)
		}
		return ExprUntypedLit(zkind, zobj)
	}
	if y.Obj.Kind() != constant.Int {
		c.Errorf("invalid shift: %v %v %v", x.Obj, op, y.Obj)
	}
	yi, exact := constant.Int64Val(y.Obj)
	yn := uint(yi)
	if !exact || yn < 0 || yi != int64(yn) {
		c.Errorf("invalid shift: %v %v %v", x.Obj, op, y.Obj)
	}
	return ExprValue(constant.Shift(x.Obj, op, yn))
}

func constantKindToUntypedLitKind(ckind constant.Kind) r.Kind {
	ret := r.Invalid
	switch ckind {
	case constant.Bool:
		ret = r.Bool
	case constant.String:
		ret = r.String
	case constant.Int:
		ret = r.Int // actually ambiguous, could be a rune - thus r.Int64
	case constant.Float:
		ret = r.Float64
	case constant.Complex:
		ret = r.Complex64
	}
	return ret
}

func (c *Comp) Shl(node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	return c.unimplementedBinaryExpr(node, x, y)
}

func (c *Comp) Shr(node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	return c.unimplementedBinaryExpr(node, x, y)
}

func (c *Comp) Land(node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	xval, xfun, xerr := x.AsPred()
	yval, yfun, yerr := y.AsPred()
	if xerr || yerr {
		return c.invalidBinaryExpr(node, x, y)
	}
	// optimize short-circuit logic
	if xfun == nil {
		if xval {
			return y
		}
		return ExprValue(false)
	}
	if yfun == nil {
		if yval {
			return x
		}
		return ExprBool(func(env *Env) bool {
			return xfun(env) && false
		})
	}
	return ExprBool(func(env *Env) bool {
		return xfun(env) && yfun(env)
	})
}

func (c *Comp) Lor(node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	xval, xfun, xerr := x.AsPred()
	yval, yfun, yerr := y.AsPred()
	if xerr || yerr {
		return c.invalidBinaryExpr(node, x, y)
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

func (c *Comp) invalidBinaryExpr(node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	return c.badBinaryExpr("invalid", node, x, y)
}

func (c *Comp) unimplementedBinaryExpr(node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	return c.badBinaryExpr("unimplemented", node, x, y)
}

func (c *Comp) badBinaryExpr(reason string, node *ast.BinaryExpr, x *Expr, y *Expr) *Expr {
	opstr := mt.String(node.Op)
	c.Errorf("%s binary operation %s between <%v> and <%v>: %v %s %v",
		reason, opstr, x.Type, y.Type, node.X, opstr, node.Y)
	return nil
}

// convert x and y to the same single-valued expression type. needed to emulate Go untyped constants
func (c *Comp) toSameFuncType(node *ast.BinaryExpr, xe *Expr, ye *Expr) {
	xe.CheckX1()
	ye.CheckX1()
	xconst, yconst := xe.Const(), ye.Const()
	if yconst {
		if xconst {
			x, y := c.constsToSameType(node, xe, ye)
			xe.SetWithFun(x)
			ye.SetWithFun(y)
		} else {
			ye.ConstTo(xe.Type)
		}
	} else {
		if xconst {
			xe.ConstTo(ye.Type)
		} else {
			if xe.Type != ye.Type {
				c.invalidBinaryExpr(node, xe, ye)
			}
		}
	}
}

func (c *Comp) constsToSameType(node *ast.BinaryExpr, xe *Expr, ye *Expr) (I, I) {
	x, y := xe.Value, ye.Value
	if x == nil {
		if y == nil {
			return x, y
		} else {
			c.invalidBinaryExpr(node, xe, ye)
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
		c.badBinaryExpr("cannot convert operands to the same type in", node, xe, ye)
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
