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

	"github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base/reflect"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (ir *Interp) Inspect(src string) {
	c := ir.Comp
	g := &c.Globals
	inspector := g.Inspector
	if inspector == nil {
		c.Errorf("no inspector set: call Interp.SetInspector() first")
		return
	}
	form := c.Parse(src)
	if _, ok := form.(ast2.AstWithSlice); ok && form.Size() == 1 {
		form = form.Get(0)
	}
	expr, xtyp := c.Expr1OrType(ast2.ToExpr(form))
	var val xr.Value
	if expr != nil {
		val, xtyp = ir.RunExpr1(expr)
	} else {
		// attempt to inspect a type: inspect the zero value of the type
		val = xr.Zero(xtyp)
	}
	typ := xtyp.ReflectType()
	if val.IsValid() && val.Kind() == r.Interface {
		// extract concrete type
		val = val.Elem()
		typ = reflect.ValueType(val)
	}
	inspector.Inspect(src, val.ReflectValue(), typ, xtyp, &ir.Comp.Globals)
}
