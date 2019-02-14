// +build windows

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
 * sys_windows.go
 *
 *  Created on May 25, 2018
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

const MMAP_SUPPORTED = true

var PAGESIZE = windows.Getpagesize()

type memarea struct {
	addr, size uintptr
}

func (asm *Asm) mmap() memarea {
	asm.Epilogue()
	if VERBOSE {
		fmt.Printf("asm: %#v\n", asm.code)
	}
	size := uintptr((len(asm.code) + PAGESIZE - 1) &^ (PAGESIZE - 1))
	mem, err := windows.VirtualAlloc(0, size, windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if err != nil {
		errorf("sys/windows.VirtualAlloc failed: %v", err)
	}
	memcpy(mem, uintptr(unsafe.Pointer(&asm.code[0])), size)
	var old uint32
	err = windows.VirtualProtect(mem, size, windows.PAGE_EXECUTE_READ, &old)
	if err != nil {
		windows.VirtualFree(mem, 0, windows.MEM_RELEASE)
		errorf("sys/windows.VirtualProtect failed: %v", err)
	}
	return memarea{mem, size}
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

func munmap(mem memarea) {
	if mem.addr != 0 {
		windows.VirtualFree(mem.addr, 0, windows.MEM_RELEASE)
	}
}
