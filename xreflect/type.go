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
 * type.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"reflect"

	"github.com/cosmos72/gomacro/typeutil"
)

func identicalType(t, u Type) bool {
	xt := unwrap(t)
	yt := unwrap(u)
	xnil := xt == nil
	ynil := yt == nil
	if xnil || ynil {
		return xnil == ynil
	}
	return xt == yt || xt.identicalTo(yt)
}

// go/types.InterfaceOf(methods, embeddeds) modifies the methods argument and adds receiver to them,
// thwarting our attempts to keep go/types.Type and reflect.Type in sync in the cache.
// So it would be safer to set enableCache = false, but...
//
// Currently, enableCache = false breaks TestFast/method_on_ptr (and possibly other tests)
// FIXME: fix TestFast/method_on_ptr, then consider setting enableCache = false
const (
	enableCache = true

	enableWarnOnCacheMismatch = false
)

func warnOnCacheMismatch(gtype types.Type, rtype reflect.Type, rcached reflect.Type) {
	if enableWarnOnCacheMismatch {
		debugf("overwriting mismatched reflect.Type found in cache for type %v:\n\tcurrent reflect.Type: %v\n\tcached  reflect.Type: %v",
			gtype, rtype, rcached)
	}
}

func (m *Types) add(t Type) {
	if enableCache {
		switch t.Kind() {
		case reflect.Func:
			t.NumIn() // check consistency

		case reflect.Interface:
			rtype := t.ReflectType()
			rkind := rtype.Kind()
			if rkind != reflect.Interface && (rkind != reflect.Ptr || rtype.Elem().Kind() != reflect.Struct) {
				errorf(t, "bug! inconsistent type <%v>: has kind = %s but its Type.Reflect() is %s\n\tinstead of interface or pointer-to-struct: <%v>", t, t.Kind(), rtype.Kind(), t.ReflectType())
			}
		}
		m.gmap.Set(t.GoType(), t)
		// debugf("added type to cache: %v <%v> <%v>", t.Kind(), t.GoType(), t.ReflectType())
	}
}

func (v *Universe) unique(t Type) Type {
	gtype := t.GoType()
	ret := v.Types.gmap.At(gtype)
	if ret != nil {
		tret := ret.(Type)
		if tret.ReflectType() == t.ReflectType() {
			return t
		}
		warnOnCacheMismatch(gtype, t.ReflectType(), tret.ReflectType())
	}
	v.add(t)
	return t
}

// all unexported methods assume lock is already held
func (v *Universe) maketype3(kind reflect.Kind, gtype types.Type, rtype reflect.Type) Type {
	if gtype == nil {
		errorf(nil, "MakeType of nil types.Type")
	} else if rtype == nil {
		errorf(nil, "MakeType of nil reflect.Type")
	}
	ret := v.Types.gmap.At(gtype)
	if ret != nil {
		t := ret.(Type)
		if t.ReflectType() == rtype {
			return t
		}
		warnOnCacheMismatch(gtype, rtype, t.ReflectType())
	}
	if v.BasicTypes == nil {
		// lazy creation of basic types
		v.init()
	}
	t := wrap(&xtype{kind: kind, gtype: gtype, rtype: rtype, universe: v})
	v.add(t)
	return t
}

func (v *Universe) maketype(gtype types.Type, rtype reflect.Type) Type {
	return v.maketype3(gtypeToKind(nil, gtype), gtype, rtype)
}

func (v *Universe) MakeType(gtype types.Type, rtype reflect.Type) Type {
	kind := gtypeToKind(nil, gtype)
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return v.maketype3(kind, gtype, rtype)
}

// GoType returns the go/types.Type corresponding to the type.
func (t *xtype) GoType() types.Type {
	return t.gtype
}

// ReflectType returns a best-effort reflect.Type that approximates the type.
// It may be inexact for the following reasons:
// 1) missing reflect.NamedOf(): no way to programmatically create named types, or to access the underlying type of a named type
// 2) missing reflect.InterfaceOf(): interface types created at runtime will be approximated by structs
// 3) missing reflect.MethodOf(): method types created at runtime will be approximated by functions
//    whose first parameter is the receiver
// 4) reflect.StructOf() does not support embedded or unexported fields
// 5) go/reflect lacks the ability to create self-referencing types:
//    references to the type itself will be replaced by interface{}.
//
// Examples:
//    after invoking at runtime type2.NewStruct() and type2.NewNamed()
//    to create the following type:
//        type List struct { Elem int; Rest *List }
//    ReflectType will return a reflect.Type equivalent to:
//        struct { Elem int; Rest interface{} }
//    i.e. the type name will be missing due to limitation 1 above,
//    and the field 'Rest' will have type interface{} instead of *List due to limitation 5.
func (t *xtype) ReflectType() reflect.Type {
	return t.rtype
}

