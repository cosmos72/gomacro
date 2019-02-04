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

uint32_t add1_32(uint32_t a) {
	return a + 1;
}
uint32_t add9_32(uint32_t a) {
	return a + 9;
}
uint32_t add255_32(uint32_t a) {
	return a + 255;
}

uint64_t add1_64(uint64_t a) {
	return a + 1;
}
uint64_t add9_64(uint64_t a) {
	return a + 9;
}
uint64_t add255_64(uint64_t a) {
	return a + 255;
}

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

void padd8(uint8_t *a, uint8_t *b, uint8_t *dst) {
	*dst = *a + *b;
}
void padd16(uint16_t *a, uint16_t *b, uint16_t *dst) {
	*dst = *a + *b;
}
void padd32(uint32_t *a, uint32_t *b, uint32_t *dst) {
	*dst = *a + *b;
}
void padd64(uint64_t *a, uint64_t *b, uint64_t *dst) {
	*dst = *a + *b;
}

