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
 * template_maker.go
 *
 *  Created on Jun 16, 2018
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"bytes"
	"fmt"
	"go/ast"
	r "reflect"

	xr "github.com/cosmos72/gomacro/xreflect"
)

type templateMaker struct {
	comp  *Comp
	sym   *Symbol
	ifun  I
	exprs []ast.Expr
	vals  []I
	types []xr.Type
	ikey  I
	name  string
}

type templateTypeCandidate struct {
	decl  TemplateTypeDecl
	vals  []I
	types []xr.Type
}

type templateFuncCandidate struct {
	decl  TemplateFuncDecl
	vals  []I
	types []xr.Type
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
	var params []string
	ifun := sym.Value
	ok = false
	if ifun != nil && sym.Desc.Class() == which {
		switch which {
		case TemplateFuncBind:
			fun, _ := ifun.(*TemplateFunc)
			ok = fun != nil
			if ok {
				params = fun.Master.Params
			}
		case TemplateTypeBind:
			typ, _ := ifun.(*TemplateType)
			ok = typ != nil
			if ok {
				params = typ.Master.Params
			}
		}
	}
	if !ok {
		c.Errorf("symbol is not a %v, cannot use #[...] on it: %s", which, name)
	}
	if n != len(params) {
		c.Errorf("%v expects exactly %d template parameters %v, found %d: %v", which, len(params), params, n, node)
	}
	vals := make([]I, n)
	types := make([]xr.Type, n)
	// slices cannot be used as map keys. use an array and reflection
	key := r.New(r.ArrayOf(n, rtypeOfInterface)).Elem()

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
			key.Index(i).Set(r.ValueOf(xr.MakeKey(t)))
		}
	}
	return &templateMaker{upc, sym, ifun, templateArgs, vals, types, key.Interface(), ""}
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

func (c *Comp) templateParams(params []ast.Expr, errlabel string, node ast.Node) ([]string, []ast.Expr) {
	names := make([]string, 0, len(params))
	var exprs []ast.Expr
	for i, param := range params {
		switch param := param.(type) {
		case *ast.Ident:
			names = append(names, param.Name)
		case *ast.BadExpr:
		case *ast.CompositeLit:
			exprs = param.Elts
		default:
			c.Errorf("invalid template %s declaration: template parameter %d should be *ast.Ident or *ast.CompositeLit, found %T: %v",
				errlabel, i, param, node)
		}
	}
	return names, exprs
}

// return the most specialized function declaration applicable to used params.
// panics if there is no single most specialized declaration.
func (maker *templateMaker) chooseFunc(fun *TemplateFunc) *templateFuncCandidate {
	candidate := templateFuncCandidate{
		decl:  fun.Master,
		vals:  maker.vals,
		types: maker.types,
	}
	/*
		for name, special := range fun.Special {
		}
	*/
	return &candidate
}

// return the most specialized type declaration applicable to used params.
// panics if there is no single most specialized declaration.
func (maker *templateMaker) chooseType(typ *TemplateType) *templateTypeCandidate {
	candidate := templateTypeCandidate{
		decl:  typ.Master,
		vals:  maker.vals,
		types: maker.types,
	}
	return &candidate
}

func (special *templateFuncCandidate) injectBinds(c *Comp) {
	for i, name := range special.decl.Params {
		t := special.types[i]
		if val := special.vals[i]; val != nil {
			c.DeclConst0(name, t, val)
		} else {
			c.declTypeAlias(name, t)
		}
	}
}

func (special *templateTypeCandidate) injectBinds(c *Comp) {
	for i, name := range special.decl.Params {
		t := special.types[i]
		if val := special.vals[i]; val != nil {
			c.DeclConst0(name, t, val)
		} else {
			c.declTypeAlias(name, t)
		}
	}
}
