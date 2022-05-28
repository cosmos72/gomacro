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
 * declaration.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"go/ast"
	r "reflect"
	"strings"

	"github.com/cosmos72/gomacro/base/genimport"

	"github.com/cosmos72/gomacro/base"
	bstrings "github.com/cosmos72/gomacro/base/strings"
)

type PackageName = genimport.PackageName

// eval a single import
func (env *Env) evalImportDecl(decl ast.Spec) (r.Value, []r.Value) {
	switch node := decl.(type) {
	case *ast.ImportSpec:
		return env.evalImport(node)
	default:
		return env.Errorf("unimplemented import: %v", decl)
	}
}

// eval a single import
func (env *Env) evalImport(imp *ast.ImportSpec) (r.Value, []r.Value) {
	path := bstrings.UnescapeString(imp.Path.Value)
	path = env.sanitizeImportPath(path)
	var name PackageName
	if imp.Name != nil {
		name = PackageName(imp.Name.Name)
	}
	pkg := env.Globals.Importer.ImportPackage(name, path, env.Globals.Options&base.OptModuleImport != 0)
	if pkg != nil {
		// if import appears *inside* a block, it is local for that block
		if name == "." {
			// dot import, i.e. import . "the/package/path"
			env.MergePackage(pkg.Package)
		} else {
			// https://golang.org/ref/spec#Package_clause states:
			// If the PackageName is omitted, it defaults to the identifier
			// specified in the package clause of the imported package
			if len(name) == 0 {
				name = pkg.DefaultName()
			}
			env.DefineConst(name.String(), r.TypeOf(pkg), r.ValueOf(pkg))
		}
	}
	return r.ValueOf(name.String()), nil
}

func (ir *ThreadGlobals) sanitizeImportPath(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	if genimport.IsLocalImportPath(path) {
		abspath, err := genimport.MakeAbsolutePathOrError(path)
		if err != nil {
			ir.Errorf("invalid import %q: conversion to absolute path failed: %v", path, err)
		}
		path = abspath.String()
	}
	l := len(path)
	if (l >= 3 && path[l-3:] == "/..") || strings.Contains(path, "/../") {
		ir.Errorf("invalid import %q: contains \"..\"", path)
	}
	if (l >= 2 && path[l-2:] == "/.") || strings.Contains(path, "/./") {
		ir.Errorf("invalid import %q: contains \".\"", path)
	}
	return path
}
