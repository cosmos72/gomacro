/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * _zero.c
 *
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

#include <stdint.h>

uint8_t zero8() {
	return 0;
}
uint16_t zero16() {
	return 0;
}
uint32_t zero32() {
	return 0;
}
uint64_t zero64() {
	return 0;
}

