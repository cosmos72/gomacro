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

package compiler

func (c *CompEnv) addBuiltins() {
	c.DefConst("false", nil, false)
	c.DefConst("true", nil, true)
	c.DefConst("nil", nil, nil)

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
	c.DeclType0("bool", typeOfBool)
	c.DeclType0("byte", typeOfByte)
	c.DeclType0("complex64", typeOfComplex64)
	c.DeclType0("complex128", typeOfComplex128)
	c.DeclType0("error", typeOfError)
	c.DeclType0("float32", typeOfFloat32)
	c.DeclType0("float64", typeOfFloat64)
	c.DeclType0("int", typeOfInt)
	c.DeclType0("int8", typeOfInt8)
	c.DeclType0("int16", typeOfInt16)
	c.DeclType0("int32", typeOfInt32)
	c.DeclType0("int64", typeOfInt64)
	c.DeclType0("rune", typeOfRune)
	c.DeclType0("string", typeOfString)
	c.DeclType0("uint", typeOfUint)
	c.DeclType0("uint8", typeOfUint8)
	c.DeclType0("uint16", typeOfUint16)
	c.DeclType0("uint32", typeOfUint32)
	c.DeclType0("uint64", typeOfUint64)
	c.DeclType0("uintptr", typeOfUintptr)

	/*
		// --------- proxies ---------
		if env.Proxies == nil {
			env.Proxies = make(map[string]r.Type)
		}
		proxies := env.Proxies

		proxies["error", typeOf(*Error_builtin)(nil)).Elem()
	*/
}
