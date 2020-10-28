/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * scanner.go
 *
 *  Created on Mar 06, 2017
 *      Author Massimiliano Ghilardi
 */

package genimport

import (
	"fmt"
	"go/token"
	"go/types"
	r "reflect"
	"sort"

	"github.com/cosmos72/gomacro/base/output"
)

type TypeVisitor func(t types.Type) bool

// implemented by *types.Pointer, *types.Array, *types.Slice, *types.Chan
type typeWithElem interface {
	Elem() types.Type
}

var depth int = 0

func traceargs(o *Output, msg ...interface{}) {
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

func trace(o *Output, caller string, name string, x interface{}) *Output {
	traceargs(o, caller, "(", name, x)
	depth++
	return o
}

func un(o *Output) {
	depth--
	traceargs(o, ")")
}

func traverseType(o *Output, name string, in types.Type, visitor TypeVisitor) {
	for {
		// defer un(trace(o, "traverseType", name, in))

		if !visitor(in) {
			return
		}
		switch t := in.(type) {
		case *types.Basic:
			break
		case *types.Named:
			// we are only interested in types visible in interface method declarations:
			// to use a named type we do not need to import the packages used by it.
			if false {
				u := t.Underlying()
				if in != u {
					name = t.Obj().Name()
					in = u
					continue
				}
			}
		case *types.Signature:
			// we are only interested in types visible in interface method declarations:
			// the receiver is omitted by Go interface{ ... } syntax
			// and thus we do not need to import its package
			if false {
				if recv := t.Recv(); recv != nil {
					u := recv.Type()
					// the receiver is often the interface containing this signature...
					// avoid infinite recursion!
					if in != u {
						if _, ok := u.(*types.Interface); !ok {
							traverseType(o, recv.Name(), u, visitor)
						}
					}
				}
			}
			tuples := []*types.Tuple{t.Params(), t.Results()}
			for _, tuple := range tuples {
				n := tuple.Len()
				for i := 0; i < n; i++ {
					v := tuple.At(i)
					traverseType(o, v.Name(), v.Type(), visitor)
				}
			}
		case *types.Interface:
			n := t.NumMethods()
			for i := 0; i < n; i++ {
				method := t.Method(i)
				traverseType(o, method.Name(), method.Type(), visitor)
			}
		case *types.Struct:
			n := t.NumFields()
			for i := 0; i < n; i++ {
				field := t.Field(i)
				traverseType(o, field.Name(), field.Type(), visitor)
			}
		case *types.Map:
			traverseType(o, "", t.Key(), visitor)
			name = ""
			in = t.Elem()
			continue
		case typeWithElem: // *types.Pointer, *types.Array, *types.Slice, *types.Chan
			name = ""
			in = t.Elem()
			continue
		default:
			o.Warnf("traverseType: unimplemented %#v <%v>", t, r.TypeOf(t))
		}
		break
	}
}

type importExtractor struct {
	imports map[string]bool
	seen    map[types.Type]bool
	o       *Output
}

func (ie *importExtractor) visitPackage(pkg *types.Package, requireAllInterfaceMethodsExported bool) {
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		t := extractInterface(obj, requireAllInterfaceMethodsExported)
		if t != nil {
			traverseType(ie.o, "", t, ie.visitType)
		}
	}
}

func (ie *importExtractor) visitType(t types.Type) bool {
	if ie.seen[t] {
		return false
	}
	switch t := t.(type) {
	case *types.Basic:
		if t.Kind() == types.UnsafePointer {
			ie.imports["unsafe"] = true
			// output.Debugf("will import %q. reason: found type %v", "unsafe", t)
		}
		// no need to visit the definition of a basic type
		return false
	case *types.Named:
		if obj := t.Obj(); obj != nil {
			if pkg := obj.Pkg(); pkg != nil {
				ie.imports[pkg.Path()] = true
				// output.Debugf("will import %q. reason: found type %v", pkg.Path(), t)
			}
		}
		// no need to visit the definition of a named type
		return false
	default:
		return true
	}
}

func extractInterface(obj types.Object, requireAllMethodsAndTypesExported bool) *types.Interface {
	if obj == nil || !obj.Exported() {
		return nil
	}
	switch obj.(type) {
	case *types.TypeName:
		u := obj.Type().Underlying()
		if u, ok := u.(*types.Interface); ok {
			// do not export proxies for empty interfaces:
			// using reflect.Value.Convert() at runtime is enough
			if u.NumMethods() != 0 && (!requireAllMethodsAndTypesExported || interfaceExported(u)) {
				return u
			}
		}
	}
	return nil
}

// return true if Object is nil, or type of Object is exported.
func objectExported(obj types.Object) bool {
	if obj == nil {
		return true
	}
	return typeExported(obj.Type())
}

// return true if tuple's elements types are exported, or tuple is nil
func tupleExported(tuple *types.Tuple) bool {
	if tuple != nil {
		n := tuple.Len()
		for i := 0; i < n; i++ {
			if !objectExported(tuple.At(i)) {
				return false
			}
		}
	}
	return true
}

