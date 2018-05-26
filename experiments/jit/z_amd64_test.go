// +build amd64

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
 * z_test_amd64.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"math/rand"
	"testing"
)

func TestLoadStoreAmd64(t *testing.T) {
	var asm Asm
	v := NewVar(0)
	ints := [1]uint64{0}
	for r := rLo; r <= rHi; r++ {
		if r == rBP || r == rSP || r == rDI {
			continue
		}
		val := int64(rand.Uint64())
		f := asm.Init().loadConst(r, val).storeReg(v, r).Func()
		f(&ints[0])
		actual := int64(ints[0])
		if actual != val {
			t.Errorf("LoadConst+Store returned %d, expecting %d", actual, val)
		}
	}
}
