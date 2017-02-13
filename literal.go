package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)

func (ir *Interpreter) evalLiteral(expr *ast.BasicLit) (reflect.Value, error) {
	ret, err := ir.evalLiteral0(expr)
	if err != nil {
		return Nil, err
	}
	return reflect.ValueOf(ret), nil
}

func (ir *Interpreter) evalLiteral0(expr *ast.BasicLit) (interface{}, error) {
	kind := expr.Kind
	str := expr.Value
	var ret interface{}

	switch kind {

	case token.INT:
		if strings.HasPrefix(str, "-") {
			return strconv.ParseInt(str, 0, 0)
		} else {
			return strconv.ParseUint(str, 0, 0)
		}

	case token.FLOAT:
		return strconv.ParseFloat(str, 64)

	case token.IMAG:
		if strings.HasSuffix(str, "i") {
			str = str[:len(str)-1]
		}
		im, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		ret = complex(0.0, im)
		// fmt.Printf("debug evalLiteral(): parsed IMAG %s -> %T %#v -> %T %#v\n", str, im, im, ret, ret)

	case token.CHAR:
		return unescapeChar(str)

	case token.STRING:
		ret = unescapeString(str)

	default:
		return nil, errors.New(fmt.Sprintf("unsupported literal Kind = %s, reflect.Value = %#v", kind, str))

	}
	return ret, nil
}
