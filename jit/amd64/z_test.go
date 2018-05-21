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
	f := asm.Init().AddInt64(Bind{Idx: 1}, Bind{Idx: 2}, Bind{Idx: 3}).Func()
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
	fsum := DeclSum()

	actual := fsum(n)
	if actual != expected {
		t.Errorf("sum(%v) returned %v, expecting %d", n, actual, expected)
	} else if verbose {
		t.Logf("sum(%v) = %v\n", n, actual)
	}
}
