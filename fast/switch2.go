// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

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
 * switch2.go
 *
 *  Created on May 06, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) switchTag(e *Expr) *Expr {
	efun := e.Fun
	var cachefun I
	var stmt Stmt
	switch e.Type.Kind() {
	case r.Bool:
		{
			var val bool
			cachefun = func(*Env) bool {
				return val
			}

			efun := efun.(func(*Env) bool)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Int:
		{
			var val int
			cachefun = func(*Env) int {
				return val
			}

			efun := efun.(func(*Env) int)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Int8:
		{
			var val int8
			cachefun = func(*Env) int8 {
				return val
			}

			efun := efun.(func(*Env) int8)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Int16:
		{
			var val int16
			cachefun = func(*Env) int16 {
				return val
			}

			efun := efun.(func(*Env) int16)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Int32:
		{
			var val int32
			cachefun = func(*Env) int32 {
				return val
			}

			efun := efun.(func(*Env) int32)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Int64:
		{
			var val int64
			cachefun = func(*Env) int64 {
				return val
			}

			efun := efun.(func(*Env) int64)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Uint:
		{
			var val uint
			cachefun = func(*Env) uint {
				return val
			}

			efun := efun.(func(*Env) uint)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Uint8:
		{
			var val uint8
			cachefun = func(*Env) uint8 {
				return val
			}

			efun := efun.(func(*Env) uint8)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	case r.Uint16:
		{
			var val uint16
			cachefun = func(*Env) uint16 {
				return val
			}

			efun := efun.(func(*Env) uint16)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Uint32:
		{
			var val uint32
			cachefun = func(*Env) uint32 {
				return val
			}

			efun := efun.(func(*Env) uint32)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Uint64:
		{
			var val uint64
			cachefun = func(*Env) uint64 {
				return val
			}

			efun := efun.(func(*Env) uint64)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Uintptr:
		{
			var val uintptr
			cachefun = func(*Env) uintptr {
				return val
			}

			efun := efun.(func(*Env) uintptr)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Float32:
		{
			var val float32
			cachefun = func(*Env) float32 {
				return val
			}

			efun := efun.(func(*Env) float32)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Float64:
		{
			var val float64
			cachefun = func(*Env) float64 {
				return val
			}

			efun := efun.(func(*Env) float64)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Complex64:
		{
			var val complex64
			cachefun = func(*Env) complex64 {
				return val
			}

			efun := efun.(func(*Env) complex64)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.Complex128:
		{
			var val complex128
			cachefun = func(*Env) complex128 {
				return val
			}

			efun := efun.(func(*Env) complex128)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}

	case r.String:
		{
			var val string
			cachefun = func(*Env) string {
				return val
			}

			efun := efun.(func(*Env) string)
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}

	default:
		val := xr.Zero(e.Type)
		cachefun = func(*Env) r.Value { return val }

		if efun, ok := efun.(func(*Env) (r.Value, []r.Value)); ok {
			stmt = func(env *Env) (Stmt, *Env) {
				val, _ = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		} else {
			efun := e.AsX1()
			stmt = func(env *Env) (Stmt, *Env) {
				val = efun(env)
				env.IP++
				return env.Code[env.IP], env
			}
		}
	}
	c.Code.Append(stmt)
	return exprFun(e.Type, cachefun)
}
func (c *Comp) switchGotoMap(tag *Expr, seen *caseHelper, ip int) {
	if seen.SomeNonConst || len(seen.Map) == 0 {
		return
	}

	var stmt Stmt
	switch efun := tag.Fun.(type) {
	case func(*Env) bool:
		m := [2]int{-1, -1}
		for k, v := range seen.Map {
			if r.ValueOf(k).Bool() {
				m[1] = v.IP
			} else {
				m[0] = v.IP
			}
		}

		stmt = func(env *Env) (Stmt, *Env) {
			val := efun(env)
			var ip int
			if val {
				ip = m[1]
			} else {
				ip = m[0]
			}

			if ip >= 0 {
				env.IP = ip
			} else {
				env.IP++
			}
			return env.Code[env.IP], env
		}
	case func(*Env) int:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[int]int, len(seen.Map))
			for k, v := range seen.Map {
				m[int(r.ValueOf(k).Int())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) int8:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[int8]int, len(seen.Map))
			for k, v := range seen.Map {
				m[int8(r.ValueOf(k).Int())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) int16:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[int16]int, len(seen.Map))
			for k, v := range seen.Map {
				m[int16(r.ValueOf(k).Int())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) int32:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[int32]int, len(seen.Map))
			for k, v := range seen.Map {
				m[int32(r.ValueOf(k).Int())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) int64:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[int64]int, len(seen.Map))
			for k, v := range seen.Map {
				m[r.ValueOf(k).Int()] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uint:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uint]int, len(seen.Map))
			for k, v := range seen.Map {
				m[uint(r.ValueOf(k).Uint())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uint8:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uint8]int, len(seen.Map))
			for k, v := range seen.Map {
				m[uint8(r.ValueOf(k).Uint())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uint16:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uint16]int, len(seen.Map))
			for k, v := range seen.Map {
				m[uint16(r.ValueOf(k).Uint())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uint32:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uint32]int, len(seen.Map))
			for k, v := range seen.Map {
				m[uint32(r.ValueOf(k).Uint())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uint64:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uint64]int, len(seen.Map))
			for k, v := range seen.Map {
				m[r.ValueOf(k).Uint()] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) uintptr:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[uintptr]int, len(seen.Map))
			for k, v := range seen.Map {
				m[uintptr(r.ValueOf(k).Uint())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) float32:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[float32]int, len(seen.Map))
			for k, v := range seen.Map {
				m[float32(r.ValueOf(k).Float())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) float64:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[float64]int, len(seen.Map))
			for k, v := range seen.Map {
				m[r.ValueOf(k).Float()] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) complex64:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[complex64]int, len(seen.Map))
			for k, v := range seen.Map {
				m[complex64(r.ValueOf(k).Complex())] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) complex128:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[complex128]int, len(seen.Map))
			for k, v := range seen.Map {
				m[r.ValueOf(k).Complex()] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) string:
		{
			stmt = c.switchGotoSlice(tag, seen)
			if stmt != nil {
				break
			}

			m := make(map[string]int, len(seen.Map))
			for k, v := range seen.Map {
				m[r.ValueOf(k).String()] = v.IP
			}

			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				if ip, ok := m[val]; ok {
					env.IP = ip
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}

	case func(*Env) (r.Value, []r.Value):
		m := make(map[interface{}]int, len(seen.Map))
		for k, v := range seen.Map {
			m[k] = v.IP
		}

		stmt = func(env *Env) (Stmt, *Env) {
			v, _ := efun(env)
			if ip, ok := m[v.Interface()]; ok {
				env.IP = ip
			} else {
				env.IP++
			}
			return env.Code[env.IP], env
		}
	default:
		fun := tag.AsX1()
		m := make(map[interface{}]int, len(seen.Map))
		for k, v := range seen.Map {
			m[k] = v.IP
		}

		stmt = func(env *Env) (Stmt, *Env) {
			val := fun(env).Interface()
			if ip, ok := m[val]; ok {
				env.IP = ip
			} else {
				env.IP++
			}
			return env.Code[env.IP], env
		}
	}
	if stmt == nil {
		return
	}

	c.Code.List[ip] = stmt
}
func (c *Comp) switchGotoSlice(tag *Expr, seen *caseHelper) Stmt {
	var stmt Stmt
	switch efun := tag.Fun.(type) {
	case func(*Env) int:
		{
			var min, max int
			for k := range seen.Map {
				key := int(r.ValueOf(k).Int())
				min = key
				max = key
				break
			}
			for k := range seen.Map {
				key := int(r.ValueOf(k).Int())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.Map) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.Map) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.Map))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.Map {
				key := int(r.ValueOf(k).Int())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		{
			var min, max int8
			for k := range seen.Map {
				key := int8(r.ValueOf(k).Int())
				min = key
				max = key
				break
			}
			for k := range seen.Map {
				key := int8(r.ValueOf(k).Int())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.Map) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.Map) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.Map))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.Map {
				key := int8(r.ValueOf(k).Int())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		{
			var min, max int16
			for k := range seen.Map {
				key := int16(r.ValueOf(k).Int())
				min = key
				max = key
				break
			}
			for k := range seen.Map {
				key := int16(r.ValueOf(k).Int())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.Map) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.Map) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.Map))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.Map {
				key := int16(r.ValueOf(k).Int())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		{
			var min, max int32
			for k := range seen.Map {
				key :=
					int32(r.ValueOf(k).Int())
				min = key
				max = key
				break
			}
			for k := range seen.Map {
				key := int32(r.ValueOf(k).Int())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.Map) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.Map) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.Map))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.Map {
				key := int32(r.ValueOf(k).Int())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		{
			var min, max int64
			for k := range seen.Map {
				key := r.ValueOf(k).Int()
				min = key
				max = key
				break
			}
			for k := range seen.Map {
				key := r.ValueOf(k).Int()
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.Map) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.Map) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.Map))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.Map {
				key := r.ValueOf(k).Int()

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		{
			var min, max uint
			for k := range seen.Map {
				key :=

					uint(r.ValueOf(k).Uint())
				min = key
				max = key
				break
			}
			for k := range seen.Map {
				key := uint(r.ValueOf(k).Uint())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.Map) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.Map) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.Map))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.Map {
				key := uint(r.ValueOf(k).Uint())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		{
			var min, max uint8
			for k := range seen.Map {
				key :=

					uint8(r.ValueOf(k).Uint())
				min = key
				max = key
				break
			}
			for k := range seen.Map {
				key := uint8(r.ValueOf(k).Uint())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.Map) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.Map) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.Map))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.Map {
				key := uint8(r.ValueOf(k).Uint())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		{
			var min, max uint16
			for k := range seen.Map {
				key :=

					uint16(r.ValueOf(k).Uint())
				min = key
				max = key
				break
			}
			for k := range seen.Map {
				key := uint16(r.ValueOf(k).Uint())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.Map) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.Map) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.Map))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.Map {
				key := uint16(r.ValueOf(k).Uint())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		{
			var min, max uint32
			for k := range seen.Map {
				key :=

					uint32(r.ValueOf(k).Uint())
				min = key
				max = key
				break
			}
			for k := range seen.Map {
				key := uint32(r.ValueOf(k).Uint())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.Map) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.Map) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.Map))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.Map {
				key := uint32(r.ValueOf(k).Uint())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		{
			var min, max uint64
			for k := range seen.Map {
				key := r.ValueOf(k).Uint()
				min = key
				max = key
				break
			}
			for k := range seen.Map {
				key := r.ValueOf(k).Uint()
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.Map) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.Map) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.Map))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.Map {
				key := r.ValueOf(k).Uint()

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		{
			var min, max uintptr
			for k := range seen.Map {
				key :=

					uintptr(r.ValueOf(k).Uint())
				min = key
				max = key
				break
			}
			for k := range seen.Map {
				key :=

					uintptr(r.ValueOf(k).Uint())
				if min > key {
					min = key
				} else if max < key {
					max = key
				}

			}

			halfrange_trunc := max/2 - min/2
			if uint64(halfrange_trunc) >= uint64(MaxInt/2-3) || int(halfrange_trunc) > len(seen.Map) {
				break
			}

			fullrange := int(max-min) + 1
			if fullrange < len(seen.Map) {
				c.Errorf("switchGotoSlice: internal error, allocated slice has len=%v: less than the %d cases", fullrange, len(seen.Map))
			}

			slice := make([]int, fullrange)
			for k, v := range seen.Map {
				key := uintptr(r.ValueOf(k).Uint())

				slice[key-min] = v.IP + 1
			}
			stmt = func(env *Env) (Stmt, *Env) {
				val := efun(env)
				ip := 0
				if val >= min && val <= max {
					ip = slice[val-min]
				}

				if ip > 0 {
					env.IP = ip - 1
				} else {
					env.IP++
				}
				return env.Code[env.IP], env
			}
		}
	}
	return stmt
}
