package main

import (
	"errors"
	"fmt"
	"go/ast"
	"reflect"
)

func (ir *Interpreter) evalIdentifier(expr *ast.Ident) (reflect.Value, error) {
	name := expr.Name

	switch name {
	case "false":
		return reflect.ValueOf(false), nil
	case "true":
		return reflect.ValueOf(true), nil
	case "iota":
		pos := ir.Fileset.Position(expr.NamePos)
		return reflect.ValueOf(pos.Line - ir.iotaOffset), nil
	default:
		for env := ir.Env; env != nil; env = env.Outer {
			bind, exists := env.Binds[name]
			if exists {
				return bind, nil
			}
		}
		return Nil, errors.New(fmt.Sprintf("undefined identifier: %s", name))
	}
}
