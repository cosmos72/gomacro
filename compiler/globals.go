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
 * globals.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package compiler

import (
	"fmt"
	"os"
	r "reflect"

	"github.com/cosmos72/gomacro/constants"
)

// Lit represents a literal value, i.e. a constant
type Lit struct {
	Type  r.Type
	Value I // may be nil when embedded in other structs that represent non-constants
}

// Expr represents an expression in the compiler
type Expr struct {
	Lit
	Types []r.Type // in case the expression produces multiple values. if nil, use Lit.Type.
	Fun   I        // function that evaluates the expression at runtime.
	isNil bool
}

func (e *Expr) Const() bool {
	return e.Value != nil || e.isNil
}

// NumOut returns the number of values that an expression will produce when evaluated
func (e *Expr) NumOut() int {
	if e.Types == nil {
		return 1
	}
	return len(e.Types)
}

// Out returns the i-th type that an expression will produce when evaluated
func (e *Expr) Out(i int) r.Type {
	if e.Types == nil {
		return e.Type
	}
	return e.Types[i]
}

// Bind represents a constant, variable, or function in the compiler
type Bind struct {
	Lit
	Index int // < 0 for literals i.e. constants
}

func (b *Bind) Const() bool {
	return b.Index < 0
}

func BindValue(value I) Bind {
	return Bind{Lit: Lit{Type: r.TypeOf(value), Value: value}, Index: -1}
}

type NamedType struct {
	Name, Path string
}

type Comp struct {
	Binds      map[string]Bind
	BindNum    int // len(Binds) - # of constants
	Types      map[string]r.Type
	NamedTypes map[r.Type]NamedType
	Outer      *Comp
	Name       string
	Path       string
}

type Env struct {
	Binds []r.Value
	Outer *Env
}

type CompEnv struct {
	*Comp
	Env *Env
}

type SReturn struct {
	result0 r.Value
	results []r.Value
}

type (
	I     interface{}
	X     func(*Env) (r.Value, []r.Value)
	XBool func(*Env) bool
	/*
		X1          func(*Env) r.Value
		XInt        func(*Env) int
		XInt8       func(*Env) int8
		XInt16      func(*Env) int16
		XInt32      func(*Env) int32
		XInt64      func(*Env) int64
		XUint       func(*Env) uint
		XUint8      func(*Env) uint8
		XUint16     func(*Env) uint16
		XUint32     func(*Env) uint32
		XUint64     func(*Env) uint64
		XUintptr    func(*Env) uintptr
		XFloat32    func(*Env) float32
		XFloat64    func(*Env) float64
		XComplex64  func(*Env) complex64
		XComplex128 func(*Env) complex128
		XString     func(*Env) string
	*/
)

type Func func(args ...r.Value) (r.Value, []r.Value)
type FuncInt func(args ...r.Value) int

type XFunc func(env *Env) Func
type XFuncInt func(env *Env) FuncInt

type runtimeError struct {
	comp   *Comp
	format string
	args   []interface{}
}

func RetOf(expr I) r.Type {
	t := TypeOf(expr)
	if t == nil || t.Kind() != r.Func {
		return t
	}
	if t.NumOut() == 1 {
		return t.Out(0)
	}
	return typeOfInterface
}

func (err runtimeError) Error() string {
	return fmt.Sprintf(err.format, err.args...)
}

func (comp *Comp) Errorf(format string, args ...interface{}) X {
	panic(runtimeError{comp, format, args})
}

func (comp *Comp) Error(err error) interface{} {
	panic(err)
}

func errorf(format string, args ...interface{}) X {
	panic(runtimeError{nil, format, args})
}

func error_(err error) interface{} {
	panic(err)
}

func (comp *Comp) Warnf(format string, args ...interface{}) {
	format = fmt.Sprintf("warning: %s\n", format)
	fmt.Fprintf(os.Stderr, format, args...)
}

var (
	False = constants.False
	True  = constants.True

	Nil  = constants.Nil
	None = constants.None

	nilEnv *Env
	NilEnv = []r.Value{r.ValueOf(nilEnv)}

	typeOfInt   = r.TypeOf(int(0))
	typeOfInt8  = r.TypeOf(int8(0))
	typeOfInt16 = r.TypeOf(int16(0))
	typeOfInt32 = r.TypeOf(int32(0))
	typeOfInt64 = r.TypeOf(int64(0))

	typeOfUint    = r.TypeOf(uint(0))
	typeOfUint8   = r.TypeOf(uint8(0))
	typeOfUint16  = r.TypeOf(uint16(0))
	typeOfUint32  = r.TypeOf(uint32(0))
	typeOfUint64  = r.TypeOf(uint64(0))
	typeOfUintptr = r.TypeOf(uintptr(0))

	typeOfFloat32    = r.TypeOf(float32(0))
	typeOfFloat64    = r.TypeOf(float64(0))
	typeOfComplex64  = r.TypeOf(complex64(0))
	typeOfComplex128 = r.TypeOf(complex128(0))

	typeOfBool      = r.TypeOf(bool(false))
	typeOfByte      = r.TypeOf(byte(0))
	typeOfRune      = r.TypeOf(rune(0))
	typeOfString    = r.TypeOf("")
	typeOfInterface = r.TypeOf((*interface{})(nil)).Elem()
	typeOfError     = r.TypeOf((*error)(nil)).Elem()

	zeroStrings = []string{}
	zeroTypes   = []r.Type{}
)
