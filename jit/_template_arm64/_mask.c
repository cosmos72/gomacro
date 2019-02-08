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
 * _mask.c
 *
 *  Created on Feb 08, 2019
 *      Author Massimiliano Ghilardi
 */

#include <stdint.h>

uint8_t mask8(uint8_t a, uint8_t b) {
	return a + b;
}

uint16_t mask16(uint16_t a, uint16_t b) {
	return a + b;
}
