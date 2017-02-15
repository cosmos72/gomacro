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
 * function.go
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

func (ir *Interpreter) evalDeclFunc(node *ast.FuncDecl) (r.Value, error) {
	name := node.Name.Name
	if name == TemporaryFunctionName {
		// do *NOT* use ir.evalBlock(), because it would create all bindings
		// in its block scope -> they are lost after ir.evalBlock() returns
		return ir.evalStatements(node.Body.List)
	}
	// TODO
	return Nil, errors.New(fmt.Sprintf("unimplemented function declaration: %#v", node))
}