func (t *xtype) UnsafeForceReflectType(rtype reflect.Type) {
	t.rtype = rtype
}

func (t *xtype) Universe() *Universe {
	return t.universe
}

// Named returns whether the type is named.
// It returns false for unnamed types.
func (t *xtype) Named() bool {
	if t != nil {
		switch t.gtype.(type) {
		case *types.Basic, *types.Named:
			return true
		}
	}
	return false
}

// Name returns the type's name within its package.
// It returns an empty string for unnamed types.
func (t *xtype) Name() string {
	if t != nil {
		switch gtype := t.gtype.(type) {
		case *types.Basic:
			return gtype.Name()
		case *types.Named:
			return gtype.Obj().Name()
		}
	}
	return ""
}

// Pkg returns a named type's package, that is, the package where it was defined.
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// Pkg will return nil.
func (t *xtype) Pkg() *Package {
	switch gtype := t.gtype.(type) {
	case *types.Named:
		return (*Package)(gtype.Obj().Pkg())
	default:
		return nil
	}
}

// PkgName returns a named type's package name, that is,
// the default name that the package provides when imported.
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// the package name will be the empty string.
func (t *xtype) PkgName() string {
	switch gtype := t.gtype.(type) {
	case *types.Named:
		return gtype.Obj().Pkg().Name()
	default:
		return ""
	}
}

// PkgPath returns a named type's package path, that is, the import path
// that uniquely identifies the package, such as "encoding/base64".
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// the package path will be the empty string.
func (t *xtype) PkgPath() string {
	switch gtype := t.gtype.(type) {
	case *types.Named:
		return gtype.Obj().Pkg().Path()
	default:
		return ""
	}
}

// Size returns the number of bytes needed to store
// a value of the given type; it is analogous to unsafe.Sizeof.
func (t *xtype) Size() uintptr {
	return t.rtype.Size()
}

// String returns a string representation of a type.
func (t *xtype) String() string {
	if t == nil {
		return "<nil>"
	}
	return t.gtype.String()
}

/*
// Underlying returns the underlying type of a type.
func (t *xtype) Underlying() Type {
	return Type{t.underlying}
}
*/

func (t *xtype) gunderlying() types.Type {
	return t.gtype.Underlying()
}

// Kind returns the specific kind of the type.
func (t *xtype) Kind() reflect.Kind {
	if t == nil {
		return reflect.Invalid
	}
	return t.kind
}

// Implements reports whether the type implements the interface type u.
// It panics if u's Kind is not Interface
func (t *xtype) Implements(u Type) bool {
	if u.Kind() != reflect.Interface {
		xerrorf(t, "Type.Implements of non-interface type: %v", u)
	}
	return t.gtype == u.GoType() || types.Implements(t.gtype, u.GoType().Underlying().(*types.Interface))
}

// IdenticalTo reports whether the type is identical to type u.
func (t *xtype) IdenticalTo(u Type) bool {
	xu := unwrap(u)
	return t == xu || t.identicalTo(xu)
}

func (t *xtype) identicalTo(u *xtype) bool {
	return typeutil.Identical(t.GoType(), u.GoType())
}

// AssignableTo reports whether a value of the type is assignable to type u.
func (t *xtype) AssignableTo(u Type) bool {
	// debugf("AssignableTo: <%v> <%v>", t, u)
	return t.gtype == u.GoType() || types.AssignableTo(t.gtype, u.GoType())
}

// ConvertibleTo reports whether a value of the type is convertible to type u.
func (t *xtype) ConvertibleTo(u Type) bool {
	return t.gtype == u.GoType() || types.ConvertibleTo(t.gtype, u.GoType())
}

// Comparable reports whether values of this type are comparable.
func (t *xtype) Comparable() bool {
	return types.Comparable(t.gtype)
}

// Zero returns a Value representing the zero value for the specified type.
func Zero(t Type) reflect.Value {
	return reflect.Zero(t.ReflectType())
}
