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
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package asm

import (
	"testing"
)

func MakeCode(instr ...uint32) Code {
	code := make(Code, len(instr)*4)
	for i, inst := range instr {
		code[4*i+0] = byte(inst >> 0)
		code[4*i+1] = byte(inst >> 8)
		code[4*i+2] = byte(inst >> 16)
		code[4*i+3] = byte(inst >> 24)
	}
	return code
}

func SameCode(actual Code, expected Code) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i := range actual {
		if actual[i] != expected[i] {
			return false
		}
	}
	return true
}

func EqUint8(t *testing.T, actual uint8, expected uint8) {
	if actual != expected {
		t.Errorf("expected %d,\tactual %d", expected, actual)
	}
}

func TestLog2(t *testing.T) {
	for shift := uint8(1); shift < 64; shift++ {
		n := uint64(1) << shift
		actual, _ := log2uint(n)
		EqUint8(t, actual, shift)
		actual, _ = log2uint(n - 1)
		EqUint8(t, actual, 0)
		actual, _ = log2uint(n + 1)
		EqUint8(t, actual, 0)
	}
}
