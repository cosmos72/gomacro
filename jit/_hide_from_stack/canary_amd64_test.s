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
 * func_test_amd64.go
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

#include "textflag.h" // for NOSPLIT
#include "funcdata.h" // for NO_LOCAL_POINTERS

TEXT ·asm_loop(SB),NOSPLIT|NOFRAME,$0-0
label_asm_loop:
    JMP  label_asm_loop
	RET

DATA  ·const_canary+0(SB)/8,$·canary(SB)
GLOBL ·const_canary+0(SB),RODATA,$8

TEXT ·asm_address_of_canary(SB),NOSPLIT|NOFRAME,$0-8
    LEAQ ·const_canary(SB), AX // closure. actual function is LEAQ ·canary(SB), AX
	MOVQ AX, ret+0(FP)
	RET

TEXT ·asm_call_canary(SB),NOSPLIT,$8-8
	NO_LOCAL_POINTERS
	MOVQ received_arg+0(FP), AX
	MOVQ AX, local_arg-8(SP)
    CALL ·canary(SB)
	RET

TEXT ·asm_call_func(SB),NOSPLIT,$8-16
	NO_LOCAL_POINTERS
	MOVQ func_address+0(FP), AX
	MOVQ received_arg+8(FP), DX
	MOVQ DX, local_arg-8(SP)
	CALL AX
	RET

// on amd64, closure itself must be passed in DX
TEXT ·asm_call_closure(SB),NOSPLIT,$8-16
	NO_LOCAL_POINTERS
	MOVQ closure+0(FP), DX
	MOVQ received_arg+8(FP), AX
	MOVQ AX, local_arg-8(SP)
	CALL 0(DX)
	RET

TEXT ·asm_hideme(SB),NOSPLIT|NOFRAME,$0-8 // must not have local variables
	MOVQ env+0(FP), AX
	MOVQ 0(AX),  DX            // closure, must be in DX
	MOVQ 8(AX),  BX            // closure arg
	MOVQ 16(AX), CX            // helper function
	MOVQ BX, local_arg-32(SP)  // write into callee stack
	// CALL ·call8(SB)
	CALL CX
	RET

/*
TEXT ·asm_hideme(SB),NOSPLIT|NOFRAME,$0-8 // must not have local variables
	MOVQ env+0(FP), AX
	MOVQ 0(AX), DX            // closure, must be in DX
	MOVQ 8(AX), BX            // closure arg
	MOVQ BX, local_arg-536(SP) // write into callee stack
	CALL ·call512(SB)
	RET
*/

