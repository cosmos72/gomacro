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
 * interface.go
 *
 *  Created on: Mar 29, 2017
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/base/reflect"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type genericInterfaceReceiverType struct{}

var genericInterfaceReceiverKey genericInterfaceReceiverType

// compile an interface definition
func (c *Comp) TypeInterface(node *ast.InterfaceType) xr.Type {
	if node.Methods == nil || len(node.Methods.List) == 0 {
		return c.TypeOfInterface()
	}
	types, names := c.TypeFields(node.Methods)

	// parser returns embedded interfaces as unnamed fields
	var methodnames []string
	var methodtypes, embeddedtypes []xr.Type
	for i, typ := range types {
		if i < len(names) && len(names[i]) != 0 {
			methodnames = append(methodnames, names[i])
			methodtypes = append(methodtypes, typ)
		} else {
			if typ.Kind() != r.Interface {
				c.Errorf("embedded interface is not an interface: %v", typ)
			}
			embeddedtypes = append(embeddedtypes, typ)
		}
	}
	universe := c.Universe
	pkg := universe.LoadPackage(c.FileComp().Path)
	return universe.InterfaceOf(pkg, methodnames, methodtypes, embeddedtypes)
}

// InterfaceProxy returns the proxy struct that implements a compiled interface
func (c *Comp) InterfaceProxy(t xr.Type) r.Type {
	ret := c.interf2proxy[t.ReflectType()]
	if ret == nil {
		c.Errorf("internal error: proxy not found for %s type <%v>", t.Kind(), t)
	}
	return ret
}

// converterToProxy compiles a conversion from 'tin' into a proxy struct that implements the interface type 'tout'
// and returns a function that performs such conversion
func (c *Comp) converterToProxy(tin xr.Type, tout xr.Type) func(val xr.Value) xr.Value {
	rtout := tout.ReflectType()       // a compiled interface
	rtproxy := c.InterfaceProxy(tout) // one of our proxies that pre-implement the compiled interface

	vtable := xr.NewR(rtproxy).Elem()
	n := rtout.NumMethod()
	for i := 0; i < n; i++ {
		mtdout := rtout.Method(i)
		mtdin, count := tin.MethodByName(mtdout.Name, mtdout.PkgPath)
		if count == 0 {
			c.Errorf("cannot convert type <%v> to interface <%v>: missing method %s %s", tin, rtout, mtdout.PkgPath, mtdout.Name)
		} else if count > 1 {
			c.Errorf("type <%v> has %d wrapper methods %s %s all at the same depth=%d - cannot convert to interface <%v>",
				tin, count, mtdout.PkgPath, mtdout.Name, len(mtdin.FieldIndex), tout)
		}
		e := c.compileMethodAsFunc(tin, mtdin)
		// c.Debugf("type %v proxy %v method %s = %v // %v", tin.Name(), tout.Name(), mtdin.Name, e.Value, e.Type)
		setProxyField(vtable.Field(i+1), xr.ValueOf(e.Value))
	}
	extractor := c.extractor(tin)
	if extractor == nil {
		return func(val xr.Value) xr.Value {
			vaddr := xr.NewR(rtproxy)
			vproxy := vaddr.Elem()
			vproxy.Set(vtable)
			vproxy.Field(0).Set(xr.ValueOf(xr.MakeInterfaceHeader(val, tin)))
			return convert(vaddr, rtout)
		}
	}
	// extract object from tin proxy or emulated interface (if any),
	// and wrap it in tout proxy
	return func(val xr.Value) xr.Value {
		v, t := extractor(val)
		vaddr := xr.NewR(rtproxy)
		vproxy := vaddr.Elem()
		vproxy.Set(vtable)
		vproxy.Field(0).Set(xr.ValueOf(xr.MakeInterfaceHeader(v, t)))
		return convert(vaddr, rtout)
	}
}

func setProxyField(place xr.Value, mtd xr.Value) {
	rtin := mtd.Type()
	rtout := place.Type()
	if rtin == rtout {
		place.Set(mtd)
	} else if rtin.ConvertibleTo(rtout) {
		place.Set(mtd.Convert(rtout))
	} else {
		rcall := r.Value.Call
		if rtout.IsVariadic() {
			rcall = r.Value.CallSlice
		}
		rmtd := mtd.ReflectValue()
		rfunc := r.MakeFunc(rtout, func(args []r.Value) []r.Value {
			args[0] = args[0].Interface().(xr.InterfaceHeader).Value().ReflectValue()
			return rcall(rmtd, args)
		})
		place.Set(xr.MakeValue(rfunc))
	}
}

