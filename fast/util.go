/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * identifier.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	"go/constant"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/base/reflect"
	"github.com/cosmos72/gomacro/base/untyped"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func eFalse(*Env) bool {
	return false
}

func eTrue(*Env) bool {
	return true
}

func eNilR(*Env) xr.Value {
	return xr.Value{}
}

func eXVNone(*Env) (xr.Value, []xr.Value) {
	return None, nil
}

func nop() {
}

var valueOfNopFunc = xr.ValueOf(nop)

// opaqueType returns an xr.Type corresponding to rtype but without fields or methods and with the given pkgpath
func (g *CompGlobals) opaqueType(rtype r.Type, pkgpath string) xr.Type {
	return g.opaqueNamedType(rtype, rtype.Name(), pkgpath)
}

// opaqueNamedType returns an xr.Type corresponding to rtype but without fields or methods and with the given name and pkgpath
func (g *CompGlobals) opaqueNamedType(rtype r.Type, name string, pkgpath string) xr.Type {
	v := g.Universe
	switch k := rtype.Kind(); k {
	case xr.Ptr:
		telem := g.opaqueType(rtype.Elem(), pkgpath)
		t := v.PtrTo(telem)
		v.ReflectTypes[rtype] = t
		return t
	case xr.Struct:
		break
	default:
		g.Errorf("internal error: unimplemented opaqueNamedType for kind=%v, expecting kind=Struct", k)
	}
	t := v.NamedOf(name, pkgpath)
	t.SetUnderlying(v.TypeOf(struct{}{}))
	t.UnsafeForceReflectType(rtype)
	v.ReflectTypes[rtype] = t // also cache Type in g.Universe.ReflectTypes
	// g.Debugf("initialized opaque type %v <%v> <%v>", t.Kind(), t.GoType(), t.ReflectType())
	return t
}

func asIdent(node ast.Expr) *ast.Ident {
	ident, _ := node.(*ast.Ident)
	return ident
}

func (e *Expr) TryAsPred() (value bool, fun func(*Env) bool, err bool) {
	if e.Untyped() {
		untyp := e.Value.(UntypedLit)
		if untyp.Kind != untyped.Bool {
			return false, nil, true
		}
		return constant.BoolVal(untyp.Val), nil, false
	}
	if e.Type.Kind() != r.Bool {
		return false, nil, true
	}
	if e.Const() {
		v := xr.ValueOf(e.Value)
		return v.Bool(), nil, false
	}
	switch fun := e.Fun.(type) {
	case func(*Env) bool:
		return false, fun, false
	case func(*Env) (xr.Value, []xr.Value):
		e.CheckX1()
		return false, func(env *Env) bool {
			ret, _ := fun(env)
			return ret.Bool()
		}, false
	default:
		fun1 := e.AsX1()
		return false, func(env *Env) bool {
			return fun1(env).Bool()
		}, false
	}
}

func (c *Comp) invalidPred(node ast.Expr, x *Expr) Stmt {
	return c.badPred("invalid", node, x)
}

func (c *Comp) badPred(reason string, node ast.Expr, x *Expr) Stmt {
	var t xr.Type = nil
	if x.NumOut() != 0 {
		t = x.Out(0)
	}
	c.Errorf("%s boolean predicate, expecting <bool> expression, found <%v>: %v",
		reason, t, node)
	return nil
}

var negativeShiftAmount = fmt.Errorf("runtime error: negative shift amount")

/*
 * used by Comp.Shl(), Comp.Shr(), Comp.varSh{l,r}Expr() and Comp.placeSh{l,r}Expr().
 * If e returns a signed integer type, the returned function
 * panics at runtime when the integer is negative.
 * This reproduces the behaviour of shift by signed integer introduced in Go 1.13
 */
