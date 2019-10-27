// +build gc

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
 * call_amd64.go
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

// #include "go_asm.h"
#include "textflag.h" // for NOSPLIT
#include "funcdata.h" // for NO_LOCAL_POINTERS

TEXT ·asm_loop(SB),NOSPLIT,$8-0
label_asm_loop:
    JMP  label_asm_loop
	RET


TEXT ·asm_address_of_canary(SB),NOSPLIT,$0-8
    MOVQ ·var_canary(SB), AX // closure. actual function is LEAQ ·canary(SB), AX
	MOVQ AX, ret+0(FP)
	RET

TEXT ·asm_call_canary(SB),NOSPLIT,$0-0
    CALL ·canary(SB)
	RET

TEXT ·asm_call_func(SB),NOSPLIT,$0-8
	MOVQ tocall+0(FP), AX
	CALL AX
	RET

TEXT ·asm_call_closure(SB),NOSPLIT,$0-8
	MOVQ tocall+0(FP), BX
	CALL 0(BX)
	RET

TEXT ·hideme(SB),NOSPLIT,$0-8
	MOVQ env+0(FP), BX
	MOVQ 0(BX), BX
	CALL ·call_0b(SB)
	RET

/**
 * call the closure stored in BX,
 * hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_func()
 */
TEXT ·call_0b(SB),NOSPLIT,$8-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), DX
	MOVQ caller_ip+8(SP), AX
	MOVQ DX, caller_ip+8(SP)
	MOVQ AX, localvar-8(SP)
	/*
	 * Go call abi: closures are pointers to
	 * struct {
	 *   func_address uintptr
	 *   closure_data [...]uint8
	 * }
	 * and struct address must be passed in BX
	 */
	CALL 0(BX)
	MOVQ localvar-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

/*
TEXT ·call_8b(SB),NOSPLIT,$0-16
	POPQ  AX
	MOVQ  target_func+8(FP), BX
	MOVQ  AX, target_func+8(FP)
	CALL  0(BX)
	MOVQ  target_func+8(FP), AX
	PUSHQ AX
	RET

TEXT ·call_16b(SB),NOSPLIT,$0-24
	POPQ  AX
	MOVQ  target_func+16(FP), BX
	MOVQ  AX, target_func+16(FP)
	CALL  0(BX)
	MOVQ  target_func+16(FP), DX
	PUSHQ AX
	RET

TEXT ·call_32b(SB),NOSPLIT,$0-40
	POPQ  AX
	MOVQ  target_func+32(FP), BX
	MOVQ  AX, target_func+32(FP)
	CALL  0(BX)
	MOVQ  target_func+32(FP), DX
	PUSHQ AX
	RET

TEXT ·call_64b(SB),NOSPLIT,$0-72
	POPQ  AX
	MOVQ  target_func+64(FP), BX
	MOVQ  AX, target_func+64(FP)
	CALL  0(BX)
	MOVQ  target_func+64(FP), DX
	PUSHQ AX
	RET

TEXT ·call_128b(SB),NOSPLIT,$0-136
	POPQ  AX
	MOVQ  target_func+128(FP), BX
	MOVQ  AX, target_func+128(FP)
	CALL  0(BX)
	MOVQ  target_func+128(FP), DX
	PUSHQ AX
	RET

TEXT ·call_256b(SB),NOSPLIT,$0-264
	POPQ  AX
	MOVQ  target_func+256(FP), BX
	MOVQ  AX, target_func+256(FP)
	CALL  0(BX)
	MOVQ  target_func+256(FP), DX
	PUSHQ AX
	RET
*/
