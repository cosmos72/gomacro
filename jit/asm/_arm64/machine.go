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
 * machine.go
 *
 *  Created on May 26, 2018
 *      Author Massimiliano Ghilardi
 */

package arm64

import (
	"fmt"
)

type Arm64 struct {
}

func (Arm64) Name() string {
	return "arm64"
}

const (
	arm64_noregid RegId = iota
	X0
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	X29
	X30
	XZR, XSP  = iota, iota // depending on context, zero register or stack pointer
	arm64_rlo = X0
	arm64_rhi = XZR
	// stack pointer
	arm64_rsp = XSP
	// suggested register to point to local variables
	arm64_rvar = X29
)

var arm64 = struct {
	regName4   []string
	regName8   []string
	alwaysLive RegIds
}{
	regName4: arm64_makeRegNames("w"),
	regName8: arm64_makeRegNames("x"),
	alwaysLive: RegIds{
		X28: 1, // pointer to goroutine-local data
		X30: 1, // return address register
		XZR: 1, // zero register / stack pointer
	},
}

func (Arm64) RegIdValid(id RegId) bool {
	return id >= arm64_rlo && id < arm64_rhi // XZR/XSP is valid only in few, hand-checked cases
}

func arm64_makeRegNames(prefix string) []string {
	name := make([]string, arm64_rhi+1)
	for id := arm64_rlo; id < arm64_rhi; id++ {
		name[id] = fmt.Sprint(prefix, int(id)-1)
	}
	name[arm64_rhi] = prefix + "zr"
	return name
}

func (Arm64) RegIdString(id RegId) string {
	var s string
	if id >= arm64_rlo && id <= arm64_rhi {
		s = arm64.regName8[id]
	}
	if s == "" {
		s = fmt.Sprintf("unknown_reg(%d)", uint8(id))
	}
	return s
}

// return the bitmask to be or-ed to the instruction
// to specify the registers width
func (k Kind) arm64_kbit() uint32 {
	return uint32(k.Size()) & 8 << 28
}

func (Arm64) RegString(r Reg) string {
	var s string
	id := r.id
	if id >= arm64_rlo && id <= arm64_rhi {
		switch r.kind.Size() {
		case 1, 2, 4:
			s = arm64.regName4[id]
		case 8:
			s = arm64.regName8[id]
		}
	}
	if s == "" {
		s = fmt.Sprintf("unknown_reg(%d,%v)", uint8(id), r.kind)
	}
	return s
}

// validate and return uint32 mask representing r.id
// note that XSP/XZR is not considered valid
func (r Reg) arm64_val() uint32 {
	r.Validate()
	return uint32(r.id) - 1
}

// validate and return uint32 mask representing r.id
// if allowX31 is true, also allows r.id == XSP/XZR
func (r Reg) arm64_valOrX31(allowX31 bool) uint32 {
	if !allowX31 || r.id != XZR {
		r.Validate()
	}
	return uint32(r.id) - 1
}

// Prologue used to add the following instruction to generated code,
// but now it does nothing, because adding ANY code is the user's responsibility:
//   ldr x29, [sp, #8]
// equivalent to:
// asm.Asm(MOV, MakeMem(8, XSP, Uint64), MakeReg(X29, Uint64))
func (Arm64) Prologue(asm *Asm) *Asm {
	// return asm.Uint32(0xf94007fd)
	// equivalent:
	// return asm.Asm(MOV, MakeMem(8, XSP, Uint64), MakeReg(X29, Uint64))
	return asm
}

func (arch Arm64) Epilogue(asm *Asm) {
	arch.Op0(asm, RET)
}

func (Arm64) ArchInit(asm *Asm, start SaveSlot, end SaveSlot) {
	s := asm.save
	s.reg = Reg{XSP, Uint64}
	s.start, s.next, s.end = start, start, end
	s.bitmap = make([]bool, end-start)
}

// print arm64 machine code as sequence of 4-byte instructions
func (c Code) String() string {
	const hexdigit string = "0123456789abcdef"
	i, j, n := 0, 0, len(c)
	buf := make([]byte, (n+3)/4*9)
	for ; i+4 <= n; i += 4 {
		buf[j+0] = hexdigit[c[i+3]>>4]
		buf[j+1] = hexdigit[c[i+3]&0xF]
		buf[j+2] = hexdigit[c[i+2]>>4]
		buf[j+3] = hexdigit[c[i+2]&0xF]
		buf[j+4] = hexdigit[c[i+1]>>4]
		buf[j+5] = hexdigit[c[i+1]&0xF]
		buf[j+6] = hexdigit[c[i+0]>>4]
		buf[j+7] = hexdigit[c[i+0]&0xF]
		buf[j+8] = ' '
		j += 9
	}
	for k := n - 1; k >= i; k-- {
		buf[j+0] = hexdigit[c[k]>>4]
		buf[j+1] = hexdigit[c[k]&0xF]
		j += 2
	}
	return string(buf[:j])
}
