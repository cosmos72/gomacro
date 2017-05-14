/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * builtin.go
 *
 *  Created on: Apr 02, 2017
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"io"
	r "reflect"
	"time"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

var (
	zeroTypes          = []xr.Type{}
	typeOfReflectType  = xr.TypeOf((*r.Type)(nil)).Elem() // inception
	rtypeOfSliceOfByte = r.TypeOf([]byte{})

/*
	typeOfPtrBool       = xr.PtrTo(xr.TypeOfBool)
	typeOfPtrInt        = xr.PtrTo(xr.TypeOfInt)
	typeOfPtrInt8       = xr.PtrTo(xr.TypeOfInt)
	typeOfPtrInt16      = xr.PtrTo(xr.TypeOfInt16)
	typeOfPtrInt32      = xr.PtrTo(xr.TypeOfInt32)
	typeOfPtrInt64      = xr.PtrTo(xr.TypeOfInt64)
	typeOfPtrUint       = xr.PtrTo(xr.TypeOfUint)
	typeOfPtrUint8      = xr.PtrTo(xr.TypeOfUint)
	typeOfPtrUint16     = xr.PtrTo(xr.TypeOfUint16)
	typeOfPtrUint32     = xr.PtrTo(xr.TypeOfUint32)
	typeOfPtrUint64     = xr.PtrTo(xr.TypeOfUint64)
	typeOfPtrUintptr    = xr.PtrTo(xr.TypeOfUintptr)
	typeOfPtrFloat32    = xr.PtrTo(xr.TypeOfFloat32)
	typeOfPtrFloat64    = xr.PtrTo(xr.TypeOfFloat64)
	typeOfPtrComplex64  = xr.PtrTo(xr.TypeOfComplex64)
	typeOfPtrComplex128 = xr.PtrTo(xr.TypeOfComplex128)
	typeOfPtrString     = xr.PtrTo(xr.TypeOfString)
*/
)

// =================================== iota ===================================

func (top *Comp) addIota() {
	// https://golang.org/ref/spec#Constants
	// "Literal constants, true, false, iota, and certain constant expressions containing only untyped constant operands are untyped."
	top.Binds["iota"] = BindUntyped(UntypedZero)
}

func (top *Comp) removeIota() {
	delete(top.Binds, "iota")
}

func (top *Comp) incrementIota() {
	uIota := top.Binds["iota"].Lit.Value.(UntypedLit).Obj
	uIota = constant.BinaryOp(uIota, token.ADD, UntypedOne.Obj)
	top.Binds["iota"] = BindUntyped(UntypedLit{Kind: r.Int, Obj: uIota})
}

// ============================== initialization ===============================

