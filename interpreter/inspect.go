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

type Inspector struct {
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
	stack := Inspector{names: []string{name}, vs: []r.Value{v}, ts: []r.Type{t}, in: in, env: env}
	stack.Show()
	stack.Repl()
}

func (env *Env) showVar(name string, v r.Value, t r.Type) {
	env.fprintf(env.Stdout, "%s\t= %v\t<%v>\n", name, v, t)
}

func (*Inspector) Help() {
	fmt.Println("// inspect commands: <number> help quit top up")
}

func (ip *Inspector) Show() {
	depth := len(ip.names)
	name := strings.Join(ip.names, ".")
	v := ip.vs[depth-1]
	t := ip.ts[depth-1]
	ip.env.showVar(name, v, t)

	v = dereferenceValue(v) // dereference pointers on-the-fly
	switch v.Kind() {
	case r.Array, r.Slice, r.String:
		ip.showIndexes(v)
	case r.Struct:
		ip.showFields(v)
	}
}

func (ip *Inspector) Repl() error {
	for len(ip.names) > 0 {
		fmt.Fprintf(ip.env.Stdout, "goinspect %s> ", strings.Join(ip.names, "."))
		cmd, err := ip.in.ReadString('\n')
		if err != nil {
			return err
		}
		cmd = strings.TrimSpace(cmd)
		err = ip.Eval(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ip *Inspector) Eval(cmd string) error {
	switch {
	case cmd == "?", startsWith("help", cmd):
		ip.Help()
	case startsWith(cmd, "quit"):
		return errors.New("user quit")
	case startsWith("top", cmd):
		ip.Top()
		ip.Show()
	case cmd == "", cmd == ".":
		ip.Show()
	case cmd == "-", startsWith("up", cmd):
		ip.Leave()
	default:
		ip.Enter(cmd)
	}
	return nil
}

func (ip *Inspector) Top() {
	ip.names = ip.names[0:1]
	ip.vs = ip.vs[0:1]
	ip.ts = ip.ts[0:1]
}

func (ip *Inspector) Leave() {
	depth := len(ip.names)
	if depth <= 0 {
		return
	}
	depth--
	ip.names = ip.names[:depth]
	ip.vs = ip.vs[:depth]
	ip.ts = ip.ts[:depth]
	if depth > 0 {
		ip.Show()
	}
}

func (ip *Inspector) showFields(v r.Value) {
	n := v.NumField()
	for i := 0; i < n; i++ {
		f := v.Field(i)
		t := typeOf(f)
		f = dereferenceValue(f)
		fmt.Fprintf(ip.env.Stdout, "    %d. ", i)
		ip.env.showVar(v.Type().Field(i).Name, f, t)
	}
}

func (ip *Inspector) showIndexes(v r.Value) {
	n := v.Len()
	for i := 0; i < n; i++ {
		f := v.Index(i)
		t := typeOf(f)
		f = dereferenceValue(f)
		fmt.Fprintf(ip.env.Stdout, "    %d. ", i)
		ip.env.showVar("", f, t)
	}
}

func (ip *Inspector) Enter(cmd string) {
	i, err := strconv.Atoi(cmd)
	if err != nil {
		fmt.Fprintf(ip.env.Stdout, "unknown inspect command \"%s\". Type ? for help\n", cmd)
		return
	}
	depth := len(ip.names)
	v := dereferenceValue(ip.vs[depth-1])
	var n int
	var fname string
	var f r.Value
	switch v.Kind() {
	case r.Array, r.Slice, r.String:
		n = v.Len()
		if !ip.validRange(i, n) {
			return
		}
		fname = fmt.Sprintf("[%s]", cmd)
		f = v.Index(i)
	case r.Struct:
		n = v.NumField()
		if !ip.validRange(i, n) {
			return
		}
		fname = v.Type().Field(i).Name
		f = v.Field(i)
	default:
		fmt.Fprintf(ip.env.Stdout, "cannot enter <%v>: expecting array, slice, string or struct\n", typeOf(v))
		return
	}
	var t r.Type
	if f != Nil && f != None {
		t = r.TypeOf(f.Interface()) // concrete type
	}

	switch dereferenceValue(f).Kind() { // dereference pointers on-the-fly
	case r.Array, r.Slice, r.String, r.Struct:
		ip.names = append(ip.names, fname)
		ip.vs = append(ip.vs, f)
		ip.ts = append(ip.ts, t)
		ip.Show()
	default:
		ip.env.showVar(fname, f, t)
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

func (ip *Inspector) validRange(i, n int) bool {
	if i < 0 || i >= n {
		fmt.Fprintf(ip.env.Stdout, "%s contains %d elements, cannot inspect element %d\n",
			strings.Join(ip.names, "."), n, i)
		return false
	}
	return true
}
