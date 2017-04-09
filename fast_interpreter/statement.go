/*
 * gomacro - A Go intepreter with Lisp-like macros
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
 * statement.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

var NilStmt = Stmt{}

var Nop = Stmt{func(env *Env) (Stmt, *Env) {
	env.IP++
	return env.Code[env.IP], env
}}

func (s Stmt) Nil() bool {
	return s.Exec == nil
}

func (c *Comp) Stmt(node ast.Stmt) Stmt {
	switch node := node.(type) {
	case *ast.AssignStmt:
		// return c.Assign(node)
	case *ast.BlockStmt:
		// return c.Block(node)
	case *ast.BranchStmt:
		// return env.Branch(node)
	case *ast.CaseClause, *ast.CommClause:
		c.Errorf("misplaced case: not inside switch or select: %v <%v>", node, r.TypeOf(node))
		return NilStmt
	case *ast.DeclStmt:
		// return c.DeclStmt(node.Decl)
	case *ast.DeferStmt:
		// return c.DeferStmt(node.Call)
	case *ast.EmptyStmt:
		return NilStmt
	case *ast.ExprStmt:
		expr := c.Expr(node.X)
		if expr.Const() {
			return NilStmt
		} else {
			return expr.AsStmt()
		}
	case *ast.ForStmt:
		// return c.For(node)
	case *ast.GoStmt:
		// return c.Go(node)
	case *ast.IfStmt:
		return c.If(node)
	case *ast.IncDecStmt:
		// return c.IncDec(node)
	case *ast.LabeledStmt:
		// return c.Label(node)
	case *ast.RangeStmt:
		// return c.Range(node)
	case *ast.ReturnStmt:
		// return c.Return(node)
	case *ast.SelectStmt:
		// return c.Select(node)
	case *ast.SendStmt:
		// return c.Send(node)
	case *ast.SwitchStmt:
		// return c.Switch(node)
	case *ast.TypeSwitchStmt:
		// return c.TypeSwitch(node)
	default:
		c.Errorf("invalid statement: %v <%v>", node, r.TypeOf(node))
		return NilStmt
	}
	c.Errorf("unimplemented statement: %v <%v>", node, r.TypeOf(node))
	return NilStmt
}

func (c *Comp) Block(node *ast.BlockStmt) Stmt {
	// TODO
	return NilStmt
}

func (c *Comp) If(node *ast.IfStmt) Stmt {
	var then, els, ret Stmt

	if node.Init != nil {
		c.Stmt(node.Init)
	}
	pred := c.Expr(node.Cond)
	flag, fun, err := pred.TryAsPred()
	if err {
		return c.invalidPred(node.Cond, pred)
	}
	then = c.Block(node.Body)
	if node.Else != nil {
		els = c.Stmt(node.Else)
	}
	// TODO "if" creates a new environment
	if fun != nil {
		ret = Stmt{func(env *Env) (Stmt, *Env) {
			if fun(env) {
				return then, env
			} else {
				return els, env
			}
		}}
	} else if flag {
		ret = then
	} else {
		ret = els
	}
	return ret
}

func For(init X, pred func(*Env) bool, post X, body X) X {
	if init == nil && post == nil {
		return func(env *Env) {
			for pred(env) {
				body(env)
			}
		}

	} else {
		if init == nil || post == nil {
			panic("invalid for(): init and post must be both present, or both omitted")
		}
		return func(env *Env) {
			for init(env); pred(env); post(env) {
				body(env)
			}
		}
	}
}

func RemoveNils(list []X) []X {
	j, n := 0, len(list)
	for i := 0; i < n; i++ {
		x := list[i]
		if x != nil {
			list[j] = x
			j++
		}
	}
	return list[0:j]
}

func Block(list ...X) X {
	list = RemoveNils(list)
	switch len(list) {
	case 0:
		return nil
	case 1:
		return list[0]
	case 2:
		return func(env *Env) {
			list[0](env)
			list[1](env)
		}
	default:
		return func(env *Env) {
			for _, x := range list {
				x(env)
			}
		}
	}
}

func Return(exprs ...X) X {
	switch n := len(exprs); n {
	case 0:
		return nil
	case 1:
		// expr := exprs[0]
		// return foo() returns *all* the values returned by foo, not just the first one
		return func(env *Env) {
			// ret, rets := expr(env)
			// panic(SReturn{ret, rets})
		}
	default:
		return func(env *Env) {
			n := len(exprs)
			rets := make([]r.Value, n)
			// for i, value := range exprs {
			//	 rets[i], _ = value(env)
			// }
			ret0 := None
			if len(rets) > 0 {
				ret0 = rets[0]
			}
			panic(SReturn{ret0, rets})
		}
	}
}
