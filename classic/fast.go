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
 * fast.go
 *
 *  Created on: Apr 02, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	r "reflect"

	"github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/fast"
)

// temporary helper to invoke the new fast interpreter
func (env *Env) fastEval(form ast2.Ast) (r.Value, []r.Value) {
	// compile phase

	var ce *fast.CompEnv
	if env.CompEnv == nil {
		ce = fast.New()
		ce.Comp.CompileOptions |= fast.CompileKeepUntyped
		env.CompEnv = ce
	} else {
		ce = env.CompEnv.(*fast.CompEnv)
	}
	ce.Comp.Stringer.Copy(&env.Stringer) // sync Fileset, Pos, Line
	fun := ce.Comp.CompileAst(form)

	// debug output
	if env.Options&base.OptShowCompile != 0 {
		env.FprintValues(env.Options, env.Stdout, r.ValueOf(fun))
	}

	// eval phase
	return ce.Run(fun)
}
