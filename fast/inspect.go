/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * inspect.go
 *
 *  Created on: Apr 20, 2017
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
)

func (ir *Interp) Inspect(src string) {
	c := ir.Comp
	g := &c.Globals
	inspector := g.Inspector
	if inspector == nil {
		c.Errorf("no inspector set: call Interp.SetInspector() first")
		return
	}
	// not ir.Compile because it only macroexpands if OptMacroExpandOnly is set
	val, xtyp := ir.RunExpr1(c.Compile(c.Parse(src)))
	typ := xtyp.ReflectType()
	if val.IsValid() && val.Kind() == r.Interface {
		val = val.Elem() // extract concrete type
		typ = typ.Elem()
	}
	inspector.Inspect(src, val.ReflectValue(), typ, xtyp, &ir.Comp.Globals)
}
