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
	"go/build"
	"go/types"
	"io/ioutil"
	"os"
	r "reflect"
	"sync"
	"sync/atomic"

	"github.com/cosmos72/gomacro/base/output"
	p "github.com/cosmos72/gomacro/base/paths"
	"github.com/cosmos72/gomacro/base/reflect"
	"github.com/cosmos72/gomacro/imports"
)

type Importer struct {
	srcDir     string
	mode       types.ImportMode
	PluginOpen r.Value // = reflect.ValueOf(plugin.Open)
	output     *Output
	BuildArgs  []string
}

func DefaultImporter(o *Output) *Importer {
	return &Importer{output: o}
}

func (imp *Importer) havePluginOpen() bool {
	if !imp.PluginOpen.IsValid() {
		imp.PluginOpen = imports.Packages["plugin"].Binds["Open"]
		if !imp.PluginOpen.IsValid() {
			imp.PluginOpen = reflect.NoneR // cache the failure
		}
	}
	return imp.PluginOpen != reflect.NoneR
}

func srcDirForPluginImport() string {
	return p.Subdir(p.GoSrcDir, "gomacro.imports", fmt.Sprintf("gomacro_pid_%d", os.Getpid()))
}

// LookupPackage returns a package if already present in cache
func LookupPackage(alias PackageName, pkgpath string) *PackageRef {
	pkg, found := imports.Packages[pkgpath]
	if !found {
		return nil
	}
	if len(pkg.Name) == 0 {
		// missing pkg.Name, initialize it
		_ = pkg.DefaultName(pkgpath)
		imports.Packages[pkgpath] = pkg
	}
	if len(alias) == 0 {
		// import "foo" => get alias from package name
		alias = pkg.DefaultName(pkgpath)
	}
	pkg.Validate(pkgpath)
	return &PackageRef{Package: pkg, Path: pkgpath}
}

func (imp *Importer) wrapImportError(pkgpaths []string, enableModule bool, err error) output.RuntimeError {
	if rerr, ok := err.(output.RuntimeError); ok {
		return rerr
	}
	if enableModule {
		return imp.output.MakeRuntimeError("error loading packages %v metadata: %v", pkgpaths, err)
	}
	return imp.output.MakeRuntimeError(
		"error loading packages %v metadata, maybe you need to download (go get), compile (go build) and install (go install) them? %v",
		pkgpaths, err)
}

// path can be a Go package path or and absolute filesystem path
func (imp *Importer) ImportPackage(alias PackageName, path string, enableModule bool) *PackageRef {
	refs, err := imp.ImportPackagesOrError(map[string]PackageName{path: alias}, enableModule)
	if err != nil {
		panic(err)
	}
	return refs[path]
}

// pathMap is a map (Go package path or absolute filesystem path) -> package alias
func (imp *Importer) ImportPackagesOrError(pathMap map[string]PackageName, enableModule bool) (map[string]*PackageRef, error) {

	refs := make(map[string]*PackageRef)
	pathMap = clone(pathMap)
	for path, alias := range pathMap {
		ref := LookupPackage(alias, path)
		if ref != nil {
			// found in package cache, no need to compile and load this package
			refs[path] = ref
			delete(pathMap, path)
		}
	}
	if len(pathMap) != 0 {
		imported, err := imp.doImportPackagesOrError(pathMap, enableModule)
		if err != nil {
			return nil, err
		}
		for path, ref := range imported {
			refs[path] = ref
		}
	}
	return refs, nil
}

