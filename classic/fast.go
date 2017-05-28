/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
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
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
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
	xr "github.com/cosmos72/gomacro/xreflect"
)

// temporary helper to invoke the new fast interpreter
func (env *Env) fastEval(form ast2.Ast) (r.Value, []r.Value, xr.Type, []xr.Type) {
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
	ce.Comp.Options = env.Options        // sync Options

	expr := ce.Comp.Compile(form)

	// debug output
	if env.Options&base.OptShowCompile != 0 {
		env.Fprintf(env.Stdout, "%v\n", expr)
	}
	// eval phase
	if expr == nil {
		return base.None, nil, nil, nil
	}
	value, values := ce.RunExpr(expr)
	return value, values, expr.Type, expr.Types
}