func (e *Expr) AsUint64() func(*Env) uint64 {
	if e == nil {
		output.Errorf("internal error in Expr.AsUint64: receiver is nil")
	} else if e.Const() {
		n, ok := constAsUint64(e.Value)
		if !ok {
			output.Errorf("invalid shift amount: %v // %v",
				e.Value, e.Type)
		}
		return func(*Env) uint64 {
			return n
		}
	}
	if e.NumOut() == 0 {
		output.Errorf("expression returns no values, cannot convert to func(env *Env) uint64")
		return nil
	} else if e.NumOut() > 1 {
		output.Warnf("expression returns %d values, using only the first one: %v", e.NumOut(), e.Types)
	}
	t := e.Type
	if t == nil {
		t = e.Types[0]
	}
	cat := reflect.Category(t.Kind())
	if cat != r.Int && cat != r.Uint {
		output.Errorf("expression returns %v, cannot convert to func(env *Env) uint64", t)
		return nil
	}
	var ret func(*Env) uint64
	switch fun := e.Fun.(type) {
	case func(*Env) xr.Value:
		if cat == r.Int {
			ret = func(env *Env) uint64 {
				i := fun(env).Int()
				if i < 0 {
					panic(negativeShiftAmount)
				}
				return uint64(i)
			}
		} else {
			ret = func(env *Env) uint64 {
				return fun(env).Uint()
			}
		}
	case func(*Env) (xr.Value, []xr.Value):
		if cat == r.Int {
			ret = func(env *Env) uint64 {
				v, _ := fun(env)
				i := v.Int()
				if i < 0 {
					panic(negativeShiftAmount)
				}
				return uint64(i)
			}
		} else {
			ret = func(env *Env) uint64 {
				v, _ := fun(env)
				return v.Uint()
			}
		}
	case func(*Env) int:
		ret = func(env *Env) uint64 {
			i := fun(env)
			if i < 0 {
				panic(negativeShiftAmount)
			}
			return uint64(i)
		}
	case func(*Env) int8:
		ret = func(env *Env) uint64 {
			i := fun(env)
			if i < 0 {
				panic(negativeShiftAmount)
			}
			return uint64(i)
		}
	case func(*Env) int16:
		ret = func(env *Env) uint64 {
			i := fun(env)
			if i < 0 {
				panic(negativeShiftAmount)
			}
			return uint64(i)
		}
	case func(*Env) int32:
		ret = func(env *Env) uint64 {
			i := fun(env)
			if i < 0 {
				panic(negativeShiftAmount)
			}
			return uint64(i)
		}
	case func(*Env) int64:
		ret = func(env *Env) uint64 {
			i := fun(env)
			if i < 0 {
				panic(negativeShiftAmount)
			}
			return uint64(i)
		}
	case func(*Env) uint:
		ret = func(env *Env) uint64 {
			return uint64(fun(env))
		}
	case func(*Env) uint8:
		ret = func(env *Env) uint64 {
			return uint64(fun(env))
		}
	case func(*Env) uint16:
		ret = func(env *Env) uint64 {
			return uint64(fun(env))
		}
	case func(*Env) uint32:
		ret = func(env *Env) uint64 {
			return uint64(fun(env))
		}
	case func(*Env) uint64:
		return fun
	case func(*Env) uintptr:
		ret = func(env *Env) uint64 {
			return uint64(fun(env))
		}
	default:
		output.Errorf("unsupported function type, cannot convert to func(*Env) uint64: %v <%T>", fun, fun)
	}
	return ret
}

func constAsUint64(any I) (uint64, bool) {
	v := xr.ValueOf(any)
	if !v.IsValid() {
		return 0, false
	}
	switch reflect.Category(v.Kind()) {
	case xr.Uint:
		return v.Uint(), true
	case xr.Int:
		if i := v.Int(); i >= 0 {
			return uint64(i), true
		}
	}
	return 0, false
}

func (e *Expr) AsX() func(*Env) {
	if e == nil || e.Const() {
		return nil
	}
	return funAsX(e.Fun)
}

func funAsX(any I) func(*Env) {
	switch fun := any.(type) {
	case nil:
	case func(*Env):
		return fun
	case func(*Env) xr.Value:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) (xr.Value, []xr.Value):
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) bool:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) int:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) int8:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) int16:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) int32:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) int64:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uint:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uint8:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uint16:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uint32:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uint64:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) float32:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) float64:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) complex64:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) complex128:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) string:
		return func(env *Env) {
			fun(env)
		}
	default:
		output.Errorf("unsupported function type, cannot convert to func(*Env): %v <%T>", any, any)
	}
	return nil
}

