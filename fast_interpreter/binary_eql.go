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
 * binary_eql.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) Eql(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	if xe.IsNil {
		if ye.IsNil {
			return c.invalidBinaryExpr(node, xe, ye)
		} else {
			// nil == expr
			return c.eqlNil(node, xe, ye)
		}
	} else if ye.IsNil {
		// expr == nil
		return c.eqlNil(node, xe, ye)
	}
	if !xe.Type.Comparable() || !xe.Type.Comparable() {
		return c.invalidBinaryExpr(node, xe, ye)
	}
	xc, yc := xe.Const(), ye.Const()
	if xe.Type.Kind() != r.Interface && ye.Type.Kind() != r.Interface {
		// comparison between different types is allowed only if at least one is an interface
		c.toSameFuncType(node, xe, ye)
	}

	// if both x and y are constants, BinaryExpr will invoke EvalConst()
	// on our return value. no need to optimize that.
	var fun func(env *Env) bool
	if xc == yc {
		x, y := xe.Fun, ye.Fun
		switch x := x.(type) {
		case func(*Env) int:
			y := y.(func(*Env) int)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) int8:
			y := y.(func(*Env) int8)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) int16:
			y := y.(func(*Env) int16)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) int32:
			y := y.(func(*Env) int32)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) int64:
			y := y.(func(*Env) int64)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) uint:
			y := y.(func(*Env) uint)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) uint8:
			y := y.(func(*Env) uint8)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) uint16:
			y := y.(func(*Env) uint16)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) uint32:
			y := y.(func(*Env) uint32)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) uint64:
			y := y.(func(*Env) uint64)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) uintptr:
			y := y.(func(*Env) uintptr)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) float32:
			y := y.(func(*Env) float32)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) float64:
			y := y.(func(*Env) float64)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) complex64:
			y := y.(func(*Env) complex64)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) complex128:
			y := y.(func(*Env) complex128)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		case func(*Env) string:
			y := y.(func(*Env) string)
			fun = func(env *Env) bool {
				return x(env) == y(env)
			}
		default:
			return c.eqlMisc(node, xe, ye)
		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		switch x := x.(type) {
		case func(*Env) int:
			y := y.(int)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) int8:
			y := y.(int8)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) int16:
			y := y.(int16)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) int32:
			y := y.(int32)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) int64:
			y := y.(int64)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) uint:
			y := y.(uint)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) uint8:
			y := y.(uint8)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) uint16:
			y := y.(uint16)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) uint32:
			y := y.(uint32)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) uint64:
			y := y.(uint64)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) uintptr:
			y := y.(uintptr)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) float32:
			y := y.(float32)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) float64:
			y := y.(float64)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) complex64:
			y := y.(complex64)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) complex128:
			y := y.(complex128)
			fun = func(env *Env) bool {
				return x(env) == y
			}
		case func(*Env) string:
			y := y.(string)
			if len(y) == 0 {
				fun = func(env *Env) bool {
					return len(x(env)) == 0
				}
			} else {
				fun = func(env *Env) bool {
					return x(env) == y
				}
			}
		default:
			return c.eqlMisc(node, xe, ye)
		}
	} else {
		x := xe.Value
		y := ye.Fun
		switch y := y.(type) {
		case func(*Env) int:
			x := x.(int)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) int8:
			x := x.(int8)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) int16:
			x := x.(int16)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) int32:
			x := x.(int32)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) int64:
			x := x.(int64)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) uint:
			x := x.(uint)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) uint8:
			x := x.(uint8)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) uint16:
			x := x.(uint16)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) uint32:
			x := x.(uint32)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) uint64:
			x := x.(uint64)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) uintptr:
			x := x.(uintptr)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) float32:
			x := x.(float32)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) float64:
			x := x.(float64)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) complex64:
			x := x.(complex64)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) complex128:
			x := x.(complex128)
			fun = func(env *Env) bool {
				return x == y(env)
			}
		case func(*Env) string:
			x := x.(string)
			if len(x) == 0 {
				fun = func(env *Env) bool {
					return len(y(env)) == 0
				}
			} else {
				fun = func(env *Env) bool {
					return x == y(env)
				}
			}
		default:
			return c.eqlMisc(node, xe, ye)
		}
	}
	return ExprBool(fun)
}

func (c *Comp) eqlMisc(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	var fun func(*Env) bool

	if xe.Type.Kind() == r.Interface || ye.Type.Kind() == r.Interface {
		// not checked yet that xe and ye return at least one value... check now
		xe.CheckX1()
		ye.CheckX1()
	}

	switch x := xe.Fun.(type) {
	case func(*Env) (r.Value, []r.Value):
		switch y := ye.Fun.(type) {
		case func(*Env) (r.Value, []r.Value):
			fun = func(env *Env) bool {
				v1, _ := x(env)
				v2, _ := y(env)
				if v1 == Nil || v2 == Nil {
					return v1 == v2
				} else {
					return v1.Interface() == v2.Interface()
				}
			}
		default:
			y1 := ye.AsX1()
			fun = func(env *Env) bool {
				v1, _ := x(env)
				v2 := y1(env)
				if v1 == Nil || v2 == Nil {
					return v1 == v2
				} else {
					return v1.Interface() == v2.Interface()
				}
			}
		}
	default:
		x1 := xe.AsX1()

		switch y := ye.Fun.(type) {
		case func(*Env) (r.Value, []r.Value):
			fun = func(env *Env) bool {
				v1 := x1(env)
				v2, _ := y(env)
				if v1 == Nil || v2 == Nil {
					return v1 == v2
				} else {
					return v1.Interface() == v2.Interface()
				}
			}
		default:
			y1 := ye.AsX1()
			fun = func(env *Env) bool {
				v1 := x1(env)
				v2 := y1(env)
				if v1 == Nil || v2 == Nil {
					return v1 == v2
				} else {
					return v1.Interface() == v2.Interface()
				}
			}
		}
	}
	return ExprBool(fun)
}

func (c *Comp) eqlNil(node *ast.BinaryExpr, xe *Expr, ye *Expr) *Expr {
	var e *Expr
	if ye.IsNil {
		e = xe
	} else {
		e = ye
	}
	// e cannot be a constant (none of the nillable types support compile-time constants) but better safe than sorry
	if e.Const() || !IsNillableKind(e.Type.Kind()) {
		return c.invalidBinaryExpr(node, xe, ye)
	}

	var fun func(env *Env) bool
	if f, ok := e.Fun.(func(env *Env) (r.Value, []r.Value)); ok {
		e.CheckX1() // to warn or error as appropriate
		fun = func(env *Env) bool {
			v, _ := f(env)
			vnil := v == Nil || IsNillableKind(v.Kind()) && v.IsNil()
			return vnil
		}
	} else {
		f := e.AsX1()
		fun = func(env *Env) bool {
			v := f(env)
			vnil := v == Nil || IsNillableKind(v.Kind()) && v.IsNil()
			return vnil
		}
	}
	return ExprBool(fun)
}
