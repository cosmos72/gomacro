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
 * controlflow.go
 *
 *  Created on: Feb 15, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"errors"
	"fmt"
	"go/ast"

	r "reflect"
)

func (ir *Interpreter) evalIf(node *ast.IfStmt) (r.Value, error) {
	if node.Init != nil {
		ir.PushEnv()
		defer ir.PopEnv()

		_, err := ir.evalStatement(node.Init)
		if err != nil {
			return Nil, err
		}
	}
	cond, err := ir.Eval(node.Cond)
	if err != nil {
		return Nil, err
	}
	if cond.Kind() != r.Bool {
		cf := cond.Interface()
		return Nil, errors.New(fmt.Sprintf("if: invalid condition type <%T> %#v, expecting <bool>", cf, cf))
	}
	if cond.Bool() {
		return ir.evalBlock(node.Body)
	} else if node.Else != nil {
		return ir.evalStatement(node.Else)
	} else {
		return Nil, nil
	}
}

func (ir *Interpreter) evalFor(node *ast.ForStmt) (r.Value, error) {
	// fmt.Printf("debug: evalFor() init = %#v, cond = %#v, post = %#v, body = %#v\n", node.Init, node.Cond, node.Post, node.Body)

	if node.Init != nil {
		ir.PushEnv()
		defer ir.PopEnv()

		_, err := ir.evalStatement(node.Init)
		if err != nil {
			return Nil, err
		}
	}
	for {
		if node.Cond != nil {
			cond, err := ir.evalExpr(node.Cond)
			if err != nil {
				return Nil, err
			}
			if cond.Kind() != r.Bool {
				cf := cond.Interface()
				return Nil, errors.New(fmt.Sprintf("for: invalid condition type <%T> %#v, expecting <bool>", cf, cf))
			}
			if !cond.Bool() {
				break
			}
		}
		_, err := ir.evalBlock(node.Body)
		if err != nil {
			return Nil, err
		}
		if node.Post != nil {
			_, err := ir.evalStatement(node.Post)
			if err != nil {
				return Nil, err
			}
		}
	}
	return Nil, nil
}
