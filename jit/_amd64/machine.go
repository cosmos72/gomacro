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
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

// ============================================================================
// register
const (
	noregid RegId = 128 + iota
	RAX
	RCX
	RDX
	RBX
	RSP
	RBP
	RSI
	RDI
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
	RLo RegId = RAX
	RHi RegId = R15
	// suggested register to point to local variables
	RVAR = RSI
)

var regName1 = [...]string{
	RAX - noregid: "%al",
	RCX - noregid: "%cl",
	RDX - noregid: "%dl",
	RBX - noregid: "%bl",
	RSP - noregid: "%spl",
	RBP - noregid: "%bpl",
	RSI - noregid: "%sil",
	RDI - noregid: "%dil",
	R8 - noregid:  "%r8b",
	R9 - noregid:  "%r9b",
	R10 - noregid: "%r10b",
	R11 - noregid: "%r11b",
	R12 - noregid: "%r12b",
	R13 - noregid: "%r13b",
	R14 - noregid: "%r14b",
	R15 - noregid: "%r15b",
}
var regName2 = [...]string{
	RAX - noregid: "%ax",
	RCX - noregid: "%cx",
	RDX - noregid: "%dx",
	RBX - noregid: "%bx",
	RSP - noregid: "%sp",
	RBP - noregid: "%bp",
	RSI - noregid: "%si",
	RDI - noregid: "%di",
	R8 - noregid:  "%r8w",
	R9 - noregid:  "%r9w",
	R10 - noregid: "%r10w",
	R11 - noregid: "%r11w",
	R12 - noregid: "%r12w",
	R13 - noregid: "%r13w",
	R14 - noregid: "%r14w",
	R15 - noregid: "%r15w",
}
var regName4 = [...]string{
	RAX - noregid: "%eax",
	RCX - noregid: "%ecx",
	RDX - noregid: "%edx",
	RBX - noregid: "%ebx",
	RSP - noregid: "%esp",
	RBP - noregid: "%ebp",
	RSI - noregid: "%esi",
	RDI - noregid: "%edi",
	R8 - noregid:  "%r8d",
	R9 - noregid:  "%r9d",
	R10 - noregid: "%r10d",
	R11 - noregid: "%r11d",
	R12 - noregid: "%r12d",
	R13 - noregid: "%r13d",
	R14 - noregid: "%r14d",
	R15 - noregid: "%r15d",
}
var regName8 = [...]string{
	RAX - noregid: "%rax",
	RCX - noregid: "%rcx",
	RDX - noregid: "%rdx",
	RBX - noregid: "%rbx",
	RSP - noregid: "%rsp",
	RBP - noregid: "%rbp",
	RSI - noregid: "%rsi",
	RDI - noregid: "%rdi",
	R8 - noregid:  "%r8",
	R9 - noregid:  "%r9",
	R10 - noregid: "%r10",
	R11 - noregid: "%r11",
	R12 - noregid: "%r12",
	R13 - noregid: "%r13",
	R14 - noregid: "%r14",
	R15 - noregid: "%r15",
}

func bits(id RegId) uint8 {
	id.Validate()
	return uint8(id - noregid)
}

func lohiId(id RegId) (uint8, uint8) {
	bits := bits(id)
	return bits & 0x7, (bits & 0x8) >> 3
}

func lohi(r Reg) (uint8, uint8) {
	return lohiId(r.RegId())
}

// return number of assembler bytes needed to encode m.off
func offlen(m Mem, id RegId) (offlen uint8, offbit uint8) {
	moffset := m.Offset()
	switch {
	// (%rbp) and (%r13) registers must use 1-byte offset even if m.off == 0
	case moffset == 0 && id != RBP && id != R13:
		return 0, 0
	case moffset == int32(int8(moffset)):
		return 1, 0x40
	default:
		return 4, 0x80
	}
}

func quirk24(asm *Asm, id RegId) *Asm {
	if id == RSP || id == R12 {
		asm.Bytes(0x24) // amd64 quirk
	}
	return asm
}
