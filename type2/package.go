/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * package.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package type2

import (
	"go/types"
	"strings"
)

func (pkg Package) Name() string {
	if pkg.impl == nil {
		return ""
	}
	return pkg.impl.Name()
}

func (pkg Package) Path() string {
	if pkg.impl == nil {
		return ""
	}
	return pkg.impl.Path()
}

func (pkg Package) String() string {
	if pkg.impl == nil {
		return ""
	}
	return pkg.impl.Name()
}

func NewPackage(path, name string) Package {
	if len(path) == 0 {
		// do not create unnamed packages
		return Package{}
	}
	if len(name) == 0 {
		name = path[1+strings.LastIndexByte(path, '/'):]
	}
	return Package{
		impl: types.NewPackage(path, name),
	}
}

func (pkg Package) GoPackage() *types.Package {
	return pkg.impl
}

func makepackage(gpkg *types.Package) Package {
	return Package{
		impl: gpkg,
	}
}
