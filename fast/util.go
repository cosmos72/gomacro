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
 * identifier.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/constant"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func eFalse(*Env) bool {
	return false
}

func eTrue(*Env) bool {
	return true
}

func eNone(*Env) {
}

func eNil(*Env) r.Value {
	return Nil
}

func nop() {
}

var valueOfNopFunc = r.ValueOf(nop)

func (e *Expr) TryAsPred() (value bool, fun func(*Env) bool, err bool) {
	if e.Untyped() {
		untyp := e.Value.(UntypedLit)
		if untyp.Kind != r.Bool {
			return false, nil, true
		}
		return constant.BoolVal(untyp.Obj), nil, false
	}
	if e.Type.Kind() != r.Bool {
		return false, nil, true
	}
	if e.Const() {
		v := r.ValueOf(e.Value)
		return v.Bool(), nil, false
	}
	switch fun := e.Fun.(type) {
	case func(*Env) bool:
		return false, fun, false
	case func(*Env) (r.Value, []r.Value):
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
	var t r.Type = nil
	if x.NumOut() != 0 {
		t = x.Out(0)
	}
	c.Errorf("%s boolean predicate, expecting <bool> expression, found <%v>: %v",
		reason, t, node)
	return nil
}

func (e *Expr) AsX() X {
	if e == nil || e.Const() {
		return nil
	}
	return funAsX(e.Fun)
}

func funAsX(any I) X {
	switch fun := any.(type) {
	case nil:
	case X:
		return fun
	case func(*Env):
		return fun
	case func(*Env) r.Value:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) (r.Value, []r.Value):
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
		Errorf("unsupported function type, cannot convert to func(*Env): %v <%v>", any, r.TypeOf(any))
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
		Errorf("expression returns no values, cannot convert to func(env *Env) r.Value")
		return
	} else if e.NumOut() > 1 {
		Warnf("expression returns %d values, using only the first one: %v", e.Types)
	}
}

func (e *Expr) AsX1() func(*Env) r.Value {
	if e.Const() {
		return valueAsX1(e.Value, e.Type, CompileDefaults)
	}
	e.CheckX1()
	return funAsX1(e.Fun, e.Type)
}

func (e *Expr) AsXV(opts CompileOptions) func(*Env) (r.Value, []r.Value) {
	if e.Const() {
		return valueAsXV(e.Value, e.Type, opts)
	}
	return funAsXV(e.Fun, e.Type)
}

func valueAsX1(any I, t r.Type, opts CompileOptions) func(*Env) r.Value {
	convertuntyped := opts&CompileKeepUntyped == 0
	if convertuntyped {
		if untyp, ok := any.(UntypedLit); ok {
			// Debugf("late conversion of untyped constant %v <%v> to <%v>", untyp, r.TypeOf(untyp), t)
			if t == typeOfUntypedLit {
				t = untyp.DefaultType()
			}
			any = untyp.ConstTo(t)
		}
	}
	v := r.ValueOf(any)
	if t != nil {
		if ValueType(v) == nil {
			v = r.Zero(t)
		} else if convertuntyped {
			v = v.Convert(t)
		}
	}
	return func(*Env) r.Value {
		return v
	}
}

func valueAsXV(any I, t r.Type, opts CompileOptions) func(*Env) (r.Value, []r.Value) {
	convertuntyped := opts&CompileKeepUntyped == 0
	if convertuntyped {
		if untyp, ok := any.(UntypedLit); ok {
			// Debugf("valueAsXV: late conversion of untyped constant %v <%v> to <%v>", untyp, r.TypeOf(untyp), t)
			if t == typeOfUntypedLit {
				t = untyp.DefaultType()
			}
			any = untyp.ConstTo(t)
		}
	}
	v := r.ValueOf(any)
	if t != nil {
		if ValueType(v) == nil {
			v = r.Zero(t)
		} else if convertuntyped {
			v = v.Convert(t)
		}
	}
	return func(*Env) (r.Value, []r.Value) {
		return v, nil
	}
}

func funAsX1(fun I, t r.Type) func(*Env) r.Value {
	// Debugf("funAsX1() %v -> %v", r.TypeOf(fun), t)
	switch fun := fun.(type) {
	case nil:
	case X:
		if fun == nil {
			break
		}
		return func(env *Env) r.Value {
			fun(env)
			return None
		}
	case func(*Env):
		if fun == nil {
			break
		}
		return func(env *Env) r.Value {
			fun(env)
			return None
		}
	case func(*Env) r.Value:
		return fun
	case func(*Env) (r.Value, []r.Value):
		return func(env *Env) r.Value {
			ret, _ := fun(env)
			return ret
		}
	case func(*Env) bool:
		if t == nil || t == TypeOfBool {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) int:
		if t == nil || t == TypeOfInt {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) int8:
		if t == nil || t == TypeOfInt8 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) int16:
		if t == nil || t == TypeOfInt16 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) int32:
		if t == nil || t == TypeOfInt32 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) int64:
		if t == nil || t == TypeOfInt64 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) uint:
		if t == nil || t == TypeOfUint {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) uint8:
		if t == nil || t == TypeOfUint8 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) uint16:
		if t == nil || t == TypeOfUint16 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) uint32:
		if t == nil || t == TypeOfUint32 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) uint64:
		if t == nil || t == TypeOfUint64 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) uintptr:
		if t == nil || t == TypeOfUintptr {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) float32:
		if t == nil || t == TypeOfFloat32 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) float64:
		if t == nil || t == TypeOfFloat64 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) complex64:
		if t == nil || t == TypeOfComplex64 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) complex128:
		if t == nil || t == TypeOfComplex128 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) string:
		if t == nil || t == TypeOfString {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *bool:
		if t == nil || t == TypeOfPtrBool {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *int:
		if t == nil || t == TypeOfPtrInt {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *int8:
		if t == nil || t == TypeOfPtrInt8 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *int16:
		if t == nil || t == TypeOfPtrInt16 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *int32:
		if t == nil || t == TypeOfPtrInt32 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *int64:
		if t == nil || t == TypeOfPtrInt64 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *uint:
		if t == nil || t == TypeOfPtrUint {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *uint8:
		if t == nil || t == TypeOfPtrUint8 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *uint16:
		if t == nil || t == TypeOfPtrUint16 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *uint32:
		if t == nil || t == TypeOfPtrUint32 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *uint64:
		if t == nil || t == TypeOfPtrUint64 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *uintptr:
		if t == nil || t == TypeOfPtrUintptr {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *float32:
		if t == nil || t == TypeOfPtrFloat32 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *float64:
		if t == nil || t == TypeOfPtrFloat64 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	case func(*Env) *complex64:
		if t == nil || t == TypeOfPtrComplex64 {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env))
			}
		} else {
			return func(env *Env) r.Value {
				return r.ValueOf(fun(env)).Convert(t)
			}
		}
	default:
		Errorf("unsupported expression type, cannot convert to func(*Env) r.Value: %v <%v>", fun, r.TypeOf(fun))
	}
	return nil
}

