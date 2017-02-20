/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * literal.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	"go/constant"
	"go/token"
	r "reflect"
	"strconv"
	"strings"
)

var Unknown = constant.MakeUnknown()

func (env *Env) evalLiteral(node *ast.BasicLit) (r.Value, []r.Value) {
	ret := env.evalLiteral0(node)
	return r.ValueOf(ret), nil
}

func (env *Env) evalLiteral0(node *ast.BasicLit) interface{} {
	kind := node.Kind
	str := node.Value
	var ret interface{}

	switch kind {

	case token.INT:
		if strings.HasPrefix(str, "-") {
			i64, err := strconv.ParseInt(str, 0, 0)
			if err != nil {
				return Errore(err)
			}
			// prefer int to int64. reason: in compiled Go,
			// type inference deduces int for all constants representable by an int
			i := int(i64)
			if int64(i) == i64 {
				return i
			}
			return i64
		} else {
			u64, err := strconv.ParseUint(str, 0, 0)
			if err != nil {
				return Errore(err)
			}
			// prefer, in order: int, int64, uint, uint64. reason: in compiled Go,
			// type inference deduces int for all constants representable by an int
			i := int(u64)
			if i >= 0 && uint64(i) == u64 {
				return i
			}
			i64 := int64(u64)
			if i64 >= 0 && uint64(i64) == u64 {
				return i64
			}
			u := uint(u64)
			if uint64(u) == u64 {
				return u
			}
			return u64
		}

	case token.FLOAT:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return Errore(err)
		}
		ret = f

	case token.IMAG:
		if strings.HasSuffix(str, "i") {
			str = str[:len(str)-1]
		}
		im, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return Errore(err)
		}
		ret = complex(0.0, im)
		// env.Debugf("evalLiteral(): parsed IMAG %s -> %T %#v -> %T %#v", str, im, im, ret, ret)

	case token.CHAR:
		return unescapeChar(str)

	case token.STRING:
		return unescapeString(str)

	default:
		env.Errorf("unimplemented basic literal: %v", node)
		ret = nil
	}
	return ret
}

func (env *Env) evalCompositeLiteral(node *ast.CompositeLit) (r.Value, []r.Value) {
	//	prefix := node.Type
	// t := env.evalType(prefix)
	return env.Errorf("unimplemented composite literal: %v", node)
}

// lambda()
func (env *Env) evalFunctionLiteral(node *ast.FuncLit) (r.Value, []r.Value) {
	// env.Debugf("func() at position %v", node.Type.Func)

	ret, _ := env.evalDeclFunction(nil, node.Type, node.Body)
	return ret, nil
}
