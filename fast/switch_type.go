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

	"github.com/cosmos72/gomacro/typeutil"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type typecaseEntry struct {
	Type xr.Type
	Pos  token.Pos
	IP   int
}

type typecaseHelper struct {
	TypeMap     typeutil.Map // map types.Type -> typecaseEntry
	ConcreteMap typeutil.Map // only contains the initial segment of non-interface types
	AllConcrete bool
}

// keep track of types in type-switch. error on duplicates
func (seen *typecaseHelper) add(c *Comp, t xr.Type, entry typecaseEntry) {
	var gtype types.Type
	if t != nil {
		gtype = t.GoType()
	}
	prev := seen.TypeMap.At(gtype)
	if prev != nil {
		c.Errorf("duplicate case <%v> in switch\n\tprevious case at %s", t, c.Fileset.Position(prev.(typecaseEntry).Pos))
		return
	}
	entry.Type = t
	seen.TypeMap.Set(gtype, entry)
	if t.Kind() == r.Interface {
		seen.AllConcrete = false
	} else if seen.AllConcrete {
		seen.ConcreteMap.Set(gtype, entry)
	}
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
		ipswitchgoto := c.Code.Len()
		seen := &typecaseHelper{AllConcrete: true} // keeps track of types in cases. errors on duplicates
		c.Append(stmtNop, node.Body.Pos())

		list := node.Body.List
		defaulti := -1
		var defaultpos token.Pos
		for _, stmt := range list {
			c.Pos = stmt.Pos()
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
		c.typeswitchGotoMap(tag, seen, ipswitchgoto)
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

	// c.Debugf("typeswitchTag: allocated bind %v, upn = %d", bind, upn)
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
				// Debugf("typeswitchTag = %v <%v>", v, ValueType(v))
				// no need to create a settable reflect.Value
				env.Binds[index] = v
				env.IP++
				return env.Code[env.IP], env
			})
		case 1:
			c.append(func(env *Env) (Stmt, *Env) {
				v := r.ValueOf(init(env).Interface()) // extract concrete type
				// Debugf("typeswitchTag = %v <%v>", v, ValueType(v))
				// no need to create a settable reflect.Value
				env.Outer.Binds[index] = v
				env.IP++
				return env.Code[env.IP], env
			})
		default:
			c.append(func(env *Env) (Stmt, *Env) {
				v := r.ValueOf(init(env).Interface()) // extract concrete type
				// Debugf("typeswitchTag = %v <%v>", v, ValueType(v))
				// no need to create a settable reflect.Value
				o := env.Outer.Outer
				for i := 2; i < upn; i++ {
					o = o.Outer
				}
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
	if seen.ConcreteMap.Len() <= 1 {
		return
	}
	m := make(map[r.Type]int) // FIXME this breaks on types declared in the interpreter
	seen.ConcreteMap.Iterate(func(k types.Type, v interface{}) {
		entry := v.(typecaseEntry)
		m[entry.Type.ReflectType()] = entry.IP
	})

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

	ibody := c.Code.Len() + 1 // body will start here
	tagfun := tag.AsX1()
	ts := make([]xr.Type, len(node.List))
	rtypes := make([]r.Type, len(node.List))

	// compile a comparison of tag against each type
	for i, enode := range node.List {
		t := c.compileTypeOrNil(enode)
		if t != nil {
			rtypes[i] = t.ReflectType()
		}
		ts[i] = t
		seen.add(c, t, typecaseEntry{Pos: enode.Pos(), IP: ibody})
	}
	// compile like "if r.TypeOf(tag) == t1 || r.TypeOf(tag) == t2 ... { }"
	// and keep track of where to jump if no expression matches
	//
	// always occupy a Code slot for type comparison, even if nothing to do.
	// reason: typeswitchGotoMap optimizer skips such slot and jumps to current body
	var iend int
	var stmt Stmt
	switch len(node.List) {
	case 0:
		// compile anyway. reachable?
		stmt = func(env *Env) (Stmt, *Env) {
			// Debugf("typeswitchCase: comparing %v against zero types", tagfun(env))
			ip := iend
			env.IP = ip
			return env.Code[ip], env
		}
	case 1:
		t := ts[0]
		rtype := rtypes[0]
		if t == nil {
			stmt = func(env *Env) (Stmt, *Env) {
				v := tagfun(env)
				// Debugf("typeswitchCase: comparing %v <%v> against nil type", v, ValueType(v))
				var ip int
				if !v.IsValid() {
					ip = env.IP + 1
				} else {
					ip = iend
				}
				env.IP = ip
				return env.Code[ip], env
			}
		} else if t.Kind() == r.Interface {
			stmt = func(env *Env) (Stmt, *Env) {
				v := tagfun(env)
				// Debugf("typeswitchCase: comparing %v <%v> against interface type %v", v, ValueType(v), rtype)
				var ip int
				if v.IsValid() && v.Type().Implements(rtype) {
					ip = env.IP + 1
				} else {
					ip = iend
				}
				env.IP = ip
				return env.Code[ip], env
			}
		} else {
			stmt = func(env *Env) (Stmt, *Env) {
				v := tagfun(env)
				// Debugf("typeswitchCase: comparing %v <%v> against concrete type %v", v, ValueType(v), rtype)
				var ip int
				if v.IsValid() && v.Type() == rtype {
					ip = env.IP + 1
				} else {
					ip = iend
				}
				env.IP = ip
				return env.Code[ip], env
			}
		}
	default:
		stmt = func(env *Env) (Stmt, *Env) {
			v := tagfun(env)
			var vt r.Type
			if v.IsValid() {
				vt = v.Type()
			}
			// Debugf("typeswitchCase: comparing %v <%v> against types %v", v, vt, rtypes)
			ip := iend
			for _, rtype := range rtypes {
				switch {
				case vt == rtype:
				case rtype != nil:
					if rtype.Kind() != r.Interface || !vt.Implements(rtype) {
						continue
					}
				default: // rtype == nil
					if v.IsValid() {
						continue
					}
				}
				// Debugf("typeswitchCase: v <%v> matches type %v", v, vt, rtype)
				ip = env.IP + 1
				break
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
	var iend int
	stmt := func(env *Env) (Stmt, *Env) {
		// Debugf("typeswitchDefault: default entered normally, skipping it")
		ip := iend
		env.IP = ip
		return env.Code[ip], env
	}
	c.append(stmt)
	// TODO declare varname
	c.switchCaseBody(node.Body, false)
	iend = c.Code.Len()
}
