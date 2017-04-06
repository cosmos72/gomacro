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
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func ExprsToX(inits []*Expr) X {
	var funs []X

	for _, init := range inits {
		if !init.Const() {
			funs = append(funs, init.AsX())
		}
	}
	return ExprStmtsToX(funs)
}

func ExprStmtsToX(funs []X) X {
	funs = RemoveNils(funs)
	switch len(funs) {
	case 0:
		return nil
	case 1:
		return funs[0]
	case 2:
		return func(env *Env) {
			funs[0](env)
			funs[1](env)
		}
	default:
		return func(env *Env) {
			for _, fun := range funs {
				fun(env)
			}
		}
	}
}

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

func (e *Expr) AsPred() (flag bool, pred func(*Env) bool, err bool) {
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

func (e *Expr) AsX() X {
	if e == nil || e.Const() {
		return nil
	}
	return ToX(e.Fun)
}

func ToX(any I) X {
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
		errorf("unsupported expression, cannot convert to func(*Env): %v <%T>", any, any)
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
		errorf("expression returns no values, cannot convert to func(env *Env) r.Value")
		return
	} else if e.NumOut() > 1 {
		warnf("expression returns %d values, using only the first one: %v", e.Types)
	}
}

func (e *Expr) AsX1() func(*Env) r.Value {
	if e.Const() {
		return ToX1(e.Value)
	}
	e.CheckX1()
	return ToX1(e.Fun)
}

func ToX1(any I) func(*Env) r.Value {
	if isLiteral(any) {
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
		errorf("unsupported expression, cannot convert to func(*Env) r.Value: %v <%T>", fun, fun)
	}
	return nil
}

func (e *Expr) AsXV() func(*Env) (r.Value, []r.Value) {
	if e.Const() {
		return ToXV(e.Value)
	} else {
		return ToXV(e.Fun)
	}
}

func ToXV(expr I) func(*Env) (r.Value, []r.Value) {
	if isLiteral(expr) {
		v := r.ValueOf(expr)
		return func(*Env) (r.Value, []r.Value) {
			return v, nil
		}
	}
	switch fun := expr.(type) {
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
		errorf("unsupported expression, cannot convert to func(*Env) (r.Value, []r.Value) : %v <%T>",
			expr, expr)
	}
	return nil
}
