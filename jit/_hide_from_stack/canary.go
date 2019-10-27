/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * canary.go
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

package hide_from_stack

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

var canary_print_stacktrace = false

func canary() {
	fmt.Println("canary called")
	if canary_print_stacktrace {
		debug.PrintStack()
	}
	runtime.GC()
}

// used by asm_address_of_canary()
var var_canary = canary

func address_of_canary() func() {
	return canary
}

func asm_address_of_canary() func()
func asm_call_canary()
func asm_call_func(tocall uintptr)
func asm_call_closure(tocall func())
func asm_loop()