func (ce *CompEnv) addBuiltins() {
	// https://golang.org/ref/spec#Constants
	// "Literal constants, true, false, iota, and certain constant expressions containing only untyped constant operands are untyped."
	ce.DeclConst("false", nil, UntypedLit{Kind: r.Bool, Obj: constant.MakeBool(false)})
	ce.DeclConst("true", nil, UntypedLit{Kind: r.Bool, Obj: constant.MakeBool(true)})

	// https://golang.org/ref/spec#Variables : "[...] the predeclared identifier nil, which has no type"
	ce.DeclConst("nil", nil, nil)

	ce.DeclBuiltin("append", Builtin{compileAppend, 1, base.MaxUint16})
	ce.DeclBuiltin("cap", Builtin{compileCap, 1, 1})
	ce.DeclBuiltin("close", Builtin{compileClose, 1, 1})
	ce.DeclBuiltin("copy", Builtin{compileCopy, 2, 2})
	ce.DeclBuiltin("complex", Builtin{compileComplex, 2, 2})
	ce.DeclBuiltin("delete", Builtin{compileDelete, 2, 2})
	ce.DeclBuiltin("imag", Builtin{compileRealImag, 1, 1})
	ce.DeclBuiltin("len", Builtin{compileLen, 1, 1})
	ce.DeclBuiltin("make", Builtin{compileMake, 1, 3})
	ce.DeclBuiltin("new", Builtin{compileNew, 1, 1})
	ce.DeclBuiltin("panic", Builtin{compilePanic, 1, 1})
	ce.DeclBuiltin("print", Builtin{compilePrint, 0, base.MaxUint16})
	ce.DeclBuiltin("println", Builtin{compilePrint, 0, base.MaxUint16})
	ce.DeclBuiltin("real", Builtin{compileRealImag, 1, 1})

	ce.DeclEnvFunc("Env", Function{callIdentity, xr.TypeOf((*func(*CompEnv) *CompEnv)(nil)).Elem()})
	ce.DeclFunc("Sleep", func(seconds float64) {
		time.Sleep(time.Duration(seconds * float64(time.Second)))
	})
	/*
		binds["Eval"] = r.ValueOf(Function{funcEval, 1})
		binds["MacroExpand"] = r.ValueOf(Function{funcMacroExpand, -1})
		binds["MacroExpand1"] = r.ValueOf(Function{funcMacroExpand1, -1})
		binds["MacroExpandCodewalk"] = r.ValueOf(Function{funcMacroExpandCodewalk, -1})
		binds["Parse"] = r.ValueOf(Function{funcParse, 1})
		binds["Read"] = r.ValueOf(ReadString)
		binds["ReadDir"] = r.ValueOf(callReadDir)
		binds["ReadFile"] = r.ValueOf(callReadFile)
		binds["ReadMultiline"] = r.ValueOf(ReadMultiline)
		binds["Slice"] = r.ValueOf(callSlice)
		binds["String"] = r.ValueOf(func(args ...interface{}) string {
			return env.toString("", args...)
		})
		// return multiple values, extracting the concrete type of each interface
		binds["Values"] = r.ValueOf(Function{funcValues, -1})
	*/
	/*
		binds["recover"] = r.ValueOf(Function{funcRecover, 0})
	*/

	// --------- types ---------
	ce.DeclType("bool", xr.TypeOfBool)
	ce.DeclType("byte", xr.TypeOfByte)
	ce.DeclType("complex64", xr.TypeOfComplex64)
	ce.DeclType("complex128", xr.TypeOfComplex128)
	ce.DeclType("error", xr.TypeOfError)
	ce.DeclType("float32", xr.TypeOfFloat32)
	ce.DeclType("float64", xr.TypeOfFloat64)
	ce.DeclType("int", xr.TypeOfInt)
	ce.DeclType("int8", xr.TypeOfInt8)
	ce.DeclType("int16", xr.TypeOfInt16)
	ce.DeclType("int32", xr.TypeOfInt32)
	ce.DeclType("int64", xr.TypeOfInt64)
	ce.DeclType("rune", xr.TypeOfRune)
	ce.DeclType("string", xr.TypeOfString)
	ce.DeclType("uint", xr.TypeOfUint)
	ce.DeclType("uint8", xr.TypeOfUint8)
	ce.DeclType("uint16", xr.TypeOfUint16)
	ce.DeclType("uint32", xr.TypeOfUint32)
	ce.DeclType("uint64", xr.TypeOfUint64)
	ce.DeclType("uintptr", xr.TypeOfUintptr)

	ce.DeclType("Duration", xr.TypeOf(time.Duration(0)))

	/*
		// --------- proxies ---------
		if env.Proxies == nil {
			env.Proxies = make(map[string]Type)
		}
		proxies := env.Proxies

		proxies["error", TypeOf(*Error_builtin)(nil)).Elem()
	*/
}

// ============================= builtin functions =============================

// --- append() ---

func compileAppend(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	n := len(node.Args)
	args := make([]*Expr, n)

	args[0] = c.Expr1(node.Args[0])
	t0 := args[0].Type
	if t0.Kind() != r.Slice {
		c.Errorf("first argument to %s must be slice; have <%s>", sym.Name, t0)
		return nil
	}
	telem := t0.Elem()

	if node.Ellipsis != token.NoPos {
		if n != 2 {
			return c.badBuiltinCallArgNum(sym.Name+"(arg1, arg2...)", 2, 2, node.Args)
		}
		telem = t0 // second argument is a slice too
	}
	for i := 1; i < n; i++ {
		argi := c.Expr1(node.Args[i])
		if argi.Const() {
			argi.ConstTo(telem)
		} else if ti := argi.Type; ti != telem && (ti == nil || !ti.AssignableTo(telem)) {
			return c.badBuiltinCallArgType(sym.Name, node.Args[i], ti, telem)
		}
		args[i] = argi
	}
	t := xr.FuncOf([]xr.Type{t0, t0}, []xr.Type{t0}, true) // compile as reflect.Append(), which is variadic
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: r.Append}, &sym)
	return &Call{
		Fun:      fun,
		Args:     args,
		OutTypes: []xr.Type{t0},
		Const:    false,
		Ellipsis: node.Ellipsis != token.NoPos,
	}
}

