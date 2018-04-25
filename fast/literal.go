/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * literal.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/constant"
	"go/token"
	"math/big"
	r "reflect"
	"unsafe"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
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
	return c.exprUntypedLit(kind, obj)
}

func constantKindToUntypedLitKind(ckind constant.Kind) r.Kind {
	ret := r.Invalid
	switch ckind {
	case constant.Bool:
		ret = r.Bool
	case constant.Int:
		ret = r.Int // actually ambiguous, could be a rune - thus r.Int32
	case constant.Float:
		ret = r.Float64
	case constant.Complex:
		ret = r.Complex128
	case constant.String:
		ret = r.String
	}
	return ret
}

func isLiteral(x interface{}) bool {
	if x == nil {
		return true
	}
	rtype := r.TypeOf(x)
	switch KindToCategory(rtype.Kind()) {
	case r.Bool, r.Int, r.Uint, r.Float64, r.Complex128, r.String:
		return true
	}
	_, ok := x.(UntypedLit)
	return ok
}

func isLiteralNumber(x I, n int64) bool {
	if x == nil {
		return false
	}
	v := r.ValueOf(x)
	switch KindToCategory(v.Kind()) {
	case r.Bool:
		return false
	case r.Int:
		return v.Int() == n
	case r.Uint:
		u := v.Uint()
		if n >= 0 {
			return u == uint64(n)
		}
		// n == -1 means "unsigned integer equals its maximum value"
		// similarly, n == -2 means "unsigned integer equals its maximum value minus 1"
		// and so on...
		un := r.ValueOf(n).Convert(v.Type()).Uint()
		return u == un
	case r.Float64:
		return v.Float() == float64(n)
	case r.Complex128:
		return v.Complex() == complex(float64(n), 0)
	case r.String:
		return false
	}
	// no luck yet... try harder
	switch x := x.(type) {
	case r.Value:
		return false
	case UntypedLit:
		return x.IsLiteralNumber(n)
	}
	Errorf("isLiteralNumber: unexpected literal type %v <%v>", x, r.TypeOf(x))
	return false
}

func (untyp *UntypedLit) IsLiteralNumber(n int64) bool {
	val := untyp.Val
	switch val.Kind() {
	case constant.Int:
		m, exact := constant.Int64Val(val)
		return exact && m == n
	case constant.Float:
		m, exact := constant.Float64Val(val)
		return exact && float64(int64(m)) == m && int64(m) == n
	case constant.Complex:
		m, exact := constant.Float64Val(constant.Imag(val))
		if !exact || m != 0.0 {
			return false
		}
		m, exact = constant.Float64Val(constant.Real(val))
		return exact && float64(int64(m)) == m && int64(m) == n
	default:
		return false
	}
}

// ================================= ConstTo =================================

// ConstTo checks that a constant Expr can be used as the given type.
// panics if not constant, or if Expr is a typed constant of different type
// actually performs type conversion (and subsequent overflow checks) ONLY on untyped constants.
func (e *Expr) ConstTo(t xr.Type) I {
	if !e.Const() {
		Errorf("internal error: expression is not constant, use Expr.To() instead of Expr.ConstTo() to convert from <%v> to <%v>", e.Type, t)
	}
	val, fun := e.Lit.ConstTo2(t)
	if fun != nil {
		// no longer a constant
		e.Lit.Value = nil
		e.Fun = fun
	} else if e.Fun != nil {
		// e.Fun is no longer valid, recompute it
		e.WithFun()
	}
	return val
}

// ConstTo checks that a Lit can be used as the given type.
// panics if Lit is a typed constant of different type
// actually performs type conversion (and subsequent overflow checks) ONLY on untyped constants.
func (lit *Lit) ConstTo(t xr.Type) I {
	val, _ := lit.ConstTo2(t)
	return val
}

func (lit *Lit) ConstTo2(t xr.Type) (I, func(*Env) r.Value) {
	value := lit.Value
	// Debugf("Lit.ConstTo(): converting constant %v <%v> (stored as <%v>) to <%v>", value, TypeOf(value), lit.Type, t)
	if t == nil {
		// only literal nil has type nil
		if value != nil {
			Errorf("cannot convert constant %v <%v> to <nil>", value, lit.Type)
		}
		return nil, nil
	}
	// stricter than t == lit.Type
	tfrom := lit.Type
	if tfrom != nil && t != nil && tfrom.IdenticalTo(t) {
		return value, nil
	}
	switch x := value.(type) {
	case UntypedLit:
		val, fun := x.ConstTo2(t)
		lit.Type = t
		lit.Value = val
		// Debugf("Lit.ConstTo(): converted untyped constant %v to %v <%v> (stored as <%v>)", x, val, TypeOf(val), t)
		return val, fun
	case nil:
		// literal nil can only be converted to nillable types
		if IsNillableKind(t.Kind()) {
			lit.Type = t
			return nil, nil
			// lit.Value = r.Zero(t).Interface()
			// return lit.Value
		}
	}
	if tfrom != nil && t != nil && (tfrom.AssignableTo(t) || t.Kind() == r.Interface && tfrom.Implements(t)) {
		lit.Type = t
		// FIXME: use (*Comp).Converter(), requires a *Comp parameter
		lit.Value = r.ValueOf(value).Convert(t.ReflectType()).Interface()
		return lit.Value, nil
	}
	Errorf("cannot convert typed constant %v <%v> to <%v>%s", value, lit.Type, t, interfaceMissingMethod(lit.Type, t))
	return nil, nil
}

// ConstTo checks that an UntypedLit can be used as the given type.
// performs actual untyped -> typed conversion and subsequent overflow checks.
// returns the constant converted to given type
func (untyp *UntypedLit) ConstTo(t xr.Type) I {
	val, _ := untyp.ConstTo2(t)
	return val
}

func (untyp *UntypedLit) ConstTo2(t xr.Type) (I, func(*Env) r.Value) {
	val := untyp.Val
	var ret I
	var fun func(*Env) r.Value
again:
	switch t.Kind() {
	case r.Bool:
		if val.Kind() == constant.Bool {
			ret = constant.BoolVal(val)
		}
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
		r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr,
		r.Float32, r.Float64, r.Complex64, r.Complex128:

		n := untyp.extractNumber(val, t)
		return convertLiteralCheckOverflow(n, t), nil
	case r.Interface:
		// this can happen too... for example in "var foo interface{} = 7"
		// and it requires to convert the untyped constant to its default type.
		// obviously, untyped constants can only implement empty interfaces
		if t.NumMethod() == 0 {
			t = untyp.DefaultType()
			goto again
		}
	case r.Slice:
		// https://golang.org/ref/spec#String_literals states:
		//
		// 4. Converting a value of a string type to a slice of bytes type
		// yields a slice whose successive elements are the bytes of the string.
		//
		// 5. Converting a value of a string type to a slice of runes type
		// yields a slice containing the individual Unicode code points of the string.
		if val.Kind() == constant.String {
			s := UnescapeString(val.ExactString())
			switch t.Elem().Kind() {
			case r.Uint8:
				ret = []byte(s)
			case r.Int32:
				ret = []rune(s)
			}
		}
	case r.String:
		switch val.Kind() {
		case constant.String:
			// untyped string -> string
			ret = UnescapeString(val.ExactString())
		case constant.Int:
			// https://golang.org/ref/spec#String_literals states:
			//
			// 1. Converting a signed or unsigned integer value to a string type yields
			// a string containing the UTF-8 representation of the integer.
			// Values outside the range of valid Unicode code points are converted to "\uFFFD".

			i, exact := constant.Int64Val(val)
			if exact {
				ret = string(i)
			} else {
				ret = "\uFFFD"
			}
		}
	case r.Ptr:
		ret, fun = untyp.toMathBig(t)
	}
	if ret == nil {
		Errorf("cannot convert untyped constant %v to <%v>", untyp, t)
		return nil, nil
	}
	v := r.ValueOf(ret)
	if v.Type() != t.ReflectType() {
		ret = v.Convert(t.ReflectType())
	}
	return ret, fun
}

