/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * interface.go
 *
 *  Created on: Mar 29, 2017
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

// "\u0080" is Unicode codepoint: Padding Character.
// reflect.StructOf() allows it as field name, while go/scanner forbids it in Go source code
const nameOfInterfaceObject = "\u0080"

func (c *Comp) TypeInterface(node *ast.InterfaceType) r.Type {
	if node.Methods == nil || len(node.Methods.List) == 0 {
		return TypeOfInterface
	}
	types, names := c.TypeFields(node.Methods)

	types = append([]r.Type{TypeOfInterface}, types...)
	names = append([]string{nameOfInterfaceObject}, names...)

	fields := c.makeStructFields(c.FileComp().Path, names, types)
	return r.StructOf(fields)
}

func isInterfaceType(t r.Type) bool {
	if t.Kind() == r.Struct && t.NumField() > 0 {
		field := t.Field(0)
		return field.Name == nameOfInterfaceObject && field.Type == TypeOfInterface
	}
	return false
}
