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
		val := r.Zero(e.Type)
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
func (c *Comp) switchGoto(tag *Expr, seen *caseHelper, ip int) {
	if seen.SomeNonConst {
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
