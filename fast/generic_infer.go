/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * generic_infer.go
 *
 *  Created on Jun 06, 2018
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	"go/token"
	r "reflect"

	"github.com/cosmos72/gomacro/base/untyped"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type inferType struct {
	Type    xr.Type
	Untyped untyped.Kind // for untyped literals
	Value   I            // in case we infer a constant, not a type
	Exact   bool
}

func (inf *inferType) String() string {
	if inf.Value != nil {
		return fmt.Sprint(inf.Value)
	}
	var s string
	if inf.Type != nil {
		s = inf.Type.String()
	} else {
		s = inf.Untyped.String()
	}
	return s
}

// type inference on generic functions
type inferFuncType struct {
	comp     *Comp
	tfun     *GenericFunc
	funcname string
	inferred map[string]inferType
	patterns []ast.Expr
	targs    []inferType
	call     *ast.CallExpr // for error messages
}

func (inf *inferFuncType) String() string {
	return inf.tfun.Signature(inf.funcname)
}

func (inf *inferFuncType) CallPos() token.Pos {
	if inf != nil && inf.call != nil && inf.call.Fun != nil {
		return inf.call.Pos()
	}
	return token.NoPos
}

func (inf *inferFuncType) DeclPos() token.Pos {
	if inf != nil && inf.tfun != nil {
		return inf.tfun.Pos()
	}
	return token.NoPos
}

func (c *Comp) inferGenericFunc(call *ast.CallExpr, fun *Expr, args []*Expr) *Expr {
	tfun, ok := fun.Value.(*GenericFunc)
	if !ok {
		c.Errorf("internal error: Comp.inferGenericFunc() invoked on non-generic function %v: %v", fun.Type, call.Fun)
	}
	// get the scope where fun is declared
	upc := tfun.DeclScope
	if upc == nil {
		c.Errorf("internal error: Comp.inferGenericFunc() failed to determine the scope containing generic function declaration: %v", call.Fun)
	}
	var funcname string
	if ident, ok := call.Fun.(*ast.Ident); ok {
		funcname = ident.Name
	} else {
		i := c.GensymN
		c.GensymN++
		funcname = fmt.Sprintf("lambda%d", i)
	}

	master := tfun.Master
	typ := master.Decl.Type

	var patterns []ast.Expr
	ellipsis := call.Ellipsis != token.NoPos
	variadic := false
	// collect generic function param types expressions
	if fields := typ.Params; fields != nil {
		if n := len(fields.List); n != 0 {
			_, variadic = fields.List[n-1].Type.(*ast.Ellipsis)
			for _, field := range fields.List {
				for _ = range field.Names {
					patterns = append(patterns, field.Type)
				}
			}
		}
	}
	if variadic && !ellipsis {
		c.Errorf("unimplemented type inference on variadic generic function: %v", call)
	} else if !variadic && ellipsis {
		c.Errorf("invalid use of ... in call to non-variadic generic function: %v", call)
	}

	// collect call arg types
	nargs := len(args)
	var targs []inferType
	if nargs == 1 && args[0].NumOut() != 1 {
		arg := args[0]
		nargs = arg.NumOut()
		targs = make([]inferType, nargs)
		for i := 0; i < nargs; i++ {
			targs[i] = inferType{Type: arg.Out(i)}
		}
	} else {
		targs = make([]inferType, nargs)
		for i, arg := range args {
			if kind := arg.UntypedKind(); kind != untyped.None {
				targs[i] = inferType{Untyped: kind}
			} else {
				targs[i] = inferType{Type: arg.Type}
			}
		}
	}
	if nargs != len(patterns) {
		c.Errorf("generic function %v has %d params, cannot call with %d values: %v", tfun, len(patterns), nargs, call)
	}
	inf := inferFuncType{
		comp: c, tfun: tfun, funcname: funcname,
		inferred: make(map[string]inferType),
		patterns: patterns, targs: targs, call: call,
	}
	vals, types := inf.args()
	maker := &genericMaker{
		comp: upc, sym: fun.Sym, ifun: fun.Sym.Value,
		exprs: nil, vals: vals, types: types,
		ikey: GenericKey(vals, types),
		pos:  inf.call.Pos(),
	}
	return c.genericFunc(maker, call)
}

