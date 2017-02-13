package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

func (ir *Interpreter) evalDecl(node *ast.GenDecl) (reflect.Value, error) {
	tok := node.Tok
	var ret reflect.Value
	var err error
	for _, decl := range node.Specs {
		switch tok {
		case token.IMPORT:
			ret, err = ir.evalImport(decl)
		case token.CONST:
			ret, err = ir.evalDeclConst(decl)
		case token.TYPE:
			ret, err = ir.evalDeclType(decl)
		case token.VAR:
			ret, err = ir.evalDeclVar(decl)
		default:
			return Nil, errors.New(fmt.Sprintf("unsupported declaration: %#v", decl))
		}
		if err != nil {
			ir.PopEnv()
			return Nil, err
		}
	}
	return ret, nil
}

func (ir *Interpreter) evalImport(node ast.Spec) (reflect.Value, error) {
	switch node := node.(type) {
	default:
		return Nil, errors.New(fmt.Sprintf("unsupported import %#v", node))
	}
}

func (ir *Interpreter) evalDeclConst(node ast.Spec) (reflect.Value, error) {
	switch node := node.(type) {
	default:
		return Nil, errors.New(fmt.Sprintf("unsupported constant declaration %#v", node))
	}
}

func (ir *Interpreter) evalDeclType(node ast.Spec) (reflect.Value, error) {
	switch node := node.(type) {
	default:
		return Nil, errors.New(fmt.Sprintf("unsupported type declaration %#v", node))
	}
}

func (ir *Interpreter) evalDeclVar(node ast.Spec) (reflect.Value, error) {
	var ret reflect.Value
	switch node := node.(type) {
	case *ast.ValueSpec:
		idents := node.Names
		values := node.Values
		t, err := ir.evalType(node.Type)
		if err != nil {
			return Nil, err
		}
		zero := reflect.Zero(t)
		for i, ident := range idents {
			if values != nil && len(values) >= i {
				value, err := ir.Eval(values[i])
				if err == nil {
					ret, err = ir.defineVar(ident.Name, t, value)
				}
				if err != nil {
					return Nil, err
				}
			} else {
				ret, err = ir.defineVar(ident.Name, t, zero)
				if err != nil {
					return Nil, err
				}
			}
		}
	default:
		return Nil, errors.New(fmt.Sprintf("unsupported variable declaration: %#v", node))
	}
	return ret, nil
}

func (ir *Interpreter) defineVarConvert(name string, t reflect.Type, value reflect.Value) (reflect.Value, error) {
	value_as_t, ok := toType(value, t)
	if !ok {
		return Nil, errors.New(fmt.Sprintf("failed to cast %#v to %v in variable declaration: var %s %v = %#v", value, t, name, t, value))
	}
	return ir.defineVar(name, t, value_as_t)
}

func (ir *Interpreter) defineVar(name string, t reflect.Type, value reflect.Value) (reflect.Value, error) {
	// fmt.Printf("debug: defineVar() var %s %v = %#v\n", name, t, value.Interface())

	addr := reflect.New(t)
	place := addr.Elem()

	_, err := ir.assign(place, token.ASSIGN, value)
	if err != nil {
		return Nil, err
	}
	ir.Binds[name] = place
	return reflect.ValueOf(name), nil
}
