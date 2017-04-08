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

package fast_interpreter

import (
	"go/constant"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

var (
	untypedZero = UntypedLit{Kind: r.Int, Obj: constant.MakeInt64(int64(0))}
	untypedOne  = UntypedLit{Kind: r.Int, Obj: constant.MakeInt64(int64(1))}
)

func (top *Comp) addIota() {
	// https://golang.org/ref/spec#Constants
	// "Literal constants, true, false, iota, and certain constant expressions containing only untyped constant operands are untyped."
	top.Binds["iota"] = BindConst(untypedZero)
}

func (top *Comp) removeIota() {
	delete(top.Binds, "iota")
}

func (top *Comp) incrementIota() {
	uIota := top.Binds["iota"].Lit.Value.(UntypedLit).Obj
	uIota = constant.BinaryOp(uIota, token.ADD, untypedOne.Obj)
	top.Binds["iota"] = BindConst(UntypedLit{Kind: r.Int, Obj: uIota})
}

func (c *CompEnv) addBuiltins() {
	// https://golang.org/ref/spec#Constants
	// "Literal constants, true, false, iota, and certain constant expressions containing only untyped constant operands are untyped."
	c.DeclConst0("false", nil, UntypedLit{Kind: r.Bool, Obj: constant.MakeBool(false)})
	c.DeclConst0("true", nil, UntypedLit{Kind: r.Bool, Obj: constant.MakeBool(true)})

	// https://golang.org/ref/spec#Variables : "[...] the predeclared identifier nil, which has no type"
	c.DeclConst0("nil", nil, nil)

	/*
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
	*/
	/*
		binds["imag"] = r.ValueOf(Function{funcImag, 1})
		binds["len"] = r.ValueOf(callLen)
		binds["make"] = r.ValueOf(Builtin{builtinMake, -1})
		binds["new"] = r.ValueOf(Builtin{builtinNew, 1})
		binds["panic"] = r.ValueOf(callPanic)
		binds["println"] = r.ValueOf(func(args ...interface{}) {
			// values := toValues(args)
			// env.FprintValues(env.Stdout, values...)
			fmt.Fprintln(env.Stdout, args...)
		})
		binds["real"] = r.ValueOf(Function{funcReal, 1})
		binds["recover"] = r.ValueOf(Function{funcRecover, 0})
	*/

	// --------- types ---------
	c.DefType("bool", TypeOfBool)
	c.DefType("byte", TypeOfByte)
	c.DefType("complex64", TypeOfComplex64)
	c.DefType("complex128", TypeOfComplex128)
	c.DefType("error", TypeOfError)
	c.DefType("float32", TypeOfFloat32)
	c.DefType("float64", TypeOfFloat64)
	c.DefType("int", TypeOfInt)
	c.DefType("int8", TypeOfInt8)
	c.DefType("int16", TypeOfInt16)
	c.DefType("int32", TypeOfInt32)
	c.DefType("int64", TypeOfInt64)
	c.DefType("rune", TypeOfRune)
	c.DefType("string", TypeOfString)
	c.DefType("uint", TypeOfUint)
	c.DefType("uint8", TypeOfUint8)
	c.DefType("uint16", TypeOfUint16)
	c.DefType("uint32", TypeOfUint32)
	c.DefType("uint64", TypeOfUint64)
	c.DefType("uintptr", TypeOfUintptr)

	/*
		// --------- proxies ---------
		if env.Proxies == nil {
			env.Proxies = make(map[string]r.Type)
		}
		proxies := env.Proxies

		proxies["error", TypeOf(*Error_builtin)(nil)).Elem()
	*/
}
