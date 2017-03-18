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

type Env struct {
	*Interpreter
	Binds      map[string]r.Value
	Types      map[string]r.Type
	Proxies    map[string]r.Type
	Outer      *Env
	funcData   *funcData
	iotaOffset int
	Name, Path string
}

type funcData struct {
	defers    []func()
	panicking *interface{} // current panic
}

type Builtin struct {
	Exec func(env *Env, args []ast.Expr) (r.Value, []r.Value)
}

type Macro struct {
	Closure func(args []r.Value) (results []r.Value)
	ArgNum  int
}

type Options uint
type whichMacroExpand uint

const (
	OptTrapPanic Options = 1 << iota
	OptShowPrompt
	OptShowAfterEval
	OptShowAfterParse
	OptShowAfterMacroExpansion
	OptShowEvalDuration
	OptDebugMacroExpand
	OptDebugQuasiquote

	cMacroExpand1 whichMacroExpand = iota
	cMacroExpand
	cMacroExpandCodewalk
)

func (m whichMacroExpand) String() string {
	switch m {
	case cMacroExpand1:
		return "MacroExpand1"
	case cMacroExpandCodewalk:
		return "MacroExpandCodewalk"
	default:
		return "MacroExpand"
	}
}

var Nil = r.Value{}

var none struct{}
var None = r.ValueOf(none)

var One = r.ValueOf(1)

var typeOfInt = r.TypeOf(int(0))
var typeOfInterface = r.TypeOf((*interface{})(nil)).Elem()
var typeOfString = r.TypeOf("")
var typeOfDeferFunc = r.TypeOf(func() {})
var zeroStrings = []string{}

const temporaryFunctionName = "gorepl_temporary_function"
