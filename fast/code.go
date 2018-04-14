/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
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
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * code.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
)

func (code *Code) Clear() {
	code.List = nil
	code.DebugPos = nil
	code.WithDefers = false
}

func (code *Code) Len() int {
	return len(code.List)
}

func (code *Code) Append(stmt Stmt, pos token.Pos) {
	if stmt != nil {
		code.List = append(code.List, stmt)
		code.DebugPos = append(code.DebugPos, pos)
	}
}

func (code *Code) AsExpr() *Expr {
	fun := code.Exec()
	if fun == nil {
		return nil
	}
	return expr0(fun)
}

// spinInterrupt is the statement executed while waiting for an interrupt to be serviced.
// To signal an interrupt, a statement must set env.ThreadGlobals.Signal to the desired signal,
// then return env.ThreadGlobals.Interrupt, env
func spinInterrupt(env *Env) (Stmt, *Env) {
	g := env.ThreadGlobals
	switch g.Signal {
	case SigNone:
		g.Signal = SigReturn
	case SigInterrupt:
		g.Signal = SigNone
		panic(SigInterrupt)
	}
	return g.Interrupt, env
}

func pushDefer(g *ThreadGlobals, deferOf *Env, panicking bool) (retg *ThreadGlobals, deferOf_ *Env, isDefer bool) {
	deferOf_ = g.DeferOfFun
	if panicking {
		g.PanicFun = deferOf
	}
	g.DeferOfFun = deferOf
	g.StartDefer = true
	return g, deferOf_, g.IsDefer
}

func popDefer(g *ThreadGlobals, deferOf *Env, isDefer bool) {
	g.DeferOfFun = deferOf
	g.StartDefer = false
	g.IsDefer = isDefer
}

func restore(g *ThreadGlobals, flag bool, interrupt Stmt) {
	g.IsDefer = flag
	g.Interrupt = interrupt
	sig := g.Signal
	g.Signal = SigNone
	if sig == SigInterrupt {
		panic(SigInterrupt)
	}
}

func maybeRepanic(g *ThreadGlobals) bool {
	if g.PanicFun != nil {
		panic(g.Panic)
	}
	// either not panicking or recover() invoked, no longer panicking
	return false
}

func (g *ThreadGlobals) interrupt() {
	g.Signal = SigInterrupt
}

// Exec returns a func(*Env) that will execute the compiled code
func (code *Code) Exec() func(*Env) {
	all := code.List
	pos := code.DebugPos
	defers := code.WithDefers

	code.Clear()
	if len(all) == 0 {
		return nil
	}
	all = append(all, spinInterrupt)

	if defers {
		// code to support defer is slower... isolate it in a separate function
		return func(env *Env) {
			execWithDefers(env, all, pos)
		}
	} else {
		return exec(all, pos)
	}
}

func exec(all []Stmt, pos []token.Pos) func(*Env) {
	return func(env *Env) {
		g := env.ThreadGlobals
		if g.IsDefer || g.StartDefer {
			// code to support defer is slower... isolate it in a separate function
			execWithDefers(env, all, pos)
			return
		}
		sig := g.Signal
		g.Signal = SigNone
		if sig == SigInterrupt {
			panic(SigInterrupt)
		}
		saveInterrupt := g.Interrupt
		g.Interrupt = nil

		stmt := all[0]
		env.IP = 0
		env.Code = all
		env.DebugPos = pos

		for j := 0; j < 5; j++ {
			if stmt, env = stmt(env); stmt != nil {
				if stmt, env = stmt(env); stmt != nil {
					if stmt, env = stmt(env); stmt != nil {
						if stmt, env = stmt(env); stmt != nil {
							if stmt, env = stmt(env); stmt != nil {
								if stmt, env = stmt(env); stmt != nil {
									if stmt, env = stmt(env); stmt != nil {
										if stmt, env = stmt(env); stmt != nil {
											if stmt, env = stmt(env); stmt != nil {
												if stmt, env = stmt(env); stmt != nil {
													if stmt, env = stmt(env); stmt != nil {
														if stmt, env = stmt(env); stmt != nil {
															if stmt, env = stmt(env); stmt != nil {
																if stmt, env = stmt(env); stmt != nil {
																	if g.Signal == SigNone {
																		continue
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
			goto finish
		}

		g.Interrupt = spinInterrupt
		for {
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)
			stmt, env = stmt(env)

			if g.Signal != SigNone {
				break
			}
		}
	finish:
		// restore env.ThreadGlobals.Interrupt and Signal before returning
		g.Interrupt = saveInterrupt
		sig = g.Signal
		g.Signal = SigNone
		if sig == SigInterrupt {
			panic(SigInterrupt)
		}
		return
	}
}

// execWithDefers executes the given compiled code, including support for defer()
func execWithDefers(env *Env, all []Stmt, pos []token.Pos) {
	g := env.ThreadGlobals

	sig := g.Signal
	g.Signal = SigNone
	if sig == SigInterrupt {
		panic(SigInterrupt)
	}
	defer restore(g, g.IsDefer, g.Interrupt) // restore g.IsDefer, g.Signal and g.Interrupt on return
	g.Interrupt = nil
	g.IsDefer = g.StartDefer
	g.StartDefer = false

	funenv := env
	stmt := all[0]
	env.IP = 0
	env.Code = all
	env.DebugPos = pos

	panicking, panicking2 := true, false
	rundefer := func(fun func()) {
		if panicking || panicking2 {
			panicking = true
			panicking2 = false
			g.Panic = recover()
		}
		defer popDefer(pushDefer(g, funenv, panicking))
		panicking2 = true // detect panics inside defer
		fun()
		panicking2 = false
		if panicking {
			panicking = maybeRepanic(g)
		}
	}

	for j := 0; j < 5; j++ {
		if stmt, env = stmt(env); stmt != nil {
			if stmt, env = stmt(env); stmt != nil {
				if stmt, env = stmt(env); stmt != nil {
					if stmt, env = stmt(env); stmt != nil {
						if stmt, env = stmt(env); stmt != nil {
							if stmt, env = stmt(env); stmt != nil {
								if stmt, env = stmt(env); stmt != nil {
									if stmt, env = stmt(env); stmt != nil {
										if stmt, env = stmt(env); stmt != nil {
											if stmt, env = stmt(env); stmt != nil {
												if stmt, env = stmt(env); stmt != nil {
													if stmt, env = stmt(env); stmt != nil {
														if stmt, env = stmt(env); stmt != nil {
															if stmt, env = stmt(env); stmt != nil {
																if g.Signal == SigNone {
																	continue
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
		if g.Signal != SigDefer {
			goto finish
		}
		g.Signal = SigNone
		fun := g.InstallDefer
		g.InstallDefer = nil
		defer rundefer(fun)
		stmt = env.Code[env.IP]
		if stmt == nil {
			goto finish
		}
		continue
	}

	g.Interrupt = spinInterrupt
	for {
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)
		stmt, env = stmt(env)

		for sig := g.Signal; sig != SigNone; {
			if sig != SigDefer {
				goto finish
			}
			g.Signal = SigNone
			fun := g.InstallDefer
			g.InstallDefer = nil
			defer rundefer(fun)
			stmt = env.Code[env.IP]
			// single step
			stmt, env = stmt(env)
		}
	}
finish:
	panicking = false
	// no need to restore g.IsDefer, g.Signal, g.Interrupt:
	// done by defer restore(g, g.IsDefer, interrupt) above
	return
}
