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

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) litValue(value I) Lit {
	return Lit{Type: c.TypeOf(value), Value: value}
}

func exprUntypedLit(kind r.Kind, value constant.Value) *Expr {
	return &Expr{Lit: Lit{Type: TypeOfUntypedLit, Value: UntypedLit{Kind: kind, Obj: value}}}
}

func (c *Comp) exprValue(t xr.Type, value I) *Expr {
	if t == nil {
		t = c.TypeOf(value)
	}
	return &Expr{Lit: Lit{Type: t, Value: value}, IsNil: value == nil}
}

func exprLit(lit Lit, sym *Symbol) *Expr {
	return &Expr{Lit: lit, Sym: sym, IsNil: lit.Value == nil}
}

func exprFun(t xr.Type, fun I) *Expr {
	return &Expr{Lit: Lit{Type: t}, Fun: fun}
}

func exprX1(t xr.Type, fun func(env *Env) r.Value) *Expr {
	return &Expr{Lit: Lit{Type: t}, Fun: fun}
}

func exprXV(types []xr.Type, fun func(env *Env) (r.Value, []r.Value)) *Expr {
	if len(types) == 1 {
		return &Expr{Lit: Lit{Type: types[0]}, Fun: fun}
	} else {
		return &Expr{Lit: Lit{Type: types[0]}, Types: types, Fun: fun}
	}
}

func expr0(fun func(env *Env)) *Expr {
	return &Expr{Types: zeroTypes, Fun: fun}
}

func exprBool(fun func(env *Env) bool) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfBool}, Fun: fun}
}

func exprInt(fun func(env *Env) int) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfInt}, Fun: fun}
}

func exprInt8(fun func(env *Env) int8) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfInt8}, Fun: fun}
}

func exprInt16(fun func(env *Env) int16) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfInt16}, Fun: fun}
}

func exprInt32(fun func(env *Env) int32) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfInt32}, Fun: fun}
}

func exprInt64(fun func(env *Env) int64) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfInt64}, Fun: fun}
}

func exprUint(fun func(env *Env) uint) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfUint}, Fun: fun}
}

func exprUint8(fun func(env *Env) uint8) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfUint8}, Fun: fun}
}

func exprUint16(fun func(env *Env) uint16) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfUint16}, Fun: fun}
}

func exprUint32(fun func(env *Env) uint32) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfUint32}, Fun: fun}
}

func exprUint64(fun func(env *Env) uint64) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfUint64}, Fun: fun}
}

func exprUintptr(fun func(env *Env) uintptr) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfUintptr}, Fun: fun}
}

func exprFloat32(fun func(env *Env) float32) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfFloat32}, Fun: fun}
}

func exprFloat64(fun func(env *Env) float64) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfFloat64}, Fun: fun}
}

func exprComplex64(fun func(env *Env) complex64) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfComplex64}, Fun: fun}
}

func exprComplex128(fun func(env *Env) complex128) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfComplex128}, Fun: fun}
}

func exprString(fun func(env *Env) string) *Expr {
	return &Expr{Lit: Lit{Type: xr.TypeOfString}, Fun: fun}
}

func (expr *Expr) EvalConst(opts CompileOptions) I {
	if expr == nil {
		return nil
	}
	if expr.Const() {
		return expr.Value
	}
	ret := expr.AsX1()(nil)
	if ret == base.None {
		base.Errorf("constant should evaluate to a single value, found no values at all")
		return nil
	}
	var value I
	if ret != base.Nil {
		value = ret.Interface()
	}
	expr.Value = value
	expr.IsNil = value == nil
	expr.Fun = nil // no longer needed, will be recreated if needed as a wrapper for the computed value
	return value
}
