/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * eval.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (env *Env) Eval(src interface{}) (r.Value, []r.Value) {
	return env.EvalAst(env.Parse(src))
}

func (env *Env) Eval1(src interface{}) r.Value {
	return env.EvalAst1(env.Parse(src))
}

func (env *Env) EvalAst1(in ast2.Ast) r.Value {
	value, extraValues := env.EvalAst(in)
	if len(extraValues) > 1 {
		env.WarnExtraValues(extraValues)
	}
	return value
}

func (env *Env) EvalAst(in ast2.Ast) (r.Value, []r.Value) {
	switch in := in.(type) {
	case ast2.AstWithNode:
		if in != nil {
			return env.EvalNode(ast2.ToNode(in))
		}
	case ast2.AstWithSlice:
		if in != nil {
			var ret r.Value
			var rets []r.Value
			n := in.Size()
			for i := 0; i < n; i++ {
				ret, rets = env.EvalNode(ast2.ToNode(in.Get(i)))
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

func (env *Env) EvalNode(node ast.Node) (r.Value, []r.Value) {
	switch node := node.(type) {
	case ast.Decl:
		env.evalDecl(node)
	case ast.Expr:
		// Go expressions *DO* return values
		return env.evalExpr(node)
	case ast.Stmt:
		env.evalStatement(node)
	case *ast.File:
		env.evalFile(node)
	default:
		return env.Errorf("unimplemented Eval for %v <%v>", node, r.TypeOf(node))
	}
	// Go declarations, statements and files do not return values
	return None, nil
}

func (env *Env) EvalNode1(node ast.Node) r.Value {
	value, extraValues := env.EvalNode(node)
	if len(extraValues) > 1 {
		env.WarnExtraValues(extraValues)
	}
	return value
}

// macroexpand + collect + eval.
// if opt&CmdOptFast != 0 calls env.fastEval(), otherwise calls env.classicEval()
func (env *Env) evalAst(form ast2.Ast, opt CmdOpt) ([]r.Value, []xr.Type) {
	var values []r.Value
	var types []xr.Type
	if form != nil {
		if opt&CmdOptFast != 0 {
			values, types = PackValuesAndTypes(env.fastEval(form))
		} else {
			values = PackValues(env.classicEval(form))
		}
	}
	return values, types
}

// macroexpand + collect + eval
func (env *Env) classicEval(form ast2.Ast) (r.Value, []r.Value) {
	// macroexpansion phase.
	form, _ = env.MacroExpandAstCodewalk(form)

	if env.Options&OptShowMacroExpand != 0 {
		env.Debugf("after macroexpansion: %v", form.Interface())
	}

	// collect phase
	if env.Options&(OptCollectDeclarations|OptCollectStatements) != 0 {
		env.CollectAst(form)
	}

	// eval phase
	if env.Options&OptMacroExpandOnly != 0 {
		return r.ValueOf(form.Interface()), nil
	} else {
		return env.EvalAst(form)
	}
}
