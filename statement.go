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

func (ir *Interpreter) evalFile(node *ast.File) (r.Value, error) {
	ir.Packagename = node.Name.Name

	// eval node.Imports
	var ret r.Value
	var err error

	for _, stmt := range node.Decls {
		ret, err = ir.evalStatement(stmt)
		if err != nil {
			return Nil, err
		}
	}
	return ret, nil
}

func (ir *Interpreter) evalBlock(block []ast.Stmt) (r.Value, error) {
	ir.PushEnv()
	defer ir.PopEnv()

	var ret r.Value
	var err error

	for _, stmt := range block {
		ret, err = ir.evalStatement(stmt)
		if err != nil {
			return Nil, err
		}
	}
	return ret, nil
}

func (ir *Interpreter) evalStatement(node ast.Node) (r.Value, error) {
	switch node := node.(type) {
	case *ast.GenDecl:
		return ir.evalDecl(node)
	case *ast.FuncDecl:
		return ir.evalDeclFunc(node)
	case *ast.AssignStmt:
		return ir.evalAssignments(node)
	default:
		return Nil, errors.New(fmt.Sprintf("unsupported statement: %#v", node))
	}
}
