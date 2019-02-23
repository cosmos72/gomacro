// +build darwin dragonfly freebsd linux netbsd openbsd

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
 * mmap_unix.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"golang.org/x/sys/unix"
)

const (
	MMAP_SUPPORTED = true
)

// allocate memory in 64k chunks
var minAllocSize = unix.Getpagesize() * 16

type MemPool struct {
	bytes  []byte
	offset int
}

type MemArea []byte

func NewMemPool(size int) *MemPool {
	bytes, err := unix.Mmap(-1, 0, (size+minAllocSize-1)&^(minAllocSize-1),
		unix.PROT_READ|unix.PROT_EXEC,
		unix.MAP_ANON|unix.MAP_PRIVATE)
	if err != nil {
		errorf("sys/unix.Mmap failed: %v", err)
	}
	return &MemPool{bytes, 0}
}

func (mem *MemPool) Size() int {
	if mem == nil {
		return 0
	}
	return len(mem.bytes) - mem.offset
}

func (mem *MemPool) protect(prot int) {
	err := unix.Mprotect(mem.bytes, prot)
	if err != nil {
		errorf("sys/unix.Mprotect failed: %v", err)
	}
}

func (mem *MemPool) SetReadonly() {
	mem.protect(unix.PROT_READ | unix.PROT_EXEC)
}

func (mem *MemPool) SetReadWrite() {
	mem.protect(unix.PROT_READ | unix.PROT_WRITE | unix.PROT_EXEC)
}

func (mem *MemPool) Copy(code MachineCode) MemArea {
	size := len(code.Bytes)
	avail := mem.Size()
	if size > avail {
		errorf("MachineCode is %d bytes, cannot copy to %d bytes MemPool", size, avail)
	}
	if MMAP_VERBOSE {
		debugf("copying %d bytes MachineCode to MemPool{addr:%p, size:%d, offset:%d}",
			size, &mem.bytes[0], len(mem.bytes), mem.offset)
	}
	mem.SetReadWrite()
	copy(mem.bytes[mem.offset:], code.Bytes)
	mem.SetReadonly()
	used := (size + 15) &^ 15
	if used >= avail {
		used = avail
	}
	ret := mem.bytes[mem.offset : mem.offset+size]
	mem.offset += used
	return ret
}
