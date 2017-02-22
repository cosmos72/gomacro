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

package main

import (
	"bufio"
	"fmt"
	"os"
	r "reflect"
	"strings"
	"time"
)

type Binds map[string]r.Value

type Env struct {
	*Interpreter
	binds      Binds
	Outer      *Env
	iotaOffset int
}

func NewEnv(outer *Env) *Env {
	env := Env{}
	env.binds = make(map[string]r.Value)
	env.iotaOffset = 1
	if outer == nil {
		env.Interpreter = NewInterpreter()
		env.addBuiltins()
		env.addInterpretedBuiltins()
	} else {
		env.Interpreter = outer.Interpreter
		env.Outer = outer
	}
	return &env
}

func (env *Env) Repl() {
	in := bufio.NewReader(os.Stdin)
	env.Repl1(in)
}

func (env *Env) Repl1(in *bufio.Reader) {
	fmt.Fprint(env.Stdout, "// Welcome to gomacro. Type :help for help\n")

	for env.ReadParseEvalPrint(in) {
	}
}

func (env *Env) ReadParseEvalPrint(in *bufio.Reader) (ret bool) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Fprintln(env.Stderr, rec)
			ret = true
		}
	}()

	fmt.Fprint(env.Stdout, "go> ")

	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Fprintln(env.Stderr, err)
		return false
	}
	return env.ParseEvalPrint(line)
}

func (env *Env) ParseEvalPrint(str string) (ret bool) {
	t1 := time.Now()
	defer func() {
		delta := time.Now().Sub(t1)
		env.Debugf("eval() elapsed time: %g s", float64(delta)/float64(time.Second))
	}()

	src := []byte(strings.TrimSpace(str))
	pos := findFirstToken(src)
	trimmed := src[pos:]
	ntrimmed := len(trimmed)

	if ntrimmed == 0 {
		env.FprintValues(env.Stdout) // no value
		return true
	} else if ntrimmed > 0 && trimmed[0] == ':' {
		_cmd := string(extractFirstIdentifier(trimmed[1:]))
		if _cmd == "quit" {
			return false
		} else if _cmd == "env" {
			env.showEnv(env.Stdout)
			return true
		} else if _cmd == "help" {
			env.showHelp(env.Stdout)
			return true
		}
	}
	ast := env.Parse(src)
	// env.FprintValue(env.Stdout, r.ValueOf(ast))
	value, values := env.EvalList(ast)
	if len(values) != 0 {
		env.FprintValues(env.Stdout, values...)
	} else {
		env.FprintValues(env.Stdout, value)
	}
	return true
}
