// +build !go1.16

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
 * importer_load_go1_15.go
 *
 *  Created on Mar 16, 2021
 *      Author Massimiliano Ghilardi
 */

package genimport

// Go < 1.16 does not require to run "go get ..." before "go list ..."
func runGoGetIfNeeded(output *Output, pkgpath string, dir string, env []string) error {
	_ = dir
	_ = env
	output.Debugf("looking for package %q ...", pkgpath)
	return nil
}

// Go < 1.16 does not require to run "go mod tidy" before "go build ..."
func runGoModTidyIfNeeded(output *Output, pkgpath string, dir string, env []string) error {
	return nil
}
