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
 * identifier.go
 *
 *  Created on: Feb 13, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"errors"
	"fmt"
	"go/ast"
	r "reflect"
)

func (ir *Interpreter) evalIdentifier(expr *ast.Ident) (r.Value, error) {
	name := expr.Name

	switch name {
	case "false":
		return r.ValueOf(false), nil
	case "true":
		return r.ValueOf(true), nil
	case "iota":
		pos := ir.Fileset.Position(expr.NamePos)
		return r.ValueOf(pos.Line - ir.iotaOffset), nil
	default:
		for env := ir.Env; env != nil; env = env.Outer {
			// fmt.Printf("debug: evalIdentifier() looking up %#v in %#v\n", name, env.Binds)
			bind, exists := env.Binds[name]
			if exists {
				return bind, nil
			}
		}
		return Nil, errors.New(fmt.Sprintf("undefined identifier: %s", name))
	}
}
