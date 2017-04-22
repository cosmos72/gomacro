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
 * call1ret1.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
)

func call1ret1(c *Call, maxdepth int) I {
	expr := c.Fun
	exprfun := expr.AsX1()
	funsym := c.Funsym
	funupn, funindex := -1, -1
	if funsym != nil {
		funupn = funsym.Upn
		funindex = funsym.Desc.Index()
	}
	karg0 := expr.Type.In(0).Kind()
	kret := expr.Type.Out(0).Kind()
	arg0 := c.Args[0]
	argfun := c.Argfuns[0]
	var cachedfunv r.Value
	var call I
	switch kret {

	case r.Bool:
		switch karg0 {
		case r.Bool:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(bool)
					switch funupn {
					case maxdepth:
						var cachedfun func(bool) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(bool) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) bool)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(bool) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(bool) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(bool) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(bool) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(bool) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int)
					switch funupn {
					case maxdepth:
						var cachedfun func(int) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(int) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(int) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(int) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int8)
					switch funupn {
					case maxdepth:
						var cachedfun func(int8) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int8) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int8) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int8) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(int8) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(int8) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(int8) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int16)
					switch funupn {
					case maxdepth:
						var cachedfun func(int16) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int16) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int16) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int16) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(int16) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(int16) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(int16) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int32)
					switch funupn {
					case maxdepth:
						var cachedfun func(int32) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int32) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int32) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int32) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(int32) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(int32) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(int32) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int64)
					switch funupn {
					case maxdepth:
						var cachedfun func(int64) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int64) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int64) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int64) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(int64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(int64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(int64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(uint) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(uint) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(uint) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint8)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint8) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint8) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint8) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint8) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(uint8) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(uint8) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(uint8) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint16)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint16) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint16) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint16) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint16) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(uint16) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(uint16) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(uint16) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint32)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint32) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint32) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint32) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint32) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(uint32) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(uint32) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(uint32) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint64)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint64) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint64) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint64) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint64) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(uint64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(uint64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(uint64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uintptr:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uintptr)
					switch funupn {
					case maxdepth:
						var cachedfun func(uintptr) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uintptr) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uintptr) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uintptr) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(uintptr) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(uintptr) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(uintptr) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float32)
					switch funupn {
					case maxdepth:
						var cachedfun func(float32) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float32) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float32) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float32) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(float32) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(float32) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(float32) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float64)
					switch funupn {
					case maxdepth:
						var cachedfun func(float64) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float64) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float64) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float64) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(float64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(float64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(float64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex64)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex64) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex64) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex64) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex64) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(complex64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(complex64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(complex64) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex128:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex128)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex128) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex128) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex128)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex128) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex128) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(complex128) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(complex128) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(complex128) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.String:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(string)
					switch funupn {
					case maxdepth:
						var cachedfun func(string) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) bool)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(string) bool

						call = func(env *Env) bool {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) bool)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) string)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(string) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(string) bool

							call = func(env *Env) bool {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) bool)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) bool {
								fun := env.Binds[funindex].Interface().(func(string) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) bool {
								fun := env.Outer.Binds[funindex].Interface().(func(string) bool)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) bool {
								fun := exprfun(env).Interface().(func(string) bool)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		default:
			call = func(env *Env) bool {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Bool()
			}
		}
	case r.Int:
		switch karg0 {
		case r.Bool:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(bool)
					switch funupn {
					case maxdepth:
						var cachedfun func(bool) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(bool) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) bool)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(bool) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(bool) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(bool) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(bool) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(bool) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int)
					switch funupn {
					case maxdepth:
						var cachedfun func(int) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(int) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(int) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(int) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int8)
					switch funupn {
					case maxdepth:
						var cachedfun func(int8) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int8) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int8) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int8) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(int8) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(int8) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(int8) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int16)
					switch funupn {
					case maxdepth:
						var cachedfun func(int16) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int16) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int16) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int16) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(int16) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(int16) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(int16) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int32)
					switch funupn {
					case maxdepth:
						var cachedfun func(int32) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int32) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int32) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int32) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(int32) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(int32) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(int32) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int64)
					switch funupn {
					case maxdepth:
						var cachedfun func(int64) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int64) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int64) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int64) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(int64) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(int64) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(int64) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(uint) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(uint) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(uint) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint8)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint8) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint8) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint8) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint8) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(uint8) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(uint8) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(uint8) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint16)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint16) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint16) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint16) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint16) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(uint16) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(uint16) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(uint16) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint32)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint32) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint32) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint32) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint32) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(uint32) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(uint32) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(uint32) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint64)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint64) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint64) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint64) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint64) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(uint64) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(uint64) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(uint64) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uintptr:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uintptr)
					switch funupn {
					case maxdepth:
						var cachedfun func(uintptr) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uintptr) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uintptr) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uintptr) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(uintptr) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(uintptr) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(uintptr) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float32)
					switch funupn {
					case maxdepth:
						var cachedfun func(float32) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float32) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float32) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float32) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(float32) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(float32) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(float32) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float64)
					switch funupn {
					case maxdepth:
						var cachedfun func(float64) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float64) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float64) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float64) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(float64) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(float64) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(float64) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex64)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex64) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex64) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex64) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex64) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(complex64) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(complex64) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(complex64) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex128:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex128)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex128) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex128) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex128)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex128) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex128) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(complex128) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(complex128) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(complex128) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.String:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(string)
					switch funupn {
					case maxdepth:
						var cachedfun func(string) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(string) int

						call = func(env *Env) int {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) string)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(string) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(string) int

							call = func(env *Env) int {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) int)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int {
								fun := env.Binds[funindex].Interface().(func(string) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int {
								fun := env.Outer.Binds[funindex].Interface().(func(string) int)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int {
								fun := exprfun(env).Interface().(func(string) int)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		default:
			call = func(env *Env) int {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return int(ret0.Int())
			}
		}
	case r.Int8:
		if karg0 == kret {

			if arg0.Const() && funsym != nil {
				argconst := arg0.Value.(int8)
				switch funupn {
				case maxdepth:
					var cachedfun func(int8) int8

					call = func(env *Env) int8 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8) int8)
						}
						return cachedfun(argconst)
					}
				case maxdepth - 1:
					var cachedfun func(int8) int8

					call = func(env *Env) int8 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8) int8)
						}
						return cachedfun(argconst)
					}
				}
			}
			if call == nil {
				argfun := arg0.WithFun().(func(env *Env) int8)
				if funsym != nil {
					switch funupn {
					case maxdepth:
						var cachedfun func(int8) int8

						call = func(env *Env) int8 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int8)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					case maxdepth - 1:
						var cachedfun func(int8) int8

						call = func(env *Env) int8 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int8)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					}
				}

				if call == nil {
					if funsym != nil && funupn == 0 {
						call = func(env *Env) int8 {
							fun := env.Binds[funindex].Interface().(func(int8) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funsym != nil && funupn == 1 {
						call = func(env *Env) int8 {
							fun := env.Outer.Binds[funindex].Interface().(func(int8) int8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int8 {
							fun := exprfun(env).Interface().(func(int8) int8)
							arg := argfun(env)
							return fun(arg)
						}
					}
				}

			}
		} else {
			switch karg0 {
			case r.Bool:

				{
					argfun := arg0.WithFun().(func(env *Env) bool)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(bool) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int:

				{
					argfun := arg0.WithFun().(func(env *Env) int)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(int) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int8:
				{
					argfun := arg0.WithFun().(func(env *Env) int8)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(int8) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int16:
				{
					argfun := arg0.WithFun().(func(env *Env) int16)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(int16) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int32:
				{
					argfun := arg0.WithFun().(func(env *Env) int32)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(int32) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int64:
				{
					argfun := arg0.WithFun().(func(env *Env) int64)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(int64) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint:
				{
					argfun := arg0.WithFun().(func(env *Env) uint)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(uint) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint8:
				{
					argfun := arg0.WithFun().(func(env *Env) uint8)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(uint8) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint16:
				{
					argfun := arg0.WithFun().(func(env *Env) uint16)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(uint16) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint32:
				{
					argfun := arg0.WithFun().(func(env *Env) uint32)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(uint32) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint64:
				{
					argfun := arg0.WithFun().(func(env *Env) uint64)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(uint64) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uintptr:
				{
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(uintptr) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Float32:
				{
					argfun := arg0.WithFun().(func(env *Env) float32)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(float32) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Float64:
				{
					argfun := arg0.WithFun().(func(env *Env) float64)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(float64) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex64:
				{
					argfun := arg0.WithFun().(func(env *Env) complex64)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(complex64) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex128:
				{
					argfun := arg0.WithFun().(func(env *Env) complex128)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(complex128) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.String:
				{
					argfun := arg0.WithFun().(func(env *Env) string)
					call = func(env *Env) int8 {
						fun := exprfun(env).Interface().(func(string) int8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			default:
				call = func(env *Env) int8 {
					funv := exprfun(env)
					argv := []r.Value{
						argfun(env),
					}

					ret0 := funv.Call(argv)[0]
					return int8(ret0.Int())
				}
			}
		}
	case r.Int16:
		if karg0 == kret {

			if arg0.Const() && funsym != nil {
				argconst := arg0.Value.(int16)
				switch funupn {
				case maxdepth:
					var cachedfun func(int16) int16

					call = func(env *Env) int16 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16) int16)
						}
						return cachedfun(argconst)
					}
				case maxdepth - 1:
					var cachedfun func(int16) int16

					call = func(env *Env) int16 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16) int16)
						}
						return cachedfun(argconst)
					}
				}
			}
			if call == nil {
				argfun := arg0.WithFun().(func(env *Env) int16)
				if funsym != nil {
					switch funupn {
					case maxdepth:
						var cachedfun func(int16) int16

						call = func(env *Env) int16 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int16)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					case maxdepth - 1:
						var cachedfun func(int16) int16

						call = func(env *Env) int16 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int16)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					}
				}

				if call == nil {
					if funsym != nil && funupn == 0 {
						call = func(env *Env) int16 {
							fun := env.Binds[funindex].Interface().(func(int16) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funsym != nil && funupn == 1 {
						call = func(env *Env) int16 {
							fun := env.Outer.Binds[funindex].Interface().(func(int16) int16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) int16 {
							fun := exprfun(env).Interface().(func(int16) int16)
							arg := argfun(env)
							return fun(arg)
						}
					}
				}

			}
		} else {
			switch karg0 {
			case r.Bool:

				{
					argfun := arg0.WithFun().(func(env *Env) bool)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(bool) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int:

				{
					argfun := arg0.WithFun().(func(env *Env) int)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(int) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int8:
				{
					argfun := arg0.WithFun().(func(env *Env) int8)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(int8) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int16:
				{
					argfun := arg0.WithFun().(func(env *Env) int16)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(int16) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int32:
				{
					argfun := arg0.WithFun().(func(env *Env) int32)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(int32) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int64:
				{
					argfun := arg0.WithFun().(func(env *Env) int64)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(int64) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint:
				{
					argfun := arg0.WithFun().(func(env *Env) uint)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(uint) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint8:
				{
					argfun := arg0.WithFun().(func(env *Env) uint8)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(uint8) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint16:
				{
					argfun := arg0.WithFun().(func(env *Env) uint16)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(uint16) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint32:
				{
					argfun := arg0.WithFun().(func(env *Env) uint32)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(uint32) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint64:
				{
					argfun := arg0.WithFun().(func(env *Env) uint64)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(uint64) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uintptr:
				{
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(uintptr) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Float32:
				{
					argfun := arg0.WithFun().(func(env *Env) float32)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(float32) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Float64:
				{
					argfun := arg0.WithFun().(func(env *Env) float64)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(float64) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex64:
				{
					argfun := arg0.WithFun().(func(env *Env) complex64)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(complex64) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex128:
				{
					argfun := arg0.WithFun().(func(env *Env) complex128)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(complex128) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.String:
				{
					argfun := arg0.WithFun().(func(env *Env) string)
					call = func(env *Env) int16 {
						fun := exprfun(env).Interface().(func(string) int16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			default:
				call = func(env *Env) int16 {
					funv := exprfun(env)
					argv := []r.Value{
						argfun(env),
					}

					ret0 := funv.Call(argv)[0]
					return int16(ret0.Int())
				}
			}
		}
	case r.Int32:
		switch karg0 {
		case r.Bool:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(bool)
					switch funupn {
					case maxdepth:
						var cachedfun func(bool) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(bool) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) bool)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(bool) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(bool) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(bool) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(bool) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(bool) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int)
					switch funupn {
					case maxdepth:
						var cachedfun func(int) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(int) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(int) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(int) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int8)
					switch funupn {
					case maxdepth:
						var cachedfun func(int8) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int8) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int8) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int8) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(int8) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(int8) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(int8) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int16)
					switch funupn {
					case maxdepth:
						var cachedfun func(int16) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int16) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int16) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int16) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(int16) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(int16) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(int16) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int32)
					switch funupn {
					case maxdepth:
						var cachedfun func(int32) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int32) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int32) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int32) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(int32) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(int32) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(int32) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int64)
					switch funupn {
					case maxdepth:
						var cachedfun func(int64) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int64) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int64) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int64) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(int64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(int64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(int64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(uint) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(uint) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint8)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint8) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint8) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint8) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint8) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(uint8) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint8) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(uint8) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint16)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint16) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint16) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint16) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint16) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(uint16) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint16) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(uint16) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint32)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint32) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint32) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint32) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint32) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(uint32) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint32) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(uint32) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint64)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint64) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint64) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint64) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint64) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(uint64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(uint64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uintptr:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uintptr)
					switch funupn {
					case maxdepth:
						var cachedfun func(uintptr) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uintptr) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uintptr) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uintptr) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(uintptr) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(uintptr) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(uintptr) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float32)
					switch funupn {
					case maxdepth:
						var cachedfun func(float32) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float32) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float32) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float32) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(float32) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(float32) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(float32) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float64)
					switch funupn {
					case maxdepth:
						var cachedfun func(float64) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float64) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float64) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float64) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(float64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(float64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(float64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex64)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex64) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex64) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex64) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex64) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(complex64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(complex64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(complex64) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex128:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex128)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex128) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex128) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex128)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex128) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex128) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(complex128) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(complex128) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(complex128) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.String:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(string)
					switch funupn {
					case maxdepth:
						var cachedfun func(string) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int32)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(string) int32

						call = func(env *Env) int32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int32)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) string)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(string) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(string) int32

							call = func(env *Env) int32 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) int32)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int32 {
								fun := env.Binds[funindex].Interface().(func(string) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int32 {
								fun := env.Outer.Binds[funindex].Interface().(func(string) int32)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int32 {
								fun := exprfun(env).Interface().(func(string) int32)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		default:
			call = func(env *Env) int32 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return int32(ret0.Int())
			}
		}
	case r.Int64:
		switch karg0 {
		case r.Bool:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(bool)
					switch funupn {
					case maxdepth:
						var cachedfun func(bool) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(bool) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) bool)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(bool) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(bool) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(bool) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(bool) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(bool) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int)
					switch funupn {
					case maxdepth:
						var cachedfun func(int) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(int) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(int) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(int) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int8)
					switch funupn {
					case maxdepth:
						var cachedfun func(int8) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int8) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int8) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int8) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(int8) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(int8) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(int8) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int16)
					switch funupn {
					case maxdepth:
						var cachedfun func(int16) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int16) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int16) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int16) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(int16) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(int16) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(int16) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int32)
					switch funupn {
					case maxdepth:
						var cachedfun func(int32) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int32) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int32) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int32) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(int32) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(int32) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(int32) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int64)
					switch funupn {
					case maxdepth:
						var cachedfun func(int64) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int64) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int64) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int64) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(int64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(int64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(int64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(uint) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(uint) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint8)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint8) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint8) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint8) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint8) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(uint8) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint8) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(uint8) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint16)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint16) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint16) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint16) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint16) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(uint16) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint16) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(uint16) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint32)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint32) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint32) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint32) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint32) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(uint32) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint32) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(uint32) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint64)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint64) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint64) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint64) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint64) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(uint64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(uint64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uintptr:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uintptr)
					switch funupn {
					case maxdepth:
						var cachedfun func(uintptr) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uintptr) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uintptr) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uintptr) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(uintptr) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uintptr) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(uintptr) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float32)
					switch funupn {
					case maxdepth:
						var cachedfun func(float32) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float32) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float32) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float32) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(float32) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(float32) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(float32) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float64)
					switch funupn {
					case maxdepth:
						var cachedfun func(float64) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float64) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float64) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float64) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(float64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(float64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(float64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex64)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex64) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex64) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex64) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex64) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(complex64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(complex64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(complex64) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex128:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex128)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex128) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex128) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex128)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex128) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex128) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(complex128) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(complex128) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(complex128) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.String:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(string)
					switch funupn {
					case maxdepth:
						var cachedfun func(string) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(string) int64

						call = func(env *Env) int64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) int64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) string)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(string) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(string) int64

							call = func(env *Env) int64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) int64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) int64 {
								fun := env.Binds[funindex].Interface().(func(string) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) int64 {
								fun := env.Outer.Binds[funindex].Interface().(func(string) int64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) int64 {
								fun := exprfun(env).Interface().(func(string) int64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		default:
			call = func(env *Env) int64 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Int()
			}
		}

	case r.Uint:
		switch karg0 {
		case r.Bool:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(bool)
					switch funupn {
					case maxdepth:
						var cachedfun func(bool) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(bool) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) bool)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(bool) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(bool) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(bool) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(bool) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(bool) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int)
					switch funupn {
					case maxdepth:
						var cachedfun func(int) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(int) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(int) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(int) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int8)
					switch funupn {
					case maxdepth:
						var cachedfun func(int8) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int8) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int8) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int8) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(int8) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(int8) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(int8) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int16)
					switch funupn {
					case maxdepth:
						var cachedfun func(int16) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int16) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int16) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int16) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(int16) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(int16) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(int16) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int32)
					switch funupn {
					case maxdepth:
						var cachedfun func(int32) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int32) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int32) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int32) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(int32) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(int32) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(int32) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int64)
					switch funupn {
					case maxdepth:
						var cachedfun func(int64) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int64) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int64) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int64) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(int64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(int64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(int64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(uint) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(uint) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(uint) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint8)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint8) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint8) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint8) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint8) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(uint8) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(uint8) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(uint8) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint16)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint16) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint16) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint16) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint16) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(uint16) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(uint16) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(uint16) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint32)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint32) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint32) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint32) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint32) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(uint32) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(uint32) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(uint32) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint64)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint64) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint64) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint64) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint64) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(uint64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(uint64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(uint64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uintptr:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uintptr)
					switch funupn {
					case maxdepth:
						var cachedfun func(uintptr) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uintptr) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uintptr) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uintptr) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(uintptr) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(uintptr) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(uintptr) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float32)
					switch funupn {
					case maxdepth:
						var cachedfun func(float32) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float32) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float32) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float32) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(float32) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(float32) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(float32) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float64)
					switch funupn {
					case maxdepth:
						var cachedfun func(float64) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float64) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float64) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float64) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(float64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(float64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(float64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex64)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex64) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex64) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex64) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex64) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(complex64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(complex64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(complex64) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex128:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex128)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex128) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex128) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex128)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex128) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex128) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(complex128) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(complex128) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(complex128) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.String:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(string)
					switch funupn {
					case maxdepth:
						var cachedfun func(string) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(string) uint

						call = func(env *Env) uint {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) string)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(string) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(string) uint

							call = func(env *Env) uint {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) uint)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint {
								fun := env.Binds[funindex].Interface().(func(string) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint {
								fun := env.Outer.Binds[funindex].Interface().(func(string) uint)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint {
								fun := exprfun(env).Interface().(func(string) uint)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		default:
			call = func(env *Env) uint {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return uint(ret0.Uint())
			}
		}

	case r.Uint8:
		if karg0 == kret {

			if arg0.Const() && funsym != nil {
				argconst := arg0.Value.(uint8)
				switch funupn {
				case maxdepth:
					var cachedfun func(uint8) uint8

					call = func(env *Env) uint8 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8) uint8)
						}
						return cachedfun(argconst)
					}
				case maxdepth - 1:
					var cachedfun func(uint8) uint8

					call = func(env *Env) uint8 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8) uint8)
						}
						return cachedfun(argconst)
					}
				}
			}
			if call == nil {
				argfun := arg0.WithFun().(func(env *Env) uint8)
				if funsym != nil {
					switch funupn {
					case maxdepth:
						var cachedfun func(uint8) uint8

						call = func(env *Env) uint8 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint8)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					case maxdepth - 1:
						var cachedfun func(uint8) uint8

						call = func(env *Env) uint8 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint8)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					}
				}

				if call == nil {
					if funsym != nil && funupn == 0 {
						call = func(env *Env) uint8 {
							fun := env.Binds[funindex].Interface().(func(uint8) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funsym != nil && funupn == 1 {
						call = func(env *Env) uint8 {
							fun := env.Outer.Binds[funindex].Interface().(func(uint8) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint8 {
							fun := exprfun(env).Interface().(func(uint8) uint8)
							arg := argfun(env)
							return fun(arg)
						}
					}
				}

			}
		} else {
			switch karg0 {
			case r.Bool:

				{
					argfun := arg0.WithFun().(func(env *Env) bool)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(bool) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int:

				{
					argfun := arg0.WithFun().(func(env *Env) int)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(int) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int8:
				{
					argfun := arg0.WithFun().(func(env *Env) int8)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(int8) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int16:
				{
					argfun := arg0.WithFun().(func(env *Env) int16)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(int16) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int32:
				{
					argfun := arg0.WithFun().(func(env *Env) int32)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(int32) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int64:
				{
					argfun := arg0.WithFun().(func(env *Env) int64)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(int64) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint:
				{
					argfun := arg0.WithFun().(func(env *Env) uint)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(uint) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint8:
				{
					argfun := arg0.WithFun().(func(env *Env) uint8)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(uint8) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint16:
				{
					argfun := arg0.WithFun().(func(env *Env) uint16)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(uint16) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint32:
				{
					argfun := arg0.WithFun().(func(env *Env) uint32)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(uint32) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint64:
				{
					argfun := arg0.WithFun().(func(env *Env) uint64)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(uint64) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uintptr:
				{
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(uintptr) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Float32:
				{
					argfun := arg0.WithFun().(func(env *Env) float32)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(float32) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Float64:
				{
					argfun := arg0.WithFun().(func(env *Env) float64)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(float64) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex64:
				{
					argfun := arg0.WithFun().(func(env *Env) complex64)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(complex64) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex128:
				{
					argfun := arg0.WithFun().(func(env *Env) complex128)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(complex128) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.String:
				{
					argfun := arg0.WithFun().(func(env *Env) string)
					call = func(env *Env) uint8 {
						fun := exprfun(env).Interface().(func(string) uint8)
						arg := argfun(env)
						return fun(arg)
					}
				}

			default:
				call = func(env *Env) uint8 {
					funv := exprfun(env)
					argv := []r.Value{
						argfun(env),
					}

					ret0 := funv.Call(argv)[0]
					return uint8(ret0.Uint())
				}
			}
		}

	case r.Uint16:
		if karg0 == kret {

			if arg0.Const() && funsym != nil {
				argconst := arg0.Value.(uint16)
				switch funupn {
				case maxdepth:
					var cachedfun func(uint16) uint16

					call = func(env *Env) uint16 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16) uint16)
						}
						return cachedfun(argconst)
					}
				case maxdepth - 1:
					var cachedfun func(uint16) uint16

					call = func(env *Env) uint16 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16) uint16)
						}
						return cachedfun(argconst)
					}
				}
			}
			if call == nil {
				argfun := arg0.WithFun().(func(env *Env) uint16)
				if funsym != nil {
					switch funupn {
					case maxdepth:
						var cachedfun func(uint16) uint16

						call = func(env *Env) uint16 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint16)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					case maxdepth - 1:
						var cachedfun func(uint16) uint16

						call = func(env *Env) uint16 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint16)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					}
				}

				if call == nil {
					if funsym != nil && funupn == 0 {
						call = func(env *Env) uint16 {
							fun := env.Binds[funindex].Interface().(func(uint16) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funsym != nil && funupn == 1 {
						call = func(env *Env) uint16 {
							fun := env.Outer.Binds[funindex].Interface().(func(uint16) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint16 {
							fun := exprfun(env).Interface().(func(uint16) uint16)
							arg := argfun(env)
							return fun(arg)
						}
					}
				}

			}
		} else {
			switch karg0 {
			case r.Bool:

				{
					argfun := arg0.WithFun().(func(env *Env) bool)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(bool) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int:

				{
					argfun := arg0.WithFun().(func(env *Env) int)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(int) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int8:
				{
					argfun := arg0.WithFun().(func(env *Env) int8)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(int8) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int16:
				{
					argfun := arg0.WithFun().(func(env *Env) int16)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(int16) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int32:
				{
					argfun := arg0.WithFun().(func(env *Env) int32)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(int32) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int64:
				{
					argfun := arg0.WithFun().(func(env *Env) int64)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(int64) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint:
				{
					argfun := arg0.WithFun().(func(env *Env) uint)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(uint) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint8:
				{
					argfun := arg0.WithFun().(func(env *Env) uint8)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(uint8) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint16:
				{
					argfun := arg0.WithFun().(func(env *Env) uint16)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(uint16) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint32:
				{
					argfun := arg0.WithFun().(func(env *Env) uint32)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(uint32) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint64:
				{
					argfun := arg0.WithFun().(func(env *Env) uint64)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(uint64) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uintptr:
				{
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(uintptr) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Float32:
				{
					argfun := arg0.WithFun().(func(env *Env) float32)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(float32) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Float64:
				{
					argfun := arg0.WithFun().(func(env *Env) float64)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(float64) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex64:
				{
					argfun := arg0.WithFun().(func(env *Env) complex64)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(complex64) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex128:
				{
					argfun := arg0.WithFun().(func(env *Env) complex128)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(complex128) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.String:
				{
					argfun := arg0.WithFun().(func(env *Env) string)
					call = func(env *Env) uint16 {
						fun := exprfun(env).Interface().(func(string) uint16)
						arg := argfun(env)
						return fun(arg)
					}
				}

			default:
				call = func(env *Env) uint16 {
					funv := exprfun(env)
					argv := []r.Value{
						argfun(env),
					}

					ret0 := funv.Call(argv)[0]
					return uint16(ret0.Uint())
				}
			}
		}

	case r.Uint32:
		if karg0 == kret {

			if arg0.Const() && funsym != nil {
				argconst := arg0.Value.(uint32)
				switch funupn {
				case maxdepth:
					var cachedfun func(uint32) uint32

					call = func(env *Env) uint32 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32) uint32)
						}
						return cachedfun(argconst)
					}
				case maxdepth - 1:
					var cachedfun func(uint32) uint32

					call = func(env *Env) uint32 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32) uint32)
						}
						return cachedfun(argconst)
					}
				}
			}
			if call == nil {
				argfun := arg0.WithFun().(func(env *Env) uint32)
				if funsym != nil {
					switch funupn {
					case maxdepth:
						var cachedfun func(uint32) uint32

						call = func(env *Env) uint32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint32)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					case maxdepth - 1:
						var cachedfun func(uint32) uint32

						call = func(env *Env) uint32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint32)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					}
				}

				if call == nil {
					if funsym != nil && funupn == 0 {
						call = func(env *Env) uint32 {
							fun := env.Binds[funindex].Interface().(func(uint32) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funsym != nil && funupn == 1 {
						call = func(env *Env) uint32 {
							fun := env.Outer.Binds[funindex].Interface().(func(uint32) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uint32 {
							fun := exprfun(env).Interface().(func(uint32) uint32)
							arg := argfun(env)
							return fun(arg)
						}
					}
				}

			}
		} else {
			switch karg0 {
			case r.Bool:

				{
					argfun := arg0.WithFun().(func(env *Env) bool)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(bool) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int:

				{
					argfun := arg0.WithFun().(func(env *Env) int)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(int) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int8:
				{
					argfun := arg0.WithFun().(func(env *Env) int8)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(int8) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int16:
				{
					argfun := arg0.WithFun().(func(env *Env) int16)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(int16) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int32:
				{
					argfun := arg0.WithFun().(func(env *Env) int32)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(int32) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int64:
				{
					argfun := arg0.WithFun().(func(env *Env) int64)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(int64) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint:
				{
					argfun := arg0.WithFun().(func(env *Env) uint)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(uint) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint8:
				{
					argfun := arg0.WithFun().(func(env *Env) uint8)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(uint8) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint16:
				{
					argfun := arg0.WithFun().(func(env *Env) uint16)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(uint16) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint32:
				{
					argfun := arg0.WithFun().(func(env *Env) uint32)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(uint32) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint64:
				{
					argfun := arg0.WithFun().(func(env *Env) uint64)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(uint64) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uintptr:
				{
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(uintptr) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Float32:
				{
					argfun := arg0.WithFun().(func(env *Env) float32)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(float32) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Float64:
				{
					argfun := arg0.WithFun().(func(env *Env) float64)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(float64) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex64:
				{
					argfun := arg0.WithFun().(func(env *Env) complex64)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(complex64) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex128:
				{
					argfun := arg0.WithFun().(func(env *Env) complex128)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(complex128) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.String:
				{
					argfun := arg0.WithFun().(func(env *Env) string)
					call = func(env *Env) uint32 {
						fun := exprfun(env).Interface().(func(string) uint32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			default:
				call = func(env *Env) uint32 {
					funv := exprfun(env)
					argv := []r.Value{
						argfun(env),
					}

					ret0 := funv.Call(argv)[0]
					return uint32(ret0.Uint())
				}
			}
		}

	case r.Uint64:
		switch karg0 {
		case r.Bool:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(bool)
					switch funupn {
					case maxdepth:
						var cachedfun func(bool) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(bool) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) bool)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(bool) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(bool) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(bool) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(bool) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(bool) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int)
					switch funupn {
					case maxdepth:
						var cachedfun func(int) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(int) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(int) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(int) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int8)
					switch funupn {
					case maxdepth:
						var cachedfun func(int8) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int8) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int8) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int8) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(int8) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(int8) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(int8) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int16)
					switch funupn {
					case maxdepth:
						var cachedfun func(int16) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int16) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int16) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int16) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(int16) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(int16) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(int16) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int32)
					switch funupn {
					case maxdepth:
						var cachedfun func(int32) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int32) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int32) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int32) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(int32) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(int32) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(int32) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int64)
					switch funupn {
					case maxdepth:
						var cachedfun func(int64) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int64) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int64) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int64) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(int64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(int64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(int64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(uint) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(uint) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint8)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint8) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint8) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint8) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint8) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(uint8) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint8) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(uint8) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint16)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint16) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint16) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint16) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint16) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(uint16) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint16) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(uint16) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint32)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint32) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint32) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint32) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint32) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(uint32) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint32) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(uint32) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint64)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint64) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint64) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint64) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint64) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(uint64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uint64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(uint64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uintptr:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uintptr)
					switch funupn {
					case maxdepth:
						var cachedfun func(uintptr) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uintptr) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uintptr) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uintptr) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(uintptr) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(uintptr) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(uintptr) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float32)
					switch funupn {
					case maxdepth:
						var cachedfun func(float32) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float32) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float32) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float32) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(float32) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(float32) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(float32) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float64)
					switch funupn {
					case maxdepth:
						var cachedfun func(float64) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float64) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float64) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float64) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(float64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(float64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(float64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex64)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex64) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex64) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex64) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex64) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(complex64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(complex64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(complex64) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex128:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex128)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex128) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex128) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex128)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex128) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex128) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(complex128) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(complex128) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(complex128) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.String:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(string)
					switch funupn {
					case maxdepth:
						var cachedfun func(string) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint64)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(string) uint64

						call = func(env *Env) uint64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) uint64)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) string)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(string) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(string) uint64

							call = func(env *Env) uint64 {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) uint64)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) uint64 {
								fun := env.Binds[funindex].Interface().(func(string) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) uint64 {
								fun := env.Outer.Binds[funindex].Interface().(func(string) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) uint64 {
								fun := exprfun(env).Interface().(func(string) uint64)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		default:
			call = func(env *Env) uint64 {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return ret0.Uint()
			}
		}

	case r.Uintptr:
		if karg0 == kret {

			if arg0.Const() && funsym != nil {
				argconst := arg0.Value.(uintptr)
				switch funupn {
				case maxdepth:
					var cachedfun func(uintptr) uintptr

					call = func(env *Env) uintptr {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr) uintptr)
						}
						return cachedfun(argconst)
					}
				case maxdepth - 1:
					var cachedfun func(uintptr) uintptr

					call = func(env *Env) uintptr {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr) uintptr)
						}
						return cachedfun(argconst)
					}
				}
			}
			if call == nil {
				argfun := arg0.WithFun().(func(env *Env) uintptr)
				if funsym != nil {
					switch funupn {
					case maxdepth:
						var cachedfun func(uintptr) uintptr

						call = func(env *Env) uintptr {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uintptr)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					case maxdepth - 1:
						var cachedfun func(uintptr) uintptr

						call = func(env *Env) uintptr {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) uintptr)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					}
				}

				if call == nil {
					if funsym != nil && funupn == 0 {
						call = func(env *Env) uintptr {
							fun := env.Binds[funindex].Interface().(func(uintptr) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funsym != nil && funupn == 1 {
						call = func(env *Env) uintptr {
							fun := env.Outer.Binds[funindex].Interface().(func(uintptr) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) uintptr {
							fun := exprfun(env).Interface().(func(uintptr) uintptr)
							arg := argfun(env)
							return fun(arg)
						}
					}
				}

			}
		} else {
			switch karg0 {
			case r.Bool:

				{
					argfun := arg0.WithFun().(func(env *Env) bool)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(bool) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int:

				{
					argfun := arg0.WithFun().(func(env *Env) int)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(int) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int8:
				{
					argfun := arg0.WithFun().(func(env *Env) int8)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(int8) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int16:
				{
					argfun := arg0.WithFun().(func(env *Env) int16)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(int16) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int32:
				{
					argfun := arg0.WithFun().(func(env *Env) int32)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(int32) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int64:
				{
					argfun := arg0.WithFun().(func(env *Env) int64)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(int64) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint:
				{
					argfun := arg0.WithFun().(func(env *Env) uint)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(uint) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint8:
				{
					argfun := arg0.WithFun().(func(env *Env) uint8)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(uint8) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint16:
				{
					argfun := arg0.WithFun().(func(env *Env) uint16)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(uint16) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint32:
				{
					argfun := arg0.WithFun().(func(env *Env) uint32)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(uint32) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint64:
				{
					argfun := arg0.WithFun().(func(env *Env) uint64)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(uint64) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uintptr:
				{
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(uintptr) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Float32:
				{
					argfun := arg0.WithFun().(func(env *Env) float32)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(float32) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Float64:
				{
					argfun := arg0.WithFun().(func(env *Env) float64)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(float64) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex64:
				{
					argfun := arg0.WithFun().(func(env *Env) complex64)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(complex64) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex128:
				{
					argfun := arg0.WithFun().(func(env *Env) complex128)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(complex128) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.String:
				{
					argfun := arg0.WithFun().(func(env *Env) string)
					call = func(env *Env) uintptr {
						fun := exprfun(env).Interface().(func(string) uintptr)
						arg := argfun(env)
						return fun(arg)
					}
				}

			default:
				call = func(env *Env) uintptr {
					funv := exprfun(env)
					argv := []r.Value{
						argfun(env),
					}

					ret0 := funv.Call(argv)[0]
					return uintptr(ret0.Uint())
				}
			}
		}

	case r.Float32:
		if karg0 == kret {

			if arg0.Const() && funsym != nil {
				argconst := arg0.Value.(float32)
				switch funupn {
				case maxdepth:
					var cachedfun func(float32) float32

					call = func(env *Env) float32 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32) float32)
						}
						return cachedfun(argconst)
					}
				case maxdepth - 1:
					var cachedfun func(float32) float32

					call = func(env *Env) float32 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32) float32)
						}
						return cachedfun(argconst)
					}
				}
			}
			if call == nil {
				argfun := arg0.WithFun().(func(env *Env) float32)
				if funsym != nil {
					switch funupn {
					case maxdepth:
						var cachedfun func(float32) float32

						call = func(env *Env) float32 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) float32)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					case maxdepth - 1:
						var cachedfun func(float32) float32

						call = func(env *Env) float32 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) float32)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					}
				}

				if call == nil {
					if funsym != nil && funupn == 0 {
						call = func(env *Env) float32 {
							fun := env.Binds[funindex].Interface().(func(float32) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funsym != nil && funupn == 1 {
						call = func(env *Env) float32 {
							fun := env.Outer.Binds[funindex].Interface().(func(float32) float32)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float32 {
							fun := exprfun(env).Interface().(func(float32) float32)
							arg := argfun(env)
							return fun(arg)
						}
					}
				}

			}
		} else {
			switch karg0 {
			case r.Bool:

				{
					argfun := arg0.WithFun().(func(env *Env) bool)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(bool) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int:

				{
					argfun := arg0.WithFun().(func(env *Env) int)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(int) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int8:
				{
					argfun := arg0.WithFun().(func(env *Env) int8)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(int8) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int16:
				{
					argfun := arg0.WithFun().(func(env *Env) int16)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(int16) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int32:
				{
					argfun := arg0.WithFun().(func(env *Env) int32)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(int32) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int64:
				{
					argfun := arg0.WithFun().(func(env *Env) int64)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(int64) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint:
				{
					argfun := arg0.WithFun().(func(env *Env) uint)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(uint) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint8:
				{
					argfun := arg0.WithFun().(func(env *Env) uint8)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(uint8) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint16:
				{
					argfun := arg0.WithFun().(func(env *Env) uint16)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(uint16) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint32:
				{
					argfun := arg0.WithFun().(func(env *Env) uint32)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(uint32) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint64:
				{
					argfun := arg0.WithFun().(func(env *Env) uint64)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(uint64) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uintptr:
				{
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(uintptr) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Float32:
				{
					argfun := arg0.WithFun().(func(env *Env) float32)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(float32) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Float64:
				{
					argfun := arg0.WithFun().(func(env *Env) float64)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(float64) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex64:
				{
					argfun := arg0.WithFun().(func(env *Env) complex64)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(complex64) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex128:
				{
					argfun := arg0.WithFun().(func(env *Env) complex128)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(complex128) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.String:
				{
					argfun := arg0.WithFun().(func(env *Env) string)
					call = func(env *Env) float32 {
						fun := exprfun(env).Interface().(func(string) float32)
						arg := argfun(env)
						return fun(arg)
					}
				}

			default:
				call = func(env *Env) float32 {
					funv := exprfun(env)
					argv := []r.Value{
						argfun(env),
					}

					ret0 := funv.Call(argv)[0]
					return float32(ret0.Float())
				}
			}
		}

	case r.Float64:
		if karg0 == kret {

			if arg0.Const() && funsym != nil {
				argconst := arg0.Value.(float64)
				switch funupn {
				case maxdepth:
					var cachedfun func(float64) float64

					call = func(env *Env) float64 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64) float64)
						}
						return cachedfun(argconst)
					}
				case maxdepth - 1:
					var cachedfun func(float64) float64

					call = func(env *Env) float64 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64) float64)
						}
						return cachedfun(argconst)
					}
				}
			}
			if call == nil {
				argfun := arg0.WithFun().(func(env *Env) float64)
				if funsym != nil {
					switch funupn {
					case maxdepth:
						var cachedfun func(float64) float64

						call = func(env *Env) float64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) float64)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					case maxdepth - 1:
						var cachedfun func(float64) float64

						call = func(env *Env) float64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) float64)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					}
				}

				if call == nil {
					if funsym != nil && funupn == 0 {
						call = func(env *Env) float64 {
							fun := env.Binds[funindex].Interface().(func(float64) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funsym != nil && funupn == 1 {
						call = func(env *Env) float64 {
							fun := env.Outer.Binds[funindex].Interface().(func(float64) float64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) float64 {
							fun := exprfun(env).Interface().(func(float64) float64)
							arg := argfun(env)
							return fun(arg)
						}
					}
				}

			}
		} else {
			switch karg0 {
			case r.Bool:

				{
					argfun := arg0.WithFun().(func(env *Env) bool)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(bool) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int:

				{
					argfun := arg0.WithFun().(func(env *Env) int)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(int) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int8:
				{
					argfun := arg0.WithFun().(func(env *Env) int8)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(int8) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int16:
				{
					argfun := arg0.WithFun().(func(env *Env) int16)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(int16) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int32:
				{
					argfun := arg0.WithFun().(func(env *Env) int32)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(int32) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int64:
				{
					argfun := arg0.WithFun().(func(env *Env) int64)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(int64) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint:
				{
					argfun := arg0.WithFun().(func(env *Env) uint)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(uint) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint8:
				{
					argfun := arg0.WithFun().(func(env *Env) uint8)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(uint8) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint16:
				{
					argfun := arg0.WithFun().(func(env *Env) uint16)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(uint16) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint32:
				{
					argfun := arg0.WithFun().(func(env *Env) uint32)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(uint32) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint64:
				{
					argfun := arg0.WithFun().(func(env *Env) uint64)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(uint64) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uintptr:
				{
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(uintptr) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Float32:
				{
					argfun := arg0.WithFun().(func(env *Env) float32)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(float32) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Float64:
				{
					argfun := arg0.WithFun().(func(env *Env) float64)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(float64) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex64:
				{
					argfun := arg0.WithFun().(func(env *Env) complex64)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(complex64) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex128:
				{
					argfun := arg0.WithFun().(func(env *Env) complex128)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(complex128) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.String:
				{
					argfun := arg0.WithFun().(func(env *Env) string)
					call = func(env *Env) float64 {
						fun := exprfun(env).Interface().(func(string) float64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			default:
				call = func(env *Env) float64 {
					funv := exprfun(env)
					argv := []r.Value{
						argfun(env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Float()
				}
			}
		}

	case r.Complex64:
		if karg0 == kret {

			if arg0.Const() && funsym != nil {
				argconst := arg0.Value.(complex64)
				switch funupn {
				case maxdepth:
					var cachedfun func(complex64) complex64

					call = func(env *Env) complex64 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64) complex64)
						}
						return cachedfun(argconst)
					}
				case maxdepth - 1:
					var cachedfun func(complex64) complex64

					call = func(env *Env) complex64 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64) complex64)
						}
						return cachedfun(argconst)
					}
				}
			}
			if call == nil {
				argfun := arg0.WithFun().(func(env *Env) complex64)
				if funsym != nil {
					switch funupn {
					case maxdepth:
						var cachedfun func(complex64) complex64

						call = func(env *Env) complex64 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) complex64)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					case maxdepth - 1:
						var cachedfun func(complex64) complex64

						call = func(env *Env) complex64 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) complex64)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					}
				}

				if call == nil {
					if funsym != nil && funupn == 0 {
						call = func(env *Env) complex64 {
							fun := env.Binds[funindex].Interface().(func(complex64) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funsym != nil && funupn == 1 {
						call = func(env *Env) complex64 {
							fun := env.Outer.Binds[funindex].Interface().(func(complex64) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex64 {
							fun := exprfun(env).Interface().(func(complex64) complex64)
							arg := argfun(env)
							return fun(arg)
						}
					}
				}

			}
		} else {
			switch karg0 {
			case r.Bool:

				{
					argfun := arg0.WithFun().(func(env *Env) bool)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(bool) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int:

				{
					argfun := arg0.WithFun().(func(env *Env) int)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(int) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int8:
				{
					argfun := arg0.WithFun().(func(env *Env) int8)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(int8) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int16:
				{
					argfun := arg0.WithFun().(func(env *Env) int16)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(int16) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int32:
				{
					argfun := arg0.WithFun().(func(env *Env) int32)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(int32) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int64:
				{
					argfun := arg0.WithFun().(func(env *Env) int64)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(int64) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint:
				{
					argfun := arg0.WithFun().(func(env *Env) uint)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(uint) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint8:
				{
					argfun := arg0.WithFun().(func(env *Env) uint8)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(uint8) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint16:
				{
					argfun := arg0.WithFun().(func(env *Env) uint16)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(uint16) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint32:
				{
					argfun := arg0.WithFun().(func(env *Env) uint32)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(uint32) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint64:
				{
					argfun := arg0.WithFun().(func(env *Env) uint64)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(uint64) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uintptr:
				{
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(uintptr) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Float32:
				{
					argfun := arg0.WithFun().(func(env *Env) float32)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(float32) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Float64:
				{
					argfun := arg0.WithFun().(func(env *Env) float64)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(float64) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex64:
				{
					argfun := arg0.WithFun().(func(env *Env) complex64)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(complex64) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex128:
				{
					argfun := arg0.WithFun().(func(env *Env) complex128)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(complex128) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.String:
				{
					argfun := arg0.WithFun().(func(env *Env) string)
					call = func(env *Env) complex64 {
						fun := exprfun(env).Interface().(func(string) complex64)
						arg := argfun(env)
						return fun(arg)
					}
				}

			default:
				call = func(env *Env) complex64 {
					funv := exprfun(env)
					argv := []r.Value{
						argfun(env),
					}

					ret0 := funv.Call(argv)[0]
					return complex64(ret0.Complex())
				}
			}
		}

	case r.Complex128:
		if karg0 == kret {

			if arg0.Const() && funsym != nil {
				argconst := arg0.Value.(complex128)
				switch funupn {
				case maxdepth:
					var cachedfun func(complex128) complex128

					call = func(env *Env) complex128 {
						funv := env.ThreadGlobals.TopEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128) complex128)
						}
						return cachedfun(argconst)
					}
				case maxdepth - 1:
					var cachedfun func(complex128) complex128

					call = func(env *Env) complex128 {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128) complex128)
						}
						return cachedfun(argconst)
					}
				}
			}
			if call == nil {
				argfun := arg0.WithFun().(func(env *Env) complex128)
				if funsym != nil {
					switch funupn {
					case maxdepth:
						var cachedfun func(complex128) complex128

						call = func(env *Env) complex128 {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) complex128)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					case maxdepth - 1:
						var cachedfun func(complex128) complex128

						call = func(env *Env) complex128 {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) complex128)
							}

							arg := argfun(env)
							return cachedfun(arg)
						}
					}
				}

				if call == nil {
					if funsym != nil && funupn == 0 {
						call = func(env *Env) complex128 {
							fun := env.Binds[funindex].Interface().(func(complex128) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else if funsym != nil && funupn == 1 {
						call = func(env *Env) complex128 {
							fun := env.Outer.Binds[funindex].Interface().(func(complex128) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					} else {
						call = func(env *Env) complex128 {
							fun := exprfun(env).Interface().(func(complex128) complex128)
							arg := argfun(env)
							return fun(arg)
						}
					}
				}

			}
		} else {
			switch karg0 {
			case r.Bool:

				{
					argfun := arg0.WithFun().(func(env *Env) bool)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(bool) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int:

				{
					argfun := arg0.WithFun().(func(env *Env) int)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(int) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int8:
				{
					argfun := arg0.WithFun().(func(env *Env) int8)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(int8) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int16:
				{
					argfun := arg0.WithFun().(func(env *Env) int16)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(int16) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int32:
				{
					argfun := arg0.WithFun().(func(env *Env) int32)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(int32) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Int64:
				{
					argfun := arg0.WithFun().(func(env *Env) int64)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(int64) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint:
				{
					argfun := arg0.WithFun().(func(env *Env) uint)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(uint) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint8:
				{
					argfun := arg0.WithFun().(func(env *Env) uint8)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(uint8) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint16:
				{
					argfun := arg0.WithFun().(func(env *Env) uint16)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(uint16) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint32:
				{
					argfun := arg0.WithFun().(func(env *Env) uint32)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(uint32) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uint64:
				{
					argfun := arg0.WithFun().(func(env *Env) uint64)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(uint64) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Uintptr:
				{
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(uintptr) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}
			case r.Float32:
				{
					argfun := arg0.WithFun().(func(env *Env) float32)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(float32) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Float64:
				{
					argfun := arg0.WithFun().(func(env *Env) float64)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(float64) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex64:
				{
					argfun := arg0.WithFun().(func(env *Env) complex64)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(complex64) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.Complex128:
				{
					argfun := arg0.WithFun().(func(env *Env) complex128)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(complex128) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}

			case r.String:
				{
					argfun := arg0.WithFun().(func(env *Env) string)
					call = func(env *Env) complex128 {
						fun := exprfun(env).Interface().(func(string) complex128)
						arg := argfun(env)
						return fun(arg)
					}
				}

			default:
				call = func(env *Env) complex128 {
					funv := exprfun(env)
					argv := []r.Value{
						argfun(env),
					}

					ret0 := funv.Call(argv)[0]
					return ret0.Complex()
				}
			}
		}

	case r.String:
		switch karg0 {
		case r.Bool:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(bool)
					switch funupn {
					case maxdepth:
						var cachedfun func(bool) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(bool) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(bool) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) bool)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(bool) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(bool) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(bool) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(bool) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(bool) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(bool) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int)
					switch funupn {
					case maxdepth:
						var cachedfun func(int) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(int) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(int) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(int) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int8)
					switch funupn {
					case maxdepth:
						var cachedfun func(int8) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int8) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int8) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int8) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int8) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int8) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(int8) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(int8) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(int8) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int16)
					switch funupn {
					case maxdepth:
						var cachedfun func(int16) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int16) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int16) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int16) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int16) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int16) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(int16) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(int16) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(int16) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int32)
					switch funupn {
					case maxdepth:
						var cachedfun func(int32) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int32) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int32) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int32) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int32) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int32) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(int32) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(int32) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(int32) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Int64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(int64)
					switch funupn {
					case maxdepth:
						var cachedfun func(int64) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(int64) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(int64) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) int64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(int64) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(int64) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(int64) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(int64) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(int64) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(int64) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(uint) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(uint) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(uint) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint8:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint8)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint8) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint8) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint8) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint8)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint8) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint8) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint8) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(uint8) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(uint8) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(uint8) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint16:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint16)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint16) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint16) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint16) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint16)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint16) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint16) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint16) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(uint16) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(uint16) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(uint16) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint32)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint32) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint32) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint32) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint32) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint32) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint32) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(uint32) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(uint32) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(uint32) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uint64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uint64)
					switch funupn {
					case maxdepth:
						var cachedfun func(uint64) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uint64) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uint64) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uint64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uint64) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uint64) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uint64) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(uint64) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(uint64) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(uint64) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Uintptr:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(uintptr)
					switch funupn {
					case maxdepth:
						var cachedfun func(uintptr) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(uintptr) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(uintptr) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) uintptr)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(uintptr) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(uintptr) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(uintptr) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(uintptr) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(uintptr) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(uintptr) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float32:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float32)
					switch funupn {
					case maxdepth:
						var cachedfun func(float32) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float32) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float32) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float32)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float32) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float32) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float32) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(float32) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(float32) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(float32) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Float64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(float64)
					switch funupn {
					case maxdepth:
						var cachedfun func(float64) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(float64) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(float64) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) float64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(float64) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(float64) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(float64) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(float64) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(float64) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(float64) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex64:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex64)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex64) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex64) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex64) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex64)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex64) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex64) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex64) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(complex64) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(complex64) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(complex64) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.Complex128:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(complex128)
					switch funupn {
					case maxdepth:
						var cachedfun func(complex128) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(complex128) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(complex128) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) complex128)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(complex128) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(complex128) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(complex128) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(complex128) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(complex128) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(complex128) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		case r.String:

			{

				if arg0.Const() && funsym != nil {
					argconst := arg0.Value.(string)
					switch funupn {
					case maxdepth:
						var cachedfun func(string) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.TopEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) string)
							}
							return cachedfun(argconst)
						}
					case maxdepth - 1:
						var cachedfun func(string) string

						call = func(env *Env) string {
							funv := env.ThreadGlobals.FileEnv.Binds[funindex]
							if cachedfunv != funv {
								cachedfunv = funv
								cachedfun = funv.Interface().(func(string) string)
							}
							return cachedfun(argconst)
						}
					}
				}
				if call == nil {
					argfun := arg0.WithFun().(func(env *Env) string)
					if funsym != nil {
						switch funupn {
						case maxdepth:
							var cachedfun func(string) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.TopEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						case maxdepth - 1:
							var cachedfun func(string) string

							call = func(env *Env) string {
								funv := env.ThreadGlobals.FileEnv.Binds[funindex]
								if cachedfunv != funv {
									cachedfunv = funv
									cachedfun = funv.Interface().(func(string) string)
								}

								arg := argfun(env)
								return cachedfun(arg)
							}
						}
					}

					if call == nil {
						if funsym != nil && funupn == 0 {
							call = func(env *Env) string {
								fun := env.Binds[funindex].Interface().(func(string) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else if funsym != nil && funupn == 1 {
							call = func(env *Env) string {
								fun := env.Outer.Binds[funindex].Interface().(func(string) string)
								arg := argfun(env)
								return fun(arg)
							}
						} else {
							call = func(env *Env) string {
								fun := exprfun(env).Interface().(func(string) string)
								arg := argfun(env)
								return fun(arg)
							}
						}
					}

				}
			}
		default:
			call = func(env *Env) string {
				funv := exprfun(env)
				argv := []r.Value{
					argfun(env),
				}

				ret0 := funv.Call(argv)[0]
				return ret0.String()
			}
		}

	}
	if call == nil {
		call = func(env *Env) r.Value {
			funv := exprfun(env)
			argv := []r.Value{
				argfun(env),
			}
			return funv.Call(argv)[0]
		}
	}
	return call
}
