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
 * globals.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	r "reflect"
)

type Binds map[string]r.Value
type Types map[string]r.Type

type Env struct {
	*Interpreter
	Binds      Binds
	Types      Types
	Proxies    Types
	Outer      *Env
	iotaOffset int
	Name, Path string
}

type Builtin struct {
	Exec func(env *Env, args ...ast.Expr) (r.Value, []r.Value)
}

type Macro struct {
	Closure func(args []r.Value) (results []r.Value)
	ArgNum  int
}

type Options uint

const (
	OptTrapPanic Options = 1 << iota
	OptShowEvalDuration
	OptShowAfterParse
	OptShowAfterMacroExpansion
	OptDebugMacroExpand
	OptDebugQuasiquote
)

var Nil = r.Value{}

var none struct{}
var None = r.ValueOf(none)

var One = r.ValueOf(1)

var typeOfInterface = r.TypeOf((*interface{})(nil)).Elem()
var typeOfString = r.TypeOf("")
var zeroStrings = []string{}

const temporaryFunctionName = "gorepl_temporary_function"
