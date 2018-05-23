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
 * binary_arith.c
 *
 *  Created on May 20, 2018
 *      Author Massimiliano Ghilardi
 */


#include "asm_template.h"

void Add(u64 *ints) {
    z(64) += a(64);
}
void Sub(u64 *ints) {
    z(64) -= a(64);
}
void Mul2_l(u64 *ints) {
    z(64) *= 0x55667788;
}
void Mul2_q(u64 *ints) {
    z(64) *= 0x5566778899aabbcc;
}
void Mul(u64 *ints) {
    z(64) *= a(64);
}


void Quo(u64 *ints) {
    z(64) /= a(64);
}
void QuoU(u64 *ints) {
    uz(64) /= ua(64);
}


void Rem(u64 *ints) {
    z(64) %= a(64);
}
void RemU(u64 *ints) {
    uz(64) %= ua(64);
}
