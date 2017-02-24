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
 * output.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"io"
	r "reflect"
	"sort"
)

type RuntimeError struct {
	*Interpreter
	Format string
	Args   []interface{}
}

func (err RuntimeError) Error() string {
	return fmt.Sprintf(err.Format, err.toPrintables(err.Args)...)
}

func Errore(err error) interface{} {
	panic(err)
}

func (p *Interpreter) Errorf(format string, args ...interface{}) (r.Value, []r.Value) {
	panic(RuntimeError{p, format, args})
}

func (p *Interpreter) PackErrorf(format string, args ...interface{}) []r.Value {
	panic(RuntimeError{p, format, args})
}

func (p *Interpreter) Warnf(format string, args ...interface{}) {
	str := p.Sprintf(format, args...)
	fmt.Fprintf(p.Stderr, "warning: %s\n", str)
}

func (p *Interpreter) Debugf(format string, args ...interface{}) {
	str := p.Sprintf(format, args...)
	fmt.Fprintf(p.Stdout, "// debug: %s\n", str)
}

func (p *Interpreter) FprintValues(out io.Writer, values ...r.Value) {
	if len(values) == 0 {
		fmt.Fprint(out, "// no value\n")
		return
	}
	for _, v := range values {
		p.FprintValue(out, v)
	}
}

func (p *Interpreter) FprintValue(out io.Writer, v r.Value) {
	var vi interface{}
	var vt r.Type
	if v == None {
		fmt.Fprint(out, "// no value\n")
		return
	} else if v == Nil {
		vi = nil
		vt = nil
	} else {
		vi = v.Interface()
		vt = v.Type()
	}
	vi = p.toPrintable(vi)
	if vi == nil && vt == nil {
		fmt.Fprint(out, "<nil>\n")
		return
	}
	switch vi := vi.(type) {
	case uint, uint8, uint32, uint64, uintptr:
		fmt.Fprintf(out, "%d <%v>\n", vi, vt)
	default:
		if vt == typeOfString {
			fmt.Fprintf(out, "%#v <%v>\n", vi, vt)
		} else {
			fmt.Fprintf(out, "%v <%v>\n", vi, vt)
		}
	}
}

func (p *Interpreter) Fprintf(out io.Writer, format string, values ...interface{}) (n int, err error) {
	values = p.toPrintables(values)
	return fmt.Fprintf(out, format, values...)
}

func (p *Interpreter) Sprintf(format string, values ...interface{}) string {
	values = p.toPrintables(values)
	return fmt.Sprintf(format, values...)
}

func (p *Interpreter) toString(separator string, values ...interface{}) string {
	if len(values) == 0 {
		return ""
	}
	values = p.toPrintables(values)
	var buf bytes.Buffer
	for i, value := range values {
		if i != 0 {
			buf.WriteString(separator)
		}
		fmt.Fprint(&buf, value)
	}
	return buf.String()
}

func (p *Interpreter) toPrintables(values []interface{}) []interface{} {
	for i, vi := range values {
		values[i] = p.toPrintable(vi)
	}
	return values
}

func (p *Interpreter) toPrintable(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	if v, ok := value.(r.Value); ok {
		return p.valueToPrintable(v)
	}
	v := r.ValueOf(value)
	k := v.Kind()
	if k == r.Array || k == r.Slice {
		n := v.Len()
		values := make([]interface{}, n)
		for i := 0; i < n; i++ {
			values[i] = p.toPrintable(v.Index(i))
		}
		return values
	}
	if node, ok := value.(ast.Node); ok {
		return p.nodeToPrintable(node)
	}
	return value
}

func (p *Interpreter) valueToPrintable(value r.Value) interface{} {
	if value == None {
		return "/*no value*/"
	} else if value == Nil {
		return nil
	} else {
		return p.toPrintable(value.Interface())
	}
}

func (p *Interpreter) nodeToPrintable(node ast.Node) interface{} {
	var buf bytes.Buffer
	err := format.Node(&buf, p.Fileset, node)
	if err != nil {
		return err
	}
	return buf.String()
}

func (p *Interpreter) showHelp(out io.Writer) {
	fmt.Fprint(out, `// interpreter commands:
:env    show available functions, variables and constants
:help   show this help
:quit   quit the interpreter
`)
}

func (env *Env) showEnv(out io.Writer) {
	binds := env.binds
	keys := make([]string, len(binds))
	i := 0
	for k := range binds {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	spaces15 := "               "
	for _, k := range keys {
		n := len(k) & 15
		fmt.Fprintf(out, "%s%s = ", k, spaces15[n:])
		env.FprintValue(out, binds[k])
	}
}