// return true if Type is exported or nil
func typeExported(typ types.Type) bool {
	for typ != nil {
		switch t := typ.(type) {
		case *types.Array:
			typ = t.Elem()
		case *types.Basic:
			return true
		case *types.Chan:
			typ = t.Elem()
		case *types.Interface:
			return interfaceExported(t)
		case *types.Map:
			if !typeExported(t.Key()) {
				return false
			}
			typ = t.Elem()
		case *types.Named:
			return t.Obj().Pkg() == nil || t.Obj().Exported()
		case *types.Pointer:
			typ = t.Elem()
		case *types.Signature:
			return signatureExported(t)
		case *types.Slice:
			typ = t.Elem()
		case *types.Struct:
			return structExported(t)
		default:
			output.Errorf("unexpected type %v", typ)
		}
	}
	return true
}

// return true if an interface type is exported. This means:
// 1. all its methods are exported,
// 2. and all the argument types and return types of its methods are exported too
func interfaceExported(intf *types.Interface) bool {
	if intf != nil {
		n := intf.NumMethods()
		for i := 0; i < n; i++ {
			m := intf.Method(i)
			if !m.Exported() {
				return false
			}
			if !signatureExported(m.Type().(*types.Signature)) {
				return false
			}
		}
	}
	return true
}

// return true if all the argument types and return types of given function signature are exported
// Note: receiver is ignored because only interface methods can be part of an exported type
func signatureExported(sig *types.Signature) bool {
	return sig == nil ||
		tupleExported(sig.Params()) &&
			tupleExported(sig.Results())
}

// return true if a struct type is exported. This means:
// 1. all its fields are exported,
// 2. and all the types of its fields are exported too
func structExported(s *types.Struct) bool {
	if s != nil {
		n := s.NumFields()
		for i := 0; i < n; i++ {
			f := s.Field(i)
			if !f.Exported() {
				return false
			}
			if !typeExported(f.Type()) {
				return false
			}
		}
	}
	return true
}

// we need to collect only the imports that actually appear in package's interfaces methods
// because Go rejects programs with unused imports.
//
// To avoid naming conflicts when importing two different packages
// that end with the same name, as for example image/draw and golang.org/x/image/draw,
// we rename conflicting packages and return a map[path]renamed
func collectPackageImportsWithRename(o *Output, pkg *types.Package, requireAllInterfaceMethodsExported bool) map[string]string {
	ie := importExtractor{
		// we always need to import the package itself
		imports: map[string]bool{pkg.Path(): true},
		o:       o,
	}
	// output.Debugf("before visitPackage: imports = %v", ie.imports)
	ie.visitPackage(pkg, requireAllInterfaceMethodsExported)
	// output.Debugf("after  visitPackage: imports = %v", ie.imports)

	// for deterministic renaming, use a sorted []string instead of a map[string]bool
	pathlist := getKeys(ie.imports)
	sort.Strings(pathlist)

	nametopath := renamePackages(pathlist)
	pathtoname := transposeKeyValue(nametopath)

	// prevent renaming the package we are scanning
	// unless it's named unsafe - see below
	preventRenamingIfPresent(pkg.Path(), nametopath, pathtoname)

	// also prevent renaming "unsafe" - easier than
	// intercepting type unsafe.Pointer and renaming it
	preventRenamingIfPresent("unsafe", nametopath, pathtoname)

	return pathtoname
}

func preventRenamingIfPresent(path string, nametopath map[string]string, pathtoname map[string]string) {
	name := packageSanitizedName(path)
	name2, ok := pathtoname[path]
	if !ok || name == name2 {
		return
	}
	// some *other* path is associated to name: swap their names
	path2, ok := nametopath[name]
	if ok {
		nametopath[name2] = path2
		pathtoname[path2] = name2
	}
	nametopath[path] = name
	pathtoname[path] = name
	// output.Debugf("swapped import aliases: will import\n\t%s %q\n\t%s %q", name, path, name2, path2)
}

// given a slice []path, return a map[name]path where all paths
// that end in the same name have been assigned unique names
func renamePackages(in []string) map[string]string {
	out := make(map[string]string)
	for _, path := range in {
		name := renamePackage(path, out)
		out[name] = path
	}
	return out
}

// given a package path and a map[name]path, extract the path last name.
// Change it (if needed) to a value that is NOT in map and return it.
func renamePackage(path string, out map[string]string) string {
	name := packageSanitizedName(path)
	if _, exists := out[name]; !exists {
		// output.Debugf("package %q will be imported as %s", path, name)
		return name
	}
	n := len(name)
	for n != 0 && isDigit(name[n-1]) {
		n--
	}
	name = name[:n]
	for i := uint64(0); i < ^uint64(0); i++ {
		namei := fmt.Sprintf("%s%d", name, i)
		if _, exists := out[namei]; !exists {
			// output.Debugf("package %q will be imported as %s because the alias %s is already used", path, namei, name)
			return namei
		}
	}
	output.Errorf("failed to find a non-conflicting rename for package %q", path)
	return "???"
}

func isReservedKeyword(s string) bool {
	return len(s) > 1 && token.Lookup(s) != token.IDENT
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

// given a map k -> v, return an *unsorted* slice of its keys
func getKeys(in map[string]bool) []string {
	keys := make([]string, len(in))
	i := 0
	for key := range in {
		keys[i] = key
		i++
	}
	return keys
}

// given a map k -> v, return a map v -> k
func transposeKeyValue(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[v] = k
	}
	return out
}
