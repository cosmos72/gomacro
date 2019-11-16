// +build !gc !amd64,!arm64

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
 * misc_test.go
 *
 *  Created on Nov 01, 2019
 *      Author Massimiliano Ghilardi
 */

package hide_from_stack

func asm_address_of_canary() func(uintptr) {
	return canary
}

func asm_call_canary(arg uintptr) {
	canary(arg)
}

func asm_call_func(func_address uintptr, arg uintptr) {
	unimplemented()
}

func asm_call_closure(tocall func(uintptr), arg uintptr) {
	tocall(arg)
}

func asm_loop() {
	for {
	}
}
