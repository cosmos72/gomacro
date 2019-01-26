/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * output.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

func panicf(format string, args ...interface{}) *Asm {
	panic(fmt.Errorf("jit/amd64 internal error: "+format, args...))
	return nil
}

func errorf(format string, args ...interface{}) *Asm {
	panic(fmt.Errorf(format, args...))
	return nil
}
