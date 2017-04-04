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
	"unsafe"
)

const (
	n int = 10000
)

/*
    -------- n =  5 --------
	BenchmarkThreadedStmtFunc0-8            	100000000	        14.3 ns/op
	BenchmarkThreadedStmtFunc1-8            	100000000	        14.8 ns/op
	BenchmarkThreadedStmtFunc2-8            	100000000	        14.7 ns/op
	BenchmarkThreadedStmtFunc3-8            	100000000	        13.4 ns/op
	BenchmarkThreadedStmtFunc4-8            	100000000	        13.4 ns/op
	BenchmarkThreadedStmtFunc4Unroll -8     	100000000	        12.5 ns/op
	BenchmarkThreadedStmtFunc4Terminate-8   	50000000	        37.0 ns/op
	BenchmarkThreadedStmtFunc4Panic-8       	10000000	       128 ns/op
	BenchmarkThreadedStmtStruct1-8          	100000000	        14.9 ns/op
	BenchmarkThreadedStmtStruct4-8          	100000000	        13.2 ns/op
	BenchmarkThreadedStmtStruct4Unroll-8    	100000000	        12.4 ns/op

    -------- n = 20 --------
	BenchmarkThreadedStmtFunc0-8            	20000000	        69.7 ns/op
	BenchmarkThreadedStmtFunc1-8            	30000000	        59.8 ns/op
	BenchmarkThreadedStmtFunc2-8            	30000000	        59.0 ns/op
	BenchmarkThreadedStmtFunc3-8            	30000000	        54.7 ns/op
	BenchmarkThreadedStmtFunc4-8            	30000000	        48.6 ns/op
	BenchmarkThreadedStmtFunc4Unroll -8     	30000000	        47.9 ns/op
	BenchmarkThreadedStmtFunc4Terminate-8   	20000000	        74.6 ns/op
	BenchmarkThreadedStmtFunc4Panic-8       	10000000	       161 ns/op
	BenchmarkThreadedStmtStruct1-8          	30000000	        59.1 ns/op
	BenchmarkThreadedStmtStruct4-8          	30000000	        48.8 ns/op
	BenchmarkThreadedStmtStruct4Unroll-8    	30000000	        46.4 ns/op

    -------- n = 100 --------
	BenchmarkThreadedStmtFunc0-8            	 3000000	       418 ns/op
	BenchmarkThreadedStmtFunc1-8            	 5000000	       313 ns/op
	BenchmarkThreadedStmtFunc2-8            	 5000000	       303 ns/op
	BenchmarkThreadedStmtFunc3-8            	 5000000	       295 ns/op
	BenchmarkThreadedStmtFunc4-8            	 5000000	       250 ns/op
	BenchmarkThreadedStmtFunc4Unroll -8     	 5000000	       242 ns/op
	BenchmarkThreadedStmtFunc4Terminate-8   	10000000	       233 ns/op
	BenchmarkThreadedStmtFunc4Panic-8       	 5000000	       345 ns/op
	BenchmarkThreadedStmtStruct1-8          	 5000000	       304 ns/op
	BenchmarkThreadedStmtStruct4-8          	 5000000	       248 ns/op
	BenchmarkThreadedStmtStruct4Unroll-8    	 5000000	       247 ns/op

    -------- n = 1000 --------
	BenchmarkThreadedStmtFunc0-8            	  300000	      4205 ns/op
	BenchmarkThreadedStmtFunc1-8            	  500000	      3031 ns/op
	BenchmarkThreadedStmtFunc2-8            	  500000	      2996 ns/op
	BenchmarkThreadedStmtFunc3-8            	  500000	      2875 ns/op
	BenchmarkThreadedStmtFunc4-8            	  500000	      2406 ns/op
	BenchmarkThreadedStmtFunc4Unroll -8     	 1000000	      2349 ns/op
	BenchmarkThreadedStmtFunc4Terminate-8   	 1000000	      2177 ns/op
	BenchmarkThreadedStmtFunc4Panic-8       	  500000	      2255 ns/op
	BenchmarkThreadedStmtStruct1-8          	  500000	      2968 ns/op
	BenchmarkThreadedStmtStruct4-8          	  500000	      2410 ns/op
	BenchmarkThreadedStmtStruct4Unroll-8    	 1000000	      2282 ns/op

    -------- n = 10000 --------
	BenchmarkThreadedStmtFunc0-8            	   30000	     42124 ns/op
	BenchmarkThreadedStmtFunc1-8            	   50000	     30382 ns/op
	BenchmarkThreadedStmtFunc2-8            	   50000	     29695 ns/op
	BenchmarkThreadedStmtFunc3-8            	   50000	     29007 ns/op
	BenchmarkThreadedStmtFunc4-8            	   50000	     24040 ns/op
	BenchmarkThreadedStmtFunc4Unroll -8     	  100000	     23418 ns/op
	BenchmarkThreadedStmtFunc4Terminate-8   	  100000	     21372 ns/op
	BenchmarkThreadedStmtFunc4Panic-8       	  100000	     21506 ns/op
	BenchmarkThreadedStmtStruct1-8          	   50000	     29901 ns/op
	BenchmarkThreadedStmtStruct4-8          	  100000	     24938 ns/op
	BenchmarkThreadedStmtStruct4Unroll-8    	  100000	     22599 ns/op
*/

