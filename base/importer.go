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

type ImportMode int

const (
	ImSharedLib ImportMode = iota
	ImBuiltin
	ImInception
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

// LookupPackage returns a package if already present in cache
func (g *Globals) LookupPackage(name, path string) (ref *PackageRef, complete bool) {
	pkg, found := imports.Packages[path]
	if !found {
		return nil, false
	}
	_, complete = pkg.GoPkg.(*types.Package)
	return &PackageRef{Package: pkg, Name: name, Path: path}, complete
}

func (g *Globals) ImportPackage(name, path string) *PackageRef {
	ref, complete := g.LookupPackage(name, path)
	if complete {
		return ref
	}
	gpkg, err := g.Importer.Import(path) // loads names and types, not the values!
	if err != nil {
		if ref != nil {
			g.Warnf("error loading package %q metadata, continuing without (wrapper methods for embedded fields may be inaccurate): %v", path, err)
			return ref
		} else {
			g.Errorf("error loading package %q metadata, maybe you need to download (go get), compile (go build) and install (go install) it? %v", path, err)
			return nil
		}
	}
	if ref != nil {
		// we have all the information. cache gpkg for future use.
		ref.Package.GoPkg = gpkg
		imports.Packages[path] = ref.Package
		return ref
	}
	var mode ImportMode
	switch name {
	case "_b":
		mode = ImBuiltin
	case "_i":
		mode = ImInception
	}
	file := g.createImportFile(path, gpkg, mode)
	if mode != ImSharedLib {
		return nil
	}

	ref = &PackageRef{Name: name, Path: path}

	if len(file) == 0 {
		// empty package. still cache it for future use.
		ref.Package.GoPkg = gpkg
		imports.Packages[path] = ref.Package
		return ref
	}
	soname := g.compilePlugin(file, g.Stdout, g.Stderr)
	ifun := g.loadPlugin(soname, "Exports")
	fun := ifun.(func() (map[string]r.Value, map[string]r.Type, map[string]r.Type))
	binds, types, proxies := fun()

	// done. cache pkg and gpkg for future use.
	ref.Package = imports.Package{
		Binds:   binds,
		Types:   types,
		Proxies: proxies,
		GoPkg:   gpkg,
	}
	imports.Packages[path] = ref.Package
	return ref
}

func (g *Globals) createImportFile(path string, pkg *types.Package, mode ImportMode) string {
	buf := bytes.Buffer{}
	isEmpty := g.writeImportFile(&buf, path, pkg, mode)
	if isEmpty {
		g.Warnf("package %q exports zero constants, functions, types and variables", path)
		return ""
	}

	file := computeImportFilename(path, mode)
	err := ioutil.WriteFile(file, buf.Bytes(), os.FileMode(0666))
	if err != nil {
		g.Errorf("error writing file %q: %v", file, err)
	}
	if mode != ImSharedLib {
		g.Warnf("created file %q, recompile gomacro to use it", file)
	} else {
		g.Debugf("created file %q...", file)
	}
	return file
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

const gomacro_dir = "github.com/cosmos72/gomacro"

func computeImportFilename(path string, mode ImportMode) string {
	srcdir := getGoSrcPath()

	switch mode {
	case ImBuiltin:
		return fmt.Sprintf("%s/%s/imports/%s.go", srcdir, gomacro_dir, sanitizeIdentifier(path))
	case ImInception:
		return fmt.Sprintf("%s/%s/x_package.go", srcdir, path)
	}

	file := path[1+strings.LastIndexByte(path, '/'):]
	file = fmt.Sprintf("%s/gomacro_imports/%s/%s.go", srcdir, path, file)
	dir := file[0 : 1+strings.LastIndexByte(file, '/')]
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		Errorf("error creating directory %q: %v", dir, err)
	}
	return file
}

func (ir *Globals) writeImportFile(out *bytes.Buffer, path string, gpkg *types.Package, mode ImportMode) (isEmpty bool) {
	scope := gpkg.Scope()
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

	pkg := path[1+strings.LastIndexByte(path, '/'):]
	pkg = sanitizeIdentifier(pkg)

	pkgSuffix := ""
	if mode != ImSharedLib {
		pkgSuffix = fmt.Sprintf("_%s", sanitizeIdentifier(path))
	}

	var alias, pkg_, filepkg string
	switch mode {
	case ImBuiltin:
		alias = "_b "
		filepkg = "imports"
		pkg_ = pkg + "."
	case ImInception:
		alias = "_i "
		filepkg = pkg
	default:
		filepkg = "main"
		pkg_ = pkg + "."
	}

	fmt.Fprintf(out, `// this file was generated by gomacro command: import %s%q
// DO NOT EDIT! Any change will be lost when the file is re-generated

package %s

import (`, alias, path, filepkg)

	var imports, reflect string
	if mode == ImInception {
		fmt.Fprintf(out, "\n\tr \"reflect\"\n\t\"github.com/cosmos72/gomacro/imports\"")
		imports = "imports."
		reflect = "r."
	} else {
		fmt.Fprintf(out, "\n\t. \"reflect\"")
	}
	for _, str := range ir.CollectPackageImports(gpkg, true) {
		if mode != ImInception || str != path {
			fmt.Fprintf(out, "\n\t%q", str)
		}
	}
	fmt.Fprintf(out, "\n)\n")

	if mode == ImSharedLib {
		fmt.Fprint(out, `
func main() {
}

func Exports() (map[string]Value, map[string]Type, map[string]Type) {
	return map[string]Value{`)
	} else {
		fmt.Fprintf(out, `
// reflection: allow interpreted code to import %q
func init() {
	%sPackages[%q] = %sPackage{
	Binds: map[string]%sValue{`, path, imports, path, imports, reflect)
	}

	for _, name := range names {
		if obj := scope.Lookup(name); obj.Exported() {
			switch obj := obj.(type) {
			case *types.Const:
				val := obj.Val()
				var conv1, conv2 string
				if t, ok := obj.Type().(*types.Basic); ok && t.Info()&types.IsUntyped != 0 {
					// untyped constants have arbitrary precision... they may overflow integers
					if val.Kind() == constant.Int {
						str := val.ExactString()
						conv1, conv2 = ir.detectIntKind(path, name, str)
					}
				}
				fmt.Fprintf(out, "\n\t\t%q:\t%sValueOf(%s%s%s%s),", name, reflect, conv1, pkg_, name, conv2)
			case *types.Var:
				fmt.Fprintf(out, "\n\t\t%q:\t%sValueOf(&%s%s).Elem(),", name, reflect, pkg_, name)
			case *types.Func:
				fmt.Fprintf(out, "\n\t\t%q:\t%sValueOf(%s%s),", name, reflect, pkg_, name)
			}
		}
	}

	if mode == ImSharedLib {
		fmt.Fprint(out, "\n\t}, map[string]Type{")
	} else {
		fmt.Fprintf(out, "\n\t},\n\tTypes: map[string]%sType{", reflect)
	}

	for _, name := range names {
		if obj := scope.Lookup(name); obj.Exported() {
			switch obj.(type) {
			case *types.TypeName:
				fmt.Fprintf(out, "\n\t\t%q:\t%sTypeOf((*%s%s)(nil)).Elem(),", name, reflect, pkg_, name)
			}
		}
	}

	if mode == ImSharedLib {
		fmt.Fprint(out, "\n\t}, map[string]Type{")
	} else {
		fmt.Fprintf(out, "\n\t},\n\tProxies: map[string]%sType{", reflect)
	}

	for _, name := range names {
		obj := scope.Lookup(name)
		if t := extractInterface(obj, true); t != nil {
			fmt.Fprintf(out, "\n\t\t%q:\t%sTypeOf((*%s%s)(nil)).Elem(),", name, reflect, name, pkgSuffix)
		}
	}

	if mode == ImSharedLib {
		fmt.Fprint(out, "\n\t}\n}\n")
	} else {
		fmt.Fprint(out, "\n\t} }\n}\n")
	}

	for _, name := range names {
		obj := scope.Lookup(name)
		if t := extractInterface(obj, true); t != nil {
			writeInterfaceProxy(out, gpkg.Path(), pkgSuffix, name, t)
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
