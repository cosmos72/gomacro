// +build !go1.12

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
 * range_map_go1_11.go
 *
 *  Created on Dec 09, 2019
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"
	"unsafe"

	"github.com/cosmos72/gomacro/base/reflect"
)

func (c *Comp) rangeMap(node *ast.RangeStmt, erange *Expr, jump *rangeJump) {
	t := erange.Type
	tkey, tval := t.Key(), t.Elem()
	tkeyslice := c.Universe.SliceOf(tkey)
	rtkeyslice := tkeyslice.ReflectType()

	// unnamed bind, contains map
	bindmap := c.DeclVar0("", nil, erange)
	idxmap := bindmap.Desc.Index()

	// unnamed bind, contains map keys
	bindkeys := c.NewBind("", VarBind, tkeyslice)
	idxkeys := bindkeys.Desc.Index()
	c.append(func(env *Env) (Stmt, *Env) {
		// convert []xr.Value slice into a []rtkey slice, to avoid reflect.Value.Interface() while iterating
		vkeys := env.Vals[idxmap].MapKeys()
		keys := xr.MakeSlice(rtkeyslice, len(vkeys), len(vkeys))
		for i, vkey := range vkeys {
			keys.Index(i).Set(vkey)
		}
		env.Vals[idxkeys] = keys
		env.IP++
		return env.Code[env.IP], env
	})

	// unnamed bind, contains iteration index
	bindnext := c.DeclVar0("", c.TypeOfInt(), nil)
	idxnext := bindnext.Desc.Index()

	placekey, placeval := c.rangeVars(node, tkey, tval)

	var bindkey *Bind
	var ekey *Expr
	if placekey != nil || placeval != nil {
		// unnamed bind, contains iteration map key
		bindkey = c.DeclVar0("", c.TypeOfInterface(), nil)
		ekey = unwrapBind(bindkey, tkey)
	}

	jump.Start = c.Code.Len()

	if bindkey == nil {
		// check iteration index against # of keys
		c.append(func(env *Env) (Stmt, *Env) {
			n := env.Vals[idxkeys].Len()
			i := *(*int)(unsafe.Pointer(&env.Ints[idxnext]))
			var ip int
			if i < n {
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
	} else {
		// check iteration index against # of keys,
		// and copy current map key into bindkey
		idxkey := bindkey.Desc.Index()
		c.append(func(env *Env) (Stmt, *Env) {
			vkeys := env.Vals[idxkeys]
			n := vkeys.Len()
			i := *(*int)(unsafe.Pointer(&env.Ints[idxnext]))
			var ip int
			if i < n {
				env.Vals[idxkey] = vkeys.Index(i)
				ip = env.IP + 1
			} else {
				ip = jump.Break
			}
			env.IP = ip
			return env.Code[ip], env
		})
	}

	if placekey != nil {
		// copy current map key into placekey
		c.SetPlace(placekey, token.ASSIGN, ekey)
	}

	if placeval == nil {
		// nothing to do
	} else if placeval.IsVar() && !reflect.IsOptimizedKind(placeval.Type.Kind()) {
		idxkey := bindkey.Desc.Index()
		idxval := placeval.Var.Desc.Index()
		upval := placeval.Var.Upn
		rtype := tval.ReflectType()
		zero := xr.Zero(rtype)
		c.append(func(env *Env) (Stmt, *Env) {
			vmap := env.Vals[idxmap]
			key := env.Vals[idxkey]
			o := env
			for j := 0; j < upval; j++ {
				o = o.Outer
			}
			val := vmap.MapIndex(key)
			if !val.IsValid() {
				val = zero
			} else if val.Type() != rtype {
				val = convert(val, rtype)
			}
			o.Vals[idxval].Set(val)
			env.IP++
			return env.Code[env.IP], env
		})
	} else {
		emap := c.Bind(bindmap)
		c.SetPlace(placeval, token.ASSIGN, c.mapIndex1(nil, emap, ekey))
	}

	// compile the body
	c.Block(node.Body)

	// "continue" is a jump to the last statement below
	jump.Continue = c.Code.Len()

	// increase iteration index and jump back to start
	c.append(func(env *Env) (Stmt, *Env) {
		(*(*int)(unsafe.Pointer(&env.Ints[idxnext])))++
		ip := jump.Start
		env.IP = ip
		return env.Code[ip], env
	})
}
