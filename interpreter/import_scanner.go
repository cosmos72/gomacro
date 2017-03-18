/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * import_scanner.go
 *
 *  Created on Mar 06, 2017
 *      Author Massimiliano Ghilardi
 */

package interpreter

import (
	"bytes"
	"fmt"
	"go/types"
	r "reflect"
	"sort"
	"strings"
	// "time"
)

type typeVisitor func(name string, t types.Type) bool

// implemented by *types.Pointer, *types.Array, *types.Slice, *types.Chan
type typeWithElem interface {
	Elem() types.Type
}

var depth int = 0

func (o *output) Trace(msg ...interface{}) {
	const dots = ". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . "
	const n = len(dots)
	i := 2 * depth
	for i > n {
		fmt.Fprint(o.Stdout, dots)
		i -= n
	}
	// i <= n
	fmt.Fprint(o.Stdout, dots[0:i])
	fmt.Fprintln(o.Stdout, msg...)
}

func trace(o *output, caller string, name string, x interface{}) *output {
	o.Trace(caller, "(", name, x)
	depth++
	return o
}

func un(o *output) {
	depth--
	o.Trace(")")
}

func (o *output) traverseType(name string, in types.Type, visitor typeVisitor) {
	for {
		// defer un(trace(o, "traverseType", name, r.TypeOf(in)))

		if !visitor(name, in) {
			return
		}
		switch t := in.(type) {
		case *types.Basic:
			break
		case *types.Named:
			u := t.Underlying()
			if in != u {
				name = t.Obj().Name()
				in = u
				continue
			}
		case *types.Signature:
			if recv := t.Recv(); recv != nil {
				u := recv.Type()
				// the receiver is often the interface containing this signature...
				// avoid infinite recursion!
				if in != u {
					if _, ok := u.(*types.Interface); !ok {
						o.traverseType(recv.Name(), u, visitor)
					}
				}
			}
			tuples := []*types.Tuple{t.Params(), t.Results()}
			for _, tuple := range tuples {
				n := tuple.Len()
				for i := 0; i < n; i++ {
					v := tuple.At(i)
					o.traverseType(v.Name(), v.Type(), visitor)
				}
			}
		case *types.Interface:
			n := t.NumMethods()
			for i := 0; i < n; i++ {
				method := t.Method(i)
				o.traverseType(method.Name(), method.Type(), visitor)
			}
		case *types.Struct:
			n := t.NumFields()
			for i := 0; i < n; i++ {
				field := t.Field(i)
				o.traverseType(field.Name(), field.Type(), visitor)
			}
		case *types.Map:
			o.traverseType("", t.Key(), visitor)
			name = ""
			in = t.Elem()
			continue
		case typeWithElem: // *types.Pointer, *types.Array, *types.Slice, *types.Chan
			name = ""
			in = t.Elem()
			continue
		default:
			o.warnf("traverseType: unimplemented %#v <%v>", t, r.TypeOf(t))
		}
		break
	}
}

type importExtractor struct {
	imports map[string]bool
	seen    map[types.Type]bool
	o       *output
}

func (ie *importExtractor) visitPackage(pkg *types.Package) {
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		t := extractTypeObjectInterface(obj)
		if t == nil {
			continue
		}
		ie.o.traverseType("", t, ie.visitType)
	}
}

func (ie *importExtractor) visitType(name string, t types.Type) bool {
	if ie.seen[t] {
		return false
	}
	switch t := t.(type) {
	case *types.Named:
		if obj := t.Obj(); obj != nil {
			if pkg := obj.Pkg(); pkg != nil {
				ie.imports[pkg.Path()] = true
			}
		}
		// no need to visit the definition of a named type
		return false
	default:
		return true
	}
}

func extractTypeObjectInterface(obj types.Object) *types.Interface {
	if obj == nil || !obj.Exported() {
		return nil
	}
	switch obj.(type) {
	case *types.TypeName:
		u := obj.Type().Underlying()
		if u, ok := u.(*types.Interface); ok {
			return u
		}
	}
	return nil
}

