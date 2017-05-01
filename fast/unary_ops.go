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
 * unary_plus.go
 *
 *  Created on Apr 07, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) UnaryPlus(node *ast.UnaryExpr, xe *Expr) *Expr {
	if !IsCategory(xe.Type.Kind(), r.Int, r.Uint, r.Float64, r.Complex128) {
		return c.invalidUnaryExpr(node, xe)
	}
	return xe
}

func (c *Comp) UnaryMinus(node *ast.UnaryExpr, xe *Expr) *Expr {
	// if xe is constant, UnaryExpr will invoke EvalConst()
	// on our return value. no need to optimize that.
	x := xe.Fun
	var fun I
	switch x := x.(type) {
	case func(env *Env) int:
		fun = func(env *Env) int {
			return -x(env)
		}
	case func(env *Env) int8:
		fun = func(env *Env) int8 {
			return -x(env)
		}
	case func(env *Env) int16:
		fun = func(env *Env) int16 {
			return -x(env)
		}
	case func(env *Env) int32:
		fun = func(env *Env) int32 {
			return -x(env)
		}
	case func(env *Env) int64:
		fun = func(env *Env) int64 {
			return -x(env)
		}
	case func(env *Env) uint:
		fun = func(env *Env) uint {
			return -x(env)
		}
	case func(env *Env) uint8:
		fun = func(env *Env) uint8 {
			return -x(env)
		}
	case func(env *Env) uint16:
		fun = func(env *Env) uint16 {
			return -x(env)
		}
	case func(env *Env) uint32:
		fun = func(env *Env) uint32 {
			return -x(env)
		}
	case func(env *Env) uint64:
		fun = func(env *Env) uint64 {
			return -x(env)
		}
	case func(env *Env) uintptr:
		fun = func(env *Env) uintptr {
			return -x(env)
		}
	case func(env *Env) float32:
		fun = func(env *Env) float32 {
			return -x(env)
		}
	case func(env *Env) float64:
		fun = func(env *Env) float64 {
			return -x(env)
		}
	case func(env *Env) complex64:
		fun = func(env *Env) complex64 {
			return -x(env)
		}
	case func(env *Env) complex128:
		fun = func(env *Env) complex128 {
			return -x(env)
		}
	default:
		return c.invalidUnaryExpr(node, xe)
	}
	return exprFun(xe.Type, fun)
}

func (c *Comp) UnaryXor(node *ast.UnaryExpr, xe *Expr) *Expr {
	// if xe is constant, UnaryExpr will invoke EvalConst()
	// on our return value. no need to optimize that.
	x := xe.Fun
	var fun I
	switch x := x.(type) {
	case func(env *Env) int:
		fun = func(env *Env) int {
			return ^x(env)
		}
	case func(env *Env) int8:
		fun = func(env *Env) int8 {
			return ^x(env)
		}
	case func(env *Env) int16:
		fun = func(env *Env) int16 {
			return ^x(env)
		}
	case func(env *Env) int32:
		fun = func(env *Env) int32 {
			return ^x(env)
		}
	case func(env *Env) int64:
		fun = func(env *Env) int64 {
			return ^x(env)
		}
	case func(env *Env) uint:
		fun = func(env *Env) uint {
			return ^x(env)
		}
	case func(env *Env) uint8:
		fun = func(env *Env) uint8 {
			return ^x(env)
		}
	case func(env *Env) uint16:
		fun = func(env *Env) uint16 {
			return ^x(env)
		}
	case func(env *Env) uint32:
		fun = func(env *Env) uint32 {
			return ^x(env)
		}
	case func(env *Env) uint64:
		fun = func(env *Env) uint64 {
			return ^x(env)
		}
	case func(env *Env) uintptr:
		fun = func(env *Env) uintptr {
			return ^x(env)
		}
	default:
		return c.invalidUnaryExpr(node, xe)
	}
	return exprFun(xe.Type, fun)
}

