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
 * binary_arith.go
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */

package amd64

// ---------------- ADD --------------------

func (asm *Asm) IncInt64c(z Bind, val int64) *Asm {
	if val == 0 {
		return asm
	} else if val == int64(int32(val)) {
		return asm.Bytes(0x48, 0x81, 0x87).Idx(z).Uint32(uint32(val)) // addq   $val,z(%rdi)
	}
	asm.load_rax_const(val)
	asm.Bytes(0x48, 0x01, 0x87).Idx(z) //   addq    %rax,z(%rdi)
	return asm
}

func (asm *Asm) AddInt64c(z Bind, a Bind, val int64) *Asm {
	if a == z {
		return asm.IncInt64c(z, val)
	} else if val == 0 {
		return asm.Set(z, a)
	}
	asm.load_rax_const(val)
	asm.Bytes(0x48, 0x03, 0x87).Idx(a) //   addq    a(%rdi),%rax
	asm.store_rax(z)
	return asm
}

func (asm *Asm) IncInt64(z, a Bind) *Asm {
	asm.load_rax(a)
	asm.Bytes(0x48, 0x01, 0x87).Idx(z) //   addq    %rax,z(%rdi)
	return asm
}

func (asm *Asm) AddInt64(z, a, b Bind) *Asm {
	if a == z {
		return asm.IncInt64(z, b)
	} else if b == z {
		return asm.IncInt64(z, a)
	}
	asm.load_rax(a)
	asm.Bytes(0x48, 0x03, 0x87).Idx(b) //   addq   b(%rdi),%rax
	asm.store_rax(z)
	return asm
}

func (asm *Asm) IncUint64c(z Bind, val uint64) *Asm {
	return asm.IncInt64c(z, int64(val))
}

func (asm *Asm) AddUint64c(z Bind, a Bind, val uint64) *Asm {
	return asm.AddInt64c(z, a, int64(val))
}

func (asm *Asm) IncUint64(z, a Bind) *Asm {
	return asm.IncInt64(z, a)
}

func (asm *Asm) AddUint64(z, a, b Bind) *Asm {
	return asm.AddInt64(z, a, b)
}

// ---------------- SUB --------------------

func (asm *Asm) DecInt64c(z Bind, val int64) *Asm {
	if val == 0 {
		return asm
	} else if val == int64(int32(val)) {
		return asm.Bytes(0x48, 0x81, 0xaf).Idx(z).Int32(int32(val)) // subq   $val,z(%rdi)
	}
	asm.load_rax_const(-val)           //   // negate val!
	asm.Bytes(0x48, 0x01, 0x87).Idx(z) //   addq    %rax,z(%rdi)
	return asm
}

func (asm *Asm) SubInt64c(z Bind, a Bind, val int64) *Asm {
	if a == z {
		return asm.DecInt64c(z, val)
	} else if val == 0 {
		return asm.Set(z, a)
	}
	asm.load_rax_const(-val)           //   // negate val!
	asm.Bytes(0x48, 0x03, 0x87).Idx(a) //   addq    a(%rdi),%rax
	asm.store_rax(z)
	return asm
}

func (asm *Asm) DecInt64(z, a Bind) *Asm {
	if a == z {
		return asm.Zero(z)
	}
	asm.load_rax(a)
	asm.Bytes(0x48, 0x29, 0x87).Idx(z) //   subq    %rax,z(%rdi)
	return asm
}

func (asm *Asm) SubInt64(z, a, b Bind) *Asm {
	if a == z {
		return asm.DecInt64(z, b)
	}
	asm.load_rax(a)
	asm.Bytes(0x48, 0x2b, 0x87).Idx(b) //   subq    b(%rdi),%rax
	asm.store_rax(z)
	return asm
}

func (asm *Asm) DecUint64c(z Bind, val uint64) *Asm {
	return asm.DecInt64c(z, int64(val))
}

func (asm *Asm) SubUint64c(z Bind, a Bind, val uint64) *Asm {
	return asm.SubInt64c(z, a, int64(val))
}

func (asm *Asm) DecUint64(z, a Bind) *Asm {
	return asm.DecInt64(z, a)
}

func (asm *Asm) SubUint64(z, a, b Bind) *Asm {
	return asm.SubInt64(z, a, b)
}

// ---------------- MUL --------------------

func (asm *Asm) MulInt64c(z, a Bind, val int64) *Asm {
	if val == 0 {
		return asm.Zero(z)
	} else if val == 1 {
		return asm.Set(z, a)
	} else if val == int64(int32(val)) {
		asm.Bytes(0x48, 0x69, 0x87).Idx(a).Int32(int32(val)) //  imul   $val,a(%rdi),%rax
	} else {
		asm.load_rax_const(val)
		asm.Bytes(0x48, 0x0f, 0xaf, 0x87).Idx(a) //    imul   a(%rdi),%rax
	}
	asm.store_rax(z)
	return asm
}

func (asm *Asm) MulInt64(z, a, b Bind) *Asm {
	asm.load_rax(a)
	asm.Bytes(0x48, 0x0f, 0xaf, 0x87).Idx(b) //  imul   b(%rdi),%rax
	asm.store_rax(z)
	return asm
}

func (asm *Asm) MulUint64c(z, a Bind, val int64) *Asm {
	return asm.MulInt64c(z, a, val)
}
func (asm *Asm) MulUint64(z, a, b Bind) *Asm {
	return asm.MulInt64(z, a, b)
}

// ---------------- QUO --------------------

// FIXME: golang remainder rules are NOT the same as C !
func (asm *Asm) QuoInt64(z, a, b Bind) *Asm {
	asm.load_rax(a)
	asm.Bytes(0x48, 0x99)              //  cqto
	asm.Bytes(0x48, 0xf7, 0xbf).Idx(b) //  idivq  b(%rdi)
	asm.store_rax(z)
	return asm
}

func (asm *Asm) QuoUint64(z, a, b Bind) *Asm {
	asm.load_rax(a)
	asm.Bytes(0x31, 0xd2)              //  xor    %edx,%edx
	asm.Bytes(0x48, 0xf7, 0xb7).Idx(b) //  divq   b(%rdi)
	asm.store_rax(z)
	return asm
}

// ---------------- REM --------------------

// FIXME: golang remainder rules are NOT the same as C !
func (asm *Asm) RemInt64(z, a, b Bind) *Asm {
	asm.load_rax(a)
	asm.Bytes(0x48, 0x99)              //  cqto
	asm.Bytes(0x48, 0xf7, 0xbf).Idx(b) //  idivq  b(%rdi)
	asm.store_rdx(z)
	return asm
}

func (asm *Asm) RemUint64(z, a, b Bind) *Asm {
	asm.load_rax(a)
	asm.Bytes(0x31, 0xd2)              //  xor    %edx,%edx
	asm.Bytes(0x48, 0xf7, 0xb7).Idx(b) //  divq   b(%rdi)
	asm.store_rdx(z)
	return asm
}
