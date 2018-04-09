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
 * a_package.go
 *
 *  Created on: Feb 28, 2017
 *      Author: Massimiliano Ghilardi
 */

package imports

import (
	. "reflect"

	"github.com/cosmos72/gomacro/imports/syscall"
)

type PackageUnderlying = struct { // unnamed
	Binds   map[string]Value
	Types   map[string]Type
	Proxies map[string]Type
	// Untypeds contains a string representation of untyped constants,
	// stored without loss of precision
	Untypeds map[string]string
	// Wrappers is the list of wrapper methods for named types.
	// Stored explicitly because reflect package cannot distinguish
	// between explicit methods and wrapper methods for embedded fields
	Wrappers map[string][]string
}

type Package PackageUnderlying // named, can have methods

type PackageMap map[string]Package // named, can have methods

var Packages = make(PackageMap)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/imports"
func init() {
	Packages["github.com/cosmos72/gomacro/imports"] = Package{
		Binds: map[string]Value{
			"Packages": ValueOf(&Packages).Elem(),
		},
		Types: map[string]Type{
			"Package": TypeOf((*Package)(nil)).Elem(),
		},
	}
	Packages.Merge(syscall.Packages)
}

func (pkg *Package) LazyInit() {
	if pkg.Binds == nil {
		pkg.Binds = make(map[string]Value)
	}
	if pkg.Types == nil {
		pkg.Types = make(map[string]Type)
	}
	if pkg.Proxies == nil {
		pkg.Proxies = make(map[string]Type)
	}
	if pkg.Untypeds == nil {
		pkg.Untypeds = make(map[string]string)
	}
	if pkg.Wrappers == nil {
		pkg.Wrappers = make(map[string][]string)
	}
}

func (pkgs PackageMap) Merge(srcs map[string]PackageUnderlying) {
	// exploit the fact that maps are actually handles
	for path, src := range srcs {
		pkgs.MergePackage(path, src)
	}
}

func (pkgs PackageMap) MergePackage(path string, src PackageUnderlying) {
	// exploit the fact that maps are actually handles
	pkg, ok := pkgs[path]
	if ok {
		pkg.Merge(src)
	} else {
		pkg = Package(src)
		pkg.LazyInit()
		pkgs[path] = pkg
	}
}

func (dst Package) Merge(src PackageUnderlying) {
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
	for k, v := range src.Untypeds {
		dst.Untypeds[k] = v
	}
	for k, v := range src.Wrappers {
		dst.Wrappers[k] = v
	}
}
