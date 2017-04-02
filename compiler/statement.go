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
 * statement.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package compiler

import (
	r "reflect"
)

func VarSet(idx int, expr X) X {
	return func(env *Env) (r.Value, []r.Value) {
		val, _ := expr(env)
		place := env.Binds[idx]
		place.Set(val)
		return place, nil
	}
}

func VarIncInt(idx int) X {
	return func(env *Env) (r.Value, []r.Value) {
		v := env.Binds[idx]
		v.SetInt(v.Int() + 1)
		return v, nil
	}
}

func If(pred XBool, then, els X) X {
	if els != nil {
		return func(env *Env) (r.Value, []r.Value) {
			if pred(env) {
				return then(env)
			} else {
				return els(env)
			}
		}
	} else {
		return func(env *Env) (r.Value, []r.Value) {
			if pred(env) {
				return then(env)
			} else {
				return None, nil
			}
		}
	}
}

func For(init X, pred XBool, post X, body X) X {
	if init == nil && post == nil {
		return func(env *Env) (r.Value, []r.Value) {
			for pred(env) {
				body(env)
			}
			return None, nil
		}

	} else {
		if init == nil || post == nil {
			panic("invalid for(): init and post must be both present, or both omitted")
		}
		return func(env *Env) (r.Value, []r.Value) {
			for init(env); pred(env); post(env) {
				body(env)
			}
			return None, nil
		}
	}
}

func (c *Comp) Block(list ...X) X {
	switch len(list) {
	case 0:
		return XNop
	case 1:
		return list[0]
	case 2:
		return func(env *Env) (r.Value, []r.Value) {
			list[0](env)
			return list[1](env)
		}
	default:
		return func(env *Env) (r.Value, []r.Value) {
			n_1 := len(list) - 1
			for i := 0; i < n_1; i++ {
				list[i](env)
			}
			return list[n_1](env)
		}
	}
}

func Return(exprs ...X) X {
	switch n := len(exprs); n {
	case 0:
		return XNop
	case 1:
		expr := exprs[0]
		// return foo() returns *all* the values returned by foo, not just the first one
		return func(env *Env) (r.Value, []r.Value) {
			ret, rets := expr(env)
			panic(SReturn{ret, rets})
		}
	default:
		return func(env *Env) (r.Value, []r.Value) {
			n := len(exprs)
			rets := make([]r.Value, n)
			for i, value := range exprs {
				rets[i], _ = value(env)
			}
			ret0 := None
			if len(rets) > 0 {
				ret0 = rets[0]
			}
			panic(SReturn{ret0, rets})
		}
	}
}