func funAsXV(fun I, t r.Type) func(*Env) (r.Value, []r.Value) {
	// Debugf("funAsXV() %v -> %v", r.TypeOf(fun), t)
	switch fun := fun.(type) {
	case nil:
	case X:
		if fun == nil {
			break
		}
		return func(env *Env) (r.Value, []r.Value) {
			fun(env)
			return None, nil
		}
	case func(*Env):
		if fun == nil {
			break
		}
		return func(env *Env) (r.Value, []r.Value) {
			fun(env)
			return None, nil
		}
	case func(*Env) r.Value:
		return func(env *Env) (r.Value, []r.Value) {
			return fun(env), nil
		}
	case func(*Env) (r.Value, []r.Value):
		return fun
	case func(*Env) bool:
		if t == nil || t == TypeOfBool {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) int:
		if t == nil || t == TypeOfInt {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) int8:
		if t == nil || t == TypeOfInt8 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) int16:
		if t == nil || t == TypeOfInt16 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) int32:
		if t == nil || t == TypeOfInt32 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) int64:
		if t == nil || t == TypeOfInt64 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) uint:
		if t == nil || t == TypeOfUint {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) uint8:
		if t == nil || t == TypeOfUint8 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) uint16:
		if t == nil || t == TypeOfUint16 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) uint32:
		if t == nil || t == TypeOfUint32 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) uint64:
		if t == nil || t == TypeOfUint64 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) uintptr:
		if t == nil || t == TypeOfUintptr {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) float32:
		if t == nil || t == TypeOfFloat32 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) float64:
		if t == nil || t == TypeOfFloat64 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) complex64:
		if t == nil || t == TypeOfComplex64 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) complex128:
		if t == nil || t == TypeOfComplex128 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) string:
		if t == nil || t == TypeOfString {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *bool:
		if t == nil || t == TypeOfPtrBool {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *int:
		if t == nil || t == TypeOfPtrInt {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *int8:
		if t == nil || t == TypeOfPtrInt8 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *int16:
		if t == nil || t == TypeOfPtrInt16 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *int32:
		if t == nil || t == TypeOfPtrInt32 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *int64:
		if t == nil || t == TypeOfPtrInt64 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *uint:
		if t == nil || t == TypeOfPtrUint {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *uint8:
		if t == nil || t == TypeOfPtrUint8 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *uint16:
		if t == nil || t == TypeOfPtrUint16 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *uint32:
		if t == nil || t == TypeOfPtrUint32 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *uint64:
		if t == nil || t == TypeOfPtrUint64 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *uintptr:
		if t == nil || t == TypeOfPtrUintptr {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *float32:
		if t == nil || t == TypeOfPtrFloat32 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *float64:
		if t == nil || t == TypeOfPtrFloat64 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	case func(*Env) *complex64:
		if t == nil || t == TypeOfPtrComplex64 {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)), nil
			}
		} else {
			return func(env *Env) (r.Value, []r.Value) {
				return r.ValueOf(fun(env)).Convert(t), nil
			}
		}
	default:
		Errorf("unsupported expression, cannot convert to func(*Env) (r.Value, []r.Value) : %v <%v>",
			fun, r.TypeOf(fun))
	}
	return nil
}

func (e *Expr) AsStmt() Stmt {
	if e == nil || e.Const() {
		return nil
	}
	return funAsStmt(e.Fun)
}

func funAsStmt(fun I) Stmt {
	var ret func(env *Env) (Stmt, *Env)

	switch fun := fun.(type) {
	case nil:
	case X:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env):
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) r.Value:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) (r.Value, []r.Value):
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

	case func(*Env) *bool:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *int:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *int8:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *int16:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *int32:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *int64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *uint:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *uint8:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *uint16:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *uint32:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *uint64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *uintptr:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *float32:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *float64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) *complex64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	default:
		Errorf("unsupported expression type, cannot convert to Stmt : %v <%v>",
			fun, r.TypeOf(fun))
	}
	return ret
}

// funTypeOuts returns the return types of given function
func funTypeOuts(fun I) []r.Type {
	t := r.TypeOf(fun)
	if t == nil || t.Kind() != r.Func {
		return []r.Type{t}
	}
	n := t.NumOut()
	ts := make([]r.Type, n)
	for i := 0; i < n; i++ {
		ts[i] = t.Out(i)
	}
	return ts
}
