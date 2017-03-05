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
 * statement.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
	r "reflect"
)

func (env *Env) evalBlock(block *ast.BlockStmt) (r.Value, []r.Value) {
	env = NewEnv(env, "{}")

	return env.evalStatements(block.List)
}

func (env *Env) evalStatements(list []ast.Stmt) (r.Value, []r.Value) {
	ret := None
	var rets []r.Value

	for i := range list {
		ret, rets = env.evalStatement(list[i])
	}
	return ret, rets
}

func (env *Env) evalStatement(node ast.Stmt) (r.Value, []r.Value) {
	switch node := node.(type) {
	case *ast.AssignStmt:
		return env.evalAssignments(node)
	case *ast.BlockStmt:
		return env.evalBlock(node)
	case *ast.BranchStmt:
		return env.evalBranch(node)
	case *ast.DeclStmt:
		return env.evalDecl(node.Decl)
	case *ast.ExprStmt:
		return env.evalExpr(node.X)
	case *ast.ForStmt:
		return env.evalFor(node)
	case *ast.IfStmt:
		return env.evalIf(node)
	case *ast.IncDecStmt:
		return env.evalIncDec(node)
	case *ast.EmptyStmt:
		return None, nil
	case *ast.ReturnStmt:
		return env.evalReturn(node)

	case *ast.CaseClause, *ast.CommClause, *ast.DeferStmt,
		*ast.GoStmt, *ast.LabeledStmt, *ast.RangeStmt,
		*ast.SelectStmt, *ast.SendStmt, *ast.SwitchStmt, *ast.TypeSwitchStmt:
		// TODO
		return env.Errorf("unimplemented statement: %v <%v>", node, r.TypeOf(node))
	default:
		return env.Errorf("unimplemented statement: %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) evalIncDec(node *ast.IncDecStmt) (r.Value, []r.Value) {
	var op token.Token
	switch node.Tok {
	case token.INC:
		op = token.ADD_ASSIGN
	case token.DEC:
		op = token.SUB_ASSIGN
	default:
		return env.Errorf("unsupported *ast.IncDecStmt operation, expecting ++ or -- : %v <%v>", node, r.TypeOf(node))
	}
	place := env.evalPlace(node.X)
	return env.assignPlace(place, op, One), nil
}
