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
 * fromreflect.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package type2

import (
	"go/token"
	"go/types"
	"reflect"
	"strings"
)

const (
	KeepReflectType     = 0
	RebuildReflectType1 = 1
)

// FromReflectType creates a Type corresponding to given reflect.Type
// Note: conversions from Type to reflect.Type and back are not exact,
// because of the reasons listed in Type.ReflectType()
// Conversions from reflect.Type to Type and back are not exact for the same reasons.
func FromReflectType(rtype reflect.Type, rebuildDepth int) Type {
	var t Type
	rebuildDepth-- // decrement ONLY here and in fromReflectPtr() when calling fromReflectInterfaceStruct()
	switch k := rtype.Kind(); k {
	case reflect.Invalid:
		return Type{}
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		t = BasicTypes[k]
	case reflect.Array:
		t = fromReflectArray(rtype, rebuildDepth)
	case reflect.Chan:
		t = fromReflectChan(rtype, rebuildDepth)
	case reflect.Func:
		t = fromReflectFunc(rtype, rebuildDepth)
	case reflect.Interface:
		t = fromReflectInterface(rtype, rebuildDepth)
	case reflect.Map:
		t = fromReflectMap(rtype, rebuildDepth)
	case reflect.Ptr:
		t = fromReflectPtr(rtype, rebuildDepth)
	case reflect.Slice:
		t = fromReflectSlice(rtype, rebuildDepth)
	case reflect.Struct:
		t = fromReflectStruct(rtype, rebuildDepth)
	// case reflect.UnsafePointer:
	default:
		errorf("unsupported reflect.Type %v", rtype)
	}
	return finish(t, rtype, rebuildDepth)
}

func finish(t Type, rtype reflect.Type, rebuildDepth int) Type {
	name := rtype.Name()
	pkgpath := rtype.PkgPath()
	if pkgpath == t.PkgPath() && name == t.Name() {
		return addmethods(t, rtype, rebuildDepth)
	}
	underlying := maketype(t.underlying(), t.rtype)
	if len(name) == 0 {
		return underlying
	}
	pkg := NewPackage(pkgpath, "")
	t = NamedOf(name, pkg)
	t.SetUnderlying(underlying)
	return addmethods(t, rtype, rebuildDepth)
}

func addmethods(t Type, rtype reflect.Type, rebuildDepth int) Type {
	n := rtype.NumMethod()
	if n == 0 {
		return t
	}
	if !t.Named() {
		errorf("cannot add methods to unnamed type %v", t)
	}
	for i := 0; i < n; i++ {
		rmethod := rtype.Method(i)
		signature := fromReflectFunc(rmethod.Type, rebuildDepth)
		t.AddMethod(rmethod.Name, signature)
	}
	return t
}

func fromReflectField(rfield *reflect.StructField, rebuildDepth int) StructField {
	t := FromReflectType(rfield.Type, rebuildDepth)
	name := rfield.Name
	anonymous := rfield.Anonymous

	if strings.HasPrefix(name, StrGensymEmbedded) {
		// this reflect.StructField emulates embedded field using our own convention.
		// eat our own dogfood and convert it back to an embedded field.
		rtype := rfield.Type
		typename := rtype.Name()
		if len(typename) == 0 {
			typename = name[len(StrGensymEmbedded):]
		}
		// rebuild the type's name and package
		t = named(t, typename, rtype.PkgPath())
		name = ""
		anonymous = true
	} else if strings.HasPrefix(name, StrGensymPrivate) {
		// this reflect.StructField emulates private (unexported) field using our own convention.
		// eat our own dogfood and convert it back to a private field.
		name = name[len(StrGensymPrivate):]
	}

	return StructField{
		Name:      name,
		Pkg:       NewPackage(rfield.PkgPath, ""),
		Type:      t,
		Tag:       rfield.Tag,
		Offset:    rfield.Offset,
		Index:     rfield.Index,
		Anonymous: anonymous,
	}
}

func fromReflectFields(rfields []reflect.StructField, rebuildDepth int) []StructField {
	fields := make([]StructField, len(rfields))
	for i := range rfields {
		fields[i] = fromReflectField(&rfields[i], rebuildDepth)
	}
	return fields
}

// named creates a new named Type based on t, having the given name and pkgpath
func named(t Type, name string, pkgpath string) Type {
	if t.Name() != name || t.PkgPath() != pkgpath {
		t2 := NamedOf(name, NewPackage(pkgpath, ""))
		t2.SetUnderlying(maketype(t.underlying(), t.rtype))
		t = t2
	}
	return t
}

