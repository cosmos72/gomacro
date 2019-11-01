// +build !gc !amd64,!arm64

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
 * hideme_stub.go
 *
 *  Created on Nov 01, 2019
 *      Author Massimiliano Ghilardi
 */

package hide_from_stack

func asm_hideme(*Env) {
	unimplemented()
}