// extract a value from a proxy struct (one of the imports.* structs) that implements an interface
// this is the inverse of the function returned by Comp.converterToProxy() above
func (g *CompGlobals) extractFromProxy(v xr.Value) (xr.Value, xr.Type) {
	// base.Debugf("type assertion: value = %v <%v>", v, base.Type(v))

	// v.Kind() is allowed also on invalid xr.Value, and it returns r.Invalid
	if v.Kind() == r.Interface {
		v = v.Elem() // extract concrete type
	}
	if !v.IsValid() || v == None {
		// cannot extract concrete type
		return v, nil
	}
	rt := v.Type()
	var xt xr.Type
	// base.Debugf("type assertion: concrete value = %v <%v>", i, t)
	if rt != nil && rt.Kind() == r.Ptr && g.proxy2interf[rt.Elem()] != nil {
		v = v.Elem().Field(0)
		if j, ok := reflect.ValueInterface(v).(xr.InterfaceHeader); ok {
			// base.Debugf("type assertion: unwrapped value = %v <%T>", j, j)
			v = j.Value()
			xt = j.Type()
		} else {
			// base.Debugf("type assertion: failed to unwrap value = %v <%T>", i, i)
			if v.Kind() == r.Interface {
				v = v.Elem() // extract concrete type
			}
		}
	}
	return v, xt
}

// converterToProxy compiles a conversion from 'tin' into the emulated interface type 'tout'
// and returns a function that performs such conversion
func (c *Comp) converterToEmulatedInterface(tin, tout xr.Type) func(val xr.Value) xr.Value {
	if !tin.Implements(tout) {
		c.Errorf("cannot convert from <%v> to <%v>", tin, tout)
	}
	n := tout.NumMethod()
	obj2methodFuncs := make([]func(xr.Value) xr.Value, n)

	tsrc := tin
	if tin.Kind() == r.Ptr {
		// xr.Type.MethodByName wants T, not *T, even for methods with pointer receiver
		tsrc = tin.Elem()
	}
	debug := c.Options&base.OptDebugMethod != 0
	for i := 0; i < n; i++ {
		mtdout := tout.Method(i)
		mtdin, count := tsrc.MethodByName(mtdout.Name, c.PackagePath) // pkgpath is ignored for exported names

		if count == 0 {
			c.Errorf("cannot convert from <%v> to <%v>: missing method %s %s", tin, tout, mtdout.Name, mtdout.Type)
		} else if count > 1 {
			c.Errorf("cannot convert from <%v> to <%v>: multiple methods match %s %s", tin, tout, mtdout.Name, mtdout.Type)
		}
		if !mtdin.Type.AssignableTo(mtdout.Type) {
			c.Errorf("cannot convert from <%v> to <%v>: mismatched method %s: expecting %v, found %v",
				tin, tout, mtdout.Name, mtdout.Type, mtdin.Type)
		}
		obj2methodFuncs[i] = c.compileObjGetMethod(tin, mtdin)
		if debug {
			c.Debugf("compiled  method conversion from %v.%s <%v> (concrete method %d) to %v.%s <%v> (interface method %d)",
				tin, mtdin.Name, mtdin.Type, mtdin.Index, tout, mtdout.Name, mtdout.Type, mtdout.Index)
		}
	}
	rtout := tout.ReflectType()

	extractor := c.extractor(tin)
	if extractor == nil {
		return func(obj xr.Value) xr.Value {
			return xr.ToEmulatedInterface(rtout, obj, tin, obj2methodFuncs)
		}
	}
	// extract object from tin proxy or emulated interface (if any),
	// and wrap it in tout emulated interface
	return func(obj xr.Value) xr.Value {
		v, t := extractor(obj)
		return xr.ToEmulatedInterface(rtout, v, t, obj2methodFuncs)
	}
}

// return a function that extracts value wrapped in a proxy or emulated interface
// returns nil if no extraction is needed
func (g *CompGlobals) extractor(tin xr.Type) func(xr.Value) (xr.Value, xr.Type) {
	if tin.Kind() != r.Interface {
		return nil
	} else if xr.IsEmulatedInterface(tin) {
		return xr.FromEmulatedInterface
	} else {
		return g.extractFromProxy
	}
}

// return the error "\n\treason: t does not implement tinterf: missing method <method>"
func interfaceMissingMethod(t, tinterf xr.Type) string {
	var s string
	if tinterf.Kind() == r.Interface {
		s = fmt.Sprintf("\n\treason: %v does not implement %v", t, tinterf)
		missingmtd := xr.MissingMethod(t, tinterf)
		if missingmtd != nil {
			s = fmt.Sprintf("%s: missing method %s", s, missingmtd.String())
		}
	}
	return s
}
