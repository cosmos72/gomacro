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

type closureHeader struct {
	funcAddress uintptr
	closureData [0]uintptr
}

func deconstruct_any_func(fun interface{}) *closureHeader {
	type interfaceHeader struct {
		typ uintptr
		val *closureHeader
	}
	return (*interfaceHeader)(unsafe.Pointer(&fun)).val
}

func deconstruct_func0(fun func()) *closureHeader {
	return *(**closureHeader)(unsafe.Pointer(&fun))
}

func deconstruct_func8(fun func(uintptr)) *closureHeader {
	return *(**closureHeader)(unsafe.Pointer(&fun))
}
