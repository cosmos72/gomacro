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
 * mul.c
 *
 *  Created on Feb 02, 2019
 *      Author Massimiliano Ghilardi
 */

#include <stdint.h>

uint32_t mul1_32(uint32_t a) {
	return a * 1;
}
uint32_t mul9_32(uint32_t a) {
	return a * 9;
}
uint32_t mul255_32(uint32_t a) {
	return a * 255;
}

uint64_t mul1_64(uint64_t a) {
	return a * 1;
}
uint64_t mul9_64(uint64_t a) {
	return a * 9;
}
uint64_t mul255_64(uint64_t a) {
	return a * 255;
}

uint8_t mul8(uint8_t a, uint8_t b) {
	return a * b;
}
uint16_t mul16(uint16_t a, uint16_t b) {
	return a * b;
}
uint32_t mul32(uint32_t a, uint32_t b) {
	return a * b;
}
uint64_t mul64(uint64_t a, uint64_t b) {
	return a * b;
}

int8_t smul8(int8_t a, int8_t b) {
	return a * b;
}
int16_t smul16(int16_t a, int16_t b) {
	return a * b;
}
int32_t smul32(int32_t a, int32_t b) {
	return a * b;
}
int64_t smul64(int64_t a, int64_t b) {
	return a * b;
}


