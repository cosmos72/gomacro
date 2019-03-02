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
	"errors"
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
}

func makeJitField(offset uintptr, kind jit.Kind) jitField {
	return jitField{
		index: jit.ConstUintptr(offset / uintptr(kind.Size())),
	}
}

var (
	jit_verbose int  = 0
	jit_enabled bool = false

	envInts jitField // description of Env.Ints struct field
	envIP   jitField // description of Env.IP   struct field
	envCode jitField // description of Env.Code struct field
	envOk   bool
)

func init() {
	if s := os.Getenv("GOMACRO_JIT_V"); s != "" {
		jit_verbose, _ = strconv.Atoi(s)
	}

	jitExtractEnvFields()

	jitCheckSupported()
}

func jitExtractEnvFields() {
	var sizeofUintptr = uintptr(jit.Uintptr.Size())

	tenv := r.TypeOf((*Env)(nil)).Elem()
	f, ok := tenv.FieldByName("Ints")
	if !ok || f.Offset%sizeofUintptr != 0 {
		return
	}
	envInts = makeJitField(f.Offset, jit.Uintptr)

	f, ok = tenv.FieldByName("IP")
	if !ok || f.Offset%f.Type.Size() != 0 {
		return
	}
	envIP = makeJitField(f.Offset, jit.Kind(f.Type.Kind()))

	f, ok = tenv.FieldByName("Code")
	if !ok || f.Offset%sizeofUintptr != 0 {
		return
	}
	envCode = makeJitField(f.Offset, jit.Uintptr)
	envOk = true
}

func jitCheckSupported() {
	if !envOk {
		if jit_verbose > 0 {
			output.Debugf("Jit: failed to extract *Env fields")
		}
		jit_enabled = false
		return
	}
	arch := jit.Archs[jit.ARCH_ID]
	if arch == nil || !jit.SUPPORTED {
		if jit_verbose > 0 {
			output.Debugf("Jit: unsupported architecture or operating system")
		}
		jit_enabled = false
		return
	}
	if !jit_enabled && os.Getenv("GOMACRO_JIT") == "" {
		if jit_verbose > 0 {
			output.Debugf("Jit: not enabled with environment variable GOMACRO_JIT")
		}
		return
	}
	jit_enabled = true
	// stmtNop = jitMakeInterpNop()
}

func NewJit() *Jit {
	if !jit_enabled {
		return nil
	}
	arch := jit.Archs[jit.ARCH_ID]
	var j Jit
	j.InitArch(arch)
	// tell jit compiler we need register RVAR
	j.Asm().RegIncUse(arch.RegIdConfig().RVAR)
	if jit_verbose > 0 {
		output.Debugf("Jit supported and enabled")
	}
	return &j
}

func (j *Jit) InitArch(arch jit.Arch) *Jit {
	j.c.InitArch(arch)
	return j
}

func (j *Jit) Comp() *jit.Comp {
	return &j.c
}

func (j *Jit) Asm() *jit.Asm {
	return j.c.Asm()
}

func (j *Jit) Code() jit.Code {
	return j.c.Code()
}

func (j *Jit) RegIdConfig() jit.RegIdConfig {
	return j.c.RegIdConfig
}

func (j *Jit) Log(e *Expr) {
	if jit_verbose > 2 {
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
		jvar, err := jit.MakeVar(idx, upn, jit.Kind(e.Type.Kind()), j.RegIdConfig())
		if err == nil {
			e.Jit = jvar
			j.Log(e)
		}
	}
	return e
}

// if supported, return a jit-compiled statement that will perform va OP= init
func (j *Jit) SetVar(va *Var, op token.Token, init *Expr) Stmt {
	if jit_verbose > 2 && j != nil {
		output.Debugf("jit to compile assignment: %v %v %v with e.Jit = %v", va, op, init, init.Jit)
	}
	if j == nil || init.Jit == nil {
		return nil
	}
	jvar, err := j.MakeVar(va)
	if err != nil {
		if jit_verbose > 0 {
			output.Debugf("jit failed: %v", err)
			return nil
		}
	}
	inst, err := jit.TokenInst2(op)
	if err != nil {
		if jit_verbose > 0 {
			output.Debugf("jit failed: %v", err)
			return nil
		}
	}
	return j.stmt0(jit.NewStmt2(inst, jvar, init.Jit))
}

var errMakeVarClass = errors.New("unimplemented: jit.MakeVar with class != IntBind")

func (j *Jit) MakeVar(va *Var) (jit.Mem, error) {
	if va.Desc.Class() != IntBind {
		return jit.Mem{}, errMakeVarClass
	}
	return jit.MakeVar(va.Desc.Index(), va.Upn, jit.Kind(va.Type.Kind()), j.RegIdConfig())
}

// if supported, return a jit-compiled Stmt that will evaluate Expr.
// return nil on failure
// TODO: optimize
func (j *Jit) AsStmt(e *Expr) Stmt {
	if j == nil || e.Jit == nil {
		return nil
	}
	var success bool

	j.preamble()
	defer j.cleanup(&success)

	// compile accumulated jit expression and discard the result.
	jc := j.Comp()
	jc.Stmt1(jit.NOP, e.Jit)

	stmt := j.makeStmt()
	success = true
	return stmt
}

