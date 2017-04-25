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
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) placeAddConst(place *Place, val I) {
	if isLiteralNumber(val, 0) || val == "" {
		c.placeForSideEffects(place)
		return
	}

	{
		if place.MapKey != nil {
			c.Errorf("unimplemented: map[key] %s= constant", token.ADD,
			)
		}

		lhsfun := place.Func()
		var ret Stmt
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() + int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() + int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() + int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() + int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() + int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() + uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() + uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() + uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() + uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() + uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() + uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() + float64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() + float64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() + complex128(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() + complex128(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case string:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetString(lhs.String() +
					val,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.ADD, place.Type)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) placeAddExpr(place *Place, fun I) {
	if place.MapKey != nil {
		c.Errorf("unimplemented: map[key] %s= expression", token.ADD,
		)
	}

	lhsfun := place.Func()
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() + int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() + int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() + int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() + int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() + int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() + uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() + uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() + uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() + uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() + uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uintptr:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() + uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetFloat(lhs.Float() + float64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetFloat(lhs.Float() + float64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetComplex(lhs.Complex() + complex128(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex128:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetComplex(lhs.Complex() + complex128(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) string:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetString(lhs.String() +
				fun(env),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.ADD, place.Type, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) placeSubConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {
		c.placeForSideEffects(place)
		return
	}

	{
		if place.MapKey != nil {
			c.Errorf("unimplemented: map[key] %s= constant", token.SUB,
			)
		}

		lhsfun := place.Func()
		var ret Stmt
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() - int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() - int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() - int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() - int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() - int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() - uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() - uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() - uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() - uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() - uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() - uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() - float64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() - float64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() - complex128(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() - complex128(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.SUB, place.Type)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) placeSubExpr(place *Place, fun I) {
	if place.MapKey != nil {
		c.Errorf("unimplemented: map[key] %s= expression", token.SUB,
		)
	}

	lhsfun := place.Func()
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() - int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() - int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() - int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() - int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() - int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() - uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() - uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() - uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() - uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() - uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uintptr:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() - uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetFloat(lhs.Float() - float64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetFloat(lhs.Float() - float64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetComplex(lhs.Complex() - complex128(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex128:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetComplex(lhs.Complex() - complex128(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SUB, place.Type, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) placeMulConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {

		c.placeSetZero(place)
		return
	} else if isLiteralNumber(val, 1) {
		c.placeForSideEffects(place)
		return
	}

	{
		if place.MapKey != nil {
			c.Errorf("unimplemented: map[key] %s= constant", token.MUL,
			)
		}

		lhsfun := place.Func()
		var ret Stmt
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() * int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() * int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() * int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() * int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() * int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() * uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() * uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() * uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() * uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() * uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() * uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() * float64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() * float64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() * complex128(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() * complex128(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.MUL, place.Type)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) placeMulExpr(place *Place, fun I) {
	if place.MapKey != nil {
		c.Errorf("unimplemented: map[key] %s= expression", token.MUL,
		)
	}

	lhsfun := place.Func()
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() * int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() * int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() * int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() * int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() * int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() * uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() * uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() * uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() * uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() * uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uintptr:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() * uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetFloat(lhs.Float() * float64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetFloat(lhs.Float() * float64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetComplex(lhs.Complex() * complex128(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex128:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetComplex(lhs.Complex() * complex128(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.MUL, place.Type, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) placeQuoConst(place *Place, val I) {
	if isLiteralNumber(val, 0) {
		c.Errorf("division by %v <%v>", val, place.Type)
		return
	} else if isLiteralNumber(val, 1) {
		c.placeForSideEffects(place)
		return
	}

	{
		if place.MapKey != nil {
			c.Errorf("unimplemented: map[key] %s= constant", token.QUO,
			)
		}

		lhsfun := place.Func()
		var ret Stmt
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() / int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() / int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() / int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() / int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() / int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() / uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() / uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() / uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() / uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() / uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() / uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() / float64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			float64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetFloat(lhs.Float() / float64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() / complex128(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			complex128:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetComplex(lhs.Complex() / complex128(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.QUO, place.Type)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) placeQuoExpr(place *Place, fun I) {
	if place.MapKey != nil {
		c.Errorf("unimplemented: map[key] %s= expression", token.QUO,
		)
	}

	lhsfun := place.Func()
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() / int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() / int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() / int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() / int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() / int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() / uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() / uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() / uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() / uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() / uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uintptr:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() / uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetFloat(lhs.Float() / float64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) float64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetFloat(lhs.Float() / float64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetComplex(lhs.Complex() / complex128(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) complex128:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetComplex(lhs.Complex() / complex128(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.QUO, place.Type, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) placeRemConst(place *Place, val I) {
	if IsCategory(place.Type.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, 0) {
			c.Errorf("division by %v <%v>", val, place.Type)
			return
		} else if isLiteralNumber(val, 1) {

			c.placeSetZero(place)
			return
		}
	}
	{
		if place.MapKey != nil {
			c.Errorf("unimplemented: map[key] %s= constant", token.REM,
			)
		}

		lhsfun := place.Func()
		var ret Stmt
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() % int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() % int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() % int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() % int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() % int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() % uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() % uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() % uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() % uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() % uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() % uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.REM, place.Type)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) placeRemExpr(place *Place, fun I) {
	if place.MapKey != nil {
		c.Errorf("unimplemented: map[key] %s= expression", token.REM,
		)
	}

	lhsfun := place.Func()
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() % int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() % int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() % int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() % int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() % int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() % uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() % uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() % uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() % uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() % uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uintptr:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() % uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.REM, place.Type, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) placeAndConst(place *Place, val I) {
	if IsCategory(place.Type.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, -1) {
			c.placeForSideEffects(place)
			return
		} else if isLiteralNumber(val, 0) {

			c.placeSetZero(place)
			return
		}
	}
	{
		if place.MapKey != nil {
			c.Errorf("unimplemented: map[key] %s= constant", token.AND,
			)
		}

		lhsfun := place.Func()
		var ret Stmt
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() & int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() & int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() & int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() & int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() & int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() & uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() & uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() & uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() & uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() & uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() & uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.AND, place.Type)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) placeAndExpr(place *Place, fun I) {
	if place.MapKey != nil {
		c.Errorf("unimplemented: map[key] %s= expression", token.AND,
		)
	}

	lhsfun := place.Func()
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() & int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() & int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() & int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() & int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() & int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() & uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() & uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() & uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() & uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() & uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uintptr:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() & uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND, place.Type, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) placeOrConst(place *Place, val I) {
	if IsCategory(place.Type.Kind(), r.Int, r.Uint) && isLiteralNumber(val, 0) {
		c.placeForSideEffects(place)
		return
	}

	{
		if place.MapKey != nil {
			c.Errorf("unimplemented: map[key] %s= constant", token.OR,
			)
		}

		lhsfun := place.Func()
		var ret Stmt
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() | int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() | int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() | int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() | int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() | int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() | uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() | uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() | uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() | uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() | uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() | uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.OR, place.Type)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) placeOrExpr(place *Place, fun I) {
	if place.MapKey != nil {
		c.Errorf("unimplemented: map[key] %s= expression", token.OR,
		)
	}

	lhsfun := place.Func()
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() | int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() | int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() | int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() | int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() | int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() | uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() | uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() | uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() | uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() | uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uintptr:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() | uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.OR, place.Type, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) placeXorConst(place *Place, val I) {
	if IsCategory(place.Type.Kind(), r.Int, r.Uint) && isLiteralNumber(val, 0) {
		c.placeForSideEffects(place)
		return
	}

	{
		if place.MapKey != nil {
			c.Errorf("unimplemented: map[key] %s= constant", token.XOR,
			)
		}

		lhsfun := place.Func()
		var ret Stmt
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() ^ int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() ^ int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() ^ int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() ^ int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() ^ int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() ^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.XOR, place.Type)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) placeXorExpr(place *Place, fun I) {
	if place.MapKey != nil {
		c.Errorf("unimplemented: map[key] %s= expression", token.XOR,
		)
	}

	lhsfun := place.Func()
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() ^ int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() ^ int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() ^ int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() ^ int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() ^ int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() ^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() ^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() ^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() ^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() ^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uintptr:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() ^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, place.Type, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) placeAndnotConst(place *Place, val I) {
	if IsCategory(place.Type.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, -1) {

			c.placeSetZero(place)
			return
		} else if isLiteralNumber(val, 0) {
			c.placeForSideEffects(place)
			return
		}
	}
	{
		if place.MapKey != nil {
			c.Errorf("unimplemented: map[key] %s= constant", token.AND_NOT,
			)
		}

		lhsfun := place.Func()
		var ret Stmt
		switch val := val.(type) {
		case int:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() &^ int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() &^ int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() &^ int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() &^ int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case int64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetInt(lhs.Int() &^ int64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint8:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint16:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint32:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uint64:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		case

			uintptr:

			ret = func(env *Env) (Stmt, *Env) {
				lhs := lhsfun(env)
				lhs.SetUint(lhs.Uint() &^ uint64(val,
				),
				)

				env.IP++
				return env.Code[env.IP], env
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.AND_NOT, place.Type)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) placeAndnotExpr(place *Place, fun I) {
	if place.MapKey != nil {
		c.Errorf("unimplemented: map[key] %s= expression", token.AND_NOT,
		)
	}

	lhsfun := place.Func()
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() &^ int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() &^ int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() &^ int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() &^ int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) int64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetInt(lhs.Int() &^ int64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() &^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint8:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() &^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint16:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() &^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint32:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() &^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uint64:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() &^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	case func(*Env) uintptr:

		ret = func(env *Env) (Stmt, *Env) {
			lhs := lhsfun(env)
			lhs.SetUint(lhs.Uint() &^ uint64(fun(env),
			),
			)

			env.IP++
			return env.Code[env.IP], env
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND_NOT, place.Type, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) SetPlace(place *Place, op token.Token, init *Expr) {
	t := place.Type
	if init.Const() {
		init.ConstTo(t)
	} else if !init.Type.AssignableTo(t) {
		c.Errorf("incompatible types in assignment: <%v> %s <%v>", t, op, init.Type)
		return
	}
	if init.Const() {
		val := init.Value
		v := r.ValueOf(val)
		if v == None || v == Nil {
			v = r.Zero(t)
			val = v.Interface()
		} else if v.Type() != t {
			v = v.Convert(t)
			val = v.Interface()
		}
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
