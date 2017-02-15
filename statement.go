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

func (ir *Interpreter) evalBlock(block *ast.BlockStmt) (r.Value, error) {
	ir.PushEnv()
	defer ir.PopEnv()

	return ir.evalStatements(block.List)
}

func (ir *Interpreter) evalStatements(list []ast.Stmt) (r.Value, error) {
	var ret r.Value
	var err error

	for _, stmt := range list {
		ret, err = ir.evalStatement(stmt)
		if err != nil {
			return Nil, err
		}
	}
	return ret, nil
}

func (ir *Interpreter) evalStatement(node ast.Stmt) (r.Value, error) {
	switch node := node.(type) {
	case *ast.AssignStmt:
		return ir.evalAssignments(node)
	case *ast.BlockStmt:
		return ir.evalBlock(node)
	case *ast.DeclStmt:
		return ir.evalDecl(node.Decl)
	case *ast.ExprStmt:
		return ir.evalExpr(node.X)
	case *ast.ForStmt:
		return ir.evalFor(node)
	case *ast.IfStmt:
		return ir.evalIf(node)
	case *ast.EmptyStmt:
		return Nil, nil
	case *ast.BranchStmt, *ast.CaseClause, *ast.CommClause, *ast.DeferStmt,
		*ast.GoStmt, *ast.IncDecStmt, *ast.LabeledStmt, *ast.RangeStmt, *ast.ReturnStmt,
		*ast.SelectStmt, *ast.SendStmt, *ast.SwitchStmt, *ast.TypeSwitchStmt:
		// TODO
		return Nil, errors.New(fmt.Sprintf("unimplemented statement: %#v", node))
	default:
		return Nil, errors.New(fmt.Sprintf("unimplemented statement: %#v", node))
	}
}
