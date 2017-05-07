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
func (t *timpl) Field(i int) StructField {
	if t.kind != reflect.Struct {
		errorf("Field of non-struct type %v", t)
	}
	gtype := t.gtype.Underlying().(*types.Struct)

	va := gtype.Field(i)
	rf := t.rtype.Field(i)
	return StructField{
		Name:      va.Name(),
		Pkg:       makepackage(va.Pkg()),
		Type:      maketype(va.Type(), rf.Type),
		Tag:       rf.Tag,
		Offset:    rf.Offset,
		Index:     rf.Index,
		Anonymous: va.Anonymous(),
	}
}

func (t *timpl) NumField() int {
	if t.kind != reflect.Struct {
		errorf("NumField of non-struct type %v", t)
	}
	gtype := t.underlying().(*types.Struct)
	return gtype.NumFields()
}

func (field *StructField) toReflectField() reflect.StructField {
	var pkgpath string
	if pkg := field.Pkg; pkg.impl != nil {
		pkgpath = pkg.Path()
	}
	return reflect.StructField{
		Name:      field.Name,
		PkgPath:   pkgpath,
		Type:      field.Type.ReflectType(),
		Tag:       field.Tag,
		Offset:    field.Offset,
		Index:     field.Index,
		Anonymous: field.Anonymous,
	}
}

func StructOf(fields []StructField) Type {
	n := len(fields)
	rfields := make([]reflect.StructField, n)
	vars := make([]*types.Var, n)
	tags := make([]string, n)
	for i, field := range fields {
		rfields[i] = field.toReflectField()
		vars[i] = types.NewField(token.NoPos, field.Pkg.impl, field.Name, field.Type.gtype, field.Anonymous)
		tags[i] = string(field.Tag)
	}
	return maketype(
		types.NewStruct(vars, tags),
		reflect.StructOf(rfields),
	)
}
