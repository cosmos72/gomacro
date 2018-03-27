/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * interface.go
 *
 *  Created on: Mar 29, 2017
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) TypeInterface(node *ast.InterfaceType) xr.Type {
	if node.Methods == nil || len(node.Methods.List) == 0 {
		return c.TypeOfInterface()
	}
	methodtypes, methodnames := c.TypeFields(node.Methods)

	// TODO embedded interfaces
	return c.Universe.InterfaceOf(methodnames, methodtypes, nil)
}

// InterfaceProxy returns the proxy struct that implements a compiled interface
func (c *Comp) InterfaceProxy(t xr.Type) r.Type {
	ret := c.interf2proxy[t.ReflectType()]
	if ret == nil {
		c.Errorf("internal error: proxy not found for %s type <%v>", t.Kind(), t)
	}
	return ret
}

// converterToInterface compiles a conversion from 'tin' into a proxy struct that implements the interface type 'tout'
// and returns a function that performs such conversion
func (c *Comp) converterToInterface(tin xr.Type, tout xr.Type) func(val r.Value, rtout r.Type) r.Value {
	rtproxy := c.InterfaceProxy(tout)
	rtout := tout.ReflectType()

	vtable := r.New(rtproxy).Elem()
	n := rtout.NumMethod()
	for i := 0; i < n; i++ {
		imtd := rtout.Method(i)
		xmtd, count := tin.MethodByName(imtd.Name, imtd.PkgPath)
		if count == 0 {
			c.Errorf("cannot convert type <%v> to interface <%v>: missing method %s %s", tin, rtout, imtd.PkgPath, imtd.Name)
		} else if count > 1 {
			c.Errorf("type <%v> has %d wrapper methods %s %s all at the same depth=%d - cannot convert to interface <%v>",
				tin, count, imtd.PkgPath, imtd.Name, len(xmtd.FieldIndex), tout)
		}
		e := c.compileMethodAsFunc(tin, xmtd)
		setProxyMethod(vtable.Field(i+1), r.ValueOf(e.Value))
	}
	return func(val r.Value, rtout r.Type) r.Value {
		vaddr := r.New(rtproxy)
		vproxy := vaddr.Elem()
		vproxy.Set(vtable)
		vproxy.Field(0).Set(r.ValueOf(xr.MakeInterfaceHeader(val, tin)))
		return vaddr.Convert(rtout)
	}
}

func setProxyMethod(place r.Value, mtd r.Value) {
	rtin := mtd.Type()
	rtout := place.Type()
	if rtin == rtout {
		place.Set(mtd)
	} else if rtin.ConvertibleTo(rtout) {
		place.Set(mtd.Convert(rtout))
	} else {
		place.Set(r.MakeFunc(rtout, func(args []r.Value) []r.Value {
			args[0] = args[0].Interface().(xr.InterfaceHeader).Value()
			return mtd.Call(args)
		}))
	}
}

// extract a value from a proxy struct (one of the imports.* structs) that implements an interface
// this is the inverse of the function returned by Comp.converterToInterface() above
func (g *CompGlobals) extractFromInterface(v r.Value) (r.Value, xr.Type) {
	// base.Debugf("type assertion: value = %v <%v>", v, base.ValueType(v))
	i := base.ValueInterface(v)
	v = r.ValueOf(i) // rebuild with concrete type
	rt := r.TypeOf(i)
	var xt xr.Type
	// base.Debugf("type assertion: concrete value = %v <%v>", i, t)
	if rt != nil && rt.Kind() == r.Ptr && g.proxy2interf[rt.Elem()] != nil {
		v = v.Elem().Field(0)
		i = base.ValueInterface(v)
		if j, ok := i.(xr.InterfaceHeader); ok {
			// base.Debugf("type assertion: unwrapped value = %v <%T>", j, j)
			v = j.Value()
			xt = j.Type()
		} else {
			// base.Debugf("type assertion: failed to unwrap value = %v <%T>", i, i)
			v = r.ValueOf(i) // rebuild with concrete type
		}
	}
	return v, xt
}
