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
 * _gen_op2.go
 *
 *  Created on Jan 28, 2019
 *      Author Massimiliano Ghilardi
 */

package main

import (
	"fmt"
	"io"

	"github.com/cosmos72/gomacro/jit/arch"
)

type genOp1 struct {
	opname, opName string
	w              io.Writer
}

func NewGenOp1(w io.Writer, opname string) *genOp1 {
	return &genOp1{
		opname: opname,
		opName: string(opname[0]-'a'+'A') + opname[1:],
		w:      w,
	}
}

func (g *genOp1) generate() {
	g.fileHeader()
	g.opReg()
	g.opMem()
}

func (g *genOp1) fileHeader() {
	fmt.Fprintf(g.w,
		`	.file	"%s.s"
	.text
`, g.opname)
}

func (g *genOp1) funcHeader(funcName string) {
	fmt.Fprintf(g.w,
		`
	.p2align 4,,15
	.globl	%s%s
	.type	%s%s, @function
%s%s:
	.cfi_startproc
`, g.opName, funcName, g.opName, funcName, g.opName, funcName)
}

func (g *genOp1) funcFooter() {
	fmt.Fprint(g.w, `	ret
	.cfi_endproc

`)
}

func (g *genOp1) opReg() {
	g.funcHeader("Reg")
	for _, k := range [...]arch.Kind{arch.Uint8, arch.Uint16, arch.Uint32, arch.Uint64} {
		fmt.Fprintf(g.w, "\t// OP reg%d\n", k.Size()*8)
		for r := arch.RLo; r <= arch.RHi; r++ {
			fmt.Fprintf(g.w, "\t%s\t%v\n", g.opname, arch.MakeReg(r, k))
		}
		fmt.Fprint(g.w, "\tnop\n")
	}
	g.funcFooter()
}

func (g *genOp1) opMem() {
	for _, k := range [...]arch.Kind{arch.Uint8, arch.Uint16, arch.Uint32, arch.Uint64} {
		g.opMemKind(k)
	}
}

func (g *genOp1) opMemKind(k arch.Kind) {
	ksuffix := map[arch.Size]string{1: "b", 2: "w", 4: "l", 8: "q"}
	klen := k.Size() * 8
	g.funcHeader(fmt.Sprintf("Mem%d", klen))
	offstr := [...]string{"", "0x7F", "0x78563412"}
	for i, offlen := range [...]uint8{0, 8, 32} {
		fmt.Fprintf(g.w, "\t// OP mem%d[off%d]\n", klen, offlen)
		for r := arch.RLo; r <= arch.RHi; r++ {
			fmt.Fprintf(g.w, "\t%s%s\t%s(%v)\n", g.opname, ksuffix[k.Size()],
				offstr[i], arch.MakeReg(r, arch.Uintptr))
		}
		fmt.Fprint(g.w, "\tnop\n")
	}
	g.funcFooter()
}
