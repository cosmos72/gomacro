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
	"hash/crc32"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	MMAP_SUPPORTED = true
)

// we must call win32 VirtualAlloc() once for each function. REASON:
// cannot call win32 FlushInstructionCache(),
// it is not available in Go
var pagesize = uintptr(windows.Getpagesize())

type MemPool struct {
	// no need to use *uint8: this C-allocated memory,
	// not Go-allocated memory, so it's outside the garbage collector reach
	addr         uintptr
	size, offset uintptr
}

type MemArea struct {
	// use *uint8 instead of uintptr to avoid garbage collector
	// freeing a MemArea created from Go-allocated MachineCode
	ptr  *uint8
	size uintptr
}

func NewMemPool(size int) *MemPool {
	poolsize := (uintptr(size) + pagesize - 1) &^ (pagesize - 1)
	addr, err := windows.VirtualAlloc(0, poolsize,
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_READWRITE)
	if err != nil {
		errorf("sys/windows.VirtualAlloc failed: %v", err)
	}
	return &MemPool{addr, poolsize, 0}
}

func (pool *MemPool) Size() int {
	if pool == nil {
		return 0
	}
	return int(pool.size - pool.offset)
}

func (pool *MemPool) protect(prot uint32) {
	var old uint32
	err := windows.VirtualProtect(pool.addr, pool.size, prot, &old)
	if err != nil {
		errorf("sys/windows.VirtualProtect failed: %v", err)
	}
}

func (pool *MemPool) SetReadonly() {
	pool.protect(windows.PAGE_EXECUTE_READ)
}

func (pool *MemPool) SetReadWrite() {
	// pool.protect(windows.PAGE_EXECUTE_READWRITE)
}

func (pool *MemPool) Copy(area MemArea) MemArea {
	size := area.size
	avail := uintptr(pool.Size())
	if size > avail {
		errorf("MemArea is %d bytes, cannot copy to %d bytes MemPool", size, avail)
	}
	if MMAP_VERBOSE {
		debugf("copying %d bytes MemArea to MemPool{addr:%#x, size:%d, offset:%d}",
			size, pool.addr, pool.size, pool.offset)
	}
	pool.SetReadWrite()
	memcpy(pool.addr, area.addr(), size)
	pool.SetReadonly()
	used := (size + pagesize - 1) &^ (pagesize - 1)
	if used >= avail {
		used = avail
	}
	ret := MemArea{(*uint8)(unsafe.Pointer(pool.addr + pool.offset)), size}
	// use all memory, because any function stored
	// in the remaining fragment will raise exception
	// mem.offset += used
	pool.offset = pool.size
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

// memory comparison. a bit slow, but avoids depending on CGO
func memcmp(lhs uintptr, rhs uintptr, size uintptr) int {
	if lhs == rhs || size == 0 {
		return 0
	}
	var i uintptr
	for ; i+8 <= size; i += 8 {
		l := *(*uint64)(unsafe.Pointer(lhs + i))
		r := *(*uint64)(unsafe.Pointer(rhs + i))
		if l < r {
			return -1
		} else if l > r {
			return 1
		}
	}
	for ; i < size; i++ {
		l := *(*uint8)(unsafe.Pointer(lhs + i))
		r := *(*uint8)(unsafe.Pointer(rhs + i))
		if l < r {
			return -1
		} else if l > r {
			return 1
		}
	}
	return 0
}

// convert MachineCode to MemArea
func (code MachineCode) MemArea() MemArea {
	size := uintptr(len(code.Bytes))
	var area MemArea
	if size != 0 {
		area.ptr = &code.Bytes[0]
		area.size = size
	}
	return area
}

func (area MemArea) Size() int {
	return int(area.size)
}

func (area MemArea) addr() uintptr {
	return uintptr(unsafe.Pointer(area.ptr))
}

func (area MemArea) Equal(other MemArea) bool {
	size := area.size
	if size != other.size {
		return false
	}
	if size == 0 {
		return true
	}
	return memcmp(area.addr(), other.addr(), size) == 0
}

var crcTable = crc32.MakeTable(crc32.Castagnoli)

func (area MemArea) Checksum() uint32 {
	// cannot use crc32.Checksum(): we do not have a []uint8 slice
	crc := ^uint32(0)
	addr := area.addr()
	for i := uintptr(0); i < area.size; i++ {
		index := uint8(crc) ^ *(*uint8)(unsafe.Pointer(addr + i))
		crc = crcTable[index] ^ crc>>8
	}
	return ^crc
}
