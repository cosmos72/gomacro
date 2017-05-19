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
 * composite.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"reflect"
)

// ChanDir returns a channel type's direction.
// It panics if the type's Kind is not Chan.
func (t *xtype) ChanDir() reflect.ChanDir {
	if t.Kind() != reflect.Chan {
		errorf("ChanDir of non-chan type %v", t)
	}
	return t.rtype.ChanDir()
}

// Elem returns a type's element type.
// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
func (t *xtype) Elem() Type {
	gtype := t.underlying()
	rtype := t.rtype
	switch gtype := gtype.(type) {
	case *types.Array:
		return MakeType(gtype.Elem(), rtype.Elem())
	case *types.Chan:
		return MakeType(gtype.Elem(), rtype.Elem())
	case *types.Map:
		return MakeType(gtype.Elem(), rtype.Elem())
	case *types.Pointer:
		return MakeType(gtype.Elem(), rtype.Elem())
	case *types.Slice:
		return MakeType(gtype.Elem(), rtype.Elem())
	default:
		errorf("Elem of invalid type %v", t)
		return nil
	}
}

// Key returns a map type's key type.
// It panics if the type's Kind is not Map.
func (t *xtype) Key() Type {
	if t.Kind() != reflect.Map {
		errorf("Key of non-map type %v", t)
	}
	gtype := t.underlying().(*types.Map)
	return MakeType(gtype.Key(), t.rtype.Key())
}

// Len returns an array type's length.
// It panics if the type's Kind is not Array.
func (t *xtype) Len() int {
	if t.Kind() != reflect.Func {
		errorf("Len of non-array type %v", t)
	}
	return t.rtype.Len()
}

func ArrayOf(count int, elem Type) Type {
	return MakeType(
		types.NewArray(elem.GoType(), int64(count)),
		reflect.ArrayOf(count, elem.ReflectType()))
}

func ChanOf(dir reflect.ChanDir, elem Type) Type {
	gdir := dirToGdir(dir)
	return MakeType(
		types.NewChan(gdir, elem.GoType()),
		reflect.ChanOf(dir, elem.ReflectType()))
}

func MapOf(key, elem Type) Type {
	return MakeType(
		types.NewMap(key.GoType(), elem.GoType()),
		reflect.MapOf(key.ReflectType(), elem.ReflectType()))
}

func PtrTo(elem Type) Type {
	return MakeType(
		types.NewPointer(elem.GoType()),
		reflect.PtrTo(elem.ReflectType()))
}

func SliceOf(elem Type) Type {
	return MakeType(
		types.NewSlice(elem.GoType()),
		reflect.SliceOf(elem.ReflectType()))
}
