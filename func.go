package main

import (
	"errors"
	"fmt"
	"go/ast"
	"reflect"
)

func (ir *Interpreter) evalDeclFunc(node *ast.FuncDecl) (reflect.Value, error) {
	name := node.Name.Name
	if name == TemporaryFunctionName {
		return ir.evalBlock(node.Body.List)
	}
	// TODO
	return Nil, errors.New(fmt.Sprintf("unsupported function declaration: %#v", node))
}
