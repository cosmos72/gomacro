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
 * main.go
 *
 *  Created on: Nov 23, 2019
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/classic"
)

func main() {
	ir := classic.New()
	g := ir.Globals
	g.Options |= OptTrapPanic | OptShowPrompt | OptShowEval | OptShowEvalType
	ir.ReplStdin()
}
