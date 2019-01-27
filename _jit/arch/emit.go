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

func (s *Save) Init(start, end uint16) {
	s.start, s.idx, s.end = start, start, end
}

func (asm *Asm) Init() *Asm {
	return asm.Init2(0, 0)
}

func (asm *Asm) Init2(saveStart, saveEnd uint16) *Asm {
	asm.code = asm.code[:0:cap(asm.code)]
	asm.RegIds.InitLive()
	asm.NextRegId = RLo
	asm.Save.Init(saveStart, saveEnd)
	return asm.Prologue()
}

func (asm *Asm) Code() Code {
	return asm.code
}

func (asm *Asm) Bytes(bytes ...uint8) *Asm {
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

/*
func (asm *Asm) pushRegs(rs *Regs) *Regs {
	var ret Regs
	v := &Var{}
	for r := Lo; r <= Hi; r++ {
		if !rs.Contains(r) {
			continue
		}
		if asm.Save.idx >= asm.Save.end {
			errorf("save area is full, cannot push registers")
		}
		v.idx = asm.save.idx
		asm.storeReg(v, r)
		asm.save.idx++
		ret.Set(r)
	}
	return &ret
}

func (asm *Asm) popRegs(rs *Regs) {
	v := &Var{}
	for r := rHi; r >= rLo; r-- {
		if !rs.Contains(r) {
			continue
		}
		if asm.save.idx <= asm.save.start {
			errorf("save area is empty, cannot pop registers")
		}
		asm.save.idx--
		v.idx = asm.save.idx
		asm.load(r, v)
	}
}
*/

func (asm *Asm) RegAlloc(kind Kind) Reg {
	var id RegId
	for {
		if asm.NextRegId > RHi {
			errorf("no free register")
		}
		id = asm.NextRegId
		asm.NextRegId++
		if asm.RegIds[id] == 0 {
			asm.RegIds[id] = 1
			break
		}
	}
	return Reg{id: id, kind: kind}
}

func (asm *Asm) Alloc(a Arg) (r Reg, allocated bool) {
	id := a.RegId()
	if id != NoRegId {
		return Reg{id: id, kind: a.Kind()}, false
	}
	return asm.RegAlloc(a.Kind()), true
}

// combined Alloc + Load
func (asm *Asm) AllocLoad(a Arg) (r Reg, allocated bool) {
	r, allocated = asm.Alloc(a)
	if allocated {
		asm.Mov(r, a)
	}
	return r, allocated
}

func (asm *Asm) RegFree(r Reg) *Asm {
	count := asm.RegIds[r.id]
	if count <= 0 {
		return asm
	}
	count--
	asm.RegIds[r.id] = count
	if count == 0 && asm.NextRegId > r.id {
		asm.NextRegId = r.id
	}
	return asm
}

func (asm *Asm) Free(r Reg, allocated bool) *Asm {
	if r.Valid() && allocated {
		asm.RegFree(r)
	}
	return asm
}

// combined Store + Free
func (asm *Asm) StoreFree(a Arg, r Reg, allocated bool) *Asm {
	if allocated {
		asm.Mov(a, r)
		asm.RegFree(r)
	}
	return asm
}
