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
 * template_function.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"
	"strings"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type TemplateFunc struct {
	Decl           *ast.FuncLit // template function declaration. use a *ast.FuncLit because we will compile it with Comp.FuncLit()
	Params         []string     // template param names
	SpecializedFor []ast.Expr   // not used yet
	// cache of instantiated and compiled functions.
	// key is [N]interface{}{T1, T2...}
	// value is *TemplateFuncInstance to allow replacing instantiated template functions at runtime (by recompiling them)
	Instances map[I]*Expr
}

func (f *TemplateFunc) String() string {
	if f == nil {
		return "<nil>"
	}
	var buf strings.Builder
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

// TemplateFuncDecl stores a template function or method declaration
// for later instantiation
func (c *Comp) TemplateFuncDecl(decl *ast.FuncDecl) {
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
	var templateTypes []ast.Expr
	if clit, _ := decl.Recv.List[1].Type.(*ast.CompositeLit); clit != nil {
		templateTypes = clit.Elts
	} else {
		c.Errorf("invalid template function or method declaration: the second receiver should be an *ast.CompositeLit, found %T: %v",
			decl.Recv.List[1].Type, decl)
	}

	typeParams := make([]string, len(templateTypes))
	for i, typ := range templateTypes {
		if ident, _ := typ.(*ast.Ident); ident != nil {
			typeParams[i] = ident.Name
		} else {
			c.Errorf("invalid template function or method declaration: template parameter %d should be an identifier, found %T: %v",
				typ, decl)
		}
	}

	bind := c.NewBind(decl.Name.Name, TemplateFuncBind, c.TypeOfPtrTemplateFunc())

	// a template function declaration has no runtime effect:
	// it merely creates the bind for on-demand instantiation by other code

	bind.Value = &TemplateFunc{
		Decl: &ast.FuncLit{
			Type: decl.Type,
			Body: decl.Body,
		},
		Params:    typeParams,
		Instances: make(map[I]*Expr),
	}
}

// TemplateFunc compiles a template function name#[T1, T2...] instantiating it if needed.
// node is used only for error messages
func (c *Comp) TemplateFunc(name string, templateArgs []ast.Expr, node *ast.IndexExpr) *Expr {
	sym, upc := c.tryResolve(name)
	if sym == nil {
		c.Errorf("undefined identifier: %v", name)
	}
	fun, _ := sym.Value.(*TemplateFunc)
	if fun == nil || sym.Desc.Class() != TemplateFuncBind {
		c.Errorf("symbol is not a template function, cannot use #[...] template arguments on it: %s", name)
	}
	n := len(templateArgs)
	if n != len(fun.Params) {
		c.Errorf("template function %q expects exactly %d template parameters %v, found %d: %v", name, len(fun.Params), fun.Params, n, templateArgs)
	}
	vals := make([]I, n)
	types := make([]xr.Type, n)
	key := r.New(r.ArrayOf(n, rTypeOfInterface)).Elem() // slices cannot be used as map keys... use an array and reflection

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

	ikey := key.Interface()
	expr, _ := fun.Instances[ikey]
	if expr == nil {
		// hard part: instantiate the template function.
		// must be instantiated in the same *Comp where it was declared!
		expr = upc.instantiateTemplateFunc(fun, vals, types, node)
		fun.Instances[ikey] = expr
	}

	efun := expr.AsX1()
	var retfun func(*Env) r.Value

	// switch to the correct *Env before evaluating expr
	switch upn := sym.Upn; upn {
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
	return exprFun(expr.Type, retfun)
}

// TemplateFunc instantiates and compiles a template function.
// node and origC are used only for error messages
func (c *Comp) instantiateTemplateFunc(fun *TemplateFunc, vals []I, types []xr.Type, node *ast.IndexExpr) *Expr {

	// create a new nested Comp, and inject template arguments in it
	c = NewComp(c, nil)
	c.UpCost = 0
	c.Depth--

	for i, name := range fun.Params {
		t := types[i]
		if val := vals[i]; val != nil {
			c.DeclConst0(name, t, val)
		} else {
			c.declTypeAlias(name, t)
		}
	}
	panicking := true
	defer func() {
		if panicking {
			c.ErrorAt(node.Pos(), "error instantiating template function: %v\n\t%v", node, recover())
		}
	}()
	// compile an expression that, when evaluated at runtime in the *Env
	// where the template function was declared, returns the instantiated function
	expr := c.FuncLit(fun.Decl)
	panicking = false
	return expr
}
