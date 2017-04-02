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

	mt "github.com/cosmos72/gomacro/token"
)

func (c *Comp) BasicLit(node *ast.BasicLit) I {
	str := node.Value
	switch node.Kind {
	case token.INT:
		if strings.HasPrefix(str, "-") {
			i64, err := strconv.ParseInt(str, 0, 64)
			if err != nil {
				return error_(err)
			}
			// prefer int to int64. reason: in compiled Go,
			// type inference deduces int for all constants representable by an int
			i := int(i64)
			if int64(i) == i64 {
				return i
			}
			return i64
		} else {
			u64, err := strconv.ParseUint(str, 0, 64)
			if err != nil {
				return error_(err)
			}
			// prefer, in order: int, int64, uint, uint64. reason: in compiled Go,
			// type inference deduces int for all constants representable by an int
			i := int(u64)
			if i >= 0 && uint64(i) == u64 {
				return i
			}
			i64 := int64(u64)
			if i64 >= 0 && uint64(i64) == u64 {
				return i64
			}
			u := uint(u64)
			if uint64(u) == u64 {
				return u
			}
			return u64
		}

	case token.FLOAT:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return error_(err)
		}
		return f

	case token.IMAG:
		if strings.HasSuffix(str, "i") {
			str = str[:len(str)-1]
		}
		im, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return error_(err)
		}
		return complex(0.0, im)
		// env.Debugf("evalLiteral(): parsed IMAG %s -> %T %#v -> %T %#v", str, im, im, ret, ret)

	case token.CHAR:
		return unescapeChar(str)

	case token.STRING:
		return unescapeString(str)

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
		return // no conversion happened
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

// convertLiteral converts a literal value to the given type.
// panics on integer overflow.
func convertLiteral(x I, t r.Type) I {
	v1 := r.ValueOf(x)
	v2 := Nil
	if t != nil {
		ok := false
		defer func() {
			if !ok {
				errorf("cannot convert %v <%T> to <%v>", x, x, t)
			}
		}()
		v2 = v1.Convert(t)
		ok = true
	}
	checkLiteralOverflow(v1, v2)
	if v2 == Nil || v2 == None {
		return nil
	}
	return v2.Interface()
}

// wrapLiteral wraps a literal value into a closure that will return it,
// converted to the given type.
func wrapLiteral(x I, t r.Type) I {
	x = convertLiteral(x, t)
	switch x := x.(type) {
	case nil:
		return func(env *Env) interface{} {
			return nil
		}
	case bool:
		return func(env *Env) bool {
			return x
		}
	case int:
		return func(env *Env) int {
			return x
		}
	case int8:
		return func(env *Env) int8 {
			return x
		}
	case int16:
		return func(env *Env) int16 {
			return x
		}
	case int32:
		return func(env *Env) int32 {
			return x
		}
	case int64:
		return func(env *Env) int64 {
			return x
		}
	case uint:
		return func(env *Env) uint {
			return x
		}
	case uint8:
		return func(env *Env) uint8 {
			return x
		}
	case uint16:
		return func(env *Env) uint16 {
			return x
		}
	case uint32:
		return func(env *Env) uint32 {
			return x
		}
	case uint64:
		return func(env *Env) uint64 {
			return x
		}
	case uintptr:
		return func(env *Env) uintptr {
			return x
		}
	case float32:
		return func(env *Env) float32 {
			return x
		}
	case float64:
		return func(env *Env) float64 {
			return x
		}
	case complex64:
		return func(env *Env) complex64 {
			return x
		}
	case complex128:
		return func(env *Env) complex128 {
			return x
		}
	case string:
		return func(env *Env) string {
			return x
		}
	default:
		if r.ValueOf(x).Kind() == r.Func {
			return x
		}
		errorf("internal error: unexpected value, expecting a literal or a func(*Env)..., found: %#v <%v>",
			x, r.TypeOf(x))
		return nil
	}
}

func literalsToSameType(op token.Token, x, y I) (I, I) {
	if x == nil {
		if y == nil {
			return x, y
		} else {
			str := mt.String(op)
			errorf("unsupported binary operation <%T> %s <%T>: %v %s %v",
				x, str, y, x, str, y)
		}
	}
	xt, yt := TypeOf(x), TypeOf(y)
	xk, yk := xt.Kind(), yt.Kind()
	if xk == yk {
		return x, y
	}
	xc, yc := kindToClass(xk), kindToClass(yk)
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
			x = complex(r.ValueOf(x).Convert(typeOfFloat64).Float(), 0.0)
		}
		if yc != r.Complex128 {
			y = complex(r.ValueOf(y).Convert(typeOfFloat64).Float(), 0.0)
		}
	} else if xc == r.Float64 || yc == r.Float64 {
		x = r.ValueOf(x).Convert(typeOfFloat64).Float()
		y = r.ValueOf(y).Convert(typeOfFloat64).Float()
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
