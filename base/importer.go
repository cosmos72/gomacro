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
 * importer.go
 *
 *  Created on Feb 27, 2017
 *      Author Massimiliano Ghilardi
 */

package base

import (
	"bytes"
	"fmt"
	"go/constant"
	"go/importer"
	"go/types"
	"io/ioutil"
	"math"
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

func DefaultImporter() *Importer {
	imp := &Importer{}
	compat := importer.Default()
	if from, ok := compat.(types.ImporterFrom); ok {
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

func (g *Globals) ImportPackage(name, path string) *PackageRef {
	if pkg, ok := imports.Packages[path]; ok {
		return &PackageRef{Package: pkg, Name: name, Path: path}
	}
	pkg, err := g.Importer.Import(path) // loads names and types, not the values!
	if err != nil {
		g.Errorf("error loading package %q metadata, maybe you need to download (go get), compile (go build) and install (go install) it? %v", path, err)
		return nil
	}
	internal := name == "__"
	filename := g.createImportFile(path, pkg, internal)
	if internal {
		return nil
	}
	if len(filename) == 0 {
		// empty package
		return &PackageRef{Name: name, Path: path}
	}

	soname := g.compilePlugin(filename, g.Stdout, g.Stderr)
	ifun := g.loadPlugin(soname, "Exports")
	fun := ifun.(func() (map[string]r.Value, map[string]r.Type, map[string]r.Type))
	binds, types, proxies := fun()
	return &PackageRef{
		Package: imports.Package{Binds: binds, Types: types, Proxies: proxies},
		Name:    name, Path: path}
}

func (g *Globals) createImportFile(path string, pkg *types.Package, internal bool) string {
	buf := bytes.Buffer{}
	isEmpty := g.writeImportFile(&buf, path, pkg, internal)
	if isEmpty {
		g.Warnf("package %q exports zero constants, functions, types and variables", path)
		return ""
	}

	filename := computeImportFilename(path, internal)
	err := ioutil.WriteFile(filename, buf.Bytes(), os.FileMode(0666))
	if err != nil {
		g.Errorf("error writing file %q: %v", filename, err)
	}
	if internal {
		g.Warnf("created file %q, recompile gomacro to use it", filename)
	} else {
		g.Debugf("created file %q...", filename)
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

	filename := path[1+strings.LastIndexByte(path, '/'):]
	filename = fmt.Sprintf("%s/gomacro_imports/%s/%s.go", srcdirname, path, filename)
	dirname := filename[0 : 1+strings.LastIndexByte(filename, '/')]
	err := os.MkdirAll(dirname, 0700)
	if err != nil {
		Errorf("error creating directory %q: %v", dirname, err)
	}
	return filename
}

func (ir *Globals) writeImportFile(out *bytes.Buffer, path string, pkg *types.Package, internal bool) (isEmpty bool) {
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

	for _, str := range ir.CollectPackageImports(pkg, true) {
		fmt.Fprintf(out, "\n\t%q", str)
	}

	if internal {
		fmt.Fprintf(out, `
)

func init() {
	Packages[%q] = Package{
	Binds: map[string]Value{`, path)
	} else {
		fmt.Fprint(out, `
)

func main() {
}

func Exports() (map[string]Value, map[string]Type, map[string]Type) {
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
					prefix, suffix = ir.detectIntKind(path, name, str)
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
		fmt.Fprint(out, "\n\t},\n\tTypes: map[string]Type{")
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
		fmt.Fprint(out, "\n\t},\n\tProxies: map[string]Type{")
	} else {
		fmt.Fprint(out, "\n\t}, map[string]Type{")
	}

	for _, name := range names {
		obj := scope.Lookup(name)
		if t := extractInterface(obj, true); t != nil {
			fmt.Fprintf(out, "\n\t\t%q:\tTypeOf((*%s%s)(nil)).Elem(),", name, name, pkgSuffix)
		}
	}

	if internal {
		fmt.Fprint(out, "\n\t} }\n}\n")
	} else {
		fmt.Fprint(out, "\n\t}\n}\n")
	}

	for _, name := range names {
		obj := scope.Lookup(name)
		if t := extractInterface(obj, true); t != nil {
			writeInterfaceProxy(out, pkg.Path(), pkgSuffix, name, t)
		}
	}
	return isEmpty
}

func (ir *Globals) detectIntKind(path, name, str string) (string, string) {
	i, err := strconv.ParseInt(str, 0, 64)
	if err == nil {
		if i == int64(int32(i)) {
			// constant fits int32. We can use the default (i.e. int)
			// on both 32-bit and 64-bit platforms
			return "", ""
		} else if i == int64(uint32(i)) {
			// constant fits uint32
			return "uint32(", ")"
		} else {
			return "int64(", ")"
		}
	}
	_, err = strconv.ParseUint(str, 0, 64)
	if err == nil {
		return "uint64(", ")"
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		// nothing fits... leave the default
		return "", ""
	} else {
		prefix := "float64"
		f = math.Abs(f)
		if f == float64(float32(f)) && f <= math.MaxFloat32 && f >= math.SmallestNonzeroFloat32 {
			// float32 loses no precision vs. float64
			prefix = "float32"
		}
		ir.Warnf("package %q: integer constant %s = %s overflows both int64 and uint64, converting to %s", path, name, str, prefix)
		return prefix + "(", ")"
	}
}
