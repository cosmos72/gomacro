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
 * eval.go
 *
 *  Created on: Apr 02, 2017
 *      Author: Massimiliano Ghilardi
 */

package compiler

import (
	r "reflect"
)

func (c *CompEnv) Run(fun X) (r.Value, []r.Value) {
	c.growEnv()
	return fun(c.Env)
}

// DefConst compiles a constant declaration, then executes it
func (c *CompEnv) DefConst(name string, t r.Type, value I) (r.Value, []r.Value) {
	fun := c.DeclConst0(name, t, value)
	c.growEnv()
	return fun(c.Env)
}

// DefVar compiles a variable declaration, then executes it
func (c *CompEnv) DefVar(name string, t r.Type, value I) (r.Value, []r.Value) {
	fun := c.DeclVar0(name, t, value)
	c.growEnv()
	return fun(c.Env)
}

func (c *CompEnv) growEnv() {
	// usually we know at Env creation how many slots are needed in c.Env.Binds
	// but here we are modifying an existing Env...
	for len(c.Env.Binds) < len(c.Binds) {
		c.Env.Binds = append(c.Env.Binds, None)
	}
}
