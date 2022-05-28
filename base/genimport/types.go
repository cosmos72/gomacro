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
 * types.go
 *
 *  Created on May 28, 2022
 *      Author Massimiliano Ghilardi
 */

package genimport

import (
	"fmt"
	"path/filepath"

	"github.com/cosmos72/gomacro/imports"
)

type ImportMode int

const (
	// ImBuiltin import mechanism is:
	// 1. write a file $GOPATH/src/github.com/cosmos72/gomacro/imports/$PKGPATH.go containing a single func init()
	//    i.e. *inside* gomacro sources
	// 2. tell the user to recompile gomacro
	ImBuiltin ImportMode = iota

	// ImThirdParty import mechanism is the same as ImBuiltin, except that files are created in a thirdparty/ subdirectory:
	// 1. write a file $GOPATH/src/github.com/cosmos72/gomacro/imports/thirdparty/$PKGPATH.go containing a single func init()
	//    i.e. *inside* gomacro sources
	// 2. tell the user to recompile gomacro
	ImThirdParty

	// ImInception import mechanism is:
	// 1. write a file $GOPATH/src/$PKGPATH/x_package.go containing a single func init()
	//    i.e. *inside* the package to be imported
	// 2. tell the user to recompile $PKGPATH
	ImInception

	// ImPlugin import mechanism is:
	// 1. write a file $GOPATH/src/gomacro.imports/.../ containing a var Packages map[string]Package
	//    and a single func init() to populate it
	// 2. invoke "go build -buildmode=plugin" on the file to create a shared library
	// 3. load such shared library with plugin.Open().Lookup("Packages")
	ImPlugin
)

type PackageName = imports.PackageName // package default name, or package alias

// the full path of a Go package
type PackagePath struct {
	str string
}

// an absolute filesystem path - can be either a file or directory
type AbsolutePath struct {
	path string
}

type PackageRef struct {
	imports.Package
	Path string
}

// --------------------- ImportMode methods -----------------------------------

func (mode ImportMode) String() string {
	switch mode {
	case ImBuiltin:
		return "Builtin"
	case ImThirdParty:
		return "ThirdParty"
	case ImInception:
		return "Inception"
	case ImPlugin:
		return "Plugin"
	default:
		return fmt.Sprintf("ImportMode(%d)", int(mode))
	}
}

// --------------------- PackagePath methods -----------------------------------

func MakePackagePathOrPanic(path string) PackagePath {
	if isLocalFilesystemPath(path) {
		panic(fmt.Errorf("path %q is not a Go package path", path))
	}
	return PackagePath{path}
}

func (pkgpath PackagePath) String() string {
	return pkgpath.str
}

func (pkgpath PackagePath) Join(subpkg PackagePath) PackagePath {
	return pkgpath.JoinString(subpkg.String())
}

func (pkgpath PackagePath) JoinString(subdir string) PackagePath {
	if len(subdir) == 0 {
		return pkgpath
	} else if len(pkgpath.str) == 0 {
		return MakePackagePathOrPanic(subdir)
	} else {
		pkgpath.str += "/"
		pkgpath.str += subdir
		return pkgpath
	}
}

func packagePathsToString(pkgpaths []PackagePath) []string {
	ret := make([]string, len(pkgpaths))
	for i, pkgpath := range pkgpaths {
		ret[i] = pkgpath.str
	}
	return ret
}

// --------------------- AbsolutePath methods -----------------------------------

func MakeAbsolutePathOrError(path string) (AbsolutePath, error) {
	var ret AbsolutePath
	if !isLocalFilesystemPath(path) {
		return ret, fmt.Errorf(
			"path %q is not a supported (absolute or relative) filesystem path. Maybe prepend \"/\" or \"./\" ?",
			path)
	}
	path, err := filepath.Abs(path)
	if err != nil {
		return ret, err
	}
	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		return ret, err
	}
	return AbsolutePath{path}, nil
}

func MakeAbsolutePathOrPanic(path string) AbsolutePath {
	ret, err := MakeAbsolutePathOrError(path)
	if err != nil {
		panic(err)
	}
	return ret
}

func (abspath AbsolutePath) String() string {
	return abspath.path
}

// --------------------- PackageRef methods ------------------------------------

func (ref *PackageRef) DefaultName() PackageName {
	return ref.Package.DefaultName(ref.Path)
}

func (ref *PackageRef) String() string {
	return fmt.Sprintf("{%s %q, %d binds, %d types}", ref.DefaultName(), ref.Path, len(ref.Binds), len(ref.Types))
}
