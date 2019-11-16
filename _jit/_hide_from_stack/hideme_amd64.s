// +build gc,amd64

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
 * hideme_amd64.s
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

#include "textflag.h" // for NOSPLIT
#include "funcdata.h" // for NO_LOCAL_POINTERS

// use same restrictions as a JIT function:
// cannot have stack map, frame, local variables;
// must use one of call* trampolines to invoke arbitrary Go functions
TEXT ·asm_hideme(SB),NOSPLIT|NOFRAME,$0-8 // must not have local variables
	MOVQ env+0(FP), AX
	MOVQ 0(AX),  DX            // closure, must be in DX
	MOVQ 8(AX),  BX            // closure arg
	MOVQ 24(AX), CX            // helper function: use call[1] == call16
	MOVQ BX, local_arg-40(SP)  // write into callee stack
	CALL CX
	RET
