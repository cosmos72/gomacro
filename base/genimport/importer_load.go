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
	"fmt"
	"go/build"
	"go/importer"
	"go/types"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Go >= 1.14 requires a valid go.mod file in the directory used for packages.Config.Dir
//dir := computeImportDir(o, []string{pkgpath}, ImPlugin)

const GoModuleSupported bool = true

func (imp *Importer) Load(dir string, pkgpaths []string, enableModule bool) (ret_pkgs map[string]*types.Package, ret_err error) {
	defer func() {
		if ret_pkgs == nil && ret_err == nil {
			r := recover()
			if rerr, ok := r.(error); ok {
				ret_err = rerr
			} else {
				ret_err = fmt.Errorf("%v", r)
			}
		}
	}()
	pkgs := make(map[string]*types.Package, len(pkgpaths))

	if !enableModule {
		for _, pkgpath := range pkgpaths {
			pkg, err := importer.Default().Import(pkgpath)
			if err != nil {
				return nil, err
			}
			pkgs[pkgpath] = pkg
		}
		return pkgs, nil
	}

	o := imp.output
	createDir(o, dir)
	removeAllFilesInDir(o, dir)
	createPluginGoModFile(o, dir)

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
	for _, pkgpath := range pkgpaths {
		list, err := packages.Load(&cfg, "pattern="+pkgpath)
		if err != nil {
			return nil, err
		}
		pkg, err := findPackage(pkgpath, list)
		if err != nil {
			return nil, err
		}
		pkgs[pkgpath] = pkg
	}
	return pkgs, nil
}

func findPackage(pkgpath string, list []*packages.Package) (*types.Package, error) {
	for _, pkg := range list {
		if pkg.PkgPath == pkgpath {
			if len(pkg.Errors) != 0 {
				return nil, errorList{pkg.Errors, mergeErrorMessages(pkg.Errors)}
			} else {
				return pkg.Types, nil
			}
		}
	}
	return nil, fmt.Errorf("packages.Load() could not find package %q", pkgpath)
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
