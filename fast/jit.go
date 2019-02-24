/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
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
	"os"
	r "reflect"
	"strconv"

	"github.com/cosmos72/gomacro/base/output"
	"github.com/cosmos72/gomacro/jit"
	"github.com/cosmos72/gomacro/jit/asm"
	xr "github.com/cosmos72/gomacro/xreflect"
)

var JIT_VERBOSE = jitVerbose()

func jitVerbose() int {
	i, _ := strconv.Atoi(os.Getenv("GOMACRO_JIT_V"))
	return i
}

// offset of Env struct fields
var (
	envInts int32 = -1 // offset of Env.Ints
	envIP   int32 = -1 // offset of Env.IP
	envCode int32 = -1 // offset of Env.Code
	envOk   bool
)

func init() {
	f, ok := r.TypeOf((*Env)(nil)).Elem().FieldByName("Ints")
	if !ok || f.Offset != uintptr(int32(f.Offset)) {
		return
	}
	envInts = int32(f.Offset)
	f, ok = r.TypeOf((*Env)(nil)).Elem().FieldByName("IP")
	if !ok || f.Offset != uintptr(int32(f.Offset)) {
		return
	}
	envIP = int32(f.Offset)
	f, ok = r.TypeOf((*Env)(nil)).Elem().FieldByName("Code")
	if !ok || f.Offset != uintptr(int32(f.Offset)) {
		return
	}
	envCode = int32(f.Offset)
	envOk = true

	// stmtNop = jitMakeInterpNop()

}

func jitMakeInterpNop() Stmt {
	jc := jit.New()
	r := jc.Asm().RegAlloc(jit.Uint64)
	s := jc.Asm().RegAlloc(jit.Uint64)
	renv := jc.Asm().RegAlloc(jit.Uint64)
	renvid := renv.RegId()
	// on amd64 and arm64, in a func(env *Env) ...
	// the parameter env is on the stack at [RSP+8]
	env := jit.MakeParam(8, jit.Uint64, jc.RegIdConfig)
	// copy env from stack to renv register
	jc.Stmt2(jit.ASSIGN, renv, env)
	// copy env.IP to r
	mip := jit.MakeMem(envIP, renvid, jit.Uint64)
	jc.Stmt2(jit.ASSIGN, r, mip)
	// increment r
	jc.Stmt1(jit.INC, r)
	// copy r to env.IP
	jc.Stmt2(jit.ASSIGN, mip, r)
	// multiply r by sizeof(Stmt)
	jc.Stmt2(jit.MUL_ASSIGN, r, jit.MakeConst(8, jit.Uint64))
	// copy &env.Code[0] to s
	jc.Stmt2(jit.ASSIGN, s, jit.MakeMem(envCode, renvid, jit.Uint64))
	// add r (== env.IP) to s (== &env.Code[0]) to get &env.Code[env.IP]
	jc.Stmt2(jit.ADD_ASSIGN, s, r)
	// dereference s to get env.Code[env.IP]
	jc.Stmt2(jit.ASSIGN, s, jit.MakeMem(0, s.RegId(), jit.Uint64))
	// copy env from renv register to stack
	jc.Stmt2(jit.ASSIGN, jit.MakeParam(24, jit.Uint64, jc.RegIdConfig), renv)
	// copy env.Code[env.IP] from s register to stack
	jc.Stmt2(jit.ASSIGN, jit.MakeParam(16, jit.Uint64, jc.RegIdConfig), s)
	var f func(*Env) (Stmt, *Env)
	jc.Func(&f)
	return f
}

func jitNew() *jit.Comp {
	arch := asm.Archs[asm.ARCH_ID]
	if arch == nil || !asm.SUPPORTED || !envOk || os.Getenv("GOMACRO_JIT") == "" {
		// unsupported architecture or operating system,
		// or not enabled from environment variable
		return nil
	}
	jc := jit.NewArch(arch)
	// tell jit compiler we need register RVAR
	jc.Asm().RegIncUse(arch.RegIdConfig().RVAR)
	return jc
}

