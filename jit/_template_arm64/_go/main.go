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
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

package main

func main() {
	println(Add8, Add16, Add32, Add64, Zero,
		Div8, Div16, Div32, Div64,
		UDiv8, UDiv16, UDiv32, UDiv64)
}