// CheckX1() panics if given expression cannot be used in single-value context,
// for example because it returns no value at all.
// It just prints a warning if expression returns multiple values.
func (e *Expr) CheckX1() {
	if e != nil && e.Const() {
		return
	}
	if e == nil || e.NumOut() == 0 {
		output.Errorf("expression returns no values, cannot convert to func(env *Env) xr.Value")
		return
	} else if e.NumOut() > 1 {
		output.Warnf("expression returns %d values, using only the first one: %v", e.NumOut(), e.Types)
	}
}

func (e *Expr) AsX1() func(*Env) xr.Value {
	if e == nil {
		return eNilR
	}
	if e.Const() {
		return valueAsX1(e.Value, e.Type, COptDefaults)
	}
	e.CheckX1()
	return funAsX1(e.Fun, e.Type)
}

func (e *Expr) AsXV(opts CompileOptions) func(*Env) (xr.Value, []xr.Value) {
	if e == nil {
		return eXVNone
	}
	if e.Const() {
		return valueAsXV(e.Value, e.Type, opts)
	}
	return funAsXV(e.Fun, e.Type)
}

func valueAsX1(any I, t xr.Type, opts CompileOptions) func(*Env) xr.Value {
	convertuntyped := opts&COptKeepUntyped == 0
	untyp, untyped := any.(UntypedLit)
	if untyped && convertuntyped {
		if t == nil || t.ReflectType() == rtypeOfUntypedLit {
			t = untyp.DefaultType()
		}
		// Debugf("late conversion of untyped constant %v <%T> to <%v>", untyp, untyp, t)
		any = untyp.Convert(t)
	}
	v := xr.ValueOf(any)
	if t != nil {
		rtype := t.ReflectType()
		if !v.IsValid() {
			v = xr.ZeroR(rtype)
		} else if convertuntyped || !untyped {
			v = convert(v, rtype)
		}
	}
	return func(*Env) xr.Value {
		return v
	}
}

func valueAsXV(any I, t xr.Type, opts CompileOptions) func(*Env) (xr.Value, []xr.Value) {
	convertuntyped := opts&COptKeepUntyped == 0
	untyp, untyped := any.(UntypedLit)
	if convertuntyped {
		if untyped {
			if t == nil || t.ReflectType() == rtypeOfUntypedLit {
				t = untyp.DefaultType()
				// Debugf("valueAsXV: late conversion of untyped constant %v <%T> to its default type <%v>", untyp, untyp, t)
			} else {
				// Debugf("valueAsXV: late conversion of untyped constant %v <%T> to <%v>", untyp, untyp, t.ReflectType())
			}
			any = untyp.Convert(t)
		}
	}
	v := xr.ValueOf(any)
	if t != nil {
		rtype := t.ReflectType()
		if reflect.ValueType(v) == nil {
			v = xr.ZeroR(rtype)
		} else if convertuntyped || !untyped {
			v = convert(v, rtype)
		}
	}
	return func(*Env) (xr.Value, []xr.Value) {
		return v, nil
	}
}

