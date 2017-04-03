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
 * literal.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package compiler

import (
	"go/ast"
	"go/token"
	r "reflect"
	"strconv"
	"strings"
)

func (c *Comp) BasicLit(node *ast.BasicLit) *Expr {
	str := node.Value
	switch node.Kind {
	case token.INT:
		if strings.HasPrefix(str, "-") {
			i64, err := strconv.ParseInt(str, 0, 64)
			if err != nil {
				c.Error(err)
				return nil
			}
			// prefer int to int64. reason: in compiled Go,
			// type inference deduces int for all constants representable by an int
			i := int(i64)
			if int64(i) == i64 {
				return ExprValue(i)
			}
			return ExprValue(i64)
		} else {
			u64, err := strconv.ParseUint(str, 0, 64)
			if err != nil {
				c.Error(err)
				return nil
			}
			// prefer, in order: int, int64, uint, uint64. reason: in compiled Go,
			// type inference deduces int for all constants representable by an int
			i := int(u64)
			if i >= 0 && uint64(i) == u64 {
				return ExprValue(i)
			}
			i64 := int64(u64)
			if i64 >= 0 && uint64(i64) == u64 {
				return ExprValue(i64)
			}
			u := uint(u64)
			if uint64(u) == u64 {
				return ExprValue(u)
			}
			return ExprValue(u64)
		}

	case token.FLOAT:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			c.Error(err)
			return nil
		}
		return ExprValue(f)

	case token.IMAG:
		if strings.HasSuffix(str, "i") {
			str = str[:len(str)-1]
		}
		im, err := strconv.ParseFloat(str, 64)
		if err != nil {
			c.Error(err)
			return nil
		}
		return ExprValue(complex(0.0, im))
		// env.Debugf("evalLiteral(): parsed IMAG %s -> %T %#v -> %T %#v", str, im, im, ret, ret)

	case token.CHAR:
		return ExprValue(unescapeChar(str))

	case token.STRING:
		return ExprValue(unescapeString(str))

	default:
		c.Errorf("unimplemented basic literal: %v", node)
		return nil
	}
}

func isLiteral(x I) bool {
	switch x.(type) {
	case nil, bool,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64, complex64, complex128, string:
		return true
	default:
		return false
	}
}

func isLiteralNumber(x I, n int64) bool {
	switch x.(type) {
	case int, int8, int16, int32, int64:
		return r.ValueOf(x).Int() == n
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return n >= 0 && r.ValueOf(x).Uint() == uint64(n)
	case float32, float64:
		return r.ValueOf(x).Float() == float64(n)
	case complex64, complex128:
		return r.ValueOf(x).Complex() == complex(float64(n), 0)
	default:
		return false
	}
}

func kindToClass(k r.Kind) r.Kind {
	switch k {
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return r.Int
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		return r.Uint
	case r.Float32, r.Float64:
		return r.Float64
	case r.Complex64, r.Complex128:
		return r.Complex128
	default:
		return k
	}
}

func isClass(k r.Kind, classes ...r.Kind) bool {
	k = kindToClass(k)
	for _, c := range classes {
		if k == c {
			return true
		}
	}
	return false
}

// checkLiteralOverflow panics if the conversion from vsrc to vdst overflowed the destination type
func checkLiteralOverflow(vsrc, vdst r.Value) {
	k1, k2 := vsrc.Kind(), vdst.Kind()
	if k1 == k2 {
		return // no conversion happenedConstTo
	}
	c1, c2 := kindToClass(k1), kindToClass(k2)
	if c2 == r.Int || c2 == r.Uint {
		if c1 == r.Float64 || c1 == r.Complex128 {
			// float-to-integer conversion. check for truncation
			t1 := ValueType(vsrc)
			vback := vdst.Convert(t1)
			if src := vsrc.Interface(); src != vback.Interface() {
				errorf("constant %v truncated to %v", src, ValueType(vdst))
			}
		} else {
			// integer-to-integer conversion. convert back and compare the interfaces for overflows
			t1 := ValueType(vsrc)
			vback := vdst.Convert(t1)
			if src := vsrc.Interface(); src != vback.Interface() {
				errorf("constant %v overflows %v", src, ValueType(vdst))
			}
		}
	}
}

// ConstTo converts a constant Expr to the given type.
// panics if not constant and on integer overflow.
func (e *Expr) ConstTo(t r.Type) I {
	if !e.Const() {
		errorf("expression is not a constant, cannot convert from <%v> to <%v>", e.Type, t)
	}
	return e.Lit.ConstTo(t)
}

// ConstTo converts a Lit to the given type.
// panics on conversion failure and on integer overflow.
func (lit *Lit) ConstTo(t r.Type) I {
	x := lit.Value
	if t == nil {
		// only literal nil has type nil
		if x != nil {
			errorf("cannot convert constant %v <%v> to <nil>", lit.Value, lit.Type)
		}
		return nil
	}
	if t == lit.Type {
		return x
	}
	v1 := r.ValueOf(x)
	ok := false
	defer func() {
		if !ok {
			errorf("cannot convert constant %v <%T> to <%v>", x, x, t)
		}
	}()
	v2 := v1.Convert(t)
	ok = true
	checkLiteralOverflow(v1, v2)
	x = v2.Interface()
	lit.Type = t
	lit.Value = x
	return x
}

// Set sets the expression value to the given constant
func (e *Expr) Set(x I) {
	e.Lit.Set(x)
	e.Types = nil
	e.Fun = nil
	e.isNil = x == nil
}

// Set sets the Lit to the given constant
func (e *Lit) Set(x I) {
	e.Type = r.TypeOf(x)
	e.Value = x
}

// SetWithFun sets the expression value to the given constant,
// then sets Expr.Fun to a closure that will return such constant
func (e *Expr) SetWithFun(x I) {
	e.Set(x)
	e.WithFun()
}

// if Expr is a constant, WithFun sets Expr.Fun to a closure that will return such constant
// in any case, it will then return Expr.Fun
func (e *Expr) WithFun() I {
	if !e.Const() {
		return e.Fun
	}
	var fun I
	switch x := e.Value.(type) {
	case nil:
		fun = XNil
	case bool:
		if x {
			fun = XTrue
		} else {
			fun = XFalse
		}
	case int:
		fun = func(env *Env) int {
			return x
		}
	case int8:
		fun = func(env *Env) int8 {
			return x
		}
	case int16:
		fun = func(env *Env) int16 {
			return x
		}
	case int32:
		fun = func(env *Env) int32 {
			return x
		}
	case int64:
		fun = func(env *Env) int64 {
			return x
		}
	case uint:
		fun = func(env *Env) uint {
			return x
		}
	case uint8:
		fun = func(env *Env) uint8 {
			return x
		}
	case uint16:
		fun = func(env *Env) uint16 {
			return x
		}
	case uint32:
		fun = func(env *Env) uint32 {
			return x
		}
	case uint64:
		fun = func(env *Env) uint64 {
			return x
		}
	case uintptr:
		fun = func(env *Env) uintptr {
			return x
		}
	case float32:
		fun = func(env *Env) float32 {
			return x
		}
	case float64:
		fun = func(env *Env) float64 {
			return x
		}
	case complex64:
		fun = func(env *Env) complex64 {
			return x
		}
	case complex128:
		fun = func(env *Env) complex128 {
			return x
		}
	case string:
		fun = func(env *Env) string {
			return x
		}
	default:
		v := r.ValueOf(x)
		fun = func(env *Env) (r.Value, []r.Value) {
			return v, nil
		}
	}
	e.Fun = fun
	return fun
}
