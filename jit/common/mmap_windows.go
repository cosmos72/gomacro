// +build windows

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * mmap_windows.go
 *
 *  Created on May 25, 2018
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	MMAP_SUPPORTED = false // crashes executing the second function call (!)
)

// allocate memory in 64k chunks
var minAllocSize = windows.Getpagesize() * 16

type MemPool struct {
	addr, size, offset uintptr
}

type MemArea struct {
	addr, size uintptr
}

func NewMemPool(size int) *MemPool {
	msize := uintptr((size + minAllocSize - 1) &^ (minAllocSize - 1))
	addr, err := windows.VirtualAlloc(0, msize,
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_EXECUTE_READ)
	if err != nil {
		errorf("sys/windows.VirtualAlloc failed: %v", err)
	}
	return &MemPool{addr, msize, 0}
}

func (mem *MemPool) Size() int {
	if mem == nil {
		return 0
	}
	return int(mem.size - mem.offset)
}

func (mem *MemPool) protect(prot uint32) {
	var old uint32
	err := windows.VirtualProtect(mem.addr, mem.size, prot, &old)
	if err != nil {
		errorf("sys/windows.VirtualProtect failed: %v", err)
	}
}

func (mem *MemPool) SetReadonly() {
	mem.protect(windows.PAGE_EXECUTE_READ)
}

func (mem *MemPool) SetReadWrite() {
	mem.protect(windows.PAGE_EXECUTE_READWRITE)
}

func (mem *MemPool) Copy(code MachineCode) MemArea {
	size := uintptr(len(code.Bytes))
	avail := uintptr(mem.Size())
	if size > avail {
		errorf("MachineCode is %d bytes, cannot copy to %d bytes MemPool", size, avail)
	}
	if MMAP_VERBOSE {
		debugf("copying %d bytes MachineCode to MemPool{addr:%#x, size:%d, offset:%d}",
			size, mem.addr, mem.size, mem.offset)
	}
	mem.SetReadWrite()
	memcpy(mem.addr, uintptr(unsafe.Pointer(&code.Bytes[0])), size)
	mem.SetReadonly()
	used := (size + 15) &^ 15
	if used >= avail {
		used = avail
	}
	ret := MemArea{mem.addr + mem.offset, size}
	mem.offset += used
	return ret
}

// memory copy. a bit slow, but avoids depending on CGO
func memcpy(dst uintptr, src uintptr, size uintptr) {
	var i uintptr
	for ; i+32 <= size; i += 32 {
		*(*uint64)(unsafe.Pointer(dst + i + 0)) = *(*uint64)(unsafe.Pointer(src + i + 0))
		*(*uint64)(unsafe.Pointer(dst + i + 8)) = *(*uint64)(unsafe.Pointer(src + i + 8))
		*(*uint64)(unsafe.Pointer(dst + i + 16)) = *(*uint64)(unsafe.Pointer(src + i + 16))
		*(*uint64)(unsafe.Pointer(dst + i + 24)) = *(*uint64)(unsafe.Pointer(src + i + 24))
	}
	for ; i+8 <= size; i += 8 {
		*(*uint64)(unsafe.Pointer(dst + i)) = *(*uint64)(unsafe.Pointer(src + i))
	}
	for ; i < size; i++ {
		*(*uint8)(unsafe.Pointer(dst + i)) = *(*uint8)(unsafe.Pointer(src + i))
	}
}
