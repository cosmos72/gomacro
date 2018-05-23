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
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

const verbose = false

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

func TestArithMem(t *testing.T) {
	testArith(t, "arith_mem", DeclArithMem())
}

func TestArithReg(t *testing.T) {
	testArith(t, "arith_reg", DeclArithReg())
}

func testArith(t *testing.T, name string, f func(*uint64)) {
	const (
		n        int = 9
		expected int = ((((n*2 + 3) | 4) &^ 5) ^ 6) / ((n & 2) | 1)
	)
	env := [3]uint64{uint64(n), 0, 0}
	f(&env[0])
	actual := int(env[1])
	if actual != expected {
		t.Errorf("%s(%v) returned %v, expecting %d", name, n, actual, expected)
	} else if verbose {
		t.Logf("%s(%v) = %v\n", name, n, actual)
	}
}
