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
 * template_func.go
 *
 *  Created on Jun 06, 2018
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"bytes"
	"go/ast"
	"go/token"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// an instantiated (and compiled) template function.
type TemplateFuncInstance struct {
	Func *func(*Env) r.Value
	Type xr.Type
}

// a template function declaration.
// either general, or partially specialized or fully specialized
type TemplateFuncDecl struct {
	Decl   *ast.FuncLit // template function declaration. use a *ast.FuncLit because we will compile it with Comp.FuncLit()
	Params []string     // template param names
	For    []ast.Expr   // partial or full specialization
}

// template function
type TemplateFunc struct {
	Master    TemplateFuncDecl            // master (i.e. non specialized) declaration
	Special   map[string]TemplateFuncDecl // partially or fully specialized declarations. key is TemplateFuncDecl.For converted to string
	Instances map[I]*TemplateFuncInstance // cache of instantiated functions. key is [N]interface{}{T1, T2...}
}

func (f *TemplateFunc) String() string {
	if f == nil {
		return "<nil>"
	}
	var buf bytes.Buffer // strings.Builder requires Go >= 1.10
	buf.WriteString("template[")
	decl := f.Master
	for i, param := range decl.Params {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param)
	}
	buf.WriteString("] ")
	(*base.Stringer).Fprintf(nil, &buf, "%v", decl.Decl.Type)
	return buf.String()
}

// DeclTemplateFunc stores a template function or method declaration
// for later instantiation
func (c *Comp) DeclTemplateFunc(decl *ast.FuncDecl) {
	n := 0
	if decl.Recv != nil {
		n = len(decl.Recv.List)
	}
	if n < 2 {
		c.Errorf("invalid template function or method declaration: expecting at least 2 receivers, found %d: %v", n, decl)
	}
	if decl.Recv.List[0] != nil {
		c.Errorf("template method declaration not yet implemented: %v", decl)
	}
	lit, _ := decl.Recv.List[1].Type.(*ast.CompositeLit)
	if lit == nil {
		c.Errorf("invalid template function or method declaration: the second receiver should be an *ast.CompositeLit, found %T: %v",
			decl.Recv.List[1].Type, decl)
	}

	params, fors := c.templateParams(lit.Elts, "function or method", decl)

	fdecl := TemplateFuncDecl{
		Decl: &ast.FuncLit{
			Type: decl.Type,
			Body: decl.Body,
		},
		Params: params,
		For:    fors,
	}
	name := decl.Name.Name

	if len(fors) == 0 {
		// master (i.e. not specialized) declaration

		if len(params) == 0 {
			c.Errorf("cannot declare template function with zero template parameters: %v", decl.Type)
		}
		bind := c.NewBind(name, TemplateFuncBind, c.TypeOfPtrTemplateFunc())

		// a template function declaration has no runtime effect:
		// it merely creates the bind for on-demand instantiation by other code
		bind.Value = &TemplateFunc{
			Master:    fdecl,
			Special:   make(map[string]TemplateFuncDecl),
			Instances: make(map[I]*TemplateFuncInstance),
		}
		return
	}

	// partially or fully specialized declaration
	bind := c.Binds[name]
	if bind == nil {
		c.Errorf("undefined identifier: %v", name)
	}
	fun, ok := bind.Value.(*TemplateFunc)
	if !ok {
		c.Errorf("symbol is not a template function, cannot declare function specializations on it: %s // %v", name, bind.Type)
	}
	key := c.Globals.Sprintf("%v", &ast.IndexExpr{X: decl.Name, Index: &ast.CompositeLit{Elts: fors}})
	if len(fun.Master.Params) != len(fors) {
		c.Errorf("template function specialization for %d parameters, expecting %d: %s", len(fors), len(fun.Master.Params), key)
	}
	if _, ok := fun.Special[key]; ok {
		c.Warnf("redefined template function specialization: %s", key)
	}
	fun.Special[key] = fdecl
}

// TemplateFunc compiles a template function name#[T1, T2...] instantiating it if needed.
func (c *Comp) TemplateFunc(node *ast.IndexExpr) *Expr {
	maker := c.templateMaker(node, TemplateFuncBind)
	if maker == nil {
		return nil
	}
	fun := maker.ifun.(*TemplateFunc)
	key := maker.ikey

	instance, _ := fun.Instances[key]
	if instance != nil {
		if c.Globals.Options&base.OptDebugTemplate != 0 {
			c.Debugf("found instantiated template function %v", maker)
		}
	} else {
		if c.Globals.Options&base.OptDebugTemplate != 0 {
			c.Debugf("instantiating template function %v", maker)
		}
		// hard part: instantiate the template function.
		// must be instantiated in the same *Comp where it was declared!
		instance = maker.instantiateFunc(fun, node)
	}

	var efun, retfun func(*Env) r.Value
	eaddr := instance.Func
	if *eaddr == nil {
		// currently instantiating it, see comment in Comp.instantiateTemplateFunc() below.
		// We must try again later to dereference instance.Func.
		efun = func(env *Env) r.Value {
			return (*eaddr)(env)
		}
	} else {
		efun = *eaddr
	}
	upn := maker.sym.Upn
	if c.Globals.Options&base.OptDebugTemplate != 0 {
		c.Debugf("template function: %v, upn = %v, instance = %v", maker, upn, instance)
	}
	// switch to the correct *Env before evaluating expr
	switch upn {
	case 0:
		retfun = efun
	case 1:
		retfun = func(env *Env) r.Value {
			return efun(env.Outer)
		}
	case 2:
		retfun = func(env *Env) r.Value {
			return efun(env.Outer.Outer)
		}
	case c.Depth - 1:
		retfun = func(env *Env) r.Value {
			return efun(env.FileEnv)
		}
	case c.Depth:
		retfun = func(env *Env) r.Value {
			return efun(env.FileEnv.Outer)
		}
	default:
		retfun = func(env *Env) r.Value {
			for i := upn; i > 0; i-- {
				env = env.Outer
			}
			return efun(env)
		}
	}
	// always return a new *Expr, in case caller modifies it
	return exprFun(instance.Type, retfun)
}

