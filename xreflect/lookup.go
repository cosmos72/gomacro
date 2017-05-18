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
	// debugf("field cache for %v <%v> = %v", unsafe.Pointer(t), t, t.fieldcache)
	field, found := t.fieldcache[pkgpath][name]
	if found {
		if field.Index == nil { // marker for ambiguous field names
			count = int(field.Offset) // reuse Offset as "number of ambiguous fields"
		} else {
			count = 1
		}
		return field, count
	}
	var tovisit []StructField

	field, count, tovisit = fieldByName(t, name, pkgpath, 0, nil)

	// breadth-first recursion
	for count == 0 && len(tovisit) != 0 {
		var next []StructField
		for _, f := range tovisit {
			efield, ecount, etovisit := fieldByName(&f.Type[0], name, pkgpath, f.Offset, f.Index)
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
		cacheFieldByName(t, name, pkgpath, &field, count)
	}
	return field, count
}

func fieldByName(t *xtype, name, pkgpath string, offset uintptr, index []int) (field StructField, count int, tovisit []StructField) {
	exported := isExportedName(name)
	gtype := t.gtype.Underlying().(*types.Struct)
	n := t.NumField()
	for i := 0; i < n; i++ {

		gfield := gtype.Field(i)
		if matchFieldByName(name, pkgpath, exported, gfield) {
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
func matchFieldByName(name, pkgpath string, exported bool, gfield *types.Var) bool {
	// always check the field's package, not the type's package
	if !exported && (gfield.Pkg() == nil || gfield.Pkg().Path() != pkgpath) {
		return false
	}
	if gfield.Name() == name {
		return true
	}
	if gfield.Anonymous() {
		switch gtype := gfield.Type().(type) {
		case *types.Basic:
			return gtype.Name() == name
		case *types.Named:
			return gtype.Obj().Name() == name
		}
	}
	return false
}

// add field to type's fieldcache. used by Type.FieldByName after a successful lookup
func cacheFieldByName(t *xtype, name, pkgpath string, field *StructField, count int) {
	if t.fieldcache == nil {
		t.fieldcache = make(map[string]map[string]StructField)
	}
	if t.fieldcache[pkgpath] == nil {
		t.fieldcache[pkgpath] = make(map[string]StructField)
	}
	if count > 1 {
		field.Index = nil             // marker for ambiguous field names
		field.Offset = uintptr(count) // reuse Offset as "number of ambiguous fields"
	}
	t.fieldcache[pkgpath][name] = *field
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
	method, found := t.methodcache[pkgpath][name]
	if found {
		if method.Index < 0 { // marker for ambiguous method names
			count = -method.Index
		} else {
			count = 1
		}
		return method, count
	}
	method, count = methodByName(t, name, pkgpath, nil)
	if count == 0 {
		tovisit := anonymousFields(t, 0, nil)
		// breadth-first recursion
		for count == 0 && len(tovisit) != 0 {
			var next []StructField
			for _, f := range tovisit {
				t = &f.Type[0]
				emethod, ecount := methodByName(t, name, pkgpath, f.Index)
				if count == 0 {
					if ecount > 0 {
						method = emethod
					} else {
						// no recursion if we found something
						next = append(next, anonymousFields(t, f.Offset, f.Index)...)
					}
				}
				count += ecount
			}
			tovisit = next
		}
	}
	if count != 0 {
		cacheMethodByName(t, name, pkgpath, &method, count)
	}
	return method, count
}

func methodByName(t *xtype, name, pkgpath string, index []int) (method Method, count int) {
	exported := isExportedName(name)
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		gmethod := t.method(i)
		if matchMethodByName(name, pkgpath, exported, gmethod) {
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
func matchMethodByName(name, pkgpath string, exported bool, gmethod *types.Func) bool {
	// always check the methods's package, not the type's package
	return gmethod.Name() == name && (exported || (gmethod.Pkg() != nil && gmethod.Pkg().Path() == pkgpath))
}

// add method to type's methodcache. used by Type.MethodByName after a successful lookup
func cacheMethodByName(t *xtype, name, pkgpath string, method *Method, count int) {
	if t.methodcache == nil {
		t.methodcache = make(map[string]map[string]Method)
	}
	if t.methodcache[pkgpath] == nil {
		t.methodcache[pkgpath] = make(map[string]Method)
	}
	if count > 1 {
		method.Index = -count // marker for ambiguous method names
	}
	t.methodcache[pkgpath][name] = *method
}
