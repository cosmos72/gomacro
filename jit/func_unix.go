// +build darwin dragonfly freebsd linux netbsd openbsd

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
 * func_unix.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"fmt"
	"syscall"
	"unsafe"
)

var PAGESIZE = syscall.Getpagesize()

func nop(*uint64) {
}

func (asm *Asm) Func() func(*uint64) {
	if len(asm.code) == 0 {
		return nop
	}
	asm.epilogue()
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
