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
 * file.go
 *
 *  Created on: Feb 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"bufio"
	"go/ast"
	"os"

	. "github.com/cosmos72/gomacro/base"
)

func (ir *Interp) EvalFile(filePath string, pkgPath string) {
	file, err := os.Open(filePath)
	if err != nil {
		ir.Errorf("error opening file '%s': %v", filePath, err)
		return
	}
	defer file.Close()

	savePath := ir.Env.ThreadGlobals.PackagePath
	saveOpts := ir.Env.Options

	ir.ChangePackage(pkgPath)
	ir.Env.Options &^= OptShowEval

	defer func() {
		ir.ChangePackage(savePath)
		ir.Env.Options = saveOpts
	}()

	in := bufio.NewReader(file)
	ir.Repl(in)
}

func (env *Env) evalFile(node *ast.File) {
	env.Name = node.Name.Name
	env.Path = env.Name
	env.PackagePath = env.Name

	for _, imp := range node.Imports {
		env.evalImport(imp)
	}

	for _, decl := range node.Decls {
		env.evalDecl(decl)
	}
}
