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
 * mov.go
 *
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

// ============================================================================
func (asm *Asm) Mov(src Arg, dst Arg) *Asm {
	assert(SizeOf(src) == SizeOf(dst))

	if dst.Const() {
		errorf("destination cannot be a constant: %v %v, %v", MOV, src, dst)
	}
	if src == dst {
		return asm
	}

	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Const:
			asm.movConstReg(src, dst)
		case Reg:
			asm.movRegReg(src, dst)
		case Mem:
			asm.Load(src, dst)
		default:
			errorf("unknown source type %T, expecting Const, Reg or Mem: %v %v, %v", src, MOV, src, dst)
		}
	case Mem:
		switch src := src.(type) {
		case Const:
			asm.movConstMem(src, dst)
		case Reg:
			asm.Store(src, dst)
		case Mem:
			asm.movMemMem(src, dst)
		default:
			errorf("unknown source type %T, expecting Const, Reg or Mem: %v %v, %v", src, MOV, src, dst)
		}
	default:
		errorf("unknown destination type %T, expecting Reg or Mem: %v %v, %v", dst, MOV, src, dst)
	}
	return asm
}

type loadstore uint32

const (
	load  loadstore = 0x39400000
	store loadstore = 0x39000000
)

func (asm *Asm) Load(src Mem, dst Reg) *Asm {
	return asm.loadstore(load, src, dst)
}

func (asm *Asm) Store(src Reg, dst Mem) *Asm {
	return asm.loadstore(store, dst, src)
}

func (asm *Asm) movRegReg(src Reg, dst Reg) *Asm {
	// arm64 implements "mov src,dst" as "orr xzr,src,dst"
	return asm.Uint32(dst.kind.kbit() | 0x2A0003E0 | src.valOrX31(true)<<16 | dst.val())
}

func (asm *Asm) movConstReg(c Const, dst Reg) *Asm {
	xzr := MakeReg(XZR, dst.kind)
	var immcval uint32
	var movk bool
	if c.val >= 0 && c.val < 0x10000 {
		immcval = 0x40<<19 | uint32(c.val)
	} else if c.val < 0 && c.val >= -0x10000 {
		immcval = uint32(^c.val)
	} else if asm.tryOp3RegConstReg(OR3, xzr, uint64(c.val), dst) {
		return asm
	} else if asm.tryOp3RegConstReg(OR3, xzr, uint64(uint32(c.val)), dst) {
		if dst.kind.Size() == 8 {
			asm.movk(uint16(c.val>>32), 32, dst)
			asm.movk(uint16(c.val>>48), 48, dst)
		}
		return asm
	} else {
		immcval = 0x40<<19 | uint32(c.val&0xFFFF)
		movk = true
	}
	asm.Uint32(dst.kind.kbit() | 0x12800000 | immcval<<5 | dst.val())
	if movk {
		asm.movk(uint16(c.val>>16), 16, dst)
		if dst.kind.Size() == 8 {
			asm.movk(uint16(c.val>>32), 32, dst)
			asm.movk(uint16(c.val>>48), 48, dst)
		}
	}
	return asm
}

// set some bits of dst, preserving others
func (asm *Asm) movk(val uint16, shift uint8, dst Reg) *Asm {
	if val != 0 {
		asm.Uint32(dst.kind.kbit() | 0xF2800000 | uint32(shift)<<17 | uint32(val)<<5 | dst.val())
	}
	return asm
}

func (asm *Asm) loadstore(op loadstore, m Mem, r Reg) *Asm {
	off := m.off
	var sizebit uint32
	switch m.Kind().Size() {
	case 1:
		if off >= 0 && off <= 4095 {
			return asm.Uint32(sizebit | uint32(op) | uint32(m.off)<<10 | m.reg.valOrX31(true)<<5 | r.valOrX31(true))
		}
	case 2:
		sizebit = 0x4 << 28
		if off >= 0 && off <= 8190 && off%2 == 0 {
			return asm.Uint32(sizebit | uint32(op) | uint32(m.off)<<9 | m.reg.valOrX31(true)<<5 | r.valOrX31(true))
		}
	case 4:
		sizebit = 0x8 << 28
		if off >= 0 && off <= 16380 && off%4 == 0 {
			return asm.Uint32(sizebit | uint32(op) | uint32(m.off)<<8 | m.reg.valOrX31(true)<<5 | r.valOrX31(true))
		}
	case 8:
		sizebit = 0xC << 28
		if off >= 0 && off <= 32760 && off%8 == 0 {
			return asm.Uint32(sizebit | uint32(op) | uint32(m.off)<<7 | m.reg.valOrX31(true)<<5 | r.valOrX31(true))
		}
	}
	// load offset in a register. we could also try "ldur" or "stur"...
	tmp := asm.RegAlloc(Uint64)
	asm.movConstReg(ConstInt64(int64(off)), tmp)

	asm.Uint32(sizebit | uint32(op^0x1206800) | tmp.val()<<16 | m.reg.valOrX31(true)<<5 | r.val())

	return asm.RegFree(tmp)
}

