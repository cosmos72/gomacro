/*
 * gomacro - A Go interpreter with Lisp-like macros
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
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/types"
	r "reflect"
	"strconv"
	"strings"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// Import compiles an import statement
func (c *Comp) Import(node ast.Spec) {
	switch node := node.(type) {
	case *ast.ImportSpec:
		str := node.Path.Value
		path, err := strconv.Unquote(str)
		if err != nil {
			c.Errorf("error unescaping import path %q: %v", str, err)
		}
		path = c.sanitizeImportPath(path)
		var name string
		if node.Name != nil {
			name = node.Name.Name
		} else {
			name = path[1+strings.LastIndexByte(path, '/'):]
		}
		c.FileComp().ImportPackage(name, path)
	default:
		c.Errorf("unimplemented import: %v", node)
	}
}

func (g *CompThreadGlobals) sanitizeImportPath(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	l := len(path)
	if path == ".." || l >= 3 && (path[:3] == "../" || path[l-3:] == "/..") || strings.Contains(path, "/../") {
		g.Errorf("invalid import %q: contains \"..\"", path)
	}
	if path == "." || l >= 2 && (path[:2] == "./" || path[l-2:] == "/.") || strings.Contains(path, "/./") {
		g.Errorf("invalid import %q: contains \".\"", path)
	}
	return path
}

// Import imports a package. Usually invoked as Comp.FileComp().ImportPackage(name, path)
// because imports are top-level statements in a source file.
func (cfile *Comp) ImportPackage(name, path string) *Import {
	g := cfile.CompThreadGlobals
	pkgref := g.ImportPackage(name, path)
	if pkgref == nil {
		return nil
	}
	pkg, _ := pkgref.GoPkg.(*types.Package)
	if pkg != nil {
		g.Universe.CachePackage(pkg)
	}

	binds, bindtypes := g.parseImportBinds(pkgref.Binds, pkg)

	imp := Import{
		Binds:     binds,
		BindTypes: bindtypes,
		Types:     g.rtypesToXTypes(pkgref.Types),
		Name:      name,
		Path:      path,
	}

	g.loadProxies(pkgref.Proxies, imp.Types)

	cfile.DeclImport0(name, imp)
	return &imp
}

// declImport0 compiles an import declaration.
// Note: does not loads proxies, use ImportPackage for that
func (cfile *Comp) DeclImport0(name string, imp Import) {
	// treat imported package as a constant,
	// because to compile code we need the declarations it contains:
	// importing them at runtime would be too late.
	t := cfile.TypeOfImport()
	bind := cfile.AddBind(name, ConstBind, t)
	bind.Value = imp // cfile.Binds[] is a map[string]*Bind => changes to *Bind propagate to the map
}

func (g *CompThreadGlobals) parseImportBinds(binds map[string]r.Value, pkg *types.Package) (map[string]r.Value, map[string]xr.Type) {
	v := g.Universe
	retbinds := make(map[string]r.Value)
	rettypes := make(map[string]xr.Type)
	var scope *types.Scope
	if pkg != nil {
		scope = pkg.Scope()
	}
	for name, bind := range binds {
		if scope != nil {
			value, typ, ok := g.parseImportConst(name, bind, scope)
			if ok {
				retbinds[name] = value
				rettypes[name] = typ
				continue
			}
		}
		// Universe.FromReflectType uses cached *types.Package if possible
		retbinds[name] = bind
		rettypes[name] = v.FromReflectType(bind.Type())
	}
	return retbinds, rettypes
}

func (g *CompThreadGlobals) parseImportConst(name string, bind r.Value, scope *types.Scope) (r.Value, xr.Type, bool) {
	if c, ok := scope.Lookup(name).(*types.Const); ok {
		switch gtype := c.Type().(type) {
		case *types.Basic:
			gkind := gtype.Kind()
			kind := xr.ToReflectKind(gkind)
			if kind == r.Invalid {
				// untyped nil
				return Nil, nil, true
			}
			lit := UntypedLit{Kind: kind, Obj: c.Val(), Universe: g.Universe}
			if xr.IsGoUntypedKind(gkind) {
				// untyped constant
				return r.ValueOf(lit), g.TypeOfUntypedLit(), true
			}
			t := g.Universe.BasicTypes[kind]
			if t == nil {
				t = lit.DefaultType()
			}
			return r.ValueOf(lit.ConstTo(t)), t, true
		}
	}
	return Nil, nil, false
}

func (g *CompThreadGlobals) rtypesToXTypes(rtypes map[string]r.Type) map[string]xr.Type {
	v := g.Universe
	xtypes := make(map[string]xr.Type)
	for name, rtype := range rtypes {
		// Universe.FromReflectType uses cached *types.Package if possible
		xtypes[name] = v.FromReflectType(rtype)
	}
	return xtypes
}

// loadProxies adds to thread-global maps the proxies found in import
func (g *CompThreadGlobals) loadProxies(proxies map[string]r.Type, xtypes map[string]xr.Type) {
	for name, proxy := range proxies {
		xtype := xtypes[name]
		if xtype == nil {
			g.Warnf("import %q: type not found for proxy <%v>", proxy.PkgPath(), proxy)
			continue
		}
		if xtype.Kind() != r.Interface {
			g.Warnf("import %q: type for proxy <%v> is not an interface: %v", proxy.PkgPath(), proxy, xtype)
			continue
		}
		rtype := xtype.ReflectType()
		g.interf2proxy[rtype] = proxy
		g.proxy2interf[proxy] = xtype
	}
}
