/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
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
 * struct.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/token"
	"go/types"
	"reflect"
)

func toGoFuncs(names []string, methods []Type) []*types.Func {
	gfuns := make([]*types.Func, len(methods))
	for i, t := range methods {
		switch gsig := t.gunderlying().(type) {
		case *types.Signature:
			gfuns[i] = types.NewFunc(token.NoPos, nil, names[i], gsig)
		default:
			errorf(t, "interface contains non-function: %s %v", names[i], t)
		}
	}
	return gfuns
}

func toGoNamedTypes(ts []Type) []*types.Named {
	gnameds := make([]*types.Named, len(ts))
	for i, t := range ts {
		if gt, ok := t.GoType().(*types.Named); ok {
			if t.Kind() == reflect.Interface {
				gnameds[i] = gt
			} else {
				errorf(t, "interface contains embedded non-interface: %v", t)
			}
		} else {
			errorf(t, "interface contains embedded interface without name: %v", t)
		}
	}
	return gnameds
}

// InterfaceOf returns a new interface for the given methods and embedded types.
// After the methods and embeddeds are fully defined, call Complete() to mark
// the interface as complete and compute wrapper methods for embedded fields.
//
// WARNING: the Type returned by InterfaceOf is not complete,
// i.e. its method set is not computed yet.
// Once you know that methods and embedded interfaces are complete,
// call Complete() to compute the method set and mark this Type as complete.
func (v *Universe) InterfaceOf(methodnames []string, methodtypes []Type, embeddeds []Type) Type {
	gmethods := toGoFuncs(methodnames, methodtypes)
	gembeddeds := toGoNamedTypes(embeddeds)

	gtype := types.NewInterface(gmethods, gembeddeds)
	gtype.Complete()

	// for reflect.Type, approximate an interface as a pointer-to-struct:
	// one field for the wrapped object: type is interface{},
	// one field for each explicit method: type is the method type i.e. a function
	rfields := make([]reflect.StructField, 2+len(methodtypes), gtype.NumMethods()+2)
	rfields[0] = approxInterfaceHeader()
	rfields[1] = approxInterfaceEmbeddeds(embeddeds)

	for i, methodtype := range methodtypes {
		rfields[i+2] = approxInterfaceMethod(methodnames[i], methodtype.ReflectType())
	}
	for _, e := range embeddeds {
		n := e.NumMethod()
		for i := 0; i < n; i++ {
			method := e.Method(i)
			rfields = append(rfields, approxInterfaceMethod(method.Name, method.Type.ReflectType()))
		}
	}
	// interfaces may have lots of methods, thus a lot of fields in the proxy struct.
	// Use a pointer to the proxy struct
	rtype := reflect.PtrTo(reflect.StructOf(rfields))
	t := v.maketype3(reflect.Interface, gtype, rtype)
	// debugf("InterfaceOf: new type %v", t)
	// debugf("           types.Type %v", gtype)
	// debugf("         reflect.Type %v", rtype)
	return t
}

// Complete marks an interface type as complete and computes wrapper methods for embedded fields.
// It must be called by users of InterfaceOf after the interface's embedded types are fully defined
// and before using the interface type in any way other than to form other types.
func (t *xtype) Complete() Type {
	if t.kind != reflect.Interface {
		xerrorf(t, "Complete of non-interface %v", t)
	}
	return wrap(t)
}

// utilities for InterfaceOf()

func approxInterfaceHeader() reflect.StructField {
	return reflect.StructField{
		Name: StrGensymInterface,
		Type: reflectTypeOfInterfaceHeader,
	}
}

func approxInterfaceEmbeddeds(embeddeds []Type) reflect.StructField {
	fields := make([]reflect.StructField, len(embeddeds))
	for i, t := range embeddeds {
		fields[i] = approxInterfaceEmbedded(t)
	}
	return reflect.StructField{
		Name: StrGensymEmbedded,
		Type: reflect.ArrayOf(0, reflect.StructOf(fields)),
	}
}

func approxInterfaceEmbedded(t Type) reflect.StructField {
	return reflect.StructField{
		Name: toExportedFieldName("", t, true),
		Type: t.ReflectType(),
	}
}

func approxInterfaceMethod(name string, rtype reflect.Type) reflect.StructField {
	// interface methods cannot be anonymous
	if len(name) == 0 {
		name = "_"
	}
	return reflect.StructField{
		Name: toExportedFieldName(name, nil, false),
		Type: rtype,
	}
}
