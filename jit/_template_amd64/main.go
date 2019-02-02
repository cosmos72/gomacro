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
 * main.go
 *
 *  Created on Jan 28, 2019
 *      Author Massimiliano Ghilardi
 */

package main

import "os"

func main() {
	for _, opname := range [...]string{"inc", "dec", "neg", "not"} {
		f, err := os.Create("_gen_" + opname + ".s")
		if err != nil {
			panic(err)
		}
		g := NewGenOp1(f, opname)
		g.generate()
		f.Close()
	}

	for _, opname := range [...]string{"add", "and", "mov", "or", "sub", "xchg", "xor"} {
		f, err := os.Create("_gen_" + opname + ".s")
		if err != nil {
			panic(err)
		}
		g := NewGenOp2(f, opname)
		g.generate()
		f.Close()
	}
}