// infer type of generic function from arguments
func (inf *inferFuncType) args() (vals []I, types []xr.Type) {
	exact := false // allow implicit type conversions

	// first pass: types and typed constants
	for i, targ := range inf.targs {
		node := inf.patterns[i]
		if targ.Type != nil {
			inf.arg(node, targ.Type, exact)
		} else if targ.Untyped != untyped.None {
			// skip untyped constant, handled below
		} else if targ.Value != nil {
			inf.constant(node, targ.Value, exact)
		} else {
			inf.fail(node, targ)
		}
	}

	// second pass: untyped constants
	for i, targ := range inf.targs {
		if targ.Type == nil && targ.Untyped != untyped.None {
			inf.untyped(inf.patterns[i], targ.Untyped)
		}
	}

	params := inf.tfun.Master.Params
	n := len(params)
	vals = make([]I, n)
	types = make([]xr.Type, n)
	for i, name := range params {
		inferred := inf.inferred[name]
		// inf.comp.Debugf("inferred  generic function type argument %v: to %v", name, inferred)
		if inferred.Type == nil && inferred.Untyped != untyped.None {
			inferred.Type = inf.comp.Universe.BasicTypes[inferred.Untyped.Reflect()]
		}
		if inferred.Type == nil {
			inf.comp.Errorf("failed to infer %v in call to generic function: %v", name, inf.call)
		}
		types[i] = inferred.Type
		vals[i] = inferred.Value
	}
	return vals, types
}

// partially infer type of generic function for a single parameter
func (inf *inferFuncType) arg(pattern ast.Expr, targ xr.Type, exact bool) {
	stars := 0
	for {
		if targ == nil {
			inf.fail(pattern, targ)
		}
		// inf.comp.Debugf("type inference: matching generic parameter <%v> with actual type <%v>", pattern, targ)

		switch node := pattern.(type) {
		case *ast.Ident:
			inf.ident(node, targ, exact)
			return
		case *ast.ArrayType:
			pattern, targ, exact = inf.arrayType(node, targ, exact)
			continue
		case *ast.ChanType:
			pattern, targ, exact = inf.chanType(node, targ, exact)
			continue
		case *ast.FuncType:
			pattern, targ, exact = inf.funcType(node, targ, exact)
			if pattern != nil {
				continue
			}
		case *ast.IndexExpr:
			// function's parameter is itself a generic
			pattern, targ, exact = inf.genericType(node, targ, exact)
			if pattern != nil {
				continue
			}
		case *ast.InterfaceType:
			pattern, targ, exact = inf.interfaceType(node, targ, exact)
			if pattern != nil {
				continue
			}
		case *ast.MapType:
			pattern, targ, exact = inf.mapType(node, targ, exact)
			continue
		case *ast.ParenExpr:
			pattern = node.X
			continue
		case *ast.SelectorExpr:
			// packagename.typename
			pattern, targ, exact = inf.selector(node, targ, exact)
			if pattern != nil {
				continue
			}
		case *ast.StarExpr:
			inf.is(pattern, targ, r.Ptr)
			pattern, targ = node.X, targ.Elem()
			if stars != 0 {
				exact = true
			}
			stars++
			continue
		case *ast.StructType:
			pattern, targ, exact = inf.structType(node, targ, exact)
			if pattern != nil {
				continue
			}
		default:
			inf.unimplemented(node, targ)
		}
		break
	}
}

// partially infer type of generic function from an array or slice parameter
func (inf *inferFuncType) arrayType(node *ast.ArrayType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	if node.Len == nil {
		inf.is(node, targ, r.Slice)
	} else {
		inf.is(node, targ, r.Array)
		if _, ok := node.Len.(*ast.Ellipsis); !ok {
			// [n]array
			inf.constant(node.Len, targ.Len(), exact)
		}
	}
	return node.Elt, targ.Elem(), true
}

// partially infer type of generic function for a channel parameter
func (inf *inferFuncType) chanType(node *ast.ChanType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	inf.is(node, targ, r.Chan)
	tdir := targ.ChanDir()
	dir := reflectChanDir(node.Dir)
	if dir&tdir == 0 || (exact && dir != tdir) {
		inf.fail(node, targ)
	}
	return node.Value, targ.Elem(), true
}

