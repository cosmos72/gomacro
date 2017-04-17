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
 * compile.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
	r "reflect"
	"strings"

	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
)

func NewInterpreterCommon() *InterpreterCommon {
	b := MakeInterpreterBase()
	return &InterpreterCommon{
		InterpreterBase: &b,
	}
}

func New() *CompEnv {
	top := NewCompEnv(nil, "builtin")
	top.growEnv(128)
	file := NewCompEnv(top, "main")
	file.growEnv(1024)
	return file
}

func NewCompEnv(outer *CompEnv, path string) *CompEnv {
	name := path[1+strings.LastIndexByte(path, '/'):]

	var outerComp *Comp
	var outerEnv *Env
	var common *InterpreterCommon
	if outer != nil {
		outerComp = outer.Comp
		outerEnv = outer.Env
		common = outer.InterpreterCommon
	}
	if common == nil {
		common = NewInterpreterCommon()
	}
	c := &CompEnv{
		Comp: &Comp{
			Outer:             outerComp,
			Name:              name,
			Path:              path,
			InterpreterCommon: common,
		},
		Env: &Env{
			Outer:  outerEnv,
			Common: common,
		},
	}
	if outer == nil {
		c.addBuiltins()
	}
	return c

}

func NewComp(outer *Comp) *Comp {
	if outer == nil {
		return &Comp{UpCost: 1}
	}
	return &Comp{
		UpCost:            1,
		Code:              outer.Code,
		Outer:             outer,
		CompileOptions:    outer.CompileOptions,
		InterpreterCommon: outer.InterpreterCommon,
	}
}

func (c *Comp) Top() *Comp {
	for ; c != nil; c = c.Outer {
		if c.Outer == nil {
			break
		}
	}
	return c
}

func (c *Comp) File() *Comp {
	for ; c != nil; c = c.Outer {
		outer := c.Outer
		if outer == nil || outer.Outer == nil {
			break
		}
	}
	return c
}

// if a function Env only declares ignored binds, it gets this scratch buffers
var ignoredBinds = []r.Value{Nil}
var ignoredIntBinds = []uint64{0}

func NewEnv(outer *Env, nbinds int, nintbinds int) *Env {
	common := outer.Common
	if env := common.allocEnv(nbinds, nintbinds); env != nil {
		env.Outer = outer
		env.IP = outer.IP
		env.Code = outer.Code
		env.Interrupt = outer.Interrupt
		return env
	}

	env := &Env{
		Outer:     outer,
		IP:        outer.IP,
		Code:      outer.Code,
		Interrupt: outer.Interrupt,
		Common:    common,
	}
	if nbinds <= 1 {
		env.Binds = ignoredBinds
	} else {
		env.Binds = make([]r.Value, nbinds)
	}
	if nintbinds <= 1 {
		env.IntBinds = ignoredIntBinds
	} else {
		env.IntBinds = make([]uint64, nintbinds)
	}
	return env
}

func NewEnv4Func(outer *Env, nbinds int, nintbinds int) *Env {
	common := outer.Common

	var env *Env
	if env = common.allocEnv(nbinds, nintbinds); env != nil {
		env.Outer = outer
		return env
	}
	env = &Env{
		Outer:  outer,
		Common: common,
	}
	if nbinds <= 1 {
		env.Binds = ignoredBinds
	} else {
		env.Binds = make([]r.Value, nbinds)
	}
	if nintbinds <= 1 {
		env.IntBinds = ignoredIntBinds
	} else {
		env.IntBinds = make([]uint64, nintbinds)
	}
	return env
}

func (common *InterpreterCommon) allocEnv(nbinds int, nintbinds int) *Env {
	pool := &common.Pool // pool is an array, do NOT copy it!
	index := common.PoolSize - 1
	if index < 0 {
		return nil
	}
	common.PoolSize = index
	env := pool[index]
	pool[index] = nil
	if cap(env.Binds) < nbinds {
		env.Binds = make([]r.Value, nbinds)
	} else if nbinds != 0 {
		env.Binds = env.Binds[0:nbinds]
	}
	if cap(env.IntBinds) < nintbinds {
		env.IntBinds = make([]uint64, nintbinds)
	} else if nbinds != 0 {
		env.IntBinds = env.IntBinds[0:nintbinds]
	}
	env.Common = common
	return env
}

func (env *Env) MarkUsedByClosure() {
	for ; env != nil && !env.UsedByClosure; env = env.Outer {
		env.UsedByClosure = true
	}
}

// FreeEnv tells the interpreter that given Env is no longer needed.
func (env *Env) FreeEnv() {
	if env.UsedByClosure {
		// in use, cannot recycle
		return
	}
	common := env.Common
	n := common.PoolSize
	if n >= PoolCapacity {
		return
	}
	if env.AddressTaken {
		env.Binds = nil
		env.IntBinds = nil
		env.AddressTaken = false
	}
	env.Outer = nil
	env.Code = nil
	env.Common = nil
	common.Pool[n] = env // pool is an array, be careful NOT to copy it!
	common.PoolSize = n + 1
}

func (env *Env) Top() *Env {
	for ; env != nil; env = env.Outer {
		if env.Outer == nil {
			break
		}
	}
	return env
}

func (env *Env) File() *Env {
	for ; env != nil; env = env.Outer {
		outer := env.Outer
		if outer == nil || outer.Outer == nil {
			break
		}
	}
	return env
}

func (c *Comp) ParseAst(src string) Ast {
	nodes := c.ParseBytes([]byte(src))
	return AnyToAst(nodes, "ParseAst")
}

func (c *Comp) CompileAst(in Ast) func(*Env) (r.Value, []r.Value) {
	for {
		switch form := in.(type) {
		case nil:
			return nil
		case AstWithNode:
			return c.Compile(form.Node())
		case AstWithSlice:
			switch n := form.Size(); n {
			case 0:
				return nil
			case 1:
				in = form.Get(0)
				continue
			default:
				var list []func(*Env) (r.Value, []r.Value)
				for i := 0; i < n; i++ {
					fun := c.CompileAst(form.Get(i))
					if fun != nil {
						list = append(list, fun)
					}
				}
				return func(env *Env) (r.Value, []r.Value) {
					n_1 := len(list) - 1
					for i := 0; i < n_1; i++ {
						list[i](env)
					}
					return list[n_1](env)
				}
			}
		}
		c.Errorf("CompileAst: unsupported value, expecting <AstWithNode> or <AstWithSlice>, found %v <%v>", in, r.TypeOf(in))
		return nil
	}
}

func (c *Comp) Compile(in ast.Node) func(*Env) (r.Value, []r.Value) {
	c.Code.Clear()
	switch node := in.(type) {
	case nil:
		return nil
	case ast.Decl:
		c.Decl(node)
	case ast.Expr:
		return c.Expr(node).AsXV(c.CompileOptions)
	case *ast.ExprStmt:
		// special case of statement
		return c.Expr(node.X).AsXV(c.CompileOptions)
	case ast.Stmt:
		c.Stmt(node)
	case *ast.File:
		c.Errorf("Compile: unimplemented <*ast.File>, found %v <%v>", in, r.TypeOf(in))
		// TODO return c.File(node)
	default:
		c.Errorf("Compile: unsupported expression, expecting <ast.Decl>, <ast.Expr>, <ast.Stmt> or <*ast.File>, found %v <%v>", in, r.TypeOf(in))
		return nil
	}
	return c.Code.AsXV()
}
