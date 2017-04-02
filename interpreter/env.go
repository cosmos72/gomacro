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
	"strings"
	"time"

	. "github.com/cosmos72/gomacro/ast2"
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
	in := bufio.NewReader(os.Stdin)
	env.Repl(in)
}

func (env *Env) Repl(in *bufio.Reader) {
	if env.Options&OptShowPrompt != 0 {
		fmt.Fprint(env.Stdout, "// Welcome to gomacro. Type :help for help\n")
	}

	for env.ReadParseEvalPrint(in) {
	}
}

func (env *Env) ReadParseEvalPrint(in *bufio.Reader) (callAgain bool) {
	str, err := ReadMultiline(in, env.Options&OptShowPrompt != 0, env.Stdout, "gomacro> ")
	if err != nil {
		if err != io.EOF {
			fmt.Fprintln(env.Stderr, err)
		}
		return false
	}

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
					fmt.Fprintln(env.Stderr, rec)
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
		case hasPrefix(":compiler", cmd):
			if len(args) > 1 {
				env.compile(args[1])
			}
			return true
		case hasPrefix(":env", cmd):
			if len(args) <= 1 {
				env.showPackage("")
			} else {
				env.showPackage(args[1])
			}
			return true
		case hasPrefix(":help", cmd):
			env.showHelp(env.Stdout)
			return true
		case hasPrefix(":inspect", cmd):
			if in == nil {
				fmt.Fprint(env.Stdout, "// not connected to user input, cannot :inspect\n")
			} else if len(args) == 1 {
				fmt.Fprint(env.Stdout, "// inspect: missing argument\n")
			} else {
				env.Inspect(in, args[1])
			}
			return true
		case hasPrefix(":options", cmd):
			if len(args) > 1 {
				env.Options ^= parseOptions(args[1])
			}
			fmt.Fprintf(env.Stdout, "// current options: %v\n", env.Options)
			fmt.Fprintf(env.Stdout, "// unset   options: %v\n", ^env.Options)
			return true
		case hasPrefix(":quit", cmd):
			return false
		case hasPrefix(":write", cmd):
			if len(args) <= 1 {
				env.writeDeclsToStream(env.Stdout)
			} else {
				env.writeDeclsToFile(args[1])
			}
			return true
		default:
			switch cmd {
			case ":const", ":func", ":import", ":type", ":var":
				// temporarily disable collection of declarations and statements
				saveOpts := env.Options
				collect := OptCollectDeclarations | OptCollectStatements
				if saveOpts&collect != 0 {
					env.Options &^= collect
					defer func() {
						env.Options = saveOpts
					}()
				}
				src = src[1:]
			}
		}
	}
	if src == "package" || hasPrefix(src, "package ") {
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

	// eval phase
	value, values := env.EvalAst(ast)

	// print phase
	if env.Options&OptShowEval != 0 {
		if len(values) != 0 {
			env.fprintValues(env.Stdout, values...)
		} else if value != None {
			env.fprintValues(env.Stdout, value)
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
		env.collectAst(form)
	}
	return form
}
