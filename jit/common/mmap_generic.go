// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!windows

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
 * mmap_generic.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"runtime"
)

const MMAP_SUPPORTED = false

type MemPool struct {
}

func (mem *MemPool) Size() int {
	return 0
}

func NewMemPool(size int) *MemPool {
	errorf("MemPool: unsupported operating system %v, cannot create executable memory", runtime.GOOS)
	return nil
}

func (mem *MemPool) SetReadonly() {
}

func (mem *MemPool) SetReadWrite() {
}

func (mem *MemPool) Copy(code MachineCode) {
	errorf("MemPool: unsupported operating system %v, cannot copy machine code", runtime.GOOS)
}