// instantiateTemplateFunc instantiates and compiles a template function.
// node is used only for error messages
func (maker *templateMaker) instantiateFunc(fun *TemplateFunc, node *ast.IndexExpr) *TemplateFuncInstance {

	// choose the specialization to use
	_, special := maker.chooseFunc(fun)

	// create a new nested Comp
	c := NewComp(maker.comp, nil)
	c.UpCost = 0
	c.Depth--

	// and inject template arguments into it
	special.injectBinds(c)

	key := maker.ikey
	panicking := true
	defer func() {
		if panicking {
			delete(fun.Instances, key)
			c.ErrorAt(node.Pos(), "error instantiating template function: %v\n\t%v", maker, recover())
		}
	}()

	if c.Globals.Options&base.OptDebugTemplate != 0 {
		c.Debugf("forward-declaring template function before instantiation: %v", maker)
	}
	// support for template recursive functions, as for example
	//   template[T] func fib(n T) T { if n <= 2 { return 1 }; return fib#[T](n-1) + fib#[T](n-2) }
	// requires to cache fib#[T] as instantiated **before** actually instantiating it.
	//
	// This is similar to the technique used for non-template recursive function, as
	//    func fib(n int) int { if n <= 2 { return 1 }; return fib(n-1) + fib(n-2) }
	// with the difference that the cache is fun.Instances[key] instead of Comp.Binds[name]

	// for such trick to work, we must:
	// 1. compute in advance the instantiated function type
	// 2. check TemplateFuncInstance.Func: if it's nil, take its address and dereference it later at runtime
	t, _, _ := c.TypeFunction(special.decl.Decl.Type)

	instance := &TemplateFuncInstance{Type: t, Func: new(func(*Env) r.Value)}
	fun.Instances[key] = instance

	// compile an expression that, when evaluated at runtime in the *Env
	// where the template function was declared, returns the instantiated function
	expr := c.FuncLit(special.decl.Decl)

	*instance.Func = expr.AsX1()
	instance.Type = expr.Type

	panicking = false
	return instance
}

/*
 *
 *
 * type inference on template functions
 *
 *
 */

type templateInferrer struct {
	comp     *Comp
	tfun     *TemplateFunc
	params   []string
	patterns []ast.Expr
	targs    []xr.Type
	matches  []xr.Type
	node     *ast.CallExpr // for error messages
}

func (c *Comp) inferTemplateFunc(node *ast.CallExpr, fun *Expr, args []*Expr) *Expr {
	tfun, ok := fun.Value.(*TemplateFunc)
	if !ok {
		c.Errorf("internal error: Comp.inferTemplateFunc() invoked on non-template function %v: %v", fun.Type, node.Fun)
	}
	master := tfun.Master
	typ := master.Decl.Type

	var patterns []ast.Expr
	ellipsis := node.Ellipsis != token.NoPos
	variadic := false
	// collect template function param types
	if params := typ.Params; params != nil {
		if n := len(params.List); n != 0 {
			_, variadic = params.List[n-1].Type.(*ast.Ellipsis)
			for _, param := range params.List {
				for _ = range param.Names {
					patterns = append(patterns, param.Type)
				}
			}
		}
	}
	if variadic && !ellipsis {
		c.Errorf("unimplemented type inference on variadic template function: %v", node)
	} else if !variadic && ellipsis {
		c.Errorf("invalid use of ... in call to non-variadic template function: %v", node)
	}

	// collect call arg types
	nargs := len(args)
	var targs []xr.Type
	if nargs == 1 {
		arg := args[0]
		nargs = arg.NumOut()
		if nargs == 1 {
			targs = []xr.Type{arg.Type}
		} else {
			targs = arg.Types
		}
	} else {
		targs = make([]xr.Type, nargs)
		for i, arg := range args {
			targs[i] = arg.Type
		}
	}
	if nargs != len(patterns) {
		c.Errorf("template function %v has %d params, cannot call with %d values: %v", tfun, len(patterns), nargs, node)
	}
	inf := templateInferrer{c, tfun, master.Params, patterns, nil, targs, node}
	return inf.inferArgs()
}

func (inf *templateInferrer) inferArgs() *Expr {
	inf.matches = make([]xr.Type, len(inf.params))
	for i := range inf.targs {
		inf.inferArg(i)
	}
	inf.comp.Errorf("unimplemented type inference on template function: %v", inf.node)
	return nil
}

func (inf *templateInferrer) inferArg(i int) {
}
