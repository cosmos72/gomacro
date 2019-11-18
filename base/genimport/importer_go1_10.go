// +build !go1.11

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
 * importer_go1_10.go
 *
 *  Created on Nov 18, 2019
 *      Author Massimiliano Ghilardi
 */

package genimport

import (
	"go/build"
	"go/importer"
	"go/types"
	"os"
)

func (imp *Importer) Load(path string, enableModule bool) (p *types.Package, err error) {
	return importer.Default().Import(path)
}

const GoModuleSupported bool = false

func environForCompiler(enableModule bool) []string {
	return append(os.Environ(),
		"GOARCH="+build.Default.GOARCH,
		"GOOS="+build.Default.GOOS,
		"GOROOT="+build.Default.GOROOT)
}