// we need to collect only the imports that actually appear in package's interfaces methods
// because Go rejects programs with unused imports
func (o *output) collectPackageImports(pkg *types.Package) []string {
	ie := importExtractor{
		// we always need to import the package itself
		imports: map[string]bool{pkg.Path(): true},
		o:       o,
	}

	ie.visitPackage(pkg)

	strings := make([]string, len(ie.imports))
	i := 0
	for str := range ie.imports {
		strings[i] = str
		i++
	}
	sort.Strings(strings)
	return strings
}

type writeTypeOpts int

const (
	writeMethodsAsFields writeTypeOpts = 1 << iota
	writeForceParamNames
	writeIncludeParamTypes
)

func writeInterfaceProxy(out *bytes.Buffer, pkgPath string, pkgSuffix string, name string, t *types.Interface) {
	fmt.Fprintf(out, "\n// --------------- proxy for %s.%s ---------------\ntype %s%s struct {", pkgPath, name, name, pkgSuffix)
	writeInterfaceMethods(out, pkgSuffix, name, t, writeMethodsAsFields)
	out.WriteString("\n}\n")
	writeInterfaceMethods(out, pkgSuffix, name, t, writeForceParamNames)
}

func writeInterfaceMethods(out *bytes.Buffer, pkgSuffix string, name string, t *types.Interface, opts writeTypeOpts) {
	n := t.NumMethods()
	for i := 0; i < n; i++ {
		writeInterfaceMethod(out, pkgSuffix, name, t.Method(i), opts)
	}
}

func writeInterfaceMethod(out *bytes.Buffer, pkgSuffix string, interfaceName string, method *types.Func, opts writeTypeOpts) {
	if !method.Exported() {
		return
	}
	sig, ok := method.Type().(*types.Signature)
	if !ok {
		return
	}
	if opts&writeMethodsAsFields != 0 {
		fmt.Fprintf(out, "\n\t%s_\tfunc", method.Name())
	} else {
		fmt.Fprintf(out, "func (Obj %s%s) %s", interfaceName, pkgSuffix, method.Name())
	}
	params := sig.Params()
	results := sig.Results()
	writeTypeTupleIn(out, params, opts)
	out.WriteString(" ")
	writeTypeTupleOut(out, results)
	if opts&writeMethodsAsFields != 0 {
		return
	}
	out.WriteString(" {\n\t")
	if results != nil && results.Len() > 0 {
		out.WriteString("return ")
	}
	fmt.Fprintf(out, "Obj.%s_(", method.Name())
	writeTypeTuple(out, params, writeForceParamNames)
	out.WriteString(")\n}\n")
}

func writeTypeTupleIn(out *bytes.Buffer, tuple *types.Tuple, opts writeTypeOpts) {
	out.WriteString("(")
	writeTypeTuple(out, tuple, opts|writeIncludeParamTypes)
	out.WriteString(")")
}

func writeTypeTupleOut(out *bytes.Buffer, tuple *types.Tuple) {
	if tuple == nil || tuple.Len() == 0 {
		return
	}
	ret0 := tuple.At(0)
	if tuple.Len() > 1 || len(ret0.Name()) > 0 {
		out.WriteString("(")
		writeTypeTuple(out, tuple, writeIncludeParamTypes)
		out.WriteString(")")
	} else {
		types.WriteType(out, ret0.Type(), packageNameQualifier)
	}
}

func writeTypeTuple(out *bytes.Buffer, tuple *types.Tuple, opts writeTypeOpts) {
	n := tuple.Len()
	for i := 0; i < n; i++ {
		if i != 0 {
			out.WriteString(", ")
		}
		writeTypeVar(out, tuple.At(i), i, opts)
	}
}

func writeTypeVar(out *bytes.Buffer, v *types.Var, index int, opts writeTypeOpts) {
	name := v.Name()
	if len(name) == 0 && opts&writeForceParamNames != 0 {
		name = fmt.Sprintf("unnamed%d", index)
	}
	out.WriteString(name)
	if opts&writeIncludeParamTypes != 0 {
		if len(name) != 0 {
			out.WriteString(" ")
		}
		types.WriteType(out, v.Type(), packageNameQualifier)
	}
}

func packageNameQualifier(pkg *types.Package) string {
	path := pkg.Path()
	return path[1+strings.LastIndexByte(path, '/'):]
}