func (c *Comp) UnaryNot(node *ast.UnaryExpr, xe *Expr) *Expr {
	// if xe is constant, UnaryExpr will invoke EvalConst()
	// on our return value. no need to optimize that.
	x := xe.Fun
	var fun I
	switch x := x.(type) {
	case func(env *Env) bool:
		fun = func(env *Env) bool {
			return !x(env)
		}
	default:
		return c.invalidUnaryExpr(node, xe)
	}
	return exprFun(xe.Type, fun)
}

// UnaryRecv compiles <-channel (returns two values: the received value and an 'ok' flag)
func (c *Comp) UnaryRecv(node *ast.UnaryExpr, xe *Expr) *Expr {
	t := xe.Type
	if t.Kind() != r.Chan {
		return c.badUnaryExpr("expecting channel, found", node, xe)
	}
	if t.ChanDir()&r.RecvDir == 0 {
		return c.badUnaryExpr("cannot receive from send-only channel", node, xe)
	}
	var fun func(env *Env) (r.Value, []r.Value)
	switch x := xe.Fun.(type) {
	case func(env *Env) (r.Value, []r.Value):
		channelfun := x
		fun = func(env *Env) (r.Value, []r.Value) {
			channel, _ := channelfun(env)
			retv, ok := channel.Recv()
			var okv r.Value
			if ok {
				okv = True
			} else {
				okv = False
			}
			return retv, []r.Value{retv, okv}
		}
	default:
		channelfun := xe.AsX1()
		fun = func(env *Env) (r.Value, []r.Value) {
			retv, ok := channelfun(env).Recv()
			var okv r.Value
			if ok {
				okv = True
			} else {
				okv = False
			}
			return retv, []r.Value{retv, okv}
		}
	}
	types := []r.Type{t.Elem(), TypeOfBool}
	return exprXV(types, fun)
}

