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
	"os"
	"testing"
	"unsafe"
)

func TestAddressOfCanary(t *testing.T) {
	if os.Stdout == nil {
		fmt.Printf("address_of_canary                = %p\n\n", address_of_canary)
	}
	fmt.Printf("canary                           = %p\n", canary)
	fmt.Printf("address_of_canary()              = %p\n", address_of_canary())
	fmt.Printf("asm_address_of_canary()          = %p\n", asm_address_of_canary())
	u := empty_func_to_uintptr(canary)
	fmt.Printf("empty_func_to_uintptr(canary)[0] = %#x\n", *(*uintptr)(unsafe.Pointer(u)))
	u = func_to_uintptr(canary)
	fmt.Printf("func_to_uintptr(canary)[0]       = %#x\n", *(*uintptr)(unsafe.Pointer(u)))
}

func TestCallCanary(t *testing.T) {
	asm_call_canary()
	asm_call_closure(asm_address_of_canary())
	hideme(&Env{canary})
}
