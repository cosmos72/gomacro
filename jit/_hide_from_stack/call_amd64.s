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

#include "textflag.h" // for NOSPLIT
#include "funcdata.h" // for NO_LOCAL_POINTERS

TEXT ·growStack(SB),0,$1152-0
	NO_LOCAL_POINTERS
	RET

/**
 * call the closure stored in DX (which expects exactly 0 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_func()
 */
TEXT ·call0(SB),NOSPLIT,$8-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), BX
	MOVQ caller_ip+8(SP), AX
	MOVQ BX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	/*
	 * Go call abi: closures are pointers to
	 * struct {
	 *   func_address uintptr
	 *   closure_data [...]uint8
	 * }
	 * and struct address must be passed in DX
	 */
	CALL 0(DX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

/**
 * call the closure stored in DX (which expects up to 8 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_func()
 *
 * writing arguments in our stack and retrieving return values from it
 * are caller's responsibility
 * (possible only in assembly: caller has to access our stack)
 */
TEXT ·call8(SB),NOSPLIT,$16-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), BX
	MOVQ caller_ip+8(SP), AX
	MOVQ BX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(DX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 16 bytes of arguments return values
TEXT ·call16(SB),NOSPLIT,$24-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), BX
	MOVQ caller_ip+8(SP), AX
	MOVQ BX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(DX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 32 bytes of arguments return values
TEXT ·call32(SB),NOSPLIT,$40-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), BX
	MOVQ caller_ip+8(SP), AX
	MOVQ BX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(DX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 64 bytes of arguments return values
TEXT ·call64(SB),NOSPLIT,$72-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), BX
	MOVQ caller_ip+8(SP), AX
	MOVQ BX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(DX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 128 bytes of arguments return values
TEXT ·call128(SB),NOSPLIT,$136-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), BX
	MOVQ caller_ip+8(SP), AX
	MOVQ BX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(DX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 256 bytes of arguments return values
TEXT ·call256(SB),NOSPLIT,$264-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), BX
	MOVQ caller_ip+8(SP), AX
	MOVQ BX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(DX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 512 bytes of arguments return values
TEXT ·call512(SB),NOSPLIT,$520-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), BX
	MOVQ caller_ip+8(SP), AX
	MOVQ BX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(DX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET

// as above, but closure can expect up to 1024 bytes of arguments return values
TEXT ·call1024(SB),NOSPLIT,$1032-0
	NO_LOCAL_POINTERS
	LEAQ ·hidden_func(SB), BX
	MOVQ caller_ip+8(SP), AX
	MOVQ BX, caller_ip+8(SP)
	MOVQ AX, save_caller_ip-8(SP)
	CALL 0(DX)
	MOVQ save_caller_ip-8(SP), AX
	MOVQ AX, caller_ip+8(SP)
	RET
