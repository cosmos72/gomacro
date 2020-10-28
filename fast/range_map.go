/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * range_map.go
 *
 *  Created on Dec 11, 2019
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"

	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) rangeMap(node *ast.RangeStmt, erange *Expr, jump *rangeJump) {
	t := erange.Type
	tkey, tval := t.Key(), t.Elem()

	// no need to save in a bind the map to iterate on:
	// we just call once the function to compute it
	mapfun := erange.AsX1()

	// unnamed bind, contains map iterator
	binditer := c.NewBind("", VarBind, c.TypeOfInterface())
	idxiter := binditer.Desc.Index()
	c.append(func(env *Env) (Stmt, *Env) {
		iter := mapfun(env).ReflectValue().MapRange()
		env.Vals[idxiter] = xr.ValueOf(iter)
		env.IP++
		return env.Code[env.IP], env
	})

	placekey, placeval := c.rangeVars(node, tkey, tval)

	var bindkey, bindval *Bind
	var idxkey, idxval int = NoIndex, NoIndex
	if placekey != nil {
		// unnamed bind, contains current key
		bindkey = c.NewBind("", VarBind, c.TypeOfInterface())
		idxkey = bindkey.Desc.Index()
	}
	if placeval != nil {
		// unnamed bind, contains current value
		bindval = c.NewBind("", VarBind, c.TypeOfInterface())
		idxval = bindval.Desc.Index()
	}

	jump.Start = c.Code.Len()

	// "continue" is a jump to the statement below
	jump.Continue = c.Code.Len()

	// advance iterator, check end-of-iteration,
	// and copy current key/val into bindkey/bindval
	c.append(func(env *Env) (Stmt, *Env) {
		var ip int
		iter := env.Vals[idxiter].Interface().(*r.MapIter)
		if iter.Next() {
			if idxkey != NoIndex {
				env.Vals[idxkey] = xr.MakeValue(iter.Key())
			}
			if idxval != NoIndex {
				env.Vals[idxval] = xr.MakeValue(iter.Value())
			}
			ip = env.IP + 1
		} else {
			ip = jump.Break
		}
		env.IP = ip
		return env.Code[ip], env
	})

	if placekey != nil {
		// copy current key into placekey
		c.SetPlace(placekey, token.ASSIGN, unwrapBind(bindkey, tkey))
	}
	if placeval != nil {
		// copy current val into placeval
		c.SetPlace(placeval, token.ASSIGN, unwrapBind(bindval, tval))
	}

	// compile the body
	c.Block(node.Body)

	// jump back to start
	c.append(func(env *Env) (Stmt, *Env) {
		ip := jump.Start
		env.IP = ip
		return env.Code[ip], env
	})
}
