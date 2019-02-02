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
 * gen_op2.go
 *
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

package main

import (
	"fmt"
	"io"

	arch "github.com/cosmos72/gomacro/jit/arm64"
)

type genOp2 struct {
	opname, opName string
	w              io.Writer
}

func newGenOp2(w io.Writer, opname string) *genOp2 {
	return &genOp2{
		opname: opname,
		opName: string(opname[0]-'a'+'A') + opname[1:],
		w:      w,
	}
}

func (g *genOp2) generate() {
	g.fileHeader()
	g.opRegReg()
	if g.opname == "mov" {
		g.opConstReg()
	}
}

func (g *genOp2) fileHeader() {
	fmt.Fprintf(g.w,
		`	.file	"%s.s"
	.text
`, g.opname)
}

func (g *genOp2) funcHeader(funcName string) {
	fmt.Fprintf(g.w,
		`
	.p2align 4,,15
	.globl	%s%s
	.type	%s%s, @function
%s%s:
	.cfi_startproc
`, g.opName, funcName, g.opName, funcName, g.opName, funcName)
}

func (g *genOp2) funcFooter() {
	fmt.Fprint(g.w, `	ret
	.cfi_endproc

`)
}

func (g *genOp2) opRegReg() {
	g.funcHeader("RegReg")
	for _, k := range [...]arch.Kind{arch.Uint32, arch.Uint64} {
		kbits := k.Size() * 8
		fmt.Fprintf(g.w, "\t// OP reg%d, reg%d\n", kbits, kbits)
		rlo := arch.MakeReg(arch.RLo, k)
		for id := arch.RLo; id < arch.RHi; id++ {
			fmt.Fprintf(g.w, "\t%s\t%v, %v\n", g.opname, arch.MakeReg(id, k), rlo)
		}
		fmt.Fprint(g.w, "\tnop\n")
		for id := arch.RLo; id < arch.RHi; id++ {
			fmt.Fprintf(g.w, "\t%s\t%v, %v\n", g.opname, rlo, arch.MakeReg(id, k))
		}
		fmt.Fprint(g.w, "\tnop\n")
	}
	g.funcFooter()
}

func (g *genOp2) opConstReg() {
	g.funcHeader("ConstReg")
	for _, k := range [...]arch.Kind{arch.Uint32, arch.Uint64} {
		kbits := k.Size() * 8
		fmt.Fprintf(g.w, "\t// OP reg%d, const\n", kbits)
		rlo := arch.MakeReg(arch.RLo, k)
		for val := 1; val <= 0x10000; val *= 2 {
			fmt.Fprintf(g.w, "\t%s\t%v, #%#x\n", g.opname, rlo, -val)
		}
		for val := 1; val < 0x10000; val *= 2 {
			fmt.Fprintf(g.w, "\t%s\t%v, #%#x\n", g.opname, rlo, val)
		}
		for id := arch.RLo; id < arch.RHi; id++ {
			fmt.Fprintf(g.w, "\t%s\t%v, #0x%x\n", g.opname, arch.MakeReg(id, k), 0)
		}
		fmt.Fprint(g.w, "\tnop\n")
	}
	g.funcFooter()
}