// EXTENSION: conversion from untyped constant to big.Int, bit.Rat, big.Float
func (untyp *UntypedLit) toMathBig(t xr.Type) (I, func(*Env) r.Value) {
	var fun func(*Env) r.Value
	var ret I
	if k := untyp.Val.Kind(); k == constant.Int || k == constant.Float {
		switch t.ReflectType() {
		case rtypeOfPtrBigInt:
			if a := untyp.BigInt(); a != nil {
				ret = a
				fun = func(*Env) r.Value {
					var b big.Int
					// make a copy of a at every evaluation
					b.Set(a)
					return r.ValueOf(&b)
				}
			}
		case rtypeOfPtrBigRat:
			if a := untyp.BigRat(); a != nil {
				ret = a
				fun = func(*Env) r.Value {
					var b big.Rat
					// make a copy of a at every evaluation
					b.Set(a)
					return r.ValueOf(&b)
				}
			}
		case rtypeOfPtrBigFloat:
			if a := untyp.BigFloat(); a != nil {
				ret = a
				fun = func(*Env) r.Value {
					var b big.Float
					// make a copy of a at every evaluation
					b.Set(a)
					return r.ValueOf(&b)
				}
			}
		}
	}
	return ret, fun
}

func (untyp *UntypedLit) BigInt() *big.Int {
	var b big.Int
	var ret *big.Int

	if i, exact := untyp.Int64(); exact {
		ret = b.SetInt64(i)
	} else if n, exact := untyp.Uint64(); exact {
		ret = b.SetUint64(n)
	} else if i := untyp.rawBigInt(); i != nil {
		ret = b.Set(i)
	} else if r := untyp.rawBigRat(); r != nil {
		if !r.IsInt() {
			return nil
		}
		ret = b.Set(r.Num())
	} else if f := untyp.rawBigFloat(); f != nil {
		if !f.IsInt() {
			return nil
		}
		if i, acc := f.Int(&b); acc == big.Exact {
			if i != &b {
				b.Set(i)
			}
			ret = &b
		}
	}
	if ret == nil {
		// no luck... try to go through string representation
		s := untyp.Val.ExactString()
		if _, ok := b.SetString(s, 0); ok {
			ret = &b
		}
	}
	return ret
}

func (untyp *UntypedLit) BigRat() *big.Rat {
	var b big.Rat
	var ret *big.Rat

	if i, exact := untyp.Int64(); exact {
		ret = b.SetInt64(i)
	} else if i := untyp.rawBigInt(); i != nil {
		ret = b.SetInt(i)
	} else if r := untyp.rawBigRat(); r != nil {
		ret = b.Set(r)
	} else if f := untyp.rawBigFloat(); f != nil {
		if f.IsInt() {
			if i, acc := f.Int(nil); acc == big.Exact {
				ret = b.SetInt(i)
			}
		}
	}

	if ret == nil {
		// no luck... try to go through string representation
		s := untyp.Val.ExactString()
		_, ok := b.SetString(s)
		if ok {
			ret = &b
		}
	}
	return ret
}

func (untyp *UntypedLit) BigFloat() *big.Float {
	var b big.Float
	var ret *big.Float

	if i, exact := untyp.Int64(); exact {
		ret = b.SetInt64(i)
		// Debugf("UntypedLit.BigFloat(): converted int64 %v to *big.Float %v", i, b)
	} else if f, exact := untyp.Float64(); exact {
		ret = b.SetFloat64(f)
		// Debugf("UntypedLit.BigFloat(): converted float64 %v to *big.Float %v", f, b)
	} else if i := untyp.rawBigInt(); i != nil {
		ret = b.SetInt(i)
		// Debugf("UntypedLit.BigFloat(): converted *big.Int %v to *big.Float %v", *i, b)
	} else if r := untyp.rawBigRat(); r != nil {
		ret = b.SetRat(r)
		// Debugf("UntypedLit.BigFloat(): converted *big.Rat %v to *big.Float %v", *r, b)
	} else if f := untyp.rawBigFloat(); f != nil {
		ret = b.Set(f)
		// Debugf("UntypedLit.BigFloat(): converted *big.Float %v to *big.Float %v", *f, b)
	}

	if ret == nil {
		// no luck... try to go through string representation
		s := untyp.Val.ExactString()
		snum, sden := Split2(s, '/')
		_, ok := b.SetString(snum)
		if ok && len(sden) != 0 {
			var b2 big.Float
			if _, ok = b2.SetString(sden); ok {
				b.Quo(&b, &b2)
			}
		}
		if ok {
			ret = &b
			// Debugf("UntypedLit.BigFloat(): converted constant.Value %v %v to *big.Float %v", untyp.Val.Kind(), s, b)
		}
	}
	return ret
}

