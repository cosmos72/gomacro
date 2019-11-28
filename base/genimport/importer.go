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
 * importer.go
 *
 *  Created on Feb 27, 2017
 *      Author Massimiliano Ghilardi
 */

package genimport

import (
	"bytes"
	"fmt"
	"go/types"
	"io/ioutil"
	"os"
	r "reflect"

	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/base/paths"
	"github.com/cosmos72/gomacro/base/reflect"
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
	// 1. write a file $GOPATH/src/gomacro.imports/$PKGPATH/$PKGNAME.go containing a var Packages map[string]Package
	//    and a single func init() to populate it
	// 2. invoke "go build -buildmode=plugin" on the file to create a shared library
	// 3. load such shared library with plugin.Open().Lookup("Packages")
	ImPlugin
)

type PackageRef struct {
	imports.Package
	Path string
}

func (ref *PackageRef) DefaultName() string {
	return ref.Package.DefaultName(ref.Path)
}

func (ref *PackageRef) String() string {
	return fmt.Sprintf("{%s %q, %d binds, %d types}", ref.DefaultName(), ref.Path, len(ref.Binds), len(ref.Types))
}

type Importer struct {
	srcDir     string
	mode       types.ImportMode
	PluginOpen r.Value // = reflect.ValueOf(plugin.Open)
	output     *Output
}

func DefaultImporter(o *Output) *Importer {
	return &Importer{output: o}
}

func (imp *Importer) havePluginOpen() bool {
	if !imp.PluginOpen.IsValid() {
		imp.PluginOpen = imports.Packages["plugin"].Binds["Open"]
		if !imp.PluginOpen.IsValid() {
			imp.PluginOpen = reflect.None // cache the failure
		}
	}
	return imp.PluginOpen != reflect.None
}

// LookupPackage returns a package if already present in cache
func LookupPackage(alias, path string) *PackageRef {
	pkg, found := imports.Packages[path]
	if !found {
		return nil
	}
	if len(pkg.Name) == 0 {
		// missing pkg.Name, initialize it
		pkg.DefaultName(path)
		imports.Packages[path] = pkg
	}
	if len(alias) == 0 {
		// import "foo" => get alias from package name
		alias = pkg.DefaultName(path)
	}
	return &PackageRef{Package: pkg, Path: path}
}

func (imp *Importer) wrapImportError(path string, enableModule bool, err error) output.RuntimeError {
	if rerr, ok := err.(output.RuntimeError); ok {
		return rerr
	}
	if enableModule {
		return imp.output.MakeRuntimeError("error loading package %q metadata: %v", path, err)
	}
	return imp.output.MakeRuntimeError(
		"error loading package %q metadata, maybe you need to download (go get), compile (go build) and install (go install) it? %v",
		path, err)
}

func (imp *Importer) ImportPackage(alias, path string, enableModule bool) *PackageRef {
	ref, err := imp.ImportPackageOrError(alias, path, enableModule)
	if err != nil {
		panic(err)
	}
	return ref
}

func (imp *Importer) ImportPackageOrError(alias, pkgpath string, enableModule bool) (*PackageRef, error) {
	ref := LookupPackage(alias, pkgpath)
	if ref != nil {
		return ref, nil
	}
	o := imp.output
	gpkg, err := imp.Load(pkgpath, enableModule) // loads names and types, not the values!
	if err != nil {
		return nil, imp.wrapImportError(pkgpath, enableModule, err)
	}
	var mode ImportMode
	switch alias {
	case "_b":
		mode = ImBuiltin
	case "_i":
		mode = ImInception
	case "_3":
		mode = ImThirdParty
	default:
		if len(alias) == 0 {
			alias = gpkg.Name()
		}
		if imp.havePluginOpen() {
			mode = ImPlugin
		} else {
			mode = ImThirdParty
		}
	}
	file := createImportFile(imp.output, pkgpath, gpkg, mode)
	ref = &PackageRef{Path: pkgpath}
	if len(file) == 0 || mode != ImPlugin {
		// either the package exports nothing, or user must rebuild gomacro.
		// in both cases, still cache it to avoid recreating the file.
		imports.Packages[pkgpath] = ref.Package
		return ref, nil
	}
	soname := compilePlugin(o, file, enableModule, o.Stdout, o.Stderr)
	ipkgs := imp.loadPluginSymbol(soname, "Packages")
	pkgs := *ipkgs.(*map[string]imports.PackageUnderlying)

	// cache *all* packages found for future use
	imports.Packages.Merge(pkgs)

	// but return only requested one
	pkg, found := imports.Packages[pkgpath]
	if !found {
		return nil, imp.output.MakeRuntimeError(
			"error loading package %q: the compiled plugin %q does not contain it! internal error? %v",
			pkgpath, soname)
	}
	ref.Package = pkg
	return ref, nil
}

