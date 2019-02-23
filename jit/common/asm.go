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
 * emit.go
 *
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package common

type Asm struct {
	code          MachineCode
	nextSoftRegId SoftRegId // first available soft register
	softRegs      SoftRegIds
	save          Save
	regIds        RegIds
	arch          Arch
	mem           *MemPool
}

func New(id ArchId) *Asm {
	var asm Asm
	return asm.InitArchId(id)
}

func NewArch(arch Arch) *Asm {
	var asm Asm
	return asm.InitArch(arch)
}

func (asm *Asm) ArchId() ArchId {
	if asm.arch == nil {
		return NOARCH
	}
	return asm.arch.Id()
}

func (asm *Asm) Arch() Arch {
	return asm.arch
}

func (asm *Asm) InitArchId(archId ArchId) *Asm {
	return asm.InitArch2(Archs[archId], 0, 0)
}

func (asm *Asm) InitArchId2(archId ArchId, saveStart SaveSlot, saveEnd SaveSlot) *Asm {
	return asm.InitArch2(Archs[archId], saveStart, saveEnd)
}

func (asm *Asm) InitArch(arch Arch) *Asm {
	return asm.InitArch2(arch, 0, 0)
}

func (asm *Asm) InitArch2(arch Arch, saveStart SaveSlot, saveEnd SaveSlot) *Asm {
	if arch == nil {
		errorf("unknown arch")
	}
	id := arch.Id()
	if Archs[id] == nil {
		Archs[id] = arch
	}
	config := arch.RegIdConfig()
	asm.arch = arch
	asm.code = MachineCode{ArchId: id}
	asm.nextSoftRegId = 0
	asm.softRegs = make(SoftRegIds)
	s := asm.save
	s.start, s.next, s.end = saveStart, saveStart, saveEnd
	s.reg = Reg{config.RSP, Uint64}
	s.bitmap = make([]bool, saveEnd-saveStart)
	asm.regIds.inuse = make(map[RegId]uint32)
	asm.regIds.first = config.RAllocFirst
	asm.regIds.curr = config.RAllocFirst
	asm.regIds.rlo = config.RLo
	asm.regIds.rhi = config.RHi
	asm.mem = nil
	arch.Init(asm, saveStart, saveEnd)
	arch.Prologue(asm)
	return asm
}

func (asm *Asm) Code() MachineCode {
	return asm.code
}

func (asm *Asm) ClearCode() *Asm {
	asm.code.Bytes = nil
	return asm
}

func (asm *Asm) Byte(b byte) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, b)
	return asm
}

func (asm *Asm) Bytes(bytes ...byte) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, bytes...)
	return asm
}

func (asm *Asm) Uint8(val uint8) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, val)
	return asm
}

func (asm *Asm) Uint16(val uint16) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, uint8(val), uint8(val>>8))
	return asm
}

func (asm *Asm) Uint32(val uint32) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, uint8(val), uint8(val>>8), uint8(val>>16), uint8(val>>24))
	return asm
}

func (asm *Asm) Uint64(val uint64) *Asm {
	asm.code.Bytes = append(asm.code.Bytes, uint8(val), uint8(val>>8), uint8(val>>16), uint8(val>>24), uint8(val>>32), uint8(val>>40), uint8(val>>48), uint8(val>>56))
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

// convert AsmCode to Arg
func (asm *Asm) Arg(x AsmCode) Arg {
	switch x := x.(type) {
	case SoftRegId:
		return asm.SoftRegId(x)
	case Arg:
		return x
	default:
		errorf("unknown argument type %T, expecting Const, Reg, Mem or SoftRegId", x)
		return nil
	}
}

// convert SoftRegId to Arg
func (asm *Asm) SoftRegId(s SoftRegId) Arg {
	a := asm.softRegs[s]
	if a == nil {
		errorf("soft register %v not allocated", s)
	}
	return a
}

// allocate a SoftRegId
func (asm *Asm) Alloc(s SoftRegId, kind Kind) Arg {
	var a Arg
	if r := asm.TryRegAlloc(kind); r.Valid() {
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

func (asm *Asm) Free(s SoftRegId) {
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