func (imp *Importer) doImportPackagesOrError(pathMap map[string]PackageName, enableModule bool) (map[string]*PackageRef, error) {
	p.GetImportsSrcDir() // warns if GOPATH or paths.ImportsDir may be wrong

	paths := pathKeys(pathMap)
	mode := imp.importModeFor(pathMap)
	o := imp.output
	dir := computeImportDir(o, paths, mode)

	// o.Debugf("compiling plugin in directory %q...", dir)

	// loads names and types, not the values!
	pkginfos, err := imp.Load(dir, paths, mode, enableModule)
	if err != nil {
		return nil, imp.wrapImportError(paths, enableModule, err)
	}
	ok := createImportFiles(imp.output, dir, pkginfos, mode, enableModule)
	refs := make(map[string]*PackageRef, len(paths))

	if !ok || mode != ImPlugin {
		// either the packages export nothing, or user must rebuild gomacro.
		// in both cases, still cache them to avoid recreating the files.
		for _, path := range paths {
			ref := &PackageRef{Path: path}
			refs[path] = ref
			imports.Packages[path] = ref.Package
		}
		return refs, nil
	}
	soname := compilePlugin(imp.BuildArgs, o, dir, enableModule, o.Stdout, o.Stderr)
	ipkgs := imp.loadPluginSymbol(soname, "Packages")
	pkgs := *ipkgs.(*map[string]imports.PackageUnderlying)

	// o.Debugf("plugin %q loaded, contains definitions for packages %v", soname, pkgs)

	// cache *all* packages found for future use
	imports.Packages.Merge(pkgs)

	// but return only requested ones
	for _, pkgpath := range paths {
		pkg, found := imports.Packages[pkgpath]
		if !found {
			return nil, imp.output.MakeRuntimeError(
				"error loading package %q: the compiled plugin %q does not contain it! internal error? %v",
				pkgpath, soname)
		}
		refs[pkgpath] = &PackageRef{Package: pkg, Path: pkgpath}
	}
	return refs, nil
}

func clone(pathMap map[string]PackageName) map[string]PackageName {
	ret := make(map[string]PackageName, len(pathMap))
	for k, v := range pathMap {
		ret[k] = v
	}
	return ret
}

func (imp *Importer) importModeFor(pathMap map[string]PackageName) ImportMode {
	var curr ImportMode
	havemode := false
	for _, alias := range pathMap {
		var mode ImportMode
		switch alias {
		case "_b":
			mode = ImBuiltin
		case "_i":
			mode = ImInception
		case "_3":
			mode = ImThirdParty
		default:
			if imp.havePluginOpen() {
				mode = ImPlugin
			} else {
				mode = ImThirdParty
			}
		}
		highest := mode
		if havemode && mode != curr {
			if highest < curr {
				highest = curr
			}
			imp.output.Warnf("attempt to import multiple packages with conflicting modes %v and %v - using numerically higher mode %v",
				curr, mode, highest)
		}
		curr = highest
		havemode = true
	}
	return curr
}

func pathKeys(pkgpathMap map[string]PackageName) []string {
	ret := make([]string, len(pkgpathMap))
	index := 0
	for pkgpath := range pkgpathMap {
		ret[index] = pkgpath
		index++
	}
	return ret
}

func createImportFiles(o *Output, dir string, pkginfos map[string]*types.Package, mode ImportMode, enableModule bool) bool {
	if mode == ImPlugin {
		createDir(o, dir)
		removeAllFilesInDirExcept(o, dir, []string{"go.mod", "go.sum"})
	}
	buf := bytes.Buffer{}
	index := 0
	allEmpty := true
	if mode == ImPlugin {
		filepath, err := createPluginMainFile(dir)
		if err != nil {
			o.Errorf("error writing file %q: %v", filepath, err)
		}
	}
	// path can be a Go package path or an absolute filesystem path
	for path, pkginfo := range pkginfos {
		filepath := p.Subdir(dir, computeImportFilename(o, path, mode, index))

		isEmpty := createImportFile(o, &buf, path, pkginfo, mode)
		if isEmpty {
			o.Warnf("package %q exports zero constants, functions, types and variables", path)
		} else {
			allEmpty = false
		}
		err := ioutil.WriteFile(filepath, buf.Bytes(), os.FileMode(0644))
		buf.Reset()
		if err != nil {
			o.Errorf("error writing file %q: %v", filepath, err)
		}
		index++
	}
	if allEmpty {
		return false
	}

	switch mode {
	case ImBuiltin, ImThirdParty:
		o.Warnf("created files in %q, recompile gomacro to use it", dir)
	case ImInception:
		o.Warnf("created files in %q, recompile such package to use it", dir)
	case ImPlugin:
		// if needed, go.mod file was created already by Importer.Load()
		env := environForCompiler(enableModule)
		runGoModTidyIfNeeded(o, dir, env)
	}
	return true
}

func createDir(o *Output, dir string) {
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		o.Errorf("error creating directory %q: %v", dir, err)
	}
}

func listDir(o *Output, dir string) []os.FileInfo {
	d, err := os.Open(dir)
	if err != nil {
		o.Errorf("error opening directory %q: %v", dir, err)
	}
	defer d.Close()
	infos, err := d.Readdir(0)
	if err != nil {
		o.Errorf("error listing directory %q: %v", dir, err)
	}
	return infos
}

