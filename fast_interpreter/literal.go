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

package fast_interpreter

import (
	"go/ast"
	"go/constant"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) BasicLit(node *ast.BasicLit) *Expr {
	str := node.Value
	var kind r.Kind
	var label string
	switch node.Kind {
	case token.INT:
		kind, label = r.Int, "integer"
	case token.FLOAT:
		kind, label = r.Float64, "float"
	case token.IMAG:
		kind, label = r.Complex128, "complex"
	case token.CHAR:
		kind, label = r.Int32, "rune"
	case token.STRING:
		kind, label = r.String, "string"
	default:
		c.Errorf("unsupported basic literal: %v", node)
		return nil
	}
	obj := constant.MakeFromLiteral(str, node.Kind, 0)
	if obj.Kind() == constant.Unknown {
		c.Errorf("invalid %s literal: %v", label, str)
		return nil
	}
	return ExprValue(UntypedLit{Kind: kind, Obj: obj})
}

func isLiteral(x I) bool {
	switch x.(type) {
	case nil, bool, int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64, complex64, complex128, string,
		UntypedLit:
		return true
	default:
		return false
	}
}

func isLiteralNumber(x I, n int64) bool {
	switch x := x.(type) {
	case int, int8, int16, int32, int64:
		return r.ValueOf(x).Int() == n
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return n >= 0 && r.ValueOf(x).Uint() == uint64(n)
	case float32, float64:
		return r.ValueOf(x).Float() == float64(n)
	case complex64, complex128:
		return r.ValueOf(x).Complex() == complex(float64(n), 0)
	case UntypedLit:
		return x.IsLiteralNumber(n)
	default:
		return false
	}
}

func (untyp *UntypedLit) IsLiteralNumber(n int64) bool {
	obj := untyp.Obj
	switch obj.Kind() {
	case constant.Int:
		m, exact := constant.Int64Val(obj)
		return exact && m == n
	case constant.Float:
		m, exact := constant.Float64Val(obj)
		return exact && float64(int64(m)) == m && int64(m) == n
	case constant.Complex:
		m, exact := constant.Float64Val(constant.Imag(obj))
		if !exact || m != 0.0 {
			return false
		}
		m, exact = constant.Float64Val(constant.Real(obj))
		return exact && float64(int64(m)) == m && int64(m) == n
	default:
		return false
	}
}

