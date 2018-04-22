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

	. "github.com/cosmos72/gomacro/base"
)

func (code *Code) Clear() {
	code.List = nil
	code.DebugPos = nil
	code.WithDefers = false
}

func (code *Code) Len() int {
	return len(code.List)
}

func (code *Code) Truncate(n int) {
	if len(code.List) > n {
		code.List = code.List[0:n]
	}
	if len(code.DebugPos) > n {
		code.DebugPos = code.DebugPos[0:n]
	}
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
	if g.Signals.Sync == SigNone {
		g.Signals.Sync = SigReturn
	}
	if g.Signals.Async == SigInterrupt {
		g.Signals.Async = SigNone
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
	g.ExecFlags.SetStartDefer(true)
	return g, deferOf_, g.ExecFlags.IsDefer()
}

func popDefer(g *ThreadGlobals, deferOf *Env, isDefer bool) {
	g.DeferOfFun = deferOf
	g.ExecFlags.SetStartDefer(false)
	g.ExecFlags.SetDefer(isDefer)
}

func restore(g *ThreadGlobals, isDefer bool, interrupt Stmt) {
	g.ExecFlags.SetDefer(isDefer)
	g.Interrupt = interrupt
	g.Signals.Sync = SigNone
	if g.Signals.Async == SigInterrupt {
		g.Signals.Async = SigNone
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
	g.Signals.Async = SigInterrupt
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
		return execWithFlags(all, pos)
	}
	return exec(all, pos)
}

func exec(all []Stmt, pos []token.Pos) func(*Env) {
	return func(env *Env) {
		g := env.ThreadGlobals
		g.Signals.Sync = SigNone
		if g.ExecFlags != 0 {
			// code to support defer and debugger is slower... isolate it in a separate function
			reExecWithFlags(env, all, pos, all[0], 0)
			return
		}
		if g.Signals.Async == SigInterrupt {
			g.Signals.Async = SigNone
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
																	if g.Signals.IsEmpty() {
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

			if !g.Signals.IsEmpty() {
				break
			}
		}
	finish:
		// restore env.ThreadGlobals.Interrupt and Signal before returning
		g.Interrupt = saveInterrupt
		if g.Signals.Async == SigInterrupt {
			g.Signals.Async = SigNone
			panic(SigInterrupt)
		}
		if g.Signals.Debug != SigNone {
			reExecWithFlags(env, all, pos, stmt, env.IP)
			return
		}
		g.Signals.Sync = SigNone
		return
	}
}

// execWithFlags returns a function that will execute the given compiled code, including support for defer() and debugger
func execWithFlags(all []Stmt, pos []token.Pos) func(*Env) {
	return func(env *Env) {
		env.ThreadGlobals.Signals.Sync = SigNone
		reExecWithFlags(env, all, pos, all[0], 0)
	}
}

func reExecWithFlags(env *Env, all []Stmt, pos []token.Pos, stmt Stmt, ip int) {
	g := env.ThreadGlobals

	if g.Signals.Async == SigInterrupt {
		g.Signals.Async = SigNone
		panic(SigInterrupt)
	}
	defer restore(g, g.ExecFlags.IsDefer(), g.Interrupt) // restore g.IsDefer, g.Signal and g.Interrupt on return
	ef := &g.ExecFlags
	ef.SetDefer(ef.StartDefer())
	ef.SetStartDefer(false)
	ef.SetDebug(g.Signals.Debug != SigNone)

	funenv := env
	env.IP = ip
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

	if stmt == nil {
		goto signal
	}
again:
	g.Interrupt = nil
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
																if g.Signals.IsEmpty() {
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
		for g.Signals.Sync == SigDefer {
			g.Signals.Sync = SigNone
			fun := g.InstallDefer
			g.InstallDefer = nil
			defer rundefer(fun)
			stmt = env.Code[env.IP]
			if stmt == nil {
				goto signal
			}
		}
		if !g.Signals.IsEmpty() {
			goto signal
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

		for g.Signals.Sync == SigDefer {
			g.Signals.Sync = SigNone
			fun := g.InstallDefer
			g.InstallDefer = nil
			defer rundefer(fun)
			// single step
			stmt = env.Code[env.IP]
			stmt, env = stmt(env)
		}
		if !g.Signals.IsEmpty() {
			goto signal
		}
	}
signal:
	for g.Signals.Debug != SigNone {
		g.Interrupt = spinInterrupt
		stmt, env = singleStep(env)
		/*DELETEME*/ g.Debugf("singleStep returned stmt = %p, env = %p with signals = %#v\n", stmt, env, g.Signals)
		// a Sync or Async signal may be pending.
		if g.Signals.Sync == SigReturn || g.Signals.Async != SigNone {
			break
		}
		if g.Signals.IsEmpty() || g.Signals.Sync == SigDefer {
			goto again
		}
	}
	panicking = false
	// no need to restore g.IsDefer, g.Signal, g.Interrupt:
	// done by defer restore(g, g.IsDefer, interrupt) above
	return
}

func singleStep(env *Env) (Stmt, *Env) {
	stmt := env.Code[env.IP]
	g := env.ThreadGlobals
	switch g.Signals.Debug {
	case SigDebugFinish:
		g.DebugDepth = env.DebugCallDepth
	case SigDebugNext:
		g.DebugDepth = env.DebugCallDepth + 1
	case SigDebugStep:
		g.DebugDepth = MaxInt
	case SigDebugRepl:
		break
	default:
		g.Signals.Debug = SigNone
		g.DebugDepth = 0
		return stmt, env // resume normal execution
	}
	// prevent further changes to g.DebugDepth
	if g.Signals.Debug != SigNone {
		g.Signals.Debug = SigDebugRepl
	}

	if env.DebugCallDepth < g.DebugDepth {
		c := env.DebugComp
		if c != nil {
			ir := Interp{c, env}
			op := ir.debug(false) // not a breakpoint
			if op == SigDebugContinue {
				g.Signals.Debug = SigNone
			} else {
				g := env.ThreadGlobals
				g.Signals.Debug = op
			}
		}
	}

	// single step
	stmt, env = stmt(env)
	if g.Signals.Debug != SigNone {
		stmt = g.Interrupt
	}
	return stmt, env
}
