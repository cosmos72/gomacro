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
 * import.go
 *
 *  Created on: Feb 28, 2017
 *      Author: Massimiliano Ghilardi
 */

package imports

import (
	. "reflect"
)

type Package struct {
	Binds   map[string]Value
	Types   map[string]Type
	Proxies map[string]Type
}

var Packages = make(map[string]Package)

// inception: allow interpreted code to import "github.com/cosmos72/gomacro/imports"
func init() {
	Packages["github.com/cosmos72/gomacro/imports"] = Package{
		map[string]Value{
			"Packages": ValueOf(&Packages).Elem(),
		},
		map[string]Type{
			"Package": TypeOf((*Package)(nil)).Elem(),
		},
		map[string]Type{},
	}
}
