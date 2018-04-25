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
 * global.go
 *
 *  Created on Apr 25, 2018
 *      Author Massimiliano Ghilardi
 */

package untyped

import (
	"fmt"
	"go/constant"
	r "reflect"

	xr "github.com/cosmos72/gomacro/xreflect"
)

// UntypedLit represents an untyped literal value, i.e. an untyped constant
type Lit struct {
	Kind       r.Kind // default type. matches Val.Kind() except for rune literals, where Kind == reflect.Int32
	Val        constant.Value
	BasicTypes *[]xr.Type
}

// pretty-print untyped constants
func (untyp Lit) String() string {
	val := untyp.Val
	var strkind, strobj interface{} = untyp.Kind, nil
	if untyp.Kind == r.Int32 {
		strkind = "rune"
		if val.Kind() == constant.Int {
			if i, exact := constant.Int64Val(val); exact {
				if i >= 0 && i <= 0x10FFFF {
					strobj = fmt.Sprintf("%q", i)
				}
			}
		}
	}
	if strobj == nil {
		strobj = val.ExactString()
	}
	return fmt.Sprintf("{%v %v}", strkind, strobj)
}

func ConstantKindToUntypedLitKind(ckind constant.Kind) r.Kind {
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
