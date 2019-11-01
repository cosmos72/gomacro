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
 * jit_test.go
 *
 *  Created on Oct 31, 2019
 *      Author Massimiliano Ghilardi
 */

package hide_from_stack

import (
	"unsafe"

	"github.com/cosmos72/gomacro/jit/asm"
)

type Env struct {
	closure func(uintptr)
	arg     uintptr
	call    [7]uintptr // call0, call16 ... call512
}

var jit_hideme func(*Env) = make_jit(asm_hideme)

var jit_fails func(*Env) = make_jit(call_closure)

// if the same assembly as this function is executed from runtime-allocated memory,
// it does not have a stackmap thus calling (almost all) Go functions crashes the garbage collector
func call_closure(env *Env) {
	env.closure(env.arg)
}

// return a copy of fun allocated in executable mmap'ed memory,
// as a JIT compiler would do.
// fun must occupy 128 bytes or less.
func make_jit(fun func(*Env)) func(*Env) {
	if !asm.ARCH_SUPPORTED {
		return nil
	}
	var ret func(*Env)
	addr := (*[128]byte)(unsafe.Pointer(deconstruct_any_func(fun).funcAddress))[:]
	asm.New().Bytes(addr...).Func(&ret)
	return ret
}
