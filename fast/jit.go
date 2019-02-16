/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * jit.go
 *
 *  Created on Feb 16, 2019
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
	r "reflect"

	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/jit"
	"github.com/cosmos72/gomacro/jit/asm"
	xr "github.com/cosmos72/gomacro/xreflect"
)

const JIT_DEBUG = false

func jitLog(e *Expr) {
	if JIT_DEBUG {
		output.Debugf("jit successful: %+v => %v", e, e.Jit)
	}
}

// HACK!
var jitArch = asm.Archs[asm.ARCH_ID]

var jitConfig = jitGetConfig()

func jitGetConfig() *jit.RegIdCfg {
	if jitArch == nil || !asm.SUPPORTED {
		// host architecture not supported by jit/asm
		return nil
	}
	config := jitArch.RegIdCfg()
	return &config
}

// if supported, set e.Jit to jit constant == e.Lit.Value
// always returns e.
func jitConst(e *Expr) *Expr {
	if e.Jit == nil && e.Const() {
		switch e.Lit.Type.Kind() {
		case r.Bool, r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
			r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr,
			r.Float32, r.Float64: // r.Complex64, r.Complex128

			jc, err := jit.ConstInterface(e.Lit.Value, e.Lit.Type.ReflectType())
			if err == nil {
				e.Jit = jc
				jitLog(e)
			}
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute xe
// always returns e.
func jitIdentity(e *Expr, xe *Expr) *Expr {
	if jitConfig != nil && e.Jit == nil {
		jitConst(xe)
		if xe.Jit != nil {
			e.Jit = xe.Jit
			jitLog(e)
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute t(xe)
// always returns e.
func jitCast(e *Expr, t xr.Type, xe *Expr) *Expr {
	if jitConfig != nil && e.Jit == nil {
		jitConst(xe)
		if xe.Jit != nil {
			jop, err := jit.KindOp1(t.Kind())
			if err == nil {
				e.Jit = jit.NewExpr1(jop, xe.Jit)
				jitLog(e)
			}
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute op xe
// always returns e.
func jitUnaryExpr(e *Expr, op token.Token, xe *Expr) *Expr {
	if jitConfig != nil && e.Jit == nil {
		jitConst(xe)
		if xe.Jit != nil {
			jop, err := jit.TokenOp1(op)
			if err == nil {
				e.Jit = jit.NewExpr1(jop, xe.Jit)
				jitLog(e)
			}
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute xe op ye
// always returns e.
func jitBinaryExpr(e *Expr, op token.Token, xe *Expr, ye *Expr) *Expr {
	if jitConfig != nil && e.Jit == nil {
		jitConst(xe)
		jitConst(ye)
		if xe.Jit != nil && ye.Jit != nil {
			jop, err := jit.TokenOp2(op)
			if err == nil {
				e.Jit = jit.NewExpr2(jop, xe.Jit, ye.Jit)
				jitLog(e)
			}
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will access local variable
// always returns e.
// currently not supported, needs access to env.Vals[idx]
// which is a reflect.Value
func jitSymbol(e *Expr, idx int, upn int) *Expr {
	return e
}

// if supported, set e.Jit to jit expression that will access local variable
// always returns e.
func jitIntSymbol(e *Expr, idx int, upn int) *Expr {
	if jitConfig != nil && e.Jit == nil {
		jvar, err := jit.MakeVar(idx, upn, jit.Kind(e.Type.Kind()), *jitConfig)
		if err == nil {
			e.Jit = jvar
			jitLog(e)
		}
	}
	return e
}