// partially infer type of generic function for a constant parameter
func (inf *inferFuncType) constant(node ast.Expr, val I, exact bool) {
	// TODO
	inf.comp.ErrorAt(inf.CallPos(), "unimplemented type inference: function with generic parameter <%v> and argument <%v>: %v",
		node, val, inf.call)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// partially infer type of generic function for a func parameter
func (inf *inferFuncType) funcType(node *ast.FuncType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	inf.is(node, targ, r.Func)
	if node.Params.NumFields() != targ.NumIn() {
		inf.fail(node, targ)
	}
	if node.Results.NumFields() != targ.NumOut() {
		inf.fail(node, targ)
	}
	if node.Params != nil {
		i := 0
		for _, param := range node.Params.List {
			end := i + max(len(param.Names), 1)
			for ; i < end; i++ {
				inf.arg(param.Type, targ.In(i), exact)
			}
		}
	}
	if node.Results != nil {
		i := 0
		for _, param := range node.Results.List {
			end := i + max(len(param.Names), 1)
			for ; i < end; i++ {
				inf.arg(param.Type, targ.Out(i), exact)
			}
		}
	}
	return nil, nil, exact
}

// partially infer type of generic function for an identifier parameter
func (inf *inferFuncType) ident(node *ast.Ident, targ xr.Type, exact bool) {
	c := inf.comp
	name := node.Name
	inferred, ok := inf.inferred[name]
	// inf.comp.Debugf("inferring generic function type argument %v: currently inferred %+v, must match %v", name, inferred, targ)
	if !ok && !inf.tfun.Master.HasParam(name) {
		// name must be an existing type
		t := c.TryResolveType(name)
		if t != nil {
			if !targ.AssignableTo(t) {
				inf.comp.ErrorAt(inf.CallPos(),
					"type inference: in %v, mismatched types for %v: %v cannot be assigned to %v: %v",
					inf, name, targ, t, inf.call)
			}
		}
		return
	}

	// inferring one of the function generic parameters
	inf.combine(node, &inferred, inferType{Type: targ, Exact: exact})
	inf.inferred[name] = inferred

}

func (inf *inferFuncType) combine(node ast.Expr, inferred *inferType, with inferType) {
	targ := with.Type
	exact := with.Exact
	if inferred.Type == nil {
		inferred.Type = targ
	} else if !inferred.Type.IdenticalTo(targ) {
		// inf.comp.Debugf("type inference: combining inferred types %+v and %+v for generic parameter %v",	*inferred, with, node)
		if exact && inferred.Exact {
			inf.fail3(node, inferred, targ)
		}
		fwd := targ.AssignableTo(inferred.Type)
		rev := inferred.Type.AssignableTo(targ)
		if inferred.Exact {
			if fwd {
				inf.fail3(node, inferred, targ)
			}
		} else if exact {
			if rev {
				inferred.Type = targ
			} else {
				inf.fail3(node, inferred, targ)
			}
		} else {
			if fwd && rev {
				if !targ.Named() {
					inferred.Type = targ
				}
			} else if fwd {
			} else if rev {
				inferred.Type = targ
			} else {
				inf.fail3(node, inferred, targ)
			}
		}
	}
	if exact {
		inferred.Exact = true
	}
}

func (inf *inferFuncType) untyped(node ast.Expr, kind untyped.Kind) {
	ident, ok := node.(*ast.Ident)
	if !ok {
		inf.fail(node, kind)
	}
	name := ident.Name
	inferred := inf.inferred[name]
	// inf.comp.Debugf("inferring generic function type argument %v: currently inferred {untyped %v}, must match {untyped %v}", name, inferred.Untyped, kind)
	inf.combineUntyped(node, &inferred, inferType{Untyped: kind})
	inf.inferred[name] = inferred
}

func (inf *inferFuncType) combineUntyped(node ast.Expr, inferred *inferType, with inferType) {
	ikind := inferred.Untyped
	wkind := with.Untyped
	if ikind == untyped.None {
		ikind = wkind
	} else if wkind != untyped.None && ikind != wkind {
		switch ikind {
		case untyped.Bool:
		case untyped.Int:
			switch wkind {
			case untyped.Rune, untyped.Float, untyped.Complex:
				ikind = wkind
			}
		case untyped.Rune:
			switch wkind {
			case untyped.Float, untyped.Complex:
				ikind = wkind
			}
		case untyped.Float:
			switch wkind {
			case untyped.Complex:
				ikind = wkind
			}
		case untyped.Complex:
		case untyped.String:
		}
	}
	// if the conversion is unsupported, it will fail later on
	inferred.Untyped = ikind
}

// partially infer type of generic function for an interface parameter
func (inf *inferFuncType) interfaceType(node *ast.InterfaceType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	// TODO
	return inf.unimplemented(node, targ)
}

// partially infer type of generic function for a map parameter
func (inf *inferFuncType) mapType(node *ast.MapType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	inf.is(node, targ, r.Map)
	inf.arg(node.Key, targ.Key(), true)
	return node.Value, targ.Elem(), true
}

// partially infer type of generic function for an imported type
func (inf *inferFuncType) selector(node *ast.SelectorExpr, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	// TODO
	return inf.unimplemented(node, targ)
}

// partially infer type of generic function for a struct parameter
func (inf *inferFuncType) structType(node *ast.StructType, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {
	// TODO
	return inf.unimplemented(node, targ)
}

// partially infer type of generic function for a generic parameter
func (inf *inferFuncType) genericType(node *ast.IndexExpr, targ xr.Type, exact bool) (ast.Expr, xr.Type, bool) {

	var sym *Symbol
	// node is foo#[T1,T2...] => set name = foo
	if name, ok := node.X.(*ast.Ident); ok {
		sym = inf.comp.TryResolve(name.Name)
	} else {
		inf.fail(node, targ)
	}
	if sym == nil {
		st := &inf.comp.Stringer
		st.ErrorAt(inf.DeclPos(), "type inference: instantiating generic %v\n\t%v: required by function call: %v\n\t%v: undefined identifier %v in expression: %v",
			inf, st.Fileset.Position(inf.CallPos()), inf.call, st.Fileset.Position(node.Pos()), node.X, node)
	}
	class := sym.Desc.Class()
	if class != GenericFuncBind && class != GenericTypeBind {
		inf.fail(node, targ)
	}
	var elts []ast.Expr
	if complit, ok := node.Index.(*ast.CompositeLit); ok && complit.Type == nil {
		// node is foo#[T1,T2...] => set elts = {T1,T2...}
		elts = complit.Elts
	} else {
		inf.fail(node, targ)
	}
	_ = elts
	return inf.unimplemented(node, targ)
}

func (inf *inferFuncType) is(node ast.Expr, targ xr.Type, kind r.Kind) {
	if targ.Kind() != kind {
		inf.fail(node, targ)
	}
}

func (inf *inferFuncType) fail(node ast.Expr, targ I) {
	st := &inf.comp.Stringer
	st.ErrorAt(inf.DeclPos(), "type inference: instantiating generic %v\n\t%v: required by function call: %v\n\t%v: generic parameter <%v> cannot match argument type <%v>",
		inf, st.Fileset.Position(inf.CallPos()), inf.call, st.Fileset.Position(node.Pos()), node, targ)
}

func (inf *inferFuncType) fail3(node ast.Expr, tinferred *inferType, targ xr.Type) {
	inf.comp.ErrorAt(inf.CallPos(),
		"type inference: in %v, generic parameter <%v> cannot match both <%v> and <%v>: %v",
		inf, node, tinferred, targ, inf.call)
}

func (inf *inferFuncType) unimplemented(node ast.Expr, targ I) (ast.Expr, xr.Type, bool) {
	inf.comp.ErrorAt(inf.CallPos(), "unimplemented type inference: in %v, generic parameter <%v> with argument type <%v>: %v",
		inf, node, targ, inf.call)
	return nil, nil, false
}

var chandirs = map[ast.ChanDir]r.ChanDir{
	ast.RECV:            r.RecvDir,
	ast.SEND:            r.SendDir,
	ast.RECV | ast.SEND: r.BothDir,
}

func reflectChanDir(dir ast.ChanDir) r.ChanDir {
	return chandirs[dir]
}