type Env struct {
	Binds []r.Value
	Outer *Env
}

func NewEnv(outer *Env) *Env {
	return &Env{
		Binds: make([]r.Value, 10),
		Outer: outer,
	}
}

func BenchmarkThreadedStmtFunc0(b *testing.B) {

	type Stmt func(env *Env, all []Stmt) Stmt

	env := NewEnv(nil)
	all := make([]Stmt, n+1)
	for i := 0; i < n; i++ {
		i := i
		all[i] = func(env *Env, all []Stmt) Stmt {
			return all[i+1]
		}
	}
	all[n] = nil

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stmt := all[0]
		for stmt != nil {
			stmt = stmt(env, all)
		}
	}
}

func BenchmarkThreadedStmtFunc1(b *testing.B) {

	type Stmt func(env *Env, next *Stmt, ip int, all []Stmt) (Stmt, int)
	var nop Stmt = func(env *Env, next *Stmt, ip int, all []Stmt) (Stmt, int) {
		return *next, ip
	}

	env := NewEnv(nil)
	all := make([]Stmt, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = nil

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ip := 0
		stmt := all[ip]
		for stmt != nil {
			ip++
			stmt, ip = stmt(env, &all[ip], ip, all)
		}
	}
}

func BenchmarkThreadedStmtFunc2(b *testing.B) {

	type Stmt func(env *Env, ip int, all []Stmt) (Stmt, int)
	var nop Stmt = func(env *Env, ip int, all []Stmt) (Stmt, int) {
		ip++
		return all[ip], ip
	}

	env := NewEnv(nil)
	all := make([]Stmt, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = nil

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ip := 0
		stmt := all[ip]
		for stmt != nil {
			stmt, ip = stmt(env, ip, all)
		}
	}
}

type (
	EnvF3 struct {
		Binds []r.Value
		Outer *Env
		Code  []StmtF3
	}
	StmtF3 func(env *EnvF3, ip int) (StmtF3, int)
)

func BenchmarkThreadedStmtFunc3(b *testing.B) {

	var nop StmtF3 = func(env *EnvF3, ip int) (StmtF3, int) {
		ip++
		return env.Code[ip], ip
	}
	env := &EnvF3{
		Binds: make([]r.Value, 10),
	}
	all := make([]StmtF3, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ip := 0
		stmt := all[ip]
		for stmt != nil {
			stmt, ip = stmt(env, ip)
		}
	}
}

type (
	EnvF4 struct {
		Binds     []r.Value
		Outer     *Env
		Code      []StmtF4
		IP        int
		Interrupt StmtF4
	}
	StmtF4 func(env *EnvF4) StmtF4
)

func BenchmarkThreadedStmtFunc4(b *testing.B) {
	var nop StmtF4 = func(env *EnvF4) StmtF4 {
		env.IP++
		return env.Code[env.IP]
	}
	env := &EnvF4{
		Binds: make([]r.Value, 10),
	}
	all := make([]StmtF4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt != nil {
			stmt = stmt(env)
		}
	}
}

