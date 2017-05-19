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

package xreflect

import (
	"go/token"
	"go/types"
	"reflect"
	"strings"
	"sync"
)

var (
	cache = &Cache{}
	lock  = sync.Mutex{}
)

// TypeOf creates a Type corresponding to reflect.TypeOf() of given value.
// Note: conversions from Type to reflect.Type and back are not exact,
// because of the reasons listed in Type.ReflectType()
// Conversions from reflect.Type to Type and back are not exact for the same reasons.
func TypeOf(rvalue interface{}) Type {
	return FromReflectType(reflect.TypeOf(rvalue), nil)
}

// FromReflectType creates a Type corresponding to given reflect.Type
// Note: conversions from Type to reflect.Type and back are not exact,
// because of the reasons listed in Type.ReflectType()
// Conversions from reflect.Type to Type and back are not exact for the same reasons.
func FromReflectType(rtype reflect.Type, cfg *Cache) Type {
	if cfg == nil {
		lock.Lock()
		cfg = cache
		cache = nil
		lock.Unlock()
		if cfg == nil {
			cfg = &Cache{}
		}
		defer func() {
			lock.Lock()
			if cache == nil {
				cache = cfg
			}
			lock.Unlock()
		}()
	}
	return cfg.FromReflectType(rtype)
}

func (cfg *Cache) TypeOf(rvalue interface{}) Type {
	return cfg.FromReflectType(reflect.TypeOf(rvalue))
}

func (cfg *Cache) FromReflectType(rtype reflect.Type) Type {
	if rtype == nil {
		return nil
	}
	t := BasicTypes[rtype.Kind()]
	if t != nil && t.ReflectType() == rtype {
		return t
	}
	if t = cfg.ReflectTypes[rtype]; t != nil {
		// debugf("found type in cache: %v -> %v (%v)", rtype, t, t.ReflectType())
		// time.Sleep(100 * time.Millisecond)
		return t
	}
	name := rtype.Name()
	tryresolve := cfg.TryResolve
	if tryresolve != nil && len(name) != 0 {
		t = tryresolve(name, rtype.PkgPath())
		if t != nil {
			return t
		}
	}
	if cfg.RebuildDepth >= 0 {
		// decrement ONLY here and in fromReflectPtr() when calling fromReflectInterfaceStruct()
		cfg.RebuildDepth--
		defer func() {
			cfg.RebuildDepth++
		}()
	}
	// when converting a named type and cfg.Importer cannot locate it,
	// immediately register it in the cache because it may reference itself,
	// as for example type List struct { Elem int; Rest *List }
	// otherwise we may get an infinite recursion
	if len(name) != 0 {
		if !cfg.rebuild() {
			if t = cfg.namedTypeFromImport(rtype); t != nil {
				// debugf("found type in import: %v -> %v", t, t.ReflectType())
				return t
			}
		}
		t = NamedOf(name, cfg.NewPackage(rtype.PkgPath(), ""))
		cfg.cache(rtype, t) // support self-refencing types
		// debugf("prepared named type %v", t)
	}

	var u Type
	switch k := rtype.Kind(); k {
	case reflect.Invalid:
		return nil
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String,
		reflect.UnsafePointer:
		u = BasicTypes[k]
	case reflect.Array:
		u = cfg.fromReflectArray(rtype)
	case reflect.Chan:
		u = cfg.fromReflectChan(rtype)
	case reflect.Func:
		u = cfg.fromReflectFunc(rtype)
	case reflect.Interface:
		u = cfg.fromReflectInterface(rtype)
	case reflect.Map:
		u = cfg.fromReflectMap(rtype)
	case reflect.Ptr:
		u = cfg.fromReflectPtr(rtype)
	case reflect.Slice:
		u = cfg.fromReflectSlice(rtype)
	case reflect.Struct:
		u = cfg.fromReflectStruct(rtype)
	default:
		errorf("unsupported reflect.Type %v", rtype)
	}
	if t == nil {
		t = u
		// cache before adding methods - otherwise we get an infinite recursion
		// if u is a pointer to named type with methods that reference the named type
		cfg.cache(rtype, t)
	} else {
		t.SetUnderlying(u)
		// t.ReflectType() is now u.ReflectType(). but we can do better... we know the exact rtype to set
		if !cfg.rebuild() {
			t[0].rtype = rtype
		}
	}
	return cfg.addmethods(t, rtype)
}

func (cfg *Cache) addmethods(t Type, rtype reflect.Type) Type {
	n := rtype.NumMethod()
	if n == 0 {
		return t
	}
	tm := t
	if !t.Named() && t.Kind() == reflect.Ptr {
		// methods on pointer-to-type. add them to the type itself
		tm = t.Elem()
	}
	if !tm.Named() {
		errorf("cannot add methods to unnamed type %v", t)
	}
	xt := &tm[0]
	if xt.methodvalues != nil {
		// prevent another infinite recursion: Type.AddMethod() may reference the type itself in its methods
		// debugf("NOT adding again %d methods to %v", n, tm)
	} else {
		// debugf("adding %d methods to %v", n, tm)
		xt.methodvalues = make([]reflect.Value, 0, n)
		nilv := reflect.Value{}
		for i := 0; i < n; i++ {
			rmethod := rtype.Method(i)
			signature := cfg.FromReflectType(rmethod.Type)
			n1 := tm.NumMethod()
			tm.AddMethod(rmethod.Name, signature)
			n2 := tm.NumMethod()
			if n1 == n2 {
				// method was already present
				continue
			}
			for len(xt.methodvalues) < n2 {
				xt.methodvalues = append(xt.methodvalues, nilv)
			}
			xt.methodvalues[n1] = rmethod.Func
		}
	}
	return t
}

