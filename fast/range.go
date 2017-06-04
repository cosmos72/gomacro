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
 * statement.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"
	"sort"

	xr "github.com/cosmos72/gomacro/xreflect"
)

type rangeJump struct {
	Start, Continue, Break int
}

// Range compiles a "for-range" statement
func (c *Comp) Range(node *ast.RangeStmt, labels []string) {
	var nbinds [2]int

	c, _ = c.pushEnvIfFlag(&nbinds, true)
	erange := c.Expr1(node.X)
	t := erange.Type
	var jump rangeJump

	sort.Strings(labels)
	// we need a fresh Comp here... created above by c.pushEnvIfLocalBinds()
	c.Loop = &LoopInfo{
		Continue:   &jump.Continue,
		Break:      &jump.Break,
		ThisLabels: labels,
	}

	switch t.Kind() {
	case r.Ptr:
		if t.Elem().Kind() != r.Array {
			c.Errorf("cannot range over %v <%v>", node.X, t)
		}
		// range on pointer to array: dereference it
		t = t.Elem()
		efun := erange.AsX1()
		erange = exprX1(t, func(env *Env) r.Value {
			return efun(env).Elem()
		})
		fallthrough
	case r.Array, r.Slice:
		c.rangeArray(node, erange, &jump)
	case r.Chan:
		c.rangeChan(node, erange, &jump)
	case r.Map:
		c.rangeMap(node, erange, &jump)
	case r.String:
		c.rangeString(node, erange, &jump)
	default:
		c.Errorf("cannot range over %v <%v>", node.X, t)
	}

	jump.Break = c.Code.Len()

	c = c.popEnvIfFlag(&nbinds, true)
}

func (c *Comp) rangeArray(node *ast.RangeStmt, erange *Expr, jump *rangeJump) {

	t := erange.Type
	var constlen int
	var elen *Expr

	if node.Value != nil || t.Kind() != r.Array {
		// Go spec: one-variable range on array ONLY evaluates the array length, not the array itself
		// save range variable in an unnamed bind
		sym := c.DeclVar0("", nil, erange).AsSymbol(0)
		erange = c.Symbol(sym)
	}

	if t.Kind() == r.Array {
		constlen = t.Len()
	} else {
		// save range length in an unnamed bind
		rangefun := erange.AsX1()
		elen0 := exprFun(c.TypeOfInt(), func(env *Env) int {
			return rangefun(env).Len()
		})
		symlen := c.DeclVar0("", nil, elen0).AsSymbol(0)
		elen = c.Symbol(symlen)
	}

	placekey, placeval := c.rangeVars(node, c.TypeOfInt(), t.Elem())

	jump.Start = c.Code.Len()

	// compile comparison against range length
	ekey := c.GetPlace(placekey)
	funkey := ekey.WithFun().(func(*Env) int)

	if t.Kind() == r.Array {
		c.append(func(env *Env) (Stmt, *Env) {
			var ip int
			if funkey(env) < constlen {
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
	} else {
		funlen := elen.WithFun().(func(*Env) int)
		c.append(func(env *Env) (Stmt, *Env) {
			var ip int
			if funkey(env) < funlen(env) {
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
	}
	if placeval != nil {
		// for error messages
		indexnode := &ast.IndexExpr{X: node.X, Lbrack: node.X.Pos(), Index: node.Key, Rbrack: node.X.Pos()}
		eindex := c.vectorIndex(indexnode, erange, ekey)
		c.SetPlace(placeval, token.ASSIGN, eindex)
	}

	// compile the body
	c.Block(node.Body)

	// increment key
	c.Pos = node.End() - 1
	one := c.exprValue(c.TypeOfInt(), 1)
	c.SetPlace(placekey, token.ADD_ASSIGN, one)

	// jump back to comparison
	c.append(func(env *Env) (Stmt, *Env) {
		ip := jump.Start
		env.IP = ip
		return env.Code[ip], env
	})
}

// rangeVars compiles the key and value iteration variables in a for-range
func (c *Comp) rangeVars(node *ast.RangeStmt, tkey xr.Type, tval xr.Type) (*Place, *Place) {
	var place [2]*Place
	t := [2]xr.Type{tkey, tval}

	for i, expr := range [2]ast.Expr{node.Key, node.Value} {
		if expr == nil {
			if i == 0 {
				// we need key bind for looping, even if user code ignores it
				place[i] = c.DeclVar0("", t[i], nil).AsVar(0, PlaceSettable).AsPlace()
			}
			continue
		}
		c.Pos = expr.Pos()
		if node.Tok == token.DEFINE {
			switch expr := expr.(type) {
			case *ast.Ident:
				name := expr.Name
				if i == 0 && name == "_" {
					// we need key bind for looping, even if user code ignores it
					name = ""
				}
				place[i] = c.DeclVar0(name, t[i], nil).AsVar(0, PlaceSettable).AsPlace()
			default:
				c.Errorf("non-name %v on left side of :=", expr)
			}
		} else {
			place[i] = c.Place(expr)
			if !t[i].AssignableTo(place[i].Type) {
				c.Errorf("cannot assign type <%v> to %v <%v> in range", t[i], expr, place[i].Type)
			}
		}
	}
	return place[0], place[1]
}

func (c *Comp) rangeChan(node *ast.RangeStmt, erange *Expr, jump *rangeJump) {
	c.Errorf("range on chan not implemented: range %v <%v>", node.X, erange.Type)
}

func (c *Comp) rangeMap(node *ast.RangeStmt, erange *Expr, jump *rangeJump) {
	c.Errorf("range on map not implemented: %v <%v>", node.X, erange.Type)
}

func (c *Comp) rangeString(node *ast.RangeStmt, erange *Expr, jump *rangeJump) {
	c.Errorf("range on string not implemented: %v <%v>", node.X, erange.Type)
}
