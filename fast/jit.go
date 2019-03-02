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
	xr "github.com/cosmos72/gomacro/xreflect"
)

// jit.Comp wrapper
type Jit struct {
	c jit.Comp
}

type jitField struct {
	index jit.Const
	kind  jit.Kind
}

func makeJitField(offset uintptr, kind jit.Kind) jitField {
	return jitField{
		index: jit.ConstUintptr(offset / uintptr(kind.Size())),
		kind:  kind,
	}
}

var (
	JIT_VERBOSE int

	envInts jitField // description of Env.Ints struct field
	envIP   jitField // description of Env.IP   struct field
	envCode jitField // description of Env.Code struct field
	envOk   bool
)

func init() {
	JIT_VERBOSE, _ = strconv.Atoi(os.Getenv("GOMACRO_JIT_V"))
	// JIT_VERBOSE = 2

	var sizeofUintptr = uintptr(jit.Uintptr.Size())

	f, ok := r.TypeOf((*Env)(nil)).Elem().FieldByName("Ints")
	if !ok || f.Offset%sizeofUintptr != 0 {
		return
	}
	envInts = makeJitField(f.Offset, jit.Uintptr)

	f, ok = r.TypeOf((*Env)(nil)).Elem().FieldByName("IP")
	if !ok || f.Offset%f.Type.Size() != 0 {
		return
	}
	envIP = makeJitField(f.Offset, jit.Kind(f.Type.Kind()))

	f, ok = r.TypeOf((*Env)(nil)).Elem().FieldByName("Code")
	if !ok || f.Offset%sizeofUintptr != 0 {
		return
	}
	envCode = makeJitField(f.Offset, jit.Uintptr)
	envOk = true

	// stmtNop = jitMakeInterpNop()

}

func jitMakeInterpNop() Stmt {
	jc := jit.New()
	renv := jc.AllocSoftReg(jit.Uint64)
	s := jc.AllocSoftReg(jit.Uintptr)
	t := jc.AllocSoftReg(envIP.kind)
	// on amd64 and arm64, in a func(env *Env) ...
	// the parameter env is on the stack at [RSP+8]
	// renv = stack[env_param]
	jc.Stmt2(jit.ASSIGN, renv, jc.MakeParam(8, jit.Uint64))
	// t = env.IP
	jc.Stmt2(jit.ASSIGN, t, jit.NewExprIdx(renv, envIP.index, envIP.kind))
	// t++
	jc.Stmt1(jit.INC, t)
	// env.IP = t
	jc.Stmt3(jit.IDX_ASSIGN, renv, envIP.index, t)
	// s = env.Code
	jc.Stmt2(jit.ASSIGN, s, jit.NewExprIdx(renv, envCode.index, jit.Uintptr))
	// s = s[t] i.e. s = env.Code[t] i.e. s = env.Code[env.IP+1]
	jc.Stmt2(jit.ASSIGN, s, jit.NewExprIdx(s, t, jit.Uintptr))
	// stack[env_result] = renv
	jc.Stmt2(jit.ASSIGN, jc.MakeParam(24, jit.Uint64), renv)
	// stack[stmt_result] = s, with s == env.Code[env.IP+1]
	jc.Stmt2(jit.ASSIGN, jc.MakeParam(16, jit.Uint64), s)
	jc.FreeSoftReg(t)
	jc.FreeSoftReg(s)
	jc.FreeSoftReg(renv)
	var f func(*Env) (Stmt, *Env)
	jc.Func(&f)
	return f
}

func NewJit() *Jit {
	arch := jit.Archs[jit.ARCH_ID]
	if arch == nil || !jit.SUPPORTED {
		if JIT_VERBOSE > 0 {
			output.Debugf("Jit: unsupported architecture or operating system")
		}
		return nil
	}
	if !envOk {
		if JIT_VERBOSE > 0 {
			output.Debugf("Jit: failed to extract *Env fields")
		}
		return nil
	}
	if os.Getenv("GOMACRO_JIT") == "" {
		if JIT_VERBOSE > 0 {
			output.Debugf("Jit: not enabled with environment variable GOMACRO_JIT")
		}
		return nil
	}
	var j Jit
	j.c.InitArch(arch)
	// tell jit compiler we need register RVAR
	j.c.Asm().RegIncUse(arch.RegIdConfig().RVAR)
	if JIT_VERBOSE > 0 {
		output.Debugf("Jit supported and enabled")
	}
	return &j
}

func (j *Jit) Log(e *Expr) {
	if JIT_VERBOSE > 2 {
		output.Debugf("jit expr: %+v => %v", e, e.Jit)
	}
}

