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

package interpreter

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
)

func (env *Env) EvalAst1(in Ast) r.Value {
	value, extraValues := env.EvalAst(in)
	if len(extraValues) > 1 {
		env.warnf("expression returned %d values, using only the first one: %v returned %v",
			len(extraValues), in.Interface(), extraValues)
	}
	return value
}

func (env *Env) EvalAst(in Ast) (r.Value, []r.Value) {
	switch in := in.(type) {
	case AstWithNode:
		if in != nil {
			return env.Eval(ToNode(in))
		}
	case AstWithSlice:
		if in != nil {
			var ret r.Value
			var rets []r.Value
			n := in.Size()
			for i := 0; i < n; i++ {
				ret, rets = env.Eval(ToNode(in.Get(i)))
			}
			return ret, rets
		}
	case nil:
		return None, nil
	default:
		return env.Errorf("EvalAst(): expecting <AstWithNode> or <AstWithSlice>, found: %v <%v>",
			in, r.TypeOf(in))
	}
	return env.Errorf("EvalAst(): expecting <AstWithNode> or <AstWithSlice>, found: nil")
}

func (env *Env) Eval(node ast.Node) (r.Value, []r.Value) {
	switch node := node.(type) {
	case ast.Decl:
		return env.evalDecl(node)
	case ast.Expr:
		return env.evalExpr(node)
	case ast.Stmt:
		return env.evalStatement(node)
	case *ast.File:
		return env.evalFile(node)
	default:
		return env.Errorf("unimplemented Eval for %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) Eval1(node ast.Node) r.Value {
	value, extraValues := env.Eval(node)
	if len(extraValues) > 1 {
		env.warnExtraValues(extraValues)
	}
	return value
}
