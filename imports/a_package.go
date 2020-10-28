/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
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

	syscall "github.com/cosmos72/gomacro/imports/syscall"
	thirdparty "github.com/cosmos72/gomacro/imports/thirdparty"
	"github.com/cosmos72/gomacro/imports/util"
)

type PackageUnderlying = struct { // unnamed
	Name    string
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
			"Package":           TypeOf((*Package)(nil)).Elem(),
			"PackageMap":        TypeOf((*PackageMap)(nil)).Elem(),
			"PackageUnderlying": TypeOf((*PackageUnderlying)(nil)).Elem(),
		},
	}
	Packages.Merge(syscall.Packages)
	Packages.Merge(thirdparty.Packages)
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
		pkg.LazyInit(path)
		pkgs[path] = pkg
	}
}

/**
 * return the default name to bind when importing Package
 *
 * https://golang.org/ref/spec#Package_clause states:
 * If the PackageName is omitted, it defaults to the identifier
 * specified in the package clause of the imported package
 *
 * So use that if known, otherwise extrapolate it from package path
 */
func (pkg *Package) DefaultName(path string) string {
	if len(pkg.Name) == 0 {
		pkg.Name = util.TailIdentifier(util.FileName(path))
	}
	return pkg.Name
}

func (pkg *Package) LazyInit(path string) {
	pkg.DefaultName(path)
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

func (dst *Package) Merge(src PackageUnderlying) {
	if len(src.Name) != 0 {
		dst.Name = src.Name
	}
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
