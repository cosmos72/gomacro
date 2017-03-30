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
	"go/printer"
	"go/token"
	"io"
	r "reflect"
	"sort"

	. "github.com/cosmos72/gomacro/ast2"
)

type stringer struct {
	Fileset    *token.FileSet
	NamedTypes map[r.Type]string
}

type output struct {
	stringer
	Stdout io.Writer
	Stderr io.Writer
}

type runtimeError struct {
	st     *stringer
	format string
	args   []interface{}
}

func (err runtimeError) Error() string {
	return fmt.Sprintf(err.format, err.st.toPrintables(err.args)...)
}

func error_(err error) interface{} {
	panic(err)
}

func errorf(format string, args ...interface{}) {
	panic(runtimeError{nil, format, args})
}

func (st *stringer) errorf(format string, args ...interface{}) (r.Value, []r.Value) {
	panic(runtimeError{st, format, args})
}

func (st *stringer) packErrorf(format string, args ...interface{}) []r.Value {
	panic(runtimeError{st, format, args})
}

func (o *output) warnf(format string, args ...interface{}) {
	str := o.sprintf(format, args...)
	fmt.Fprintf(o.Stderr, "warning: %s\n", str)
}

func (o *output) debugf(format string, args ...interface{}) {
	str := o.sprintf(format, args...)
	fmt.Fprintf(o.Stdout, "// debug: %s\n", str)
}

func (env *Env) showStack() {
	frames := env.CallStack.Frames
	n := len(frames)
	for i := 1; i < n; i++ {
		frame := &frames[i]
		name := ""
		if frame.FuncEnv != nil {
			name = frame.FuncEnv.Name
		}
		if frame.panicking {
			env.debugf("%d:\t     %v, runningDefers = %v, panic = %v", i, name, frame.runningDefers, frame.panick)
		} else {
			env.debugf("%d:\t     %v, runningDefers = %v, panic is nil", i, name, frame.runningDefers)
		}
	}
}

func (st *stringer) fprintValues(out io.Writer, values ...r.Value) {
	if len(values) == 0 {
		fmt.Fprint(out, "// no value\n")
		return
	}
	for _, v := range values {
		st.fprintValue(out, v)
	}
}

func (st *stringer) fprintValue(out io.Writer, v r.Value) {
	var vi, vt interface{}
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
	vi = st.toPrintable(vi)
	if vi == nil && vt == nil {
		fmt.Fprint(out, "<nil>\n")
		return
	}
	vt = st.toPrintable(vt)
	switch vi := vi.(type) {
	case uint, uint8, uint32, uint64, uintptr:
		fmt.Fprintf(out, "%d <%v>\n", vi, vt)
	default:
		if vt == typeOfString {
			fmt.Fprintf(out, "%q <%v>\n", vi, vt)
		} else {
			fmt.Fprintf(out, "%v <%v>\n", vi, vt)
		}
	}
}

func (st *stringer) fprintf(out io.Writer, format string, values ...interface{}) (n int, err error) {
	values = st.toPrintables(values)
	return fmt.Fprintf(out, format, values...)
}

func (st *stringer) sprintf(format string, values ...interface{}) string {
	values = st.toPrintables(values)
	return fmt.Sprintf(format, values...)
}

func (st *stringer) toString(separator string, values ...interface{}) string {
	if len(values) == 0 {
		return ""
	}
	values = st.toPrintables(values)
	var buf bytes.Buffer
	for i, value := range values {
		if i != 0 {
			buf.WriteString(separator)
		}
		fmt.Fprint(&buf, value)
	}
	return buf.String()
}

func (st *stringer) toPrintables(values []interface{}) []interface{} {
	for i, vi := range values {
		values[i] = st.toPrintable(vi)
	}
	return values
}

