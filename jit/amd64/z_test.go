/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * z_test.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

import (
	_ "fmt"
	"math/rand"
	_ "runtime"
	"testing"
	_ "unsafe"
)

const verbose = false

func TestLoadStore(t *testing.T) {
	var asm Asm
	v := NewVar(0)
	ints := [1]uint64{0}
	for _, reg := range [...]Reg{AX, CX, DX, BX /*BP,SP*/, SI /*DI,*/, R8, R9, R10, R11, R12, R13, R14, R15} {
		val := int64(rand.Uint64())
		f := asm.Init().LoadConst(reg, val).Store(v, reg).Func()
		f(&ints[0])
		actual := int64(ints[0])
		if actual != val {
			t.Errorf("LoadConst+Store returned %d, expecting %d", actual, val)
		}
	}
}

func TestSum(t *testing.T) {
	const (
		n        = 10
		expected = n * (n + 1) / 2
	)
	f := DeclSum()

	actual := f(n)
	if actual != expected {
		t.Errorf("sum(%v) returned %v, expecting %d", n, actual, expected)
	} else if verbose {
		t.Logf("sum(%v) = %v\n", n, actual)
	}
}

/*
func TestAdd(t *testing.T) {
	var asm Asm
	v1, v2, v3 := NewVar(1), NewVar(2), NewVar(3)
	f := asm.Init().Set(v1, v2).Add(v1, v3).Func()
	code := asm.Code
	runtime.GC()

	mem := **(**[]uint8)(unsafe.Pointer(&f))
	if verbose {
		fmt.Printf("f    = %p\n", f)
		fmt.Printf("addr = %p\n", mem)
		fmt.Printf("mem  = %v\n", mem)
		fmt.Printf("code = %v\n", code)
	}
	const (
		a = 7
		b = 11
		c = a + b
	)

	ints := [...]uint64{2: a, 3: b}
	f(&ints[0])
	if ints[1] != c {
		t.Errorf("Add returned %v, expecting %d", ints[1], c)
	} else if verbose {
		t.Logf("ints = %v\n", ints)
	}

}

func TestArith(t *testing.T) {
	const (
		n        int = 9
		expected int = ((((n*2 + 3) | 4) &^ 5) ^ 6) / ((n & 2) | 1)
	)
	env := [5]uint64{uint64(n), 0, 0}
	f := DeclArith(len(env))

	f(&env[0])
	actual := int(env[1])
	if actual != expected {
		t.Errorf("arith(%d) returned %d, expecting %d", n, actual, expected)
	} else if verbose {
		t.Logf("arith(%d) = %d\n", n, actual)
	}
}
*/
