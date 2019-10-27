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
 * hide.go
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */
package hide_from_stack

type Env struct {
	tocall func()
}

func hideme(env *Env)

func hidden_func(env *Env) {
	panic("hidden_func called! this should not happen")
}
