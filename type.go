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
 *  Created on: Feb 13, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"errors"
	"fmt"
	"go/ast"
	r "reflect"
)

func (ir *Interpreter) evalType(node ast.Expr) (r.Type, error) {

	switch node := node.(type) {
	case *ast.Ident:
		return ir.evalTypeIdentifier(node.Name)
	case *ast.FuncType:
		return ir.evalTypeFunction(node)
	default:
		// TODO *ast.InterfaceType and many others
		if node == nil {
			// type can be omitted in many case - then we must perform type inference
			return nil, nil
		}
		return nil, errors.New(fmt.Sprintf("unimplemented type: %#v", node))
	}
}

func (ir *Interpreter) evalTypeFunction(node *ast.FuncType) (r.Type, error) {

	params, err := ir.evalTypeFields(node.Params)
	if err != nil {
		return nil, err
	}
	results, err := ir.evalTypeFields(node.Results)
	if err != nil {
		return nil, err
	}
	return r.FuncOf(params, results, false /* TODO variadic*/), nil
}

func (ir *Interpreter) evalTypeFields(fields *ast.FieldList) ([]r.Type, error) {
	ts := make([]r.Type, 0)
	if fields == nil || len(fields.List) == 0 {
		return ts, nil
	}
	for _, f := range fields.List {

		t, err := ir.evalType(f.Type)
		if err != nil {
			return nil, err
		}
		for range f.Names {
			ts = append(ts, t)
			// fmt.Printf("debug: evalTypeFields() %v %v -> %v\n", f.Names[i], f.Type, t)
		}
	}
	return ts, nil
}

func (ir *Interpreter) evalTypeFieldsNames(fields *ast.FieldList) []string {
	names := make([]string, 0)
	if fields == nil || len(fields.List) == 0 {
		return names
	}
	for _, f := range fields.List {
		for _, ident := range f.Names {
			names = append(names, ident.Name)
		}
	}
	return names
}

func (ir *Interpreter) evalTypeIdentifier(name string) (r.Type, error) {
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
	case "rune", "uint32":
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
		return nil, errors.New(fmt.Sprintf("unimplemented type identifier: %v", name))
	}
	return r.TypeOf(v), nil
}

func toType(value r.Value, t r.Type) (r.Value, bool) {
	if value.Type().ConvertibleTo(t) {
		return value.Convert(t), true
	}
	return Nil, false
}

func primitiveTypeOverflows(value r.Value, place r.Value) bool {
	kv, kp := value.Type().Kind(), place.Type().Kind()
	if isIntKind(kv) {
		v := value.Int()
		if isIntKind(kp) {
			return place.OverflowInt(v)
		} else if isUintKind(kp) {
			return v < 0 || place.OverflowUint(uint64(v))
		} else if isFloatKind(kp) {
			return place.OverflowFloat(float64(v))
		} else if isComplexKind(kp) {
			return place.OverflowComplex(complex(float64(v), 0.0))
		} else {
			return false
		}
	} else if isUintKind(kv) {
		v := value.Uint()
		if isIntKind(kp) {
			return v > 0x7fffffffffffffff || place.OverflowInt(int64(v))
		} else if isUintKind(kp) {
			return place.OverflowUint(v)
		} else if isFloatKind(kp) {
			return place.OverflowFloat(float64(v))
		} else if isComplexKind(kp) {
			return place.OverflowComplex(complex(float64(v), 0.0))
		} else {
			return false
		}
	} else if isFloatKind(kv) {
		v := value.Float()
		if isIntKind(kp) {
			return float64(int64(v)) != v || place.OverflowInt(int64(v))
		} else if isUintKind(kp) {
			return float64(int64(v)) != v || place.OverflowUint(uint64(v))
		} else if isFloatKind(kp) {
			return place.OverflowFloat(v)
		} else if isComplexKind(kp) {
			return place.OverflowComplex(complex(v, 0.0))
		} else {
			return false
		}
	} else if isComplexKind(kv) {
		v := value.Complex()
		re := real(v)
		im := imag(v)
		if isIntKind(kp) {
			return im != 0.0 || float64(int64(re)) != re || place.OverflowInt(int64(re))
		} else if isUintKind(kp) {
			return im != 0.0 || float64(uint64(re)) != re || place.OverflowUint(uint64(re))
		} else if isFloatKind(kp) {
			return im != 0.0 || place.OverflowFloat(re)
		} else if isComplexKind(kp) {
			return place.OverflowComplex(v)
		} else {
			return false
		}
	}
	return false
}

func isPrimitiveKind(kind r.Kind) bool {
	return isIntKind(kind) || isUintKind(kind) || isFloatKind(kind) || isComplexKind(kind)
}

func isIntKind(kind r.Kind) bool {
	switch kind {
	case r.Int, r.Int8, r.Int16, r.Int32, r.Int64:
		return true
	default:
		return false
	}
}

func isUintKind(kind r.Kind) bool {
	switch kind {
	case r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr:
		return true
	default:
		return false
	}
}

func isFloatKind(kind r.Kind) bool {
	switch kind {
	case r.Float32, r.Float64:
		return true
	default:
		return false
	}
}

func isComplexKind(kind r.Kind) bool {
	switch kind {
	case r.Complex64, r.Complex128:
		return true
	default:
		return false
	}
}
