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
	var rfunc reflect.Value
	var rfunctype reflect.Type
	if len(t.methodvalues) > i && t.methodvalues[i].Kind() == reflect.Func {
		rfunc = t.methodvalues[i]
		rfunctype = rfunc.Type()
	} else {
		rmethod, _ := t.rtype.MethodByName(gfun.Name())
		rfunc = rmethod.Func
		rfunctype = rmethod.Type
	}
	return t.universe.makemethod(i, gfun, rfunc, rfunctype)
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

func (v *Universe) makemethod(index int, gfun *types.Func, rfunc reflect.Value, rfunctype reflect.Type) Method {
	return Method{
		Name:  gfun.Name(),
		Pkg:   (*Package)(gfun.Pkg()),
		Type:  v.MakeType(gfun.Type(), rfunctype),
		Func:  rfunc,
		Index: index,
	}
}

func (v *Universe) NamedOf(name, pkgpath string) Type {
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return v.namedOf(name, pkgpath)
}

func (v *Universe) namedOf(name, pkgpath string) Type {
	underlying := v.TypeOfInterface
	pkg := v.newPackage(pkgpath, "")
	typename := types.NewTypeName(token.NoPos, (*types.Package)(pkg), name, underlying.GoType())
	return v.maketype3(
		reflect.Invalid, // incomplete type! will be fixed by SetUnderlying
		types.NewNamed(typename, underlying.GoType(), nil),
		underlying.ReflectType(),
	)
}

// SetUnderlying sets the underlying type of a named type and marks t as complete.
// It panics if the type is unnamed, or if the underlying type is named,
// or if SetUnderlying() was already invoked on the named type.
func (t *xtype) SetUnderlying(underlying Type) {
	switch gtype := t.gtype.(type) {
	case *types.Named:
		v := t.universe
		if t.kind != reflect.Invalid || gtype.Underlying() != v.TypeOfInterface.GoType() || t.rtype != v.TypeOfInterface.ReflectType() {
			errorf("SetUnderlying invoked multiple times on named type %v", t)
		}
		gunderlying := underlying.GoType().Underlying() // in case underlying is named
		t.kind = gtypeToKind(gunderlying)
		gtype.SetUnderlying(gunderlying)
		t.rtype = underlying.ReflectType()
	default:
		errorf("SetUnderlying of unnamed type %v", t)
	}
}

// AddMethod adds method 'name' to type, unless it is already in the method list.
// It panics if the type is unnamed, or if the signature is not a function-with-receiver type.
// Returns the method index, or < 0 in case of errors
func (t *xtype) AddMethod(name string, signature Type) int {
	gtype, ok := t.gtype.(*types.Named)
	if !ok {
		errorf("AddMethod on unnamed type %v", t)
	}
	if signature.Kind() != reflect.Func {
		errorf("AddMethod on <%v> of non-func signature: %v", t, signature)
	}
	gsig := signature.underlying().(*types.Signature)
	// we accept both signatures "non-nil receiver" and "nil receiver, use the first parameter as receiver"
	grecv := gsig.Recv()
	if grecv == nil && gsig.Params().Len() != 0 {
		grecv = gsig.Params().At(0)
	}
	if grecv == nil {
		errorf("AddMethod on <%v> of function with no receiver and no parameters: %v", t, gsig)
	}
	if !types.IdenticalIgnoreTags(grecv.Type(), gtype) &&
		!types.IdenticalIgnoreTags(grecv.Type(), gtype.Underlying()) &&
		!types.IdenticalIgnoreTags(grecv.Type(), types.NewPointer(gtype)) {

		label := "receiver"
		if gsig.Recv() == nil {
			label = "first parameter"
		}
		debugf("AddMethod on <%v> of function <%v> with mismatched %s type: %v", t, gsig, label, grecv)
	}

	gpkg := gtype.Obj().Pkg()
	gfun := types.NewFunc(token.NoPos, gpkg, name, gsig)

	n1 := gtype.NumMethods()
	gtype.AddMethod(gfun)
	n2 := gtype.NumMethods()

	for len(t.methodvalues) < n2 {
		t.methodvalues = append(t.methodvalues, reflect.Value{})
	}

	if n2 == n1+1 {
		return n1
	}
	pkgpath := (*Package)(gpkg).Path()
	for i := 0; i < n2; i++ {
		gmethod := gtype.Method(i)
		if matchMethodByName(name, pkgpath, gmethod) {
			return i
		}
	}
	return -1 // method not found??
}

// GetMethodAddr returns the pointer to the method value for given method index.
// It panics if the type is unnamed, or if the index is outside [0,NumMethod()-1]
func (t *xtype) GetMethodAddr(index int) *reflect.Value {
	if !t.Named() {
		errorf("SetMethodValue on unnamed type %v", t)
	}
	return &t.methodvalues[index]
}
