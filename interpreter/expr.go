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

package interpreter

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
		for i := range nodes {
			rets[i] = env.evalExpr1(nodes[i])
		}
		return rets
	}
}

func (env *Env) evalExpr1(node ast.Expr) r.Value {
	value, extraValues := env.evalExpr(node)
	if len(extraValues) > 1 {
		env.Warnf("function returned %d values, using only the first one: %v returned %v",
			len(extraValues), node, extraValues)
	}
	return value
}

func (env *Env) evalExpr(in ast.Expr) (r.Value, []r.Value) {
	for {
		// env.Debugf("evalExpr() %v", node)
		switch node := in.(type) {
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
			in = node.X
			continue

		case *ast.UnaryExpr:
			return env.evalUnaryExpr(node)

		case *ast.SelectorExpr:
			return env.evalSelectorExpr(node)

		case *ast.TypeAssertExpr:
			return env.evalTypeAssertExpr(node)

		case *ast.KeyValueExpr, *ast.SliceExpr:
			// TODO
		}
		return env.Errorf("unimplemented Eval() for: %v <%v>", in, r.TypeOf(in))
	}
}

func (env *Env) evalIndexExpr(node *ast.IndexExpr) (r.Value, []r.Value) {
	index := env.evalExpr1(node.Index)
	obj := env.evalExpr1(node.X)

	switch obj.Kind() {
	case r.Map:
		return obj.MapIndex(index), nil
	case r.Array, r.Slice, r.String:
		i, ok := env.toInt(index)
		if !ok {
			return env.Errorf("invalid index, expecting an int: %v <%v>", index, index.Type())
		}
		return obj.Index(int(i)), nil
	default:
		return env.Errorf("unsupported index operation: %v [ %v ]. not an array, map, slice or string: %v <%v>", node.X, index, obj, obj.Type())
	}
}

func (env *Env) evalSelectorExpr(node *ast.SelectorExpr) (r.Value, []r.Value) {
	obj := env.evalExpr1(node.X)
	name := node.Sel.Name
	if obj.Kind() == r.Ptr {
		obj = obj.Elem()
	}
	switch obj.Kind() {
	case r.Struct:
		val := obj.FieldByName(name)
		if val == Nil {
			val = obj.MethodByName(name)
		}
		if val == Nil {
			return env.Errorf("%v <%v> has no field or method %s", obj.Interface(), obj.Type(), name)
		}
		return val, nil
	default:
		return env.Errorf("not a struct: %v <%v> has no field or method %s", obj.Interface(), obj.Type(), name)
	}
}

func (env *Env) evalTypeAssertExpr(node *ast.TypeAssertExpr) (r.Value, []r.Value) {
	val := env.evalExpr1(node.X)
	fval := val.Interface()
	t1 := r.TypeOf(fval) // extract the actual runtime type of fval

	t2 := env.evalType(node.Type)
	if t1.AssignableTo(t2) {
		return r.ValueOf(fval).Convert(t2), nil
	}
	return env.Errorf("type assertion failed: %v <%v> is not a <%v>", fval, t1, t2)
}