// --- cap() ---

func callCap(val r.Value) int {
	return val.Cap()
}

func compileCap(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	// argument of builtin cap() cannot be a literal
	arg := c.Expr1(node.Args[0])
	tin := arg.Type
	tout := xr.TypeOfInt
	switch tin.Kind() {
	// no cap() on r.Map, see
	// https://golang.org/ref/spec#Length_and_capacity
	// and https://golang.org/pkg/reflect/#Value.Cap
	case r.Array, r.Chan, r.Slice:
		// ok
	case r.Ptr:
		if tin.Elem().Kind() == r.Array {
			// cap() on pointer to array
			arg = c.Deref(arg)
			tin = arg.Type
			break
		}
		fallthrough
	default:
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "array, channel, slice, pointer to array")
	}
	t := xr.FuncOf([]xr.Type{tin}, []xr.Type{tout}, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: callCap}, &sym)
	// capacity of arrays is part of their type: cannot change at runtime, we could optimize it.
	// TODO https://golang.org/ref/spec#Length_and_capacity specifies
	// when the array passed to cap() is evaluated and when is not...
	return newCall1(fun, arg, arg.Const(), tout)
}

// --- close() ---

func callClose(val r.Value) {
	val.Close()
}

func compileClose(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	arg := c.Expr1(node.Args[0])
	tin := arg.Type
	if tin.Kind() != r.Chan {
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "channel")
	}
	t := xr.FuncOf([]xr.Type{tin}, zeroTypes, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: callClose}, &sym)
	return newCall1(fun, arg, false)
}

// --- complex() ---

func callComplex64(re float32, im float32) complex64 {
	return complex(re, im)
}

func callComplex128(re float64, im float64) complex128 {
	return complex(re, im)
}

func compileComplex(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	re := c.Expr1(node.Args[0])
	im := c.Expr1(node.Args[1])
	if re.Untyped() {
		if im.Untyped() {
			re.ConstTo(xr.TypeOfFloat64)
			im.ConstTo(xr.TypeOfFloat64)
		} else {
			re.ConstTo(im.Type)
		}
	} else if im.Untyped() {
		im.ConstTo(re.Type)
	}
	c.toSameFuncType(node, re, im)
	kre := base.KindToCategory(re.Type.Kind())
	if re.Const() && kre != r.Float64 {
		re.ConstTo(xr.TypeOfFloat64)
		kre = r.Float64
	}
	kim := base.KindToCategory(im.Type.Kind())
	if im.Const() && kim != r.Float64 {
		im.ConstTo(xr.TypeOfFloat64)
		kim = r.Float64
	}
	if kre != r.Float64 {
		c.Errorf("invalid operation: %v (arguments have type %v, expected floating-point)",
			node, re.Type)
	}
	if kim != r.Float64 {
		c.Errorf("invalid operation: %v (arguments have type %v, expected floating-point)",
			node, im.Type)
	}
	tin := re.Type
	k := re.Type.Kind()
	var tout xr.Type
	var call I
	switch k {
	case r.Float32:
		tout = xr.TypeOfComplex64
		call = callComplex64
	case r.Float64:
		tout = xr.TypeOfComplex128
		call = callComplex128
	default:
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "floating point")
	}
	touts := []xr.Type{tout}
	t := xr.FuncOf([]xr.Type{tin}, touts, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: call}, &sym)
	// complex() of two constants is constant: it can be computed at compile time
	return &Call{Fun: fun, Args: []*Expr{re, im}, Const: re.Const() && im.Const(), OutTypes: touts}
}

// --- copy() ---

func copyStringToBytes(dst []byte, src string) int {
	// reflect.Copy does not support this case... use the compiler support
	return copy(dst, src)
}

