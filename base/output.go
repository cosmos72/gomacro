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

package base

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"io"
	r "reflect"

	. "github.com/cosmos72/gomacro/ast2"
)

type Stringer struct {
	Fileset    *token.FileSet
	NamedTypes map[r.Type]string
}

type Output struct {
	Stringer
	Stdout io.Writer
	Stderr io.Writer
}

type RuntimeError struct {
	st     *Stringer
	format string
	args   []interface{}
}

func (err RuntimeError) Error() string {
	return fmt.Sprintf(err.format, err.st.toPrintables(err.args)...)
}

func Error(err error) interface{} {
	panic(err)
}

func (o *Output) Error(err error) interface{} {
	panic(err)
}

func Errorf(format string, args ...interface{}) {
	panic(RuntimeError{nil, format, args})
}

func (st *Stringer) Errorf(format string, args ...interface{}) (r.Value, []r.Value) {
	panic(RuntimeError{st, format, args})
}

func Warnf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	fmt.Printf("// warning: %s\n", str)
}

func (o *Output) Warnf(format string, args ...interface{}) {
	str := o.Sprintf(format, args...)
	fmt.Fprintf(o.Stderr, "// warning: %s\n", str)
}

func (o *Output) WarnExtraValues(extraValues []r.Value) {
	o.Warnf("expression returned %d values, using only the first one: %v",
		len(extraValues), extraValues)
}

func Debugf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	fmt.Printf("// debug: %s\n", str)
}

func (o *Output) Debugf(format string, args ...interface{}) {
	str := o.Sprintf(format, args...)
	fmt.Fprintf(o.Stdout, "// debug: %s\n", str)
}

func (st *Stringer) FprintValues(opts Options, out io.Writer, values ...r.Value) {
	if len(values) == 0 {
		fmt.Fprint(out, "// no value\n")
		return
	}
	for _, v := range values {
		st.FprintValue(opts, out, v)
	}
}

func (st *Stringer) FprintValue(opts Options, out io.Writer, v r.Value) {
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
		if opts&OptShowEvalType != 0 {
			fmt.Fprint(out, "<nil>\n")
		}
		return
	}
	switch vi := vi.(type) {
	case uint, uint8, uint32, uint64, uintptr:
		if opts&OptShowEvalType != 0 {
			vt = st.toPrintable(vt)
			fmt.Fprintf(out, "%d <%v>\n", vi, vt)
		} else {
			fmt.Fprintf(out, "%d\n", vi)
		}
	default:
		if opts&OptShowEvalType != 0 {
			vt = st.toPrintable(vt)
			if vt == TypeOfString {
				fmt.Fprintf(out, "%q <%v>\n", vi, vt)
			} else {
				fmt.Fprintf(out, "%v <%v>\n", vi, vt)
			}
		} else {
			if vt == TypeOfString {
				fmt.Fprintf(out, "%q\n", vi)
			} else {
				fmt.Fprintf(out, "%v\n", vi)
			}
		}
	}
}

func (st *Stringer) Fprintf(out io.Writer, format string, values ...interface{}) (n int, err error) {
	values = st.toPrintables(values)
	return fmt.Fprintf(out, format, values...)
}

func (st *Stringer) Sprintf(format string, values ...interface{}) string {
	values = st.toPrintables(values)
	return fmt.Sprintf(format, values...)
}

func (st *Stringer) ToString(separator string, values ...interface{}) string {
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

func (st *Stringer) toPrintables(values []interface{}) []interface{} {
	for i, vi := range values {
		values[i] = st.toPrintable(vi)
	}
	return values
}

func (st *Stringer) toPrintable(value interface{}) (ret interface{}) {
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

func (st *Stringer) valueToPrintable(value r.Value) interface{} {
	if value == None {
		return "/*no value*/"
	} else if value == Nil {
		return nil
	} else {
		return st.toPrintable(value.Interface())
	}
}

func (st *Stringer) typeToPrintable(t r.Type) interface{} {
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

func (st *Stringer) nodeToPrintable(node ast.Node) interface{} {
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
