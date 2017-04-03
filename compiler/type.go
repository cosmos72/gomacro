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
 * type.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package compiler

import (
	"fmt"
	"go/ast"
	r "reflect"
)

// DeclType compiles a type declaration.
func (c *Comp) DeclType(node ast.Spec) X {
	switch node := node.(type) {
	case *ast.TypeSpec:
		name := node.Name.Name
		t := c.CompileType(node.Type)
		c.DeclType0(name, t)

		vt := r.ValueOf(&t).Elem() // return a reflect.Type, not the concrete type
		return func(*Env) (r.Value, []r.Value) {
			return vt, nil
		}

	default:
		return c.Errorf("Compile: unexpected type declaration, expecting <*ast.TypeSpec>, found: %v <%v>", node, r.TypeOf(node))
	}
}

// DeclType0 executes a type declaration
// in Go, types are computed only at compile time - no need for a runtime *Env
func (c *Comp) DeclType0(name string, t r.Type) r.Type {
	if name == "_" {
		// never define bindings for "_"
		return t
	}
	if _, ok := c.Types[name]; ok {
		c.Warnf("redefined type: %v", name)
	} else if c.Types == nil {
		c.Types = make(map[string]r.Type)
	}
	c.Types[name] = t
	if _, ok := c.NamedTypes[t]; !ok {
		c.NamedTypes[t] = NamedType{Name: name, Path: c.Path}
	}
	return t
}

// CompileType compiles a type expression.
func (c *Comp) CompileType(node ast.Expr) r.Type {
	t, _ := c.compileType2(node, false)
	return t
}

// compileTypeOrNil compiles a type expression. as a special case used by type switch, compiles *ast.Ident{Name:"nil"} to nil
func (c *Comp) compileTypeOrNil(node ast.Expr) r.Type {
	for {
		switch expr := node.(type) {
		case *ast.ParenExpr:
			node = expr.X
			continue
		case *ast.Ident:
			if expr.Name == "nil" {
				_, bind, ok := c.tryResolve(expr.Name)
				if ok && bind.Type == nil {
					return nil
				}
			}
		}
		break
	}
	t, _ := c.compileType2(node, false)
	return t
}

// compileType2 compiles a type expression.
// if allowEllipsis is true, it supports the special case &ast.Ellipsis{/*expression*/}
// that represents ellipsis in the last argument of a function declaration.
// The second return value is true both in the case above, and for array types whose length is [...]
func (c *Comp) compileType2(node ast.Expr, allowEllipsis bool) (t r.Type, ellipsis bool) {
	stars := 0
	for {
		switch expr := node.(type) {
		case *ast.StarExpr:
			stars++
			node = expr.X
			continue
		case *ast.ParenExpr:
			node = expr.X
			continue
		case *ast.Ellipsis:
			if allowEllipsis {
				node = expr.Elt
				ellipsis = true
				continue
			}
		}
		break
	}

	switch node := node.(type) {
	case *ast.ArrayType: // also for slices
		var ellipsis2 bool
		t, ellipsis2 = c.TypeArray(node)
		if !ellipsis {
			ellipsis = ellipsis2
		}
	case *ast.ChanType:
		t = c.CompileType(node.Value)
		dir := r.BothDir
		if node.Dir == ast.SEND {
			dir = r.SendDir
		} else if node.Dir == ast.RECV {
			dir = r.RecvDir
		}
		t = r.ChanOf(dir, t)
	case *ast.FuncType:
		t, _, _ = c.TypeFunction(node)
	case *ast.Ident:
		t = c.TypeIdent(node.Name)
	case *ast.InterfaceType:
		t = c.TypeInterface(node)
	case *ast.MapType:
		kt := c.CompileType(node.Key)
		vt := c.CompileType(node.Value)
		t = r.MapOf(kt, vt)
	case *ast.SelectorExpr:
		if _, ok := node.X.(*ast.Ident); ok {
			/*
				pkgv := c.Identifier(pkgIdent)
				if pkg, ok := pkgv.Interface().(*PackageRef); ok {
					name := node.Sel.Name
					if t, ok = pkg.Types[name]; !ok {
						c.Errorf("not a type: %v <%v>", node, r.TypeOf(node))
					}
				} else {
					c.Errorf("not a package: %v = %v <%v>", pkgIdent, pkgv, typeOf(pkgv))
				}
			*/
			c.Errorf("types from another package are not implemented: %v <%v>", node, r.TypeOf(node))
		} else {
			c.Errorf("invalid qualified type, expecting packageName.identifier, found: %v <%v>", node, r.TypeOf(node))
		}
	case *ast.StructType:
		// c.Debugf("evalType() struct declaration: %v <%v>", node, r.TypeOf(node))
		types, names := c.TypeFields(node.Fields)
		// c.Debugf("evalType() struct names and types: %v %v", types, names)
		fields := makeStructFields(c.File().Path, names, types)
		// c.Debugf("evalType() struct fields: %#v", fields)
		t = r.StructOf(fields)
	case nil:
		// type can be omitted in many case - then we must perform type inference
		break
	default:
		// TODO which types are still missing?
		c.Errorf("unimplemented type: %v <%v>", node, r.TypeOf(node))
	}
	for i := 0; i < stars; i++ {
		t = r.PtrTo(t)
	}
	if allowEllipsis && ellipsis {
		t = r.SliceOf(t)
	}
	return t, ellipsis
}

