/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * api.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"
)

const (
	S        = 8 // sizeof(uint64)
	PAGESIZE = 4096
	VERBOSE  = false
)

type Code []uint8

type Save struct {
	Start, Idx, End uint16 // memory area where spill registers can be saved
}

type Asm struct {
	Code     Code
	LiveRegs Regs
	Save     Save
}

type Desc struct {
	Idx   uint16
	Upn   uint16
	Const bool
}

type Var struct {
	Val  int64
	Kind reflect.Kind
	Desc
}

func NewVar(idx uint16) *Var {
	return &Var{Desc: Desc{Idx: idx}}
}

func Int64(val int64) *Var {
	return &Var{Val: val, Desc: Desc{Const: true}}
}

func (s *Save) Init(start, end uint16) {
	s.Start, s.Idx, s.End = start, start, end
}

func (asm *Asm) Bytes(bytes ...uint8) *Asm {
	asm.Code = append(asm.Code, bytes...)
	return asm
}

func (asm *Asm) Uint32(val uint32) *Asm {
	asm.Code = append(asm.Code, uint8(val), uint8(val>>8), uint8(val>>16), uint8(val>>24))
	return asm
}

func (asm *Asm) Uint64(val uint64) *Asm {
	return asm.Uint32(uint32(val)).Uint32(uint32(val >> 32))
}

func (asm *Asm) Int32(val int32) *Asm {
	return asm.Uint32(uint32(val))
}

func (asm *Asm) Int64(val int64) *Asm {
	return asm.Uint64(uint64(val))
}

func (asm *Asm) Idx(a *Var) *Asm {
	return asm.Uint32(uint32(a.Idx) * S)
}

func (asm *Asm) PushRegs(rs Regs) Regs {
	var ret Regs
	v := &Var{}
	for r := FirstReg; r <= LastReg; r++ {
		mask := r.mask()
		if !rs.Contains(r) || !asm.LiveRegs.Contains(r) {
			continue
		}
		if asm.Save.Idx >= asm.Save.End {
			errorf("save area is full, cannot push registers")
		}
		v.Idx = asm.Save.Idx
		asm.Store(v, r)
		// asm.LiveRegs &^= mask
		asm.Save.Idx++

		ret |= mask
	}
	return ret
}

func (asm *Asm) PopRegs(rs Regs) {
	v := &Var{}
	for r := LastReg; r >= FirstReg; r-- {
		mask := r.mask()
		if !rs.Contains(r) {
			continue
		}
		if asm.Save.Idx <= asm.Save.Start {
			errorf("save area is empty, cannot pop registers")
		}
		asm.Save.Idx--
		v.Idx = asm.Save.Idx
		asm.Load(r, v)
		asm.LiveRegs |= mask
	}
}

func (asm *Asm) ToReg(a *Var) Reg {
	r := asm.LiveRegs.Alloc()
	asm.Load(r, a)
	return r
}

func (asm *Asm) ToReg2(a *Var, ret *Reg) *Asm {
	*ret = asm.ToReg(a)
	return asm
}

func (asm *Asm) FreeReg(r Reg) *Asm {
	if r.Valid() {
		asm.LiveRegs.Free(r)
	}
	return asm
}

func (asm *Asm) Init() *Asm {
	return asm.Init2(0, 0)
}

func (asm *Asm) Init2(saveStart, saveEnd uint16) *Asm {
	asm.Code = asm.Code[:0]
	asm.LiveRegs.Init()
	asm.Save.Init(saveStart, saveEnd)
	return asm.Bytes(0x48, 0x8b, 0x7c, 0x24, 0x08) // movq 0x8(%rsp), %rdi
}

func (asm *Asm) ret() *Asm {
	return asm.Bytes(0xc3) // ret
}

func (asm *Asm) Func() func(*uint64) {
	asm.ret()
	if VERBOSE {
		fmt.Printf("asm: %#v\n", asm.Code)
	}
	mem, err := syscall.Mmap(-1, 0, (len(asm.Code)+PAGESIZE-1)&^(PAGESIZE-1),
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANONYMOUS|syscall.MAP_PRIVATE)
	if err != nil {
		errorf("syscall.Mmap failed: %v", err)
	}
	copy(mem, asm.Code)
	err = syscall.Mprotect(mem, syscall.PROT_EXEC|syscall.PROT_READ)
	if err != nil {
		syscall.Munmap(mem)
		errorf("syscall.Mprotect failed: %v", err)
	}
	ptr := new(func(*uint64))
	*(**[]uint8)(unsafe.Pointer(ptr)) = &mem
	// runtime.SetFinalizer(ptr, munmap)
	return *ptr
}

func munmap(obj interface{}) {
	ptr, ok := obj.(*func(*uint64))
	if ok && ptr != nil && *ptr != nil {
		mem := **(**[]uint8)(unsafe.Pointer(ptr))
		syscall.Munmap(mem)
	}
}
