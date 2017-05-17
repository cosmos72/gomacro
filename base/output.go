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
	"strings"

	. "github.com/cosmos72/gomacro/ast2"
	mt "github.com/cosmos72/gomacro/token"
)

type Stringer struct {
	Fileset    *mt.FileSet
	Pos        token.Pos
	Line       int
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

func (st *Stringer) Copy(other *Stringer) {
	st.Fileset = other.Fileset
	st.Pos = other.Pos
	st.Line = other.Line
}

func (err RuntimeError) Error() string {
	args := err.args
	var prefix string
	if st := err.st; st != nil {
		args = st.toPrintables(args)
		prefix = st.Position().String()
	}
	msg := fmt.Sprintf(err.format, args...)
	if prefix != "" && prefix != "-" {
		msg = fmt.Sprintf("%s: %s", prefix, msg)
	}
	return msg
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

var warnExtraValues = 5

func (o *Output) WarnExtraValues(extraValues []r.Value) {
	if warnExtraValues > 0 {
		o.Warnf("expression returned %d values, using only the first one: %v",
			len(extraValues), extraValues)
		warnExtraValues--
		if warnExtraValues == 0 {
			o.Warnf("suppressing further similar warnings")
		}
	}
}

func Debugf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	fmt.Printf("// debug: %s\n", str)
}

func (o *Output) Debugf(format string, args ...interface{}) {
	str := o.Sprintf(format, args...)
	fmt.Fprintf(o.Stdout, "// debug: %s\n", str)
}

func (st *Stringer) IncLine(src string) {
	st.Line += strings.Count(src, "\n")
}

func (st *Stringer) IncLineBytes(src []byte) {
	st.Line += bytes.Count(src, []byte("\n"))
}

func (st *Stringer) Position() token.Position {
	if st == nil || st.Fileset == nil {
		return token.Position{}
	}
	return st.Fileset.Position(st.Pos)
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

var typeOfReflectValue = r.TypeOf(r.Value{})

func (st *Stringer) FprintValue(opts Options, out io.Writer, v r.Value) {
	var vi interface{}
	var vt r.Type
	if v == None {
		fmt.Fprint(out, "// no value\n")
		return
	} else if v == Nil {
		vi = nil
		vt = nil
	} else {
		// print the actual type, but unwrap values inside reflect.Value
		vt = v.Type()
		for v.Type() == typeOfReflectValue {
			v = v.Interface().(r.Value)
		}
		vi = v.Interface()
	}
	vi = st.toPrintable(vi)
	if vi == nil && vt == nil {
		if opts&OptShowEvalType != 0 {
			fmt.Fprint(out, "<nil>\n")
		}
		return
	}
	var typestr string
	if opts&OptShowEvalType != 0 {
		typestr = fmt.Sprintf(" // %v", st.toPrintable(vt))
	}

	if vkind := v.Kind(); KindToType(vkind) == v.Type() {
		switch vkind {
		case r.Uint, r.Uint8, r.Uint32, r.Uint64, r.Uintptr:
			fmt.Fprintf(out, "%d%s\n", vi, typestr)
			return
		case r.String:
			fmt.Fprintf(out, "%#v%s\n", vi, typestr)
			return
		}
	}
	// recompute v, because vi = st.toPrintable(vi) may have extracted a non-struct from a struct
	v = r.ValueOf(vi)
	if v.Kind() == r.Struct {
		st.fprintStruct(out, v, typestr)
	} else {
		fmt.Fprintf(out, "%v%s\n", vi, typestr)
	}
}

func (st *Stringer) fprintStruct(out io.Writer, v r.Value, typestr string) {
	buf := bytes.Buffer{}
	n := v.NumField()
	t := v.Type()
	ch := '{'
	for i := 0; i < n; i++ {
		fmt.Fprintf(&buf, "%c%s:%v", ch, t.Field(i).Name, v.Field(i))
		ch = ' '
	}
	fmt.Fprintf(&buf, "}%s\n", typestr)
	buf.WriteTo(out)
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
	rets := make([]interface{}, len(values))
	for i, vi := range values {
		rets[i] = st.toPrintable(vi)
	}
	return rets
}

func (st *Stringer) toPrintable(value interface{}) (ret interface{}) {
	if value == nil {
		return nil
	}
	defer func() {
		if rec := recover(); rec != nil {
			ret = fmt.Sprintf("error pretty-printing %v", value, r.TypeOf(value))
		}
	}()
	switch value := value.(type) {
	case r.Value:
		return st.valueToPrintable(value)
	case AstWithNode:
		return st.nodeToPrintable(value.Node())
	case Ast:
		return st.toPrintable(value.Interface())
	case ast.Node:
		return st.nodeToPrintable(value)
	case r.Type:
		return st.typeToPrintable(value)
	case fmt.Stringer:
		return value.String()
	case fmt.GoStringer:
		return value.GoString()
	}
	v := r.ValueOf(value)
	k := v.Kind()
	if k == r.Array || k == r.Slice {
		n := v.Len()
		values := make([]interface{}, n)
		converted := false
		for i := 0; i < n; i++ {
			vi := v.Index(i)
			if vi == Nil {
				values[i] = nil
			} else if !vi.CanInterface() {
				values[i] = vi
			} else {
				valuei := vi.Interface()
				values[i] = st.toPrintable(valuei)
				converted = converted || !vi.Type().Comparable() || valuei != values[i]
			}
		}
		// return []interface{} only if we actually converted some element
		if converted {
			return values
		} else {
			return value
		}
	}
	return value
}

func (st *Stringer) valueToPrintable(value r.Value) interface{} {
	if value == None {
		return "/*no value*/"
	} else if value == Nil {
		return nil
	} else if value.CanInterface() {
		return st.toPrintable(value.Interface())
	} else {
		return value
	}
}

func (st *Stringer) typeToPrintable(t r.Type) interface{} {
	if t == nil {
		return "nil" // because fmt.Printf("%v", nil) prints <nil> i.e adds extra <>
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
	var fset *mt.FileSet
	if st != nil {
		fset = st.Fileset
	}
	if fset == nil {
		fset = mt.NewFileSet()
	}
	var buf bytes.Buffer
	err := config.Fprint(&buf, &fset.FileSet, node)
	if err != nil {
		return err
	}
	return buf.String()
}
