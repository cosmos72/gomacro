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
	"strings"
	"time"
)

func New() *Env {
	top := NewEnv(nil, "builtin")
	return NewEnv(top, "main")
}

func NewEnv(outer *Env, path string) *Env {
	env := &Env{
		Binds:      make(map[string]r.Value),
		Types:      nil,
		Proxies:    nil,
		iotaOffset: 1,
		Outer:      outer,
		Name:       path,
		Path:       path,
	}
	if outer == nil {
		env.Interpreter = NewInterpreter()
		env.addBuiltins()
		env.addInterpretedBuiltins()
	} else {
		env.Interpreter = outer.Interpreter
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

func (env *Env) FuncEnv() *Env {
	for ; env != nil; env = env.Outer {
		if env.funcData != nil {
			break
		}
	}
	return env
}

func (env *Env) CallerEnv() *Env {
	return env.FuncEnv().Outer.FuncEnv()
}

func (env *Env) ReplStdin() {
	in := bufio.NewReader(os.Stdin)
	env.Repl(in)
}

func (env *Env) Repl(in *bufio.Reader) {
	fmt.Fprint(env.Stdout, "// Welcome to gomacro. Type :help for help\n")

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

	if env.Options&OptTrapPanic != 0 {
		defer func() {
			if rec := recover(); rec != nil {
				fmt.Fprintln(env.Stderr, rec)
				callAgain = true
			}
		}()
	}
	return env.ParseEvalPrint(str, in)
}

func (env *Env) ParseEvalPrint(str string, in *bufio.Reader) (callAgain bool) {
	if env.Options&OptShowEvalDuration != 0 {
		t1 := time.Now()
		defer func() {
			delta := time.Now().Sub(t1)
			env.Debugf("eval time %.6f s", float32(delta)/float32(time.Second))
		}()
	}

	src := strings.TrimSpace(str)
	n := len(src)

	if n == 0 {
		if env.Options&OptShowAfterEval != 0 {
			env.FprintValues(env.Stdout) // no value
		}
		return true
	} else if n > 0 && src[0] == ':' {
		args := strings.SplitN(src, " ", 2)
		cmd := args[0]
		switch {
		case isPrefix(cmd, ":quit"):
			return false
		case isPrefix(cmd, ":inspect"):
			if in == nil {
				fmt.Fprint(env.Stdout, "// not connected to user input, cannot :inspect\n")
			} else if len(args) == 1 {
				fmt.Fprint(env.Stdout, "// inspect: missing argument\n")
			} else {
				env.Inspect(in, args[1])
			}
			return true
		case isPrefix(cmd, ":env"):
			if len(args) <= 1 {
				env.showPackage(env.Stdout, "")
			} else {
				env.showPackage(env.Stdout, args[1])
			}
			return true
		case isPrefix(cmd, ":help"):
			env.showHelp(env.Stdout)
			return true
		}
	}
	// parse phase
	ast := env.ParseAst(src)
	if env.Options&OptShowAfterParse != 0 {
		env.Debugf("after parse: %v", ast.Interface())
	}

	// macroexpansion phase.
	ast, _ = env.MacroExpandAstCodewalk(ast)

	if env.Options&OptShowAfterMacroExpansion != 0 {
		env.Debugf("after macroexpansion: %v", ast.Interface())
	}

	// eval phase
	value, values := env.EvalAst(ast)

	// print phase
	if env.Options&OptShowAfterEval != 0 {
		if len(values) != 0 {
			env.FprintValues(env.Stdout, values...)
		} else {
			env.FprintValues(env.Stdout, value)
		}
	}
	return true
}

func isPrefix(prefix, str string) bool {
	n := len(prefix)
	return n <= len(str) && prefix == str[:n]
}
