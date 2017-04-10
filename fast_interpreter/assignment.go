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
	r "reflect"
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
	place := c.PlaceVar(name)
	if init.Const() {
		c.placeSetConst(place, init)
	} else {
		c.placeSetExpr(place, init)
	}
}

type Place struct {
	Upn  int
	Desc BindDescriptor
	Type r.Type
	// Fun is nil for identifiers. returns settable reflect.Value
	// (except for map[key], where it may be non-settable).
	// call it only once, it may have side effects!
	Fun func(*Env) r.Value
	// used only for map[key], returns key. call it only once, it may have side effects!
	MapKey func(*Env) r.Value
}

// Place compiles the left-hand-side of an assignment
func (c *Comp) Place(lhs ast.Expr) *Place {
	switch lhs := lhs.(type) {
	case *ast.Ident:
		name := lhs.Name
		if name == "_" {
			return &Place{}
		}
		upn, bind := c.Resolve(name)
		class := bind.Desc.Class()
		if class != VarBind && class != IntBind {
			c.Errorf("cannot assign to: %v", name)
		}
		return &Place{Upn: upn, Desc: bind.Desc, Type: bind.Type}
	default:
		c.Errorf("unimplemented: assignment of non-identifier: %v", lhs)
		return nil
	}
}

// PlaceVar compiles the left-hand-side of an assignment, in case it's an identifier (i.e. a variable name)
func (c *Comp) PlaceVar(name string) *Place {
	if name == "_" {
		return &Place{}
	}
	upn, bind := c.Resolve(name)
	class := bind.Desc.Class()
	if class != VarBind && class != IntBind {
		c.Errorf("cannot assign to %v: %v", class, name)
	}
	return &Place{Upn: upn, Desc: bind.Desc, Type: bind.Type}
}