func funAsX1(fun I, t xr.Type) func(*Env) xr.Value {
	// output.Debugf("funAsX1() %T -> %v", fun, t)
	var rt r.Type
	if t != nil {
		rt = t.ReflectType()
	}
	switch fun := fun.(type) {
	case nil:
	case func(*Env):
		if fun == nil {
			break
		}
		return func(env *Env) xr.Value {
			fun(env)
			return None
		}
	case func(*Env) xr.Value:
		return fun
	case func(*Env) (xr.Value, []xr.Value):
		return func(env *Env) xr.Value {
			ret, _ := fun(env)
			return ret
		}
	case func(*Env) bool:
		if rt == nil || rt == base.TypeOfBool {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) int:
		if rt == nil || rt == base.TypeOfInt {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) int8:
		if rt == nil || rt == base.TypeOfInt8 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) int16:
		if rt == nil || rt == base.TypeOfInt16 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) int32:
		if rt == nil || rt == base.TypeOfInt32 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) int64:
		if rt == nil || rt == base.TypeOfInt64 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uint:
		if rt == nil || rt == base.TypeOfUint {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uint8:
		if rt == nil || rt == base.TypeOfUint8 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uint16:
		if rt == nil || rt == base.TypeOfUint16 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uint32:
		if rt == nil || rt == base.TypeOfUint32 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uint64:
		if rt == nil || rt == base.TypeOfUint64 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uintptr:
		if rt == nil || rt == base.TypeOfUintptr {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) float32:
		if rt == nil || rt == base.TypeOfFloat32 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) float64:
		if rt == nil || rt == base.TypeOfFloat64 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) complex64:
		if rt == nil || rt == base.TypeOfComplex64 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) complex128:
		if rt == nil || rt == base.TypeOfComplex128 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) string:
		if rt == nil || rt == base.TypeOfString {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *bool:
		if rt == nil || rt == base.TypeOfPtrBool {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *int:
		if rt == nil || rt == base.TypeOfPtrInt {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *int8:
		if rt == nil || rt == base.TypeOfPtrInt8 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *int16:
		if rt == nil || rt == base.TypeOfPtrInt16 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *int32:
		if rt == nil || rt == base.TypeOfPtrInt32 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *int64:
		if rt == nil || rt == base.TypeOfPtrInt64 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *uint:
		if rt == nil || rt == base.TypeOfPtrUint {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *uint8:
		if rt == nil || rt == base.TypeOfPtrUint8 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *uint16:
		if rt == nil || rt == base.TypeOfPtrUint16 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *uint32:
		if rt == nil || rt == base.TypeOfPtrUint32 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *uint64:
		if rt == nil || rt == base.TypeOfPtrUint64 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *uintptr:
		if rt == nil || rt == base.TypeOfPtrUintptr {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *float32:
		if rt == nil || rt == base.TypeOfPtrFloat32 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *float64:
		if rt == nil || rt == base.TypeOfPtrFloat64 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) *complex64:
		if rt == nil || rt == base.TypeOfPtrComplex64 {
			return func(env *Env) xr.Value {
				return xr.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) xr.Value {
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	default:
		output.Errorf("unsupported expression type, cannot convert to func(*Env) xr.Value: %v <%T>", fun, fun)
	}
	return nil
}

func funAsXV(fun I, t xr.Type) func(*Env) (xr.Value, []xr.Value) {
	// output.Debugf("funAsXV() %T -> %v", fun, t)
	var rt r.Type
	if t != nil {
		rt = t.ReflectType()
	}
	switch fun := fun.(type) {
	case nil:
	case func(*Env):
		if fun == nil {
			break
		}
		return func(env *Env) (xr.Value, []xr.Value) {
			fun(env)
			return None, nil
		}
	case func(*Env) xr.Value:
		return func(env *Env) (xr.Value, []xr.Value) {
			return fun(env), nil
		}
	case func(*Env) (xr.Value, []xr.Value):
		return fun
	case func(*Env) bool:
		if rt == nil || rt == base.TypeOfBool {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) int:
		if rt == nil || rt == base.TypeOfInt {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) int8:
		if rt == nil || rt == base.TypeOfInt8 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) int16:
		if rt == nil || rt == base.TypeOfInt16 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) int32:
		if rt == nil || rt == base.TypeOfInt32 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) int64:
		if rt == nil || rt == base.TypeOfInt64 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) uint:
		if rt == nil || rt == base.TypeOfUint {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) uint8:
		if rt == nil || rt == base.TypeOfUint8 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) uint16:
		if rt == nil || rt == base.TypeOfUint16 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) uint32:
		if rt == nil || rt == base.TypeOfUint32 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) uint64:
		if rt == nil || rt == base.TypeOfUint64 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) uintptr:
		if rt == nil || rt == base.TypeOfUintptr {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) float32:
		if rt == nil || rt == base.TypeOfFloat32 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) float64:
		if rt == nil || rt == base.TypeOfFloat64 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) complex64:
		if rt == nil || rt == base.TypeOfComplex64 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) complex128:
		if rt == nil || rt == base.TypeOfComplex128 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) string:
		if rt == nil || rt == base.TypeOfString {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *bool:
		if rt == nil || rt == base.TypeOfPtrBool {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *int:
		if rt == nil || rt == base.TypeOfPtrInt {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *int8:
		if rt == nil || rt == base.TypeOfPtrInt8 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *int16:
		if rt == nil || rt == base.TypeOfPtrInt16 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *int32:
		if rt == nil || rt == base.TypeOfPtrInt32 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *int64:
		if rt == nil || rt == base.TypeOfPtrInt64 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *uint:
		if rt == nil || rt == base.TypeOfPtrUint {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *uint8:
		if rt == nil || rt == base.TypeOfPtrUint8 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *uint16:
		if rt == nil || rt == base.TypeOfPtrUint16 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *uint32:
		if rt == nil || rt == base.TypeOfPtrUint32 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *uint64:
		if rt == nil || rt == base.TypeOfPtrUint64 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *uintptr:
		if rt == nil || rt == base.TypeOfPtrUintptr {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *float32:
		if rt == nil || rt == base.TypeOfPtrFloat32 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *float64:
		if rt == nil || rt == base.TypeOfPtrFloat64 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	case func(*Env) *complex64:
		if rt == nil || rt == base.TypeOfPtrComplex64 {
			return func(env *Env) (xr.Value, []xr.Value) {
				return xr.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (xr.Value, []xr.Value) {
				return convert(xr.ValueOf(fun(env)), rt), nil
			}
		}
	default:
		output.Errorf("unsupported expression, cannot convert to func(*Env) (xr.Value, []xr.Value) : %v <%T>",
			fun, fun)
	}
	return nil
}

func (e *Expr) exprXVAsI() *Expr {
	// Debugf("exprXVAsI() %v -> %v", e.Types, e.Type)
	e.CheckX1()
	if e.NumOut() <= 1 {
		return e
	}
	fun := e.Fun.(func(*Env) (xr.Value, []xr.Value))
	t := e.Type
	var ret I
	switch t.Kind() {
	case xr.Bool:
		ret = func(env *Env) bool {
			v, _ := fun(env)
			return v.Bool()
		}
	case xr.Int:
		ret = func(env *Env) int {
			v, _ := fun(env)
			return int(v.Int())
		}
	case xr.Int8:
		ret = func(env *Env) int8 {
			v, _ := fun(env)
			return int8(v.Int())
		}
	case xr.Int16:
		ret = func(env *Env) int16 {
			v, _ := fun(env)
			return int16(v.Int())
		}
	case xr.Int32:
		ret = func(env *Env) int32 {
			v, _ := fun(env)
			return int32(v.Int())
		}
	case xr.Int64:
		ret = func(env *Env) int64 {
			v, _ := fun(env)
			return v.Int()
		}
	case xr.Uint:
		ret = func(env *Env) uint {
			v, _ := fun(env)
			return uint(v.Uint())
		}
	case xr.Uint8:
		ret = func(env *Env) uint8 {
			v, _ := fun(env)
			return uint8(v.Uint())
		}
	case xr.Uint16:
		ret = func(env *Env) uint16 {
			v, _ := fun(env)
			return uint16(v.Uint())
		}
	case xr.Uint32:
		ret = func(env *Env) uint32 {
			v, _ := fun(env)
			return uint32(v.Uint())
		}
	case xr.Uint64:
		ret = func(env *Env) uint64 {
			v, _ := fun(env)
			return v.Uint()
		}
	case xr.Uintptr:
		ret = func(env *Env) uintptr {
			v, _ := fun(env)
			return uintptr(v.Uint())
		}
	case xr.Float32:
		ret = func(env *Env) float32 {
			v, _ := fun(env)
			return float32(v.Float())
		}
	case xr.Float64:
		ret = func(env *Env) float64 {
			v, _ := fun(env)
			return v.Float()
		}
	case xr.Complex64:
		ret = func(env *Env) complex64 {
			v, _ := fun(env)
			return complex64(v.Complex())
		}
	case xr.Complex128:
		ret = func(env *Env) complex128 {
			v, _ := fun(env)
			return v.Complex()
		}
	case xr.String:
		ret = func(env *Env) string {
			v, _ := fun(env)
			return v.String()
		}
	default:
		ret = func(env *Env) xr.Value {
			v, _ := fun(env)
			return v
		}
	}
	return exprFun(t, ret)
}

func (e *Expr) AsStmt(c *Comp) Stmt {
	if e == nil || e.Const() {
		return nil
	}
	if stmt := c.Jit.AsStmt(e); stmt != nil {
		return stmt
	}
	return funAsStmt(e.Fun)
}

func funAsStmt(fun I) Stmt {
	var ret func(env *Env) (Stmt, *Env)

	switch fun := fun.(type) {
	case nil:
	case func(*Env):
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) xr.Value:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) (xr.Value, []xr.Value):
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) bool:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int8:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int16:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int32:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint8:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint16:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint32:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uintptr:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float32:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex128:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) string:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	default:

		output.Errorf("unsupported expression type, cannot convert to Stmt : %v <%T>",
			fun, fun)
	}
	return ret
}

