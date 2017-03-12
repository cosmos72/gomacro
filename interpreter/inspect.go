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
 * inspect.go
 *
 *  Created on: Feb 11, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"bufio"
	"errors"
	"fmt"
	r "reflect"
	"strconv"
	"strings"
)

type Stack struct {
	names []string
	vs    []r.Value
	ts    []r.Type
	in    *bufio.Reader
	env   *Env
}

func (env *Env) Inspect(in *bufio.Reader, name string) {
	ast := env.ParseAst(name)
	v := env.EvalAst1(ast)
	var t r.Type
	if v != Nil && v != None {
		t = r.TypeOf(v.Interface()) // show concrete type
	}
	switch dereferenceValue(v).Kind() {
	case r.Array, r.Slice, r.String, r.Struct:
		break
	default:
		env.showVar(name, v, t)
		return
	}
	stack := Stack{names: []string{name}, vs: []r.Value{v}, ts: []r.Type{t}, in: in, env: env}
	stack.Show()
	stack.Repl()
}

func (env *Env) showVar(name string, v r.Value, t r.Type) {
	env.Fprintf(env.Stdout, "%s\t= %v\t<%v>\n", name, v, t)
}

func (*Stack) Help() {
	fmt.Println("// inspect commands: <number> help quit top up")
}

func (stack *Stack) Show() {
	depth := len(stack.names)
	name := strings.Join(stack.names, ".")
	v := stack.vs[depth-1]
	t := stack.ts[depth-1]
	stack.env.showVar(name, v, t)

	v = dereferenceValue(v) // dereference pointers on-the-fly
	switch v.Kind() {
	case r.Array, r.Slice, r.String:
		stack.showIndexes(v)
	case r.Struct:
		stack.showFields(v)
	}
}

func (stack *Stack) Repl() error {
	for len(stack.names) > 0 {
		fmt.Printf("goinspect %s> ", strings.Join(stack.names, "."))
		cmd, err := stack.in.ReadString('\n')
		if err != nil {
			return err
		}
		cmd = strings.TrimSpace(cmd)
		err = stack.Eval(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (stack *Stack) Eval(cmd string) error {
	switch {
	case cmd == "?", isPrefix(cmd, "help"):
		stack.Help()
	case isPrefix(cmd, "quit"):
		return errors.New("user quit")
	case isPrefix(cmd, "top"):
		stack.Top()
		stack.Show()
	case cmd == "", cmd == ".":
		stack.Show()
	case cmd == "-", isPrefix(cmd, "up"):
		stack.Leave()
	default:
		stack.Enter(cmd)
	}
	return nil
}

func (stack *Stack) Top() {
	stack.names = stack.names[0:1]
	stack.vs = stack.vs[0:1]
	stack.ts = stack.ts[0:1]
}

func (stack *Stack) Leave() {
	depth := len(stack.names)
	if depth <= 0 {
		return
	}
	depth--
	stack.names = stack.names[:depth]
	stack.vs = stack.vs[:depth]
	stack.ts = stack.ts[:depth]
	if depth > 0 {
		stack.Show()
	}
}

func (stack *Stack) showFields(v r.Value) {
	n := v.NumField()
	for i := 0; i < n; i++ {
		f := v.Field(i)
		t := TypeOf(f)
		f = dereferenceValue(f)
		fmt.Printf("    %d. ", i)
		stack.env.showVar(v.Type().Field(i).Name, f, t)
	}
}

func (stack *Stack) showIndexes(v r.Value) {
	n := v.Len()
	for i := 0; i < n; i++ {
		f := v.Index(i)
		t := TypeOf(f)
		f = dereferenceValue(f)
		fmt.Printf("    %d. ", i)
		stack.env.showVar("", f, t)
	}
}

func (stack *Stack) Enter(cmd string) {
	i, err := strconv.Atoi(cmd)
	if err != nil {
		fmt.Fprintf(stack.env.Stdout, "unknown inspect command \"%s\". Type ? for help\n", cmd)
		return
	}
	depth := len(stack.names)
	v := dereferenceValue(stack.vs[depth-1])
	var n int
	var fname string
	var f r.Value
	switch v.Kind() {
	case r.Array, r.Slice, r.String:
		n = v.Len()
		if !stack.validRange(i, n) {
			return
		}
		fname = fmt.Sprintf("[%s]", cmd)
		f = v.Index(i)
	case r.Struct:
		n = v.NumField()
		if !stack.validRange(i, n) {
			return
		}
		fname = v.Type().Field(i).Name
		f = v.Field(i)
	default:
		fmt.Fprintf(stack.env.Stdout, "cannot enter <%v>: expecting array, slice, string or struct\n", TypeOf(v))
		return
	}
	var t r.Type
	if f != Nil && f != None {
		t = r.TypeOf(f.Interface()) // concrete type
	}

	switch dereferenceValue(f).Kind() { // dereference pointers on-the-fly
	case r.Array, r.Slice, r.String, r.Struct:
		stack.names = append(stack.names, fname)
		stack.vs = append(stack.vs, f)
		stack.ts = append(stack.ts, t)
		stack.Show()
	default:
		stack.env.showVar(fname, f, t)
	}
}

func dereferenceValue(v r.Value) r.Value {
	for {
		switch v.Kind() {
		case r.Ptr:
			v = v.Elem()
			continue
		case r.Interface:
			v = r.ValueOf(v.Interface())
			continue
		}
		break
	}
	return v
}

func (stack *Stack) validRange(i, n int) bool {
	if i < 0 || i >= n {
		fmt.Fprintf(stack.env.Stdout, "%s contains %d elements, cannot inspect element %d\n",
			strings.Join(stack.names, "."), n, i)
		return false
	}
	return true
}
