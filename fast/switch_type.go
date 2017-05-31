/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 *
 * switch.go
 *
 *  Created on May 06, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	"go/types"
	r "reflect"
	"sort"

	. "github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/typeutil"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type typecaseEntry struct {
	Type xr.Type
	Pos  token.Pos
	IP   int
}

type typecaseHelper struct {
	Map typeutil.Map // map types.Type -> typecaseEntry
}

// keep track of types in type-switch. error on duplicates
func (seen *typecaseHelper) add(c *Comp, t xr.Type, entry typecaseEntry) {
	var gtype types.Type
	if t != nil {
		gtype = t.GoType()
	}
	prev := seen.Map.At(gtype)
	if prev != nil {
		c.Errorf("duplicate case <%v> in switch\n\tprevious case at %s", t, c.Fileset.Position(prev.(typecaseEntry).Pos))
		return
	}
	entry.Type = t
	seen.Map.Set(gtype, entry)
}

/*
func (c *Comp) TypeSwitch(node *ast.TypeSwitchStmt, labels []string) {
	c.Errorf("unimplemented statement: %v <%v>", node, r.TypeOf(node))
}
*/

func (c *Comp) TypeSwitch(node *ast.TypeSwitchStmt, labels []string) {
	initLocals := false
	var initBinds [2]int
	c, initLocals = c.pushEnvIfLocalBinds(&initBinds, node.Init, node.Assign)
	if node.Init != nil {
		c.Stmt(node.Init)
	}
	var ibreak int
	sort.Strings(labels)
	c.Loop = &LoopInfo{
		Break:      &ibreak,
		ThisLabels: labels,
	}

	tagstmt := node.Assign
	var varname string // empty, or name of variable in 'switch varname := expression.(type)'
	var tagnode ast.Expr
	switch tagstmt := tagstmt.(type) {
	case *ast.AssignStmt:
		if len(tagstmt.Lhs) == 1 && len(tagstmt.Rhs) == 1 && tagstmt.Tok == token.DEFINE {
			if lhs, ok := tagstmt.Lhs[0].(*ast.Ident); ok {
				varname = lhs.Name
				tagnode = tagstmt.Rhs[0]
			}
		}
	case *ast.ExprStmt:
		expr := tagstmt.X
		for {
			switch e := expr.(type) {
			case *ast.ParenExpr:
				expr = e.X
				continue
			case *ast.TypeAssertExpr:
				if e.Type != nil {
					c.Errorf("invalid type switch: expecting '.(type)', found type assertion: %v", node)
				}
				tagnode = e.X
			default:
				tagnode = e
			}
			break
		}
	}
	if tagnode == nil {
		c.Errorf("expected type-switch expression, found: %v", node)
	} else {
		c.Debugf("type-switch on %v <%v>", tagnode, r.TypeOf(tagnode))
	}
	tag := c.Expr1(tagnode)

	if tag.Type.Kind() != r.Interface {
		c.Errorf("cannot type switch on non-interface type <%v>: %v", tag.Type, tagnode)
	}
	if !tag.Const() {
		// just like Comp.Switch, we cannot invoke tag.Fun() multiple times because
		// side effects must be applied only once!
		// typeswitchTag saves the result of tag.Fun() in a runtime bind
		// and returns an expression that retrieves it, with its concrete type
		tag = c.typeswitchTag(tag)
	}

	if node.Body != nil {
		// reserve a code slot for typeSwitchGotoMap optimizer
		icasemap := c.Code.Len()
		seen := &typecaseHelper{} // keeps track of types in cases. errors on duplicates
		c.Append(stmtNop, node.Body.Pos())

		list := node.Body.List
		defaulti := -1
		var defaultpos token.Pos
		for _, stmt := range list {
			switch clause := stmt.(type) {
			case *ast.CaseClause:
				if clause.List == nil {
					if defaulti >= 0 {
						c.Errorf("multiple defaults in switch (first at %s)", c.Fileset.Position(defaultpos))
					}
					defaulti = c.Code.Len()
					defaultpos = clause.Pos()
					c.typeswitchDefault(clause, tagnode, tag, varname)
				} else {
					c.typeswitchCase(clause, tagnode, tag, varname, seen)
				}
			default:
				c.Errorf("invalid statement inside switch: expecting case or default, found: %v <%v>", stmt, r.TypeOf(stmt))
			}
		}
		// default is executed as last, if no other case matches
		if defaulti >= 0 {
			// +1 to skip its "never matches" header
			c.Append(func(env *Env) (Stmt, *Env) {
				ip := defaulti + 1
				env.IP = ip
				return env.Code[ip], env
			}, defaultpos)
		}
		// try to optimize
		c.typeswitchGotoMap(tag, seen, icasemap)
	}
	// we finally know this
	ibreak = c.Code.Len()

	c = c.popEnvIfLocalBinds(initLocals, &initBinds, node.Init)
}