func createImportFile(o *Output, pkgpath string, pkg *types.Package, mode ImportMode) string {
	file := computeImportFilename(pkgpath, mode)

	buf := bytes.Buffer{}
	isEmpty := writeImportFile(o, &buf, pkgpath, pkg, mode)
	if isEmpty {
		o.Warnf("package %q exports zero constants, functions, types and variables", pkgpath)
		return ""
	}

	err := ioutil.WriteFile(file, buf.Bytes(), os.FileMode(0644))
	if err != nil {
		o.Errorf("error writing file %q: %v", file, err)
	}
	switch mode {
	case ImBuiltin, ImThirdParty:
		o.Warnf("created file %q, recompile gomacro to use it", file)
	case ImInception:
		o.Warnf("created file %q, recompile %s to use it", file, pkgpath)
	case ImPlugin:
		if GoModuleSupported {
			gomod := paths.Subdir(paths.DirName(file), "go.mod")
			err := ioutil.WriteFile(gomod, []byte("module gomacro.imports/"+pkgpath+"\n"), os.FileMode(0644))
			if err != nil {
				o.Errorf("error writing file %q: %v", gomod, err)
			}
		}
	}
	return file
}

func packageSanitizedName(path string) string {
	return sanitizeIdent(paths.FileName(path))
}

func sanitizeIdent(str string) string {
	return sanitizeIdent2(str, '_')
}

func sanitizeIdent2(str string, replacement rune) string {
	runes := []rune(str)
	for i, ch := range runes {
		if ch >= 'a' && ch <= 'z' {
			continue
		} else if ch >= 'A' && ch <= 'Z' {
			if i == 0 {
				// first rune must be lowercase to avoid conflict
				// with Packages, ValueOf, TypeOf
				runes[i] = ch - 'A' + 'a'
			}
			continue
		} else if i > 0 && (ch == '_' || (ch >= '0' && ch <= '9')) {
			continue
		}
		runes[i] = replacement
	}
	str = string(runes)
	if isReservedKeyword(str) {
		runes = append(runes, '_')
		str = string(runes)
	}
	return str
}

func computeImportFilename(path string, mode ImportMode) string {
	switch mode {
	case ImBuiltin:
		// user will need to recompile gomacro
		return paths.Subdir(paths.ImportsSrcDir, sanitizeIdent(path)+".go")
	case ImInception:
		// user will need to recompile gosrcdir / path
		return paths.Subdir(paths.GoSrcDir, path, "x_package.go")
	case ImThirdParty:
		// either plugin.Open is not available, or user explicitly requested import _3 "package".
		// In both cases, user will need to recompile gomacro
		return paths.Subdir(paths.ImportsSrcDir, "thirdparty", sanitizeIdent(path)+".go")
	}

	file := paths.FileName(path) + ".go"
	file = paths.Subdir(paths.GoSrcDir, "gomacro.imports", path, file)
	dir := paths.DirName(file)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		output.Errorf("error creating directory %q: %v", dir, err)
	}
	return file
}
