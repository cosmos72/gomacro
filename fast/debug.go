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
	r "reflect"
)

type DebugOpt uint

const (
	DebugCont DebugOpt = iota
	DebugNext
	DebugStep
	DebugFinish
)

func isBreakpoint(expr *Expr) bool {
	return expr.Const() && expr.UntypedKind() == r.String && expr.Value.(UntypedLit).Val.ExactString() == `"break"`
}

func makeStmtBreakpoint(c *Comp) Stmt {

	return func(env *Env) (Stmt, *Env) {
		// create an inner Comp to preserve existing binds
		// create an inner Env to preserve compiled code and IP
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

func (ir *Interp) debug() DebugOpt {
	g := ir.Comp.Globals

	env := ir.env.Outer
	pos := env.DebugPos
	if env.IP < len(pos) && g.Fileset != nil {
		g.Fprintf(g.Stdout, "// breakpoint at %s\n", g.Fileset.Position(pos[env.IP]))
	} else {
		g.Fprintf(g.Stdout, "// breakpoint\n")
	}

	prompt := g.Prompt
	g.Prompt = "godebug> "
	defer func() {
		g.Prompt = prompt
	}()

	g.Line = 0
	for ir.ReadParseEvalPrint() {
		g.Line = 0
	}
	return DebugCont
}