// typeswitchTag takes the expression immediately following a type-switch,
// compiles it to a statement that evaluates it and saves its result and its type
// in two runtime bindings (interpreter local variables),
// finally returns another expression that retrieves the expression value
// with its concrete type
func (c *Comp) typeswitchTag(e *Expr) *Expr {
	var upn, o = 0, c
	// try to piggyback the binding to a Comp that already has some bindings,
	// but do not cross function boundaries
	for o.BindNum == 0 && o.IntBindNum == 0 && o.Func == nil && o.Outer != nil {
		upn += o.UpCost
		o = o.Outer
	}
	bind := o.AddBind("", VarBind, c.TypeOfInterface())

	// c.Debugf("switchTag: allocated bind %v, upn = %d", bind, upn)
	switch bind.Desc.Class() {
	case VarBind:
		// cannot use c.DeclVar0 because the variable is declared in o
		// cannot use o.DeclVar0 because the initializer must be evaluated in c
		// so initialize the binding manually
		index := bind.Desc.Index()
		init := e.AsX1()
		switch upn {
		case 0:
			c.append(func(env *Env) (Stmt, *Env) {
				v := r.ValueOf(init(env).Interface()) // extract concrete type
				// no need to create a settable reflect.Value
				env.Binds[index] = v
				env.IP++
				return env.Code[env.IP], env
			})
		case 1:
			c.append(func(env *Env) (Stmt, *Env) {
				v := r.ValueOf(init(env).Interface()) // extract concrete type
				// no need to create a settable reflect.Value
				env.Outer.Binds[index] = v
				env.IP++
				return env.Code[env.IP], env
			})
		default:
			c.append(func(env *Env) (Stmt, *Env) {
				o := env
				for i := 0; i < upn; i++ {
					o = o.Outer
				}
				v := r.ValueOf(init(env).Interface()) // extract concrete type
				// no need to create a settable reflect.Value
				o.Binds[index] = v
				env.IP++
				return env.Code[env.IP], env
			})
		}
	default:
		c.Errorf("internal error! Comp.AddBind(name=%q, class=VarBind, type=%v) returned class=%v, expecting VarBind",
			"", bind.Type, bind.Desc.Class())
		return nil
	}
	sym := bind.AsSymbol(upn)
	return c.Symbol(sym)
}

// typeswitchGotoMap tries to optimize the dispatching of a type-switch
func (c *Comp) typeswitchGotoMap(tag *Expr, seen *typecaseHelper, ip int) {
	fun := tag.AsX1()
	m := make(map[r.Type]int) // FIXME this breaks on types declared in the interpreter
	seen.Map.Iterate(func(k types.Type, v interface{}) {
		// useless to put interface types here
		entry := v.(typecaseEntry)
		if t := entry.Type; t.Kind() != r.Interface {
			m[t.ReflectType()] = entry.IP
		}
	})
	if len(m) == 0 {
		return
	}

	stmt := func(env *Env) (Stmt, *Env) {
		// FIXME this breaks on types declared in the interpreter
		rtype := fun(env).Type() // concrete reflect.Type already extracted by fun(env)
		if ip, found := m[rtype]; found {
			env.IP = ip
		} else {
			env.IP++
		}
		return env.Code[env.IP], env
	}
	c.Code.List[ip] = stmt
}

