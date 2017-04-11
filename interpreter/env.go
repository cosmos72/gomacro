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
 * env.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	r "reflect"
	"runtime/debug"
	"strings"
	"time"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/imports"
)

type Env struct {
	*InterpreterCommon
	imports.Package
	Outer      *Env
	CallStack  *CallStack
	iotaOffset int
	Name, Path string
}

func New() *Env {
	top := NewEnv(nil, "builtin")
	return NewEnv(top, "main")
}

func NewEnv(outer *Env, path string) *Env {
	env := &Env{
		Package:    imports.Package{},
		iotaOffset: 1,
		Outer:      outer,
		Name:       path,
		Path:       path,
	}
	if outer == nil {
		env.InterpreterCommon = NewInterpreterCommon()
		env.CallStack = &CallStack{Frames: []CallFrame{CallFrame{}}}
		env.addBuiltins()
		env.addInterpretedBuiltins()
	} else {
		env.InterpreterCommon = outer.InterpreterCommon
		env.CallStack = outer.CallStack
	}
	return env
}

func (env *Env) TopEnv() *Env {
	for ; env != nil; env = env.Outer {
		if env.Outer == nil {
			break
		}
	}
	return env
}

func (env *Env) FileEnv() *Env {
	for ; env != nil; env = env.Outer {
		outer := env.Outer
		if outer == nil || outer.Outer == nil {
			break
		}
	}
	return env
}

// CurrentFrame returns the CallFrame representing the current function call
func (env *Env) CurrentFrame() *CallFrame {
	if env != nil {
		frames := env.CallStack.Frames
		if n := len(frames); n > 0 {
			return &frames[n-1]
		}
	}
	return nil
}

// CallerFrame returns the CallFrame representing the caller's function.
// needed by recover()
func (env *Env) CallerFrame() *CallFrame {
	if env != nil {
		frames := env.CallStack.Frames
		if n := len(frames); n > 1 {
			return &frames[n-2]
		}
	}
	return nil
}

func (env *Env) ReplStdin() {
	if env.Options&OptShowPrompt != 0 {
		fmt.Fprint(env.Stdout, "// Welcome to gomacro. Type :help for help\n")
	}
	in := bufio.NewReader(os.Stdin)
	env.Repl(in)
}

func (env *Env) Repl(in *bufio.Reader) {
	for env.ReadParseEvalPrint(in) {
	}
}

func (env *Env) ReadParseEvalPrint(in *bufio.Reader) (callAgain bool) {
	var readopts ReadOptions
	if env.Options&OptShowPrompt != 0 {
		readopts |= ReadOptShowPrompt
	}

	str, _ := env.ReadMultiline(in, readopts)
	return len(str) > 0 && env.ParseEvalPrintRecover(str, in)
}

func (env *Env) ReadMultiline(in *bufio.Reader, readopts ReadOptions) (str string, comments string) {
	str, comments, err := ReadMultiline(in, readopts, env.Stdout, "gomacro> ")
	if err != nil && err != io.EOF {
		fmt.Fprintln(env.Stderr, err)
	}
	return str, comments
}

func (env *Env) ParseEvalPrintRecover(str string, in *bufio.Reader) (callAgain bool) {

	trap := env.Options&OptTrapPanic != 0
	duration := env.Options&OptShowTime != 0
	if trap || duration {
		var t1 time.Time
		if duration {
			t1 = time.Now()
		}
		defer func() {
			if trap {
				if rec := recover(); rec != nil {
					if env.Options&OptPanicStackTrace != 0 {
						fmt.Fprintf(env.Stderr, "%s\n%s", rec, debug.Stack())
					} else {
						fmt.Fprintf(env.Stderr, "%s\n", rec)
					}
					callAgain = true
				}
			}
			if duration {
				delta := time.Now().Sub(t1)
				env.Debugf("eval time %.6f s", float32(delta)/float32(time.Second))
			}
		}()
	}
	return env.ParseEvalPrint(str, in)
}

