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
 * call.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"bytes"
	"fmt"
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

type Call struct {
	OutTypes []r.Type
	Fun      *Expr
	Funvar   *Var // in case function is an identifier (a variable or function name)
	Args     []*Expr
	Argvars  []*Var // in case arguments are variable or function names
	Argfuns  []func(*Env) r.Value
}

// CallExpr compiles a function call
func (c *Comp) CallExpr(node *ast.CallExpr) *Expr {
	call := c.callExpr(node)
	expr := &Expr{}
	switch len(call.OutTypes) {
	case 0:
		expr.Types = call.OutTypes
		expr.Fun = call_ret0(call)
	case 1:
		expr.Type = call.OutTypes[0]
		expr.Fun = call_ret1(call)
	default:
		expr.Types = call.OutTypes
		expr.Fun = call_ret2plus(call)
	}
	return expr
}

// callExpr compiles the common part between CallExpr and Go statement
func (c *Comp) callExpr(node *ast.CallExpr) *Call {
	fun := c.Expr(node.Fun)
	t := fun.Type
	if t.Kind() != r.Func {
		c.Errorf("call of non-function: %v <%v>", node.Fun, t)
		return nil
	}
	if t.IsVariadic() {
		c.Errorf("unimplemented: call to variadic function: %v <%v>", node.Fun, t)
		return nil
	}
	var funvar *Var
	if ident, ok := UnwrapTrivialNode(node.Fun).(*ast.Ident); ok {
		upn, bind := c.Resolve(ident.Name)
		if bind.Desc.Class() != ConstBind {
			funvar = &Var{Upn: upn, Desc: bind.Desc, Type: bind.Type}
		}
	}

	args := c.Exprs(node.Args)
	n := t.NumIn()
	// TODO support funcAcceptsNArgs(funcReturnsNValues())
	if n != len(args) {
		return c.badCallArgNum(node.Fun, t, args)
	}
	argvars := make([]*Var, n)
	argfuns := make([]func(*Env) r.Value, n)
	for i, arg := range args {
		ti := t.In(i)
		if arg.Const() {
			arg.ConstTo(ti)
		} else if arg.Type != ti && !arg.Type.AssignableTo(ti) {
			c.Errorf("cannot use <%v> as <%v> in argument to %v", arg.Type, ti, node.Fun)
		}
		argfuns[i] = arg.AsX1()
		if ident, ok := UnwrapTrivialNode(node.Args[i]).(*ast.Ident); ok {
			upn, bind := c.Resolve(ident.Name)
			if bind.Desc.Class() != ConstBind {
				argvars[i] = &Var{Upn: upn, Desc: bind.Desc, Type: bind.Type}
			}
		}
	}
	ret := &Call{Fun: fun, Funvar: funvar, Args: args, Argfuns: argfuns, Argvars: argvars}
	nout := t.NumOut()
	types := make([]r.Type, nout)
	for i := 0; i < nout; i++ {
		types[i] = t.Out(i)
	}
	ret.OutTypes = types
	return ret
}

// cannot optimize much here... fast_interpreter ASSUMES that expressions
// returning multiple values actually return (reflect.Value, []reflect.Value)
func call_ret2plus(callexpr *Call) I {
	expr := callexpr.Fun
	exprfun := expr.AsX1()
	argfuns := callexpr.Argfuns
	var call func(*Env) (r.Value, []r.Value)
	// slightly optimize fun() (tret0, tret1)
	switch expr.Type.NumIn() {
	case 0:
		call = func(env *Env) (r.Value, []r.Value) {
			funv := exprfun(env)
			retv := funv.Call(ZeroValues)
			return retv[0], retv
		}
	case 1:
		argfun := argfuns[0]
		call = func(env *Env) (r.Value, []r.Value) {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			retv := funv.Call(argv)
			return retv[0], retv
		}
	case 2:
		call = func(env *Env) (r.Value, []r.Value) {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
			}
			retv := funv.Call(argv)
			return retv[0], retv
		}
	case 3:
		call = func(env *Env) (r.Value, []r.Value) {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
				argfuns[2](env),
			}
			retv := funv.Call(argv)
			return retv[0], retv
		}
	default:
		// general case
		call = func(env *Env) (r.Value, []r.Value) {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}
			retv := funv.Call(argv)
			return retv[0], retv
		}
	}
	return call
}

func (c *Comp) badCallArgNum(fun ast.Expr, t r.Type, args []*Expr) *Call {
	prefix := "not enough"
	n := t.NumIn()
	nargs := len(args)
	if nargs > n {
		prefix = "too many"
	}
	have := bytes.Buffer{}
	for i, arg := range args {
		if i == 0 {
			fmt.Fprintf(&have, "%v", arg.Type)
		} else {
			fmt.Fprintf(&have, ", %v", arg.Type)
		}
	}
	want := bytes.Buffer{}
	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Fprintf(&want, "%v", t.In(i))
		} else {
			fmt.Fprintf(&want, ", %v", t.In(i))
		}
	}
	c.Errorf("%s arguments in call to %v:\n\thave (%s)\n\twant (%s)", prefix, fun, have.Bytes(), want.Bytes())
	return nil
}
