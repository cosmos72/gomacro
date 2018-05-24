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
 * asm.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	S        = 8 // sizeof(uint64)
	PAGESIZE = 4096
	VERBOSE  = false
)

func (s *Save) Init(start, end uint16) {
	s.start, s.idx, s.end = start, start, end
}

func (asm *Asm) Init() *Asm {
	return asm.Init2(0, 0)
}

func (asm *Asm) Init2(saveStart, saveEnd uint16) *Asm {
	asm.code = asm.code[:0]
	asm.swRegs = make(map[Reg]hwReg)
	asm.liveRegs.InitLive()
	asm.save.Init(saveStart, saveEnd)
	return asm.Bytes(0x48, 0x8b, 0x7c, 0x24, 0x08) // movq 0x8(%rsp), %rdi
}

func (asm *Asm) Bytes(bytes ...uint8) *Asm {
	asm.code = append(asm.code, bytes...)
	return asm
}

func (asm *Asm) Uint32(val uint32) *Asm {
	asm.code = append(asm.code, uint8(val), uint8(val>>8), uint8(val>>16), uint8(val>>24))
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
	return asm.Uint32(uint32(a.idx) * S)
}

func (asm *Asm) PushRegs(rs *hwRegs) *hwRegs {
	var ret hwRegs
	v := &Var{}
	for r := rLo; r <= rHi; r++ {
		if !rs.Contains(r) || !asm.liveRegs.Contains(r) {
			continue
		}
		if asm.save.idx >= asm.save.end {
			errorf("save area is full, cannot push registers")
		}
		v.idx = asm.save.idx
		asm.Store(v, r)
		asm.save.idx++
		ret.Set(r)
	}
	return &ret
}

func (asm *Asm) PopRegs(rs *hwRegs) {
	v := &Var{}
	for r := rHi; r >= rLo; r-- {
		if !rs.Contains(r) {
			continue
		}
		if asm.save.idx <= asm.save.start {
			errorf("save area is empty, cannot pop registers")
		}
		asm.save.idx--
		v.idx = asm.save.idx
		asm.Load(r, v)
	}
}

func (asm *Asm) hwAlloc(a Arg) (r hwReg, allocated bool) {
	r = a.Reg()
	if r != noReg {
		return r, false
	}
	r = asm.liveRegs.Alloc()
	asm.Load(r, a)
	return r, true
}

func (asm *Asm) hwAllocConst(val int64) hwReg {
	r := asm.liveRegs.Alloc()
	asm.LoadConst(r, val)
	return r
}

func (asm *Asm) hwAlloc3(a Arg, ret *hwReg, allocated *bool) *Asm {
	*ret, *allocated = asm.hwAlloc(a)
	return asm
}

func (asm *Asm) hwFree(r hwReg, allocated bool) *Asm {
	if r.Valid() && allocated {
		asm.liveRegs.Free(r)
	}
	return asm
}

func (asm *Asm) ret() *Asm {
	return asm.Bytes(0xc3) // ret
}

func (asm *Asm) Func() func(*uint64) {
	asm.ret()
	if VERBOSE {
		fmt.Printf("asm: %#v\n", asm.code)
	}
	mem, err := syscall.Mmap(-1, 0, (len(asm.code)+PAGESIZE-1)&^(PAGESIZE-1),
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANONYMOUS|syscall.MAP_PRIVATE)
	if err != nil {
		errorf("syscall.Mmap failed: %v", err)
	}
	copy(mem, asm.code)
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
