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
 * reg.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

import (
	"reflect"
)

// ------------------- Reg -------------------------

const (
	noReg hwReg = iota
	rAX
	rCX
	rDX
	rBX
	rSP
	rBP
	rSI
	rDI
	rR8
	rR9
	rR10
	rR11
	rR12
	rR13
	rR14
	rR15
	rLo hwReg = rAX
	rHi hwReg = rR15
)

// implement Arg interface
func (r hwReg) Reg() hwReg {
	return r
}

func (r hwReg) Const() bool {
	return false
}

func (r hwReg) Kind() reflect.Kind {
	// update after implementing MMX and XMM
	return reflect.Int64
}

// ---------

func (r hwReg) Valid() bool {
	return r >= rAX && r <= rR15
}

func (r hwReg) Validate() {
	if !r.Valid() {
		errorf("invalid register: %d", r)
	}
}

func (r hwReg) bits() uint8 {
	r.Validate()
	return uint8(r) - 1
}

func (r hwReg) lo() uint8 {
	return r.bits() & 0x7
}

func (r hwReg) hi() uint8 {
	return (r.bits() & 0x8) >> 3
}

func (r hwReg) lohi() (uint8, uint8) {
	bits := r.bits()
	return bits & 0x7, (bits & 0x8) >> 3
}

// ------------------- Regs -------------------------

func newHwRegs(rs ...hwReg) *hwRegs {
	var ret hwRegs
	for _, r := range rs {
		ret.Set(r)
	}
	return &ret
}

func (rs *hwRegs) InitLive() {
	*rs = hwRegs{rSP: 1, rBP: 1, rDI: 1}
}

func (rs *hwRegs) Set(r hwReg) {
	if r >= rLo && r <= rHi {
		rs[r]++
	}
}

func (rs *hwRegs) Unset(r hwReg) {
	if rs.Contains(r) {
		rs[r]--
	}
}

func (rs *hwRegs) Contains(r hwReg) bool {
	return r >= rLo && r <= rHi && rs[r] != 0
}

func (rs *hwRegs) Alloc() hwReg {
	for r := rLo; r <= rHi; r++ {
		if rs[r] == 0 {
			rs[r]++
			return r
		}
	}
	errorf("no free registers")
	return noReg
}

func (rs *hwRegs) Free(r hwReg) {
	rs.Unset(r)
}
