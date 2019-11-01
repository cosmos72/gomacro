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
 * func_test.go
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

package hide_from_stack

import (
	"fmt"
	"testing"
)

/*
func TestAsmLoop(t *testing.T) {
	go asm_loop() // causes runtime.GC() to wait indefinitely
	runtime.GC()
}
*/

func TestAddressOfCanary(t *testing.T) {
	if false {
		fmt.Printf("canary                       = %p\n", canary)
		fmt.Printf("address_of_canary()          = %p\n", address_of_canary())
		fmt.Printf("asm_address_of_canary()      = %p\n", asm_address_of_canary())
		header := deconstruct_func8(canary)
		fmt.Printf("deconstruct_func8(canary)    = %#x\n", *header)
		header = deconstruct_any_func(canary)
		fmt.Printf("deconstruct_any_func(canary) = %#x\n", *header)
	}
}

func TestCallCanary(t *testing.T) {
	asm_call_canary(0)
	asm_call_func(deconstruct_any_func(canary).funcAddress, 1)
	asm_call_closure(asm_address_of_canary(), 2)
	grow_stack()
	env := &Env{canary, 3, MakeCallArray()}
	asm_hideme(env)
	if jit_hideme != nil {
		env.arg++
		jit_hideme(env)
	}
}

func _TestCallParrot(t *testing.T) {
	parrot := make_parrot(123456)
	asm_call_closure(parrot, 0)
	env := &Env{parrot, 1, MakeCallArray()}
	asm_hideme(env)
	if jit_hideme != nil {
		env.arg++
		jit_hideme(env)
	}
}