// UnaryRecv1 compiles <-channel (returns a single value: the received value)
// mandatory optimization: fast_interpreter ASSUMES that expressions
// returning bool, int, uint, float, complex, string do NOT wrap them in reflect.Value
func (c *Comp) UnaryRecv1(node *ast.UnaryExpr, xe *Expr) *Expr {
	t := xe.Type
	if t.Kind() != r.Chan {
		return c.badUnaryExpr("expecting channel, found", node, xe)
	}
	if t.ChanDir()&r.RecvDir == 0 {
		return c.badUnaryExpr("cannot receive from send-only channel", node, xe)
	}
	telem := t.Elem()
	var fun I
	switch x := xe.Fun.(type) {
	case func(env *Env) (r.Value, []r.Value):
		channelfun := x
		switch telem.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return int(retv.Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return int8(retv.Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return int16(retv.Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return int32(retv.Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return uint(retv.Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return uint8(retv.Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return uint16(retv.Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return uint32(retv.Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return uintptr(retv.Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return float32(retv.Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return complex64(retv.Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv.String()
			}
		default:
			fun = func(env *Env) r.Value {
				channel, _ := channelfun(env)
				retv, _ := channel.Recv()
				return retv
			}
		}
	default:
		recvonly := t.ChanDir() == r.RecvDir
		channelfun := xe.AsX1()
		switch telem.Kind() {
		case r.Bool:
			if telem != TypeOfBool {
				fun = func(env *Env) bool {
					retv, _ := channelfun(env).Recv()
					return retv.Bool()
				}
			} else if recvonly {
				fun = func(env *Env) bool {
					channel := channelfun(env).Interface().(<-chan bool)
					return <-channel
				}
			} else {
				fun = func(env *Env) bool {
					channel := channelfun(env).Interface().(chan bool)
					return <-channel
				}
			}
		case r.Int:
			if telem != TypeOfInt {
				fun = func(env *Env) int {
					retv, _ := channelfun(env).Recv()
					return int(retv.Int())
				}
			} else if recvonly {
				fun = func(env *Env) int {
					channel := channelfun(env).Interface().(<-chan int)
					return <-channel
				}
			} else {
				fun = func(env *Env) int {
					channel := channelfun(env).Interface().(chan int)
					return <-channel
				}
			}
		case r.Int8:
			if telem != TypeOfInt8 {
				fun = func(env *Env) int8 {
					retv, _ := channelfun(env).Recv()
					return int8(retv.Int())
				}
			} else if recvonly {
				fun = func(env *Env) int8 {
					channel := channelfun(env).Interface().(<-chan int8)
					return <-channel
				}
			} else {
				fun = func(env *Env) int8 {
					channel := channelfun(env).Interface().(chan int8)
					return <-channel
				}
			}
		case r.Int16:
			if telem != TypeOfInt16 {
				fun = func(env *Env) int16 {
					retv, _ := channelfun(env).Recv()
					return int16(retv.Int())
				}
			} else if recvonly {
				fun = func(env *Env) int16 {
					channel := channelfun(env).Interface().(<-chan int16)
					return <-channel
				}
			} else {
				fun = func(env *Env) int16 {
					channel := channelfun(env).Interface().(chan int16)
					return <-channel
				}
			}
		case r.Int32:
			if telem != TypeOfInt32 {
				fun = func(env *Env) int32 {
					retv, _ := channelfun(env).Recv()
					return int32(retv.Int())
				}
			} else if recvonly {
				fun = func(env *Env) int32 {
					channel := channelfun(env).Interface().(<-chan int32)
					return <-channel
				}
			} else {
				fun = func(env *Env) int32 {
					channel := channelfun(env).Interface().(chan int32)
					return <-channel
				}
			}
		case r.Int64:
			if telem != TypeOfInt64 {
				fun = func(env *Env) int64 {
					retv, _ := channelfun(env).Recv()
					return retv.Int()
				}
			} else if recvonly {
				fun = func(env *Env) int64 {
					channel := channelfun(env).Interface().(<-chan int64)
					return <-channel
				}
			} else {
				fun = func(env *Env) int64 {
					channel := channelfun(env).Interface().(chan int64)
					return <-channel
				}
			}
		case r.Uint:
			if telem != TypeOfUint {
				fun = func(env *Env) uint {
					retv, _ := channelfun(env).Recv()
					return uint(retv.Uint())
				}
			} else if recvonly {
				fun = func(env *Env) uint {
					channel := channelfun(env).Interface().(<-chan uint)
					return <-channel
				}
			} else {
				fun = func(env *Env) uint {
					channel := channelfun(env).Interface().(chan uint)
					return <-channel
				}
			}
		case r.Uint8:
			if telem != TypeOfUint8 {
				fun = func(env *Env) uint8 {
					retv, _ := channelfun(env).Recv()
					return uint8(retv.Uint())
				}
			} else if recvonly {
				fun = func(env *Env) uint8 {
					channel := channelfun(env).Interface().(<-chan uint8)
					return <-channel
				}
			} else {
				fun = func(env *Env) uint8 {
					channel := channelfun(env).Interface().(chan uint8)
					return <-channel
				}
			}
		case r.Uint16:
			if telem != TypeOfUint16 {
				fun = func(env *Env) uint16 {
					retv, _ := channelfun(env).Recv()
					return uint16(retv.Uint())
				}
			} else if recvonly {
				fun = func(env *Env) uint16 {
					channel := channelfun(env).Interface().(<-chan uint16)
					return <-channel
				}
			} else {
				fun = func(env *Env) uint16 {
					channel := channelfun(env).Interface().(chan uint16)
					return <-channel
				}
			}
		case r.Uint32:
			if telem != TypeOfUint32 {
				fun = func(env *Env) uint32 {
					retv, _ := channelfun(env).Recv()
					return uint32(retv.Uint())
				}
			} else if recvonly {
				fun = func(env *Env) uint32 {
					channel := channelfun(env).Interface().(<-chan uint32)
					return <-channel
				}
			} else {
				fun = func(env *Env) uint32 {
					channel := channelfun(env).Interface().(chan uint32)
					return <-channel
				}
			}
		case r.Uint64:
			if telem != TypeOfUint64 {
				fun = func(env *Env) uint64 {
					retv, _ := channelfun(env).Recv()
					return retv.Uint()
				}
			} else if recvonly {
				fun = func(env *Env) uint64 {
					channel := channelfun(env).Interface().(<-chan uint64)
					return <-channel
				}
			} else {
				fun = func(env *Env) uint64 {
					channel := channelfun(env).Interface().(chan uint64)
					return <-channel
				}
			}
		case r.Uintptr:
			if telem != TypeOfUintptr {
				fun = func(env *Env) uintptr {
					retv, _ := channelfun(env).Recv()
					return uintptr(retv.Uint())
				}
			} else if recvonly {
				fun = func(env *Env) uintptr {
					channel := channelfun(env).Interface().(<-chan uintptr)
					return <-channel
				}
			} else {
				fun = func(env *Env) uintptr {
					channel := channelfun(env).Interface().(chan uintptr)
					return <-channel
				}
			}
		case r.Float32:
			if telem != TypeOfFloat32 {
				fun = func(env *Env) float32 {
					retv, _ := channelfun(env).Recv()
					return float32(retv.Float())
				}
			} else if recvonly {
				fun = func(env *Env) float32 {
					channel := channelfun(env).Interface().(<-chan float32)
					return <-channel
				}
			} else {
				fun = func(env *Env) float32 {
					channel := channelfun(env).Interface().(chan float32)
					return <-channel
				}
			}
		case r.Float64:
			if telem != TypeOfFloat32 {
				fun = func(env *Env) float64 {
					retv, _ := channelfun(env).Recv()
					return retv.Float()
				}
			} else if recvonly {
				fun = func(env *Env) float64 {
					channel := channelfun(env).Interface().(<-chan float64)
					return <-channel
				}
			} else {
				fun = func(env *Env) float64 {
					channel := channelfun(env).Interface().(chan float64)
					return <-channel
				}
			}
		case r.Complex64:
			if telem != TypeOfComplex64 {
				fun = func(env *Env) complex64 {
					retv, _ := channelfun(env).Recv()
					return complex64(retv.Complex())
				}
			} else if recvonly {
				fun = func(env *Env) complex64 {
					channel := channelfun(env).Interface().(<-chan complex64)
					return <-channel
				}
			} else {
				fun = func(env *Env) complex64 {
					channel := channelfun(env).Interface().(chan complex64)
					return <-channel
				}
			}
		case r.Complex128:
			if telem != TypeOfComplex128 {
				fun = func(env *Env) complex128 {
					retv, _ := channelfun(env).Recv()
					return retv.Complex()
				}
			} else if recvonly {
				fun = func(env *Env) complex128 {
					channel := channelfun(env).Interface().(<-chan complex128)
					return <-channel
				}
			} else {
				fun = func(env *Env) complex128 {
					channel := channelfun(env).Interface().(chan complex128)
					return <-channel
				}
			}
		case r.String:
			if telem != TypeOfString {
				fun = func(env *Env) string {
					retv, _ := channelfun(env).Recv()
					return retv.String()
				}
			} else if recvonly {
				fun = func(env *Env) string {
					channel := channelfun(env).Interface().(<-chan string)
					return <-channel
				}
			} else {
				fun = func(env *Env) string {
					channel := channelfun(env).Interface().(chan string)
					return <-channel
				}
			}
		default:
			fun = func(env *Env) r.Value {
				retv, _ := channelfun(env).Recv()
				return retv
			}
		}
	}
	return exprFun(telem, fun)
}

// StarExpr compiles unary operator * i.e. pointer dereference
func (c *Comp) StarExpr(node *ast.StarExpr) *Expr {
	addr := c.Expr1(node.X) // panics if addr returns zero values, warns if returns multiple values
	taddr := addr.Type
	if taddr.Kind() != r.Ptr {
		c.Errorf("unary operation * on non-pointer <%v>: %v", taddr, node)
	}
	return c.Deref(addr)
}

// Deref compiles unary operator * i.e. pointer dereference
func (c *Comp) Deref(addr *Expr) *Expr {
	taddr := addr.Type
	if taddr.Kind() != r.Ptr {
		c.Errorf("unary operation * on non-pointer <%v>", taddr)
	}
	x1 := addr.AsX1() // panics if addr returns zero values, warns if returns multiple values
	t := taddr.Elem()
	x := addr.Fun
	var fun I
	// fast interpreter expects that Exprs returning primitive types or string
	// do NOT wrap them into reflect.Value
	switch x := x.(type) {
	case func(env *Env) *bool:
		fun = func(env *Env) bool {
			return *x(env)
		}
	case func(env *Env) *int:
		fun = func(env *Env) int {
			return *x(env)
		}
	case func(env *Env) *int8:
		fun = func(env *Env) int8 {
			return *x(env)
		}
	case func(env *Env) *int16:
		fun = func(env *Env) int16 {
			return *x(env)
		}
	case func(env *Env) *int32:
		fun = func(env *Env) int32 {
			return *x(env)
		}
	case func(env *Env) *int64:
		fun = func(env *Env) int64 {
			return *x(env)
		}
	case func(env *Env) *uint:
		fun = func(env *Env) uint {
			return *x(env)
		}
	case func(env *Env) *uint8:
		fun = func(env *Env) uint8 {
			return *x(env)
		}
	case func(env *Env) *uint16:
		fun = func(env *Env) uint16 {
			return *x(env)
		}
	case func(env *Env) *uint32:
		fun = func(env *Env) uint32 {
			return *x(env)
		}
	case func(env *Env) *uint64:
		fun = func(env *Env) uint64 {
			return *x(env)
		}
	case func(env *Env) *uintptr:
		fun = func(env *Env) uintptr {
			return *x(env)
		}
	case func(env *Env) *float32:
		fun = func(env *Env) float32 {
			return *x(env)
		}
	case func(env *Env) *float64:
		fun = func(env *Env) float64 {
			return *x(env)
		}
	case func(env *Env) *complex64:
		fun = func(env *Env) complex64 {
			return *x(env)
		}
	default:
		fun = c.derefUnwrap(t, x1)
	}
	return exprFun(t, fun)
}

// deref0Unwrap compiles unary operator * on reflect.Value - unwraps reflect.Value.Elem() if possible
func (c *Comp) derefUnwrap(t r.Type, x1 func(*Env) r.Value) I {
	var fun I
	switch t.Kind() {
	case r.Bool:
		fun = func(env *Env) bool {
			return x1(env).Elem().Bool()
		}
	case r.Int:
		fun = func(env *Env) int {
			return int(x1(env).Elem().Int())
		}
	case r.Int8:
		fun = func(env *Env) int8 {
			return int8(x1(env).Elem().Int())
		}
	case r.Int16:
		fun = func(env *Env) int16 {
			return int16(x1(env).Elem().Int())
		}
	case r.Int32:
		fun = func(env *Env) int32 {
			return int32(x1(env).Elem().Int())
		}
	case r.Int64:
		fun = func(env *Env) int64 {
			return x1(env).Elem().Int()
		}
	case r.Uint:
		fun = func(env *Env) uint {
			return uint(x1(env).Elem().Uint())
		}
	case r.Uint8:
		fun = func(env *Env) uint8 {
			return uint8(x1(env).Elem().Uint())
		}
	case r.Uint16:
		fun = func(env *Env) uint16 {
			return uint16(x1(env).Elem().Uint())
		}
	case r.Uint32:
		fun = func(env *Env) uint32 {
			return uint32(x1(env).Elem().Uint())
		}
	case r.Uint64:
		fun = func(env *Env) uint64 {
			return x1(env).Elem().Uint()
		}
	case r.Uintptr:
		fun = func(env *Env) uintptr {
			return uintptr(x1(env).Elem().Uint())
		}
	case r.Float32:
		fun = func(env *Env) float32 {
			return float32(x1(env).Elem().Float())
		}
	case r.Float64:
		fun = func(env *Env) float64 {
			return x1(env).Elem().Float()
		}
	case r.Complex64:
		fun = func(env *Env) complex64 {
			return complex64(x1(env).Elem().Complex())
		}
	case r.Complex128:
		fun = func(env *Env) complex128 {
			return x1(env).Elem().Complex()
		}
	case r.String:
		fun = func(env *Env) string {
			return x1(env).Elem().String()
		}
	default:
		fun = func(env *Env) r.Value {
			return x1(env).Elem()
		}
	}
	return fun
}
