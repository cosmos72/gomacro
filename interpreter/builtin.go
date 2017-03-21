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
 *  Created on: Feb 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"fmt"
	"go/ast"
	"io/ioutil"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
)

func funcAppend(env *Env, args []r.Value) (r.Value, []r.Value) {
	n := len(args)
	if n < 1 {
		return env.errorf("builtin append() expects at least one argument, found %d", n)
	}
	return r.Append(args[0], args[1:]...), nil
}

func callCap(arg interface{}) int {
	return r.ValueOf(arg).Cap()
}

func callClose(channel interface{}) {
	r.ValueOf(channel).Close()
}

func funcComplex(env *Env, args []r.Value) (r.Value, []r.Value) {
	rv, iv := args[0], args[1]
	r_, rok := env.toFloat(rv)
	i_, iok := env.toFloat(iv)
	if !rok {
		return env.errorf("builtin complex(): not a float: %v <%v>", rv, typeOf(rv))
	}
	if !iok {
		return env.errorf("builtin complex(): not a float: %v <%v>", iv, typeOf(iv))
	}
	cplx := complex(r_, i_)
	var ret interface{}
	if rv.Kind() == r.Float32 && iv.Kind() == r.Float32 {
		ret = complex64(cplx)
	} else {
		ret = cplx
	}
	return r.ValueOf(ret), nil
}

func callCopy(dst, src interface{}) int {
	return r.Copy(r.ValueOf(dst), r.ValueOf(src))
}

func callDelete(m interface{}, key interface{}) {
	r.ValueOf(m).SetMapIndex(r.ValueOf(key), Nil)
}

func funcEnv(env *Env, args []r.Value) (r.Value, []r.Value) {
	return r.ValueOf(env), nil
}

func funcEval(env *Env, args []r.Value) (r.Value, []r.Value) {
	node := toInterface(args[0])
	form := AnyToAst(node, "Eval")
	return env.EvalAst(form)
}

func funcImag(env *Env, args []r.Value) (r.Value, []r.Value) {
	cv := args[0]
	c_, ok := env.toComplex(cv)
	if !ok {
		return env.errorf("builtin imag(): not a complex: %v <%v>", cv, typeOf(cv))
	}
	i_ := imag(c_)
	var ret interface{}
	if cv.Kind() == r.Complex64 {
		ret = float32(i_)
	} else {
		ret = i_
	}
	return r.ValueOf(ret), nil
}

func callLen(arg interface{}) int {
	return r.ValueOf(arg).Len()
}

//
// --------- macroexpansion ----------
//

func funcMacroExpand(env *Env, args []r.Value) (r.Value, []r.Value) {
	return callMacroExpand(env, args, cMacroExpand)
}

func funcMacroExpand1(env *Env, args []r.Value) (r.Value, []r.Value) {
	return callMacroExpand(env, args, cMacroExpand1)
}

func funcMacroExpandCodewalk(env *Env, args []r.Value) (r.Value, []r.Value) {
	return callMacroExpand(env, args, cMacroExpandCodewalk)
}

func callMacroExpand(env *Env, args []r.Value, which whichMacroExpand) (r.Value, []r.Value) {
	n := len(args)
	if n < 1 || n > 2 {
		return env.errorf("builtin %v() expects one or two arguments, found %d: %v", which, n, args)
	}
	val := args[0]
	if val == Nil || val == None {
		return val, nil
	}
	node := AnyToAstWithNode(val.Interface(), which.String()).Node()
	if n == 2 {
		e := args[1]
		if e != Nil && e != None {
			env = e.Interface().(*Env)
		}
	}
	var expanded bool
	switch which {
	case cMacroExpand1:
		node, expanded = env.MacroExpand1(node)
	case cMacroExpandCodewalk:
		node, expanded = env.MacroExpandCodewalk(node)
	default:
		node, expanded = env.MacroExpand(node)
	}
	nodev := r.ValueOf(node)
	return nodev, []r.Value{nodev, r.ValueOf(expanded)}
}

