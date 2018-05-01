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
 * global.go
 *
 *  Created on Apr 21, 2018
 *      Author Massimiliano Ghilardi
 */

package debug

import (
	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/fast"
)

type DebugOp = base.DebugOp

const (
	DebugContinue = base.SigDebugContinue
	DebugFinish   = base.SigDebugFinish
	DebugNext     = base.SigDebugNext
	DebugStep     = base.SigDebugStep
	DebugRepl     = base.SigDebugRepl
)

type Debugger struct {
	interp  *fast.Interp
	env     *fast.Env
	globals *base.Globals
	lastcmd string
}

func (d *Debugger) Breakpoint(interp *fast.Interp, env *fast.Env) DebugOp {
	return d.main(interp, env, true)
}

func (d *Debugger) At(interp *fast.Interp, env *fast.Env) DebugOp {
	return d.main(interp, env, false)
}

func (d *Debugger) main(interp *fast.Interp, env *fast.Env, breakpoint bool) DebugOp {
	// create an inner Interp to preserve existing Binds, compiled Code and IP
	//
	// this is needed to allow compiling and evaluating code at a breakpoint or single step
	// without disturbing the code being debugged
	d.interp = fast.NewInnerInterp(interp, "debug", "debug")
	d.env = env
	d.globals = &env.Run.Globals
	if !d.Show(breakpoint) {
		// skip synthetic statements
		return DebugRepl
	}
	return d.Repl()
}
