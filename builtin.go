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

package main

import (
	"go/ast"
	"io/ioutil"
	r "reflect"
)

func callAppend(slice interface{}, elems ...interface{}) interface{} {
	elemsv := toValues(elems)
	return r.Append(r.ValueOf(slice), elemsv...)
}

func callCap(arg interface{}) int {
	return r.ValueOf(arg).Cap()
}

func callClose(channel interface{}) {
	r.ValueOf(channel).Close()
}

func callCopy(dst, src interface{}) int {
	return r.Copy(r.ValueOf(dst), r.ValueOf(src))
}

func callDelete(m interface{}, key interface{}) {
	r.ValueOf(m).SetMapIndex(r.ValueOf(key), Nil)
}

func callLen(arg interface{}) int {
	return r.ValueOf(arg).Len()
}

func callPanic(arg interface{}) {
	panic(Panic{arg})
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

func callRecover() interface{} {
	return recover()
}

func callSlice(args ...interface{}) []interface{} {
	return args
}

func (env *Env) addBuiltins() {
	binds := env.binds

	binds["append"] = r.ValueOf(callAppend)
	binds["cap"] = r.ValueOf(callCap)
	binds["close"] = r.ValueOf(callClose)
	binds["copy"] = r.ValueOf(callCopy)
	binds["DeepEqual"] = r.ValueOf(r.DeepEqual)
	binds["delete"] = r.ValueOf(callDelete)
	binds["Eval"] = r.ValueOf(func(node ast.Node) interface{} {
		return toInterface(env.Eval1(node))
	})
	binds["EvalN"] = r.ValueOf(func(node ast.Node) []interface{} {
		return toInterfaces(packValues(env.Eval(node)))
	})
	binds["len"] = r.ValueOf(callLen)
	// binds["new"] = r.ValueOf(callNew) // should be handled specially, its argument is a type
	binds["nil"] = Nil
	binds["panic"] = r.ValueOf(callPanic)
	binds["Parse"] = r.ValueOf(func(src interface{}) ast.Node {
		return env.Parse(src)
	})
	binds["println"] = r.ValueOf(func(args ...interface{}) {
		values := toValues(args)
		env.FprintValues(env.Stdout, values...)
	})
	binds["ReadDir"] = r.ValueOf(callReadDir)
	binds["ReadFile"] = r.ValueOf(callReadFile)
	// binds["recover"] = r.ValueOf(callRecover) // does not work! recover() works only inside a deferred function (but not any function called by it)
	binds["Slice"] = r.ValueOf(callSlice)
	binds["String"] = r.ValueOf(func(args ...interface{}) string {
		return env.toString("", args...)
	})
}

func (env *Env) addInterpretedBuiltins() {
	line := "func not(flag bool) bool { if flag { return false } else { return true } }"
	ast := env.Parse(line)
	env.Eval(ast)
}