// if supported, replace e.Fun with a jit-compiled equivalent function.
// always returns e.
func (j *Jit) Fun(e *Expr) *Expr {
	if jit_verbose > 2 && j != nil {
		output.Debugf("jit to compile expr: %v with e.Jit = %v", e, e.Jit)
	}
	if j == nil || e.Jit == nil {
		return e
	}
	kind := jit.Kind(e.Type.Kind())
	if kind.Size() == 0 {
		return e
	}
	fun := j.fun0(e, kind)
	if fun != nil {
		e.Fun = fun
		e.Jit = nil // in case we are invoked again on the same Expr
	}
	return e
}

// implementation of Jit.Fun
func (j *Jit) fun0(e *Expr, kind jit.Kind) I {
	var success bool

	j.preamble()
	defer j.cleanup(&success)

	// compile accumulated jit expression and copy result to stack.
	// on amd64 and arm64, in a func(env *Env) ...
	// the return value is on the stack at [RSP+16]
	jc := j.Comp()
	jc.Stmt2(jit.ASSIGN, jc.MakeParam(16, e.Jit.Kind()), e.Jit)
	if jit_verbose > 1 {
		output.Debugf("jit compiled:  %v", jc.Code())
		output.Debugf("jit assembled: %v", jc.Assemble())
	}
	fun := j.makeFun(kind)
	success = true
	return fun
}

// implementation of Jit.Stmt
func (j *Jit) stmt0(t jit.Stmt) Stmt {
	var success bool

	j.preamble()
	defer j.cleanup(&success)

	jc := j.Comp()
	jc.Stmt(t)

	stmt := j.makeStmt()
	success = true
	return stmt
}

func (j *Jit) preamble() {
	jc := j.Comp()
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
}

func (j *Jit) cleanup(success *bool) {
	jc := j.Comp()
	jc.ClearCode()
	jc.ClearRegs()
	if !*success {
		err := recover()
		if jit_verbose > 0 {
			output.Debugf("jit failed: %v", err)
		}
	}
}

func (j *Jit) makeFun(kind jit.Kind) I {
	jc := j.Comp()
	switch kind {
	case jit.Bool:
		var fun func(*Env) bool
		jc.Func(&fun)
		return fun
	case jit.Int:
		var fun func(*Env) int
		jc.Func(&fun)
		return fun
	case jit.Int8:
		var fun func(*Env) int8
		jc.Func(&fun)
		return fun
	case jit.Int16:
		var fun func(*Env) int16
		jc.Func(&fun)
		return fun
	case jit.Int32:
		var fun func(*Env) int32
		jc.Func(&fun)
		return fun
	case jit.Int64:
		var fun func(*Env) int64
		jc.Func(&fun)
		return fun
	case jit.Uint:
		var fun func(*Env) uint
		jc.Func(&fun)
		return fun
	case jit.Uint8:
		var fun func(*Env) uint8
		jc.Func(&fun)
		return fun
	case jit.Uint16:
		var fun func(*Env) uint16
		jc.Func(&fun)
		return fun
	case jit.Uint32:
		var fun func(*Env) uint32
		jc.Func(&fun)
		return fun
	case jit.Uint64:
		var fun func(*Env) uint64
		jc.Func(&fun)
		return fun
	case jit.Uintptr:
		var fun func(*Env) uintptr
		jc.Func(&fun)
		return fun
	case jit.Float32:
		var fun func(*Env) float32
		jc.Func(&fun)
		return fun
	case jit.Float64:
		var fun func(*Env) float64
		jc.Func(&fun)
		return fun
	/*
		case jit.Complex64:
		case jit.Complex128:
	*/
	default:
		return nil
	}
}

func (j *Jit) makeStmt() Stmt {
	// jit-generate the following
	/*
		func(env *Env) (Stmt, *Env) {
			fun(env)
			ip := env.IP + 1
			env.IP = ip
			return env.Code[ip], env
		}
	*/

	jc := j.Comp()
	renv := jc.NewSoftReg(jit.Uint64)
	s := jc.NewSoftReg(jit.Uint64)
	t := jc.NewSoftReg(jit.Uint64)
	// on amd64 and arm64, in a func(env *Env) ...
	// the parameter env is on the stack at [RSP+8]
	source := jit.Source{
		// renv = stack[env_param]
		jit.ASSIGN, renv, jc.MakeParam(8, jit.Uint64),
		// t = env.IP
		jit.ASSIGN, t, jit.NewExprIdx(renv, envIP.index, jit.Uint64),
		// t++
		jit.INC, t,
		// env.IP = t
		jit.IDX_ASSIGN, renv, envIP.index, t,
		// s = env.Code
		jit.ASSIGN, s, jit.NewExprIdx(renv, envCode.index, jit.Uint64),
		// s = s[t] i.e. s = env.Code[t] i.e. s = env.Code[env.IP+1]
		jit.ASSIGN, s, jit.NewExprIdx(s, t, jit.Uintptr),
		// stack[env_result] = renv
		jit.ASSIGN, jc.MakeParam(24, jit.Uint64), renv,
		// stack[stmt_result] = s, with s == env.Code[env.IP+1]
		jit.ASSIGN, jc.MakeParam(16, jit.Uint64), s,
		jit.FREE, renv,
		jit.FREE, s,
		jit.FREE, t,
	}
	jc.Compile(source)
	if jit_verbose > 1 {
		output.Debugf("jit compiled:  %v", jc.Code())
		output.Debugf("jit assembled: %v", jc.Assemble())
	}
	var f func(*Env) (Stmt, *Env)
	jc.Func(&f)
	return f
}
