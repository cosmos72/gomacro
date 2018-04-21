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
}

func NewDebugger(interp *fast.Interp, env *fast.Env) *Debugger {
	return &Debugger{
		interp:  interp,
		env:     env.Outer,
		globals: interp.Comp.Globals,
	}
}