func removeAllFilesInDir(o *Output, dir string) {
	removeAllFilesInDirExcept(o, dir, nil)
}

func removeAllFilesInDirExcept(o *Output, dir string, exceptList []string) {
	for _, info := range listDir(o, dir) {
		if info.IsDir() {
			continue
		}
		name := info.Name()
		for _, exceptName := range exceptList {
			if name == exceptName {
				name = ""
				break
			}
		}
		if name == "" {
			continue
		}
		filepath := p.Subdir(dir, name)
		if err := os.Remove(filepath); err != nil {
			o.Errorf("error removing file %q: %v", filepath, err)
		}
	}
}

func createPluginGoModFile(o *Output, dir string, goModReplaceDirective map[PackagePath]AbsolutePath) string {
	var buf bytes.Buffer
	buf.WriteString("module gomacro.imports/")
	buf.WriteString(p.FileName(dir))
	buf.WriteRune('\n')
	buf.WriteRune('\n')

	for key, value := range goModReplaceDirective {
		buf.WriteString("replace ")
		buf.WriteString(key.String())
		buf.WriteString(" => ")
		buf.WriteString(value.String())
		buf.WriteRune('\n')
	}
	goModPath := p.Subdir(dir, "go.mod")
	err := ioutil.WriteFile(goModPath, buf.Bytes(), os.FileMode(0644))
	if err != nil {
		o.Errorf("error writing file %q: %v", goModPath, err)
	}
	return goModPath
}

func packageSanitizedName(pkgpath string) string {
	return sanitizeIdent(p.FileName(pkgpath))
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

var removeOnceGomacroImports sync.Once

func computeImportDir(o *Output, paths []string, mode ImportMode) string {
	switch mode {
	case ImBuiltin:
		// user will need to recompile gomacro
		return p.GetImportsSrcDir()
	case ImThirdParty:
		// either plugin.Open is not available, or user explicitly requested import _3 "package".
		// In both cases, user will need to recompile gomacro
		return p.Subdir(p.GetImportsSrcDir(), "thirdparty")
	case ImInception:
		if len(paths) != 1 {
			o.Errorf("import in Inception mode only supports a single package at time - attempted to import %d packages instead",
				len(paths))
		}
		for _, path := range paths {
			// user will need to recompile the package being imported
			if isLocalFilesystemPath(path) {
				return MakeAbsolutePathOrPanic(path).String()
			}
			for _, srcdir := range p.GoSrcDirs {
				dir := p.Subdir(srcdir, path)
				if _, err := os.Stat(dir); err == nil {
					return p.Subdir(srcdir, path)
				}
			}
			o.Errorf("unable to locate package %q in $GOPATH/src ($GOPATH=%s)",
				path, build.Default.GOPATH)
		}
	case ImPlugin:
		{
			baseDir := srcDirForPluginImport()
			toRemoveDir := p.DirName(baseDir)
			// Go has not atexit(), so remove directory $GOPATH/src/gomacro.imports once,
			// before the first attempt to import packages that requires compiling a plugin
			if p.FileName(toRemoveDir) == "gomacro.imports" {
				removeOnceGomacroImports.Do(func() {
					_ = os.RemoveAll(toRemoveDir)
				})
			}
			return p.Subdir(baseDir, fmt.Sprintf("import_%d", nextImportCounter()))
		}
	default:
		o.Errorf("unknown import mode: %v", mode)
	}
	return ""
}

func computeImportFilename(o *Output, pkgpath string, mode ImportMode, index int) string {
	switch mode {
	case ImBuiltin:
		// user will need to recompile gomacro
		return sanitizeIdent(pkgpath) + ".go"
	case ImThirdParty:
		// either plugin.Open is not available, or user explicitly requested import _3 "package".
		// In both cases, user will need to recompile gomacro
		return sanitizeIdent(pkgpath) + ".go"
	case ImInception:
		// user will need to recompile package being imported
		return "x_package.go"
	case ImPlugin:
		// avoid collision if multiple packages imported at once have the same sanitizeIdent(pkgpath)
		// and avoid file names starting with '_' because go build would ignore them
		return fmt.Sprintf("x_%s_%d.go", sanitizeIdent(pkgpath), index)
	default:
		o.Errorf("unknown import mode: %v", mode)
		return ""
	}
}

var importCounter uintptr

func nextImportCounter() uintptr {
	return atomic.AddUintptr(&importCounter, 1)
}
