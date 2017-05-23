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
	GoPkg   interface{} // contains correspoding *go/types.Package if available
}

var Packages = make(map[string]Package)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/imports"
func init() {
	Packages["github.com/cosmos72/gomacro/imports"] = Package{
		Binds: map[string]Value{
			"Packages": ValueOf(&Packages).Elem(),
		},
		Types: map[string]Type{
			"Package": TypeOf((*Package)(nil)).Elem(),
		},
		Proxies: map[string]Type{},
	}
}

func (pkg *Package) Init() {
	pkg.Binds = make(map[string]Value)
	pkg.Types = make(map[string]Type)
	pkg.Proxies = make(map[string]Type)
}

func (pkg Package) SaveToPackages(path string) {
	// exploit the fact that maps are actually handles
	dst, ok := Packages[path]
	if !ok {
		dst = Package{}
		dst.Init()
		Packages[path] = dst
	}
	dst.Merge(pkg)
}

func (dst Package) Merge(src Package) {
	// exploit the fact that maps are actually handles
	for k, v := range src.Binds {
		dst.Binds[k] = v
	}
	for k, v := range src.Types {
		dst.Types[k] = v
	}
	for k, v := range src.Proxies {
		dst.Proxies[k] = v
	}
}
