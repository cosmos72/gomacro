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

	. "github.com/cosmos72/gomacro/base"
)

type stubDebugger struct{}

func (s stubDebugger) Breakpoint(ir *Interp, env *Env) DebugOp {
	return SigDebugContinue
}

func (s stubDebugger) At(ir *Interp, env *Env) DebugOp {
	return SigDebugContinue
}

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
		ir := Interp{c, env}
		op := ir.debug(true)
		env.IP++
		stmt := env.Code[env.IP]
		if op != SigNone {
			g := env.ThreadGlobals
			stmt = g.Interrupt
			if g.Options&OptDebugDebugger != 0 {
				g.Debugf("after breakpoint: single-stepping with stmt = %p, env = %p, IP = %v, execFlags = %v, signals = %#v", stmt, env, env.IP, g.ExecFlags, g.Signals)
			}
		}
		return stmt, env
	}
}

func singleStep(env *Env) (Stmt, *Env) {
	stmt := env.Code[env.IP]
	g := env.ThreadGlobals
	if g.Signals.Debug == SigNone {
		return stmt, env // resume normal execution
	}

	if env.CallDepth < g.DebugDepth {
		if g.Options&OptDebugDebugger != 0 {
			g.Debugf("single-stepping: stmt = %p, env = %p, IP = %v, env.CallDepth = %d, g.DebugDepth = %d", stmt, env, env.IP, env.CallDepth, g.DebugDepth)
		}
		c := env.DebugComp
		if c != nil {
			ir := Interp{c, env}
			op := ir.debug(false) // not a breakpoint
			if op != SigNone {
				g := env.ThreadGlobals
				g.Signals.Debug = op
			}
		}
	}

	// single step
	stmt, env = stmt(env)
	if g.Signals.Debug != SigNone {
		stmt = g.Interrupt
	}
	return stmt, env
}

func (ir *Interp) debug(breakpoint bool) DebugOp {
	g := ir.env.ThreadGlobals
	if g.Debugger == nil {
		ir.Comp.Warnf("// breakpoint: no debugger set with Interp.SetDebugger(), resuming execution (warned only once)")
		g.Debugger = stubDebugger{}
	}
	var op DebugOp
	if breakpoint {
		op = g.Debugger.Breakpoint(ir, ir.env)
	} else {
		op = g.Debugger.At(ir, ir.env)
	}
	if g.Options&OptDebugDebugger != 0 {
		g.Debugf("Debugger returned op = %v", op)
	}
	return ir.env.applyDebugSignal(op)
}

func (env *Env) applyDebugSignal(op DebugOp) DebugOp {
	g := env.ThreadGlobals
	saveOp := op
	saveDepth := g.DebugDepth
	switch op {
	case SigDebugFinish:
		g.DebugDepth = env.CallDepth
	case SigDebugNext:
		g.DebugDepth = env.CallDepth + 1
	case SigDebugStep:
		g.DebugDepth = MaxInt
	case SigDebugRepl:
		break
	default:
		op = SigNone
		g.DebugDepth = 0
	}
	// prevent further changes to g.DebugDepth
	if g.Options&OptDebugDebugger != 0 {
		if op == saveOp {
			g.Debugf("applyDebugSignal: op = %v, updated g.DebugDepth from %v to %v", op, saveDepth, g.DebugDepth)
		} else {
			g.Debugf("applyDebugSignal: op = %v, replaced with %v and updated g.DebugDepth from %v to %v", saveOp, op, saveDepth, g.DebugDepth)
		}
	}
	if op != SigNone {
		op = SigDebugRepl
	}
	g.ExecFlags.SetDebug(op != SigNone)
	g.Signals.Debug = op
	return op
}
