/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * function.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	xr "github.com/cosmos72/gomacro/xreflect"
)

type funcMaker struct {
	nbinds      int
	nintbinds   int
	parambinds  []*Bind
	resultbinds []*Bind
	resultfuns  []I
	funcbody    func(*Env)
}

// DeclFunc compiles a function or method declaration
// For closure declarations, use FuncLit()
func (c *Comp) FuncDecl(funcdecl *ast.FuncDecl) {
	if funcdecl.Recv != nil {
		c.methodDecl(funcdecl)
		return
	}
	functype := funcdecl.Type
	t, paramnames, resultnames := c.TypeFunction(functype)

	// declare the function name and type before compiling its body: allows recursive functions/methods
	funcname := funcdecl.Name.Name
	funcbind := c.AddBind(funcname, FuncBind, t)
	funcclass := funcbind.Desc.Class()
	if funcclass != FuncBind {
		c.Errorf("internal error! function bind should have class '%v', found class '%v' for: %s <%v>", FuncBind, funcclass, funcname, t)
	}

	cf := NewComp(c)
	info, resultfuns := cf.funcBinds(functype, t, paramnames, resultnames)
	cf.Func = info

	body := funcdecl.Body
	if body != nil && len(body.List) != 0 {
		// in Go, function arguments/results and function body are in the same scope
		cf.List(body.List)
	}

	funcindex := funcbind.Desc.Index()
	if funcindex == NoIndex {
		// unnamed function. still compile it (to check for compile errors) but discard the compiled code
		return
	}

	// do NOT keep a reference to compile environment!
	funcbody := cf.Code.Exec()

	f := cf.funcCreate(t, info, resultfuns, funcbody)

	// a function declaration is a statement:
	// executing it creates the function in the runtime environment
	stmt := func(env *Env) (Stmt, *Env) {
		fun := f(env)
		// Debugf("setting env.Binds[%d] = %v <%v>", funcindex, fun.Interface(), fun.Type())
		env.Binds[funcindex] = fun
		env.IP++
		return env.Code[env.IP], env
	}
	c.Code.Append(stmt)
}

func (c *Comp) methodAdd(funcdecl *ast.FuncDecl, t xr.Type) (methodindex int, methods *[]r.Value) {
	name := funcdecl.Name.Name
	trecv := t.In(0)
	if trecv.Kind() == r.Ptr && !trecv.Named() {
		// receiver is an unnamed pointer type. add the method to its element type
		trecv = trecv.Elem()
	}

	panicking := true
	defer func() {
		if panicking {
			rec := recover()
			c.Errorf("error adding method %s <%v> to type <%v>\n\t%v", name, t, trecv, rec)
		}
	}()
	n1 := trecv.NumMethod()
	methodindex = trecv.AddMethod(name, t)
	n2 := trecv.NumMethod()
	if n1 == n2 {
		c.Warnf("redefined method: %s.%s", trecv.Name(), name)
	}
	methods = trecv.GetMethods()
	panicking = false
	return
}

// methodDecl compiles a method declaration
func (c *Comp) methodDecl(funcdecl *ast.FuncDecl) {
	n := len(funcdecl.Recv.List)
	if n != 1 {
		c.Errorf("invalid function/method declaration: expecting one receiver or nil, found %d receivers: func %v %s(/*...*/)",
			n, funcdecl.Recv, funcdecl.Name)
		return
	}
	recvdecl := funcdecl.Recv.List[0]

	functype := funcdecl.Type
	t, paramnames, resultnames := c.TypeFunctionOrMethod(recvdecl, functype)

	// declare the method name and type before compiling its body: allows recursive methods
	methodindex, methods := c.methodAdd(funcdecl, t)

	cf := NewComp(c)
	info, resultfuns := cf.funcBinds(functype, t, paramnames, resultnames)
	cf.Func = info

	body := funcdecl.Body
	if body != nil && len(body.List) != 0 {
		// in Go, function arguments/results and function body are in the same scope
		cf.List(body.List)
	}
	// do NOT keep a reference to compile environment!
	funcbody := cf.Code.Exec()
	f := cf.funcCreate(t, info, resultfuns, funcbody)

	// a method declaration is a statement:
	// executing it sets the method value in the receiver type
	stmt := func(env *Env) (Stmt, *Env) {
		(*methods)[methodindex] = f(env)
		env.IP++
		return env.Code[env.IP], env
	}
	c.Code.Append(stmt)
}

// FuncLit compiles a function literal, i.e. a closure.
// For functions or methods declarations, use FuncDecl()
func (c *Comp) FuncLit(funclit *ast.FuncLit) *Expr {
	functype := funclit.Type
	t, paramnames, resultnames := c.TypeFunction(functype)

	cf := NewComp(c)
	info, resultfuns := cf.funcBinds(functype, t, paramnames, resultnames)
	cf.Func = info

	body := funclit.Body
	if body != nil && len(body.List) != 0 {
		// in Go, function arguments/results and function body are in the same scope
		cf.List(body.List)
	}
	// do NOT keep a reference to compile environment!
	funcbody := cf.Code.Exec()

	f := cf.funcCreate(t, info, resultfuns, funcbody)

	// a function literal is an expression:
	// executing it returns the function
	return exprX1(t, f)
}

