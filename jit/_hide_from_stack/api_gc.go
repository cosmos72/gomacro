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
 * call the closure stored in DX (which expects exactly 0 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * on AMD64, caller is replaced with a fake entry hidden_jit_func()
 * on ARM64, caller completely disappears from stack trace
 */
func call0()

/**
 * call the closure stored in DX, hiding the caller from runtime stack:
 * on AMD64, caller is replaced with a fake entry hidden_jit_func()
 * on ARM64, caller completely disappears from stack trace

 * the closure can expect up to 16 bytes of arguments + return values
 *
 * writing arguments in call16() stack and retrieving return values from it
 * are caller's responsibility
 * (possible only in assembly: caller has to access call16() stack)
 */
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

// ensure stack has > 512 free bytes
func grow_stack()

// on AMD64, hidden JIT functions will be replaced by this function in the stack trace
func hidden_jit_func(uintptr)
