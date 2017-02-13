// unaryexpr
package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

func unsupportedUnaryExpr(x interface{}, op token.Token) error {
	return errors.New(fmt.Sprintf("unsupported unary expression %s on %T: %s %#v\n", op, x, op, x))
}

func (ir *Interpreter) evalUnaryExpr(expr *ast.UnaryExpr) (reflect.Value, error) {
	x, err := ir.Eval(expr.X)
	if err != nil {
		return Nil, err
	}
	op := expr.Op
	switch x := x.Interface().(type) {
	case bool:
		return ir.evalUnaryExprBool(x, op)
	case int64:
		return ir.evalUnaryExprInt(x, op)
	case uint64:
		return ir.evalUnaryExprUint(x, op)
	case float64:
		return ir.evalUnaryExprFloat(x, op)
	case complex128:
		return ir.evalUnaryExprComplex(x, op)
	default:
		return Nil, unsupportedUnaryExpr(x, op)
	}
}

func (ir *Interpreter) evalUnaryExprBool(x bool, op token.Token) (reflect.Value, error) {
	var ret bool
	switch op {
	case token.NOT:
		ret = !x
	default:
		return Nil, unsupportedUnaryExpr(x, op)
	}
	return reflect.ValueOf(ret), nil
}

func (ir *Interpreter) evalUnaryExprUint(x uint64, op token.Token) (reflect.Value, error) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -int64(x)
	case token.XOR:
		ret = ^x
	default:
		return Nil, unsupportedUnaryExpr(x, op)
	}
	return reflect.ValueOf(ret), nil
}

func (ir *Interpreter) evalUnaryExprInt(x int64, op token.Token) (reflect.Value, error) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -x
	case token.XOR:
		ret = ^x
	default:
		return Nil, unsupportedUnaryExpr(x, op)
	}
	return reflect.ValueOf(ret), nil
}

func (ir *Interpreter) evalUnaryExprFloat(x float64, op token.Token) (reflect.Value, error) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -x
	default:
		return Nil, unsupportedUnaryExpr(x, op)
	}
	return reflect.ValueOf(ret), nil
}

func (ir *Interpreter) evalUnaryExprComplex(x complex128, op token.Token) (reflect.Value, error) {
	var ret interface{}
	switch op {
	case token.ADD:
		ret = x
	case token.SUB:
		ret = -x
	default:
		return Nil, unsupportedUnaryExpr(x, op)
	}
	return reflect.ValueOf(ret), nil
}
