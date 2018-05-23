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

void And2_l(u64 *ints) {
    z(64) &= 0x55667788;
}
void And2_q(u64 *ints) {
    z(64) &= 0x5566778899aabbccll;
}
void And_l(u64 *ints) {
    z(64) = a(64) & ~0x55667788ll;
}
void And_q(u64 *ints) {
    z(64) = a(64) & ~0x5566778899aabbccll;
}
void And(u64 *ints) {
    z(64) = a(64) & b(64);
}


void Or2_l(u64 *ints) {
    z(64) |= 0x55667788;
}
void Or2_q(u64 *ints) {
    z(64) |= 0x5566778899aabbccll;
}
void Or_l(u64 *ints) {
    z(64) = a(64) | 0x55667788;
}
void Or_q(u64 *ints) {
    z(64) = a(64) | 0x5566778899aabbcc;
}
void Or(u64 *ints) {
    z(64) = a(64) | b(64);
}


void Xor2_l(u64 *ints) {
    z(64) ^= 0x55667788;
}
void Xor2_q(u64 *ints) {
    z(64) ^= 0x5566778899aabbccll;
}
void Xor_l(u64 *ints) {
    z(64) = a(64) ^ 0x55667788;
}
void Xor_q(u64 *ints) {
    z(64) = a(64) ^ 0x5566778899aabbcc;
}
void Xor(u64 *ints) {
    z(64) = a(64) ^ b(64);
}



void Andnot2_l(u64 *ints) {
    z(64) &= ~0x55667788;
}
void Andnot2_q(u64 *ints) {
    z(64) &= ~0x5566778899aabbccll;
}
void Andnot_l(u64 *ints) {
    z(64) = a(64) & ~0x55667788;
}
void Andnot_1l(u64 *ints) {
    z(64) = 0x55667788 & ~a(64);
}
void Andnot_1q(u64 *ints) {
    z(64) = 0x5566778899aabbcc & ~a(64);
}
void Andnot_q(u64 *ints) {
    z(64) = a(64) & ~0x5566778899aabbcc;
}
void Andnot(u64 *ints) {
    z(64) = a(64) & ~b(64);
}
