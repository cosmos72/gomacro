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
 * inspect.go
 *
 *  Created on: Feb 11, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (env *Env) Inspect(str string) {
	inspector := env.Globals.Inspector
	if inspector == nil {
		env.Errorf("no inspector set: call Interp.SetInspector() first")
	}

	form := env.Parse(str)
	v := env.EvalAst1(form)
	var t r.Type
	if v.IsValid() && v != None {
		if v.CanInterface() {
			t = r.TypeOf(v.Interface()) // show concrete type
		} else {
			t = v.Type()
		}
	}
	inspector(str, v, t, nil, env.Globals)
}
