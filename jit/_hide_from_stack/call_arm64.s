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

TEXT ·grow_stack(SB),0,$1152-0
	NO_LOCAL_POINTERS
	RET

// hidden JIT functions will be replaced by this function in the stacktrace
TEXT ·hidden_jit_func(SB),NOSPLIT|NOFRAME,$0-8 // same signature as asm_hideme
	NO_LOCAL_POINTERS
	RET

/**
 * call the closure stored in R26 (which expects exactly 0 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_func()
 */
TEXT ·call0(SB),NOSPLIT,$8-0
	NO_LOCAL_POINTERS
	RET

/**
 * call the closure stored in R26 (which expects up to 8 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * caller is replaced with a fake entry hidden_jit_func()
 *
 * writing arguments in our stack and retrieving return values from it
 * are caller's responsibility
 * (possible only in assembly: caller has to access our stack)
 */
TEXT ·call8(SB),NOSPLIT,$16-0
    // on arm64, we cannot emulate a FRAME function using NOFRAME,
	// even if the function is the same byte-by-byte:
    // runtime.GC() crashes with "runtime: unexpected return pc ..."
	//
	// thus we must allow Go assembler inject standard prologue and epilogue,
	// then alter the copies of X30 stored on stack.
	NO_LOCAL_POINTERS

    // load caller's caller's return_addr, prepared by caller
	MOVD hidden_return_addr-8(SP), R0
    // and store it as caller's return address, hiding the latter
	MOVD R0, return_addr-24(SP)
	// finally save caller return_addr R30 to a "hidden" location
	MOVD R30, hidden_return_addr-8(SP)

	/*
	 * Go call abi: closures are pointers to
	 * struct {
	 *   func_address uintptr
	 *   closure_data [...]uint8
	 * }
	 * and struct address must be passed in R26
	 */
	MOVD 0(R26), R1
	CALL R1

    // override standard epilogue:
	// retrieve frame_reg R29 from standard location
	MOVD frame_reg-32(SP), R29 // $0xf85f83fd        ldur    x29, [sp, #-8]
	// retrieve caller return_addr R30 from "hidden" location
	MOVD hidden_return_addr-8(SP), R30
	// copy back caller's caller's return_addr from standard to "hidden" location,
	// were our caller will pick it
	MOVD return_addr-24(SP), R0
	MOVD R0, hidden_return_addr-8(SP)

	// add 32 to stack
    WORD $0x910083FF        // add     sp, sp, #0x20
    WORD $0xD65F03C0        // ret


// as above, but closure can expect up to 16 bytes of arguments return values
TEXT ·call16(SB),NOSPLIT,$24-0
	NO_LOCAL_POINTERS
	RET

// as above, but closure can expect up to 32 bytes of arguments return values
TEXT ·call32(SB),NOSPLIT,$40-0
	NO_LOCAL_POINTERS
	RET

// as above, but closure can expect up to 64 bytes of arguments return values
TEXT ·call64(SB),NOSPLIT,$72-0
	NO_LOCAL_POINTERS
	RET

// as above, but closure can expect up to 128 bytes of arguments return values
TEXT ·call128(SB),NOSPLIT,$136-0
	NO_LOCAL_POINTERS
	RET

// as above, but closure can expect up to 256 bytes of arguments return values
TEXT ·call256(SB),NOSPLIT,$264-0
	NO_LOCAL_POINTERS
	RET

// as above, but closure can expect up to 512 bytes of arguments return values
TEXT ·call512(SB),NOSPLIT,$520-0
	NO_LOCAL_POINTERS
	RET

// as above, but closure can expect up to 1024 bytes of arguments return values
TEXT ·call1024(SB),NOSPLIT,$1032-0
	NO_LOCAL_POINTERS
	RET
