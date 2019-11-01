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
 * canary.go
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

package hide_from_stack

import (
	"fmt"
	"runtime"
	"strings"
)

func canary(arg uintptr) {
	fmt.Printf("canary(%d) called", arg)
	pcs := make([]uintptr, 3)
	n := runtime.Callers(2, pcs)
	if n > 0 {
		frames := runtime.CallersFrames(pcs)
		for _, pc := range pcs {
			frame, more := frames.Next()
			if frame.PC != 0 {
				function := frame.Function[1+strings.LastIndexByte(frame.Function, '.'):]
				fmt.Printf(" by %s()%s", function, "               "[(len(function)+15)&15:])
			} else {
				fmt.Printf(" by 0x%08x\t", pc)
			}
			if !more {
				break
			}
		}
	}
	fmt.Println()
	runtime.GC()
}

// return a closure
func make_parrot(arg0 uintptr) func(uintptr) {
	return func(arg1 uintptr) {
		fmt.Printf("parrot(%d) called, closure data = %d\n", arg1, arg0)
		// debug.PrintStack()
		runtime.GC()
	}
}

func address_of_canary() func(uintptr) {
	return canary
}

type Env struct {
	closure func(uintptr)
	arg     uintptr
	call    [7]uintptr // call0, call16 ... call512
}

func MakeCallArray() [7]uintptr {
	var ret [7]uintptr
	for i, call := range [...]func(){
		call0, call16, call32, call64, call128, call256, call512,
	} {
		ret[i] = deconstruct_func0(call).funcAddress
	}
	return ret
}

func asm_address_of_canary() func(uintptr)
func asm_call_canary(arg uintptr)
func asm_call_func(func_address uintptr, arg uintptr)
func asm_call_closure(tocall func(uintptr), arg uintptr)
func asm_loop()
func asm_hideme(env *Env)
