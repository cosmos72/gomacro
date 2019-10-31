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
 * canary_jit_test.go
 *
 *  Created on Oct 31, 2019
 *      Author Massimiliano Ghilardi
 */

package hide_from_stack

import (
	"unsafe"

	"github.com/cosmos72/gomacro/jit/asm"
)

var jit_hideme func(*Env) = make_jit_hideme()

var jit_fails func(*Env) = make_jit_fails()

func make_jit_hideme() func(*Env) {
	if !asm.ARCH_SUPPORTED {
		return nil
	}
	var ret func(*Env)
	addr := (*[128]byte)(unsafe.Pointer(deconstruct_any_func(asm_hideme).funcAddress))[:]
	asm.New().Bytes(addr...).Func(&ret)
	return ret
}

// if the same assembly as this function is executed from runtime-allocated memory,
// it does not have a stackmap thus calling (almost all) Go functions crashes
func fails(env *Env) {
	env.closure(env.arg)
}

func make_jit_fails() func(*Env) {
	if !asm.ARCH_SUPPORTED {
		return nil
	}
	var ret func(*Env)
	addr := (*[128]byte)(unsafe.Pointer(deconstruct_any_func(fails).funcAddress))[:]
	asm.New().Bytes(addr...).Func(&ret)
	return ret
}