func jitLog(e *Expr) {
	if JIT_VERBOSE > 1 {
		output.Debugf("jit expr: %+v => %v", e, e.Jit)
	}
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
func (g *CompGlobals) jitIdentity(e *Expr, xe *Expr) *Expr {
	if g.JitComp != nil && e.Jit == nil {
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
func (g *CompGlobals) jitCast(e *Expr, t xr.Type, xe *Expr) *Expr {
	if g.JitComp != nil && e.Jit == nil {
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

// if supported, set e.Jit to jit expression that will compute *xe
// always returns e.
// unimplemented.
func (g *CompGlobals) jitDeref(e *Expr) *Expr {
	return e
}

// if supported, set e.Jit to jit expression that will compute op xe
// always returns e.
func (g *CompGlobals) jitUnaryExpr(e *Expr, op token.Token, xe *Expr) *Expr {
	if g.JitComp != nil && e.Jit == nil {
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
func (g *CompGlobals) jitBinaryExpr(e *Expr, op token.Token, xe *Expr, ye *Expr) *Expr {
	if g.JitComp != nil && e.Jit == nil {
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
func (g *CompGlobals) jitSymbol(e *Expr, idx int, upn int) *Expr {
	return e
}

// if supported, set e.Jit to jit expression that will access local variable
// always returns e.
func (g *CompGlobals) jitIntSymbol(e *Expr, idx int, upn int) *Expr {
	if g.JitComp != nil && e.Jit == nil {
		jvar, err := jit.MakeVar(idx, upn, jit.Kind(e.Type.Kind()), g.JitComp.RegIdConfig)
		if err == nil {
			e.Jit = jvar
			jitLog(e)
		}
	}
	return e
}

// if supported, return a jit-compiled statement that will perform va OP= init
// TODO: Comp.SetVar() should call this function
func (g *CompGlobals) jitAssignStmt(va *Var, op token.Token, init *Expr) Stmt {
	return nil
}

// if supported, replace e.Fun with a jit-compiled equivalent function.
// always returns e.
func (g *CompGlobals) jitFun(e *Expr) *Expr {
	if JIT_VERBOSE > 0 && g.JitComp != nil {
		output.Debugf("jit to compile: %v with e.Jit = %v", e, e.Jit)
	}
	if g.JitComp != nil && e.Jit != nil {
		kind := jit.Kind(e.Type.Kind())
		if kind.Size() != 0 {
			fun := g.jitFun0(e, kind)
			if fun != nil {
				e.Fun = fun
				e.Jit = nil // in case we are invoked again on the same Expr
			}
		}
	}
	return e
}

// implementation of jitFun
func (g *CompGlobals) jitFun0(e *Expr, kind jit.Kind) I {
	jc := g.JitComp
	asm := jc.Asm()
	jc.ClearCode()
	// on amd64 and arm64, in a func(env *Env) ...
	// the parameter env is on the stack at [RSP+8]
	env := jit.MakeParam(8, jit.Uint64, jc.RegIdConfig)
	rvarid := jc.RegIdConfig.RVAR
	rvar := jit.MakeReg(rvarid, jit.Uint64)
	// copy env from stack to RVAR
	jc.Stmt2(jit.ASSIGN, rvar, env)
	// copy &env.Ints[0] to RVAR
	jc.Stmt2(jit.ASSIGN, rvar, jit.MakeMem(envInts, rvarid, jit.Uint64))
	// compile accumulated jit expression
	val, softval := jc.Expr(e.Jit)
	// copy result to stack.
	// on amd64 and arm64, in a func(env *Env) ...
	// the return value is on the stack at [RSP+16]
	ret := jit.MakeParam(16, val.Kind(), jc.RegIdConfig)
	jc.Stmt2(jit.ASSIGN, ret, val)
	jc.FreeSoftReg(softval)
	if JIT_VERBOSE > 0 {
		output.Debugf("jit compiled: %v", jc.Code())
	}
	var assembled bool
	defer func() {
		jc.ClearCode()
		if !assembled {
			err := recover()
			if JIT_VERBOSE > 0 {
				output.Debugf("%v", err)
			}
		}
	}()
	machinecode := jc.Assemble()
	if JIT_VERBOSE > 0 {
		output.Debugf("jit compiled: %v", jc.Code())
		output.Debugf("jit assembled: %v", machinecode)
	}
	return jitMakeFun(asm, kind)
}

func jitMakeFun(asm *jit.Asm, kind jit.Kind) I {
	switch kind {
	case jit.Bool:
		var fun func(*Env) bool
		asm.Func(&fun)
		return fun
	case jit.Int:
		var fun func(*Env) int
		asm.Func(&fun)
		return fun
	case jit.Int8:
		var fun func(*Env) int8
		asm.Func(&fun)
		return fun
	case jit.Int16:
		var fun func(*Env) int16
		asm.Func(&fun)
		return fun
	case jit.Int32:
		var fun func(*Env) int32
		asm.Func(&fun)
		return fun
	case jit.Int64:
		var fun func(*Env) int64
		asm.Func(&fun)
		return fun
	case jit.Uint:
		var fun func(*Env) uint
		asm.Func(&fun)
		return fun
	case jit.Uint8:
		var fun func(*Env) uint8
		asm.Func(&fun)
		return fun
	case jit.Uint16:
		var fun func(*Env) uint16
		asm.Func(&fun)
		return fun
	case jit.Uint32:
		var fun func(*Env) uint32
		asm.Func(&fun)
		return fun
	case jit.Uint64:
		var fun func(*Env) uint64
		asm.Func(&fun)
		return fun
	case jit.Uintptr:
		var fun func(*Env) uintptr
		asm.Func(&fun)
		return fun
	case jit.Float32:
		var fun func(*Env) float32
		asm.Func(&fun)
		return fun
	case jit.Float64:
		var fun func(*Env) float64
		asm.Func(&fun)
		return fun
	/*
		case jit.Complex64:
		case jit.Complex128:
	*/
	default:
		return nil
	}
}
