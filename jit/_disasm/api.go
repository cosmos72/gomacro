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
 * api.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"fmt"

	"github.com/bnagy/gapstone"
)

type Engine = gapstone.Engine

type Arch int

const (
	AMD64 = Arch(gapstone.CS_ARCH_X86)
	ARM64 = Arch(gapstone.CS_ARCH_ARM64)
)

func New(arch Arch) (Engine, error) {
	var mode uint = gapstone.CS_MODE_64
	if arch == ARM64 {
		mode = gapstone.CS_MODE_ARM // | gapstone.CS_MODE_V8
	}
	engine, err := gapstone.New(
		int(arch),
		mode,
	)
	if err != nil {
		return engine, err
	}
	engine.SetOption(gapstone.CS_OPT_SYNTAX, gapstone.CS_OPT_SYNTAX_ATT)
	return engine, nil
}

func Show(insn gapstone.Instruction) {
	fmt.Printf("0x%x:\t%x%s%s\t%s\n", insn.Address, insn.Bytes, spaces(2*len(insn.Bytes)), insn.Mnemonic, insn.OpStr)
}

func spaces(n int) string {
	return "                "[n%16:]
}

func Disasm(arch Arch, code []uint8) ([]gapstone.Instruction, error) {
	engine, err := New(arch)
	if err != nil {
		return nil, err
	}
	return engine.Disasm(code, 0x10000, 0)
}

func PrintDisasm(arch Arch, code []uint8) {
	insns, err := Disasm(arch, code)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Disasm:\n")
		for _, insn := range insns {
			Show(insn)
		}
	}
}
