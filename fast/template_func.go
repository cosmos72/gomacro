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
	"fmt"
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type TemplateFuncInstance struct {
	Func *func(*Env) r.Value
	Type xr.Type
}

type TemplateFunc struct {
	Decl           *ast.FuncLit // template function declaration. use a *ast.FuncLit because we will compile it with Comp.FuncLit()
	Params         []string     // template param names
	SpecializedFor []ast.Expr   // not used yet
	// cache of instantiated and compiled functions.
	// key is [N]interface{}{T1, T2...}
	// value is *TemplateFuncInstance to allow replacing instantiated functions at runtime (by recompiling them)
	Instances map[I]*TemplateFuncInstance
}

type templateMaker struct {
	comp  *Comp
	sym   *Symbol
	ifun  I
	vals  []I
	types []xr.Type
	ikey  I
	name  string
}

func (maker *templateMaker) injectBinds(c *Comp, names []string) {
	for i, name := range names {
		t := maker.types[i]
		if val := maker.vals[i]; val != nil {
			c.DeclConst0(name, t, val)
		} else {
			c.declTypeAlias(name, t)
		}
	}
}

// return the qualified name of the function or type to instantiate, for example "Pair#[int,string]"
func (maker *templateMaker) Name() string {
	if len(maker.name) != 0 {
		return maker.name
	}
	var buf bytes.Buffer
	buf.WriteString(maker.sym.Name)
	buf.WriteString("#[")

	for i, val := range maker.vals {
		if i != 0 {
			buf.WriteByte(',')
		}
		if val == nil {
			val = maker.types[i].ReflectType()
		}
		fmt.Fprint(&buf, val)
	}
	buf.WriteByte(']')
	maker.name = buf.String()
	return maker.name
}

func (f *TemplateFunc) String() string {
	if f == nil {
		return "<nil>"
	}
	var buf bytes.Buffer // strings.Builder requires Go >= 1.10

	buf.WriteString("template[")
	for i, param := range f.Params {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param)
	}
	buf.WriteString("] ")
	(*base.Stringer).Fprintf(nil, &buf, "%v", f.Decl.Type)
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

	paramNames := c.templateParamNames(lit.Elts, "function or method", decl)

	bind := c.NewBind(decl.Name.Name, TemplateFuncBind, c.TypeOfPtrTemplateFunc())

	// a template function declaration has no runtime effect:
	// it merely creates the bind for on-demand instantiation by other code

	bind.Value = &TemplateFunc{
		Decl: &ast.FuncLit{
			Type: decl.Type,
			Body: decl.Body,
		},
		Params:    paramNames,
		Instances: make(map[I]*TemplateFuncInstance),
	}
}

func (c *Comp) templateParamNames(params []ast.Expr, errlabel string, node ast.Node) []string {
	names := make([]string, len(params))
	for i, param := range params {
		if ident, _ := param.(*ast.Ident); ident != nil {
			names[i] = ident.Name
		} else {
			c.Errorf("invalid template %s declaration: template parameter %d should be an identifier, found %T: %v",
				errlabel, i, param, node)
		}
	}
	return names
}

func splitTemplateArgs(node *ast.IndexExpr) (string, []ast.Expr, bool) {
	if ident, _ := node.X.(*ast.Ident); ident != nil {
		cindex, _ := node.Index.(*ast.CompositeLit)
		if cindex != nil && cindex.Type == nil {
			return ident.Name, cindex.Elts, true
		}
	}
	return "", nil, false
}

func (c *Comp) templateMaker(node *ast.IndexExpr, which BindClass) *templateMaker {
	name, templateArgs, ok := splitTemplateArgs(node)
	if !ok {
		return nil
	}
	sym, upc := c.tryResolve(name)
	if sym == nil {
		c.Errorf("undefined identifier: %v", name)
	}
	n := len(templateArgs)
	var paramNames []string
	ifun := sym.Value
	ok = false
	if ifun != nil && sym.Desc.Class() == which {
		switch which {
		case TemplateFuncBind:
			fun, _ := ifun.(*TemplateFunc)
			ok = fun != nil
			paramNames = fun.Params
		case TemplateTypeBind:
			typ, _ := ifun.(*TemplateType)
			ok = typ != nil
			paramNames = typ.Params
		}
	}
	if !ok {
		c.Errorf("symbol is not a %v, cannot use #[...] on it: %s", which, name)
	}
	if n != len(paramNames) {
		c.Errorf("%v expects exactly %d template parameters %v, found %d: %v", which, len(paramNames), paramNames, n, node)
	}
	vals := make([]I, n)
	types := make([]xr.Type, n)
	// slices cannot be used as map keys. use an array and reflection
	key := r.New(r.ArrayOf(n, rTypeOfInterface)).Elem()

	for i, templateArg := range templateArgs {
		e, t := c.Expr1OrType(templateArg)
		if e != nil {
			if !e.Const() {
				c.Errorf("argument of template function %q is not a constant: %v", name, templateArg)
			}
			// UntypedLit is unsuitable as map key, because its == is not usable
			vals[i] = e.EvalConst(COptDefaults)
			types[i] = e.Type // also remember the type
			key.Index(i).Set(r.ValueOf(vals[i]))
		} else {
			types[i] = t
			key.Index(i).Set(r.ValueOf(t.ReflectType()))
		}
	}
	return &templateMaker{upc, sym, ifun, vals, types, key.Interface(), ""}
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
			c.Debugf("found instantiated template function %v", node)
		}
	} else {
		if c.Globals.Options&base.OptDebugTemplate != 0 {
			c.Debugf("instantiating template function %v", node)
		}
		// hard part: instantiate the template function.
		// must be instantiated in the same *Comp where it was declared!
		instance = maker.comp.instantiateTemplateFunc(maker, fun, node)
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
		c.Debugf("template function: %v, upn = %v, instance = %v", node, upn, instance)
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
func (c *Comp) instantiateTemplateFunc(maker *templateMaker, fun *TemplateFunc, node *ast.IndexExpr) *TemplateFuncInstance {

	// create a new nested Comp
	c = NewComp(c, nil)
	c.UpCost = 0
	c.Depth--

	// and inject template arguments into it
	maker.injectBinds(c, fun.Params)

	key := maker.ikey
	panicking := true
	defer func() {
		if panicking {
			delete(fun.Instances, key)
			c.ErrorAt(node.Pos(), "error instantiating template function: %v\n\t%v", node, recover())
		}
	}()

	if c.Globals.Options&base.OptDebugTemplate != 0 {
		c.Debugf("forward-declaring template function before instantiation: %v", node)
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
	t, _, _ := c.TypeFunction(fun.Decl.Type)

	instance := &TemplateFuncInstance{Type: t, Func: new(func(*Env) r.Value)}
	fun.Instances[key] = instance

	// compile an expression that, when evaluated at runtime in the *Env
	// where the template function was declared, returns the instantiated function
	expr := c.FuncLit(fun.Decl)

	*instance.Func = expr.AsX1()
	instance.Type = expr.Type

	panicking = false
	return instance
}
