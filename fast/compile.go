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
 * compile.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"

	"github.com/cosmos72/gls"
	. "github.com/cosmos72/gomacro/ast2"
	. "github.com/cosmos72/gomacro/base"
)

func NewComp(outer *Comp, code *Code) *Comp {
	if outer == nil {
		return &Comp{UpCost: 1}
	}
	c := Comp{
		UpCost:      1,
		Depth:       outer.Depth + 1,
		Outer:       outer,
		CompGlobals: outer.CompGlobals,
	}
	// Debugf("NewComp(%p->%p) %s", outer, &c, debug.Stack())
	if code != nil {
		c.Code = *code
	}
	return &c
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

func NewIrGlobals() *IrGlobals {
	return &IrGlobals{
		Globals: *NewGlobals(),
		gls:     make(map[uintptr]*ThreadGlobals),
	}
}

func (g *IrGlobals) glsGet(goid uintptr) *ThreadGlobals {
	g.lock.Lock()
	ret := g.gls[goid]
	g.lock.Unlock()
	return ret
}

func (tg *ThreadGlobals) getThreadGlobals4Goid(goid uintptr) *ThreadGlobals {
	g := tg.IrGlobals
	ret := g.glsGet(goid)
	if ret == nil {
		ret = tg.new(goid)
		ret.glsStore()
	}
	return ret
}

func (tg *ThreadGlobals) glsStore() {
	g := tg.IrGlobals
	goid := tg.goid
	g.lock.Lock()
	g.gls[goid] = tg
	g.lock.Unlock()
}

func (tg *ThreadGlobals) glsDel() {
	g := tg.IrGlobals
	goid := tg.goid
	g.lock.Lock()
	delete(g.gls, goid)
	g.lock.Unlock()
}

func (tg *ThreadGlobals) new(goid uintptr) *ThreadGlobals {
	return &ThreadGlobals{
		IrGlobals: tg.IrGlobals,
		goid:      goid,
		FileEnv:   tg.FileEnv,
		TopEnv:    tg.TopEnv,
		// Interrupt, Signal, PoolSize and Pool are zero-initialized, fine with that
	}
}

// if a function Env only declares ignored binds, it gets this scratch buffers
var ignoredBinds = []r.Value{Nil}
var ignoredIntBinds = []uint64{0}

// common part between NewEnv(), newEnv4Func() and newEnv4Go()
func newEnv(tg *ThreadGlobals, outer *Env, nbind int, nintbind int) *Env {
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
	if nbind <= 1 {
		env.Vals = ignoredBinds
	} else if cap(env.Vals) < nbind {
		env.Vals = make([]r.Value, nbind)
	} else {
		env.Vals = env.Vals[0:nbind]
	}
	if nintbind <= 1 {
		env.Ints = ignoredIntBinds
	} else if cap(env.Ints) < nintbind {
		env.Ints = make([]uint64, nintbind)
	} else {
		env.Ints = env.Ints[0:nintbind]
	}
	env.Outer = outer
	env.ThreadGlobals = tg
	return env
}

// return a new, nested Env with given number of binds and intbinds
func NewEnv(outer *Env, nbind int, nintbind int) *Env {
	tg := outer.ThreadGlobals
	env := newEnv(tg, outer, nbind, nintbind)

	env.IP = outer.IP
	env.Code = outer.Code
	env.DebugPos = outer.DebugPos
	env.CallDepth = outer.CallDepth
	// this is a nested *Env, not a function body: to obtain the caller function,
	// follow env.Outer.Outer... chain until you find an *Env with non-nil Caller
	// env.Caller = nil
	// DebugCallStack Debugf("NewEnv(%p->%p) nbind=%d nintbind=%d calldepth: %d->%d", outer, env, nbind, nintbind, outer.CallDepth, env.CallDepth)
	tg.CurrEnv = env
	return env
}

func newEnv4Func(outer *Env, nbind int, nintbind int, debugComp *Comp) *Env {
	goid := gls.GoID()
	tg := outer.ThreadGlobals
	if tg.goid != goid {
		// no luck... get the correct ThreadGlobals for goid
		tg = tg.getThreadGlobals4Goid(goid)
	}
	env := newEnv(tg, outer, nbind, nintbind)

	env.DebugComp = debugComp
	caller := tg.CurrEnv
	env.Caller = caller
	if caller == nil {
		env.CallDepth = 1
	} else {
		env.CallDepth = caller.CallDepth + 1
	}
	// DebugCallStack Debugf("newEnv4Func(%p->%p) nbind=%d nintbind=%d calldepth: %d->%d", caller, env, nbind, nintbind, env.CallDepth-1, env.CallDepth)
	tg.CurrEnv = env
	return env
}

func (env *Env) MarkUsedByClosure() {
	for ; env != nil && !env.UsedByClosure; env = env.Outer {
		env.UsedByClosure = true
	}
}

// FreeEnv tells the interpreter that given nested *Env is no longer needed.
func (env *Env) FreeEnv() {
	g := env.ThreadGlobals
	g.CurrEnv = env.Outer
	env.freeEnv(g)
}

// freeEnv4Func tells the interpreter that given function body *Env is no longer needed.
func (env *Env) freeEnv4Func() {
	g := env.ThreadGlobals
	g.CurrEnv = env.Caller
	env.freeEnv(g)
}

func (env *Env) freeEnv(g *ThreadGlobals) {
	// DebugCallStack Debugf("FreeEnv(%p->%p), calldepth: %d->%d", env, caller, env.CallDepth, caller.CallDepth)
	if env.UsedByClosure {
		// in use, cannot recycle
		return
	}
	n := g.PoolSize
	if n >= PoolCapacity {
		return
	}
	if env.IntAddressTaken {
		env.Ints = nil
		env.IntAddressTaken = false
	}
	env.Outer = nil
	env.Code = nil
	env.DebugPos = nil
	env.DebugComp = nil
	env.Caller = nil
	env.ThreadGlobals = nil
	g.Pool[n] = env // pool is an array, be careful NOT to copy it!
	g.PoolSize = n + 1
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

// combined Parse + MacroExpandCodeWalk
func (c *Comp) Parse(src string) Ast {
	c.Line = 0
	nodes := c.ParseBytes([]byte(src))
	forms := anyToAst(nodes, "Parse")

	forms, _ = c.MacroExpandCodewalk(forms)
	if c.Options&OptShowMacroExpand != 0 {
		c.Debugf("after macroexpansion: %v", forms.Interface())
	}
	return forms
}

func (c *Comp) Compile(in Ast) *Expr {
	switch form := in.(type) {
	case nil:
		return nil
	case AstWithNode:
		return c.CompileNode(form.Node())
	case AstWithSlice:
		n := form.Size()
		var list []*Expr
		for i := 0; i < n; i++ {
			e := c.Compile(form.Get(i))
			if e != nil {
				list = append(list, e)
			}
		}
		return exprList(list, c.CompileOptions())
	}
	c.Errorf("unsupported Ast node, expecting <AstWithNode> or <AstWithSlice>, found %v <%v>", in, r.TypeOf(in))
	return nil
}

// compileExpr is a wrapper for Compile
// that guarantees Code does not get clobbered/cleared.
// Used by Comp.Quasiquote
func (c *Comp) compileExpr(in Ast) *Expr {
	cf := NewComp(c, nil)
	cf.UpCost = 0
	cf.Depth--
	return cf.Compile(in)
}

func (c *Comp) CompileNode(node ast.Node) *Expr {
	if n := c.Code.Len(); n != 0 {
		c.Warnf("Compile: discarding %d previously compiled statements from code buffer", n)
	}
	c.Code.Clear()
	if node == nil {
		return nil
	}
	c.Pos = node.Pos()
	switch node := node.(type) {
	case ast.Decl:
		c.Decl(node)
	case ast.Expr:
		return c.Expr(node, nil)
	case *ast.ExprStmt:
		// special case of statement
		return c.Expr(node.X, nil)
	case ast.Stmt:
		c.Stmt(node)
	case *ast.File:
		c.File(node)
	default:
		c.Errorf("unsupported node type, expecting <ast.Decl>, <ast.Expr>, <ast.Stmt> or <*ast.File>, found %v <%v>", node, r.TypeOf(node))
		return nil
	}
	return c.Code.AsExpr()
}

func (c *Comp) File(node *ast.File) {
	c.Name = node.Name.Name
	for _, decl := range node.Decls {
		c.Decl(decl)
	}
}

func (c *Comp) Append(stmt Stmt, pos token.Pos) {
	c.Code.Append(stmt, pos)
}

func (c *Comp) append(stmt Stmt) {
	c.Code.Append(stmt, c.Pos)
}
