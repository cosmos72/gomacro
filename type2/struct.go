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

package type2

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
		Type:      maketype(va.Type(), rf.Type),
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
	return maketype(
		types.NewStruct(vars, tags),
		reflect.StructOf(rfields),
	)
}