func (asm *Asm) movConstMem(c Const, dst Mem) *Asm {
	if c.val == 0 {
		return asm.zeroMem(dst)
	}
	r := asm.RegAlloc(dst.Kind())
	return asm.movConstReg(c, r).Store(r, dst).RegFree(r)
}

func (asm *Asm) movMemMem(src Mem, dst Mem) *Asm {
	r := asm.RegAlloc(src.Kind())
	return asm.Load(src, r).Store(r, dst).RegFree(r)
}

// ============================================================================
func (asm *Asm) Cast(src Arg, dst Arg) *Asm {
	if src == dst {
		return asm
	} else if SizeOf(src) == SizeOf(dst) {
		return asm.Op2(MOV, src, dst)
	}

	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Reg:
			asm.castRegReg(src, dst)
		case Mem:
			asm.castMemReg(src, dst)
		case Const:
			src = src.Cast(dst.kind)
			asm.movConstReg(src, dst)
		default:
			errorf("unsupported source type %T, expecting Reg, Mem or Const: %v %v %v", src, CAST, src, dst)
		}
	case Mem:
		switch src := src.(type) {
		case Reg:
			asm.castRegMem(src, dst)
		case Mem:
			asm.castMemMem(src, dst)
		case Const:
			src = src.Cast(dst.Kind())
			asm.movConstMem(src, dst)
		default:
			errorf("unsupported source type %T, expecting Reg, Mem or Const: %v %v %v", src, CAST, src, dst)
		}
	case Const:
		errorf("destination cannot be a constant: %v %v %v", CAST, src, dst)
	default:
		errorf("unsupported destination type %T, expecting Reg or Mem: %v %v %v", dst, CAST, src, dst)
	}
	return asm
}

func (asm *Asm) castMemMem(src Mem, dst Mem) *Asm {
	r1 := asm.RegAlloc(src.Kind())
	r2 := MakeReg(r1.id, dst.Kind())
	return asm.Load(src, r1).castRegReg(r1, r2).Store(r2, dst).RegFree(r1)
}

func (asm *Asm) castMemReg(src Mem, dst Reg) *Asm {
	r := MakeReg(dst.id, src.Kind())
	asm.Load(src, r)
	return asm.castRegReg(r, dst)
}

func (asm *Asm) castRegMem(src Reg, dst Mem) *Asm {
	r := MakeReg(src.id, dst.Kind())
	if SizeOf(src) < SizeOf(dst) {
		// extend src. we can safely overwrite its high bits: they are junk
		return asm.castRegReg(src, r).Store(r, dst)
	} else {
		// just ignore src high bits
		return asm.Store(r, dst)
	}
}

func (asm *Asm) castRegReg(src Reg, dst Reg) *Asm {
	skind := src.kind
	dkind := dst.kind
	ssize := skind.Size()
	dsize := dkind.Size()
	if ssize >= dsize {
		// truncate. easy, just ignore src high bits
		return asm.Mov(MakeReg(src.id, dst.kind), dst)
	} else if skind.Signed() {
		// sign-extend: use one of
		// use "sxtb	src, dst"
		// or  "sxth	src, dst"
		// or  "sxtw	src, dst"
		errorf("unimplemented sign-extend: %v %v, %v", CAST, src, dst)
		return asm
	} else {
		// zero-extend
		if ssize == 4 {
			// zero-extend 32 bit -> 64 bit: use
			// "mov dst, src"
			// must be kept even if src == dst to zero high bits,
			// so use Asm.movRegReg() instead of too smart Asm.Mov()
			return asm.movRegReg(src, MakeReg(dst.id, skind))
		}
		// zero-extend, src is 8 bit or 16 bit. use one of:
		// "and dst, src, #0xff"
		// "and dst, src, #0xffff"
		if dsize <= 4 {
			dkind = Uint32
		}
		r := MakeReg(src.id, dkind)
		c := MakeConst(int64(0xffff)>>(16-ssize*8), dkind)
		return asm.op3RegConstReg(AND3, r, c, dst)
	}
}
