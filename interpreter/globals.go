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
	"sort"
	"strings"

	"github.com/cosmos72/gomacro/constants"
	"github.com/cosmos72/gomacro/imports"
)

type CallStack struct {
	Frames []CallFrame
}

type CallFrame struct {
	FuncEnv       *Env
	InnerEnv      *Env          // innermost Env
	CurrentCall   *ast.CallExpr // call currently in progress
	defers        []func()
	panick        interface{} // current panic
	panicking     bool
	runningDefers bool
}

type Builtin struct {
	Exec   func(env *Env, args []ast.Expr) (r.Value, []r.Value)
	ArgNum int // if negative, do not check
}

type Function struct {
	Exec   func(env *Env, args []r.Value) (r.Value, []r.Value)
	ArgNum int // if negative, do not check
}

type Macro struct {
	Closure func(args []r.Value) (results []r.Value)
	ArgNum  int
}

type PackageRef struct {
	imports.Package
	Name, Path string
}

type TypedValue struct {
	Type  r.Type
	Value r.Value
}

/**
 * inside Methods, each string is the method name
 * and each TypedValue is {
 *   Type: the method signature, i.e. the type of a func() *without* the receiver (to allow comparison with Interface methods)
 *   Value: the method implementation, i.e. a func() whose first argument is the receiver,
 * }
 */
type Methods map[string]TypedValue

/**
 * Interface is the interpreted version of Golang interface values.
 * Each Interface contains {
 *   Type:  the interface type. returned by Env.evalInterface(), i.e. the type of a struct {  interface{}; functions... }
 *   Value: the datum implementing the interface. Value.Type() must be its concrete type, i.e. == r.TypeOf(Value.Interface())
 * }
 */
type Interface TypedValue

type Options uint
type whichMacroExpand uint

const (
	OptTrapPanic Options = 1 << iota
	OptShowPrompt
	OptShowParse
	OptShowMacroExpand
	OptShowCompile
	OptShowEval
	OptShowTime
	OptDebugMacroExpand
	OptDebugQuasiquote
	OptDebugCallStack
	OptDebugPanicRecover
	OptCollectDeclarations
	OptCollectStatements

	cMacroExpand1 whichMacroExpand = iota
	cMacroExpand
	cMacroExpandCodewalk
)

var optNames = map[Options]string{
	OptTrapPanic:           "TrapPanic",
	OptShowPrompt:          "Prompt",
	OptShowParse:           "Parse",
	OptShowMacroExpand:     "MacroExpand",
	OptShowCompile:         "Compile",
	OptShowEval:            "Eval",
	OptShowTime:            "Time",
	OptDebugMacroExpand:    "?MacroExpand",
	OptDebugQuasiquote:     "?Quasiquote",
	OptDebugCallStack:      "?CallStack",
	OptDebugPanicRecover:   "?PanicRecover",
	OptCollectDeclarations: "Declarations",
	OptCollectStatements:   "Statements",
}

var optValues = map[string]Options{}

func init() {
	for k, v := range optNames {
		optValues[v] = k
	}
}

func (o Options) String() string {
	names := make([]string, 0)
	for k, v := range optNames {
		if k&o != 0 {
			names = append(names, v)
		}
	}
	sort.Strings(names)
	return strings.Join(names, " ")
}

func parseOptions(str string) Options {
	var opts Options
	for _, name := range strings.Split(str, " ") {
		if opt, ok := optValues[name]; ok {
			opts ^= opt
		} else if len(name) != 0 {
			for k, v := range optNames {
				if hasPrefix(v, name) {
					opts ^= k
				}
			}
		}
	}
	return opts
}

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

var Nil = constants.Nil
var None = constants.None // used to indicate "no value"

var one = r.ValueOf(1)

var typeOfInt = r.TypeOf(int(0))
var typeOfRune = r.TypeOf(rune(0))
var typeOfInterface = r.TypeOf((*interface{})(nil)).Elem()
var typeOfString = r.TypeOf("")
var typeOfDeferFunc = r.TypeOf(func() {})

var valueOfFalse = r.ValueOf(false)
var valueOfTrue = r.ValueOf(true)

var zeroStrings = []string{}
