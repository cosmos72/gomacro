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
 * convert.go
 *
 *  Created on Apr 30, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"
)

// Convert compiles a type conversion
func (c *Comp) Convert(node ast.Expr, t r.Type) *Expr {
	e := c.Expr1(node)
	if e.Const() {
		e.ConstTo(t)
		return e
	} else if e.Type == t {
		return e
	} else if !e.Type.ConvertibleTo(t) {
		c.Errorf("cannot convert <%v> to <%v>: %v", node)
		return nil
	}
	fun := e.AsX1()
	var ret I
	switch t.Kind() {
	case r.Bool:
		ret = func(env *Env) bool {
			val := fun(env).Convert(t)
			return val.Bool()
		}
	case r.Int:
		ret = func(env *Env) int {
			val := fun(env).Convert(t)
			return int(val.Int())
		}
	case r.Int8:
		ret = func(env *Env) int8 {
			val := fun(env).Convert(t)
			return int8(val.Int())
		}
	case r.Int16:
		ret = func(env *Env) int16 {
			val := fun(env).Convert(t)
			return int16(val.Int())
		}
	case r.Int32:
		ret = func(env *Env) int32 {
			val := fun(env).Convert(t)
			return int32(val.Int())
		}
	case r.Int64:
		ret = func(env *Env) int64 {
			val := fun(env).Convert(t)
			return val.Int()
		}
	case r.Uint:
		ret = func(env *Env) uint {
			val := fun(env).Convert(t)
			return uint(val.Uint())
		}
	case r.Uint16:
		ret = func(env *Env) uint16 {
			val := fun(env).Convert(t)
			return uint16(val.Uint())
		}
	case r.Uint32:
		ret = func(env *Env) uint32 {
			val := fun(env).Convert(t)
			return uint32(val.Uint())
		}
	case r.Uint64:
		ret = func(env *Env) uint64 {
			val := fun(env).Convert(t)
			return val.Uint()
		}
	case r.Uintptr:
		ret = func(env *Env) uintptr {
			val := fun(env).Convert(t)
			return uintptr(val.Uint())
		}
	case r.Float32:
		ret = func(env *Env) float32 {
			val := fun(env).Convert(t)
			return float32(val.Float())
		}
	case r.Float64:
		ret = func(env *Env) float64 {
			val := fun(env).Convert(t)
			return val.Float()
		}
	case r.Complex64:
		ret = func(env *Env) complex64 {
			val := fun(env).Convert(t)
			return complex64(val.Complex())
		}
	case r.Complex128:
		ret = func(env *Env) complex128 {
			val := fun(env).Convert(t)
			return val.Complex()
		}
	case r.String:
		ret = func(env *Env) string {
			val := fun(env).Convert(t)
			return val.String()
		}
	default:
		ret = func(env *Env) r.Value {
			return fun(env).Convert(t)
		}
	}
	return exprFun(t, ret)
}
