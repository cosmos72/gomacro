/*
 * gomacro - A Go interpreter with Lisp-like macros
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
 * struct.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package type2

import (
	"go/token"
	"go/types"
	"reflect"
)

func toGoFuncs(names []string, methods []Type) []*types.Func {
	gfuns := make([]*types.Func, len(methods))
	for i, t := range methods {
		switch gsig := t.underlying().(type) {
		case *types.Signature:
			gfuns[i] = types.NewFunc(token.NoPos, nil, names[i], gsig)
		default:
			errorf("InterfaceOf: %d-th 'method' argument is not a function type: %v", i, t)
		}
	}
	return gfuns
}

func toGoNamedTypes(ts []Type) []*types.Named {
	gnameds := make([]*types.Named, len(ts))
	for i, t := range ts {
		switch gt := t.gtype.(type) {
		case *types.Named:
			gnameds[i] = gt
		default:
			errorf("InterfaceOf: %d-th 'embedded' argument is not a named type: %v", i, t)
		}
	}
	return gnameds
}

func InterfaceOf(methodnames []string, methods []Type, embeddeds []Type) Type {
	gmethods := toGoFuncs(methodnames, methods)
	gembeddeds := toGoNamedTypes(embeddeds)
	fields := toStructFields(methodnames, methods)
	rfields := toReflectFields(fields, true)
	return maketype(
		types.NewInterface(gmethods, gembeddeds),
		reflect.StructOf(rfields),
	)
}
