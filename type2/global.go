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

package type2

import (
	"go/types"
	"reflect"
)

type timpl struct {
	kind  reflect.Kind
	gtype types.Type
	rtype reflect.Type
	// underlying *timpl // needed? computing it for elements
	//   of arrays, maps, slices, structs is painful.
	//   for function/method arguments is even worse.
}

type Package struct {
	impl *types.Package
}

type Type struct {
	*timpl
}

type Method struct {
	Name  string
	Pkg   Package
	Type  Type          // method type
	Func  reflect.Value // func with receiver as first argument
	Index int           // index for Type.Method
}

type StructField struct {
	// Name is the field name.
	Name string
	// Pkg is the package that qualifies a lower case (unexported)
	// field name. It may be nil for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	Pkg       Package
	Type      Type              // field type
	Tag       reflect.StructTag // field tag string
	Offset    uintptr           // offset within struct, in bytes
	Index     []int             // index sequence for Type.FieldByIndex
	Anonymous bool              // is an embedded field
}
