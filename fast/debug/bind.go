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
 * bind.go
 *
 *  Created on Apr 27, 2018
 *      Author Massimiliano Ghilardi
 */

package debug

import (
	r "reflect"
	"sort"

	"github.com/cosmos72/gomacro/fast"
)

func (d *Debugger) Vars() {
	env := d.env
	var envs []*fast.Env
	for env != nil {
		envs = append(envs, env)
		if env.Caller != nil {
			// function body
			break
		}
		// nested env
		env = env.Outer
	}
	d.showEnvs(envs)
}

func (d *Debugger) showEnvs(envs []*fast.Env) {
	// show outermost scope
	for i := len(envs) - 1; i >= 0; i-- {
		d.showEnv(envs[i])
	}
}

func (d *Debugger) showEnv(env *fast.Env) {
	c := env.DebugComp
	if c == nil || (c.BindNum == 0 && c.IntBindNum == 0) {
		return
	}
	g := d.globals
	g.Fprintf(g.Stdout, "// ------------------------\n")
	binds := make([]*fast.Bind, len(c.Binds))
	i := 0
	for _, bind := range c.Binds {
		binds[i] = bind
		i++
	}
	sort.Slice(binds, func(i, j int) bool {
		return binds[i].Name < binds[j].Name
	})
	for _, bind := range binds {
		value := getBindValue(c, env, bind)
		g.Fprintf(g.Stdout, "%s\t= %v\t// %v\n", bind.Name, value, bind.Type)
	}
}

// =============================================================================

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

func (d *Debugger) showBinds(c *fast.Comp, env *fast.Env, binds []*fast.Bind) {
	g := d.globals
	for i, bind := range binds {
		if i != 0 {
			g.Fprintf(g.Stdout, ", ")
		}
		d.showBind(c, env, bind)
	}
}

func (d *Debugger) showBind(c *fast.Comp, env *fast.Env, bind *fast.Bind) {
	value := getBindValue(c, env, bind)
	var ivalue interface{} = value
	if !value.IsValid() {
		ivalue = "nil"
	}

	g := d.globals
	if name := bind.Name; len(name) != 0 {
		g.Fprintf(g.Stdout, "%s=%v <%v>", name, ivalue, bind.Type)
	} else {
		g.Fprintf(g.Stdout, "%v <%v>", ivalue, bind.Type)
	}
}
