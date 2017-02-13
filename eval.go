package main

import (
	"errors"
	"fmt"
	"go/ast"
	"reflect"
)

func (ir *Interpreter) Eval(node ast.Node) (reflect.Value, error) {
	switch node := node.(type) {

	case *ast.BasicLit:
		return ir.evalLiteral(node)

	case *ast.Ident:
		return ir.evalIdentifier(node)

	case *ast.UnaryExpr:
		return ir.evalUnaryExpr(node)

	case *ast.BinaryExpr:
		return ir.evalBinaryExpr(node)

	case *ast.File:
		return ir.evalFile(node)

	default:
		return Nil, errors.New(fmt.Sprintf("unsupported expression: %#v\n", node))
	}
}
