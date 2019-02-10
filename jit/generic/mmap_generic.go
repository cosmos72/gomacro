// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!windows

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
 * os_generic.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"runtime"
)

const MMAP_SUPPORTED = false

type memarea struct {
}

func (asm *Asm) mmap() memarea {
	errorf("Asm: unsupported operating system %v, cannot mmap() created code", runtime.GOOS)
	return memarea{}
}