// funTypeOut returns the first return type of given function
func funTypeOut(fun I) r.Type {
	rt := r.TypeOf(fun)
	if rt == nil || rt.Kind() != r.Func || rt.NumOut() == 0 {
		return nil
	}
	return rt.Out(0)
}

// funTypeOuts returns the return types of given function
func funTypeOuts(fun I) []r.Type {
	rt := r.TypeOf(fun)
	if rt == nil || rt.Kind() != r.Func {
		return []r.Type{rt}
	}
	n := rt.NumOut()
	rts := make([]r.Type, n)
	for i := 0; i < n; i++ {
		rts[i] = rt.Out(i)
	}
	return rts
}

// exprList merges together a list of expressions,
// and returns an expression that evaluates each one
func exprList(list []*Expr, opts CompileOptions) *Expr {
	// skip constant expressions (except the last one)
	var n int
	for i, ni := 0, len(list)-1; i <= ni; i++ {
		// preserve the last expression even if constant
		// because it will be returned to the user
		if i == ni || !list[i].Const() {
			list[n] = list[i]
			n++
		}
	}
	switch n {
	case 0:
		return nil
	case 1:
		return list[0]
	}
	list = list[:n]

	funs := make([]func(*Env), n-1)
	for i := range funs {
		funs[i] = list[i].AsX()
	}
	return &Expr{
		Lit:   Lit{Type: list[n-1].Type},
		Types: list[n-1].Types,
		Fun:   funList(funs, list[n-1], opts),
	}
}

