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
 * import.go
 *
 *  Created on Feb 27, 2017
 *      Author Massimiliano Ghilardi
 */

package interpreter

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/constant"
	"go/importer"
	"go/types"
	"io/ioutil"
	"os"
	r "reflect"
	"strconv"
	"strings"

	"github.com/cosmos72/gomacro/imports"
)

type Importer struct {
	from   types.ImporterFrom
	compat types.Importer
	srcDir string
	mode   types.ImportMode
}

func DefaultImporter() Importer {
	imp := Importer{}
	compat := importer.Default()
	if from, ok := imp.compat.(types.ImporterFrom); ok {
		imp.from = from
	} else {
		imp.compat = compat
	}
	return imp
}

func (imp *Importer) Import(path string) (*types.Package, error) {
	if imp.from != nil {
		return imp.from.ImportFrom(path, imp.srcDir, imp.mode)
	} else {
		return imp.compat.Import(path)
	}
}

func (imp *Importer) ImportFrom(path string, srcDir string, mode types.ImportMode) (*types.Package, error) {
	if imp.from != nil {
		return imp.from.ImportFrom(path, srcDir, mode)
	} else {
		return imp.compat.Import(path)
	}
}

// eval a single import
func (env *Env) evalImport(node ast.Spec) (r.Value, []r.Value) {
	switch node := node.(type) {
	case *ast.ImportSpec:
		path := unescapeString(node.Path.Value)
		var name string
		if node.Name != nil {
			name = node.Name.Name
		} else {
			name = path[1+strings.LastIndexByte(path, '/'):]
		}
		newEnv := env.ImportPackage(name, path)
		if newEnv != nil {
			fileEnv := env.FileEnv()
			newEnv.Outer = fileEnv.TopEnv()

			value := r.ValueOf(newEnv)
			fileEnv.defineConst(name, TypeOf(value), value)
		}
		return r.ValueOf(path), nil
	default:
		return env.Errorf("unimplemented import: %v", node)
	}
}

func (ir *Interpreter) ImportPackage(name, path string) *Env {
	if binds, ok := imports.Binds[path]; ok {
		if types, ok := imports.Types[path]; ok {
			return &Env{Binds: binds, Types: types, Name: name, Path: path}
		}
	}
	pkg, err := ir.Importer.Import(path) // loads names and types, not the values!
	if err != nil {
		ir.Errorf("error loading package %q metadata, maybe you need to download (go get), compile (go build) and install (go install) it? %v", path, err)
		return nil
	}
	internal := false
	filename := ir.createImportFile(path, pkg, internal)
	if internal {
		return nil
	}
	if len(filename) == 0 {
		// empty package
		return &Env{Binds: map[string]r.Value{}, Types: map[string]r.Type{}, Proxies: map[string]r.Type{}, Name: name, Path: path}
	}

	soname := ir.compilePlugin(filename, ir.Stdout, ir.Stderr)
	ifun := loadPlugin(soname, "Package")
	fun := ifun.(func() (map[string]r.Value, map[string]r.Type, map[string]r.Type))
	binds, types, proxies := fun()
	return &Env{Binds: binds, Types: types, Proxies: proxies, Name: name, Path: path}
}

func (ir *Interpreter) createImportFile(path string, pkg *types.Package, internal bool) string {
	buf := bytes.Buffer{}
	isEmpty := ir.writeImportFile(&buf, path, pkg, internal)
	if isEmpty {
		ir.Warnf("package %q exports zero constants, functions, types and variables", path)
		return ""
	}

	filename := computeImportFilename(path, internal)
	err := ioutil.WriteFile(filename, buf.Bytes(), os.FileMode(0666))
	if err != nil {
		Errorf("error writing file %q: %v", filename, err)
	}
	if internal {
		ir.Warnf("created file %q, recompile gomacro to use it", filename)
	} else {
		ir.Debugf("created file %q...", filename)
	}
	return filename
}

func sanitizeIdentifier(str string) string {
	return sanitizeIdentifier2(str, '_')
}

func sanitizeIdentifier2(str string, replacement rune) string {
	runes := []rune(str)
	for i, ch := range runes {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || ch == '_' ||
			(i != 0 && ch >= '0' && ch <= '9') {
			continue
		}
		runes[i] = replacement
	}
	return string(runes)
}

