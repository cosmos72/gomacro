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
 *  Created on: Feb 15, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	r "reflect"
)

func callAppend(slice interface{}, elems ...interface{}) interface{} {
	var elemsv []r.Value
	for i := range elems {
		elemsv[i] = r.ValueOf(elems[i])
	}
	return r.Append(r.ValueOf(slice), elemsv...)
}

func callCap(arg interface{}) interface{} {
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

func callLen(arg interface{}) interface{} {
	return r.ValueOf(arg).Len()
}

func callPanic(arg interface{}) {
	panic(arg)
}

func callRecover() interface{} {
	return recover()
}

func addBuiltins(binds Binds) {
	binds["append"] = r.ValueOf(callAppend)
	binds["cap"] = r.ValueOf(callCap)
	binds["close"] = r.ValueOf(callClose)
	binds["copy"] = r.ValueOf(callCopy)
	binds["delete"] = r.ValueOf(callDelete)
	binds["len"] = r.ValueOf(callLen)
	// binds["new"] = r.ValueOf(callNew) // should be handled specially, its argument is a type
	binds["panic"] = r.ValueOf(callPanic)
	// binds["recover"] = r.ValueOf(callRecover) // does not work! recover() works only inside a deferred function (but not any function called by it)
}
