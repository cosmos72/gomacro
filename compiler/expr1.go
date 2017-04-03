/*
 * gomacro - A Go intepreter with Lisp-like macros
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
 * expr1.go
 *
 *  Created on Apr 03, 2017
 *      Author Massimiliano Ghilardi
 */

package compiler

import (
	r "reflect"
)

func LitValue(value I) Lit {
	return Lit{Type: r.TypeOf(value), Value: value}
}

func ExprValue(value I) *Expr {
	return &Expr{Lit: Lit{Type: r.TypeOf(value), Value: value}, isNil: value == nil}
}

func ExprLit(lit Lit) *Expr {
	return &Expr{Lit: lit, isNil: lit.Value == nil}
}

func Expr1(t r.Type, fun func(env *Env) (r.Value, []r.Value)) *Expr {
	return &Expr{Lit: Lit{Type: t}, Fun: fun}
}

func ExprBool(fun func(env *Env) bool) *Expr {
	return &Expr{Lit: Lit{Type: typeOfBool}, Fun: fun}
}

func ExprInt(fun func(env *Env) int) *Expr {
	return &Expr{Lit: Lit{Type: typeOfInt}, Fun: fun}
}

func ExprInt8(fun func(env *Env) int8) *Expr {
	return &Expr{Lit: Lit{Type: typeOfInt8}, Fun: fun}
}

func ExprInt16(fun func(env *Env) int16) *Expr {
	return &Expr{Lit: Lit{Type: typeOfInt16}, Fun: fun}
}

func ExprInt32(fun func(env *Env) int32) *Expr {
	return &Expr{Lit: Lit{Type: typeOfInt32}, Fun: fun}
}

func ExprInt64(fun func(env *Env) int64) *Expr {
	return &Expr{Lit: Lit{Type: typeOfInt64}, Fun: fun}
}

func ExprUint(fun func(env *Env) uint) *Expr {
	return &Expr{Lit: Lit{Type: typeOfUint}, Fun: fun}
}

func ExprUint8(fun func(env *Env) uint8) *Expr {
	return &Expr{Lit: Lit{Type: typeOfUint8}, Fun: fun}
}

func ExprUint16(fun func(env *Env) uint16) *Expr {
	return &Expr{Lit: Lit{Type: typeOfUint16}, Fun: fun}
}

func ExprUint32(fun func(env *Env) uint32) *Expr {
	return &Expr{Lit: Lit{Type: typeOfUint32}, Fun: fun}
}

func ExprUint64(fun func(env *Env) uint64) *Expr {
	return &Expr{Lit: Lit{Type: typeOfUint64}, Fun: fun}
}

func ExprUintptr(fun func(env *Env) uintptr) *Expr {
	return &Expr{Lit: Lit{Type: typeOfUintptr}, Fun: fun}
}

func ExprFloat32(fun func(env *Env) float32) *Expr {
	return &Expr{Lit: Lit{Type: typeOfFloat32}, Fun: fun}
}

func ExprFloat64(fun func(env *Env) float64) *Expr {
	return &Expr{Lit: Lit{Type: typeOfFloat64}, Fun: fun}
}

func ExprComplex64(fun func(env *Env) complex64) *Expr {
	return &Expr{Lit: Lit{Type: typeOfComplex64}, Fun: fun}
}

func ExprComplex128(fun func(env *Env) complex128) *Expr {
	return &Expr{Lit: Lit{Type: typeOfComplex128}, Fun: fun}
}

func ExprString(fun func(env *Env) string) *Expr {
	return &Expr{Lit: Lit{Type: typeOfString}, Fun: fun}
}

func (expr *Expr) EvalConst() I {
	if expr == nil {
		return nil
	}
	if expr.Const() {
		return expr.Value
	}
	ret, rets := ToX(expr.Fun)(nil)
	if ret == None {
		errorf("constant should evaluate to a single value, found no values at all")
		return nil
	}
	if len(rets) > 1 {
		errorf("constant should evaluate to a single value, found %d values: %v", len(rets), rets)
		return nil
	}
	t1 := expr.Type
	t2 := ValueType(ret)
	if t1 != t2 {
		errorf("constant should evaluate to <%v>, found: %v <%v>", t1, t2, ret)
		return nil
	}
	var value I
	if ret != Nil {
		value = ret.Interface()
	}
	expr.Value = value
	expr.isNil = value == nil
	expr.Fun = nil // no longer needed.
	return value
}
