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
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func Nop(env *Env) (Stmt, *Env) {
	env.IP++
	return env.Code[env.IP], env
}

// code.go needs the address of Interrupt
var Interrupt Stmt = func(env *Env) (Stmt, *Env) {
	return env.Interrupt, env
}

func PopEnv(env *Env) (Stmt, *Env) {
	outer := env.Outer
	outer.IP = env.IP + 1
	// Debugf("PopEnv, IP = %d of %d", outer.IP, len(outer.Code))
	return outer.Code[outer.IP], outer
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
	case *ast.ForStmt:
		c.For(node)
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

// Block compiles a block statement, i.e. { ... }
func (c *Comp) Block(block *ast.BlockStmt) {
	if block == nil || len(block.List) == 0 {
		return
	}
	c.Block0(block.List)
}

// Block0 compiles a block statement, i.e. { ... }
func (c *Comp) Block0(list []ast.Stmt) {
	if len(list) == 0 {
		return
	}
	var nbinds [2]int // # of binds in the block

	c2, locals := c.PushEnvIfLocalBinds(&nbinds)

	for _, node := range list {
		c2.Stmt(node)
	}

	c2.PopEnvIfLocalBinds(locals, &nbinds)

	// c.Debugf("Block compiled. inner *Comp = %#v", c2)
}

// containLocalBinds return true if one or more of the given statements (but not their contents:
// blocks are not examined) contain some function/variable declaration.
// ignores types, constants and anything named "_"
func containLocalBinds(list ...ast.Stmt) bool {
	for _, node := range list {
		switch node := node.(type) {
		case *ast.AssignStmt:
			if node.Tok == token.DEFINE {
				return true
			}
		case *ast.DeclStmt:
			switch decl := node.Decl.(type) {
			case *ast.FuncDecl:
				// Go compiler forbids local functions... we allow them
				if decl.Name != nil && decl.Name.Name != "_" {
					return true
				}
			case *ast.GenDecl:
				if decl.Tok != token.VAR {
					continue
				}
				// found local variables... bail out unless they are all named "_"
				for _, spec := range decl.Specs {
					switch spec := spec.(type) {
					case *ast.ValueSpec:
						for _, ident := range spec.Names {
							if ident.Name != "_" {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}

// PushEnvIfLocalBinds compiles a PushEnv statement if list contains local binds
// returns the *Comp to use to compile statement list.
func (c *Comp) PushEnvIfLocalBinds(nbinds *[2]int, list ...ast.Stmt) (inner *Comp, locals bool) {
	// optimization: examine statements. if none of them is a function/variable declaration,
	// no need to create a new *Env at runtime
	// note: we still create a new *Comp at compile time to handle constant/type declarations
	locals = containLocalBinds(list...)
	if locals {
		// push new *Env at runtime. we will know # of binds in the block only later, so use a closure on them
		c.Code.Append(func(env *Env) (Stmt, *Env) {
			inner := NewEnv(env, nbinds[0], nbinds[1])
			inner.IP++
			// Debugf("PushEnv, IP = %d of %d, pushed %d binds and %d intbinds", inner.IP, nbinds[0], nbinds[1])
			return inner.Code[inner.IP], inner
		})
	}
	inner = NewComp(c)
	if !locals {
		inner.UpCost = 0
	}
	return inner, locals
}

// PopEnvIfLocalBinds compiles a PopEnv statement if locals is true. also sets *nbinds and *nintbinds
func (inner *Comp) PopEnvIfLocalBinds(locals bool, nbinds *[2]int, list ...ast.Stmt) *Comp {
	c := inner.Outer
	c.Code = inner.Code       // copy back accumulated code
	nbinds[0] = inner.BindNum // we finally know these
	nbinds[1] = inner.IntBindNum

	if locals != (inner.BindNum != 0 || inner.IntBindNum != 0) {
		c.Errorf(`internal error: containLocalBinds() returned %t, but block actually defined %d Binds and %d IntBinds:
	Binds = %v
	Block = %v
`, locals, inner.BindNum, inner.IntBindNum, inner.Binds, list)
		return nil
	}

	if locals {
		// pop *Env at runtime
		c.Code.Append(PopEnv)
	}
	return c
}

// If compiles an "if" statement
func (c *Comp) If(node *ast.IfStmt) {
	var ithen, ielse, iend int

	initLocals := false
	var initBinds [2]int
	if node.Init != nil {
		c, initLocals = c.PushEnvIfLocalBinds(&initBinds, node.Init)
		c.Stmt(node.Init)
	}
	pred := c.Expr(node.Cond)
	flag, fun, err := pred.TryAsPred()
	if err {
		c.invalidPred(node.Cond, pred)
		return
	}
	if fun != nil {
		c.Code.Append(func(env *Env) (Stmt, *Env) {
			var ip int
			if fun(env) {
				ip = ithen
			} else {
				ip = ielse
			}
			env.IP = ip
			return env.Code[ip], env
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
		// parser should guarantee Else to be a block or another "if"
		// but macroexpansion can optimize away the block if it contains no declarations.
		// still, better be safe and wrap the Else again in a block because:
		// 1) catches improper macroexpander optimizations
		// 2) there is no runtime performance penalty
		xelse := node.Else
		_, ok1 := xelse.(*ast.BlockStmt)
		_, ok2 := xelse.(*ast.IfStmt)
		if ok1 || ok2 {
			c.Stmt(xelse)
		} else {
			c.Block(&ast.BlockStmt{List: []ast.Stmt{xelse}})
		}
		if fun == nil && flag {
			// 'else' branch is never executed...
			// still compiled above (to check for errors) but drop the generated code
			c.Code.List = c.Code.List[0:ielse]
		}
	}
	iend = c.Code.Len()

	if node.Init != nil {
		c = c.PopEnvIfLocalBinds(initLocals, &initBinds)
	}
}

// For compiles a "for" statement
func (c *Comp) For(node *ast.ForStmt) {
	initLocals := false
	var initBinds [2]int
	if node.Init != nil {
		c, initLocals = c.PushEnvIfLocalBinds(&initBinds, node.Init)
		c.Stmt(node.Init)
	}
	flag, fun, err := true, (func(*Env) bool)(nil), false // "for { }" without a condition means "for true { }"
	if node.Cond != nil {
		pred := c.Expr(node.Cond)
		flag, fun, err = pred.TryAsPred()
		if err {
			c.invalidPred(node.Cond, pred)
			return
		}
	}
	var ibody, ibreak int
	// compile the condition, if not a constant
	icond := c.Code.Len()
	if fun != nil {
		c.Code.Append(func(env *Env) (Stmt, *Env) {
			var ip int
			if fun(env) {
				ip = ibody
				// Debugf("for: condition = true, iterating. IntBinds = %v", env.IntBinds)
			} else {
				// Debugf("for: condition = false, exiting. IntBinds = %v", env.IntBinds)
				ip = ibreak
			}
			env.IP = ip
			return env.Code[ip], env
		})
	}
	// compile the body
	// TODO implement break and continue
	ibody = c.Code.Len()
	c.Block(node.Body)
	if fun == nil && !flag {
		// body is never executed...
		// still compiled above (to check for errors) but drop the generated code
		c.Code.List = c.Code.List[0:ibody]
	}
	// compile the post. trust it's not a declaration...
	if node.Post != nil {
		c.Stmt(node.Post)
	}
	c.Code.Append(func(env *Env) (Stmt, *Env) {
		// jump back to the condition
		// Debugf("for: body executed, jumping back to condition. IntBinds = %v", env.IntBinds)
		// time.Sleep(time.Second / 10)
		env.IP = icond
		return env.Code[icond], env
	})
	if fun == nil && !flag {
		// "for false { }" means that body, post and jump back to condition are never executed...
		// still compiled above (to check for errors) but drop the generated code
		c.Code.List = c.Code.List[0:icond]
	}
	ibreak = c.Code.Len()
	if node.Init != nil {
		c = c.PopEnvIfLocalBinds(initLocals, &initBinds, node.Init)
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