func (c *Comp) TypeArray(node *ast.ArrayType) (t r.Type, ellipsis bool) {
	t = c.CompileType(node.Elt)
	n := node.Len
	switch n := n.(type) {
	case *ast.Ellipsis:
		t = r.SliceOf(t)
		ellipsis = true
	case nil:
		t = r.SliceOf(t)
	default:
		init := c.Expr(n)
		if !init.Const() {
			c.Errorf("array length is not a constant: %v", node)
			return
		}
		value := init.EvalConst()
		count := int(r.ValueOf(value).Int())
		t = r.ArrayOf(count, t)
	}
	return t, ellipsis
}

func (c *Comp) TypeFunction(node *ast.FuncType) (t r.Type, argNames []string, resultNames []string) {
	tFunc, _, argNames, resultNames := c.TypeFunctionOrMethod(nil, node)
	return tFunc, argNames, resultNames
}

func (c *Comp) TypeFunctionOrMethod(recv *ast.Field, node *ast.FuncType) (tFunc r.Type, tFuncOrMethod r.Type, argNames []string, resultNames []string) {
	argTypes, argNames, variadic := c.typeFieldOrParamList(node.Params, true)
	resultTypes, resultNames := c.TypeFields(node.Results)
	tFunc = r.FuncOf(argTypes, resultTypes, variadic)

	if recv != nil {
		recvTypes, recvNames, _ := c.typeFieldsOrParams([]*ast.Field{recv}, false)
		argTypes = append(recvTypes, argTypes...)
		argNames = append(recvNames, argNames...)
		tFuncOrMethod = r.FuncOf(argTypes, resultTypes, variadic)
	} else {
		tFuncOrMethod = tFunc
	}
	return tFunc, tFuncOrMethod, argNames, resultNames
}

func (c *Comp) TypeFields(fields *ast.FieldList) (types []r.Type, names []string) {
	types, names, _ = c.typeFieldOrParamList(fields, false)
	return types, names
}

func (c *Comp) typeFieldOrParamList(fields *ast.FieldList, allowEllipsis bool) (types []r.Type, names []string, ellipsis bool) {
	var list []*ast.Field
	if fields != nil {
		list = fields.List
	}
	return c.typeFieldsOrParams(list, allowEllipsis)
}

func (c *Comp) typeFieldsOrParams(list []*ast.Field, allowEllipsis bool) (types []r.Type, names []string, ellipsis bool) {
	types = make([]r.Type, 0)
	names = zeroStrings
	n := len(list)
	if n == 0 {
		return types, names, ellipsis
	}
	var t r.Type
	for i, f := range list {
		t, ellipsis = c.compileType2(f.Type, i == n-1)
		if len(f.Names) == 0 {
			types = append(types, t)
			names = append(names, "_")
			// c.Debugf("evalTypeFields() %v -> %v", f.Type, t)
		} else {
			for _, ident := range f.Names {
				types = append(types, t)
				names = append(names, ident.Name)
				// Debugf("evalTypeFields() %v %v -> %v", ident.Name, f.Type, t)
			}
		}
	}
	return types, names, ellipsis
}

func (c *Comp) TypeIdent(name string) r.Type {
	for co := c; co != nil; co = co.Outer {
		if t, ok := co.Types[name]; ok {
			return t
		}
	}
	c.Errorf("undefined identifier: %v", name)
	return nil
}

func makeStructFields(pkgPath string, names []string, types []r.Type) []r.StructField {
	// pkgIdentifier := sanitizeIdentifier(pkgPath)
	fields := make([]r.StructField, len(names))
	for i, name := range names {
		fields[i] = r.StructField{
			Name:      toExportedName(name), // Go 1.8 reflect.StructOf() supports *only* exported fields
			Type:      types[i],
			Tag:       "",
			Anonymous: false,
		}
	}
	return fields
}

func toExportedName(name string) string {
	if len(name) == 0 {
		return name
	}
	ch := name[0]
	if ch >= 'a' && ch <= 'z' {
		ch -= 'a' - 'A'
	} else if ch == '_' {
		ch = 'X'
	} else {
		return name
	}
	return fmt.Sprintf("%c%s", ch, name[1:])
}
