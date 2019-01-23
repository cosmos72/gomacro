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

func newRegs(rs ...Reg) *Regs {
	var ret Regs
	for _, r := range rs {
		ret.Set(r)
	}
	return &ret
}

func (rs *Regs) InitLive() {
	*rs = alwaysLiveRegs
}

func (rs *Regs) Contains(r Reg) bool {
	return r >= RLo && r <= RHi && rs[r] != 0
}

func (rs *Regs) Set(r Reg) {
	if r >= RLo && r <= RHi {
		rs[r]++
	}
}

func (rs *Regs) Unset(r Reg) {
	if rs.Contains(r) {
		rs[r]--
	}
}
