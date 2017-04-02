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

package interpreter

import (
	r "reflect"

	"github.com/cosmos72/gomacro/compiler"
)

// temporary helper to invoke the new closure compiler
func (env *Env) compile(src string) {
	// parse + macroexpansion phase
	ast := env.ParseAst(src)

	// compile phase
	comp := env.CompEnv
	if comp == nil {
		comp = compiler.New()
		env.CompEnv = comp
	}
	fun := comp.CompileAst(ast)

	// print phase
	if env.Options&OptShowCompile != 0 {
		env.fprintValues(env.Stdout, r.ValueOf(fun))
	}

	// eval phase
	value, values := comp.Run(fun)

	// print phase
	if env.Options&OptShowEval != 0 {
		if len(values) != 0 {
			env.fprintValues(env.Stdout, values...)
		} else if value != None {
			env.fprintValues(env.Stdout, value)
		}
	}
}
