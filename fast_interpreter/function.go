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

package fast_interpreter

import (
	"go/ast"
	r "reflect"
	"unsafe"

	_ "github.com/cosmos72/gomacro/base"
)

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
	parambinds := make([]Bind, n)
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
	resultbinds := make([]Bind, n)
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
			c.Errorf("")
		}
		bind := cf.DeclVar0(name, t.Out(i), nil)
		resultbinds[i] = bind
		// compile the extraction of results from runtime env
		resultfuns[i] = c.IdentBind(0, bind).WithFun()
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
	funcbody := cf.Code.AsX()
	nbinds := cf.BindNum
	nintbinds := cf.IntBindNum

	makefunc := c.funcOptimized(nbinds, nintbinds, t,
		paramnames, parambinds,
		resultbinds, resultfuns, funcbody)

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

// TODO actually optimize, i.e. create specialized functions without using reflect.MakeFunc
func (c *Comp) funcOptimized(nbinds int, nintbinds int, t r.Type,
	paramnames []string, parambinds []Bind,
	resultbinds []Bind, resultfuns []I,
	funcbody func(*Env)) func(*Env) r.Value {

	switch r.Zero(t).Interface().(type) {
	case func(int) int:
		param0index := parambinds[0].Desc.Index()
		result0index := resultbinds[0].Desc.Index()
		if funcbody == nil {
			return func(env *Env) r.Value {
				// function is closed over the env used to DECLARE it
				return r.ValueOf(func(int) int {
					return 0
				})
			}
		}
		if param0index == NoIndex {
			return func(env *Env) r.Value {
				// function is closed over the env used to DECLARE it
				return r.ValueOf(func(int) int {
					env := NewEnv4Func(env, nbinds, nintbinds)
					// arg0 is ignored

					// execute the body
					funcbody(env)

					// read results from allocated binds and return them
					ret0 := *(*int)(unsafe.Pointer(&env.IntBinds[result0index]))

					env.FreeEnv()
					return ret0
				})
			}
		}
		return func(env *Env) r.Value {
			// function is closed over the env used to DECLARE it
			return r.ValueOf(func(arg0 int) int {
				env := NewEnv4Func(env, nbinds, nintbinds)

				// copy runtime arguments into allocated binds
				*(*int)(unsafe.Pointer(&env.IntBinds[param0index])) = arg0

				// execute the body
				funcbody(env)

				// read results from allocated binds and return them
				ret0 := *(*int)(unsafe.Pointer(&env.IntBinds[result0index]))

				env.FreeEnv()
				return ret0
			})
		}
	}

	paramdecls := make([]func(*Env, r.Value), len(parambinds))
	for i, bind := range parambinds {
		if bind.Desc.Index() != NoIndex {
			paramdecls[i] = c.DeclBindRuntimeValue(paramnames[i], parambinds[i])
		}
	}
	resultexprs := make([]func(*Env) r.Value, len(resultfuns))
	for i, resultfun := range resultfuns {
		resultexprs[i] = FunAsX1(resultfun, CompileDefaults)
	}

	return func(env *Env) r.Value {
		// function is closed over the env used to DECLARE it
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
