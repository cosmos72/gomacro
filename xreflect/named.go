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
 * named.go
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

// NumMethod returns the number of explicitly declared methods of named type or interface t.
// Wrapper methods for embedded fields or embedded interfaces are not counted.
func (t *xtype) NumMethod() int {
	num := 0
	if gtype, ok := t.gtype.Underlying().(*types.Interface); ok {
		num = gtype.NumExplicitMethods()
	} else if gtype, ok := t.gtype.(*types.Named); ok {
		num = gtype.NumMethods()
	}
	return num
}

// Method return the i-th explicitly declared method of named type or interface t.
// Wrapper methods for embedded fields are not counted
func (t *xtype) Method(i int) Method {
	gfun := t.method(i)
	rmethod, _ := t.rtype.MethodByName(gfun.Name())
	return makemethod(i, gfun, &rmethod)
}

func (t *xtype) method(i int) *types.Func {
	var gfun *types.Func
	if gtype, ok := t.gtype.Underlying().(*types.Interface); ok {
		gfun = gtype.ExplicitMethod(i)
	} else if gtype, ok := t.gtype.(*types.Named); ok {
		gfun = gtype.Method(i)
	} else {
		errorf("Method on invalid type %v", t)
	}
	return gfun
}

func makemethod(index int, gfun *types.Func, rmethod *reflect.Method) Method {
	return Method{
		Name:  gfun.Name(),
		Pkg:   (*Package)(gfun.Pkg()),
		Type:  MakeType(gfun.Type(), rmethod.Type),
		Func:  rmethod.Func,
		Index: index,
	}
}

// NamedOf returns a new named type for the given type name.
// Initially, the underlying type is set to interface{} - use SetUnderlying to change it.
// These two steps are separate to allow creating self-referencing types,
// as for example type List struct { Elem int; Rest *List }
func NamedOf(name string, pkg *Package) Type {
	underlying := TypeOfInterface
	typename := types.NewTypeName(token.NoPos, (*types.Package)(pkg), name, underlying.GoType())
	return wrap(&xtype{
		kind:  reflect.Invalid, // incomplete type! will be fixed by SetUnderlying
		gtype: types.NewNamed(typename, underlying.GoType(), nil),
		rtype: underlying.ReflectType(),
	})
}

// SetUnderlying sets the underlying type of a named type and marks t as complete.
// It panics if the type is unnamed, or if the underlying type is named,
// or if SetUnderlying() was already invoked on the named type.
func (t *xtype) SetUnderlying(underlying Type) {
	switch gtype := t.gtype.(type) {
	case *types.Named:
		if t.kind != reflect.Invalid || gtype.Underlying() != TypeOfInterface.GoType() || t.rtype != TypeOfInterface.ReflectType() {
			errorf("SetUnderlying invoked multiple times on named type %v", t)
		}
		gunderlying := underlying.GoType()
		t.kind = gtypeToKind(gunderlying)
		gtype.SetUnderlying(gunderlying)
		t.rtype = underlying.ReflectType()
	default:
		errorf("SetUnderlying of unnamed type %v", t)
	}
}
