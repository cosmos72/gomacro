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
 * binary_shly.go
 *
 *  Created on Apr 08, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

// define all combinations. quite tedious...

func shlYInt(x int, y I) I {
	switch y := y.(type) {
	case func(*Env) uint:
		return func(env *Env) int {
			return x << y(env)
		}
	case func(*Env) uint8:
		return func(env *Env) int {
			return x << y(env)
		}
	case func(*Env) uint16:
		return func(env *Env) int {
			return x << y(env)
		}
	case func(*Env) uint32:
		return func(env *Env) int {
			return x << y(env)
		}
	case func(*Env) uint64:
		return func(env *Env) int {
			return x << y(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) int {
			return x << y(env)
		}
	default:
		return nil
	}
}

func shlYInt8(x int8, y I) I {
	switch y := y.(type) {
	case func(*Env) uint:
		return func(env *Env) int8 {
			return x << y(env)
		}
	case func(*Env) uint8:
		return func(env *Env) int8 {
			return x << y(env)
		}
	case func(*Env) uint16:
		return func(env *Env) int8 {
			return x << y(env)
		}
	case func(*Env) uint32:
		return func(env *Env) int8 {
			return x << y(env)
		}
	case func(*Env) uint64:
		return func(env *Env) int8 {
			return x << y(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) int8 {
			return x << y(env)
		}
	default:
		return nil
	}
}

func shlYInt16(x int16, y I) I {
	switch y := y.(type) {
	case func(*Env) uint:
		return func(env *Env) int16 {
			return x << y(env)
		}
	case func(*Env) uint8:
		return func(env *Env) int16 {
			return x << y(env)
		}
	case func(*Env) uint16:
		return func(env *Env) int16 {
			return x << y(env)
		}
	case func(*Env) uint32:
		return func(env *Env) int16 {
			return x << y(env)
		}
	case func(*Env) uint64:
		return func(env *Env) int16 {
			return x << y(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) int16 {
			return x << y(env)
		}
	default:
		return nil
	}
}

func shlYInt32(x int32, y I) I {
	switch y := y.(type) {
	case func(*Env) uint:
		return func(env *Env) int32 {
			return x << y(env)
		}
	case func(*Env) uint8:
		return func(env *Env) int32 {
			return x << y(env)
		}
	case func(*Env) uint16:
		return func(env *Env) int32 {
			return x << y(env)
		}
	case func(*Env) uint32:
		return func(env *Env) int32 {
			return x << y(env)
		}
	case func(*Env) uint64:
		return func(env *Env) int32 {
			return x << y(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) int32 {
			return x << y(env)
		}
	default:
		return nil
	}
}

func shlYInt64(x int64, y I) I {
	switch y := y.(type) {
	case func(*Env) uint:
		return func(env *Env) int64 {
			return x << y(env)
		}
	case func(*Env) uint8:
		return func(env *Env) int64 {
			return x << y(env)
		}
	case func(*Env) uint16:
		return func(env *Env) int64 {
			return x << y(env)
		}
	case func(*Env) uint32:
		return func(env *Env) int64 {
			return x << y(env)
		}
	case func(*Env) uint64:
		return func(env *Env) int64 {
			return x << y(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) int64 {
			return x << y(env)
		}
	default:
		return nil
	}
}

func shlYUint(x uint, y I) I {
	switch y := y.(type) {
	case func(*Env) uint:
		return func(env *Env) uint {
			return x << y(env)
		}
	case func(*Env) uint8:
		return func(env *Env) uint {
			return x << y(env)
		}
	case func(*Env) uint16:
		return func(env *Env) uint {
			return x << y(env)
		}
	case func(*Env) uint32:
		return func(env *Env) uint {
			return x << y(env)
		}
	case func(*Env) uint64:
		return func(env *Env) uint {
			return x << y(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) uint {
			return x << y(env)
		}
	default:
		return nil
	}
}

func shlYUint8(x uint8, y I) I {
	switch y := y.(type) {
	case func(*Env) uint:
		return func(env *Env) uint8 {
			return x << y(env)
		}
	case func(*Env) uint8:
		return func(env *Env) uint8 {
			return x << y(env)
		}
	case func(*Env) uint16:
		return func(env *Env) uint8 {
			return x << y(env)
		}
	case func(*Env) uint32:
		return func(env *Env) uint8 {
			return x << y(env)
		}
	case func(*Env) uint64:
		return func(env *Env) uint8 {
			return x << y(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) uint8 {
			return x << y(env)
		}
	default:
		return nil
	}
}

func shlYUint16(x uint16, y I) I {
	switch y := y.(type) {
	case func(*Env) uint:
		return func(env *Env) uint16 {
			return x << y(env)
		}
	case func(*Env) uint8:
		return func(env *Env) uint16 {
			return x << y(env)
		}
	case func(*Env) uint16:
		return func(env *Env) uint16 {
			return x << y(env)
		}
	case func(*Env) uint32:
		return func(env *Env) uint16 {
			return x << y(env)
		}
	case func(*Env) uint64:
		return func(env *Env) uint16 {
			return x << y(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) uint16 {
			return x << y(env)
		}
	default:
		return nil
	}
}

func shlYUint32(x uint32, y I) I {
	switch y := y.(type) {
	case func(*Env) uint:
		return func(env *Env) uint32 {
			return x << y(env)
		}
	case func(*Env) uint8:
		return func(env *Env) uint32 {
			return x << y(env)
		}
	case func(*Env) uint16:
		return func(env *Env) uint32 {
			return x << y(env)
		}
	case func(*Env) uint32:
		return func(env *Env) uint32 {
			return x << y(env)
		}
	case func(*Env) uint64:
		return func(env *Env) uint32 {
			return x << y(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) uint32 {
			return x << y(env)
		}
	default:
		return nil
	}
}

func shlYUint64(x uint64, y I) I {
	switch y := y.(type) {
	case func(*Env) uint:
		return func(env *Env) uint64 {
			return x << y(env)
		}
	case func(*Env) uint8:
		return func(env *Env) uint64 {
			return x << y(env)
		}
	case func(*Env) uint16:
		return func(env *Env) uint64 {
			return x << y(env)
		}
	case func(*Env) uint32:
		return func(env *Env) uint64 {
			return x << y(env)
		}
	case func(*Env) uint64:
		return func(env *Env) uint64 {
			return x << y(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) uint64 {
			return x << y(env)
		}
	default:
		return nil
	}
}

func shlYUintptr(x uintptr, y I) I {
	switch y := y.(type) {
	case func(*Env) uint:
		return func(env *Env) uintptr {
			return x << y(env)
		}
	case func(*Env) uint8:
		return func(env *Env) uintptr {
			return x << y(env)
		}
	case func(*Env) uint16:
		return func(env *Env) uintptr {
			return x << y(env)
		}
	case func(*Env) uint32:
		return func(env *Env) uintptr {
			return x << y(env)
		}
	case func(*Env) uint64:
		return func(env *Env) uintptr {
			return x << y(env)
		}
	case func(*Env) uintptr:
		return func(env *Env) uintptr {
			return x << y(env)
		}
	default:
		return nil
	}
}
