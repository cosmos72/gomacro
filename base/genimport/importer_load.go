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
 * importer_load.go
 *
 *  Created on Nov 16, 2019
 *      Author Massimiliano Ghilardi
 */

package genimport

import (
	"bufio"
	"fmt"
	"go/build"
	"go/importer"
	"go/types"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Go >= 1.14 requires a valid go.mod file in the directory used for packages.Config.Dir
//dir := computeImportDir(o, []string{pkgpath}, ImPlugin)

const GoModuleSupported bool = true

// Return the exported declarations of requested packages.
// If needed, prepare a Go module in 'dir' to list them.
//
// Note: paths can contain both Go package paths (corresponding to published Go packages)
// and filesystem absolute paths (corrisponding to local Go packages)
//
// Each returned map entry {string, *types.Package} will contain:
// 1. paths[i] as map key
// 2. *types.Package as map value, which contains both the Go package name and its Go package path
//
// allowing the caller to match each filesystem absolute path to the corresponding Go package
func (imp *Importer) Load(dir string, paths []string, enableModule bool) (
	retPkgs map[string]*types.Package, retErr error) {

	defer func() {
		if retPkgs == nil && retErr == nil {
			r := recover()
			if rerr, ok := r.(error); ok {
				retErr = rerr
			} else {
				retErr = fmt.Errorf("%v", r)
			}
		}
	}()
	pkgs := make(map[string]*types.Package, len(paths))

	if !enableModule {
		for _, path := range paths {
			pkg, err := importer.Default().Import(path)
			if err != nil {
				return nil, err
			}
			pkgs[path] = pkg
		}
		return pkgs, nil
	}

	goModReplaceDirective, pkgpaths, err := findLocalPackagesOnDisk(paths)
	if err != nil {
		return nil, err
	}

	o := imp.output
	createDir(o, dir)
	removeAllFilesInDir(o, dir)
	createPluginGoModFile(o, dir, goModReplaceDirective)

	env := environForCompiler(enableModule)

	// Go >= 1.16 usually requires running "go get ..." before "go list ..."
	// to start updating go.mod
	if err := runGoGetIfNeeded(o, dir, pkgpaths, env); err != nil {
		return nil, err
	}

	cfg := packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedImports | packages.NeedModule,
		Env:  env,
		Dir:  dir,
		Logf: nil, // imp.output.Debugf,
	}
	for i, pkgpath := range pkgpaths {
		list, err := packages.Load(&cfg, "pattern="+pkgpath.String())
		if err != nil {
			return nil, err
		}
		pkg, err := findPackageInList(pkgpath, list)
		if err != nil {
			return nil, err
		}
		// paths[i] can be either a Go package path, or an absolute filesystem path
		pkgs[paths[i]] = pkg
	}
	return pkgs, nil
}

func findPackageInList(pkgpath PackagePath, list []*packages.Package) (*types.Package, error) {
	for _, pkg := range list {
		if pkg.PkgPath == pkgpath.String() {
			if len(pkg.Errors) != 0 {
				return nil, errorList{pkg.Errors, mergeErrorMessages(pkg.Errors)}
			} else {
				return pkg.Types, nil
			}
		}
	}
	return nil, fmt.Errorf("packages.Load() could not find package %q", pkgpath.String())
}

type errorList struct {
	errors []packages.Error
	str    string
}

func (e errorList) Error() string {
	return e.str
}

func mergeErrorMessages(errors []packages.Error) string {
	str := make([]string, len(errors))
	for i, err := range errors {
		str[i] = err.Error()
	}
	return strings.Join(str, "\n")
}

func environForCompiler(enableModule bool) []string {
	env := append(os.Environ(),
		"GOARCH="+build.Default.GOARCH,
		"GOOS="+build.Default.GOOS,
		"GOROOT="+build.Default.GOROOT)
	if enableModule {
		env = append(env, "GO111MODULE=on")
	} else {
		env = append(env, "GO111MODULE=off")
	}
	return env
}

// for each path in pkgpaths that starts with '/'
// use it as filesystem path and scan it and its parents for a 'go.mod' file
func findLocalPackagesOnDisk(paths []string) (
	goModReplaceDirective map[PackagePath]AbsolutePath, pkgpaths []PackagePath, err error) {

	pkgpaths = make([]PackagePath, len(paths))
	for i, path := range paths {
		if !isLocalFilesystemPath(path) {
			pkgpaths[i] = MakePackagePathOrPanic(path)
			continue
		}
		abspath := MakeAbsolutePathOrPanic(path)
		onDiskModuleDir, goModulePkgpath, relativePkgpath, isErr := findLocalPackageOnDisk(abspath)
		if isErr != nil {
			return nil, nil, isErr
		}
		pkgpaths[i] = goModulePkgpath.Join(relativePkgpath)
		if goModReplaceDirective == nil {
			goModReplaceDirective = make(map[PackagePath]AbsolutePath)
		}
		goModReplaceDirective[goModulePkgpath] = onDiskModuleDir
	}
	return goModReplaceDirective, pkgpaths, nil
}

// given an absolute filesystem path, find a 'go.mod' file in it or its parent paths,
// and if found, return the absolute path of the directory containing the 'go.mod' file,
// the name of the the module name written inside the go.mod file,
// and the relative package path to be imported - i.e. relative to the module
func findLocalPackageOnDisk(onDiskDir AbsolutePath) (onDiskModuleDir AbsolutePath, goModulePkgpath PackagePath, relativePkgpath PackagePath, retErr error) {
	dir := onDiskDir
	var err error
	for {
		goModFile := filepath.Join(dir.String(), "go.mod")
		goModulePkgpath, err = readGoModFile(goModFile)
		if err == nil {
			return dir, goModulePkgpath, relativePkgpath, nil
		}
		relativePkgpath = relativePkgpath.JoinString(filepath.Base(dir.String()))

		parentDir := filepath.Dir(dir.String())
		if len(parentDir) == 0 || parentDir == dir.String() {
			break
		}
		dir = AbsolutePath{parentDir}
	}
	return AbsolutePath{}, PackagePath{}, PackagePath{},
		fmt.Errorf("packages.Load() could not find \"go.mod\" file in directory %q and all its parent directories",
			onDiskDir)
}

func readGoModFile(filepath string) (PackagePath, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return PackagePath{}, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return PackagePath{}, err
		}
		if !startsWith(line, "module ") {
			continue
		}
		pkgpath := MakePackagePathOrPanic(strings.TrimSpace(line[7:]))
		if len(pkgpath.String()) == 0 {
			return pkgpath, fmt.Errorf("packages.Load() found file %q, but it contains no package name after \"module \"")
		}
		return pkgpath, nil
	}
}

func isLocalFilesystemPath(path string) bool {
	return path == "/" || path == "." || path == ".." ||
		startsWith(path, "/") || startsWith(path, "./") || startsWith(path, "../")
}

func startsWith(str string, prefix string) bool {
	return len(str) >= len(prefix) && str[0:len(prefix)] == prefix
}

func IsLocalImportPath(pkgpath string) bool {
	return isLocalFilesystemPath(pkgpath)
}
