package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

func unsupportedBinaryExpr(x interface{}, op token.Token, y interface{}) error {
	return errors.New(fmt.Sprintf("unsupported binary operation %s between %T and %T: %#v %s %#v\n", op, x, y, x, op, y))
}

func (ir *Interpreter) evalBinaryExpr(expr *ast.BinaryExpr) (reflect.Value, error) {
	x, err := ir.Eval(expr.X)
	if err != nil {
		return Nil, err
	}
	y, err := ir.Eval(expr.Y)
	if err != nil {
		return Nil, err
	}
	op := expr.Op
	xf := x.Interface()
	yf := y.Interface()
	switch xf := xf.(type) {
	case bool:
		return ir.evalBinaryExprBool(xf, op, yf)
	case int64:
		return ir.evalBinaryExprInt(xf, op, yf)
	case uint64:
		return ir.evalBinaryExprUint(xf, op, yf)
	case float64:
		return ir.evalBinaryExprFloat(xf, op, yf)
	case complex128:
		return ir.evalBinaryExprComplex(xf, op, yf)
	default:
		return Nil, unsupportedBinaryExpr(xf, op, yf)
	}
}

func (ir *Interpreter) evalBinaryExprBool(x bool, op token.Token, y interface{}) (reflect.Value, error) {
	var ret bool
	switch y := y.(type) {
	case bool:
		switch op {
		case token.EQL:
			ret = x == y
		case token.NEQ:
			ret = x != y
		default:
			return Nil, unsupportedBinaryExpr(x, op, y)
		}
		return reflect.ValueOf(ret), nil
	default:
		return Nil, unsupportedBinaryExpr(x, op, y)
	}
}

func (ir *Interpreter) evalBinaryExprUint(x uint64, op token.Token, y interface{}) (reflect.Value, error) {
	switch y := y.(type) {

	case uint64:
		var ret interface{}
		switch op {
		case token.ADD:
			ret = x + y
		case token.SUB:
			ret = x - y
		case token.MUL:
			ret = x * y
		case token.QUO:
			ret = x / y
		case token.REM:
			ret = x % y
		case token.AND:
			ret = x & y
		case token.OR:
			ret = x | y
		case token.XOR:
			ret = x ^ y
		case token.SHL:
			ret = x << y
		case token.SHR:
			ret = x >> y
		case token.AND_NOT:
			ret = x &^ y
		case token.EQL:
			ret = x == y
		case token.LSS:
			ret = x < y
		case token.GTR:
			ret = x < y
		case token.NEQ:
			ret = x != y
		case token.LEQ:
			ret = x <= y
		case token.GEQ:
			ret = x >= y
		default:
			return Nil, unsupportedBinaryExpr(x, op, y)
		}
		return reflect.ValueOf(ret), nil

	case int64:
		return ir.evalBinaryExprInt(int64(x), op, y)

	case float64:
		return ir.evalBinaryExprFloat(float64(x), op, y)

	case complex128:
		return ir.evalBinaryExprComplex(complex(float64(x), 0.0), op, y)
	}
	return Nil, unsupportedBinaryExpr(x, op, y)
}

func (ir *Interpreter) evalBinaryExprInt(x int64, op token.Token, y interface{}) (reflect.Value, error) {
	var ret interface{}
	if y, ok := toInt(y); ok {
		switch op {
		case token.ADD:
			ret = x + y
		case token.SUB:
			ret = x - y
		case token.MUL:
			ret = x * y
		case token.QUO:
			ret = x / y
		case token.REM:
			ret = x % y
		case token.AND:
			ret = x & y
		case token.OR:
			ret = x | y
		case token.XOR:
			ret = x ^ y
		case token.SHL:
			// in Go, x << y and x >> y require y to be unsigned
			ret = x << uint64(y)
		case token.SHR:
			ret = x >> uint64(y)
		case token.AND_NOT:
			ret = x &^ y
		case token.EQL:
			ret = x == y
		case token.LSS:
			ret = x < y
		case token.GTR:
			ret = x < y
		case token.NEQ:
			ret = x != y
		case token.LEQ:
			ret = x <= y
		case token.GEQ:
			ret = x >= y
		default:
			return Nil, unsupportedBinaryExpr(x, op, y)
		}
		return reflect.ValueOf(ret), nil
	}
	switch y := y.(type) {
	case float64:
		return ir.evalBinaryExprFloat(float64(x), op, y)
	case complex128:
		return ir.evalBinaryExprComplex(complex(float64(x), 0.0), op, y)
	}
	return Nil, unsupportedBinaryExpr(x, op, y)
}

func (ir *Interpreter) evalBinaryExprFloat(x float64, op token.Token, y interface{}) (reflect.Value, error) {
	var ret interface{}
	if y, ok := toFloat(y); ok {
		switch op {
		case token.ADD:
			ret = x + y
		case token.SUB:
			ret = x - y
		case token.MUL:
			ret = x * y
		case token.QUO:
			ret = x / y
		case token.EQL:
			ret = x == y
		case token.LSS:
			ret = x < y
		case token.GTR:
			ret = x < y
		case token.NEQ:
			ret = x != y
		case token.LEQ:
			ret = x <= y
		case token.GEQ:
			ret = x >= y
		default:
			return Nil, unsupportedBinaryExpr(x, op, y)
		}
		return reflect.ValueOf(ret), nil
	}
	if y, ok := y.(complex128); ok {
		return ir.evalBinaryExprComplex(complex(x, 0.0), op, y)
	}
	return Nil, unsupportedBinaryExpr(x, op, y)
}

func (ir *Interpreter) evalBinaryExprComplex(x complex128, op token.Token, y interface{}) (reflect.Value, error) {
	var ret interface{}
	if y, ok := toComplex(y); ok {
		switch op {
		case token.ADD:
			ret = x + y
		case token.SUB:
			ret = x - y
		case token.MUL:
			ret = x * y
		case token.QUO:
			ret = x / y
		case token.EQL:
			ret = x == y
		case token.NEQ:
			ret = x != y
		default:
			return Nil, unsupportedBinaryExpr(x, op, y)
		}
		return reflect.ValueOf(ret), nil
	}
	return Nil, unsupportedBinaryExpr(x, op, y)
}
