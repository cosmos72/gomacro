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
	for _, opname := range [...]string{"mov", "neg", "mvn"} {
		f, err := os.Create("_gen_" + opname + ".s")
		if err != nil {
			panic(err)
		}
		g := newGenOp2(f, opname)
		g.generate()
		f.Close()
	}

	for _, opname := range [...]string{"adc", "add", "sub", "sbc", "mul", "and", "orr", "eor", "lsl", "lsr", "asr"} {
		f, err := os.Create("_gen_" + opname + ".s")
		if err != nil {
			panic(err)
		}
		g := newGenOp3(f, opname)
		g.generate()
		f.Close()
	}
}