// prepare the function parameter binds, result binds and FuncInfo
func (c *Comp) funcBinds(functype *ast.FuncType, t xr.Type, paramnames, resultnames []string) (info *FuncInfo, resultfuns []I) {

	parambinds := c.funcParamBinds(functype, t, paramnames)

	resultbinds, resultfuns := c.funcResultBinds(functype, t, resultnames)
	namedresults := true
	for _, resultname := range resultnames {
		if len(resultname) == 0 {
			namedresults = false
		}
	}
	return &FuncInfo{
		Params:       parambinds,
		Results:      resultbinds,
		NamedResults: namedresults,
	}, resultfuns
}

// prepare the function parameter binds
func (c *Comp) funcParamBinds(functype *ast.FuncType, t xr.Type, names []string) []*Bind {
	n := t.NumIn()
	binds := make([]*Bind, n)
	var namedparams, unnamedparams bool
	for i := 0; i < n; i++ {
		// names[i] == "" means that argument is unnamed, and thus ignored inside the function.
		// change to "_" so that AddBind will not allocate a bind for it - correct optimization...
		// just remember to check for such case when creating the function
		name := names[i]
		if name == "" {
			name = "_"
			unnamedparams = true
		} else {
			namedparams = true
		}
		if namedparams && unnamedparams {
			c.Errorf("cannot mix named and unnamed parameters in function declaration: %v", functype)
		}
		bind := c.AddBind(name, VarBind, t.In(i))
		binds[i] = bind
	}
	return binds
}

// prepare the function result binds
func (c *Comp) funcResultBinds(functype *ast.FuncType, t xr.Type, names []string) (binds []*Bind, funs []I) {
	n := t.NumOut()
	binds = make([]*Bind, n)
	funs = make([]I, n)
	var namedresults, unnamedresults bool
	for i, n := 0, t.NumOut(); i < n; i++ {
		// names[i] == "" means that result is unnamed.
		// we must still allocate a bind for it.
		name := names[i]
		if name == "" {
			unnamedresults = true
		} else {
			namedresults = true
		}
		if namedresults && unnamedresults {
			c.Errorf("cannot mix named and unnamed results in function declaration: %v", functype)
		}
		bind := c.DeclVar0(name, t.Out(i), nil)
		binds[i] = bind
		// compile the extraction of results from runtime env
		sym := bind.AsSymbol(0)
		funs[i] = c.Symbol(sym).WithFun()
	}
	return
}

// actually create the function
func (c *Comp) funcCreate(t xr.Type, info *FuncInfo, resultfuns []I, funcbody func(*Env)) func(*Env) r.Value {
	m := &funcMaker{
		nbinds:      c.BindNum,
		nintbinds:   c.IntBindNum,
		parambinds:  info.Params,
		resultbinds: info.Results,
		resultfuns:  resultfuns,
		funcbody:    funcbody,
	}

	var fun func(*Env) r.Value
	switch t.NumIn() {
	case 0:
		switch t.NumOut() {
		case 0:
			fun = c.func0ret0(t, m)
		case 1:
			fun = c.func0ret1(t, m)
		}
	case 1:
		switch t.NumOut() {
		case 0:
			fun = c.func1ret0(t, m)
		case 1:
			fun = c.func1ret1(t, m)
		}
	case 2:
		switch t.NumOut() {
		case 0:
			fun = c.func2ret0(t, m)
		}
	}
	if fun == nil {
		fun = c.funcGeneric(t, m)
	}
	return fun
}

// fallback: create a non-optimized function
func (c *Comp) funcGeneric(t xr.Type, m *funcMaker) func(*Env) r.Value {

	paramdecls := make([]func(*Env, r.Value), len(m.parambinds))
	for i, bind := range m.parambinds {
		if bind.Desc.Index() != NoIndex {
			paramdecls[i] = c.DeclBindRuntimeValue(bind)
		}
	}
	resultexprs := make([]func(*Env) r.Value, len(m.resultfuns))
	for i, resultfun := range m.resultfuns {
		resultexprs[i] = funAsX1(resultfun, m.resultbinds[i].Type)
	}

	// do NOT keep a reference to funcMaker
	nbinds := m.nbinds
	nintbinds := m.nintbinds
	funcbody := m.funcbody
	rtype := t.ReflectType()

	return func(env *Env) r.Value {
		// function is closed over the env used to DECLARE it
		env.MarkUsedByClosure()
		return r.MakeFunc(rtype, func(args []r.Value) []r.Value {
			env := NewEnv(env, nbinds, nintbinds)

			if funcbody != nil {
				// copy runtime arguments into allocated binds
				for i, decl := range paramdecls {
					if decl != nil {
						// decl == nil means the argument is ignored inside the function
						decl(env, args[i])
					}
				}
				// execute the body
				funcbody(env)
			}
			// read results from allocated binds and return them
			rets := make([]r.Value, len(resultexprs))
			for i, expr := range resultexprs {
				rets[i] = expr(env)
			}
			env.FreeEnv()
			return rets
		})
	}
}

func declBindRuntimeValueNop(*Env, r.Value) {
}