func builtinMake(env *Env, args []ast.Expr) (r.Value, []r.Value) {
	n := len(args)
	if n < 1 || n > 3 {
		return env.errorf("builtin make() expects one, two or three arguments, found %d", n)
	}
	t := env.evalType(args[0])
	values := env.evalExprs(args[1:])
	n--
	ret := Nil
	switch t.Kind() {
	case r.Chan:
		buffer := 0
		if n > 0 {
			buffer = int(values[0].Int())
		}
		ret = r.MakeChan(t, buffer)
	case r.Map:
		ret = r.MakeMap(t)
	case r.Slice:
		length := 0
		if n > 0 {
			length = int(values[0].Int())
		}
		capacity := length
		if n > 1 {
			capacity = int(values[1].Int())
		}
		ret = r.MakeSlice(t, length, capacity)
	}
	return ret, nil
}

func builtinNew(env *Env, args []ast.Expr) (r.Value, []r.Value) {
	t := env.evalType(args[0])
	return r.New(t), nil
}

func funcParse(env *Env, args []r.Value) (r.Value, []r.Value) {
	var in interface{}
	if arg := args[0]; arg != Nil && arg != None {
		in = arg.Interface()
	}
	out := env.ParseAst(in)
	if out != nil {
		return r.ValueOf(out.Interface()), nil
	}
	return Nil, nil
}

func callPanic(arg interface{}) {
	panic(arg)
}

func funcReal(env *Env, args []r.Value) (r.Value, []r.Value) {
	n := len(args)
	if n != 1 {
		return env.errorf("builtin real() expects exactly one argument, found %d", n)
	}
	cv := args[0]
	c_, ok := env.toComplex(cv)
	if !ok {
		return env.errorf("builtin real(): not a complex: %v <%v>", cv, typeOf(cv))
	}
	i_ := real(c_)
	var ret interface{}
	if cv.Kind() == r.Complex64 {
		ret = float32(i_)
	} else {
		ret = i_
	}
	return r.ValueOf(ret), nil
}

func callReadFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		callPanic(err)
	}
	return string(bytes)
}

func callReadDir(dirname string) []string {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		callPanic(err)
	}
	n := len(files)
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = files[i].Name()
	}
	return names
}

func funcRecover(env *Env, args []r.Value) (r.Value, []r.Value) {
	// Go specs: "Executing a call to recover inside a deferred function
	// (but not any function called by it) stops the panicking sequence
	// by restoring normal execution and retrieves the error value passed to the call of panic"
	//
	// thus recover() is invoked inside deferred functions: find their caller's env
	ret := Nil

	trace := env.Options&OptDebugPanicRecover != 0
	caller := env.CallerFrame()
	if trace {
		env.debugf("recover(): env = %v, stack is:", env.Name)
		env.showStack()
		curr := env.CurrentFrame()
		if curr != nil {
			env.debugf("           frame = %v, runningDefers = %v", curr.FuncEnv.Name, curr.runningDefers)
		} else {
			env.debugf("           frame = nil")
		}
		if caller != nil {
			env.debugf("           caller = %v, runningDefers = %v", caller.FuncEnv.Name, caller.runningDefers)
		} else {
			env.debugf("           caller = nil")
		}
	}

	if caller != nil {
		if caller.runningDefers && caller.panicking {
			// consume current panic
			if trace {
				env.debugf("           consuming current panic = %#v", caller.panick)
			}
			ret = r.ValueOf(caller.panick)
			caller.panick = nil
			caller.panicking = false
		} else if trace {
			env.debugf("           no panic to consume: caller.runningDefers = %q, caller.panicking = %q",
				caller.runningDefers, caller.panicking)
		}
	}
	return ret, nil
}

func callSlice(args ...interface{}) []interface{} {
	return args
}