func compileCopy(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	args := []*Expr{
		c.Expr1(node.Args[0]),
		c.Expr1(node.Args[1]),
	}
	if args[1].Const() {
		// we also accept a string literal as second argument
		args[1].ConstTo(args[1].DefaultType())
	}
	t0, t1 := args[0].Type, args[1].Type
	var funCopy I = r.Copy
	if t0 == nil || t0.Kind() != r.Slice || !t0.AssignableTo(xr.SliceOf(t0.Elem())) {
		// https://golang.org/ref/spec#Appending_and_copying_slices
		// copy [...] arguments must have identical element type T and must be assignable to a slice of type []T.
		c.Errorf("first argument to copy should be slice; have %v <%v>", node.Args[0], t0)
		return nil
	} else if t0.Elem().Kind() == r.Uint8 && t1.Kind() == r.String {
		// [...] As a special case, copy also accepts a destination argument assignable to type []byte
		// with a source argument of a string type. This form copies the bytes from the string into the byte slice.
		funCopy = copyStringToBytes
	} else if t1 == nil || t1.Kind() != r.Slice || !t1.AssignableTo(xr.SliceOf(t1.Elem())) {
		c.Errorf("second argument to copy should be slice or string; have %v <%v>", node.Args[1], t1)
		return nil
	} else if t0.Elem() != t1.Elem() {
		c.Errorf("arguments to copy have different element types: <%v> and <%v>", t0.Elem(), t1.Elem())
	}
	outtypes := []xr.Type{xr.TypeOfInt}
	t := xr.FuncOf([]xr.Type{t0, t1}, outtypes, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: funCopy}, &sym)
	return &Call{Fun: fun, Args: args, OutTypes: outtypes, Const: false}
}

// --- delete() ---

// use whatever calling convention is convenient: reflect.Values, interface{}s, primitive types...
// as long as call_builtin supports it, we're fine
func callDelete(vmap r.Value, vkey r.Value) {
	vmap.SetMapIndex(vkey, base.Nil)
}

func compileDelete(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	emap := c.Expr1(node.Args[0])
	ekey := c.Expr1(node.Args[1])
	tmap := emap.Type
	if tmap.Kind() != r.Map {
		c.Errorf("first argument to delete must be map; have %v", tmap)
		return nil
	}
	tkey := tmap.Key()
	if ekey.Const() {
		ekey.ConstTo(tkey)
	} else if ekey.Type == nil || !ekey.Type.AssignableTo(tkey) {
		c.Errorf("cannot use %v <%v> as type <%v> in delete", node.Args[1], ekey.Type, tkey)
	}
	t := xr.FuncOf([]xr.Type{tmap, tkey}, zeroTypes, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: callDelete}, &sym)
	return &Call{Fun: fun, Args: []*Expr{emap, ekey}, OutTypes: zeroTypes, Const: false}
}

// --- Env() ---

func callIdentity(v r.Value) r.Value {
	return v
}

// --- len() ---

func callLenValue(val r.Value) int {
	return val.Len()
}

func callLenString(val string) int {
	return len(val)
}

func compileLen(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	arg := c.Expr1(node.Args[0])
	if arg.Const() {
		arg.ConstTo(arg.DefaultType())
	}
	tin := arg.Type
	tout := xr.TypeOfInt
	switch tin.Kind() {
	case r.Array, r.Chan, r.Map, r.Slice, r.String:
		// ok
	case r.Ptr:
		if tin.Elem().Kind() == r.Array {
			// len() on pointer to array
			arg = c.Deref(arg)
			tin = arg.Type
			break
		}
		fallthrough
	default:
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "array, channel, map, slice, string, pointer to array")
	}
	t := xr.FuncOf([]xr.Type{tin}, []xr.Type{tout}, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: callLenValue}, &sym)
	if tin.Kind() == r.String {
		fun.Value = callLenString // optimization
	}
	// length of arrays is part of their type: cannot change at runtime, we could optimize it.
	// TODO https://golang.org/ref/spec#Length_and_capacity specifies
	// when the array passed to len() is evaluated and when is not...
	return newCall1(fun, arg, arg.Const(), tout)
}

// --- make() ---

func makeChan1(t r.Type) r.Value {
	return r.MakeChan(t, 0)
}

func makeMap2(t r.Type, n int) r.Value {
	// reflect.MakeMap cannot specify initial capacity
	return r.MakeMap(t)
}

func makeSlice2(t r.Type, n int) r.Value {
	// reflect.MakeSlice requires capacity
	return r.MakeSlice(t, n, n)
}

