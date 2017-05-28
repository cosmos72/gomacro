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
 * interpreter_common.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	r "reflect"

	"github.com/cosmos72/gomacro/base"
)

type ThreadGlobals struct {
	base.Globals
	AllMethods map[r.Type]Methods // methods implemented by interpreted code
	CompEnv    interface{}        // *fast.CompEnv // temporary...
}

func NewThreadGlobals() *ThreadGlobals {
	tg := &ThreadGlobals{
		AllMethods: make(map[r.Type]Methods),
	}
	tg.Globals.Init()
	return tg
}
