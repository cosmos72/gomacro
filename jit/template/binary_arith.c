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

void IncInt8(u64 *ints) {
    z(8) += a(8);
}
void IncInt16(u64 *ints) {
    z(16) += a(16);
}
void IncInt32(u64 *ints) {
    z(32) += a(32);
}
void IncInt64(u64 *ints) {
    z(64) += a(64);
}

void AddInt8(u64 *ints) {
    z(8) = a(8) + b(8);
}
void AddInt16(u64 *ints) {
    z(16) = a(16) + b(16);
}
void AddInt32(u64 *ints) {
    z(32) = a(32) + b(32);
}
void AddInt64(u64 *ints) {
    z(64) = a(64) + b(64);
}

void DecInt8(u64 *ints) {
    z(8) -= a(8);
}
void DecInt16(u64 *ints) {
    z(16) -= a(16);
}
void DecInt32(u64 *ints) {
    z(32) -= a(32);
}
void DecInt64(u64 *ints) {
    z(64) -= a(64);
}

void SubInt8(u64 *ints) {
    z(8) = a(8) - b(8);
}
void SubInt16(u64 *ints) {
    z(16) = a(16) - b(16);
}
void SubInt32(u64 *ints) {
    z(32) = a(32) - b(32);
}
void SubInt64(u64 *ints) {
    z(64) = a(64) - b(64);
}

void MulInt8(u64 *ints) {
    z(8) = a(8) * b(8);
}
void MulInt16(u64 *ints) {
    z(16) = a(16) * b(16);
}
void MulInt32(u64 *ints) {
    z(32) = a(32) * b(32);
}
void MulInt64(u64 *ints) {
    z(64) = a(64) * b(64);
}

void QuoInt8(u64 *ints) {
    z(8) = a(8) / b(8);
}
void QuoInt16(u64 *ints) {
    z(16) = a(16) / b(16);
}
void QuoInt32(u64 *ints) {
    z(32) = a(32) / b(32);
}
void QuoInt64(u64 *ints) {
    z(64) = a(64) / b(64);
}

void QuoUint8(u64 *ints) {
    uz(8) = ua(8) / ub(8);
}
void QuoUint16(u64 *ints) {
    uz(16) = ua(16) / ub(16);
}
void QuoUint32(u64 *ints) {
    uz(32) = ua(32) / ub(32);
}
void QuoUint64(u64 *ints) {
    uz(64) = ua(64) / ub(64);
}

void RemInt8(u64 *ints) {
    z(8) = a(8) % b(8);
}
void RemInt16(u64 *ints) {
    z(16) = a(16) % b(16);
}
void RemInt32(u64 *ints) {
    z(32) = a(32) % b(32);
}
void RemInt64(u64 *ints) {
    z(64) = a(64) % b(64);
}

void RemUint8(u64 *ints) {
    uz(8) = ua(8) % ub(8);
}
void RemUint16(u64 *ints) {
    uz(16) = ua(16) % ub(16);
}
void RemUint32(u64 *ints) {
    uz(32) = ua(32) % ub(32);
}
void RemUint64(u64 *ints) {
    uz(64) = ua(64) % ub(64);
}