// typeswitchCase compiles a case in a type-switch.
func (c *Comp) typeswitchCase(node *ast.CaseClause, tagnode ast.Expr, tag *Expr, varname string, seen *typecaseHelper) {
	cmpfuns := make([]func(*Env) bool, len(node.List))
	cmpnode := &ast.BinaryExpr{Op: token.EQL, X: tagnode} // for error messages, and Comp.BinaryExpr1 dispatches on its Op

	ibody := c.Code.Len() + 1 // body will start here
	tagfun := tag.AsX1()
	// compile a comparison of tag against each type
	for i, enode := range node.List {
		t := c.compileTypeOrNil(enode)
		var rtype r.Type
		if t != nil {
			rtype = t.ReflectType()
		}
		cmpnode.OpPos = enode.Pos()
		cmpnode.Y = enode
		if rtype == nil {
			// in a type-switch, 'case nil:' means 'is the value a nil interface ?'
			cmpfuns[i] = func(env *Env) bool {
				v := tagfun(env)
				return v == Nil || v.Kind() == r.Interface && v.IsNil()
			}
		} else if rtype.Kind() == r.Interface {
			cmpfuns[i] = func(env *Env) bool {
				return tagfun(env).Type().Implements(rtype)
			}
		} else {
			cmpfuns[i] = func(env *Env) bool {
				return rtype == tagfun(env).Type()
			}
		}
		seen.add(c, t, typecaseEntry{Pos: enode.Pos(), IP: ibody})
	}
	// compile like "if r.TypeOf(tag) == t1 || r.TypeOf(tag) == t2 ... { }"
	// and keep track of where to jump if no expression matches
	//
	// always occupy a Code slot for cmpfuns, even if nothing to do.
	// reason: both caseMap optimizer and fallthrough from previous case
	// skip such slot and jump to current body
	var iend int
	var stmt Stmt
	switch len(cmpfuns) {
	case 0:
		// compile anyway. reachable?
		stmt = func(env *Env) (Stmt, *Env) {
			ip := iend
			env.IP = ip
			return env.Code[ip], env
		}
	case 1:
		cmpfun := cmpfuns[0]
		stmt = func(env *Env) (Stmt, *Env) {
			var ip int
			if cmpfun(env) {
				ip = env.IP + 1
			} else {
				ip = iend
			}
			env.IP = ip
			return env.Code[ip], env
		}
	case 2:
		cmpfuns := [...]func(*Env) bool{
			cmpfuns[0],
			cmpfuns[1],
		}
		stmt = func(env *Env) (Stmt, *Env) {
			var ip int
			if cmpfuns[0](env) || cmpfuns[1](env) {
				ip = env.IP + 1
			} else {
				ip = iend
			}
			env.IP = ip
			return env.Code[ip], env
		}
	default:
		stmt = func(env *Env) (Stmt, *Env) {
			ip := iend
			for _, cmpfun := range cmpfuns {
				if cmpfun(env) {
					ip = env.IP + 1
					break
				}
			}
			env.IP = ip
			return env.Code[ip], env
		}
	}
	c.Append(stmt, node.Pos())
	// TODO declare varname
	c.switchCaseBody(node.Body, false)
	// we finally know where to jump if match fails
	iend = c.Code.Len()
}

// typeswitchCase compiles the default case in a type-switch.
func (c *Comp) typeswitchDefault(node *ast.CaseClause, tagnode ast.Expr, tag *Expr, varname string) {
	// TODO declare varname
	c.switchCaseBody(node.Body, false)
}