func compileMake(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	nargs := len(node.Args)
	var nmin, nmax uint16 = 1, 2
	tin := c.Type(node.Args[0])
	var funMakes [4]I
	switch tin.Kind() {
	case r.Chan:
		funMakes[1] = makeChan1
		funMakes[2] = r.MakeChan
	case r.Map:
		funMakes[1] = r.MakeMap
		funMakes[2] = makeMap2
	case r.Slice:
		nmin, nmax = 2, 3
		funMakes[2] = makeSlice2
		funMakes[3] = r.MakeSlice
	default:
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "channel, map, slice")
	}
	if nargs < int(nmin) || nargs > int(nmax) {
		return c.badBuiltinCallArgNum(sym.Name+"()", nmin, nmax, node.Args)
	}
	args := make([]*Expr, nargs)
	argtypes := make([]xr.Type, nargs)
	args[0] = c.exprValue(typeOfReflectType, tin.ReflectType())
	argtypes[0] = typeOfReflectType
	te := xr.TypeOfInt
	for i := 1; i < nargs; i++ {
		argi := c.Expr1(node.Args[i])
		if argi.Const() {
			argi.ConstTo(te)
		} else if ti := argi.Type; ti == nil || (ti != te && !ti.AssignableTo(te)) {
			return c.badBuiltinCallArgType(sym.Name, node.Args[i], ti, te)
		}
		args[i] = argi
		argtypes[i] = te
	}
	outtypes := []xr.Type{tin}
	t := xr.FuncOf(argtypes, outtypes, false)
	sym.Type = t
	funMake := funMakes[nargs]
	if funMake == nil {
		c.Errorf("internal error: no make() alternative to call for %v with %d arguments", tin, nargs)
		return nil
	}
	fun := exprLit(Lit{Type: t, Value: funMake}, &sym)
	return &Call{Fun: fun, Args: args, OutTypes: outtypes, Const: false}
}

// --- new() ---

func compileNew(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	tin := c.Type(node.Args[0])
	tout := xr.PtrTo(tin)
	t := xr.FuncOf([]xr.Type{typeOfReflectType}, []xr.Type{tout}, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: r.New}, &sym)
	arg := c.exprValue(typeOfReflectType, tin.ReflectType())
	return newCall1(fun, arg, false, tout)
}

// --- panic() ---

func callPanic(arg interface{}) {
	panic(arg)
}

var typeOfBuiltinPanic = xr.TypeOf(callPrint)

func compilePanic(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	arg := c.Expr1(node.Args[0])
	if arg.Const() {
		arg.ConstTo(arg.DefaultType())
	}

	t := typeOfBuiltinPanic
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: callPanic}, &sym)
	return newCall1(fun, arg, false)
}

// --- print(), println() ---

func callPrint(out io.Writer, args ...interface{}) {
	fmt.Fprint(out, args...)
}

func callPrintln(out io.Writer, args ...interface{}) {
	fmt.Fprintln(out, args...)
}

func getStdout(env *Env) r.Value {
	return r.ValueOf(env.ThreadGlobals.Stdout)
}

var (
	typeOfIoWriter     = xr.TypeOf((*io.Writer)(nil)).Elem()
	typeOfBuiltinPrint = xr.TypeOf(callPrint)
)

func compilePrint(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	args := c.Exprs(node.Args)
	for _, arg := range args {
		if arg.Const() {
			arg.ConstTo(arg.DefaultType())
		}
	}
	arg0 := exprFun(typeOfIoWriter, getStdout)
	args = append([]*Expr{arg0}, args...)

	t := typeOfBuiltinPrint
	sym.Type = t
	call := callPrint
	if sym.Name == "println" {
		call = callPrintln
	}
	fun := exprLit(Lit{Type: t, Value: call}, &sym)
	return &Call{Fun: fun, Args: args, OutTypes: zeroTypes, Const: false, Ellipsis: node.Ellipsis != token.NoPos}
}

// --- real() and imag() ---

func callReal32(val complex64) float32 {
	return real(val)
}

func callReal64(val complex128) float64 {
	return real(val)
}

func callImag32(val complex64) float32 {
	return imag(val)
}

func callImag64(val complex128) float64 {
	return imag(val)
}

func compileRealImag(c *Comp, sym Symbol, node *ast.CallExpr) *Call {
	arg := c.Expr1(node.Args[0])
	if arg.Const() {
		arg.ConstTo(arg.DefaultType())
	}
	tin := arg.Type
	var tout xr.Type
	var call I
	switch tin.Kind() {
	case r.Complex64:
		tout = xr.TypeOfFloat32
		if sym.Name == "real" {
			call = callReal32
		} else {
			call = callImag32
		}
	case r.Complex128:
		tout = xr.TypeOfFloat64
		if sym.Name == "real" {
			call = callReal64
		} else {
			call = callImag64
		}
	default:
		return c.badBuiltinCallArgType(sym.Name, node.Args[0], tin, "complex")
	}
	t := xr.FuncOf([]xr.Type{tin}, []xr.Type{tout}, false)
	sym.Type = t
	fun := exprLit(Lit{Type: t, Value: call}, &sym)
	// real() and imag() of a constant are constants: they can be computed at compile time
	return newCall1(fun, arg, arg.Const(), tout)
}

