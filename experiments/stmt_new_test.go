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
 * stmt_test.go
 *
 *  Created on Apr 04, 2017
 *      Author Massimiliano Ghilardi
 */

package experiments

import (
	r "reflect"
	"testing"
)

const (
	n int = 10000
)

type (
	Env5 struct {
		Binds []r.Value
		IP    int
		Code  []Stmt5
		Outer *Env5
	}
	Stmt5 func(*Env5) (Stmt5, *Env5)
)

func BenchmarkThreadedStmtFunc5(b *testing.B) {

	var nop Stmt5 = func(env *Env5) (Stmt5, *Env5) {
		env.IP++
		return env.Code[env.IP], env
	}

	env := &Env5{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt5, n+1)
	for i := 0; i < n; i++ {
		i := i
		all[i] = nop
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt != nil {
			stmt, env = stmt(env)
		}
	}
}

type (
	Env6 struct {
		Binds []r.Value
		IP    int
		Code  []Stmt6
		Outer *Env6
	}
	Stmt6 func(**Env6) Stmt6
)

func BenchmarkThreadedStmtFunc6(b *testing.B) {

	var nop Stmt6 = func(penv **Env6) Stmt6 {
		env := *penv
		env.IP++
		return env.Code[env.IP]
	}

	env := &Env6{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt6, n+1)
	for i := 0; i < n; i++ {
		i := i
		all[i] = nop
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt != nil {
			stmt = stmt(&env)
		}
	}
}
