// +build amd64

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
 * asm_amd64_mov.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

func (asm *Asm) Mov(dst Arg, src Arg) *Asm {
	return asm.Op2(MOV, dst, src)
}

// %reg_dst = const
func (asm *Asm) movRegConst(dst Reg, c Const) *Asm {
	if c.val == 0 {
		return asm.xorRegSelf(dst)
	}
	dlo, dhi := dst.lohi()
	// 32-bit signed immediate constants, use mov
	if c.val == int64(int32(c.val)) {
		return asm.Bytes(0x48|dhi, 0xC7, 0xC0|dlo).Int32(int32(c.val))
	}
	// 64-bit constant, must use movabs
	return asm.Bytes(0x48|dhi, 0xB8|dlo).Int64(c.val)
}

// %reg_dst ^= %reg_dst // compact way to zero a register
func (asm *Asm) xorRegSelf(dst Reg) *Asm {
	dlo, dhi := dst.lohi()
	if dhi == 0 {
		return asm.Bytes(0x31, 0xC0|dlo|dlo<<3)
	} else {
		return asm.Bytes(0x48|dhi<<1|dhi<<2, 0x31, 0xC0|dlo|dlo<<3)
	}
}

// movsx, movzx or mov
func (asm *Asm) Cast(dst Arg, src Arg) *Asm {
	if src == dst {
		return asm
	} else if src.Kind().Size() == dst.Kind().Size() {
		return asm.Op2(MOV, dst, src)
	}
	switch dst := dst.(type) {
	case Reg:
		switch src := src.(type) {
		case Reg:
			asm.castRegReg(dst, src)
		case Mem:
			asm.castRegMem(dst, src)
		case Const:
			src = src.Cast(dst.kind)
			asm.movRegConst(dst, src)
		default:
			errorf("Cast: unsupported source type, expecting Reg, Mem or Const: %v %v %v", CAST, dst, src)
		}
	case Mem:
		switch src := src.(type) {
		case Reg:
			// assume that user code cannot use the same register
			// twice with different kinds
			r := MakeReg(src.id, dst.Kind())
			asm.castRegReg(r, src)
			asm.op2MemReg(MOV, dst, r)
		case Mem:
			r := asm.alloc(dst.Kind())
			asm.castRegMem(r, src)
			asm.op2MemReg(MOV, dst, r)
			asm.free(r)
		case Const:
			src = src.Cast(dst.Kind())
			asm.op2MemConst(MOV, dst, src)
		default:
			panicf("Cast: unsupported source type, expecting Reg, Mem or Const: %v %v %v", CAST, dst, src)
		}
	case Const:
		panicf("Cast: destination cannot be a constant: %v %v %v", CAST, dst, src)
	default:
		panicf("Cast: unsupported destination type, expecting Reg or Mem: %v %v %v", CAST, dst, src)
	}
	return asm
}

func (asm *Asm) castRegReg(dst Reg, src Reg) *Asm {
	// TODO
	return asm
}

func (asm *Asm) castRegMem(dst Reg, src Mem) *Asm {
	// TODO
	return asm
}
