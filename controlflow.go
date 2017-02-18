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
	"go/ast"
	"go/token"
	r "reflect"
)

type Break struct {
	Label string
}

type Continue struct {
	Label string
}

type Return struct {
	Values []r.Value
}

func (_ Break) Error() string {
	return "break outside for or switch"
}

func (_ Continue) Error() string {
	return "continue outside for"
}

func (_ Return) Error() string {
	return "return outside function"
}

func (env *Env) evalBranch(node *ast.BranchStmt) (r.Value, []r.Value) {
	var label string
	if node.Label != nil {
		label = node.Label.Name
	}
	switch node.Tok {
	case token.BREAK:
		panic(Break{label})
	case token.CONTINUE:
		panic(Continue{label})
	case token.GOTO:
		return Errorf("unimplemented: goto")
	case token.FALLTHROUGH:
		return Errorf("unimplemented: fallthrough")
	default:
		return Errorf("unimplemented branch: %#v", node)
	}
}

func (env *Env) evalReturn(node *ast.ReturnStmt) (r.Value, []r.Value) {
	rets := env.evalExprs(node.Results)
	panic(Return{rets})
}

func (env *Env) evalIf(node *ast.IfStmt) (r.Value, []r.Value) {
	if node.Init != nil {
		env = NewEnv(env)
		_, _ = env.evalStatement(node.Init)
	}
	cond, _ := env.Eval(node.Cond)
	if cond.Kind() != r.Bool {
		cf := cond.Interface()
		return Errorf("if: invalid condition type <%T> %#v, expecting <bool>", cf, cf)
	}
	if cond.Bool() {
		return env.evalBlock(node.Body)
	} else if node.Else != nil {
		return env.evalStatement(node.Else)
	} else {
		return Nil, nil
	}
}

func (env *Env) evalFor(node *ast.ForStmt) (r.Value, []r.Value) {
	// fmt.Printf("debug: evalFor() init = %#v, cond = %#v, post = %#v, body = %#v\n", node.Init, node.Cond, node.Post, node.Body)

	if node.Init != nil {
		env = NewEnv(env)
		env.evalStatement(node.Init)
	}
	for {
		if node.Cond != nil {
			cond, _ := env.evalExpr(node.Cond)
			if cond.Kind() != r.Bool {
				cf := cond.Interface()
				return Errorf("for: invalid condition type <%T> %#v, expecting <bool>", cf, cf)
			}
			if !cond.Bool() {
				break
			}
		}
		if !env.evalForBodyOnce(node.Body) {
			break
		}
		if node.Post != nil {
			env.evalStatement(node.Post)
		}
	}
	return Nil, nil
}

func (env *Env) evalForBodyOnce(node *ast.BlockStmt) (cont bool) {
	defer func() {
		if rec := recover(); rec != nil {
			switch rec := rec.(type) {
			case Break:
				cont = false
			case Continue:
				cont = true
			default:
				panic(rec)
			}
		}
	}()
	env.evalBlock(node)
	return true
}
