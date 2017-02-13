package main

import (
	"errors"
	"fmt"
	"go/ast"
	"reflect"
)

func (ir *Interpreter) evalType(node ast.Expr) (reflect.Type, error) {

	switch node := node.(type) {
	case *ast.Ident:
		return ir.evalSimpleType(node.Name)
	default:
		return nil, errors.New(fmt.Sprintf("unsupported type: %#v", node))
	}
}

func (ir *Interpreter) evalSimpleType(name string) (t reflect.Type, err error) {
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
		return nil, errors.New(fmt.Sprintf("unsupported type: %v", name))
	}
	return reflect.TypeOf(v), nil
}

func toType(value reflect.Value, t reflect.Type) (reflect.Value, bool) {
	if value.Type().ConvertibleTo(t) {
		return value.Convert(t), true
	}
	return Nil, false
}

func primitiveTypeOverflows(value reflect.Value, place reflect.Value) bool {
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

func isPrimitiveKind(kind reflect.Kind) bool {
	return isIntKind(kind) || isUintKind(kind) || isFloatKind(kind) || isComplexKind(kind)
}

func isIntKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func isUintKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	default:
		return false
	}
}

func isFloatKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func isComplexKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Complex64, reflect.Complex128:
		return true
	default:
		return false
	}
}