func (st *stringer) toPrintable(value interface{}) (ret interface{}) {
	if value == nil {
		return nil
	}
	defer func() {
		if rec := recover(); rec != nil {
			ret = fmt.Sprintf("error pretty-printing %#v <%v>", value, r.TypeOf(value))
		}
	}()
	if v, ok := value.(r.Value); ok {
		return st.valueToPrintable(v)
	}
	v := r.ValueOf(value)
	k := v.Kind()
	if k == r.Array || k == r.Slice {
		n := v.Len()
		values := make([]interface{}, n)
		for i := 0; i < n; i++ {
			values[i] = st.toPrintable(v.Index(i))
		}
		return values
	}
	switch v := value.(type) {
	case AstWithNode:
		return st.nodeToPrintable(v.Node())
	case Ast:
		return st.toPrintable(v.Interface())
	case ast.Node:
		return st.nodeToPrintable(v)
	case r.Type:
		return st.typeToPrintable(v)
	}
	return value
}

func (st *stringer) valueToPrintable(value r.Value) interface{} {
	if value == None {
		return "/*no value*/"
	} else if value == Nil {
		return nil
	} else {
		return st.toPrintable(value.Interface())
	}
}

func (st *stringer) typeToPrintable(t r.Type) interface{} {
	if t == nil {
		return nil
	}
	if st != nil {
		if name, ok := st.NamedTypes[t]; ok {
			return name
		}
	}
	return t
}

var config = printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}

func (st *stringer) nodeToPrintable(node ast.Node) interface{} {
	if node == nil {
		return nil
	}
	var fset *token.FileSet
	if st != nil {
		fset = st.Fileset
	}
	if fset == nil {
		fset = token.NewFileSet()
	}
	var buf bytes.Buffer
	err := config.Fprint(&buf, fset, node)
	if err != nil {
		return err
	}
	return buf.String()
}

func (st *stringer) showHelp(out io.Writer) {
	fmt.Fprint(out, `// interpreter commands:
:env [name]     show available functions, variables and constants
                in current package, or from imported package "name"
:help           print this help
:inspect EXPR   inspect expression interactively
:options [OPTS] show or toggle interpreter options
:quit           quit the interpreter
:write [FILE]   write collected declarations and/or statements to standard output or to file
                use :o Declarations and/or :o Statements to start collecting them
`)
}

func (env *Env) showPackage(packageName string) {
	out := env.Stdout
	e := env
	path := env.Path
	pkg := &env.Package
	if len(packageName) != 0 {
		bind := env.evalIdentifier(&ast.Ident{Name: packageName})
		if bind == None || bind == Nil {
			env.warnf("not an imported package: %q", packageName)
			return
		}
		switch val := bind.Interface().(type) {
		case *PackageRef:
			e = nil
			pkg = &val.Package
			path = packageName
		default:
			env.warnf("not an imported package: %q = %v <%v>", packageName, val, typeOf(bind))
			return
		}
	}
	spaces15 := "               "
Loop:
	binds := pkg.Binds
	if len(binds) > 0 {
		fmt.Fprintf(out, "// ----- %s binds -----\n", path)

		keys := make([]string, len(binds))
		i := 0
		for k := range binds {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			n := len(k) & 15
			fmt.Fprintf(out, "%s%s = ", k, spaces15[n:])
			bind := binds[k]
			if bind != Nil {
				switch bind := bind.Interface().(type) {
				case *Env:
					fmt.Fprintf(out, "%p <%v>\n", bind, r.TypeOf(bind))
					continue
				}
			}
			env.fprintValue(out, bind)
		}
		fmt.Fprintln(out)
	}
	types := pkg.Types
	if len(types) > 0 {
		fmt.Fprintf(out, "// ----- %s types -----\n", path)

		keys := make([]string, len(types))
		i := 0
		for k := range types {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			n := len(k) & 15
			t := types[k]
			fmt.Fprintf(out, "%s%s %v <%v>\n", k, spaces15[n:], t.Kind(), t)
		}
		fmt.Fprintln(out)
	}
	if e != nil {
		if e = e.Outer; e != nil {
			path = e.Path
			pkg = &e.Package
			goto Loop
		}
	}
}
