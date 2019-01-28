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
 * op2misc.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

type Op2Misc uint8

const (
	PUSH Op2Misc = iota
	POP
)

var op2MiscName = map[Op2Misc]string{
	PUSH: "PUSH",
	POP:  "POP",
}

func (op Op2Misc) String() string {
	s, ok := op2MiscName[op]
	if !ok {
		s = fmt.Sprintf("Op2Misc(%d)", int(op))
	}
	return s
}

func (asm *Asm) Op2Misc(op Op2Misc, arg1 interface{}, arg2 interface{}) *Asm {
	switch op {
	case PUSH:
		asm.Push(arg1.(RegId), arg2.(*bool))
	case POP:
		asm.Pop(arg1.(RegId), arg2.(bool))
	default:
		errorf("unknown Op2Misc operation: %v %v %v", op, arg1, arg2)
	}
	return asm
}

func (asm *Asm) Push(id RegId, pushed *bool) *Asm {
	id.Validate()
	if asm.RegIds[id] == 0 {
		// mark in use, caller wants this register
		asm.RegIds[id]++
		*pushed = false
		return asm
	}
	s := asm.save
	if s.idx >= s.end {
		errorf("register save area is full, cannot push register %v", id)
	}
	asm.archPush(id)
	s.idx++
	*pushed = false
	return asm
}

func (asm *Asm) Pop(id RegId, pushed bool) *Asm {
	id.Validate()
	if !pushed {
		if asm.RegIds[id] > 0 {
			asm.RegIds[id]--
		}
		return asm
	}
	s := asm.save
	if s.idx <= s.start {
		errorf("register save area is empty, cannot pop register %v", id)
	}
	s.idx--
	asm.archPop(id)
	return asm
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
