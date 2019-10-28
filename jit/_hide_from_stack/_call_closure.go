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
 * call_closure.go
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

package hide_from_stack

// go:nosplit
func call00_closure(f func()) {
	f()
}

// go:nosplit
func call10_closure(f func(uintptr)) {
	f(0)
}

// go:nosplit
func call01_closure(f func() uintptr) uintptr {
	return f()
}
