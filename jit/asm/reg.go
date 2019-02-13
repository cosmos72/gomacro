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
 *  Created on May 24, 2018
 *      Author Massimiliano Ghilardi
 */

package asm

// machine register
type RegId uint8

func (id RegId) Arch() Arch {
	return Archs[ArchId(1+id>>7)]
}

func (id RegId) String() string {
	return id.Arch().RegIdString(id)
}

func (id RegId) Valid() bool {
	return id.Arch().RegIdValid(id)
}

func (id RegId) Validate() {
	if !id.Valid() {
		errorf("invalid register: %v", id)
	}
}

// ===================================

type RegIdCfg struct {
	RLo, RHi, RSP, RVAR RegId
}

// register + kind
type Reg struct {
	id   RegId
	kind Kind // defines width and signedness
}

func MakeReg(id RegId, kind Kind) Reg {
	return Reg{id: id, kind: kind}
}

// implement Arg interface
func (r Reg) RegId() RegId {
	return r.id
}

func (r Reg) Kind() Kind {
	return r.kind
}

func (r Reg) Const() bool {
	return false
}

func (r Reg) Valid() bool {
	return r.id.Valid()
}

func (r Reg) Validate() {
	r.id.Validate()
}

// ===================================

type RegIds struct {
	list []uint32 // RegId -> use count
	rlo  RegId
}

func (rs *RegIds) IsUsed(r RegId) bool {
	return r.Valid() && rs.list[r-rs.rlo] != 0
}

// return new use count
func (rs *RegIds) IncUse(r RegId) uint32 {
	if r.Valid() {
		addr := &rs.list[r-rs.rlo]
		if *addr < ^uint32(0) {
			*addr++
		}
		return *addr
	}
	return 0
}

// return new use count
func (rs *RegIds) DecUse(r RegId) uint32 {
	if r.Valid() {
		addr := &rs.list[r-rs.rlo]
		if *addr > 0 {
			*addr--
		}
		return *addr
	}
	return 0
}

// ===================================

func (asm *Asm) RegIsUsed(id RegId) bool {
	return asm.regIds.IsUsed(id)
}

// return new use count
func (asm *Asm) RegIncUse(id RegId) uint32 {
	return asm.regIds.IncUse(id)
}

// return new use count
func (asm *Asm) RegDecUse(id RegId) uint32 {
	return asm.regIds.DecUse(id)
}