func BenchmarkThreadedStmtFunc4Unroll(b *testing.B) {
	var nop StmtF4 = func(env *EnvF4) StmtF4 {
		env.IP++
		return env.Code[env.IP]
	}
	env := &EnvF4{
		Binds: make([]r.Value, 10),
	}
	all := make([]StmtF4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt != nil {
			if stmt = stmt(env); stmt != nil {
				if stmt = stmt(env); stmt != nil {
					if stmt = stmt(env); stmt != nil {
						if stmt = stmt(env); stmt != nil {
							if stmt = stmt(env); stmt != nil {
								if stmt = stmt(env); stmt != nil {
									if stmt = stmt(env); stmt != nil {
										if stmt = stmt(env); stmt != nil {
											if stmt = stmt(env); stmt != nil {
												if stmt = stmt(env); stmt != nil {
													if stmt = stmt(env); stmt != nil {
														if stmt = stmt(env); stmt != nil {
															if stmt = stmt(env); stmt != nil {
																if stmt = stmt(env); stmt != nil {
																	stmt = stmt(env)
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func BenchmarkThreadedStmtFunc4Terminate(b *testing.B) {
	var interrupt StmtF4
	interrupt = func(env *EnvF4) StmtF4 {
		return interrupt
	}
	unsafeInterrupt := *(**uintptr)(unsafe.Pointer(&interrupt))

	var nop StmtF4 = func(env *EnvF4) StmtF4 {
		env.IP++
		return env.Code[env.IP]
	}
	env := &EnvF4{
		Binds: make([]r.Value, 10),
	}
	all := make([]StmtF4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = interrupt
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for {
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)

			if x := stmt; *(**uintptr)(unsafe.Pointer(&x)) == unsafeInterrupt {
				break
			}
		}
	}
}

func BenchmarkThreadedStmtFunc4Adaptive(b *testing.B) {
	var nop StmtF4 = func(env *EnvF4) StmtF4 {
		env.IP++
		return env.Code[env.IP]
	}
	var interrupt StmtF4 = func(env *EnvF4) StmtF4 {
		return env.Interrupt
	}
	unsafeInterrupt := *(**uintptr)(unsafe.Pointer(&interrupt))
	_ = unsafeInterrupt
	env := &EnvF4{
		Binds: make([]r.Value, 10),
	}
	all := make([]StmtF4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = nil
	env.Code = all

	b.ResetTimer()
outer:
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		if stmt == nil {
			continue outer
		}
		for j := 0; j < 10; j++ {
			if stmt = stmt(env); stmt != nil {
				if stmt = stmt(env); stmt != nil {
					if stmt = stmt(env); stmt != nil {
						if stmt = stmt(env); stmt != nil {
							if stmt = stmt(env); stmt != nil {
								if stmt = stmt(env); stmt != nil {
									if stmt = stmt(env); stmt != nil {
										if stmt = stmt(env); stmt != nil {
											if stmt = stmt(env); stmt != nil {
												if stmt = stmt(env); stmt != nil {
													if stmt = stmt(env); stmt != nil {
														if stmt = stmt(env); stmt != nil {
															if stmt = stmt(env); stmt != nil {
																if stmt = stmt(env); stmt != nil {
																	break
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			continue outer
		}
		all[n] = interrupt
		env.Interrupt = interrupt
		for {
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)
			stmt = stmt(env)

			if x := stmt; *(**uintptr)(unsafe.Pointer(&x)) == unsafeInterrupt {
				continue outer
			}
		}
	}
}

func BenchmarkThreadedStmtFunc4Panic(b *testing.B) {
	var terminate StmtF4 = func(env *EnvF4) StmtF4 {
		panic("end of code")
	}

	var nop StmtF4 = func(env *EnvF4) StmtF4 {
		env.IP++
		return env.Code[env.IP]
	}
	env := &EnvF4{
		Binds: make([]r.Value, 10),
	}
	all := make([]StmtF4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = terminate
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runThreadedStmtFunc4Panic(env)
	}
}

func runThreadedStmtFunc4Panic(env *EnvF4) {
	env.IP = 0
	stmt := env.Code[0]
	defer func() {
		recover()
	}()
	for {
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
		stmt = stmt(env)
	}
}

func BenchmarkThreadedStmtStruct1(b *testing.B) {

	type Stmt struct {
		Exec func(env *Env, ip int, all []Stmt) (Stmt, int)
	}
	var nop = Stmt{func(env *Env, ip int, all []Stmt) (Stmt, int) {
		ip++
		return all[ip], ip
	}}

	env := NewEnv(nil)
	all := make([]Stmt, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = Stmt{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ip := 0
		stmt := all[ip]
		for stmt.Exec != nil {
			stmt, ip = stmt.Exec(env, ip, all)
		}
	}
}

type (
	Env4 struct {
		Binds []r.Value
		Outer *Env
		Code  []Stmt4
		IP    int
	}
	Stmt4 struct {
		Exec func(env *Env4) Stmt4
	}
)

func BenchmarkThreadedStmtStruct4(b *testing.B) {

	var nop Stmt4 = Stmt4{func(env *Env4) Stmt4 {
		env.IP++
		return env.Code[env.IP]
	}}
	env := &Env4{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = Stmt4{}
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt.Exec != nil {
			stmt = stmt.Exec(env)
		}
	}
}

func BenchmarkThreadedStmtStruct4Unroll(b *testing.B) {

	var nop Stmt4 = Stmt4{func(env *Env4) Stmt4 {
		env.IP++
		return env.Code[env.IP]
	}}
	env := &Env4{
		Binds: make([]r.Value, 10),
	}
	all := make([]Stmt4, n+1)
	for i := 0; i < n; i++ {
		all[i] = nop
	}
	all[n] = Stmt4{}
	env.Code = all

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env.IP = 0
		stmt := all[0]
		for stmt.Exec != nil {
			stmt = stmt.Exec(env)
			if stmt.Exec != nil {
				stmt = stmt.Exec(env)
				if stmt.Exec != nil {
					stmt = stmt.Exec(env)
					if stmt.Exec != nil {
						stmt = stmt.Exec(env)
						if stmt.Exec != nil {
							stmt = stmt.Exec(env)
							if stmt.Exec != nil {
								stmt = stmt.Exec(env)
								if stmt.Exec != nil {
									stmt = stmt.Exec(env)
									if stmt.Exec != nil {
										stmt = stmt.Exec(env)
										if stmt.Exec != nil {
											stmt = stmt.Exec(env)
											if stmt.Exec != nil {
												stmt = stmt.Exec(env)
												if stmt.Exec != nil {
													stmt = stmt.Exec(env)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
