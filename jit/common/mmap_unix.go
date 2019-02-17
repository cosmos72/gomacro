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
 * os_unix.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"fmt"

	"golang.org/x/sys/unix"
)

const MMAP_SUPPORTED = true

var PAGESIZE = unix.Getpagesize()

type memarea []byte

func (asm *Asm) mmap() memarea {
	asm.Epilogue()
	if VERBOSE {
		fmt.Printf("asm: %#v\n", asm.code)
	}
	mem, err := unix.Mmap(-1, 0, (len(asm.code)+PAGESIZE-1)&^(PAGESIZE-1),
		unix.PROT_READ|unix.PROT_WRITE, unix.MAP_ANON|unix.MAP_PRIVATE)
	if err != nil {
		errorf("sys/unix.Mmap failed: %v", err)
	}
	copy(mem, asm.code)
	err = unix.Mprotect(mem, unix.PROT_EXEC|unix.PROT_READ)
	if err != nil {
		unix.Munmap(mem)
		errorf("sys/unix.Mprotect failed: %v", err)
	}
	return mem
}

func munmap(mem memarea) {
	if len(mem) != 0 {
		unix.Munmap(mem)
	}
}