func (cfg *Cache) fromReflectField(rfield *reflect.StructField) StructField {
	t := cfg.FromReflectType(rfield.Type)
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
		t = cfg.named(t, typename, rtype.PkgPath())
		name = ""
		anonymous = true
	} else if strings.HasPrefix(name, StrGensymPrivate) {
		// this reflect.StructField emulates private (unexported) field using our own convention.
		// eat our own dogfood and convert it back to a private field.
		name = name[len(StrGensymPrivate):]
	}

	return StructField{
		Name:      name,
		Pkg:       cfg.NewPackage(rfield.PkgPath, ""),
		Type:      t,
		Tag:       rfield.Tag,
		Offset:    rfield.Offset,
		Index:     rfield.Index,
		Anonymous: anonymous,
	}
}

func (cfg *Cache) fromReflectFields(rfields []reflect.StructField) []StructField {
	fields := make([]StructField, len(rfields))
	for i := range rfields {
		fields[i] = cfg.fromReflectField(&rfields[i])
	}
	return fields
}

// named creates a new named Type based on t, having the given name and pkgpath
func (cfg *Cache) named(t Type, name string, pkgpath string) Type {
	if t.Name() != name || t.PkgPath() != pkgpath {
		t2 := NamedOf(name, cfg.NewPackage(pkgpath, ""))
		t2.SetUnderlying(MakeType(t.underlying(), t.ReflectType()))
		t = t2
	}
	return t
}

// fromReflectArray converts a reflect.Type with Kind reflect.Array into a Type
func (cfg *Cache) fromReflectArray(rtype reflect.Type) Type {
	count := rtype.Len()
	elem := cfg.FromReflectType(rtype.Elem())
	if cfg.rebuild() {
		rtype = reflect.ArrayOf(count, elem.ReflectType())
	}
	return MakeType(types.NewArray(elem.GoType(), int64(count)), rtype)
}

// fromReflectChan converts a reflect.Type with Kind reflect.Chan into a Type
func (cfg *Cache) fromReflectChan(rtype reflect.Type) Type {
	dir := rtype.ChanDir()
	elem := cfg.FromReflectType(rtype.Elem())
	if cfg.rebuild() {
		rtype = reflect.ChanOf(dir, elem.ReflectType())
	}
	gdir := dirToGdir(dir)
	return MakeType(types.NewChan(gdir, elem.GoType()), rtype)
}

// fromReflectFunc converts a reflect.Type with Kind reflect.Func into a function Type
func (cfg *Cache) fromReflectFunc(rtype reflect.Type) Type {
	nin, nout := rtype.NumIn(), rtype.NumOut()
	in := make([]Type, nin)
	out := make([]Type, nout)
	for i := 0; i < nin; i++ {
		in[i] = cfg.FromReflectType(rtype.In(i))
	}
	for i := 0; i < nout; i++ {
		out[i] = cfg.FromReflectType(rtype.Out(i))
	}
	gin := toGoTuple(in)
	gout := toGoTuple(out)
	variadic := rtype.IsVariadic()

	if cfg.rebuild() {
		rin := toReflectTypes(in)
		rout := toReflectTypes(out)
		rtype = reflect.FuncOf(rin, rout, variadic)
	}
	return MakeType(
		types.NewSignature(nil, gin, gout, variadic),
		rtype,
	)
}

