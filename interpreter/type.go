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
	"go/ast"
	r "reflect"
)

func TypeOf(value r.Value) r.Type {
	if value == None || value == Nil {
		return typeOfInterface
	}
	return value.Type()
}

func (env *Env) evalType(node ast.Expr) r.Type {
	stars := 0
	for {
		if expr, ok := node.(*ast.StarExpr); ok {
			stars++
			node = expr.X
		} else {
			break
		}
	}
	var t r.Type
	switch node := node.(type) {
	case *ast.Ident:
		t = env.evalTypeIdentifier(node.Name)
	case *ast.FuncType:
		t, _, _ = env.evalTypeFunction(node)
	case *ast.InterfaceType:
		t, _ = env.evalTypeInterface(node)
	case *ast.SelectorExpr:
		if pkgIdent, ok := node.X.(*ast.Ident); ok {
			pkgv := env.evalIdentifier(pkgIdent)
			if pkg, ok := pkgv.Interface().(*Env); ok {
				name := node.Sel.Name
				if t, ok = pkg.Types[name]; !ok {
					env.Errorf("not a type: %v <%v>", node, r.TypeOf(node))
				}
			} else {
				env.Errorf("not a package: %v = %v <%v>", pkgIdent, pkgv, TypeOf(pkgv))
			}
		} else {
			env.Errorf("unimplemented qualified type, expecting packageName.identifier: %v <%v>", node, r.TypeOf(node))
		}
	default:
		// TODO *ast.StructType and many others
		// type can be omitted in many case - then we must perform type inference
		if node != nil {
			env.Errorf("evalType(): unimplemented type: %v <%v>", node, r.TypeOf(node))
		}
	}
	for i := 0; i < stars; i++ {
		t = r.PtrTo(t)
	}
	return t
}

func (env *Env) evalTypeInterface(node *ast.InterfaceType) (t r.Type, methodNames []string) {
	if node.Methods != nil && len(node.Methods.List) != 0 {
		env.Errorf("unimplemented interface { /*methods*/ }: %#v", node.Methods.List)
		return nil, nil
	}
	return typeOfInterface, zeroStrings
}

func (env *Env) evalTypeFunction(node *ast.FuncType) (t r.Type, argNames []string, resultNames []string) {
	argTypes, argNames := env.evalTypeFields(node.Params)
	resultTypes, resultNames := env.evalTypeFields(node.Results)
	return r.FuncOf(argTypes, resultTypes, false /* TODO variadic*/), argNames, resultNames
}

func (env *Env) evalTypeFields(fields *ast.FieldList) ([]r.Type, []string) {
	types := make([]r.Type, 0)
	names := zeroStrings
	if fields == nil || len(fields.List) == 0 {
		return types, names
	}
	for _, f := range fields.List {

		t := env.evalType(f.Type)
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
	return types, names
}

func (env *Env) evalTypeIdentifier(name string) r.Type {
	var v interface{}
	switch name {
	case "bool":
		v = false
	case "string":
		v = ""
	case "int":
		v = int(0)
	case "int8":
		v = int8(0)
	case "int16":
		v = int16(0)
	case "int32":
		v = int32(0)
	case "int64":
		v = int64(0)
	case "uint":
		v = uint(0)
	case "uint8", "byte":
		v = uint8(0)
	case "uint16":
		v = uint16(0)
	case "uint32", "rune":
		v = uint32(0)
	case "uint64":
		v = uint64(0)
	case "float32":
		v = float32(0)
	case "float64":
		v = float64(0)
	case "complex64":
		v = complex(float32(0), float32(0))
	case "complex128":
		v = complex(float64(0), float64(0))
	default:
		env.Errorf("unimplemented type identifier: %v", name)
		return nil
	}
	return r.TypeOf(v)
}

func (env *Env) valueToType(value r.Value, t r.Type) r.Value {
	vt := TypeOf(value)
	if !vt.AssignableTo(t) && !vt.ConvertibleTo(t) {
		ret, _ := env.Errorf("failed to convert %#v to %v", value, t)
		return ret
	}
	newValue := value.Convert(t)
	if differentIntegerValues(value, newValue) {
		env.Warnf("value %d overflows %v, truncated to %d", value, t, newValue)
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
