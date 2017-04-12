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
 * declaration.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
	"go/token"
)

// Assign compiles an *ast.AssignStmt into an assignment to one or more place
func (c *Comp) Assign(node *ast.AssignStmt) {
	lhs, rhs := node.Lhs, node.Rhs
	ln, rn := len(lhs), len(rhs)
	if ln > 1 && rn == 1 {
		c.Errorf("unimplemented: assignment of multiple places with a single multi-valued expression: %v", node)
	} else if ln != rn {
		c.Errorf("invalid assignment, cannot assign %d values to %d places: %v", node)
	} else {
		for i, l := range lhs {
			c.Assign1(l, node.Tok, rhs[i])
		}
	}
}

// Assign1 compiles a single assignment to a place
func (c *Comp) Assign1(lhs ast.Expr, op token.Token, rhs ast.Expr) {
	place := c.Place(lhs)
	init := c.Expr(rhs)
	if init.Const() {
		switch op {
		case token.ASSIGN:
			c.placeSetConst(place, init)
		case token.ADD, token.ADD_ASSIGN:
			c.placeAddConst(place, init)
		case token.SUB, token.SUB_ASSIGN:
			c.placeSubConst(place, init)
		default:
			c.Errorf("unimplemented assignment operator '%v' : %v %v %v", op, lhs, op, rhs)
		}
	} else {
		switch op {
		case token.ASSIGN:
			c.placeSetExpr(place, init)
		case token.ADD, token.ADD_ASSIGN:
			c.placeAddExpr(place, init)
		default:
			c.Errorf("unimplemented assignment operator '%v' : %v %v %v", op, lhs, op, rhs)
		}
	}
}

// AssignVar0 compiles an assignment to a variable.
// not used by gomacro, provided as convenience for applications
func (c *Comp) AssignVar0(name string, init *Expr) {
	var_ := c.Var(name)
	if init.Const() {
		c.placeSetConst(&Place{Var: *var_}, init)
	} else {
		c.placeSetExpr(&Place{Var: *var_}, init)
	}
}

// Place compiles the left-hand-side of an assignment
func (c *Comp) Place(lhs ast.Expr) *Place {
	return c.PlaceOrAddress(lhs, false)
}

// PlaceOrAddress compiles the left-hand-side of an assignment or the location of an address-of
func (c *Comp) PlaceOrAddress(lhs ast.Expr, addressof bool) *Place {
	switch lhs := lhs.(type) {
	case *ast.Ident:
		name := lhs.Name
		if name == "_" {
			if addressof {
				c.Errorf("cannot take the address of _")
				return nil
			}
			return &Place{}
		}
		upn, bind := c.Resolve(name)
		class := bind.Desc.Class()
		if class != VarBind && class != IntBind {
			if addressof {
				c.Errorf("cannot take the address of %s: %v", class, name)
			} else {
				c.Errorf("cannot assign to %s: %v", class, name)
			}
			return nil
		}
		return &Place{Var: Var{Upn: upn, Desc: bind.Desc, Type: bind.Type}}
	default:
		if addressof {
			c.Errorf("unimplemented: address of non-identifier: %v", lhs)
		} else {
			c.Errorf("unimplemented: assignment of non-identifier: %v", lhs)
		}
		return nil
	}
}

// Var compiles the left-hand-side of an assignment, in case it's an identifier (i.e. a variable name)
func (c *Comp) Var(name string) *Var {
	if name == "_" {
		return &Var{}
	}
	upn, bind := c.Resolve(name)
	class := bind.Desc.Class()
	if class != VarBind && class != IntBind {
		c.Errorf("cannot assign to %s: %v", class, name)
	}
	return &Var{Upn: upn, Desc: bind.Desc, Type: bind.Type}
}
