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
	"go/ast"
	r "reflect"
)

func (env *Env) evalExprs(nodes []ast.Expr) []r.Value {
	switch n := len(nodes); n {
	case 0:
		return nil
	case 1:
		ret, _ := env.evalExpr(nodes[0])
		return []r.Value{ret}
	default:
		rets := make([]r.Value, n)
		for i, node := range nodes {
			ret, _ := env.evalExpr(node)
			rets[i] = ret
		}
		return rets
	}
}

func (env *Env) evalExpr(node ast.Expr) (r.Value, []r.Value) {
	// fmt.Printf("debug: evalExpr() %#v\n", node)

	switch node := node.(type) {
	case *ast.BasicLit:
		return env.evalLiteral(node)

	case *ast.BinaryExpr:
		return env.evalBinaryExpr(node)

	case *ast.CallExpr:
		return env.evalCall(node)

	case *ast.Ident:
		return env.evalIdentifier(node)

	case *ast.ParenExpr:
		return env.evalExpr(node.X)

	case *ast.UnaryExpr:
		return env.evalUnaryExpr(node)

	case *ast.CompositeLit, *ast.FuncLit, *ast.IndexExpr, *ast.KeyValueExpr,
		*ast.SelectorExpr, *ast.SliceExpr, *ast.TypeAssertExpr:

		// TODO
		return Errorf("unimplemented expression %#v", node)

	default:
		return Errorf("unimplemented expression %#v", node)
	}
}
