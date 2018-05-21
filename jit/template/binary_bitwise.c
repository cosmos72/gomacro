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
 * binary_bitwise.c
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */


#include "asm_template.h"


void AndInt8(u64 *ints) {
    z(8) = a(8) & b(8);
}
void AndInt16(u64 *ints) {
    z(16) = a(16) & b(16);
}
void AndInt32(u64 *ints) {
    z(32) = a(32) & b(32);
}
void AndInt64(u64 *ints) {
    z(64) = a(64) & b(64);
}

void OrInt8(u64 *ints) {
    z(8) = a(8) | b(8);
}
void OrInt16(u64 *ints) {
    z(16) = a(16) | b(16);
}
void OrInt32(u64 *ints) {
    z(32) = a(32) | b(32);
}
void OrInt64(u64 *ints) {
    z(64) = a(64) | b(64);
}

void XorInt8(u64 *ints) {
    z(8) = a(8) ^ b(8);
}
void XorInt16(u64 *ints) {
    z(16) = a(16) ^ b(16);
}
void XorInt32(u64 *ints) {
    z(32) = a(32) ^ b(32);
}
void XorInt64(u64 *ints) {
    z(64) = a(64) ^ b(64);
}

void AndnotInt8(u64 *ints) {
    z(8) = a(8) & ~b(8);
}
void AndnotInt16(u64 *ints) {
    z(16) = a(16) & ~b(16);
}
void AndnotInt32(u64 *ints) {
    z(32) = a(32) & ~b(32);
}
void AndnotInt64(u64 *ints) {
    z(64) = a(64) & ~b(64);
}
