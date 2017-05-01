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

package fast

import (
	"go/constant"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func litValue(value I) Lit {
	return Lit{Type: r.TypeOf(value), Value: value}
}

func exprUntypedLit(kind r.Kind, value constant.Value) *Expr {
	return &Expr{Lit: Lit{Type: r.TypeOf(value), Value: UntypedLit{Kind: kind, Obj: value}}}
}

func exprValue(value I) *Expr {
	return &Expr{Lit: Lit{Type: r.TypeOf(value), Value: value}, IsNil: value == nil}
}

func exprLit(lit Lit, sym *Symbol) *Expr {
	return &Expr{Lit: lit, Sym: sym, IsNil: lit.Value == nil}
}

func exprFun(t r.Type, fun I) *Expr {
	return &Expr{Lit: Lit{Type: t}, Fun: fun}
}

func exprX1(t r.Type, fun func(env *Env) r.Value) *Expr {
	return &Expr{Lit: Lit{Type: t}, Fun: fun}
}

func exprXV(types []r.Type, fun func(env *Env) (r.Value, []r.Value)) *Expr {
	if len(types) == 1 {
		return &Expr{Lit: Lit{Type: types[0]}, Fun: fun}
	} else {
		return &Expr{Lit: Lit{Type: types[0]}, Types: types, Fun: fun}
	}
}

func exprBool(fun func(env *Env) bool) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfBool}, Fun: fun}
}

func exprInt(fun func(env *Env) int) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfInt}, Fun: fun}
}

func exprInt8(fun func(env *Env) int8) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfInt8}, Fun: fun}
}

func exprInt16(fun func(env *Env) int16) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfInt16}, Fun: fun}
}

func exprInt32(fun func(env *Env) int32) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfInt32}, Fun: fun}
}

func exprInt64(fun func(env *Env) int64) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfInt64}, Fun: fun}
}

func exprUint(fun func(env *Env) uint) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUint}, Fun: fun}
}

func exprUint8(fun func(env *Env) uint8) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUint8}, Fun: fun}
}

func exprUint16(fun func(env *Env) uint16) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUint16}, Fun: fun}
}

func exprUint32(fun func(env *Env) uint32) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUint32}, Fun: fun}
}

func exprUint64(fun func(env *Env) uint64) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUint64}, Fun: fun}
}

func exprUintptr(fun func(env *Env) uintptr) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUintptr}, Fun: fun}
}

func exprFloat32(fun func(env *Env) float32) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfFloat32}, Fun: fun}
}

func exprFloat64(fun func(env *Env) float64) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfFloat64}, Fun: fun}
}

func exprComplex64(fun func(env *Env) complex64) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfComplex64}, Fun: fun}
}

func exprComplex128(fun func(env *Env) complex128) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfComplex128}, Fun: fun}
}

func exprString(fun func(env *Env) string) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfString}, Fun: fun}
}

func (expr *Expr) EvalConst(opts CompileOptions) I {
	if expr == nil {
		return nil
	}
	if expr.Const() {
		return expr.Value
	}
	ret, rets := funAsXV(expr.Fun, opts)(nil)
	if ret == None {
		Errorf("constant should evaluate to a single value, found no values at all")
		return nil
	}
	if len(rets) > 1 {
		Errorf("constant should evaluate to a single value, found %d values: %v", len(rets), rets)
		return nil
	}
	t1 := expr.Type
	t2 := ValueType(ret)
	if t1 != t2 {
		Errorf("constant should evaluate to <%v>, found: %v <%v>", t1, t2, ret)
		return nil
	}
	var value I
	if ret != Nil {
		value = ret.Interface()
	}
	expr.Value = value
	expr.IsNil = value == nil
	expr.Fun = nil // no longer needed, will be recreated if needed as a wrapper for the computed value
	return value
}
