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
 * lookup.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
)

// FieldByName returns the (possibly embedded) struct field with given name,
// and the number of fields found at the same (shallowest) depth: 0 if not found.
// Private fields are returned only if they were declared in pkgpath.
func (t *xtype) FieldByName(name, pkgpath string) (field StructField, count int) {
	if name == "_" {
		return
	}
	// debugf("field cache for %v <%v> = %v", unsafe.Pointer(t), t, t.fieldcache)
	qname := QName2(name, pkgpath)

	field, found := t.fieldcache[qname]
	if found {
		if field.Index == nil { // marker for ambiguous field names
			count = int(field.Offset) // reuse Offset as "number of ambiguous fields"
		} else {
			count = 1
		}
		return field, count
	}
	var tovisit []StructField

	field, count, tovisit = fieldByName(t, qname, 0, nil)

	// breadth-first recursion
	for count == 0 && len(tovisit) != 0 {
		var next []StructField
		for _, f := range tovisit {
			efield, ecount, etovisit := fieldByName(unwrap(f.Type), qname, f.Offset, f.Index)
			if count == 0 {
				if ecount > 0 {
					field = efield
				} else {
					// no recursion if we found something
					next = append(next, etovisit...)
				}
			}
			count += ecount
		}
		tovisit = next
	}
	if count > 0 {
		cacheFieldByName(t, qname, &field, count)
	}
	return field, count
}

func fieldByName(t *xtype, qname QName, offset uintptr, index []int) (field StructField, count int, tovisit []StructField) {
	gtype := t.gtype.Underlying().(*types.Struct)
	n := t.NumField()
	for i := 0; i < n; i++ {

		gfield := gtype.Field(i)
		if matchFieldByName(qname, gfield) {
			if count == 0 {
				field = t.Field(i)
				field.Offset += offset
				field.Index = concat(index, field.Index) // make a copy of index
				// debugf("fieldByName: %d-th field of <%v> matches: %#v", i, t.rtype, field)
			}
			count++
		} else if count == 0 && gfield.Anonymous() {
			efield := t.Field(i)
			efield.Offset += offset
			efield.Index = concat(index, efield.Index) // make a copy of index
			// debugf("fieldByName: %d-th field of <%v> is anonymous: %#v", i, t.rtype, efield)
			tovisit = append(tovisit, efield)
		}
	}
	return
}

// return true if gfield name matches given name, or if it's anonymous and its *type* name matches given name
func matchFieldByName(qname QName, gfield *types.Var) bool {
	// always check the field's package, not the type's package
	if qname == QNameGo(gfield) {
		return true
	}
	if gfield.Anonymous() {
		switch gtype := gfield.Type().(type) {
		case *types.Basic:
			// is it possible to embed basic types?
			// yes, and they work as unexported embedded fields,
			// i.e. in the same package as the struct that includes them
			return qname == QNameGo2(gtype.Name(), gfield.Pkg())
		case *types.Named:
			// gtype.Obj().Pkg() and gfield.Pkg() should be identical for *unexported* fields
			// (they are ignored for exported fields)
			return qname == QNameGo2(gtype.Obj().Name(), gfield.Pkg())
		}
	}
	return false
}

// add field to type's fieldcache. used by Type.FieldByName after a successful lookup
func cacheFieldByName(t *xtype, qname QName, field *StructField, count int) {
	if t.fieldcache == nil {
		t.fieldcache = make(map[QName]StructField)
	}
	if count > 1 {
		field.Index = nil             // marker for ambiguous field names
		field.Offset = uintptr(count) // reuse Offset as "number of ambiguous fields"
	}
	t.fieldcache[qname] = *field
}

// anonymousFields returns the anonymous fields of a (named or unnamed) struct type
func anonymousFields(t *xtype, offset uintptr, index []int) []StructField {
	var tovisit []StructField
	gtype := t.gtype.Underlying().(*types.Struct)
	n := t.NumField()
	for i := 0; i < n; i++ {
		gfield := gtype.Field(i)
		if gfield.Anonymous() {
			field := t.Field(i)
			field.Offset += offset
			field.Index = concat(index, field.Index) // make a copy of index
			tovisit = append(tovisit, field)
		}
	}
	return tovisit
}

// MethodByName returns the method with given name (including wrapper methods for embedded fields)
// and the number of methods found at the same (shallowest) depth: 0 if not found.
// Private methods are returned only if they were declared in pkgpath.
func (t *xtype) MethodByName(name, pkgpath string) (method Method, count int) {
	// debugf("method cache for %v <%v> = %v", unsafe.Pointer(t), t, t.methodcache)
	if name == "_" {
		return
	}
	qname := QName2(name, pkgpath)
	method, found := t.methodcache[qname]
	if found {
		index := method.Index
		if index < 0 { // marker for ambiguous method names
			count = -index
		} else {
			count = 1
		}
		return method, count
	}
	method, count = methodByName(t, qname, nil)
	if count == 0 {
		tovisit := anonymousFields(t, 0, nil)
		// breadth-first recursion
		for count == 0 && len(tovisit) != 0 {
			var next []StructField
			for _, f := range tovisit {
				et := unwrap(f.Type)
				emethod, ecount := methodByName(et, qname, f.Index)
				if count == 0 {
					if ecount > 0 {
						method = emethod
					} else {
						// no recursion if we found something
						next = append(next, anonymousFields(et, f.Offset, f.Index)...)
					}
				}
				count += ecount
			}
			tovisit = next
		}
	}
	if count > 0 {
		cacheMethodByName(t, qname, &method, count)
	}
	return method, count
}

func methodByName(t *xtype, qname QName, index []int) (method Method, count int) {
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		gmethod := t.method(i)
		if matchMethodByName(qname, gmethod) {
			if count == 0 {
				method = t.Method(i)
				method.FieldIndex = concat(index, method.FieldIndex) // make a copy of index
				// debugf("methodByName: %d-th method of <%v> matches: %#v", i, t.rtype, method)
			}
			count++
		}
	}
	return
}

// return true if gmethod name matches given name
func matchMethodByName(qname QName, gmethod *types.Func) bool {
	// always check the methods's package, not the type's package
	return qname == QNameGo(gmethod)
}

// add method to type's methodcache. used by Type.MethodByName after a successful lookup
func cacheMethodByName(t *xtype, qname QName, method *Method, count int) {
	if t.methodcache == nil {
		t.methodcache = make(map[QName]Method)
	}
	if count > 1 {
		method.Index = -count // marker for ambiguous method names
	}
	t.methodcache[qname] = *method
}

func invalidateCache(gtype types.Type, t interface{}) {
	if t, ok := t.(Type); ok {
		t := unwrap(t)
		t.fieldcache = nil
		t.methodcache = nil
	}
}

// clears all xtype.fieldcache and xtype.methodcache.
// invoked by NamedOf() when a type is redefined.
func (v *Universe) invalidateCache() {
	v.gmap.Iterate(invalidateCache)
}
