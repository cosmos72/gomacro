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
 * code.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	r "reflect"
	"unsafe"

	. "github.com/cosmos72/gomacro/base"
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

// more wrapping (thus slower) than needed... but only used by REPL
func (code *Code) AsXV() func(*Env) (r.Value, []r.Value) {
	fun := code.Exec()
	if fun == nil {
		return nil
	}
	return func(env *Env) (r.Value, []r.Value) {
		fun(env)
		return None, nil
	}
}

// Exec returns a func(*Env) that will execute the compiled code
func (code *Code) Exec() X {
	if code.Len() == 0 {
		code.Clear()
		return nil
	}
	all := code.List
	code.Clear()
	all = append(all, nil)

	if len(all) == 2 {
		return func(env *Env) {
			env.Interrupt = nil
			env.IP = 0
			env.Code = all
			stmt := all[0]
			all[1] = nil
			for stmt != nil {
				stmt, env = stmt(env)
			}
		}
	}
	return func(env *Env) {
		stmt := all[0]
		if stmt == nil {
			return
		}

		n := len(all) - 1
		all[n] = nil
		env.Interrupt = nil
		env.IP = 0
		env.Code = all

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
			return
		}

		unsafeInterrupt := *(**uintptr)(unsafe.Pointer(&Interrupt))
		all[n] = Interrupt
		env.Interrupt = Interrupt
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
				return
			}
		}
	}
}
