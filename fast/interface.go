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

// ToInterface wraps a reflect.Value of type 'rtin' into a proxy struct of type 'rtproxy'
// that implements the interface type 'rtout'
// and returns the proxy value, converted to the interface type 'rtout'
func ToInterface(vin r.Value, tin xr.Type, rtproxy r.Type, rtout r.Type) r.Value {
	vaddrproxy := r.New(rtproxy)
	vproxy := vaddrproxy.Elem()
	vproxy.Field(0).Set(r.ValueOf(xr.MakeInterfaceHeader(vin, tin)))
	n := rtout.NumMethod()
	for i := 0; i < n; i++ {
		imtd := rtout.Method(i)
		xmtd, count := tin.MethodByName(imtd.Name, imtd.PkgPath)
		switch count {
		case 0:
			base.Errorf("cannot convert type <%v> to interface <%v>: missing method %s %s", tin, rtout, imtd.PkgPath, imtd.Name)
		case 1:
			/*
				if len(xmtd.FieldIndex) != 0 {
					Errorf("unimplemented: conversion of type <%v> to interface <%v> requires wrapping method %s %s from embedded field",
						tin, rtout, imtd.PkgPath, imtd.Name)
				}
			*/
			setProxyMethod(vproxy.Field(i+1), (*xmtd.Funs)[xmtd.Index])
		default:
			base.Errorf("cannot convert type <%v> to interface <%v>: ambiguous method %s %s", tin, rtout, imtd.PkgPath, imtd.Name)
		}
	}
	return vaddrproxy.Convert(rtout)
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
