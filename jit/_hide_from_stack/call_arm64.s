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
 * call_arm64.s
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

#include "textflag.h" // for NOSPLIT
#include "funcdata.h" // for NO_LOCAL_POINTERS

TEXT ·GrowStack(SB),0,$640-0
	NO_LOCAL_POINTERS
	RET

// unused on arm64
TEXT ·hidden_jit_func(SB),NOSPLIT|NOFRAME,$0-8
	NO_LOCAL_POINTERS
	RET

/**
 * call the closure stored in R26 (which expects exactly 0 bytes of arguments
 * + return values), hiding the caller from runtime stack:
 * caller completely disappears from stack trace
 *
 * writing arguments in our stack and retrieving return values from it
 * are caller's responsibility
 * (possible only in assembly: caller has to access our stack)
 *
 * call0 stack contains, from top to bottom:
 *   hidden_return_addr-0x8(SP):  R30, stored by call0 itself
 *   caller_return_addr-0x10(SP): R30, stored&read by caller
 * and is then followed by standard data:
 *   return_addr-0x18(SP):        R30, stored by call0 itself
 *   frame_reg-0x20(SP):          R29, stored by call0 prologue
 */
TEXT ·call0(SB),NOSPLIT,$0x10-0
	// on arm64, we cannot emulate a FRAME function using NOFRAME,
	// even if the function is the same byte-by-byte:
	// runtime.GC() crashes with "runtime: unexpected return pc ..."
	//
	// thus we must allow Go assembler inject standard prologue,
	// then alter the copies of X30 stored on stack,
	// finally implement our own epilogue.
	NO_LOCAL_POINTERS

	// save our return_addr R30 to the "hidden" location
	MOVD R30, hidden_return_addr-0x8(SP)
	// load caller's return_addr, prepared by caller
	MOVD caller_return_addr-0x10(SP), R0
	// and store it as our return address, overwriting the value set by prologue
	MOVD R0, return_addr-0x18(SP)

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
	WORD $0xF85F83FD        // ldur    x29, [sp, #-8]
	// retrieve our return_addr R30 from "hidden" location
	MOVD hidden_return_addr-0x8(SP), R30
	// add 32 to stack
	WORD $0x910083FF        // add     sp, sp, #0x20
	// return
	WORD $0xD65F03C0        // ret


/*
 * as above, but closure can expect up to 16 bytes of arguments return values
 * and stack does not have padding:
 *   hidden_return_addr-0x8(SP):  R30, stored by call16 itself
 *   caller_return_addr-0x10(SP): R30, stored&read by caller
 *   closure_arg_or_ret-0x18(SP):      stored&read by caller
 *   closure_arg_or_ret-0x20(SP):      stored&read by caller
 * and is then followed by standard data:
 *   return_addr-0x28(SP):        R30, stored by call16 itself
 *   frame_reg-0x30(SP):          R29, stored by call16 prologue
 */
TEXT ·call16(SB),NOSPLIT,$0x20-0
	// same technique as call0
	NO_LOCAL_POINTERS
	MOVD R30, hidden_return_addr-0x8(SP)
	MOVD caller_return_addr-0x10(SP), R0
	MOVD R0, return_addr-0x28(SP)
	MOVD 0(R26), R1
	CALL R1
	WORD $0xF85F83FD        // ldur    x29, [sp, #-8]
	MOVD hidden_return_addr-0x8(SP), R30
	WORD $0x9100C3FF        // add     sp, sp, #0x30
	WORD $0xD65F03C0        // ret

// as above, but closure can expect up to 32 bytes of arguments return values
TEXT ·call32(SB),NOSPLIT,$0x30-0
	// same technique as call0
	NO_LOCAL_POINTERS
	MOVD R30, hidden_return_addr-0x8(SP)
	MOVD caller_return_addr-0x10(SP), R0
	MOVD R0, return_addr-0x38(SP)
	MOVD 0(R26), R1
	CALL R1
	WORD $0xF85F83FD        // ldur    x29, [sp, #-8]
	MOVD hidden_return_addr-0x8(SP), R30
	WORD $0x910103ff        // add     sp, sp, #0x40
	WORD $0xD65F03C0        // ret

// as above, but closure can expect up to 64 bytes of arguments return values
TEXT ·call64(SB),NOSPLIT,$0x50-0
	// same technique as call0
	NO_LOCAL_POINTERS
	MOVD R30, hidden_return_addr-0x8(SP)
	MOVD caller_return_addr-0x10(SP), R0
	MOVD R0, return_addr-0x58(SP)
	MOVD 0(R26), R1
	CALL R1
	WORD $0xF85F83FD        // ldur    x29, [sp, #-8]
	MOVD hidden_return_addr-0x8(SP), R30
	WORD $0x910183ff        // add     sp, sp, #0x60
	WORD $0xD65F03C0        // ret

// as above, but closure can expect up to 128 bytes of arguments return values
TEXT ·call128(SB),NOSPLIT,$0x90-0
	// same technique as call0
	NO_LOCAL_POINTERS
	MOVD R30, hidden_return_addr-0x8(SP)
	MOVD caller_return_addr-0x10(SP), R0
	MOVD R0, return_addr-0x98(SP)
	MOVD 0(R26), R1
	CALL R1
	WORD $0xF85F83FD        // ldur    x29, [sp, #-8]
	MOVD hidden_return_addr-0x8(SP), R30
	WORD $0x910283ff        // add     sp, sp, #0xA0
	WORD $0xD65F03C0        // ret

// as above, but closure can expect up to 256 bytes of arguments return values
TEXT ·call256(SB),NOSPLIT,$0x110-0
	// same technique as call0
	NO_LOCAL_POINTERS
	MOVD R30, hidden_return_addr-0x8(SP)
	MOVD caller_return_addr-0x10(SP), R0
	MOVD R0, return_addr-0x118(SP)
	MOVD 0(R26), R1
	CALL R1
	WORD $0xF85F83FD        // ldur    x29, [sp, #-8]
	MOVD hidden_return_addr-0x8(SP), R30
	WORD $0x910483ff        // add     sp, sp, #0x120
	WORD $0xD65F03C0        // ret

// as above, but closure can expect up to 512 bytes of arguments return values
TEXT ·call512(SB),NOSPLIT,$0x210-0
	// same technique as call0
	NO_LOCAL_POINTERS
	MOVD R30, hidden_return_addr-0x8(SP)
	MOVD caller_return_addr-0x10(SP), R0
	MOVD R0, return_addr-0x218(SP)
	MOVD 0(R26), R1
	CALL R1
	WORD $0xF85F83FD        // ldur    x29, [sp, #-8]
	MOVD hidden_return_addr-0x8(SP), R30
	WORD $0x910883ff        // add     sp, sp, #0x220
	WORD $0xD65F03C0        // ret
