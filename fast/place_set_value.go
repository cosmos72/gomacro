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
 * place_set_value.go
 *
 *  Created on May 29, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"github.com/cosmos72/gomacro/base/reflect"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// placeSetValue compiles 'place = value' where value is a reflect.Value passed at runtime.
// Used to assign places with the result of multi-valued expressions,
// and to implement multiple assignment place1, place2... = expr1, expr2...
func (c *Comp) placeSetValue(place *Place) func(lhs, key, val xr.Value) {
	rtype := place.Type.ReflectType()

	if place.MapKey != nil {
		zero := xr.ZeroR(rtype)
		return func(lhs, key, val xr.Value) {
			if !val.IsValid() || val == None {
				val = zero
			} else if val.Type() != rtype {
				val = val.Convert(rtype)
			}
			lhs.SetMapIndex(key, val)
		}
	}
	var ret func(xr.Value, xr.Value, xr.Value)
	switch reflect.Category(rtype.Kind()) {
	case xr.Bool:
		ret = func(lhs, key, val xr.Value) {
			lhs.SetBool(val.Bool())
		}
	case xr.Int:
		ret = func(lhs, key, val xr.Value) {
			lhs.SetInt(val.Int())
		}
	case xr.Uint:
		ret = func(lhs, key, val xr.Value) {
			lhs.SetUint(val.Uint())
		}
	case xr.Float64:
		ret = func(lhs, key, val xr.Value) {
			lhs.SetFloat(val.Float())
		}
	case xr.Complex128:
		ret = func(lhs, key, val xr.Value) {
			lhs.SetComplex(val.Complex())
		}
	case xr.String:
		ret = func(lhs, key, val xr.Value) {
			lhs.SetString(val.String())
		}
	default:
		zero := xr.ZeroR(rtype)
		ret = func(lhs, key, val xr.Value) {
			if !val.IsValid() || val == None {
				val = zero
			} else if val.Type() != rtype {
				val = val.Convert(rtype)
			}
			lhs.Set(val)
		}
	}
	return ret
}
