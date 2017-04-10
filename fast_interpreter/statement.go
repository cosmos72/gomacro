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

var Nop Stmt = func(env *Env) (Stmt, *Env) {
	env.IP++
	return env.Code[env.IP], env
}

var Interrupt Stmt = func(env *Env) (Stmt, *Env) {
	return env.Interrupt, env
}

func (c *Comp) Stmt(node ast.Stmt) {
	switch node := node.(type) {
	case nil:
	case *ast.AssignStmt:
		c.Assign(node)
	case *ast.BlockStmt:
		c.Block(node)
	// case *ast.BranchStmt:
	//   c.Branch(node)
	case *ast.CaseClause, *ast.CommClause:
		c.Errorf("misplaced case/default: not inside switch or select: %v <%v>", node, r.TypeOf(node))
	case *ast.DeclStmt:
		c.Decl(node.Decl)
	// case *ast.DeferStmt:
	//   c.DeferStmt(node.Call)
	case *ast.EmptyStmt:
	case *ast.ExprStmt:
		expr := c.Expr(node.X)
		if !expr.Const() {
			c.Code.Append(expr.AsStmt())
		}
	// case *ast.ForStmt:
	//   c.For(node)
	// case *ast.GoStmt:
	//   c.Go(node)
	case *ast.IfStmt:
		c.If(node)
	// case *ast.IncDecStmt:
	//   c.IncDec(node)
	// case *ast.LabeledStmt:
	//   c.Label(node)
	// case *ast.RangeStmt:
	//   c.Range(node)
	// case *ast.ReturnStmt:
	//   c.Return(node)
	// case *ast.SelectStmt:
	//   c.Select(node)
	// case *ast.SendStmt:
	//   c.Send(node)
	// case *ast.SwitchStmt:
	//   c.Switch(node)
	// case *ast.TypeSwitchStmt:
	//   c.TypeSwitch(node)
	default:
		c.Errorf("unimplemented statement: %v <%v>", node, r.TypeOf(node))
	}
}

func (c *Comp) Block(node *ast.BlockStmt) {
	if node == nil || len(node.List) == 0 {
		return
	}
	// TODO block creates a new environment
	for _, stmt := range node.List {
		c.Stmt(stmt)
	}
}

func (c *Comp) If(node *ast.IfStmt) {
	var ithen, ielse, iend int

	if node.Init != nil {
		c.Stmt(node.Init)
	}
	pred := c.Expr(node.Cond)
	flag, fun, err := pred.TryAsPred()
	if err {
		c.invalidPred(node.Cond, pred)
		return
	}
	// TODO "if" creates a new environment
	if fun != nil {
		c.Code.Append(func(env *Env) (Stmt, *Env) {
			if fun(env) {
				env.IP = ithen
				return env.Code[ithen], env
			} else {
				env.IP = ielse
				return env.Code[ielse], env
			}
		})
	}
	// compile 'then' branch
	ithen = c.Code.Len()
	c.Block(node.Body)
	if fun == nil && !flag {
		// 'then' branch is never executed...
		// still compiled above (to check for errors) but drop the generated code
		c.Code.List = c.Code.List[0:ithen]
	}
	// compile a 'goto' between 'then' and 'else' branches
	if fun != nil && node.Else != nil {
		c.Code.Append(func(env *Env) (Stmt, *Env) {
			// after executing 'then' branch, we must skip 'else' branch
			env.IP = iend
			return env.Code[iend], env
		})
	}
	// compile 'else' branch
	ielse = c.Code.Len()
	if node.Else != nil {
		c.Stmt(node.Else)
		if fun == nil && flag {
			// 'else' branch is never executed...
			// still compiled above (to check for errors) but drop the generated code
			c.Code.List = c.Code.List[0:ielse]
		}
	}
	iend = c.Code.Len()
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
			//  rets[i], _ = value(env)
			// }
			ret0 := None
			if len(rets) > 0 {
				ret0 = rets[0]
			}
			panic(SReturn{ret0, rets})
		}
	}
}
