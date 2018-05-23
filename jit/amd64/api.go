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
	"reflect"
	"syscall"
	"unsafe"
)

const (
	S        = 8 // sizeof(uint64)
	PAGESIZE = 4096
)

type Code []uint8

type Asm struct {
	Code Code
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

func (asm *Asm) Idx(bind *Var) *Asm {
	return asm.Uint32(uint32(bind.Idx) * S)
}

func (asm *Asm) Init() *Asm {
	asm.Code = asm.Code[:0]
	return asm.Bytes(0x48, 0x8b, 0x7c, 0x24, 0x08) // movq 0x8(%rsp), %rdi
}

func (asm *Asm) ret() *Asm {
	return asm.Bytes(0xc3) // ret
}

func (asm *Asm) Func() func(*uint64) {
	asm.ret()

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
