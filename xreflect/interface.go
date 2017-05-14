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

package xtype

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
		switch gt := t.GoType().(type) {
		case *types.Named:
			gnameds[i] = gt
		default:
			errorf("InterfaceOf: %d-th 'embedded' argument is not a named type: %v", i, t)
		}
	}
	return gnameds
}

// InterfaceOf returns a new interface for the given methods and embedded types.
// After the methods and embeddeds are fully defined, call Complete() to mark the interface as complete
// and compute wrapper methods for embedded fields.
func InterfaceOf(methodnames []string, methods []Type, embeddeds []Type) Type {
	gmethods := toGoFuncs(methodnames, methods)
	gembeddeds := toGoNamedTypes(embeddeds)

	// for reflect.Type, approximate an interface as a struct:
	// one field for the wrapped object: type is interface{},
	// one field for each embedded interface: type is the embedded interface type
	// one field for each method: type is the method type i.e. a function
	nemb := len(embeddeds)
	rfields := make([]reflect.StructField, 1+nemb+len(methods))
	rfields[0] = approxInterfaceHeader()
	for i, emb := range embeddeds {
		rfields[i+1] = approxInterfaceEmbedded(emb.Name(), emb.ReflectType())
	}
	for i, method := range methods {
		rfields[i+nemb+1] = approxInterfaceMethod(methodnames[i], method.ReflectType())
	}
	return maketype(
		types.NewInterface(gmethods, gembeddeds),
		// interfaces may have lots of methods, thus a lot of fields in the proxy struct.
		// Then use a pointer to the proxy struct
		reflect.PtrTo(reflect.StructOf(rfields)),
	)
}

// Complete marks an interface type as complete and computes wrapper methods for embedded fields.
// It must be called by users of InterfaceOf after the interface's embedded types are fully defined
// and before using the interface type in any way other than to form other types.
// Complete returns the receiver.
func (t *xtype) Complete() Type {
	if t.kind != reflect.Interface {
		errorf("Complete of non-interface %v", t)
	}
	gtype := t.gtype.Underlying().(*types.Interface)
	gtype.Complete()
	return t
}

// utilities for InterfaceOf()

func approxInterfaceHeader() reflect.StructField {
	return reflect.StructField{
		Name: StrGensymInterface,
		Type: reflectTypeOfInterfaceHeader,
	}
}

func approxInterfaceEmbedded(typename string, rtype reflect.Type) reflect.StructField {
	// embedded interfaces are always anonymous
	return reflect.StructField{
		Name: toExportedFieldName("", typename, true),
		Type: rtype,
	}
}

func approxInterfaceMethod(name string, rtype reflect.Type) reflect.StructField {
	// interface methods cannot be anonymous
	if len(name) == 0 {
		name = "_"
	}
	return reflect.StructField{
		Name: toExportedFieldName(name, "", false),
		Type: rtype,
	}
}
