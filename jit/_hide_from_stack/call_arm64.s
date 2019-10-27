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
 * call_arm64.go
 *
 *  Created on Oct 27, 2019
 *      Author Massimiliano Ghilardi
 */

// +build gc

#include "go_asm.h"
#include "textflag.h" // for NOSPLIT

TEXT ·call_0b(SB),NOSPLIT,$0-8
	POP   R1
	MOVD  target_ip+0(FP), R0
	MOVD  R1, target_ip+0(FP)
	CALL  R0
	MOVD  target_ip+0(FP), R1
	PUSH  R1
	RET

