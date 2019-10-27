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
 * func.go
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

package hide_from_stack

import (
	"unsafe"
)

func func_to_uintptr(closure interface{}) uintptr {
	type interfaceHeader struct {
		typ uintptr
		val uintptr
	}
	return (*interfaceHeader)(unsafe.Pointer(&closure)).val
}

func empty_func_to_uintptr(closure func()) uintptr {
	return *(*uintptr)(unsafe.Pointer(&closure))
}
