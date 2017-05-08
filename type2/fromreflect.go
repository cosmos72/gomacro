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

type FromReflectOpts int

const (
	Defaults           FromReflectOpts = 0
	RebuildReflectType FromReflectOpts = 1
)

// FromReflectType creates a Type corresponding to given reflect.Type
// Note: conversions from Type to reflect.Type and back are not exact,
// because of the reasons listed in Type.ReflectType()
// Conversions from reflect.Type to Type and back are not exact for the same reasons.
func FromReflectType(rtype reflect.Type, opts FromReflectOpts) Type {
	var t Type
	switch k := rtype.Kind(); k {
	case reflect.Invalid:
		return Type{}
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		t = BasicTypes[k]
	case reflect.Array:
		t = fromReflectArray(rtype, opts)
	case reflect.Chan:
		t = fromReflectChan(rtype, opts)
	case reflect.Func:
		t = fromReflectFunc(rtype, opts)
	case reflect.Interface:
		t = fromReflectInterface(rtype, opts)
	case reflect.Map:
		t = fromReflectMap(rtype, opts)
	case reflect.Ptr:
		t = fromReflectPtr(rtype, opts)
	case reflect.Slice:
		t = fromReflectSlice(rtype, opts)
	case reflect.Struct:
		t = fromReflectStruct(rtype, opts)
	// case reflect.UnsafePointer:
	default:
		errorf("unsupported reflect.Type %v", rtype)
	}
	return finish(t, rtype, opts)
}

func finish(t Type, rtype reflect.Type, opts FromReflectOpts) Type {
	name := rtype.Name()
	pkgpath := rtype.PkgPath()
	if pkgpath == t.PkgPath() && name == t.Name() {
		return addmethods(t, rtype, opts)
	}
	underlying := maketype(t.underlying(), t.rtype)
	if len(name) == 0 {
		return underlying
	}
	pkg := NewPackage(pkgpath, "")
	t = NamedOf(name, pkg)
	t.SetUnderlying(underlying)
	return addmethods(t, rtype, opts)
}

func addmethods(t Type, rtype reflect.Type, opts FromReflectOpts) Type {
	n := rtype.NumMethod()
	if n == 0 {
		return t
	}
	if !t.Named() {
		errorf("cannot add methods to unnamed type %v", t)
	}
	for i := 0; i < n; i++ {
		rmethod := rtype.Method(i)
		signature := fromReflectFunc(rmethod.Type, opts)
		t.AddMethod(rmethod.Name, signature)
	}
	return t
}

func fromReflectArray(rtype reflect.Type, opts FromReflectOpts) Type {
	count := rtype.Len()
	elem := FromReflectType(rtype.Elem(), opts)
	if opts&RebuildReflectType != 0 {
		rtype = reflect.ArrayOf(count, elem.rtype)
	}
	return maketype(types.NewArray(elem.gtype, int64(count)), rtype)
}

func fromReflectChan(rtype reflect.Type, opts FromReflectOpts) Type {
	dir := rtype.ChanDir()
	elem := FromReflectType(rtype.Elem(), opts)
	if opts&RebuildReflectType != 0 {
		rtype = reflect.ChanOf(dir, elem.rtype)
	}
	gdir := dirToGdir(dir)
	return maketype(types.NewChan(gdir, elem.gtype), rtype)
}

func fromReflectFunc(rtype reflect.Type, opts FromReflectOpts) Type {
	nin, nout := rtype.NumIn(), rtype.NumOut()
	in := make([]Type, nin)
	out := make([]Type, nout)
	for i := 0; i < nin; i++ {
		in[i] = FromReflectType(rtype.In(i), opts)
	}
	for i := 0; i < nout; i++ {
		out[i] = FromReflectType(rtype.Out(i), opts)
	}
	gin := toGoTuple(in)
	gout := toGoTuple(out)
	variadic := rtype.IsVariadic()

	if opts&RebuildReflectType != 0 {
		rin := toReflectTypes(in)
		rout := toReflectTypes(out)
		rtype = reflect.FuncOf(rin, rout, variadic)
	}
	return maketype(
		types.NewSignature(nil, gin, gout, variadic),
		rtype,
	)
}

func fromReflectInterface(rtype reflect.Type, opts FromReflectOpts) Type {
	if rtype == TypeOfInterface.rtype {
		return TypeOfInterface
	}
	n := rtype.NumMethod()
	gmethods := make([]*types.Func, n)
	for i := 0; i < n; i++ {
		rmethod := rtype.Method(i)
		method := fromReflectFunc(rmethod.Type, opts)
		gmethods[i] = types.NewFunc(token.NoPos, NewPackage(rmethod.PkgPath, "").impl, rmethod.Name, method.gtype.(*types.Signature))
	}
	// no way to extract embedded interfaces from reflect.Type
	if opts&RebuildReflectType != 0 {
		rfields := make([]reflect.StructField, 1+n)
		rfields[0] = approxInterfaceSelf()
		for i := 0; i < n; i++ {
			rmethod := rtype.Method(i)
			rfields[i+1] = approxInterfaceMethod(rmethod.Name, rmethod.Type)

		}
		rtype = reflect.StructOf(rfields)
	}
	return maketype(types.NewInterface(gmethods, nil), rtype)
}

func fromReflectMap(rtype reflect.Type, opts FromReflectOpts) Type {
	key := FromReflectType(rtype.Key(), opts)
	elem := FromReflectType(rtype.Elem(), opts)
	if opts&RebuildReflectType != 0 {
		rtype = reflect.MapOf(key.rtype, elem.rtype)
	}
	return maketype(types.NewMap(key.gtype, elem.gtype), rtype)
}

func fromReflectPtr(rtype reflect.Type, opts FromReflectOpts) Type {
	elem := FromReflectType(rtype.Elem(), opts)
	if opts&RebuildReflectType != 0 {
		rtype = reflect.PtrTo(elem.rtype)
	}
	return maketype(types.NewPointer(elem.gtype), rtype)
}

func fromReflectSlice(rtype reflect.Type, opts FromReflectOpts) Type {
	elem := FromReflectType(rtype.Elem(), opts)
	if opts&RebuildReflectType != 0 {
		rtype = reflect.SliceOf(elem.rtype)
	}
	return maketype(types.NewSlice(elem.gtype), rtype)
}

func fromReflectStruct(rtype reflect.Type, opts FromReflectOpts) Type {
	n := rtype.NumField()
	fields := make([]StructField, n)
	for i := 0; i < n; i++ {
		rfield := rtype.Field(i)
		name := rfield.Name
		if name == StrGensymInterface {
			errorf("unimplemented: fromReflectStruct() of generated struct that emulates interface: %v", rtype)
		} else if strings.HasPrefix(name, StrGensymEmbedded) {
			errorf("unimplemented: fromReflectStruct() of generated struct that emulates embedded fields: %v", rtype)
		} else if strings.HasPrefix(name, StrGensymPrivate) {
			errorf("unimplemented: fromReflectStruct() of generated struct that emulates private fields: %v", rtype)
		}
		fields[i] = fromReflectField(&rfield, opts)
	}
	vars := toGoFields(fields)
	tags := toTags(fields)

	// use reflect.StructOf to recreate reflect.Type only if requested, because it's not 100% accurate:
	// reflect.StructOf does not support unexported or anonymous fields,
	// and go/reflect cannot create named types, interfaces and self-referencing types
	if opts&RebuildReflectType != 0 {
		rfields := toReflectFields(fields, true)
		rtype = reflect.StructOf(rfields)
	}
	return maketype(types.NewStruct(vars, tags), rtype)
}
