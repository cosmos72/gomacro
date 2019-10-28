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

TEXT ·asm_call_closure(SB),NOSPLIT,$8-16
	NO_LOCAL_POINTERS
	MOVQ closure+0(FP), BX
	MOVQ received_arg+8(FP), DX
	MOVQ DX, local_arg-8(SP)
	CALL 0(BX)
	RET

TEXT ·growStack(SB),0,$1152-0
	NO_LOCAL_POINTERS
	RET

/*
TEXT ·hideme(SB),NOSPLIT,$0-8 // must not have local variables
	MOVQ env+0(FP), AX
	MOVQ 0(AX), BX            // closure, must be in BX
	MOVQ 8(AX), DX            // closure arg
	MOVQ DX, local_arg-32(SP) // write into callee stack
	CALL ·call8(SB)
	RET
*/

TEXT ·hideme(SB),NOSPLIT,$0-8 // must not have local variables
	MOVQ env+0(FP), AX
	MOVQ 0(AX), BX            // closure, must be in BX
	MOVQ 8(AX), DX            // closure arg
	MOVQ DX, local_arg-536(SP) // write into callee stack
	CALL ·call512(SB)
	RET

/**
 * call the closure stored in BX (which expects exactly 0 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_func()
 */
TEXT ·call0(SB),NOSPLIT,$8-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), DX
	MOVQ caller_ip+8(SP), AX
	MOVQ DX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	/*
	 * Go call abi: closures are pointers to
	 * struct {
	 *   func_address uintptr
	 *   closure_data [...]uint8
	 * }
	 * and struct address must be passed in BX
	 */
	CALL 0(BX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

/**
 * call the closure stored in BX (which expects up to 8 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_func()
 *
 * writing arguments in our stack and retrieving return values from it
 * are caller's responsibility
 * (possible only in assembly: caller has to access our stack)
 */
TEXT ·call8(SB),NOSPLIT,$16-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), DX
	MOVQ caller_ip+8(SP), AX
	MOVQ DX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(BX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 16 bytes of arguments return values
TEXT ·call16(SB),NOSPLIT,$24-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), DX
	MOVQ caller_ip+8(SP), AX
	MOVQ DX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(BX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 32 bytes of arguments return values
TEXT ·call32(SB),NOSPLIT,$40-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), DX
	MOVQ caller_ip+8(SP), AX
	MOVQ DX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(BX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 64 bytes of arguments return values
TEXT ·call64(SB),NOSPLIT,$72-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), DX
	MOVQ caller_ip+8(SP), AX
	MOVQ DX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(BX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 128 bytes of arguments return values
TEXT ·call128(SB),NOSPLIT,$136-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), DX
	MOVQ caller_ip+8(SP), AX
	MOVQ DX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(BX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 256 bytes of arguments return values
TEXT ·call256(SB),NOSPLIT,$264-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), DX
	MOVQ caller_ip+8(SP), AX
	MOVQ DX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(BX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 512 bytes of arguments return values
TEXT ·call512(SB),NOSPLIT,$520-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), DX
	MOVQ caller_ip+8(SP), AX
	MOVQ DX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(BX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 1024 bytes of arguments return values
TEXT ·call1024(SB),NOSPLIT,$1032-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), DX
	MOVQ caller_ip+8(SP), AX
	MOVQ DX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(BX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET
