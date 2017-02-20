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
 * expr.go
 *
 *  Created on: Feb 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	r "reflect"
)

func (env *Env) evalExprsMultipleValues(nodes []ast.Expr, expectedValuesN int) []r.Value {
	n := len(nodes)
	var values []r.Value
	if n != expectedValuesN {
		if n != 1 {
			return env.PackErrorf("value count mismatch: cannot assign %d values to %d places: %v",
				n, expectedValuesN, nodes)
		}
		node := nodes[0]
		// collect multiple values
		values = packValues(env.Eval(node))
		n = len(values)
		if n < expectedValuesN {
			return env.PackErrorf("value count mismatch: function returned %d values, cannot assign them to %d places: %v returned %v",
				n, expectedValuesN, node, values)
		} else if n > expectedValuesN {
			env.Warnf("function returned %d values, using only %d of them: %v returned %v",
				n, expectedValuesN, node, values)
		}
	} else {
		values = env.evalExprs(nodes)
	}
	return values
}

func (env *Env) evalExprs(nodes []ast.Expr) []r.Value {
	switch n := len(nodes); n {
	case 0:
		return nil
	case 1:
		ret := env.evalExpr1(nodes[0])
		return []r.Value{ret}
	default:
		rets := make([]r.Value, n)
		for i, node := range nodes {
			rets[i] = env.evalExpr1(node)
		}
		return rets
	}
}

func (env *Env) evalExpr1(expr ast.Expr) r.Value {
	value, extraValues := env.evalExpr(expr)
	if len(extraValues) > 1 {
		env.Warnf("function returned %d values, using only the first one: %v returned %v",
			len(extraValues), expr, extraValues)
	}
	return value
}

func (env *Env) evalExpr(expr ast.Expr) (r.Value, []r.Value) {
	// Debugf("evalExpr() %#v", node)

	for {
		switch node := expr.(type) {
		case *ast.BasicLit:
			return env.evalLiteral(node)

		case *ast.BinaryExpr:
			return env.evalBinaryExpr(node)

		case *ast.CallExpr:
			return env.evalCall(node)

		case *ast.CompositeLit:
			return env.evalCompositeLiteral(node)

		case *ast.FuncLit:
			return env.evalFunctionLiteral(node)

		case *ast.Ident:
			return env.evalIdentifier(node)

		case *ast.IndexExpr:
			return env.evalIndexExpr(node)

		case *ast.ParenExpr:
			expr = node.X
			continue

		case *ast.UnaryExpr:
			return env.evalUnaryExpr(node)

		case *ast.KeyValueExpr, *ast.SelectorExpr, *ast.SliceExpr, *ast.TypeAssertExpr:

			// TODO
			return env.Errorf("unimplemented expression %#v", node)

		default:
			return env.Errorf("unimplemented expression %#v", node)
		}
	}
}

func (env *Env) evalIndexExpr(node *ast.IndexExpr) (r.Value, []r.Value) {
	index := env.evalExpr1(node.Index)
	container := env.evalExpr1(node.X)

	switch container.Kind() {
	case r.Map:
		return container.MapIndex(index), nil
	case r.Array, r.Slice, r.String:
		i, ok := env.toInt(index)
		if !ok {
			return env.Errorf("invalid index, expecting an int: %v <%v>", index, index.Type())
		}
		return container.Index(int(i)), nil
	default:
		return env.Errorf("unsupported index operation: %v [ %v ]. not an array, map, slice or string: %v <%v>", node.X, index, container, container.Type())
	}
}