func (untyp *UntypedLit) Int64() (int64, bool) {
	if c := untyp.Val; c.Kind() == constant.Int {
		return constant.Int64Val(c)
	}
	return 0, false
}

func (untyp *UntypedLit) Uint64() (uint64, bool) {
	if c := untyp.Val; c.Kind() == constant.Int {
		return constant.Uint64Val(c)
	}
	return 0, false
}

func (untyp *UntypedLit) Float64() (float64, bool) {
	if c := untyp.Val; c.Kind() == constant.Float {
		return constant.Float64Val(c)
	}
	return 0, false
}

func (untyp *UntypedLit) rawBigInt() *big.Int {
	if untyp.Val.Kind() != constant.Int {
		return nil
	}
	v := r.ValueOf(untyp.Val)
	if v.Kind() == r.Struct {
		v = v.Field(0)
	}
	if v.Type() != r.TypeOf((*big.Int)(nil)) {
		return nil
	}
	return (*big.Int)(unsafe.Pointer(v.Pointer()))
}

func (untyp *UntypedLit) rawBigRat() *big.Rat {
	if untyp.Val.Kind() != constant.Float {
		return nil
	}
	v := r.ValueOf(untyp.Val)
	if v.Kind() == r.Struct {
		v = v.Field(0)
	}
	if v.Type() != r.TypeOf((*big.Rat)(nil)) {
		return nil
	}
	return (*big.Rat)(unsafe.Pointer(v.Pointer()))
}

func (untyp *UntypedLit) rawBigFloat() *big.Float {
	if untyp.Val.Kind() != constant.Float {
		return nil
	}
	v := r.ValueOf(untyp.Val)
	if v.Kind() == r.Struct {
		v = v.Field(0)
	}
	if v.Type() != r.TypeOf((*big.Float)(nil)) {
		return nil
	}
	return (*big.Float)(unsafe.Pointer(v.Pointer()))
}

// ================================= DefaultType =================================

// DefaultType returns the default type of an expression.
func (e *Expr) DefaultType() xr.Type {
	if e.Untyped() {
		return e.Lit.DefaultType()
	}
	return e.Type
}

// DefaultType returns the default type of a constant.
func (lit *Lit) DefaultType() xr.Type {
	switch x := lit.Value.(type) {
	case UntypedLit:
		return x.DefaultType()
	default:
		return lit.Type
	}
}

// DefaultType returns the default type of an untyped constant.
func (untyp *UntypedLit) DefaultType() xr.Type {
	switch untyp.Kind {
	case r.Bool, r.Int32, r.Int, r.Uint, r.Float64, r.Complex128, r.String:
		if basicTypes := untyp.BasicTypes; basicTypes == nil {
			Errorf("UntypedLit.DefaultType(): malformed untyped constant %v, has nil BasicTypes!", untyp)
			return nil
		} else {
			return (*basicTypes)[untyp.Kind]
		}

	default:
		Errorf("unexpected untyped constant %v, its default type is not known", untyp)
		return nil
	}
}

// ======================= utilities for ConstTo and ConstToDefaultType =======================

