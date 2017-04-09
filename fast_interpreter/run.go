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

package fast_interpreter

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (c *CompEnv) Run(fun func(*Env) (r.Value, []r.Value)) (r.Value, []r.Value) {
	c.growEnv(128)
	if fun != nil {
		return fun(c.Env)
	} else {
		return None, nil
	}
}

// DefConst compiles a constant declaration, then executes it
func (c *CompEnv) DefConst(name string, t r.Type, value I) {
	c.DeclConst0(name, t, value)

}

// DefVar compiles a variable declaration, then executes it
func (c *CompEnv) DefVar(name string, t r.Type, value I) {
	c.Code.Clear()
	c.DeclVar0(name, t, ExprValue(value))
	fun := c.Code.AsXV()
	if fun != nil {
		c.growEnv(128)
		fun(c.Env)
	}
}

// DefType compiles a type declaration
func (c *CompEnv) DefType(name string, t r.Type) {
	c.DeclType0(name, t)
}

func (c *CompEnv) growEnv(minDelta int) {
	// usually we know at Env creation how many slots are needed in c.Env.Binds
	// but here we are modifying an existing Env...
	if minDelta < 0 {
		minDelta = 0
	}
	capacity, min := cap(c.Env.Binds), c.BindNum
	if capacity < min {
		if capacity <= min/2 {
			capacity = min
		} else {
			capacity *= 2
		}
		if capacity-min < minDelta {
			capacity = min + minDelta
		}
		binds := make([]r.Value, min, capacity)
		copy(binds, c.Env.Binds)
		c.Env.Binds = binds
	}
	if len(c.Env.Binds) < min {
		c.Env.Binds = c.Env.Binds[0:min]
	}

	capacity, min = cap(c.Env.IntBinds), c.IntBindNum
	if capacity < min {
		if capacity <= min/2 {
			capacity = min
		} else {
			capacity *= 2
		}
		if capacity-min < minDelta {
			capacity = min + minDelta
		}
		binds := make([]uint64, min, capacity)
		copy(binds, c.Env.IntBinds)
		c.Env.IntBinds = binds
	}
	if len(c.Env.IntBinds) < min {
		c.Env.IntBinds = c.Env.IntBinds[0:min]
	}
}
