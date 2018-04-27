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
	r "reflect"

	"github.com/cosmos72/gomacro/base"
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
	params := getBindValues(c, env, m.Param)
	results := getBindValues(c, env, m.Result)

	g.Fprintf(g.Stdout, "%p\tfunc %s(", env, m.Name)
	showBinds(g, m.Param, params)
	g.Fprintf(g.Stdout, ") ")
	if len(results) > 1 {
		g.Fprintf(g.Stdout, "(")
	}
	showBinds(g, m.Result, results)
	if len(results) > 1 {
		g.Fprintf(g.Stdout, ")\n")
	} else {
		g.Fprintf(g.Stdout, "\n")
	}
}

func getBindValues(c *fast.Comp, env *fast.Env, binds []*fast.Bind) []r.Value {
	values := make([]r.Value, len(binds))
	for i, bind := range binds {
		values[i] = getBindValue(c, env, bind)
	}
	return values
}

func getBindValue(c *fast.Comp, env *fast.Env, bind *fast.Bind) r.Value {
	e := c.Symbol(bind.AsSymbol(0))
	var value r.Value
	if e.Const() {
		value = r.ValueOf(e.Value)
	} else {
		value = e.AsX1()(env)
	}
	return value
}

func showBinds(g *base.Globals, binds []*fast.Bind, values []r.Value) {
	for i, bind := range binds {
		if i != 0 {
			g.Fprintf(g.Stdout, ", ")
		}
		showBind(g, bind, values[i])
	}
}

func showBind(g *base.Globals, bind *fast.Bind, value r.Value) {
	if name := bind.Name; len(name) != 0 {
		g.Fprintf(g.Stdout, "%s=%v <%v>", name, value, bind.Type)
	} else {
		g.Fprintf(g.Stdout, "%v <%v>", value, bind.Type)
	}
}
