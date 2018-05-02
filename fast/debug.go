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
	return DebugOpContinue
}

func (s stubDebugger) At(ir *Interp, env *Env) DebugOp {
	return DebugOpContinue
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
		sig := ir.debug(true)
		env.IP++
		stmt := env.Code[env.IP]
		if sig != SigNone {
			g := env.Run
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
	run := env.Run
	if run.Signals.Debug == SigNone {
		return stmt, env // resume normal execution
	}

	if env.CallDepth < run.DebugDepth {
		if run.Options&OptDebugDebugger != 0 {
			run.Debugf("single-stepping: stmt = %p, env = %p, IP = %v, env.CallDepth = %d, g.DebugDepth = %d", stmt, env, env.IP, env.CallDepth, run.DebugDepth)
		}
		c := env.DebugComp
		if c != nil {
			ir := Interp{c, env}
			sig := ir.debug(false) // not a breakpoint
			if sig != SigNone {
				run := env.Run
				run.Signals.Debug = sig
			}
		}
	}

	// single step
	stmt, env = stmt(env)
	if run.Signals.Debug != SigNone {
		stmt = run.Interrupt
	}
	return stmt, env
}

func (ir *Interp) debug(breakpoint bool) Signal {
	g := ir.env.Run
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
	return ir.env.applyDebugOp(op)
}

func (env *Env) applyDebugOp(op DebugOp) Signal {
	run := env.Run
	saveOp := op
	var sig Signal
	if op.Depth > 0 {
		sig = SigDebug
	} else {
		sig = SigNone
		op.Depth = 0
	}
	if run.Options&OptDebugDebugger != 0 {
		if op == saveOp {
			run.Debugf("applyDebugSignal: op = %v, updated run.DebugDepth from %v to %v", op, run.DebugDepth, op.Depth)
		} else {
			run.Debugf("applyDebugSignal: op = %v, replaced with %v and updated run.DebugDepth from %v to %v", saveOp, op, run.DebugDepth, op.Depth)
		}
	}
	run.DebugDepth = op.Depth
	run.ExecFlags.SetDebug(sig != SigNone)
	run.Signals.Debug = sig
	return sig
}
