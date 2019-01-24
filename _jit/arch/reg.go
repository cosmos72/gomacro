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

package arch

// hardware register. implementation is architecture-dependent
type RegId uint8

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

// ===================================

type RegIds [RHi + 1]uint32 // Reg -> use count

func newRegs(ids ...RegId) *RegIds {
	var ret RegIds
	for _, id := range ids {
		ret.Set(id)
	}
	return &ret
}

func (rs *RegIds) InitLive() {
	*rs = alwaysLiveRegIds
}

func (rs *RegIds) Contains(r RegId) bool {
	return r >= RLo && r <= RHi && rs[r] != 0
}

func (rs *RegIds) Set(r RegId) {
	if r >= RLo && r <= RHi {
		rs[r]++
	}
}

func (rs *RegIds) Unset(r RegId) {
	if rs.Contains(r) {
		rs[r]--
	}
}
