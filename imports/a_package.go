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
 * a_package.go
 *
 *  Created on: Feb 28, 2017
 *      Author: Massimiliano Ghilardi
 */

package imports

import (
	fmt "fmt"
	. "reflect"

	syscall "github.com/cosmos72/gomacro/imports/syscall"
	thirdparty "github.com/cosmos72/gomacro/imports/thirdparty"
	"github.com/cosmos72/gomacro/imports/util"
)

type PackageUnderlying = struct { // unnamed
	Name    string
	Binds   map[string]Value
	Types   map[string]Type
	Proxies map[string]Type
	// Untypeds contains a string representation of untyped constants,
	// stored without loss of precision
	Untypeds map[string]string
	// Wrappers is the list of wrapper methods for named types.
	// Stored explicitly because reflect package cannot distinguish
	// between explicit methods and wrapper methods for embedded fields
	Wrappers map[string][]string
}

type PackageName string // package default name, or package alias

type Package PackageUnderlying // named, can have methods

type PackageMap map[string]Package // named, can have methods

var Packages = make(PackageMap)

func (name PackageName) String() string {
	return string(name)
}

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/imports"
func init() {
	Packages["github.com/cosmos72/gomacro/imports"] = Package{
		Binds: map[string]Value{
			"Packages": ValueOf(&Packages).Elem(),
		},
		Types: map[string]Type{
			"Package":           TypeOf((*Package)(nil)).Elem(),
			"PackageMap":        TypeOf((*PackageMap)(nil)).Elem(),
			"PackageName":       TypeOf((*PackageName)(nil)).Elem(),
			"PackageUnderlying": TypeOf((*PackageUnderlying)(nil)).Elem(),
		},
	}
	Packages.Merge(syscall.Packages)
	Packages.Merge(thirdparty.Packages)
}

func (pkgs PackageMap) Merge(srcs map[string]PackageUnderlying) {
	// exploit the fact that maps are actually handles
	for path, src := range srcs {
		pkgs.MergePackage(path, src)
	}
}

func (pkgs PackageMap) MergePackage(path string, src PackageUnderlying) {
	pkg := Package(src)
	pkg.Validate(path)

	curr, ok := pkgs[path]
	if ok {
		curr.Merge(src)
	} else {
		pkg.LazyInit(path)
		// exploit the fact that maps are actually handles
		pkgs[path] = pkg
	}
}

/**
 * return the default name to bind when importing Package
 *
 * https://golang.org/ref/spec#Package_clause states:
 * If the PackageName is omitted, it defaults to the identifier
 * specified in the package clause of the imported package
 *
 * So use that if known, otherwise extrapolate it from package path
 */
func (pkg *Package) DefaultName(path string) PackageName {
	if len(pkg.Name) == 0 {
		pkg.Name = util.TailIdentifier(util.FileName(path))
	}
	return PackageName(pkg.Name)
}

func (pkg *Package) LazyInit(path string) {
	pkg.DefaultName(path)
	if pkg.Binds == nil {
		pkg.Binds = make(map[string]Value)
	}
	if pkg.Types == nil {
		pkg.Types = make(map[string]Type)
	}
	if pkg.Proxies == nil {
		pkg.Proxies = make(map[string]Type)
	}
	if pkg.Untypeds == nil {
		pkg.Untypeds = make(map[string]string)
	}
	if pkg.Wrappers == nil {
		pkg.Wrappers = make(map[string][]string)
	}
}

func (dst *Package) Merge(src PackageUnderlying) {
	if len(src.Name) != 0 {
		dst.Name = src.Name
	}
	for k, v := range src.Binds {
		dst.Binds[k] = v
	}
	for k, v := range src.Types {
		dst.Types[k] = v
		// when overwriting a type, also overwrite its proxy and wrapper list
		delete(dst.Proxies, k)
		delete(dst.Wrappers, k)
	}
	for k, v := range src.Proxies {
		dst.Proxies[k] = v
	}
	for k, v := range src.Untypeds {
		dst.Untypeds[k] = v
	}
	for k, v := range src.Wrappers {
		dst.Wrappers[k] = v
	}
}

func (pkg *Package) Validate(path string) {
	for name, typ := range pkg.Types {
		if typ.Kind() == Interface {
			validateProxy(path, name, typ, pkg.Proxies[name])
		}
	}
	for name := range pkg.Proxies {
		if pkg.Types[name] == nil {
			errorf("error loading package %q: interface %s is not defined by the package, cannot define a proxy for it",
				path, name)
		}
	}
}

