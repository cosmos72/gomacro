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
 * debug.go
 *
 *  Created on Apr 20, 2018
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
)

// debugger operation
type DebugOp uint

const (
	DebugCont DebugOp = iota
	DebugFinish
	DebugNext
	DebugStep
	DebugRepl
)

// return true if statement is either "break" or _ = "break"
func isBreakpoint(stmt ast.Stmt) bool {
	switch node := stmt.(type) {
	case *ast.ExprStmt:
		return isBreakLiteral(node.X)
	case *ast.AssignStmt:
		if node.Tok == token.ASSIGN && len(node.Lhs) == 1 && len(node.Rhs) == 1 {
			return isUnderscore(node.Lhs[0]) && isBreakLiteral(node.Rhs[0])
		}
	}
	return false
}

func isUnderscore(node ast.Expr) bool {
	switch node := node.(type) {
	case *ast.Ident:
		return node.Name == "_"
	}
	return false
}

func isBreakLiteral(node ast.Expr) bool {
	switch node := node.(type) {
	case *ast.BasicLit:
		return node.Kind == token.STRING && node.Value == `"break"`
	}
	return false
}

func (c *Comp) breakpoint() Stmt {
	return func(env *Env) (Stmt, *Env) {
		// create an inner Comp to preserve existing Binds
		// create an inner Env to preserve compiled Code and IP
		ir := Interp{NewComp(c, nil), NewEnv(env, 0, 0)}
		var stmt Stmt
		switch ir.debug() {
		default:
			env.IP++
			stmt = env.Code[env.IP]
		}
		return stmt, env
	}
}

func (ir *Interp) debug() DebugOp {
	if Debugger == nil {
		ir.Comp.Warnf("// breakpoint: no debugger installed, resuming execution (warned only once)")
		Debugger = stubDebugger
	}
	return Debugger(ir, ir.env)
}

var Debugger func(ir *Interp, env *Env) DebugOp

func stubDebugger(ir *Interp, env *Env) DebugOp {
	return DebugCont
}
