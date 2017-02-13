package main

import (
	"errors"
	"fmt"
	"go/ast"
	"reflect"
)

func (ir *Interpreter) evalFile(node *ast.File) (reflect.Value, error) {
	ir.Packagename = node.Name.Name

	// eval node.Imports
	var ret reflect.Value
	var err error

	for _, stmt := range node.Decls {
		ret, err = ir.evalStatement(stmt)
		if err != nil {
			return Nil, err
		}
	}
	return ret, nil
}

func (ir *Interpreter) evalBlock(block []ast.Stmt) (reflect.Value, error) {
	ir.PushEnv()
	defer ir.PopEnv()

	var ret reflect.Value
	var err error

	for _, stmt := range block {
		ret, err = ir.evalStatement(stmt)
		if err != nil {
			return Nil, err
		}
	}
	return ret, nil
}

func (ir *Interpreter) evalStatement(node ast.Node) (reflect.Value, error) {
	switch node := node.(type) {
	case *ast.GenDecl:
		return ir.evalDecl(node)
	case *ast.FuncDecl:
		return ir.evalDeclFunc(node)
	case *ast.AssignStmt:
		return ir.evalAssignments(node)
	default:
		return Nil, errors.New(fmt.Sprintf("unsupported statement: %#v", node))
	}
}