// funList merges together a list of functions,
// and returns a function that evaluates each one
func funList(funs []func(*Env), last *Expr, opts CompileOptions) I {
	var rt r.Type
	if last.Type != nil {
		// keep untyped constants only if requested
		if opts != COptKeepUntyped && last.Untyped() {
			last.ConstTo(last.DefaultType())
		}
		rt = last.Type.ReflectType()
	}
	switch fun := last.WithFun().(type) {
	case nil:
		return func(env *Env) {
			for _, f := range funs {
				f(env)
			}
		}
	case func(*Env):
		return func(env *Env) {
			for _, f := range funs {
				f(env)
			}
			fun(env)
		}
	case func(*Env) xr.Value:
		return func(env *Env) xr.Value {
			for _, f := range funs {
				f(env)
			}
			return fun(env)
		}
	case func(*Env) (xr.Value, []xr.Value):
		return func(env *Env) (xr.Value, []xr.Value) {
			for _, f := range funs {
				f(env)
			}
			return fun(env)
		}
	case func(*Env) bool:
		if rt == nil || rt == base.TypeOfBool {
			return func(env *Env) bool {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) int:
		if rt == nil || rt == base.TypeOfInt {
			return func(env *Env) int {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) int8:
		if rt == nil || rt == base.TypeOfInt8 {
			return func(env *Env) int8 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) int16:
		if rt == nil || rt == base.TypeOfInt16 {
			return func(env *Env) int16 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) int32:
		if rt == nil || rt == base.TypeOfInt32 {
			return func(env *Env) int32 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) int64:
		if rt == nil || rt == base.TypeOfInt64 {
			return func(env *Env) int64 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uint:
		if rt == nil || rt == base.TypeOfUint {
			return func(env *Env) uint {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uint8:
		if rt == nil || rt == base.TypeOfUint8 {
			return func(env *Env) uint8 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uint16:
		if rt == nil || rt == base.TypeOfUint16 {
			return func(env *Env) uint16 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uint32:
		if rt == nil || rt == base.TypeOfUint32 {
			return func(env *Env) uint32 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uint64:
		if rt == nil || rt == base.TypeOfUint64 {
			return func(env *Env) uint64 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) uintptr:
		if rt == nil || rt == base.TypeOfUintptr {
			return func(env *Env) uintptr {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) float32:
		if rt == nil || rt == base.TypeOfFloat32 {
			return func(env *Env) float32 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) float64:
		if rt == nil || rt == base.TypeOfFloat64 {
			return func(env *Env) float64 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) complex64:
		if rt == nil || rt == base.TypeOfComplex64 {
			return func(env *Env) complex64 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) complex128:
		if rt == nil || rt == base.TypeOfComplex128 {
			return func(env *Env) complex128 {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	case func(*Env) string:
		if rt == nil || rt == base.TypeOfString {
			return func(env *Env) string {
				for _, f := range funs {
					f(env)
				}
				return fun(env)
			}
		} else {
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				return convert(xr.ValueOf(fun(env)), rt)
			}
		}
	default:
		switch last.NumOut() {
		case 0:
			fun := last.AsX()
			return func(env *Env) {
				for _, f := range funs {
					f(env)
				}
				fun(env)
			}
		case 1:
			var zero xr.Value
			if rt != nil {
				zero = xr.ZeroR(rt)
			}
			fun := last.AsX1()
			return func(env *Env) xr.Value {
				for _, f := range funs {
					f(env)
				}
				ret := fun(env)
				if !ret.IsValid() {
					ret = zero
				} else if rt != nil && rt != ret.Type() {
					ret = convert(ret, rt)
				}
				return ret
			}
		default:
			var zero []xr.Value
			var rt []r.Type
			for i, t := range last.Types {
				if t != nil {
					rt[i] = t.ReflectType()
					zero[i] = xr.ZeroR(rt[i])
				}
			}
			fun := last.AsXV(opts)
			return func(env *Env) (xr.Value, []xr.Value) {
				for _, f := range funs {
					f(env)
				}
				_, rets := fun(env)
				for i, ret := range rets {
					if !ret.IsValid() {
						rets[i] = zero[i]
					} else if rt != nil && rt[i] != ret.Type() {
						rets[i] = convert(ret, rt[i])
					}
				}
				return rets[0], rets
			}
		}
	}
}

// unwrapBind compiles a conversion from a "mis-typed" bind stored in env.Binds[] as reflect.Value
// into a correctly-typed expression
func unwrapBind(bind *Bind, t xr.Type) *Expr {
	idx := bind.Desc.Index()
	var ret I
	switch t.Kind() {
	case xr.Bool:
		ret = func(env *Env) bool {
			return env.Vals[idx].Bool()
		}
	case xr.Int:
		ret = func(env *Env) int {
			return int(env.Vals[idx].Int())
		}
	case xr.Int8:
		ret = func(env *Env) int8 {
			return int8(env.Vals[idx].Int())
		}
	case xr.Int16:
		ret = func(env *Env) int16 {
			return int16(env.Vals[idx].Int())
		}
	case xr.Int32:
		ret = func(env *Env) int32 {
			return int32(env.Vals[idx].Int())
		}
	case xr.Int64:
		ret = func(env *Env) int64 {
			return env.Vals[idx].Int()
		}
	case xr.Uint:
		ret = func(env *Env) uint {
			return uint(env.Vals[idx].Uint())
		}
	case xr.Uint8:
		ret = func(env *Env) uint8 {
			return uint8(env.Vals[idx].Uint())
		}
	case xr.Uint16:
		ret = func(env *Env) uint16 {
			return uint16(env.Vals[idx].Uint())
		}
	case xr.Uint32:
		ret = func(env *Env) uint32 {
			return uint32(env.Vals[idx].Uint())
		}
	case xr.Uint64:
		ret = func(env *Env) uint64 {
			return env.Vals[idx].Uint()
		}
	case xr.Uintptr:
		ret = func(env *Env) uintptr {
			return uintptr(env.Vals[idx].Uint())
		}
	case xr.Float32:
		ret = func(env *Env) float32 {
			return float32(env.Vals[idx].Float())
		}
	case xr.Float64:
		ret = func(env *Env) float64 {
			return env.Vals[idx].Float()
		}
	case xr.Complex64:
		ret = func(env *Env) complex64 {
			return complex64(env.Vals[idx].Complex())
		}
	case xr.Complex128:
		ret = func(env *Env) complex128 {
			return env.Vals[idx].Complex()
		}
	case xr.String:
		ret = func(env *Env) string {
			return env.Vals[idx].String()
		}
	default:
		rtype := t.ReflectType()
		zero := xr.ZeroR(rtype)
		ret = func(env *Env) xr.Value {
			v := env.Vals[idx]
			if !v.IsValid() {
				v = zero
			} else if v.Type() != rtype {
				v = convert(v, rtype)
			}
			return v
		}
	}
	return exprFun(t, ret)
}

// unwrapBindUp1 compiles a conversion from a "mis-typed" bind stored in env.Outer.Binds[] as reflect.Value
// into a correctly-typed expression
func unwrapBindUp1(bind *Bind, t xr.Type) *Expr {
	idx := bind.Desc.Index()
	var ret I
	switch t.Kind() {
	case xr.Bool:
		ret = func(env *Env) bool {
			return env.Outer.Vals[idx].Bool()
		}
	case xr.Int:
		ret = func(env *Env) int {
			return int(env.Outer.Vals[idx].Int())
		}
	case xr.Int8:
		ret = func(env *Env) int8 {
			return int8(env.Outer.Vals[idx].Int())
		}
	case xr.Int16:
		ret = func(env *Env) int16 {
			return int16(env.Outer.Vals[idx].Int())
		}
	case xr.Int32:
		ret = func(env *Env) int32 {
			return int32(env.Outer.Vals[idx].Int())
		}
	case xr.Int64:
		ret = func(env *Env) int64 {
			return env.Outer.Vals[idx].Int()
		}
	case xr.Uint:
		ret = func(env *Env) uint {
			return uint(env.Outer.Vals[idx].Uint())
		}
	case xr.Uint8:
		ret = func(env *Env) uint8 {
			return uint8(env.Outer.Vals[idx].Uint())
		}
	case xr.Uint16:
		ret = func(env *Env) uint16 {
			return uint16(env.Outer.Vals[idx].Uint())
		}
	case xr.Uint32:
		ret = func(env *Env) uint32 {
			return uint32(env.Outer.Vals[idx].Uint())
		}
	case xr.Uint64:
		ret = func(env *Env) uint64 {
			return env.Outer.Vals[idx].Uint()
		}
	case xr.Uintptr:
		ret = func(env *Env) uintptr {
			return uintptr(env.Outer.Vals[idx].Uint())
		}
	case xr.Float32:
		ret = func(env *Env) float32 {
			return float32(env.Outer.Vals[idx].Float())
		}
	case xr.Float64:
		ret = func(env *Env) float64 {
			return env.Outer.Vals[idx].Float()
		}
	case xr.Complex64:
		ret = func(env *Env) complex64 {
			return complex64(env.Outer.Vals[idx].Complex())
		}
	case xr.Complex128:
		ret = func(env *Env) complex128 {
			return env.Outer.Vals[idx].Complex()
		}
	case xr.String:
		ret = func(env *Env) string {
			return env.Outer.Vals[idx].String()
		}
	default:
		rtype := t.ReflectType()
		zero := xr.ZeroR(rtype)
		ret = func(env *Env) xr.Value {
			v := env.Outer.Vals[idx]
			if !v.IsValid() {
				v = zero
			} else if v.Type() != rtype {
				v = convert(v, rtype)
			}
			return v
		}
	}
	return exprFun(t, ret)
}
