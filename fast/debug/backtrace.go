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
 * backtrace.go
 *
 *  Created on Apr 27, 2018
 *      Author Massimiliano Ghilardi
 */

package debug

import (
	"github.com/cosmos72/gomacro/fast"
)

func (d *Debugger) Backtrace(arg string) DebugOp {
	env := d.env
	var calls []*fast.Env
	for env != nil {
		if env.Caller != nil {
			// function body
			calls = append(calls, env)
			env = env.Caller
		} else {
			// nested env
			env = env.Outer
		}
	}
	d.showFunctionCalls(calls)
	return DebugRepl
}

func (d *Debugger) showFunctionCalls(calls []*fast.Env) {
	// show outermost stack frame first
	for i := len(calls) - 1; i >= 0; i-- {
		d.showFunctionCall(calls[i])
	}
}

func (d *Debugger) showFunctionCall(env *fast.Env) {
	g := d.globals
	c := env.DebugComp
	if c == nil || c.FuncMaker == nil {
		g.Fprintf(g.Stdout, "%p\tfunc (???) ???\n", env)
		return
	}
	m := c.FuncMaker

	g.Fprintf(g.Stdout, "%p\tfunc %s(", env, m.Name)
	d.showBinds(c, env, m.Param)
	g.Fprintf(g.Stdout, ") ")
	if len(m.Result) > 1 {
		g.Fprintf(g.Stdout, "(")
	}
	d.showBinds(c, env, m.Result)
	if len(m.Result) > 1 {
		g.Fprintf(g.Stdout, ")\n")
	} else {
		g.Fprintf(g.Stdout, "\n")
	}
}
