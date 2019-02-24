/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * func.go
 *
 *  Created on Feb 07, 2019
 *      Author Massimiliano Ghilardi
 */

package common

import (
	"reflect"
	"unsafe"
)

const MMAP_VERBOSE = false

type interfaceHeader struct {
	typ  uintptr
	addr **MemArea
}

/**
 * convert code created by the programmer to a callable function.
 * funcaddr must be a non-nil pointer to function.
 *
 * function type MUST match the code created by the programmer,
 * or BAD things will happen: crash, memory corruption, undefined behaviour...
 *
 * Obviously, code created by the programmer must be for the same architecture
 * the program is currently running on...
 */
func (asm *Asm) Func(funcaddr interface{}) {
	v := reflect.ValueOf(funcaddr)
	if !v.IsValid() || v.Kind() != reflect.Ptr || v.IsNil() || !v.Elem().CanSet() || v.Elem().Kind() != reflect.Func {
		errorf("Asm.Func() argument must be non-nil, settable pointer to function, received %p // %T", funcaddr, funcaddr)
	}
	header := *(*interfaceHeader)(unsafe.Pointer(&funcaddr))
	/*
		fmt.Printf("header = %+v\n", header)
		fmt.Printf("funcaddr = %p\n", header.addr)
		fmt.Printf("*funcaddr = %p\n", *header.addr)
	*/
	mem := asm.mmap()
	*header.addr = &mem
	/*
		fmt.Printf("&mem = %p\n", &mem)
		fmt.Printf("header = %+v\n", header)
		fmt.Printf("funcaddr = %p\n", header.addr)
		fmt.Printf("*funcaddr = %p\n", *header.addr)
	*/
}

// return MemArea containing executable machine code
func (asm *Asm) mmap() MemArea {
	asm.Epilogue()
	if MMAP_VERBOSE {
		debugf("asm: %#v", asm.code)
	}
	area := asm.code.MemArea()
	size := area.Size()
	if ret := asm.cache.Lookup(area); ret.Size() == size {
		return ret
	}
	if asm.pool.Size() < size {
		// we waste asm.mem.Size() bytes of mmapped memory...
		asm.pool = NewMemPool(size)
	}
	ret := asm.pool.Copy(area)
	if asm.cache == nil {
		asm.cache = make(Cache)
	}
	asm.cache.Add(ret)
	return ret
}
