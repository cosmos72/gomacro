/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * template_type.go
 *
 *  Created on Jun 06, 2018
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type TemplateType struct {
	Decl           ast.Expr      // type declaration body. use an ast.Expr because we will compile it with Comp.Type()
	Alias          bool          // true if declaration is an alias: 'type Foo = ...'
	Params         []string      // template param names
	SpecializedFor []ast.Expr    // not used yet
	Instances      map[I]xr.Type // cache of instantiated types. key is [N]interface{}{T1, T2...}
}

// DeclTemplateType stores a template type declaration
// for later instantiation
func (c *Comp) DeclTemplateType(spec *ast.TypeSpec) {

	lit, _ := spec.Type.(*ast.CompositeLit)
	if lit == nil {
		c.Errorf("invalid template type declaration: expecting an *ast.CompositeLit, found %T: %v",
			spec.Type, spec)
	}
	expr := lit.Type
	if _, ok := expr.(*ast.CompositeLit); ok {
		c.Errorf("invalid template type declaration: expecting an *ast.CompositeLit, found &ast.CompositeLit{Type: &ast.CompositeLit{}}: %v",
			spec)
	}
	paramNames := c.templateParamNames(lit.Elts, "type", spec)

	bind := c.NewBind(spec.Name.Name, TemplateTypeBind, c.TypeOfPtrTemplateFunc())

	// a template type declaration has no runtime effect:
	// it merely creates the bind for on-demand instantiation by other code

	bind.Value = &TemplateType{
		Decl:      lit.Type,
		Alias:     spec.Assign != token.NoPos,
		Params:    paramNames,
		Instances: make(map[I]xr.Type),
	}
}

// TemplateType compiles a template type name#[T1, T2...] instantiating it if needed.
func (c *Comp) TemplateType(node *ast.IndexExpr) xr.Type {
	maker := c.compileTemplateArgs(node, TemplateTypeBind)
	if maker == nil {
		return nil
	}
	typ := maker.ifun.(*TemplateType)
	key := maker.ikey

	instance, _ := typ.Instances[key]
	if instance != nil {
		if c.Globals.Options&base.OptDebugTemplate != 0 {
			c.Debugf("found instantiated template type %v", node)
		}
	} else {
		if c.Globals.Options&base.OptDebugTemplate != 0 {
			c.Debugf("instantiating template type %v", node)
		}
		// hard part: instantiate the template type.
		// must be instantiated in the same *Comp where it was declared!
		instance = maker.comp.instantiateTemplateType(maker, typ, node)
	}
	return instance
}

// instantiateTemplateType instantiates and compiles a template function.
// node is used only for error messages
func (c *Comp) instantiateTemplateType(maker *templateMaker, typ *TemplateType, node *ast.IndexExpr) xr.Type {

	// create a new nested Comp
	c = NewComp(c, nil)
	c.UpCost = 0
	c.Depth--

	// and inject template arguments in it
	maker.injectBinds(c, typ.Params)

	key := maker.ikey

	panicking := true
	defer func() {
		if panicking {
			delete(typ.Instances, key) // remove the cached instance if present
			c.ErrorAt(node.Pos(), "error instantiating template type: %v\n\t%v", node, recover())
		}
	}()
	// compile the type instantiation
	//
	debug := c.Globals.Options&base.OptDebugTemplate != 0
	var t xr.Type
	if !typ.Alias && maker.sym.Name != "_" {
		if debug {
			c.Debugf("forward-declaring template type before instantiation: %v", node)
		}
		// support for recursive types, as for example
		//   template[T] type List struct { First T; Rest *List#[T] }
		// requires to cache List#[T] as instantiated **before** actually instantiating it.
		//
		// This is similar to the technique used for non-template recursive types, as
		//    type List struct { First int; Rest *List }
		// with the difference that the cache is typ.Instances[key] instead of Comp.Types[name]
		t = c.Universe.NamedOf(maker.TypeString(), c.FileComp().Path, r.Invalid /*kind not yet known*/)
		typ.Instances[key] = t
	}
	u := c.Type(typ.Decl)
	if t == nil { // t == nil means it's either an alias, or name == "_" (discards the result of type declaration)
		t = u
	} else {
		c.SetUnderlyingType(t, u)
	}
	panicking = false
	return t
}
