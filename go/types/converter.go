// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file converts objects from go/types to github.com/cosmos72/go/etypes

package types

import (
	"fmt"
	"go/types"
)

type Converter struct {
	pkg map[string]*Package
	// use pointer identity to compare types, not types.Identical.
	// faster although less accurate.
	cache        map[types.Type]Type
	toaddmethods map[*Named]*types.Named
	tocomplete   []*Interface
}

type funcOption bool

const (
	funcIgnoreRecv funcOption = false
	funcSetRecv    funcOption = true
)

func (c *Converter) mkpackage(g *types.Package) *Package {
	if g == nil {
		return nil
	}
	path := g.Path()
	if p := c.pkg[path]; p != nil {
		return p
	}
	p := NewPackage(path, g.Name())
	if c.pkg == nil {
		c.pkg = make(map[string]*Package)
	}
	c.pkg[path] = p
	return p
}

func (c *Converter) universe() *Package {
	if p := c.pkg[""]; p != nil {
		return p
	}
	p := NewPackage("", "")
	if c.pkg == nil {
		c.pkg = make(map[string]*Package)
	}
	c.pkg[""] = p
	return p
}

func (c *Converter) mktypename(g *types.TypeName) (*TypeName, bool) {
	pkg := c.mkpackage(g.Pkg())
	if pkg == nil {
		pkg = c.universe()
	}
	scope := pkg.Scope()
	obj := scope.Lookup(g.Name())
	if typename, ok := obj.(*TypeName); ok {
		return typename, true
	}
	typename := NewTypeName(g.Pos(), pkg, g.Name(), nil)
	pkg.Scope().Insert(typename)
	return typename, false
}

func (c *Converter) mkfield(g *types.Var) *Var {
	return NewField(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()), g.Embedded())
}

func (c *Converter) mkparam(g *types.Var) *Var {
	if g == nil {
		return nil
	}
	return NewParam(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()))
}

func (c *Converter) mkparams(g *types.Tuple) *Tuple {
	if g == nil {
		return nil
	}
	n := g.Len()
	v := make([]*Var, n)
	for i := 0; i < n; i++ {
		v[i] = c.mkparam(g.At(i))
	}
	return NewTuple(v...)
}

func (c *Converter) mkvar(g *types.Var) *Var {
	if g == nil {
		return nil
	}
	return NewVar(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()))
}

func (c *Converter) Type(g types.Type) Type {
	if c.cache == nil {
		c.cache = make(map[types.Type]Type)
	}
	ret := c.typ(g)
	for _, t := range c.tocomplete {
		t.Complete()
	}
	c.tocomplete = c.tocomplete[0:0:cap(c.tocomplete)]

	for t, g := range c.toaddmethods {
		c.addmethods(t, g)
		delete(c.toaddmethods, t)
	}
	return ret
}

func (c *Converter) typ(g types.Type) Type {
	t := c.cache[g]
	if t != nil {
		return t
	}
	switch g := g.(type) {
	case *types.Array:
		elem := c.typ(g.Elem())
		t = NewArray(elem, g.Len())
	case *types.Basic:
		return Typ[BasicKind(g.Kind())]
	case *types.Chan:
		elem := c.typ(g.Elem())
		t = NewChan(ChanDir(g.Dir()), elem)
	case *types.Interface:
		t = c.mkinterface(g)
	case *types.Map:
		t = c.mkmap(g)
	case *types.Named:
		t = c.mknamed(g)
	case *types.Pointer:
		elem := c.typ(g.Elem())
		t = NewPointer(elem)
	case *types.Signature:
		t = c.mksignature(g, funcSetRecv)
	case *types.Slice:
		elem := c.typ(g.Elem())
		t = NewSlice(elem)
	case *types.Struct:
		t = c.mkstruct(g)
	default:
		panic(fmt.Errorf("Converter.Type(): unsupported types.Type: %T", g))
	}
	c.cache[g] = t
	return t
}

func (c *Converter) mkinterface(g *types.Interface) *Interface {
	n := g.NumExplicitMethods()
	fs := make([]*Func, n)
	for i := 0; i < n; i++ {
		fs[i] = c.mkfunc(g.ExplicitMethod(i), funcIgnoreRecv)
	}
	n = g.NumEmbeddeds()
	es := make([]Type, n)
	for i := 0; i < n; i++ {
		es[i] = c.typ(g.EmbeddedType(i))
	}
	t := NewInterfaceType(fs, es)
	c.tocomplete = append(c.tocomplete, t)
	return t
}

func (c *Converter) mkmap(g *types.Map) *Map {
	key := c.typ(g.Key())
	elem := c.typ(g.Elem())
	return NewMap(key, elem)
}

func (c *Converter) mknamed(g *types.Named) *Named {
	typename, found := c.mktypename(g.Obj())
	if found {
		return typename.Type().(*Named)
	}
	t := NewNamed(typename, nil, nil)
	u := c.typ(g.Underlying())
	t.SetUnderlying(u)
	if g.NumMethods() != 0 {
		if c.toaddmethods == nil {
			c.toaddmethods = make(map[*Named]*types.Named)
		}
		c.toaddmethods[t] = g
	}
	return t
}

func (c *Converter) mksignature(g *types.Signature, opt funcOption) *Signature {
	var recv *Var
	if opt == funcSetRecv {
		recv = c.mkparam(g.Recv())
	}
	return NewSignature(
		recv,
		c.mkparams(g.Params()),
		c.mkparams(g.Results()),
		g.Variadic(),
	)
}

func (c *Converter) mkstruct(g *types.Struct) *Struct {
	n := g.NumFields()
	fields := make([]*Var, n)
	tags := make([]string, n)
	for i := 0; i < n; i++ {
		fields[i] = c.mkfield(g.Field(i))
		tags[i] = g.Tag(i)
	}
	return NewStruct(fields, tags)
}

func (c *Converter) addmethods(t *Named, g *types.Named) {
	n := g.NumMethods()
	for i := 0; i < n; i++ {
		m := c.mkfunc(g.Method(i), funcSetRecv)
		t.AddMethod(m)
	}
}

func (c *Converter) mkfunc(m *types.Func, opt funcOption) *Func {
	sig := c.mksignature(m.Type().(*types.Signature), opt)
	return NewFunc(m.Pos(), c.mkpackage(m.Pkg()), m.Name(), sig)
}
