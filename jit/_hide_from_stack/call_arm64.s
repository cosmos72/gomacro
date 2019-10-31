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
 * call the closure stored in R26 (which expects exactly 0 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_func()
 */
TEXT ·call0(SB),NOSPLIT,$8-0
	NO_LOCAL_POINTERS
	MOVD 0(R26), R0
	MOVD LR, save_caller_ip-8(SP)
	MOVD ·some_hidden_func(SB), LR
	/*
	 * Go call abi: closures are pointers to
	 * struct {
	 *   func_address uintptr
	 *   closure_data [...]uint8
	 * }
	 * and struct address must be passed in R26
	 */
	CALL R0
	MOVD save_caller_ip-8(SP), LR
	RET

/**
 * call the closure stored in R26 (which expects up to 8 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_func()
 *
 * writing arguments in our stack and retrieving return values from it
 * are caller's responsibility
 * (possible only in assembly: caller has to access our stack)
 */
TEXT ·call8(SB),NOSPLIT,$16-0
	NO_LOCAL_POINTERS
	MOVD 0(R26), R0
	CALL R0
	MOVD save_caller_ip-8(SP), R30
	RET

// as above, but closure can expect up to 16 bytes of arguments return values
TEXT ·call16(SB),NOSPLIT,$24-0
	NO_LOCAL_POINTERS
	MOVD 0(R26), R0
	CALL R0
	MOVD save_caller_ip-8(SP), LR
	RET

// as above, but closure can expect up to 32 bytes of arguments return values
TEXT ·call32(SB),NOSPLIT,$40-0
	NO_LOCAL_POINTERS
	MOVD 0(R26), R0
	CALL R0
	MOVD save_caller_ip-8(SP), LR
	RET

// as above, but closure can expect up to 64 bytes of arguments return values
TEXT ·call64(SB),NOSPLIT,$72-0
	NO_LOCAL_POINTERS
	MOVD 0(R26), R0
	CALL R0
	MOVD save_caller_ip-8(SP), LR
	RET

// as above, but closure can expect up to 128 bytes of arguments return values
TEXT ·call128(SB),NOSPLIT,$136-0
	NO_LOCAL_POINTERS
	MOVD 0(R26), R0
	CALL R0
	MOVD save_caller_ip-8(SP), LR
	RET

// as above, but closure can expect up to 256 bytes of arguments return values
TEXT ·call256(SB),NOSPLIT,$264-0
	NO_LOCAL_POINTERS
	MOVD 0(R26), R0
	CALL R0
	MOVD save_caller_ip-8(SP), LR
	RET

// as above, but closure can expect up to 512 bytes of arguments return values
TEXT ·call512(SB),NOSPLIT,$520-0
	NO_LOCAL_POINTERS
	MOVD 0(R26), R0
	CALL R0
	MOVD save_caller_ip-8(SP), LR
	RET

// as above, but closure can expect up to 1024 bytes of arguments return values
TEXT ·call1024(SB),NOSPLIT,$1032-0
	NO_LOCAL_POINTERS
	MOVD 0(R26), R0
	CALL R0
	MOVD save_caller_ip-8(SP), LR
	RET
