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
 * global.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"reflect"
)

type Package types.Package

// InterfaceHeader is the internal header of interpreted interfaces
type InterfaceHeader struct {
	// val and typ must be private! otherwise interpreted code may mess with them and break type safety
	val reflect.Value
	typ Type
}

type Method struct {
	Name       string
	Pkg        *Package
	Type       Type             // method type
	Funs       *[]reflect.Value // (*Funs)[Index] is the method, with receiver as first argument
	Index      int              // index for Type.Method
	FieldIndex []int            // embedded fields index sequence for reflect.Type.FieldByIndex or reflect.Value.FieldByIndex
}

type StructField struct {
	// Name is the field name.
	Name string
	// Pkg is the package that qualifies a lower case (unexported)
	// field name. It may be nil for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	Pkg       *Package
	Type      Type              // field type
	Tag       reflect.StructTag // field tag string
	Offset    uintptr           // offset within struct, in bytes
	Index     []int             // index sequence for reflect.Type.FieldByIndex or reflect.Value.FieldByIndex
	Anonymous bool              // is an embedded field
}

type xtype struct {
	kind         reflect.Kind
	gtype        types.Type
	rtype        reflect.Type
	universe     *Universe
	methodvalues []reflect.Value
	fieldcache   map[string]map[string]StructField // index by pkgpath, then by name
	methodcache  map[string]map[string]Method      // index by pkgpath, then by name
}
