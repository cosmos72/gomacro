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
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

type Call struct {
	Fun      *Expr
	Args     []*Expr
	OutTypes []r.Type
}

func newCall1(fun *Expr, arg *Expr, outtypes ...r.Type) *Call {
	return &Call{
		Fun:      fun,
		Args:     []*Expr{arg},
		OutTypes: outtypes,
	}
}

func (call *Call) MakeArgfuns() []func(*Env) r.Value {
	args := call.Args
	argfuns := make([]func(*Env) r.Value, len(args))
	for i, arg := range args {
		argfuns[i] = arg.AsX1()
	}
	return argfuns
}

// CallExpr compiles a function call
func (c *Comp) CallExpr(node *ast.CallExpr) *Expr {
	call := c.callExpr(node)
	maxdepth := c.Depth
	expr := &Expr{}
	switch len(call.OutTypes) {
	case 0:
		expr.Types = call.OutTypes
		expr.Fun = call_ret0(call, maxdepth)
	case 1:
		expr.Type = call.OutTypes[0]
		expr.Fun = call_ret1(call, maxdepth)
	default:
		expr.Types = call.OutTypes
		expr.Fun = call_ret2plus(call, maxdepth)
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
	args := c.Exprs(node.Args)
	n := t.NumIn()
	// TODO support funcAcceptsNArgs(funcReturnsNValues())
	if node.Ellipsis != token.NoPos {
		c.Errorf("unimplemented: variadic call, i.e. '...' after last argument: %v", node)
	}
	if n != len(args) {
		return c.badCallArgNum(node.Fun, t, args)
	}
	for i, arg := range args {
		ti := t.In(i)
		if arg.Const() {
			arg.ConstTo(ti)
		} else if arg.Type != ti && !arg.Type.AssignableTo(ti) {
			c.Errorf("cannot use <%v> as <%v> in argument to %v", arg.Type, ti, node.Fun)
		}
	}

	nout := t.NumOut()
	types := make([]r.Type, nout)
	for i := 0; i < nout; i++ {
		types[i] = t.Out(i)
	}
	return &Call{Fun: fun, Args: args, OutTypes: types}
}

/*
func (c *Comp) extractSymbols(args []ast.Expr) []*Symbol {
	n := len(args)
	syms := make([]*Symbol, n)
	for i, arg := range args {
		syms[i] = c.extractSymbol(arg)
	}
	return syms
}

func (c *Comp) extractSymbol(arg ast.Expr) *Symbol {
	if ident, ok := UnwrapTrivialNode(arg).(*ast.Ident); ok {
		return c.TryResolve(ident.Name)
	}
	return nil
}
*/

// mandatory optimization: fast_interpreter ASSUMES that expressions
// returning bool, int, uint, float, complex, string do NOT wrap them in reflect.Value
func call_ret0(c *Call, maxdepth int) func(env *Env) {
	exprfun := c.Fun.AsX1()
	var call func(*Env)
	// optimize fun(t1, t2)
	switch c.Fun.Type.NumIn() {
	case 0:
		call = call0ret0(c, maxdepth)
	case 1:
		call = call1ret0(c, maxdepth)
	case 2:
		call = call2ret0(c, maxdepth)
	case 3:
		argfuns := c.MakeArgfuns()
		call = func(env *Env) {
			funv := exprfun(env)
			argv := []r.Value{
				argfuns[0](env),
				argfuns[1](env),
				argfuns[2](env),
			}
			funv.Call(argv)
		}
	}
	if call == nil {
		argfuns := c.MakeArgfuns()
		call = func(env *Env) {
			funv := exprfun(env)
			argv := make([]r.Value, len(argfuns))
			for i, argfun := range argfuns {
				argv[i] = argfun(env)
			}
			funv.Call(argv)
		}
	}
	return call
}

// mandatory optimization: fast_interpreter ASSUMES that expressions
// returning bool, int, uint, float, complex, string do NOT wrap them in reflect.Value
func call_ret1(c *Call, maxdepth int) I {
	expr := c.Fun
	var call I
	// optimize fun(tret) tret
	switch expr.Type.NumIn() {
	case 0:
		call = call0ret1(c, maxdepth)
	case 1:
		call = call1ret1(c, maxdepth)
	case 2:
		call = call2ret1(c, maxdepth)
	default:
		call = callnret1(c, maxdepth)
	}
	return call
}

// cannot optimize much here... fast_interpreter ASSUMES that expressions
// returning multiple values actually return (reflect.Value, []reflect.Value)
func call_ret2plus(callexpr *Call, maxdepth int) I {
	expr := callexpr.Fun
	exprfun := expr.AsX1()
	argfuns := callexpr.MakeArgfuns()
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

func (c *Comp) badBuiltinCallArgNum(name string, n1 int, n2 int, args []ast.Expr) *Call {
	prefix := "not enough"
	nargs := len(args)
	if nargs > n2 {
		prefix = "too many"
	}
	str := fmt.Sprintf("%d", n1)
	if n2 != n1 {
		str = fmt.Sprintf("%s or %d", str, n2)
	}
	c.Errorf("%s arguments in call to builtin %s(): expecting %s, found %d: %v", prefix, name, str, nargs, args)
	return nil
}

func (c *Comp) badBuiltinCallArgType(name string, arg ast.Expr, t r.Type) *Call {
	c.Errorf("invalid argument %v <%v> for builtin %s()", arg, t, name)
	return nil
}
