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
 * call the closure stored in BX (which expects exactly 0 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_func()
 */
func call0()

/**
 * call the closure stored in BX (which expects up to 8 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_func()
 *
 * writing arguments in call8() stack and retrieving return values from it
 * are caller's responsibility
 * (possible only in assembly: caller has to access call8() stack)
 */
func call8()

// as above, but closure can expect up to 16 bytes of arguments return values
func call16()

// as above, but closure can expect up to 32 bytes of arguments return values
func call32()

// as above, but closure can expect up to 64 bytes of arguments return values
func call64()

// as above, but closure can expect up to 128 bytes of arguments return values
func call128()

// as above, but closure can expect up to 256 bytes of arguments return values
func call256()

// as above, but closure can expect up to 512 bytes of arguments return values
func call512()

// as above, but closure can expect up to 1024 bytes of arguments return values
func call1024()

// ensure stack has > 1024 free bytes
func growStack()
