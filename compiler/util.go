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

package compiler

import (
	r "reflect"
)

// TypeOf() is a nil-safe version of reflect.TypeOf()
func TypeOf(i interface{}) r.Type {
	if i == nil {
		return nil
	}
	return r.TypeOf(i)
}

// ValueType() is a zero-value-safe version of reflect.Value.Type()
func ValueType(v r.Value) r.Type {
	if v == Nil || v == None {
		return nil
	}
	return v.Type()
}

func XNop(env *Env) (r.Value, []r.Value) {
	return None, nil
}

func XNil(env *Env) (r.Value, []r.Value) {
	return Nil, nil
}

func XTrue(env *Env) (r.Value, []r.Value) {
	return True, nil
}

func XFalse(env *Env) (r.Value, []r.Value) {
	return False, nil
}

func ToX(expr I) X {
	switch x := expr.(type) {
	case nil:
		return XNil
	case bool:
		if x {
			return XTrue
		} else {
			return XFalse
		}
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64, complex64, complex128, string:
		v := r.ValueOf(expr)
		return func(env *Env) (r.Value, []r.Value) {
			return v, nil
		}
	case X:
		return x
	case func(*Env) (r.Value, []r.Value):
		return x
	case func(*Env) bool:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) int:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) int8:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) int16:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) int32:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) int64:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) uint:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) uint8:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) uint16:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) uint32:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) uint64:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) uintptr:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) float32:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) float64:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) complex64:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) complex128:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	case func(*Env) string:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(x(env)), nil
		}
	default:
		errorf("unsupported expression, cannot convert to func(*Env) (r.Value[], []r.Value): %v <%T>",
			expr, expr)
		return nil
	}
}

func PackValues(val0 r.Value, vals []r.Value) []r.Value {
	if len(vals) == 0 && val0 != None {
		vals = []r.Value{val0}
	}
	return vals
}

func UnpackValues(vals []r.Value) (r.Value, []r.Value) {
	val0 := None
	if len(vals) > 0 {
		val0 = vals[0]
	}
	return val0, vals
}
