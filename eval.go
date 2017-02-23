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
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	r "reflect"
)

func (env *Env) EvalList(nodes []ast.Node) (r.Value, []r.Value) {
	var ret r.Value
	var rets []r.Value
	for _, node := range nodes {
		ret, rets = env.Eval(node)
	}
	return ret, rets
}

func (env *Env) Eval(node ast.Node) (r.Value, []r.Value) {
	switch node := node.(type) {
	case ast.Expr:
		temp := node
		return env.evalExpr(&temp)
	case ast.Decl:
		return env.evalDecl(node)
	case ast.Stmt:
		temp := node
		return env.evalStatement(&temp)
	case *ast.File:
		return env.evalFile(node)
	}
	return env.Errorf("unimplemented Eval for %v <%v>", node, r.TypeOf(node))
}

func (env *Env) Eval1(node ast.Node) r.Value {
	value, extraValues := env.Eval(node)
	if len(extraValues) > 1 {
		env.Warnf("function returned %d values, only the first one will be used: %v", len(extraValues), node)
	}
	return value
}
