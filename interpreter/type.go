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
 * type.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"fmt"
	"go/ast"
	r "reflect"
)

func typeOf(value r.Value) r.Type {
	if value == None || value == Nil {
		return typeOfInterface
	}
	return value.Type()
}

func (env *Env) evalExpr1OrType(node ast.Expr) (val r.Value, t r.Type) {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case runtimeError:
				t = env.evalType(node)
			default:
				panic(r)
			}
		}
	}()
	val = env.evalExpr1(node)
	return val, nil
}

func (env *Env) evalType(node ast.Expr) r.Type {
	t, _ := env.evalTypeEllipsis(node, false)
	return t
}

func (env *Env) evalTypeEllipsis(node ast.Expr, allowEllipsis bool) (t r.Type, ellipsis bool) {
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
		t = env.evalTypeArray(node)
	case *ast.ChanType:
		t = env.evalType(node.Value)
		dir := r.BothDir
		if node.Dir == ast.SEND {
			dir = r.SendDir
		} else if node.Dir == ast.RECV {
			dir = r.RecvDir
		}
		t = r.ChanOf(dir, t)
	case *ast.FuncType:
		t, _, _ = env.evalTypeFunction(node)
	case *ast.Ident:
		t = env.evalTypeIdentifier(node.Name)
	case *ast.InterfaceType:
		t, _ = env.evalTypeInterface(node)
	case *ast.MapType:
		kt := env.evalType(node.Key)
		vt := env.evalType(node.Value)
		t = r.MapOf(kt, vt)
	case *ast.SelectorExpr:
		if pkgIdent, ok := node.X.(*ast.Ident); ok {
			pkgv := env.evalIdentifier(pkgIdent)
			if pkg, ok := pkgv.Interface().(*PackageRef); ok {
				name := node.Sel.Name
				if t, ok = pkg.Types[name]; !ok {
					env.errorf("not a type: %v <%v>", node, r.TypeOf(node))
				}
			} else {
				env.errorf("not a package: %v = %v <%v>", pkgIdent, pkgv, typeOf(pkgv))
			}
		} else {
			env.errorf("unimplemented qualified type, expecting packageName.identifier: %v <%v>", node, r.TypeOf(node))
		}
	case *ast.StructType:
		// env.Debugf("evalType() struct declaration: %v <%v>", node, r.TypeOf(node))
		types, names := env.evalTypeFields(node.Fields)
		// env.Debugf("evalType() struct names and types: %v %v", types, names)
		fields := makeStructFields(env.FileEnv().Path, names, types)
		// env.Debugf("evalType() struct fields: %#v", fields)
		t = r.StructOf(fields)
	case nil:
		// type can be omitted in many case - then we must perform type inference
		break
	default:
		// TODO which types are still missing?
		env.errorf("unimplemented type: %v <%v>", node, r.TypeOf(node))
	}
	for i := 0; i < stars; i++ {
		t = r.PtrTo(t)
	}
	if ellipsis {
		t = r.SliceOf(t)
	}
	return t, ellipsis
}

func (env *Env) evalTypeArray(node *ast.ArrayType) r.Type {
	t := env.evalType(node.Elt)
	n := node.Len
	switch n := n.(type) {
	case *ast.Ellipsis:
		env.errorf("evalType(): unimplemented array type with ellipsis: %v", node, r.TypeOf(node))
	case nil:
		t = r.SliceOf(t)
	default:
		count := env.evalExpr1(n).Int()
		t = r.ArrayOf(int(count), t)
	}
	return t
}

func (env *Env) evalTypeFunction(node *ast.FuncType) (t r.Type, argNames []string, resultNames []string) {
	argTypes, argNames, variadic := env.evalTypeFieldsOrParams(node.Params, true)
	resultTypes, resultNames := env.evalTypeFields(node.Results)
	return r.FuncOf(argTypes, resultTypes, variadic), argNames, resultNames
}

func (env *Env) evalTypeFields(fields *ast.FieldList) (types []r.Type, names []string) {
	types, names, _ = env.evalTypeFieldsOrParams(fields, false)
	return types, names
}

