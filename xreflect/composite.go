/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * composite.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	r "reflect"

	"github.com/cosmos72/gomacro/go/types"
)

// ChanDir returns a channel type's direction.
// It panics if the type's Kind is not Chan.
func (t *xtype) ChanDir() r.ChanDir {
	if t.Kind() != r.Chan {
		xerrorf(t, "ChanDir of non-chan type %v", t)
	}
	gdir := t.gunderlying().(*types.Chan).Dir()
	return gdirTodir(gdir)
}

// Elem returns a type's element type.
// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
func (t *xtype) Elem() Type {
	v := t.universe
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return t.elem()
}

func (t *xtype) elem() Type {
	var gelem types.Type
	switch gtype := t.gunderlying().(type) {
	case *types.Array:
		gelem = gtype.Elem()
	case *types.Chan:
		gelem = gtype.Elem()
	case *types.Map:
		gelem = gtype.Elem()
	case *types.Pointer:
		gelem = gtype.Elem()
	case *types.Slice:
		gelem = gtype.Elem()
	default:
		xerrorf(t, "Elem of invalid type %v", t)
	}
	rtype := t.rtype
	// if rtype is xreflect.Forward due to contagion,
	// we do not know the element type -> use xreflect.Forward
	switch rtype.Kind() {
	case r.Array, r.Chan, r.Map, r.Ptr, r.Slice:
		rtype = rtype.Elem()
	}
	return t.universe.maketype(gelem, rtype, t.option)
}

// Key returns a map type's key type.
// It panics if the type's Kind is not Map.
func (t *xtype) Key() Type {
	if t.Kind() != r.Map {
		xerrorf(t, "Key of non-map type %v", t)
	}
	gkey := t.gunderlying().(*types.Map).Key()
	rtype := t.rtype
	// if rtype is xreflect.Forward due to contagion,
	// we do not know the element type -> use xreflect.Forward
	if rtype != rTypeOfForward {
		rtype = rtype.Key()
	}
	return t.universe.MakeType(gkey, rtype, t.option)
}

// Len returns an array type's length.
// It panics if the type's Kind is not Array.
func (t *xtype) Len() int {
	if t.Kind() != r.Array {
		xerrorf(t, "Len of non-array type %v", t)
	}
	return int(t.gunderlying().(*types.Array).Len())
}

func (v *Universe) ArrayOf(count int, elem Type) Type {
	e := unwrap(elem)
	return v.MakeType(
		types.NewArray(e.gtype, int64(count)),
		propagateFwd(e,
			func(rt r.Type) r.Type {
				return r.ArrayOf(count, rt)
			}),
		e.option)
}

func (v *Universe) ChanOf(dir r.ChanDir, elem Type) Type {
	e := unwrap(elem)
	gdir := dirToGdir(dir)
	return v.MakeType(
		types.NewChan(gdir, e.gtype),
		r.ChanOf(dir, e.approxReflectType()),
		e.option)
}

func (v *Universe) MapOf(key, elem Type) Type {
	k := unwrap(key)
	e := unwrap(elem)
	return v.MakeType(
		types.NewMap(k.gtype, e.gtype),
		r.MapOf(k.approxReflectType(), e.approxReflectType()),
		k.option|e.option)
}

func (v *Universe) PtrTo(elem Type) Type {
	e := unwrap(elem)
	return v.MakeType(
		types.NewPointer(e.gtype),
		propagateFwd(e, r.PtrTo),
		e.option)
}

func (v *Universe) SliceOf(elem Type) Type {
	e := unwrap(elem)
	return v.MakeType(
		types.NewSlice(e.gtype),
		propagateFwd(e, r.SliceOf),
		e.option)
}

func propagateFwd(xt *xtype, maker func(r.Type) r.Type) r.Type {
	rtype := xt.rtype
	if xt.option != OptDefault || rtype == rTypeOfForward {
		xt.option = OptRecursive
		rtype = rTypeOfForward
	} else {
		rtype = maker(rtype)
	}
	return rtype
}
