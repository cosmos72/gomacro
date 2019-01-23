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
 * const.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
	"reflect"
)

func (c Const) String() string {
	return fmt.Sprintf("0x%x/*%v*/", c.kind, c.val)
}

// implement Arg interface
func (c Const) Reg() Reg {
	return NoReg
}

func (c Const) Kind() reflect.Kind {
	return c.kind
}

func (c Const) Const() bool {
	return true
}

func Int64(val int64) Const {
	return Const{val: val, kind: reflect.Int64}
}
