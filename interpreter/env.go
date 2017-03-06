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
	"os"
	r "reflect"
	"strings"
	"time"
)

func NewEnv(outer *Env, path string) *Env {
	env := &Env{
		Binds:      make(map[string]r.Value),
		Types:      make(map[string]r.Type),
		iotaOffset: 1,
		Outer:      outer,
		Name:       path,
		Path:       path,
	}
	if outer == nil {
		env.Interpreter = NewInterpreter()
		env.addBuiltins()
		env.addInterpretedBuiltins()

		/*
			type Foo struct{ a, b int }
			type Bar struct{ a, b int }
			var foo Foo
			var bar Bar = Bar(foo)
			t := r.TypeOf(bar)
			var tf interface{} = t
			fmt.Printf("typeof(bar) = %v - actually %#v\n", t, tf)
			fmt.Printf("typeof(bar) Name = %v, Kind = %v\n", t.Name(), t.Kind())
		*/

	} else {
		env.Interpreter = outer.Interpreter
	}

	// fmt.Printf("NewEnv(): env = %p %q, outer = %p\n", env, env.Path, env.Outer)
	return env
}

func (env *Env) TopEnv() *Env {
	for {
		outer := env.Outer
		// fmt.Printf("TopEnv(): env = %p %q, outer = %p\n", env, env.Path, outer)
		// time.Sleep(time.Second)
		if outer == nil {
			break
		}
		env = outer
	}
	return env
}

func (env *Env) FileEnv() *Env {
	for {
		outer := env.Outer
		// fmt.Printf("FileEnv(): env = %p %q, outer = %p\n", env, env.Path, outer)
		// time.Sleep(time.Second)
		if outer == nil || outer.Outer == nil {
			break
		}
		env = outer
	}
	return env
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

func (env *Env) ReadParseEvalPrint(in *bufio.Reader) (ret bool) {
	if env.Options&OptTrapPanic != 0 {
		defer func() {
			if rec := recover(); rec != nil {
				fmt.Fprintln(env.Stderr, rec)
				ret = true
			}
		}()
	}

	fmt.Fprint(env.Stdout, "gomacro> ")

	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Fprintln(env.Stderr, err)
		return false
	}
	return env.ParseEvalPrint(line)
}

func (env *Env) ParseEvalPrint(str string) (ret bool) {
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
		env.FprintValues(env.Stdout) // no value
		return true
	} else if n > 0 && src[0] == ':' {
		args := strings.SplitN(src, " ", 2)
		switch args[0] {
		case ":q", ":quit":
			return false
		case ":e", ":env":
			if len(args) <= 1 {
				env.showPackage(env.Stdout, "")
			} else {
				env.showPackage(env.Stdout, args[1])
			}
			return true
		case ":h", ":help":
			env.showHelp(env.Stdout)
			return true
		}
	}
	// parse phase
	list := env.Parse(src)
	if env.Options&OptShowAfterParse != 0 {
		env.Debugf("after parse: %v", list)
	}

	// macroexpansion phase
	for i, elt := range list {
		list[i], _ = env.MacroExpandCodewalk(elt)
	}
	if env.Options&OptShowAfterMacroExpansion != 0 {
		env.Debugf("after macroexpansion: %v", list)
	}

	// eval phase
	value, values := env.EvalList(list)

	// print phase
	if len(values) != 0 {
		env.FprintValues(env.Stdout, values...)
	} else {
		env.FprintValues(env.Stdout, value)
	}
	return true
}
