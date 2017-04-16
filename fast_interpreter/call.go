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

package fast_interpreter

import (
	"bytes"
	"fmt"
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

// CallExpr compiles a function call
func (c *Comp) CallExpr(node *ast.CallExpr) *Expr {
	expr := c.Expr(node.Fun)
	t := expr.Type
	if t.Kind() != r.Func {
		c.Errorf("call of non-function: %v <%v>", node.Fun, t)
		return nil
	}
	if t.IsVariadic() {
		c.Errorf("unimplemented: call to variadic function: %v <%v>", node.Fun, t)
		return nil
	}
	// TODO support funcAcceptsNArgs(funcReturnsNValues())
	args := c.Exprs(node.Args)
	n := t.NumIn()
	if n != len(args) {
		return c.badCallArgNum(node.Fun, t, args)
	}
	// TODO optimize for bool, int, uint... argument types and return types
	// and also for calls to identifiers
	argfuns := make([]func(*Env) r.Value, n)
	for i, arg := range args {
		ti := t.In(i)
		if arg.Const() {
			arg.ConstTo(ti)
		} else if arg.Type != ti && !arg.Type.AssignableTo(ti) {
			c.Errorf("cannot use <%v> as <%v> in argument to %v", arg.Type, ti, node.Fun)
		}
		argfuns[i] = arg.AsX1()
	}
	ret := &Expr{}
	nout := t.NumOut()
	switch nout {
	case 0:
		ret.Types = ZeroTypes
		ret.Fun = call_ret0(expr, args, argfuns)
	case 1:
		ret.Type = t.Out(0)
		ret.Fun = call_ret1(expr, args, argfuns)
	default:
		ret.Types = []r.Type{t.Out(0), t.Out(1)}
		ret.Fun = call_ret2plus(expr, args, argfuns)
	}
	return ret
}

// cannot optimize much here... fast_interpreter ASSUMES that expressions
// returning multiple values actually return (reflect.Value, []reflect.Value)
func call_ret2plus(expr *Expr, args []*Expr, argfuns []func(*Env) r.Value) I {
	exprfun := expr.AsX1()
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

func (c *Comp) badCallArgNum(fun ast.Expr, t r.Type, args []*Expr) *Expr {
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
