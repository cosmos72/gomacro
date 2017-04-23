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
 * indexexpr.go
 *
 *  Created on Apr 23, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

// IndexExpr compiles obj[idx]
func (c *Comp) IndexExpr(node *ast.IndexExpr) *Expr {
	obj := c.Expr1(node.X)
	idx := c.Expr1(node.Index)
	if obj.Untyped() {
		obj.ConstTo(obj.DefaultType())
	}
	t := obj.Type
	var ret *Expr
	switch t.Kind() {
	case r.Array, r.Slice, r.String:
		ret = c.vectorIndex(node, obj, idx)
	case r.Map:
		//ret = c.mapIndex(node, obj, idx)
		c.Errorf("unimplemented: indexing on map <%v>: %v", t, node)
	case r.Ptr:
		if t.Elem().Kind() == r.Array {
			//ret = c.vectorPointerIndex(node, obj, idx)
			c.Errorf("unimplemented: indexing on pointer to array <%v>: %v", t, node)
			break
		}
		fallthrough
	default:
		c.Errorf("invalid operation: %v (type %v does not support indexing)", node, t)
	}
	if obj.Const() && idx.Const() {
		// constant propagation
		ret.EvalConst(CompileKeepUntyped)
	}
	return ret
}

// vectorIndex compiles obj[idx] where obj is an array or slice
func (c *Comp) vectorIndex(node *ast.IndexExpr, obj *Expr, idx *Expr) *Expr {
	idxconst := idx.Const()
	if idxconst {
		idx.ConstTo(TypeOfInt)
	} else if !idx.Type.AssignableTo(TypeOfInt) {
		c.Errorf("non-integer %s index: %v <%v>", obj.Type.Kind(), node.Index, idx.Type)
	}
	t := obj.Type
	if t.Kind() == r.String {
		return c.stringIndex(node, obj, idx)
	}
	t = t.Elem()
	objfun := obj.AsX1()
	var fun I
	if idxconst {
		i := idx.Value.(int)
		switch t.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				objv := objfun(env)
				return objv.Index(i).Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				objv := objfun(env)
				return int(objv.Index(i).Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				objv := objfun(env)
				return int8(objv.Index(i).Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				objv := objfun(env)
				return int16(objv.Index(i).Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				objv := objfun(env)
				return int32(objv.Index(i).Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				objv := objfun(env)
				return objv.Index(i).Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				objv := objfun(env)
				return uint(objv.Index(i).Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				objv := objfun(env)
				return uint8(objv.Index(i).Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				objv := objfun(env)
				return uint16(objv.Index(i).Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				objv := objfun(env)
				return uint32(objv.Index(i).Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				objv := objfun(env)
				return objv.Index(i).Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				objv := objfun(env)
				return uintptr(objv.Index(i).Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				objv := objfun(env)
				return float32(objv.Index(i).Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				objv := objfun(env)
				return objv.Index(i).Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				objv := objfun(env)
				return complex64(objv.Index(i).Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				objv := objfun(env)
				return objv.Index(i).Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				objv := objfun(env)
				return objv.Index(i).String()
			}
		default:
			fun = func(env *Env) r.Value {
				objv := objfun(env)
				return objv.Index(i)
			}
		}
	} else {
		idxfun := idx.WithFun().(func(*Env) int)
		switch t.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				objv := objfun(env)
				i := idxfun(env)
				return int(objv.Index(i).Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				objv := objfun(env)
				i := idxfun(env)
				return int8(objv.Index(i).Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				objv := objfun(env)
				i := idxfun(env)
				return int16(objv.Index(i).Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				objv := objfun(env)
				i := idxfun(env)
				return int32(objv.Index(i).Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				objv := objfun(env)
				i := idxfun(env)
				return uint(objv.Index(i).Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				objv := objfun(env)
				i := idxfun(env)
				return uint8(objv.Index(i).Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				objv := objfun(env)
				i := idxfun(env)
				return uint16(objv.Index(i).Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				objv := objfun(env)
				i := idxfun(env)
				return uint32(objv.Index(i).Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				objv := objfun(env)
				i := idxfun(env)
				return uintptr(objv.Index(i).Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				objv := objfun(env)
				i := idxfun(env)
				return float32(objv.Index(i).Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				objv := objfun(env)
				i := idxfun(env)
				return complex64(objv.Index(i).Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i).String()
			}
		default:
			fun = func(env *Env) r.Value {
				objv := objfun(env)
				i := idxfun(env)
				return objv.Index(i)
			}
		}
	}
	return exprFun(t, fun)
}

// stringIndex compiles obj[idx] where obj is a string
func (c *Comp) stringIndex(node *ast.IndexExpr, obj *Expr, idx *Expr) *Expr {
	idxfun := idx.WithFun().(func(*Env) int)
	objfun := obj.WithFun().(func(*Env) string)
	var fun func(env *Env) uint8
	if obj.Const() {
		str := obj.Value.(string)
		fun = func(env *Env) uint8 {
			i := idxfun(env)
			return str[i]
		}
	} else if idx.Const() {
		i := idx.Value.(int)
		fun = func(env *Env) uint8 {
			str := objfun(env)
			return str[i]
		}
	} else {
		fun = func(env *Env) uint8 {
			str := objfun(env)
			i := idxfun(env)
			return str[i]
		}
	}
	return exprUint8(fun)
}