// ============================ support functions =============================

// call_builtin compiles a call to a builtin function: append, cap, copy, delete, len, make, new...
func call_builtin(c *Call) I {
	// builtin functions are always literals, i.e. funindex == NoIndex thus not stored in Env.Binds[]
	// we must retrieve them directly from c.Fun.Value
	if !c.Fun.Const() {
		base.Errorf("internal error: call_builtin() invoked for non-constant function %#v. use one of the callXretY() instead", c.Fun)
	}
	var name string
	if c.Fun.Sym != nil {
		name = c.Fun.Sym.Name
	}
	args := c.Args
	argfuns := make([]I, len(args))
	for i, arg := range args {
		argfuns[i] = arg.WithFun()
	}
	if false {
		argtypes := make([]xr.Type, len(args))
		for i, arg := range args {
			argtypes[i] = arg.Type
		}
		// Debugf("compiling builtin %s() <%v> with arg types %v", name, TypeOf(c.Fun.Value), argtypes)
	}
	var call I
	switch fun := c.Fun.Value.(type) {
	case func(float32, float32) complex64: // complex
		arg0fun := argfuns[0].(func(*Env) float32)
		arg1fun := argfuns[1].(func(*Env) float32)
		if name == "complex" {
			if args[0].Const() {
				arg0 := args[0].Value.(float32)
				call = func(env *Env) complex64 {
					arg1 := arg1fun(env)
					return complex(arg0, arg1)
				}
			} else if args[1].Const() {
				arg1 := args[1].Value.(float32)
				call = func(env *Env) complex64 {
					arg0 := arg0fun(env)
					return complex(arg0, arg1)
				}
			} else {
				call = func(env *Env) complex64 {
					arg0 := arg0fun(env)
					arg1 := arg1fun(env)
					return complex(arg0, arg1)
				}
			}
		} else {
			call = func(env *Env) complex64 {
				arg0 := arg0fun(env)
				arg1 := arg1fun(env)
				return fun(arg0, arg1)
			}
		}
	case func(float64, float64) complex128: // complex
		arg0fun := argfuns[0].(func(*Env) float64)
		arg1fun := argfuns[1].(func(*Env) float64)
		if name == "complex" {
			if args[0].Const() {
				arg0 := args[0].Value.(float64)
				call = func(env *Env) complex128 {
					arg1 := arg1fun(env)
					return complex(arg0, arg1)
				}
			} else if args[1].Const() {
				arg1 := args[1].Value.(float64)
				call = func(env *Env) complex128 {
					arg0 := arg0fun(env)
					return complex(arg0, arg1)
				}
			} else {
				call = func(env *Env) complex128 {
					arg0 := arg0fun(env)
					arg1 := arg1fun(env)
					return complex(arg0, arg1)
				}
			}
		} else {
			call = func(env *Env) complex128 {
				arg0 := arg0fun(env)
				arg1 := arg1fun(env)
				return fun(arg0, arg1)
			}
		}
	case func(complex64) float32: // real(), imag()
		argfun := argfuns[0].(func(*Env) complex64)
		if name == "real" {
			call = func(env *Env) float32 {
				arg := argfun(env)
				return real(arg)
			}
		} else if name == "imag" {
			call = func(env *Env) float32 {
				arg := argfun(env)
				return imag(arg)
			}
		} else {
			call = func(env *Env) float32 {
				arg := argfun(env)
				return fun(arg)
			}
		}
	case func(complex128) float64: // real(), imag()
		argfun := argfuns[0].(func(*Env) complex128)
		if name == "real" {
			call = func(env *Env) float64 {
				arg := argfun(env)
				return real(arg)
			}
		} else if name == "imag" {
			call = func(env *Env) float64 {
				arg := argfun(env)
				return imag(arg)
			}
		} else {
			call = func(env *Env) float64 {
				arg := argfun(env)
				return fun(arg)
			}
		}
	case func(string) int: // len(string)
		argfun := argfuns[0].(func(*Env) string)
		if name == "len" {
			call = func(env *Env) int {
				arg := argfun(env)
				return len(arg)
			}
		} else {
			call = func(env *Env) int {
				arg := argfun(env)
				return fun(arg)
			}
		}
	case func([]byte, string) int: // copy([]byte, string)
		arg0fun := args[0].AsX1()
		if args[1].Const() {
			// string is a literal
			arg1const := args[1].Value.(string)
			call = func(env *Env) int {
				// arg0 is "assignable to []byte"
				arg0 := arg0fun(env)
				if arg0.Type() != rtypeOfSliceOfByte {
					arg0 = arg0.Convert(rtypeOfSliceOfByte)
				}
				return fun(arg0.Interface().([]byte), arg1const)
			}
		} else {
			arg1fun := args[1].Fun.(func(*Env) string)
			call = func(env *Env) int {
				// arg0 is "assignable to []byte"
				arg0 := arg0fun(env)
				if arg0.Type() != rtypeOfSliceOfByte {
					arg0 = arg0.Convert(rtypeOfSliceOfByte)
				}
				arg1 := arg1fun(env)
				return fun(arg0.Interface().([]byte), arg1)
			}
		}
	case func(interface{}): // panic()
		argfunsX1 := c.MakeArgfunsX1()
		argfun := argfunsX1[0]
		if name == "panic" {
			call = func(env *Env) {
				arg := argfun(env).Interface()
				panic(arg)
			}
		} else {
			call = func(env *Env) {
				arg := argfun(env).Interface()
				fun(arg)
			}
		}
	case func(r.Value): // close()
		argfunsX1 := c.MakeArgfunsX1()
		argfun := argfunsX1[0]
		if name == "close" {
			call = func(env *Env) {
				arg := argfun(env)
				arg.Close()
			}
		} else {
			call = func(env *Env) {
				arg := argfun(env)
				fun(arg)
			}
		}
	case func(r.Value) int: // cap(), len()
		argfunsX1 := c.MakeArgfunsX1()
		argfun := argfunsX1[0]
		call = func(env *Env) int {
			arg := argfun(env)
			return fun(arg)
		}
	case func(r.Value) r.Value: // Env()
		argfunsX1 := c.MakeArgfunsX1()
		argfun := argfunsX1[0]
		if name == "Env" {
			call = func(env *Env) r.Value {
				arg0 := argfun(env)
				return arg0
			}
		} else {
			call = func(env *Env) r.Value {
				arg0 := argfun(env)
				return fun(arg0)
			}
		}
	case func(r.Value, r.Value): // delete()
		argfunsX1 := c.MakeArgfunsX1()
		call = func(env *Env) {
			arg0 := argfunsX1[0](env)
			arg1 := argfunsX1[1](env)
			fun(arg0, arg1)
		}
	case func(r.Value, r.Value) int: // copy()
		argfunsX1 := c.MakeArgfunsX1()
		call = func(env *Env) int {
			arg0 := argfunsX1[0](env)
			arg1 := argfunsX1[1](env)
			return fun(arg0, arg1)
		}
	case func(io.Writer, ...interface{}): // print, println()
		argfunsX1 := c.MakeArgfunsX1()
		if c.Ellipsis {
			argfunsX1 := [2]func(*Env) r.Value{
				argfunsX1[0],
				argfunsX1[1],
			}
			call = func(env *Env) {
				arg0 := argfunsX1[0](env).Interface()
				argslice := argfunsX1[1](env).Interface().([]interface{})
				fun(arg0.(io.Writer), argslice...)
			}
		} else {
			call = func(env *Env) {
				args := make([]interface{}, len(argfunsX1))
				for i, argfun := range argfunsX1 {
					args[i] = argfun(env).Interface()
				}
				fun(args[0].(io.Writer), args[1:]...)
			}
		}
	case func(r.Value, ...r.Value) r.Value: // append()
		argfunsX1 := c.MakeArgfunsX1()
		if c.Ellipsis {
			argfunsX1 := [2]func(*Env) r.Value{
				argfunsX1[0],
				argfunsX1[1],
			}
			if name == "append" {
				call = func(env *Env) r.Value {
					arg0 := argfunsX1[0](env)
					arg1 := argfunsX1[1](env)
					argslice := unwrapSlice(arg1)
					return r.Append(arg0, argslice...)
				}
			} else {
				call = func(env *Env) r.Value {
					arg0 := argfunsX1[0](env)
					arg1 := argfunsX1[1](env)
					argslice := unwrapSlice(arg1)
					return fun(arg0, argslice...)
				}
			}
		} else {
			if name == "append" {
				call = func(env *Env) r.Value {
					args := make([]r.Value, len(argfunsX1))
					for i, argfun := range argfunsX1 {
						args[i] = argfun(env)
					}
					return r.Append(args[0], args[1:]...)
				}
			} else {
				call = func(env *Env) r.Value {
					args := make([]r.Value, len(argfunsX1))
					for i, argfun := range argfunsX1 {
						args[i] = argfun(env)
					}
					return fun(args[0], args[1:]...)
				}
			}
		}
	case func(r.Type) r.Value: // new(), make()
		arg0 := args[0].Value.(r.Type)
		if name == "new" {
			call = func(env *Env) r.Value {
				return r.New(arg0)
			}
		} else {
			call = func(env *Env) r.Value {
				return fun(arg0)
			}
		}
	case func(r.Type, int) r.Value: // make()
		arg0 := args[0].Value.(r.Type)
		arg1fun := argfuns[1].(func(*Env) int)
		call = func(env *Env) r.Value {
			arg1 := arg1fun(env)
			return fun(arg0, arg1)
		}
	case func(r.Type, int, int) r.Value: // make()
		arg0 := args[0].Value.(r.Type)
		arg1fun := argfuns[1].(func(*Env) int)
		arg2fun := argfuns[2].(func(*Env) int)
		call = func(env *Env) r.Value {
			arg1 := arg1fun(env)
			arg2 := arg2fun(env)
			return fun(arg0, arg1, arg2)
		}
	default:
		base.Errorf("unimplemented call_builtin() for function type %v", r.TypeOf(fun))
	}
	return call
}

