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
 * emit.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package arch

const (
	VERBOSE = false
)

func NewAsm() *Asm {
	var asm Asm
	return asm.Init()
}

func (asm *Asm) Init() *Asm {
	return asm.Init2(0, 0)
}

func (asm *Asm) Init2(saveStart, saveEnd SaveSlot) *Asm {
	asm.code = asm.code[:0:cap(asm.code)]
	asm.nextRegId = RLo
	asm.softRegs = make(SoftRegs)
	asm.save.ArchInit(saveStart, saveEnd)
	asm.regIds.InitLive()
	return asm.Prologue()
}

func (asm *Asm) Code() Code {
	return asm.code
}

func (asm *Asm) Byte(b byte) *Asm {
	asm.code = append(asm.code, b)
	return asm
}

func (asm *Asm) Bytes(bytes ...byte) *Asm {
	asm.code = append(asm.code, bytes...)
	return asm
}

func (asm *Asm) Uint8(val uint8) *Asm {
	asm.code = append(asm.code, val)
	return asm
}

func (asm *Asm) Uint16(val uint16) *Asm {
	asm.code = append(asm.code, uint8(val), uint8(val>>8))
	return asm
}

func (asm *Asm) Uint32(val uint32) *Asm {
	asm.code = append(asm.code, uint8(val), uint8(val>>8), uint8(val>>16), uint8(val>>24))
	return asm
}

func (asm *Asm) Uint64(val uint64) *Asm {
	asm.code = append(asm.code, uint8(val), uint8(val>>8), uint8(val>>16), uint8(val>>24), uint8(val>>32), uint8(val>>40), uint8(val>>48), uint8(val>>56))
	return asm
}

func (asm *Asm) Int8(val int8) *Asm {
	return asm.Uint8(uint8(val))
}

func (asm *Asm) Int16(val int16) *Asm {
	return asm.Uint16(uint16(val))
}

func (asm *Asm) Int32(val int32) *Asm {
	return asm.Uint32(uint32(val))
}

func (asm *Asm) Int64(val int64) *Asm {
	return asm.Uint64(uint64(val))
}

// ===================================

func (asm *Asm) tryRegAlloc(kind Kind) Reg {
	var id RegId
	for {
		if asm.nextRegId > RHi {
			return Reg{}
		}
		id = asm.nextRegId
		asm.nextRegId++
		if asm.regIds[id] == 0 {
			asm.regIds[id] = 1
			break
		}
	}
	return Reg{id: id, kind: kind}
}

func (asm *Asm) RegAlloc(kind Kind) Reg {
	r := asm.tryRegAlloc(kind)
	if !r.Valid() {
		errorf("no free register")
	}
	return r
}

func (asm *Asm) RegFree(r Reg) *Asm {
	count := asm.regIds[r.id]
	if count <= 0 {
		return asm
	}
	count--
	asm.regIds[r.id] = count
	if count == 0 && asm.nextRegId > r.id {
		asm.nextRegId = r.id
	}
	return asm
}

// ===================================

// convert SoftReg to Arg
func (asm *Asm) Arg(x interface{}) Arg {
	switch x := x.(type) {
	case SoftReg:
		return asm.SoftReg(x)
	case Arg:
		return x
	default:
		errorf("unknown argument type %T, expecting Const, Reg, Mem or SoftReg", x)
		return nil
	}
}

// convert SoftReg to Arg
func (asm *Asm) SoftReg(s SoftReg) Arg {
	a := asm.softRegs[s]
	if a == nil {
		errorf("soft register %v not allocated", s)
	}
	return a
}

func (asm *Asm) Alloc(s SoftReg, kind Kind) Arg {
	var a Arg
	if r := asm.tryRegAlloc(kind); r.Valid() {
		a = r
	} else {
		idx := asm.save.Alloc()
		if idx == InvalidSlot {
			errorf("no free register, and save area is full. Cannot allocate soft register %v", s)
		}
		a = MakeMem(int32(idx)*8, asm.save.reg.id, kind)
	}
	asm.softRegs[s] = a
	return a
}

func (asm *Asm) Free(s SoftReg) {
	a := asm.softRegs[s]
	if a == nil {
		errorf("cannot free unallocated soft register %v", s)
	}
	switch a := a.(type) {
	case Reg:
		asm.RegFree(a)
	case Mem:
		asm.save.Free(SaveSlot(a.off / 8))
	default:
		errorf("soft register %v is mapped to unknown type %T, expecting Reg or Mem", s, a)
	}
	delete(asm.softRegs, s)
}