func funcValues(env *Env, args []r.Value) (r.Value, []r.Value) {
	for i, arg := range args {
		// go through interface{} to forget any "static" compile-time type information
		if arg != None && arg != Nil {
			args[i] = r.ValueOf(arg.Interface())
		}
	}
	return unpackValues(args)
}

func (env *Env) addBuiltins() {
	if env.Binds == nil {
		env.Binds = make(map[string]r.Value)
	}
	binds := env.Binds

	binds["Env"] = r.ValueOf(Function{funcEnv, 0})
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

	binds["append"] = r.ValueOf(Function{funcAppend, -1})
	binds["cap"] = r.ValueOf(callCap)
	binds["close"] = r.ValueOf(callClose)
	binds["complex"] = r.ValueOf(Function{funcComplex, 2})
	binds["copy"] = r.ValueOf(callCopy)
	binds["delete"] = r.ValueOf(callDelete)
	binds["false"] = r.ValueOf(false)
	binds["imag"] = r.ValueOf(Function{funcImag, 1})
	binds["len"] = r.ValueOf(callLen)
	binds["make"] = r.ValueOf(Builtin{builtinMake, -1})
	binds["new"] = r.ValueOf(Builtin{builtinNew, 1})
	binds["nil"] = Nil
	binds["panic"] = r.ValueOf(callPanic)
	binds["println"] = r.ValueOf(func(args ...interface{}) {
		// values := toValues(args)
		// env.FprintValues(env.Stdout, values...)
		fmt.Fprintln(env.Stdout, args...)
	})
	binds["real"] = r.ValueOf(Function{funcReal, 1})
	binds["recover"] = r.ValueOf(Function{funcRecover, 0})
	binds["true"] = r.ValueOf(true)

	// --------- types ---------
	if env.Types == nil {
		env.Types = make(map[string]r.Type)
	}
	types := env.Types

	types["bool"] = r.TypeOf(false)
	types["byte"] = r.TypeOf(byte(0))
	types["complex64"] = r.TypeOf(complex64(0))
	types["complex128"] = r.TypeOf(complex128(0))
	types["error"] = r.TypeOf((*error)(nil)).Elem()
	types["float32"] = r.TypeOf(float32(0))
	types["float64"] = r.TypeOf(float64(0))
	types["int"] = r.TypeOf(int(0))
	types["int8"] = r.TypeOf(int8(0))
	types["int16"] = r.TypeOf(int16(0))
	types["int32"] = r.TypeOf(int32(0))
	types["int64"] = r.TypeOf(int64(0))
	types["rune"] = r.TypeOf(rune(0))
	types["string"] = r.TypeOf("")
	types["uint"] = r.TypeOf(uint(0))
	types["uint8"] = r.TypeOf(uint8(0))
	types["uint16"] = r.TypeOf(uint16(0))
	types["uint32"] = r.TypeOf(uint32(0))
	types["uint64"] = r.TypeOf(uint64(0))
	types["uintptr"] = r.TypeOf(uintptr(0))

	// --------- proxies ---------
	if env.Proxies == nil {
		env.Proxies = make(map[string]r.Type)
	}
	proxies := env.Proxies

	proxies["error"] = r.TypeOf((*Error_builtin)(nil)).Elem()
}

type Error_builtin struct {
	Object interface{}
	Error_ func() string
}

func (Proxy Error_builtin) Error() string {
	return Proxy.Error_()
}

func (env *Env) addInterpretedBuiltins() {
	if false {
		line := "func not(flag bool) bool { if flag { return false } else { return true } }"
		env.EvalAst(env.ParseAst(line))
	}
	if false {
		// Factorial(1000000): eval() elapsed time: 1.233714899 s
		line := "func Factorial(n int) int { t := 1; for i := 2; i <= n; i=i+1 { t = t * i }; t }"
		env.EvalAst(env.ParseAst(line))
	}
}