func (env *Env) evalTypeFieldsOrParams(fields *ast.FieldList, allowEllipsis bool) (types []r.Type, names []string, ellipsis bool) {
	types = make([]r.Type, 0)
	names = zeroStrings
	if fields == nil || len(fields.List) == 0 {
		return types, names, ellipsis
	}
	list := fields.List
	n := len(list)
	var t r.Type
	for i, f := range list {
		t, ellipsis = env.evalTypeEllipsis(f.Type, i == n-1)
		if len(f.Names) == 0 {
			types = append(types, t)
			names = append(names, "_")
			// env.Debugf("evalTypeFields() %v -> %v", f.Type, t)
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

func (env *Env) evalTypeIdentifier(name string) r.Type {
	for e := env; e != nil; e = e.Outer {
		if t, ok := e.Types[name]; ok {
			return t
		}
	}
	env.errorf("undefined identifier: %v", name)
	return nil
}

func (env *Env) evalTypeInterface(node *ast.InterfaceType) (t r.Type, methodNames []string) {
	if node.Methods != nil && len(node.Methods.List) != 0 {
		env.errorf("unimplemented type: %v <%v>", node, r.TypeOf(node))
		return nil, nil
	}
	return typeOfInterface, zeroStrings
}

func makeStructFields(pkgPath string, names []string, types []r.Type) []r.StructField {
	// pkgIdentifier := sanitizeIdentifier(pkgPath)
	fields := make([]r.StructField, len(names))
	var offset, next uintptr
	for i, name := range names {
		t := types[i]
		offset, next = alignStructField(next, t)
		fields[i] = r.StructField{
			Name:      toExportedName(name), // Go 1.8 reflect.StructOf() supports *only* exported fields
			Type:      t,
			Tag:       "",
			Offset:    offset,
			Index:     []int{i},
			Anonymous: false,
		}
	}
	return fields
}

func alignStructField(offset uintptr, t r.Type) (uintptr, uintptr) {
	align := uintptr(t.FieldAlign())
	// fmt.Printf("alignStructField(offset = %d, type = <%v>) -> align = %d", offset, t, align)
	if align > 0 {
		offset = (offset + align - 1) / align * align
	}
	// fmt.Printf(", offset = %d, next = %d\n", offset, offset+t.Size())
	return offset, offset + t.Size()
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

func (env *Env) valueToType(value r.Value, t r.Type) r.Value {
	if value == None || value == Nil {
		switch t.Kind() {
		case r.Chan, r.Func, r.Interface, r.Map, r.Ptr, r.Slice:
			return r.Zero(t)
		}
	}
	vt := typeOf(value)
	if !vt.AssignableTo(t) && !vt.ConvertibleTo(t) {
		ret, _ := env.errorf("failed to convert %v <%v> to <%v>", value, vt, t)
		return ret
	}
	newValue := value.Convert(t)
	if differentIntegerValues(value, newValue) {
		env.warnf("value %d overflows <%v>, truncated to %d", value, t, newValue)
	}
	return newValue
}

func differentIntegerValues(v1 r.Value, v2 r.Value) bool {
	k1, k2 := v1.Kind(), v2.Kind()
	switch k1 {
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		n1 := v1.Int()
		switch k2 {
		case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
			return n1 != v2.Int()
		case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
			return n1 < 0 || uint64(n1) != v2.Uint()
		default:
			return false
		}
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		n1 := v1.Uint()
		switch k2 {
		case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
			n2 := v2.Int()
			return n2 < 0 || uint64(n2) != n1
		case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
			return n1 != v2.Uint()
		default:
			return false
		}
	default:
		return false
	}
}

func toValues(args []interface{}) []r.Value {
	n := len(args)
	values := make([]r.Value, n)
	for i := 0; i < n; i++ {
		values[i] = r.ValueOf(args[i])
	}
	return values
}

func toInterfaces(values []r.Value) []interface{} {
	n := len(values)
	rets := make([]interface{}, n)
	for i := 0; i < n; i++ {
		rets[i] = toInterface(values[i])
	}
	return rets
}

func toInterface(value r.Value) interface{} {
	if value != Nil && value != None {
		return value.Interface()
	}
	return nil
}
