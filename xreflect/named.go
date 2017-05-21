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
	"unsafe"
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
	v := t.universe
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return t.method(i)
}

func (t *xtype) method(i int) Method {
	gfunc := t.gmethod(i)
	resizemethodvalues(t)

	var rfunctype reflect.Type
	rfuncs := &t.methodvalues
	rfunc := t.methodvalues[i]
	if rfunc.Kind() == reflect.Func {
		rfunctype = rfunc.Type()
	} else {
		rmethod, _ := t.rtype.MethodByName(gfunc.Name())
		rfunc = rmethod.Func
		t.methodvalues[i] = rfunc
		rfunctype = rmethod.Type
	}
	return t.universe.makemethod(i, gfunc, rfuncs, rfunctype) // lock already held
}

func (t *xtype) gmethod(i int) *types.Func {
	var gfun *types.Func
	if gtype, ok := t.gtype.Underlying().(*types.Interface); ok {
		gfun = gtype.ExplicitMethod(i)
	} else if gtype, ok := t.gtype.(*types.Named); ok {
		gfun = gtype.Method(i)
	} else {
		xerrorf(t, "Method on invalid type %v", t)
	}
	return gfun
}

func (v *Universe) makemethod(index int, gfun *types.Func, rfuns *[]reflect.Value, rfunctype reflect.Type) Method {
	return Method{
		Name:  gfun.Name(),
		Pkg:   (*Package)(gfun.Pkg()),
		Type:  v.maketype(gfun.Type(), rfunctype), // lock already held
		Funs:  rfuns,
		Index: index,
	}
}

func resizemethodvalues(t *xtype) {
	n := t.NumMethod()
	if cap(t.methodvalues) < n {
		slice := make([]reflect.Value, n, n+n/2+4)
		copy(slice, t.methodvalues)
		t.methodvalues = slice
	} else if len(t.methodvalues) < n {
		t.methodvalues = t.methodvalues[0:n]
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
			// redefined type. try really hard to support it.
			v.invalidateCache()
			// xerrorf(t, "SetUnderlying invoked multiple times on named type %v", t)
		}
		gunderlying := underlying.GoType().Underlying() // in case underlying is named
		t.kind = gtypeToKind(t, gunderlying)
		gtype.SetUnderlying(gunderlying)
		t.rtype = underlying.ReflectType()
	default:
		xerrorf(t, "SetUnderlying of unnamed type %v", t)
	}
}

// AddMethod adds method 'name' to type, unless it is already in the method list.
// It panics if the type is unnamed, or if the signature is not a function type,
// Returns the method index, or < 0 in case of errors
func (t *xtype) AddMethod(name string, signature Type) int {
	gtype, ok := t.gtype.(*types.Named)
	if !ok {
		xerrorf(t, "AddMethod on unnamed type %v", t)
	}
	kind := gtypeToKind(t, gtype.Underlying())
	if kind == reflect.Ptr || kind == reflect.Interface {
		xerrorf(t, "AddMethod: cannot add methods to named %s type: <%v>", kind, t)
	}
	if signature.Kind() != reflect.Func {
		xerrorf(t, "AddMethod on <%v> of non-function: %v", t, signature)
	}
	gsig := signature.underlying().(*types.Signature)
	// accept both signatures "non-nil receiver" and "nil receiver, use the first parameter as receiver"
	grecv := gsig.Recv()
	if grecv == nil && gsig.Params().Len() != 0 {
		grecv = gsig.Params().At(0)
	}
	if grecv == nil {
		xerrorf(t, "AddMethod on <%v> of function with no receiver and no parameters: %v", t, gsig)
	}
	if !types.IdenticalIgnoreTags(grecv.Type(), gtype) &&
		// !types.IdenticalIgnoreTags(grecv.Type(), gtype.Underlying()) &&
		!types.IdenticalIgnoreTags(grecv.Type(), types.NewPointer(gtype)) {

		label := "receiver"
		if gsig.Recv() == nil {
			label = "first parameter"
		}
		xerrorf(t, "AddMethod on <%v> of function <%v> with mismatched %s type: %v", t, gsig, label, grecv.Type())
	}

	gpkg := gtype.Obj().Pkg()
	gfun := types.NewFunc(token.NoPos, gpkg, name, gsig)

	index := unsafeAddMethod(gtype, gfun)
	n := gtype.NumMethods()

	for len(t.methodvalues) < n {
		t.methodvalues = append(t.methodvalues, reflect.Value{})
	}

	return index
}

// internal representation of go/types.Named
type unsafeNamed struct {
	obj        *types.TypeName
	underlying types.Type
	methods    []*types.Func
}

// patched version of go/types.Named.AddMethod() that *overwrites* matching methods
// (the original does not)
func unsafeAddMethod(gtype *types.Named, gfun *types.Func) int {
	if gfun.Name() == "_" {
		return -1
	}
	gt := (*unsafeNamed)(unsafe.Pointer(gtype))
	qname := QNameGo(gfun)
	for i, m := range gt.methods {
		if qname == QNameGo(m) {
			gt.methods[i] = gfun
			return i
		}
	}
	gt.methods = append(gt.methods, gfun)
	return len(gt.methods) - 1
}

// GetMethods returns the pointer to the method values.
// It panics if the type is unnamed
func (t *xtype) GetMethods() *[]reflect.Value {
	if !t.Named() {
		xerrorf(t, "GetMethods on unnamed type %v", t)
	}
	resizemethodvalues(t)
	return &t.methodvalues
}
