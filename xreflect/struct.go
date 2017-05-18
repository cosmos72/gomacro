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
 * struct.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/token"
	"go/types"
	"reflect"
)

// Field returns a struct type's i'th field.
// It panics if the type's Kind is not Struct.
// It panics if i is not in the range [0, NumField()).
func (t *xtype) Field(i int) StructField {
	if t.kind != reflect.Struct {
		errorf("Field of non-struct type %v", t)
	}
	gtype := t.gtype.Underlying().(*types.Struct)

	va := gtype.Field(i)
	rf := t.rtype.Field(i)

	return StructField{
		Name:      va.Name(),
		Pkg:       (*Package)(va.Pkg()),
		Type:      MakeType(va.Type(), rf.Type),
		Tag:       rf.Tag,
		Offset:    rf.Offset,
		Index:     rf.Index,
		Anonymous: va.Anonymous(),
	}
}

// NumField returns a struct type's field count.
// It panics if the type's Kind is not Struct.
func (t *xtype) NumField() int {
	if t.kind != reflect.Struct {
		errorf("NumField of non-struct type %v", t)
	}
	gtype := t.underlying().(*types.Struct)
	return gtype.NumFields()
}

// FieldByName returns the (possibly embedded) struct field with the given name
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

func concat(a, b []int) []int {
	na := len(a)
	c := make([]int, na+len(b))
	copy(c, a)
	copy(c[na:], b)
	return c
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

func isExportedName(name string) bool {
	if len(name) == 0 {
		return true
	}
	ch := name[0]
	return ch != '_' && (ch < 'a' || ch > 'z')
}

func (field *StructField) toReflectField(forceExported bool) reflect.StructField {
	var pkgpath string
	if pkg := field.Pkg; pkg != nil && !forceExported {
		pkgpath = pkg.Path()
	}
	name := field.Name
	if forceExported {
		name = toExportedFieldName(field.Name, field.Type.Name(), field.Anonymous)
	}
	return reflect.StructField{
		Name:      name,
		PkgPath:   pkgpath,
		Type:      field.Type.ReflectType(),
		Tag:       field.Tag,
		Offset:    field.Offset,
		Index:     field.Index,
		Anonymous: field.Anonymous,
	}
}

func toReflectFields(fields []StructField, forceExported bool) []reflect.StructField {
	rfields := make([]reflect.StructField, len(fields))
	for i := range fields {
		rfields[i] = fields[i].toReflectField(forceExported)
	}
	return rfields
}

func (field *StructField) toGoField() *types.Var {
	return types.NewField(token.NoPos, (*types.Package)(field.Pkg), field.Name, field.Type.GoType(), field.Anonymous)
}

func toGoFields(fields []StructField) []*types.Var {
	vars := make([]*types.Var, len(fields))
	for i := range fields {
		vars[i] = fields[i].toGoField()
	}
	return vars
}

func (field *StructField) toTag() string {
	return string(field.Tag)
}

func toTags(fields []StructField) []string {
	tags := make([]string, len(fields))
	for i := range fields {
		tags[i] = fields[i].toTag()
	}
	return tags
}

func toExportedFieldName(name, typename string, anonymous bool) string {
	if anonymous || len(name) == 0 {
		if len(name) == 0 {
			name = typename
		}
		return GensymEmbedded(name)
	}
	if ch := name[0]; ch >= 'a' && ch <= 'z' || ch == '_' {
		return GensymPrivate(name)
	}
	return name
}

func StructOf(fields []StructField) Type {
	vars := toGoFields(fields)
	tags := toTags(fields)
	rfields := toReflectFields(fields, true)
	return MakeType(
		types.NewStruct(vars, tags),
		reflect.StructOf(rfields),
	)
}