// if supported, set e.Jit to jit constant == e.Lit.Value
// always returns e.
func (j *Jit) Const(e *Expr) *Expr {
	if j != nil && e.Jit == nil && e.Const() {
		switch e.Lit.Type.Kind() {
		case r.Bool, r.Int, r.Int8, r.Int16, r.Int32, r.Int64,
			r.Uint, r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uintptr,
			r.Float32, r.Float64: // r.Complex64, r.Complex128

			c, err := jit.ConstInterface(e.Lit.Value, e.Lit.Type.ReflectType())
			if err == nil {
				e.Jit = c
				j.Log(e)
			}
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute xe
// always returns e.
func (j *Jit) Identity(e *Expr, xe *Expr) *Expr {
	if j != nil && e.Jit == nil {
		j.Const(xe)
		if xe.Jit != nil {
			e.Jit = xe.Jit
			j.Log(e)
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute t(xe)
// always returns e.
func (j *Jit) Cast(e *Expr, t xr.Type, xe *Expr) *Expr {
	if j != nil && e.Jit == nil {
		j.Const(xe)
		if xe.Jit != nil {
			jop, err := jit.KindOp1(t.Kind())
			if err == nil {
				e.Jit = jit.NewExpr1(jop, xe.Jit)
				j.Log(e)
			}
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute *xe
// always returns e.
// unimplemented.
func (j *Jit) Deref(e *Expr) *Expr {
	return e
}

// if supported, set e.Jit to jit expression that will compute op xe
// always returns e.
func (j *Jit) UnaryExpr(e *Expr, op token.Token, xe *Expr) *Expr {
	if j != nil && e.Jit == nil {
		j.Const(xe)
		if xe.Jit != nil {
			jop, err := jit.TokenOp1(op)
			if err == nil {
				e.Jit = jit.NewExpr1(jop, xe.Jit)
				j.Log(e)
			}
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will compute xe op ye
// always returns e.
func (j *Jit) BinaryExpr(e *Expr, op token.Token, xe *Expr, ye *Expr) *Expr {
	if j != nil && e.Jit == nil {
		j.Const(xe)
		j.Const(ye)
		if xe.Jit != nil && ye.Jit != nil {
			jop, err := jit.TokenOp2(op)
			if err == nil {
				e.Jit = jit.NewExpr2(jop, xe.Jit, ye.Jit)
				j.Log(e)
			}
		}
	}
	return e
}

// if supported, set e.Jit to jit expression that will access local variable
// always returns e.
// currently not supported, needs access to env.Vals[idx]
// which is a reflect.Value
func (j *Jit) Symbol(e *Expr, idx int, upn int) *Expr {
	return e
}

// if supported, set e.Jit to jit expression that will access local variable
// always returns e.
func (j *Jit) IntSymbol(e *Expr, idx int, upn int) *Expr {
	if j != nil && e.Jit == nil {
		jvar, err := jit.MakeVar(idx, upn, jit.Kind(e.Type.Kind()), j.c.RegIdConfig)
		if err == nil {
			e.Jit = jvar
			j.Log(e)
		}
	}
	return e
}

// if supported, return a jit-compiled statement that will perform va OP= init
// TODO: not yet implemented
// TODO: Comp.SetVar() should call this function
func (j *Jit) AssignStmt(va *Var, op token.Token, init *Expr) Stmt {
	return nil
}

// if supported, replace e.Fun with a jit-compiled equivalent function.
// always returns e.
func (j *Jit) Fun(e *Expr) *Expr {
	if JIT_VERBOSE > 2 && j != nil {
		output.Debugf("jit to compile: %v with e.Jit = %v", e, e.Jit)
	}
	if j != nil && e.Jit != nil {
		kind := jit.Kind(e.Type.Kind())
		if kind.Size() != 0 {
			fun := j.fun0(e, kind)
			if fun != nil {
				e.Fun = fun
				e.Jit = nil // in case we are invoked again on the same Expr
			}
		}
	}
	return e
}

// implementation of Jit.Fun
func (j *Jit) fun0(e *Expr, kind jit.Kind) I {
	jc := j.c
	jc.ClearCode()
	jc.ClearRegs()
	jc.Asm().RegIncUse(jc.RegIdConfig.RVAR)
	// on amd64 and arm64, in a func(env *Env) ...
	// the parameter env is on the stack at [RSP+8]
	rvar := jit.MakeReg(jc.RegIdConfig.RVAR, jit.Uint64)
	// env = stack[env_param]
	jc.Stmt2(jit.ASSIGN, rvar, jc.MakeParam(8, jit.Uint64))
	// rvar = env.Ints equivalent to rvar = &env.Ints[0]
	jc.Stmt2(jit.ASSIGN, rvar, jit.NewExprIdx(rvar, envInts.index, jit.Uint64))
	// compile accumulated jit expression and copy result to stack.
	// on amd64 and arm64, in a func(env *Env) ...
	// the return value is on the stack at [RSP+16]
	jc.Stmt2(jit.ASSIGN, jc.MakeParam(16, e.Jit.Kind()), e.Jit)
	var assembled bool
	defer func() {
		jc.ClearCode()
		jc.ClearRegs()
		if !assembled {
			err := recover()
			if JIT_VERBOSE > 0 {
				output.Debugf("jit failed: %v", err)
			}
		}
	}()
	machinecode := jc.Assemble()
	if JIT_VERBOSE > 1 {
		output.Debugf("jit compiled:  %v", jc.Code())
		output.Debugf("jit assembled: %v", machinecode)
	}
	fun := j.makeFun(kind)
	assembled = true
	return fun
}

func (j *Jit) makeFun(kind jit.Kind) I {
	asm := j.c.Asm()
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