// fromReflectInterface converts a reflect.Type with Kind reflect.Interface into a Type
func (cfg *Cache) fromReflectInterface(rtype reflect.Type) Type {
	if rtype == TypeOfInterface.ReflectType() {
		return TypeOfInterface
	}
	n := rtype.NumMethod()
	gmethods := make([]*types.Func, n)
	pkgs := make(map[string]*types.Package)
	for i := 0; i < n; i++ {
		rmethod := rtype.Method(i)
		method := cfg.FromReflectType(rmethod.Type)
		pkg := pkgs[rmethod.PkgPath]
		if pkg == nil {
			pkg = (*types.Package)(cfg.NewPackage(rmethod.PkgPath, ""))
		}
		gmethods[i] = types.NewFunc(token.NoPos, pkg, rmethod.Name, method.GoType().(*types.Signature))
	}
	// no way to extract embedded interfaces from reflect.Type
	if cfg.rebuild() {
		rfields := make([]reflect.StructField, 1+n)
		rfields[0] = approxInterfaceHeader()
		for i := 0; i < n; i++ {
			rmethod := rtype.Method(i)
			rmethodtype := rmethod.Type
			if cfg.RebuildDepth >= 1 {
				// needed? method := cfg.FromReflectType(rmethod.Type) above
				// should already rebuild rmethod.Type.ReflectType()
				rmethodtype = cfg.FromReflectType(rmethod.Type).ReflectType()
			}
			rfields[i+1] = approxInterfaceMethod(rmethod.Name, rmethodtype)
		}
		// interfaces may have lots of methods, thus a lot of fields in the proxy struct.
		// Then use a pointer to the proxy struct: InterfaceOf() does that, and we must behave identically
		rtype = reflect.PtrTo(reflect.StructOf(rfields))
	}
	return MakeType(types.NewInterface(gmethods, nil), rtype)
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
func (cfg *Cache) fromReflectInterfaceStruct(rtype reflect.Type) Type {
	rebuild := cfg.rebuild()
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
			t := cfg.FromReflectType(rfield.Type)
			if t.Kind() != reflect.Interface {
				errorf("FromReflectType: reflect.Type <%v> is an emulated interface containing the embedded interface <%v>.\n\tExtracting the latter returned a non-interface: %v", t)
			}
			t = cfg.named(t, typename, rfield.Type.PkgPath())
			gembeddeds = append(gembeddeds, t.GoType().(*types.Named))
			if rebuild {
				rebuildfields[i] = approxInterfaceEmbedded(typename, t.ReflectType())
			}
		} else {
			if strings.HasPrefix(name, StrGensymPrivate) {
				name = name[len(StrGensymPrivate):]
			}
			t := cfg.fromReflectFunc(rfield.Type)
			if t.Kind() != reflect.Func {
				errorf("FromReflectType: reflect.Type <%v> is an emulated interface containing the method <%v>.\n\tExtracting the latter returned a non-function: %v", t)
			}
			gtype := t.GoType()
			if t.Named() {
				gtype = gtype.Underlying()
			}
			pkg := cfg.Pkgs[rfield.PkgPath]
			if pkg == nil {
				pkg = cfg.NewPackage(rfield.PkgPath, "")
			}
			gmethods = append(gmethods, types.NewFunc(token.NoPos, (*types.Package)(pkg), name, gtype.(*types.Signature)))
			if rebuild {
				rebuildfields[i] = approxInterfaceMethod(name, t.ReflectType())
			}
		}
	}
	if rebuild {
		rtype = reflect.StructOf(rebuildfields)
	}
	return MakeType(types.NewInterface(gmethods, gembeddeds), rtype)
}

// fromReflectMap converts a reflect.Type with Kind reflect.map into a Type
func (cfg *Cache) fromReflectMap(rtype reflect.Type) Type {
	key := cfg.FromReflectType(rtype.Key())
	elem := cfg.FromReflectType(rtype.Elem())
	if cfg.rebuild() {
		rtype = reflect.MapOf(key.ReflectType(), elem.ReflectType())
	}
	return MakeType(types.NewMap(key.GoType(), elem.GoType()), rtype)
}

// fromReflectPtr converts a reflect.Type with Kind reflect.Ptr into a Type
func (cfg *Cache) fromReflectPtr(rtype reflect.Type) Type {
	relem := rtype.Elem()
	var elem Type
	var gtype types.Type
	rebuild := cfg.rebuild()
	if isReflectInterfaceStruct(relem) {
		if rebuild {
			cfg.RebuildDepth--
			defer func() {
				cfg.RebuildDepth++
			}()
		}
		elem = cfg.fromReflectInterfaceStruct(relem)
		gtype = elem.GoType()
	} else {
		elem = cfg.FromReflectType(relem)
		gtype = types.NewPointer(elem.GoType())
	}
	if rebuild {
		rtype = reflect.PtrTo(elem.ReflectType())
	}
	return MakeType(gtype, rtype)
}

// fromReflectPtr converts a reflect.Type with Kind reflect.Slice into a Type
func (cfg *Cache) fromReflectSlice(rtype reflect.Type) Type {
	elem := cfg.FromReflectType(rtype.Elem())
	if cfg.rebuild() {
		rtype = reflect.SliceOf(elem.ReflectType())
	}
	return MakeType(types.NewSlice(elem.GoType()), rtype)
}

// fromReflectStruct converts a reflect.Type with Kind reflect.Struct into a Type
func (cfg *Cache) fromReflectStruct(rtype reflect.Type) Type {
	n := rtype.NumField()
	fields := make([]StructField, n)
	for i := 0; i < n; i++ {
		rfield := rtype.Field(i)
		fields[i] = cfg.fromReflectField(&rfield)
	}
	vars := toGoFields(fields)
	tags := toTags(fields)

	// use reflect.StructOf to recreate reflect.Type only if requested, because it's not 100% accurate:
	// reflect.StructOf does not support unexported or anonymous fields,
	// and go/reflect cannot create named types, interfaces and self-referencing types
	if cfg.rebuild() {
		rfields := toReflectFields(fields, true)
		rtype = reflect.StructOf(rfields)
	}
	return MakeType(types.NewStruct(vars, tags), rtype)
}
