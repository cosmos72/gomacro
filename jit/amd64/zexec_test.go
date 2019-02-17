// +build amd64

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
 * z_exec_test.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package amd64

import (
	"fmt"
	"math/rand"
	"testing"
	"unsafe"
)

var verbose = false

func Init(asm *Asm) *Asm {
	asm.InitArch(Amd64{})
	asm.RegIncUse(RSI)
	asm.Load(MakeMem(8, RSP, Uint64), MakeReg(RSI, Uint64))
	return asm
}

func TestNop(t *testing.T) {
	var asm Asm
	var f func()
	asm.InitArch(Amd64{}).Func(&f)
	f()
}

func TestMov(t *testing.T) {
	var f func(*uint64)

	m := Var(0)
	binds := [...]uint64{0}
	var asm Asm
	for id := RLo; id <= RHi; id++ {
		Init(&asm)
		if asm.RegIsUsed(id) {
			continue
		}
		r := MakeReg(id, Int64)
		c := ConstInt64(int64(rand.Uint64()))
		asm.Mov(c, r).Mov(r, m).Func(&f)
		f(&binds[0])
		actual := int64(binds[0])
		if actual != c.Val() {
			t.Errorf("Mov returned %d, expecting %d", actual, c.Val())
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
  jit-compiled version of:

	func sum(n int) int {
		total := 0
		for i := 1; i <= n; i++ {
			total += i
		}
		return total
	}
*/
func DeclSum() func(arg int64) int64 {
	var init, next, loop func(*uint64)
	const n, total, i = 0, 1, 2
	_, Total, I := Var(n), Var(total), Var(i)

	var asm Asm
	Init(&asm).Mov(ConstInt64(1), I).Func(&init)
	pred := func(env *[3]uint64) bool {
		return int64(env[i]) <= int64(env[n])
	}
	Init(&asm).Op2(ADD, ConstInt64(1), I).Func(&next)
	Init(&asm).Op2(ADD, I, Total).Func(&loop)

	return func(arg int64) int64 {
		env := [3]uint64{n: uint64(arg)}

		for init(&env[0]); pred(&env); next(&env[0]) {
			loop(&env[0])
		}
		return int64(env[total])
	}
}

func TestAdd(t *testing.T) {
	var f func(*uint64)
	var asm Asm
	v1, v2, v3 := Var(0), Var(1), Var(2)

	for id := RLo; id <= RHi; id++ {
		Init(&asm)
		if asm.RegIsUsed(id) {
			continue
		}
		r := MakeReg(id, Int64)
		asm.Asm(MOV, v1, r, //
			NEG, r, //
			NOT, r, //
			ADD, v2, r, //
			NOT, r, //
			NEG, r, //
			MOV, r, v3, //
		).Func(&f)

		if verbose {
			code := asm.Code()
			mem := *(**[]uint8)(unsafe.Pointer(&f))
			fmt.Printf("f    = %p\n", f)
			fmt.Printf("addr = %p\n", mem)
			fmt.Printf("mem  = %v\n", *mem)
			fmt.Printf("code = %#v\n", code)
		}
		const (
			a = 7
			b = 11
			c = a + b
		)

		ints := [3]uint64{0: a, 1: b}
		f(&ints[0])
		if ints[2] != c {
			t.Errorf("Add returned %v, expecting %d", ints[1], c)
		} else if verbose {
			t.Logf("ints = %v\n", ints)
		}
	}
}

func TestCast(t *testing.T) {
	var f func(*uint64)
	var asm Asm
	Init(&asm)

	const n = uint64(0xEFCDAB8967452301)
	const hi = ^uint64(0)
	actual := [...]uint64{n, hi, hi, hi, hi, hi, hi}
	expected := [...]uint64{
		n,
		uint64(uint8(n & 0xFF)), uint64(uint16(n & 0xFFFF)), uint64(uint32(n & 0xFFFFFFFF)),
		uint64(int8(n & 0xFF)), uint64(int16(n & 0xFFFF)), uint64(int32(n & 0xFFFFFFFF)),
	}
	N := [...]Mem{
		VarK(0, Uint64),
		VarK(0, Uint8), VarK(0, Uint16), VarK(0, Uint32),
		VarK(0, Int8), VarK(0, Int16), VarK(0, Int32),
	}
	V := [...]Mem{
		VarK(0, Uint64),
		VarK(1, Uint64), VarK(2, Uint64), VarK(3, Uint64),
		VarK(4, Uint64), VarK(5, Uint64), VarK(6, Uint64),
	}
	r := asm.RegAlloc(Uint64)
	asm.Asm(
		CAST, N[1], V[1],
		CAST, N[2], V[2],
		CAST, N[3], V[3],
		CAST, N[4], V[4],
		CAST, N[5], V[5],
		CAST, N[6], V[6],
	).RegFree(r)
	asm.Func(&f)
	f(&actual[0])
	if actual != expected {
		t.Errorf("CAST returned %v, expecting %v", actual, expected)
	}
}

func TestLea(t *testing.T) {
	var f func(*uint64)
	const (
		n, m     int64 = 1020304, 9
		expected int64 = n * m
	)
	N := Var(0)
	env := [...]uint64{uint64(n)}

	var asm Asm
	Init(&asm).Asm(MUL, ConstInt64(m), N).Func(&f)
	f(&env[0])

	actual := int64(env[0])
	if actual != expected {
		t.Errorf("MUL %d 5 returned %d, expecting %d", n, actual, expected)
	} else if verbose {
		t.Logf("MUL %d 5 = %d\n", n, actual)
	}
}

/*
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