func (env *Env) ParseEvalPrint(src string, in *bufio.Reader) (callAgain bool) {

	src = strings.TrimSpace(src)
	n := len(src)
	if n == 0 {
		return true // no input. don't print anything
	}
	if n > 0 && src[0] == ':' {
		args := strings.SplitN(src, " ", 2)
		cmd := args[0]
		switch {
		case strings.HasPrefix(":compiler", cmd):
			if len(args) > 1 {
				env.compile(args[1])
			}
			return true
		case strings.HasPrefix(":env", cmd):
			if len(args) <= 1 {
				env.showPackage("")
			} else {
				env.showPackage(args[1])
			}
			return true
		case strings.HasPrefix(":help", cmd):
			env.showHelp(env.Stdout)
			return true
		case strings.HasPrefix(":inspect", cmd):
			if in == nil {
				fmt.Fprint(env.Stdout, "// not connected to user input, cannot :inspect\n")
			} else if len(args) == 1 {
				fmt.Fprint(env.Stdout, "// inspect: missing argument\n")
			} else {
				env.Inspect(in, args[1])
			}
			return true
		case strings.HasPrefix(":options", cmd):
			if len(args) > 1 {
				env.Options ^= ParseOptions(args[1])
			}
			fmt.Fprintf(env.Stdout, "// current options: %v\n", env.Options)
			fmt.Fprintf(env.Stdout, "// unset   options: %v\n", ^env.Options)
			return true
		case strings.HasPrefix(":quit", cmd):
			return false
		case strings.HasPrefix(":write", cmd):
			if len(args) <= 1 {
				env.WriteDeclsToStream(env.Stdout)
			} else {
				env.WriteDeclsToFile(args[1])
			}
			return true
		default:
			// temporarily disable collection of declarations and statements,
			// and temporarily disable macroexpandonly (i.e. re-enable eval)
			saved := env.Options
			todisable := OptMacroExpandOnly | OptCollectDeclarations | OptCollectStatements
			if saved&todisable != 0 {
				env.Options &^= todisable
				defer func() {
					env.Options = saved
				}()
			}
			src = src[1:]
		}
	}
	if src == "package" || strings.HasPrefix(src, "package ") {
		arg := ""
		space := strings.IndexByte(src, ' ')
		if space >= 0 {
			arg = strings.TrimSpace(src[1+space:])
		}
		if len(arg) == 0 {
			fmt.Fprintf(env.Stdout, "// current package: %v\n", env.Packagename)
		} else {
			env = env.ChangePackage(arg)
		}
		return true
	}

	// parse + macroexpansion phase
	ast := env.ParseAst(src)

	var value r.Value
	var values []r.Value

	// eval phase
	if env.Options&OptMacroExpandOnly == 0 {
		value, values = env.EvalAst(ast)
	} else if ast != nil {
		value = r.ValueOf(ast.Interface())
	}

	// print phase
	if env.Options&OptShowEval != 0 {
		if len(values) != 0 {
			env.FprintValues(env.Options, env.Stdout, values...)
		} else if value != None {
			env.FprintValues(env.Options, env.Stdout, value)
		}
	}
	return true
}

func (env *Env) ParseAst(src interface{}) Ast {
	bytes := ReadBytes(src)
	nodes := env.ParseBytes(bytes)

	if env.Options&OptShowParse != 0 {
		env.Debugf("after parse: %v", nodes)
	}

	var form Ast
	switch len(nodes) {
	case 0:
		return nil
	case 1:
		form = ToAst(nodes[0])
	default:
		form = NodeSlice{X: nodes}
	}

	// macroexpansion phase.
	form, _ = env.MacroExpandAstCodewalk(form)

	if env.Options&OptShowMacroExpand != 0 {
		env.Debugf("after macroexpansion: %v", form.Interface())
	}
	if env.Options&(OptCollectDeclarations|OptCollectStatements) != 0 {
		env.CollectAst(form)
	}
	return form
}
