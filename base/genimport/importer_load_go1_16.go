// +build go1.16

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
 * importer_load_go1_16.go
 *
 *  Created on Mar 16, 2021
 *      Author Massimiliano Ghilardi
 */

package genimport

import (
	"fmt"
	"os/exec"
)

// if the package and the version to import is NOT listed in go.mod,
// Go >= 1.16 requires to run "go get pkg/to/be/imported" or "go install ..."
// before "go list ..." in order to update go.mod
// We cannot know the version beforehand, so we always run "go get ..."
func runGoGetIfNeeded(output *Output, pkgpath string, dir string, env []string) error {

	output.Debugf("downloading package %q ...", pkgpath)

	gocmd := chooseGoCmd()

	cmd := exec.Command(gocmd, "get", pkgpath)
	cmd.Dir = dir
	cmd.Env = env
	cmd.Stdin = nil
	cmd.Stdout = output.Stdout
	cmd.Stderr = output.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing \"%s get %s\" in directory %q: %v", gocmd, pkgpath, dir, err)
	}
	return nil
}
