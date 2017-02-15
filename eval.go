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
 * eval.go
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

func (ir *Interpreter) Eval(node ast.Node) (r.Value, error) {
	switch node := node.(type) {

	case *ast.BasicLit:
		return ir.evalLiteral(node)

	case *ast.Ident:
		return ir.evalIdentifier(node)

	case *ast.UnaryExpr:
		return ir.evalUnaryExpr(node)

	case *ast.BinaryExpr:
		return ir.evalBinaryExpr(node)

	case *ast.File:
		return ir.evalFile(node)

	default:
		return Nil, errors.New(fmt.Sprintf("unsupported expression: %#v\n", node))
	}
}
