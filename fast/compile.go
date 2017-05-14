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

package fast

import (
	"go/ast"
	r "reflect"
	"strings"

	. "github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base"
)

func NewThreadGlobals() *ThreadGlobals {
	return &ThreadGlobals{
		Globals: base.NewGlobals(),
	}
}

func New() *CompEnv {
	top := NewCompEnvTop("builtin")
	top.env.UsedByClosure = true // do not free this *Env
	file := NewCompEnv(top, "main")
	file.env.UsedByClosure = true // do not free this *Env
	return file
}

func NewCompEnvTop(path string) *CompEnv {
	name := path[1+strings.LastIndexByte(path, '/'):]

	globals := base.NewGlobals()
	envGlobals := &ThreadGlobals{Globals: globals}
	compGlobals := &CompThreadGlobals{
		Globals: globals,
	}
	c := &CompEnv{
		Comp: &Comp{
			UpCost:            1,
			Depth:             0,
			Outer:             nil,
			Name:              name,
			Path:              path,
			CompThreadGlobals: compGlobals,
		},
		env: &Env{
			Outer:         nil,
			ThreadGlobals: envGlobals,
		},
	}
	envGlobals.TopEnv = c.env
	c.addBuiltins()
	return c
}

func NewCompEnv(outer *CompEnv, path string) *CompEnv {
	name := path[1+strings.LastIndexByte(path, '/'):]

	compGlobals := outer.Comp.CompThreadGlobals
	envGlobals := outer.env.ThreadGlobals
	c := &CompEnv{
		Comp: &Comp{
			UpCost:            1,
			Depth:             outer.Comp.Depth + 1,
			Outer:             outer.Comp,
			Name:              name,
			Path:              path,
			CompThreadGlobals: compGlobals,
		},
		env: &Env{
			Outer:         outer.env,
			ThreadGlobals: envGlobals,
		},
	}
	if outer.env.Outer == nil {
		envGlobals.FileEnv = c.env
	}
	return c
}

func NewComp(outer *Comp) *Comp {
	if outer == nil {
		return &Comp{UpCost: 1}
	}
	return &Comp{
		UpCost:            1,
		Depth:             outer.Depth + 1,
		Code:              outer.Code,
		Outer:             outer,
		CompileOptions:    outer.CompileOptions,
		CompThreadGlobals: outer.CompThreadGlobals,
	}
}

func (c *Comp) TopComp() *Comp {
	for ; c != nil; c = c.Outer {
		if c.Outer == nil {
			break
		}
	}
	return c
}

func (c *Comp) FileComp() *Comp {
	for ; c != nil; c = c.Outer {
		outer := c.Outer
		if outer == nil || outer.Outer == nil {
			break
		}
	}
	return c
}

// if a function Env only declares ignored binds, it gets this scratch buffers
var ignoredBinds = []r.Value{base.Nil}
var ignoredIntBinds = []uint64{0}

func NewEnv(outer *Env, nbinds int, nintbinds int) *Env {
	tg := outer.ThreadGlobals
	pool := &tg.Pool // pool is an array, do NOT copy it!
	index := tg.PoolSize - 1
	var env *Env
	if index >= 0 {
		tg.PoolSize = index
		env = pool[index]
		pool[index] = nil
	} else {
		env = &Env{}
	}
	if nbinds <= 1 {
		env.Binds = ignoredBinds
	} else if cap(env.Binds) < nbinds {
		env.Binds = make([]r.Value, nbinds)
	} else {
		env.Binds = env.Binds[0:nbinds]
	}
	if nintbinds <= 1 {
		env.IntBinds = ignoredIntBinds
	} else if cap(env.IntBinds) < nintbinds {
		env.IntBinds = make([]uint64, nintbinds)
	} else {
		env.IntBinds = env.IntBinds[0:nintbinds]
	}
	env.Outer = outer
	env.IP = outer.IP
	env.Code = outer.Code
	env.ThreadGlobals = tg
	return env
}

func NewEnv4Func(outer *Env, nbinds int, nintbinds int) *Env {
	tg := outer.ThreadGlobals
	pool := &tg.Pool // pool is an array, do NOT copy it!
	index := tg.PoolSize - 1
	var env *Env
	if index >= 0 {
		tg.PoolSize = index
		env = pool[index]
		pool[index] = nil
	} else {
		env = &Env{}
	}
	if nbinds <= 1 {
		env.Binds = ignoredBinds
	} else if cap(env.Binds) < nbinds {
		env.Binds = make([]r.Value, nbinds)
	} else {
		env.Binds = env.Binds[0:nbinds]
	}
	if nintbinds <= 1 {
		env.IntBinds = ignoredIntBinds
	} else if cap(env.IntBinds) < nintbinds {
		env.IntBinds = make([]uint64, nintbinds)
	} else {
		env.IntBinds = env.IntBinds[0:nintbinds]
	}
	env.Outer = outer
	env.ThreadGlobals = tg
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
	common := env.ThreadGlobals
	n := common.PoolSize
	if n >= PoolCapacity {
		return
	}
	if env.AddressTaken {
		env.IntBinds = nil
		env.AddressTaken = false
	}
	env.Outer = nil
	env.Code = nil
	env.ThreadGlobals = nil
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

func (c *Comp) Parse(src string) Ast {
	c.Line = 0
	nodes := c.ParseBytes([]byte(src))
	return AnyToAst(nodes, "ParseAst")
}

func (c *Comp) Compile(in Ast) func(*Env) (r.Value, []r.Value) {
	for {
		switch form := in.(type) {
		case nil:
			return nil
		case AstWithNode:
			return c.CompileNode(form.Node())
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
					fun := c.Compile(form.Get(i))
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
		c.Errorf("Compile: unsupported value, expecting <AstWithNode> or <AstWithSlice>, found %v <%v>", in, r.TypeOf(in))
		return nil
	}
}

func (c *Comp) CompileNode(node ast.Node) func(*Env) (r.Value, []r.Value) {
	c.Code.Clear()
	if node == nil {
		return nil
	}
	c.Pos = node.Pos()
	switch node := node.(type) {
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
		c.File(node)
	default:
		c.Errorf("Compile: unsupported expression, expecting <ast.Decl>, <ast.Expr>, <ast.Stmt> or <*ast.File>, found %v <%v>", node, r.TypeOf(node))
		return nil
	}
	return c.Code.AsXV()
}

func (c *Comp) File(node *ast.File) {
	c.Name = node.Name.Name
	for _, decl := range node.Decls {
		c.Decl(decl)
	}
}