func computeImportFilename(path string, internal bool) string {
	if internal {
		return fmt.Sprintf("imports/%s.go", sanitizeIdentifier(path))
	}

	srcdirname := getGoSrcPath()
	fmt.Printf("computeImportFilename(): srcdirname = %q", srcdirname)

	filename := path[1+strings.LastIndexByte(path, '/'):]
	filename = fmt.Sprintf("%s/gomacro_imports/%s/%s.go", srcdirname, path, filename)
	dirname := filename[0 : 1+strings.LastIndexByte(filename, '/')]
	err := os.MkdirAll(dirname, 0700)
	if err != nil {
		Errorf("error creating directory %q: %v", dirname, err)
	}
	return filename
}

func (ir *Interpreter) writeImportFile(out *bytes.Buffer, path string, pkg *types.Package, internal bool) (isEmpty bool) {
	scope := pkg.Scope()
	names := scope.Names()

	isEmpty = true
	for _, name := range names {
		if obj := scope.Lookup(name); obj.Exported() {
			switch obj.(type) {
			case *types.Const, *types.Var, *types.Func, *types.TypeName:
				isEmpty = false
				break
			}
		}
	}
	if isEmpty {
		return isEmpty
	}

	pkgName := path[1+strings.LastIndexByte(path, '/'):]
	pkgName = sanitizeIdentifier(pkgName)

	pkgSuffix := ""
	if internal {
		pkgSuffix = fmt.Sprintf("_%s", sanitizeIdentifier(path))
	}

	var thisPkgName = "main"
	if internal {
		thisPkgName = "imports"
	}

	fmt.Fprintf(out, `// this file was generated by gomacro command: import %q
// DO NOT EDIT! Any change will be lost when the file is re-generated

package %s

import (
	. "reflect"`, path, thisPkgName)

	for _, str := range ir.collectPackageImports(pkg) {
		fmt.Fprintf(out, "\n\t%q", str)
	}

	if internal {
		fmt.Fprintf(out, `
)

func init() {
	Binds[%q] = map[string]Value{`, path)
	} else {
		fmt.Fprint(out, `
)

func main() {
}

func Package() (map[string]Value, map[string]Type, map[string]Type) {
	return map[string]Value{`)
	}

	for _, name := range names {
		if obj := scope.Lookup(name); obj.Exported() {
			switch obj := obj.(type) {
			case *types.Const:
				val := obj.Val()
				var prefix, suffix string
				if val.Kind() == constant.Int {
					str := val.ExactString()
					_, err := strconv.ParseInt(str, 0, 0)
					if err != nil {
						_, err := strconv.ParseUint(str, 0, 0)
						if err == nil {
							prefix = "uint64("
							suffix = ")"
						} else {
							ir.Warnf("package %q: integer constant %s = %s overflows both int64 and uint64, expect compile errors", path, name, str)
						}
					}
				}
				fmt.Fprintf(out, "\n\t\t%q:\tValueOf(%s%s.%s%s),", name, prefix, pkgName, name, suffix)
			case *types.Var:
				fmt.Fprintf(out, "\n\t\t%q:\tValueOf(&%s.%s).Elem(),", name, pkgName, name)
			case *types.Func:
				fmt.Fprintf(out, "\n\t\t%q:\tValueOf(%s.%s),", name, pkgName, name)
			}
		}
	}

	if internal {
		fmt.Fprintf(out, "\n\t}\n\tTypes[%q] = map[string]Type{", path)
	} else {
		fmt.Fprint(out, "\n\t}, map[string]Type{")
	}

	for _, name := range names {
		if obj := scope.Lookup(name); obj.Exported() {
			switch obj.(type) {
			case *types.TypeName:
				fmt.Fprintf(out, "\n\t\t%q:\tTypeOf((*%s.%s)(nil)).Elem(),", name, pkgName, name)
			}
		}
	}

	if internal {
		fmt.Fprintf(out, "\n\t}\n\tProxies[%q] = map[string]Type{", path)
	} else {
		fmt.Fprint(out, "\n\t}, map[string]Type{")
	}

	for _, name := range names {
		obj := scope.Lookup(name)
		if t := extractTypeObjectInterface(obj); t != nil {
			fmt.Fprintf(out, "\n\t\t%q:\tTypeOf((*%s%s)(nil)).Elem(),", name, name, pkgSuffix)
		}
	}

	fmt.Fprint(out, "\n\t}\n}\n")

	for _, name := range names {
		obj := scope.Lookup(name)
		if t := extractTypeObjectInterface(obj); t != nil {
			writeInterfaceProxy(out, pkg.Path(), pkgSuffix, name, t)
		}
	}
	return isEmpty
}
