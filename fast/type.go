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

package fast

import (
	"fmt"
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

// DeclType compiles a type declaration.
func (c *Comp) DeclType(node ast.Spec) {
	switch node := node.(type) {
	case *ast.TypeSpec:
		name := node.Name.Name
		t := c.Type(node.Type)
		c.DeclType0(name, t)
	default:
		c.Errorf("Compile: unexpected type declaration, expecting <*ast.TypeSpec>, found: %v <%v>", node, r.TypeOf(node))
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
	if c.NamedTypes == nil {
		c.NamedTypes = make(map[r.Type]NamedType)
	}
	c.NamedTypes[t] = NamedType{Name: name, Path: c.Path}
	return t
}

// Type compiles a type expression.
func (c *Comp) Type(node ast.Expr) r.Type {
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
				sym := c.TryResolve(expr.Name)
				if sym != nil && sym.Type == nil {
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
	if node != nil {
		c.Pos = node.Pos()
	}

	switch node := node.(type) {
	case *ast.ArrayType: // also for slices
		var ellipsis2 bool
		t, ellipsis2 = c.TypeArray(node)
		if !ellipsis {
			ellipsis = ellipsis2
		}
	case *ast.ChanType:
		t = c.Type(node.Value)
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
		kt := c.Type(node.Key)
		vt := c.Type(node.Value)
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
	t = c.Type(node.Elt)
	n := node.Len
	switch n := n.(type) {
	case *ast.Ellipsis:
		t = r.SliceOf(t)
		ellipsis = true
	case nil:
		t = r.SliceOf(t)
	default:
		// as stated by https://golang.org/ref/spec#Array_types
		// "The length is part of the array's type; it must evaluate to a non-negative constant
		// representable by a value of type int. "
		var count int
		init := c.Expr(n)
		if !init.Const() {
			c.Errorf("array length is not a constant: %v", node)
			return
		} else if init.Untyped() {
			count = init.ConstTo(TypeOfInt).(int)
		} else {
			count = convertLiteralCheckOverflow(init.Value, TypeOfInt).(int)
		}
		if count < 0 {
			c.Errorf("array length [%v] is negative: %v", count, node)
		}
		t = r.ArrayOf(count, t)
	}
	return t, ellipsis
}

func (c *Comp) TypeFunction(node *ast.FuncType) (t r.Type, paramNames []string, resultNames []string) {
	tFunc, _, paramNames, resultNames := c.TypeFunctionOrMethod(nil, node)
	return tFunc, paramNames, resultNames
}

func (c *Comp) TypeFunctionOrMethod(recv *ast.Field, node *ast.FuncType) (tFunc r.Type, tFuncOrMethod r.Type, paramNames []string, resultNames []string) {
	paramTypes, paramNames, variadic := c.typeFieldOrParamList(node.Params, true)
	resultTypes, resultNames := c.TypeFields(node.Results)
	tFunc = r.FuncOf(paramTypes, resultTypes, variadic)

	if recv != nil {
		recvTypes, recvNames, _ := c.typeFieldsOrParams([]*ast.Field{recv}, false)
		paramTypes = append(recvTypes, paramTypes...)
		paramNames = append(recvNames, paramNames...)
		tFuncOrMethod = r.FuncOf(paramTypes, resultTypes, variadic)
	} else {
		tFuncOrMethod = tFunc
	}
	return tFunc, tFuncOrMethod, paramNames, resultNames
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
	names = ZeroStrings
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

// TypeAssert2 compiles a multi-valued type assertion
func (c *Comp) TypeAssert2(node *ast.TypeAssertExpr) *Expr {
	val := c.Expr1(node.X)
	tin := val.Type
	tout := c.Type(node.Type)
	kout := tout.Kind()
	if tin == nil || tin.Kind() != r.Interface {
		c.Errorf("invalid type assertion: %v (non-interface type <%v> on left)", node, tin)
		return nil
	}
	if tout.Kind() != r.Interface && !tout.Implements(tin) {
		c.Errorf("impossible type assertion: <%v> does not implement <%v>", tout, tin)
	}
	fun := val.Fun.(func(*Env) r.Value)    // val returns an interface... must be already wrapped in a reflect.Value
	fail := []r.Value{r.Zero(tout), False} // returned by type assertion in case of failure

	var ret func(env *Env) (r.Value, []r.Value)

	if IsOptimizedKind(kout) {
		ret = func(env *Env) (r.Value, []r.Value) {
			v := fun(env)
			v = r.ValueOf(v.Interface()) // rebuild reflect.Value with concrete type
			if v.Type() != tout {
				return fail[0], fail
			}
			return v, []r.Value{v, True}
		}
	} else if tout == TypeOfInterface {
		// special case, nil is a valid interface{}
		ret = func(env *Env) (r.Value, []r.Value) {
			v := fun(env).Convert(TypeOfInterface)
			return v, []r.Value{v, True}
		}
	} else if tin.Implements(tout) {
		ret = func(env *Env) (r.Value, []r.Value) {
			v := fun(env)
			// nil is not a valid tout, check for it.
			// IsNil() can be invoked only on nillable types...
			// but v.Type().Kind() should be r.Interface, which is nillable :)
			if v.IsNil() {
				return fail[0], fail
			}
			v = v.Convert(tout)
			return v, []r.Value{v, True}
		}
	} else {
		ret = func(env *Env) (r.Value, []r.Value) {
			v := fun(env)
			// nil is not a valid tout, check for it.
			// IsNil() can be invoked only on nillable types...
			// but v.Type().Kind() should be r.Interface, which is nillable :)
			if v.IsNil() {
				return fail[0], fail
			}
			v = r.ValueOf(v.Interface()) // rebuild reflect.Value with concrete type
			tconcr := v.Type()
			if tconcr != tout && !tconcr.Implements(tout) {
				return fail[0], fail
			}
			v = v.Convert(tout)
			return v, []r.Value{v, True}
		}
	}
	return exprXV([]r.Type{tout, TypeOfBool}, ret)
}

// TypeAssert1 compiles a single-valued type assertion
func (c *Comp) TypeAssert1(node *ast.TypeAssertExpr) *Expr {
	val := c.Expr1(node.X)
	tin := val.Type
	tout := c.Type(node.Type)
	kout := tout.Kind()
	if tin == nil || tin.Kind() != r.Interface {
		c.Errorf("invalid type assertion: %v (non-interface type <%v> on left)", node, tin)
		return nil
	}
	if tout.Kind() != r.Interface && !tout.Implements(tin) {
		c.Errorf("impossible type assertion: <%v> does not implement <%v>", tout, tin)
	}
	fun := val.Fun.(func(*Env) r.Value) // val returns an interface... must be already wrapped in a reflect.Value
	var ret I
	switch kout {
	case r.Bool:
		ret = func(env *Env) bool {
			return fun(env).Interface().(bool)
		}
	case r.Int:
		ret = func(env *Env) int {
			return fun(env).Interface().(int)
		}
	case r.Int8:
		ret = func(env *Env) int8 {
			return fun(env).Interface().(int8)
		}
	case r.Int16:
		ret = func(env *Env) int16 {
			return fun(env).Interface().(int16)
		}
	case r.Int32:
		ret = func(env *Env) int32 {
			return fun(env).Interface().(int32)
		}
	case r.Int64:
		ret = func(env *Env) int64 {
			return fun(env).Interface().(int64)
		}
	case r.Uint:
		ret = func(env *Env) uint {
			return fun(env).Interface().(uint)
		}
	case r.Uint8:
		ret = func(env *Env) uint8 {
			return fun(env).Interface().(uint8)
		}
	case r.Uint16:
		ret = func(env *Env) uint16 {
			return fun(env).Interface().(uint16)
		}
	case r.Uint32:
		ret = func(env *Env) uint32 {
			return fun(env).Interface().(uint32)
		}
	case r.Uint64:
		ret = func(env *Env) uint64 {
			return fun(env).Interface().(uint64)
		}
	case r.Uintptr:
		ret = func(env *Env) uintptr {
			return fun(env).Interface().(uintptr)
		}
	case r.Float32:
		ret = func(env *Env) float32 {
			return fun(env).Interface().(float32)
		}
	case r.Float64:
		ret = func(env *Env) float64 {
			return fun(env).Interface().(float64)
		}
	case r.Complex64:
		ret = func(env *Env) complex64 {
			return fun(env).Interface().(complex64)
		}
	case r.Complex128:
		ret = func(env *Env) complex128 {
			return fun(env).Interface().(complex128)
		}
	case r.String:
		ret = func(env *Env) string {
			return fun(env).Interface().(string)
		}
	default:
		if tout == TypeOfInterface {
			// special case, nil is a valid interface{}
			ret = func(env *Env) r.Value {
				return fun(env).Convert(TypeOfInterface)
			}
			break
		}
		if tin.Implements(tout) {
			ret = func(env *Env) r.Value {
				v := fun(env)
				// nil is not a valid tout, check for it.
				// IsNil() can be invoked only on nillable types...
				// but v.Type().Kind() should be r.Interface, which is nillable :)
				if v.IsNil() {
					panic(&TypeAssertionError{
						Interface: tin,
						Concrete:  nil,
						Asserted:  tout,
					})
				}
				return v.Convert(tout)
			}
			break
		}
		ret = func(env *Env) r.Value {
			v := fun(env)
			// nil is not a valid tout, check for it.
			// IsNil() can be invoked only on nillable types...
			// but v.Type().Kind() should be r.Interface, which is nillable :)
			if v.IsNil() {
				panic(&TypeAssertionError{
					Interface: tin,
					Concrete:  nil,
					Asserted:  tout,
				})
			}
			v = r.ValueOf(v.Interface()) // rebuild reflect.Value with concrete type
			tconcr := v.Type()
			if tconcr != tout && !tconcr.Implements(tout) {
				panic(&TypeAssertionError{
					Interface: tin,
					Concrete:  tconcr,
					Asserted:  tout,
				})
			}
			return v.Convert(tout)
		}
	}
	return exprFun(tout, ret)
}

// A TypeAssertionError explains a failed type assertion.
type TypeAssertionError struct {
	Interface     r.Type
	Concrete      r.Type
	Asserted      r.Type
	MissingMethod string // one method needed by Interface, missing from Concrete
}

func (*TypeAssertionError) RuntimeError() {}

func (e *TypeAssertionError) Error() string {
	in := e.Interface
	concr := e.Concrete
	if concr == nil {
		return fmt.Sprintf("interface conversion: <%v> is nil, not <%v>", in, e.Asserted)
	}
	if len(e.MissingMethod) == 0 {
		return fmt.Sprintf("interface conversion: <%v> is <%v>, not <%v>", in, concr, e.Asserted.String())
	}
	return fmt.Sprintf("interface conversion: <%v> is not <%v>: missing method ", concr, e.Asserted, e.MissingMethod)
}