// fromReflectArray converts a reflect.Type with Kind reflect.Array into a Type
func fromReflectArray(rtype reflect.Type, rebuildDepth int) Type {
	count := rtype.Len()
	elem := FromReflectType(rtype.Elem(), rebuildDepth)
	if rebuildDepth >= 0 {
		rtype = reflect.ArrayOf(count, elem.rtype)
	}
	return maketype(types.NewArray(elem.gtype, int64(count)), rtype)
}

// fromReflectChan converts a reflect.Type with Kind reflect.Chan into a Type
func fromReflectChan(rtype reflect.Type, rebuildDepth int) Type {
	dir := rtype.ChanDir()
	elem := FromReflectType(rtype.Elem(), rebuildDepth)
	if rebuildDepth >= 0 {
		rtype = reflect.ChanOf(dir, elem.rtype)
	}
	gdir := dirToGdir(dir)
	return maketype(types.NewChan(gdir, elem.gtype), rtype)
}

// fromReflectFunc converts a reflect.Type with Kind reflect.Func into a Type
func fromReflectFunc(rtype reflect.Type, rebuildDepth int) Type {
	nin, nout := rtype.NumIn(), rtype.NumOut()
	in := make([]Type, nin)
	out := make([]Type, nout)
	for i := 0; i < nin; i++ {
		in[i] = FromReflectType(rtype.In(i), rebuildDepth)
	}
	for i := 0; i < nout; i++ {
		out[i] = FromReflectType(rtype.Out(i), rebuildDepth)
	}
	gin := toGoTuple(in)
	gout := toGoTuple(out)
	variadic := rtype.IsVariadic()

	if rebuildDepth >= 0 {
		rin := toReflectTypes(in)
		rout := toReflectTypes(out)
		rtype = reflect.FuncOf(rin, rout, variadic)
	}
	return maketype(
		types.NewSignature(nil, gin, gout, variadic),
		rtype,
	)
}

// fromReflectInterface converts a reflect.Type with Kind reflect.Interface into a Type
func fromReflectInterface(rtype reflect.Type, rebuildDepth int) Type {
	if rtype == TypeOfInterface.rtype {
		return TypeOfInterface
	}
	n := rtype.NumMethod()
	gmethods := make([]*types.Func, n)
	for i := 0; i < n; i++ {
		rmethod := rtype.Method(i)
		method := fromReflectFunc(rmethod.Type, rebuildDepth)
		gmethods[i] = types.NewFunc(token.NoPos, NewPackage(rmethod.PkgPath, "").impl, rmethod.Name, method.gtype.(*types.Signature))
	}
	// no way to extract embedded interfaces from reflect.Type
	if rebuildDepth >= 0 {
		rfields := make([]reflect.StructField, 1+n)
		rfields[0] = approxInterfaceHeader()
		for i := 0; i < n; i++ {
			rmethod := rtype.Method(i)
			rmethodtype := rmethod.Type
			if rebuildDepth >= 1 {
				rmethodtype = FromReflectType(rmethod.Type, rebuildDepth).rtype
			}
			rfields[i+1] = approxInterfaceMethod(rmethod.Name, rmethodtype)
		}
		// interfaces may have lots of methods, thus a lot of fields in the proxy struct.
		// Then use a pointer to the proxy struct: InterfaceOf() does that, and we must behave identically
		rtype = reflect.PtrTo(reflect.StructOf(rfields))
	}
	return maketype(types.NewInterface(gmethods, nil), rtype)
}

// isReflectInterfaceStruct returns true if rtype is a reflect.Type with Kind reflect.Struct,
// that contains our own conventions to emulate an interface
func isReflectInterfaceStruct(rtype reflect.Type) bool {
	if rtype.Kind() == reflect.Struct {
		if n := rtype.NumField(); n != 0 {
			rfield := rtype.Field(0)
			return rfield.Name == StrGensymInterface && rfield.Type == reflectTypeOfInterfaceHeader
		}
	}
	return false
}

