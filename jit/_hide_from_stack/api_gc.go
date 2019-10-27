// +build gc,amd64 gc,arm64

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
 * api_gc.go
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

package hide_from_stack

/**
 * call closure stored in register BX,
 * hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_func()
 */
func call_0b()
func call_8b(arg [1]uintptr)
func call_16b(arg [2]uintptr)
func call_32b(arg [4]uintptr)
func call_64b(arg [8]uintptr)
func call_128b(arg [16]uintptr)
func call_256b(arg [16]uintptr)
