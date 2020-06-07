/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * func0ret0.go
 *
 *  Created on Apr 16, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) func0ret0(t xr.Type, m *funcMaker) func(env *Env) xr.Value {
	funcbody := m.funcbody
	if funcbody == nil {
		return func(env *Env) xr.Value {
			return valueOfNopFunc
		}
	}
	var debugC *Comp
	if c.Globals.Options&base.OptDebugger != 0 {
		debugC = c
	}

	nbind := m.nbind
	nintbind := m.nintbind
	return func(env *Env) xr.Value {
		// function is closed over the env used to DECLARE it
		env.MarkUsedByClosure()
		return xr.ValueOf(func() {
			env := newEnv4Func(env, nbind, nintbind, debugC)
			// execute the body
			funcbody(env)

			env.freeEnv4Func()
		})
	}
}
