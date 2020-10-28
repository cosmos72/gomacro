/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * paths.go
 *
 *  Created on: Jun 24, 2018
 *      Author: Massimiliano Ghilardi
 */

package paths

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/cosmos72/gomacro/imports"
	"github.com/cosmos72/gomacro/imports/util"
)

// return the string after last '/' in path
func FileName(path string) string {
	return util.FileName(path)
}

// return the string up to (and including) last '/' in path
func DirName(path string) string {
	return path[0 : 1+strings.LastIndexByte(path, '/')]
}

// remove last byte from string
func RemoveLastByte(s string) string {
	if n := len(s); n != 0 {
		s = s[:n-1]
	}
	return s
}

// always use forward slashes. they work also on Windows...
func unixpath(path string) string {
	if os.PathSeparator != '/' && len(path) != 0 {
		path = strings.Replace(path, string(os.PathSeparator), "/", -1)
	}
	return path
}

// find user's home directory, see https://stackoverflow.com/questions/2552416/how-can-i-find-the-users-home-dir-in-a-cross-platform-manner-using-c
// without importing "os/user" - which requires cgo to work thus makes cross-compile difficult, see https://github.com/golang/go/issues/11797
func UserHomeDir() string {
	home := os.Getenv("HOME")
	if len(home) == 0 {
		home = os.Getenv("USERPROFILE")
		if len(home) == 0 {
			home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		}
	}
	return unixpath(home)
}

func Subdir(dirs ...string) string {
	// should use string(os.PathSeparator) instead of "/', but:
	// 1) package names use '/', not os.PathSeparator
	// 2) it would complicate DirName()
	return strings.Join(dirs, "/")
}

var (
	GoSrcDirs = goSrcDirs()
	GoSrcDir  = GoSrcDirs[0]

	// where to find the Go compiler used to compile gomacro.
	// needed to build compatible plugins
	GoRootDir = build.Default.GOROOT

	// a value whose type is defined in the imports package.
	// used to locate the source directory where to write import files.
	SymbolFromImportsPackage interface{} = imports.Packages

	// directory where to write import files - computed lazily.
	// Current autodetection mechanism already supports vendored or cloned copies of gomacro.
	// To change it, set the variable SymbolFromImportsPackage.
	importsSrcDir string
)

func goSrcDirs() []string {
	gopath := build.Default.GOPATH
	if len(gopath) == 0 {
		// GOARCH=wasm reports empty GOPATH
		return []string{"src"}
	}
	var srcdirs []string
	for _, path := range filepath.SplitList(gopath) {
		srcdirs = append(srcdirs, filepath.Join(path, "src"))
	}
	return srcdirs
}

// lazily compute the directory where to write imports
func GetImportsSrcDir() string {
	if importsSrcDir == "" {
		importsSrcDir = findPkgSrcDir(SymbolFromImportsPackage)
	}
	return importsSrcDir
}

// return the source directory inside GOPATH
// where the package containing the declaration of x's type
// should be located.
func findPkgSrcDir(x interface{}) string {
	pkg := reflect.TypeOf(x).PkgPath()
	pkg = filepath.Join(strings.Split(pkg, "/")...)
	gopath := build.Default.GOPATH
	for _, dir := range filepath.SplitList(gopath) {
		path := filepath.Join(dir, "src", pkg)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	defaultDir := filepath.Join(GoSrcDir, pkg)
	if false {
		// disable this warning, it unnecessarily worries users
		fmt.Printf(`// warning: could not find package %q in $GOPATH = %q
//          command 'import _b "path/to/some/package"' may not work correctly.

`, pkg, gopath)
	}
	return defaultDir
}