func kindToCategory(k r.Kind) r.Kind {
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

func isCategory(k r.Kind, categories ...r.Kind) bool {
	k = kindToCategory(k)
	for _, c := range categories {
		if k == c {
			return true
		}
	}
	return false
}

// ================================= ConstTo =================================

// ConstTo checks that a constant Expr can be used as the given type.
// panics if not constant, or if Expr is a typed constant of different type
// actually performs type conversion (and subsequent overflow checks) ONLY on untyped constants.
func (e *Expr) ConstTo(t r.Type) I {
	if !e.Const() {
		Errorf("expression is not a constant, cannot convert from <%v> to <%v>", e.Type, t)
	}
	return e.Lit.ConstTo(t)
}

// ConstTo checks that a Lit can be used as the given type.
// panics if Lit is a typed constant of different type
// actually performs type conversion (and subsequent overflow checks) ONLY on untyped constants.
func (lit *Lit) ConstTo(t r.Type) I {
	value := lit.Value
	if t == nil {
		// only literal nil has type nil
		if value != nil {
			Errorf("cannot convert constant %v <%v> to <nil>", value, lit.Type)
		}
		return nil
	}
	if t == lit.Type {
		return value
	}
	switch x := value.(type) {
	case UntypedLit:
		value = x.ConstTo(t)
		lit.Value = value
		lit.Type = t
	default:
		Errorf("cannot convert typed constant %v <%v> to <%v>", value, r.TypeOf(value), t)
	}
	return value
}

// ConstTo checks that an UntypedLit can be used as the given type.
// performs actual untyped -> typed conversion and subsequent overflow checks.
// returns the constant converted to given type
func (untyp *UntypedLit) ConstTo(t r.Type) I {
	obj := untyp.Obj
again:
	// Debugf("converting untyped constant %v to <%v>", untyp, t)
	switch t.Kind() {
	case r.Bool:
		if obj.Kind() != constant.Bool {
			Errorf("cannot convert untyped constant %v to <%v>", untyp, t)
		}
		return constant.BoolVal(obj)
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
		r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr,
		r.Float32, r.Float64, r.Complex64, r.Complex128:

		n := untyp.extractNumber(obj, t)
		return convertLiteralCheckOverflow(n, t)
	case r.String:
		if untyp.Obj.Kind() != constant.String {
			Errorf("cannot convert untyped constant %v to <%v>", untyp, t)
		}
		return UnescapeString(obj.ExactString())
	case r.Interface:
		// this can happen too... for example in "var foo interface{} = 7"
		// and it requites to convert the untyped constant to its default type.
		// obviously, untyped constants can only implement empty interfaces
		if t.NumMethod() == 0 {
			t = untyp.DefaultType()
			goto again
		}
		fallthrough
	default:
		Errorf("cannot convert untyped constant %v to <%v>", untyp, t)
		return nil
	}
}

// ================================= DefaultType =================================

// DefaultType returns the default type of an expression.
func (e *Expr) DefaultType() r.Type {
	if e.Untyped() {
		return e.Lit.DefaultType()
	}
	return e.Type
}

// DefaultType returns the default type of a constant.
func (lit *Lit) DefaultType() r.Type {
	switch x := lit.Value.(type) {
	case UntypedLit:
		return x.DefaultType()
	default:
		return lit.Type
	}
}

// DefaultType returns the default type of an untyped constant.
func (untyp *UntypedLit) DefaultType() r.Type {
	switch untyp.Kind {
	case r.Bool:
		return TypeOfBool
	case r.Int32: // rune
		return TypeOfInt32
	case r.Int:
		return TypeOfInt
	case r.Uint:
		return TypeOfUint
	case r.Float64:
		return TypeOfFloat64
	case r.Complex128:
		return TypeOfComplex128
	case r.String:
		return TypeOfString
	default:
		Errorf("unexpected untyped constant %v, its default type is not known", untyp)
		return nil
	}
}

// ======================= utilities for ConstTo and ConstToDefaultType =======================

// extractNumber converts the untyped constant src to an integer, float or complex.
// panics if src has different kind from constant.Int, constant.Float and constant.Complex
// the receiver (untyp UntypedLit) and the second argument (t reflect.Type) are only used to pretty-print the panic error message
func (untyp *UntypedLit) extractNumber(src constant.Value, t r.Type) interface{} {
	var n interface{}
	var exact bool
	switch src.Kind() {
	case constant.Int:
		n, exact = constant.Int64Val(src)
	case constant.Float:
		n, exact = constant.Float64Val(src)
	case constant.Complex:
		re := untyp.extractNumber(constant.Real(src), t)
		im := untyp.extractNumber(constant.Imag(src), t)
		rfloat := r.ValueOf(re).Convert(TypeOfFloat64).Interface().(float64)
		ifloat := r.ValueOf(im).Convert(TypeOfFloat64).Interface().(float64)
		n = complex(rfloat, ifloat)
		exact = true
	default:
		Errorf("cannot convert untyped constant %v to <%v>", untyp, t)
		return nil
	}
	// allow inexact conversions to float64 and complex128:
	// floating point is intrinsically inexact, and Go compiler allows them too
	if !exact && src.Kind() == constant.Int {
		Errorf("untyped constant %v overflows <%v>", untyp, t)
		return nil
	}
	return n
}

// convertLiteralCheckOverflow converts a literal to type t and returns the converted value.
// panics if the conversion overflows the given type
func convertLiteralCheckOverflow(src interface{}, t r.Type) interface{} {
	vsrc := r.ValueOf(src)
	vdst := vsrc.Convert(t)

	k1, k2 := vsrc.Kind(), vdst.Kind()
	if k1 == k2 {
		return vdst.Interface() // no conversion happened
	}
	c1, c2 := kindToCategory(k1), kindToCategory(k2)
	if c2 == r.Int || c2 == r.Uint {
		if c1 == r.Float64 || c1 == r.Complex128 {
			// float-to-integer conversion. check for truncation
			t1 := ValueType(vsrc)
			vback := vdst.Convert(t1)
			if src != vback.Interface() {
				Errorf("constant %v truncated to %v", src, t)
				return nil
			}
		} else {
			// integer-to-integer conversion. convert back and compare the interfaces for overflows
			t1 := ValueType(vsrc)
			vback := vdst.Convert(t1)
			if src != vback.Interface() {
				Errorf("constant %v overflows %v", src, t)
				return nil
			}
		}
	}
	return vdst.Interface()
}

// Set sets the expression value to the given (typed or untyped) constant
func (e *Expr) Set(x I) {
	e.Lit.Set(x)
	e.Types = nil
	e.Fun = nil
	e.IsNil = x == nil
}

// Set sets the Lit to the given typed constant
func (e *Lit) Set(x I) {
	t := r.TypeOf(x)
	if !isLiteral(x) {
		Errorf("cannot set Lit to non-literal value %v <%v>", x, t)
	}
	e.Type = t
	e.Value = x
}

// WithFun ensures that Expr.Fun is a closure that will return the expression result:
//
// if Expr is an untyped constant, WithFun converts the constant to its default type (panics on overflows),
//    then sets Expr.Fun to a closure that will return such constant.
// if Expr is a typed constant, WithFun sets Expr.Fun to a closure that will return such constant.
// if Expr is not a constant, WithFun does nothing (Expr.Fun must be set already)
func (e *Expr) WithFun() {
	if !e.Const() {
		return
	}
	var fun I
again:
	value := e.Value
	t := e.Type
	if t != r.TypeOf(value) {
		Errorf("internal error: constant %v <%v> was assumed to have type <%v>", value, r.TypeOf(value), e.Type)
	}
	switch x := value.(type) {
	case nil:
		fun = XVNil
	case bool:
		if x {
			fun = ITrue
		} else {
			fun = IFalse
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
	case UntypedLit:
		e.ConstTo(x.DefaultType())
		goto again
	default:
		v := r.ValueOf(x)
		fun = func(env *Env) r.Value {
			return v
		}
	}
	e.Fun = fun
}
