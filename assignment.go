package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

func (ir *Interpreter) evalAssignments(node *ast.AssignStmt) (reflect.Value, error) {
	lhs := node.Lhs
	rhs := node.Rhs
	op := node.Tok
	var ret reflect.Value
	var err error
	for i, n := 0, len(lhs); i < n; i++ {
		ret, err = ir.evalAssignment(lhs[i], op, rhs[i])
		if err != nil {
			return Nil, err
		}
	}
	return ret, nil
}

func (ir *Interpreter) evalAssignment(lhs ast.Expr, op token.Token, rhs ast.Expr) (reflect.Value, error) {
	// fmt.Printf("debug: evalAssignment() [%v] %s [%v]\n", lhs, op, rhs)
	place, err := ir.Eval(lhs)
	if err != nil {
		return Nil, err
	}
	if !place.CanSet() {
		return Nil, errors.New(fmt.Sprintf("cannot assign to read-only location: %#v %s %#v", lhs, op, rhs))
	}
	value, err := ir.Eval(rhs)
	if err != nil {
		return Nil, err
	}
	return ir.assign(place, token.ASSIGN, value)
}

func (ir *Interpreter) assign(place reflect.Value, op token.Token, value reflect.Value) (reflect.Value, error) {
	t := place.Type()
	if !value.Type().AssignableTo(t) {
		if value.CanSet() {
			return Nil, errors.New(fmt.Sprintf("failed to assign %#v to %v", value, t))
		}
		if !value.Type().ConvertibleTo(t) {
			return Nil, errors.New(fmt.Sprintf("failed to convert %#v to %v", value, t))
		}
		if primitiveTypeOverflows(value, place) {
			return Nil, errors.New(fmt.Sprintf("constant %#v overflows %v", value, t))
		}
		value = value.Convert(t)
	}
	// TODO op can be = += -= *= ...
	place.Set(value)
	return value, nil
}
