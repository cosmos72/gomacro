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

package fast_interpreter

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
	if i == 0 && e.Types == nil {
		return e.Type
	}
	return e.Types[i]
}

// Outs returns the types that an expression will produce when evaluated
func (e *Expr) Outs() []r.Type {
	if e.Types == nil {
		return []r.Type{e.Type}
	}
	return e.Types
}

type BindClass int

// the zero value of BindDescriptor is a valid descriptor for variables or functions named "_", and also for all constants
type BindDescriptor BindClass

const (
	ConstBind BindClass = iota
	FuncBind
	VarBind
	IntBind
	NoBind           BindClass      = ConstBind // NoBind and ConstBind are aliases
	bindClassMask    BindClass      = 0x3
	bindIndexShift                  = 2
	NoBindDescriptor BindDescriptor = 0
)

func MakeBindDescriptor(class BindClass, index int) BindDescriptor {
	class &= bindClassMask
	// debugf("MakeBindDescriptor() class=%v index=%v", class, index)
	if class != NoBind {
		index++
	}
	return BindDescriptor(index<<bindIndexShift | int(class))
}

// IntBind returns true if BindIndex refers to a slot in Env.IntBinds (the default is a slot in Env.Binds)
func (desc BindDescriptor) Class() BindClass {
	return BindClass(desc) & bindClassMask
}

func (desc BindDescriptor) Index() int {
	index := int(desc>>bindIndexShift) - 1
	// debugf("BindDescriptor=%v, class=%v, index=%v", desc, desc.Class(), index)
	return index
}

func (desc BindDescriptor) Settable() bool {
	class := desc.Class()
	return class == IntBind || class == VarBind
}

// Bind represents a constant, variable, or function in the compiler
type Bind struct {
	Lit
	Desc BindDescriptor
}

func (bind *Bind) Const() bool {
	return bind.Desc == NoBindDescriptor
}

func BindValue(value I) Bind {
	return Bind{Lit: Lit{Type: r.TypeOf(value), Value: value}, Desc: NoBindDescriptor}
}

type NamedType struct {
	Name, Path string
}

type Comp struct {
	Binds      map[string]Bind
	BindNum    int // len(Binds) == BindNum + IntBindNum + # of constants
	IntBindNum int
	Types      map[string]r.Type
	NamedTypes map[r.Type]NamedType
	Outer      *Comp
	Name       string
	Path       string
}

type Env struct {
	Binds    []r.Value
	IntBinds []uint64
	Outer    *Env
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
	I interface{}
	X func(*Env)
	/*
		XBool func(*Env) bool
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
		XV          func(*Env) (r.Value, []r.Value)
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

func (c *Comp) Errorf(format string, args ...interface{}) X {
	panic(runtimeError{c, format, args})
}

func (c *Comp) Error(err error) interface{} {
	panic(err)
}

func errorf(format string, args ...interface{}) X {
	panic(runtimeError{nil, format, args})
}

func error_(err error) interface{} {
	panic(err)
}

func (c *Comp) Warnf(format string, args ...interface{}) {
	format = fmt.Sprintf("// warning: %s\n", format)
	fmt.Fprintf(os.Stdout, format, args...)
}

func warnf(format string, args ...interface{}) {
	format = fmt.Sprintf("// warning: %s\n", format)
	fmt.Fprintf(os.Stdout, format, args...)
}

func (c *Comp) Debugf(format string, args ...interface{}) {
	format = fmt.Sprintf("// debug: %s\n", format)
	fmt.Fprintf(os.Stdout, format, args...)
}

func debugf(format string, args ...interface{}) {
	format = fmt.Sprintf("// debug: %s\n", format)
	fmt.Fprintf(os.Stdout, format, args...)
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