// unwrapSlice accepts a reflect.Value with kind == reflect.Array, Slice or String
// and returns slice of its elements, each wrapped in a reflect.Value
func unwrapSlice(arg r.Value) []r.Value {
	n := arg.Len()
	slice := make([]r.Value, n)
	for i := range slice {
		slice[i] = arg.Index(i)
	}
	return slice
}

// callBuiltin invokes the appropriate compiler for a call to a builtin function: cap, copy, len, make, new...
func (c *Comp) callBuiltin(node *ast.CallExpr, fun *Expr) *Call {
	builtin := fun.Value.(Builtin)
	if fun.Sym == nil {
		c.Errorf("invalid call to non-name builtin: %v", node)
		return nil
	}
	nmin := builtin.ArgMin
	nmax := builtin.ArgMax
	n := len(node.Args)
	if n < int(nmin) || n > int(nmax) {
		return c.badBuiltinCallArgNum(fun.Sym.Name+"()", nmin, nmax, node.Args)
	}
	return builtin.Compile(c, *fun.Sym, node)
}

// callFunction compiles a call to a function that accesses interpreter's *CompEnv
func (c *Comp) callFunction(node *ast.CallExpr, fun *Expr) (newfun *Expr, lastarg *Expr) {
	function := fun.Value.(Function)
	t := function.Type
	var sym *Symbol
	if fun.Sym != nil {
		symcopy := *fun.Sym
		symcopy.Type = t
		sym = &symcopy
	}
	newfun = exprLit(Lit{Type: t, Value: function.Fun}, sym)
	if len(node.Args) < t.NumIn() {
		lastarg = exprX1(typeOfCompEnv, func(env *Env) r.Value {
			return r.ValueOf(&CompEnv{Comp: c, env: env})
		})
	}
	return newfun, lastarg
}

func (c *Comp) badBuiltinCallArgNum(name interface{}, nmin uint16, nmax uint16, args []ast.Expr) *Call {
	prefix := "not enough"
	nargs := len(args)
	if nargs > int(nmax) {
		prefix = "too many"
	}
	str := fmt.Sprintf("%d", nmin)
	if nmax <= nmin {
	} else if nmax == nmin+1 {
		str = fmt.Sprintf("%s or %d", str, nmax)
	} else if nmax < base.MaxUint16 {
		str = fmt.Sprintf("%s to %d", str, nmax)
	} else {
		str = fmt.Sprintf("%s or more", str)
	}
	c.Errorf("%s arguments in call to builtin %v: expecting %s, found %d: %v", prefix, name, str, nargs, args)
	return nil
}

func (c *Comp) badBuiltinCallArgType(name string, arg ast.Expr, tactual xr.Type, texpected interface{}) *Call {
	c.Errorf("cannot use %v <%v> as %v in builtin %s()", arg, tactual, texpected, name)
	return nil
}
