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
 * unaryexpr.go
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

func (ir *Interpreter) evalExpr(node ast.Expr) (r.Value, error) {
	// fmt.Printf("debug: evalExpr() %#v\n", node)

	switch node := node.(type) {
	case *ast.BasicLit:
		return ir.evalLiteral(node)

	case *ast.BinaryExpr:
		return ir.evalBinaryExpr(node)

	case *ast.Ident:
		return ir.evalIdentifier(node)

	case *ast.ParenExpr:
		return ir.evalExpr(node.X)

	case *ast.UnaryExpr:
		return ir.evalUnaryExpr(node)

	case *ast.CallExpr, *ast.CompositeLit, *ast.FuncLit, *ast.IndexExpr, *ast.KeyValueExpr,
		*ast.SelectorExpr, *ast.SliceExpr, *ast.TypeAssertExpr:

		// TODO
		return Nil, errors.New(fmt.Sprintf("unimplemented expression %#v\n", node))

	default:
		return Nil, errors.New(fmt.Sprintf("unimplemented expression %#v\n", node))
	}
}
