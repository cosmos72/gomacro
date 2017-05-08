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
	"go/types"
	"reflect"
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
		t = basictypes[k]
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
	return ArrayOf(count, elem)
}

func fromReflectChan(rtype reflect.Type, opts FromReflectOpts) Type {
	dir := rtype.ChanDir()
	elem := FromReflectType(rtype.Elem(), opts)
	return ChanOf(dir, elem)
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
	return FuncOf(in, out, rtype.IsVariadic())
}

func fromReflectInterface(rtype reflect.Type, opts FromReflectOpts) Type {
	if rtype == TypeOfInterface.rtype {
		return TypeOfInterface
	}
	errorf("unimplemented: fromReflectInterface()")
	return Type{}
}

func fromReflectMap(rtype reflect.Type, opts FromReflectOpts) Type {
	key := FromReflectType(rtype.Key(), opts)
	elem := FromReflectType(rtype.Elem(), opts)
	return MapOf(key, elem)
}

func fromReflectPtr(rtype reflect.Type, opts FromReflectOpts) Type {
	elem := FromReflectType(rtype.Elem(), opts)
	return PtrTo(elem)
}

func fromReflectSlice(rtype reflect.Type, opts FromReflectOpts) Type {
	elem := FromReflectType(rtype.Elem(), opts)
	return SliceOf(elem)
}

func fromReflectStruct(rtype reflect.Type, opts FromReflectOpts) Type {
	n := rtype.NumField()
	fields := make([]StructField, n)
	for i := 0; i < n; i++ {
		rfield := rtype.Field(i)
		fields[i] = fromReflectField(&rfield, opts)
	}
	vars := toGoFields(fields)
	tags := toTags(fields)
	gtype := types.NewStruct(vars, tags)

	// use reflect.StructOf to recreate reflect.Type only if requested, because it's not 100% accurate:
	// reflect.StructOf does not support unexported or anonymous fields, and go/reflect cannot create named types
	if opts&RebuildReflectType != 0 {
		rfields := toReflectFields(fields, true)
		rtype = reflect.StructOf(rfields)
	}
	return maketype(gtype, rtype)
}