// fromReflectInterfaceStruct converts a reflect.Type with Kind reflect.Struct,
// that contains our own conventions to emulate an interface, into a Type
func fromReflectInterfaceStruct(rtype reflect.Type, rebuildDepth int) Type {
	rebuild := rebuildDepth >= 0
	n := rtype.NumField()
	// skip rtype.Field(0), it is just approxInterfaceSelf()
	var gmethods []*types.Func
	var gembeddeds []*types.Named
	var rebuildfields []reflect.StructField
	if rebuild {
		rebuildfields = make([]reflect.StructField, n)
		rebuildfields[0] = approxInterfaceHeader()
	}
	for i := 1; i < n; i++ {
		rfield := rtype.Field(i)
		name := rfield.Name
		if strings.HasPrefix(name, StrGensymEmbedded) {
			typename := name[len(StrGensymEmbedded):]
			t := FromReflectType(rfield.Type, rebuildDepth)
			if t.Kind() != reflect.Interface {
				errorf("FromReflectType: reflect.Type <%v> is an emulated interface containing the embedded interface <%v>.\n\tExtracting the latter returned a non-interface: %v", t)
			}
			t = named(t, typename, rfield.Type.PkgPath())
			gembeddeds = append(gembeddeds, t.gtype.(*types.Named))
			if rebuild {
				rebuildfields[i] = approxInterfaceEmbedded(typename, t.rtype)
			}
		} else {
			if strings.HasPrefix(name, StrGensymPrivate) {
				name = name[len(StrGensymPrivate):]
			}
			t := fromReflectFunc(rfield.Type, rebuildDepth)
			if t.Kind() != reflect.Func {
				errorf("FromReflectType: reflect.Type <%v> is an emulated interface containing the method <%v>.\n\tExtracting the latter returned a non-function: %v", t)
			}
			gmethods = append(gmethods, types.NewFunc(token.NoPos, NewPackage(rfield.PkgPath, "").impl, name, t.gtype.(*types.Signature)))
			if rebuild {
				rebuildfields[i] = approxInterfaceMethod(name, t.rtype)
			}
		}
	}
	if rebuild {
		rtype = reflect.StructOf(rebuildfields)
	}
	return maketype(types.NewInterface(gmethods, gembeddeds), rtype)
}

// fromReflectMap converts a reflect.Type with Kind reflect.map into a Type
func fromReflectMap(rtype reflect.Type, rebuildDepth int) Type {
	key := FromReflectType(rtype.Key(), rebuildDepth)
	elem := FromReflectType(rtype.Elem(), rebuildDepth)
	if rebuildDepth >= 0 {
		rtype = reflect.MapOf(key.rtype, elem.rtype)
	}
	return maketype(types.NewMap(key.gtype, elem.gtype), rtype)
}

// fromReflectPtr converts a reflect.Type with Kind reflect.Ptr into a Type
func fromReflectPtr(rtype reflect.Type, rebuildDepth int) Type {
	relem := rtype.Elem()
	var elem Type
	var gtype types.Type
	if isReflectInterfaceStruct(relem) {
		elem = fromReflectInterfaceStruct(relem, rebuildDepth-1)
		gtype = elem.gtype
	} else {
		elem = FromReflectType(relem, rebuildDepth)
		gtype = types.NewPointer(elem.gtype)
	}
	if rebuildDepth >= 0 {
		rtype = reflect.PtrTo(elem.rtype)
	}
	return maketype(gtype, rtype)
}

// fromReflectPtr converts a reflect.Type with Kind reflect.Slice into a Type
func fromReflectSlice(rtype reflect.Type, rebuildDepth int) Type {
	elem := FromReflectType(rtype.Elem(), rebuildDepth)
	if rebuildDepth >= 0 {
		rtype = reflect.SliceOf(elem.rtype)
	}
	return maketype(types.NewSlice(elem.gtype), rtype)
}

// fromReflectStruct converts a reflect.Type with Kind reflect.Struct into a Type
func fromReflectStruct(rtype reflect.Type, rebuildDepth int) Type {
	n := rtype.NumField()
	fields := make([]StructField, n)
	for i := 0; i < n; i++ {
		rfield := rtype.Field(i)
		fields[i] = fromReflectField(&rfield, rebuildDepth)
	}
	vars := toGoFields(fields)
	tags := toTags(fields)

	// use reflect.StructOf to recreate reflect.Type only if requested, because it's not 100% accurate:
	// reflect.StructOf does not support unexported or anonymous fields,
	// and go/reflect cannot create named types, interfaces and self-referencing types
	if rebuildDepth >= 0 {
		rfields := toReflectFields(fields, true)
		rtype = reflect.StructOf(rfields)
	}
	return maketype(types.NewStruct(vars, tags), rtype)
}
