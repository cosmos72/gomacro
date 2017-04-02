/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * expr.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package compiler

import (
	"go/ast"
	r "reflect"

	"github.com/cosmos72/gomacro/constants"
)

func (c *Comp) ExprsMultipleValues(nodes []ast.Expr, expectedValuesN int) (inits []I) {
	n := len(nodes)
	if n != expectedValuesN {
		if n != 1 {
			c.Errorf("value count mismatch: cannot assign %d values to %d places: %v",
				n, expectedValuesN, nodes)
			return nil
		}
		node := nodes[0]
		inits = []I{c.Expr(node)}
	} else {
		inits = c.Exprs(nodes)
	}
	return inits
}

// Exprs compiles multiple expressions
func (c *Comp) Exprs(nodes []ast.Expr) []I {
	var rets []I
	if n := len(nodes); n != 0 {
		rets = make([]I, n)
		for i := range nodes {
			rets[i] = c.Expr(nodes[i])
		}
	}
	return rets
}

// Expr compiles an expression
func (c *Comp) Expr(in ast.Expr) I {
	for {
		// env.Debugf("evalExpr() %v", node)
		switch node := in.(type) {
		case *ast.BasicLit:
			return c.BasicLit(node)
		case *ast.BinaryExpr:
			return c.BinaryExpr(node)
		case *ast.CallExpr:
		case *ast.CompositeLit:
		case *ast.FuncLit:
		case *ast.Ident:
			return c.Ident(node.Name)
		case *ast.IndexExpr:
		case *ast.ParenExpr:
			in = node.X
			continue
		case *ast.UnaryExpr:
		case *ast.SelectorExpr:
		case *ast.SliceExpr:
		case *ast.StarExpr:
		case *ast.TypeAssertExpr:
		}
		c.Errorf("unimplemented Compile() for: %#v <%v>", in, r.TypeOf(in))
		return nil
	}
}

func (c *Comp) CallInt(fun X, args ...X) func(*Env) int {
	return func(env *Env) int {
		fvalue, _ := fun(env)
		f := fvalue.Interface().(FuncInt)
		n := len(args)
		values := make([]r.Value, n)
		for i, arg := range args {
			values[i], _ = arg(env)
		}
		return f(values...)
	}
}

func (c *Comp) evalConst(expr I) I {
	if expr == nil {
		return nil
	}
	exprv := r.ValueOf(expr)
	if exprv.Kind() != r.Func {
		return expr
	}
	rets := exprv.Call(NilEnv)
	if len(rets) == 0 {
		return nil
	}
	ret0 := rets[0]
	if ret0 == constants.Nil || ret0 == None {
		return nil
	}
	return ret0.Interface()
}
