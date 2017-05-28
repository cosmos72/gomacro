/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
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
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 *
 * code.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"unsafe"
)

func (code *Code) Clear() {
	code.List = nil
}

func (code *Code) Len() int {
	return len(code.List)
}

func (code *Code) Append(stmt Stmt) {
	if stmt != nil {
		code.List = append(code.List, stmt)
	}
}

func (code *Code) AsExpr() *Expr {
	fun := code.Exec()
	if fun == nil {
		return nil
	}
	return expr0(fun)
}

// Exec returns a func(*Env) that will execute the compiled code
func (code *Code) Exec() func(*Env) {
	if code.Len() == 0 {
		code.Clear()
		return nil
	}
	all := code.List
	code.Clear()
	all = append(all, nil)

	if len(all) == 2 {
		return func(env *Env) {
			env.IP = 0
			env.Code = all
			interrupt := env.ThreadGlobals.Interrupt
			env.ThreadGlobals.Interrupt = nil
			stmt := all[0]
			all[1] = nil
			for stmt != nil {
				stmt, env = stmt(env)
			}
			// restore env.ThreadGlobals.Interrupt before returning
			env.ThreadGlobals.Interrupt = interrupt
		}
	}
	return func(env *Env) {
		stmt := all[0]
		if stmt == nil {
			return
		}

		n := len(all) - 1
		all[n] = nil
		env.IP = 0
		env.Code = all
		interrupt := env.ThreadGlobals.Interrupt
		env.ThreadGlobals.Interrupt = nil

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
			// restore env.ThreadGlobals.Interrupt before returning
			env.ThreadGlobals.Interrupt = interrupt
			return
		}

		unsafeInterrupt := *(**uintptr)(unsafe.Pointer(&Interrupt))
		all[n] = Interrupt
		env.ThreadGlobals.Interrupt = Interrupt
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

			if x := stmt; *(**uintptr)(unsafe.Pointer(&x)) == unsafeInterrupt {
				// restore env.ThreadGlobals.Interrupt before returning
				env.ThreadGlobals.Interrupt = interrupt
				return
			}
		}
	}
}
