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
 * slice.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

// SliceExpr compiles slice[lo:hi] and slice[lo:hi:max]
func (c *Comp) SliceExpr(node *ast.SliceExpr) *Expr {
	e := c.Expr1(node.X)
	if e.Const() {
		e.ConstTo(e.DefaultType())
	}
	if e.Type.Kind() == r.Array {
		c.sliceArrayMustBeAddressable(node, e)
	}
	var lo, hi, max *Expr
	if node.Low != nil {
		lo = c.Expr1(node.Low)
		if lo.Const() {
			lo.ConstTo(TypeOfInt)
		} else if !lo.Type.AssignableTo(TypeOfInt) {
			c.Errorf("invalid slice index: expecting integer, found: %v <%v>", lo.Type, node.Low)
		}
	}
	if node.High != nil {
		hi = c.Expr1(node.High)
		if hi.Const() {
			hi.ConstTo(TypeOfInt)
		} else if !hi.Type.AssignableTo(TypeOfInt) {
			c.Errorf("invalid slice index: expecting integer, found: %v <%v>", hi.Type, node.High)
		}
	}
	if node.Max != nil {
		max = c.Expr1(node.Max)
		if max.Const() {
			max.ConstTo(TypeOfInt)
		} else if !max.Type.AssignableTo(TypeOfInt) {
			c.Errorf("invalid slice index: expecting integer, found: %v <%v>", max.Type, node.Max)
		}
	}
	var ret *Expr
	if node.Slice3 {
		ret = c.slice3(node, e, lo, hi, max)
	} else {
		ret = c.slice2(node, e, lo, hi)
	}
	// constant propagation
	if e.Const() && (lo == nil || lo.Const()) && (hi == nil || hi.Const()) && (max == nil || max.Const()) {
		ret.EvalConst(CompileKeepUntyped)
	}
	return ret
}

// slice2 compiles slice[lo:hi]
func (c *Comp) slice2(node *ast.SliceExpr, e, lo, hi *Expr) *Expr {
	t := e.Type
	switch t.Kind() {
	case r.String:
		return c.sliceString(e, lo, hi)
	case r.Ptr:
		if t.Elem().Kind() != r.Array {
			break
		}
		fallthrough
	case r.Slice, r.Array:
		if t.Kind() == r.Ptr {
			t = t.Elem()
			objfun := e.AsX1()
			e = exprX1(t, func(env *Env) r.Value {
				return objfun(env).Elem()
			})
		}
		objfun := e.AsX1()
		if lo == nil {
			lo = exprValue(0)
		}
		var fun func(env *Env) r.Value
		if lo.Const() {
			lo := lo.Value.(int)
			if hi == nil {
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					return obj.Slice(lo, obj.Len())
				}
			} else if hi.Const() {
				hi := hi.Value.(int)
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					return obj.Slice(lo, hi)
				}
			} else {
				hifun := hi.WithFun().(func(*Env) int)
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					hi := hifun(env)
					return obj.Slice(lo, hi)
				}
			}
		} else {
			lofun := lo.WithFun().(func(*Env) int)
			if hi == nil {
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					lo := lofun(env)
					return obj.Slice(lo, obj.Len())
				}
			} else if hi.Const() {
				hi := hi.Value.(int)
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					lo := lofun(env)
					return obj.Slice(lo, hi)
				}
			} else {
				hifun := hi.WithFun().(func(*Env) int)
				fun = func(env *Env) r.Value {
					obj := objfun(env)
					lo := lofun(env)
					hi := hifun(env)
					return obj.Slice(lo, hi)
				}
			}
		}
		tout := r.SliceOf(t.Elem())
		return exprX1(tout, fun)
	}
	c.Errorf("cannot slice %v: %v", t, node)
	return nil
}

// sliceString compiles string[lo:hi]
func (c *Comp) sliceString(e, lo, hi *Expr) *Expr {
	objfun := e.WithFun().(func(*Env) string)
	var fun func(env *Env) string
	if lo == nil {
		if hi == nil {
			fun = objfun
		} else if hi.Const() {
			hi := hi.Value.(int)
			fun = func(env *Env) string {
				obj := objfun(env)
				return obj[:hi]
			}
		} else {
			hifun := hi.WithFun().(func(*Env) int)
			fun = func(env *Env) string {
				obj := objfun(env)
				hi := hifun(env)
				return obj[:hi]
			}
		}
	} else if lo.Const() {
		lo := lo.Value.(int)
		if hi == nil {
			fun = func(env *Env) string {
				obj := objfun(env)
				return obj[lo:]
			}
		} else if hi.Const() {
			hi := hi.Value.(int)
			fun = func(env *Env) string {
				obj := objfun(env)
				return obj[lo:hi]
			}
		} else {
			hifun := hi.WithFun().(func(*Env) int)
			fun = func(env *Env) string {
				obj := objfun(env)
				hi := hifun(env)
				return obj[lo:hi]
			}
		}
	} else {
		lofun := lo.WithFun().(func(*Env) int)
		if hi == nil {
			fun = func(env *Env) string {
				obj := objfun(env)
				lo := lofun(env)
				return obj[lo:]
			}
		} else if hi.Const() {
			hi := hi.Value.(int)
			fun = func(env *Env) string {
				obj := objfun(env)
				lo := lofun(env)
				return obj[lo:hi]
			}
		} else {
			hifun := hi.WithFun().(func(*Env) int)
			fun = func(env *Env) string {
				obj := objfun(env)
				lo := lofun(env)
				hi := hifun(env)
				return obj[lo:hi]
			}
		}
	}
	return exprString(fun)
}

// slice3 compiles slice[lo:hi:max]
func (c *Comp) slice3(node *ast.SliceExpr, e, lo, hi, max *Expr) *Expr {
	if lo == nil {
		lo = exprValue(0)
	}
	if hi == nil {
		c.Errorf("final index required in 3-index slice: %v", node)
	}
	if max == nil {
		c.Errorf("final index required in 3-index slice: %v", node)
	}
	t := e.Type
	switch t.Kind() {
	case r.String:
		c.Errorf("invalid operation %v (3-index slice of string)", node)
		return nil
	case r.Ptr:
		if t.Elem().Kind() != r.Array {
			break
		}
		fallthrough
	case r.Slice, r.Array:
		objfun := e.AsX1()
		lofun := lo.WithFun().(func(*Env) int)
		hifun := hi.WithFun().(func(*Env) int)
		maxfun := max.WithFun().(func(*Env) int)
		var fun func(env *Env) r.Value
		if t.Kind() == r.Ptr {
			t = t.Elem()
			fun = func(env *Env) r.Value {
				obj := objfun(env).Elem()
				lo := lofun(env)
				hi := hifun(env)
				max := maxfun(env)
				return obj.Slice3(lo, hi, max)
			}
		} else {
			fun = func(env *Env) r.Value {
				obj := objfun(env)
				lo := lofun(env)
				hi := hifun(env)
				max := maxfun(env)
				return obj.Slice3(lo, hi, max)
			}
		}
		tout := r.SliceOf(t.Elem())
		return exprX1(tout, fun)
	}
	c.Errorf("cannot slice %v: %v", t, node)
	return nil
}

func (c *Comp) sliceArrayMustBeAddressable(node *ast.SliceExpr, e *Expr) {
	c.Place(node.X)
}
