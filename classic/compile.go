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
 * compile.go
 *
 *  Created on: Apr 02, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/fast"
)

// temporary helper to invoke the new closure compiler
func (env *Env) compile(src string) {
	// parse + macroexpansion phase
	ast := env.ParseAst(src)

	var value r.Value
	var values []r.Value

	if env.Options&OptMacroExpandOnly == 0 {
		// compile phase

		var ce *fast.CompEnv
		if env.CompEnv == nil {
			ce = fast.New()
			ce.Comp.CompileOptions |= fast.CompileKeepUntyped
			env.CompEnv = ce
		} else {
			ce = env.CompEnv.(*fast.CompEnv)
		}
		fun := ce.Comp.CompileAst(ast)

		// print phase
		if env.Options&OptShowCompile != 0 {
			env.FprintValues(env.Options, env.Stdout, r.ValueOf(fun))
		}

		// eval phase
		value, values = ce.Run(fun)
	} else {
		value = r.ValueOf(ast.Interface())
	}

	// print phase
	if env.Options&OptShowEval != 0 {
		if len(values) != 0 {
			env.FprintValues(env.Options, env.Stdout, values...)
		} else if value != None {
			env.FprintValues(env.Options, env.Stdout, value)
		}
	}
}
