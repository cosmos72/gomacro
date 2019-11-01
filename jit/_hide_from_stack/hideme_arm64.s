// +build gc,arm64

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * hideme_arm64.s
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

#include "textflag.h" // for NOSPLIT
#include "funcdata.h" // for NO_LOCAL_POINTERS

// use same restrictions as a JIT function:
// cannot have stack map, frame, local variables;
// must use one of call* trampolines to invoke arbitrary Go functions
TEXT ·asm_hideme(SB),NOSPLIT|NOFRAME,$0-8
	// manually save return_addr R30 in callee's stack
	MOVD LR, return_addr-24(SP)

	MOVD env+0(FP), R0           // received argument *Env
	MOVD 0(R0),  R26             // closure, must be in R26
	MOVD 8(R0),  R1              // closure arg
	MOVD 24(R0), R2              // helper function: call[1] == call16
	MOVD R1, arg-40(SP)          // store closure arg on callee's stack
	CALL R2

	// manually load return_addr R30 from callee's stack
	MOVD return_addr-24(SP), LR
	RET
