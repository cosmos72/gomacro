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
 * init.go
 *
 *  Created on Apr 21, 2018
 *      Author Massimiliano Ghilardi
 */

package debug

import (
	"github.com/cosmos72/gomacro/fast"
)

func init() {
	// install debugger
	fast.CurrentDebugger = debuggerInterf{}
}

type debuggerInterf struct{}

func (debuggerInterf) Breakpoint(ir *fast.Interp, env *fast.Env) DebugOp {
	d := NewDebugger(ir, env)
	d.Show(true)
	return d.Repl()
}

func (debuggerInterf) At(ir *fast.Interp, env *fast.Env) DebugOp {
	d := NewDebugger(ir, env)
	d.Show(false)
	return d.Repl()
}
