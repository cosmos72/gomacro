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
 * api_stub.go
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

package hide_from_stack

// Return an array of function addresses {call0, call16 ... call512}
// If unsupported on this CPU or compiler, all values will be zero.
func MakeCallArray() [7]uintptr {
	return [7]uintptr{}
}

func call0() {
	unimplemented()
}
func call16() {
	unimplemented()
}
func call32() {
	unimplemented()
}
func call64() {
	unimplemented()
}
func call128() {
	unimplemented()
}
func call256() {
	unimplemented()
}

// ensure stack has > 512 free bytes
func GrowStack() {
	unimplemented()
}

func hidden_jit_func(uintptr) {
	unimplemented()
}

// go:noinline
func unimplemented() {
	panic("hide_from_stack: not implemented on this CPU and/or compiler")
}
