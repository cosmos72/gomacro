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
 *  Created on Apr 11, 2017
 *      Author Massimiliano Ghilardi
 */

package base

import (
	r "reflect"
)

func KindToCategory(k r.Kind) r.Kind {
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

func IsCategory(k r.Kind, categories ...r.Kind) bool {
	k = KindToCategory(k)
	for _, c := range categories {
		if k == c {
			return true
		}
	}
	return false
}

// IsOptimizedKind returns true if fast interpreter expects optimized expressions for given Kind
func IsOptimizedKind(k r.Kind) bool {
	switch k {
	case r.Bool, r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
		r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr,
		r.Float32, r.Float64, r.Complex64, r.Complex128, r.String:
		return true
	}
	return false
}

// ConvertValue converts a value to type t and returns the converted value.
// extends reflect.Value.Convert(t) by allowing conversions from/to complex numbers.
// does not check for overflows or truncation.
func ConvertValue(v r.Value, to r.Type) r.Value {
	t := ValueType(v)
	if t == to {
		return v
	}
	if !t.ConvertibleTo(to) {
		// reflect.Value does not allow conversions from/to complex types
		k := v.Kind()
		kto := to.Kind()
		if IsCategory(kto, r.Complex128) {
			if IsCategory(k, r.Int, r.Uint, r.Float64) {
				temp := v.Convert(TypeOfFloat64).Float()
				v = r.ValueOf(complex(temp, 0.0))
			}
		} else if IsCategory(k, r.Complex128) {
			if IsCategory(k, r.Int, r.Uint, r.Float64) {
				temp := real(v.Complex())
				v = r.ValueOf(temp)
			}
		}
	}
	return v.Convert(to)
}
