/*
 * gomacro - A Go intepreter with Lisp-like macros
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
 * identifier.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
	"go/constant"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func IFalse() bool {
	return false
}

func ITrue() bool {
	return true
}

func INone() {
}

func X1False(*Env) r.Value {
	return False
}

func X1True(*Env) r.Value {
	return True
}

func X1None(*Env) r.Value {
	return None
}

func X1Nil(*Env) r.Value {
	return Nil
}

func XVFalse(*Env) (r.Value, []r.Value) {
	return False, nil
}

func XVTrue(*Env) (r.Value, []r.Value) {
	return True, nil
}

func XVNone() (r.Value, []r.Value) {
	return None, nil
}

func XVNil() (r.Value, []r.Value) {
	return Nil, nil
}

func (e *Expr) TryAsPred() (value bool, fun func(*Env) bool, err bool) {
	if e.Untyped() {
		untyp := e.Value.(UntypedLit)
		if untyp.Kind != r.Bool {
			return false, nil, true
		}
		return constant.BoolVal(untyp.Obj), nil, false
	}
	if e.Type != TypeOfBool {
		return false, nil, true
	}
	if value, ok := e.Value.(bool); ok {
		return value, nil, false
	}
	switch fun := e.Fun.(type) {
	case func(*Env) bool:
		return false, fun, false
	case func(*Env) (r.Value, []r.Value):
		e.CheckX1()
		return false, func(env *Env) bool {
			ret, _ := fun(env)
			return ret.Bool()
		}, false
	default:
		fun1 := e.AsX1()
		return false, func(env *Env) bool {
			return fun1(env).Bool()
		}, false
	}
}

func (c *Comp) invalidPred(node ast.Expr, x *Expr) Stmt {
	return c.badPred("invalid", node, x)
}

func (c *Comp) badPred(reason string, node ast.Expr, x *Expr) Stmt {
	var t r.Type = nil
	if x.NumOut() != 0 {
		t = x.Out(0)
	}
	c.Errorf("%s boolean predicate, expecting <bool> expression, found <%v>: %v",
		reason, t, node)
	return nil
}

func (e *Expr) AsX() X {
	if e == nil || e.Const() {
		return nil
	}
	return AsX(e.Fun)
}

func AsX(any I) X {
	if isLiteral(any) {
		return nil
	}
	switch fun := any.(type) {
	case nil:
	case X:
		return fun
	case func(*Env):
		return fun
	case func(*Env) r.Value:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) (r.Value, []r.Value):
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) bool:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) int:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) int8:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) int16:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) int32:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) int64:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uint:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uint8:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uint16:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uint32:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uint64:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) float32:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) float64:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) complex64:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) complex128:
		return func(env *Env) {
			fun(env)
		}
	case func(*Env) string:
		return func(env *Env) {
			fun(env)
		}
	default:
		Errorf("unsupported expression, cannot convert to func(*Env): %v <%T>", any, any)
	}
	return nil
}

// CheckX1() panics if given expression cannot be used in single-value context,
// for example because it returns no value at all.
// It just prints a warning if expression returns multiple values.
func (e *Expr) CheckX1() {
	if e != nil && e.Const() {
		return
	}
	if e == nil || e.NumOut() == 0 {
		Errorf("expression returns no values, cannot convert to func(env *Env) r.Value")
		return
	} else if e.NumOut() > 1 {
		Warnf("expression returns %d values, using only the first one: %v", e.Types)
	}
}

func (e *Expr) AsX1() func(*Env) r.Value {
	if e.Const() {
		return AsX1(e.Value, CompileDefaults)
	}
	e.CheckX1()
	return AsX1(e.Fun, CompileDefaults)
}

func AsX1(any I, opts CompileOptions) func(*Env) r.Value {
	if isLiteral(any) {
		if opts&CompileKeepUntyped == 0 {
			if untyp, ok := any.(UntypedLit); ok {
				// late conversion of untyped constants to their default type
				any = untyp.ConstTo(untyp.DefaultType())
			}
		}
		v := r.ValueOf(any)
		return func(*Env) r.Value {
			return v
		}
	}
	switch fun := any.(type) {
	case nil:
		return nil
	case X:
		if fun == nil {
			break
		}
		return func(env *Env) r.Value {
			fun(env)
			return None
		}
	case func(*Env):
		if fun == nil {
			break
		}
		return func(env *Env) r.Value {
			fun(env)
			return None
		}
	case func(*Env) r.Value:
		return fun
	case func(*Env) (r.Value, []r.Value):
		return func(env *Env) r.Value {
			ret, _ := fun(env)
			return ret
		}
	case func(*Env) bool:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) int:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) int8:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) int16:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) int32:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) int64:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) uint:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) uint8:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) uint16:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) uint32:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) uint64:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) uintptr:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) float32:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) float64:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) complex64:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) complex128:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	case func(*Env) string:
		return func(env *Env) r.Value {
			return r.ValueOf(fun(env))
		}
	default:
		Errorf("unsupported expression, cannot convert to func(*Env) r.Value: %v <%T>", fun, fun)
	}
	return nil
}

func (e *Expr) AsXV(opts CompileOptions) func(*Env) (r.Value, []r.Value) {
	if e.Const() {
		return AsXV(e.Value, opts)
	} else {
		return AsXV(e.Fun, opts)
	}
}

func AsXV(any I, opts CompileOptions) func(*Env) (r.Value, []r.Value) {
	if isLiteral(any) {
		if opts&CompileKeepUntyped == 0 {
			if untyp, ok := any.(UntypedLit); ok {
				// late conversion of untyped constants to their default type
				any = untyp.ConstTo(untyp.DefaultType())
			}
		}
		v := r.ValueOf(any)
		return func(*Env) (r.Value, []r.Value) {
			return v, nil
		}
	}
	switch fun := any.(type) {
	case nil:
	case X:
		if fun == nil {
			break
		}
		return func(env *Env) (r.Value, []r.Value) {
			fun(env)
			return None, nil
		}
	case func(*Env):
		if fun == nil {
			break
		}
		return func(env *Env) (r.Value, []r.Value) {
			fun(env)
			return None, nil
		}
	case func(*Env) r.Value:
		return func(env *Env) (r.Value, []r.Value) {
			return fun(env), nil
		}
	case func(*Env) (r.Value, []r.Value):
		return fun
	case func(*Env) bool:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) int:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) int8:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) int16:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) int32:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) int64:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) uint:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) uint8:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) uint16:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) uint32:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) uint64:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) uintptr:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) float32:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) float64:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) complex64:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) complex128:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	case func(*Env) string:
		return func(env *Env) (r.Value, []r.Value) {
			return r.ValueOf(fun(env)), nil
		}
	default:
		Errorf("unsupported expression, cannot convert to func(*Env) (r.Value, []r.Value) : %v <%T>",
			any, any)
	}
	return nil
}

func (e *Expr) AsStmt() Stmt {
	if e == nil || e.Const() {
		return nil
	}
	return AsStmt(e.Fun)
}

func AsStmt(any I) Stmt {
	if isLiteral(any) {
		return nil
	}
	var ret func(env *Env) (Stmt, *Env)

	switch fun := any.(type) {
	case nil:
	case X:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env):
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) r.Value:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) (r.Value, []r.Value):
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) bool:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int8:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int16:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int32:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint8:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint16:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint32:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uintptr:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float32:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex64:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex128:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) string:
		ret = func(env *Env) (Stmt, *Env) {
			fun(env)
			env.IP++
			return env.Code[env.IP], env
		}
	default:
		Errorf("unsupported expression, cannot convert to Stmt: %v <%v>", any, r.TypeOf(any))
	}
	return ret
}
