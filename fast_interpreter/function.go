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
 * function.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
	r "reflect"
)

func (c *Comp) DeclFunc(decl_or_nil *ast.FuncDecl, funcType *ast.FuncType, body *ast.BlockStmt) {
	var declRecv *ast.Field
	if decl_or_nil != nil && decl_or_nil.Recv != nil {
		n := len(decl_or_nil.Recv.List)
		if n != 1 {
			c.Errorf("invalid function/method declaration: expecting one receiver or nil, found %d receivers: func %v %s(/*...*/)",
				n, decl_or_nil.Recv, decl_or_nil.Name)
			return
		}
		declRecv = decl_or_nil.Recv.List[0]
	}
	_, t, paramNames, resultNames := c.TypeFunctionOrMethod(declRecv, funcType)
	cf := NewComp(c)
	for i, n := 0, t.NumIn(); i < n; i++ {
		// paramNames[i] == "_" means that argument is not used inside the function.
		// DeclVar0 will not allocate a bind for it - correct optimization...
		// just remember to check for such case when compiling calls to the function
		cf.DeclVar0(paramNames[i], t.In(i), nil)
	}
	for i, n := 0, t.NumOut(); i < n; i++ {
		// resultNames[i] == "_" means that result is unnamed.
		// we must still allocate a bind for it.
		resultName := resultNames[i]
		if resultName == "_" {
			resultName = ""
		}
		cf.DeclVar0(resultName, t.Out(i), nil)
	}
	// TODO copy runtime arguments into allocated binds
	if body != nil {
		// in Go, function arguments/results and function body are in the same scope
		cf.Block0(body.List...)
	}
	// TODO read allocated returns and return them

	c.Errorf("function declaration is not implemented, found: %v <%v>", decl_or_nil, r.TypeOf(decl_or_nil))
}

/*
func (c *Comp) DeclFunc(name string, paramTypes []r.Type, body X) X {
	idx := c.AddBind(name, typeOfInterface) // FIXME need accurate function type
	xf := c.MakeFunc(paramTypes, body)
	return func(env *Env) (r.Value, []r.Value) {
		f := xf(env)
		env.Binds[idx] = r.ValueOf(f)
		return r.ValueOf(f), nil
	}
}

func (c *Comp) MakeFunc(paramTypes []r.Type, body X) XFunc {
	return func(env *Env) Func {
		return func(args ...r.Value) (ret r.Value, rets []r.Value) {
			fenv := NewEnv(env, 10)
			panicking := true // use a flag to distinguish non-panic from panic(nil)
			defer func() {
				if panicking {
					pan := recover()
					switch p := pan.(type) {
					case SReturn:
						// return is implemented with a panic(SReturn{})
						ret = p.result0
						rets = p.results
					default:
						panic(pan)
					}
				}
			}()
			for i, paramType := range paramTypes {
				place := r.New(paramType).Elem()
				place.Set(args[i].convert(paramType))
				fenv.Binds[i] = place
			}
			ret, rets = body(fenv)
			panicking = false
			return ret, rets
		}
	}
}

func MakeFuncInt(paramTypes []r.Type, body X) XFuncInt {
	return func(env *Env) FuncInt {
		return func(args ...r.Value) (ret int) {
			fenv := NewEnv(env, 10)
			panicking := true // use a flag to distinguish non-panic from panic(nil)
			defer func() {
				if panicking {
					pan := recover()
					switch p := pan.(type) {
					case SReturn:
						// return is implemented with a panic(cReturn{})
						ret = int(p.result0.Int())
					default:
						panic(pan)
					}
				}
			}()
			for i, paramType := range paramTypes {
				place := r.New(paramType).Elem()
				place.Set(args[i].convert(paramType))
				fenv.Binds[i] = place
			}
			ret0, _ := body(fenv)
			panicking = false
			return int(ret0.Int())
		}
	}
}
*/
