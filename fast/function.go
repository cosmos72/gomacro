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
)

type funcMaker struct {
	nbinds      int
	nintbinds   int
	parambinds  []*Bind
	resultbinds []*Bind
	resultfuns  []I
	funcbody    func(*Env)
}

// DeclFunc compiles a function or method declaration (does not support closure declarations)
func (c *Comp) DeclFunc(funcdecl *ast.FuncDecl, functype *ast.FuncType, body *ast.BlockStmt) {
	var recvdecl *ast.Field
	if funcdecl.Recv != nil {
		n := len(funcdecl.Recv.List)
		if n != 1 {
			c.Errorf("invalid function/method declaration: expecting one receiver or nil, found %d receivers: func %v %s(/*...*/)",
				n, funcdecl.Recv, funcdecl.Name)
			return
		}
		recvdecl = funcdecl.Recv.List[0]
	}
	ismethod := recvdecl != nil
	if ismethod {
		c.Errorf("unimplemented: method declaration: %v", functype)
		return
	}
	_, t, paramnames, resultnames := c.TypeFunctionOrMethod(recvdecl, functype)

	// declare the function name and type before compiling its body: allows recursive functions
	funcname := funcdecl.Name.Name
	funcbind := c.AddBind(funcname, FuncBind, t)
	funcclass := funcbind.Desc.Class()
	if funcclass != FuncBind {
		c.Errorf("internal error! function bind should have class '%v', found class '%v' for: %s <%v>", FuncBind, funcclass, funcname, t)
	}

	cf := NewComp(c)

	// prepare the function parameters
	n := t.NumIn()
	parambinds := make([]*Bind, n)
	for i := 0; i < n; i++ {
		// paramNames[i] == "_" means that argument is ignored inside the function.
		// AddBind will not allocate a bind for it - correct optimization...
		// just remember to check for such case below
		name := paramnames[i]
		bind := cf.AddBind(name, VarBind, t.In(i))
		parambinds[i] = bind
		if bind.Desc.Index() == NoIndex {
			continue
		}
	}

	// prepare the function results
	n = t.NumOut()
	resultbinds := make([]*Bind, n)
	resultfuns := make([]I, n)
	var namedresults, unnamedresults bool
	for i, n := 0, t.NumOut(); i < n; i++ {
		// resultnames[i] == "_" means that result is unnamed.
		// we must still allocate a bind for it.
		name := resultnames[i]
		if name == "_" {
			name = ""
			unnamedresults = true
		} else {
			namedresults = true
		}
		if namedresults && unnamedresults {
			c.Errorf("cannot mix named and unnamed results in function declaration: %v", functype)
		}
		bind := cf.DeclVar0(name, t.Out(i), nil)
		resultbinds[i] = bind
		// compile the extraction of results from runtime env
		sym := bind.AsSymbol(0)
		resultfuns[i] = c.Symbol(sym).WithFun()
	}
	cf.Func = &FuncInfo{
		Params:  parambinds,
		Results: resultbinds,
	}

	if body != nil && len(body.List) != 0 {
		// in Go, function arguments/results and function body are in the same scope
		cf.Block0(body.List...)
	}

	funcindex := funcbind.Desc.Index()
	if funcindex == NoIndex {
		// unnamed function. still compile it (to check for compile errors) but discard the compiled code
		return
	}

	// do NOT keep a reference to compile environment!
	funcbody := cf.Code.Exec()
	nbinds := cf.BindNum
	nintbinds := cf.IntBindNum

	m := funcMaker{
		nbinds:      nbinds,
		nintbinds:   nintbinds,
		parambinds:  parambinds,
		resultbinds: resultbinds,
		resultfuns:  resultfuns,
		funcbody:    funcbody,
	}
	makefunc := cf.declFunc(t, &m)

	// a function declaration is a statement:
	// executing it creates the function in the runtime environment
	stmt := func(env *Env) (Stmt, *Env) {
		fun := makefunc(env)
		// Debugf("setting env.Binds[%d] = %v <%v>", funcindex, fun.Interface(), fun.Type())
		env.Binds[funcindex] = fun
		env.IP++
		return env.Code[env.IP], env
	}
	c.Code.Append(stmt)
}

func (c *Comp) declFunc(t r.Type, m *funcMaker) func(*Env) r.Value {
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
func (c *Comp) funcGeneric(t r.Type, m *funcMaker) func(*Env) r.Value {

	paramdecls := make([]func(*Env, r.Value), len(m.parambinds))
	for i, bind := range m.parambinds {
		if bind.Desc.Index() != NoIndex {
			paramdecls[i] = c.DeclBindRuntimeValue(bind)
		}
	}
	resultexprs := make([]func(*Env) r.Value, len(m.resultfuns))
	for i, resultfun := range m.resultfuns {
		resultexprs[i] = funAsX1(resultfun, CompileDefaults)
	}

	// do NOT keep a reference to funcMaker
	nbinds := m.nbinds
	nintbinds := m.nintbinds
	funcbody := m.funcbody

	return func(env *Env) r.Value {
		// function is closed over the env used to DECLARE it
		env.MarkUsedByClosure()
		return r.MakeFunc(t, func(args []r.Value) []r.Value {
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
