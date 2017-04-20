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

package fast

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (ce *CompEnv) Exec1(expr *Expr) r.Value {
	ce.PrepareEnv()
	if expr != nil {
		return expr.AsX1()(ce.env)
	} else {
		return None
	}
}

func (ce *CompEnv) Exec(expr *Expr) (r.Value, []r.Value) {
	ce.PrepareEnv()
	if expr != nil {
		return expr.AsXV(CompileDefaults)(ce.env)
	} else {
		return None, nil
	}
}

func (ce *CompEnv) Run(fun func(*Env) (r.Value, []r.Value)) (r.Value, []r.Value) {
	ce.PrepareEnv()
	if fun != nil {
		return fun(ce.env)
	} else {
		return None, nil
	}
}

// combined ParseAst + CompileAst + Run
func (ce *CompEnv) Eval(src string) (r.Value, []r.Value) {
	c := ce.Comp
	return ce.Run(c.CompileAst(c.ParseAst(src)))
}

// DeclConst compiles a constant declaration
func (ce *CompEnv) DeclConst(name string, t r.Type, value I) {
	ce.Comp.DeclConst0(name, t, value)
}

// DeclFunc compiles a function declaration
func (ce *CompEnv) DeclFunc(name string, fun I) {
	ce.Comp.DeclFunc0(name, fun)
}

// DeclType compiles a type declaration
func (ce *CompEnv) DeclType(name string, t r.Type) {
	ce.Comp.DeclType0(name, t)
}

// DeclVar compiles a variable declaration
func (ce *CompEnv) DeclVar(name string, t r.Type, value I) {
	ce.Comp.DeclVar0(name, t, exprValue(value))
}

// Apply executes the compiled declarations, statements and expressions,
// then clears the compiled buffer
func (ce *CompEnv) Apply() {
	exec := ce.Comp.Code.Exec()
	if exec != nil {
		exec(ce.PrepareEnv())
	}
}

// AddressOfVar compiles the expression &name, then executes it
func (ce *CompEnv) AddressOfVar(name string) r.Value {
	expr := ce.Comp.AddressOfVar(name)
	return ce.Exec1(expr)
}

// ValueOf retrieves the value of a constant, function or variable
func (ce *CompEnv) ValueOf(name string) r.Value {
	expr := ce.Comp.Ident(name)
	return ce.Exec1(expr)
}

func (ce *CompEnv) PrepareEnv() *Env {
	return ce.prepareEnv(128)
}

func (ce *CompEnv) prepareEnv(minDelta int) *Env {
	c := ce.Comp
	env := ce.env
	// usually we know at Env creation how many slots are needed in c.Env.Binds
	// but here we are modifying an existing Env...
	if minDelta < 0 {
		minDelta = 0
	}
	capacity, min := cap(env.Binds), c.BindNum
	// c.Debugf("prepareEnv() before: c.BindNum = %v, minDelta = %v, len(env.Binds) = %v, cap(env.Binds) = %v, env = %p", c.BindNum, minDelta, len(env.Binds), cap(env.Binds), env)

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
		copy(binds, env.Binds)
		env.Binds = binds
	}
	if len(env.Binds) < min {
		env.Binds = env.Binds[0:min:cap(env.Binds)]
	}
	// c.Debugf("prepareEnv() after:  c.BindNum = %v, minDelta = %v, len(env.Binds) = %v, cap(env.Binds) = %v, env = %p", c.BindNum, minDelta, len(env.Binds), cap(env.Binds), env)

	capacity, min = cap(env.IntBinds), c.IntBindNum
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
		copy(binds, env.IntBinds)
		env.IntBinds = binds
	}
	if len(env.IntBinds) < min {
		env.IntBinds = env.IntBinds[0:min:cap(env.IntBinds)]
	}
	return env
}
