// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

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
 * place_setops.go
 *
 *  Created on Apr 25, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
	_ "reflect"
	_ "unsafe"

	_ "github.com/cosmos72/gomacro/base"
)

func (c *Comp) placeAddConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {
		return
	}
	c.Errorf("unimplemented: place += constant")

}
func (c *Comp) placeAddExpr(place *Place, fun I) { c.Errorf("unimplemented: place += expression") }
func (c *Comp) placeSubConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {
		return
	}
	c.Errorf("unimplemented: place -= constant")

}
func (c *Comp) placeSubExpr(place *Place, fun I) { c.Errorf("unimplemented: place -= expression") }
func (c *Comp) placeMulConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {

		c.placeSetZero(place)
		return
	} else if isLiteralNumber(val, 1) {
		return
	}
	c.Errorf("unimplemented: place *= constant")

}
func (c *Comp) placeMulExpr(place *Place, fun I) { c.Errorf("unimplemented: place *= expression") }
func (c *Comp) placeQuoConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {
		c.Errorf("division by %v <%v>", val, place.Type)
		return
	} else if isLiteralNumber(val, 1) {
		return
	}
	c.Errorf("unimplemented: place /= constant")

}
func (c *Comp) placeQuoExpr(place *Place, fun I) { c.Errorf("unimplemented: place /= expression") }
func (c *Comp) placeRemConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {
		c.Errorf("division by %v <%v>", val, place.Type)
		return
	} else if isLiteralNumber(val, 1) {

		c.placeSetZero(place)
		return
	}
	c.Errorf("unimplemented: place %= constant")

}
func (c *Comp) placeRemExpr(place *Place, fun I) { c.Errorf("unimplemented: place %= expression") }
func (c *Comp) placeAndConst(place *Place, val I) {
	if isLiteralNumber(val, -1) {
		return
	} else if isLiteralNumber(val, 0) {

		c.placeSetZero(place)
		return
	}
	c.Errorf("unimplemented: place &= constant")

}
func (c *Comp) placeAndExpr(place *Place, fun I) { c.Errorf("unimplemented: place &= expression") }
func (c *Comp) placeOrConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {
		return
	}
	c.Errorf("unimplemented: place |= constant")

}
func (c *Comp) placeOrExpr(place *Place, fun I) { c.Errorf("unimplemented: place |= expression") }
func (c *Comp) placeXorConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {
		return
	}
	c.Errorf("unimplemented: place ^= constant")

}
func (c *Comp) placeXorExpr(place *Place, fun I) { c.Errorf("unimplemented: place ^= expression") }
func (c *Comp) placeAndnotConst(place *Place, val I) {
	if isLiteralNumber(val, -1) {

		c.placeSetZero(place)
		return
	} else if isLiteralNumber(val, 0) {
		return
	}
	c.Errorf("unimplemented: place &^= constant")

}
func (c *Comp) placeAndnotExpr(place *Place, fun I) { c.Errorf("unimplemented: place &^= expression") }
func (c *Comp) SetPlace(place *Place, op token.Token, init *Expr) {
	t := place.Type
	if init.Const() {
		init.ConstTo(t)
	} else if init.Type != t {
		if t.Kind() != init.Type.Kind() || !init.Type.AssignableTo(t) {
			c.Errorf("incompatible types in assignment: <%v> %s <%v>", t, op, init.Type)
			return
		}
	}
	if init.Const() {
		val := init.Value
		switch op {
		case token.ASSIGN:
			c.placeSetConst(place, val)
		case token.ADD, token.ADD_ASSIGN:
			c.placeAddConst(place, val)
		case token.SUB, token.SUB_ASSIGN:
			c.placeSubConst(place, val)
		case token.MUL, token.MUL_ASSIGN:
			c.placeMulConst(place, val)
		case token.QUO, token.QUO_ASSIGN:
			c.placeQuoConst(place, val)
		case token.REM, token.REM_ASSIGN:
			c.placeRemConst(place, val)
		case token.AND, token.AND_ASSIGN:
			c.placeAndConst(place, val)
		case token.OR, token.OR_ASSIGN:
			c.placeOrConst(place, val)
		case token.XOR, token.XOR_ASSIGN:
			c.placeAndConst(place, val)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			c.placeAndnotConst(place, val)
		default:
			c.Errorf("operator %s is not implemented", op)
		}
	} else {
		fun := init.Fun
		switch op {
		case token.ASSIGN:
			c.placeSetExpr(place, fun)
		case token.ADD, token.ADD_ASSIGN:
			c.placeAddExpr(place, fun)
		case token.SUB, token.SUB_ASSIGN:
			c.placeSubExpr(place, fun)
		case token.MUL, token.MUL_ASSIGN:
			c.placeMulExpr(place, fun)
		case token.QUO, token.QUO_ASSIGN:
			c.placeQuoExpr(place, fun)
		case token.REM, token.REM_ASSIGN:
			c.placeRemExpr(place, fun)
		case token.AND, token.AND_ASSIGN:
			c.placeAndExpr(place, fun)
		case token.OR, token.OR_ASSIGN:
			c.placeOrExpr(place, fun)
		case token.XOR, token.XOR_ASSIGN:
			c.placeAndExpr(place, fun)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			c.placeAndnotExpr(place, fun)
		default:
			c.Errorf("operator %s is not implemented", op)
		}
	}
}