// extractNumber converts the untyped constant src to an integer, float or complex.
// panics if src has different kind from constant.Int, constant.Float and constant.Complex
// the receiver (untyp UntypedLit) and the second argument (t reflect.Type) are only used to pretty-print the panic error message
func (untyp *UntypedLit) extractNumber(src constant.Value, t xr.Type) interface{} {
	var n interface{}
	cat := KindToCategory(t.Kind())
	var exact bool
	switch src.Kind() {
	case constant.Int:
		switch cat {
		case r.Int:
			n, exact = constant.Int64Val(src)
		case r.Uint:
			n, exact = constant.Uint64Val(src)
		default:
			n, exact = constant.Int64Val(src)
			if !exact {
				n, exact = constant.Uint64Val(src)
			}
		}
	case constant.Float:
		n, exact = constant.Float64Val(src)
	case constant.Complex:
		re := untyp.extractNumber(constant.Real(src), t)
		im := untyp.extractNumber(constant.Imag(src), t)
		rfloat := r.ValueOf(re).Convert(TypeOfFloat64).Float()
		ifloat := r.ValueOf(im).Convert(TypeOfFloat64).Float()
		n = complex(rfloat, ifloat)
		exact = true
	default:
		Errorf("cannot convert untyped constant %v to <%v>", untyp, t)
		return nil
	}
	// allow inexact conversions to float64 and complex128:
	// floating point is intrinsically inexact, and Go compiler allows them too
	if !exact && (cat == r.Int || cat == r.Uint) {
		Errorf("untyped constant %v overflows <%v>", untyp, t)
		return nil
	}
	return n
}

// convertLiteralCheckOverflow converts a literal to type t and returns the converted value.
// panics if the conversion overflows the given type
func convertLiteralCheckOverflow(src interface{}, to xr.Type) interface{} {
	v := r.ValueOf(src)
	rto := to.ReflectType()
	vto := ConvertValue(v, rto)

	k, kto := v.Kind(), vto.Kind()
	if k == kto {
		return vto.Interface() // no numeric conversion happened
	}
	c, cto := KindToCategory(k), KindToCategory(kto)
	if cto == r.Int || cto == r.Uint {
		if c == r.Float64 || c == r.Complex128 {
			// float-to-integer conversion. check for truncation
			t1 := ValueType(v)
			vback := ConvertValue(vto, t1)
			if src != vback.Interface() {
				Errorf("constant %v truncated to %v", src, to)
				return nil
			}
		} else {
			// integer-to-integer conversion. convert back and compare the interfaces for overflows
			t1 := ValueType(v)
			vback := vto.Convert(t1)
			if src != vback.Interface() {
				Errorf("constant %v overflows <%v>", src, to)
				return nil
			}
		}
	}
	return vto.Interface()
}

// SetTypes sets the expression result types
func (e *Expr) SetTypes(tout []xr.Type) {
	switch len(tout) {
	case 0:
		e.Type = nil
		e.Types = tout
	case 1:
		e.Type = tout[0]
		e.Types = nil
	default:
		e.Type = tout[0]
		e.Types = tout
	}
}

/* used?

// Set sets the expression value to the given (typed or untyped) constant
func (e *Expr) Set(x I) {
	e.Lit.Set(x)
	e.Types = nil
	e.Fun = nil
	e.IsNil = x == nil
}

// Set sets the Lit to the given typed constant
func (lit *Lit) Set(x I) {
	t := TypeOf(x)
	if !isLiteral(x) {
		Errorf("cannot set Lit to non-literal value %v <%v>", x, t)
	}
	lit.Type = t
	lit.Value = x
}
*/

// To checks that an Expr can be used as (i.e. is assignable to) the given type,
// and converts Expr to the given type.
// panics if Expr has an incompatible type.
func (e *Expr) To(c *Comp, t xr.Type) {
	if e.Const() {
		e.ConstTo(t)
		return
	}
	if e.Type.IdenticalTo(t) {
		return
	}
	if !e.Type.AssignableTo(t) {
		c.Errorf("cannot use <%v> as <%v>", e.Type, t)
	}
	k := e.Type.Kind()
	if IsOptimizedKind(k) {
		if k == t.Kind() {
			// same optimized representation
			e.Type = t
			return
		} else if t.Kind() == r.Interface {
			e.Fun = e.AsX1()
			e.Type = t
			return
		}
		c.Errorf("internal error: cannot use <%v> as <%v> (should not happen, <%v> is assignable to <%v>", e.Type, t, e.Type, t)
	}
	fun := e.AsX1()
	rtype := t.ReflectType()
	zero := r.Zero(rtype)

	if conv := c.Converter(e.Type, t); conv == nil {
		e.Fun = func(env *Env) r.Value {
			v := fun(env)
			if !v.IsValid() {
				v = zero
			}
			return v
		}
	} else {
		e.Fun = func(env *Env) r.Value {
			v := fun(env)
			if !v.IsValid() {
				v = zero
			} else {
				v = conv(v)
			}
			return v
		}
	}
	e.Type = t
}

