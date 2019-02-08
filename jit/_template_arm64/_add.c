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
 * add.c
 *
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

#include <stdint.h>

uint8_t add8(uint8_t a, uint8_t b) {
	return a + b;
}
uint16_t add16(uint16_t a, uint16_t b) {
	return a + b;
}
uint32_t add32(uint32_t a, uint32_t b) {
	return a + b;
}
uint64_t add64(uint64_t a, uint64_t b) {
	return a + b;
}


int8_t sadd8(int8_t a, int8_t b) {
	return a + b;
}
int16_t sadd16(int16_t a, int16_t b) {
	return a + b;
}
int32_t sadd32(int32_t a, int32_t b) {
	return a + b;
}
int64_t sadd64(int64_t a, int64_t b) {
	return a + b;
}