func validateProxy(path string, name string, typ Type, proxy Type) {
	if proxy == nil {
		// no proxy for interface typ. This is allowed:
		// interfaces with unexported methods cannot be implemented by other packages
		// => there's no way to create a proxy implementing such interface
		// (short of "cheating" and embedding the original interface, which is not useful in this case)
		return
	}
	// fmt.Printf("package %q:\tvalidating proxy for interface %s\n", path, name)

	if proxy.Kind() != Struct {
		errorf("error loading package %q: proxy for interface %s is invalid: expecting a Struct, found a %v",
			path, name, proxy.Kind())
	}
	proxy = PtrTo(proxy)
	if !proxy.Implements(typ) {
		errorf("error loading package %q: proxy for interface %s is invalid: type <%v> does not implement the interface %s",
			path, name, proxy)
	}
	typMethodN := typ.NumMethod()
	proxyMethodN := proxy.NumMethod()
	if typMethodN != proxyMethodN {
		errorf("error loading package %q: proxy for interface %s is invalid: type <%v> has %d methods, expecting %d",
			path, name, proxy, proxyMethodN, typMethodN)
	}
	if proxy.Elem().NumField() != proxyMethodN+1 {
		errorf("error loading package %q: proxy for interface %s is invalid: type <%v> has %d fields, expecting %d i.e. 1 + number of methods",
			path, name, proxy, proxy.Elem().NumField(), proxyMethodN+1)
	}
	validateProxyField0(path, name, typ, proxy)
	for i := 0; i < typMethodN; i++ {
		validateProxyFieldAndMethod(path, name, typ, proxy, i)
	}
}

var rTypeOfInterface = TypeOf((*interface{})(nil)).Elem()

func validateProxyField0(path string, name string, typ Type, proxy Type) {
	field := proxy.Elem().Field(0)
	fieldName := qname(field.PkgPath, field.Name)
	if fieldName != "Object" {
		errorf("error loading package %q: proxy for interface %s is invalid: type <%v> has field[0] name %q, expecting name %q",
			path, name, proxy.Elem(), fieldName, "Object")
	}
	if field.Type != rTypeOfInterface {
		errorf("error loading package %q: proxy for interface %s is invalid: type <%v> has field[0] type <%v>, expecting type <%v>",
			path, name, proxy.Elem(), field.Type, rTypeOfInterface)
	}
}

func validateProxyFieldAndMethod(path string, name string, typ Type, proxy Type, i int) {
	typMethod := typ.Method(i)
	proxyMethod := proxy.Method(i)
	typMethodName := qname(typMethod.PkgPath, typMethod.Name)
	proxyMethodName := qname(proxyMethod.PkgPath, proxyMethod.Name)
	if typMethodName != proxyMethodName {
		errorf("error loading package %q: proxy for interface %s is invalid: type <%v> has method[%d] name %q, expecting name %q",
			path, name, proxy, i, proxyMethodName, typMethodName)
	}
	expectedType := addFuncFirstParam(typMethod.Type, proxy)
	if proxyMethod.Type != expectedType {
		errorf("error loading package %q: proxy for interface %s is invalid: type <%v> has method[%d] type <%v>, expecting type <%v>",
			path, name, proxy, i, proxyMethod.Type, expectedType)
	}

	field := proxy.Elem().Field(i + 1) // skip field 0 "Obiect"
	fieldName := qname(field.PkgPath, field.Name)
	if typMethodName+"_" != fieldName {
		errorf("error loading package %q: proxy for interface %s is invalid: type <%v> has field[%d] name %q, expecting name %q",
			path, name, proxy.Elem(), i+1, fieldName, typMethodName+"_")
	}
	expectedType = addFuncFirstParam(typMethod.Type, rTypeOfInterface)
	if field.Type != expectedType {
		errorf("error loading package %q: proxy for interface %s is invalid: type <%v> has field[%d] type <%v>, expecting type <%v>",
			path, name, proxy.Elem(), i+1, field.Type, expectedType)
	}
}

// return package-qualified name of a method
func qname(pkgPath string, name string) string {
	if len(pkgPath) == 0 {
		return name
	}
	return pkgPath + "." + name
}

// return a function type with additional first input parameter of given type
func addFuncFirstParam(ftype Type, in0 Type) Type {
	n := ftype.NumIn() + 1
	in := make([]Type, n)
	in[0] = in0
	for i := 1; i < n; i++ {
		in[i] = ftype.In(i - 1)
	}
	n = ftype.NumOut()
	out := make([]Type, n)
	for i := 0; i < n; i++ {
		out[i] = ftype.Out(i)
	}
	return FuncOf(in, out, ftype.IsVariadic())
}

func errorf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}