// WithFun ensures that Expr.Fun is a closure that will return the expression result:
//
// if Expr is an untyped constant, WithFun converts the constant to its default type (panics on overflows),
//    then sets Expr.Fun to a closure that will return such constant.
// if Expr is a typed constant, WithFun sets Expr.Fun to a closure that will return such constant.
// if Expr is not a constant, WithFun does nothing (Expr.Fun must be set already)
func (e *Expr) WithFun() I {
	if !e.Const() {
		return e.Fun
	}
	var fun I
again:
	value := e.Value
	v := r.ValueOf(value)
	t := e.Type
	if t == nil {
		e.Fun = eNil
		return eNil
	}
	if value == nil {
		if !IsNillableKind(t.Kind()) {
			Errorf("internal error: constant of type <%v> cannot be nil", t)
		}
		zero := r.Zero(t.ReflectType())
		fun = func(*Env) r.Value {
			return zero
		}
		e.Fun = fun
		return fun
	}
	rtactual := r.TypeOf(value)
	rtexpected := t.ReflectType()
	if rtexpected != rtactual {
		if rtexpected.Kind() == r.Interface && rtactual.Implements(rtexpected) {
			v = v.Convert(rtexpected)
		} else {
			Errorf("internal error: constant %v <%v> was assumed to have type <%v>", value, rtactual, rtexpected)
		}
	}
	switch v.Kind() {
	case r.Invalid:
		fun = eNil
	case r.Bool:
		if v.Bool() {
			fun = eTrue
		} else {
			fun = eFalse
		}
	case r.Int:
		x := int(v.Int())
		fun = func(env *Env) int {
			return x
		}
	case r.Int8:
		x := int8(v.Int())
		fun = func(env *Env) int8 {
			return x
		}
	case r.Int16:
		x := int16(v.Int())
		fun = func(env *Env) int16 {
			return x
		}
	case r.Int32:
		x := int32(v.Int())
		fun = func(env *Env) int32 {
			return x
		}
	case r.Int64:
		x := v.Int()
		fun = func(env *Env) int64 {
			return x
		}
	case r.Uint:
		x := uint(v.Uint())
		fun = func(env *Env) uint {
			return x
		}
	case r.Uint8:
		x := uint8(v.Uint())
		fun = func(env *Env) uint8 {
			return x
		}
	case r.Uint16:
		x := uint16(v.Uint())
		fun = func(env *Env) uint16 {
			return x
		}
	case r.Uint32:
		x := uint32(v.Uint())
		fun = func(env *Env) uint32 {
			return x
		}
	case r.Uint64:
		x := v.Uint()
		fun = func(env *Env) uint64 {
			return x
		}
	case r.Uintptr:
		x := uintptr(v.Uint())
		fun = func(env *Env) uintptr {
			return x
		}
	case r.Float32:
		x := float32(v.Float())
		fun = func(env *Env) float32 {
			return x
		}
	case r.Float64:
		x := v.Float()
		fun = func(env *Env) float64 {
			return x
		}
	case r.Complex64:
		x := complex64(v.Complex())
		fun = func(env *Env) complex64 {
			return x
		}
	case r.Complex128:
		x := v.Complex()
		fun = func(env *Env) complex128 {
			return x
		}
	case r.String:
		x := v.String()
		fun = func(env *Env) string {
			return x
		}
	default:
		if t.ReflectType() == rtypeOfUntypedLit {
			e.ConstTo(e.DefaultType())
			goto again
		}
		fun = func(env *Env) r.Value {
			return v
		}
	}
	e.Fun = fun
	return fun
}
