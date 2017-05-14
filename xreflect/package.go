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

package xtype

import (
	"go/types"
	"strings"
)

func newPackage(path, name string) *types.Package {
	if len(path) == 0 {
		// do not create unnamed packages
		return nil
	}
	if len(name) == 0 {
		name = path[1+strings.LastIndexByte(path, '/'):]
	}
	return types.NewPackage(path, name)
}

func NewPackage(path, name string) *Package {
	return (*Package)(newPackage(path, name))
}

func (pkg *Package) GoPackage() *types.Package {
	return (*types.Package)(pkg)
}

func (pkg *Package) Name() string {
	if pkg == nil {
		return ""
	}
	return (*types.Package)(pkg).Name()
}

func (pkg *Package) Path() string {
	if pkg == nil {
		return ""
	}
	return (*types.Package)(pkg).Path()
}
