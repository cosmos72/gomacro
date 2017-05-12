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

package classic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	r "reflect"
	"runtime/debug"
	"strings"
	"time"

	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/imports"
)

type Env struct {
	*ThreadGlobals
	Binds      BindMap
	Types      TypeMap
	Proxies    TypeMap
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
		iotaOffset: 1,
		Outer:      outer,
		Name:       path,
		Path:       path,
	}
	if outer == nil {
		env.ThreadGlobals = NewThreadGlobals()
		env.CallStack = &CallStack{Frames: []CallFrame{CallFrame{}}}
		env.addBuiltins()
		env.addInterpretedBuiltins()
	} else {
		env.ThreadGlobals = outer.ThreadGlobals
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

func (env *Env) AsPackage() imports.Package {
	return imports.Package{
		Binds:   env.Binds.AsMap(),
		Types:   env.Types.AsMap(),
		Proxies: env.Proxies.AsMap(),
	}
}

func (env *Env) MergePackage(pkg imports.Package) {
	env.Binds.Ensure().Merge(pkg.Binds)
	env.Types.Ensure().Merge(pkg.Types)
	env.Proxies.Ensure().Merge(pkg.Proxies)
}

func (env *Env) ChangePackage(name string) *Env {
	fenv := env.FileEnv()
	currname := fenv.ThreadGlobals.Packagename
	if name == currname {
		return env
	}
	fenv.AsPackage().SaveToPackages(currname)

	nenv := NewEnv(fenv.TopEnv(), name)
	nenv.MergePackage(imports.Packages[name])
	nenv.ThreadGlobals.Packagename = name

	return nenv
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

// ValueOf returns the value of a constant, function or variable.
// for variables, the returned reflect.Value is settable and addressable
// returns the zero reflect.Value if not found
func (env *Env) ValueOf(name string) (value r.Value) {
	found := false
	for e := env; e != nil; e = e.Outer {
		if value, found = e.Binds.Get(name); found {
			break
		}
	}
	return
}

func (env *Env) ReplStdin() {
	if env.Options&OptShowPrompt != 0 {
		fmt.Fprint(env.Stdout, `// GOMACRO, an interactive Go interpreter with macros <https://github.com/cosmos72/gomacro>
// Copyright (C) 2017 Massimiliano Ghilardi
// License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
// This is free software with ABSOLUTELY NO WARRANTY.
//
// Type :help for help
`)
	}
	in := bufio.NewReader(os.Stdin)

	env.Line = 0
	for env.ReadParseEvalPrint(in) {
		env.Line = 0
	}
}

func (env *Env) Repl(in *bufio.Reader) {
	for env.ReadParseEvalPrint(in) {
	}
}

func (env *Env) ReadParseEvalPrint(in *bufio.Reader) (callAgain bool) {
	var opts ReadOptions
	if env.Options&OptShowPrompt != 0 {
		opts |= ReadOptShowPrompt
	}
	str, firstToken := env.ReadMultiline(in, opts)
	if firstToken < 0 {
		// skip comments and continue, but fail on EOF or other errors
		env.IncLine(str)
		return len(str) > 0
	} else if firstToken > 0 {
		env.IncLine(str[0:firstToken])
	}
	return env.ParseEvalPrint(str[firstToken:], in)
}

func (env *Env) ReadMultiline(in *bufio.Reader, opts ReadOptions) (str string, firstToken int) {
	str, firstToken, err := ReadMultiline(in, opts, env.Stdout, "gomacro> ")
	if err != nil && err != io.EOF {
		fmt.Fprintf(env.Stderr, "// read error: %s\n", err)
	}
	return str, firstToken
}

func (env *Env) ParseEvalPrint(str string, in *bufio.Reader) (callAgain bool) {
	var t1 time.Time
	trap := env.Options&OptTrapPanic != 0
	duration := env.Options&OptShowTime != 0
	if duration {
		t1 = time.Now()
	}
	defer func() {
		env.IncLine(str)
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
	return env.parseEvalPrint(str, in)
}

func (env *Env) parseEvalPrint(src string, in *bufio.Reader) (callAgain bool) {

	src = strings.TrimSpace(src)
	n := len(src)
	if n == 0 {
		return true // no input. don't print anything
	}
	fast := env.Options&OptFastInterpreter != 0 // use the fast interpreter?

	if n > 0 && src[0] == ':' {
		args := strings.SplitN(src, " ", 2)
		cmd := args[0]
		switch {
		case strings.HasPrefix(":classic", cmd):
			if len(args) <= 1 {
				if env.Options&OptFastInterpreter != 0 {
					env.Debugf("switched to classic interpreter")
				}
				env.Options &^= OptFastInterpreter
				return true
			}
			// temporary override
			src = strings.TrimSpace(args[1])
			fast = false
		case strings.HasPrefix(":env", cmd):
			if len(args) <= 1 {
				env.showPackage("")
			} else {
				env.showPackage(args[1])
			}
			return true
		case strings.HasPrefix(":fast", cmd):
			if len(args) <= 1 {
				if env.Options&OptFastInterpreter == 0 {
					env.Debugf("switched to fast interpreter (incomplete)")
				}
				env.Options |= OptFastInterpreter
				return true
			}
			// temporary override
			src = strings.TrimSpace(args[1])
			fast = true
		case strings.HasPrefix(":help", cmd):
			env.showHelp(env.Stdout)
			return true
		case strings.HasPrefix(":inspect", cmd):
			if in == nil {
				fmt.Fprint(env.Stdout, "// not connected to user input, cannot :inspect\n")
			} else if len(args) == 1 {
				fmt.Fprint(env.Stdout, "// inspect: missing argument\n")
			} else {
				env.Inspect(in, args[1], fast)
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
			src = " " + src[1:] // slower than src = src[1:], but gives accurate column positions in error messages
		}
	}
	if src == "package" || src == " package" || strings.HasPrefix(src, "package ") || strings.HasPrefix(src, " package ") {
		arg := ""
		src = strings.TrimSpace(src)
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
	form := env.Parse(src)

	var value r.Value
	var values []r.Value

	// eval phase
	if form != nil {
		if env.Options&OptMacroExpandOnly != 0 {
			value = r.ValueOf(form.Interface())
		} else if fast {
			value, values = env.fastEval(form)
		} else {
			value, values = env.EvalAst(form)
		}
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
